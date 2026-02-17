package models

import (
	"time"
)

type Tool struct {
	Id         int       `orm:"pk;auto"`
	Title      string    `orm:"size(100)"`
	Category   string    `orm:"size(50)"`
	Folder     string    `orm:"size(255)"`
	Status     int       `orm:"default(2)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}
