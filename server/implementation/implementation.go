package implementation

import (
	"SpeedLoggerz/server/database"
	"SpeedLoggerz/server/models"
	_ "github.com/go-sql-driver/mysql"
)

func LoadAppGlobals() error {

	return nil
}

func CreateOrUpdateUser(user models.User) (error, int64) {
	connection := database.Config.Connection
	tx := connection.Db.MustBegin()
	id, err := user.Insert(tx)
	if err != nil {
		return err, 0
	}
	return err, id
}
