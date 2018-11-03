// basic connection and disconnection functions

package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)


// dbinfo declaration
// to be changed to an elections specific user to only give them access to nessecary DB parts
const (
	
)

// connection function
func Login() bool {

	// putting everything needed into one string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// opening the sql connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()

	// error check and returning if it connected
	if err != nil {
		return false
	} else {
		return true
	}
}
