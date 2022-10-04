package logic

import (
	"cloud-drive/core/define"
	"cloud-drive/core/models"
	"cloud-drive/core/utils"
	"context"
	"errors"
	"time"

	"cloud-drive/core/internal/svc"
	"cloud-drive/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendResponse, err error) {
	// check if email exists
	cnt, err := l.svcCtx.Engine.Where("email=?", req.Email).Count(new(models.User))
	if err != nil {
		return
	}
	if cnt > 0 {
		err = errors.New("email already exists")
		return
	}

	code := utils.GenerateRandCode()

	// save code to redis
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpireTime))

	// send code to email
	err = utils.MailCodeSend(req.Email, code)
	if err != nil {
		return nil, err
	}

	return
}
