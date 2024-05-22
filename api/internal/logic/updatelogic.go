package logic

import (
	"context"

	"mymod/godemo/api/internal/svc"
	"mymod/godemo/api/internal/types"
	"mymod/godemo/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.StudentRpc.Update(l.ctx, &student.UpdateRequest{
		Id:     req.Id,
		Name:   req.Name,
		StuNo: req.StuNo,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateResponse{}, nil
}
