package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare("insert into posts (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (repository Posts) GetByID(postID uint64) (models.Post, error) {
	line, err := repository.db.Query(`
		SELECT p.*, u.nick
		FROM posts p INNER JOIN users u
		ON u.id = p.author_id
		WHERE p.id = ?`, postID)
	if err != nil {
		return models.Post{}, err
	}
	defer line.Close()

	var post models.Post
	if line.Next() {
		if err = line.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository Posts) Get(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		SELECT DISTINCT p.*, u.nick
		FROM posts p INNER JOIN users u
		ON u.id = p.author_id
		INNER JOIN followers f
		ON p.author_id = f.user_id
		WHERE u.id = ? OR f.follower_id = ?
		ORDER BY 1 DESC`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post
	for lines.Next() {
		var post models.Post
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) Update(postID uint64, post models.Post) error {
	statement, err := repository.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) Delete(postID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) GetByUser(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		SELECT p.*, u.nick
		FROM posts p INNER JOIN users u
		ON u.id = p.author_id
		WHERE p.author_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post
	for lines.Next() {
		var post models.Post
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) Like(postID uint64) error {
	statement, err := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository Posts) Unlike(postID uint64) error {
	statement, err := repository.db.Prepare(`
		UPDATE posts SET likes =
		CASE
			WHEN likes > 0 THEN likes - 1
			ELSE 0
		END
		WHERE id = ?
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}
