package main

import (
	"fmt"
	"log"

	"github.com/example/testing/cmd/app"
	"github.com/example/testing/config"
)

func main() {
	fmt.Println("Starting application...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app := app.NewApp(cfg)
	app.Run()

	// // Initialize logger
	// logger.Init(logger.LogConfig{
	// 	Level:            "info",
	// 	Format:           "json",
	// 	EnableCaller:     true,
	// 	EnableStacktrace: false,
	// })

	// // Connect to the database
	// dbConnection, err := database.ConnectDb(&database.DBConfig{
	// 	Driver:                   "postgres",
	// 	Host:                     "localhost",
	// 	Port:                     5432,
	// 	User:                     "prabhat",
	// 	Password:                 "prabhat",
	// 	DBName:                   "ChatAppGo",
	// 	MaxIdleConnection:        10,
	// 	MaxOpenConnection:        50,
	// 	ConnectionLifeTimeMinute: 30,
	// 	Logging:                  true,
	// })

	// if err != nil {
	// 	logger.Error(context.Background(), "database connection failed")
	// 	return
	// }

	// fmt.Printf("DB Connection: %v\n", dbConnection)

	// // Close database connection when done
	// if err := database.CloseDb(); err != nil {
	// 	logger.Error(context.Background(), "failed to close database")
	// 	return
	// }

	// fmt.Println("Database connection closed successfully.")

	// cfg := &cacheConfig.Config{
	// 	Driver:   "redis",
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// }
	// c, err := cache.Init(cfg)
	// if err != nil {
	// 	panic(err)
	// }

	// ctx := context.Background()

	// // Set a value
	// err = c.Set(ctx, "greeting:checking", "Hello Prabhat", 10*time.Second)
	// if err != nil {
	// 	fmt.Println("Set error:", err)
	// }

	// // Get a value
	// val, err := c.Get(ctx, "greeting:checking")
	// if err != nil {
	// 	fmt.Println("Get error:", err)
	// } else {
	// 	fmt.Println("Cache Value:", val)
	// }

	// if err := cache.Close(); err != nil {
	// 	logger.Error(context.Background(), "failed to close redis")
	// 	return
	// }
}
