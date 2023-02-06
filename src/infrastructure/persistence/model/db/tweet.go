package db

// ツイート情報のDBモデル
type Tweet struct {
	ID        string
	Content   string
	CreatedAt string
	TagID     uint
	UserName  string
	Name      string
	AvatarURL string
}

func NewTweet(id string, content string, createdAt string, tagID uint, userName string, name string, avatarURL string) *Tweet {
	return &Tweet{
		ID:        id,
		Content:   content,
		CreatedAt: createdAt,
		TagID:     tagID,
		UserName:  userName,
		Name:      name,
		AvatarURL: avatarURL,
	}
}
