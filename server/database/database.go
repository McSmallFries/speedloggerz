package database

import (
  "github.com/jmoiron/sqlx"
)

/* Only needed for when multiple schemas are in play. */
//type DatabaseConnection struct {
//  Host     string
//  IsMaster bool
//  Db       *sqlx.DB
//}

type Configuration struct {
  User        string
  Password    string
  Connections Connection
}

type Connection struct {
  Host     string
  IsMaster bool // will have a master db schema.
  Db       *sqlx.DB
}

var config = Configuration{
  User:     "root",
  Password: "admin",
}

var connection = Connection{
  Host:     "localhost",
  IsMaster: false,
}

func (d *Connection) OpenDatabase() (err error) {
  d.Db, err = sqlx.Open("mysql", "root:@tcp(localhost:3306)/speedloggerz")
  return err
}

func InitialiseDatabaseConn() (error, Connection) {
  err := connection.OpenDatabase()
  return err, connection
}
