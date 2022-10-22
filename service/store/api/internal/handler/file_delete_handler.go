package handler

import (
	"net/http"

	"go-filedrop/service/store/api/internal/logic"
	"go-filedrop/service/store/api/internal/svc"
	"go-filedrop/service/store/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewFileDeleteLogic(r.Context(), svcCtx)
		resp, err := l.FileDelete(&req, r.Header.Get("UserId"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
