package http

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/models"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/middleware"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/app/services/api/task"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/errors"
	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	useCase task.UseCase
}

func CreateHandler(router *echo.Echo, useCase task.UseCase, mw *middleware.Middleware) {
	handler := &TaskHandler{useCase: useCase}
	router.POST("boards/:bid/columns/:cid/tasks", handler.Create, mw.Sanitize,
		mw.CheckAuth, mw.CheckBoardMemberPermission, mw.CheckColInBoard)
	router.GET("boards/:bid/columns/:cid/tasks/:tid", handler.Get,
		mw.CheckAuth, mw.CheckBoardMemberPermission, mw.CheckColInBoard)
	router.PUT("boards/:bid/columns/:cid/tasks/:tid", handler.Update, mw.Sanitize,
		mw.CheckAuth, mw.CheckBoardMemberPermission, mw.CheckColInBoard, mw.CheckTaskInCol)
	router.DELETE("boards/:bid/columns/:cid/tasks/:tid", handler.Delete,
		mw.CheckAuth, mw.CheckBoardMemberPermission, mw.CheckColInBoard, mw.CheckTaskInCol)
	router.POST("boards/:bid/columns/:cid/tasks/:tid/members/:uid", handler.Assign,
		mw.CheckAuth, mw.CheckBoardMemberPermission, mw.CheckColInBoard, mw.CheckTaskInCol, mw.CheckUserForAssignInBoard)
	router.DELETE("boards/:bid/columns/:cid/tasks/:tid/members/:uid", handler.Unassign,
		mw.CheckAuth, mw.CheckBoardMemberPermission, mw.CheckColInBoard, mw.CheckTaskInCol, mw.CheckUserForAssignInBoard)
}

func (taskHandler *TaskHandler) Create(ctx echo.Context) error {
	var tsk models.Task
	body := ctx.Get("body").([]byte)
	err := tsk.UnmarshalJSON(body)
	if err != nil {
		logger.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	tsk.Cid = ctx.Get("cid").(uint)
	err = taskHandler.useCase.Create(&tsk)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	resp, err := tsk.MarshalJSON()
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, string(resp))
}

func (taskHandler *TaskHandler) Get(ctx echo.Context) error {
	cid := ctx.Get("cid").(uint)
	var tid uint
	_, err := fmt.Sscan(ctx.Param("tid"), &tid)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	tsk, err := taskHandler.useCase.Get(cid, tid)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	resp, err := tsk.MarshalJSON()
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.String(http.StatusOK, string(resp))
}

func (taskHandler *TaskHandler) Update(ctx echo.Context) error {
	var tsk models.Task
	body := ctx.Get("body").([]byte)
	err := tsk.UnmarshalJSON(body)
	if err != nil {
		logger.Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	tsk.ID = ctx.Get("tid").(uint)
	err = taskHandler.useCase.Update(tsk)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (taskHandler *TaskHandler) Delete(ctx echo.Context) error {
	tid := ctx.Get("tid").(uint)
	err := taskHandler.useCase.Delete(tid)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (taskHandler *TaskHandler) Assign(ctx echo.Context) error {
	tid := ctx.Get("tid").(uint)
	assignUid := ctx.Get("uid_for_assign").(uint)
	err := taskHandler.useCase.Assign(tid, assignUid)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}

func (taskHandler *TaskHandler) Unassign(ctx echo.Context) error {
	tid := ctx.Get("tid").(uint)
	assignUid := ctx.Get("uid_for_assign").(uint)
	err := taskHandler.useCase.Unassign(tid, assignUid)
	if err != nil {
		logger.Error(err)
		return ctx.String(errors.ResolveErrorToCode(err), err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}