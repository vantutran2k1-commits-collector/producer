package main

import (
	"github.com/vantutran2k1-commits-collector/producer/app"
	"github.com/vantutran2k1-commits-collector/producer/config"
	"log"
)

func main() {
	a := app.InitApp()
	if err := a.Router.Run(":" + config.AppEnv.AppPort); err != nil {
		log.Fatal(err)
		return
	}
}
