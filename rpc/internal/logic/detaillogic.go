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

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *student.DetailRequest) (*student.DetailResponse, error) {
	// todo: add your logic here and delete this line

	newStudent := model.Student{}
	// 是否存在
	res := l.svcCtx.MySqlDb.First(&newStudent, in.Id)
	err := res.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(100, "不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &student.DetailResponse{
		Id:    newStudent.Id,
		Name:  newStudent.Name,
		StuNo: newStudent.StuNo,
	}, err
}
