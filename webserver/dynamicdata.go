package webserver

type vote struct {
	BillId string
	BillName string
	Author string
}

var Votestest = []vote{
	{"M-1", "An Act to Test a System", "thehowlinggreywolf"},
}
