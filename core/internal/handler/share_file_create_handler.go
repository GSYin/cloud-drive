package handler

import (
	"net/http"

	"cloud-drive/core/internal/logic"
	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShareFileCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareFileCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShareFileCreateLogic(r.Context(), svcCtx)
		resp, err := l.ShareFileCreate(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
