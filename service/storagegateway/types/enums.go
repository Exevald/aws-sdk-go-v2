// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActiveDirectoryStatus string

// Enum values for ActiveDirectoryStatus
const (
	ActiveDirectoryStatusAccessDenied            ActiveDirectoryStatus = "ACCESS_DENIED"
	ActiveDirectoryStatusDetached                ActiveDirectoryStatus = "DETACHED"
	ActiveDirectoryStatusJoined                  ActiveDirectoryStatus = "JOINED"
	ActiveDirectoryStatusJoining                 ActiveDirectoryStatus = "JOINING"
	ActiveDirectoryStatusNetworkError            ActiveDirectoryStatus = "NETWORK_ERROR"
	ActiveDirectoryStatusTimeout                 ActiveDirectoryStatus = "TIMEOUT"
	ActiveDirectoryStatusUnknownError            ActiveDirectoryStatus = "UNKNOWN_ERROR"
	ActiveDirectoryStatusInsufficientPermissions ActiveDirectoryStatus = "INSUFFICIENT_PERMISSIONS"
)

// Values returns all known values for ActiveDirectoryStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ActiveDirectoryStatus) Values() []ActiveDirectoryStatus {
	return []ActiveDirectoryStatus{
		"ACCESS_DENIED",
		"DETACHED",
		"JOINED",
		"JOINING",
		"NETWORK_ERROR",
		"TIMEOUT",
		"UNKNOWN_ERROR",
		"INSUFFICIENT_PERMISSIONS",
	}
}

type AutomaticUpdatePolicy string

// Enum values for AutomaticUpdatePolicy
const (
	AutomaticUpdatePolicyAllVersions           AutomaticUpdatePolicy = "ALL_VERSIONS"
	AutomaticUpdatePolicyEmergencyVersionsOnly AutomaticUpdatePolicy = "EMERGENCY_VERSIONS_ONLY"
)

// Values returns all known values for AutomaticUpdatePolicy. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AutomaticUpdatePolicy) Values() []AutomaticUpdatePolicy {
	return []AutomaticUpdatePolicy{
		"ALL_VERSIONS",
		"EMERGENCY_VERSIONS_ONLY",
	}
}

type AvailabilityMonitorTestStatus string

// Enum values for AvailabilityMonitorTestStatus
const (
	AvailabilityMonitorTestStatusComplete AvailabilityMonitorTestStatus = "COMPLETE"
	AvailabilityMonitorTestStatusFailed   AvailabilityMonitorTestStatus = "FAILED"
	AvailabilityMonitorTestStatusPending  AvailabilityMonitorTestStatus = "PENDING"
)

// Values returns all known values for AvailabilityMonitorTestStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AvailabilityMonitorTestStatus) Values() []AvailabilityMonitorTestStatus {
	return []AvailabilityMonitorTestStatus{
		"COMPLETE",
		"FAILED",
		"PENDING",
	}
}

type CacheReportFilterName string

// Enum values for CacheReportFilterName
const (
	CacheReportFilterNameUploadState         CacheReportFilterName = "UploadState"
	CacheReportFilterNameUploadFailureReason CacheReportFilterName = "UploadFailureReason"
)

// Values returns all known values for CacheReportFilterName. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CacheReportFilterName) Values() []CacheReportFilterName {
	return []CacheReportFilterName{
		"UploadState",
		"UploadFailureReason",
	}
}

type CacheReportStatus string

// Enum values for CacheReportStatus
const (
	CacheReportStatusInProgress CacheReportStatus = "IN_PROGRESS"
	CacheReportStatusCompleted  CacheReportStatus = "COMPLETED"
	CacheReportStatusCanceled   CacheReportStatus = "CANCELED"
	CacheReportStatusFailed     CacheReportStatus = "FAILED"
	CacheReportStatusError      CacheReportStatus = "ERROR"
)

// Values returns all known values for CacheReportStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CacheReportStatus) Values() []CacheReportStatus {
	return []CacheReportStatus{
		"IN_PROGRESS",
		"COMPLETED",
		"CANCELED",
		"FAILED",
		"ERROR",
	}
}

type CaseSensitivity string

// Enum values for CaseSensitivity
const (
	CaseSensitivityClientSpecified CaseSensitivity = "ClientSpecified"
	CaseSensitivityCaseSensitive   CaseSensitivity = "CaseSensitive"
)

// Values returns all known values for CaseSensitivity. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CaseSensitivity) Values() []CaseSensitivity {
	return []CaseSensitivity{
		"ClientSpecified",
		"CaseSensitive",
	}
}

