package svc

import (
	"mymod/godemo/rpc/internal/config"

	"mymod/godemo/rpc/internal/database"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	MySqlDb *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:  c,
		MySqlDb: database.NewMysqlDB(c.Mysql.DataSource),
	}
}
