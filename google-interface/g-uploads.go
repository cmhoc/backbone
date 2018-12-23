package google

import (
	"backbone/tools"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Upload struct {
	Timestamp string //Time it was submitted
	Author    string //The Author of the bill
	Sponsor   string //Co-Sponsor if applicable
	Link      string //Link to the bill or motion
	Type      string //Bill or Motion
	Slot      string //The slot it's using (Gov, Specific party, ect)
	Meta      string //metadata in the format of Bot Submission: {{discord name}}
}

func GoogleBillUp(bill Upload) error {
	/* IMPORTANT NOTE READ IF CONSISTENT ERRORS
	The Google API creates Service accounts that act as the account for the bot thats using it,
	therefor if there are consistent errors with the spreadsheet not being found ensure that the service account
	the bot is using actually has access to the sheet.
	Email can be found on the google credentials site.
	*/
	spreadsheet, err := sheet.FetchSpreadsheet(tools.Conf.GetString("legsubsheet"))
	if err != nil {
		return fmt.Errorf("spreadsheet not found")
	}

	workingSheet, err := spreadsheet.SheetByIndex(0) //Uses the first sheet, I only need the first sheet.
	if err != nil {
		return fmt.Errorf("sheet not found")
	}

	rows := len(workingSheet.Rows)
	cols := 7
	//adding an extra row
	sheet.ExpandSheet(workingSheet, uint(rows+1), uint(cols))
	workingSheet.Synchronize() //ensuring we're working with our new expanded sheet

	//adding the data to the new row
	workingSheet.Update(rows, 0, bill.Timestamp)
	workingSheet.Update(rows, 1, bill.Author)
	workingSheet.Update(rows, 2, bill.Sponsor)
	workingSheet.Update(rows, 3, bill.Link)
	workingSheet.Update(rows, 4, bill.Type)
	workingSheet.Update(rows, 5, bill.Slot)
	workingSheet.Update(rows, 6, bill.Meta)

	//a final update for good measure
	workingSheet.Synchronize()

	//log the submission
	tools.Log.WithFields(logrus.Fields{
		"Timestamp": bill.Timestamp,
		"Author":    bill.Author,
		"Sponsor":   bill.Sponsor,
		"Type":      bill.Type,
		"Slot":      bill.Slot,
		"Meta":      bill.Meta,
	}).Info("Bill Submission")

	return nil
}
