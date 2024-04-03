package routes

import (
	"SpeedLoggerz/server/implementation"
	"SpeedLoggerz/server/models"
	"SpeedLoggerz/server/utils"
	"github.com/labstack/echo/v4"
)

func HookupRoutes(e *echo.Echo) {
	g := e.Group("")
	// all route handler funcs
	// dont send creds thru get method lmao
	g.POST("/register/", func(c echo.Context) error {
		// query params
		var user models.User
		err := c.Bind(&user)
		if err != nil {
			utils.HTTPErrorHandler(err, c)
		}
		err, _ = implementation.CreateOrUpdateUser(user)
		if err != nil {
			return err
		}
		return nil
	})
}
