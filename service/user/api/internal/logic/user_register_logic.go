package logic

import (
	"context"
	"errors"
	"log"

	"go-filedrop/service/user/api/helper"
	"go-filedrop/service/user/api/internal/svc"
	"go-filedrop/service/user/api/internal/types"
	"go-filedrop/service/user/api/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	// todo: add your logic here and delete this line
	// 判断用户名是否已存在
	var cnt int64
	err = l.svcCtx.DB.Model(&models.UserBasic{}).Where("name = ?", req.Name).Count(&cnt).Error
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("用户名已存在")
		return
	}
	// 数据入库
	user := &models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	result := l.svcCtx.DB.Create(&user)
	if result.Error != nil {
		return nil, err
	}
	log.Println("insert user row:", result.RowsAffected)
	return &types.UserRegisterReply{Code: 200, Msg: "注册成功"}, err
}
