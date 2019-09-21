package championmastery

import (
	"github.com/andresperezl/gol/auth"
	"github.com/andresperezl/gol/errors"
	"github.com/andresperezl/gol/lolapi/client/champion_mastery"
	"github.com/andresperezl/gol/lolapi/models"
)

type Client struct {
	cmClient *champion_mastery.Client
}

func NewClient(cmClient *champion_mastery.Client) *Client {
	golCmClient := &Client{
		cmClient: cmClient,
	}
	return golCmClient
}

func (c *Client) GetChampionMasteryScore(encryptedSummonerID string) (int32, *errors.APIError) {
	params := champion_mastery.NewGetChampionMasteryScoreParams().WithEncryptedSummonerID(encryptedSummonerID)
	res, err := c.cmClient.GetChampionMasteryScore(params, auth.GetAuthInfo())
	if err != nil {
		er := err.(errors.ErrorResponse)
		return -1, &errors.APIError{APIError: *er.GetPayload()}
	}
	return res.Payload, nil
}

func (c *Client) GetAllChampionMasteries(encryptedSummonerID string) ([]*models.ChampionMastery, *errors.APIError) {
	params := champion_mastery.NewGetAllChampionMasteriesParams().WithEncryptedSummonerID(encryptedSummonerID)
	res, err := c.cmClient.GetAllChampionMasteries(params, auth.GetAuthInfo())
	if err != nil {
		er := err.(errors.ErrorResponse)
		return nil, &errors.APIError{APIError: *er.GetPayload()}
	}
	return res.Payload, nil
}

func (c *Client) GetChampionMastery(encryptedSummonerID string, championID int64) (*models.ChampionMastery, *errors.APIError) {
	params := champion_mastery.NewGetChampionMasteryParams().WithEncryptedSummonerID(encryptedSummonerID).WithChampionID(championID)
	res, err := c.cmClient.GetChampionMastery(params, auth.GetAuthInfo())
	if err != nil {
		er := err.(errors.ErrorResponse)
		return nil, &errors.APIError{APIError: *er.GetPayload()}
	}
	return res.Payload, nil
}
