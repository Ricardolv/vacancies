package main

import (
	"github.com/Ricardolv/vacancies/config"
	"github.com/Ricardolv/vacancies/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// initialize Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// initialize Router
	router.Initialize()
}
