package models

type Person struct {
    Id 				int64 `gorm:"primary_key" json:"id"`
	Fullname		string `gorm:"type:varchar(300)" json:"fullname"`
	Email 			string `gorm:"type:varchar(300)" json:"email"`
	Address 		string `gorm:"type:varchar(300)" json:"address"`
}