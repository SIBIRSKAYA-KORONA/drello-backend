package repository

import (
	"math/rand"
	"time"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/board"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type BoardStore struct {
	DB      *gorm.DB
	linkLen uint
}

func CreateRepository(db *gorm.DB) board.Repository {
	return &BoardStore{DB: db, linkLen: 32}
}

func (boardStore *BoardStore) GenerateInviteLink(size uint) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, size)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func (boardStore *BoardStore) Create(uid uint, board *models.Board) error {
	board.InviteLink = boardStore.GenerateInviteLink(boardStore.linkLen)
	err := boardStore.DB.Create(board).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrConflict
	}
	err = boardStore.DB.Model(board).Association("Admins").Append(&models.User{ID: uid}).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrConflict
	}
	return nil
}

func (boardStore *BoardStore) GetBoardsByUser(uid uint) (models.Boards, models.Boards, error) {
	var adminsBoards []models.Board
	usr := &models.User{ID: uid}
	err := boardStore.DB.Model(usr).Preload("Admins").Related(&adminsBoards, "Admin").Error
	if err != nil {
		logger.Error(err)
		return nil, nil, errors.ErrUserNotFound
	}
	var membersBoards []models.Board
	err = boardStore.DB.Model(usr).Preload("Members").Related(&membersBoards, "Member").Error
	if err != nil {
		logger.Error(err)
		return nil, nil, errors.ErrBoardNotFound
	}
	// TODO: изменить запрос или вынести в отдельную функцию
	for i := range adminsBoards {
		for j := range adminsBoards[i].Admins {
			adminsBoards[i].Admins[j].Email = ""
			adminsBoards[i].Admins[j].Password = nil
		}
		for j := range adminsBoards[i].Members {
			adminsBoards[i].Members[j].Email = ""
			adminsBoards[i].Members[j].Password = nil
		}
	}
	for i := range membersBoards {
		for j := range membersBoards[i].Admins {
			membersBoards[i].Admins[j].Email = ""
			membersBoards[i].Admins[j].Password = nil
		}
		for j := range membersBoards[i].Members {
			membersBoards[i].Members[j].Email = ""
			membersBoards[i].Members[j].Password = nil
		}
	}
	//
	return adminsBoards, membersBoards, nil
}

func (boardStore *BoardStore) Get(bid uint) (*models.Board, error) {
	brd := new(models.Board)
	err := boardStore.DB.First(brd, bid).Error
	if err != nil {
		logger.Error(err)
		return nil, errors.ErrBoardNotFound
	}
	err = boardStore.DB.Model(brd).Select("id, name, surname, nickname, avatar").Related(&brd.Admins, "Admins").Error
	if err != nil {
		logger.Error(err)
		return nil, errors.ErrDbBadOperation
	}
	err = boardStore.DB.Model(brd).Select("id, name, surname, nickname, avatar").Related(&brd.Members, "Members").Error
	if err != nil {
		logger.Error(err)
		return nil, errors.ErrDbBadOperation
	}
	return brd, nil
}

func (boardStore *BoardStore) GetLabelsByID(bid uint) (models.Labels, error) {
	var lbls []models.Label
	err := boardStore.DB.Model(&models.Board{ID: bid}).Order("id").Related(&lbls, "bid").Error
	if err != nil {
		logger.Error(err)
		return nil, errors.ErrBoardNotFound
	}
	return lbls, nil
}

func (boardStore *BoardStore) GetColumnsByID(bid uint) (models.Columns, error) {
	var cols []models.Column
	err := boardStore.DB.Model(&models.Board{ID: bid}).Related(&cols, "bid").Error
	if err != nil {
		logger.Error(err)
		return nil, errors.ErrBoardNotFound
	}
	return cols, nil
}

func (boardStore *BoardStore) Update(newBoard *models.Board) error {
	oldBoard := new(models.Board)
	err := boardStore.DB.First(oldBoard, newBoard.ID).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrBoardNotFound
	}
	oldBoard.Name = newBoard.Name
	err = boardStore.DB.Save(oldBoard).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrConflict
	}
	return nil
}

