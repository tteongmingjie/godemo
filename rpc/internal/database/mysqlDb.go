package database

import (
	"mymod/godemo/model"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB(dataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dataSource))
	if err != nil {
		panic("连接数据库失败")
	}

	// 迁移 schema
	err = db.AutoMigrate(&model.Student{})
	if err != nil {
		panic(err)
	}

	return db
}
