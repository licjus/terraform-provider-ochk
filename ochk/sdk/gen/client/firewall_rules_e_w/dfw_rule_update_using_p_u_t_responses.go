// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_e_w

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/chmurakrajowa/terraform-provider-ochk/ochk/sdk/gen/models"
)

// DfwRuleUpdateUsingPUTReader is a Reader for the DfwRuleUpdateUsingPUT structure.
type DfwRuleUpdateUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DfwRuleUpdateUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDfwRuleUpdateUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDfwRuleUpdateUsingPUTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDfwRuleUpdateUsingPUTOK creates a DfwRuleUpdateUsingPUTOK with default headers values
func NewDfwRuleUpdateUsingPUTOK() *DfwRuleUpdateUsingPUTOK {
	return &DfwRuleUpdateUsingPUTOK{}
}

/*DfwRuleUpdateUsingPUTOK handles this case with default header values.

OK
*/
type DfwRuleUpdateUsingPUTOK struct {
	Payload *models.UpdateDFWRuleResponse
}

func (o *DfwRuleUpdateUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /network/firewall/security-policies/{securityPolicyId}/rules/{ruleId}][%d] dfwRuleUpdateUsingPUTOK  %+v", 200, o.Payload)
}

func (o *DfwRuleUpdateUsingPUTOK) GetPayload() *models.UpdateDFWRuleResponse {
	return o.Payload
}

func (o *DfwRuleUpdateUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDFWRuleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDfwRuleUpdateUsingPUTBadRequest creates a DfwRuleUpdateUsingPUTBadRequest with default headers values
func NewDfwRuleUpdateUsingPUTBadRequest() *DfwRuleUpdateUsingPUTBadRequest {
	return &DfwRuleUpdateUsingPUTBadRequest{}
}

/*DfwRuleUpdateUsingPUTBadRequest handles this case with default header values.

Bad request, error occurred. For more details see log messages.
*/
type DfwRuleUpdateUsingPUTBadRequest struct {
}

func (o *DfwRuleUpdateUsingPUTBadRequest) Error() string {
	return fmt.Sprintf("[PUT /network/firewall/security-policies/{securityPolicyId}/rules/{ruleId}][%d] dfwRuleUpdateUsingPUTBadRequest ", 400)
}

func (o *DfwRuleUpdateUsingPUTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
