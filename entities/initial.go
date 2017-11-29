package entities

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var mydb  *gorm.DB

func init() {
  db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
    if err != nil {
        panic(err)
    }
  defer db.Close()
}


func GetDB() {
    return mydb
}