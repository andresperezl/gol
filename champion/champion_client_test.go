package champion

import (
	"testing"

	"github.com/andresperezl/gol/auth"
	"github.com/andresperezl/gol/utils"
)

const CIRawData = `{
    "freeChampionIds": [
        8,
        15,
        20,
        38,
        48,
        57,
        78,
        82,
        91,
        96,
        126,
        143,
        161,
        245,
        421
    ],
    "freeChampionIdsForNewPlayers": [
        18,
        81,
        92,
        141,
        37,
        238,
        19,
        45,
        25,
        64
    ],
    "maxNewPlayerLevel": 10
}`

func TestChampionClient(t *testing.T) {
	ts, lc := utils.GetTestServerAndClient(true, CIRawData)
	defer ts.Close()

	c := NewClient(lc.Champion)
	auth.SetAuthInfo("RGAPI")
	if _, err := c.GetChampionInfo(); err != nil {
		t.Errorf("Request GetChampionInfo() failed when it should have succeded")
	}
	auth.SetAuthInfo("")
	if _, err := c.GetChampionInfo(); err == nil {
		t.Errorf("Request GetChampionInfo() needed to fail in this case")
	} else if err.Code() != 401 {
		t.Errorf("This requested should have been %d, got %d", 401, err.Code())
	}
}
