package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"ys-back/mocks"
	"ys-back/src/interfaces/validation"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNeedsHandler_UpdateNeedsTweet(tt *testing.T) {
	tt.Run(
		"正常系: エラーなし",
		func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/needs/update", nil)

			e := echo.New()
			e.Validator = &validation.CustomValidator{V: validator.New()}
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			srv := new(mocks.ITwitterApiService)
			srv.On("FetchNeedsTweet").Return(nil)

			hdr := NewNeedsHandler(srv, e.Logger)
			err := hdr.UpdateNeedsTweet(ctx)

			require.NoError(t, err)
			assert.Equal(t, 200, rec.Code)
			srv.AssertExpectations(t)
		})
}
