package operations

import (
	"encoding/json"
	"fmt"

	"github.com/infra-rdc/rudder-go/internal/utils"
	"github.com/infra-rdc/rudder-go/models/components"
)

// GetAllSettingsResult - Result of the request
type GetAllSettingsResult string

const (
	GetAllSettingsResultSuccess GetAllSettingsResult = "success"
	GetAllSettingsResultError   GetAllSettingsResult = "error"
)

func (e GetAllSettingsResult) ToPointer() *GetAllSettingsResult {
	return &e
}
func (e *GetAllSettingsResult) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetAllSettingsResult(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllSettingsResult: %v", v)
	}
}

// GetAllSettingsAction - The id of the action
type GetAllSettingsAction string

const (
	GetAllSettingsActionGetAllSettings GetAllSettingsAction = "getAllSettings"
)

func (e GetAllSettingsAction) ToPointer() *GetAllSettingsAction {
	return &e
}
func (e *GetAllSettingsAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getAllSettings":
		*e = GetAllSettingsAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetAllSettingsAction: %v", v)
	}
}

type GetAllSettingsAllowedNetworks struct {
	// Rudder id of the policy server
	ID *string `json:"id,omitempty"`
	// List of allowed networks
	AllowedNetworks []string `json:"allowed_networks,omitempty"`
}

func (o *GetAllSettingsAllowedNetworks) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetAllSettingsAllowedNetworks) GetAllowedNetworks() []string {
	if o == nil {
		return nil
	}
	return o.AllowedNetworks
}

// GlobalPolicyMode - Define the default setting for global policy mode
type GlobalPolicyMode string

const (
	GlobalPolicyModeEnforce GlobalPolicyMode = "enforce"
	GlobalPolicyModeAudit   GlobalPolicyMode = "audit"
)

func (e GlobalPolicyMode) ToPointer() *GlobalPolicyMode {
	return &e
}
func (e *GlobalPolicyMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enforce":
		fallthrough
	case "audit":
		*e = GlobalPolicyMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GlobalPolicyMode: %v", v)
	}
}

// RelayServerSynchronizationMethod - Method used to synchronize data between server and relays, either "classic" (agent protocol, default), "rsync" (use rsync to synchronize, that you'll need to be manually set up), or "disabled" (use third party system to transmit data)
type RelayServerSynchronizationMethod string

const (
	RelayServerSynchronizationMethodClassic  RelayServerSynchronizationMethod = "classic"
	RelayServerSynchronizationMethodRsync    RelayServerSynchronizationMethod = "rsync"
	RelayServerSynchronizationMethodDisabled RelayServerSynchronizationMethod = "disabled"
)

func (e RelayServerSynchronizationMethod) ToPointer() *RelayServerSynchronizationMethod {
	return &e
}
func (e *RelayServerSynchronizationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "classic":
		fallthrough
	case "rsync":
		fallthrough
	case "disabled":
		*e = RelayServerSynchronizationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RelayServerSynchronizationMethod: %v", v)
	}
}

// RudderReportProtocolDefault - Default reporting protocol used
type RudderReportProtocolDefault string

const (
	RudderReportProtocolDefaultHTTPS  RudderReportProtocolDefault = "HTTPS"
	RudderReportProtocolDefaultSyslog RudderReportProtocolDefault = "SYSLOG"
)

func (e RudderReportProtocolDefault) ToPointer() *RudderReportProtocolDefault {
	return &e
}
func (e *RudderReportProtocolDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HTTPS":
		fallthrough
	case "SYSLOG":
		*e = RudderReportProtocolDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RudderReportProtocolDefault: %v", v)
	}
}

// ReportingMode - Compliance reporting mode
type ReportingMode string

const (
	ReportingModeFullCompliance  ReportingMode = "full-compliance"
	ReportingModeChangesOnly     ReportingMode = "changes-only"
	ReportingModeReportsDisabled ReportingMode = "reports-disabled"
)

