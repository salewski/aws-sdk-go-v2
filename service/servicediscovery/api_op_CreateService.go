// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateServiceInput struct {
	_ struct{} `type:"structure"`

	// A unique string that identifies the request and that allows failed CreateService
	// requests to be retried without the risk of executing the operation twice.
	// CreatorRequestId can be any unique string, for example, a date/time stamp.
	CreatorRequestId *string `type:"string" idempotencyToken:"true"`

	// A description for the service.
	Description *string `type:"string"`

	// A complex type that contains information about the Amazon Route 53 records
	// that you want AWS Cloud Map to create when you register an instance.
	DnsConfig *DnsConfig `type:"structure"`

	// Public DNS and HTTP namespaces only. A complex type that contains settings
	// for an optional Route 53 health check. If you specify settings for a health
	// check, AWS Cloud Map associates the health check with all the Route 53 DNS
	// records that you specify in DnsConfig.
	//
	// If you specify a health check configuration, you can specify either HealthCheckCustomConfig
	// or HealthCheckConfig but not both.
	//
	// For information about the charges for health checks, see AWS Cloud Map Pricing
	// (http://aws.amazon.com/cloud-map/pricing/).
	HealthCheckConfig *HealthCheckConfig `type:"structure"`

	// A complex type that contains information about an optional custom health
	// check.
	//
	// If you specify a health check configuration, you can specify either HealthCheckCustomConfig
	// or HealthCheckConfig but not both.
	//
	// You can't add, update, or delete a HealthCheckCustomConfig configuration
	// from an existing service.
	HealthCheckCustomConfig *HealthCheckCustomConfig `type:"structure"`

	// The name that you want to assign to the service.
	//
	// If you want AWS Cloud Map to create an SRV record when you register an instance,
	// and if you're using a system that requires a specific SRV format, such as
	// HAProxy (http://www.haproxy.org/), specify the following for Name:
	//
	//    * Start the name with an underscore (_), such as _exampleservice
	//
	//    * End the name with ._protocol, such as ._tcp
	//
	// When you register an instance, AWS Cloud Map creates an SRV record and assigns
	// a name to the record by concatenating the service name and the namespace
	// name, for example:
	//
	// _exampleservice._tcp.example.com
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The ID of the namespace that you want to use to create the service.
	NamespaceId *string `type:"string"`

	// The tags to add to the service. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length
	// of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateServiceInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.DnsConfig != nil {
		if err := s.DnsConfig.Validate(); err != nil {
			invalidParams.AddNested("DnsConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.HealthCheckConfig != nil {
		if err := s.HealthCheckConfig.Validate(); err != nil {
			invalidParams.AddNested("HealthCheckConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.HealthCheckCustomConfig != nil {
		if err := s.HealthCheckCustomConfig.Validate(); err != nil {
			invalidParams.AddNested("HealthCheckCustomConfig", err.(aws.ErrInvalidParams))
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

type CreateServiceOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about the new service.
	Service *Service `type:"structure"`
}

// String returns the string representation
func (s CreateServiceOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateService = "CreateService"

// CreateServiceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Creates a service, which defines the configuration for the following entities:
//
//    * For public and private DNS namespaces, one of the following combinations
//    of DNS records in Amazon Route 53: A AAAA A and AAAA SRV CNAME
//
//    * Optionally, a health check
//
// After you create the service, you can submit a RegisterInstance (https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html)
// request, and AWS Cloud Map uses the values in the configuration to create
// the specified entities.
//
// For the current limit on the number of instances that you can register using
// the same namespace and using the same service, see AWS Cloud Map Limits (https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html)
// in the AWS Cloud Map Developer Guide.
//
//    // Example sending a request using CreateServiceRequest.
//    req := client.CreateServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/CreateService
func (c *Client) CreateServiceRequest(input *CreateServiceInput) CreateServiceRequest {
	op := &aws.Operation{
		Name:       opCreateService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServiceInput{}
	}

	req := c.newRequest(op, input, &CreateServiceOutput{})

	return CreateServiceRequest{Request: req, Input: input, Copy: c.CreateServiceRequest}
}

// CreateServiceRequest is the request type for the
// CreateService API operation.
type CreateServiceRequest struct {
	*aws.Request
	Input *CreateServiceInput
	Copy  func(*CreateServiceInput) CreateServiceRequest
}

// Send marshals and sends the CreateService API request.
func (r CreateServiceRequest) Send(ctx context.Context) (*CreateServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateServiceResponse{
		CreateServiceOutput: r.Request.Data.(*CreateServiceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateServiceResponse is the response type for the
// CreateService API operation.
type CreateServiceResponse struct {
	*CreateServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateService request.
func (r *CreateServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}