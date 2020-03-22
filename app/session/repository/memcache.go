package repository

import (
	"errors"
	"fmt"
	_ "time"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/session"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/uuid"
)

type SessionStore struct {
	DB *memcache.Client
}

func CreateRepository(db *memcache.Client) session.Repository {
	return &SessionStore{db}
}

func (sessionStore *SessionStore) Create(session *models.Session) (string, error) {
	if session == nil {
		return "", errors.New("internal error")
	}
	session.SID = uuid.New().String()
	err := sessionStore.DB.Set(&memcache.Item{
		Key:        session.SID,
		Value:      []byte(fmt.Sprintf("%d", session.ID)),
		Expiration: int32(session.Expires.Unix()),
	})
	//
	fmt.Println(sessionStore.DB.Get(session.SID))
	//
	if err != nil {
		return "", err
	}
	return session.SID, nil
}

func (sessionStore *SessionStore) Get(sid string) (uint, bool) {
	idByte, err := sessionStore.DB.Get(sid)
	if err != nil {
		return 0, false
	}
	var id uint
	_, err = fmt.Sscan(string(idByte.Value), &id)
	if err != nil {
		return 0, false
	}
	return id, true
}

func (sessionStore *SessionStore) Delete(sid string) error {
	return sessionStore.DB.Delete(sid)
}