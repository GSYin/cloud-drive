package logic

import (
	"cloud-drive/core/models"
	"cloud-drive/core/utils"
	"context"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareResponse, err error) {
	rp := new(models.RepositoryPool)
	has, err := l.svcCtx.Engine.Where("filehash = ?", req.Md5).Get(rp)
	if err != nil {
		return nil, err
	}
	resp = new(types.FileUploadPrepareResponse)
	if has {
		// upload success
		resp.Identity = rp.Identity
	} else {
		// get the upload id and key,and upload to cos
		key, uploadId, err := utils.CosInitPartUpload(req.FileExt)
		if err != nil {
			return nil, err
		}
		resp.Key = key
		resp.UploadId = uploadId
	}

	return
}
