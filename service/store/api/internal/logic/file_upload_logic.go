package logic

import (
	"context"
	"fmt"
	"time"

	"go-filedrop/service/store/api/helper"
	"go-filedrop/service/store/api/internal/svc"
	"go-filedrop/service/store/api/internal/types"
	"go-filedrop/service/store/api/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.FileUploadReply)
	// 上传文件
	err, u, filename, key := helper.UploadFile(l.svcCtx.DB, req.File, l.svcCtx.Config.COS.Bucket, l.svcCtx.Config.COS.SecretID, l.svcCtx.Config.COS.SecretKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 保存到数据库
	fileUpload := models.FileUpload{Size: req.File.Size, URL: u.String() + "/" + key, Filename: filename, UploadTime: time.Now(), IP: "127.0.0.1", Key: key}
	err = l.svcCtx.DB.Model(&models.FileUpload{}).Create(&fileUpload).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Code = 200
	resp.Msg = "上传成功"
	return
}
