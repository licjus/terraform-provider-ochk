// Code generated by go-swagger; DO NOT EDIT.

package routers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

// RouterListUsingGETReader is a Reader for the RouterListUsingGET structure.
type RouterListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RouterListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRouterListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRouterListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRouterListUsingGETOK creates a RouterListUsingGETOK with default headers values
func NewRouterListUsingGETOK() *RouterListUsingGETOK {
	return &RouterListUsingGETOK{}
}

/*RouterListUsingGETOK handles this case with default header values.

OK
*/
type RouterListUsingGETOK struct {
	Payload *models.RouterListResponse
}

func (o *RouterListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /network/routers][%d] routerListUsingGETOK  %+v", 200, o.Payload)
}

func (o *RouterListUsingGETOK) GetPayload() *models.RouterListResponse {
	return o.Payload
}

func (o *RouterListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouterListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRouterListUsingGETBadRequest creates a RouterListUsingGETBadRequest with default headers values
func NewRouterListUsingGETBadRequest() *RouterListUsingGETBadRequest {
	return &RouterListUsingGETBadRequest{}
}

/*RouterListUsingGETBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type RouterListUsingGETBadRequest struct {
}

func (o *RouterListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /network/routers][%d] routerListUsingGETBadRequest ", 400)
}

func (o *RouterListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
