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

// GroupV2IndividualGroupInviteCancelReader is a Reader for the GroupV2IndividualGroupInviteCancel structure.
type GroupV2IndividualGroupInviteCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupV2IndividualGroupInviteCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGroupV2IndividualGroupInviteCancelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupV2IndividualGroupInviteCancelOK creates a GroupV2IndividualGroupInviteCancelOK with default headers values
func NewGroupV2IndividualGroupInviteCancelOK() *GroupV2IndividualGroupInviteCancelOK {
	return &GroupV2IndividualGroupInviteCancelOK{}
}

/*GroupV2IndividualGroupInviteCancelOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type GroupV2IndividualGroupInviteCancelOK struct {
	Payload *models.GroupV2IndividualGroupInviteCancelOKBody
}

func (o *GroupV2IndividualGroupInviteCancelOK) Error() string {
	return fmt.Sprintf("[POST /GroupV2/{groupId}/Members/IndividualInviteCancel/{membershipType}/{membershipId}/][%d] groupV2IndividualGroupInviteCancelOK  %+v", 200, o.Payload)
}

func (o *GroupV2IndividualGroupInviteCancelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupV2IndividualGroupInviteCancelOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
