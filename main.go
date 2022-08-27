package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.SugaredLogger
)

func init() {
	logger, err := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      false,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(zap.Fields(zapcore.Field{
		Type: zapcore.StringType,
		Key:  "BuildInfo",
	}))
	if err != nil {
		fmt.Printf("Error loading logger: %+v", err)
		os.Exit(1)
	}
	Log = logger.Sugar()
}

func main() {
	defer func() {
		if err := Log.Sync(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	router := echo.New()

	router.Validator = &CustomValidator{*validator.New()}
	registerRoutes(router)

	log.Fatalln(router.Start(":6666"))
}

func registerRoutes(router *echo.Echo) {
	api := router.Group("/api")
	{
		api.GET(fmt.Sprintf("/add/:%s/:%s", first, second), Add)
		api.GET(fmt.Sprintf("/sub/:%s/:%s", first, second), Sub)
		api.GET(fmt.Sprintf("/multiply/:%s/:%s", first, second), Multiply)
		api.GET(fmt.Sprintf("/divide/:%s/:%s", first, second), Divide)
	}
}

const (
	first  = "first"
	second = "second"
)

type CustomValidator struct {
	Validator validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func Add(ctx echo.Context) (err error) {
	firstArg := ctx.Param(first)
	secondArg := ctx.Param(second)

	firstArgInt, err := strconv.Atoi(firstArg)
	if err != nil {
		Log.Warnw(err.Error())
		return ctx.JSON(400, echo.Map{"msg": "wrong first argument"})
	}
	secondArgInt, err := strconv.Atoi(secondArg)
	if err != nil {
		Log.Warnw(err.Error())
		return ctx.JSON(400, echo.Map{"msg": "wrong second argument"})
	}

	return ctx.JSON(200, echo.Map{"result": firstArgInt + secondArgInt})
}

func Sub(ctx echo.Context) (err error) {
	firstArg := ctx.Param(first)
	secondArg := ctx.Param(second)

	firstArgInt, err := strconv.Atoi(firstArg)
	if err != nil {
		Log.Warnw(err.Error())
		return ctx.JSON(400, echo.Map{"msg": "wrong first argument"})
	}
	secondArgInt, err := strconv.Atoi(secondArg)
	if err != nil {
		Log.Warnw(err.Error())
		return ctx.JSON(400, echo.Map{"msg": "wrong second argument"})
	}

	return ctx.JSON(200, echo.Map{"result": firstArgInt - secondArgInt})
}

func Multiply(ctx echo.Context) (err error) {
	firstArg := ctx.Param(first)
	secondArg := ctx.Param(second)

	firstArgInt, err := strconv.Atoi(firstArg)
	if err != nil {
		Log.Warnw(err.Error())
		return ctx.JSON(400, echo.Map{"msg": "wrong first argument"})
	}
	secondArgInt, err := strconv.Atoi(secondArg)
	if err != nil {
		Log.Warnw(err.Error())
		return ctx.JSON(400, echo.Map{"msg": "wrong second argument"})
	}

	return ctx.JSON(200, echo.Map{"result": firstArgInt * secondArgInt})
}

func Divide(ctx echo.Context) (err error) {
	firstArg := ctx.Param(first)
	secondArg := ctx.Param(second)

	firstArgInt, err := strconv.Atoi(firstArg)
	if err != nil {
		Log.Warnw(err.Error())
		return ctx.JSON(400, echo.Map{"msg": "wrong first argument"})
	}
	secondArgInt, err := strconv.Atoi(secondArg)
	if err != nil {
		Log.Warnw(err.Error())
		return ctx.JSON(400, echo.Map{"msg": "wrong second argument"})
	}

	return ctx.JSON(200, echo.Map{"result": firstArgInt / secondArgInt})
}
