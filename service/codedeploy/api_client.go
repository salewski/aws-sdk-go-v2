// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
	"time"
)

const ServiceID = "CodeDeploy"

// AWS CodeDeploy AWS CodeDeploy is a deployment service that automates application
// deployments to Amazon EC2 instances, on-premises instances running in your own
// facility, serverless AWS Lambda functions, or applications in an Amazon ECS
// service. You can deploy a nearly unlimited variety of application content, such
// as an updated Lambda function, updated applications in an Amazon ECS service,
// code, web and configuration files, executables, packages, scripts, multimedia
// files, and so on. AWS CodeDeploy can deploy application content stored in Amazon
// S3 buckets, GitHub repositories, or Bitbucket repositories. You do not need to
// make changes to your existing code before you can use AWS CodeDeploy. AWS
// CodeDeploy makes it easier for you to rapidly release new features, helps you
// avoid downtime during application deployment, and handles the complexity of
// updating your applications, without many of the risks associated with
// error-prone manual deployments. AWS CodeDeploy Components Use the information in
// this guide to help you work with the following AWS CodeDeploy components:
//
//     *
// Application: A name that uniquely identifies the application you want to deploy.
// AWS CodeDeploy uses this name, which functions as a container, to ensure the
// correct combination of revision, deployment configuration, and deployment group
// are referenced during a deployment.
//
//     * Deployment group: A set of individual
// instances, CodeDeploy Lambda deployment configuration settings, or an Amazon ECS
// service and network details. A Lambda deployment group specifies how to route
// traffic to a new version of a Lambda function. An Amazon ECS deployment group
// specifies the service created in Amazon ECS to deploy, a load balancer, and a
// listener to reroute production traffic to an updated containerized application.
// An EC2/On-premises deployment group contains individually tagged instances,
// Amazon EC2 instances in Amazon EC2 Auto Scaling groups, or both. All deployment
// groups can specify optional trigger, alarm, and rollback settings.
//
//     *
// Deployment configuration: A set of deployment rules and deployment success and
// failure conditions used by AWS CodeDeploy during a deployment.
//
//     *
// Deployment: The process and the components used when updating a Lambda function,
// a containerized application in an Amazon ECS service, or of installing content
// on one or more instances.
//
//     * Application revisions: For an AWS Lambda
// deployment, this is an AppSpec file that specifies the Lambda function to be
// updated and one or more functions to validate deployment lifecycle events. For
// an Amazon ECS deployment, this is an AppSpec file that specifies the Amazon ECS
// task definition, container, and port where production traffic is rerouted. For
// an EC2/On-premises deployment, this is an archive file that contains source
// content—source code, webpages, executable files, and deployment scripts—along
// with an AppSpec file. Revisions are stored in Amazon S3 buckets or GitHub
// repositories. For Amazon S3, a revision is uniquely identified by its Amazon S3
// object key and its ETag, version, or both. For GitHub, a revision is uniquely
// identified by its commit ID.
//
// This guide also contains information to help you
// get details about the instances in your deployments, to make on-premises
// instances available for AWS CodeDeploy deployments, to get details about a
// Lambda function deployment, and to get details about Amazon ECS service
// deployments. AWS CodeDeploy Information Resources
//
//     * AWS CodeDeploy User
// Guide (https://docs.aws.amazon.com/codedeploy/latest/userguide)
//
//     * AWS
// CodeDeploy API Reference Guide
// (https://docs.aws.amazon.com/codedeploy/latest/APIReference/)
//
//     * AWS CLI
// Reference for AWS CodeDeploy
// (https://docs.aws.amazon.com/cli/latest/reference/deploy/index.html)
//
//     * AWS
// CodeDeploy Developer Forum
// (https://forums.aws.amazon.com/forum.jspa?forumID=179)
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []APIOptionFunc

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// An integer value representing the logging level.
	LogLevel aws.LogLevel

	// The logger writer interface to write logging messages to.
	Logger aws.Logger

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

func (o Options) GetCredentials() aws.CredentialsProvider {
	return o.Credentials
}

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetHTTPSignerV4() HTTPSignerV4 {
	return o.HTTPSignerV4
}

func (o Options) GetLogLevel() aws.LogLevel {
	return o.LogLevel
}

func (o Options) GetLogger() aws.Logger {
	return o.Logger
}

func (o Options) GetRegion() string {
	return o.Region
}

func (o Options) GetRetryer() retry.Retryer {
	return o.Retryer
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]APIOptionFunc, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

type APIOptionFunc func(*middleware.Stack) error

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		LogLevel:    cfg.LogLevel,
		Logger:      cfg.Logger,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
	}
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("codedeploy")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}
