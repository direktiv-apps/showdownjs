// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"showdownjs/models"
)

// PostOKCode is the HTTP code returned for type PostOK
const PostOKCode int = 200

/*PostOK nice greeting

swagger:response postOK
*/
type PostOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostOK creates PostOK with default headers values
func NewPostOK() *PostOK {

	return &PostOK{}
}

// WithPayload adds the payload to the post o k response
func (o *PostOK) WithPayload(payload interface{}) *PostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post o k response
func (o *PostOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PostDefault generic error response

swagger:response postDefault
*/
type PostDefault struct {
	_statusCode int
	/*

	 */
	DirektivErrorCode string `json:"Direktiv-ErrorCode"`
	/*

	 */
	DirektivErrorMessage string `json:"Direktiv-ErrorMessage"`

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostDefault creates PostDefault with default headers values
func NewPostDefault(code int) *PostDefault {
	if code <= 0 {
		code = 500
	}

	return &PostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post default response
func (o *PostDefault) WithStatusCode(code int) *PostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post default response
func (o *PostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithDirektivErrorCode adds the direktivErrorCode to the post default response
func (o *PostDefault) WithDirektivErrorCode(direktivErrorCode string) *PostDefault {
	o.DirektivErrorCode = direktivErrorCode
	return o
}

// SetDirektivErrorCode sets the direktivErrorCode to the post default response
func (o *PostDefault) SetDirektivErrorCode(direktivErrorCode string) {
	o.DirektivErrorCode = direktivErrorCode
}

// WithDirektivErrorMessage adds the direktivErrorMessage to the post default response
func (o *PostDefault) WithDirektivErrorMessage(direktivErrorMessage string) *PostDefault {
	o.DirektivErrorMessage = direktivErrorMessage
	return o
}

// SetDirektivErrorMessage sets the direktivErrorMessage to the post default response
func (o *PostDefault) SetDirektivErrorMessage(direktivErrorMessage string) {
	o.DirektivErrorMessage = direktivErrorMessage
}

// WithPayload adds the payload to the post default response
func (o *PostDefault) WithPayload(payload *models.Error) *PostDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post default response
func (o *PostDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Direktiv-ErrorCode

	direktivErrorCode := o.DirektivErrorCode
	if direktivErrorCode != "" {
		rw.Header().Set("Direktiv-ErrorCode", direktivErrorCode)
	}

	// response header Direktiv-ErrorMessage

	direktivErrorMessage := o.DirektivErrorMessage
	if direktivErrorMessage != "" {
		rw.Header().Set("Direktiv-ErrorMessage", direktivErrorMessage)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
