package result

import "github.com/gofiber/fiber/v2"

type ApiResult struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func New(code uint, message string, data any) *ApiResult {
	return &ApiResult{
		code,
		message,
		data,
	}
}

func Result(ctx *fiber.Ctx, result *ApiResult) error {
	return ctx.JSON(result)
}

func SuccessEmpty(ctx *fiber.Ctx) error {
	return Success(ctx, fiber.Map{})
}

func Success(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(New(0, "success", data))
}

func Fail(ctx *fiber.Ctx, message string) error {
	return ctx.JSON(New(1, message, fiber.Map{}))
}
