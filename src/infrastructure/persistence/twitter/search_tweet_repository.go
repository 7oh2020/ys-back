package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"ys-back/src/domain/repository"
	"ys-back/src/infrastructure/persistence/model/remote"

	"github.com/labstack/echo/v4"
)

type SearchTweetRepository struct {
	// APIコールに必要なBearer Token
	bearer string

	logger echo.Logger
}

func NewSearchTweetRepository(bearer string, logger echo.Logger) repository.ISearchTweetRepository {
	return &SearchTweetRepository{bearer, logger}
}

func (repo *SearchTweetRepository) Search(keyword string, count uint, sinceID string) (*remote.SearchTweetResult, error) {
	body, err := repo.fetchFromApi(keyword, count, sinceID)
	if err != nil {
		return nil, err
	}

	meta, err := repo.unmarshalTweetResultMeta(body)
	if err != nil {
		return nil, err
	}

	result, err := repo.unmarshalTweetResult(body, uint(meta.Meta.Count))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Twitter API v2を使用してツイートを検索する
func (repo *SearchTweetRepository) fetchFromApi(keyword string, count uint, sinceID string) ([]byte, error) {
	// 検索キーワードはエスケープする必要がある
	keyword = url.QueryEscape(keyword)
	target := fmt.Sprintf("https://api.twitter.com/2/tweets/search/recent?tweet.fields=created_at,entities&user.fields=profile_image_url&expansions=author_id&max_results=%d&query=%s", count, keyword)

	// since_idを指定するとそのツイートIDより新しいツイートが取得対象になる
	if sinceID != "" {
		target += fmt.Sprintf("&since_id=%s", sinceID)
	}
	repo.logger.Debug("target: ", target)

	client := &http.Client{
		// リソース節約のためにクライアントにタイムアウトを設定する
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		repo.logger.Error(err)
		return nil, err
	}
	// 認証のためにBearer Tokenをセットする
	req.Header.Set("Authorization", "Bearer "+repo.bearer)

	resp, err := client.Do(req)
	if err != nil {
		repo.logger.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		repo.logger.Error(err)
		return nil, err
	}

	repo.logger.Debug("response code: ", resp.StatusCode)
	repo.logger.Debug("response body length: ", len(body))

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to get response from twitter-api")
	}
	return body, nil
}

// 検索結果のメタ情報をパースする
func (repo *SearchTweetRepository) unmarshalTweetResultMeta(body []byte) (*remote.SearchTweetResultOnlyMeta, error) {
	meta := &remote.SearchTweetResultOnlyMeta{}
	if err := json.Unmarshal(body, &meta); err != nil {
		repo.logger.Error(err)
		return nil, err
	}
	repo.logger.Debug("result count: ", meta.Meta.Count)
	return meta, nil
}

// 検索結果をパースする
func (repo *SearchTweetRepository) unmarshalTweetResult(body []byte, count uint) (*remote.SearchTweetResult, error) {
	// 必要な件数分だけスライスを作成する
	items := make([]*remote.Tweet, count)
	result := remote.NewSearchTweetResult(items)

	if err := json.Unmarshal(body, result); err != nil {
		repo.logger.Error(err)
		return nil, err
	}

	// includesのユーザー情報とauthorIDをバインドする
	result.BindUserName()

	return result, nil
}
