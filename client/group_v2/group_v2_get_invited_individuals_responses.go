// Code generated by go-swagger; DO NOT EDIT.

package group_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/zhirsch/destiny2-api/models"
)

// GroupV2GetInvitedIndividualsReader is a Reader for the GroupV2GetInvitedIndividuals structure.
type GroupV2GetInvitedIndividualsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupV2GetInvitedIndividualsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGroupV2GetInvitedIndividualsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupV2GetInvitedIndividualsOK creates a GroupV2GetInvitedIndividualsOK with default headers values
func NewGroupV2GetInvitedIndividualsOK() *GroupV2GetInvitedIndividualsOK {
	return &GroupV2GetInvitedIndividualsOK{}
}

/*GroupV2GetInvitedIndividualsOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type GroupV2GetInvitedIndividualsOK struct {
	Payload *models.GroupV2GetInvitedIndividualsOKBody
}

func (o *GroupV2GetInvitedIndividualsOK) Error() string {
	return fmt.Sprintf("[GET /GroupV2/{groupId}/Members/InvitedIndividuals/][%d] groupV2GetInvitedIndividualsOK  %+v", 200, o.Payload)
}

func (o *GroupV2GetInvitedIndividualsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupV2GetInvitedIndividualsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
