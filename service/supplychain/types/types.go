// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The BillOfMaterialsImportJob details.
type BillOfMaterialsImportJob struct {

	// The BillOfMaterialsImportJob instanceId.
	//
	// This member is required.
	InstanceId *string

	// The BillOfMaterialsImportJob jobId.
	//
	// This member is required.
	JobId *string

	// The S3 URI from which the CSV is read.
	//
	// This member is required.
	S3uri *string

	// The BillOfMaterialsImportJob ConfigurationJobStatus.
	//
	// This member is required.
	Status ConfigurationJobStatus

	// When the BillOfMaterialsImportJob has reached a terminal state, there will be a
	// message.
	Message *string

	noSmithyDocumentSerde
}

// The data integration event details.
type DataIntegrationEvent struct {

	// Event identifier (for example, orderId for InboundOrder) used for data sharding
	// or partitioning.
	//
	// This member is required.
	EventGroupId *string

	// The unique event identifier.
	//
	// This member is required.
	EventId *string

	// The event timestamp (in epoch seconds).
	//
	// This member is required.
	EventTimestamp *time.Time

	// The data event type.
	//
	// This member is required.
	EventType DataIntegrationEventType

	// The AWS Supply Chain instance identifier.
	//
	// This member is required.
	InstanceId *string

	// The target dataset details for a DATASET event type.
	DatasetTargetDetails *DataIntegrationEventDatasetTargetDetails

	noSmithyDocumentSerde
}

// The target dataset load execution details.
type DataIntegrationEventDatasetLoadExecutionDetails struct {

	// The event load execution status to target dataset.
	//
	// This member is required.
	Status DataIntegrationEventDatasetLoadStatus

	// The failure message (if any) of failed event load execution to dataset.
	Message *string

	noSmithyDocumentSerde
}

// The target dataset configuration for a DATASET event type.
type DataIntegrationEventDatasetTargetConfiguration struct {

	// The datalake dataset ARN identifier.
	//
	// This member is required.
	DatasetIdentifier *string

	// The target dataset load operation type.
	//
	// This member is required.
	OperationType DataIntegrationEventDatasetOperationType

	noSmithyDocumentSerde
}

// The target dataset details for a DATASET event type.
type DataIntegrationEventDatasetTargetDetails struct {

	// The datalake dataset ARN identifier.
	//
	// This member is required.
	DatasetIdentifier *string

	// The target dataset load execution.
	//
	// This member is required.
	DatasetLoadExecution *DataIntegrationEventDatasetLoadExecutionDetails

	// The target dataset load operation type. The available options are:
	//
	//   - APPEND - Add new records to the dataset. Noted that this operation type
	//   will just try to append records as-is without any primary key or partition
	//   constraints.
	//
	//   - UPSERT - Modify existing records in the dataset with primary key
	//   configured, events for datasets without primary keys are not allowed. If event
	//   data contains primary keys that match records in the dataset within same
	//   partition, then those existing records (in that partition) will be updated. If
	//   primary keys do not match, new records will be added. Note that if dataset
	//   contain records with duplicate primary key values in the same partition, those
	//   duplicate records will be deduped into one updated record.
	//
	//   - DELETE - Remove existing records in the dataset with primary key
	//   configured, events for datasets without primary keys are not allowed. If event
	//   data contains primary keys that match records in the dataset within same
	//   partition, then those existing records (in that partition) will be deleted. If
	//   primary keys do not match, no actions will be done. Note that if dataset contain
	//   records with duplicate primary key values in the same partition, all those
	//   duplicates will be removed.
	//
	// This member is required.
	OperationType DataIntegrationEventDatasetOperationType

	noSmithyDocumentSerde
}

