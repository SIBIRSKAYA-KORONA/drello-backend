package usecase

import (
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/session"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/user"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"
)

type SessionUseCase struct {
	sessionRepo session.Repository
	userRepo    user.Repository
}

func CreateUseCase(sessionRepo_ session.Repository, userRepo_ user.Repository) session.UseCase {
	return &SessionUseCase{
		sessionRepo: sessionRepo_,
		userRepo:    userRepo_,
	}
}

func (sessionUseCase *SessionUseCase) Create(user *models.User, sessionExpires int32) (string, error) {
	if user == nil {
		return "", errors.ErrInternal
	}
	realUser, err := sessionUseCase.userRepo.GetByNickname(user.Nickname)
	if err != nil {
		logger.Error(err)
		return "", errors.ErrUserNotFound
	}
	if realUser != nil && sessionUseCase.userRepo.CheckPassword(realUser.ID, user.Password) {
		ses := &models.Session{
			SID:     "",
			ID:      realUser.ID,
			Expires: sessionExpires,
		}
		sid, err := sessionUseCase.sessionRepo.Create(ses)
		if err != nil {
			logger.Error(err)
			return "", err
		}
		return sid, nil
	}
	return "", errors.ErrWrongPassword
}

func (sessionUseCase *SessionUseCase) Get(sid string) (uint, bool) {
	id, has := sessionUseCase.sessionRepo.Get(sid)
	if !has {
		logger.Info("not has sid", sid)
		return 0, false
	}
	return id, true
}

func (sessionUseCase *SessionUseCase) Delete(sid string) error {
	err := sessionUseCase.sessionRepo.Delete(sid)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
