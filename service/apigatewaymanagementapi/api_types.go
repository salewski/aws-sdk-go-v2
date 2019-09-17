// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewaymanagementapi

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

type Identity struct {
	_ struct{} `type:"structure"`

	// The source IP address of the TCP connection making the request to API Gateway.
	//
	// SourceIp is a required field
	SourceIp *string `locationName:"sourceIp" type:"string" required:"true"`

	// The User Agent of the API caller.
	//
	// UserAgent is a required field
	UserAgent *string `locationName:"userAgent" type:"string" required:"true"`
}

// String returns the string representation
func (s Identity) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Identity) MarshalFields(e protocol.FieldEncoder) error {
	if s.SourceIp != nil {
		v := *s.SourceIp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sourceIp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserAgent != nil {
		v := *s.UserAgent

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "userAgent", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}
