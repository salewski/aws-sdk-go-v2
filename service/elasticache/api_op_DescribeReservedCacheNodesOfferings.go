// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeReservedCacheNodesOfferings operation.
type DescribeReservedCacheNodesOfferingsInput struct {
	_ struct{} `type:"structure"`

	// The cache node type filter value. Use this parameter to show only the available
	// offerings matching the specified cache node type.
	//
	// The following node types are supported by ElastiCache. Generally speaking,
	// the current generation types provide more memory and computational power
	// at lower cost when compared to their equivalent previous generation counterparts.
	//
	//    * General purpose: Current generation: M5 node types: cache.m5.large,
	//    cache.m5.xlarge, cache.m5.2xlarge, cache.m5.4xlarge, cache.m5.12xlarge,
	//    cache.m5.24xlarge M4 node types: cache.m4.large, cache.m4.xlarge, cache.m4.2xlarge,
	//    cache.m4.4xlarge, cache.m4.10xlarge T3 node types: cache.t3.micro, cache.t3.small,
	//    cache.t3.medium T2 node types: cache.t2.micro, cache.t2.small, cache.t2.medium
	//    Previous generation: (not recommended) T1 node types: cache.t1.micro M1
	//    node types: cache.m1.small, cache.m1.medium, cache.m1.large, cache.m1.xlarge
	//    M3 node types: cache.m3.medium, cache.m3.large, cache.m3.xlarge, cache.m3.2xlarge
	//
	//    * Compute optimized: Previous generation: (not recommended) C1 node types:
	//    cache.c1.xlarge
	//
	//    * Memory optimized: Current generation: R5 node types: cache.r5.large,
	//    cache.r5.xlarge, cache.r5.2xlarge, cache.r5.4xlarge, cache.r5.12xlarge,
	//    cache.r5.24xlarge R4 node types: cache.r4.large, cache.r4.xlarge, cache.r4.2xlarge,
	//    cache.r4.4xlarge, cache.r4.8xlarge, cache.r4.16xlarge Previous generation:
	//    (not recommended) M2 node types: cache.m2.xlarge, cache.m2.2xlarge, cache.m2.4xlarge
	//    R3 node types: cache.r3.large, cache.r3.xlarge, cache.r3.2xlarge, cache.r3.4xlarge,
	//    cache.r3.8xlarge
	//
	// Additional node type info
	//
	//    * All current generation instance types are created in Amazon VPC by default.
	//
	//    * Redis append-only files (AOF) are not supported for T1 or T2 instances.
	//
	//    * Redis Multi-AZ with automatic failover is not supported on T1 instances.
	//
	//    * Redis configuration variables appendonly and appendfsync are not supported
	//    on Redis version 2.8.22 and later.
	CacheNodeType *string `type:"string"`

	// Duration filter value, specified in years or seconds. Use this parameter
	// to show only reservations for a given duration.
	//
	// Valid Values: 1 | 3 | 31536000 | 94608000
	Duration *string `type:"string"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The offering type filter value. Use this parameter to show only the available
	// offerings matching the specified offering type.
	//
	// Valid Values: "Light Utilization"|"Medium Utilization"|"Heavy Utilization"
	OfferingType *string `type:"string"`

	// The product description filter value. Use this parameter to show only the
	// available offerings matching the specified product description.
	ProductDescription *string `type:"string"`

	// The offering identifier filter value. Use this parameter to show only the
	// available offering that matches the specified reservation identifier.
	//
	// Example: 438012d3-4052-4cc7-b2e3-8d3372e0e706
	ReservedCacheNodesOfferingId *string `type:"string"`
}

// String returns the string representation
func (s DescribeReservedCacheNodesOfferingsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeReservedCacheNodesOfferings operation.
type DescribeReservedCacheNodesOfferingsOutput struct {
	_ struct{} `type:"structure"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`

	// A list of reserved cache node offerings. Each element in the list contains
	// detailed information about one offering.
	ReservedCacheNodesOfferings []ReservedCacheNodesOffering `locationNameList:"ReservedCacheNodesOffering" type:"list"`
}

// String returns the string representation
func (s DescribeReservedCacheNodesOfferingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeReservedCacheNodesOfferings = "DescribeReservedCacheNodesOfferings"

// DescribeReservedCacheNodesOfferingsRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Lists available reserved cache node offerings.
//
//    // Example sending a request using DescribeReservedCacheNodesOfferingsRequest.
//    req := client.DescribeReservedCacheNodesOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeReservedCacheNodesOfferings
func (c *Client) DescribeReservedCacheNodesOfferingsRequest(input *DescribeReservedCacheNodesOfferingsInput) DescribeReservedCacheNodesOfferingsRequest {
	op := &aws.Operation{
		Name:       opDescribeReservedCacheNodesOfferings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeReservedCacheNodesOfferingsInput{}
	}

	req := c.newRequest(op, input, &DescribeReservedCacheNodesOfferingsOutput{})

	return DescribeReservedCacheNodesOfferingsRequest{Request: req, Input: input, Copy: c.DescribeReservedCacheNodesOfferingsRequest}
}

// DescribeReservedCacheNodesOfferingsRequest is the request type for the
// DescribeReservedCacheNodesOfferings API operation.
type DescribeReservedCacheNodesOfferingsRequest struct {
	*aws.Request
	Input *DescribeReservedCacheNodesOfferingsInput
	Copy  func(*DescribeReservedCacheNodesOfferingsInput) DescribeReservedCacheNodesOfferingsRequest
}

// Send marshals and sends the DescribeReservedCacheNodesOfferings API request.
func (r DescribeReservedCacheNodesOfferingsRequest) Send(ctx context.Context) (*DescribeReservedCacheNodesOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReservedCacheNodesOfferingsResponse{
		DescribeReservedCacheNodesOfferingsOutput: r.Request.Data.(*DescribeReservedCacheNodesOfferingsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReservedCacheNodesOfferingsRequestPaginator returns a paginator for DescribeReservedCacheNodesOfferings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReservedCacheNodesOfferingsRequest(input)
//   p := elasticache.NewDescribeReservedCacheNodesOfferingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReservedCacheNodesOfferingsPaginator(req DescribeReservedCacheNodesOfferingsRequest) DescribeReservedCacheNodesOfferingsPaginator {
	return DescribeReservedCacheNodesOfferingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReservedCacheNodesOfferingsInput
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

// DescribeReservedCacheNodesOfferingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReservedCacheNodesOfferingsPaginator struct {
	aws.Pager
}

func (p *DescribeReservedCacheNodesOfferingsPaginator) CurrentPage() *DescribeReservedCacheNodesOfferingsOutput {
	return p.Pager.CurrentPage().(*DescribeReservedCacheNodesOfferingsOutput)
}

// DescribeReservedCacheNodesOfferingsResponse is the response type for the
// DescribeReservedCacheNodesOfferings API operation.
type DescribeReservedCacheNodesOfferingsResponse struct {
	*DescribeReservedCacheNodesOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReservedCacheNodesOfferings request.
func (r *DescribeReservedCacheNodesOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}