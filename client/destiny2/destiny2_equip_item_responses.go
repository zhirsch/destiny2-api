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

// Destiny2EquipItemReader is a Reader for the Destiny2EquipItem structure.
type Destiny2EquipItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Destiny2EquipItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDestiny2EquipItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestiny2EquipItemOK creates a Destiny2EquipItemOK with default headers values
func NewDestiny2EquipItemOK() *Destiny2EquipItemOK {
	return &Destiny2EquipItemOK{}
}

/*Destiny2EquipItemOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type Destiny2EquipItemOK struct {
	Payload *models.Destiny2EquipItemOKBody
}

func (o *Destiny2EquipItemOK) Error() string {
	return fmt.Sprintf("[POST /Destiny2/Actions/Items/EquipItem/][%d] destiny2EquipItemOK  %+v", 200, o.Payload)
}

func (o *Destiny2EquipItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Destiny2EquipItemOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
