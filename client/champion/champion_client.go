package champion

import (
	"github.com/andresperezl/gol/lolapi/client/champion"
	"github.com/andresperezl/gol/lolapi/models"
	"github.com/go-openapi/runtime"
)

type Client struct {
	cClient  *champion.Client
	authInfo runtime.ClientAuthInfoWriter
}

func NewClient(cmClient *champion.Client) *Client {
	golCmClient := &Client{
		cClient: cmClient,
	}
	return golCmClient
}

func (c *Client) SetAuthInfo(authInfo runtime.ClientAuthInfoWriter) {
	c.authInfo = authInfo
}

func (c *Client) GetChampionInfo() *models.ChampionInfo {
	res, err := c.cClient.GetChampionInfo(nil, c.authInfo)
	if err != nil {
		panic(err)
	}
	return res.Payload
}
