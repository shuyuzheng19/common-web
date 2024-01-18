package repository

import (
	"common-web-framework/models"
	"gorm.io/gorm"
)

type FileRepository interface {
	// Save 添加一个用户
	Save(file models.FileInfo) error
	// FindById 用过ID查询一个用户
	FindById(id int) *models.FileInfo
	// FindByMd5 通过md5查询url
	FindByMd5(md5 string) string
	// FindAll 查询所有用户
	FindAll() []models.FileInfo
}

type FileRepositoryImpl struct {
	db *gorm.DB
}

func (u FileRepositoryImpl) FindByMd5(md5 string) string {
	var r string
	u.db.Model(&models.FileMd5Info{}).Select("url").Where("md5 = ?", md5).Scan(&r)
	return r
}

func (u FileRepositoryImpl) Save(file models.FileInfo) error {
	return u.db.Model(&models.FileInfo{}).Create(&file).Error
}

func (u FileRepositoryImpl) FindById(id int) *models.FileInfo {
	var result models.FileInfo

	if err := u.db.Preload("FileMd5").First(&result, "id = ?", id).Error; err != nil {
		return nil
	}

	return &result
}

func (u FileRepositoryImpl) FindAll() []models.FileInfo {
	var result = make([]models.FileInfo, 0)

	u.db.Preload("FileMd5").Find(&result)

	return result
}

func NewFileInfoRepository(db *gorm.DB) FileRepository {
	db.AutoMigrate(&models.FileInfo{}, &models.FileMd5Info{})
	return FileRepositoryImpl{db: db}
}
