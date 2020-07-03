package status

import (
	"xorm.io/xorm"
)

var db *xorm.Engine

func Init(DB *xorm.Engine) {
	db = DB
}
