package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-filedrop/service/user/api/internal/logic"
	"go-filedrop/service/user/api/internal/svc"
	"go-filedrop/service/user/api/internal/types"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UserRegister(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}