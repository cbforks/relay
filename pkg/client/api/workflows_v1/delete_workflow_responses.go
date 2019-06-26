// Code generated by go-swagger; DO NOT EDIT.

package workflows_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// DeleteWorkflowReader is a Reader for the DeleteWorkflow structure.
type DeleteWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteWorkflowOK creates a DeleteWorkflowOK with default headers values
func NewDeleteWorkflowOK() *DeleteWorkflowOK {
	return &DeleteWorkflowOK{}
}

/*DeleteWorkflowOK handles this case with default header values.

Deletion successful
*/
type DeleteWorkflowOK struct {
	Payload *models.DeleteResponse
}

func (o *DeleteWorkflowOK) Error() string {
	return fmt.Sprintf("[DELETE /api/workflows/{workflow_name}][%d] deleteWorkflowOK  %+v", 200, o.Payload)
}

func (o *DeleteWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
