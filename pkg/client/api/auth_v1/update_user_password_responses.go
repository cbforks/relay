// Code generated by go-swagger; DO NOT EDIT.

package auth_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// UpdateUserPasswordReader is a Reader for the UpdateUserPassword structure.
type UpdateUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateUserPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewUpdateUserPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserPasswordOK creates a UpdateUserPasswordOK with default headers values
func NewUpdateUserPasswordOK() *UpdateUserPasswordOK {
	return &UpdateUserPasswordOK{}
}

/*UpdateUserPasswordOK handles this case with default header values.

User password update successful
*/
type UpdateUserPasswordOK struct {
	Payload *models.GenericSuccess
}

func (o *UpdateUserPasswordOK) Error() string {
	return fmt.Sprintf("[POST /auth/users/{id}/change-password][%d] updateUserPasswordOK  %+v", 200, o.Payload)
}

func (o *UpdateUserPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericSuccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserPasswordForbidden creates a UpdateUserPasswordForbidden with default headers values
func NewUpdateUserPasswordForbidden() *UpdateUserPasswordForbidden {
	return &UpdateUserPasswordForbidden{}
}

/*UpdateUserPasswordForbidden handles this case with default header values.

Unauthorized to access this resource
*/
type UpdateUserPasswordForbidden struct {
}

func (o *UpdateUserPasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /auth/users/{id}/change-password][%d] updateUserPasswordForbidden ", 403)
}

func (o *UpdateUserPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
