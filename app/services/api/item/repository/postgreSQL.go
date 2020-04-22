package repository

import (
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/item"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ItemStore struct {
	DB *gorm.DB
}

func CreateRepository(db *gorm.DB) item.Repository {
	return &ItemStore{DB: db}
}

func (itemStore *ItemStore) Create(item *models.Item) error {
	err := itemStore.DB.Create(item).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrDbBadOperation
	}
	return nil
}

func (itemStore *ItemStore) Update(newItem *models.Item) error {
	return errors.ErrDbBadOperation

	// var oldItem models.Item
	// if err := itemStore.DB.Where("id = ?", newItem.ID).First(&oldItem).Error; err != nil {
	// 	logger.Error(err)
	// 	return errors.ErrUserNotFound
	// }

	// oldItem.Text = newItem.Text
	// oldItem.IsDone = newItem.IsDone

	// if err := userStore.DB.Save(oldUser).Error; err != nil {
	// 	logger.Error(err)
	// 	return errors.ErrDbBadOperation
	// }

}

func (itemStore *ItemStore) Delete(itid uint) error {
	err := itemStore.DB.Delete(&models.Item{ID: itid}).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrBoardNotFound
	}
	return nil
}
