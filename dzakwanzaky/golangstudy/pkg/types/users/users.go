package users

import "test.com/dzakwanzaky/golangstudy/pkg/db"

type Users []User

type User db.Users

func (u *User) TableName() string {
	return "users"
}
