package user

import (
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
)

//go:generate mockgen -source=usecase.go -package=mocks -destination=./mocks/user_usecase_mock.go
type UseCase interface {
	Create(user *models.User) error
	GetByID(uid uint) (*models.User, error)
	GetByNickname(nickname string) (*models.User, error)
	CheckPassword(uid uint, pass []byte) bool
	Update(oldPass []byte, newUser models.User) error
	Delete(uid uint) error
	GetUsersByNicknamePart(nicknamePart string, limit uint) ([]models.User, error)
}
