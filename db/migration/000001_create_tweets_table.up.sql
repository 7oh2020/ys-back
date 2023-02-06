
CREATE TABLE IF NOT EXISTS tweets (
    id varchar(22) PRIMARY KEY,
    content text NOT NULL,
    created_at varchar(255) NOT NULL,
        tag_id smallint DEFAULT 0,
    user_name varchar(15) NOT NULL,
    name varchar(50) NOT NULL,
    avatar_url text NOT NULL
                );

CREATE INDEX tweets_tag_ID_index ON tweets(tag_id);
