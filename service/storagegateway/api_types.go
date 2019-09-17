// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Describes an iSCSI cached volume.
type CachediSCSIVolume struct {
	_ struct{} `type:"structure"`

	// The date the volume was created. Volumes created prior to March 28, 2017
	// don’t have this time stamp.
	CreatedDate *time.Time `type:"timestamp"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server
	// side encryption. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// If the cached volume was created from a snapshot, this field contains the
	// snapshot ID used, e.g. snap-78e22663. Otherwise, this field is not included.
	SourceSnapshotId *string `type:"string"`

	// The name of the iSCSI target used by an initiator to connect to a volume
	// and used as a suffix for the target ARN. For example, specifying TargetName
	// as myvolume results in the target ARN of arn:aws:storagegateway:us-east-2:111122223333:gateway/sgw-12A3456B/target/iqn.1997-05.com.amazon:myvolume.
	// The target name must be unique across all volumes on a gateway.
	//
	// If you don't specify a value, Storage Gateway uses the value that was previously
	// used for this volume as the new target name.
	TargetName *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the storage volume.
	VolumeARN *string `min:"50" type:"string"`

	// A value that indicates whether a storage volume is attached to or detached
	// from a gateway. For more information, see Moving Your Volumes to a Different
	// Gateway (https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-volumes.html#attach-detach-volume).
	VolumeAttachmentStatus *string `min:"3" type:"string"`

	// The unique identifier of the volume, e.g. vol-AE4B946D.
	VolumeId *string `min:"12" type:"string"`

	// Represents the percentage complete if the volume is restoring or bootstrapping
	// that represents the percent of data transferred. This field does not appear
	// in the response if the cached volume is not restoring or bootstrapping.
	VolumeProgress *float64 `type:"double"`

	// The size, in bytes, of the volume capacity.
	VolumeSizeInBytes *int64 `type:"long"`

	// One of the VolumeStatus values that indicates the state of the storage volume.
	VolumeStatus *string `min:"3" type:"string"`

	// One of the VolumeType enumeration values that describes the type of the volume.
	VolumeType *string `min:"3" type:"string"`

	// The size of the data stored on the volume in bytes. This value is calculated
	// based on the number of blocks that are touched, instead of the actual amount
	// of data written. This value can be useful for sequential write patterns but
	// less accurate for random write patterns. VolumeUsedInBytes is different from
	// the compressed size of the volume, which is the value that is used to calculate
	// your bill.
	//
	// This value is not available for volumes created prior to May 13, 2015, until
	// you store data on the volume.
	VolumeUsedInBytes *int64 `type:"long"`

	// An VolumeiSCSIAttributes object that represents a collection of iSCSI attributes
	// for one stored volume.
	VolumeiSCSIAttributes *VolumeiSCSIAttributes `type:"structure"`
}

// String returns the string representation
func (s CachediSCSIVolume) String() string {
	return awsutil.Prettify(s)
}

// Describes Challenge-Handshake Authentication Protocol (CHAP) information
// that supports authentication between your gateway and iSCSI initiators.
type ChapInfo struct {
	_ struct{} `type:"structure"`

	// The iSCSI initiator that connects to the target.
	InitiatorName *string `min:"1" type:"string"`

	// The secret key that the initiator (for example, the Windows client) must
	// provide to participate in mutual CHAP with the target.
	SecretToAuthenticateInitiator *string `min:"1" type:"string"`

	// The secret key that the target must provide to participate in mutual CHAP
	// with the initiator (e.g. Windows client).
	SecretToAuthenticateTarget *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the volume.
	//
	// Valid Values: 50 to 500 lowercase letters, numbers, periods (.), and hyphens
	// (-).
	TargetARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s ChapInfo) String() string {
	return awsutil.Prettify(s)
}

// Lists iSCSI information about a VTL device.
type DeviceiSCSIAttributes struct {
	_ struct{} `type:"structure"`

	// Indicates whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled *bool `type:"boolean"`

	// The network interface identifier of the VTL device.
	NetworkInterfaceId *string `type:"string"`

	// The port used to communicate with iSCSI VTL device targets.
	NetworkInterfacePort *int64 `type:"integer"`

	// Specifies the unique Amazon Resource Name (ARN) that encodes the iSCSI qualified
	// name(iqn) of a tape drive or media changer target.
	TargetARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DeviceiSCSIAttributes) String() string {
	return awsutil.Prettify(s)
}

// Represents a gateway's local disk.
type Disk struct {
	_ struct{} `type:"structure"`

	// The iSCSI qualified name (IQN) that is defined for a disk. This field is
	// not included in the response if the local disk is not defined as an iSCSI
	// target. The format of this field is targetIqn::LUNNumber::region-volumeId.
	DiskAllocationResource *string `type:"string"`

	// One of the DiskAllocationType enumeration values that identifies how a local
	// disk is used. Valid values: UPLOAD_BUFFER, CACHE_STORAGE
	DiskAllocationType *string `min:"3" type:"string"`

	// A list of values that represents attributes of a local disk.
	DiskAttributeList []string `type:"list"`

	// The unique device ID or other distinguishing data that identifies a local
	// disk.
	DiskId *string `min:"1" type:"string"`

	// The device node of a local disk as assigned by the virtualization environment.
	DiskNode *string `type:"string"`

	// The path of a local disk in the gateway virtual machine (VM).
	DiskPath *string `type:"string"`

	// The local disk size in bytes.
	DiskSizeInBytes *int64 `type:"long"`

	// A value that represents the status of a local disk.
	DiskStatus *string `type:"string"`
}

// String returns the string representation
func (s Disk) String() string {
	return awsutil.Prettify(s)
}

// Describes a file share.
type FileShareInfo struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string `min:"50" type:"string"`

	// The ID of the file share.
	FileShareId *string `min:"12" type:"string"`

	// The status of the file share. Possible values are CREATING, UPDATING, AVAILABLE
	// and DELETING.
	FileShareStatus *string `min:"3" type:"string"`

	// The type of the file share.
	FileShareType FileShareType `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s FileShareInfo) String() string {
	return awsutil.Prettify(s)
}

