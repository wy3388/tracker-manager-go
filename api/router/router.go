package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"tracker-manager-go/api"
	"tracker-manager-go/conf"
)

func New() *fiber.App {
	var app *fiber.App
	if conf.App.Debug {
		app = fiber.New()
	} else {
		app = fiber.New(fiber.Config{
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				if err != nil {
					return ctx.JSON(fiber.Map{
						"code":    1,
						"message": err.Error(),
						"data":    fiber.Map{},
					})
				}
				return nil
			},
		})
		app.Use(recover2.New())
	}

	app.Use(cors.New())
	v1 := app.Group("api/v1")

	var source api.SourceApi
	v1.Get("source", source.List)
	v1.Put("source", source.UpdateById)
	v1.Get("source/:id", source.GetById)
	v1.Post("source/", source.Save)
	v1.Delete("source/:id", source.DeleteById)

	var tracker api.TrackerApi
	v1.Get("tracker", tracker.GetSourceList)

	var sync api.SyncApi
	v1.Post("sync/syncTracker", sync.SyncTracker)
	v1.Get("sync", sync.List)
	v1.Post("sync/deleteByIds", sync.DeleteByIds)
	v1.Get("sync/clear", sync.Clear)
	return app
}
