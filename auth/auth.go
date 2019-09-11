package auth

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

const RiotAPIHeader = "X-Riot-Token"

var authInfo runtime.ClientAuthInfoWriter

func GetAuthInfo() runtime.ClientAuthInfoWriter {
	return authInfo
}

func SetAuthInfo(apiKey string) {
	authInfo = client.APIKeyAuth(RiotAPIHeader, "header", apiKey)
}
