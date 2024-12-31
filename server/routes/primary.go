package routes

import (
	"SpeedLoggerz/server/implementation"
	"SpeedLoggerz/server/models"
	"SpeedLoggerz/server/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func HookupRoutes(e *echo.Echo) {
	web := e.Group("/web")
	clientProcess := e.Group("/clientretriever")
	// all route handler funcs
	// dont send creds thru get method lmao
	web.POST("/register", func(context echo.Context) error {
		// query params
		var user models.User
		err := context.Bind(&user)
		if err != nil {
			utils.HTTPErrorHandler(err, context)
		}
		err, id := implementation.CreateOrUpdateUser(user)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, "")
		}
		return context.JSON(http.StatusOK, strconv.FormatInt(id, 10))
	})

	clientProcess.POST("/sync", func(context echo.Context) error {
		logFiles := make([]models.SpeedyLogFile, 0)
		err := context.Bind(&logFiles)
		if err != nil {
			utils.HTTPErrorHandler(err, context)
		}
		id, err := implementation.SyncLogFiles(logFiles)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, "")
		}
		return context.JSON(http.StatusOK, strconv.FormatInt(id, 10))
	})
}
