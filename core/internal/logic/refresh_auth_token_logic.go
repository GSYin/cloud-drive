package logic

import (
	"cloud-drive/core/define"
	"cloud-drive/core/utils"
	"context"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthTokenLogic {
	return &RefreshAuthTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthTokenLogic) RefreshAuthToken(req *types.RefreshAuthTokenRequest, authorization string) (resp *types.RefreshAuthTokenResponse, err error) {
	// Analysis authorization and get userclaims
	uc, err := utils.AnalyzeToken(authorization)
	if err != nil {
		return
	}
	// genreate new token based on userclaims
	token, err := utils.GenerateToken(uc.Id, uc.Identity, uc.Username, define.TokenExpireTime)
	if err != nil {
		return
	}
	// genetate new refresh token
	refreshToken, err := utils.GenerateToken(uc.Id, uc.Identity, uc.Username, define.RefreshTokenExpireTime)
	if err != nil {
		return
	}
	resp = new(types.RefreshAuthTokenResponse)
	resp.RefreshToken = refreshToken
	resp.Token = token
	return
}