// Describes a gateway object.
type GatewayInfo struct {
	_ struct{} `type:"structure"`

	// The ID of the Amazon EC2 instance that was used to launch the gateway.
	Ec2InstanceId *string `type:"string"`

	// The AWS Region where the Amazon EC2 instance is located.
	Ec2InstanceRegion *string `type:"string"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`

	// The unique identifier assigned to your gateway during activation. This ID
	// becomes part of the gateway Amazon Resource Name (ARN), which you use as
	// input for other operations.
	GatewayId *string `min:"12" type:"string"`

	// The name of the gateway.
	GatewayName *string `type:"string"`

	// The state of the gateway.
	//
	// Valid Values: DISABLED or ACTIVE
	GatewayOperationalState *string `min:"2" type:"string"`

	// The type of the gateway.
	GatewayType *string `min:"2" type:"string"`
}

// String returns the string representation
func (s GatewayInfo) String() string {
	return awsutil.Prettify(s)
}

// Describes Network File System (NFS) file share default values. Files and
// folders stored as Amazon S3 objects in S3 buckets don't, by default, have
// Unix file permissions assigned to them. Upon discovery in an S3 bucket by
// Storage Gateway, the S3 objects that represent files and folders are assigned
// these default Unix permissions. This operation is only supported for file
// gateways.
type NFSFileShareDefaults struct {
	_ struct{} `type:"structure"`

	// The Unix directory mode in the form "nnnn". For example, "0666" represents
	// the default access mode for all directories inside the file share. The default
	// value is 0777.
	DirectoryMode *string `min:"1" type:"string"`

	// The Unix file mode in the form "nnnn". For example, "0666" represents the
	// default file mode inside the file share. The default value is 0666.
	FileMode *string `min:"1" type:"string"`

	// The default group ID for the file share (unless the files have another group
	// ID specified). The default value is nfsnobody.
	GroupId *int64 `type:"long"`

	// The default owner ID for files in the file share (unless the files have another
	// owner ID specified). The default value is nfsnobody.
	OwnerId *int64 `type:"long"`
}

