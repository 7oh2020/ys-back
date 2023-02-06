package service

import (
	"ys-back/src/domain/model"
	"ys-back/src/domain/repository"

	"github.com/labstack/echo/v4"
)

// Twitter APIとの通信を行うサービス
type ITwitterApiService interface {
	// Twitter APIを使用して要望ツイートを取得してデータベースに保存する
	FetchNeedsTweet() error
}

type TwitterApiService struct {
	twitter repository.ISearchTweetRepository
	store   repository.ITweetRepository
	logger  echo.Logger
}

func NewTwitterApiService(twitter repository.ISearchTweetRepository, store repository.ITweetRepository, logger echo.Logger) ITwitterApiService {
	return &TwitterApiService{twitter, store, logger}
}

func (srv *TwitterApiService) FetchNeedsTweet() error {
	// 最後に保存したツイートのIDを取得する
	latestID, _ := srv.store.FetchLatestID()
	params := model.NewSearchTweetRequest(latestID)

	// 要望ツイートを取得する
	result, err := srv.twitter.Search(params.Keyword(), params.Count(), params.SinceID())
	if err != nil {
		return err
	}

	// ツイート内容に対応したタグIDを付加する
	tag := model.NewTagRoot()
	for _, v := range result.Items {
		v.TagID = tag.GetIDFromTweet(v.Content)
	}

	// 全ツイートの件数を取得する
	total, err := srv.store.FetchCount()
	if err != nil {
		return err
	}

	// 今回の追加で最大保存件数を超える場合は古い順に削除する
	purge := model.NewPurgeInfo()
	if purge.IsOverflow(total + uint(len(result.Items))) {
		srv.logger.Debug("tweet is overflow")
		if err := srv.store.PurgeTweets(uint(len(result.Items))); err != nil {
			return err
		}
	}

	// ツイートを保存する
	tweet := remoteTweetToDB(result)
	if err := srv.store.StoreTweets(tweet); err != nil {
		return err
	}
	return nil
}
