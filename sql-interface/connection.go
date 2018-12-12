// basic connection and disconnection functions

package database

import (
	"database/sql"
	"discordbot/logging"
	"fmt"
	_ "github.com/lib/pq"
)

// dbinfo declaration
// to be changed to an elections specific user to only give them access to nessecary DB parts
const (
	host     = ""
	port     =
	user     = ""
	password = ""
	dbname   = ""
)

// connection function
func Login() *sql.DB {

	// putting everything needed into one string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// opening the sql connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Log.Panic(err)
	}
	logger.Log.Info("Database Logged In")
	return db
}
