package store

import (
	"testing"
	"ys-back/src/infrastructure/persistence/model/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTweetRepository_FetchTweets(tt *testing.T) {

	tt.Run(
		"正常系: 全ツイートが対象",
		func(t *testing.T) {
			items := []*db.Tweet{
				db.NewTweet("1", "c1", "date", 0, "alice", "アリス", "avatar"),
				db.NewTweet("2", "c2", "date", 0, "alice", "アリス", "avatar"),
			}
			tweet := db.NewTweetRoot(items)
			count := uint(2)
			offset := uint(0)

			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`SELECT id, content, created_at, user_name, name, avatar_url, tag_id FROM tweets`).
				ExpectQuery().
				WithArgs(count, offset).
				WillReturnRows(sqlmock.NewRows([]string{"id", "content", "created_at", "user_name", "name", "avatar_url", "tag_id"}).AddRow(tweet.Items[0].ID, tweet.Items[0].Content, tweet.Items[0].CreatedAt, tweet.Items[0].UserName, tweet.Items[0].Name, tweet.Items[0].AvatarURL, tweet.Items[0].TagID).AddRow(tweet.Items[1].ID, tweet.Items[1].Content, tweet.Items[1].CreatedAt, tweet.Items[1].UserName, tweet.Items[1].Name, tweet.Items[1].AvatarURL, tweet.Items[1].TagID))

			e := echo.New()
			repo := NewTweetRepository(db, e.Logger)
			ret, err := repo.FetchTweets(count, offset)

			require.NoError(t, err)
			assert.ElementsMatch(t, items, ret.Items)
			require.NoError(t, mock.ExpectationsWereMet())
		})
}

func TestTweetRepository_FetchTweetsWithTagID(tt *testing.T) {
	tt.Run(
		"正常系: タグIDに一致するツイートが対象",
		func(t *testing.T) {
			items := []*db.Tweet{
				db.NewTweet("1", "c1", "date", 0, "alice", "アリス", "avatar"),
				db.NewTweet("2", "c2", "date", 0, "alice", "アリス", "avatar"),
			}
			tweet := db.NewTweetRoot(items)
			tagID := uint(1)
			count := uint(2)
			offset := uint(0)

			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`SELECT id, content, created_at, user_name, name, avatar_url, tag_id FROM tweets`).
				ExpectQuery().
				WithArgs(tagID, count, offset).
				WillReturnRows(sqlmock.NewRows([]string{"id", "content", "created_at", "user_name", "name", "avatar_url", "tag_id"}).AddRow(tweet.Items[0].ID, tweet.Items[0].Content, tweet.Items[0].CreatedAt, tweet.Items[0].UserName, tweet.Items[0].Name, tweet.Items[0].AvatarURL, tweet.Items[0].TagID).AddRow(tweet.Items[1].ID, tweet.Items[1].Content, tweet.Items[1].CreatedAt, tweet.Items[1].UserName, tweet.Items[1].Name, tweet.Items[1].AvatarURL, tweet.Items[1].TagID))

			e := echo.New()
			repo := NewTweetRepository(db, e.Logger)
			ret, err := repo.FetchTweetsWithTagID(tagID, count, offset)

			require.NoError(t, err)
			assert.ElementsMatch(t, items, ret.Items)
			require.NoError(t, mock.ExpectationsWereMet())
		})
}

func TestTweetRepository_FetchLatestID(tt *testing.T) {
	tt.Run(
		"正常系: 最新ツイートのIDが取得できること",
		func(t *testing.T) {
			id := "1"

			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`SELECT id FROM tweets`).
				ExpectQuery().
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

			e := echo.New()
			repo := NewTweetRepository(db, e.Logger)
			ret, err := repo.FetchLatestID()

			require.NoError(t, err)
			assert.Equal(t, id, ret)
			require.NoError(t, mock.ExpectationsWereMet())
		})
}

func TestTweetRepository_StoreTweets(tt *testing.T) {
	tt.Run(
		"正常系: ツイートが保存できること",
		func(t *testing.T) {
			items := []*db.Tweet{
				db.NewTweet("1", "c1", "date", 0, "alice", "アリス", "avatar"),
			}
			tweet := db.NewTweetRoot(items)

			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`INSERT INTO tweets`).
				ExpectExec().
				WithArgs(tweet.Items[0].ID, tweet.Items[0].Content, tweet.Items[0].CreatedAt, tweet.Items[0].UserName, tweet.Items[0].Name, tweet.Items[0].AvatarURL, tweet.Items[0].TagID).
				WillReturnResult(sqlmock.NewResult(1, 1))

			e := echo.New()
			repo := NewTweetRepository(db, e.Logger)
			ret := repo.StoreTweets(tweet)

			require.NoError(t, ret)
			require.NoError(t, mock.ExpectationsWereMet())
		})
}

func TestTweetRepository_PurgeTweets(tt *testing.T) {
	tt.Run(
		"正常系: ツイートが削除できること",
		func(t *testing.T) {
			cnt := uint(10)

			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(`DELETE FROM tweets`).
				ExpectExec().
				WithArgs(cnt).
				WillReturnResult(sqlmock.NewResult(int64(cnt), 1))

			e := echo.New()
			repo := NewTweetRepository(db, e.Logger)
			ret := repo.PurgeTweets(cnt)

			require.NoError(t, ret)
			require.NoError(t, mock.ExpectationsWereMet())
		})
}
