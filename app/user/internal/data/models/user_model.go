package models

type UserModel struct {
	Id       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"type:varchar(50);not null" json:"username"`
	Nickname string `gorm:"type:varchar(50)" json:"nickname"`
	Password string `gorm:"size:255;not null" json:"password"`
}
