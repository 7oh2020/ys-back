package remote

// ツイートの作成ユーザーのAPIモデル
type TweetUser struct {
	// 数値ID
	ID string `json:"id"`

	// スクリーンネーム
	UserName string `json:"username"`

	// 名前
	Name string `json:"name"`

	// アイコン画像URL
	AvatarURL string `json:"profile_image_url"`
}
