package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTagRoot_GetIDFromTweet(tt *testing.T) {
	tt.Run(
		"正常系: 文字列にタグが含まれる場合",
		func(t *testing.T) {
			s := "これはアプリに一致します。イベントも含まれてます。"

			tag := NewTagRoot()
			ret := tag.GetIDFromTweet(s)

			assert.Equal(t, uint(1), ret)
		})
	tt.Run(
		"正常系: 文字列にタグが含まれない場合",
		func(t *testing.T) {
			s := "どれにも一致しない。"

			tag := NewTagRoot()
			ret := tag.GetIDFromTweet(s)

			assert.Equal(t, uint(0), ret)
		})
}

func TestTagRoot_IsTargetedAllTweets(tt *testing.T) {
	tt.Run(
		"正常系: タグIDが0の場合",
		func(t *testing.T) {
			tagID := uint(0)

			tag := NewTagRoot()
			ret := tag.IsTargetedAllTweets(tagID)

			require.Equal(t, true, ret)
		})
	tt.Run(
		"正常系: タグIDが0以外の場合",
		func(t *testing.T) {
			tagID := uint(1)

			tag := NewTagRoot()
			ret := tag.IsTargetedAllTweets(tagID)

			require.Equal(t, false, ret)
		})
}
