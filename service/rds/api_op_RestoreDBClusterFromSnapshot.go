// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RestoreDBClusterFromSnapshotInput struct {
	_ struct{} `type:"structure"`

	// Provides the list of Availability Zones (AZs) where instances in the restored
	// DB cluster can be created.
	AvailabilityZones []string `locationNameList:"AvailabilityZone" type:"list"`

	// The target backtrack window, in seconds. To disable backtracking, set this
	// value to 0.
	//
	// Default: 0
	//
	// Constraints:
	//
	//    * If specified, this value must be set to a number from 0 to 259,200 (72
	//    hours).
	BacktrackWindow *int64 `type:"long"`

	// A value that indicates whether to copy all tags from the restored DB cluster
	// to snapshots of the restored DB cluster. The default is not to copy them.
	CopyTagsToSnapshot *bool `type:"boolean"`

	// The name of the DB cluster to create from the DB snapshot or DB cluster snapshot.
	// This parameter isn't case-sensitive.
	//
	// Constraints:
	//
	//    * Must contain from 1 to 63 letters, numbers, or hyphens
	//
	//    * First character must be a letter
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: my-snapshot-id
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`

	// The name of the DB cluster parameter group to associate with this DB cluster.
	// If this argument is omitted, the default DB cluster parameter group for the
	// specified engine is used.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing default DB cluster parameter
	//    group.
	//
	//    * Must be 1 to 255 letters, numbers, or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Can't end with a hyphen or contain two consecutive hyphens.
	DBClusterParameterGroupName *string `type:"string"`

	// The name of the DB subnet group to use for the new DB cluster.
	//
	// Constraints: If supplied, must match the name of an existing DB subnet group.
	//
	// Example: mySubnetgroup
	DBSubnetGroupName *string `type:"string"`

	// The database name for the restored DB cluster.
	DatabaseName *string `type:"string"`

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool `type:"boolean"`

	// The list of logs that the restored DB cluster is to export to Amazon CloudWatch
	// Logs. The values in the list depend on the DB engine being used. For more
	// information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)
	// in the Amazon Aurora User Guide.
	EnableCloudwatchLogsExports []string `type:"list"`

	// A value that indicates whether to enable mapping of AWS Identity and Access
	// Management (IAM) accounts to database accounts. By default, mapping is disabled.
	//
	// For more information, see IAM Database Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html)
	// in the Amazon Aurora User Guide.
	EnableIAMDatabaseAuthentication *bool `type:"boolean"`

	// The database engine to use for the new DB cluster.
	//
	// Default: The same as source
	//
	// Constraint: Must be compatible with the engine of the source
	//
	// Engine is a required field
	Engine *string `type:"string" required:"true"`

	// The DB engine mode of the DB cluster, either provisioned, serverless, parallelquery,
	// global, or multimaster.
	EngineMode *string `type:"string"`

	// The version of the database engine to use for the new DB cluster.
	//
	// To list all of the available engine versions for aurora (for MySQL 5.6-compatible
	// Aurora), use the following command:
	//
	// aws rds describe-db-engine-versions --engine aurora --query "DBEngineVersions[].EngineVersion"
	//
	// To list all of the available engine versions for aurora-mysql (for MySQL
	// 5.7-compatible Aurora), use the following command:
	//
	// aws rds describe-db-engine-versions --engine aurora-mysql --query "DBEngineVersions[].EngineVersion"
	//
	// To list all of the available engine versions for aurora-postgresql, use the
	// following command:
	//
	// aws rds describe-db-engine-versions --engine aurora-postgresql --query "DBEngineVersions[].EngineVersion"
	//
	// Aurora MySQL
	//
	// Example: 5.6.10a, 5.6.mysql_aurora.1.19.2, 5.7.12, 5.7.mysql_aurora.2.04.5
	//
	// Aurora PostgreSQL
	//
	// Example: 9.6.3, 10.7
	EngineVersion *string `type:"string"`

	// The AWS KMS key identifier to use when restoring an encrypted DB cluster
	// from a DB snapshot or DB cluster snapshot.
	//
	// The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryption
	// key. If you are restoring a DB cluster with the same AWS account that owns
	// the KMS encryption key used to encrypt the new DB cluster, then you can use
	// the KMS key alias instead of the ARN for the KMS encryption key.
	//
	// If you don't specify a value for the KmsKeyId parameter, then the following
	// occurs:
	//
	//    * If the DB snapshot or DB cluster snapshot in SnapshotIdentifier is encrypted,
	//    then the restored DB cluster is encrypted using the KMS key that was used
	//    to encrypt the DB snapshot or DB cluster snapshot.
	//
	//    * If the DB snapshot or DB cluster snapshot in SnapshotIdentifier is not
	//    encrypted, then the restored DB cluster is not encrypted.
	KmsKeyId *string `type:"string"`

	// The name of the option group to use for the restored DB cluster.
	OptionGroupName *string `type:"string"`

	// The port number on which the new DB cluster accepts connections.
	//
	// Constraints: This value must be 1150-65535
	//
	// Default: The same port as the original DB cluster.
	Port *int64 `type:"integer"`

	// For DB clusters in serverless DB engine mode, the scaling properties of the
	// DB cluster.
	ScalingConfiguration *ScalingConfiguration `type:"structure"`

	// The identifier for the DB snapshot or DB cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify
	// a DB cluster snapshot. However, you can use only the ARN to specify a DB
	// snapshot.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing Snapshot.
	//
	// SnapshotIdentifier is a required field
	SnapshotIdentifier *string `type:"string" required:"true"`

	// The tags to be assigned to the restored DB cluster.
	Tags []Tag `locationNameList:"Tag" type:"list"`

	// A list of VPC security groups that the new DB cluster will belong to.
	VpcSecurityGroupIds []string `locationNameList:"VpcSecurityGroupId" type:"list"`
}

// String returns the string representation
func (s RestoreDBClusterFromSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreDBClusterFromSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreDBClusterFromSnapshotInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if s.Engine == nil {
		invalidParams.Add(aws.NewErrParamRequired("Engine"))
	}

	if s.SnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreDBClusterFromSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Aurora DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters, StopDBCluster,
	// and StartDBCluster actions.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s RestoreDBClusterFromSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreDBClusterFromSnapshot = "RestoreDBClusterFromSnapshot"

// RestoreDBClusterFromSnapshotRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
//
// If a DB snapshot is specified, the target DB cluster is created from the
// source DB snapshot with a default configuration and default security group.
//
// If a DB cluster snapshot is specified, the target DB cluster is created from
// the source DB cluster restore point with the same configuration as the original
// source DB cluster, except that the new DB cluster is created with the default
// security group.
//
// For more information on Amazon Aurora, see What Is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide.
//
// This action only applies to Aurora DB clusters.
//
//    // Example sending a request using RestoreDBClusterFromSnapshotRequest.
//    req := client.RestoreDBClusterFromSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBClusterFromSnapshot
func (c *Client) RestoreDBClusterFromSnapshotRequest(input *RestoreDBClusterFromSnapshotInput) RestoreDBClusterFromSnapshotRequest {
	op := &aws.Operation{
		Name:       opRestoreDBClusterFromSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreDBClusterFromSnapshotInput{}
	}

	req := c.newRequest(op, input, &RestoreDBClusterFromSnapshotOutput{})
	return RestoreDBClusterFromSnapshotRequest{Request: req, Input: input, Copy: c.RestoreDBClusterFromSnapshotRequest}
}

// RestoreDBClusterFromSnapshotRequest is the request type for the
// RestoreDBClusterFromSnapshot API operation.
type RestoreDBClusterFromSnapshotRequest struct {
	*aws.Request
	Input *RestoreDBClusterFromSnapshotInput
	Copy  func(*RestoreDBClusterFromSnapshotInput) RestoreDBClusterFromSnapshotRequest
}

// Send marshals and sends the RestoreDBClusterFromSnapshot API request.
func (r RestoreDBClusterFromSnapshotRequest) Send(ctx context.Context) (*RestoreDBClusterFromSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBClusterFromSnapshotResponse{
		RestoreDBClusterFromSnapshotOutput: r.Request.Data.(*RestoreDBClusterFromSnapshotOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBClusterFromSnapshotResponse is the response type for the
// RestoreDBClusterFromSnapshot API operation.
type RestoreDBClusterFromSnapshotResponse struct {
	*RestoreDBClusterFromSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBClusterFromSnapshot request.
func (r *RestoreDBClusterFromSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
