package logic

import (
	"cloud-drive/core/models"
	"context"
	"errors"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetialLogic {
	return &UserDetialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetialLogic) UserDetial(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	resp = &types.UserDetailResponse{}
	userInfo := new(models.User)
	has, err := models.Engine.Where("identity=?", req.Identity).Get(userInfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user not found")
	}
	resp.Username = userInfo.Username
	resp.Email = userInfo.Email
	return
}
