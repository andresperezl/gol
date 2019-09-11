package champion

import (
	"github.com/andresperezl/gol/auth"
	"github.com/andresperezl/gol/errors"
	"github.com/andresperezl/gol/lolapi/client/champion"
	"github.com/andresperezl/gol/lolapi/models"
)

type Client struct {
	cClient *champion.Client
}

func NewClient(cmClient *champion.Client) *Client {
	golCmClient := &Client{
		cClient: cmClient,
	}
	return golCmClient
}

func (c *Client) GetChampionInfo() (*models.ChampionInfo, *errors.APIError) {
	res, err := c.cClient.GetChampionInfo(nil, auth.GetAuthInfo())
	if err != nil {
		er := err.(errors.ErrorResponse)
		return nil, &errors.APIError{*er.GetPayload()}
	}
	return res.Payload, nil
}
