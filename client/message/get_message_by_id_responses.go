// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/karixtech/go-karix/models"
)

// GetMessageByIDReader is a Reader for the GetMessageByID structure.
type GetMessageByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMessageByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMessageByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetMessageByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetMessageByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetMessageByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetMessageByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMessageByIDOK creates a GetMessageByIDOK with default headers values
func NewGetMessageByIDOK() *GetMessageByIDOK {
	return &GetMessageByIDOK{}
}

/*GetMessageByIDOK handles this case with default header values.

Message data
*/
type GetMessageByIDOK struct {
	Payload *models.GetMessageByIDOKBody
}

func (o *GetMessageByIDOK) Error() string {
	return fmt.Sprintf("[GET /message/{uid}/][%d] getMessageByIdOK  %+v", 200, o.Payload)
}

func (o *GetMessageByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMessageByIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessageByIDUnauthorized creates a GetMessageByIDUnauthorized with default headers values
func NewGetMessageByIDUnauthorized() *GetMessageByIDUnauthorized {
	return &GetMessageByIDUnauthorized{}
}

/*GetMessageByIDUnauthorized handles this case with default header values.

Authentication information is missing or invalid
*/
type GetMessageByIDUnauthorized struct {
	WWWAuthenticate string
}

func (o *GetMessageByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /message/{uid}/][%d] getMessageByIdUnauthorized ", 401)
}

func (o *GetMessageByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header WWW_Authenticate
	o.WWWAuthenticate = response.GetHeader("WWW_Authenticate")

	return nil
}

// NewGetMessageByIDForbidden creates a GetMessageByIDForbidden with default headers values
func NewGetMessageByIDForbidden() *GetMessageByIDForbidden {
	return &GetMessageByIDForbidden{}
}

/*GetMessageByIDForbidden handles this case with default header values.

User not authorized or blocked
*/
type GetMessageByIDForbidden struct {
	Payload *models.GetMessageByIDForbiddenBody
}

func (o *GetMessageByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /message/{uid}/][%d] getMessageByIdForbidden  %+v", 403, o.Payload)
}

func (o *GetMessageByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMessageByIDForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessageByIDNotFound creates a GetMessageByIDNotFound with default headers values
func NewGetMessageByIDNotFound() *GetMessageByIDNotFound {
	return &GetMessageByIDNotFound{}
}

/*GetMessageByIDNotFound handles this case with default header values.

Resource not found
*/
type GetMessageByIDNotFound struct {
	Payload *models.GetMessageByIDNotFoundBody
}

func (o *GetMessageByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /message/{uid}/][%d] getMessageByIdNotFound  %+v", 404, o.Payload)
}

func (o *GetMessageByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMessageByIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMessageByIDInternalServerError creates a GetMessageByIDInternalServerError with default headers values
func NewGetMessageByIDInternalServerError() *GetMessageByIDInternalServerError {
	return &GetMessageByIDInternalServerError{}
}

/*GetMessageByIDInternalServerError handles this case with default header values.

Error
*/
type GetMessageByIDInternalServerError struct {
	Payload *models.GetMessageByIDInternalServerErrorBody
}

func (o *GetMessageByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /message/{uid}/][%d] getMessageByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMessageByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMessageByIDInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
