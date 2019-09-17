// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// JoinDomainInput
type JoinDomainInput struct {
	_ struct{} `type:"structure"`

	// List of IPv4 addresses, NetBIOS names, or host names of your domain server.
	// If you need to specify the port number include it after the colon (“:”).
	// For example, mydc.mydomain.com:389.
	DomainControllers []string `type:"list"`

	// The name of the domain that you want the gateway to join.
	//
	// DomainName is a required field
	DomainName *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// The organizational unit (OU) is a container in an Active Directory that can
	// hold users, groups, computers, and other OUs and this parameter specifies
	// the OU that the gateway will join within the AD domain.
	OrganizationalUnit *string `min:"1" type:"string"`

	// Sets the password of the user who has permission to add the gateway to the
	// Active Directory domain.
	//
	// Password is a required field
	Password *string `min:"1" type:"string" required:"true"`

	// Sets the user name of user who has permission to add the gateway to the Active
	// Directory domain.
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s JoinDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *JoinDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "JoinDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 1))
	}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if s.OrganizationalUnit != nil && len(*s.OrganizationalUnit) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationalUnit", 1))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}
	if s.Password != nil && len(*s.Password) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Password", 1))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// JoinDomainOutput
type JoinDomainOutput struct {
	_ struct{} `type:"structure"`

	// The unique Amazon Resource Name (ARN) of the gateway that joined the domain.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s JoinDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opJoinDomain = "JoinDomain"

// JoinDomainRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Adds a file gateway to an Active Directory domain. This operation is only
// supported for file gateways that support the SMB file protocol.
//
//    // Example sending a request using JoinDomainRequest.
//    req := client.JoinDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/JoinDomain
func (c *Client) JoinDomainRequest(input *JoinDomainInput) JoinDomainRequest {
	op := &aws.Operation{
		Name:       opJoinDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &JoinDomainInput{}
	}

	req := c.newRequest(op, input, &JoinDomainOutput{})
	return JoinDomainRequest{Request: req, Input: input, Copy: c.JoinDomainRequest}
}

// JoinDomainRequest is the request type for the
// JoinDomain API operation.
type JoinDomainRequest struct {
	*aws.Request
	Input *JoinDomainInput
	Copy  func(*JoinDomainInput) JoinDomainRequest
}

// Send marshals and sends the JoinDomain API request.
func (r JoinDomainRequest) Send(ctx context.Context) (*JoinDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &JoinDomainResponse{
		JoinDomainOutput: r.Request.Data.(*JoinDomainOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// JoinDomainResponse is the response type for the
// JoinDomain API operation.
type JoinDomainResponse struct {
	*JoinDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// JoinDomain request.
func (r *JoinDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
