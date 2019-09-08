package region

import (
	"fmt"
	"strings"
)

const BASE_ENDPOINT = "api.riotgames.com"

type Region int

const (
	BR1 Region = iota
	EUN1
	EUW1
	JP1
	KR
	LA1
	LA2
	NA
	NA1
	OC1
	TR
	RU
	PBE
	Americas
	Europe
	Asia
)

func (region Region) GetRegionalEndpoint() string {
	return fmt.Sprintf("%s.%s", strings.ToLower(region.String()), BASE_ENDPOINT)
}

func (region Region) String() string {
	ep := [...]string{
		"BR1",
		"EUN1",
		"EUW1",
		"JP1",
		"KR",
		"LA1",
		"LA2",
		"NA",
		"NA1",
		"OC1",
		"TR1",
		"RU",
		"PBE1",
		"Americas",
		"Europe",
		"Asia",
	}
	return ep[region]
}

func (region Region) ServiceProxy() bool {
	switch region {
	case Americas, Europe, Asia:
		return false
	default:
		return true
	}
}

func (region Region) RegionalProxy() bool {
	switch region {
	case Americas, Europe, Asia:
		return true
	default:
		return false
	}
}
