package gol

import (
	"github.com/andresperezl/gol/client/champion"
	"github.com/andresperezl/gol/client/champion_mastery"
	"github.com/andresperezl/gol/client/region"
	lolclient "github.com/andresperezl/gol/lolapi/client"
	"github.com/go-openapi/runtime"
	goapiclient "github.com/go-openapi/runtime/client"
)

type GoLClient struct {
	lolClient       *lolclient.LeagueOfLegends
	authInfo        runtime.ClientAuthInfoWriter
	Champion        *champion.Client
	ChampionMastery *champion_mastery.Client
}

const (
	RiotAPIHeader = "X-Riot-Token"
)

func NewClient(apiKey string, region region.Region) *GoLClient {
	golClient := &GoLClient{}
	golClient.SetRegion(region)
	golClient.SetApiKey(apiKey)
	return golClient
}

func (c *GoLClient) SetRegion(region region.Region) {
	tCfg := lolclient.DefaultTransportConfig()
	tCfg.Host = region.GetRegionalEndpoint()
	c.lolClient = lolclient.NewHTTPClientWithConfig(nil, tCfg)
	c.Champion = champion.NewClient(c.lolClient.Champion)
	c.ChampionMastery = champion_mastery.NewClient(c.lolClient.ChampionMastery)
}

func (c *GoLClient) SetApiKey(apiKey string) {
	c.authInfo = goapiclient.APIKeyAuth(RiotAPIHeader, "header", apiKey)
	c.Champion.SetAuthInfo(c.authInfo)
	c.ChampionMastery.SetAuthInfo(c.authInfo)
}
