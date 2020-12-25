package editor

import "github.com/jinzhu/gorm"

//  Типы редакций
type EditorType struct {
	gorm.Model
	Name string `gorm:"type:varchar(1024);unique;not null"`
}

// Связка статей с типоми редакции
type Editor struct {
	gorm.Model
	EditorType uint `gorm:"foreignkey:u_id;association_foreignkey:id" json:"editor_type_id, db:editor_type_id"`
	ArticleID  uint `gorm:"foreignkey:u_id;association_foreignkey:id" json:"article_id, db:article_id"`
}
