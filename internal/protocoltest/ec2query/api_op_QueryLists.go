// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This test serializes simple and complex lists.
func (c *Client) QueryLists(ctx context.Context, params *QueryListsInput, optFns ...func(*Options)) (*QueryListsOutput, error) {
	stack := middleware.NewStack("QueryLists", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opQueryLists(options.Region), middleware.Before)

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
			OperationName: "QueryLists",
			Err:           err,
		}
	}
	out := result.(*QueryListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryListsInput struct {
	ListArg                  []*string
	ComplexListArg           []*types.GreetingStruct
	ListArgWithXmlNameMember []*string
	ListArgWithXmlName       []*string
}

type QueryListsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opQueryLists(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "EC2 Protocol",
		ServiceID:      "ec2protocol",
		EndpointPrefix: "ec2protocol",
		OperationName:  "QueryLists",
	}
}
