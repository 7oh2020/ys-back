package dto

// Static Generateのためのメタ情報
type MetaData struct {
	// 1ページあたりの表示件数
	CountPerPage uint `json:"count_per_page"`

	// タグ毎のメタ情報
	Tags []*MetaTagData `json:"tags"`
}

func NewMetaData(count uint, tags []*MetaTagData) *MetaData {
	return &MetaData{
		CountPerPage: count,
		Tags:         tags}
}