type EncryptionType string

// Enum values for EncryptionType
const (
	EncryptionTypeSseS3   EncryptionType = "SseS3"
	EncryptionTypeSseKms  EncryptionType = "SseKms"
	EncryptionTypeDsseKms EncryptionType = "DsseKms"
)

// Values returns all known values for EncryptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionType) Values() []EncryptionType {
	return []EncryptionType{
		"SseS3",
		"SseKms",
		"DsseKms",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeActivationKeyExpired              ErrorCode = "ActivationKeyExpired"
	ErrorCodeActivationKeyInvalid              ErrorCode = "ActivationKeyInvalid"
	ErrorCodeActivationKeyNotFound             ErrorCode = "ActivationKeyNotFound"
	ErrorCodeGatewayInternalError              ErrorCode = "GatewayInternalError"
	ErrorCodeGatewayNotConnected               ErrorCode = "GatewayNotConnected"
	ErrorCodeGatewayNotFound                   ErrorCode = "GatewayNotFound"
	ErrorCodeGatewayProxyNetworkConnectionBusy ErrorCode = "GatewayProxyNetworkConnectionBusy"
	ErrorCodeAuthenticationFailure             ErrorCode = "AuthenticationFailure"
	ErrorCodeBandwidthThrottleScheduleNotFound ErrorCode = "BandwidthThrottleScheduleNotFound"
	ErrorCodeBlocked                           ErrorCode = "Blocked"
	ErrorCodeCannotExportSnapshot              ErrorCode = "CannotExportSnapshot"
	ErrorCodeChapCredentialNotFound            ErrorCode = "ChapCredentialNotFound"
	ErrorCodeDiskAlreadyAllocated              ErrorCode = "DiskAlreadyAllocated"
	ErrorCodeDiskDoesNotExist                  ErrorCode = "DiskDoesNotExist"
	ErrorCodeDiskSizeGreaterThanVolumeMaxSize  ErrorCode = "DiskSizeGreaterThanVolumeMaxSize"
	ErrorCodeDiskSizeLessThanVolumeSize        ErrorCode = "DiskSizeLessThanVolumeSize"
	ErrorCodeDiskSizeNotGigAligned             ErrorCode = "DiskSizeNotGigAligned"
	ErrorCodeDuplicateCertificateInfo          ErrorCode = "DuplicateCertificateInfo"
	ErrorCodeDuplicateSchedule                 ErrorCode = "DuplicateSchedule"
	ErrorCodeEndpointNotFound                  ErrorCode = "EndpointNotFound"
	ErrorCodeIAMNotSupported                   ErrorCode = "IAMNotSupported"
	ErrorCodeInitiatorInvalid                  ErrorCode = "InitiatorInvalid"
	ErrorCodeInitiatorNotFound                 ErrorCode = "InitiatorNotFound"
	ErrorCodeInternalError                     ErrorCode = "InternalError"
	ErrorCodeInvalidGateway                    ErrorCode = "InvalidGateway"
	ErrorCodeInvalidEndpoint                   ErrorCode = "InvalidEndpoint"
	ErrorCodeInvalidParameters                 ErrorCode = "InvalidParameters"
	ErrorCodeInvalidSchedule                   ErrorCode = "InvalidSchedule"
	ErrorCodeLocalStorageLimitExceeded         ErrorCode = "LocalStorageLimitExceeded"
	ErrorCodeLunAlreadyAllocated               ErrorCode = "LunAlreadyAllocated "
	ErrorCodeLunInvalid                        ErrorCode = "LunInvalid"
	ErrorCodeJoinDomainInProgress              ErrorCode = "JoinDomainInProgress"
	ErrorCodeMaximumContentLengthExceeded      ErrorCode = "MaximumContentLengthExceeded"
	ErrorCodeMaximumTapeCartridgeCountExceeded ErrorCode = "MaximumTapeCartridgeCountExceeded"
	ErrorCodeMaximumVolumeCountExceeded        ErrorCode = "MaximumVolumeCountExceeded"
	ErrorCodeNetworkConfigurationChanged       ErrorCode = "NetworkConfigurationChanged"
	ErrorCodeNoDisksAvailable                  ErrorCode = "NoDisksAvailable"
	ErrorCodeNotImplemented                    ErrorCode = "NotImplemented"
	ErrorCodeNotSupported                      ErrorCode = "NotSupported"
	ErrorCodeOperationAborted                  ErrorCode = "OperationAborted"
	ErrorCodeOutdatedGateway                   ErrorCode = "OutdatedGateway"
	ErrorCodeParametersNotImplemented          ErrorCode = "ParametersNotImplemented"
	ErrorCodeRegionInvalid                     ErrorCode = "RegionInvalid"
	ErrorCodeRequestTimeout                    ErrorCode = "RequestTimeout"
	ErrorCodeServiceUnavailable                ErrorCode = "ServiceUnavailable"
	ErrorCodeSnapshotDeleted                   ErrorCode = "SnapshotDeleted"
	ErrorCodeSnapshotIdInvalid                 ErrorCode = "SnapshotIdInvalid"
	ErrorCodeSnapshotInProgress                ErrorCode = "SnapshotInProgress"
	ErrorCodeSnapshotNotFound                  ErrorCode = "SnapshotNotFound"
	ErrorCodeSnapshotScheduleNotFound          ErrorCode = "SnapshotScheduleNotFound"
	ErrorCodeStagingAreaFull                   ErrorCode = "StagingAreaFull"
	ErrorCodeStorageFailure                    ErrorCode = "StorageFailure"
	ErrorCodeTapeCartridgeNotFound             ErrorCode = "TapeCartridgeNotFound"
	ErrorCodeTargetAlreadyExists               ErrorCode = "TargetAlreadyExists"
	ErrorCodeTargetInvalid                     ErrorCode = "TargetInvalid"
	ErrorCodeTargetNotFound                    ErrorCode = "TargetNotFound"
	ErrorCodeUnauthorizedOperation             ErrorCode = "UnauthorizedOperation"
	ErrorCodeVolumeAlreadyExists               ErrorCode = "VolumeAlreadyExists"
	ErrorCodeVolumeIdInvalid                   ErrorCode = "VolumeIdInvalid"
	ErrorCodeVolumeInUse                       ErrorCode = "VolumeInUse"
	ErrorCodeVolumeNotFound                    ErrorCode = "VolumeNotFound"
	ErrorCodeVolumeNotReady                    ErrorCode = "VolumeNotReady"
)

