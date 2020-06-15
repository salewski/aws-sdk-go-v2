// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListBackupPlansInput struct {
	_ struct{} `type:"structure"`

	// A Boolean value with a default value of FALSE that returns deleted backup
	// plans when set to TRUE.
	IncludeDeleted *bool `location:"querystring" locationName:"includeDeleted" type:"boolean"`

	// The maximum number of items to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBackupPlansInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBackupPlansInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBackupPlansInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBackupPlansInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.IncludeDeleted != nil {
		v := *s.IncludeDeleted

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "includeDeleted", protocol.BoolValue(v), metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListBackupPlansOutput struct {
	_ struct{} `type:"structure"`

	// An array of backup plan list items containing metadata about your saved backup
	// plans.
	BackupPlansList []BackupPlansListMember `type:"list"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListBackupPlansOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListBackupPlansOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupPlansList != nil {
		v := s.BackupPlansList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "BackupPlansList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListBackupPlans = "ListBackupPlans"

// ListBackupPlansRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns metadata of your saved backup plans, including Amazon Resource Names
// (ARNs), plan IDs, creation and deletion dates, version IDs, plan names, and
// creator request IDs.
//
//    // Example sending a request using ListBackupPlansRequest.
//    req := client.ListBackupPlansRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListBackupPlans
func (c *Client) ListBackupPlansRequest(input *ListBackupPlansInput) ListBackupPlansRequest {
	op := &aws.Operation{
		Name:       opListBackupPlans,
		HTTPMethod: "GET",
		HTTPPath:   "/backup/plans/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListBackupPlansInput{}
	}

	req := c.newRequest(op, input, &ListBackupPlansOutput{})

	return ListBackupPlansRequest{Request: req, Input: input, Copy: c.ListBackupPlansRequest}
}

// ListBackupPlansRequest is the request type for the
// ListBackupPlans API operation.
type ListBackupPlansRequest struct {
	*aws.Request
	Input *ListBackupPlansInput
	Copy  func(*ListBackupPlansInput) ListBackupPlansRequest
}

// Send marshals and sends the ListBackupPlans API request.
func (r ListBackupPlansRequest) Send(ctx context.Context) (*ListBackupPlansResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBackupPlansResponse{
		ListBackupPlansOutput: r.Request.Data.(*ListBackupPlansOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListBackupPlansRequestPaginator returns a paginator for ListBackupPlans.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListBackupPlansRequest(input)
//   p := backup.NewListBackupPlansRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListBackupPlansPaginator(req ListBackupPlansRequest) ListBackupPlansPaginator {
	return ListBackupPlansPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListBackupPlansInput
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

// ListBackupPlansPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListBackupPlansPaginator struct {
	aws.Pager
}

func (p *ListBackupPlansPaginator) CurrentPage() *ListBackupPlansOutput {
	return p.Pager.CurrentPage().(*ListBackupPlansOutput)
}

// ListBackupPlansResponse is the response type for the
// ListBackupPlans API operation.
type ListBackupPlansResponse struct {
	*ListBackupPlansOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBackupPlans request.
func (r *ListBackupPlansResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}