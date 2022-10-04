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

type SharedFilesSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSharedFilesSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SharedFilesSaveLogic {
	return &SharedFilesSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SharedFilesSaveLogic) SharedFilesSave(req *types.SharedFilesSaveRequest, userIdentity string) (resp *types.SharedFilesSaveResponse, err error) {
	// get file detail
	rp := new(models.RepositoryPool)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.RepositoryIdentity).Get(rp)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("repository file not found")
	}

	// user file repository save
	ur := &models.UserRepository{
		Identity:           utils.GenerateUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Filename:           rp.Filename,
		Fileext:            rp.Fileext,
	}
	_, err = l.svcCtx.Engine.Insert(ur)
	resp = new(types.SharedFilesSaveResponse)
	resp.Identity = ur.Identity
	return
}
