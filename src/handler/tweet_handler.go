package handler

import (
	"ys-back/src/domain/service"
	"ys-back/src/interfaces/dto"

	"github.com/labstack/echo/v4"
)

type ITweetHandler interface {
	// 指定したタグIDに一致するツイートを取得する
	GetTweets(ctx echo.Context) error

	// ツイート件数などのメタ情報を取得する
	GetMeta(ctx echo.Context) error
}

type TweetHandler struct {
	srv    service.ITweetManageService
	logger echo.Logger
}

func NewTweetHandler(srv service.ITweetManageService, logger echo.Logger) ITweetHandler {
	return &TweetHandler{srv, logger}
}

func (hdr *TweetHandler) GetTweets(ctx echo.Context) error {
	// リクエストパラメータをチェックする
	input := new(dto.GetTweetsInput)
	if err := ctx.Bind(input); err != nil {
		return echo.NewHTTPError(400, "failed to bind request")
	}
	if err := ctx.Validate(input); err != nil {
		return echo.NewHTTPError(400, "failed to validate request")
	}

	result, err := hdr.srv.GetTweets(uint(input.TagID), uint(input.Page))
	if err != nil {
		return echo.NewHTTPError(404, "no result")
	}
	return ctx.JSON(200, result)
}

func (hdr *TweetHandler) GetMeta(ctx echo.Context) error {
	result, err := hdr.srv.GetMeta()
	if err != nil {
		ctx.Echo().Logger.Error(err)
		return echo.NewHTTPError(500, "failed to get tweet-meta")
	}
	return ctx.JSON(200, result)
}
