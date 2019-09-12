package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/andresperezl/gol/auth"
	lolapi "github.com/andresperezl/gol/lolapi/client"
	"github.com/go-openapi/runtime/client"
)

func GetTestServerAndClient(authenticate bool, response string) (*httptest.Server, *lolapi.LeagueOfLegends) {
	ts := httptest.NewTLSServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			if authenticate {
				if riotToken := r.Header.Get(auth.RiotAPIHeader); len(riotToken) > 0 && strings.HasPrefix(riotToken, "RGAPI") {
					w.Write([]byte(response))
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte(`{ "status": {"message": "Unauthorized", "status_code": 401 } }`))
				}
			}
		}),
	)

	tc := ts.Client()
	serverURL, err := url.Parse(ts.URL)
	if err != nil {
		panic(err)
	}

	rtClient := client.NewWithClient(fmt.Sprintf("%s:%s", serverURL.Hostname(), serverURL.Port()), lolapi.DefaultBasePath, []string{"https"}, tc)

	return ts, lolapi.New(rtClient, nil)
}
