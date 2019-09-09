package main

import (
	"fmt"

	golClient "github.com/andresperezl/gol/wrapper"
	"github.com/andresperezl/gol/wrapper/region"
)

func main() {
	// authInfo := client.APIKeyAuth("X-Riot-Token", "header", "RGAPI-fdb89aec-1c66-4f90-945e-d92643e6bf45")
	// tCfg := golClient.DefaultTransportConfig()
	// tCfg.Host = "na1.api.riotgames.com"
	// lolClient := golClient.NewHTTPClientWithConfig(nil, tCfg)
	// cm := lolClient.ChampionMastery
	// params := champion_mastery.NewGetChampionMasteryScoreParams()
	// params.EncryptedSummonerID = "zdDFROpUz2Zt6ynxXuHUv4oMmuKRejXmALAwt79nI2sXp3w"
	// res, err := cm.GetChampionMasteryScore(params, authInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Score: ", res.Payload)
	client := golClient.NewClient("RGAPI-fdb89aec-1c66-4f90-945e-d92643e6bf45", region.NA1)
	fmt.Printf("%+v\n", client.ChampionMastery.GetChampionMastery("zdDFROpUz2Zt6ynxXuHUv4oMmuKRejXmALAwt79nI2sXp3w", 12))
}
