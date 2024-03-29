// Code generated by go-swagger; DO NOT EDIT.

package actor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAllActorsOKCode is the HTTP code returned for type GetAllActorsOK
const GetAllActorsOKCode int = 200

/*
GetAllActorsOK Список актеров

swagger:response getAllActorsOK
*/
type GetAllActorsOK struct {

	/*
	  In: Body
	*/
	Payload []*GetAllActorsOKBodyItems0 `json:"body,omitempty"`
}

// NewGetAllActorsOK creates GetAllActorsOK with default headers values
func NewGetAllActorsOK() *GetAllActorsOK {

	return &GetAllActorsOK{}
}

// WithPayload adds the payload to the get all actors o k response
func (o *GetAllActorsOK) WithPayload(payload []*GetAllActorsOKBodyItems0) *GetAllActorsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all actors o k response
func (o *GetAllActorsOK) SetPayload(payload []*GetAllActorsOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllActorsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*GetAllActorsOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAllActorsBadRequestCode is the HTTP code returned for type GetAllActorsBadRequest
const GetAllActorsBadRequestCode int = 400

/*
GetAllActorsBadRequest Ошибка

swagger:response getAllActorsBadRequest
*/
type GetAllActorsBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetAllActorsBadRequest creates GetAllActorsBadRequest with default headers values
func NewGetAllActorsBadRequest() *GetAllActorsBadRequest {

	return &GetAllActorsBadRequest{}
}

// WithPayload adds the payload to the get all actors bad request response
func (o *GetAllActorsBadRequest) WithPayload(payload string) *GetAllActorsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all actors bad request response
func (o *GetAllActorsBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllActorsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
