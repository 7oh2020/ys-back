package remote

// ツイート検索結果のメタ情報のAPIモデル
type TweetMeta struct {
	// 検索結果のツイート件数
	Count uint `json:"result_count"`
}
