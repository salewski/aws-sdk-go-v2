// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTestCasesInput struct {
	_ struct{} `type:"structure"`

	// A TestCaseFilter object used to filter the returned reports.
	Filter *TestCaseFilter `locationName:"filter" type:"structure"`

	// The maximum number of paginated test cases returned per response. Use nextToken
	// to iterate pages in the list of returned TestCase objects. The default value
	// is 100.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// During a previous call, the maximum number of items that can be returned
	// is the value specified in maxResults. If there more items in the list, then
	// a unique string called a nextToken is returned. To get the next batch of
	// items in the list, call this operation again, adding the next token to the
	// call. To get all of the items in the list, keep calling this operation with
	// each subsequent next token that is returned, until no more next tokens are
	// returned.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The ARN of the report for which test cases are returned.
	//
	// ReportArn is a required field
	ReportArn *string `locationName:"reportArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTestCasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTestCasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTestCasesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ReportArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReportArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTestCasesOutput struct {
	_ struct{} `type:"structure"`

	// During a previous call, the maximum number of items that can be returned
	// is the value specified in maxResults. If there more items in the list, then
	// a unique string called a nextToken is returned. To get the next batch of
	// items in the list, call this operation again, adding the next token to the
	// call. To get all of the items in the list, keep calling this operation with
	// each subsequent next token that is returned, until no more next tokens are
	// returned.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The returned list of test cases.
	TestCases []TestCase `locationName:"testCases" type:"list"`
}

// String returns the string representation
func (s DescribeTestCasesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTestCases = "DescribeTestCases"

// DescribeTestCasesRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Returns a list of details about test cases for a report.
//
//    // Example sending a request using DescribeTestCasesRequest.
//    req := client.DescribeTestCasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/DescribeTestCases
func (c *Client) DescribeTestCasesRequest(input *DescribeTestCasesInput) DescribeTestCasesRequest {
	op := &aws.Operation{
		Name:       opDescribeTestCases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeTestCasesInput{}
	}

	req := c.newRequest(op, input, &DescribeTestCasesOutput{})

	return DescribeTestCasesRequest{Request: req, Input: input, Copy: c.DescribeTestCasesRequest}
}

// DescribeTestCasesRequest is the request type for the
// DescribeTestCases API operation.
type DescribeTestCasesRequest struct {
	*aws.Request
	Input *DescribeTestCasesInput
	Copy  func(*DescribeTestCasesInput) DescribeTestCasesRequest
}

// Send marshals and sends the DescribeTestCases API request.
func (r DescribeTestCasesRequest) Send(ctx context.Context) (*DescribeTestCasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTestCasesResponse{
		DescribeTestCasesOutput: r.Request.Data.(*DescribeTestCasesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTestCasesRequestPaginator returns a paginator for DescribeTestCases.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTestCasesRequest(input)
//   p := codebuild.NewDescribeTestCasesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTestCasesPaginator(req DescribeTestCasesRequest) DescribeTestCasesPaginator {
	return DescribeTestCasesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeTestCasesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeTestCasesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTestCasesPaginator struct {
	aws.Pager
}

func (p *DescribeTestCasesPaginator) CurrentPage() *DescribeTestCasesOutput {
	return p.Pager.CurrentPage().(*DescribeTestCasesOutput)
}

// DescribeTestCasesResponse is the response type for the
// DescribeTestCases API operation.
type DescribeTestCasesResponse struct {
	*DescribeTestCasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTestCases request.
func (r *DescribeTestCasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}