func (e ReportingMode) ToPointer() *ReportingMode {
	return &e
}
func (e *ReportingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "full-compliance":
		fallthrough
	case "changes-only":
		fallthrough
	case "reports-disabled":
		*e = ReportingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ReportingMode: %v", v)
	}
}

// NodeOnacceptDefaultState - Set default state for node when they are accepted within Rudder
type NodeOnacceptDefaultState string

const (
	NodeOnacceptDefaultStateEnabled       NodeOnacceptDefaultState = "enabled"
	NodeOnacceptDefaultStateIgnored       NodeOnacceptDefaultState = "ignored"
	NodeOnacceptDefaultStateEmptyPolicies NodeOnacceptDefaultState = "empty-policies"
	NodeOnacceptDefaultStateInitializing  NodeOnacceptDefaultState = "initializing"
	NodeOnacceptDefaultStatePreparingEol  NodeOnacceptDefaultState = "preparing-eol"
)

func (e NodeOnacceptDefaultState) ToPointer() *NodeOnacceptDefaultState {
	return &e
}
func (e *NodeOnacceptDefaultState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enabled":
		fallthrough
	case "ignored":
		fallthrough
	case "empty-policies":
		fallthrough
	case "initializing":
		fallthrough
	case "preparing-eol":
		*e = NodeOnacceptDefaultState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeOnacceptDefaultState: %v", v)
	}
}

// NodeOnacceptDefaultPolicyMode - Default policy mode for accepted node
type NodeOnacceptDefaultPolicyMode string

const (
	NodeOnacceptDefaultPolicyModeDefault NodeOnacceptDefaultPolicyMode = "default"
	NodeOnacceptDefaultPolicyModeEnforce NodeOnacceptDefaultPolicyMode = "enforce"
	NodeOnacceptDefaultPolicyModeAudit   NodeOnacceptDefaultPolicyMode = "audit"
)

func (e NodeOnacceptDefaultPolicyMode) ToPointer() *NodeOnacceptDefaultPolicyMode {
	return &e
}
func (e *NodeOnacceptDefaultPolicyMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "default":
		fallthrough
	case "enforce":
		fallthrough
	case "audit":
		*e = NodeOnacceptDefaultPolicyMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NodeOnacceptDefaultPolicyMode: %v", v)
	}
}

