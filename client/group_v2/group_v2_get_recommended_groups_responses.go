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

// GroupV2GetRecommendedGroupsReader is a Reader for the GroupV2GetRecommendedGroups structure.
type GroupV2GetRecommendedGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupV2GetRecommendedGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGroupV2GetRecommendedGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupV2GetRecommendedGroupsOK creates a GroupV2GetRecommendedGroupsOK with default headers values
func NewGroupV2GetRecommendedGroupsOK() *GroupV2GetRecommendedGroupsOK {
	return &GroupV2GetRecommendedGroupsOK{}
}

/*GroupV2GetRecommendedGroupsOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type GroupV2GetRecommendedGroupsOK struct {
	Payload *models.GroupV2GetRecommendedGroupsOKBody
}

func (o *GroupV2GetRecommendedGroupsOK) Error() string {
	return fmt.Sprintf("[POST /GroupV2/Recommended/][%d] groupV2GetRecommendedGroupsOK  %+v", 200, o.Payload)
}

func (o *GroupV2GetRecommendedGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupV2GetRecommendedGroupsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
