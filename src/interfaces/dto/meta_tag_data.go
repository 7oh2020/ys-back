package dto

// タグ毎のメタ情報
type MetaTagData struct {
	// タグID
	TagID uint `json:"tag_id"`

	// タグIDに一致するツイート件数
	Count uint `json:"count"`
}

func NewMetaTagData(tagID uint, count uint) *MetaTagData {
	return &MetaTagData{
		TagID: tagID,
		Count: count,
	}
}
