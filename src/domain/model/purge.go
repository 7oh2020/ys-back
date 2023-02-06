package model

// ツイートの削除情報
type PurgeInfo struct {
	// データベースに保存できる最大ツイート件数
	maxCount uint
}

func NewPurgeInfo() *PurgeInfo {
	return &PurgeInfo{
		maxCount: 10000,
	}
}

func (p PurgeInfo) MaxCount() uint {
	return p.maxCount
}

// 最大件数を超えているかどうか
func (p PurgeInfo) IsOverflow(total uint) bool {
	return total > p.maxCount
}
