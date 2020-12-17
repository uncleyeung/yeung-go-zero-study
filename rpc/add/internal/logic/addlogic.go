package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ulule/deepcopier"
	"github.com/uncleyeung/yeung-go-zero-study/rpc/model"

	"github.com/uncleyeung/yeung-go-zero-study/rpc/add/add"
	"github.com/uncleyeung/yeung-go-zero-study/rpc/add/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (logic *AddLogic) Add(in *add.AddReq) (*add.AddResp, error) {
	modelDb := logic.svcCtx.Model
	one, err := modelDb.FindOne(in.Book)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if one != nil {
		return nil, fmt.Errorf("书籍已经存在不需要添加name：%s", in.Book)
	}

	var book model.Book
	_ = deepcopier.Copy(*in).To(&book)
	logic.Info(fmt.Sprintf("创建书籍开始，%+v", book))
	dbResult, err := modelDb.Insert(book)
	if err != nil {
		return nil, err
	}

	id, _ := dbResult.LastInsertId()
	affected, _ := dbResult.RowsAffected()

	logic.Info(fmt.Sprintf("创建书籍完成，%+v 创建结果：id-%d row-%d", in, id, affected))

	resp := add.AddResp{}
	resp.Ok = true
	return &resp, nil
}
