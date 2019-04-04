// Code generated by go-swagger; DO NOT EDIT.

package todo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	modules "github.com/chinajuanbob/helloworld/pkg/gen/modules"
)

// AddTodoReader is a Reader for the AddTodo structure.
type AddTodoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTodoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddTodoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddTodoOK creates a AddTodoOK with default headers values
func NewAddTodoOK() *AddTodoOK {
	return &AddTodoOK{}
}

/*AddTodoOK handles this case with default header values.

successful operation
*/
type AddTodoOK struct {
	Payload *modules.ServiceTodoResult
}

func (o *AddTodoOK) Error() string {
	return fmt.Sprintf("[POST /todo][%d] addTodoOK  %+v", 200, o.Payload)
}

func (o *AddTodoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(modules.ServiceTodoResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
