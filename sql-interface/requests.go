package database

import (
	"database/sql"
	"discordbot/logging"
	"discordbot/webserver"
	"github.com/sirupsen/logrus"
)

var (
	BillId []string
	BillName []string
	BillSlot []string
	Author []string
	Sponsor []string
	Parliament []string
	bills []webserver.Bill
)

func Billsr(db *sql.DB) []webserver.Bill {
	rows, err := db.Query("SELECT * FROM bill_info;")
	if err != nil {logger.Log.Panic("Could not load SQL Bill Data")}
	defer rows.Close()
	rows.Scan(&BillId, &BillName, &BillSlot, &Author, &Sponsor, &Parliament)
	logger.Log.WithFields(logrus.Fields{
		"BillId": BillId,
		"BillName": BillName,
		"BillSlot": BillSlot,
		"Author": Author,
		"Sponsor": Sponsor,
		"Parl": Parliament,
	}).Trace("Bills Scanned")
	for i := 0; i < len(BillId); i++ {
		bills = []webserver.Bill{
			{BillId: BillId[i],
			BillName: BillName[i],
			BillSlot: BillSlot[i],
			Author: Author[i],
			Sponsor: Sponsor[i],
			Parliament: Parliament[i]}}
	}
	return bills
}
