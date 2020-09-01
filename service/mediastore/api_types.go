// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// This section describes operations that you can perform on an AWS Elemental
// MediaStore container.
type Container struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the container. The ARN has the following
	// format:
	//
	// arn:aws:<region>:<account that owns this container>:container/<name of container>
	//
	// For example: arn:aws:mediastore:us-west-2:111122223333:container/movies
	ARN *string `min:"1" type:"string"`

	// The state of access logging on the container. This value is false by default,
	// indicating that AWS Elemental MediaStore does not send access logs to Amazon
	// CloudWatch Logs. When you enable access logging on the container, MediaStore
	// changes this value to true, indicating that the service delivers access logs
	// for objects stored in that container to CloudWatch Logs.
	AccessLoggingEnabled *bool `type:"boolean"`

	// Unix timestamp.
	CreationTime *time.Time `type:"timestamp"`

	// The DNS endpoint of the container. Use the endpoint to identify the specific
	// container when sending requests to the data plane. The service assigns this
	// value when the container is created. Once the value has been assigned, it
	// does not change.
	Endpoint *string `min:"1" type:"string"`

	// The name of the container.
	Name *string `min:"1" type:"string"`

	// The status of container creation or deletion. The status is one of the following:
	// CREATING, ACTIVE, or DELETING. While the service is creating the container,
	// the status is CREATING. When the endpoint is available, the status changes
	// to ACTIVE.
	Status ContainerStatus `min:"1" type:"string" enum:"true"`
}

// String returns the string representation
func (s Container) String() string {
	return awsutil.Prettify(s)
}

// A rule for a CORS policy. You can add up to 100 rules to a CORS policy. If
// more than one rule applies, the service uses the first applicable rule listed.
type CorsRule struct {
	_ struct{} `type:"structure"`

	// Specifies which headers are allowed in a preflight OPTIONS request through
	// the Access-Control-Request-Headers header. Each header name that is specified
	// in Access-Control-Request-Headers must have a corresponding entry in the
	// rule. Only the headers that were requested are sent back.
	//
	// This element can contain only one wildcard character (*).
	//
	// AllowedHeaders is a required field
	AllowedHeaders []string `type:"list" required:"true"`

	// Identifies an HTTP method that the origin that is specified in the rule is
	// allowed to execute.
	//
	// Each CORS rule must contain at least one AllowedMethods and one AllowedOrigins
	// element.
	AllowedMethods []MethodName `min:"1" type:"list"`

	// One or more response headers that you want users to be able to access from
	// their applications (for example, from a JavaScript XMLHttpRequest object).
	//
	// Each CORS rule must have at least one AllowedOrigins element. The string
	// value can include only one wildcard character (*), for example, http://*.example.com.
	// Additionally, you can specify only one wildcard character to allow cross-origin
	// access for all origins.
	//
	// AllowedOrigins is a required field
	AllowedOrigins []string `min:"1" type:"list" required:"true"`

	// One or more headers in the response that you want users to be able to access
	// from their applications (for example, from a JavaScript XMLHttpRequest object).
	//
	// This element is optional for each rule.
	ExposeHeaders []string `type:"list"`

	// The time in seconds that your browser caches the preflight response for the
	// specified resource.
	//
	// A CORS rule can have only one MaxAgeSeconds element.
	MaxAgeSeconds *int64 `type:"integer"`
}

// String returns the string representation
func (s CorsRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CorsRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CorsRule"}

	if s.AllowedHeaders == nil {
		invalidParams.Add(aws.NewErrParamRequired("AllowedHeaders"))
	}
	if s.AllowedMethods != nil && len(s.AllowedMethods) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AllowedMethods", 1))
	}

	if s.AllowedOrigins == nil {
		invalidParams.Add(aws.NewErrParamRequired("AllowedOrigins"))
	}
	if s.AllowedOrigins != nil && len(s.AllowedOrigins) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AllowedOrigins", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The metric policy that is associated with the container. A metric policy
// allows AWS Elemental MediaStore to send metrics to Amazon CloudWatch. In
// the policy, you must indicate whether you want MediaStore to send container-level
// metrics. You can also include rules to define groups of objects that you
// want MediaStore to send object-level metrics for.
//
// To view examples of how to construct a metric policy for your use case, see
// Example Metric Policies (https://docs.aws.amazon.com/mediastore/latest/ug/policies-metric-examples.html).
type MetricPolicy struct {
	_ struct{} `type:"structure"`

	// A setting to enable or disable metrics at the container level.
	//
	// ContainerLevelMetrics is a required field
	ContainerLevelMetrics ContainerLevelMetrics `type:"string" required:"true" enum:"true"`

	// A parameter that holds an array of rules that enable metrics at the object
	// level. This parameter is optional, but if you choose to include it, you must
	// also include at least one rule. By default, you can include up to five rules.
	// You can also request a quota increase (https://console.aws.amazon.com/servicequotas/home?region=us-east-1#!/services/mediastore/quotas)
	// to allow up to 300 rules per policy.
	MetricPolicyRules []MetricPolicyRule `min:"1" type:"list"`
}

// String returns the string representation
func (s MetricPolicy) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricPolicy) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricPolicy"}
	if len(s.ContainerLevelMetrics) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ContainerLevelMetrics"))
	}
	if s.MetricPolicyRules != nil && len(s.MetricPolicyRules) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MetricPolicyRules", 1))
	}
	if s.MetricPolicyRules != nil {
		for i, v := range s.MetricPolicyRules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MetricPolicyRules", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A setting that enables metrics at the object level. Each rule contains an
// object group and an object group name. If the policy includes the MetricPolicyRules
// parameter, you must include at least one rule. Each metric policy can include
// up to five rules by default. You can also request a quota increase (https://console.aws.amazon.com/servicequotas/home?region=us-east-1#!/services/mediastore/quotas)
// to allow up to 300 rules per policy.
type MetricPolicyRule struct {
	_ struct{} `type:"structure"`

	// A path or file name that defines which objects to include in the group. Wildcards
	// (*) are acceptable.
	//
	// ObjectGroup is a required field
	ObjectGroup *string `min:"1" type:"string" required:"true"`

	// A name that allows you to refer to the object group.
	//
	// ObjectGroupName is a required field
	ObjectGroupName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s MetricPolicyRule) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MetricPolicyRule) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MetricPolicyRule"}

	if s.ObjectGroup == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectGroup"))
	}
	if s.ObjectGroup != nil && len(*s.ObjectGroup) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ObjectGroup", 1))
	}

	if s.ObjectGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectGroupName"))
	}
	if s.ObjectGroupName != nil && len(*s.ObjectGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ObjectGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A collection of tags associated with a container. Each tag consists of a
// key:value pair, which can be anything you define. Typically, the tag key
// represents a category (such as "environment") and the tag value represents
// a specific value within that category (such as "test," "development," or
// "production"). You can add up to 50 tags to each container. For more information
// about tagging, including naming and usage conventions, see Tagging Resources
// in MediaStore (https://docs.aws.amazon.com/mediastore/latest/ug/tagging.html).
type Tag struct {
	_ struct{} `type:"structure"`

	// Part of the key:value pair that defines a tag. You can use a tag key to describe
	// a category of information, such as "customer." Tag keys are case-sensitive.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// Part of the key:value pair that defines a tag. You can use a tag value to
	// describe a specific value within a category, such as "companyA" or "companyB."
	// Tag values are case-sensitive.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}