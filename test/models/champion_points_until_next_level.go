// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// ChampionPointsUntilNextLevel Number of points needed to achieve next level. Zero if player reached maximum champion level for this champion
// swagger:model ChampionPointsUntilNextLevel
type ChampionPointsUntilNextLevel int64

// Validate validates this champion points until next level
func (m ChampionPointsUntilNextLevel) Validate(formats strfmt.Registry) error {
	return nil
}