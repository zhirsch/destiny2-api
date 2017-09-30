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

// GroupV2EditClanBannerReader is a Reader for the GroupV2EditClanBanner structure.
type GroupV2EditClanBannerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GroupV2EditClanBannerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGroupV2EditClanBannerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGroupV2EditClanBannerOK creates a GroupV2EditClanBannerOK with default headers values
func NewGroupV2EditClanBannerOK() *GroupV2EditClanBannerOK {
	return &GroupV2EditClanBannerOK{}
}

/*GroupV2EditClanBannerOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type GroupV2EditClanBannerOK struct {
	Payload *models.GroupV2EditClanBannerOKBody
}

func (o *GroupV2EditClanBannerOK) Error() string {
	return fmt.Sprintf("[POST /GroupV2/{groupId}/EditClanBanner/][%d] groupV2EditClanBannerOK  %+v", 200, o.Payload)
}

func (o *GroupV2EditClanBannerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupV2EditClanBannerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
