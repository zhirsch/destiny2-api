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

// GroupV2EditOptionalConversationReader is a Reader for the GroupV2EditOptionalConversation structure.
type GroupV2EditOptionalConversationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupV2EditOptionalConversationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGroupV2EditOptionalConversationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupV2EditOptionalConversationOK creates a GroupV2EditOptionalConversationOK with default headers values
func NewGroupV2EditOptionalConversationOK() *GroupV2EditOptionalConversationOK {
	return &GroupV2EditOptionalConversationOK{}
}

/*GroupV2EditOptionalConversationOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type GroupV2EditOptionalConversationOK struct {
	Payload *models.GroupV2EditOptionalConversationOKBody
}

func (o *GroupV2EditOptionalConversationOK) Error() string {
	return fmt.Sprintf("[POST /GroupV2/{groupId}/OptionalConversations/Edit/{conversationId}/][%d] groupV2EditOptionalConversationOK  %+v", 200, o.Payload)
}

func (o *GroupV2EditOptionalConversationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupV2EditOptionalConversationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
