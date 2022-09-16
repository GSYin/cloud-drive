package logic

import (
	"cloud-drive/core/models"
	"cloud-drive/core/util"
	"context"
	"errors"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 1. 从数据库中查询当前用户
	user := new(models.User)
	has, err := models.Engine.Where("username = ? AND password = ?", req.Username, util.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或密码错误")
	}
	// 2. 生成token
	token, err := util.GenerateToken(user.Id, user.Identity, user.Username)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Token = token

	return
}
