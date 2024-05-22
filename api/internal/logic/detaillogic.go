package logic

import (
	"context"

	"mymod/godemo/api/internal/svc"
	"mymod/godemo/api/internal/types"
	"mymod/godemo/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.StudentRpc.Detail(l.ctx, &student.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:    res.Id,
		Name:  res.Name,
		StuNo: res.StuNo,
	}, nil
}
