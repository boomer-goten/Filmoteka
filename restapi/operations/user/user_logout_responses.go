// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UserLogoutOKCode is the HTTP code returned for type UserLogoutOK
const UserLogoutOKCode int = 200

/*
UserLogoutOK Успешно

swagger:response userLogoutOK
*/
type UserLogoutOK struct {
}

// NewUserLogoutOK creates UserLogoutOK with default headers values
func NewUserLogoutOK() *UserLogoutOK {

	return &UserLogoutOK{}
}

// WriteResponse to the client
func (o *UserLogoutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
