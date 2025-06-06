// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttributeMatchingModel string

// Enum values for AttributeMatchingModel
const (
	AttributeMatchingModelOneToOne   AttributeMatchingModel = "ONE_TO_ONE"
	AttributeMatchingModelManyToMany AttributeMatchingModel = "MANY_TO_MANY"
)

// Values returns all known values for AttributeMatchingModel. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AttributeMatchingModel) Values() []AttributeMatchingModel {
	return []AttributeMatchingModel{
		"ONE_TO_ONE",
		"MANY_TO_MANY",
	}
}

type DeleteUniqueIdErrorType string

// Enum values for DeleteUniqueIdErrorType
const (
	DeleteUniqueIdErrorTypeServiceError    DeleteUniqueIdErrorType = "SERVICE_ERROR"
	DeleteUniqueIdErrorTypeValidationError DeleteUniqueIdErrorType = "VALIDATION_ERROR"
)

// Values returns all known values for DeleteUniqueIdErrorType. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeleteUniqueIdErrorType) Values() []DeleteUniqueIdErrorType {
	return []DeleteUniqueIdErrorType{
		"SERVICE_ERROR",
		"VALIDATION_ERROR",
	}
}

type DeleteUniqueIdStatus string

// Enum values for DeleteUniqueIdStatus
const (
	DeleteUniqueIdStatusCompleted DeleteUniqueIdStatus = "COMPLETED"
	DeleteUniqueIdStatusAccepted  DeleteUniqueIdStatus = "ACCEPTED"
)

// Values returns all known values for DeleteUniqueIdStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeleteUniqueIdStatus) Values() []DeleteUniqueIdStatus {
	return []DeleteUniqueIdStatus{
		"COMPLETED",
		"ACCEPTED",
	}
}

type IdMappingType string

// Enum values for IdMappingType
const (
	IdMappingTypeProvider  IdMappingType = "PROVIDER"
	IdMappingTypeRuleBased IdMappingType = "RULE_BASED"
)

// Values returns all known values for IdMappingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IdMappingType) Values() []IdMappingType {
	return []IdMappingType{
		"PROVIDER",
		"RULE_BASED",
	}
}

type IdMappingWorkflowRuleDefinitionType string

// Enum values for IdMappingWorkflowRuleDefinitionType
const (
	IdMappingWorkflowRuleDefinitionTypeSource IdMappingWorkflowRuleDefinitionType = "SOURCE"
	IdMappingWorkflowRuleDefinitionTypeTarget IdMappingWorkflowRuleDefinitionType = "TARGET"
)

// Values returns all known values for IdMappingWorkflowRuleDefinitionType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IdMappingWorkflowRuleDefinitionType) Values() []IdMappingWorkflowRuleDefinitionType {
	return []IdMappingWorkflowRuleDefinitionType{
		"SOURCE",
		"TARGET",
	}
}

type IdNamespaceType string

// Enum values for IdNamespaceType
const (
	IdNamespaceTypeSource IdNamespaceType = "SOURCE"
	IdNamespaceTypeTarget IdNamespaceType = "TARGET"
)

// Values returns all known values for IdNamespaceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IdNamespaceType) Values() []IdNamespaceType {
	return []IdNamespaceType{
		"SOURCE",
		"TARGET",
	}
}

type IncrementalRunType string

// Enum values for IncrementalRunType
const (
	IncrementalRunTypeImmediate IncrementalRunType = "IMMEDIATE"
)

// Values returns all known values for IncrementalRunType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IncrementalRunType) Values() []IncrementalRunType {
	return []IncrementalRunType{
		"IMMEDIATE",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusRunning   JobStatus = "RUNNING"
	JobStatusSucceeded JobStatus = "SUCCEEDED"
	JobStatusFailed    JobStatus = "FAILED"
	JobStatusQueued    JobStatus = "QUEUED"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"RUNNING",
		"SUCCEEDED",
		"FAILED",
		"QUEUED",
	}
}

type MatchPurpose string

// Enum values for MatchPurpose
const (
	MatchPurposeIdentifierGeneration MatchPurpose = "IDENTIFIER_GENERATION"
	MatchPurposeIndexing             MatchPurpose = "INDEXING"
)

// Values returns all known values for MatchPurpose. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MatchPurpose) Values() []MatchPurpose {
	return []MatchPurpose{
		"IDENTIFIER_GENERATION",
		"INDEXING",
	}
}

type ProcessingType string

// Enum values for ProcessingType
const (
	ProcessingTypeConsistent       ProcessingType = "CONSISTENT"
	ProcessingTypeEventual         ProcessingType = "EVENTUAL"
	ProcessingTypeEventualNoLookup ProcessingType = "EVENTUAL_NO_LOOKUP"
)

// Values returns all known values for ProcessingType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProcessingType) Values() []ProcessingType {
	return []ProcessingType{
		"CONSISTENT",
		"EVENTUAL",
		"EVENTUAL_NO_LOOKUP",
	}
}

type RecordMatchingModel string

