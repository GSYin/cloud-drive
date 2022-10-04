package logic

import (
	"cloud-drive/core/define"
	"cloud-drive/core/models"
	"cloud-drive/core/utils"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(req *types.FileUploadChunkCompleteRequest) (resp *types.FileUploadChunkCompleteResponse, err error) {
	co := make([]cos.Object, 0)
	for _, v := range req.CosObjects {
		co = append(co, cos.Object{
			ETag:       v.Etag,
			PartNumber: int(v.ChunkNumber),
		})
	}
	err = utils.CosPartUploadComplete(req.Key, req.UploadId, co)
	rp := &models.RepositoryPool{
		Identity: utils.GenerateUUID(),
		Filehash: req.Md5,
		Filename: req.Filename,
		Fileext:  req.Fileext,
		Filesize: req.Filesize,
		Filepath: define.CosAddr + "/" + req.Key,
	}
	_, err = l.svcCtx.Engine.Insert(rp)
	return
}