// The DataIntegrationFlow details.
type DataIntegrationFlow struct {

	// The DataIntegrationFlow creation timestamp.
	//
	// This member is required.
	CreatedTime *time.Time

	// The DataIntegrationFlow instance ID.
	//
	// This member is required.
	InstanceId *string

	// The DataIntegrationFlow last modified timestamp.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The DataIntegrationFlow name.
	//
	// This member is required.
	Name *string

	// The DataIntegrationFlow source configurations.
	//
	// This member is required.
	Sources []DataIntegrationFlowSource

	// The DataIntegrationFlow target configuration.
	//
	// This member is required.
	Target *DataIntegrationFlowTarget

	// The DataIntegrationFlow transformation configurations.
	//
	// This member is required.
	Transformation *DataIntegrationFlowTransformation

	noSmithyDocumentSerde
}

// The dataset options used in dataset source and target configurations.
type DataIntegrationFlowDatasetOptions struct {

	// The option to perform deduplication on data records sharing same primary key
	// values. If disabled, transformed data with duplicate primary key values will
	// ingest into dataset, for datasets within asc namespace, such duplicates will
	// cause ingestion fail. If enabled without dedupeStrategy, deduplication is done
	// by retaining a random data record among those sharing the same primary key
	// values. If enabled with dedupeStragtegy, the deduplication is done following the
	// strategy.
	//
	// Note that target dataset may have partition configured, when dedupe is enabled,
	// it only dedupe against primary keys and retain only one record out of those
	// duplicates regardless of its partition status.
	DedupeRecords *bool

	// The deduplication strategy to dedupe the data records sharing same primary key
	// values of the target dataset. This strategy only applies to target dataset with
	// primary keys and with dedupeRecords option enabled. If transformed data still
	// got duplicates after the dedupeStrategy evaluation, a random data record is
	// chosen to be retained.
	DedupeStrategy *DataIntegrationFlowDedupeStrategy

	// The target dataset's data load type. This only affects how source S3 files are
	// selected in the S3-to-dataset flow.
	//
	//   - REPLACE - Target dataset will get replaced with the new file added under
	//   the source s3 prefix.
	//
	//   - INCREMENTAL - Target dataset will get updated with the up-to-date content
	//   under S3 prefix incorporating any file additions or removals there.
	LoadType DataIntegrationFlowLoadType

	noSmithyDocumentSerde
}

// The details of a flow execution with dataset source.
type DataIntegrationFlowDatasetSource struct {

	// The ARN of the dataset source.
	//
	// This member is required.
	DatasetIdentifier *string

	noSmithyDocumentSerde
}

// The dataset DataIntegrationFlow source configuration parameters.
type DataIntegrationFlowDatasetSourceConfiguration struct {

	// The ARN of the dataset.
	//
	// This member is required.
	DatasetIdentifier *string

	// The dataset DataIntegrationFlow source options.
	Options *DataIntegrationFlowDatasetOptions

	noSmithyDocumentSerde
}

// The dataset DataIntegrationFlow target configuration parameters.
type DataIntegrationFlowDatasetTargetConfiguration struct {

	// The dataset ARN.
	//
	// This member is required.
	DatasetIdentifier *string

	// The dataset DataIntegrationFlow target options.
	Options *DataIntegrationFlowDatasetOptions

	noSmithyDocumentSerde
}

// The deduplication strategy details.
type DataIntegrationFlowDedupeStrategy struct {

	// The type of the deduplication strategy.
	//
	//   - FIELD_PRIORITY - Field priority configuration for the deduplication
	//   strategy specifies an ordered list of fields used to tie-break the data records
	//   sharing the same primary key values. Fields earlier in the list have higher
	//   priority for evaluation. For each field, the sort order determines whether to
	//   retain data record with larger or smaller field value.
	//
	// This member is required.
	Type DataIntegrationFlowDedupeStrategyType

	// The field priority deduplication strategy.
	FieldPriority *DataIntegrationFlowFieldPriorityDedupeStrategyConfiguration

	noSmithyDocumentSerde
}

