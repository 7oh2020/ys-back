package dto

// ツイート一覧を持つDTO
type TweetRootData struct {
	Items []*TweetData `json:"items"`
}

func NewTweetRootData(items []*TweetData) *TweetRootData {
	return &TweetRootData{
		Items: items,
	}
}
