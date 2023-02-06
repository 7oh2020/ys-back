package service

import (
	"testing"
	"ys-back/mocks"
	"ys-back/src/domain/model"
	"ys-back/src/infrastructure/persistence/model/remote"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestTwitterApiService_GetNeedsTweet(tt *testing.T) {
	tt.Run(
		"正常系: 最大保存件数を超えない場合",
		func(t *testing.T) {
			rt := remote.NewSearchTweetResult([]*remote.Tweet{
				remote.NewTweetData("1", "c1", "date", 0, "alice", "アリス", "avatar"),
				remote.NewTweetData("2", "c2", "date", 0, "alice", "アリス", "avatar"),
			})
			mt := remoteTweetToDB(rt)
			sinceID := "1"
			params := model.NewSearchTweetRequest(sinceID)
			total := uint(10)

			store := new(mocks.ITweetRepository)
			store.On("FetchLatestID").Return(sinceID, nil)
			store.On("FetchCount").Return(total, nil)
			store.On("StoreTweets", mt).Return(nil)

			twitter := new(mocks.ISearchTweetRepository)
			twitter.On("Search", params.Keyword(), params.Count(), params.SinceID()).Return(rt, nil)

			e := echo.New()
			srv := NewTwitterApiService(twitter, store, e.Logger)
			ret := srv.FetchNeedsTweet()

			require.NoError(t, ret)

		})
	tt.Run(
		"正常系: 最大保存件数を超える場合",
		func(t *testing.T) {
			rt := remote.NewSearchTweetResult([]*remote.Tweet{
				remote.NewTweetData("1", "c1", "date", 0, "alice", "アリス", "avatar"),
				remote.NewTweetData("2", "c2", "date", 0, "alice", "アリス", "avatar"),
			})
			mt := remoteTweetToDB(rt)
			sinceID := "1"
			params := model.NewSearchTweetRequest(sinceID)
			purge := model.NewPurgeInfo()
			total := uint(purge.MaxCount())

			store := new(mocks.ITweetRepository)
			store.On("FetchLatestID").Return(sinceID, nil)
			store.On("FetchCount").Return(total, nil)
			store.On("PurgeTweets", uint(len(mt.Items))).Return(nil)
			store.On("StoreTweets", mt).Return(nil)

			twitter := new(mocks.ISearchTweetRepository)
			twitter.On("Search", params.Keyword(), params.Count(), params.SinceID()).Return(rt, nil)

			e := echo.New()
			srv := NewTwitterApiService(twitter, store, e.Logger)
			ret := srv.FetchNeedsTweet()

			require.NoError(t, ret)
			store.AssertExpectations(t)
			twitter.AssertExpectations(t)
		})

}
