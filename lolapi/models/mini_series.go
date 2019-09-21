// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MiniSeries mini series
// swagger:model MiniSeries
type MiniSeries struct {

	// losses
	// Read Only: true
	Losses int32 `json:"losses,omitempty"`

	// progress
	// Read Only: true
	Progress string `json:"progress,omitempty"`

	// target
	// Read Only: true
	Target int32 `json:"target,omitempty"`

	// wins
	// Read Only: true
	Wins int32 `json:"wins,omitempty"`
}

// Validate validates this mini series
func (m *MiniSeries) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MiniSeries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MiniSeries) UnmarshalBinary(b []byte) error {
	var res MiniSeries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
