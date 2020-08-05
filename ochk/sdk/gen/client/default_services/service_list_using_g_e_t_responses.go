// Code generated by go-swagger; DO NOT EDIT.

package default_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

// ServiceListUsingGETReader is a Reader for the ServiceListUsingGET structure.
type ServiceListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceListUsingGETOK creates a ServiceListUsingGETOK with default headers values
func NewServiceListUsingGETOK() *ServiceListUsingGETOK {
	return &ServiceListUsingGETOK{}
}

/*ServiceListUsingGETOK handles this case with default header values.

OK
*/
type ServiceListUsingGETOK struct {
	Payload *models.ServiceListResponse
}

func (o *ServiceListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/default-services][%d] serviceListUsingGETOK  %+v", 200, o.Payload)
}

func (o *ServiceListUsingGETOK) GetPayload() *models.ServiceListResponse {
	return o.Payload
}

func (o *ServiceListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceListUsingGETBadRequest creates a ServiceListUsingGETBadRequest with default headers values
func NewServiceListUsingGETBadRequest() *ServiceListUsingGETBadRequest {
	return &ServiceListUsingGETBadRequest{}
}

/*ServiceListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type ServiceListUsingGETBadRequest struct {
}

func (o *ServiceListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/default-services][%d] serviceListUsingGETBadRequest ", 400)
}

func (o *ServiceListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
