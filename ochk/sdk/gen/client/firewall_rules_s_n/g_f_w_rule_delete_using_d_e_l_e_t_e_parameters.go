// Code generated by go-swagger; DO NOT EDIT.

package firewall_rules_s_n

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
)

// NewGFWRuleDeleteUsingDELETEParams creates a new GFWRuleDeleteUsingDELETEParams object
// with the default values initialized.
func NewGFWRuleDeleteUsingDELETEParams() *GFWRuleDeleteUsingDELETEParams {
	var ()
	return &GFWRuleDeleteUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGFWRuleDeleteUsingDELETEParamsWithTimeout creates a new GFWRuleDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGFWRuleDeleteUsingDELETEParamsWithTimeout(timeout time.Duration) *GFWRuleDeleteUsingDELETEParams {
	var ()
	return &GFWRuleDeleteUsingDELETEParams{

		timeout: timeout,
	}
}

// NewGFWRuleDeleteUsingDELETEParamsWithContext creates a new GFWRuleDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewGFWRuleDeleteUsingDELETEParamsWithContext(ctx context.Context) *GFWRuleDeleteUsingDELETEParams {
	var ()
	return &GFWRuleDeleteUsingDELETEParams{

		Context: ctx,
	}
}

// NewGFWRuleDeleteUsingDELETEParamsWithHTTPClient creates a new GFWRuleDeleteUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGFWRuleDeleteUsingDELETEParamsWithHTTPClient(client *http.Client) *GFWRuleDeleteUsingDELETEParams {
	var ()
	return &GFWRuleDeleteUsingDELETEParams{
		HTTPClient: client,
	}
}

/*GFWRuleDeleteUsingDELETEParams contains all the parameters to send to the API endpoint
for the g f w rule delete using d e l e t e operation typically these are written to a http.Request
*/
type GFWRuleDeleteUsingDELETEParams struct {

	/*GatewayPolicyID
	  gatewayPolicyId

	*/
	GatewayPolicyID string
	/*RuleID
	  ruleId

	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) WithTimeout(timeout time.Duration) *GFWRuleDeleteUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) WithContext(ctx context.Context) *GFWRuleDeleteUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) WithHTTPClient(client *http.Client) *GFWRuleDeleteUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayPolicyID adds the gatewayPolicyID to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) WithGatewayPolicyID(gatewayPolicyID string) *GFWRuleDeleteUsingDELETEParams {
	o.SetGatewayPolicyID(gatewayPolicyID)
	return o
}

// SetGatewayPolicyID adds the gatewayPolicyId to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) SetGatewayPolicyID(gatewayPolicyID string) {
	o.GatewayPolicyID = gatewayPolicyID
}

// WithRuleID adds the ruleID to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) WithRuleID(ruleID string) *GFWRuleDeleteUsingDELETEParams {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the g f w rule delete using d e l e t e params
func (o *GFWRuleDeleteUsingDELETEParams) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *GFWRuleDeleteUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gatewayPolicyId
	if err := r.SetPathParam("gatewayPolicyId", o.GatewayPolicyID); err != nil {
		return err
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
