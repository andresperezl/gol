// Code generated by go-swagger; DO NOT EDIT.

package champion_mastery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetChampionMasteryURL generates an URL for the get champion mastery operation
type GetChampionMasteryURL struct {
	ChampionID          int64
	EncryptedSummonerID string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetChampionMasteryURL) WithBasePath(bp string) *GetChampionMasteryURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetChampionMasteryURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetChampionMasteryURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/champion-mastery/v4/champion-masteries/by-summoner/{encryptedSummonerId}/by-champion/{championId}"

	championID := swag.FormatInt64(o.ChampionID)
	if championID != "" {
		_path = strings.Replace(_path, "{championId}", championID, -1)
	} else {
		return nil, errors.New("championId is required on GetChampionMasteryURL")
	}

	encryptedSummonerID := o.EncryptedSummonerID
	if encryptedSummonerID != "" {
		_path = strings.Replace(_path, "{encryptedSummonerId}", encryptedSummonerID, -1)
	} else {
		return nil, errors.New("encryptedSummonerId is required on GetChampionMasteryURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/lol"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetChampionMasteryURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetChampionMasteryURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetChampionMasteryURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetChampionMasteryURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetChampionMasteryURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetChampionMasteryURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}