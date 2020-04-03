package user

import (
	"mime/multipart"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
)

//go:generate mockgen -source=repository.go -package=mocks -destination=./mocks/user_repo_mock.go

type Repository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByNickname(nickname string) (*models.User, error)
	Update(oldPass string, newUser *models.User, avatarFileDescriptor *multipart.FileHeader) error
	Delete(id uint) error
}
