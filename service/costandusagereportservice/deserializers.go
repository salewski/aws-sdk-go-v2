// Code generated by smithy-go-codegen DO NOT EDIT.

package costandusagereportservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice/types"
	smithy "github.com/awslabs/smithy-go"
	smithyio "github.com/awslabs/smithy-go/io"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"strings"
)

type awsAwsjson11_deserializeOpDeleteReportDefinition struct {
}

func (*awsAwsjson11_deserializeOpDeleteReportDefinition) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpDeleteReportDefinition) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsAwsjson11_deserializeOpErrorDeleteReportDefinition(response)
	}
	output := &DeleteReportDefinitionOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson11_deserializeDocumentDeleteReportDefinitionOutput(&output, decoder)
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

func awsAwsjson11_deserializeOpErrorDeleteReportDefinition(response *smithyhttp.Response) error {
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
	case strings.EqualFold("InternalErrorException", errorCode):
		return awsAwsjson11_deserializeErrorInternalErrorException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsAwsjson11_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson11_deserializeOpDescribeReportDefinitions struct {
}

func (*awsAwsjson11_deserializeOpDescribeReportDefinitions) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpDescribeReportDefinitions) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsAwsjson11_deserializeOpErrorDescribeReportDefinitions(response)
	}
	output := &DescribeReportDefinitionsOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson11_deserializeDocumentDescribeReportDefinitionsOutput(&output, decoder)
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

func awsAwsjson11_deserializeOpErrorDescribeReportDefinitions(response *smithyhttp.Response) error {
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
	case strings.EqualFold("InternalErrorException", errorCode):
		return awsAwsjson11_deserializeErrorInternalErrorException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson11_deserializeOpModifyReportDefinition struct {
}

func (*awsAwsjson11_deserializeOpModifyReportDefinition) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpModifyReportDefinition) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsAwsjson11_deserializeOpErrorModifyReportDefinition(response)
	}
	output := &ModifyReportDefinitionOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson11_deserializeDocumentModifyReportDefinitionOutput(&output, decoder)
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

func awsAwsjson11_deserializeOpErrorModifyReportDefinition(response *smithyhttp.Response) error {
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
	case strings.EqualFold("InternalErrorException", errorCode):
		return awsAwsjson11_deserializeErrorInternalErrorException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsAwsjson11_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

type awsAwsjson11_deserializeOpPutReportDefinition struct {
}

func (*awsAwsjson11_deserializeOpPutReportDefinition) ID() string {
	return "OperationDeserializer"
}

func (m *awsAwsjson11_deserializeOpPutReportDefinition) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
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
		return out, metadata, awsAwsjson11_deserializeOpErrorPutReportDefinition(response)
	}
	output := &PutReportDefinitionOutput{}
	out.Result = output

	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(response.Body, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsAwsjson11_deserializeDocumentPutReportDefinitionOutput(&output, decoder)
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

func awsAwsjson11_deserializeOpErrorPutReportDefinition(response *smithyhttp.Response) error {
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
	case strings.EqualFold("DuplicateReportNameException", errorCode):
		return awsAwsjson11_deserializeErrorDuplicateReportNameException(response, errorBody)

	case strings.EqualFold("InternalErrorException", errorCode):
		return awsAwsjson11_deserializeErrorInternalErrorException(response, errorBody)

	case strings.EqualFold("ReportLimitReachedException", errorCode):
		return awsAwsjson11_deserializeErrorReportLimitReachedException(response, errorBody)

	case strings.EqualFold("ValidationException", errorCode):
		return awsAwsjson11_deserializeErrorValidationException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsAwsjson11_deserializeErrorDuplicateReportNameException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.DuplicateReportNameException{}
	err := awsAwsjson11_deserializeDocumentDuplicateReportNameException(&output, decoder)

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

func awsAwsjson11_deserializeErrorInternalErrorException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.InternalErrorException{}
	err := awsAwsjson11_deserializeDocumentInternalErrorException(&output, decoder)

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

func awsAwsjson11_deserializeErrorReportLimitReachedException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.ReportLimitReachedException{}
	err := awsAwsjson11_deserializeDocumentReportLimitReachedException(&output, decoder)

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

func awsAwsjson11_deserializeErrorValidationException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	var buff [1024]byte
	ringBuffer := smithyio.NewRingBuffer(buff[:])

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	output := &types.ValidationException{}
	err := awsAwsjson11_deserializeDocumentValidationException(&output, decoder)

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

func awsAwsjson11_deserializeDocumentAdditionalArtifactList(v *[]types.AdditionalArtifact, decoder *json.Decoder) error {
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

	var cv []types.AdditionalArtifact
	if *v == nil {
		cv = []types.AdditionalArtifact{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col types.AdditionalArtifact
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(string)
			if !ok {
				return fmt.Errorf("expected AdditionalArtifact to be of type string, got %T instead", val)
			}
			col = types.AdditionalArtifact(jtv)
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

func awsAwsjson11_deserializeDocumentDuplicateReportNameException(v **types.DuplicateReportNameException, decoder *json.Decoder) error {
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

	var sv *types.DuplicateReportNameException
	if *v == nil {
		sv = &types.DuplicateReportNameException{}
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
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", val)
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

func awsAwsjson11_deserializeDocumentInternalErrorException(v **types.InternalErrorException, decoder *json.Decoder) error {
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

	var sv *types.InternalErrorException
	if *v == nil {
		sv = &types.InternalErrorException{}
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
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", val)
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

func awsAwsjson11_deserializeDocumentReportDefinition(v **types.ReportDefinition, decoder *json.Decoder) error {
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

	var sv *types.ReportDefinition
	if *v == nil {
		sv = &types.ReportDefinition{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "AdditionalArtifacts":
			if err := awsAwsjson11_deserializeDocumentAdditionalArtifactList(&sv.AdditionalArtifacts, decoder); err != nil {
				return err
			}

		case "AdditionalSchemaElements":
			if err := awsAwsjson11_deserializeDocumentSchemaElementList(&sv.AdditionalSchemaElements, decoder); err != nil {
				return err
			}

		case "Compression":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected CompressionFormat to be of type string, got %T instead", val)
				}
				sv.Compression = types.CompressionFormat(jtv)
			}

		case "Format":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected ReportFormat to be of type string, got %T instead", val)
				}
				sv.Format = types.ReportFormat(jtv)
			}

		case "RefreshClosedReports":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(bool)
				if !ok {
					return fmt.Errorf("expected RefreshClosedReports to be of type *bool, got %T instead", val)
				}
				sv.RefreshClosedReports = &jtv
			}

		case "ReportName":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected ReportName to be of type string, got %T instead", val)
				}
				sv.ReportName = &jtv
			}

		case "ReportVersioning":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected ReportVersioning to be of type string, got %T instead", val)
				}
				sv.ReportVersioning = types.ReportVersioning(jtv)
			}

		case "S3Bucket":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected S3Bucket to be of type string, got %T instead", val)
				}
				sv.S3Bucket = &jtv
			}

		case "S3Prefix":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected S3Prefix to be of type string, got %T instead", val)
				}
				sv.S3Prefix = &jtv
			}

		case "S3Region":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected AWSRegion to be of type string, got %T instead", val)
				}
				sv.S3Region = types.AWSRegion(jtv)
			}

		case "TimeUnit":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected TimeUnit to be of type string, got %T instead", val)
				}
				sv.TimeUnit = types.TimeUnit(jtv)
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

