package league

import (
	"testing"

	"github.com/andresperezl/gol/auth"
	"github.com/andresperezl/gol/utils"
)

const leagueEntriesRawData = `[
    {
        "queueType": "RANKED_SOLO_5x5",
        "summonerName": "Santorin",
        "hotStreak": false,
        "wins": 372,
        "veteran": true,
        "losses": 250,
        "rank": "I",
        "tier": "CHALLENGER",
        "inactive": false,
        "freshBlood": false,
        "leagueId": "974b70e3-28eb-3b60-9e9f-82a8efa19f10",
        "summonerId": "71OBtFCJIIj6-1CUa-gLGyEWvG8bHTiBWXKY65R8dK1vqy4",
        "leaguePoints": 1574
    },
    {
        "queueType": "RANKED_SOLO_5x5",
        "summonerName": "TSM BB",
        "hotStreak": false,
        "wins": 485,
        "veteran": true,
        "losses": 380,
        "rank": "I",
        "tier": "CHALLENGER",
        "inactive": false,
        "freshBlood": false,
        "leagueId": "974b70e3-28eb-3b60-9e9f-82a8efa19f10",
        "summonerId": "2nthUpJ8yOR89GXvhcMJiWwl0WwRZhCRlaynWtmjuP1dF7tb",
        "leaguePoints": 1432
	}
]`

func TestLeagueExp(t *testing.T) {
	ts, lc := utils.GetTestServerAndClient(true, leagueEntriesRawData)
	defer ts.Close()
	queue := RankedSolo5x5
	tier := ExpChallenger
	division := I
	c := NewClient(lc.League)
	auth.SetAuthInfo("RGAPI")
	if r, err := c.GetLeagueEntriesExp(queue, tier, division, nil); err != nil {
		t.Errorf("Request GetChampionInfo() failed when it should have succeded: %s", err.Error())
	} else if len(r) != 2 {
		t.Errorf("The response should contain only 2 League Entries")
	}
	auth.SetAuthInfo("")
	if _, err := c.GetLeagueEntriesExp(queue, tier, division, nil); err == nil {
		t.Errorf("Request GetLeagueEntriesExp() needed to fail in this case")
	} else if err.Code() != 401 {
		t.Errorf("This requested should have been %d, got %d", 401, err.Code())
	}
}
