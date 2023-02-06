package main

import (
	"fmt"
	"os"
	"ys-back/src/infrastructure/persistence/twitter"

	"github.com/labstack/echo/v4"
)

// Twitter APIとの疎通確認を行う
func main() {
	e := echo.New()
	repo := twitter.NewSearchTweetRepository(os.Getenv("TWITTER_BEARER"), e.Logger)
	result, err := repo.Search(`Go言語 lang:ja`, 10, "")
	if err != nil {
		panic(err)
	}
	for _, v := range result.Items {
		fmt.Println(v.Name, v.UserName)
		fmt.Println(v.Content)
		fmt.Println("---")
	}
}
