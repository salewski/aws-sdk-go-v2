// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This example uses all query string types.
func (c *Client) AllQueryStringTypes(ctx context.Context, params *AllQueryStringTypesInput, optFns ...func(*Options)) (*AllQueryStringTypesOutput, error) {
	stack := middleware.NewStack("AllQueryStringTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAllQueryStringTypes(options.Region), middleware.Before)

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
			OperationName: "AllQueryStringTypes",
			Err:           err,
		}
	}
	out := result.(*AllQueryStringTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllQueryStringTypesInput struct {
	QueryString        *string
	QueryStringList    []*string
	QueryStringSet     []*string
	QueryByte          *int8
	QueryShort         *int16
	QueryInteger       *int32
	QueryIntegerList   []*int32
	QueryIntegerSet    []*int32
	QueryLong          *int64
	QueryFloat         *float32
	QueryDouble        *float64
	QueryDoubleList    []*float64
	QueryBoolean       *bool
	QueryBooleanList   []*bool
	QueryTimestamp     *time.Time
	QueryTimestampList []*time.Time
	QueryEnum          types.FooEnum
	QueryEnumList      []types.FooEnum
}

type AllQueryStringTypesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opAllQueryStringTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "AllQueryStringTypes",
	}
}
