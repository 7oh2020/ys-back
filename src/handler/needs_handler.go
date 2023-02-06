package handler

import (
	"ys-back/src/domain/service"

	"github.com/labstack/echo/v4"
)

type INeedsHandler interface {
	// 要望ツイートの更新を行う
	UpdateNeedsTweet(ctx echo.Context) error
}

type NeedsHandler struct {
	srv    service.ITwitterApiService
	logger echo.Logger
}

func NewNeedsHandler(srv service.ITwitterApiService, logger echo.Logger) INeedsHandler {
	return &NeedsHandler{srv, logger}
}

func (hdr *NeedsHandler) UpdateNeedsTweet(ctx echo.Context) error {
	if err := hdr.srv.FetchNeedsTweet(); err != nil {
		ctx.Echo().Logger.Error(err)
		return echo.NewHTTPError(500, "failed to update needs-tweet")
	}
	return ctx.String(200, `ok`)
}
