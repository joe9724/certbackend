// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Router router
// swagger:model Router
type Router struct {
	ID *int64 `json:"id" gorm:"AUTO_INCREMENT"`
	// name
	// Required: true
	Name *string `json:"name"`

	// router Url
	// Required: true
	RouterURL *string `json:"routerUrl"`
}

// Validate validates this router
func (m *Router) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRouterURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Router) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Router) validateRouterURL(formats strfmt.Registry) error {

	if err := validate.Required("routerUrl", "body", m.RouterURL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Router) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Router) UnmarshalBinary(b []byte) error {
	var res Router
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}


// Router router
// swagger:model Router
type DynamicRouter struct {
	ID *int64 `json:"id" gorm:"AUTO_INCREMENT"`
	// name
	// Required: true
	Name *string `json:"name"`

	// router Url
	// Required: true
	Router_URL *string `json:"router_url"`

	Parent_Name * string `json:"parent_name"`

	Parent_Id * int64 `json:"parent_id"`

	Children *[] ChildrenRouter `json:"children"`
}

type ChildrenRouter struct {
	ID *int64 `json:"id" gorm:"AUTO_INCREMENT"`
	// name
	// Required: true
	Name *string `json:"name"`

	// router Url
	// Required: true
	Router_URL *string `json:"router_url"`

	Parent_Name * string `json:"parent_name"`

	Parent_Id * int64 `json:"parent_id"`
}