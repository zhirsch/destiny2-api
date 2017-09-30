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

// Destiny2GetDestinyEntityDefinitionReader is a Reader for the Destiny2GetDestinyEntityDefinition structure.
type Destiny2GetDestinyEntityDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Destiny2GetDestinyEntityDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDestiny2GetDestinyEntityDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestiny2GetDestinyEntityDefinitionOK creates a Destiny2GetDestinyEntityDefinitionOK with default headers values
func NewDestiny2GetDestinyEntityDefinitionOK() *Destiny2GetDestinyEntityDefinitionOK {
	return &Destiny2GetDestinyEntityDefinitionOK{}
}

/*Destiny2GetDestinyEntityDefinitionOK handles this case with default header values.

Provides common properties for destiny definitions.
*/
type Destiny2GetDestinyEntityDefinitionOK struct {
	Payload *models.Destiny2GetDestinyEntityDefinitionOKBody
}

func (o *Destiny2GetDestinyEntityDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /Destiny2/Manifest/{entityType}/{hashIdentifier}/][%d] destiny2GetDestinyEntityDefinitionOK  %+v", 200, o.Payload)
}

func (o *Destiny2GetDestinyEntityDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Destiny2GetDestinyEntityDefinitionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
