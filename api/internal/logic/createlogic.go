package logic

import (
	"context"

	"mymod/godemo/api/internal/svc"
	"mymod/godemo/api/internal/types"
	"mymod/godemo/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.StudentRpc.Create(l.ctx, &student.CreateRequest{
		Name:   req.Name,
		StuNo:   req.StuNo,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
