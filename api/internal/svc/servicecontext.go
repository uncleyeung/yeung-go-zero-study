package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/uncleyeung/yeung-go-zero-study/api/internal/config"
	"github.com/uncleyeung/yeung-go-zero-study/rpc/add/adder"
	"github.com/uncleyeung/yeung-go-zero-study/rpc/check/checker"
)

type ServiceContext struct {
	Config config.Config

	//rpc
	Adder   adder.Adder
	Checker checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
