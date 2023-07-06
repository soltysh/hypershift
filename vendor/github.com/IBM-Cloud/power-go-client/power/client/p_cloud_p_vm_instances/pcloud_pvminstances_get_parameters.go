// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

// NewPcloudPvminstancesGetParams creates a new PcloudPvminstancesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesGetParams() *PcloudPvminstancesGetParams {
	return &PcloudPvminstancesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesGetParamsWithTimeout creates a new PcloudPvminstancesGetParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesGetParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesGetParams {
	return &PcloudPvminstancesGetParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesGetParamsWithContext creates a new PcloudPvminstancesGetParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesGetParamsWithContext(ctx context.Context) *PcloudPvminstancesGetParams {
	return &PcloudPvminstancesGetParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesGetParamsWithHTTPClient creates a new PcloudPvminstancesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesGetParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesGetParams {
	return &PcloudPvminstancesGetParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesGetParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances get operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesGetParams) WithDefaults() *PcloudPvminstancesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) WithContext(ctx context.Context) *PcloudPvminstancesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesGetParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances get params
func (o *PcloudPvminstancesGetParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
