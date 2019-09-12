package champion_mastery

import (
	"testing"

	"github.com/andresperezl/gol/auth"
	"github.com/andresperezl/gol/utils"
)

const (
	masteryScore = "237"

	allCMRawData = `[
      {
          "championLevel": 6,
          "chestGranted": true,
          "championPoints": 75378,
          "championPointsSinceLastLevel": 53778,
          "championPointsUntilNextLevel": 0,
          "summonerId": "test",
          "tokensEarned": 1,
          "championId": 12,
          "lastPlayTime": 1567724912000
      },
      {
          "championLevel": 5,
          "chestGranted": false,
          "championPoints": 45702,
          "championPointsSinceLastLevel": 24102,
          "championPointsUntilNextLevel": 0,
          "summonerId": "test",
          "tokensEarned": 0,
          "championId": 31,
          "lastPlayTime": 1562631217000
      }
  ]`

	cmRawData = `{
      "championLevel": 6,
      "chestGranted": true,
      "championPoints": 75378,
      "championPointsSinceLastLevel": 53778,
      "championPointsUntilNextLevel": 0,
      "summonerId": "test",
      "tokensEarned": 1,
      "championId": 12,
      "lastPlayTime": 1567724912000
  }`

	summonerID = "test"
	championID = 12
)

func TestMasteryScore(t *testing.T) {
	ts, lc := utils.GetTestServerAndClient(true, masteryScore)
	defer ts.Close()

	c := NewClient(lc.ChampionMastery)
	auth.SetAuthInfo("RGAPI")
	if _, err := c.GetChampionMasteryScore("test"); err != nil {
		t.Errorf("Request GetChampionMasteryScore() failed when it should have succeded")
	}
	auth.SetAuthInfo("")
	if _, err := c.GetChampionMasteryScore("test"); err == nil {
		t.Errorf("Request GetChampionMasteryScore() needed to fail in this case")
	} else if err.Code() != 401 {
		t.Errorf("This requested should have been %d, got %d", 401, err.Code())
	}
}

func TestAllChampionMasteries(t *testing.T) {
	ts, lc := utils.GetTestServerAndClient(true, allCMRawData)
	defer ts.Close()

	c := NewClient(lc.ChampionMastery)
	auth.SetAuthInfo("RGAPI")
	if r, err := c.GetAllChampionMasteries("test"); err != nil {
		t.Errorf("Request GetAllChampionMasteries() failed when it should have succeded")
	} else {
		if len(r) != 2 {
			t.Errorf("This should have returned 2 items")
		}
		for _, e := range r {
			if e.SummonerID != "test" {
				t.Errorf("All the elements in the response should contain the %s summonerID", summonerID)
				break
			}
		}
	}
	auth.SetAuthInfo("")
	if _, err := c.GetAllChampionMasteries("test"); err == nil {
		t.Errorf("Request GetAllChampionMasteries() needed to fail in this case")
	} else if err.Code() != 401 {
		t.Errorf("This requested should have been %d, got %d", 401, err.Code())
	}
}

func TestChampionMastery(t *testing.T) {
	ts, lc := utils.GetTestServerAndClient(true, cmRawData)
	defer ts.Close()

	c := NewClient(lc.ChampionMastery)
	auth.SetAuthInfo("RGAPI")
	if r, err := c.GetChampionMastery(summonerID, championID); err != nil {
		t.Errorf("Request GetChampionMastery() failed when it should have succeded")
	} else {
		if r.SummonerID != summonerID {
			t.Errorf("The response should contain the %s summonerID", summonerID)
		}
		if r.ChampionID != championID {
			t.Errorf("The response should contain the %d championID", championID)
		}

	}
	auth.SetAuthInfo("")
	if _, err := c.GetChampionMastery(summonerID, championID); err == nil {
		t.Errorf("Request GetChampionMastery() needed to fail in this case")
	} else if err.Code() != 401 {
		t.Errorf("This requested should have been %d, got %d", 401, err.Code())
	}
}
