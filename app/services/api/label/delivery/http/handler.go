package http

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/label"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/middleware"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"

	"github.com/labstack/echo/v4"
)

type LabelHandler struct {
	useCase label.UseCase
}

/*
// GET присылать лейблы вместе с таской
*/

func CreateHandler(router *echo.Echo, useCase label.UseCase, mw *middleware.GoMiddleware) {
	handler := &LabelHandler{
		useCase: useCase,
	}
	router.POST("/boards/:bid/labels", handler.Create, mw.Sanitize, mw.CheckAuth, mw.CheckBoardMemberPermission)
	router.GET("/boards/:bid/labels/:lid", handler.Get, mw.CheckAuth, mw.CheckBoardMemberPermission)
	router.PUT("/boards/:bid/labels/:lid", handler.Update, mw.Sanitize, mw.CheckAuth,
		mw.CheckBoardMemberPermission, mw.CheckLabelInBoard)
	router.DELETE("/boards/:bid/labels/:lid", handler.Delete, mw.CheckAuth,
		mw.CheckBoardMemberPermission, mw.CheckLabelInBoard)
	router.POST("/boards/:bid/columns/:cid/tasks/:tid/labels/:lid", handler.AddLabelOnTask, mw.CheckAuth,
		mw.CheckBoardMemberPermission, mw.CheckLabelInBoard, mw.CheckColInBoard, mw.CheckTaskInCol)
	router.DELETE("/boards/:bid/columns/:cid/tasks/:tid/labels/:lid", handler.RemoveLabelFromTask, mw.CheckAuth,
		mw.CheckBoardMemberPermission, mw.CheckLabelInBoard, mw.CheckColInBoard, mw.CheckTaskInCol)
}

func (labelHandler *LabelHandler) Create(ctx echo.Context) error {
	var lbl models.Label
	body := ctx.Get("body").([]byte)
	err := lbl.UnmarshalJSON(body)
	if err != nil {
		logger.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	lbl.Bid = ctx.Get("bid").(uint)
	err = labelHandler.useCase.Create(&lbl)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	resp, err := lbl.MarshalJSON()
	if err != nil {
		logger.Error(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, string(resp))
}

func (labelHandler *LabelHandler) Get(ctx echo.Context) error {
	bid := ctx.Get("bid").(uint)
	var lid uint
	_, err := fmt.Sscan(ctx.Param("lid"), &lid)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	lbl, err := labelHandler.useCase.Get(bid, lid)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	resp, err := lbl.MarshalJSON()
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, string(resp))
}

func (labelHandler *LabelHandler) Update(ctx echo.Context) error {
	var lbl models.Label
	body := ctx.Get("body").([]byte)
	err := lbl.UnmarshalJSON(body)
	if err != nil {
		logger.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	lbl.ID = ctx.Get("lid").(uint)
	lbl.Bid = ctx.Get("bid").(uint)
	err = labelHandler.useCase.Update(lbl)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (labelHandler *LabelHandler) Delete(ctx echo.Context) error {
	lid := ctx.Get("lid").(uint)
	err := labelHandler.useCase.Delete(lid)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (labelHandler *LabelHandler) AddLabelOnTask(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (labelHandler *LabelHandler) RemoveLabelFromTask(ctx echo.Context) error {
    return ctx.NoContent(http.StatusOK)
}
