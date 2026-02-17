package models

import (
	"time"
)

type Article struct {
	Id         int       `orm:"pk;auto"`
	Title      string    `orm:"size(100)"`
	Category   string    `orm:"size(50)"`
	ContentMd  string    `orm:"type(text)"`
	Status     int       `orm:"default(2)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}
