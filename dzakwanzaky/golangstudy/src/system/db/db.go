package db

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func Connect(host string, port string, user string, pass string, database string, options string) (db *xorm.Engine, err error) {
	return xorm.NewEngine("mysql", "root:@tcp(localhost:3306)/golang?charset-utf8")
}

func Find(DB *xorm.Engine, findBy interface{}, objects interface{}) error {
	return DB.Find(objects, findBy)
}

func FindBy(DB *xorm.Engine, model interface{}) (err error) {
	_, err = DB.Get(model)
	return
}

func Exists(DB *xorm.Engine, model interface{}) (bool, error) {
	return DB.Get(model)
}

func Update(DB *xorm.Engine, id int64, model interface{}) (err error) {
	_, err = DB.ID(id).Update(model)
	return
}

func Store(DB *xorm.Engine, model interface{}) (err error) {
	_, err = DB.Insert(model)
	return
}

func Destroy(DB *xorm.Engine, id int64, model interface{}) (err error) {
	_, err = DB.ID(id).Delete(model)
	return
}
