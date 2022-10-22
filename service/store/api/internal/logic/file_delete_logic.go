package logic

import (
	"context"
	"fmt"

	"go-filedrop/service/store/api/helper"
	"go-filedrop/service/store/api/internal/svc"
	"go-filedrop/service/store/api/internal/types"
	"go-filedrop/service/store/api/models"
	"go-filedrop/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteLogic {
	return &FileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDeleteLogic) FileDelete(req *types.FileDeleteRequest, userId string) (resp *types.FileDeleteReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.FileDeleteReply)
	// 这里使用RPC查询文件是否属于用户
	FileIsUserReply, err := l.svcCtx.UserRpc.FileIsUser(l.ctx, &user.FileIsUserReq{})
	if err != nil {
		fmt.Println(err)
		return
	}
	if !FileIsUserReply.IsUser {
		resp.Code = 500
		resp.Msg = "文件不属于用户"
		return
	}
	// 查询Key
	var fileUpload models.FileUpload
	err = l.svcCtx.DB.Model(&models.FileUpload{}).Where("id = ?", req.ID).First(&fileUpload).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	// 删除数据库
	err = l.svcCtx.DB.Model(&models.FileUpload{}).Delete("id = ?", fileUpload.ID).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	// 删除文件
	err = helper.DeleteFile(fileUpload.Key, l.svcCtx.Config.COS.Bucket, l.svcCtx.Config.COS.SecretID, l.svcCtx.Config.COS.SecretKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Code = 200
	resp.Msg = "删除成功"
	return
}
