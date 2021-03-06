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

// Destiny2GetHistoricalStatsForAccountReader is a Reader for the Destiny2GetHistoricalStatsForAccount structure.
type Destiny2GetHistoricalStatsForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Destiny2GetHistoricalStatsForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDestiny2GetHistoricalStatsForAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestiny2GetHistoricalStatsForAccountOK creates a Destiny2GetHistoricalStatsForAccountOK with default headers values
func NewDestiny2GetHistoricalStatsForAccountOK() *Destiny2GetHistoricalStatsForAccountOK {
	return &Destiny2GetHistoricalStatsForAccountOK{}
}

/*Destiny2GetHistoricalStatsForAccountOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type Destiny2GetHistoricalStatsForAccountOK struct {
	Payload *models.Destiny2GetHistoricalStatsForAccountOKBody
}

func (o *Destiny2GetHistoricalStatsForAccountOK) Error() string {
	return fmt.Sprintf("[GET /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/][%d] destiny2GetHistoricalStatsForAccountOK  %+v", 200, o.Payload)
}

func (o *Destiny2GetHistoricalStatsForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Destiny2GetHistoricalStatsForAccountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