// Values returns all known values for ErrorCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"ActivationKeyExpired",
		"ActivationKeyInvalid",
		"ActivationKeyNotFound",
		"GatewayInternalError",
		"GatewayNotConnected",
		"GatewayNotFound",
		"GatewayProxyNetworkConnectionBusy",
		"AuthenticationFailure",
		"BandwidthThrottleScheduleNotFound",
		"Blocked",
		"CannotExportSnapshot",
		"ChapCredentialNotFound",
		"DiskAlreadyAllocated",
		"DiskDoesNotExist",
		"DiskSizeGreaterThanVolumeMaxSize",
		"DiskSizeLessThanVolumeSize",
		"DiskSizeNotGigAligned",
		"DuplicateCertificateInfo",
		"DuplicateSchedule",
		"EndpointNotFound",
		"IAMNotSupported",
		"InitiatorInvalid",
		"InitiatorNotFound",
		"InternalError",
		"InvalidGateway",
		"InvalidEndpoint",
		"InvalidParameters",
		"InvalidSchedule",
		"LocalStorageLimitExceeded",
		"LunAlreadyAllocated ",
		"LunInvalid",
		"JoinDomainInProgress",
		"MaximumContentLengthExceeded",
		"MaximumTapeCartridgeCountExceeded",
		"MaximumVolumeCountExceeded",
		"NetworkConfigurationChanged",
		"NoDisksAvailable",
		"NotImplemented",
		"NotSupported",
		"OperationAborted",
		"OutdatedGateway",
		"ParametersNotImplemented",
		"RegionInvalid",
		"RequestTimeout",
		"ServiceUnavailable",
		"SnapshotDeleted",
		"SnapshotIdInvalid",
		"SnapshotInProgress",
		"SnapshotNotFound",
		"SnapshotScheduleNotFound",
		"StagingAreaFull",
		"StorageFailure",
		"TapeCartridgeNotFound",
		"TargetAlreadyExists",
		"TargetInvalid",
		"TargetNotFound",
		"UnauthorizedOperation",
		"VolumeAlreadyExists",
		"VolumeIdInvalid",
		"VolumeInUse",
		"VolumeNotFound",
		"VolumeNotReady",
	}
}

type FileShareType string

// Enum values for FileShareType
const (
	FileShareTypeNfs FileShareType = "NFS"
	FileShareTypeSmb FileShareType = "SMB"
)