// The flow execution details.
type DataIntegrationFlowExecution struct {

	// The flow executionId.
	//
	// This member is required.
	ExecutionId *string

	// The flow execution's flowName.
	//
	// This member is required.
	FlowName *string

	// The flow execution's instanceId.
	//
	// This member is required.
	InstanceId *string

	// The flow execution end timestamp.
	EndTime *time.Time

	// The failure message (if any) of failed flow execution.
	Message *string

	// The flow execution output metadata.
	OutputMetadata *DataIntegrationFlowExecutionOutputMetadata

	// The source information for a flow execution.
	SourceInfo *DataIntegrationFlowExecutionSourceInfo

	// The flow execution start timestamp.
	StartTime *time.Time

	// The status of flow execution.
	Status DataIntegrationFlowExecutionStatus

	noSmithyDocumentSerde
}

// The output metadata of the flow execution.
type DataIntegrationFlowExecutionOutputMetadata struct {

	// The S3 URI under which all diagnostic files (such as deduped records if any)
	// are stored.
	DiagnosticReportsRootS3URI *string

	noSmithyDocumentSerde
}

// The source information of a flow execution.
type DataIntegrationFlowExecutionSourceInfo struct {

	// The data integration flow execution source type.
	//
	// This member is required.
	SourceType DataIntegrationFlowSourceType

	// The source details of a flow execution with dataset source.
	DatasetSource *DataIntegrationFlowDatasetSource

	// The source details of a flow execution with S3 source.
	S3Source *DataIntegrationFlowS3Source

	noSmithyDocumentSerde
}

// The field used in the field priority deduplication strategy.
type DataIntegrationFlowFieldPriorityDedupeField struct {

	// The name of the deduplication field. Must exist in the dataset and not be a
	// primary key.
	//
	// This member is required.
	Name *string

	// The sort order for the deduplication field.
	//
	// This member is required.
	SortOrder DataIntegrationFlowFieldPriorityDedupeSortOrder

	noSmithyDocumentSerde
}

// The field priority deduplication strategy details.
type DataIntegrationFlowFieldPriorityDedupeStrategyConfiguration struct {

	// The list of field names and their sort order for deduplication, arranged in
	// descending priority from highest to lowest.
	//
	// This member is required.
	Fields []DataIntegrationFlowFieldPriorityDedupeField

	noSmithyDocumentSerde
}

// The Amazon S3 options used in S3 source and target configurations.
type DataIntegrationFlowS3Options struct {

	// The Amazon S3 file type in S3 options.
	FileType DataIntegrationFlowFileType

	noSmithyDocumentSerde
}

// The details of a flow execution with S3 source.
type DataIntegrationFlowS3Source struct {

	// The S3 bucket name of the S3 source.
	//
	// This member is required.
	BucketName *string

	// The S3 object key of the S3 source.
	//
	// This member is required.
	Key *string

	noSmithyDocumentSerde
}

// The S3 DataIntegrationFlow source configuration parameters.
type DataIntegrationFlowS3SourceConfiguration struct {

	// The bucketName of the S3 source objects.
	//
	// This member is required.
	BucketName *string

	// The prefix of the S3 source objects. To trigger data ingestion, S3 files need
	// to be put under s3://bucketName/prefix/ .
	//
	// This member is required.
	Prefix *string

	// The other options of the S3 DataIntegrationFlow source.
	Options *DataIntegrationFlowS3Options

	noSmithyDocumentSerde
}

// The S3 DataIntegrationFlow target configuration parameters.
type DataIntegrationFlowS3TargetConfiguration struct {

	// The bucketName of the S3 target objects.
	//
	// This member is required.
	BucketName *string

	// The prefix of the S3 target objects.
	//
	// This member is required.
	Prefix *string

	// The S3 DataIntegrationFlow target options.
	Options *DataIntegrationFlowS3Options

	noSmithyDocumentSerde
}

