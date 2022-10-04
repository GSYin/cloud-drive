package logic

import (
	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"
	"cloud-drive/core/models"
	"cloud-drive/core/utils"
	"context"
	"errors"
	"log"

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

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// check code is valid
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("mail code is invalid")
	}
	if code != req.Code {
		err = errors.New("code is invalid")
		return
	}

	// check is username exists
	cnt, err := l.svcCtx.Engine.Where("username=?", req.Username).Count(new(models.User))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("username already exists")
		return
	}
	// insert user data into database
	user := &models.User{
		Identity: utils.GenerateUUID(),
		Username: req.Username,
		Password: utils.Md5(req.Password),
		Email:    req.Email,
	}
	one, err := l.svcCtx.Engine.InsertOne(user)
	if err != nil {
		return nil, err
	}
	log.Println(one)
	return
}
