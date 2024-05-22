package logic

import (
	"context"

	"mymod/godemo/model"
	"mymod/godemo/rpc/internal/svc"
	"mymod/godemo/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *student.CreateRequest) (*student.CreateResponse, error) {
	// todo: add your logic here and delete this line
	newStudent := model.Student{
		Name:  in.Name,
		StuNo: in.StuNo,
	}
	res := l.svcCtx.MySqlDb.Create(&newStudent)
	err := res.Error
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &student.CreateResponse{
		Id: newStudent.Id,
	}, nil
}
