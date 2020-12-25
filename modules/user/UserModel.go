package user

import (
	"github.com/jinzhu/gorm"
	"knowledge/database"
	"time"
)

// Пользователи
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name, db:name"`
	RoleID   uint   `gorm:"type:int;foreignkey:u_id;association_foreignkey:id" json:"role_id, db:role_id"`
	Login    string `gorm:"type:varchar(100);unique_index" json:"login, db:login"`
	Password string `gorm:"type:varchar(255)" json:"password, db:password"`
}

// Ттекущие сессмм
type Session struct {
	gorm.Model
	UserID  uint `gorm:"foreignkey:u_id;association_foreignkey:id" json:"user_id, db:user_id"`
	RoleID  uint `gorm:"foreignkey:u_id;association_foreignkey:id" json:"role_id, db:role_id"`
	Token   string
	Expires time.Time
}

func CreateUser(user *User) *gorm.DB {
	db := database.Connect()
	defer db.Close()
	res := db.Create(user)
	return res
}

func GetAllUsers() User {
	db := database.Connect()
	defer db.Close()
	users := User{}
	db.Find(&users)
	return users
}

func GetUserFromDB(userid uint) (User, *gorm.DB) {
	db := database.Connect()
	defer db.Close()

	user := User{}
	res := db.Where("id = ?", userid).First(&user)
	return user, res
}

func ChangeUser(userid uint, user User) (User, *gorm.DB) {
	db := database.Connect()
	defer db.Close()

	newuser := &User{}
	res := db.Where("id = %v", userid).Update(&user).First(newuser)
	return *newuser, res
}

func DeleteUserFromDB(userid uint) (User, *gorm.DB) {
	db := database.Connect()
	defer db.Close()

	user, db := GetUserFromDB(userid)
	if db.Error == nil {
		db = db.Delete(&user).First(user)
	}
	return user, db
}
