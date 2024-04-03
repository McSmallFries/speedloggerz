package models

import "github.com/jmoiron/sqlx"

type User struct {
	IdUser   int64  `json:"IdUser" db:"UserId"`
	Username string `json:"Username" db:"Username"`
	Password string `json:"Password" db:"Password"`
	// store hashed in database..
	// AccountCode string;
	// LastPasswordID: number;
}

func (u *User) Insert(db *sqlx.Tx) (id int64, err error) {
	query := `INSERT INTO users_test (Username, Password)
    VALUES (:Username, :Password);`

	defer db.Rollback()
	stmt, err := db.PrepareNamed(query)
	if err != nil {

	}
	defer stmt.Close()

	res, err := stmt.Exec(u)
	if err != nil {

	}
	u.IdUser, err = res.LastInsertId()
	return u.IdUser, err
}
