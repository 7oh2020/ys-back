package service

import (
	"ys-back/src/domain/model"
	"ys-back/src/domain/repository"
	"ys-back/src/infrastructure/persistence/model/db"
	"ys-back/src/interfaces/dto"

	"github.com/labstack/echo/v4"
)

// 保存されているツイートの管理を行う
type ITweetManageService interface {
	// タグIDに一致するツイートを取得する
	GetTweets(tagID uint, page uint) (*dto.TweetRootData, error)

	// ページングのためのメタ情報を取得する
	GetMeta() (*dto.MetaData, error)
}

type TweetManageService struct {
	store  repository.ITweetRepository
	logger echo.Logger
}

func NewTweetManageService(store repository.ITweetRepository, logger echo.Logger) ITweetManageService {
	return &TweetManageService{store, logger}

}

func (srv *TweetManageService) GetTweets(tagID uint, page uint) (*dto.TweetRootData, error) {
	// ページからoffsetを計算する
	p := model.NewPageInfo()
	offset := (page - 1) * p.CountPerPage()

	var result *db.TweetRoot
	var err error

	t := model.NewTagRoot()
	if t.IsTargetedAllTweets(tagID) {
		// 全ツイートが取得対象
		result, err = srv.store.FetchTweets(p.CountPerPage(), offset)
	} else {
		// タグIDに一致するツイートが取得対象
		result, err = srv.store.FetchTweetsWithTagID(tagID, p.CountPerPage(), offset)
	}
	if err != nil {
		return nil, err
	}
	return dbTweetToDTO(result), nil
}

func (srv *TweetManageService) GetMeta() (*dto.MetaData, error) {
	t := model.NewTagRoot()
	tags := make([]*dto.MetaTagData, len(t.Items())+1)

	// 全ツイートの件数を取得する
	total, err := srv.store.FetchCount()
	if err != nil {
		return nil, err
	}
	tags[0] = dto.NewMetaTagData(0, total)

	// タグ毎の件数を取得する
	for i, v := range t.Items() {
		cnt, err := srv.store.FetchCountWithTagID(v.ID())
		if err != nil {
			return nil, err
		}
		tags[i+1] = dto.NewMetaTagData(v.ID(), cnt)
	}

	p := model.NewPageInfo()
	meta := dto.NewMetaData(p.CountPerPage(), tags)

	return meta, nil
}
