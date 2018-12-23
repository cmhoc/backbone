package reddit

import (
	"backbone/tools"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Count(https string) (map[string]map[string]int, []string, error) {
	https = https + ".json"

	tempClient := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest(http.MethodGet, https, nil)
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("User-Agent", tools.Conf.GetString("redditagent"))
	response, err := tempClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	var thread []thread
	err = json.Unmarshal(body, &thread)
	if err != nil {
		return nil, nil, err
	}

	//for holding yea or nay votes
	votes := make(map[string]map[string]int)

	//getting individual bill titles of the vote
	var (
		billtitle string
		bills     []string
	)

	//creating a slice with the bill titles in it using the standardized format of the post titles
	temp1 := strings.Split(thread[0].TopData.Children[0].Data.Title, "| ")
	bills = strings.Split(temp1[len(temp1)-1], ", ")

	for i := 0; i < len(thread[1].TopData.Children); i++ {
		//if its the automod whip go to the next
		if thread[1].TopData.Children[i].Data.Author == "AutoModerator" {
			continue
		}
		var (
			temp    int
			replies []string
			body    string
		)

		replies = strings.Split(thread[1].TopData.Children[i].Data.CommentBody, "\n")

		//Some people do two spaces between replies
		for y := 0; y < len(replies); y++ {
			if replies[y] == "" {
				if y+1 < len(replies) {
					replies[y] = replies[y+1]
				}
			}
		}

		for x := 0; x < len(bills); x++ {
			billtitle = bills[x]
			if votes[billtitle] == nil {
				votes[billtitle] = make(map[string]int)
			}
			billtitle = bills[x]
			if strings.HasPrefix(replies[x], billtitle+": ") {
				body = strings.ToLower(strings.TrimPrefix(replies[x], billtitle+": "))
			} else { // if its an amendment
				temp1 = strings.Split(replies[x], ": ")
				body = strings.ToLower(temp1[len(temp1)-1])
			}
			if body == "yea" || body == "oui" {
				temp = 0
			} else if body == "nay" || body == "non" {
				temp = 1
			} else if body == "abstain" || body == "abstention" {
				temp = 2
			} else {
				temp = 3
			}
			votes[billtitle][strings.ToLower(thread[1].TopData.Children[i].Data.Author)] = temp
		}
	}
	return votes, bills, nil
}
