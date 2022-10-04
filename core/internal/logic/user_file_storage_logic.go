package logic

import (
	"cloud-drive/core/models"
	"cloud-drive/core/utils"
	"context"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileStorageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileStorageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileStorageLogic {
	return &UserFileStorageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileStorageLogic) UserFileStorage(req *types.UserFileStorageRequest, userIdentity string) (resp *types.UserFileStorageResponse, err error) {
	ur := &models.UserRepository{
		Identity:           utils.GenerateUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Fileext:            req.Fileext,
		Filename:           req.Filename,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	if err != nil {
		return nil, err
	}
	return
}
