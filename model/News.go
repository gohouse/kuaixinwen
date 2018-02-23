package model

import (
	"github.com/devfeel/dotweb"
	"strconv"
	. "github.com/gohouse/kuaixinwen/bootstrap"
)

func NewsAddOrEdit(ctx dotweb.Context) int {
	// 如果传了新闻的主键id则修改, 否则增加
	news := DB.Table("news").
		Data(map[string]string{"content": ctx.FormValue("content")})

	// 增加
	if ctx.FormValue("id") == "" {
		res, err := news.Insert()
		if err != nil {
			return 0
		}
		return res
	}

	// 修改
	res, err := news.
		Where("id", ctx.FormValue("id")).
		Update()
	if err != nil {
		return 0
	}
	return res
}

// 删除
func NewsDel(ctx dotweb.Context) int {
	res, err := DB.Table("news").
		Where("id", ctx.FormValue("id")).
		Delete()
	if err != nil {
		return 0
	}
	return res
}

// 获取列表
func GetNewsList(ctx dotweb.Context) (interface{}, int) {
	var page int = 1
	var limit int = 10
	pageParam := ctx.FormValue("page")
	if pageParam != "" {
		pageInt, err := strconv.Atoi(pageParam)
		if err != nil {
			return err, 0
		}
		page = pageInt
	}

	limitParam := ctx.FormValue("limit")
	if limitParam != "" {
		limitInt, err := strconv.Atoi(limitParam)
		if err != nil {
			return err, 0
		}
		limit = limitInt
	}
	count, err := DB.Table("news").
		Where("status", 1).
		Count()
	if err != nil {
		return nil, 0
	}

	// 判断数据是否超限
	if count < (limit * (page - 1)) {
		return nil, count
	}

	list, err := DB.Table("news").
		Where("status", 1).
		Order("id desc").
		Limit(limit).
		Page(page).
		Get()
	if err != nil {
		return nil, 0
	}

	return list, count
}

// 根据主键id获取一条数据
func GetNewsById(ctx dotweb.Context) interface{} {
	res, err := DB.Table("news").
		Where("id", ctx.FormValue("id")).
		First()
	if err != nil {
		return ""
	}
	return res
}
