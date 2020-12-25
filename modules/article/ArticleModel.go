package article

import (
	"github.com/jinzhu/gorm"
	"knowledge/database"
	"strings"
)

// TODO:: Все развязочные таблицы называем с префиксом "Rel" (relation)

// Шаблон статьи
type ArticleTemplate struct {
	gorm.Model
	Bkackword     string `gorm:"type:varchar(1024)"`
	ArticleTypeID uint   `gorm:"foreignkey:u_id;association_foreignkey:id" json:"article_type_id, db:article_type_id"`
	Permissions   uint   `gorm:"type:int;not null"`
}

// Статья
type Article struct {
	gorm.Model
	Bkackword     string `gorm:"type:varchar(1024)" json:"blackword, db:blackword"`
	ArticleTypeID uint   //  `gorm:"foreignkey:u_id;association_foreignkey:id" json:"article_type_id, db:article_type_id"`
	Permissions   uint   `gorm:"type:int;not null"`
	Content       string `gorm:"type:text"`
}

//// TODO:: составляющая часть статьи. Пока под вопросом
//type Part struct {
//	gorm.Model
//	ArticleId uint `gorm:"foreignkey:u_id;association_foreignkey:id" json:"article_id, db:article_id"`
//	Content string `gorm:"type:text"`
//}

// Тип статьи ( отраслевая / словарная / ... )
type ArticleType struct {
	gorm.Model
	InternalName string `gorm:"type:varchar(1024);unique;not null"`
	FullName     string `gorm:"type:varchar(1024);unique;not null"`
	Description  string `gorm:"type:varchar(1024)"`
	HeadingID    uint   //TODO::  `gorm:"foreignkey:u_id;association_foreignkey:id" json:"heading_id, db:heading_id"`
	GroupID      uint   //TODO::  `gorm:"foreignkey:u_id;association_foreignkey:id" json:"group_id, db:group_id"` // Связь с группой
}

// Группы типов статей. Связано по id и ArticleType.GroupID
type ArticleTypeGroup struct {
	gorm.Model
	Type        uint   `gorm:"type:int;not nul"`
	Name        string `gorm:"type:varchar(1024);not nul"`
	Nickname    string `gorm:"type:varchar(1024);not nul"`
	Description string `gorm:"type:varchar(1024);not nul"`
}

//TODO:: Под вопросом ------------------------------------
//type Heading struct {
//	gorm.Model
//	Name         string `gorm:"type:varchar(1024);unique;not null"`
//	ShortName    string `gorm:"type:varchar(1024);unique;not null"`
//	DepartmentID uint   `gorm:"foreignkey:u_id;association_foreignkey:id" json:"department_id, db:department_id"`
//}

func CreateArticleTypeGroup(grouptype uint, name, nickname, descripton string) {
	db := database.Connect()
	defer db.Close()
	db.Create(&ArticleTypeGroup{Type: grouptype, Name: name, Nickname: nickname, Description: descripton})
}

func CreateArticleType(name, description string, headingid, GroupID uint) *gorm.DB {
	db := database.Connect()
	var internalname string
	if len(name) > 2 {
		internalname = strings.ToUpper(name[0:3])
	} else {
		internalname = strings.ToUpper(name)
	}
	defer db.Close()
	res := db.Create(&ArticleType{InternalName: internalname, FullName: name, Description: description, HeadingID: headingid, GroupID: GroupID})
	return res
}

func InsertArticle(article *Article) (error, uint) {
	db := database.Connect()
	defer db.Close()
	db.Create(article)
	if db.Error != nil {
		return db.Error, 0
	}
	return nil, 9
}

func FetchAllArticles() (error error, articles []Article) {
	var arts []Article
	db := database.Connect()
	defer db.Close()
	res := db.Find(&arts)
	return res.Error, arts
}

func FetchArticle(articleId int) (Article, *gorm.DB) {
	db := database.Connect()
	defer db.Close()

	data := &Article{}
	res := db.Where("id = ?", articleId).First(&data)
	return *data, res
}
