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

// GroupV2GetGroupReader is a Reader for the GroupV2GetGroup structure.
type GroupV2GetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupV2GetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGroupV2GetGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupV2GetGroupOK creates a GroupV2GetGroupOK with default headers values
func NewGroupV2GetGroupOK() *GroupV2GetGroupOK {
	return &GroupV2GetGroupOK{}
}

/*GroupV2GetGroupOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type GroupV2GetGroupOK struct {
	Payload *models.GroupV2GetGroupOKBody
}

func (o *GroupV2GetGroupOK) Error() string {
	return fmt.Sprintf("[GET /GroupV2/{groupId}/][%d] groupV2GetGroupOK  %+v", 200, o.Payload)
}

func (o *GroupV2GetGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupV2GetGroupOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}