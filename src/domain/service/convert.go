package service

import (
	"ys-back/src/infrastructure/persistence/model/db"
	"ys-back/src/infrastructure/persistence/model/remote"
	"ys-back/src/interfaces/dto"
)

// Twitter APIに特化したモデルからDBに特化したモデルへ変換する
func remoteTweetToDB(tweet *remote.SearchTweetResult) *db.TweetRoot {
	items := make([]*db.Tweet, len(tweet.Items))
	for i, v := range tweet.Items {
		items[i] = db.NewTweet(v.ID, v.Content, v.CreatedAt, v.TagID, v.UserName, v.Name, v.AvatarURL)
	}
	return db.NewTweetRoot(items)
}

// DBに特化したモデルからアプリケーションに特化したDTOへ変換する
func dbTweetToDTO(tweet *db.TweetRoot) *dto.TweetRootData {
	items := make([]*dto.TweetData, len(tweet.Items))
	for i, v := range tweet.Items {
		items[i] = dto.NewTweetData(v.ID, v.Content, v.CreatedAt, v.TagID, v.UserName, v.Name, v.AvatarURL)
	}
	return dto.NewTweetRootData(items)
}
