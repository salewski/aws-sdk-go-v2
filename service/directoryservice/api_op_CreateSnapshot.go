// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the CreateSnapshot operation.
type CreateSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory of which to take a snapshot.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The descriptive name to apply to the snapshot.
	Name *string `type:"string"`
}

// String returns the string representation
func (s CreateSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSnapshotInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the results of the CreateSnapshot operation.
type CreateSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the snapshot that was created.
	SnapshotId *string `type:"string"`
}

// String returns the string representation
func (s CreateSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSnapshot = "CreateSnapshot"

// CreateSnapshotRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Creates a snapshot of a Simple AD or Microsoft AD directory in the AWS cloud.
//
// You cannot take snapshots of AD Connector directories.
//
//    // Example sending a request using CreateSnapshotRequest.
//    req := client.CreateSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateSnapshot
func (c *Client) CreateSnapshotRequest(input *CreateSnapshotInput) CreateSnapshotRequest {
	op := &aws.Operation{
		Name:       opCreateSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSnapshotInput{}
	}

	req := c.newRequest(op, input, &CreateSnapshotOutput{})

	return CreateSnapshotRequest{Request: req, Input: input, Copy: c.CreateSnapshotRequest}
}

// CreateSnapshotRequest is the request type for the
// CreateSnapshot API operation.
type CreateSnapshotRequest struct {
	*aws.Request
	Input *CreateSnapshotInput
	Copy  func(*CreateSnapshotInput) CreateSnapshotRequest
}

// Send marshals and sends the CreateSnapshot API request.
func (r CreateSnapshotRequest) Send(ctx context.Context) (*CreateSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSnapshotResponse{
		CreateSnapshotOutput: r.Request.Data.(*CreateSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSnapshotResponse is the response type for the
// CreateSnapshot API operation.
type CreateSnapshotResponse struct {
	*CreateSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSnapshot request.
func (r *CreateSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}