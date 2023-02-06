package dto

// ツイート情報のDTO
type TweetData struct {
	// ツイートID
	ID string `json:"id"`

	// ツイート本文
	Content string `json:"text"`

	// ツイートの投稿日時
	CreatedAt string `json:"created_at"`

	// スクリーンネーム
	UserName string `json:"user_name"`

	// ユーザー名
	Name string `json:"name"`

	// アイコン画像URL
	AvatarURL string `json:"avatar_url"`

	// タグID
	TagID uint `json:"-"`
}

func NewTweetData(id string, content string, createdat string, tagid uint, username string, name string, avatar string) *TweetData {
	return &TweetData{
		ID:        id,
		Content:   content,
		CreatedAt: createdat,
		TagID:     tagid,
		UserName:  username,
		Name:      name,
		AvatarURL: avatar,
	}
}