// The DataIntegrationFlow source parameters.
type DataIntegrationFlowSource struct {

	// The DataIntegrationFlow source name that can be used as table alias in SQL
	// transformation query.
	//
	// This member is required.
	SourceName *string

	// The DataIntegrationFlow source type.
	//
	// This member is required.
	SourceType DataIntegrationFlowSourceType

	// The dataset DataIntegrationFlow source.
	DatasetSource *DataIntegrationFlowDatasetSourceConfiguration

	// The S3 DataIntegrationFlow source.
	S3Source *DataIntegrationFlowS3SourceConfiguration

	noSmithyDocumentSerde
}

// The SQL DataIntegrationFlow transformation configuration parameters.
type DataIntegrationFlowSQLTransformationConfiguration struct {

	// The transformation SQL query body based on SparkSQL.
	//
	// This member is required.
	Query *string

	noSmithyDocumentSerde
}

// The DataIntegrationFlow target parameters.
type DataIntegrationFlowTarget struct {

	// The DataIntegrationFlow target type.
	//
	// This member is required.
	TargetType DataIntegrationFlowTargetType

	// The dataset DataIntegrationFlow target. Note that for AWS Supply Chain dataset
	// under asc namespace, it has a connection_id internal field that is not allowed
	// to be provided by client directly, they will be auto populated.
	DatasetTarget *DataIntegrationFlowDatasetTargetConfiguration

	// The S3 DataIntegrationFlow target.
	S3Target *DataIntegrationFlowS3TargetConfiguration

	noSmithyDocumentSerde
}

// The DataIntegrationFlow transformation parameters.
type DataIntegrationFlowTransformation struct {

	// The DataIntegrationFlow transformation type.
	//
	// This member is required.
	TransformationType DataIntegrationFlowTransformationType

	// The SQL DataIntegrationFlow transformation configuration.
	SqlTransformation *DataIntegrationFlowSQLTransformationConfiguration

	noSmithyDocumentSerde
}

// The data lake dataset details.
type DataLakeDataset struct {

	// The arn of the dataset.
	//
	// This member is required.
	Arn *string

	// The creation time of the dataset.
	//
	// This member is required.
	CreatedTime *time.Time

	// The Amazon Web Services Supply Chain instance identifier.
	//
	// This member is required.
	InstanceId *string

	// The last modified time of the dataset.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The name of the dataset. For asc namespace, the name must be one of the
	// supported data entities under [https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html].
	//
	// [https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html]: https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html
	//
	// This member is required.
	Name *string

	// The namespace of the dataset, besides the custom defined namespace, every
	// instance comes with below pre-defined namespaces:
	//
	//   - asc - For information on the Amazon Web Services Supply Chain supported
	//   datasets see [https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html].
	//
	//   - default - For datasets with custom user-defined schemas.
	//
	// [https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html]: https://docs.aws.amazon.com/aws-supply-chain/latest/userguide/data-model-asc.html
	//
	// This member is required.
	Namespace *string

	// The schema of the dataset.
	//
	// This member is required.
	Schema *DataLakeDatasetSchema

	// The description of the dataset.
	Description *string

	// The partition specification for a dataset.
	PartitionSpec *DataLakeDatasetPartitionSpec

	noSmithyDocumentSerde
}

// The detail of the partition field.
type DataLakeDatasetPartitionField struct {

	// The name of the partition field.
	//
	// This member is required.
	Name *string

	// The transformation of the partition field. A transformation specifies how to
	// partition on a given field. For example, with timestamp you can specify that
	// you'd like to partition fields by day, e.g. data record with value
	// 2025-01-03T00:00:00Z in partition field is in 2025-01-03 partition. Also noted
	// that data record without any value in optional partition field is in NULL
	// partition.
	//
	// This member is required.
	Transform *DataLakeDatasetPartitionFieldTransform

	noSmithyDocumentSerde
}

// The detail of the partition field transformation.
type DataLakeDatasetPartitionFieldTransform struct {

	// The type of partitioning transformation for this field. The available options
	// are:
	//
	//   - IDENTITY - Partitions data on a given field by its exact values.
	//
	//   - YEAR - Partitions data on a timestamp field using year granularity.
	//
	//   - MONTH - Partitions data on a timestamp field using month granularity.
	//
	//   - DAY - Partitions data on a timestamp field using day granularity.
	//
	//   - HOUR - Partitions data on a timestamp field using hour granularity.
	//
	// This member is required.
	Type DataLakeDatasetPartitionTransformType

	noSmithyDocumentSerde
}

