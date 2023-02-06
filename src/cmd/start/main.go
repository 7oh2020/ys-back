package main

import (
	"database/sql"
	"os"
	"ys-back/src/interfaces/di"
	"ys-back/src/interfaces/validation"

	"github.com/go-playground/validator"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// RESTサーバーの起動を行う
func main() {
	// ポート番号の取得
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	// Bearer Tokenの取得
	var bearer string
	if bearer = os.Getenv("TWITTER_BEARER"); bearer == "" {
		panic("bearer-token is empty")
	}

	// データベースURLの取得
	var dbURL string
	if dbURL = os.Getenv("DATABASE_URL"); dbURL == "" {
		panic("database-url is empty")
	}

	// postgresに接続
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Echoインスタンスの作成
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Validator = &validation.CustomValidator{V: validator.New()}
	e.Use(middleware.Logger())
	if mode := os.Getenv("APP_MODE"); mode == "DEBUG" {
		e.Logger.SetLevel(log.DEBUG)
	}

	// ルートとハンドラーを関連付ける
	SetNeedsRoutes(e, bearer, db)
	SetTweetRoutes(e, db)

	// サーバーを起動する
	e.Logger.Fatal(e.Start(":" + port))
}

func SetNeedsRoutes(e *echo.Echo, bearer string, db *sql.DB) {
	needs := di.InitNeeds(bearer, db, e.Logger)
	e.GET("/needs/update", needs.UpdateNeedsTweet)
}

func SetTweetRoutes(e *echo.Echo, db *sql.DB) {
	tweet := di.InitTweet(db, e.Logger)
	e.GET("/tweet/index", tweet.GetTweets)
	e.GET("/tweet/meta", tweet.GetMeta)
}
