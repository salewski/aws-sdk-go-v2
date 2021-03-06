// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudhsm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteLunaClientInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the client to delete.
	//
	// ClientArn is a required field
	ClientArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLunaClientInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLunaClientInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLunaClientInput"}

	if s.ClientArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteLunaClientOutput struct {
	_ struct{} `type:"structure"`

	// The status of the action.
	//
	// Status is a required field
	Status *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLunaClientOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLunaClient = "DeleteLunaClient"

// DeleteLunaClientRequest returns a request value for making API operation for
// Amazon CloudHSM.
//
// This is documentation for AWS CloudHSM Classic. For more information, see
// AWS CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/),
// the AWS CloudHSM Classic User Guide (http://docs.aws.amazon.com/cloudhsm/classic/userguide/),
// and the AWS CloudHSM Classic API Reference (http://docs.aws.amazon.com/cloudhsm/classic/APIReference/).
//
// For information about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide (http://docs.aws.amazon.com/cloudhsm/latest/userguide/),
// and the AWS CloudHSM API Reference (http://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
//
// Deletes a client.
//
//    // Example sending a request using DeleteLunaClientRequest.
//    req := client.DeleteLunaClientRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudhsm-2014-05-30/DeleteLunaClient
func (c *Client) DeleteLunaClientRequest(input *DeleteLunaClientInput) DeleteLunaClientRequest {
	op := &aws.Operation{
		Name:       opDeleteLunaClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLunaClientInput{}
	}

	req := c.newRequest(op, input, &DeleteLunaClientOutput{})

	return DeleteLunaClientRequest{Request: req, Input: input, Copy: c.DeleteLunaClientRequest}
}

// DeleteLunaClientRequest is the request type for the
// DeleteLunaClient API operation.
type DeleteLunaClientRequest struct {
	*aws.Request
	Input *DeleteLunaClientInput
	Copy  func(*DeleteLunaClientInput) DeleteLunaClientRequest
}

// Send marshals and sends the DeleteLunaClient API request.
func (r DeleteLunaClientRequest) Send(ctx context.Context) (*DeleteLunaClientResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLunaClientResponse{
		DeleteLunaClientOutput: r.Request.Data.(*DeleteLunaClientOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLunaClientResponse is the response type for the
// DeleteLunaClient API operation.
type DeleteLunaClientResponse struct {
	*DeleteLunaClientOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLunaClient request.
func (r *DeleteLunaClientResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
