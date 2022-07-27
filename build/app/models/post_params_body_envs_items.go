// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostParamsBodyEnvsItems post params body envs items
//
// swagger:model postParamsBodyEnvsItems
type PostParamsBodyEnvsItems struct {

	// Name of the variable.
	Name string `json:"name,omitempty"`

	// Value of the variable.
	Value string `json:"value,omitempty"`
}

// Validate validates this post params body envs items
func (m *PostParamsBodyEnvsItems) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post params body envs items based on context it is used
func (m *PostParamsBodyEnvsItems) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBodyEnvsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBodyEnvsItems) UnmarshalBinary(b []byte) error {
	var res PostParamsBodyEnvsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}