// Code generated by go-swagger; DO NOT EDIT.

package products

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

	"github.com/subhankardas/go-microservices/products-service/models"
)

// NewPostProductsParams creates a new PostProductsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostProductsParams() *PostProductsParams {
	return &PostProductsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostProductsParamsWithTimeout creates a new PostProductsParams object
// with the ability to set a timeout on a request.
func NewPostProductsParamsWithTimeout(timeout time.Duration) *PostProductsParams {
	return &PostProductsParams{
		timeout: timeout,
	}
}

// NewPostProductsParamsWithContext creates a new PostProductsParams object
// with the ability to set a context for a request.
func NewPostProductsParamsWithContext(ctx context.Context) *PostProductsParams {
	return &PostProductsParams{
		Context: ctx,
	}
}

// NewPostProductsParamsWithHTTPClient creates a new PostProductsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostProductsParamsWithHTTPClient(client *http.Client) *PostProductsParams {
	return &PostProductsParams{
		HTTPClient: client,
	}
}

/*
PostProductsParams contains all the parameters to send to the API endpoint

	for the post products operation.

	Typically these are written to a http.Request.
*/
type PostProductsParams struct {

	/* Product.

	   Product Details
	*/
	Product *models.DataProduct

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProductsParams) WithDefaults() *PostProductsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post products params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostProductsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post products params
func (o *PostProductsParams) WithTimeout(timeout time.Duration) *PostProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post products params
func (o *PostProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post products params
func (o *PostProductsParams) WithContext(ctx context.Context) *PostProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post products params
func (o *PostProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post products params
func (o *PostProductsParams) WithHTTPClient(client *http.Client) *PostProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post products params
func (o *PostProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProduct adds the product to the post products params
func (o *PostProductsParams) WithProduct(product *models.DataProduct) *PostProductsParams {
	o.SetProduct(product)
	return o
}

// SetProduct adds the product to the post products params
func (o *PostProductsParams) SetProduct(product *models.DataProduct) {
	o.Product = product
}

// WriteToRequest writes these params to a swagger request
func (o *PostProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Product != nil {
		if err := r.SetBodyParam(o.Product); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}