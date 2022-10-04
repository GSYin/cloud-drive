package logic

import (
	"cloud-drive/core/models"
	"context"
	"errors"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateResponse, err error) {
	// current filename exist
	cnt, err := l.svcCtx.Engine.Where("filename = ? AND parent_id = (SELECT parent_id from user_repository ur where ur.identity = ?)", req.Filename, req.Identity).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("filename exist")
	}
	// update user filename
	data := models.UserRepository{Filename: req.Filename}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).Update(data)
	if err != nil {
		return nil, err
	}
	return
}
