// Code generated by go-swagger; DO NOT EDIT.

package access_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// CreateInviteReader is a Reader for the CreateInvite structure.
type CreateInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInviteCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateInviteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateInviteCreated creates a CreateInviteCreated with default headers values
func NewCreateInviteCreated() *CreateInviteCreated {
	return &CreateInviteCreated{}
}

/*CreateInviteCreated handles this case with default header values.

The newly created invite representation
*/
type CreateInviteCreated struct {
	Payload *CreateInviteCreatedBody
}

func (o *CreateInviteCreated) Error() string {
	return fmt.Sprintf("[POST /api/invites][%d] createInviteCreated  %+v", 201, o.Payload)
}

func (o *CreateInviteCreated) GetPayload() *CreateInviteCreatedBody {
	return o.Payload
}

func (o *CreateInviteCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateInviteCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInviteDefault creates a CreateInviteDefault with default headers values
func NewCreateInviteDefault(code int) *CreateInviteDefault {
	return &CreateInviteDefault{
		_statusCode: code,
	}
}

/*CreateInviteDefault handles this case with default header values.

An error occurred
*/
type CreateInviteDefault struct {
	_statusCode int

	Payload *CreateInviteDefaultBody
}

// Code gets the status code for the create invite default response
func (o *CreateInviteDefault) Code() int {
	return o._statusCode
}

func (o *CreateInviteDefault) Error() string {
	return fmt.Sprintf("[POST /api/invites][%d] createInvite default  %+v", o._statusCode, o.Payload)
}

func (o *CreateInviteDefault) GetPayload() *CreateInviteDefaultBody {
	return o.Payload
}

func (o *CreateInviteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateInviteDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateInviteBody Request type for creating a new invite
swagger:model CreateInviteBody
*/
type CreateInviteBody struct {

	// The email address of the invitee
	// Required: true
	Email *string `json:"email"`

	// The full name of the invitee
	// Required: true
	Name *string `json:"name"`

	// The initial roles to grant to the user
	Roles []*models.RoleIdentifier `json:"roles"`
}

// Validate validates this create invite body
func (o *CreateInviteBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateInviteBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *CreateInviteBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateInviteBody) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(o.Roles) { // not required
		return nil
	}

	for i := 0; i < len(o.Roles); i++ {
		if swag.IsZero(o.Roles[i]) { // not required
			continue
		}

		if o.Roles[i] != nil {
			if err := o.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateInviteBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateInviteBody) UnmarshalBinary(b []byte) error {
	var res CreateInviteBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateInviteCreatedBody create invite created body
swagger:model CreateInviteCreatedBody
*/
type CreateInviteCreatedBody struct {
	models.InviteEntity

	// invite
	Invite *models.Invite `json:"invite,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *CreateInviteCreatedBody) UnmarshalJSON(raw []byte) error {
	// CreateInviteCreatedBodyAO0
	var createInviteCreatedBodyAO0 models.InviteEntity
	if err := swag.ReadJSON(raw, &createInviteCreatedBodyAO0); err != nil {
		return err
	}
	o.InviteEntity = createInviteCreatedBodyAO0

	// CreateInviteCreatedBodyAO1
	var dataCreateInviteCreatedBodyAO1 struct {
		Invite *models.Invite `json:"invite,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataCreateInviteCreatedBodyAO1); err != nil {
		return err
	}

	o.Invite = dataCreateInviteCreatedBodyAO1.Invite

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o CreateInviteCreatedBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	createInviteCreatedBodyAO0, err := swag.WriteJSON(o.InviteEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, createInviteCreatedBodyAO0)

	var dataCreateInviteCreatedBodyAO1 struct {
		Invite *models.Invite `json:"invite,omitempty"`
	}

	dataCreateInviteCreatedBodyAO1.Invite = o.Invite

	jsonDataCreateInviteCreatedBodyAO1, errCreateInviteCreatedBodyAO1 := swag.WriteJSON(dataCreateInviteCreatedBodyAO1)
	if errCreateInviteCreatedBodyAO1 != nil {
		return nil, errCreateInviteCreatedBodyAO1
	}
	_parts = append(_parts, jsonDataCreateInviteCreatedBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this create invite created body
func (o *CreateInviteCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.InviteEntity
	if err := o.InviteEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInvite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateInviteCreatedBody) validateInvite(formats strfmt.Registry) error {

	if swag.IsZero(o.Invite) { // not required
		return nil
	}

	if o.Invite != nil {
		if err := o.Invite.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createInviteCreated" + "." + "invite")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateInviteCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateInviteCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateInviteCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateInviteDefaultBody Error response
swagger:model CreateInviteDefaultBody
*/
type CreateInviteDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this create invite default body
func (o *CreateInviteDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateInviteDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createInvite default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateInviteDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateInviteDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateInviteDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}