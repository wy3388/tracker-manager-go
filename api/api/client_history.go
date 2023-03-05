package api

import (
	"github.com/gofiber/fiber/v2"
	"tracker-manager-go/common/result"
	"tracker-manager-go/service"
)

type ClientHistoryApi struct {
	service.ClientHistoryService
}

func (c ClientHistoryApi) List(ctx *fiber.Ctx) error {
	return result.Success(ctx, c.ClientHistoryService.List())
}

func (c ClientHistoryApi) DeleteByIds(ctx *fiber.Ctx) error {
	var ids []string
	if err := ctx.BodyParser(&ids); err != nil {
		return result.Fail(ctx, err.Error())
	}
	if err := c.ClientHistoryService.DeleteByIds(ids); err != nil {
		return result.Fail(ctx, err.Error())
	}
	return result.SuccessEmpty(ctx)
}
