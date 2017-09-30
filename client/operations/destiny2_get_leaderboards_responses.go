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

// Destiny2GetLeaderboardsReader is a Reader for the Destiny2GetLeaderboards structure.
type Destiny2GetLeaderboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Destiny2GetLeaderboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDestiny2GetLeaderboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestiny2GetLeaderboardsOK creates a Destiny2GetLeaderboardsOK with default headers values
func NewDestiny2GetLeaderboardsOK() *Destiny2GetLeaderboardsOK {
	return &Destiny2GetLeaderboardsOK{}
}

/*Destiny2GetLeaderboardsOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type Destiny2GetLeaderboardsOK struct {
	Payload *models.Destiny2GetLeaderboardsOKBody
}

func (o *Destiny2GetLeaderboardsOK) Error() string {
	return fmt.Sprintf("[GET /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/Leaderboards/][%d] destiny2GetLeaderboardsOK  %+v", 200, o.Payload)
}

func (o *Destiny2GetLeaderboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Destiny2GetLeaderboardsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
