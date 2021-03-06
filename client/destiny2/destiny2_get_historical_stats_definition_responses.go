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

// Destiny2GetHistoricalStatsDefinitionReader is a Reader for the Destiny2GetHistoricalStatsDefinition structure.
type Destiny2GetHistoricalStatsDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Destiny2GetHistoricalStatsDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDestiny2GetHistoricalStatsDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestiny2GetHistoricalStatsDefinitionOK creates a Destiny2GetHistoricalStatsDefinitionOK with default headers values
func NewDestiny2GetHistoricalStatsDefinitionOK() *Destiny2GetHistoricalStatsDefinitionOK {
	return &Destiny2GetHistoricalStatsDefinitionOK{}
}

/*Destiny2GetHistoricalStatsDefinitionOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type Destiny2GetHistoricalStatsDefinitionOK struct {
	Payload *models.Destiny2GetHistoricalStatsDefinitionOKBody
}

func (o *Destiny2GetHistoricalStatsDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /Destiny2/Stats/Definition/][%d] destiny2GetHistoricalStatsDefinitionOK  %+v", 200, o.Payload)
}

func (o *Destiny2GetHistoricalStatsDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Destiny2GetHistoricalStatsDefinitionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
