/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ComponentsObservation struct {
	Component *string `json:"component,omitempty" tf:"component,omitempty"`

	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	KafkaAuthenticationMethod *string `json:"kafkaAuthenticationMethod,omitempty" tf:"kafka_authentication_method,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	Route *string `json:"route,omitempty" tf:"route,omitempty"`

	SSL *bool `json:"ssl,omitempty" tf:"ssl,omitempty"`

	Usage *string `json:"usage,omitempty" tf:"usage,omitempty"`
}

type ComponentsParameters struct {
}

type MigrationObservation struct {
}

type MigrationParameters struct {

	// Database name for bootstrapping the initial connection
	// +kubebuilder:validation:Optional
	Dbname *string `json:"dbname,omitempty" tf:"dbname,omitempty"`

	// Hostname or IP address of the server where to migrate data from
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// Comma-separated list of databases, which should be ignored during migration (supported by MySQL only at the moment)
	// +kubebuilder:validation:Optional
	IgnoreDbs *string `json:"ignoreDbs,omitempty" tf:"ignore_dbs,omitempty"`

	// The migration method to be used (currently supported only by Redis and MySQL service types)
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// Password for authentication with the server where to migrate data from
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// Port number of the server where to migrate data from
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The server where to migrate data from is secured with SSL
	// +kubebuilder:validation:Optional
	SSL *string `json:"ssl,omitempty" tf:"ssl,omitempty"`

	// User name for authentication with the server where to migrate data from
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type PgObservation struct {

	// Service component information objects
	Components []ComponentsObservation `json:"components,omitempty" tf:"components,omitempty"`

	// The maximum disk space of the service, possible values depend on the service type, the cloud provider and the project.
	DiskSpaceCap *string `json:"diskSpaceCap,omitempty" tf:"disk_space_cap,omitempty"`

	// The default disk space of the service, possible values depend on the service type, the cloud provider and the project. Its also the minimum value for `disk_space`
	DiskSpaceDefault *string `json:"diskSpaceDefault,omitempty" tf:"disk_space_default,omitempty"`

	// The default disk space step of the service, possible values depend on the service type, the cloud provider and the project. `disk_space` needs to increment from `disk_space_default` by increments of this size.
	DiskSpaceStep *string `json:"diskSpaceStep,omitempty" tf:"disk_space_step,omitempty"`

	// Disk space that service is currently using
	DiskSpaceUsed *string `json:"diskSpaceUsed,omitempty" tf:"disk_space_used,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// PostgreSQL specific server provided values
	// +kubebuilder:validation:Optional
	Pg []PgPgObservation `json:"pg,omitempty" tf:"pg,omitempty"`

	// The hostname of the service.
	ServiceHost *string `json:"serviceHost,omitempty" tf:"service_host,omitempty"`

	// The port of the service
	ServicePort *float64 `json:"servicePort,omitempty" tf:"service_port,omitempty"`

	// Aiven internal service type code
	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`

	// Username used for connecting to the service, if applicable
	ServiceUsername *string `json:"serviceUsername,omitempty" tf:"service_username,omitempty"`

	// Service state. One of `POWEROFF`, `REBALANCING`, `REBUILDING` or `RUNNING`
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type PgParameters struct {

	// Additional disk space. Possible values depend on the service type, the cloud provider and the project. Therefore, reducing will result in the service rebalancing.
	// +kubebuilder:validation:Optional
	AdditionalDiskSpace *string `json:"additionalDiskSpace,omitempty" tf:"additional_disk_space,omitempty"`

	// Defines where the cloud provider and region where the service is hosted in. This can be changed freely after service is created. Changing the value will trigger a potentially lengthy migration process for the service. Format is cloud provider name (`aws`, `azure`, `do` `google`, `upcloud`, etc.), dash, and the cloud provider specific region name. These are documented on each Cloud provider's own support articles, like [here for Google](https://cloud.google.com/compute/docs/regions-zones/) and [here for AWS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).
	// +kubebuilder:validation:Optional
	CloudName *string `json:"cloudName,omitempty" tf:"cloud_name,omitempty"`

	// Service disk space. Possible values depend on the service type, the cloud provider and the project. Therefore, reducing will result in the service rebalancing.
	// +kubebuilder:validation:Optional
	DiskSpace *string `json:"diskSpace,omitempty" tf:"disk_space,omitempty"`

	// Day of week when maintenance operations should be performed. One monday, tuesday, wednesday, etc.
	// +kubebuilder:validation:Optional
	MaintenanceWindowDow *string `json:"maintenanceWindowDow,omitempty" tf:"maintenance_window_dow,omitempty"`

	// Time of day when maintenance operations should be performed. UTC time in HH:mm:ss format.
	// +kubebuilder:validation:Optional
	MaintenanceWindowTime *string `json:"maintenanceWindowTime,omitempty" tf:"maintenance_window_time,omitempty"`

	// PostgreSQL specific server provided values
	// +kubebuilder:validation:Optional
	Pg []PgPgParameters `json:"pg,omitempty" tf:"pg,omitempty"`

	// Pg user configurable settings
	// +kubebuilder:validation:Optional
	PgUserConfig []PgUserConfigParameters `json:"pgUserConfig,omitempty" tf:"pg_user_config,omitempty"`

	// Defines what kind of computing resources are allocated for the service. It can be changed after creation, though there are some restrictions when going to a smaller plan such as the new plan must have sufficient amount of disk space to store all current data and switching to a plan with fewer nodes might not be supported. The basic plan names are `hobbyist`, `startup-x`, `business-x` and `premium-x` where `x` is (roughly) the amount of memory on each node (also other attributes like number of CPUs and amount of disk space varies but naming is based on memory). The available options can be seem from the [Aiven pricing page](https://aiven.io/pricing).
	// +kubebuilder:validation:Optional
	Plan *string `json:"plan,omitempty" tf:"plan,omitempty"`

	// Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// Specifies the VPC the service should run in. If the value is not set the service is not run inside a VPC. When set, the value should be given as a reference to set up dependencies correctly and the VPC must be in the same cloud and region as the service itself. Project can be freely moved to and from VPC after creation but doing so triggers migration to new servers so the operation can take significant amount of time to complete if the service has a lot of data.
	// +kubebuilder:validation:Optional
	ProjectVPCID *string `json:"projectVpcId,omitempty" tf:"project_vpc_id,omitempty"`

	// Service integrations to specify when creating a service. Not applied after initial service creation
	// +kubebuilder:validation:Optional
	ServiceIntegrations []ServiceIntegrationsParameters `json:"serviceIntegrations,omitempty" tf:"service_integrations,omitempty"`

	// Specifies the actual name of the service. The name cannot be changed later without destroying and re-creating the service so name should be picked based on intended service usage rather than current attributes.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// Static IPs that are going to be associated with this service. Please assign a value using the 'toset' function. Once a static ip resource is in the 'assigned' state it cannot be unbound from the node again
	// +kubebuilder:validation:Optional
	StaticIps []*string `json:"staticIps,omitempty" tf:"static_ips,omitempty"`

	// Tags are key-value pairs that allow you to categorize services.
	// +kubebuilder:validation:Optional
	Tag []TagParameters `json:"tag,omitempty" tf:"tag,omitempty"`

	// Prevents the service from being deleted. It is recommended to set this to `true` for all production services to prevent unintentional service deletion. This does not shield against deleting databases or topics but for services with backups much of the content can at least be restored from backup in case accidental deletion is done.
	// +kubebuilder:validation:Optional
	TerminationProtection *bool `json:"terminationProtection,omitempty" tf:"termination_protection,omitempty"`
}

type PgPgObservation struct {

	// Primary PostgreSQL database name
	Dbname *string `json:"dbname,omitempty" tf:"dbname,omitempty"`

	// PostgreSQL master node host IP or name
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// Connection limit
	MaxConnections *float64 `json:"maxConnections,omitempty" tf:"max_connections,omitempty"`

	// PostgreSQL port
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// PostgreSQL sslmode setting (currently always "require")
	Sslmode *string `json:"sslmode,omitempty" tf:"sslmode,omitempty"`

	// PostgreSQL admin user name
	User *string `json:"user,omitempty" tf:"user,omitempty"`
}

type PgPgParameters struct {

	// PostgreSQL master connection URI
	// +kubebuilder:validation:Optional
	URISecretRef *v1.SecretKeySelector `json:"uriSecretRef,omitempty" tf:"-"`
}

type PgUserConfigObservation struct {
}

type PgUserConfigParameters struct {

	// Additional Cloud Regions for Backup Replication
	// +kubebuilder:validation:Optional
	AdditionalBackupRegions []*string `json:"additionalBackupRegions,omitempty" tf:"additional_backup_regions,omitempty"`

	// Custom password for admin user. Defaults to random string. This must be set only when a new service is being created.
	// +kubebuilder:validation:Optional
	AdminPasswordSecretRef *v1.SecretKeySelector `json:"adminPasswordSecretRef,omitempty" tf:"-"`

	// Custom username for admin user. This must be set only when a new service is being created.
	// +kubebuilder:validation:Optional
	AdminUsername *string `json:"adminUsername,omitempty" tf:"admin_username,omitempty"`

	// The hour of day (in UTC) when backup for the service is started. New backup is only started if previous backup has already completed.
	// +kubebuilder:validation:Optional
	BackupHour *string `json:"backupHour,omitempty" tf:"backup_hour,omitempty"`

	// The minute of an hour when backup for the service is started. New backup is only started if previous backup has already completed.
	// +kubebuilder:validation:Optional
	BackupMinute *string `json:"backupMinute,omitempty" tf:"backup_minute,omitempty"`

	// Enable IPv6
	// +kubebuilder:validation:Optional
	EnableIPv6 *string `json:"enableIpv6,omitempty" tf:"enable_ipv6,omitempty"`

	// IP filter
	// +kubebuilder:validation:Optional
	IPFilter []*string `json:"ipFilter,omitempty" tf:"ip_filter,omitempty"`

	// Migrate data from existing server
	// +kubebuilder:validation:Optional
	Migration []MigrationParameters `json:"migration,omitempty" tf:"migration,omitempty"`

	// postgresql.conf configuration values
	// +kubebuilder:validation:Optional
	Pg []PgUserConfigPgParameters `json:"pg,omitempty" tf:"pg,omitempty"`

	// Should the service which is being forked be a read replica (deprecated, use read_replica service integration instead).
	// +kubebuilder:validation:Optional
	PgReadReplica *string `json:"pgReadReplica,omitempty" tf:"pg_read_replica,omitempty"`

	// Name of the PG Service from which to fork (deprecated, use service_to_fork_from). This has effect only when a new service is being created.
	// +kubebuilder:validation:Optional
	PgServiceToForkFrom *string `json:"pgServiceToForkFrom,omitempty" tf:"pg_service_to_fork_from,omitempty"`

	// Enable pg_stat_monitor extension if available for the current cluster
	// +kubebuilder:validation:Optional
	PgStatMonitorEnable *string `json:"pgStatMonitorEnable,omitempty" tf:"pg_stat_monitor_enable,omitempty"`

	// PostgreSQL major version
	// +kubebuilder:validation:Optional
	PgVersion *string `json:"pgVersion,omitempty" tf:"pg_version,omitempty"`

	// PGBouncer connection pooling settings
	// +kubebuilder:validation:Optional
	Pgbouncer []PgbouncerParameters `json:"pgbouncer,omitempty" tf:"pgbouncer,omitempty"`

	// PGLookout settings
	// +kubebuilder:validation:Optional
	Pglookout []PglookoutParameters `json:"pglookout,omitempty" tf:"pglookout,omitempty"`

	// Allow access to selected service ports from private networks
	// +kubebuilder:validation:Optional
	PrivateAccess []PrivateAccessParameters `json:"privateAccess,omitempty" tf:"private_access,omitempty"`

	// Allow access to selected service components through Privatelink
	// +kubebuilder:validation:Optional
	PrivatelinkAccess []PrivatelinkAccessParameters `json:"privatelinkAccess,omitempty" tf:"privatelink_access,omitempty"`

	// Name of another project to fork a service from. This has effect only when a new service is being created.
	// +kubebuilder:validation:Optional
	ProjectToForkFrom *string `json:"projectToForkFrom,omitempty" tf:"project_to_fork_from,omitempty"`

	// Allow access to selected service ports from the public Internet
	// +kubebuilder:validation:Optional
	PublicAccess []PublicAccessParameters `json:"publicAccess,omitempty" tf:"public_access,omitempty"`

	// Recovery target time when forking a service. This has effect only when a new service is being created.
	// +kubebuilder:validation:Optional
	RecoveryTargetTime *string `json:"recoveryTargetTime,omitempty" tf:"recovery_target_time,omitempty"`

	// Name of another service to fork from. This has effect only when a new service is being created.
	// +kubebuilder:validation:Optional
	ServiceToForkFrom *string `json:"serviceToForkFrom,omitempty" tf:"service_to_fork_from,omitempty"`

	// shared_buffers_percentage
	// +kubebuilder:validation:Optional
	SharedBuffersPercentage *string `json:"sharedBuffersPercentage,omitempty" tf:"shared_buffers_percentage,omitempty"`

	// Static IP addresses
	// +kubebuilder:validation:Optional
	StaticIps *string `json:"staticIps,omitempty" tf:"static_ips,omitempty"`

	// Synchronous replication type. Note that the service plan also needs to support synchronous replication.
	// +kubebuilder:validation:Optional
	SynchronousReplication *string `json:"synchronousReplication,omitempty" tf:"synchronous_replication,omitempty"`

	// TimescaleDB extension configuration values
	// +kubebuilder:validation:Optional
	Timescaledb []TimescaledbParameters `json:"timescaledb,omitempty" tf:"timescaledb,omitempty"`

	// Variant of the PostgreSQL service, may affect the features that are exposed by default
	// +kubebuilder:validation:Optional
	Variant *string `json:"variant,omitempty" tf:"variant,omitempty"`

	// work_mem
	// +kubebuilder:validation:Optional
	WorkMem *string `json:"workMem,omitempty" tf:"work_mem,omitempty"`
}

type PgUserConfigPgObservation struct {
}

type PgUserConfigPgParameters struct {

	// autovacuum_analyze_scale_factor
	// +kubebuilder:validation:Optional
	AutovacuumAnalyzeScaleFactor *string `json:"autovacuumAnalyzeScaleFactor,omitempty" tf:"autovacuum_analyze_scale_factor,omitempty"`

	// autovacuum_analyze_threshold
	// +kubebuilder:validation:Optional
	AutovacuumAnalyzeThreshold *string `json:"autovacuumAnalyzeThreshold,omitempty" tf:"autovacuum_analyze_threshold,omitempty"`

	// autovacuum_freeze_max_age
	// +kubebuilder:validation:Optional
	AutovacuumFreezeMaxAge *string `json:"autovacuumFreezeMaxAge,omitempty" tf:"autovacuum_freeze_max_age,omitempty"`

	// autovacuum_max_workers
	// +kubebuilder:validation:Optional
	AutovacuumMaxWorkers *string `json:"autovacuumMaxWorkers,omitempty" tf:"autovacuum_max_workers,omitempty"`

	// autovacuum_naptime
	// +kubebuilder:validation:Optional
	AutovacuumNaptime *string `json:"autovacuumNaptime,omitempty" tf:"autovacuum_naptime,omitempty"`

	// autovacuum_vacuum_cost_delay
	// +kubebuilder:validation:Optional
	AutovacuumVacuumCostDelay *string `json:"autovacuumVacuumCostDelay,omitempty" tf:"autovacuum_vacuum_cost_delay,omitempty"`

	// autovacuum_vacuum_cost_limit
	// +kubebuilder:validation:Optional
	AutovacuumVacuumCostLimit *string `json:"autovacuumVacuumCostLimit,omitempty" tf:"autovacuum_vacuum_cost_limit,omitempty"`

	// autovacuum_vacuum_scale_factor
	// +kubebuilder:validation:Optional
	AutovacuumVacuumScaleFactor *string `json:"autovacuumVacuumScaleFactor,omitempty" tf:"autovacuum_vacuum_scale_factor,omitempty"`

	// autovacuum_vacuum_threshold
	// +kubebuilder:validation:Optional
	AutovacuumVacuumThreshold *string `json:"autovacuumVacuumThreshold,omitempty" tf:"autovacuum_vacuum_threshold,omitempty"`

	// bgwriter_delay
	// +kubebuilder:validation:Optional
	BgwriterDelay *string `json:"bgwriterDelay,omitempty" tf:"bgwriter_delay,omitempty"`

	// bgwriter_flush_after
	// +kubebuilder:validation:Optional
	BgwriterFlushAfter *string `json:"bgwriterFlushAfter,omitempty" tf:"bgwriter_flush_after,omitempty"`

	// bgwriter_lru_maxpages
	// +kubebuilder:validation:Optional
	BgwriterLruMaxpages *string `json:"bgwriterLruMaxpages,omitempty" tf:"bgwriter_lru_maxpages,omitempty"`

	// bgwriter_lru_multiplier
	// +kubebuilder:validation:Optional
	BgwriterLruMultiplier *string `json:"bgwriterLruMultiplier,omitempty" tf:"bgwriter_lru_multiplier,omitempty"`

	// deadlock_timeout
	// +kubebuilder:validation:Optional
	DeadlockTimeout *string `json:"deadlockTimeout,omitempty" tf:"deadlock_timeout,omitempty"`

	// default_toast_compression
	// +kubebuilder:validation:Optional
	DefaultToastCompression *string `json:"defaultToastCompression,omitempty" tf:"default_toast_compression,omitempty"`

	// idle_in_transaction_session_timeout
	// +kubebuilder:validation:Optional
	IdleInTransactionSessionTimeout *string `json:"idleInTransactionSessionTimeout,omitempty" tf:"idle_in_transaction_session_timeout,omitempty"`

	// jit
	// +kubebuilder:validation:Optional
	Jit *string `json:"jit,omitempty" tf:"jit,omitempty"`

	// log_autovacuum_min_duration
	// +kubebuilder:validation:Optional
	LogAutovacuumMinDuration *string `json:"logAutovacuumMinDuration,omitempty" tf:"log_autovacuum_min_duration,omitempty"`

	// log_error_verbosity
	// +kubebuilder:validation:Optional
	LogErrorVerbosity *string `json:"logErrorVerbosity,omitempty" tf:"log_error_verbosity,omitempty"`

	// log_line_prefix
	// +kubebuilder:validation:Optional
	LogLinePrefix *string `json:"logLinePrefix,omitempty" tf:"log_line_prefix,omitempty"`

	// log_min_duration_statement
	// +kubebuilder:validation:Optional
	LogMinDurationStatement *string `json:"logMinDurationStatement,omitempty" tf:"log_min_duration_statement,omitempty"`

	// log_temp_files
	// +kubebuilder:validation:Optional
	LogTempFiles *string `json:"logTempFiles,omitempty" tf:"log_temp_files,omitempty"`

	// max_files_per_process
	// +kubebuilder:validation:Optional
	MaxFilesPerProcess *string `json:"maxFilesPerProcess,omitempty" tf:"max_files_per_process,omitempty"`

	// max_locks_per_transaction
	// +kubebuilder:validation:Optional
	MaxLocksPerTransaction *string `json:"maxLocksPerTransaction,omitempty" tf:"max_locks_per_transaction,omitempty"`

	// max_logical_replication_workers
	// +kubebuilder:validation:Optional
	MaxLogicalReplicationWorkers *string `json:"maxLogicalReplicationWorkers,omitempty" tf:"max_logical_replication_workers,omitempty"`

	// max_parallel_workers
	// +kubebuilder:validation:Optional
	MaxParallelWorkers *string `json:"maxParallelWorkers,omitempty" tf:"max_parallel_workers,omitempty"`

	// max_parallel_workers_per_gather
	// +kubebuilder:validation:Optional
	MaxParallelWorkersPerGather *string `json:"maxParallelWorkersPerGather,omitempty" tf:"max_parallel_workers_per_gather,omitempty"`

	// max_pred_locks_per_transaction
	// +kubebuilder:validation:Optional
	MaxPredLocksPerTransaction *string `json:"maxPredLocksPerTransaction,omitempty" tf:"max_pred_locks_per_transaction,omitempty"`

	// max_prepared_transactions
	// +kubebuilder:validation:Optional
	MaxPreparedTransactions *string `json:"maxPreparedTransactions,omitempty" tf:"max_prepared_transactions,omitempty"`

	// max_replication_slots
	// +kubebuilder:validation:Optional
	MaxReplicationSlots *string `json:"maxReplicationSlots,omitempty" tf:"max_replication_slots,omitempty"`

	// max_slot_wal_keep_size
	// +kubebuilder:validation:Optional
	MaxSlotWalKeepSize *string `json:"maxSlotWalKeepSize,omitempty" tf:"max_slot_wal_keep_size,omitempty"`

	// max_stack_depth
	// +kubebuilder:validation:Optional
	MaxStackDepth *string `json:"maxStackDepth,omitempty" tf:"max_stack_depth,omitempty"`

	// max_standby_archive_delay
	// +kubebuilder:validation:Optional
	MaxStandbyArchiveDelay *string `json:"maxStandbyArchiveDelay,omitempty" tf:"max_standby_archive_delay,omitempty"`

	// max_standby_streaming_delay
	// +kubebuilder:validation:Optional
	MaxStandbyStreamingDelay *string `json:"maxStandbyStreamingDelay,omitempty" tf:"max_standby_streaming_delay,omitempty"`

	// max_wal_senders
	// +kubebuilder:validation:Optional
	MaxWalSenders *string `json:"maxWalSenders,omitempty" tf:"max_wal_senders,omitempty"`

	// max_worker_processes
	// +kubebuilder:validation:Optional
	MaxWorkerProcesses *string `json:"maxWorkerProcesses,omitempty" tf:"max_worker_processes,omitempty"`

	// pg_partman_bgw.interval
	// +kubebuilder:validation:Optional
	PgPartmanBgwDotInterval *string `json:"pgPartmanBgwDotInterval,omitempty" tf:"pg_partman_bgw__dot__interval,omitempty"`

	// pg_partman_bgw.role
	// +kubebuilder:validation:Optional
	PgPartmanBgwDotRole *string `json:"pgPartmanBgwDotRole,omitempty" tf:"pg_partman_bgw__dot__role,omitempty"`

	// pg_stat_statements.track
	// +kubebuilder:validation:Optional
	PgStatStatementsDotTrack *string `json:"pgStatStatementsDotTrack,omitempty" tf:"pg_stat_statements__dot__track,omitempty"`

	// temp_file_limit
	// +kubebuilder:validation:Optional
	TempFileLimit *string `json:"tempFileLimit,omitempty" tf:"temp_file_limit,omitempty"`

	// timezone
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// track_activity_query_size
	// +kubebuilder:validation:Optional
	TrackActivityQuerySize *string `json:"trackActivityQuerySize,omitempty" tf:"track_activity_query_size,omitempty"`

	// track_commit_timestamp
	// +kubebuilder:validation:Optional
	TrackCommitTimestamp *string `json:"trackCommitTimestamp,omitempty" tf:"track_commit_timestamp,omitempty"`

	// track_functions
	// +kubebuilder:validation:Optional
	TrackFunctions *string `json:"trackFunctions,omitempty" tf:"track_functions,omitempty"`

	// track_io_timing
	// +kubebuilder:validation:Optional
	TrackIoTiming *string `json:"trackIoTiming,omitempty" tf:"track_io_timing,omitempty"`

	// wal_sender_timeout
	// +kubebuilder:validation:Optional
	WalSenderTimeout *string `json:"walSenderTimeout,omitempty" tf:"wal_sender_timeout,omitempty"`

	// wal_writer_delay
	// +kubebuilder:validation:Optional
	WalWriterDelay *string `json:"walWriterDelay,omitempty" tf:"wal_writer_delay,omitempty"`
}

type PgbouncerObservation struct {
}

type PgbouncerParameters struct {

	// If the automatically created database pools have been unused this many seconds, they are freed. If 0 then timeout is disabled. [seconds]
	// +kubebuilder:validation:Optional
	AutodbIdleTimeout *string `json:"autodbIdleTimeout,omitempty" tf:"autodb_idle_timeout,omitempty"`

	// Do not allow more than this many server connections per database (regardless of user). Setting it to 0 means unlimited.
	// +kubebuilder:validation:Optional
	AutodbMaxDBConnections *string `json:"autodbMaxDbConnections,omitempty" tf:"autodb_max_db_connections,omitempty"`

	// PGBouncer pool mode
	// +kubebuilder:validation:Optional
	AutodbPoolMode *string `json:"autodbPoolMode,omitempty" tf:"autodb_pool_mode,omitempty"`

	// If non-zero then create automatically a pool of that size per user when a pool doesn't exist.
	// +kubebuilder:validation:Optional
	AutodbPoolSize *string `json:"autodbPoolSize,omitempty" tf:"autodb_pool_size,omitempty"`

	// List of parameters to ignore when given in startup packet
	// +kubebuilder:validation:Optional
	IgnoreStartupParameters []*string `json:"ignoreStartupParameters,omitempty" tf:"ignore_startup_parameters,omitempty"`

	// Add more server connections to pool if below this number. Improves behavior when usual load comes suddenly back after period of total inactivity. The value is effectively capped at the pool size.
	// +kubebuilder:validation:Optional
	MinPoolSize *string `json:"minPoolSize,omitempty" tf:"min_pool_size,omitempty"`

	// If a server connection has been idle more than this many seconds it will be dropped. If 0 then timeout is disabled. [seconds]
	// +kubebuilder:validation:Optional
	ServerIdleTimeout *string `json:"serverIdleTimeout,omitempty" tf:"server_idle_timeout,omitempty"`

	// The pooler will close an unused server connection that has been connected longer than this. [seconds]
	// +kubebuilder:validation:Optional
	ServerLifetime *string `json:"serverLifetime,omitempty" tf:"server_lifetime,omitempty"`

	// Run server_reset_query (DISCARD ALL) in all pooling modes
	// +kubebuilder:validation:Optional
	ServerResetQueryAlways *string `json:"serverResetQueryAlways,omitempty" tf:"server_reset_query_always,omitempty"`
}

type PglookoutObservation struct {
}

type PglookoutParameters struct {

	// max_failover_replication_time_lag
	// +kubebuilder:validation:Optional
	MaxFailoverReplicationTimeLag *string `json:"maxFailoverReplicationTimeLag,omitempty" tf:"max_failover_replication_time_lag,omitempty"`
}

type PrivateAccessObservation struct {
}

type PrivateAccessParameters struct {

	// Allow clients to connect to pg with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	// +kubebuilder:validation:Optional
	Pg *string `json:"pg,omitempty" tf:"pg,omitempty"`

	// Allow clients to connect to pgbouncer with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	// +kubebuilder:validation:Optional
	Pgbouncer *string `json:"pgbouncer,omitempty" tf:"pgbouncer,omitempty"`

	// Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	// +kubebuilder:validation:Optional
	Prometheus *string `json:"prometheus,omitempty" tf:"prometheus,omitempty"`
}

type PrivatelinkAccessObservation struct {
}

type PrivatelinkAccessParameters struct {

	// Enable pg
	// +kubebuilder:validation:Optional
	Pg *string `json:"pg,omitempty" tf:"pg,omitempty"`

	// Enable pgbouncer
	// +kubebuilder:validation:Optional
	Pgbouncer *string `json:"pgbouncer,omitempty" tf:"pgbouncer,omitempty"`

	// Enable prometheus
	// +kubebuilder:validation:Optional
	Prometheus *string `json:"prometheus,omitempty" tf:"prometheus,omitempty"`
}

type PublicAccessObservation struct {
}

type PublicAccessParameters struct {

	// Allow clients to connect to pg from the public internet for service nodes that are in a project VPC or another type of private network
	// +kubebuilder:validation:Optional
	Pg *string `json:"pg,omitempty" tf:"pg,omitempty"`

	// Allow clients to connect to pgbouncer from the public internet for service nodes that are in a project VPC or another type of private network
	// +kubebuilder:validation:Optional
	Pgbouncer *string `json:"pgbouncer,omitempty" tf:"pgbouncer,omitempty"`

	// Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network
	// +kubebuilder:validation:Optional
	Prometheus *string `json:"prometheus,omitempty" tf:"prometheus,omitempty"`
}

type ServiceIntegrationsObservation struct {
}

type ServiceIntegrationsParameters struct {

	// Type of the service integration. The only supported value at the moment is `read_replica`
	// +kubebuilder:validation:Required
	IntegrationType *string `json:"integrationType" tf:"integration_type,omitempty"`

	// Name of the source service
	// +kubebuilder:validation:Required
	SourceServiceName *string `json:"sourceServiceName" tf:"source_service_name,omitempty"`
}

type TagObservation struct {
}

type TagParameters struct {

	// Service tag key
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Service tag value
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type TimescaledbObservation struct {
}

type TimescaledbParameters struct {

	// timescaledb.max_background_workers
	// +kubebuilder:validation:Optional
	MaxBackgroundWorkers *string `json:"maxBackgroundWorkers,omitempty" tf:"max_background_workers,omitempty"`
}

// PgSpec defines the desired state of Pg
type PgSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PgParameters `json:"forProvider"`
}

// PgStatus defines the observed state of Pg.
type PgStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PgObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Pg is the Schema for the Pgs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aiven}
type Pg struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PgSpec   `json:"spec"`
	Status            PgStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PgList contains a list of Pgs
type PgList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pg `json:"items"`
}

// Repository type metadata.
var (
	Pg_Kind             = "Pg"
	Pg_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Pg_Kind}.String()
	Pg_KindAPIVersion   = Pg_Kind + "." + CRDGroupVersion.String()
	Pg_GroupVersionKind = CRDGroupVersion.WithKind(Pg_Kind)
)

func init() {
	SchemeBuilder.Register(&Pg{}, &PgList{})
}
