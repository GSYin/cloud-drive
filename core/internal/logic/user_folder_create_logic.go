package logic

import (
	"cloud-drive/core/models"
	"cloud-drive/core/utils"
	"context"
	"errors"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	// current filename exist
	cnt, err := l.svcCtx.Engine.Where("filename = ? AND parent_id = ?", req.Filename, req.ParentId).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("filename exist")
	}
	// create user folder
	data := &models.UserRepository{
		Identity:     utils.GenerateUUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Filename:     req.Filename,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}
	return
}
