// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the ID of the gateway to delete.
type DeleteGatewayInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGatewayInput"}

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

// A JSON object containing the ID of the deleted gateway.
type DeleteGatewayOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DeleteGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteGateway = "DeleteGateway"

// DeleteGatewayRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Deletes a gateway. To specify which gateway to delete, use the Amazon Resource
// Name (ARN) of the gateway in your request. The operation deletes the gateway;
// however, it does not delete the gateway virtual machine (VM) from your host
// computer.
//
// After you delete a gateway, you cannot reactivate it. Completed snapshots
// of the gateway volumes are not deleted upon deleting the gateway, however,
// pending snapshots will not complete. After you delete a gateway, your next
// step is to remove it from your environment.
//
// You no longer pay software charges after the gateway is deleted; however,
// your existing Amazon EBS snapshots persist and you will continue to be billed
// for these snapshots. You can choose to remove all remaining Amazon EBS snapshots
// by canceling your Amazon EC2 subscription. If you prefer not to cancel your
// Amazon EC2 subscription, you can delete your snapshots using the Amazon EC2
// console. For more information, see the AWS Storage Gateway Detail Page (http://aws.amazon.com/storagegateway).
//
//    // Example sending a request using DeleteGatewayRequest.
//    req := client.DeleteGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DeleteGateway
func (c *Client) DeleteGatewayRequest(input *DeleteGatewayInput) DeleteGatewayRequest {
	op := &aws.Operation{
		Name:       opDeleteGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteGatewayInput{}
	}

	req := c.newRequest(op, input, &DeleteGatewayOutput{})
	return DeleteGatewayRequest{Request: req, Input: input, Copy: c.DeleteGatewayRequest}
}

// DeleteGatewayRequest is the request type for the
// DeleteGateway API operation.
type DeleteGatewayRequest struct {
	*aws.Request
	Input *DeleteGatewayInput
	Copy  func(*DeleteGatewayInput) DeleteGatewayRequest
}

// Send marshals and sends the DeleteGateway API request.
func (r DeleteGatewayRequest) Send(ctx context.Context) (*DeleteGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGatewayResponse{
		DeleteGatewayOutput: r.Request.Data.(*DeleteGatewayOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGatewayResponse is the response type for the
// DeleteGateway API operation.
type DeleteGatewayResponse struct {
	*DeleteGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGateway request.
func (r *DeleteGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
