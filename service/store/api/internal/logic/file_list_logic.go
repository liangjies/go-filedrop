package logic

import (
	"context"
	"fmt"
	"math"

	"go-filedrop/service/store/api/internal/svc"
	"go-filedrop/service/store/api/internal/types"
	"go-filedrop/service/store/api/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileListLogic {
	return &FileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileListLogic) FileList(req *types.FileListRequest) (resp *types.FileListReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.FileListReply)
	var fileUpload []models.FileUpload
	if req.Page == 0 {
		req.Page = 1
	}
	// 总数
	err = l.svcCtx.DB.Model(&models.FileUpload{}).Count(&resp.Total).Error
	if err != nil {
		return
	}

	limit := 10
	offset := 10 * (req.Page - 1)
	err = l.svcCtx.DB.Model(&models.FileUpload{}).Order("id desc").Limit(limit).Offset(offset).Find(&fileUpload).Error
	if err != nil {
		return resp, err
	}
	fmt.Println(fileUpload)
	resp.List = fileUpload
	resp.Page = req.Page
	resp.PageSize = int(math.Ceil(float64(resp.Total) / float64(limit)))

	return resp, err
}
