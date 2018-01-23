package model

import (
	"github.com/devfeel/dotweb"
	"strconv"
	"github.com/gohouse/utils"
	. "github.com/gohouse/kuaixinwen/bootstrap"
)

func NewsAddOrEdit(ctx dotweb.Context) int {
	// 如果传了新闻的主键id则修改, 否则增加

	// 增加
	if ctx.FormValue("id")==""{
		return DB.Table("news").
			Data(map[string]string{"content": ctx.FormValue("content")}).
			Insert()
	}

	// 修改
	return DB.Table("news").
		Data(map[string]string{"content": ctx.FormValue("content")}).
		Where("id", ctx.FormValue("id")).
		Update()
}
// 删除
func NewsDel(ctx dotweb.Context) int {
	return DB.Table("news").
		Where("id", ctx.FormValue("id")).
		Delete()
}
// 获取列表
func GetNewsList(ctx dotweb.Context) interface{} {
	var page int = 1
	var limit int = 10
	pageParam := ctx.FormValue("page")
	if pageParam !=""{
		pageInt,err := strconv.Atoi(pageParam)
		utils.CheckErr(err)
		page = pageInt
	}

	limitParam := ctx.FormValue("limit")
	if limitParam !=""{
		limitInt,err := strconv.Atoi(limitParam)
		utils.CheckErr(err)
		limit = limitInt
	}

	return DB.Table("news").
		Where("status", 1).
		Order("id desc").
		Limit(limit).
		Page(page).
		Get()
}
// 根据主键id获取一条数据
func GetNewsById(ctx dotweb.Context) interface{} {
	return DB.Table("news").
		Where("id", ctx.FormValue("id")).
		First()
}
