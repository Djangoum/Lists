package repositories

import (
	"database/sql"
	"log"

	"home.com/lists/backend/entities"
)

type UsersMysqlRepository struct {
	db *sql.DB
}

func NewUsersMysqlRepository(db *sql.DB) *UsersMysqlRepository {
	return &UsersMysqlRepository{
		db: db,
	}
}

func (repo *UsersMysqlRepository) ListUsers() ([]*entities.User, error) {
	stmt, err := repo.db.Prepare(`SELECT id, email, first_name, last_name, created_at, updated_at FROM users`)

	if err != nil {
		return nil, err
	}

	users := []*entities.User{}
	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user entities.User
		rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)

		users = append(users, &user)
	}

	return users, nil
}

func (repo *UsersMysqlRepository) CreateUser(user *entities.User) (int64, error) {
	stmt, err := repo.db.Prepare(`INSERT INTO users (email, password, first_name, last_name, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.CreatedAt,
		user.UpdatedAt,
	)

	log.Println(result)

	stmt.Close()
	userID, err := result.LastInsertId()
	return userID, nil
}

func (repo *UsersMysqlRepository) DoesUserEmailExists(email string) (bool, error) {
	stmt, err := repo.db.Prepare(`SELECT COUNT(1) FROM users WHERE email = ?`)

	if err != nil {
		return false, err
	}

	var count int
	rows, err := stmt.Query(email)

	for rows.Next() {
		rows.Scan(&count)
		return count >= 1, nil
	}

	stmt.Close()

	return false, nil
}
