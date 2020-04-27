package http

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/column"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/middleware"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"

	"github.com/labstack/echo/v4"
)

type ColumnHandler struct {
	useCase column.UseCase
}

func CreateHandler(router *echo.Echo, useCase column.UseCase, mw *middleware.Middleware) {
	handler := &ColumnHandler{useCase: useCase}
	// TODO: обсудить кто может создавать колонки
	router.POST("/boards/:bid/columns", handler.Create, mw.Sanitize, mw.CheckAuth, mw.CheckBoardMemberPermission)
	router.GET("/boards/:bid/columns/:cid", handler.Get, mw.CheckAuth, mw.CheckBoardMemberPermission)
	router.GET("/boards/:bid/columns/:cid/tasks", handler.GetTasks,
		mw.CheckAuth, mw.CheckBoardMemberPermission, mw.CheckColInBoard)
	router.PUT("/boards/:bid/columns/:cid", handler.Update,
		mw.CheckAuth, mw.CheckBoardMemberPermission, mw.CheckColInBoard)
	router.DELETE("/boards/:bid/columns/:cid", handler.Delete,
		mw.CheckAuth, mw.CheckBoardMemberPermission, mw.CheckColInBoard)
}

func (columnHandler *ColumnHandler) Create(ctx echo.Context) error {
	var col models.Column
	body := ctx.Get("body").([]byte)
	err := col.UnmarshalJSON(body)
	if err != nil {
		logger.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	col.Bid = ctx.Get("bid").(uint)
	err = columnHandler.useCase.Create(&col)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	resp, err := col.MarshalJSON()
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, string(resp))
}

func (columnHandler *ColumnHandler) Get(ctx echo.Context) error {
	bid := ctx.Get("bid").(uint)
	var cid uint
	_, err := fmt.Sscan(ctx.Param("cid"), &cid)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	col, err := columnHandler.useCase.Get(bid, cid)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	resp, err := col.MarshalJSON()
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, string(resp))
}

func (columnHandler *ColumnHandler) GetTasks(ctx echo.Context) error {
	cid := ctx.Get("cid").(uint)
	tsks, err := columnHandler.useCase.GetTasksByID(cid)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	resp, err := tsks.MarshalJSON()
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, string(resp))
}

func (columnHandler *ColumnHandler) Update(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

// TODO: проверить удаление колонки, если в ней есть таски
func (columnHandler *ColumnHandler) Delete(ctx echo.Context) error {
	cid := ctx.Get("cid").(uint)
	err := columnHandler.useCase.Delete(cid)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}