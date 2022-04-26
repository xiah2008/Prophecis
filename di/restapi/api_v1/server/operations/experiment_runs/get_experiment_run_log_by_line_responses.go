// Code generated by go-swagger; DO NOT EDIT.

package experiment_runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	restmodels "webank/DI/restapi/api_v1/restmodels"
)

// GetExperimentRunLogByLineOKCode is the HTTP code returned for type GetExperimentRunLogByLineOK
const GetExperimentRunLogByLineOKCode int = 200

/*GetExperimentRunLogByLineOK OK

swagger:response getExperimentRunLogByLineOK
*/
type GetExperimentRunLogByLineOK struct {
}

// NewGetExperimentRunLogByLineOK creates GetExperimentRunLogByLineOK with default headers values
func NewGetExperimentRunLogByLineOK() *GetExperimentRunLogByLineOK {

	return &GetExperimentRunLogByLineOK{}
}

// WriteResponse to the client
func (o *GetExperimentRunLogByLineOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetExperimentRunLogByLineUnauthorizedCode is the HTTP code returned for type GetExperimentRunLogByLineUnauthorized
const GetExperimentRunLogByLineUnauthorizedCode int = 401

/*GetExperimentRunLogByLineUnauthorized Unauthorized

swagger:response getExperimentRunLogByLineUnauthorized
*/
type GetExperimentRunLogByLineUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewGetExperimentRunLogByLineUnauthorized creates GetExperimentRunLogByLineUnauthorized with default headers values
func NewGetExperimentRunLogByLineUnauthorized() *GetExperimentRunLogByLineUnauthorized {

	return &GetExperimentRunLogByLineUnauthorized{}
}

// WithPayload adds the payload to the get experiment run log by line unauthorized response
func (o *GetExperimentRunLogByLineUnauthorized) WithPayload(payload *restmodels.Error) *GetExperimentRunLogByLineUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get experiment run log by line unauthorized response
func (o *GetExperimentRunLogByLineUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetExperimentRunLogByLineUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetExperimentRunLogByLineNotFoundCode is the HTTP code returned for type GetExperimentRunLogByLineNotFound
const GetExperimentRunLogByLineNotFoundCode int = 404

/*GetExperimentRunLogByLineNotFound Model create failed

swagger:response getExperimentRunLogByLineNotFound
*/
type GetExperimentRunLogByLineNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewGetExperimentRunLogByLineNotFound creates GetExperimentRunLogByLineNotFound with default headers values
func NewGetExperimentRunLogByLineNotFound() *GetExperimentRunLogByLineNotFound {

	return &GetExperimentRunLogByLineNotFound{}
}

// WithPayload adds the payload to the get experiment run log by line not found response
func (o *GetExperimentRunLogByLineNotFound) WithPayload(payload *restmodels.Error) *GetExperimentRunLogByLineNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get experiment run log by line not found response
func (o *GetExperimentRunLogByLineNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetExperimentRunLogByLineNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
