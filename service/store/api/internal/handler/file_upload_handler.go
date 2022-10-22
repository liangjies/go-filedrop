package handler

import (
	"net/http"

	"go-filedrop/service/store/api/internal/logic"
	"go-filedrop/service/store/api/internal/svc"
	"go-filedrop/service/store/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		_, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		req.File = fileHeader
		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
