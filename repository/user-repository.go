package repository

import (
	"common-web-framework/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	// Save 添加一个用户
	Save(user models.User) error
	// Update 修改用户信息
	Update(user models.User) error
	// FindById 用过ID查询一个用户
	FindById(id int) *models.User
	// FindAll 查询所有用户
	FindAll() []models.User
	// FindByUsernameAndPassword 通过账号和密码查询
	FindByUsernameAndPassword(username, password string) *models.User
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindByUsernameAndPassword(username, password string) *models.User {
	var result models.User

	if err := u.db.Preload("Role").First(&result, "username = ? and password = ?", username, password).Error; err != nil {
		return nil
	}

	return &result
}

func (u UserRepositoryImpl) Save(user models.User) error {
	return u.db.Model(&models.User{}).Create(&user).Error
}

func (u UserRepositoryImpl) Update(user models.User) error {
	return u.db.Model(&models.User{}).Save(&user).Error
}

func (u UserRepositoryImpl) FindById(id int) *models.User {
	var result models.User

	if err := u.db.Preload("Role").First(&result, "id = ?", id).Error; err != nil {
		return nil
	}

	return &result
}

func (u UserRepositoryImpl) FindAll() []models.User {
	var result = make([]models.User, 0)

	u.db.Preload("Role").Find(&result)

	return result
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&models.Role{}, &models.User{})
	return UserRepositoryImpl{db: db}
}