// Enum values for RecordMatchingModel
const (
	RecordMatchingModelOneSourceToOneTarget  RecordMatchingModel = "ONE_SOURCE_TO_ONE_TARGET"
	RecordMatchingModelManySourceToOneTarget RecordMatchingModel = "MANY_SOURCE_TO_ONE_TARGET"
)

// Values returns all known values for RecordMatchingModel. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RecordMatchingModel) Values() []RecordMatchingModel {
	return []RecordMatchingModel{
		"ONE_SOURCE_TO_ONE_TARGET",
		"MANY_SOURCE_TO_ONE_TARGET",
	}
}

type ResolutionType string

// Enum values for ResolutionType
const (
	ResolutionTypeRuleMatching ResolutionType = "RULE_MATCHING"
	ResolutionTypeMlMatching   ResolutionType = "ML_MATCHING"
	ResolutionTypeProvider     ResolutionType = "PROVIDER"
)

// Values returns all known values for ResolutionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolutionType) Values() []ResolutionType {
	return []ResolutionType{
		"RULE_MATCHING",
		"ML_MATCHING",
		"PROVIDER",
	}
}

type SchemaAttributeType string

// Enum values for SchemaAttributeType
const (
	SchemaAttributeTypeName              SchemaAttributeType = "NAME"
	SchemaAttributeTypeNameFirst         SchemaAttributeType = "NAME_FIRST"
	SchemaAttributeTypeNameMiddle        SchemaAttributeType = "NAME_MIDDLE"
	SchemaAttributeTypeNameLast          SchemaAttributeType = "NAME_LAST"
	SchemaAttributeTypeAddress           SchemaAttributeType = "ADDRESS"
	SchemaAttributeTypeAddressStreet1    SchemaAttributeType = "ADDRESS_STREET1"
	SchemaAttributeTypeAddressStreet2    SchemaAttributeType = "ADDRESS_STREET2"
	SchemaAttributeTypeAddressStreet3    SchemaAttributeType = "ADDRESS_STREET3"
	SchemaAttributeTypeAddressCity       SchemaAttributeType = "ADDRESS_CITY"
	SchemaAttributeTypeAddressState      SchemaAttributeType = "ADDRESS_STATE"
	SchemaAttributeTypeAddressCountry    SchemaAttributeType = "ADDRESS_COUNTRY"
	SchemaAttributeTypeAddressPostalcode SchemaAttributeType = "ADDRESS_POSTALCODE"
	SchemaAttributeTypePhone             SchemaAttributeType = "PHONE"
	SchemaAttributeTypePhoneNumber       SchemaAttributeType = "PHONE_NUMBER"
	SchemaAttributeTypePhoneCountrycode  SchemaAttributeType = "PHONE_COUNTRYCODE"
	SchemaAttributeTypeEmailAddress      SchemaAttributeType = "EMAIL_ADDRESS"
	SchemaAttributeTypeUniqueId          SchemaAttributeType = "UNIQUE_ID"
	SchemaAttributeTypeDate              SchemaAttributeType = "DATE"
	SchemaAttributeTypeString            SchemaAttributeType = "STRING"
	SchemaAttributeTypeProviderId        SchemaAttributeType = "PROVIDER_ID"
	SchemaAttributeTypeIpv4              SchemaAttributeType = "IPV4"
	SchemaAttributeTypeIpv6              SchemaAttributeType = "IPV6"
	SchemaAttributeTypeMaid              SchemaAttributeType = "MAID"
)

// Values returns all known values for SchemaAttributeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SchemaAttributeType) Values() []SchemaAttributeType {
	return []SchemaAttributeType{
		"NAME",
		"NAME_FIRST",
		"NAME_MIDDLE",
		"NAME_LAST",
		"ADDRESS",
		"ADDRESS_STREET1",
		"ADDRESS_STREET2",
		"ADDRESS_STREET3",
		"ADDRESS_CITY",
		"ADDRESS_STATE",
		"ADDRESS_COUNTRY",
		"ADDRESS_POSTALCODE",
		"PHONE",
		"PHONE_NUMBER",
		"PHONE_COUNTRYCODE",
		"EMAIL_ADDRESS",
		"UNIQUE_ID",
		"DATE",
		"STRING",
		"PROVIDER_ID",
		"IPV4",
		"IPV6",
		"MAID",
	}
}

type ServiceType string

// Enum values for ServiceType
const (
	ServiceTypeAssignment ServiceType = "ASSIGNMENT"
	ServiceTypeIdMapping  ServiceType = "ID_MAPPING"
)

// Values returns all known values for ServiceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ServiceType) Values() []ServiceType {
	return []ServiceType{
		"ASSIGNMENT",
		"ID_MAPPING",
	}
}

type StatementEffect string

// Enum values for StatementEffect
const (
	StatementEffectAllow StatementEffect = "Allow"
	StatementEffectDeny  StatementEffect = "Deny"
)

// Values returns all known values for StatementEffect. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StatementEffect) Values() []StatementEffect {
	return []StatementEffect{
		"Allow",
		"Deny",
	}
}
