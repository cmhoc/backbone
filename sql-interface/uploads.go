package database

import "database/sql"

func Billsu(db *sql.DB, bill Bill) {
	statement := `INSERT INTO bill_info (bill_number, bill_name, bill_slot, bill_author, bill_sponsor, parliament)
				  VALUES ($1, $2, $3, $4, $5, $6)`
	db.Exec(statement, bill.BillId, bill.BillName, bill.BillSlot, bill.Author, bill.Sponsor, bill.Parliament)
}
