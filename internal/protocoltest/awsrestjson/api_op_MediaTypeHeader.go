// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example ensures that mediaType strings are base64 encoded in headers.
func (c *Client) MediaTypeHeader(ctx context.Context, params *MediaTypeHeaderInput, optFns ...func(*Options)) (*MediaTypeHeaderOutput, error) {
	stack := middleware.NewStack("MediaTypeHeader", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opMediaTypeHeader(options.Region), middleware.Before)
	addawsRestjson1_serdeOpMediaTypeHeaderMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "MediaTypeHeader",
			Err:           err,
		}
	}
	out := result.(*MediaTypeHeaderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MediaTypeHeaderInput struct {
	Json *string
}

type MediaTypeHeaderOutput struct {
	Json *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpMediaTypeHeaderMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpMediaTypeHeader{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpMediaTypeHeader{}, middleware.After)
}

func newServiceMetadataMiddleware_opMediaTypeHeader(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Json Protocol",
		ServiceID:      "restjsonprotocol",
		EndpointPrefix: "restjsonprotocol",
		OperationName:  "MediaTypeHeader",
	}
}