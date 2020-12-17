package logic

import (
	"context"
	"github.com/uncleyeung/yeung-go-zero-study/rpc/add/adder"

	"github.com/uncleyeung/yeung-go-zero-study/api/internal/svc"
	"github.com/uncleyeung/yeung-go-zero-study/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (logic *AddLogic) Add(req types.AddReq) (*types.AddResp, error) {
	//rpc调用
	adderService := logic.svcCtx.Adder
	addReq := adder.AddReq{
		Book:  req.Book,
		Price: req.Price,
	}
	addResult, err := adderService.Add(logic.ctx, &addReq)
	if err != nil {
		return nil, err
	}

	return &types.AddResp{
		Ok: addResult.Ok,
	}, nil
}
