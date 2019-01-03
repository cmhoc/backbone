package database

import (
	"backbone/tools"
	"database/sql"
	"time"
)

//A function that will routinely check if the JSON uploaded is out of data and if so uploads a new copy
func DBUpdating(db *sql.DB){
	var prevCheckTime = time.Now()
	//infinite loop
	for {
		//checking if x amount of time has passed
		if prevCheckTime.Add(time.Minute * tools.Conf.GetDuration("updatetime")).Unix() < time.Now().Unix() {
			//updating bills
			Billsr(db)
			tools.Log.WithField("Bill", Bills).Info("Bills Loaded")

			//updating the check time
			prevCheckTime = time.Now()
		}
	}
}