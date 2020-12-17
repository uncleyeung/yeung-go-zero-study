package svc

import "github.com/uncleyeung/yeung-go-zero-study/rpc/check/internal/config"

type ServiceContext struct {
	c config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c: c,
	}
}
