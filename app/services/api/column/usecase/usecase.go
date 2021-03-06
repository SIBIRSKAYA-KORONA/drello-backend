package usecase

import (
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/column"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"
)

type ColumnUseCase struct {
	columnRepo column.Repository
}

func CreateUseCase(columnRepo_ column.Repository) column.UseCase {
	return &ColumnUseCase{columnRepo: columnRepo_}
}

func (columnUseCase *ColumnUseCase) Create(column *models.Column) error {
	return columnUseCase.columnRepo.Create(column)
}

func (columnUseCase *ColumnUseCase) Get(bid uint, cid uint) (*models.Column, error) {
	col, err := columnUseCase.columnRepo.Get(cid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if col.Bid != bid {
		return nil, errors.ErrNoPermission
	}
	return col, nil
}

func (columnUseCase *ColumnUseCase) GetTasksByID(cid uint) (models.Tasks, error) {
	tsks, err := columnUseCase.columnRepo.GetTasksByID(cid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return tsks, nil
}

func (columnUseCase *ColumnUseCase) Update(newCol models.Column) error {
	err := columnUseCase.columnRepo.Update(newCol)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (columnUseCase *ColumnUseCase) Delete(cid uint) error {
	err := columnUseCase.columnRepo.Delete(cid)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
