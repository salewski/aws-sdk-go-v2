// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodbstreams

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeStream operation.
type DescribeStreamInput struct {
	_ struct{} `type:"structure"`

	// The shard ID of the first item that this operation will evaluate. Use the
	// value that was returned for LastEvaluatedShardId in the previous operation.
	ExclusiveStartShardId *string `min:"28" type:"string"`

	// The maximum number of shard objects to return. The upper limit is 100.
	Limit *int64 `min:"1" type:"integer"`

	// The Amazon Resource Name (ARN) for the stream.
	//
	// StreamArn is a required field
	StreamArn *string `min:"37" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStreamInput"}
	if s.ExclusiveStartShardId != nil && len(*s.ExclusiveStartShardId) < 28 {
		invalidParams.Add(aws.NewErrParamMinLen("ExclusiveStartShardId", 28))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.StreamArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamArn"))
	}
	if s.StreamArn != nil && len(*s.StreamArn) < 37 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamArn", 37))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DescribeStream operation.
type DescribeStreamOutput struct {
	_ struct{} `type:"structure"`

	// A complete description of the stream, including its creation date and time,
	// the DynamoDB table associated with the stream, the shard IDs within the stream,
	// and the beginning and ending sequence numbers of stream records within the
	// shards.
	StreamDescription *StreamDescription `type:"structure"`
}

// String returns the string representation
func (s DescribeStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStream = "DescribeStream"

// DescribeStreamRequest returns a request value for making API operation for
// Amazon DynamoDB Streams.
//
// Returns information about a stream, including the current status of the stream,
// its Amazon Resource Name (ARN), the composition of its shards, and its corresponding
// DynamoDB table.
//
// You can call DescribeStream at a maximum rate of 10 times per second.
//
// Each shard in the stream has a SequenceNumberRange associated with it. If
// the SequenceNumberRange has a StartingSequenceNumber but no EndingSequenceNumber,
// then the shard is still open (able to receive more stream records). If both
// StartingSequenceNumber and EndingSequenceNumber are present, then that shard
// is closed and can no longer receive more data.
//
//    // Example sending a request using DescribeStreamRequest.
//    req := client.DescribeStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/streams-dynamodb-2012-08-10/DescribeStream
func (c *Client) DescribeStreamRequest(input *DescribeStreamInput) DescribeStreamRequest {
	op := &aws.Operation{
		Name:       opDescribeStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStreamInput{}
	}

	req := c.newRequest(op, input, &DescribeStreamOutput{})

	return DescribeStreamRequest{Request: req, Input: input, Copy: c.DescribeStreamRequest}
}

// DescribeStreamRequest is the request type for the
// DescribeStream API operation.
type DescribeStreamRequest struct {
	*aws.Request
	Input *DescribeStreamInput
	Copy  func(*DescribeStreamInput) DescribeStreamRequest
}

// Send marshals and sends the DescribeStream API request.
func (r DescribeStreamRequest) Send(ctx context.Context) (*DescribeStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStreamResponse{
		DescribeStreamOutput: r.Request.Data.(*DescribeStreamOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStreamResponse is the response type for the
// DescribeStream API operation.
type DescribeStreamResponse struct {
	*DescribeStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStream request.
func (r *DescribeStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
