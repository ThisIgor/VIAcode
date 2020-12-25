package department

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	Name string `gorm:"type:varchar(1024);unique;not null"`
}
