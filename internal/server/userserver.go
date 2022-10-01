// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"homeT/internal/logic"
	"homeT/internal/svc"
	"homeT/proto/types/proto"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	proto.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *proto.IdRequest) (*proto.UserInfoResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}