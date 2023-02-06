package remote

// ツイートのAPIモデル
type Tweet struct {
	// ツイートID
	ID string `json:"id"`

	// ツイート本文
	Content string `json:"text"`

	// ツイートの投稿日時
	CreatedAt string `json:"created_at"`

	// ユーザーID
	AuthorID string `json:"author_id"`

	// スクリーンネーム
	UserName string `json:"-"`

	// ユーザー名
	Name string `json:"-"`

	// アイコン画像URL
	AvatarURL string `json:"-"`

	// タグID
	TagID uint `json:"-"`
}

func NewTweetData(id string, content string, createdat string, tagid uint, username string, name string, avatar string) *Tweet {
	return &Tweet{
		ID:        id,
		Content:   content,
		CreatedAt: createdat,
		TagID:     tagid,
		UserName:  username,
		Name:      name,
		AvatarURL: avatar,
	}
}
