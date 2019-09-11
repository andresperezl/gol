package region

import (
	"fmt"
	"strings"
)

const BaseEndpoint = "api.riotgames.com"

// Region represents one of the supported Riot Regions
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

// String returns the string representation of the region
func (region Region) String() string {
	// The other of this should be the same as the constants above
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

// GetRegionalEndpoint returns the endpoint of this region
func (region Region) GetRegionalEndpoint() string {
	return fmt.Sprintf("%s.%s", strings.ToLower(region.String()), BaseEndpoint)
}

// ServiceProxy returns true is this region is part of the Service Proxies
func (region Region) ServiceProxy() bool {
	switch region {
	case Americas, Europe, Asia:
		return false
	default:
		return true
	}
}

// RegionalProxy returns true is this region is part of the Regional Proxies
func (region Region) RegionalProxy() bool {
	switch region {
	case Americas, Europe, Asia:
		return true
	default:
		return false
	}
}
