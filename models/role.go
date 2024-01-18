package models

const roleTableName = "roles"

type Role struct {
	//角色ID
	Id int `gorm:"primaryKey"`
	//角色的名称
	Name string `gorm:"column:name;not null;unique"`
	//角色的描述
	Description string `gorm:"column:description;not null"`
}

func (*Role) TableName() string { return roleTableName }