type Settings struct {
	// List of allowed networks for each policy server (root and relays)
	AllowedNetworks []GetAllSettingsAllowedNetworks `json:"allowed_networks,omitempty"`
	// Define the default setting for global policy mode
	GlobalPolicyMode *GlobalPolicyMode `json:"global_policy_mode,omitempty"`
	// Allow overrides on this default setting
	GlobalPolicyModeOverridable *bool `json:"global_policy_mode_overridable,omitempty"`
	// Agent run schedule - time between agent runs (in minutes)
	RunFrequency *int64 `json:"run_frequency,omitempty"`
	// First agent run time - hour
	FirstRunHour *int64 `json:"first_run_hour,omitempty"`
	// First agent run time - minute
	FirstRunMinute *int64 `json:"first_run_minute,omitempty"`
	// Maximum delay after scheduled run time (random interval)
	SplayTime *int64 `json:"splay_time,omitempty"`
	// Number of days to retain modified files
	ModifiedFileTTL *int64 `json:"modified_file_ttl,omitempty"`
	// Number of days to retain agent output files
	OutputFileTTL *int64 `json:"output_file_ttl,omitempty"`
	// Require time synchronization between nodes and policy server
	RequireTimeSynchronization *bool `json:"require_time_synchronization,omitempty"`
	// Method used to synchronize data between server and relays, either "classic" (agent protocol, default), "rsync" (use rsync to synchronize, that you'll need to be manually set up), or "disabled" (use third party system to transmit data)
	RelayServerSynchronizationMethod *RelayServerSynchronizationMethod `json:"relay_server_synchronization_method,omitempty"`
	// If **rsync** is set as a synchronization method, use rsync to synchronize policies between Rudder server and relays. If false, you'll have to synchronize policies yourself.
	RelayServerSynchronizePolicies *bool `json:"relay_server_synchronize_policies,omitempty"`
	// If **rsync** is set as a synchronization method, use rsync to synchronize shared files between Rudder server and relays. If false, you'll have to synchronize shared files yourself.
	RelayServerSynchronizeSharedFiles *bool `json:"relay_server_synchronize_shared_files,omitempty"`
	// Default reporting protocol used
	RudderReportProtocolDefault *RudderReportProtocolDefault `json:"rudder_report_protocol_default,omitempty"`
	// Compliance reporting mode
	ReportingMode *ReportingMode `json:"reporting_mode,omitempty"`
	// Send heartbeat every heartbeat_frequency runs (only on **changes-only** compliance mode)
	HeartbeatFrequency *int64 `json:"heartbeat_frequency,omitempty"`
	// Enable change audit logs
	EnableChangeMessage *bool `json:"enable_change_message,omitempty"`
	// Make message mandatory
	MandatoryChangeMessage *bool `json:"mandatory_change_message,omitempty"`
	// Explanation to display
	ChangeMessagePrompt *string `json:"change_message_prompt,omitempty"`
	// Enable Change Requests
	EnableChangeRequest *bool `json:"enable_change_request,omitempty"`
	// Allow self validation
	EnableSelfValidation *bool `json:"enable_self_validation,omitempty"`
	// Allow self deployment
	EnableSelfDeployment *bool `json:"enable_self_deployment,omitempty"`
	// Display changes graphs
	DisplayRecentChangesGraphs *bool `json:"display_recent_changes_graphs,omitempty"`
	// Enable script evaluation in Directives
	EnableJavascriptDirectives *string `json:"enable_javascript_directives,omitempty"`
	// Send anonymous usage statistics
	SendMetrics *string `json:"send_metrics,omitempty"`
	// Allow acceptation of a pending node when another one with the same hostname is already accepted
	NodeAcceptDuplicatedHostname *bool `default:"false" json:"node_accept_duplicated_hostname"`
	// Set default state for node when they are accepted within Rudder
	NodeOnacceptDefaultState *NodeOnacceptDefaultState `json:"node_onaccept_default_state,omitempty"`
	// Default policy mode for accepted node
	NodeOnacceptDefaultPolicyMode *NodeOnacceptDefaultPolicyMode `json:"node_onaccept_default_policyMode,omitempty"`
	// Allows multiple reports for configuration based on a multivalued variable
	UnexpectedUnboundVarValues *bool `default:"true" json:"unexpected_unbound_var_values"`
	// Compute list of changes (repaired reports) per rule
	RudderComputeChanges *bool `default:"true" json:"rudder_compute_changes"`
	// Recompute all dynamic groups at the start of policy generation
	RudderGenerationComputeDyngroups *bool `default:"true" json:"rudder_generation_compute_dyngroups"`
	// Set the parallelism to compute dynamic group, as a number of thread (i.e. 4), or a multiplicative of the number of core (x0.5)
	RudderComputeDyngroupsMaxParallelism *string `default:"1" json:"rudder_compute_dyngroups_max_parallelism"`
	// Store all compliance levels in database
	RudderSaveDbComplianceLevels *bool `default:"true" json:"rudder_save_db_compliance_levels"`
	// Store all compliance details in database
	RudderSaveDbComplianceDetails *bool `default:"false" json:"rudder_save_db_compliance_details"`
	// Set the policy generation parallelism, either as an number of thread (i.e. 4), or a multiplicative of the number of core (x0.5)
	RudderGenerationMaxParallelism *string `default:"x0.5" json:"rudder_generation_max_parallelism"`
	// Policy generation JS evaluation of directive parameter timeout in seconds
	RudderGenerationJsTimeout *int64 `default:"30" json:"rudder_generation_js_timeout"`
	// Policy generation continues on error during NodeConfiguration evaluation
	RudderGenerationContinueOnError *bool `default:"false" json:"rudder_generation_continue_on_error"`
	// Set a delay before starting policy generation, this will allow you to accumulate changes before they are deployed to Nodes, and can also lessen webapp resources needs. The value is a number followed by the time unit needed (seconds/s, minutes/m, hours/h ...), ie "5m" for 5 minutes
	RudderGenerationDelay *string `default:"0 seconds" json:"rudder_generation_delay"`
	// Should policy generation be triggered automatically after a change (value 'all'), or should we wait until a manual launch (through api or UI, value 'onlyManual') or even no policy generation at all (value "none")
	RudderGenerationPolicy *string `default:"all" json:"rudder_generation_policy"`
}

