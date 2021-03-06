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

// Destiny2EquipItemsReader is a Reader for the Destiny2EquipItems structure.
type Destiny2EquipItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Destiny2EquipItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDestiny2EquipItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestiny2EquipItemsOK creates a Destiny2EquipItemsOK with default headers values
func NewDestiny2EquipItemsOK() *Destiny2EquipItemsOK {
	return &Destiny2EquipItemsOK{}
}

/*Destiny2EquipItemsOK handles this case with default header values.

The results of a bulk Equipping operation performed through the Destiny API.
*/
type Destiny2EquipItemsOK struct {
	Payload *models.Destiny2EquipItemsOKBody
}

func (o *Destiny2EquipItemsOK) Error() string {
	return fmt.Sprintf("[POST /Destiny2/Actions/Items/EquipItems/][%d] destiny2EquipItemsOK  %+v", 200, o.Payload)
}

func (o *Destiny2EquipItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Destiny2EquipItemsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
