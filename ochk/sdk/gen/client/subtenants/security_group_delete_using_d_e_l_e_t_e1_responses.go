// Code generated by go-swagger; DO NOT EDIT.

package subtenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// SecurityGroupDeleteUsingDELETE1Reader is a Reader for the SecurityGroupDeleteUsingDELETE1 structure.
type SecurityGroupDeleteUsingDELETE1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityGroupDeleteUsingDELETE1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityGroupDeleteUsingDELETE1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSecurityGroupDeleteUsingDELETE1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSecurityGroupDeleteUsingDELETE1OK creates a SecurityGroupDeleteUsingDELETE1OK with default headers values
func NewSecurityGroupDeleteUsingDELETE1OK() *SecurityGroupDeleteUsingDELETE1OK {
	return &SecurityGroupDeleteUsingDELETE1OK{}
}

/*SecurityGroupDeleteUsingDELETE1OK handles this case with default header values.

OK
*/
type SecurityGroupDeleteUsingDELETE1OK struct {
	Payload *models.SubtenantDeleteResponse
}

func (o *SecurityGroupDeleteUsingDELETE1OK) Error() string {
	return fmt.Sprintf("[DELETE /subtenants/{subtenantId}][%d] securityGroupDeleteUsingDELETE1OK  %+v", 200, o.Payload)
}

func (o *SecurityGroupDeleteUsingDELETE1OK) GetPayload() *models.SubtenantDeleteResponse {
	return o.Payload
}

func (o *SecurityGroupDeleteUsingDELETE1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubtenantDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityGroupDeleteUsingDELETE1BadRequest creates a SecurityGroupDeleteUsingDELETE1BadRequest with default headers values
func NewSecurityGroupDeleteUsingDELETE1BadRequest() *SecurityGroupDeleteUsingDELETE1BadRequest {
	return &SecurityGroupDeleteUsingDELETE1BadRequest{}
}

/*SecurityGroupDeleteUsingDELETE1BadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type SecurityGroupDeleteUsingDELETE1BadRequest struct {
}

func (o *SecurityGroupDeleteUsingDELETE1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /subtenants/{subtenantId}][%d] securityGroupDeleteUsingDELETE1BadRequest ", 400)
}

func (o *SecurityGroupDeleteUsingDELETE1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
