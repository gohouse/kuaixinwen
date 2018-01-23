package controller

import (
	"github.com/devfeel/dotweb"
	"github.com/gohouse/utils"
	"github.com/gohouse/kuaixinwen/model"
)

func GetNewsList(ctx dotweb.Context) error {
	res := model.GetNewsList(ctx)

	ctx.WriteJson(utils.SuccessReturn(res))

	return nil
}

func GetNewsById(ctx dotweb.Context) error {
	res := model.GetNewsById(ctx)

	ctx.WriteJson(utils.SuccessReturn(res))

	return nil
}

func NewsAddOrEdit(ctx dotweb.Context) error {
	res := model.NewsAddOrEdit(ctx)

	if res > 0 {
		ctx.WriteJson(utils.SuccessReturn(res))
	} else {
		ctx.WriteJson(utils.FailReturn("插入或更新失败!!!"))
	}

	return nil
}
func NewsDel(ctx dotweb.Context) error {
	res := model.NewsDel(ctx)

	if res > 0 {
		ctx.WriteJson(utils.SuccessReturn())
	} else {
		ctx.WriteJson(utils.FailReturn("删除失败!!!"))
	}

	return nil
}
