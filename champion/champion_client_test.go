package champion

import (
	"os"
	"testing"

	"github.com/andresperezl/gol/auth"
	"github.com/andresperezl/gol/utils"
)

func TestChampionClient(t *testing.T) {
	host := os.Getenv("RIOT_API_TEST_HOST")
	if len(host) == 0 {
		host = "127.0.0.1:8443"
	}
	lc := utils.GetTestAPIClient(host)
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
