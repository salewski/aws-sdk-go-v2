// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateInstanceInput struct {
	_ struct{} `type:"structure"`

	// The default AWS OpsWorks Stacks agent version. You have the following options:
	//
	//    * INHERIT - Use the stack's default agent version setting.
	//
	//    * version_number - Use the specified agent version. This value overrides
	//    the stack's default setting. To update the agent version, edit the instance
	//    configuration and specify a new version. AWS OpsWorks Stacks then automatically
	//    installs that version on the instance.
	//
	// The default setting is INHERIT. To specify an agent version, you must use
	// the complete version number, not the abbreviated number shown on the console.
	// For a list of available agent version numbers, call DescribeAgentVersions.
	// AgentVersion cannot be set to Chef 12.2.
	AgentVersion *string `type:"string"`

	// A custom AMI ID to be used to create the instance. The AMI should be based
	// on one of the supported operating systems. For more information, see Using
	// Custom AMIs (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html).
	//
	// If you specify a custom AMI, you must set Os to Custom.
	AmiId *string `type:"string"`

	// The instance architecture. The default option is x86_64. Instance types do
	// not necessarily support both architectures. For a list of the architectures
	// that are supported by the different instance types, see Instance Families
	// and Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
	Architecture Architecture `type:"string" enum:"true"`

	// For load-based or time-based instances, the type. Windows stacks can use
	// only time-based instances.
	AutoScalingType AutoScalingType `type:"string" enum:"true"`

	// The instance Availability Zone. For more information, see Regions and Endpoints
	// (https://docs.aws.amazon.com/general/latest/gr/rande.html).
	AvailabilityZone *string `type:"string"`

	// An array of BlockDeviceMapping objects that specify the instance's block
	// devices. For more information, see Block Device Mapping (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html).
	// Note that block device mappings are not supported for custom AMIs.
	BlockDeviceMappings []BlockDeviceMapping `type:"list"`

	// Whether to create an Amazon EBS-optimized instance.
	EbsOptimized *bool `type:"boolean"`

	// The instance host name.
	Hostname *string `type:"string"`

	// Whether to install operating system and package updates when the instance
	// boots. The default value is true. To control when updates are installed,
	// set this value to false. You must then update your instances manually by
	// using CreateDeployment to run the update_dependencies stack command or by
	// manually running yum (Amazon Linux) or apt-get (Ubuntu) on the instances.
	//
	// We strongly recommend using the default value of true to ensure that your
	// instances have the latest security updates.
	InstallUpdatesOnBoot *bool `type:"boolean"`

	// The instance type, such as t2.micro. For a list of supported instance types,
	// open the stack in the console, choose Instances, and choose + Instance. The
	// Size list contains the currently supported types. For more information, see
	// Instance Families and Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html).
	// The parameter values that you use to specify the various types are in the
	// API Name column of the Available Instance Types table.
	//
	// InstanceType is a required field
	InstanceType *string `type:"string" required:"true"`

	// An array that contains the instance's layer IDs.
	//
	// LayerIds is a required field
	LayerIds []string `type:"list" required:"true"`

	// The instance's operating system, which must be set to one of the following.
	//
	//    * A supported Linux operating system: An Amazon Linux version, such as
	//    Amazon Linux 2018.03, Amazon Linux 2017.09, Amazon Linux 2017.03, Amazon
	//    Linux 2016.09, Amazon Linux 2016.03, Amazon Linux 2015.09, or Amazon Linux
	//    2015.03.
	//
	//    * A supported Ubuntu operating system, such as Ubuntu 16.04 LTS, Ubuntu
	//    14.04 LTS, or Ubuntu 12.04 LTS.
	//
	//    * CentOS Linux 7
	//
	//    * Red Hat Enterprise Linux 7
	//
	//    * A supported Windows operating system, such as Microsoft Windows Server
	//    2012 R2 Base, Microsoft Windows Server 2012 R2 with SQL Server Express,
	//    Microsoft Windows Server 2012 R2 with SQL Server Standard, or Microsoft
	//    Windows Server 2012 R2 with SQL Server Web.
	//
	//    * A custom AMI: Custom.
	//
	// For more information about the supported operating systems, see AWS OpsWorks
	// Stacks Operating Systems (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html).
	//
	// The default option is the current Amazon Linux version. If you set this parameter
	// to Custom, you must use the CreateInstance action's AmiId parameter to specify
	// the custom AMI that you want to use. Block device mappings are not supported
	// if the value is Custom. For more information about supported operating systems,
	// see Operating Systems (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-os.html)For
	// more information about how to use custom AMIs with AWS OpsWorks Stacks, see
	// Using Custom AMIs (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-custom-ami.html).
	Os *string `type:"string"`

	// The instance root device type. For more information, see Storage for the
	// Root Device (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ComponentsAMIs.html#storage-for-the-root-device).
	RootDeviceType RootDeviceType `type:"string" enum:"true"`

	// The instance's Amazon EC2 key-pair name.
	SshKeyName *string `type:"string"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`

	// The ID of the instance's subnet. If the stack is running in a VPC, you can
	// use this parameter to override the stack's default subnet ID value and direct
	// AWS OpsWorks Stacks to launch the instance in a different subnet.
	SubnetId *string `type:"string"`

	// The instance's tenancy option. The default option is no tenancy, or if the
	// instance is running in a VPC, inherit tenancy settings from the VPC. The
	// following are valid values for this parameter: dedicated, default, or host.
	// Because there are costs associated with changes in tenancy options, we recommend
	// that you research tenancy options before choosing them for your instances.
	// For more information about dedicated hosts, see Dedicated Hosts Overview
	// (http://aws.amazon.com/ec2/dedicated-hosts/) and Amazon EC2 Dedicated Hosts
	// (http://aws.amazon.com/ec2/dedicated-hosts/). For more information about
	// dedicated instances, see Dedicated Instances (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/dedicated-instance.html)
	// and Amazon EC2 Dedicated Instances (http://aws.amazon.com/ec2/purchasing-options/dedicated-instances/).
	Tenancy *string `type:"string"`

	// The instance's virtualization type, paravirtual or hvm.
	VirtualizationType *string `type:"string"`
}

// String returns the string representation
func (s CreateInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateInstanceInput"}

	if s.InstanceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceType"))
	}

	if s.LayerIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerIds"))
	}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a CreateInstance request.
type CreateInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The instance ID.
	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s CreateInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateInstance = "CreateInstance"

// CreateInstanceRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Creates an instance in a specified stack. For more information, see Adding
// an Instance to a Layer (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-add.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using CreateInstanceRequest.
//    req := client.CreateInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateInstance
func (c *Client) CreateInstanceRequest(input *CreateInstanceInput) CreateInstanceRequest {
	op := &aws.Operation{
		Name:       opCreateInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateInstanceInput{}
	}

	req := c.newRequest(op, input, &CreateInstanceOutput{})

	return CreateInstanceRequest{Request: req, Input: input, Copy: c.CreateInstanceRequest}
}

// CreateInstanceRequest is the request type for the
// CreateInstance API operation.
type CreateInstanceRequest struct {
	*aws.Request
	Input *CreateInstanceInput
	Copy  func(*CreateInstanceInput) CreateInstanceRequest
}

// Send marshals and sends the CreateInstance API request.
func (r CreateInstanceRequest) Send(ctx context.Context) (*CreateInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInstanceResponse{
		CreateInstanceOutput: r.Request.Data.(*CreateInstanceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInstanceResponse is the response type for the
// CreateInstance API operation.
type CreateInstanceResponse struct {
	*CreateInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInstance request.
func (r *CreateInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}