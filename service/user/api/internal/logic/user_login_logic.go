package logic

import (
	"context"
	"errors"

	"go-filedrop/service/user/api/define"
	"go-filedrop/service/user/api/helper"
	"go-filedrop/service/user/api/internal/svc"
	"go-filedrop/service/user/api/internal/types"
	"go-filedrop/service/user/api/models"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	// 1、从数据库中查询当前用户
	user := models.UserBasic{}
	err = l.svcCtx.DB.Model(&user).Where("name = ? AND password = ?", req.Name, helper.Md5(req.Password)).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名或密码错误")
	}
	if err != nil {
		return nil, err
	}
	// 2、生成token
	token, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	// 3、生成用于刷新token的token
	refreshToken, err := helper.GenerateToken(user.Id, user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
