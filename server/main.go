package main

import (
	"SpeedLoggerz/server/database"
	"SpeedLoggerz/server/routes"
	"SpeedLoggerz/server/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	version := "0.0.0.1"
	port := 555
	portString := fmt.Sprintf("%s%d", ":", port)

	_ = utils.Load()

	utils.LoadGlobals()

	database.InitialiseDatabaseConn()

	e := echo.New()
	e.HTTPErrorHandler = utils.HTTPErrorHandler
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
	}))

	routes.HookupRoutes(e) // initialise all routes

	/* Routes below that must be declared after HookupRoutes */
	e.GET("/", func(c echo.Context) error {
		fmt.Println()
		return c.String(http.StatusOK, "SpeedLoggerz Server")
	})

	e.GET("/version/", func(c echo.Context) error {
		return c.String(http.StatusOK, version)
	})

	// e.GET("/ws/", ws.WsHandler)  // might need websocket, might not.

	e.Logger.Fatal(e.Start(portString))
	http.ListenAndServe(portString, nil)
}
