package user

import (
	"api-digital-scoring/internal/entity"
	"context"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	u, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) GetById(ctx context.Context, id string) (*entity.User, error) {
	u, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Service) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	u := &entity.User{
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	object, err := s.repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func (s *Service) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	u := &entity.User{
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	object, err := s.repo.Update(ctx, u)
	if err != nil {
		return nil, err
	}

	return object, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
