package logic

import (
	"cloud-drive/core/define"
	"cloud-drive/core/models"
	"context"
	"time"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	uf := make([]*types.UserFile, 0)

	resp = new(types.UserFileListResponse)
	size := req.Size
	if size == 0 {
		size = int64(define.PageSize)
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	err = l.svcCtx.Engine.Table("user_repository").Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.fileext,"+
			"user_repository.filename,repository_pool.filepath,repository_pool.filesize").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at is null", time.Time{}.Format(define.DateTime)).
		Limit(int(size), int(offset)).Find(&uf)
	if err != nil {
		return nil, err
	}
	total, err := l.svcCtx.Engine.Where("parent_id = ? AND user_identity = ?", req.Id, userIdentity).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	resp.List = uf
	resp.Total = total

	return
}
