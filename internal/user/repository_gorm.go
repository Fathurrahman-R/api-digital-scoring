package user

import (
	"api-digital-scoring/internal/entity"
	"api-digital-scoring/internal/helper"
	"context"
	"errors"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) Repository {
	return &userRepo{db}
}

func (u userRepo) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User

	err := u.db.WithContext(ctx).
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u userRepo) GetById(ctx context.Context, id string) (entity.User, error) {
	var user entity.User

	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error

	if err != nil {
		return helper.UserNotFound(err)
	}

	return user, nil
}

func (u userRepo) Create(ctx context.Context, user entity.User) (entity.User, error) {
	uEmail := u.db.WithContext(ctx).Where("email = ?", &user.Email).First(&user).Error

	if uEmail == nil {
		return user, errors.New("email already in use")
	}

	err := u.db.WithContext(ctx).Create(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u userRepo) Update(ctx context.Context, user entity.User) (entity.User, error) {
	err := u.db.WithContext(ctx).Save(&user).Error

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u userRepo) Delete(ctx context.Context, id string) error {
	var user entity.User

	err := u.db.WithContext(ctx).Where("id = ?", id).Delete(&user).Error

	if err != nil {
		return err
	}

	return nil
}
