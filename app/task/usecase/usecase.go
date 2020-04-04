package usecase

import (
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/task"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"
)

type TaskUseCase struct {
	taskRepo task.Repository
}

func CreateUseCase(taskRepo_ task.Repository) task.UseCase {
	return &TaskUseCase{taskRepo: taskRepo_}
}

func (taskUseCase *TaskUseCase) Create(tsk *models.Task) error {
	return taskUseCase.taskRepo.Create(tsk)
}

func (taskUseCase *TaskUseCase) Get(cid uint, tid uint) (*models.Task, error) {
	tsk, err := taskUseCase.taskRepo.Get(tid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	if tsk.Cid != cid {
		return nil, errors.ErrBoardsNotFound // TODO: TaskNotFound
	}
	return tsk, nil
}
