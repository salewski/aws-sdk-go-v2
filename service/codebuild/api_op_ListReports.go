// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListReportsInput struct {
	_ struct{} `type:"structure"`

	// A ReportFilter object used to filter the returned reports.
	Filter *ReportFilter `locationName:"filter" type:"structure"`

	// The maximum number of paginated reports returned per response. Use nextToken
	// to iterate pages in the list of returned Report objects. The default value
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

	// Specifies the sort order for the list of returned reports. Valid values are:
	//
	//    * ASCENDING: return reports in chronological order based on their creation
	//    date.
	//
	//    * DESCENDING: return reports in the reverse chronological order based
	//    on their creation date.
	SortOrder SortOrderType `locationName:"sortOrder" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListReportsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListReportsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListReportsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListReportsOutput struct {
	_ struct{} `type:"structure"`

	// During a previous call, the maximum number of items that can be returned
	// is the value specified in maxResults. If there more items in the list, then
	// a unique string called a nextToken is returned. To get the next batch of
	// items in the list, call this operation again, adding the next token to the
	// call. To get all of the items in the list, keep calling this operation with
	// each subsequent next token that is returned, until no more next tokens are
	// returned.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of returned ARNs for the reports in the current AWS account.
	Reports []string `locationName:"reports" min:"1" type:"list"`
}

// String returns the string representation
func (s ListReportsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListReports = "ListReports"

// ListReportsRequest returns a request value for making API operation for
// AWS CodeBuild.
//
// Returns a list of ARNs for the reports in the current AWS account.
//
//    // Example sending a request using ListReportsRequest.
//    req := client.ListReportsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06/ListReports
func (c *Client) ListReportsRequest(input *ListReportsInput) ListReportsRequest {
	op := &aws.Operation{
		Name:       opListReports,
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
		input = &ListReportsInput{}
	}

	req := c.newRequest(op, input, &ListReportsOutput{})

	return ListReportsRequest{Request: req, Input: input, Copy: c.ListReportsRequest}
}

// ListReportsRequest is the request type for the
// ListReports API operation.
type ListReportsRequest struct {
	*aws.Request
	Input *ListReportsInput
	Copy  func(*ListReportsInput) ListReportsRequest
}

// Send marshals and sends the ListReports API request.
func (r ListReportsRequest) Send(ctx context.Context) (*ListReportsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListReportsResponse{
		ListReportsOutput: r.Request.Data.(*ListReportsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListReportsRequestPaginator returns a paginator for ListReports.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListReportsRequest(input)
//   p := codebuild.NewListReportsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListReportsPaginator(req ListReportsRequest) ListReportsPaginator {
	return ListReportsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListReportsInput
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

// ListReportsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListReportsPaginator struct {
	aws.Pager
}

func (p *ListReportsPaginator) CurrentPage() *ListReportsOutput {
	return p.Pager.CurrentPage().(*ListReportsOutput)
}

// ListReportsResponse is the response type for the
// ListReports API operation.
type ListReportsResponse struct {
	*ListReportsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListReports request.
func (r *ListReportsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}