// String returns the string representation
func (s NFSFileShareDefaults) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NFSFileShareDefaults) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NFSFileShareDefaults"}
	if s.DirectoryMode != nil && len(*s.DirectoryMode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DirectoryMode", 1))
	}
	if s.FileMode != nil && len(*s.FileMode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FileMode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The Unix file permissions and ownership information assigned, by default,
// to native S3 objects when file gateway discovers them in S3 buckets. This
// operation is only supported in file gateways.
type NFSFileShareInfo struct {
	_ struct{} `type:"structure"`

	// The list of clients that are allowed to access the file gateway. The list
	// must contain either valid IP addresses or valid CIDR blocks.
	ClientList []string `min:"1" type:"list"`

	// The default storage class for objects put into an Amazon S3 bucket by the
	// file gateway. Possible values are S3_STANDARD, S3_STANDARD_IA, or S3_ONEZONE_IA.
	// If this field is not populated, the default value S3_STANDARD is used. Optional.
	DefaultStorageClass *string `min:"5" type:"string"`

	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string `min:"50" type:"string"`

	// The ID of the file share.
	FileShareId *string `min:"12" type:"string"`

	// The status of the file share. Possible values are CREATING, UPDATING, AVAILABLE
	// and DELETING.
	FileShareStatus *string `min:"3" type:"string"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`

	// A value that enables guessing of the MIME type for uploaded objects based
	// on file extensions. Set this value to true to enable MIME type guessing,
	// and otherwise to false. The default value is true.
	GuessMIMETypeEnabled *bool `type:"boolean"`

	// True to use Amazon S3 server side encryption with your own AWS KMS key, or
	// false to use a key managed by Amazon S3. Optional.
	KMSEncrypted *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server
	// side encryption. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// The ARN of the backend storage used for storing file data.
	LocationARN *string `min:"16" type:"string"`

	// Describes Network File System (NFS) file share default values. Files and
	// folders stored as Amazon S3 objects in S3 buckets don't, by default, have
	// Unix file permissions assigned to them. Upon discovery in an S3 bucket by
	// Storage Gateway, the S3 objects that represent files and folders are assigned
	// these default Unix permissions. This operation is only supported for file
	// gateways.
	NFSFileShareDefaults *NFSFileShareDefaults `type:"structure"`

	// A value that sets the access control list permission for objects in the S3
	// bucket that a file gateway puts objects into. The default value is "private".
	ObjectACL ObjectACL `type:"string" enum:"true"`

	// The file share path used by the NFS client to identify the mount point.
	Path *string `type:"string"`

	// A value that sets the write status of a file share. This value is true if
	// the write status is read-only, and otherwise false.
	ReadOnly *bool `type:"boolean"`

	// A value that sets who pays the cost of the request and the cost associated
	// with data download from the S3 bucket. If this value is set to true, the
	// requester pays the costs. Otherwise the S3 bucket owner pays. However, the
	// S3 bucket owner always pays the cost of storing data.
	//
	// RequesterPays is a configuration for the S3 bucket that backs the file share,
	// so make sure that the configuration on the file share is the same as the
	// S3 bucket configuration.
	RequesterPays *bool `type:"boolean"`

	// The ARN of the IAM role that file gateway assumes when it accesses the underlying
	// storage.
	Role *string `min:"20" type:"string"`

	// The user mapped to anonymous user. Valid options are the following:
	//
	//    * RootSquash - Only root is mapped to anonymous user.
	//
	//    * NoSquash - No one is mapped to anonymous user
	//
	//    * AllSquash - Everyone is mapped to anonymous user.
	Squash *string `min:"5" type:"string"`

	// A list of up to 50 tags assigned to the NFS file share, sorted alphabetically
	// by key name. Each tag is a key-value pair. For a gateway with more than 10
	// tags assigned, you can view all tags using the ListTagsForResource API operation.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s NFSFileShareInfo) String() string {
	return awsutil.Prettify(s)
}

// Describes a gateway's network interface.
type NetworkInterface struct {
	_ struct{} `type:"structure"`

	// The Internet Protocol version 4 (IPv4) address of the interface.
	Ipv4Address *string `type:"string"`

	// The Internet Protocol version 6 (IPv6) address of the interface. Currently
	// not supported.
	Ipv6Address *string `type:"string"`

	// The Media Access Control (MAC) address of the interface.
	//
	// This is currently unsupported and will not be returned in output.
	MacAddress *string `type:"string"`
}

// String returns the string representation
func (s NetworkInterface) String() string {
	return awsutil.Prettify(s)
}

// The Windows file permissions and ownership information assigned, by default,
// to native S3 objects when file gateway discovers them in S3 buckets. This
// operation is only supported for file gateways.
type SMBFileShareInfo struct {
	_ struct{} `type:"structure"`

	// A list of users or groups in the Active Directory that have administrator
	// rights to the file share. A group must be prefixed with the @ character.
	// For example @group1. Can only be set if Authentication is set to ActiveDirectory.
	AdminUserList []string `type:"list"`

	// The authentication method of the file share.
	//
	// Valid values are ActiveDirectory or GuestAccess. The default is ActiveDirectory.
	Authentication *string `min:"5" type:"string"`

	// The default storage class for objects put into an Amazon S3 bucket by the
	// file gateway. Possible values are S3_STANDARD, S3_STANDARD_IA, or S3_ONEZONE_IA.
	// If this field is not populated, the default value S3_STANDARD is used. Optional.
	DefaultStorageClass *string `min:"5" type:"string"`

	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string `min:"50" type:"string"`

	// The ID of the file share.
	FileShareId *string `min:"12" type:"string"`

	// The status of the file share. Possible values are CREATING, UPDATING, AVAILABLE
	// and DELETING.
	FileShareStatus *string `min:"3" type:"string"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`

	// A value that enables guessing of the MIME type for uploaded objects based
	// on file extensions. Set this value to true to enable MIME type guessing,
	// and otherwise to false. The default value is true.
	GuessMIMETypeEnabled *bool `type:"boolean"`

	// A list of users or groups in the Active Directory that are not allowed to
	// access the file share. A group must be prefixed with the @ character. For
	// example @group1. Can only be set if Authentication is set to ActiveDirectory.
	InvalidUserList []string `type:"list"`

	// True to use Amazon S3 server-side encryption with your own AWS KMS key, or
	// false to use a key managed by Amazon S3. Optional.
	KMSEncrypted *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server
	// side encryption. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// The ARN of the backend storage used for storing file data.
	LocationARN *string `min:"16" type:"string"`

	// A value that sets the access control list permission for objects in the S3
	// bucket that a file gateway puts objects into. The default value is "private".
	ObjectACL ObjectACL `type:"string" enum:"true"`

	// The file share path used by the SMB client to identify the mount point.
	Path *string `type:"string"`

	// A value that sets the write status of a file share. This value is true if
	// the write status is read-only, and otherwise false.
	ReadOnly *bool `type:"boolean"`

	// A value that sets who pays the cost of the request and the cost associated
	// with data download from the S3 bucket. If this value is set to true, the
	// requester pays the costs. Otherwise the S3 bucket owner pays. However, the
	// S3 bucket owner always pays the cost of storing data.
	//
	// RequesterPays is a configuration for the S3 bucket that backs the file share,
	// so make sure that the configuration on the file share is the same as the
	// S3 bucket configuration.
	RequesterPays *bool `type:"boolean"`

	// The ARN of the IAM role that file gateway assumes when it accesses the underlying
	// storage.
	Role *string `min:"20" type:"string"`

	// If this value is set to "true", indicates that ACL (access control list)
	// is enabled on the SMB file share. If it is set to "false", it indicates that
	// file and directory permissions are mapped to the POSIX permission.
	//
	// For more information, see https://docs.aws.amazon.com/storagegateway/latest/userguide/smb-acl.html
	// in the Storage Gateway User Guide.
	SMBACLEnabled *bool `type:"boolean"`

	// A list of up to 50 tags assigned to the SMB file share, sorted alphabetically
	// by key name. Each tag is a key-value pair. For a gateway with more than 10
	// tags assigned, you can view all tags using the ListTagsForResource API operation.
	Tags []Tag `type:"list"`

	// A list of users or groups in the Active Directory that are allowed to access
	// the file share. A group must be prefixed with the @ character. For example
	// @group1. Can only be set if Authentication is set to ActiveDirectory.
	ValidUserList []string `type:"list"`
}

// String returns the string representation
func (s SMBFileShareInfo) String() string {
	return awsutil.Prettify(s)
}

// Provides additional information about an error that was returned by the service
// as an or. See the errorCode and errorDetails members for more information
// about the error.
type StorageGatewayError struct {
	_ struct{} `type:"structure"`

	// Additional information about the error.
	ErrorCode ErrorCode `locationName:"errorCode" type:"string" enum:"true"`

	// Human-readable text that provides detail about the error that occurred.
	ErrorDetails map[string]string `locationName:"errorDetails" type:"map"`
}

// String returns the string representation
func (s StorageGatewayError) String() string {
	return awsutil.Prettify(s)
}

// Describes an iSCSI stored volume.
type StorediSCSIVolume struct {
	_ struct{} `type:"structure"`

	// The date the volume was created. Volumes created prior to March 28, 2017
	// don’t have this time stamp.
	CreatedDate *time.Time `type:"timestamp"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server
	// side encryption. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// Indicates if when the stored volume was created, existing data on the underlying
	// local disk was preserved.
	//
	// Valid Values: true, false
	PreservedExistingData *bool `type:"boolean"`

	// If the stored volume was created from a snapshot, this field contains the
	// snapshot ID used, e.g. snap-78e22663. Otherwise, this field is not included.
	SourceSnapshotId *string `type:"string"`

	// The name of the iSCSI target used by an initiator to connect to a volume
	// and used as a suffix for the target ARN. For example, specifying TargetName
	// as myvolume results in the target ARN of arn:aws:storagegateway:us-east-2:111122223333:gateway/sgw-12A3456B/target/iqn.1997-05.com.amazon:myvolume.
	// The target name must be unique across all volumes on a gateway.
	//
	// If you don't specify a value, Storage Gateway uses the value that was previously
	// used for this volume as the new target name.
	TargetName *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the storage volume.
	VolumeARN *string `min:"50" type:"string"`

	// A value that indicates whether a storage volume is attached to, detached
	// from, or is in the process of detaching from a gateway. For more information,
	// see Moving Your Volumes to a Different Gateway (https://docs.aws.amazon.com/storagegateway/latest/userguide/managing-volumes.html#attach-detach-volume).
	VolumeAttachmentStatus *string `min:"3" type:"string"`

	// The ID of the local disk that was specified in the CreateStorediSCSIVolume
	// operation.
	VolumeDiskId *string `min:"1" type:"string"`

	// The unique identifier of the volume, e.g. vol-AE4B946D.
	VolumeId *string `min:"12" type:"string"`

	// Represents the percentage complete if the volume is restoring or bootstrapping
	// that represents the percent of data transferred. This field does not appear
	// in the response if the stored volume is not restoring or bootstrapping.
	VolumeProgress *float64 `type:"double"`

	// The size of the volume in bytes.
	VolumeSizeInBytes *int64 `type:"long"`

	// One of the VolumeStatus values that indicates the state of the storage volume.
	VolumeStatus *string `min:"3" type:"string"`

	// One of the VolumeType enumeration values describing the type of the volume.
	VolumeType *string `min:"3" type:"string"`

	// The size of the data stored on the volume in bytes. This value is calculated
	// based on the number of blocks that are touched, instead of the actual amount
	// of data written. This value can be useful for sequential write patterns but
	// less accurate for random write patterns. VolumeUsedInBytes is different from
	// the compressed size of the volume, which is the value that is used to calculate
	// your bill.
	//
	// This value is not available for volumes created prior to May 13, 2015, until
	// you store data on the volume.
	VolumeUsedInBytes *int64 `type:"long"`

	// An VolumeiSCSIAttributes object that represents a collection of iSCSI attributes
	// for one stored volume.
	VolumeiSCSIAttributes *VolumeiSCSIAttributes `type:"structure"`
}

// String returns the string representation
func (s StorediSCSIVolume) String() string {
	return awsutil.Prettify(s)
}

// A key-value pair that helps you manage, filter, and search for your resource.
// Allowed characters: letters, white space, and numbers, representable in UTF-8,
// and the following characters: + - = . _ : /
type Tag struct {
	_ struct{} `type:"structure"`

	// Tag key (String). The key can't start with aws:.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// Value of the tag key.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
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

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a virtual tape object.
type Tape struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server
	// side encryption. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// The ID of the pool that contains tapes that will be archived. The tapes in
	// this pool are archived in the S3 storage class that is associated with the
	// pool. When you use your backup application to eject the tape, the tape is
	// archived directly into the storage class (Glacier or Deep Archive) that corresponds
	// to the pool.
	//
	// Valid values: "GLACIER", "DEEP_ARCHIVE"
	PoolId *string `min:"1" type:"string"`

	// For archiving virtual tapes, indicates how much data remains to be uploaded
	// before archiving is complete.
	//
	// Range: 0 (not started) to 100 (complete).
	Progress *float64 `type:"double"`

	// The Amazon Resource Name (ARN) of the virtual tape.
	TapeARN *string `min:"50" type:"string"`

	// The barcode that identifies a specific virtual tape.
	TapeBarcode *string `min:"7" type:"string"`

	// The date the virtual tape was created.
	TapeCreatedDate *time.Time `type:"timestamp"`

	// The size, in bytes, of the virtual tape capacity.
	TapeSizeInBytes *int64 `type:"long"`

	// The current state of the virtual tape.
	TapeStatus *string `type:"string"`

	// The size, in bytes, of data stored on the virtual tape.
	//
	// This value is not available for tapes created prior to May 13, 2015.
	TapeUsedInBytes *int64 `type:"long"`

	// The virtual tape library (VTL) device that the virtual tape is associated
	// with.
	VTLDevice *string `min:"50" type:"string"`
}

// String returns the string representation
func (s Tape) String() string {
	return awsutil.Prettify(s)
}

// Represents a virtual tape that is archived in the virtual tape shelf (VTS).
type TapeArchive struct {
	_ struct{} `type:"structure"`

	// The time that the archiving of the virtual tape was completed.
	//
	// The default time stamp format is in the ISO8601 extended YYYY-MM-DD'T'HH:MM:SS'Z'
	// format.
	CompletionTime *time.Time `type:"timestamp"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server
	// side encryption. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// The ID of the pool that was used to archive the tape. The tapes in this pool
	// are archived in the S3 storage class that is associated with the pool.
	//
	// Valid values: "GLACIER", "DEEP_ARCHIVE"
	PoolId *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the tape gateway that the virtual tape
	// is being retrieved to.
	//
	// The virtual tape is retrieved from the virtual tape shelf (VTS).
	RetrievedTo *string `min:"50" type:"string"`

	// The Amazon Resource Name (ARN) of an archived virtual tape.
	TapeARN *string `min:"50" type:"string"`

	// The barcode that identifies the archived virtual tape.
	TapeBarcode *string `min:"7" type:"string"`

	// The date the virtual tape was created.
	TapeCreatedDate *time.Time `type:"timestamp"`

	// The size, in bytes, of the archived virtual tape.
	TapeSizeInBytes *int64 `type:"long"`

	// The current state of the archived virtual tape.
	TapeStatus *string `type:"string"`

	// The size, in bytes, of data stored on the virtual tape.
	//
	// This value is not available for tapes created prior to May 13, 2015.
	TapeUsedInBytes *int64 `type:"long"`
}

// String returns the string representation
func (s TapeArchive) String() string {
	return awsutil.Prettify(s)
}

// Describes a virtual tape.
type TapeInfo struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`

	// The ID of the pool that you want to add your tape to for archiving. The tape
	// in this pool is archived in the S3 storage class that is associated with
	// the pool. When you use your backup application to eject the tape, the tape
	// is archived directly into the storage class (Glacier or Deep Archive) that
	// corresponds to the pool.
	//
	// Valid values: "GLACIER", "DEEP_ARCHIVE"
	PoolId *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of a virtual tape.
	TapeARN *string `min:"50" type:"string"`

	// The barcode that identifies a specific virtual tape.
	TapeBarcode *string `min:"7" type:"string"`

	// The size, in bytes, of a virtual tape.
	TapeSizeInBytes *int64 `type:"long"`

	// The status of the tape.
	TapeStatus *string `type:"string"`
}

// String returns the string representation
func (s TapeInfo) String() string {
	return awsutil.Prettify(s)
}

// Describes a recovery point.
type TapeRecoveryPointInfo struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the virtual tape.
	TapeARN *string `min:"50" type:"string"`

	// The time when the point-in-time view of the virtual tape was replicated for
	// later recovery.
	//
	// The default time stamp format of the tape recovery point time is in the ISO8601
	// extended YYYY-MM-DD'T'HH:MM:SS'Z' format.
	TapeRecoveryPointTime *time.Time `type:"timestamp"`

	// The size, in bytes, of the virtual tapes to recover.
	TapeSizeInBytes *int64 `type:"long"`

	// The status of the virtual tapes.
	TapeStatus *string `type:"string"`
}

// String returns the string representation
func (s TapeRecoveryPointInfo) String() string {
	return awsutil.Prettify(s)
}

// Represents a device object associated with a tape gateway.
type VTLDevice struct {
	_ struct{} `type:"structure"`

	// A list of iSCSI information about a VTL device.
	DeviceiSCSIAttributes *DeviceiSCSIAttributes `type:"structure"`

	// Specifies the unique Amazon Resource Name (ARN) of the device (tape drive
	// or media changer).
	VTLDeviceARN *string `min:"50" type:"string"`

	// Specifies the model number of device that the VTL device emulates.
	VTLDeviceProductIdentifier *string `type:"string"`

	// Specifies the type of device that the VTL device emulates.
	VTLDeviceType *string `type:"string"`

	// Specifies the vendor of the device that the VTL device object emulates.
	VTLDeviceVendor *string `type:"string"`
}

// String returns the string representation
func (s VTLDevice) String() string {
	return awsutil.Prettify(s)
}

// Describes a storage volume object.
type VolumeInfo struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`

	// The unique identifier assigned to your gateway during activation. This ID
	// becomes part of the gateway Amazon Resource Name (ARN), which you use as
	// input for other operations.
	//
	// Valid Values: 50 to 500 lowercase letters, numbers, periods (.), and hyphens
	// (-).
	GatewayId *string `min:"12" type:"string"`

	// The Amazon Resource Name (ARN) for the storage volume. For example, the following
	// is a valid ARN:
	//
	// arn:aws:storagegateway:us-east-2:111122223333:gateway/sgw-12A3456B/volume/vol-1122AABB
	//
	// Valid Values: 50 to 500 lowercase letters, numbers, periods (.), and hyphens
	// (-).
	VolumeARN *string `min:"50" type:"string"`

	// One of the VolumeStatus values that indicates the state of the storage volume.
	VolumeAttachmentStatus *string `min:"3" type:"string"`

	// The unique identifier assigned to the volume. This ID becomes part of the
	// volume Amazon Resource Name (ARN), which you use as input for other operations.
	//
	// Valid Values: 50 to 500 lowercase letters, numbers, periods (.), and hyphens
	// (-).
	VolumeId *string `min:"12" type:"string"`

	// The size of the volume in bytes.
	//
	// Valid Values: 50 to 500 lowercase letters, numbers, periods (.), and hyphens
	// (-).
	VolumeSizeInBytes *int64 `type:"long"`

	// One of the VolumeType enumeration values describing the type of the volume.
	VolumeType *string `min:"3" type:"string"`
}

