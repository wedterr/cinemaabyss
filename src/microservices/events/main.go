//go:generate go run github.com/lerenn/asyncapi-codegen/cmd/asyncapi-codegen@latest -i ./asyncapi.yaml -p handlers -o ./handlers/asyncapi.gen.go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lerenn/asyncapi-codegen/pkg/extensions/brokers/kafka"
	"github.com/lerenn/asyncapi-codegen/pkg/extensions/loggers"
	"github.com/lerenn/asyncapi-codegen/pkg/extensions/middlewares"
	"github.com/wedterr/cinemaabyss/handlers"
)

func main() {
	e := echo.New()

	kafkaURL := getEnv("KAFKA_BROKERS", "localhost:9092")
	logger := loggers.NewText()
	broker, err := kafka.NewController([]string{kafkaURL}, kafka.WithLogger(logger), kafka.WithGroupID("events-service"))
	if err != nil {
		panic(err)
	}

	// Create app controller
	ctrl, err := handlers.NewAppController(
		broker,
		handlers.WithLogger(logger), // Attach an internal logger
		handlers.WithMiddlewares(middlewares.Logging(logger))) // Attach a middleware to log messages
	if err != nil {
		panic(err)
	}
	defer ctrl.Close(context.Background())

	c, err := handlers.NewContainer(ctrl)
	if err != nil {
		panic(err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CreateMovieEvent - Создание события фильма
	e.POST("/api/events/movie", c.CreateMovieEvent)

	// CreatePaymentEvent - Создание события платежа
	e.POST("/api/events/payment", c.CreatePaymentEvent)

	// CreateUserEvent - Создание события пользователя
	e.POST("/api/events/user", c.CreateUserEvent)

	// GetEventsServiceHealth - Проверка работоспособности микросервиса событий
	e.GET("/api/events/health", c.GetEventsServiceHealth)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", getEnv("PORT", "8082"))))
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
