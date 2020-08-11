// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_e_w

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ochk/terraform-provider-ochk/ochk/sdk/gen/models"
)

// NewDFWRuleUpdateUsingPUTParams creates a new DFWRuleUpdateUsingPUTParams object
// with the default values initialized.
func NewDFWRuleUpdateUsingPUTParams() *DFWRuleUpdateUsingPUTParams {
	var ()
	return &DFWRuleUpdateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDFWRuleUpdateUsingPUTParamsWithTimeout creates a new DFWRuleUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDFWRuleUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *DFWRuleUpdateUsingPUTParams {
	var ()
	return &DFWRuleUpdateUsingPUTParams{

		timeout: timeout,
	}
}

// NewDFWRuleUpdateUsingPUTParamsWithContext creates a new DFWRuleUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewDFWRuleUpdateUsingPUTParamsWithContext(ctx context.Context) *DFWRuleUpdateUsingPUTParams {
	var ()
	return &DFWRuleUpdateUsingPUTParams{

		Context: ctx,
	}
}

// NewDFWRuleUpdateUsingPUTParamsWithHTTPClient creates a new DFWRuleUpdateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDFWRuleUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *DFWRuleUpdateUsingPUTParams {
	var ()
	return &DFWRuleUpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/*DFWRuleUpdateUsingPUTParams contains all the parameters to send to the API endpoint
for the d f w rule update using p u t operation typically these are written to a http.Request
*/
type DFWRuleUpdateUsingPUTParams struct {

	/*DfwRule
	  dfwRule

	*/
	DfwRule *models.DFWRule
	/*RuleID
	  ruleId

	*/
	RuleID string
	/*SecurityPolicyID
	  securityPolicyId

	*/
	SecurityPolicyID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) WithTimeout(timeout time.Duration) *DFWRuleUpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) WithContext(ctx context.Context) *DFWRuleUpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) WithHTTPClient(client *http.Client) *DFWRuleUpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDfwRule adds the dfwRule to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) WithDfwRule(dfwRule *models.DFWRule) *DFWRuleUpdateUsingPUTParams {
	o.SetDfwRule(dfwRule)
	return o
}

// SetDfwRule adds the dfwRule to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) SetDfwRule(dfwRule *models.DFWRule) {
	o.DfwRule = dfwRule
}

// WithRuleID adds the ruleID to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) WithRuleID(ruleID string) *DFWRuleUpdateUsingPUTParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WithSecurityPolicyID adds the securityPolicyID to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) WithSecurityPolicyID(securityPolicyID string) *DFWRuleUpdateUsingPUTParams {
	o.SetSecurityPolicyID(securityPolicyID)
	return o
}

// SetSecurityPolicyID adds the securityPolicyId to the d f w rule update using p u t params
func (o *DFWRuleUpdateUsingPUTParams) SetSecurityPolicyID(securityPolicyID string) {
	o.SecurityPolicyID = securityPolicyID
}

// WriteToRequest writes these params to a swagger request
func (o *DFWRuleUpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DfwRule != nil {
		if err := r.SetBodyParam(o.DfwRule); err != nil {
			return err
		}
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	// path param securityPolicyId
	if err := r.SetPathParam("securityPolicyId", o.SecurityPolicyID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
