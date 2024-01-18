package models

import "time"

// FileInfo 文件信息
type FileInfo struct {
	//文件ID
	Id int `gorm:"primaryKey"`
	//文件原名字
	OldName string `gorm:"column:old_name"`
	//文件新名字
	NewName string `gorm:"column:new_name"`
	//上传文件的用户ID
	UserId *int `gorm:"column:user_id"`
	//用户信息
	User User
	//上传日期
	CreateTime time.Time `gorm:"column:create_time"`
	//文件后缀
	Suffix string `gorm:"column:suffix"`
	//文件大小
	Size int64 `gorm:"column:size"`
	//文件全路径
	AbsolutePath string `gorm:"column:absolute_path"`
	//文件md5
	FileMD5 string `gorm:"column:md5;unique;not null"` // 关联文件表的 MD5 字段
	// 其他关联表的字段...
	FileMd5Info FileMd5Info `gorm:"foreignKey:MD5;references:md5"`
}

type FileMd5Info struct {
	//文件的md5
	MD5 string `gorm:"column:md5;primaryKey"`
	//文件的url
	Url string `gorm:"column:url;unique"`
}
