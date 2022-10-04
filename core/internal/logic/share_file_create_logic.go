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

type ShareFileCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFileCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFileCreateLogic {
	return &ShareFileCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFileCreateLogic) ShareFileCreate(req *types.ShareFileCreateRequest, userIdentity string) (resp *types.ShareFileCreateResponse, err error) {
	uuid := utils.GenerateUUID()
	ur := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.UserRepositoryIdentity).Get(ur)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user repository file not found")
	}
	data := &models.Share{
		Identity:               uuid,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     ur.RepositoryIdentity,
		ExpiredTime:            int(req.ExpiredTime),
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}
	resp = &types.ShareFileCreateResponse{
		Identity: uuid,
	}
	return
}