// Values returns all known values for FileShareType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FileShareType) Values() []FileShareType {
	return []FileShareType{
		"NFS",
		"SMB",
	}
}

type GatewayCapacity string

// Enum values for GatewayCapacity
const (
	GatewayCapacitySmall  GatewayCapacity = "Small"
	GatewayCapacityMedium GatewayCapacity = "Medium"
	GatewayCapacityLarge  GatewayCapacity = "Large"
)

// Values returns all known values for GatewayCapacity. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (GatewayCapacity) Values() []GatewayCapacity {
	return []GatewayCapacity{
		"Small",
		"Medium",
		"Large",
	}
}

type HostEnvironment string

// Enum values for HostEnvironment
const (
	HostEnvironmentVmware   HostEnvironment = "VMWARE"
	HostEnvironmentHyperV   HostEnvironment = "HYPER-V"
	HostEnvironmentEc2      HostEnvironment = "EC2"
	HostEnvironmentKvm      HostEnvironment = "KVM"
	HostEnvironmentOther    HostEnvironment = "OTHER"
	HostEnvironmentSnowball HostEnvironment = "SNOWBALL"
)

// Values returns all known values for HostEnvironment. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (HostEnvironment) Values() []HostEnvironment {
	return []HostEnvironment{
		"VMWARE",
		"HYPER-V",
		"EC2",
		"KVM",
		"OTHER",
		"SNOWBALL",
	}
}

type ObjectACL string

// Enum values for ObjectACL
const (
	ObjectACLPrivate                ObjectACL = "private"
	ObjectACLPublicRead             ObjectACL = "public-read"
	ObjectACLPublicReadWrite        ObjectACL = "public-read-write"
	ObjectACLAuthenticatedRead      ObjectACL = "authenticated-read"
	ObjectACLBucketOwnerRead        ObjectACL = "bucket-owner-read"
	ObjectACLBucketOwnerFullControl ObjectACL = "bucket-owner-full-control"
	ObjectACLAwsExecRead            ObjectACL = "aws-exec-read"
)

// Values returns all known values for ObjectACL. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ObjectACL) Values() []ObjectACL {
	return []ObjectACL{
		"private",
		"public-read",
		"public-read-write",
		"authenticated-read",
		"bucket-owner-read",
		"bucket-owner-full-control",
		"aws-exec-read",
	}
}

type PoolStatus string

// Enum values for PoolStatus
const (
	PoolStatusActive  PoolStatus = "ACTIVE"
	PoolStatusDeleted PoolStatus = "DELETED"
)

// Values returns all known values for PoolStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PoolStatus) Values() []PoolStatus {
	return []PoolStatus{
		"ACTIVE",
		"DELETED",
	}
}

type RetentionLockType string

// Enum values for RetentionLockType
const (
	RetentionLockTypeCompliance RetentionLockType = "COMPLIANCE"
	RetentionLockTypeGovernance RetentionLockType = "GOVERNANCE"
	RetentionLockTypeNone       RetentionLockType = "NONE"
)

// Values returns all known values for RetentionLockType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RetentionLockType) Values() []RetentionLockType {
	return []RetentionLockType{
		"COMPLIANCE",
		"GOVERNANCE",
		"NONE",
	}
}

type SMBSecurityStrategy string

// Enum values for SMBSecurityStrategy
const (
	SMBSecurityStrategyClientSpecified             SMBSecurityStrategy = "ClientSpecified"
	SMBSecurityStrategyMandatorySigning            SMBSecurityStrategy = "MandatorySigning"
	SMBSecurityStrategyMandatoryEncryption         SMBSecurityStrategy = "MandatoryEncryption"
	SMBSecurityStrategyMandatoryEncryptionNoAes128 SMBSecurityStrategy = "MandatoryEncryptionNoAes128"
)

// Values returns all known values for SMBSecurityStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SMBSecurityStrategy) Values() []SMBSecurityStrategy {
	return []SMBSecurityStrategy{
		"ClientSpecified",
		"MandatorySigning",
		"MandatoryEncryption",
		"MandatoryEncryptionNoAes128",
	}
}

type TapeStorageClass string

// Enum values for TapeStorageClass
const (
	TapeStorageClassDeepArchive TapeStorageClass = "DEEP_ARCHIVE"
	TapeStorageClassGlacier     TapeStorageClass = "GLACIER"
)

// Values returns all known values for TapeStorageClass. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TapeStorageClass) Values() []TapeStorageClass {
	return []TapeStorageClass{
		"DEEP_ARCHIVE",
		"GLACIER",
	}
}
