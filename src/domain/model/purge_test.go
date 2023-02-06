package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPurgeInfo_IsOverflow(tt *testing.T) {
	tt.Run(
		"正常系: 最大件数を超える場合",
		func(t *testing.T) {
			purge := NewPurgeInfo()
			count := uint(purge.maxCount + 1)

			ret := purge.IsOverflow(count)

			assert.Equal(t, true, ret)
		})
	tt.Run(
		"正常系: 最大件数を超えない場合",
		func(t *testing.T) {
			purge := NewPurgeInfo()
			count := uint(1)

			ret := purge.IsOverflow(count)

			assert.Equal(t, false, ret)
		})
}
