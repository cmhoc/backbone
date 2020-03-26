// basic connection and disconnection functions
package database

import (
	"backbone/tools"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// dbinfo declaration
// to be changed to an elections specific user to only give them access to necessary DB parts
var (
	host     = tools.Conf.GetString("host")
	port     = tools.Conf.GetInt("port")
	user     = tools.Conf.GetString("user")
	password = tools.Conf.GetString("password")
	dbname   = tools.Conf.GetString("dbname")
)

// connection function
func Login() *sqlx.DB {

	// putting everything needed into one string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// opening the sql connection
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		tools.Log.Panic(err)
	}

	tools.Log.Info("Database Logged In")
	return db
}
