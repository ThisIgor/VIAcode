package storage

import (
	"github.com/jinzhu/gorm"
	"knowledge/database"
)

type Upload struct {
	gorm.Model
	Path           string `gorm:"type:varchar(1024);not null" json:"name, db:name"`
	Comment        string `gorm:"type:varchar(1024);not null" json:"name, db:name"`
	DocumentNumber string `gorm:"type:varchar(1024);not null" json:"name, db:name"`
}

func CreateFileInfo(data *Upload) *gorm.DB {
	db := database.Connect()
	defer db.Close()
	res := db.Create(data)

	return res
}

func GetFileInfo(fileid uint) *Upload {
	db := database.Connect()
	defer db.Close()
	upl := &Upload{}
	_ = db.Where("id =?", fileid).First(&upl)

	return upl
}

func UpdateFileInfo(fileid uint, Comment string) {
	db := database.Connect()
	defer db.Close()
	upl := &Upload{}
	_ = db.Where("id = ?", fileid).First(&upl)
	upl.Comment = Comment
	db.Save(&upl)
}
