// Code generated by go-swagger; DO NOT EDIT.

package experiments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	restmodels "webank/DI/restapi/api_v1/restmodels"
)

// NewUpdateExperimentParams creates a new UpdateExperimentParams object
// no default values defined in spec.
func NewUpdateExperimentParams() UpdateExperimentParams {

	return UpdateExperimentParams{}
}

// UpdateExperimentParams contains all the bound params for the update experiment operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateExperiment
type UpdateExperimentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The Experiment Put Request.
	  Required: true
	  In: body
	*/
	Experiment *restmodels.ProphecisExperimentPutRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateExperimentParams() beforehand.
func (o *UpdateExperimentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body restmodels.ProphecisExperimentPutRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("experiment", "body"))
			} else {
				res = append(res, errors.NewParseError("experiment", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Experiment = &body
			}
		}
	} else {
		res = append(res, errors.Required("experiment", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
