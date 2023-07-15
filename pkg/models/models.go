package models

type Note struct {
	Id     int64  `json:"id" gorm:"primarykey;auto_increment"`
	Userid int64  `json:"userid"`
	Note   string `json:"note"`
}
