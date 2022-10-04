package logic

import (
	"cloud-drive/core/models"
	"cloud-drive/core/utils"
	"context"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {
	// save the file to database
	rp := &models.RepositoryPool{
		Identity: utils.GenerateUUID(),
		Filehash: req.Filehash,
		Filename: req.Filename,
		Filesize: req.Filesize,
		Fileext:  req.Fileext,
		Filepath: req.Filepath,
	}
	_, err = l.svcCtx.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}
	resp = new(types.FileUploadResponse)
	resp.Identity = rp.Identity
	resp.Fileext = rp.Fileext
	resp.Filename = rp.Filename
	return
}