func (s Settings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Settings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Settings) GetAllowedNetworks() []GetAllSettingsAllowedNetworks {
	if o == nil {
		return nil
	}
	return o.AllowedNetworks
}

func (o *Settings) GetGlobalPolicyMode() *GlobalPolicyMode {
	if o == nil {
		return nil
	}
	return o.GlobalPolicyMode
}

func (o *Settings) GetGlobalPolicyModeOverridable() *bool {
	if o == nil {
		return nil
	}
	return o.GlobalPolicyModeOverridable
}

func (o *Settings) GetRunFrequency() *int64 {
	if o == nil {
		return nil
	}
	return o.RunFrequency
}

func (o *Settings) GetFirstRunHour() *int64 {
	if o == nil {
		return nil
	}
	return o.FirstRunHour
}

func (o *Settings) GetFirstRunMinute() *int64 {
	if o == nil {
		return nil
	}
	return o.FirstRunMinute
}

func (o *Settings) GetSplayTime() *int64 {
	if o == nil {
		return nil
	}
	return o.SplayTime
}

func (o *Settings) GetModifiedFileTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.ModifiedFileTTL
}

func (o *Settings) GetOutputFileTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.OutputFileTTL
}

func (o *Settings) GetRequireTimeSynchronization() *bool {
	if o == nil {
		return nil
	}
	return o.RequireTimeSynchronization
}

func (o *Settings) GetRelayServerSynchronizationMethod() *RelayServerSynchronizationMethod {
	if o == nil {
		return nil
	}
	return o.RelayServerSynchronizationMethod
}

func (o *Settings) GetRelayServerSynchronizePolicies() *bool {
	if o == nil {
		return nil
	}
	return o.RelayServerSynchronizePolicies
}

func (o *Settings) GetRelayServerSynchronizeSharedFiles() *bool {
	if o == nil {
		return nil
	}
	return o.RelayServerSynchronizeSharedFiles
}

func (o *Settings) GetRudderReportProtocolDefault() *RudderReportProtocolDefault {
	if o == nil {
		return nil
	}
	return o.RudderReportProtocolDefault
}

func (o *Settings) GetReportingMode() *ReportingMode {
	if o == nil {
		return nil
	}
	return o.ReportingMode
}

func (o *Settings) GetHeartbeatFrequency() *int64 {
	if o == nil {
		return nil
	}
	return o.HeartbeatFrequency
}

func (o *Settings) GetEnableChangeMessage() *bool {
	if o == nil {
		return nil
	}
	return o.EnableChangeMessage
}

func (o *Settings) GetMandatoryChangeMessage() *bool {
	if o == nil {
		return nil
	}
	return o.MandatoryChangeMessage
}

func (o *Settings) GetChangeMessagePrompt() *string {
	if o == nil {
		return nil
	}
	return o.ChangeMessagePrompt
}

func (o *Settings) GetEnableChangeRequest() *bool {
	if o == nil {
		return nil
	}
	return o.EnableChangeRequest
}

func (o *Settings) GetEnableSelfValidation() *bool {
	if o == nil {
		return nil
	}
	return o.EnableSelfValidation
}

func (o *Settings) GetEnableSelfDeployment() *bool {
	if o == nil {
		return nil
	}
	return o.EnableSelfDeployment
}

func (o *Settings) GetDisplayRecentChangesGraphs() *bool {
	if o == nil {
		return nil
	}
	return o.DisplayRecentChangesGraphs
}

