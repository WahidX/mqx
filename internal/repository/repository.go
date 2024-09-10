package repository

import "database/sql"

type Repository interface {
	// PushMessage(ctx context.Context, msg *entities.Message) error
	// GetMessage(context.Context, string) (*entities.Message, error)
}

type respository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &respository{
		db: db,
	}
}
