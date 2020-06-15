// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteConfigurationAggregatorInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration aggregator.
	//
	// ConfigurationAggregatorName is a required field
	ConfigurationAggregatorName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationAggregatorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationAggregatorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationAggregatorInput"}

	if s.ConfigurationAggregatorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationAggregatorName"))
	}
	if s.ConfigurationAggregatorName != nil && len(*s.ConfigurationAggregatorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfigurationAggregatorName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteConfigurationAggregatorOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationAggregatorOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConfigurationAggregator = "DeleteConfigurationAggregator"

// DeleteConfigurationAggregatorRequest returns a request value for making API operation for
// AWS Config.
//
// Deletes the specified configuration aggregator and the aggregated data associated
// with the aggregator.
//
//    // Example sending a request using DeleteConfigurationAggregatorRequest.
//    req := client.DeleteConfigurationAggregatorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteConfigurationAggregator
func (c *Client) DeleteConfigurationAggregatorRequest(input *DeleteConfigurationAggregatorInput) DeleteConfigurationAggregatorRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationAggregator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConfigurationAggregatorInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigurationAggregatorOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteConfigurationAggregatorRequest{Request: req, Input: input, Copy: c.DeleteConfigurationAggregatorRequest}
}

// DeleteConfigurationAggregatorRequest is the request type for the
// DeleteConfigurationAggregator API operation.
type DeleteConfigurationAggregatorRequest struct {
	*aws.Request
	Input *DeleteConfigurationAggregatorInput
	Copy  func(*DeleteConfigurationAggregatorInput) DeleteConfigurationAggregatorRequest
}

// Send marshals and sends the DeleteConfigurationAggregator API request.
func (r DeleteConfigurationAggregatorRequest) Send(ctx context.Context) (*DeleteConfigurationAggregatorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationAggregatorResponse{
		DeleteConfigurationAggregatorOutput: r.Request.Data.(*DeleteConfigurationAggregatorOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationAggregatorResponse is the response type for the
// DeleteConfigurationAggregator API operation.
type DeleteConfigurationAggregatorResponse struct {
	*DeleteConfigurationAggregatorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationAggregator request.
func (r *DeleteConfigurationAggregatorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}