package database

import (
	"backbone/tools"
	"database/sql"
	"github.com/sirupsen/logrus"
)

type Bill struct {
	BillId     string `json:"billid"`
	BillName   string `json:"billname"`
	BillSlot   string `json:"billslot"`
	Author     string `json:"author"`
	Sponsor    string `json:"sponsor"`
	Parliament string `json:"parliament"`
}

var Bills = make(map[int]Bill)

func Billsr(db *sql.DB) map[int]Bill {
	var (
		BillId     string
		BillName   string
		BillSlot   string
		Author     string
		Sponsor    string
		Parliament string
	)
	i := 0
	rows, err := db.Query("SELECT * FROM bill_info;")
	if err == sql.ErrNoRows {
		tools.Log.Error("No Rows were Returned - Bills")
	}
	if err != nil {
		tools.Log.Panic("Could not load SQL Bill Data")
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&BillId, &BillName, &BillSlot, &Author, &Sponsor, &Parliament)
		tools.Log.WithFields(logrus.Fields{
			"BillId":   BillId,
			"BillName": BillName,
			"BillSlot": BillSlot,
			"Author":   Author,
			"Sponsor":  Sponsor,
			"Parl":     Parliament,
		}).Trace("Bills Scanned")
		Bills[i] = Bill{
			BillId, BillName, BillSlot, Author, Sponsor, Parliament,
		}
		i += 1
	}
	return Bills
}
