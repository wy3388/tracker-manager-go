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
	app.Static("/", "./web")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", nil)
	})

	v1 := app.Group("api/v1")
	{
		sourceGroup := v1.Group("source")
		{
			var s api.SourceApi
			sourceGroup.Get("", s.List)
			sourceGroup.Put("", s.UpdateById)
			sourceGroup.Get(":id", s.GetById)
			sourceGroup.Post("", s.Save)
			sourceGroup.Delete(":id", s.DeleteById)
		}

		trackerGroup := v1.Group("tracker")
		{
			var t api.TrackerApi
			trackerGroup.Get("", t.GetSourceList)
		}

		syncGroup := v1.Group("sync")
		{
			var s api.SyncApi
			syncGroup.Post("syncTracker", s.SyncTracker)
			syncGroup.Get("", s.List)
			syncGroup.Post("deleteByIds", s.DeleteByIds)
			syncGroup.Get("clear", s.Clear)
		}

		clientGroup := v1.Group("client")
		{
			var c api.ClientApi
			clientGroup.Get("", c.List)
			clientGroup.Post("", c.Save)
			clientGroup.Get(":id", c.GetById)
			clientGroup.Put("", c.UpdateById)
			clientGroup.Delete(":id", c.DeleteById)
		}
	}

	app.Get("/root/*", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/")
	})

	return app
}
