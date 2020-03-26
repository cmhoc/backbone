package reddit

type thread struct {
	Kind    string  `json:"kind"`
	TopData topdata `json:"data"`
}

type topdata struct {
	Modhash  string     `json:"modhash"`
	Dist     int        `json:"dist"`
	Children []children `json:"children"`
	After    string     `json:"after"`
	Before   string     `json:"before"`
}

type children struct {
	Kind string `json:"kind"`
	Data data   `json:"data"`
}

type data struct {
	ApprovedTime int      `json:"approved_at_utc"`
	Subreddit    string   `json:"subreddit"`
	BodyText     string   `json:"selftext"`
	Reports      []string `json:"user_reports"`
	Title        string   `json:"title"`
	Author       string   `json:"author"`
	Id           string   `json:"id"`
	CommentBody  string   `json:"body"`
	//TODO: Put in all Reddit JSON Data for future use.
}
