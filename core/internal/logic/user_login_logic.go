package logic

import (
	"cloud-drive/core/define"
	"cloud-drive/core/models"
	"cloud-drive/core/utils"
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
	has, err := l.svcCtx.Engine.Where("username = ? AND password = ?", req.Username, utils.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或密码错误")
	}
	// 2. 生成token
	token, err := utils.GenerateToken(user.Id, user.Identity, user.Username, define.TokenExpireTime)
	if err != nil {
		return nil, err
	}

	//3. generate refresh token
	refreshToken, err := utils.GenerateToken(user.Id, user.Identity, user.Username, define.RefreshTokenExpireTime)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResponse)
	resp.Token = token
	resp.RefreshToken = refreshToken

	return
}
