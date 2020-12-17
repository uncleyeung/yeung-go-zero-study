package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/uncleyeung/yeung-go-zero-study/rpc/add/internal/config"
	"github.com/uncleyeung/yeung-go-zero-study/rpc/model"
)

type ServiceContext struct {
	c     config.Config
	Model model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:     c,
		Model: model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
