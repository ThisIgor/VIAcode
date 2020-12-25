package rbac

import (
	"github.com/jinzhu/gorm"
	"knowledge/database"
)

// Роль с фронта. Имя и набор ID разрешений
type FERole struct {
	Name          string
	PermissionsID []uint
}

type RolePermisson struct {
	CategoryID  uint
	Permissions uint64
}

// Права (битовая маска)
type PermissionCategory struct {
	gorm.Model
	CategoryID uint   `gorm:"type:int;not null"`
	Name       string `gorm:"type:varchar(100);unique;not null"`
}

type Permission struct {
	gorm.Model
	Name       string `gorm:"type:varchar(100);unique;not null"`
	Permission uint64 `gorm:"type:int;not null"`
	CategoryID uint   `gorm:"type:int;not null"`
}

// Роли

type Role struct {
	gorm.Model
	Name           string          `gorm:"type:varchar(100);unique;not null" json:"Name"`
	Permissions    []RolePermisson `gorm:"-" json:"Permissions"`
	PermissionsStr string          `gorm:"type:varchar(4096)"`
	PermissionsID  string          `gorm:"type:varchar(4096)"`
}

func CreateRole(name string, permissions []uint) error {
	db := database.Connect()
	defer db.Close()
	//	db.Create(&Role{Name: name, Permissions: permissions})
	return db.Error
}
