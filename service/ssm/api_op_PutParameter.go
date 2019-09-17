// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutParameterInput struct {
	_ struct{} `type:"structure"`

	// A regular expression used to validate the parameter value. For example, for
	// String types with values restricted to numbers, you can specify the following:
	// AllowedPattern=^\d+$
	AllowedPattern *string `type:"string"`

	// Information about the parameter that you want to add to the system. Optional
	// but recommended.
	//
	// Do not enter personally identifiable information in this field.
	Description *string `type:"string"`

	// The KMS Key ID that you want to use to encrypt a parameter. Either the default
	// AWS Key Management Service (AWS KMS) key automatically assigned to your AWS
	// account or a custom key. Required for parameters that use the SecureString
	// data type.
	//
	// If you don't specify a key ID, the system uses the default key associated
	// with your AWS account.
	//
	//    * To use your default AWS KMS key, choose the SecureString data type,
	//    and do not specify the Key ID when you create the parameter. The system
	//    automatically populates Key ID with your default KMS key.
	//
	//    * To use a custom KMS key, choose the SecureString data type with the
	//    Key ID parameter.
	KeyId *string `min:"1" type:"string"`

	// The fully qualified name of the parameter that you want to add to the system.
	// The fully qualified name includes the complete hierarchy of the parameter
	// path and name. For example: /Dev/DBServer/MySQL/db-string13
	//
	// Naming Constraints:
	//
	//    * Parameter names are case sensitive.
	//
	//    * A parameter name must be unique within an AWS Region
	//
	//    * A parameter name can't be prefixed with "aws" or "ssm" (case-insensitive).
	//
	//    * Parameter names can include only the following symbols and letters:
	//    a-zA-Z0-9_.-/
	//
	//    * A parameter name can't include spaces.
	//
	//    * Parameter hierarchies are limited to a maximum depth of fifteen levels.
	//
	// For additional information about valid values for parameter names, see Requirements
	// and Constraints for Parameter Names (http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html)
	// in the AWS Systems Manager User Guide.
	//
	// The maximum length constraint listed below includes capacity for additional
	// system attributes that are not part of the name. The maximum length for the
	// fully qualified parameter name is 1011 characters.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Overwrite an existing parameter. If not specified, will default to "false".
	Overwrite *bool `type:"boolean"`

	// One or more policies to apply to a parameter. This action takes a JSON array.
	// Parameter Store supports the following policy types:
	//
	// Expiration: This policy deletes the parameter after it expires. When you
	// create the policy, you specify the expiration date. You can update the expiration
	// date and time by updating the policy. Updating the parameter does not affect
	// the expiration date and time. When the expiration time is reached, Parameter
	// Store deletes the parameter.
	//
	// ExpirationNotification: This policy triggers an event in Amazon CloudWatch
	// Events that notifies you about the expiration. By using this policy, you
	// can receive notification before or after the expiration time is reached,
	// in units of days or hours.
	//
	// NoChangeNotification: This policy triggers a CloudWatch event if a parameter
	// has not been modified for a specified period of time. This policy type is
	// useful when, for example, a secret needs to be changed within a period of
	// time, but it has not been changed.
	//
	// All existing policies are preserved until you send new policies or an empty
	// policy. For more information about parameter policies, see Working with Parameter
	// Policies (http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-su-policies.html).
	Policies *string `min:"1" type:"string"`

	// Optional metadata that you assign to a resource. Tags enable you to categorize
	// a resource in different ways, such as by purpose, owner, or environment.
	// For example, you might want to tag a Systems Manager parameter to identify
	// the type of resource to which it applies, the environment, or the type of
	// configuration data referenced by the parameter. In this case, you could specify
	// the following key name/value pairs:
	//
	//    * Key=Resource,Value=S3bucket
	//
	//    * Key=OS,Value=Windows
	//
	//    * Key=ParameterType,Value=LicenseKey
	//
	// To add tags to an existing Systems Manager parameter, use the AddTagsToResource
	// action.
	Tags []Tag `type:"list"`

	// The parameter tier to assign to a parameter.
	//
	// Parameter Store offers a standard tier and an advanced tier for parameters.
	// Standard parameters have a content size limit of 4 KB and can't be configured
	// to use parameter policies. You can create a maximum of 10,000 standard parameters
	// for each Region in an AWS account. Standard parameters are offered at no
	// additional cost.
	//
	// Advanced parameters have a content size limit of 8 KB and can be configured
	// to use parameter policies. You can create a maximum of 100,000 advanced parameters
	// for each Region in an AWS account. Advanced parameters incur a charge. For
	// more information, see About Advanced Parameters (http://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html)
	// in the AWS Systems Manager User Guide.
	//
	// You can change a standard parameter to an advanced parameter any time. But
	// you can't revert an advanced parameter to a standard parameter. Reverting
	// an advanced parameter to a standard parameter would result in data loss because
	// the system would truncate the size of the parameter from 8 KB to 4 KB. Reverting
	// would also remove any policies attached to the parameter. Lastly, advanced
	// parameters use a different form of encryption than standard parameters.
	//
	// If you no longer need an advanced parameter, or if you no longer want to
	// incur charges for an advanced parameter, you must delete it and recreate
	// it as a new standard parameter.
	//
	// Using the Default Tier Configuration
	//
	// In PutParameter requests, you can specify the tier to create the parameter
	// in. Whenever you specify a tier in the request, Parameter Store creates or
	// updates the parameter according to that request. However, if you do not specify
	// a tier in a request, Parameter Store assigns the tier based on the current
	// Parameter Store default tier configuration.
	//
	// The default tier when you begin using Parameter Store is the standard-parameter
	// tier. If you use the advanced-parameter tier, you can specify one of the
	// following as the default:
	//
	//    * Advanced: With this option, Parameter Store evaluates all requests as
	//    advanced parameters.
	//
	//    * Intelligent-Tiering: With this option, Parameter Store evaluates each
	//    request to determine if the parameter is standard or advanced. If the
	//    request doesn't include any options that require an advanced parameter,
	//    the parameter is created in the standard-parameter tier. If one or more
	//    options requiring an advanced parameter are included in the request, Parameter
	//    Store create a parameter in the advanced-parameter tier. This approach
	//    helps control your parameter-related costs by always creating standard
	//    parameters unless an advanced parameter is necessary.
	//
	// Options that require an advanced parameter include the following:
	//
	//    * The content size of the parameter is more than 4 KB.
	//
	//    * The parameter uses a parameter policy.
	//
	//    * More than 10,000 parameters already exist in your AWS account in the
	//    current Region.
	//
	// For more information about configuring the default tier option, see Specifying
	// a Default Parameter Tier (http://docs.aws.amazon.com/systems-manager/latest/userguide/ps-default-tier.html)
	// in the AWS Systems Manager User Guide.
	Tier ParameterTier `type:"string" enum:"true"`

	// The type of parameter that you want to add to the system.
	//
	// Items in a StringList must be separated by a comma (,). You can't use other
	// punctuation or special character to escape items in the list. If you have
	// a parameter value that requires a comma, then use the String data type.
	//
	// SecureString is not currently supported for AWS CloudFormation templates
	// or in the China Regions.
	//
	// Type is a required field
	Type ParameterType `type:"string" required:"true" enum:"true"`

	// The parameter value that you want to add to the system. Standard parameters
	// have a value limit of 4 KB. Advanced parameters have a value limit of 8 KB.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutParameterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutParameterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutParameterInput"}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Policies != nil && len(*s.Policies) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Policies", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutParameterOutput struct {
	_ struct{} `type:"structure"`

	// The new version number of a parameter. If you edit a parameter value, Parameter
	// Store automatically creates a new version and assigns this new version a
	// unique ID. You can reference a parameter version ID in API actions or in
	// Systems Manager documents (SSM documents). By default, if you don't specify
	// a specific version, the system returns the latest parameter value when a
	// parameter is called.
	Version *int64 `type:"long"`
}

// String returns the string representation
func (s PutParameterOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutParameter = "PutParameter"

// PutParameterRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Add a parameter to the system.
//
//    // Example sending a request using PutParameterRequest.
//    req := client.PutParameterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/PutParameter
func (c *Client) PutParameterRequest(input *PutParameterInput) PutParameterRequest {
	op := &aws.Operation{
		Name:       opPutParameter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutParameterInput{}
	}

	req := c.newRequest(op, input, &PutParameterOutput{})
	return PutParameterRequest{Request: req, Input: input, Copy: c.PutParameterRequest}
}

// PutParameterRequest is the request type for the
// PutParameter API operation.
type PutParameterRequest struct {
	*aws.Request
	Input *PutParameterInput
	Copy  func(*PutParameterInput) PutParameterRequest
}

// Send marshals and sends the PutParameter API request.
func (r PutParameterRequest) Send(ctx context.Context) (*PutParameterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutParameterResponse{
		PutParameterOutput: r.Request.Data.(*PutParameterOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutParameterResponse is the response type for the
// PutParameter API operation.
type PutParameterResponse struct {
	*PutParameterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutParameter request.
func (r *PutParameterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
