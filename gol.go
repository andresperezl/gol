package gol

import (
	"github.com/andresperezl/gol/auth"
	"github.com/andresperezl/gol/champion"
	"github.com/andresperezl/gol/championmastery"
	"github.com/andresperezl/gol/league"
	"github.com/andresperezl/gol/region"
	"github.com/andresperezl/gol/utils"
)

// GoLClient represents the League of Legends API Client
type GoLClient struct {
	Champion        *champion.Client
	ChampionMastery *championmastery.Client
	League          *league.Client
	CurrentRegion   region.Region
	CurrentAPIKey   string
}

// NewClient create a new GolClient authenticated with the specified apiKey,
// and the region endpoint which the request will be sent
func NewClient(apiKey string, region region.Region) *GoLClient {
	golClient := &GoLClient{}
	golClient.SetRegion(region)
	golClient.SetAPIKey(apiKey)
	return golClient
}

// SetRegion changes the region to which this clients will send the request to
func (c *GoLClient) SetRegion(region region.Region) {
	lolClient := utils.GetInternalAPIClient(region)
	c.CurrentRegion = region
	c.Champion = champion.NewClient(lolClient.Champion)
	c.ChampionMastery = championmastery.NewClient(lolClient.ChampionMastery)
	c.League = league.NewClient(lolClient.League)
}

// SetAPIKey changes the key used for authentication with apiKey
func (c *GoLClient) SetAPIKey(apiKey string) {
	auth.SetAuthInfo(apiKey)
	c.CurrentAPIKey = apiKey
}
