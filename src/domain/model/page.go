package model

// ページングのためのページ情報
type PageInfo struct {
	// 1ページあたりの表示件数
	countPerPage uint
}

func NewPageInfo() *PageInfo {
	return &PageInfo{
		countPerPage: 15,
	}
}

func (data PageInfo) CountPerPage() uint {
	return data.countPerPage
}
