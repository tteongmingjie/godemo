// Code generated by goctl. DO NOT EDIT.
// Source: student.proto

package server

import (
	"context"

	"mymod/godemo/rpc/internal/logic"
	"mymod/godemo/rpc/internal/svc"
	"mymod/godemo/rpc/types/student"
)

type StudentServer struct {
	svcCtx *svc.ServiceContext
	student.UnimplementedStudentServer
}

func NewStudentServer(svcCtx *svc.ServiceContext) *StudentServer {
	return &StudentServer{
		svcCtx: svcCtx,
	}
}

func (s *StudentServer) Create(ctx context.Context, in *student.CreateRequest) (*student.CreateResponse, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *StudentServer) Update(ctx context.Context, in *student.UpdateRequest) (*student.UpdateResponse, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *StudentServer) Remove(ctx context.Context, in *student.RemoveRequest) (*student.RemoveResponse, error) {
	l := logic.NewRemoveLogic(ctx, s.svcCtx)
	return l.Remove(in)
}

func (s *StudentServer) Detail(ctx context.Context, in *student.DetailRequest) (*student.DetailResponse, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}
