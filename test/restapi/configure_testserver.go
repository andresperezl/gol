// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strings"

	errors "github.com/andresperezl/gol/test/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	champion_impl "github.com/andresperezl/gol/test/impl/champion"
	champion_mastery_impl "github.com/andresperezl/gol/test/impl/champion_mastery"
	"github.com/andresperezl/gol/test/restapi/operations"
	"github.com/andresperezl/gol/test/restapi/operations/champion_mastery"
)

//go:generate swagger generate server --target ../../test --name Testserver --spec ../../specs/api.yaml --default-scheme https

func configureFlags(api *operations.TestserverAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TestserverAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-Riot-Token" header is set
	api.APIKeyHeaderAuth = func(token string) (interface{}, error) {
		if strings.HasPrefix(token, "RGAPI") {
			return token, nil
		}
		return token, errors.New(http.StatusForbidden, "Invalid api key")
	}
	// Applies when the "api_key" query is set
	api.APIKeyQueryParamAuth = func(token string) (interface{}, error) {
		if strings.HasPrefix(token, "RGAPI") {
			return token, nil
		}
		return token, errors.New(http.StatusForbidden, "Invalid api key")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.ChampionMasteryGetAllChampionMasteriesHandler = champion_mastery_impl.GetAllChampionMasteriesHandlerFunc()

	api.ChampionGetChampionInfoHandler = champion_impl.GetChampionInfoHandlerFunc()

	if api.ChampionMasteryGetChampionMasteryHandler == nil {
		api.ChampionMasteryGetChampionMasteryHandler = champion_mastery.GetChampionMasteryHandlerFunc(func(params champion_mastery.GetChampionMasteryParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation champion_mastery.GetChampionMastery has not yet been implemented")
		})
	}
	if api.ChampionMasteryGetChampionMasteryScoreHandler == nil {
		api.ChampionMasteryGetChampionMasteryScoreHandler = champion_mastery.GetChampionMasteryScoreHandlerFunc(func(params champion_mastery.GetChampionMasteryScoreParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation champion_mastery.GetChampionMasteryScore has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
