package service

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/helper"
	"common-web-framework/models"
	"common-web-framework/repository"
	"common-web-framework/response"
	"common-web-framework/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"
)

type FileService interface {
	// AddFile 添加一个文件
	AddFile(fileInfo models.FileInfo)
	// GlobalUploadFWile 全局上传文件
	GlobalUploadFile(ctx *gin.Context, isImage bool, uid *int) []response.SimpleFileResponse
	// UploadImageFile 上传图片文件
	UploadImageFile(ctx *gin.Context) []response.SimpleFileResponse
	// UploadAvatarFile 上传头像文件
	UploadAvatarFile(ctx *gin.Context) []response.SimpleFileResponse
	// UploadFile 上传其他文件
	UploadFile(ctx *gin.Context) []response.SimpleFileResponse
}

type FileServiceImpl struct {
	repository repository.FileRepository
}

func GetFiles(ctx *gin.Context) []*multipart.FileHeader {
	var files, err = ctx.MultipartForm()

	if err != nil {
		ctx.JSON(http.StatusOK, common.AutoFail(common.NoFile))
		return nil
	}

	return files.File["files"]
}

var mb = 1024 * 1024

func (fs FileServiceImpl) GlobalUploadFile(ctx *gin.Context, isImage bool, uid *int) (frs []response.SimpleFileResponse) {

	var u = config.CONFIG.Upload

	var files = GetFiles(ctx)

	var infos []models.FileInfo

	for _, file := range files {

		var size = file.Size

		var fileName = file.Filename

		var suffix = filepath.Ext(fileName)

		var create = time.Now()

		if isImage {
			if !utils.IsImageFile(suffix) {
				frs = append(frs, response.SimpleFileResponse{
					Status:  "fail",
					Message: "这不是一个正确的图片格式",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				})
				continue
			} else if size > int64(u.MaxImageSize*mb) {
				frs = append(frs, response.SimpleFileResponse{
					Status:  "fail",
					Message: "图片文件大小超出",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				})
				continue
			}
		} else {
			if size > int64(u.MaxFileSize*mb) {
				frs = append(frs, response.SimpleFileResponse{
					Status:  "fail",
					Message: "文件大小超出",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				})
				continue
			}
		}

		var f, err = file.Open()

		if err != nil {
			frs = append(frs, response.SimpleFileResponse{
				Status:  "fail",
				Message: "文件打开失败",
				Name:    fileName,
				Create:  utils.FormatDate(create),
			})
			continue
		}

		var md5 = utils.GetFileMd5(f)

		var newName = md5 + suffix

		var saveFilePath = u.Path + "/" + newName

		var url string

		if dbUrl := fs.repository.FindByMd5(md5); dbUrl == "" {
			url = u.Uri + "/" + newName
			var uploadError = ctx.SaveUploadedFile(file, saveFilePath)
			if uploadError != nil {
				frs = append(frs, response.SimpleFileResponse{
					Status:  "fail",
					Message: "文件上传失败",
					Name:    fileName,
					Create:  utils.FormatDate(create),
				})
				continue
			}
		} else {
			url = dbUrl
		}

		infos = append(infos, models.FileInfo{
			OldName:      fileName,
			NewName:      newName,
			CreateTime:   create,
			Suffix:       suffix,
			Size:         size,
			UserId:       uid,
			AbsolutePath: saveFilePath,
			FileMD5:      md5,
			FileMd5Info: models.FileMd5Info{
				MD5: md5,
				Url: url,
			},
		})

		frs = append(frs, response.SimpleFileResponse{
			Status:  "ok",
			Message: "上传成功",
			Name:    fileName,
			Create:  utils.FormatDate(create),
			Url:     url,
		})
	}

	return frs
}

func (fs FileServiceImpl) UploadAvatarFile(ctx *gin.Context) []response.SimpleFileResponse {
	var uid *int = nil

	return fs.GlobalUploadFile(ctx, true, uid)
}

func (fs FileServiceImpl) UploadImageFile(ctx *gin.Context) []response.SimpleFileResponse {

	var user = utils.GetUserInfo(ctx)

	return fs.GlobalUploadFile(ctx, true, &user.Id)
}

func (fs FileServiceImpl) UploadFile(ctx *gin.Context) []response.SimpleFileResponse {
	var user = utils.GetUserInfo(ctx)

	return fs.GlobalUploadFile(ctx, false, &user.Id)
}

func (fs FileServiceImpl) AddFile(fileInfo models.FileInfo) {
	if err := fs.repository.Save(fileInfo); err != nil {
		helper.ErrorCommonF(common.AutoFail(common.AddFileFail))
	}
}

func NewFileService() FileService {
	var repository = repository.NewFileInfoRepository(config.DB)
	return FileServiceImpl{repository: repository}
}
