// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Specifies where to store data classification results, and the encryption
// settings to use when storing results in that location. Currently, you can
// store classification results only in an S3 bucket.
type PutClassificationExportConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Specifies where to store data classification results, and the encryption
	// settings to use when storing results in that location. Currently, you can
	// store classification results only in an S3 bucket.
	//
	// Configuration is a required field
	Configuration *ClassificationExportConfiguration `locationName:"configuration" type:"structure" required:"true"`
}

// String returns the string representation
func (s PutClassificationExportConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutClassificationExportConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutClassificationExportConfigurationInput"}

	if s.Configuration == nil {
		invalidParams.Add(aws.NewErrParamRequired("Configuration"))
	}
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			invalidParams.AddNested("Configuration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutClassificationExportConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Configuration != nil {
		v := s.Configuration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configuration", v, metadata)
	}
	return nil
}

// Provides information about updated settings for storing data classification
// results.
type PutClassificationExportConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Specifies where to store data classification results, and the encryption
	// settings to use when storing results in that location. Currently, you can
	// store classification results only in an S3 bucket.
	Configuration *ClassificationExportConfiguration `locationName:"configuration" type:"structure"`
}

// String returns the string representation
func (s PutClassificationExportConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutClassificationExportConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Configuration != nil {
		v := s.Configuration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "configuration", v, metadata)
	}
	return nil
}

const opPutClassificationExportConfiguration = "PutClassificationExportConfiguration"

// PutClassificationExportConfigurationRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Creates or updates the configuration settings for storing data classification
// results.
//
//    // Example sending a request using PutClassificationExportConfigurationRequest.
//    req := client.PutClassificationExportConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/PutClassificationExportConfiguration
func (c *Client) PutClassificationExportConfigurationRequest(input *PutClassificationExportConfigurationInput) PutClassificationExportConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutClassificationExportConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/classification-export-configuration",
	}

	if input == nil {
		input = &PutClassificationExportConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutClassificationExportConfigurationOutput{})

	return PutClassificationExportConfigurationRequest{Request: req, Input: input, Copy: c.PutClassificationExportConfigurationRequest}
}

// PutClassificationExportConfigurationRequest is the request type for the
// PutClassificationExportConfiguration API operation.
type PutClassificationExportConfigurationRequest struct {
	*aws.Request
	Input *PutClassificationExportConfigurationInput
	Copy  func(*PutClassificationExportConfigurationInput) PutClassificationExportConfigurationRequest
}

// Send marshals and sends the PutClassificationExportConfiguration API request.
func (r PutClassificationExportConfigurationRequest) Send(ctx context.Context) (*PutClassificationExportConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutClassificationExportConfigurationResponse{
		PutClassificationExportConfigurationOutput: r.Request.Data.(*PutClassificationExportConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutClassificationExportConfigurationResponse is the response type for the
// PutClassificationExportConfiguration API operation.
type PutClassificationExportConfigurationResponse struct {
	*PutClassificationExportConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutClassificationExportConfiguration request.
func (r *PutClassificationExportConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}