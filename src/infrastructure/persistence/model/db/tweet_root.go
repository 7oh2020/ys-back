package db

// ツイートリストを持つDBモデル
type TweetRoot struct {
	Items []*Tweet
}

func NewTweetRoot(items []*Tweet) *TweetRoot {
	return &TweetRoot{items}
}
