package store

import (
	"database/sql"
	"ys-back/src/domain/repository"
	"ys-back/src/infrastructure/persistence/model/db"

	"github.com/labstack/echo/v4"
)

type TweetRepository struct {
	db     *sql.DB
	logger echo.Logger
}

func NewTweetRepository(db *sql.DB, logger echo.Logger) repository.ITweetRepository {
	return &TweetRepository{db, logger}
}

func (repo *TweetRepository) FetchTweets(count uint, offset uint) (*db.TweetRoot, error) {
	stmt, err := repo.db.Prepare(`SELECT id, content, created_at, user_name, name, avatar_url, tag_id FROM tweets ORDER BY id DESC LIMIT $1 OFFSET $2`)
	if err != nil {
		repo.logger.Error(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(count, offset)
	if err != nil {
		repo.logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	tweets := make([]*db.Tweet, count)
	var cnt = uint(0)
	for rows.Next() {
		var id string
		var content string
		var createdAt string
		var userName string
		var name string
		var avatarURL string
		var tagID uint

		if err := rows.Scan(&id, &content, &createdAt, &userName, &name, &avatarURL, &tagID); err != nil {
			repo.logger.Error(err)
		}

		tweets[cnt] = db.NewTweet(id, content, createdAt, tagID, userName, name, avatarURL)
		cnt++
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	repo.logger.Debug("result count ", cnt)

	if cnt < count {
		return db.NewTweetRoot(tweets[:cnt]), nil
	}
	return db.NewTweetRoot(tweets), nil
}

func (repo *TweetRepository) FetchTweetsWithTagID(tagID uint, count uint, offset uint) (*db.TweetRoot, error) {
	stmt, err := repo.db.Prepare(`SELECT id, content, created_at, user_name, name, avatar_url, tag_id FROM tweets WHERE tag_id = $1 ORDER BY id DESC LIMIT $2 OFFSET $3`)
	if err != nil {
		repo.logger.Error(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(tagID, count, offset)
	if err != nil {
		repo.logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	tweets := make([]*db.Tweet, count)
	var cnt = uint(0)
	for rows.Next() {
		var id string
		var content string
		var createdAt string
		var userName string
		var name string
		var avatarURL string
		var tagID uint

		if err := rows.Scan(&id, &content, &createdAt, &userName, &name, &avatarURL, &tagID); err != nil {
			repo.logger.Error(err)
		}

		tweets[cnt] = db.NewTweet(id, content, createdAt, tagID, userName, name, avatarURL)
		cnt++
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	repo.logger.Debug("result count ", cnt)

	if cnt < count {
		return db.NewTweetRoot(tweets[:cnt]), nil
	}
	return db.NewTweetRoot(tweets), nil
}

func (repo *TweetRepository) FetchLatestID() (string, error) {
	stmt, err := repo.db.Prepare(`SELECT id FROM tweets ORDER BY id DESC LIMIT 1`)
	if err != nil {
		repo.logger.Error(err)
		return "", err
	}
	defer stmt.Close()

	var id string
	if err := stmt.QueryRow().Scan(&id); err != nil {
		repo.logger.Error(err)
		return "", nil
	}
	return id, nil
}

func (repo *TweetRepository) FetchCount() (uint, error) {
	stmt, err := repo.db.Prepare(`SELECT COUNT(id) FROM tweets`)
	if err != nil {
		repo.logger.Error(err)
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		repo.logger.Error(err)
		return 0, err
	}
	defer rows.Close()

	var count uint
	if err := stmt.QueryRow().Scan(&count); err != nil {
		repo.logger.Error(err.Error())
		return 0, nil
	}
	return count, nil
}

func (repo *TweetRepository) FetchCountWithTagID(tagID uint) (uint, error) {
	stmt, err := repo.db.Prepare(`SELECT COUNT(id) FROM tweets WHERE tag_id = $1`)
	if err != nil {
		repo.logger.Error(err)
		return 0, err
	}
	defer stmt.Close()

	var count uint
	if err := stmt.QueryRow(tagID).Scan(&count); err != nil {
		repo.logger.Error(err.Error())
		return 0, err
	}
	return count, nil
}

func (repo *TweetRepository) StoreTweets(tweets *db.TweetRoot) error {
	stmt, err := repo.db.Prepare(`INSERT INTO tweets (id, content, created_at, user_name, name, avatar_url, tag_id) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		repo.logger.Error(err)
		return err
	}
	defer stmt.Close()

	for _, t := range tweets.Items {
		if _, err := stmt.Exec(t.ID, t.Content, t.CreatedAt, t.UserName, t.Name, t.AvatarURL, t.TagID); err != nil {
			repo.logger.Error(err)
			return err
		}
	}
	return nil
}

func (repo *TweetRepository) PurgeTweets(count uint) error {
	stmt, err := repo.db.Prepare(`DELETE FROM tweets WHERE id IN (SELECT id FROM tweets ORDER BY id LIMIT $1)`)
	if err != nil {
		repo.logger.Error(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(count); err != nil {
		repo.logger.Error(err)
		return err
	}
	return nil
}
