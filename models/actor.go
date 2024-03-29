// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Actor actor
//
// swagger:model Actor
type Actor struct {

	// date of birthday
	// Example: 1963-06-09
	// Required: true
	// Format: date
	DateOfBirthday *strfmt.Date `json:"date_of_birthday"`

	// name
	// Example: Johnny Depp
	// Required: true
	Name *string `json:"name"`

	// sex
	// Example: M
	// Required: true
	// Enum: ["M","F"]
	Sex *string `json:"sex"`
}

// Validate validates this actor
func (m *Actor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateOfBirthday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Actor) validateDateOfBirthday(formats strfmt.Registry) error {

	if err := validate.Required("date_of_birthday", "body", m.DateOfBirthday); err != nil {
		return err
	}

	if err := validate.FormatOf("date_of_birthday", "body", "date", m.DateOfBirthday.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Actor) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var actorTypeSexPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["M","F"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actorTypeSexPropEnum = append(actorTypeSexPropEnum, v)
	}
}

const (

	// ActorSexM captures enum value "M"
	ActorSexM string = "M"

	// ActorSexF captures enum value "F"
	ActorSexF string = "F"
)

// prop value enum
func (m *Actor) validateSexEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, actorTypeSexPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Actor) validateSex(formats strfmt.Registry) error {

	if err := validate.Required("sex", "body", m.Sex); err != nil {
		return err
	}

	// value enum
	if err := m.validateSexEnum("sex", "body", *m.Sex); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this actor based on context it is used
func (m *Actor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Actor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Actor) UnmarshalBinary(b []byte) error {
	var res Actor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
