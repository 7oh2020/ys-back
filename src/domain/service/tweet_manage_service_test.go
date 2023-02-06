package service

import (
	"testing"
	"ys-back/mocks"
	"ys-back/src/domain/model"
	"ys-back/src/infrastructure/persistence/model/db"
	"ys-back/src/interfaces/dto"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTweetManageService_GetTweets(tt *testing.T) {
	tt.Run(
		"正常系: タグIDが0の場合は全ツイートが対象",
		func(t *testing.T) {
			tweet := db.NewTweetRoot([]*db.Tweet{
				db.NewTweet("1", "c1", "date", 0, "", "アリス", "avatar"),
				db.NewTweet("2", "c2", "date", 0, "", "アリス", "avatar"),
			})
			dt := dbTweetToDTO(tweet)
			tagID := uint(0)
			page := uint(1)
			count := model.NewPageInfo().CountPerPage()
			offset := uint(0)

			repo := new(mocks.ITweetRepository)
			repo.On("FetchTweets", count, offset).Return(tweet, nil)

			e := echo.New()
			srv := NewTweetManageService(repo, e.Logger)
			ret, err := srv.GetTweets(tagID, page)

			require.NoError(t, err)
			assert.ElementsMatch(t, dt.Items, ret.Items)
			repo.AssertExpectations(t)
		})
	tt.Run(
		"正常系: タグIDが0以外の場合はタグIDに一致するツイートが対象",
		func(t *testing.T) {
			tweet := db.NewTweetRoot([]*db.Tweet{
				db.NewTweet("1", "c1", "date", 0, "", "アリス", "avatar"),
				db.NewTweet("2", "c2", "date", 0, "", "アリス", "avatar"),
			})
			dt := dbTweetToDTO(tweet)
			tagID := uint(1)
			page := uint(1)
			count := model.NewPageInfo().CountPerPage()
			offset := uint(0)

			repo := new(mocks.ITweetRepository)
			repo.On("FetchTweetsWithTagID", tagID, count, offset).Return(tweet, nil)

			e := echo.New()
			srv := NewTweetManageService(repo, e.Logger)
			ret, err := srv.GetTweets(tagID, page)

			require.NoError(t, err)
			assert.ElementsMatch(t, dt.Items, ret.Items)
			repo.AssertExpectations(t)
		})

}

func TestTweetManageService_GetMeta(tt *testing.T) {
	tt.Run(
		"正常系: メタ情報が取得できること",
		func(t *testing.T) {
			total := uint(10)
			count := model.NewPageInfo().CountPerPage()
			meta := dto.NewMetaData(count, []*dto.MetaTagData{
				dto.NewMetaTagData(0, total),
				dto.NewMetaTagData(1, 10),
				dto.NewMetaTagData(2, 20),
				dto.NewMetaTagData(3, 30),
				dto.NewMetaTagData(4, 40),
			})

			repo := new(mocks.ITweetRepository)
			repo.On("FetchCount").Return(total, nil)
			repo.On("FetchCountWithTagID", meta.Tags[1].TagID).Return(meta.Tags[1].Count, nil)
			repo.On("FetchCountWithTagID", meta.Tags[2].TagID).Return(meta.Tags[2].Count, nil)
			repo.On("FetchCountWithTagID", meta.Tags[3].TagID).Return(meta.Tags[3].Count, nil)
			repo.On("FetchCountWithTagID", meta.Tags[4].TagID).Return(meta.Tags[4].Count, nil)

			e := echo.New()
			srv := NewTweetManageService(repo, e.Logger)
			ret, err := srv.GetMeta()

			require.NoError(t, err)
			assert.Equal(t, meta.CountPerPage, ret.CountPerPage)
			assert.ElementsMatch(t, meta.Tags, ret.Tags)
			repo.AssertExpectations(t)
		})
}
