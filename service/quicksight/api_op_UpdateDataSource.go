// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateDataSourceInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The credentials that QuickSight that uses to connect to your underlying source.
	// Currently, only credentials based on user name and password are supported.
	Credentials *DataSourceCredentials `type:"structure" sensitive:"true"`

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	//
	// DataSourceId is a required field
	DataSourceId *string `location:"uri" locationName:"DataSourceId" type:"string" required:"true"`

	// The parameters that QuickSight uses to connect to your underlying source.
	DataSourceParameters *DataSourceParameters `type:"structure"`

	// A display name for the data source.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Secure Socket Layer (SSL) properties that apply when QuickSight connects
	// to your underlying source.
	SslProperties *SslProperties `type:"structure"`

	// Use this parameter only when you want QuickSight to use a VPC connection
	// when connecting to your underlying source.
	VpcConnectionProperties *VpcConnectionProperties `type:"structure"`
}

// String returns the string representation
func (s UpdateDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDataSourceInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Credentials != nil {
		if err := s.Credentials.Validate(); err != nil {
			invalidParams.AddNested("Credentials", err.(aws.ErrInvalidParams))
		}
	}
	if s.DataSourceParameters != nil {
		if err := s.DataSourceParameters.Validate(); err != nil {
			invalidParams.AddNested("DataSourceParameters", err.(aws.ErrInvalidParams))
		}
	}
	if s.VpcConnectionProperties != nil {
		if err := s.VpcConnectionProperties.Validate(); err != nil {
			invalidParams.AddNested("VpcConnectionProperties", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDataSourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Credentials != nil {
		v := s.Credentials

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Credentials", v, metadata)
	}
	if s.DataSourceParameters != nil {
		v := s.DataSourceParameters

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DataSourceParameters", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SslProperties != nil {
		v := s.SslProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SslProperties", v, metadata)
	}
	if s.VpcConnectionProperties != nil {
		v := s.VpcConnectionProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VpcConnectionProperties", v, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSourceId != nil {
		v := *s.DataSourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateDataSourceOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the data source.
	Arn *string `type:"string"`

	// The ID of the data source. This ID is unique per AWS Region for each AWS
	// account.
	DataSourceId *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The update status of the data source's last update.
	UpdateStatus ResourceStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDataSourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSourceId != nil {
		v := *s.DataSourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DataSourceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.UpdateStatus) > 0 {
		v := s.UpdateStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UpdateStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateDataSource = "UpdateDataSource"

// UpdateDataSourceRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates a data source.
//
//    // Example sending a request using UpdateDataSourceRequest.
//    req := client.UpdateDataSourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateDataSource
func (c *Client) UpdateDataSourceRequest(input *UpdateDataSourceInput) UpdateDataSourceRequest {
	op := &aws.Operation{
		Name:       opUpdateDataSource,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sources/{DataSourceId}",
	}

	if input == nil {
		input = &UpdateDataSourceInput{}
	}

	req := c.newRequest(op, input, &UpdateDataSourceOutput{})

	return UpdateDataSourceRequest{Request: req, Input: input, Copy: c.UpdateDataSourceRequest}
}

// UpdateDataSourceRequest is the request type for the
// UpdateDataSource API operation.
type UpdateDataSourceRequest struct {
	*aws.Request
	Input *UpdateDataSourceInput
	Copy  func(*UpdateDataSourceInput) UpdateDataSourceRequest
}

// Send marshals and sends the UpdateDataSource API request.
func (r UpdateDataSourceRequest) Send(ctx context.Context) (*UpdateDataSourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDataSourceResponse{
		UpdateDataSourceOutput: r.Request.Data.(*UpdateDataSourceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDataSourceResponse is the response type for the
// UpdateDataSource API operation.
type UpdateDataSourceResponse struct {
	*UpdateDataSourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDataSource request.
func (r *UpdateDataSourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}