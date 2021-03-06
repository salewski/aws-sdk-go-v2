// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package health

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEntityAggregatesInput struct {
	_ struct{} `type:"structure"`

	// A list of event ARNs (unique identifiers). For example: "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	EventArns []string `locationName:"eventArns" min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeEntityAggregatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEntityAggregatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEntityAggregatesInput"}
	if s.EventArns != nil && len(s.EventArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEntityAggregatesOutput struct {
	_ struct{} `type:"structure"`

	// The number of entities that are affected by each of the specified events.
	EntityAggregates []EntityAggregate `locationName:"entityAggregates" type:"list"`
}

// String returns the string representation
func (s DescribeEntityAggregatesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEntityAggregates = "DescribeEntityAggregates"

// DescribeEntityAggregatesRequest returns a request value for making API operation for
// AWS Health APIs and Notifications.
//
// Returns the number of entities that are affected by each of the specified
// events. If no events are specified, the counts of all affected entities are
// returned.
//
//    // Example sending a request using DescribeEntityAggregatesRequest.
//    req := client.DescribeEntityAggregatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/health-2016-08-04/DescribeEntityAggregates
func (c *Client) DescribeEntityAggregatesRequest(input *DescribeEntityAggregatesInput) DescribeEntityAggregatesRequest {
	op := &aws.Operation{
		Name:       opDescribeEntityAggregates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEntityAggregatesInput{}
	}

	req := c.newRequest(op, input, &DescribeEntityAggregatesOutput{})

	return DescribeEntityAggregatesRequest{Request: req, Input: input, Copy: c.DescribeEntityAggregatesRequest}
}

// DescribeEntityAggregatesRequest is the request type for the
// DescribeEntityAggregates API operation.
type DescribeEntityAggregatesRequest struct {
	*aws.Request
	Input *DescribeEntityAggregatesInput
	Copy  func(*DescribeEntityAggregatesInput) DescribeEntityAggregatesRequest
}

// Send marshals and sends the DescribeEntityAggregates API request.
func (r DescribeEntityAggregatesRequest) Send(ctx context.Context) (*DescribeEntityAggregatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEntityAggregatesResponse{
		DescribeEntityAggregatesOutput: r.Request.Data.(*DescribeEntityAggregatesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEntityAggregatesResponse is the response type for the
// DescribeEntityAggregates API operation.
type DescribeEntityAggregatesResponse struct {
	*DescribeEntityAggregatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEntityAggregates request.
func (r *DescribeEntityAggregatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
