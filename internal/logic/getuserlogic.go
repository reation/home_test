package logic

import (
	"context"

	"homeT/internal/svc"
	"homeT/proto/types/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *proto.IdRequest) (*proto.UserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &proto.UserInfoResponse{
		Id:     1,
		Name:   "reation",
		Gender: 0,
	}, nil
}
