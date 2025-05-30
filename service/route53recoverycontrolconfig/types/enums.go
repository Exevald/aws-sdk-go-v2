// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type NetworkType string

// Enum values for NetworkType
const (
	NetworkTypeIpv4      NetworkType = "IPV4"
	NetworkTypeDualstack NetworkType = "DUALSTACK"
)

// Values returns all known values for NetworkType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (NetworkType) Values() []NetworkType {
	return []NetworkType{
		"IPV4",
		"DUALSTACK",
	}
}

type RuleType string

// Enum values for RuleType
const (
	RuleTypeAtleast RuleType = "ATLEAST"
	RuleTypeAnd     RuleType = "AND"
	RuleTypeOr      RuleType = "OR"
)

// Values returns all known values for RuleType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RuleType) Values() []RuleType {
	return []RuleType{
		"ATLEAST",
		"AND",
		"OR",
	}
}

type Status string

// Enum values for Status
const (
	StatusPending         Status = "PENDING"
	StatusDeployed        Status = "DEPLOYED"
	StatusPendingDeletion Status = "PENDING_DELETION"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"PENDING",
		"DEPLOYED",
		"PENDING_DELETION",
	}
}
