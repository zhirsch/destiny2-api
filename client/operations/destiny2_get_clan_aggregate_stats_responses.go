// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/zhirsch/destiny2-api/models"
)

// Destiny2GetClanAggregateStatsReader is a Reader for the Destiny2GetClanAggregateStats structure.
type Destiny2GetClanAggregateStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Destiny2GetClanAggregateStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDestiny2GetClanAggregateStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestiny2GetClanAggregateStatsOK creates a Destiny2GetClanAggregateStatsOK with default headers values
func NewDestiny2GetClanAggregateStatsOK() *Destiny2GetClanAggregateStatsOK {
	return &Destiny2GetClanAggregateStatsOK{}
}

/*Destiny2GetClanAggregateStatsOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type Destiny2GetClanAggregateStatsOK struct {
	Payload *models.Destiny2GetClanAggregateStatsOKBody
}

func (o *Destiny2GetClanAggregateStatsOK) Error() string {
	return fmt.Sprintf("[GET /Destiny2/Stats/AggregateClanStats/{groupId}/][%d] destiny2GetClanAggregateStatsOK  %+v", 200, o.Payload)
}

func (o *Destiny2GetClanAggregateStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Destiny2GetClanAggregateStatsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}