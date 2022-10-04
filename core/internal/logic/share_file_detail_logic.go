package logic

import (
	"context"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFileDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFileDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFileDetailLogic {
	return &ShareFileDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFileDetailLogic) ShareFileDetail(req *types.ShareFileDetailRequest) (resp *types.ShareFileDetailResponse, err error) {
	// Add + 1 to the number of clicks on the shared record
	_, err = l.svcCtx.Engine.Exec("update share set click_num = click_num + 1 where identity = ?", req.Identity)
	if err != nil {
		return nil, err
	}
	// get file detail
	resp = new(types.ShareFileDetailResponse)
	_, err = l.svcCtx.Engine.Table("share").
		Select("share.repository_identity,user_repository.filename,repository_pool.filesize,repository_pool.fileext,repository_pool.filepath").
		Join("LEFT", "repository_pool", "share.repository_identity = repository_pool.identity").
		Join("LEFT", "user_repository", "user_repository.identity = share.user_repository_identity").
		Where("share.identity = ?", req.Identity).Get(resp)
	if err != nil {
		return nil, err
	}
	return
}
