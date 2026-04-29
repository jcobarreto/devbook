package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), nil
}

func (repository Users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repository.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick,
		nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Created_At); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) GetByID(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, name, nick, email, created_at FROM users WHERE id = ?",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Created_At); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (repository Users) GetByEmail(email string) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, password FROM users WHERE email = ?",
		email,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) Follow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository Users) Unfollow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository Users) GetFollowers(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM users u
		INNER JOIN followers f ON u.id = f.follower_id
		WHERE f.user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Created_At); err != nil {
			return nil, err
		}

		followers = append(followers, user)
	}

	return followers, nil
}

func (repository Users) GetFollowing(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM users u
		INNER JOIN followers f ON u.id = f.user_id
		WHERE f.follower_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var following []models.User

	for lines.Next() {
		var user models.User
		if err = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Created_At); err != nil {
			return nil, err
		}

		following = append(following, user)
	}

	return following, nil
}

func (repository Users) GetPassword(userID uint64) (string, error) {
	line, err := repository.db.Query(
		"SELECT password FROM users WHERE id = ?",
		userID,
	)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository Users) UpdatePassword(userID uint64, newPassword string) error {
	statement, err := repository.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(newPassword, userID); err != nil {
		return err
	}

	return nil
}
