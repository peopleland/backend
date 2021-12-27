package v1

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ErrorResponse error response
//
// swagger:model ErrorResponse
type ErrorResponse struct {

	// error
	// Required: true
	Error *string `json:"error"`
}

// Validate validates this error response
func (m *ErrorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponse) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error response based on context it is used
func (m *ErrorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponse) UnmarshalBinary(b []byte) error {
	var res ErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoginPayLoad login pay load
//
// swagger:model LoginPayLoad
type LoginPayLoad struct {

	// address
	// Required: true
	Address *string `json:"address"`

	// origin message
	// Required: true
	OriginMessage *string `json:"origin_message"`

	// signature
	// Required: true
	Signature *string `json:"signature"`
}

// Validate validates this login pay load
func (m *LoginPayLoad) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginPayLoad) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *LoginPayLoad) validateOriginMessage(formats strfmt.Registry) error {

	if err := validate.Required("origin_message", "body", m.OriginMessage); err != nil {
		return err
	}

	return nil
}

func (m *LoginPayLoad) validateSignature(formats strfmt.Registry) error {

	if err := validate.Required("signature", "body", m.Signature); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login pay load based on context it is used
func (m *LoginPayLoad) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginPayLoad) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginPayLoad) UnmarshalBinary(b []byte) error {
	var res LoginPayLoad
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LoginResponse login response
//
// swagger:model LoginResponse
type LoginResponse struct {

	// jwt
	// Required: true
	Jwt *string `json:"jwt"`
}

// Validate validates this login response
func (m *LoginResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJwt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginResponse) validateJwt(formats strfmt.Registry) error {

	if err := validate.Required("jwt", "body", m.Jwt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this login response based on context it is used
func (m *LoginResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginResponse) UnmarshalBinary(b []byte) error {
	var res LoginResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PutProfilePayLoad put profile pay load
//
// swagger:model PutProfilePayLoad
type PutProfilePayLoad struct {

	// twitter
	Twitter string `json:"twitter,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this put profile pay load
func (m *PutProfilePayLoad) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put profile pay load based on context it is used
func (m *PutProfilePayLoad) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutProfilePayLoad) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutProfilePayLoad) UnmarshalBinary(b []byte) error {
	var res PutProfilePayLoad
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserProfile user profile
//
// swagger:model UserProfile
type UserProfile struct {

	// address
	Address string `json:"address,omitempty"`

	// discord
	Discord string `json:"discord,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	Name string `json:"name,omitempty"`

	// twitter
	Twitter string `json:"twitter,omitempty"`
}

// Validate validates this user profile
func (m *UserProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserProfile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user profile based on context it is used
func (m *UserProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserProfile) UnmarshalBinary(b []byte) error {
	var res UserProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
