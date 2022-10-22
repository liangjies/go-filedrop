package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-filedrop/service/user/api/internal/logic"
	"go-filedrop/service/user/api/internal/svc"
	"go-filedrop/service/user/api/internal/types"
)

func UserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.UserDetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}