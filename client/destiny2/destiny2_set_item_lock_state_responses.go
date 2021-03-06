// Code generated by go-swagger; DO NOT EDIT.

package destiny2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/zhirsch/destiny2-api/models"
)

// Destiny2SetItemLockStateReader is a Reader for the Destiny2SetItemLockState structure.
type Destiny2SetItemLockStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Destiny2SetItemLockStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDestiny2SetItemLockStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestiny2SetItemLockStateOK creates a Destiny2SetItemLockStateOK with default headers values
func NewDestiny2SetItemLockStateOK() *Destiny2SetItemLockStateOK {
	return &Destiny2SetItemLockStateOK{}
}

/*Destiny2SetItemLockStateOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type Destiny2SetItemLockStateOK struct {
	Payload *models.Destiny2SetItemLockStateOKBody
}

func (o *Destiny2SetItemLockStateOK) Error() string {
	return fmt.Sprintf("[POST /Destiny2/Actions/Items/SetLockState/][%d] destiny2SetItemLockStateOK  %+v", 200, o.Payload)
}

func (o *Destiny2SetItemLockStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Destiny2SetItemLockStateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
