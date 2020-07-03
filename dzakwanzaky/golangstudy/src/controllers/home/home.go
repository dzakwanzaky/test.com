package home

import (
	"xorm.io/xorm"
)

var DB *xorm.Engine

func Init(DB *xorm.Engine) {
	DB = DB
}
