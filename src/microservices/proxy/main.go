package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	toggleMigration, err := strconv.ParseBool(getEnv("GRADUAL_MIGRATION", "true"))
	if err != nil {
		e.Logger.Fatal(err)
	}
	moviesMigrationPercent, err := strconv.Atoi(getEnv("MOVIES_MIGRATION_PERCENT", "50"))
	if err != nil {
		e.Logger.Fatal(err)
	}
	monolithUrl, err := url.Parse(getEnv("MONOLITH_URL", "http://monolith:8080"))
	if err != nil {
		e.Logger.Fatal(err)
	}
	moviesUrl, err := url.Parse(getEnv("MOVIES_SERVICE_URL", "http://movies-service:8081"))
	if err != nil {
		e.Logger.Fatal(err)
	}
	eventsUrl, err := url.Parse(getEnv("EVENTS_SERVICE_URL", "http://events-service:8082"))
	if err != nil {
		e.Logger.Fatal(err)
	}

	usersTargets := []*middleware.ProxyTarget{
		{
			URL: monolithUrl,
		},
	}

	eventsTargets := []*middleware.ProxyTarget{
		{
			URL: eventsUrl,
		},
	}

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "Strangler Fig Proxy is healthy")
	})
	proxyMovies := e.Group("/api/movies")
	proxyUsers := e.Group("/api/users")
	proxyEvents := e.Group("/api/events")

	// Apply the proxy middleware to the group
	if toggleMigration {
		moviesTargets := []*middleware.ProxyTarget{
			{
				URL: monolithUrl,
			},
			{
				URL: moviesUrl,
			},
		}
		migrationBalancer := NewMigrationBalancer(moviesTargets, moviesMigrationPercent)
		proxyMovies.Use(middleware.Proxy(migrationBalancer))
	} else {
		moviesTargets := []*middleware.ProxyTarget{
			{
				URL: monolithUrl,
			},
		}
		proxyMovies.Use(middleware.Proxy(middleware.NewRandomBalancer(moviesTargets)))
	}
	proxyUsers.Use(middleware.Proxy(middleware.NewRandomBalancer(usersTargets)))
	proxyEvents.Use(middleware.Proxy(middleware.NewRandomBalancer(eventsTargets)))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", getEnv("PORT", "8000"))))
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
