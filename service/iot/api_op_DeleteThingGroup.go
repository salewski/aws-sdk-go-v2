// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteThingGroupInput struct {
	_ struct{} `type:"structure"`

	// The expected version of the thing group to delete.
	ExpectedVersion *int64 `location:"querystring" locationName:"expectedVersion" type:"long"`

	// The name of the thing group to delete.
	//
	// ThingGroupName is a required field
	ThingGroupName *string `location:"uri" locationName:"thingGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteThingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteThingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteThingGroupInput"}

	if s.ThingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingGroupName"))
	}
	if s.ThingGroupName != nil && len(*s.ThingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteThingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ThingGroupName != nil {
		v := *s.ThingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExpectedVersion != nil {
		v := *s.ExpectedVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "expectedVersion", protocol.Int64Value(v), metadata)
	}
	return nil
}

type DeleteThingGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteThingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteThingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteThingGroup = "DeleteThingGroup"

// DeleteThingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes a thing group.
//
//    // Example sending a request using DeleteThingGroupRequest.
//    req := client.DeleteThingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteThingGroupRequest(input *DeleteThingGroupInput) DeleteThingGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteThingGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/thing-groups/{thingGroupName}",
	}

	if input == nil {
		input = &DeleteThingGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteThingGroupOutput{})

	return DeleteThingGroupRequest{Request: req, Input: input, Copy: c.DeleteThingGroupRequest}
}

// DeleteThingGroupRequest is the request type for the
// DeleteThingGroup API operation.
type DeleteThingGroupRequest struct {
	*aws.Request
	Input *DeleteThingGroupInput
	Copy  func(*DeleteThingGroupInput) DeleteThingGroupRequest
}

// Send marshals and sends the DeleteThingGroup API request.
func (r DeleteThingGroupRequest) Send(ctx context.Context) (*DeleteThingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteThingGroupResponse{
		DeleteThingGroupOutput: r.Request.Data.(*DeleteThingGroupOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteThingGroupResponse is the response type for the
// DeleteThingGroup API operation.
type DeleteThingGroupResponse struct {
	*DeleteThingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteThingGroup request.
func (r *DeleteThingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
