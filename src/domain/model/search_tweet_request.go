package model

// ツイート検索のリクエストパラメータ
type SearchTweetRequest struct {
	// 検索キーワード
	keyword string

	//取得件数
	count uint

	// 最後に取得したツイートのID
	sinceID string
}

func NewSearchTweetRequest(sinceID string) *SearchTweetRequest {
	return &SearchTweetRequest{
		// 要望ツイートを検索するためのキーワードを作成する。RTや無関係なツイートは極力除外する
		keyword: `(アプリ OR 機能 OR サービス OR ツール OR ソフト OR イベント OR サイト OR ゲーム OR システム) (あればな OR あればなぁ OR あればいいのに OR あればいいのにな OR あればいいのになぁ OR あったらいいのに OR あったらいいのになぁ OR あったらなぁ OR あったらいいな OR あったらいいなぁ) -"定期" -"ビジネス" -"副業" -is:retweet lang:ja`,

		count:   100,
		sinceID: sinceID,
	}
}

func (req SearchTweetRequest) Keyword() string {
	return req.keyword
}
func (req SearchTweetRequest) Count() uint {
	return req.count
}
func (req SearchTweetRequest) SinceID() string {
	return req.sinceID
}
