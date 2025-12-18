package user

import (
	"api-digital-scoring/internal/entity"
	"api-digital-scoring/internal/user/dto"
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	u, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Service) GetById(ctx context.Context, id string) (entity.User, error) {
	u, err := s.repo.GetById(ctx, id)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s *Service) Create(ctx context.Context, user entity.User) (entity.User, error) {
	u := entity.User{
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	object, err := s.repo.Create(ctx, u)
	if err != nil {
		return object, err
	}

	return object, nil
}

func (s *Service) Update(ctx context.Context, user entity.User) (entity.User, error) {
	u := entity.User{
		Fullname: user.Fullname,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	object, err := s.repo.Update(ctx, u)
	if err != nil {
		return object, err
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

func BindRequest(request *dto.UserRequest) entity.User {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	return entity.User{
		Fullname: request.Fullname,
		Username: request.Username,
		Email:    request.Email,
		Password: string(hashed),
	}
}

func BindResponse(user *entity.User) []dto.Response {
	return []dto.Response{
		{
			ID:       user.ID,
			Fullname: user.Fullname,
			Username: user.Username,
			Email:    user.Email,
		},
	}
}
