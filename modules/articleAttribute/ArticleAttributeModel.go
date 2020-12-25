package articleAttribute

import (
	"github.com/jinzhu/gorm"
	"knowledge/database"
)

//  Атрибуты статьи
type ArticleAttribute struct {
	gorm.Model
	Name string `gorm:"type:varchar(1024);unique;not null"`
}

// связка статей с атрибутами
type RelArticleAttribute struct {
	gorm.Model
	AttributeID uint   `gorm:"foreignkey:u_id;association_foreignkey:id" json:"attribute_id, db:attribute_id"`
	ArticleID   uint   `gorm:"foreignkey:u_id;association_foreignkey:id" json:"article_id, db:article_id"`
	Value       string `gorm:"type:varchar(1024);not null"`
}

func CreateArticleAttrib(articleAttribute *ArticleAttribute) error {
	db := database.Connect()
	defer db.Close()
	res := db.Create(articleAttribute)
	return res.Error
}

func AddArticleAttrib(attributeid, articleid uint, value string) {
	db := database.Connect()
	defer db.Close()
	db.Create(&RelArticleAttribute{AttributeID: attributeid, ArticleID: articleid, Value: value})
}

func GetArticleAttrib(attributeid uint) (error error, Name, Value string) {
	db := database.Connect()
	defer db.Close()
	relattr := RelArticleAttribute{}
	db.Where("attribute_id = %v", attributeid).First(&relattr)

	attr := ArticleAttribute{}
	db.Where("attribute_id = %v", attributeid).First(&attr)

	return db.Error, attr.Name, relattr.Value
}

func GetAllArticleAttrib(documentid uint) (RelArticleAttribute, *gorm.DB) {
	db := database.Connect()
	defer db.Close()

	relattr := RelArticleAttribute{}
	res := db.Where("document_id = %v", documentid).Find(&relattr)

	return relattr, res
}

func ChangeArticleAttrib(attributeid uint, attribute *ArticleAttribute) (ArticleAttribute, *gorm.DB) {
	db := database.Connect()
	defer db.Close()

	data := ArticleAttribute{}
	res := db.Where("attribute_id = ?", attributeid).Update(attribute).First(data)
	return data, res
}

func DeleteArticleAttrib(attributeid uint) (RelArticleAttribute, *gorm.DB) {
	db := database.Connect()
	defer db.Close()

	relattr := RelArticleAttribute{}
	res := db.Where("attribute_id = %v", attributeid).First(&relattr)
	if res.Error == nil {
		res = db.Delete(&relattr).First(&relattr)
	}
	return relattr, res
}

func ChangeArticleAttribName(attributeid uint, name string) error {
	db := database.Connect()
	defer db.Close()

	db.Where("attribute_id = %v", attributeid).Update("name", name)
	return db.Error
}
