package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/middleware"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/user"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/message"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type UserHandler struct {
	useCase user.UseCase
}

func CreateHandler(router *echo.Echo, useCase user.UseCase, mw *middleware.GoMiddleware) {
	handler := &UserHandler{
		useCase: useCase,
	}
	router.POST("/settings", handler.Create)
	router.GET("/profile/:id_or_nickname", handler.Get)
	router.GET("/settings", handler.GetAll, mw.CheckAuth) // получ все настройки
	router.GET("/boards", handler.GetBoards, mw.CheckAuth)
	router.PUT("/settings", handler.Update, mw.CheckAuth, mw.CSRFmiddle)
	router.DELETE("/settings", handler.Delete, mw.CheckAuth)
}

func (userHandler *UserHandler) Create(ctx echo.Context) error {
	usr := models.CreateUser(ctx)
	if usr == nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	usr.Avatar = fmt.Sprintf("%s://%s:%s%s",
		viper.GetString("frontend.protocol"),
		viper.GetString("frontend.ip"),
		viper.GetString("frontend.port"),
		viper.GetString("frontend.default_avatar"))
	sessionExpires := time.Now().AddDate(1, 0, 0)
	sid, err := userHandler.useCase.Create(usr, sessionExpires)
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}
	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sid,
		Path:    "/",
		Expires: sessionExpires,
	}
	ctx.SetCookie(cookie)
	return ctx.NoContent(http.StatusOK)
}

func (userHandler *UserHandler) Get(ctx echo.Context) error {
	userData, err := userHandler.useCase.GetByNickname(ctx.Param("id_or_nickname"))
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}
	body, err := message.GetBody(message.Pair{Name: "user", Data: *userData})
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, body)
}

func (userHandler *UserHandler) GetAll(ctx echo.Context) error {
	userID := ctx.Get("userID").(uint)
	userData, err := userHandler.useCase.GetByID(userID)
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}
	body, err := message.GetBody(message.Pair{Name: "user", Data: *userData})
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, body)
}

func (userHandler *UserHandler) GetBoards(ctx echo.Context) error {
	userID := ctx.Get("userID").(uint)
	bAdmin, bMember, err := userHandler.useCase.GetBoardsByID(userID)
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}
	body, err := message.GetBody(message.Pair{Name: "admin", Data: bAdmin}, message.Pair{Name: "member", Data: bMember})
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, body)
}

func (userHandler *UserHandler) Update(ctx echo.Context) error {
	userID := ctx.Get("userID").(uint)
	newUser := new(models.User)
	newUser.Name = ctx.FormValue("newName")
	newUser.Surname = ctx.FormValue("newSurname")
	newUser.Nickname = ctx.FormValue("newNickname")
	newUser.Email = ctx.FormValue("newEmail")
	newUser.Password = ctx.FormValue("newPassword")
	newUser.ID = userID
	oldPass := ctx.FormValue("oldPassword")
	avatarFileDescriptor, err := ctx.FormFile("avatar")
	if err != nil {
		logger.Error(err)
	}
	if useErr := userHandler.useCase.Update(oldPass, newUser, avatarFileDescriptor); useErr != nil {
		logger.Error(useErr)
		return ctx.JSON(errors.ResolveErrorToCode(useErr), message.ResponseError{Message: useErr.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}

func (userHandler *UserHandler) Delete(ctx echo.Context) error {
	sessionID := ctx.Get("sessionID").(string)
	userID := ctx.Get("userID").(uint)
	if userHandler.useCase.Delete(userID, sessionID) != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	newCookie := http.Cookie{Name: "session_id", Value: sessionID, Expires: time.Now().AddDate(-1, 0, 0)}
	ctx.SetCookie(&newCookie)
	return ctx.NoContent(http.StatusOK)
}
