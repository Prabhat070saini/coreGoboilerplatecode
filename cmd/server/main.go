package main

import (
	"log"

	"github.com/example/testing/cmd/app"
	"github.com/example/testing/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	
	app := app.NewApp(cfg)
	app.Run()
	

}
