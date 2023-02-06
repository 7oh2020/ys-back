package remote

// 検索結果の追加情報のAPIモデル
type TweetInclude struct {
	Users []*TweetUser `json:"users"`
}
