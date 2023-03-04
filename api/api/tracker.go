package api

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"tracker-manager-go/common/result"
	"tracker-manager-go/service"
)

type TrackerApi struct {
	service.SyncSourceService
}

func (t TrackerApi) GetSourceList(ctx *fiber.Ctx) error {
	syncSources := t.SyncSourceService.List()
	text := ""
	for _, syncSource := range syncSources {
		content := syncSource.Content
		if strings.Contains(content, "\n\n") {
			content = strings.ReplaceAll(content, "\n\n", "\n")
		}
		text += content
	}
	return result.Success(ctx, text)
}
