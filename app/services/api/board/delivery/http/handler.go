package http

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/board"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/middleware"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/message"

	"github.com/labstack/echo/v4"
)

type BoardHandler struct {
	useCase board.UseCase
}

func CreateHandler(router *echo.Echo, useCase board.UseCase, mw *middleware.GoMiddleware) {
	handler := &BoardHandler{
		useCase: useCase,
	}
	router.POST("/boards", handler.Create, mw.CheckAuth)
	router.GET("/boards", handler.GetBoardsByUser, mw.CheckAuth)
	router.GET("/boards/:bid", handler.Get, mw.CheckAuth)
	router.GET("/boards/:bid/columns", handler.GetColumns, mw.CheckAuth, mw.CheckBoardMemberPermission)
	router.PUT("/boards/:bid", handler.Update, mw.CheckAuth, mw.CheckBoardAdminPermission)
	router.DELETE("/boards/:bid", handler.Delete, mw.CheckAuth, mw.CheckBoardAdminPermission) // TODO: что если есть другие админы
	router.POST("/boards/:bid/members/:uid", handler.InviteMember, mw.CheckAuth, mw.CheckBoardMemberPermission)
	router.DELETE("/boards/:bid/members/:uid", handler.DeleteMember, mw.CheckAuth, mw.CheckBoardAdminPermission)
	//GET /board/{bid}/search_for_invite?nickname={part_of_nickname}
	router.GET("/boards/:bid/search_for_invite", handler.GetUsersForInvite, mw.CheckAuth, mw.CheckBoardMemberPermission)
}

func (boardHandler *BoardHandler) Create(ctx echo.Context) error {
	uid := ctx.Get("uid").(uint)
	brd := models.CreateBoard(ctx)
	if brd == nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err := boardHandler.useCase.Create(uid, brd)
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}
	body, err := message.GetBody(message.Pair{Name: "board", Data: *brd})
	if err != nil {
		logger.Error(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, body)
}

func (boardHandler *BoardHandler) GetBoardsByUser(ctx echo.Context) error {
	uid := ctx.Get("uid").(uint)
	bAdmin, bMember, err := boardHandler.useCase.GetBoardsByUser(uid)
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

func (boardHandler *BoardHandler) Get(ctx echo.Context) error {
	uid := ctx.Get("uid").(uint)
	var bid uint
	_, err := fmt.Sscan(ctx.Param("bid"), &bid)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	brd, err := boardHandler.useCase.Get(uid, bid, false)
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}
	body, err := message.GetBody(message.Pair{Name: "board", Data: *brd})
	if err != nil {
		logger.Error(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, body)
}

func (boardHandler *BoardHandler) GetColumns(ctx echo.Context) error {
	bid := ctx.Get("bid").(uint)
	cols, err := boardHandler.useCase.GetColumnsByID(bid)
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}
	body, err := message.GetBody(message.Pair{Name: "columns", Data: cols})
	if err != nil {
		logger.Error(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, body)
}

func (boardHandler *BoardHandler) Update(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (boardHandler *BoardHandler) Delete(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (boardHandler *BoardHandler) InviteMember(ctx echo.Context) error {
	bid := ctx.Get("bid").(uint)

	var uid uint
	_, err := fmt.Sscan(ctx.Param("uid"), &uid)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = boardHandler.useCase.InviteMember(bid, uid)
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}

func (boardHandler *BoardHandler) DeleteMember(ctx echo.Context) error {
	bid := ctx.Get("bid").(uint)

	var uid uint
	_, err := fmt.Sscan(ctx.Param("uid"), &uid)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = boardHandler.useCase.DeleteMember(bid, uid)
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}
	return ctx.NoContent(http.StatusOK)
}

func (boardHandler *BoardHandler) GetUsersForInvite(ctx echo.Context) error {
	nicknamePart := ctx.QueryParam("nickname")
	if nicknamePart == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}

	bid := ctx.Get("bid").(uint)

	var limit uint
	_, err := fmt.Sscan(ctx.QueryParam("limit"), &limit)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	usersData, err := boardHandler.useCase.GetUsersForInvite(bid, nicknamePart, limit)
	if err != nil {
		logger.Error(err)
		return ctx.JSON(errors.ResolveErrorToCode(err), message.ResponseError{Message: err.Error()})
	}

	body, err := message.GetBody(message.Pair{Name: "user", Data: usersData})
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, body)
}
