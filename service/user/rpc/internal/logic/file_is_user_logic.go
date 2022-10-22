package logic

import (
	"context"
	"errors"

	"go-filedrop/service/user/rpc/internal/svc"
	"go-filedrop/service/user/rpc/models"
	"go-filedrop/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileIsUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileIsUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileIsUserLogic {
	return &FileIsUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileIsUserLogic) FileIsUser(in *user.FileIsUserReq) (res *user.FileIsUserReply, err error) {
	// todo: add your logic here and delete this line
	res = new(user.FileIsUserReply)
	res.IsUser = true
	// 查询用户是否存在
	user := new(models.UserBasic)
	err = l.svcCtx.DB.Model(&models.UserBasic{}).Where("id = ? ", in.UserId).First(&user).Error
	if err != nil {
		res.IsUser = false
		return res, errors.New("用户不存在")
	}
	// 查询文件是否属于用户
	file := new(models.FileUpload)
	err = l.svcCtx.DB.Model(&models.FileUpload{}).Where("uid = ? AND id = ?", in.UserId, in.FID).First(&file).Error
	if err != nil {
		res.IsUser = false
		return res, errors.New("文件不属于用户")
	}

	return res, nil
}