func (boardStore *BoardStore) Delete(bid uint) error {
	//TODO: трэш, переделать под каскад, файлики здесь игнорируются
	var columns []models.Column
	err := boardStore.DB.Model(&models.Board{ID: bid}).Related(&columns, "bid").Error
	if err != nil {
		logger.Error(err)
		return errors.ErrDbBadOperation
	}

	for columnID := range columns {
		var tasks []models.Task
		errQuery := boardStore.DB.Model(&models.Column{ID: columns[columnID].ID}).Related(&tasks, "cid").Error
		if errQuery != nil {
			logger.Error(errQuery)
			return errors.ErrDbBadOperation
		}
		for taskID := range tasks {
			errQuery := boardStore.DB.Delete(&models.Task{ID: tasks[taskID].ID}).Error
			if errQuery != nil {
				logger.Error(errQuery)
				return errors.ErrDbBadOperation
			}
		}
		err = boardStore.DB.Delete(&models.Column{ID: columns[columnID].ID}).Error
		if err != nil {
			logger.Error(err)
			return errors.ErrDbBadOperation
		}
	}

	err = boardStore.DB.Delete(&models.Board{ID: bid}).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrBoardNotFound
	}
	return nil
}

func (boardStore *BoardStore) InviteMember(bid uint, member *models.User) error {
	brd := new(models.Board)
	err := boardStore.DB.First(brd, bid).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrBoardNotFound
	}
	err = boardStore.DB.Model(&brd).Association("Members").Append(member).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrDbBadOperation
	}
	return nil
}

func (boardStore *BoardStore) DeleteMember(bid uint, member *models.User) error {
	brd := new(models.Board)
	err := boardStore.DB.First(brd, bid).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrBoardNotFound
	}
	err = boardStore.DB.Model(&brd).Association("Members").Delete(member).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrDbBadOperation
	}

	// чистим таски текущей доски от текущего пользователя
	var columns []models.Column
	err = boardStore.DB.Model(&models.Board{ID: bid}).Related(&columns, "bid").Error
	if err != nil {
		logger.Error(err)
		return errors.ErrDbBadOperation
	}
	for columnID := range columns {
		var tasks []models.Task
		errQuery := boardStore.DB.Model(&models.Column{ID: columns[columnID].ID}).Related(&tasks, "cid").Error
		if errQuery != nil {
			logger.Error(errQuery)
			return errors.ErrDbBadOperation
		}
		for taskID := range tasks {
			errQuery := boardStore.DB.Model(&models.Task{ID: tasks[taskID].ID}).Association("Members").Delete(member).Error
			if errQuery != nil {
				logger.Error(errQuery)
				return errors.ErrDbBadOperation
			}
		}
	}

	return nil
}

func (boardStore *BoardStore) GetUsersForInvite(bid uint, nicknamePart string, limit uint) (models.Users, error) {
	var users []models.User
	brd, err := boardStore.Get(bid)
	if err != nil {
		logger.Error(err)
		return nil, errors.ErrBoardNotFound
	}
	var boardMembersAndAdminsIDs []uint
	for _, member := range brd.Members {
		boardMembersAndAdminsIDs = append(boardMembersAndAdminsIDs, member.ID)
	}
	for _, admin := range brd.Admins {
		boardMembersAndAdminsIDs = append(boardMembersAndAdminsIDs, admin.ID)
	}
	err = boardStore.DB.Select("id, name, surname, nickname, avatar").
		Limit(limit).
		Where("nickname LIKE ?", nicknamePart+"%").
		Not("id", boardMembersAndAdminsIDs).
		Find(&users).Error
	if err != nil {
		logger.Error(err)
		return nil, errors.ErrUserNotFound
	}
	return users, nil
}

func (boardStore *BoardStore) InviteMemberByLink(usr models.User, link string) (*models.Board, error) {
	brd := new(models.Board)
	err := boardStore.DB.Where("invite_link = ?", link).First(brd).Error
	if err != nil {
		logger.Error(err)
		return nil, errors.ErrBoardNotFound
	}
	err = boardStore.DB.Model(brd).Select("id").Related(&brd.Admins, "Admins").Error
	if err != nil {
		logger.Error(err)
	}
	for _, member := range brd.Admins {
		if member.ID == usr.ID {
			return brd, errors.ErrConflict
		}
	}
	err = boardStore.DB.Model(brd).Select("id").Related(&brd.Members, "Members").Error
	if err != nil {
		logger.Error(err)
	}
	for _, member := range brd.Members {
		if member.ID == usr.ID {
			return brd, errors.ErrConflict
		}
	}
	err = boardStore.DB.Model(&brd).Association("Members").Append(usr).Error
	if err != nil {
		logger.Error(err)
		return brd, errors.ErrConflict
	}
	return brd, nil
}

func (boardStore *BoardStore) UpdateInviteLink(bid uint) error {
	inviteLink := boardStore.GenerateInviteLink(boardStore.linkLen)
	err := boardStore.DB.Model(models.Board{}).Where("id = ? ", bid).
		UpdateColumn("invite_link", inviteLink).Error
	if err != nil {
		logger.Error(err)
		return errors.ErrDbBadOperation
	}
	return nil
}
