package main

import (
	"fmt"
	"log"
	"os"

	"broken_calc/consts"
	api_v1 "broken_calc/v1"
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
	apiV1Controller := api_v1.ApiV1Controller{Log: Log}

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET(fmt.Sprintf("/add/:%s/:%s", consts.First, consts.Second), apiV1Controller.Add)
			v1.GET(fmt.Sprintf("/sub/:%s/:%s", consts.First, consts.Second), apiV1Controller.Sub)
			v1.GET(fmt.Sprintf("/multiply/:%s/:%s", consts.First, consts.Second), apiV1Controller.Multiply)
			v1.GET(fmt.Sprintf("/divide/:%s/:%s", consts.First, consts.Second), apiV1Controller.Divide)
		}
	}
}

type CustomValidator struct {
	Validator validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
