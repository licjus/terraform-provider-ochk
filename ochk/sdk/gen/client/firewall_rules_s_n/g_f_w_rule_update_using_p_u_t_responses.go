// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_s_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// GFWRuleUpdateUsingPUTReader is a Reader for the GFWRuleUpdateUsingPUT structure.
type GFWRuleUpdateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GFWRuleUpdateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGFWRuleUpdateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGFWRuleUpdateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGFWRuleUpdateUsingPUTOK creates a GFWRuleUpdateUsingPUTOK with default headers values
func NewGFWRuleUpdateUsingPUTOK() *GFWRuleUpdateUsingPUTOK {
	return &GFWRuleUpdateUsingPUTOK{}
}

/*GFWRuleUpdateUsingPUTOK handles this case with default header values.

OK
*/
type GFWRuleUpdateUsingPUTOK struct {
	Payload *models.UpdateGFWRuleResponse
}

func (o *GFWRuleUpdateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /network/firewall/gateway-policies/{gatewayPolicyId}/rules/{ruleId}][%d] gFWRuleUpdateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *GFWRuleUpdateUsingPUTOK) GetPayload() *models.UpdateGFWRuleResponse {
	return o.Payload
}

func (o *GFWRuleUpdateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateGFWRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGFWRuleUpdateUsingPUTBadRequest creates a GFWRuleUpdateUsingPUTBadRequest with default headers values
func NewGFWRuleUpdateUsingPUTBadRequest() *GFWRuleUpdateUsingPUTBadRequest {
	return &GFWRuleUpdateUsingPUTBadRequest{}
}

/*GFWRuleUpdateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type GFWRuleUpdateUsingPUTBadRequest struct {
}

func (o *GFWRuleUpdateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /network/firewall/gateway-policies/{gatewayPolicyId}/rules/{ruleId}][%d] gFWRuleUpdateUsingPUTBadRequest ", 400)
}

func (o *GFWRuleUpdateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
