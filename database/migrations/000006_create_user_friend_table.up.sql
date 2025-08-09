DROP TABLE IF EXISTS user_friends;
CREATE TABLE user_friends (
	id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    friend_id INT NOT NULL REFERENCES users(id),
);

CREATE UNIQUE INDEX idx_user_friends_unique ON user_friends (user_id, friend_id);
