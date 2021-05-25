package service

import (
	"context"
	"github.com/JoyZF/blog_gin/global"
	"github.com/JoyZF/blog_gin/internal/dao"
	otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service  {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx,global.DBEngine))
	return svc
}
