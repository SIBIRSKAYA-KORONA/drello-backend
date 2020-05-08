package notification

import "github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"

type Repository interface {
	GetEvents(uid uint) (models.Events, bool)
}
