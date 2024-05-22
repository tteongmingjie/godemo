package logic

import (
	"context"

	"mymod/godemo/rpc/internal/svc"
	"mymod/godemo/rpc/types/student"
	"mymod/godemo/model"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *student.UpdateRequest) (*student.UpdateResponse, error) {
	// todo: add your logic here and delete this line
	newStudent := model.Student{}
	// 查询是否存在
	res := l.svcCtx.MySqlDb.First(&newStudent, in.Id)
	err := res.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(100, "不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Name != "" {
		newStudent.Name = in.Name
	}
	if in.StuNo != "" {
		newStudent.StuNo = in.StuNo
	}

	res = l.svcCtx.MySqlDb.Model(&newStudent).Where("id = ?", in.Id).Updates(model.Student{Name: in.Name, StuNo: in.StuNo})
	err = res.Error
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &student.UpdateResponse{}, nil
}
