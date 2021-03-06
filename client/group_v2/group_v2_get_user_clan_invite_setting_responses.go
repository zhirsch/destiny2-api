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

// GroupV2GetUserClanInviteSettingReader is a Reader for the GroupV2GetUserClanInviteSetting structure.
type GroupV2GetUserClanInviteSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupV2GetUserClanInviteSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGroupV2GetUserClanInviteSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupV2GetUserClanInviteSettingOK creates a GroupV2GetUserClanInviteSettingOK with default headers values
func NewGroupV2GetUserClanInviteSettingOK() *GroupV2GetUserClanInviteSettingOK {
	return &GroupV2GetUserClanInviteSettingOK{}
}

/*GroupV2GetUserClanInviteSettingOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type GroupV2GetUserClanInviteSettingOK struct {
	Payload *models.GroupV2GetUserClanInviteSettingOKBody
}

func (o *GroupV2GetUserClanInviteSettingOK) Error() string {
	return fmt.Sprintf("[GET /GroupV2/GetUserClanInviteSetting/{mType}/][%d] groupV2GetUserClanInviteSettingOK  %+v", 200, o.Payload)
}

func (o *GroupV2GetUserClanInviteSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupV2GetUserClanInviteSettingOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
