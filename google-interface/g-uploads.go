/*
IMPORTANT NOTE READ IF CONSISTENT ERRORS
The Google API creates Service accounts that act as the account for the bot thats using it,
therefor if there are consistent errors with the spreadsheet not being found ensure that the service account
the bot is using actually has access to the sheet.
Email can be found on the google credentials site.
*/
package google

import (
	"backbone/tools"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type Upload struct {
	Timestamp string //Time it was submitted
	Author    string //The Author of the bill
	Sponsor   string //Co-Sponsor if applicable
	Link      string //Link to the bill or motion
	Type      string //Bill or Motion
	Slot      string //The slot it's using (Gov, Specific party, ect)
	Meta      string //metadata in the format of Bot Submission: {{.discord name}}
}

type placement struct {
	member string //the member
	row    int    //the row on the sheet
}

func GoogleBillUp(bill Upload) error {
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
	err = sheet.ExpandSheet(workingSheet, uint(rows+1), uint(cols))
	if err != nil {
		return fmt.Errorf("sheet could not be expanded")
	}
	err = workingSheet.Synchronize() //ensuring we're working with our new expanded sheet
	if err != nil {
		return fmt.Errorf("sheet could not be synced")
	}

	//adding the data to the new row
	workingSheet.Update(rows, 0, bill.Timestamp)
	workingSheet.Update(rows, 1, bill.Author)
	workingSheet.Update(rows, 2, bill.Sponsor)
	workingSheet.Update(rows, 3, bill.Link)
	workingSheet.Update(rows, 4, bill.Type)
	workingSheet.Update(rows, 5, bill.Slot)
	workingSheet.Update(rows, 6, bill.Meta)

	//a final update for good measure
	err = workingSheet.Synchronize()
	if err != nil {
		return fmt.Errorf("sheet could not be synced")
	}

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

func GoogleVotesUp(votes map[string]map[string]int, billtitles []string) error {
	var (
		place []placement
		col   int
	)

	spreadsheet, err := sheet.FetchSpreadsheet(tools.Conf.GetString("mainsheet"))
	if err != nil {
		return fmt.Errorf("spreadsheet not found")
	}

	workingSheet, err := spreadsheet.SheetByIndex(1) //Uses the second sheet
	if err != nil {
		return fmt.Errorf("sheet not found")
	}

	//finding the first col to input to
	for i := 0; i < len(workingSheet.Columns); i++ {
		if workingSheet.Columns[i][1].Value == "" {
			col = i
			break
		}
	}

	//getting the placement of all the members on the sheet
	for i := 2; i < len(workingSheet.Rows)-5; i++ {
		place = append(place, placement{strings.ToLower(workingSheet.Rows[i][2].Value), i})
	}

	for y := 0; y < len(billtitles); y++ {
		//getting the all the voters
		var voters []string
		for k := range votes[billtitles[y]] {
			voters = append(voters, k)
		}

		//Putting the bill title on the top line
		workingSheet.Update(1, col, billtitles[y])

		//Finding the Voters and appending results
		for j := 0; j < len(place); j++ {
			for i := 0; i < len(voters); i++ {
				if place[j].member == voters[i] {
					if votes[billtitles[y]][voters[i]] == 0 {
						//yea vote
						workingSheet.Update(place[j].row, col, "Y // O")
					} else if votes[billtitles[y]][voters[i]] == 1 {
						//nay vote
						workingSheet.Update(place[j].row, col, "N // N")
					} else if votes[billtitles[y]][voters[i]] == 2 {
						//abstention
						workingSheet.Update(place[j].row, col, "A // A")
					} else if votes[billtitles[y]][voters[i]] == 3 {
						//Case when the bot cant determine the vote
						workingSheet.Update(place[j].row, col, "Err")
					}
				}
			}
		}
		err = workingSheet.Synchronize() //updating the spreadsheet
		if err != nil {
			return fmt.Errorf("sheet could not be synced")
		}
		col++                      //moving to the next column for the next bill
	}

	return nil
}
