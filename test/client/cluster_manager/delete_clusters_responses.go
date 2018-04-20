// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DeleteClustersReader is a Reader for the DeleteClusters structure.
type DeleteClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteClustersOK creates a DeleteClustersOK with default headers values
func NewDeleteClustersOK() *DeleteClustersOK {
	return &DeleteClustersOK{}
}

/*DeleteClustersOK handles this case with default header values.

DeleteClustersOK delete clusters o k
*/
type DeleteClustersOK struct {
	Payload *models.OpenpitrixDeleteClustersResponse
}

func (o *DeleteClustersOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/delete][%d] deleteClustersOK  %+v", 200, o.Payload)
}

func (o *DeleteClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDeleteClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
