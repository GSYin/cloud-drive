package handler

import (
	"net/http"

	"cloud-drive/core/internal/logic"
	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserDetialHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserDetialLogic(r.Context(), svcCtx)
		resp, err := l.UserDetial(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
