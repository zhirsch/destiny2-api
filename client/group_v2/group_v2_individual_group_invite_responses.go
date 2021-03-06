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

// GroupV2IndividualGroupInviteReader is a Reader for the GroupV2IndividualGroupInvite structure.
type GroupV2IndividualGroupInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupV2IndividualGroupInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGroupV2IndividualGroupInviteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupV2IndividualGroupInviteOK creates a GroupV2IndividualGroupInviteOK with default headers values
func NewGroupV2IndividualGroupInviteOK() *GroupV2IndividualGroupInviteOK {
	return &GroupV2IndividualGroupInviteOK{}
}

/*GroupV2IndividualGroupInviteOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type GroupV2IndividualGroupInviteOK struct {
	Payload *models.GroupV2IndividualGroupInviteOKBody
}

func (o *GroupV2IndividualGroupInviteOK) Error() string {
	return fmt.Sprintf("[POST /GroupV2/{groupId}/Members/IndividualInvite/{membershipType}/{membershipId}/][%d] groupV2IndividualGroupInviteOK  %+v", 200, o.Payload)
}

func (o *GroupV2IndividualGroupInviteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupV2IndividualGroupInviteOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
