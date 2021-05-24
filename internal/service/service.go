package service

import (
	"context"
	"github.com/JoyZF/blog_gin/global"
	"github.com/JoyZF/blog_gin/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service  {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
