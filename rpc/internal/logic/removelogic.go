package logic

import (
	"context"

	"mymod/godemo/model"
	"mymod/godemo/rpc/internal/svc"
	"mymod/godemo/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *student.RemoveRequest) (*student.RemoveResponse, error) {
	// todo: add your logic here and delete this line
	newStudent := model.Student{}
	// 查询是否存在
	res := l.svcCtx.MySqlDb.Where("id = ?", in.Id).First(&newStudent)
	err := res.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(100, "不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	res = l.svcCtx.MySqlDb.Delete(&newStudent, in.Id)
	err = res.Error
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &student.RemoveResponse{}, err
}
