// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryChangeMySQLServiceRequest inventory change my SQL service request
// swagger:model inventoryChangeMySQLServiceRequest
type InventoryChangeMySQLServiceRequest struct {

	// Unique Service identifier.
	ID int64 `json:"id,omitempty"`

	// Unique user-defined Service name.
	Name string `json:"name,omitempty"`
}

// Validate validates this inventory change my SQL service request
func (m *InventoryChangeMySQLServiceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryChangeMySQLServiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryChangeMySQLServiceRequest) UnmarshalBinary(b []byte) error {
	var res InventoryChangeMySQLServiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
