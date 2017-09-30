// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/zhirsch/destiny2-api/models"
)

// Destiny2GetVendorReader is a Reader for the Destiny2GetVendor structure.
type Destiny2GetVendorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Destiny2GetVendorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDestiny2GetVendorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDestiny2GetVendorOK creates a Destiny2GetVendorOK with default headers values
func NewDestiny2GetVendorOK() *Destiny2GetVendorOK {
	return &Destiny2GetVendorOK{}
}

/*Destiny2GetVendorOK handles this case with default header values.

A response containing all of the components for a vendor.
*/
type Destiny2GetVendorOK struct {
	Payload *models.Destiny2GetVendorOKBody
}

func (o *Destiny2GetVendorOK) Error() string {
	return fmt.Sprintf("[GET /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/Vendors/{vendorHash}/][%d] destiny2GetVendorOK  %+v", 200, o.Payload)
}

func (o *Destiny2GetVendorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Destiny2GetVendorOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}