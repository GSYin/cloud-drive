package handler

import (
	"cloud-drive/core/utils"
	"errors"
	"net/http"

	"cloud-drive/core/internal/logic"
	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadChunkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadChunkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// check key,uploadId,chunkNumber are not empty
		if r.PostForm.Get("key") == "" {
			httpx.Error(w, errors.New("key is empty"))
			return
		}
		if r.PostForm.Get("upload_id") == "" {
			httpx.Error(w, errors.New("upload id is empty"))
			return
		}
		if r.PostForm.Get("chunk_number") == "" {
			httpx.Error(w, errors.New("chunk number is empty"))
			return
		}
		etag, err := utils.CosPartUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileUploadChunkLogic(r.Context(), svcCtx)
		resp, err := l.FileUploadChunk(&req)
		resp = new(types.FileUploadChunkResponse)
		resp.Etag = etag
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
