package department

import "github.com/jinzhu/gorm"

@startuml
type Department struct {
	gorm.Model
	//.. Simple Getter ..
	Name string `gorm:"type:varchar(1024);unique;not null"` //.. Simple Getter 1 ..
}
@enduml
