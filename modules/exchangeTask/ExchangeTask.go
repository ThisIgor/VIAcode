package exchangeTask

import (
	"github.com/jinzhu/gorm"
	"time"
)

type EditorialType struct {
	gorm.Model
	Type uint   `gorm:"type:int;not null"`
	Name string `gorm:"type:varchar(1024);unique;not null"`
}

type EditorialOffer struct {
	gorm.Model
	EditorialType     uint      `gorm:"foreignkey:u_id;association_foreignkey:id" json:"editorial_type_id, db:editorial_type_id"`
	EditorialMart     uint      `gorm:"foreignkey:u_id;association_foreignkey:id" json:"editorial_mart_id, db:editorial_mart_id"`
	ArticleType       uint      `gorm:"foreignkey:u_id;association_foreignkey:id" json:"article_type_id, db:article_type_id"`
	EditorialOfficeId uint      `gorm:"foreignkey:u_id;association_foreignkey:id" json:"editorial_office_id, db:editorial_office_id"`
	Volume            uint      `gorm:"type:int;not null"`
	StartDate         time.Time `gorm:"type:time"`
	EndDate           time.Time `gorm:"type:time"`
	EditorID          uint      `gorm:"foreignkey:u_id;association_foreignkey:id" json:"user_id, db:user_id"` //[id редактора, имеющего право согласиться с предложением]
}

type EditorialOffice struct {
	gorm.Model
	Description string `gorm:"type:varchar(1024)"`
	Floor       int    `gorm:"type:int;not null"`
	Room        string `gorm:"type:varchar(16);not null"`
}

type EditorialMart struct {
	gorm.Model
	Description string `gorm:"type:varchar(1024)"`
}
