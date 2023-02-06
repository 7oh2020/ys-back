package model

import "strings"

// タグリストを持つモデル
type TagRoot struct {
	items []*Tag
}

func NewTagRoot() *TagRoot {
	return &TagRoot{
		items: []*Tag{
			// 1. アプリ・ソフトウェア
			newTag(1, []string{"アプリ", "ソフト", "ツール", "システム", "機能"}),

			// 2. サービス
			newTag(2, []string{"サービス"}),

			// 3. ゲーム
			newTag(3, []string{"ゲーム"}),

			// 4. イベント
			newTag(4, []string{"イベント"}),
		},
	}
}

func (t *TagRoot) Items() []*Tag {
	return t.items
}

// 文字列に対応するタグIDを返す
// どれにも一致しない場合は0を返す
func (t TagRoot) GetIDFromTweet(body string) uint {
	for _, item := range t.items {
		for _, name := range item.Names() {
			if strings.Contains(body, name) {
				return item.ID()
			}
		}
	}
	return 0
}

// 全てのツイートが取得対象かどうか
func (t TagRoot) IsTargetedAllTweets(tagID uint) bool {
	return tagID == 0
}
