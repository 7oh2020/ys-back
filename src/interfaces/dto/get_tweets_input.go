package dto

// ツイート取得のリクエストパラメータ
type GetTweetsInput struct {
	// 現在のページ
	Page int `query:"page" validate:"required,numeric,min=1,max=9999"`

	// タグID
	TagID int `query:"tag_id" validate:"numeric,min=0,max=4"`
}
