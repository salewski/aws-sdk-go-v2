// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSMBSettingsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSMBSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSMBSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSMBSettingsInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSMBSettingsOutput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that the gateway is joined to.
	DomainName *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`

	// This value is true if a password for the guest user “smbguest” is set,
	// and otherwise false.
	SMBGuestPasswordSet *bool `type:"boolean"`

	// The type of security strategy that was specified for file gateway.
	//
	// ClientSpecified: if you use this option, requests are established based on
	// what is negotiated by the client. This option is recommended when you want
	// to maximize compatibility across different clients in your environment.
	//
	// MandatorySigning: if you use this option, file gateway only allows connections
	// from SMBv2 or SMBv3 clients that have signing enabled. This option works
	// with SMB clients on Microsoft Windows Vista, Windows Server 2008 or newer.
	//
	// MandatoryEncryption: if you use this option, file gateway only allows connections
	// from SMBv3 clients that have encryption enabled. This option is highly recommended
	// for environments that handle sensitive data. This option works with SMB clients
	// on Microsoft Windows 8, Windows Server 2012 or newer.
	SMBSecurityStrategy SMBSecurityStrategy `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeSMBSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSMBSettings = "DescribeSMBSettings"

// DescribeSMBSettingsRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Gets a description of a Server Message Block (SMB) file share settings from
// a file gateway. This operation is only supported for file gateways.
//
//    // Example sending a request using DescribeSMBSettingsRequest.
//    req := client.DescribeSMBSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeSMBSettings
func (c *Client) DescribeSMBSettingsRequest(input *DescribeSMBSettingsInput) DescribeSMBSettingsRequest {
	op := &aws.Operation{
		Name:       opDescribeSMBSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSMBSettingsInput{}
	}

	req := c.newRequest(op, input, &DescribeSMBSettingsOutput{})
	return DescribeSMBSettingsRequest{Request: req, Input: input, Copy: c.DescribeSMBSettingsRequest}
}

// DescribeSMBSettingsRequest is the request type for the
// DescribeSMBSettings API operation.
type DescribeSMBSettingsRequest struct {
	*aws.Request
	Input *DescribeSMBSettingsInput
	Copy  func(*DescribeSMBSettingsInput) DescribeSMBSettingsRequest
}

// Send marshals and sends the DescribeSMBSettings API request.
func (r DescribeSMBSettingsRequest) Send(ctx context.Context) (*DescribeSMBSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSMBSettingsResponse{
		DescribeSMBSettingsOutput: r.Request.Data.(*DescribeSMBSettingsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSMBSettingsResponse is the response type for the
// DescribeSMBSettings API operation.
type DescribeSMBSettingsResponse struct {
	*DescribeSMBSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSMBSettings request.
func (r *DescribeSMBSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
