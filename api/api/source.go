package api

import (
	"github.com/gofiber/fiber/v2"
	"tracker-manager-go/common/result"
	"tracker-manager-go/model"
	"tracker-manager-go/service"
)

type SourceApi struct {
	service.SourceService
}

func (s SourceApi) List(ctx *fiber.Ctx) error {
	return result.Success(ctx, s.SourceService.List())
}

func (s SourceApi) UpdateById(ctx *fiber.Ctx) error {
	var source model.Source
	if e := ctx.BodyParser(&source); e != nil {
		return result.Fail(ctx, e.Error())
	}
	if e := s.SourceService.UpdateById(source.UUID, map[string]any{
		"author":     source.Author,
		"key":        source.Key,
		"name":       source.Name,
		"url":        source.Url,
		"is_checked": source.IsChecked,
	}); e != nil {
		return result.Fail(ctx, "更新失败")

	}
	return result.SuccessEmpty(ctx)
}

func (s SourceApi) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	if id == "" {
		return result.Fail(ctx, "参数错误")
	}
	return result.Success(ctx, s.SourceService.GetById(id))
}

func (s SourceApi) Save(ctx *fiber.Ctx) error {
	var source model.Source
	if e := ctx.BodyParser(&source); e != nil {
		return result.Fail(ctx, e.Error())
	}
	if e := s.SourceService.Save(source); e != nil {
		return result.Fail(ctx, "保存失败")
	}
	return result.SuccessEmpty(ctx)
}

func (s SourceApi) DeleteById(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	if id == "" {
		return result.Fail(ctx, "参数错误")
	}
	if e := s.SourceService.DeleteById(id); e != nil {
		return result.Fail(ctx, "删除失败")
	}
	return result.SuccessEmpty(ctx)
}
