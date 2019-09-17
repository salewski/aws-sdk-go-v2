// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreatePipeline action.
type CreatePipelineInput struct {
	_ struct{} `type:"structure"`

	// Represents the structure of actions and stages to be performed in the pipeline.
	//
	// Pipeline is a required field
	Pipeline *PipelineDeclaration `locationName:"pipeline" type:"structure" required:"true"`

	// The tags for the pipeline.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreatePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePipelineInput"}

	if s.Pipeline == nil {
		invalidParams.Add(aws.NewErrParamRequired("Pipeline"))
	}
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			invalidParams.AddNested("Pipeline", err.(aws.ErrInvalidParams))
		}
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

// Represents the output of a CreatePipeline action.
type CreatePipelineOutput struct {
	_ struct{} `type:"structure"`

	// Represents the structure of actions and stages to be performed in the pipeline.
	Pipeline *PipelineDeclaration `locationName:"pipeline" type:"structure"`

	// Specifies the tags applied to the pipeline.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreatePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePipeline = "CreatePipeline"

// CreatePipelineRequest returns a request value for making API operation for
// AWS CodePipeline.
//
// Creates a pipeline.
//
// In the pipeline structure, you must include either artifactStore or artifactStores
// in your pipeline, but you cannot use both. If you create a cross-region action
// in your pipeline, you must use artifactStores.
//
//    // Example sending a request using CreatePipelineRequest.
//    req := client.CreatePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09/CreatePipeline
func (c *Client) CreatePipelineRequest(input *CreatePipelineInput) CreatePipelineRequest {
	op := &aws.Operation{
		Name:       opCreatePipeline,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePipelineInput{}
	}

	req := c.newRequest(op, input, &CreatePipelineOutput{})
	return CreatePipelineRequest{Request: req, Input: input, Copy: c.CreatePipelineRequest}
}

// CreatePipelineRequest is the request type for the
// CreatePipeline API operation.
type CreatePipelineRequest struct {
	*aws.Request
	Input *CreatePipelineInput
	Copy  func(*CreatePipelineInput) CreatePipelineRequest
}

// Send marshals and sends the CreatePipeline API request.
func (r CreatePipelineRequest) Send(ctx context.Context) (*CreatePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePipelineResponse{
		CreatePipelineOutput: r.Request.Data.(*CreatePipelineOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePipelineResponse is the response type for the
// CreatePipeline API operation.
type CreatePipelineResponse struct {
	*CreatePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePipeline request.
func (r *CreatePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
