// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to set the attributes that control how bounce and complaint events
// are processed.
type PutEmailIdentityFeedbackAttributesInput struct {
	_ struct{} `type:"structure"`

	// Sets the feedback forwarding configuration for the identity.
	//
	// If the value is true, Amazon Pinpoint sends you email notifications when
	// bounce or complaint events occur. Amazon Pinpoint sends this notification
	// to the address that you specified in the Return-Path header of the original
	// email.
	//
	// When you set this value to false, Amazon Pinpoint sends notifications through
	// other mechanisms, such as by notifying an Amazon SNS topic or another event
	// destination. You're required to have a method of tracking bounces and complaints.
	// If you haven't set up another mechanism for receiving bounce or complaint
	// notifications, Amazon Pinpoint sends an email notification when these events
	// occur (even if this setting is disabled).
	EmailForwardingEnabled *bool `type:"boolean"`

	// The email identity that you want to configure bounce and complaint feedback
	// forwarding for.
	//
	// EmailIdentity is a required field
	EmailIdentity *string `location:"uri" locationName:"EmailIdentity" type:"string" required:"true"`
}

// String returns the string representation
func (s PutEmailIdentityFeedbackAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEmailIdentityFeedbackAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEmailIdentityFeedbackAttributesInput"}

	if s.EmailIdentity == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailIdentity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEmailIdentityFeedbackAttributesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.EmailForwardingEnabled != nil {
		v := *s.EmailForwardingEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EmailForwardingEnabled", protocol.BoolValue(v), metadata)
	}
	if s.EmailIdentity != nil {
		v := *s.EmailIdentity

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EmailIdentity", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutEmailIdentityFeedbackAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutEmailIdentityFeedbackAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutEmailIdentityFeedbackAttributesOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutEmailIdentityFeedbackAttributes = "PutEmailIdentityFeedbackAttributes"

// PutEmailIdentityFeedbackAttributesRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Used to enable or disable feedback forwarding for an identity. This setting
// determines what happens when an identity is used to send an email that results
// in a bounce or complaint event.
//
// When you enable feedback forwarding, Amazon Pinpoint sends you email notifications
// when bounce or complaint events occur. Amazon Pinpoint sends this notification
// to the address that you specified in the Return-Path header of the original
// email.
//
// When you disable feedback forwarding, Amazon Pinpoint sends notifications
// through other mechanisms, such as by notifying an Amazon SNS topic. You're
// required to have a method of tracking bounces and complaints. If you haven't
// set up another mechanism for receiving bounce or complaint notifications,
// Amazon Pinpoint sends an email notification when these events occur (even
// if this setting is disabled).
//
//    // Example sending a request using PutEmailIdentityFeedbackAttributesRequest.
//    req := client.PutEmailIdentityFeedbackAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutEmailIdentityFeedbackAttributes
func (c *Client) PutEmailIdentityFeedbackAttributesRequest(input *PutEmailIdentityFeedbackAttributesInput) PutEmailIdentityFeedbackAttributesRequest {
	op := &aws.Operation{
		Name:       opPutEmailIdentityFeedbackAttributes,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/identities/{EmailIdentity}/feedback",
	}

	if input == nil {
		input = &PutEmailIdentityFeedbackAttributesInput{}
	}

	req := c.newRequest(op, input, &PutEmailIdentityFeedbackAttributesOutput{})

	return PutEmailIdentityFeedbackAttributesRequest{Request: req, Input: input, Copy: c.PutEmailIdentityFeedbackAttributesRequest}
}

// PutEmailIdentityFeedbackAttributesRequest is the request type for the
// PutEmailIdentityFeedbackAttributes API operation.
type PutEmailIdentityFeedbackAttributesRequest struct {
	*aws.Request
	Input *PutEmailIdentityFeedbackAttributesInput
	Copy  func(*PutEmailIdentityFeedbackAttributesInput) PutEmailIdentityFeedbackAttributesRequest
}

// Send marshals and sends the PutEmailIdentityFeedbackAttributes API request.
func (r PutEmailIdentityFeedbackAttributesRequest) Send(ctx context.Context) (*PutEmailIdentityFeedbackAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutEmailIdentityFeedbackAttributesResponse{
		PutEmailIdentityFeedbackAttributesOutput: r.Request.Data.(*PutEmailIdentityFeedbackAttributesOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutEmailIdentityFeedbackAttributesResponse is the response type for the
// PutEmailIdentityFeedbackAttributes API operation.
type PutEmailIdentityFeedbackAttributesResponse struct {
	*PutEmailIdentityFeedbackAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutEmailIdentityFeedbackAttributes request.
func (r *PutEmailIdentityFeedbackAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}