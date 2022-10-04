package handler

import (
	"cloud-drive/core/models"
	"cloud-drive/core/utils"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-drive/core/internal/logic"
	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		// Determine whether the file exists or not
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUploadResponse{Identity: rp.Identity, Fileext: rp.Fileext, Filename: rp.Filename})
			return
		}

		// save the file into tencent cos
		cosPath, err := utils.CosUploadFile(r)
		if err != nil {
			return
		}

		// transfer request to logic
		req.Filename = fileHeader.Filename
		req.Fileext = path.Ext(fileHeader.Filename)
		req.Filesize = fileHeader.Size
		req.Filehash = hash
		req.Filepath = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
