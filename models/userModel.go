package models


type User struct {
	ID 			int64	`gorm:"primary_key" json:"id"`
	Username 	string `gorm:"type:varchar(100);uniqueIndex" json:"username"`
	Password 	string `gorm:"type:varchar(200)" json:"-"`
	PersonID 	int64  `gorm:"uniqueIndex" json:"-"`
	Person		Person `gorm:"foreignKey:PersonID;constraint:OnUpdate:CASCADE,onDelete:CASCADE;" json:"-"`
}