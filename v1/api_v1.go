package api_v1

import (
	"errors"
	"strconv"

	"broken_calc/consts"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ApiV1Controller struct {
	Log *zap.SugaredLogger
}

func (c ApiV1Controller) Add(ctx echo.Context) (err error) {
	firstArg := ctx.Param(consts.First)
	secondArg := ctx.Param(consts.Second)

	firstArgInt, secondArgInt, err := convertPathParamsToInt(firstArg, secondArg)
	if err != nil {
		c.Log.Warnw(err.Error())
		return ctx.JSON(400, err.Error())
	}

	c.Log.Infow(firstArg + "added to " + secondArg)
	return ctx.JSON(200, echo.Map{"result": firstArgInt + secondArgInt})
}

func (c ApiV1Controller) Sub(ctx echo.Context) (err error) {
	firstArg := ctx.Param(consts.First)
	secondArg := ctx.Param(consts.Second)

	firstArgInt, secondArgInt, err := convertPathParamsToInt(firstArg, secondArg)
	if err != nil {
		c.Log.Warnw(err.Error())
		return ctx.JSON(400, err.Error())
	}

	c.Log.Infow(secondArg + " subtracted from " + firstArg)
	return ctx.JSON(200, echo.Map{"result": firstArgInt - secondArgInt})
}

func (c ApiV1Controller) Multiply(ctx echo.Context) (err error) {
	firstArg := ctx.Param(consts.First)
	secondArg := ctx.Param(consts.Second)

	firstArgInt, secondArgInt, err := convertPathParamsToInt(firstArg, secondArg)
	if err != nil {
		c.Log.Warnw(err.Error())
		return ctx.JSON(400, err.Error())
	}

	c.Log.Warnw(firstArg + " multiply " + secondArg)
	return ctx.JSON(200, echo.Map{"result": firstArgInt * secondArgInt})
}

func (c ApiV1Controller) Divide(ctx echo.Context) (err error) {
	firstArg := ctx.Param(consts.First)
	secondArg := ctx.Param(consts.Second)

	firstArgInt, secondArgInt, err := convertPathParamsToInt(firstArg, secondArg)
	if err != nil {
		c.Log.Warnw(err.Error())
		return ctx.JSON(400, err.Error())
	}

	c.Log.Infow(firstArg + " divided by " + secondArg)
	return ctx.JSON(200, echo.Map{"result": firstArgInt / secondArgInt})
}

func convertPathParamsToInt(first, second string) (fInt, sInt int, err error) {
	fInt, err = strconv.Atoi(first)
	if err != nil {
		return fInt, sInt, errors.New("wrong first parameter")
	}

	sInt, err = strconv.Atoi(second)
	if err != nil {
		return fInt, sInt, errors.New("wrong second parameter")
	}

	return
}
