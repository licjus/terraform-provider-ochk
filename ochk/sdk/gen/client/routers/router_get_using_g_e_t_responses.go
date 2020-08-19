// Code generated by go-swagger; DO NOT EDIT.

package routers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// RouterGetUsingGETReader is a Reader for the RouterGetUsingGET structure.
type RouterGetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouterGetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouterGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRouterGetUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRouterGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRouterGetUsingGETOK creates a RouterGetUsingGETOK with default headers values
func NewRouterGetUsingGETOK() *RouterGetUsingGETOK {
	return &RouterGetUsingGETOK{}
}

/*RouterGetUsingGETOK handles this case with default header values.

OK
*/
type RouterGetUsingGETOK struct {
	Payload *models.RouterGetResponse
}

func (o *RouterGetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/routers/{routerId}][%d] routerGetUsingGETOK  %+v", 200, o.Payload)
}

func (o *RouterGetUsingGETOK) GetPayload() *models.RouterGetResponse {
	return o.Payload
}

func (o *RouterGetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouterGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouterGetUsingGETBadRequest creates a RouterGetUsingGETBadRequest with default headers values
func NewRouterGetUsingGETBadRequest() *RouterGetUsingGETBadRequest {
	return &RouterGetUsingGETBadRequest{}
}

/*RouterGetUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type RouterGetUsingGETBadRequest struct {
}

func (o *RouterGetUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/routers/{routerId}][%d] routerGetUsingGETBadRequest ", 400)
}

func (o *RouterGetUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRouterGetUsingGETNotFound creates a RouterGetUsingGETNotFound with default headers values
func NewRouterGetUsingGETNotFound() *RouterGetUsingGETNotFound {
	return &RouterGetUsingGETNotFound{}
}

/*RouterGetUsingGETNotFound handles this case with default header values.

Entity not found.
*/
type RouterGetUsingGETNotFound struct {
}

func (o *RouterGetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /network/routers/{routerId}][%d] routerGetUsingGETNotFound ", 404)
}

func (o *RouterGetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
