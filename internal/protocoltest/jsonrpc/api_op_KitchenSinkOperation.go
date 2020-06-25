// Code generated by smithy-go-codegen DO NOT EDIT.
package jsonrpc

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

func (c *Client) KitchenSinkOperation(ctx context.Context, params *KitchenSinkOperationInput, optFns ...func(*Options)) (*KitchenSinkOperationOutput, error) {
	stack := middleware.NewStack("KitchenSinkOperation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddlewares(stack, options)
	stack.Initialize.Add(newServiceMetadataMiddleware_opKitchenSinkOperation(options.Region), middleware.Before)

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
			OperationName: "KitchenSinkOperation",
			Err:           err,
		}
	}
	out := result.(*KitchenSinkOperationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type KitchenSinkOperationInput struct {
	Blob                   []byte
	Boolean                *bool
	Double                 *float64
	EmptyStruct            *types.EmptyStruct
	Float                  *float32
	HttpdateTimestamp      *time.Time
	Integer                *int32
	Iso8601Timestamp       *time.Time
	JsonValue              *string
	ListOfLists            [][]*string
	ListOfMapsOfStrings    []map[string]*string
	ListOfStrings          []*string
	ListOfStructs          []*types.SimpleStruct
	Long                   *int64
	MapOfListsOfStrings    map[string][]*string
	MapOfMaps              map[string]map[string]*string
	MapOfStrings           map[string]*string
	MapOfStructs           map[string]*types.SimpleStruct
	RecursiveList          []*types.KitchenSink
	RecursiveMap           map[string]*types.KitchenSink
	RecursiveStruct        *types.KitchenSink
	SimpleStruct           *types.SimpleStruct
	String_                *string
	StructWithLocationName *types.StructWithLocationName
	Timestamp              *time.Time
	UnixTimestamp          *time.Time
}

type KitchenSinkOperationOutput struct {
	Blob                   []byte
	Boolean                *bool
	Double                 *float64
	EmptyStruct            *types.EmptyStruct
	Float                  *float32
	HttpdateTimestamp      *time.Time
	Integer                *int32
	Iso8601Timestamp       *time.Time
	JsonValue              *string
	ListOfLists            [][]*string
	ListOfMapsOfStrings    []map[string]*string
	ListOfStrings          []*string
	ListOfStructs          []*types.SimpleStruct
	Long                   *int64
	MapOfListsOfStrings    map[string][]*string
	MapOfMaps              map[string]map[string]*string
	MapOfStrings           map[string]*string
	MapOfStructs           map[string]*types.SimpleStruct
	RecursiveList          []*types.KitchenSink
	RecursiveMap           map[string]*types.KitchenSink
	RecursiveStruct        *types.KitchenSink
	SimpleStruct           *types.SimpleStruct
	String_                *string
	StructWithLocationName *types.StructWithLocationName
	Timestamp              *time.Time
	UnixTimestamp          *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opKitchenSinkOperation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Json Protocol",
		ServiceID:      "jsonprotocol",
		EndpointPrefix: "jsonprotocol",
		SigningName:    "foo",
		OperationName:  "KitchenSinkOperation",
	}
}
