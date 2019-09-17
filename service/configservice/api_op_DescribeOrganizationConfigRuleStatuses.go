// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeOrganizationConfigRuleStatusesInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `type:"integer"`

	NextToken *string `type:"string"`

	OrganizationConfigRuleNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeOrganizationConfigRuleStatusesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeOrganizationConfigRuleStatusesOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `type:"string"`

	OrganizationConfigRuleStatuses []OrganizationConfigRuleStatus `type:"list"`
}

// String returns the string representation
func (s DescribeOrganizationConfigRuleStatusesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeOrganizationConfigRuleStatuses = "DescribeOrganizationConfigRuleStatuses"

// DescribeOrganizationConfigRuleStatusesRequest returns a request value for making API operation for
// AWS Config.
//
//    // Example sending a request using DescribeOrganizationConfigRuleStatusesRequest.
//    req := client.DescribeOrganizationConfigRuleStatusesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeOrganizationConfigRuleStatuses
func (c *Client) DescribeOrganizationConfigRuleStatusesRequest(input *DescribeOrganizationConfigRuleStatusesInput) DescribeOrganizationConfigRuleStatusesRequest {
	op := &aws.Operation{
		Name:       opDescribeOrganizationConfigRuleStatuses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeOrganizationConfigRuleStatusesInput{}
	}

	req := c.newRequest(op, input, &DescribeOrganizationConfigRuleStatusesOutput{})
	return DescribeOrganizationConfigRuleStatusesRequest{Request: req, Input: input, Copy: c.DescribeOrganizationConfigRuleStatusesRequest}
}

// DescribeOrganizationConfigRuleStatusesRequest is the request type for the
// DescribeOrganizationConfigRuleStatuses API operation.
type DescribeOrganizationConfigRuleStatusesRequest struct {
	*aws.Request
	Input *DescribeOrganizationConfigRuleStatusesInput
	Copy  func(*DescribeOrganizationConfigRuleStatusesInput) DescribeOrganizationConfigRuleStatusesRequest
}

// Send marshals and sends the DescribeOrganizationConfigRuleStatuses API request.
func (r DescribeOrganizationConfigRuleStatusesRequest) Send(ctx context.Context) (*DescribeOrganizationConfigRuleStatusesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOrganizationConfigRuleStatusesResponse{
		DescribeOrganizationConfigRuleStatusesOutput: r.Request.Data.(*DescribeOrganizationConfigRuleStatusesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeOrganizationConfigRuleStatusesResponse is the response type for the
// DescribeOrganizationConfigRuleStatuses API operation.
type DescribeOrganizationConfigRuleStatusesResponse struct {
	*DescribeOrganizationConfigRuleStatusesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOrganizationConfigRuleStatuses request.
func (r *DescribeOrganizationConfigRuleStatusesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}