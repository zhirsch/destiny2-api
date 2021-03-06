// Code generated by go-swagger; DO NOT EDIT.

package forum

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/zhirsch/destiny2-api/models"
)

// ForumApproveFireteamThreadReader is a Reader for the ForumApproveFireteamThread structure.
type ForumApproveFireteamThreadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ForumApproveFireteamThreadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewForumApproveFireteamThreadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewForumApproveFireteamThreadOK creates a ForumApproveFireteamThreadOK with default headers values
func NewForumApproveFireteamThreadOK() *ForumApproveFireteamThreadOK {
	return &ForumApproveFireteamThreadOK{}
}

/*ForumApproveFireteamThreadOK handles this case with default header values.

Look at the Response property for more information about the nature of this response
*/
type ForumApproveFireteamThreadOK struct {
	Payload *models.ForumApproveFireteamThreadOKBody
}

func (o *ForumApproveFireteamThreadOK) Error() string {
	return fmt.Sprintf("[POST /Forum/Recruit/Approve/{topicId}/][%d] forumApproveFireteamThreadOK  %+v", 200, o.Payload)
}

func (o *ForumApproveFireteamThreadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForumApproveFireteamThreadOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
