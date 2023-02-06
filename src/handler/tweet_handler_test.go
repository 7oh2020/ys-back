package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"ys-back/mocks"
	"ys-back/src/domain/model"
	"ys-back/src/interfaces/dto"
	"ys-back/src/interfaces/validation"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTweetHandler_GetTweets(tt *testing.T) {
	tt.Run(
		"正常系: 1ページ目",
		func(t *testing.T) {
			tagID := uint(1)
			page := uint(1)
			url := fmt.Sprintf("/tweet/index?tag_id=%d&page=%d", tagID, page)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			tweet := dto.NewTweetRootData([]*dto.TweetData{
				dto.NewTweetData("1", "c1", "date", 1, "alice", "Alice", "avatar"),
				dto.NewTweetData("2", "c2", "date", 2, "alice", "Alice", "avatar"),
			})

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			srv := new(mocks.ITweetManageService)
			srv.On("GetTweets", tagID, page).Return(tweet, nil)

			hdr := NewTweetHandler(srv, e.Logger)
			err := hdr.GetTweets(ctx)

			require.NoError(t, err)
			assert.Equal(t, 200, rec.Code)
			srv.AssertExpectations(t)
		})
	tt.Run(
		"準正常系: 0ページ目は不正",
		func(t *testing.T) {
			tagID := uint(1)
			page := uint(0)
			url := fmt.Sprintf("/tweet/index?tag_id=%d&page=%d", tagID, page)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			srv := new(mocks.ITweetManageService)

			hdr := NewTweetHandler(srv, e.Logger)
			err := hdr.GetTweets(ctx)

			require.EqualError(t, err, "code=400, message=failed to validate request")
			assert.Equal(t, 200, rec.Code)
		})

}

func TestTweetHandler_GetMeta(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/tweet/meta", nil)
			count := model.NewPageInfo().CountPerPage()
			meta := dto.NewMetaData(count, []*dto.MetaTagData{
				dto.NewMetaTagData(0, 100),
				dto.NewMetaTagData(1, 10),
				dto.NewMetaTagData(2, 20),
				dto.NewMetaTagData(3, 30),
				dto.NewMetaTagData(4, 40),
			})

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			srv := new(mocks.ITweetManageService)
			srv.On("GetMeta").Return(meta, nil)

			hdr := NewTweetHandler(srv, e.Logger)
			err := hdr.GetMeta(ctx)

			require.NoError(t, err)
			assert.Equal(t, 200, rec.Code)
			srv.AssertExpectations(t)
		})
}
