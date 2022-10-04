package logic

import (
	"cloud-drive/core/models"
	"context"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteResponse, err error) {
	_, err = l.svcCtx.Engine.Where("user_identity = ? AND identity = ?", userIdentity, req.Identity).Delete(new(models.UserRepository))
	return
}
