// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a get repository triggers operation.
type GetRepositoryTriggersInput struct {
	_ struct{} `type:"structure"`

	// The name of the repository for which the trigger is configured.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRepositoryTriggersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRepositoryTriggersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRepositoryTriggersInput"}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a get repository triggers operation.
type GetRepositoryTriggersOutput struct {
	_ struct{} `type:"structure"`

	// The system-generated unique ID for the trigger.
	ConfigurationId *string `locationName:"configurationId" type:"string"`

	// The JSON block of configuration information for each trigger.
	Triggers []RepositoryTrigger `locationName:"triggers" type:"list"`
}

// String returns the string representation
func (s GetRepositoryTriggersOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRepositoryTriggers = "GetRepositoryTriggers"

// GetRepositoryTriggersRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Gets information about triggers configured for a repository.
//
//    // Example sending a request using GetRepositoryTriggersRequest.
//    req := client.GetRepositoryTriggersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetRepositoryTriggers
func (c *Client) GetRepositoryTriggersRequest(input *GetRepositoryTriggersInput) GetRepositoryTriggersRequest {
	op := &aws.Operation{
		Name:       opGetRepositoryTriggers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRepositoryTriggersInput{}
	}

	req := c.newRequest(op, input, &GetRepositoryTriggersOutput{})

	return GetRepositoryTriggersRequest{Request: req, Input: input, Copy: c.GetRepositoryTriggersRequest}
}

// GetRepositoryTriggersRequest is the request type for the
// GetRepositoryTriggers API operation.
type GetRepositoryTriggersRequest struct {
	*aws.Request
	Input *GetRepositoryTriggersInput
	Copy  func(*GetRepositoryTriggersInput) GetRepositoryTriggersRequest
}

// Send marshals and sends the GetRepositoryTriggers API request.
func (r GetRepositoryTriggersRequest) Send(ctx context.Context) (*GetRepositoryTriggersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRepositoryTriggersResponse{
		GetRepositoryTriggersOutput: r.Request.Data.(*GetRepositoryTriggersOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRepositoryTriggersResponse is the response type for the
// GetRepositoryTriggers API operation.
type GetRepositoryTriggersResponse struct {
	*GetRepositoryTriggersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRepositoryTriggers request.
func (r *GetRepositoryTriggersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}