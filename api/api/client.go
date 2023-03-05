package api

import (
	"github.com/gofiber/fiber/v2"
	"tracker-manager-go/common/result"
	"tracker-manager-go/model"
	"tracker-manager-go/service"
)

type ClientApi struct {
	service.ClientService
}

func (c ClientApi) List(ctx *fiber.Ctx) error {
	return result.Success(ctx, c.ClientService.List())
}

func (c ClientApi) Save(ctx *fiber.Ctx) error {
	var client model.Client
	if e := ctx.BodyParser(&client); e != nil {
		return result.Fail(ctx, e.Error())
	}
	if e := c.ClientService.Save(client); e != nil {
		return result.Fail(ctx, "保存失败")
	}
	return result.SuccessEmpty(ctx)
}

func (c ClientApi) GetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	if id == "" {
		return result.Fail(ctx, "参数错误")
	}
	return result.Success(ctx, c.ClientService.GetById(id))
}

func (c ClientApi) UpdateById(ctx *fiber.Ctx) error {
	var client model.Client
	if e := ctx.BodyParser(&client); e != nil {
		return result.Fail(ctx, e.Error())
	}
	if e := c.ClientService.UpdateById(client.UUID, map[string]any{
		"name":     client.Name,
		"url":      client.Url,
		"username": client.Username,
		"password": client.Password,
	}); e != nil {
		return result.Fail(ctx, "更新失败")

	}
	return result.SuccessEmpty(ctx)
}

func (c ClientApi) DeleteById(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	if id == "" {
		return result.Fail(ctx, "参数错误")
	}
	if e := c.ClientService.DeleteById(id); e != nil {
		return result.Fail(ctx, "删除失败")
	}
	return result.SuccessEmpty(ctx)
}

func (c ClientApi) Sync(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")
	if id == "" {
		return result.Fail(ctx, "参数错误")
	}
	if err := c.ClientService.Sync(id); err != nil {
		return result.Fail(ctx, err.Error())
	}
	return result.SuccessEmpty(ctx)
}
