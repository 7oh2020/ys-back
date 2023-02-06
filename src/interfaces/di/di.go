package di

import (
	"database/sql"
	"ys-back/src/domain/service"
	"ys-back/src/handler"
	"ys-back/src/infrastructure/persistence/store"
	"ys-back/src/infrastructure/persistence/twitter"

	"github.com/labstack/echo/v4"
)

func InitNeeds(bearer string, db *sql.DB, logger echo.Logger) handler.INeedsHandler {
	tr := twitter.NewSearchTweetRepository(bearer, logger)
	sr := store.NewTweetRepository(db, logger)
	srv := service.NewTwitterApiService(tr, sr, logger)
	return handler.NewNeedsHandler(srv, logger)
}

func InitTweet(db *sql.DB, logger echo.Logger) handler.ITweetHandler {
	repo := store.NewTweetRepository(db, logger)
	srv := service.NewTweetManageService(repo, logger)
	return handler.NewTweetHandler(srv, logger)
}