// The partition specification for a dataset.
type DataLakeDatasetPartitionSpec struct {

	// The fields on which to partition a dataset. The partitions will be applied
	// hierarchically based on the order of this list.
	//
	// This member is required.
	Fields []DataLakeDatasetPartitionField

	noSmithyDocumentSerde
}

// The detail of the primary key field.
type DataLakeDatasetPrimaryKeyField struct {

	// The name of the primary key field.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// The schema details of the dataset. Note that for AWS Supply Chain dataset under
// asc namespace, it may have internal fields like connection_id that will be auto
// populated by data ingestion methods.
type DataLakeDatasetSchema struct {

	// The list of field details of the dataset schema.
	//
	// This member is required.
	Fields []DataLakeDatasetSchemaField

	// The name of the dataset schema.
	//
	// This member is required.
	Name *string

	// The list of primary key fields for the dataset. Primary keys defined can help
	// data ingestion methods to ensure data uniqueness: CreateDataIntegrationFlow's
	// dedupe strategy will leverage primary keys to perform records deduplication
	// before write to dataset; SendDataIntegrationEvent's UPSERT and DELETE can only
	// work with dataset with primary keys. For more details, refer to those data
	// ingestion documentations.
	//
	// Note that defining primary keys does not necessarily mean the dataset cannot
	// have duplicate records, duplicate records can still be ingested if
	// CreateDataIntegrationFlow's dedupe disabled or through
	// SendDataIntegrationEvent's APPEND operation.
	PrimaryKeys []DataLakeDatasetPrimaryKeyField

	noSmithyDocumentSerde
}

// The dataset field details.
type DataLakeDatasetSchemaField struct {

	// Indicate if the field is required or not.
	//
	// This member is required.
	IsRequired *bool

	// The dataset field name.
	//
	// This member is required.
	Name *string

	// The dataset field type.
	//
	// This member is required.
	Type DataLakeDatasetSchemaFieldType

	noSmithyDocumentSerde
}

// The data lake namespace details.
type DataLakeNamespace struct {

	// The arn of the namespace.
	//
	// This member is required.
	Arn *string

	// The creation time of the namespace.
	//
	// This member is required.
	CreatedTime *time.Time

	// The Amazon Web Services Supply Chain instance identifier.
	//
	// This member is required.
	InstanceId *string

	// The last modified time of the namespace.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The name of the namespace.
	//
	// This member is required.
	Name *string

	// The description of the namespace.
	Description *string

	noSmithyDocumentSerde
}

// The details of the instance.
type Instance struct {

	// The Amazon Web Services account ID that owns the instance.
	//
	// This member is required.
	AwsAccountId *string

	// The Amazon Web Services Supply Chain instance identifier.
	//
	// This member is required.
	InstanceId *string

	// The state of the instance.
	//
	// This member is required.
	State InstanceState

	// The instance creation timestamp.
	CreatedTime *time.Time

	// The Amazon Web Services Supply Chain instance error message. If the instance
	// results in an unhealthy state, customers need to check the error message, delete
	// the current instance, and recreate a new one based on the mitigation from the
	// error message.
	ErrorMessage *string

	// The Amazon Web Services Supply Chain instance description.
	InstanceDescription *string

	// The Amazon Web Services Supply Chain instance name.
	InstanceName *string

	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you
	// optionally provided for encryption. If you did not provide anything here, AWS
	// Supply Chain uses the Amazon Web Services owned KMS key and nothing is returned.
	KmsKeyArn *string

	// The instance last modified timestamp.
	LastModifiedTime *time.Time

	// The version number of the instance.
	VersionNumber *float64

	// The WebApp DNS domain name of the instance.
	WebAppDnsDomain *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
