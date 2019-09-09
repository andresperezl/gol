package champion_mastery

import (
	"github.com/andresperezl/gol/lolapi/client/champion_mastery"
	"github.com/andresperezl/gol/lolapi/models"
	"github.com/go-openapi/runtime"
)

type Client struct {
	cmClient *champion_mastery.Client
	authInfo runtime.ClientAuthInfoWriter
}

func NewClient(cmClient *champion_mastery.Client) *Client {
	golCmClient := &Client{
		cmClient: cmClient,
	}
	return golCmClient
}

func (c *Client) SetAuthInfo(authInfo runtime.ClientAuthInfoWriter) {
	c.authInfo = authInfo
}

func (c *Client) GetChampionMasteryScore(encryptedSummonerID string) int32 {
	params := champion_mastery.NewGetChampionMasteryScoreParams().WithEncryptedSummonerID(encryptedSummonerID)
	res, err := c.cmClient.GetChampionMasteryScore(params, c.authInfo)
	if err != nil {
		panic(err)
	}
	return res.Payload
}

func (c *Client) GetAllChampionMasteries(encryptedSummonerID string) models.ListChampionMastery {
	params := champion_mastery.NewGetAllChampionMasteriesParams().WithEncryptedSummonerID(encryptedSummonerID)
	res, err := c.cmClient.GetAllChampionMasteries(params, c.authInfo)
	if err != nil {
		panic(err)
	}
	return res.Payload
}

func (c *Client) GetChampionMastery(encryptedSummonerID string, championID int64) *models.ChampionMastery {
	params := champion_mastery.NewGetChampionMasteryParams().WithEncryptedSummonerID(encryptedSummonerID).WithChampionID(championID)
	res, err := c.cmClient.GetChampionMastery(params, c.authInfo)
	if err != nil {
		panic(err)
	}
	return res.Payload
}
