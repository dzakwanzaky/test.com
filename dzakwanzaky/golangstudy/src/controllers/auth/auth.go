package session

import (
	Users "test.com/dzakwanzaky/golangstudy/pkg/types/users"

	"xorm.io/xorm"
)

var db *xorm.Engine

type LoginData struct {
	Token string     `json:"token"`
	User  Users.User `json:"user"`
}

func Init(DB *xorm.Engine) {
	db = DB
}
