package main

import (
	shine "github.com/liangguifeng/shine"
	"github.com/liangguifeng/shine/app"
	"shine-skeleton/startup"
)

const APP_NAME = "shine-skeleton"

func main() {
	runer, err := app.NewRunner(&shine.Application{
		Name:       APP_NAME,
		Type:       shine.APP_TYPE_GRPC,
		LoadConfig: startup.LoadConfig,
		SetupVars:  startup.SetupVars,
	})
}