func awsAwsjson11_deserializeDocumentReportDefinitionList(v *[]*types.ReportDefinition, decoder *json.Decoder) error {
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

	var cv []*types.ReportDefinition
	if *v == nil {
		cv = []*types.ReportDefinition{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col *types.ReportDefinition
		if err := awsAwsjson11_deserializeDocumentReportDefinition(&col, decoder); err != nil {
			return err
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

func awsAwsjson11_deserializeDocumentReportLimitReachedException(v **types.ReportLimitReachedException, decoder *json.Decoder) error {
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

	var sv *types.ReportLimitReachedException
	if *v == nil {
		sv = &types.ReportLimitReachedException{}
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
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", val)
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

func awsAwsjson11_deserializeDocumentSchemaElementList(v *[]types.SchemaElement, decoder *json.Decoder) error {
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

	var cv []types.SchemaElement
	if *v == nil {
		cv = []types.SchemaElement{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col types.SchemaElement
		val, err := decoder.Token()
		if err != nil {
			return err
		}
		if val != nil {
			jtv, ok := val.(string)
			if !ok {
				return fmt.Errorf("expected SchemaElement to be of type string, got %T instead", val)
			}
			col = types.SchemaElement(jtv)
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

func awsAwsjson11_deserializeDocumentValidationException(v **types.ValidationException, decoder *json.Decoder) error {
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

	var sv *types.ValidationException
	if *v == nil {
		sv = &types.ValidationException{}
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
					return fmt.Errorf("expected ErrorMessage to be of type string, got %T instead", val)
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

func awsAwsjson11_deserializeDocumentDeleteReportDefinitionOutput(v **DeleteReportDefinitionOutput, decoder *json.Decoder) error {
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

	var sv *DeleteReportDefinitionOutput
	if *v == nil {
		sv = &DeleteReportDefinitionOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "ResponseMessage":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected DeleteResponseMessage to be of type string, got %T instead", val)
				}
				sv.ResponseMessage = &jtv
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

func awsAwsjson11_deserializeDocumentDescribeReportDefinitionsOutput(v **DescribeReportDefinitionsOutput, decoder *json.Decoder) error {
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

	var sv *DescribeReportDefinitionsOutput
	if *v == nil {
		sv = &DescribeReportDefinitionsOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "NextToken":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected GenericString to be of type string, got %T instead", val)
				}
				sv.NextToken = &jtv
			}

		case "ReportDefinitions":
			if err := awsAwsjson11_deserializeDocumentReportDefinitionList(&sv.ReportDefinitions, decoder); err != nil {
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

func awsAwsjson11_deserializeDocumentModifyReportDefinitionOutput(v **ModifyReportDefinitionOutput, decoder *json.Decoder) error {
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

	var sv *ModifyReportDefinitionOutput
	if *v == nil {
		sv = &ModifyReportDefinitionOutput{}
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

func awsAwsjson11_deserializeDocumentPutReportDefinitionOutput(v **PutReportDefinitionOutput, decoder *json.Decoder) error {
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

	var sv *PutReportDefinitionOutput
	if *v == nil {
		sv = &PutReportDefinitionOutput{}
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
