package region

import (
	"fmt"
	"strings"
	"testing"
)

func TestRegion(t *testing.T) {
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
	rs := [...]Region{
		BR1,
		EUN1,
		EUW1,
		JP1,
		KR,
		LA1,
		LA2,
		NA,
		NA1,
		OC1,
		TR,
		RU,
		PBE,
		Americas,
		Europe,
		Asia,
	}
	for _, r := range rs {
		if got := r.String(); got != ep[r] {
			t.Errorf("Expected region string to be %s, got %s", ep[r], got)
		}
	}
	for _, r := range rs {
		exp := strings.ToLower(r.String())
		if got := r.GetRegionalEndpoint(); !strings.HasSuffix(got, BaseEndpoint) && !strings.HasPrefix(got, exp) {
			t.Errorf("Expected region endpoint to be %s, got %s", fmt.Sprintf("%s.%s", exp, BaseEndpoint), got)
		}
	}

	for _, r := range rs {
		got := r.ServiceProxy()
		switch r {
		case Americas, Europe, Asia:
			if got {
				t.Errorf("%s should not be marked as a Service Proxy", r.String())
			}
		default:
			if !got {
				t.Errorf("%s should be marked as a Service Proxy", r.String())
			}
		}
	}

	for _, r := range rs {
		got := r.RegionalProxy()
		switch r {
		case Americas, Europe, Asia:
			if !got {
				t.Errorf("%s should not be marked as a Regional Proxy", r.String())
			}
		default:
			if got {
				t.Errorf("%s should be marked as a Regional Proxy", r.String())
			}
		}
	}
}
