package repositories

import (
	"time"

	"github.com/snowlynxsoftware/na-api/server/database"
	"github.com/snowlynxsoftware/na-api/server/models"
)

type UserEntity struct {
	ID            int64      `json:"id" db:"id"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	ModifiedAt    *time.Time `json:"modified_at" db:"modified_at"`
	IsArchived    bool       `json:"is_archived" db:"is_archived"`
	Email         string     `json:"email" db:"email"`
	Username      string     `json:"username" db:"username"`
	IsVerified    bool       `json:"is_verified" db:"is_verified"`
	UserTypeID    *int       `json:"user_type_id" db:"user_type_id"`
	PasswordHash  *string    `json:"-" db:"password_hash"`
	LastLogin     *time.Time `json:"last_login" db:"last_login"`
	IsBanned      bool       `json:"is_banned" db:"is_banned"`
	BanReason     *string    `json:"ban_reason" db:"ban_reason"`
	CanUseAPIKeys bool       `json:"can_use_api_keys" db:"can_use_api_keys"`
}

type UserRepository struct {
	DB *database.AppDataSource
}

func NewUserRepository(db *database.AppDataSource) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetUserById(id int) (*UserEntity, error) {
	user := &UserEntity{}
	sql := `SELECT
		*
	FROM users
	WHERE id = $1`
	err := r.DB.DB.Get(user, sql, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*UserEntity, error) {
	user := &UserEntity{}
	sql := `SELECT
		*
	FROM users
	WHERE email = $1`
	err := r.DB.DB.Get(user, sql, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) CreateNewUser(dto *models.UserCreateDTO) (*UserEntity, error) {
	sql := `INSERT INTO users (email, username, password_hash)
    VALUES ($1, $2, $3)
    RETURNING id;`
	row := r.DB.DB.QueryRow(sql, dto.Email, dto.Username, dto.Password)
	var insertedId int
	err := row.Scan(&insertedId)
	if err != nil {
		return nil, err
	}

	user, err := r.GetUserById(insertedId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) MarkUserVerified(userId *int) (bool, error) {
	sql := `UPDATE users SET is_verified = true WHERE id = $1;`
	_, err := r.DB.DB.Exec(sql, &userId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) UpdateUserLastLogin(userId *int) (bool, error) {
	sql := `UPDATE users SET last_login = NOW() WHERE id = $1;`
	_, err := r.DB.DB.Exec(sql, &userId)
	if err != nil {
		return false, err
	}

	return true, nil
}
