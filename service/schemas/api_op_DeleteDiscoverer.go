// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteDiscovererInput struct {
	_ struct{} `type:"structure"`

	// DiscovererId is a required field
	DiscovererId *string `location:"uri" locationName:"discovererId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDiscovererInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDiscovererInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDiscovererInput"}

	if s.DiscovererId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiscovererId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDiscovererInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DiscovererId != nil {
		v := *s.DiscovererId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "discovererId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteDiscovererOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDiscovererOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDiscovererOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDiscoverer = "DeleteDiscoverer"

// DeleteDiscovererRequest returns a request value for making API operation for
// Schemas.
//
// Deletes a discoverer.
//
//    // Example sending a request using DeleteDiscovererRequest.
//    req := client.DeleteDiscovererRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/DeleteDiscoverer
func (c *Client) DeleteDiscovererRequest(input *DeleteDiscovererInput) DeleteDiscovererRequest {
	op := &aws.Operation{
		Name:       opDeleteDiscoverer,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/discoverers/id/{discovererId}",
	}

	if input == nil {
		input = &DeleteDiscovererInput{}
	}

	req := c.newRequest(op, input, &DeleteDiscovererOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteDiscovererRequest{Request: req, Input: input, Copy: c.DeleteDiscovererRequest}
}

// DeleteDiscovererRequest is the request type for the
// DeleteDiscoverer API operation.
type DeleteDiscovererRequest struct {
	*aws.Request
	Input *DeleteDiscovererInput
	Copy  func(*DeleteDiscovererInput) DeleteDiscovererRequest
}

// Send marshals and sends the DeleteDiscoverer API request.
func (r DeleteDiscovererRequest) Send(ctx context.Context) (*DeleteDiscovererResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDiscovererResponse{
		DeleteDiscovererOutput: r.Request.Data.(*DeleteDiscovererOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDiscovererResponse is the response type for the
// DeleteDiscoverer API operation.
type DeleteDiscovererResponse struct {
	*DeleteDiscovererOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDiscoverer request.
func (r *DeleteDiscovererResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
