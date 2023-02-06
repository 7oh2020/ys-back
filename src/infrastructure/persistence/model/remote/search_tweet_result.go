package remote

// ツイート検索結果のAPIモデル
type SearchTweetResult struct {
	Items   []*Tweet      `json:"data"`
	Include *TweetInclude `json:"includes"`
}

func NewSearchTweetResult(items []*Tweet) *SearchTweetResult {
	return &SearchTweetResult{
		Items: items}
}

// includesのユーザー情報とauthorIDをバインドする
func (tweets *SearchTweetResult) BindUserName() {
	for _, t := range tweets.Items {
		for _, u := range tweets.Include.Users {
			if t.AuthorID == u.ID {
				t.UserName = u.UserName
				t.Name = u.Name
				t.AvatarURL = u.AvatarURL
				break
			}
		}
	}
}
