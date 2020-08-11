// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_s_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GFWRuleDeleteUsingDELETEReader is a Reader for the GFWRuleDeleteUsingDELETE structure.
type GFWRuleDeleteUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GFWRuleDeleteUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGFWRuleDeleteUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGFWRuleDeleteUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGFWRuleDeleteUsingDELETEOK creates a GFWRuleDeleteUsingDELETEOK with default headers values
func NewGFWRuleDeleteUsingDELETEOK() *GFWRuleDeleteUsingDELETEOK {
	return &GFWRuleDeleteUsingDELETEOK{}
}

/*GFWRuleDeleteUsingDELETEOK handles this case with default header values.

OK
*/
type GFWRuleDeleteUsingDELETEOK struct {
	Payload *models.DeleteGFWRuleResponse
}

func (o *GFWRuleDeleteUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /network/firewall/gateway-policies/{gatewayPolicyId}/rules/{ruleId}][%d] gFWRuleDeleteUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *GFWRuleDeleteUsingDELETEOK) GetPayload() *models.DeleteGFWRuleResponse {
	return o.Payload
}

func (o *GFWRuleDeleteUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteGFWRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGFWRuleDeleteUsingDELETEBadRequest creates a GFWRuleDeleteUsingDELETEBadRequest with default headers values
func NewGFWRuleDeleteUsingDELETEBadRequest() *GFWRuleDeleteUsingDELETEBadRequest {
	return &GFWRuleDeleteUsingDELETEBadRequest{}
}

/*GFWRuleDeleteUsingDELETEBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GFWRuleDeleteUsingDELETEBadRequest struct {
}

func (o *GFWRuleDeleteUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /network/firewall/gateway-policies/{gatewayPolicyId}/rules/{ruleId}][%d] gFWRuleDeleteUsingDELETEBadRequest ", 400)
}

func (o *GFWRuleDeleteUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
