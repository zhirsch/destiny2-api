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

// CommunityContentGetCommunityContentReader is a Reader for the CommunityContentGetCommunityContent structure.
type CommunityContentGetCommunityContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommunityContentGetCommunityContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCommunityContentGetCommunityContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCommunityContentGetCommunityContentOK creates a CommunityContentGetCommunityContentOK with default headers values
func NewCommunityContentGetCommunityContentOK() *CommunityContentGetCommunityContentOK {
	return &CommunityContentGetCommunityContentOK{}
}

/*CommunityContentGetCommunityContentOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type CommunityContentGetCommunityContentOK struct {
	Payload *models.CommunityContentGetCommunityContentOKBody
}

func (o *CommunityContentGetCommunityContentOK) Error() string {
	return fmt.Sprintf("[GET /CommunityContent/Get/{sort}/{mediaFilter}/{page}/][%d] communityContentGetCommunityContentOK  %+v", 200, o.Payload)
}

func (o *CommunityContentGetCommunityContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommunityContentGetCommunityContentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
