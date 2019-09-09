package champion

import (
	"github.com/andresperezl/gol/lolapi/models"
	"github.com/andresperezl/gol/test/restapi/operations/champion"
	"github.com/go-openapi/runtime/middleware"
)

const RawData = `{
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

func GetChampionInfoHandlerFunc() champion.GetChampionInfoHandlerFunc {
	return champion.GetChampionInfoHandlerFunc(func(params champion.GetChampionInfoParams, principal interface{}) middleware.Responder {
		res := champion.NewGetChampionInfoOK()
		ci := &models.ChampionInfo{}
		ci.UnmarshalBinary([]byte(RawData))
		res.SetPayload(ci)
		return res
	})
}
