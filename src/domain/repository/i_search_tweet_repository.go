package repository

import "ys-back/src/infrastructure/persistence/model/remote"

// Twitter APIを使用してツイート検索を行う
type ISearchTweetRepository interface {
	// 指定されたパラメータに対応した検索結果を取得する
	Search(keyword string, count uint, sinceID string) (*remote.SearchTweetResult, error)
}
