package api

import (
	"github.com/gofiber/fiber/v2"
	"tracker-manager-go/common/result"
	"tracker-manager-go/service"
)

type SyncApi struct {
	service.SyncService
	service.SyncLogService
}

func (s SyncApi) SyncTracker(ctx *fiber.Ctx) error {
	if err := s.SyncService.Sync(); err != nil {
		return result.Fail(ctx, "同步失败")
	}
	return result.SuccessEmpty(ctx)
}

func (s SyncApi) List(ctx *fiber.Ctx) error {
	return result.Success(ctx, s.SyncLogService.List())
}

func (s SyncApi) DeleteByIds(ctx *fiber.Ctx) error {
	var ids []string
	if err := ctx.BodyParser(&ids); err != nil {
		return result.Fail(ctx, err.Error())
	}
	if err := s.SyncLogService.DeleteByIds(ids); err != nil {
		return result.Fail(ctx, err.Error())
	}
	return result.SuccessEmpty(ctx)
}

func (s SyncApi) Clear(ctx *fiber.Ctx) error {
	if err := s.SyncService.Clear(); err != nil {
		return result.Fail(ctx, "清空失败")
	}
	return result.SuccessEmpty(ctx)
}
