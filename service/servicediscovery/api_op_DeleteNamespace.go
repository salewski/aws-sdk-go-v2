// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteNamespaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the namespace that you want to delete.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNamespaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNamespaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNamespaceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteNamespaceOutput struct {
	_ struct{} `type:"structure"`

	// A value that you can use to determine whether the request completed successfully.
	// To get the status of the operation, see GetOperation (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html).
	OperationId *string `type:"string"`
}

// String returns the string representation
func (s DeleteNamespaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteNamespace = "DeleteNamespace"

// DeleteNamespaceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Deletes a namespace from the current account. If the namespace still contains
// one or more services, the request fails.
//
//    // Example sending a request using DeleteNamespaceRequest.
//    req := client.DeleteNamespaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/DeleteNamespace
func (c *Client) DeleteNamespaceRequest(input *DeleteNamespaceInput) DeleteNamespaceRequest {
	op := &aws.Operation{
		Name:       opDeleteNamespace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNamespaceInput{}
	}

	req := c.newRequest(op, input, &DeleteNamespaceOutput{})

	return DeleteNamespaceRequest{Request: req, Input: input, Copy: c.DeleteNamespaceRequest}
}

// DeleteNamespaceRequest is the request type for the
// DeleteNamespace API operation.
type DeleteNamespaceRequest struct {
	*aws.Request
	Input *DeleteNamespaceInput
	Copy  func(*DeleteNamespaceInput) DeleteNamespaceRequest
}

// Send marshals and sends the DeleteNamespace API request.
func (r DeleteNamespaceRequest) Send(ctx context.Context) (*DeleteNamespaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNamespaceResponse{
		DeleteNamespaceOutput: r.Request.Data.(*DeleteNamespaceOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNamespaceResponse is the response type for the
// DeleteNamespace API operation.
type DeleteNamespaceResponse struct {
	*DeleteNamespaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNamespace request.
func (r *DeleteNamespaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
