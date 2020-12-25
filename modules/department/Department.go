package department

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	// 1
	Name string `gorm:"type:varchar(1024);unique;not null"` // 2
	Init(data string)
}
