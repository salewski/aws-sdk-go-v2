// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/types"
	smithy "github.com/awslabs/smithy-go"
	smithyio "github.com/awslabs/smithy-go/io"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"strings"
	"time"
)

type awsAwsjson10_deserializeOpEmptyInputAndEmptyOutput struct {
}

func (*awsAwsjson10_deserializeOpEmptyInputAndEmptyOutput) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpEmptyInputAndEmptyOutput) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorEmptyInputAndEmptyOutput(response)
	}
	output := &EmptyInputAndEmptyOutputOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson10_deserializeDocumentEmptyInputAndEmptyOutputOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorEmptyInputAndEmptyOutput(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson10_deserializeOpGreetingWithErrors struct {
}

func (*awsAwsjson10_deserializeOpGreetingWithErrors) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpGreetingWithErrors) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorGreetingWithErrors(response)
	}
	output := &GreetingWithErrorsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson10_deserializeDocumentGreetingWithErrorsOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorGreetingWithErrors(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("ComplexError", errorCode):
		return awsAwsjson10_deserializeErrorComplexError(response, errorBody)

	case strings.EqualFold("FooError", errorCode):
		return awsAwsjson10_deserializeErrorFooError(response, errorBody)

	case strings.EqualFold("InvalidGreeting", errorCode):
		return awsAwsjson10_deserializeErrorInvalidGreeting(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson10_deserializeOpJsonUnions struct {
}

func (*awsAwsjson10_deserializeOpJsonUnions) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpJsonUnions) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorJsonUnions(response)
	}
	output := &JsonUnionsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson10_deserializeDocumentJsonUnionsOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorJsonUnions(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson10_deserializeOpNoInputAndNoOutput struct {
}

func (*awsAwsjson10_deserializeOpNoInputAndNoOutput) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpNoInputAndNoOutput) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorNoInputAndNoOutput(response)
	}
	output := &NoInputAndNoOutputOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson10_deserializeDocumentNoInputAndNoOutputOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorNoInputAndNoOutput(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson10_deserializeOpNoInputAndOutput struct {
}

func (*awsAwsjson10_deserializeOpNoInputAndOutput) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson10_deserializeOpNoInputAndOutput) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsAwsjson10_deserializeOpErrorNoInputAndOutput(response)
	}
	output := &NoInputAndOutputOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson10_deserializeDocumentNoInputAndOutputOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsAwsjson10_deserializeOpErrorNoInputAndOutput(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsAwsjson10_deserializeErrorComplexError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.ComplexError{}
	err := awsAwsjson10_deserializeDocumentComplexError(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeErrorFooError(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.FooError{}
	err := awsAwsjson10_deserializeDocumentFooError(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeErrorInvalidGreeting(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.InvalidGreeting{}
	err := awsAwsjson10_deserializeDocumentInvalidGreeting(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	return output
}

func awsAwsjson10_deserializeDocumentComplexError(v **types.ComplexError, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.ComplexError
	if *v == nil {
		sv = &types.ComplexError{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Nested":
			if err := awsAwsjson10_deserializeDocumentComplexNestedErrorData(&sv.Nested, decoder); err != nil {
				return err
			}

		case "TopLevel":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.TopLevel = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentComplexNestedErrorData(v **types.ComplexNestedErrorData, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.ComplexNestedErrorData
	if *v == nil {
		sv = &types.ComplexNestedErrorData{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Fooooo":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.Foo = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentFooError(v **types.FooError, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.FooError
	if *v == nil {
		sv = &types.FooError{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentInvalidGreeting(v **types.InvalidGreeting, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.InvalidGreeting
	if *v == nil {
		sv = &types.InvalidGreeting{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "Message":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.Message = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentMyUnion(v *types.MyUnion, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var uv types.MyUnion
	t, err := decoder.Token()
	if err != nil {
		return err
	}
	switch t {
	case "blobValue":
		var mv []byte
		err := decoder.Decode(&mv)
		if err != nil {
			return err
		}
		uv = &types.MyUnionMemberBlobValue{Value: mv}

	case "booleanValue":
		var mv *bool
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(bool)
			if !ok {
				return fmt.Errorf("expected Boolean to be of type *bool, got %T instead", val)
			}
			mv = &jtv
		}
		uv = &types.MyUnionMemberBooleanValue{Value: *mv}

	case "enumValue":
		var mv types.FooEnum
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(string)
			if !ok {
				return fmt.Errorf("expected FooEnum to be of type string, got %T instead", val)
			}
			mv = types.FooEnum(jtv)
		}
		uv = &types.MyUnionMemberEnumValue{Value: mv}

	case "listValue":
		var mv []*string
		if err := awsAwsjson10_deserializeDocumentStringList(&mv, decoder); err != nil {
			return err
		}
		uv = &types.MyUnionMemberListValue{Value: mv}

	case "mapValue":
		var mv map[string]*string
		if err := awsAwsjson10_deserializeDocumentStringMap(&mv, decoder); err != nil {
			return err
		}
		uv = &types.MyUnionMemberMapValue{Value: mv}

	case "numberValue":
		var mv *int32
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(json.Number)
			if !ok {
				return fmt.Errorf("expected Integer to be json.Number, got %T instead", val)
			}
			i64, err := jtv.Int64()
			if err != nil {
				return err
			}
			mv = ptr.Int32(int32(i64))
		}
		uv = &types.MyUnionMemberNumberValue{Value: *mv}

	case "stringValue":
		var mv *string
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", val)
			}
			mv = &jtv
		}
		uv = &types.MyUnionMemberStringValue{Value: *mv}

	case "structureValue":
		var mv *types.GreetingStruct
		if err := awsAwsjson10_deserializeDocumentGreetingStruct(&mv, decoder); err != nil {
			return err
		}
		uv = &types.MyUnionMemberStructureValue{Value: mv}

	case "timestampValue":
		var mv *time.Time
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(json.Number)
			if !ok {
				return fmt.Errorf("expected Timestamp to be json.Number, got %T instead", val)
			}
			f64, err := jtv.Float64()
			if err != nil {
				return err
			}
			mv = ptr.Time(smithytime.ParseEpochSeconds(f64))
		}
		uv = &types.MyUnionMemberTimestampValue{Value: *mv}

	default:
		tagString, ok := t.(string)
		if !ok {
			return fmt.Errorf("expected string key, found %T", t)
		}
		value, err := smithyjson.CollectUnknownField(decoder)
		if err != nil {
			return err
		}
		uv = &types.UnknownUnionMember{Tag: tagString, Value: value}

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = uv
	return nil
}

func awsAwsjson10_deserializeDocumentGreetingStruct(v **types.GreetingStruct, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.GreetingStruct
	if *v == nil {
		sv = &types.GreetingStruct{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "hi":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.Hi = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentStringList(v *[]*string, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '[' {
		return fmt.Errorf("expect `[` as start token")
	}

	var cv []*string
	if *v == nil {
		cv = []*string{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col *string
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", val)
			}
			col = &jtv
		}
		cv = append(cv, col)

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != ']' {
		return fmt.Errorf("expect `]` as end token")
	}

	*v = cv
	return nil
}

func awsAwsjson10_deserializeDocumentStringMap(v *map[string]*string, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var mv map[string]*string
	if *v == nil {
		mv = map[string]*string{}
	} else {
		mv = *v
	}

	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			return err
		}

		key, ok := token.(string)
		if !ok {
			return fmt.Errorf("expected map-key of type string, found type %T", token)
		}

		var parsedVal *string
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(string)
			if !ok {
				return fmt.Errorf("expected String to be of type string, got %T instead", val)
			}
			parsedVal = &jtv
		}
		mv[key] = parsedVal

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = mv
	return nil
}

func awsAwsjson10_deserializeDocumentEmptyInputAndEmptyOutputOutput(v **EmptyInputAndEmptyOutputOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *EmptyInputAndEmptyOutputOutput
	if *v == nil {
		sv = &EmptyInputAndEmptyOutputOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentGreetingWithErrorsOutput(v **GreetingWithErrorsOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *GreetingWithErrorsOutput
	if *v == nil {
		sv = &GreetingWithErrorsOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "greeting":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected String to be of type string, got %T instead", val)
				}
				sv.Greeting = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentJsonUnionsOutput(v **JsonUnionsOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *JsonUnionsOutput
	if *v == nil {
		sv = &JsonUnionsOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "contents":
			if err := awsAwsjson10_deserializeDocumentMyUnion(&sv.Contents, decoder); err != nil {
				return err
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentNoInputAndNoOutputOutput(v **NoInputAndNoOutputOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *NoInputAndNoOutputOutput
	if *v == nil {
		sv = &NoInputAndNoOutputOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsAwsjson10_deserializeDocumentNoInputAndOutputOutput(v **NoInputAndOutputOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *NoInputAndOutputOutput
	if *v == nil {
		sv = &NoInputAndOutputOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}
