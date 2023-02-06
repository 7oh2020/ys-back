package model

// ツイートを分類するためのタグ情報
type Tag struct {
	// タグid
	id uint

	// ツイートから検索する文字列リスト
	names []string
}

func newTag(id uint, names []string) *Tag {
	return &Tag{
		id:    id,
		names: names,
	}
}

func (t *Tag) ID() uint {
	return t.id
}

func (t *Tag) Names() []string {
	return t.names
}
