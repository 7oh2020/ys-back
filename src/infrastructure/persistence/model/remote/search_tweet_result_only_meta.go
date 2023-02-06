package remote

// ツイート検索結果のメタ情報のみのAPIモデル
type SearchTweetResultOnlyMeta struct {
	Meta TweetMeta `json:"meta"`
}
