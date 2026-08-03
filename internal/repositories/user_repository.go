package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/kongali1720/KongPay/internal/models"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(user *models.User) error {

	query := `
	INSERT INTO users (
		id,
		full_name,
		email,
		password_hash,
		phone,
		status
	)
	VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		user.ID,
		user.FullName,
		user.Email,
		user.PasswordHash,
		user.Phone,
		user.Status,
	)

	return err
}

func (r *Repository) FindByID(id uuid.UUID) (*models.User, error) {

	user := &models.User{}

	query := `
	SELECT
		id,
		full_name,
		email,
		password_hash,
		phone,
		status,
		created_at,
		updated_at
	FROM users
	WHERE id=$1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.PasswordHash,
		&user.Phone,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) FindByEmail(email string) (*models.User, error) {

	user := &models.User{}

	query := `
	SELECT
		id,
		full_name,
		email,
		password_hash,
		phone,
		status,
		created_at,
		updated_at
	FROM users
	WHERE email=$1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.PasswordHash,
		&user.Phone,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
