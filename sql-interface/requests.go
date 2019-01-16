package database

import (
	"backbone/tools"
	"github.com/jmoiron/sqlx"
)

type (
	Bill struct {
		BillId     string `json:"billid" db:"bill_number"`
		BillName   string `json:"billname" db:"bill_name"`
		BillSlot   string `json:"billslot" db:"bill_slot"`
		Author     string `json:"author" db:"bill_author"`
		Sponsor    string `json:"sponsor" db:"bill_sponsor"`
		Parliament int    `json:"parliament" db:"parliament"`
		Link       string `json:"link" db:"link"`
	}

	Party struct {
		NameEnglish     string `json:"nameenglish" db:"name_english"`
		NameFrench      string `json:"namefrench" db:"name_french"`
		ShortEnglish    string `json:"shortenglish" db:"short_english"`
		ShortFrench     string `json:"shortfrench" db:"short_french"`
		DescriptEnglish string `json:"descriptenglish" db:"descript_english"`
		DescriptFrench  string `json:"descriptfrench" db:"descript_french"`
		HexCode         string `json:"hexcode" db:"hex_colour"`
		Leader          string `json:"leader" db:"leader"`
		DeputyLeader    string `json:"deputyleader" db:"deputy_leader"`
	}

	Vote struct {
		Member     string `json:"member" db:"member"`
		Bill       string `json:"bill" db:"bill_number"`
		Vote       string `json:"vote" db:"vote"`
		Parliament int    `json:"parliament" db:"parliament"`
	}

)

var (
	Bills   []Bill
	Parties []Party
	Votes   []Vote
)

func Billsr(db *sqlx.DB) error {
	Bills = nil

	err := db.Select(&Bills, "SELECT * FROM bill_info")
	if err != nil {
		return err
	}

	tools.Log.WithField("#", len(Bills)).Trace("Bills Loaded")

	return nil
}

func Partiesr(db *sqlx.DB) error {
	Parties = nil

	err := db.Select(&Parties, "SELECT * FROM parties")
	if err != nil {
		return err
	}

	tools.Log.WithField("#", len(Parties)).Trace("Parties Loaded")

	return nil
}

func Votesr(db *sqlx.DB) error {
	Votes = nil

	err := db.Select(&Votes, "SELECT * FROM votes")
	if err != nil {
		return err
	}

	tools.Log.WithField("#", len(Votes)).Trace("Votes Loaded")

	return nil
}

func AggrVotes() (map[string]map[string]string, error) {

	votes := make(map[string]map[string]string)

	for i := 0; i < len(Votes); i++ {
		//making the map if nil
		if votes[Votes[i].Member] == nil {
			votes[Votes[i].Member] = make(map[string]string)
		}

		votes[Votes[i].Member][Votes[i].Bill] = Votes[i].Vote

	}

	return votes, nil
}