// String returns the string representation
func (s VolumeInfo) String() string {
	return awsutil.Prettify(s)
}

// Describes a storage volume recovery point object.
type VolumeRecoveryPointInfo struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the volume target.
	VolumeARN *string `min:"50" type:"string"`

	// The time the recovery point was taken.
	VolumeRecoveryPointTime *string `type:"string"`

	// The size of the volume in bytes.
	VolumeSizeInBytes *int64 `type:"long"`

	// The size of the data stored on the volume in bytes.
	//
	// This value is not available for volumes created prior to May 13, 2015, until
	// you store data on the volume.
	VolumeUsageInBytes *int64 `type:"long"`
}

// String returns the string representation
func (s VolumeRecoveryPointInfo) String() string {
	return awsutil.Prettify(s)
}

// Lists iSCSI information about a volume.
type VolumeiSCSIAttributes struct {
	_ struct{} `type:"structure"`

	// Indicates whether mutual CHAP is enabled for the iSCSI target.
	ChapEnabled *bool `type:"boolean"`

	// The logical disk number.
	LunNumber *int64 `min:"1" type:"integer"`

	// The network interface identifier.
	NetworkInterfaceId *string `type:"string"`

	// The port used to communicate with iSCSI targets.
	NetworkInterfacePort *int64 `type:"integer"`

	// The Amazon Resource Name (ARN) of the volume target.
	TargetARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s VolumeiSCSIAttributes) String() string {
	return awsutil.Prettify(s)
}
