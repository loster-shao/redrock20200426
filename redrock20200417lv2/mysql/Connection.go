package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DbConn() *gorm.DB {
	userName := "root"
	password := ""
	host     := "127.0.0.1"
	dbName   := "20200417"
	port     := 3306
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName,password, host, port, dbName )
	db, err  := gorm.Open("mysql", connArgs)

	if err != nil {
		fmt.Println("err:", err)
		return nil
	}

	db.SingularTable(true)
	return db
}
