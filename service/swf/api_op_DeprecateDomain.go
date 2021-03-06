// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeprecateDomainInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain to deprecate.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeprecateDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeprecateDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeprecateDomainInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeprecateDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeprecateDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeprecateDomain = "DeprecateDomain"

// DeprecateDomainRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Deprecates the specified domain. After a domain has been deprecated it cannot
// be used to create new workflow executions or register new types. However,
// you can still use visibility actions on this domain. Deprecating a domain
// also deprecates all activity and workflow types registered in the domain.
// Executions that were started before the domain was deprecated continues to
// run.
//
// This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using DeprecateDomainRequest.
//    req := client.DeprecateDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeprecateDomainRequest(input *DeprecateDomainInput) DeprecateDomainRequest {
	op := &aws.Operation{
		Name:       opDeprecateDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeprecateDomainInput{}
	}

	req := c.newRequest(op, input, &DeprecateDomainOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeprecateDomainRequest{Request: req, Input: input, Copy: c.DeprecateDomainRequest}
}

// DeprecateDomainRequest is the request type for the
// DeprecateDomain API operation.
type DeprecateDomainRequest struct {
	*aws.Request
	Input *DeprecateDomainInput
	Copy  func(*DeprecateDomainInput) DeprecateDomainRequest
}

// Send marshals and sends the DeprecateDomain API request.
func (r DeprecateDomainRequest) Send(ctx context.Context) (*DeprecateDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeprecateDomainResponse{
		DeprecateDomainOutput: r.Request.Data.(*DeprecateDomainOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeprecateDomainResponse is the response type for the
// DeprecateDomain API operation.
type DeprecateDomainResponse struct {
	*DeprecateDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeprecateDomain request.
func (r *DeprecateDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
