package main

import (
	_hook "task/golang/task3/7hook"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	////_baseCrud.Run(db)
	//_transferAmount.Run(db)
	//_And6reference.Run(db)
	_hook.Run(db)
	//db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}
	//_sqlxTest.Run(db)
	//_book.Run(db)

}