func (o *Settings) GetEnableJavascriptDirectives() *string {
	if o == nil {
		return nil
	}
	return o.EnableJavascriptDirectives
}

func (o *Settings) GetSendMetrics() *string {
	if o == nil {
		return nil
	}
	return o.SendMetrics
}

func (o *Settings) GetNodeAcceptDuplicatedHostname() *bool {
	if o == nil {
		return nil
	}
	return o.NodeAcceptDuplicatedHostname
}

func (o *Settings) GetNodeOnacceptDefaultState() *NodeOnacceptDefaultState {
	if o == nil {
		return nil
	}
	return o.NodeOnacceptDefaultState
}

func (o *Settings) GetNodeOnacceptDefaultPolicyMode() *NodeOnacceptDefaultPolicyMode {
	if o == nil {
		return nil
	}
	return o.NodeOnacceptDefaultPolicyMode
}

func (o *Settings) GetUnexpectedUnboundVarValues() *bool {
	if o == nil {
		return nil
	}
	return o.UnexpectedUnboundVarValues
}

func (o *Settings) GetRudderComputeChanges() *bool {
	if o == nil {
		return nil
	}
	return o.RudderComputeChanges
}

func (o *Settings) GetRudderGenerationComputeDyngroups() *bool {
	if o == nil {
		return nil
	}
	return o.RudderGenerationComputeDyngroups
}

func (o *Settings) GetRudderComputeDyngroupsMaxParallelism() *string {
	if o == nil {
		return nil
	}
	return o.RudderComputeDyngroupsMaxParallelism
}

func (o *Settings) GetRudderSaveDbComplianceLevels() *bool {
	if o == nil {
		return nil
	}
	return o.RudderSaveDbComplianceLevels
}

func (o *Settings) GetRudderSaveDbComplianceDetails() *bool {
	if o == nil {
		return nil
	}
	return o.RudderSaveDbComplianceDetails
}

func (o *Settings) GetRudderGenerationMaxParallelism() *string {
	if o == nil {
		return nil
	}
	return o.RudderGenerationMaxParallelism
}

func (o *Settings) GetRudderGenerationJsTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.RudderGenerationJsTimeout
}

func (o *Settings) GetRudderGenerationContinueOnError() *bool {
	if o == nil {
		return nil
	}
	return o.RudderGenerationContinueOnError
}

func (o *Settings) GetRudderGenerationDelay() *string {
	if o == nil {
		return nil
	}
	return o.RudderGenerationDelay
}

func (o *Settings) GetRudderGenerationPolicy() *string {
	if o == nil {
		return nil
	}
	return o.RudderGenerationPolicy
}

// GetAllSettingsData - Information about the setting
type GetAllSettingsData struct {
	Settings Settings `json:"settings"`
}

func (o *GetAllSettingsData) GetSettings() Settings {
	if o == nil {
		return Settings{}
	}
	return o.Settings
}

// GetAllSettingsResponseBody - Settings
type GetAllSettingsResponseBody struct {
	// Result of the request
	Result GetAllSettingsResult `json:"result"`
	// The id of the action
	Action GetAllSettingsAction `json:"action"`
	// Information about the setting
	Data GetAllSettingsData `json:"data"`
}

func (o *GetAllSettingsResponseBody) GetResult() GetAllSettingsResult {
	if o == nil {
		return GetAllSettingsResult("")
	}
	return o.Result
}

func (o *GetAllSettingsResponseBody) GetAction() GetAllSettingsAction {
	if o == nil {
		return GetAllSettingsAction("")
	}
	return o.Action
}

func (o *GetAllSettingsResponseBody) GetData() GetAllSettingsData {
	if o == nil {
		return GetAllSettingsData{}
	}
	return o.Data
}

type GetAllSettingsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Settings
	Object *GetAllSettingsResponseBody
}

func (o *GetAllSettingsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAllSettingsResponse) GetObject() *GetAllSettingsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
