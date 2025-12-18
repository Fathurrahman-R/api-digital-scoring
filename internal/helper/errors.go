package helper

import (
	"api-digital-scoring/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type ErrorMessages struct {
	Message []string `json:"message"`
}

var (
	ErrorUserNotFound = errors.New("pengguna tidak ditemukan")
)

func NewErrorResponse(messages ...string) ErrorMessages {
	return ErrorMessages{
		Message: messages,
	}
}

func UserNotFound(err error) (entity.User, error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, ErrorUserNotFound
	}
	return entity.User{}, err
}
