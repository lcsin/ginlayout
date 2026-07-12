package repository

import (
	"context"

	"ginlayout/internal/model"
	"ginlayout/internal/repository/query"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id uint) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
	q  *query.Query
}

func NewUserRepository(db *gorm.DB) UserRepository {
	// Auto migrate the table for demo purposes
	_ = db.AutoMigrate(&model.User{})
	return &userRepository{
		db: db,
		q:  query.Use(db),
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	u := r.q.User
	return u.WithContext(ctx).Create(user)
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	u := r.q.User
	return u.WithContext(ctx).Where(u.ID.Eq(id)).First()
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	u := r.q.User
	return u.WithContext(ctx).Where(u.Username.Eq(username)).First()
}
