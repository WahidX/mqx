package repository

type Repository interface {
	// PushMessage(ctx context.Context, msg *entities.Message) error
	// GetMessage(context.Context, string) (*entities.Message, error)
}

type respository struct {
}

func NewRepository() Repository {
	return &respository{}
}
