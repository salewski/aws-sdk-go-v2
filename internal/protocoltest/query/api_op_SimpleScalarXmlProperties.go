// Code generated by smithy-go-codegen DO NOT EDIT.
package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) SimpleScalarXmlProperties(ctx context.Context, params *SimpleScalarXmlPropertiesInput, optFns ...func(*Options)) (*SimpleScalarXmlPropertiesOutput, error) {
	stack := middleware.NewStack("SimpleScalarXmlProperties", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSimpleScalarXmlProperties(options.Region), middleware.Before)

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
			OperationName: "SimpleScalarXmlProperties",
			Err:           err,
		}
	}
	out := result.(*SimpleScalarXmlPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SimpleScalarXmlPropertiesInput struct {
}

type SimpleScalarXmlPropertiesOutput struct {
	StringValue       *string
	EmptyStringValue  *string
	TrueBooleanValue  *bool
	FalseBooleanValue *bool
	ByteValue         *int8
	ShortValue        *int16
	IntegerValue      *int32
	LongValue         *int64
	FloatValue        *float32
	DoubleValue       *float64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opSimpleScalarXmlProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "SimpleScalarXmlProperties",
	}
}