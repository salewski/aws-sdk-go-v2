// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateIPSetInput struct {
	_ struct{} `type:"structure"`

	// The updated Boolean value that specifies whether the IPSet is active or not.
	Activate *bool `locationName:"activate" type:"boolean"`

	// The detectorID that specifies the GuardDuty service whose IPSet you want
	// to update.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// The unique ID that specifies the IPSet that you want to update.
	//
	// IpSetId is a required field
	IpSetId *string `location:"uri" locationName:"ipSetId" type:"string" required:"true"`

	// The updated URI of the file that contains the IPSet. For example: https://s3.us-west-2.amazonaws.com/my-bucket/my-object-key.
	Location *string `locationName:"location" min:"1" type:"string"`

	// The unique ID that specifies the IPSet that you want to update.
	Name *string `locationName:"name" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateIPSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateIPSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateIPSetInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if s.IpSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IpSetId"))
	}
	if s.Location != nil && len(*s.Location) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Location", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateIPSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Activate != nil {
		v := *s.Activate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "activate", protocol.BoolValue(v), metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "location", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IpSetId != nil {
		v := *s.IpSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ipSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateIPSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateIPSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateIPSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateIPSet = "UpdateIPSet"

// UpdateIPSetRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Updates the IPSet specified by the IPSet ID.
//
//    // Example sending a request using UpdateIPSetRequest.
//    req := client.UpdateIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/UpdateIPSet
func (c *Client) UpdateIPSetRequest(input *UpdateIPSetInput) UpdateIPSetRequest {
	op := &aws.Operation{
		Name:       opUpdateIPSet,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/ipset/{ipSetId}",
	}

	if input == nil {
		input = &UpdateIPSetInput{}
	}

	req := c.newRequest(op, input, &UpdateIPSetOutput{})

	return UpdateIPSetRequest{Request: req, Input: input, Copy: c.UpdateIPSetRequest}
}

// UpdateIPSetRequest is the request type for the
// UpdateIPSet API operation.
type UpdateIPSetRequest struct {
	*aws.Request
	Input *UpdateIPSetInput
	Copy  func(*UpdateIPSetInput) UpdateIPSetRequest
}

// Send marshals and sends the UpdateIPSet API request.
func (r UpdateIPSetRequest) Send(ctx context.Context) (*UpdateIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateIPSetResponse{
		UpdateIPSetOutput: r.Request.Data.(*UpdateIPSetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateIPSetResponse is the response type for the
// UpdateIPSet API operation.
type UpdateIPSetResponse struct {
	*UpdateIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateIPSet request.
func (r *UpdateIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}