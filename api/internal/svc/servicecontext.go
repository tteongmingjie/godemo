package svc

import (
	"mymod/godemo/api/internal/config"
	"mymod/godemo/rpc/studentclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	StudentRpc studentclient.Student
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		StudentRpc: studentclient.NewStudent(zrpc.MustNewClient(c.StudentRpc)),
	}
}
