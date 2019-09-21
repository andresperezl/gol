package league

import (
	"github.com/andresperezl/gol/auth"
	"github.com/andresperezl/gol/errors"
	"github.com/andresperezl/gol/lolapi/client/league"
	"github.com/andresperezl/gol/lolapi/models"
)

// Client holds the internal lolapi League client
type Client struct {
	lCLient *league.Client
}

// Creates a new client from the lolapi League client
func NewClient(cmClient *league.Client) *Client {
	golCmClient := &Client{
		lCLient: cmClient,
	}
	return golCmClient
}

// GetLeagueEntriesExp is the same as GetLeagueEntries with the addition of the Apex Tiers (Challenger, Grandmaster, and Master)
func (c *Client) GetLeagueEntriesExp(queue Queue, tier TierExp, division Division, page *int32) ([]*models.LeagueEntry, *errors.APIError) {

	params := league.NewGetLeagueEntriesExpParams().WithQueue(queue.String()).WithTier(tier.String()).WithDivision(division.String()).WithPage(page)
	res, err := c.lCLient.GetLeagueEntriesExp(params, auth.GetAuthInfo())
	if err != nil {
		er := err.(errors.ErrorResponse)
		return nil, &errors.APIError{APIError: *er.GetPayload()}
	}
	return res.Payload, nil
}
