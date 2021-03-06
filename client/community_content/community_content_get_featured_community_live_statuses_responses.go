// Code generated by go-swagger; DO NOT EDIT.

package community_content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/zhirsch/destiny2-api/models"
)

// CommunityContentGetFeaturedCommunityLiveStatusesReader is a Reader for the CommunityContentGetFeaturedCommunityLiveStatuses structure.
type CommunityContentGetFeaturedCommunityLiveStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommunityContentGetFeaturedCommunityLiveStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCommunityContentGetFeaturedCommunityLiveStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCommunityContentGetFeaturedCommunityLiveStatusesOK creates a CommunityContentGetFeaturedCommunityLiveStatusesOK with default headers values
func NewCommunityContentGetFeaturedCommunityLiveStatusesOK() *CommunityContentGetFeaturedCommunityLiveStatusesOK {
	return &CommunityContentGetFeaturedCommunityLiveStatusesOK{}
}

/*CommunityContentGetFeaturedCommunityLiveStatusesOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type CommunityContentGetFeaturedCommunityLiveStatusesOK struct {
	Payload *models.CommunityContentGetFeaturedCommunityLiveStatusesOKBody
}

func (o *CommunityContentGetFeaturedCommunityLiveStatusesOK) Error() string {
	return fmt.Sprintf("[GET /CommunityContent/Live/Featured/{partnershipType}/{sort}/{page}/][%d] communityContentGetFeaturedCommunityLiveStatusesOK  %+v", 200, o.Payload)
}

func (o *CommunityContentGetFeaturedCommunityLiveStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommunityContentGetFeaturedCommunityLiveStatusesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
