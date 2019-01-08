package database

import (
	"backbone/tools"
	"github.com/jmoiron/sqlx"
	"time"
)

var prevCheckTime time.Time

//A function that will routinely check if the JSON uploaded is out of data and if so uploads a new copy
func DBUpdating(db *sqlx.DB) {
	prevCheckTime = time.Now()
	//infinite loop
	for {
		//checking if x amount of time has passed
		if prevCheckTime.Add(time.Minute*tools.Conf.GetDuration("updatetime")).Unix() < time.Now().Unix() {
			//updating bills
			Billsr(db)
			tools.Log.WithField("Bill", Bills).Info("Bills Loaded")

			//updating the check time
			prevCheckTime = time.Now()
		}
	}
}

//Will force an update for ALL parameters
func ForceUpdate(db *sqlx.DB) error {
	prevCheckTime = time.Now()

	err := Billsr(db)
	if err != nil {
		return err
	}

	err = Votesr(db)
	if err != nil {
		return err
	}

	err = Partiesr(db)
	if err != nil {
		return err
	}

	return nil
}
