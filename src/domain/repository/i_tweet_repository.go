package repository

import "ys-back/src/infrastructure/persistence/model/db"

// ツイート情報の永続化を行う
type ITweetRepository interface {
	// ツイートを取得する
	FetchTweets(count uint, offset uint) (*db.TweetRoot, error)

	// タグIDに一致するツイートを取得する
	FetchTweetsWithTagID(tagID uint, count uint, offset uint) (*db.TweetRoot, error)

	// 最後に保存したツイートのIDを取得する
	FetchLatestID() (string, error)

	// ツイートの総件数を取得する
	FetchCount() (uint, error)

	// タグIDに一致するツイートの件数を取得する
	FetchCountWithTagID(tagID uint) (uint, error)

	// ツイートを保存する
	StoreTweets(tweet *db.TweetRoot) error

	// 古いツイートを削除する
	PurgeTweets(count uint) error
}
