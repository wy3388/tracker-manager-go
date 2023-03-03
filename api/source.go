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
	if s.SourceService.UpdateById(source) {
		return result.SuccessEmpty(ctx)
	}
	return result.Fail(ctx, "更新失败")
}
