package label

import (
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
)

//go:generate mockgen -source=usecase.go -package=mocks -destination=./mocks/label_usecase_mock.go
type UseCase interface {
	Create(lbl *models.Label) error
	Get(bid uint, lid uint) (*models.Label, error)
	Update(lbl models.Label) error
	Delete(lid uint) error
	AddLabelOnTask(lid uint, tid uint) error
	RemoveLabelFromTask(lid uint, tid uint) error
}
