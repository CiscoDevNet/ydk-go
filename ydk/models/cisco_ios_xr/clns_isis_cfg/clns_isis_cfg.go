// This module contains a collection of YANG definitions
// for Cisco IOS-XR clns-isis package configuration.
// 
// This module contains definitions
// for the following management objects:
//   isis: IS-IS configuration for all instances
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package clns_isis_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package clns_isis_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-clns-isis-cfg isis}", reflect.TypeOf(Isis{}))
    ydk.RegisterEntity("Cisco-IOS-XR-clns-isis-cfg:isis", reflect.TypeOf(Isis{}))
}

// IsisSnpAuth represents Isis snp auth
type IsisSnpAuth string

const (
    // Authenticate SNP send only
    IsisSnpAuth_send_only IsisSnpAuth = "send-only"

    // Authenticate SNP send and recv
    IsisSnpAuth_full IsisSnpAuth = "full"
)

// IsisMibMaxAreaAddressMismatchBoolean represents Isis mib max area address mismatch boolean
type IsisMibMaxAreaAddressMismatchBoolean string

const (
    // Disable
    IsisMibMaxAreaAddressMismatchBoolean_false IsisMibMaxAreaAddressMismatchBoolean = "false"

    // Enable
    IsisMibMaxAreaAddressMismatchBoolean_true IsisMibMaxAreaAddressMismatchBoolean = "true"
)

// IsisMibLspTooLargeToPropagateBoolean represents Isis mib lsp too large to propagate boolean
type IsisMibLspTooLargeToPropagateBoolean string

const (
    // Disable
    IsisMibLspTooLargeToPropagateBoolean_false IsisMibLspTooLargeToPropagateBoolean = "false"

    // Enable
    IsisMibLspTooLargeToPropagateBoolean_true IsisMibLspTooLargeToPropagateBoolean = "true"
)

// IsisMibSequenceNumberSkipBoolean represents Isis mib sequence number skip boolean
type IsisMibSequenceNumberSkipBoolean string

const (
    // Disable
    IsisMibSequenceNumberSkipBoolean_false IsisMibSequenceNumberSkipBoolean = "false"

    // Enable
    IsisMibSequenceNumberSkipBoolean_true IsisMibSequenceNumberSkipBoolean = "true"
)

// IsisInterfaceFrrTiebreaker represents Isis interface frr tiebreaker
type IsisInterfaceFrrTiebreaker string

const (
    // Prefer node protecting backup path
    IsisInterfaceFrrTiebreaker_node_protecting IsisInterfaceFrrTiebreaker = "node-protecting"

    // Prefer SRLG disjoint backup path
    IsisInterfaceFrrTiebreaker_srlg_disjoint IsisInterfaceFrrTiebreaker = "srlg-disjoint"
)

// IsisAuthenticationAlgorithm represents Isis authentication algorithm
type IsisAuthenticationAlgorithm string

const (
    // Cleartext password
    IsisAuthenticationAlgorithm_cleartext IsisAuthenticationAlgorithm = "cleartext"

    // HMAC-MD5 checksum
    IsisAuthenticationAlgorithm_hmac_md5 IsisAuthenticationAlgorithm = "hmac-md5"

    // Key Chain authentication
    IsisAuthenticationAlgorithm_keychain IsisAuthenticationAlgorithm = "keychain"
)

// IsisOverloadBitMode represents Isis overload bit mode
type IsisOverloadBitMode string

const (
    // Set always
    IsisOverloadBitMode_permanently_set IsisOverloadBitMode = "permanently-set"

    // Set during the startup period
    IsisOverloadBitMode_startup_period IsisOverloadBitMode = "startup-period"

    // Set until BGP comverges
    IsisOverloadBitMode_wait_for_bgp IsisOverloadBitMode = "wait-for-bgp"
)

// IsisMibRejectedAdjacencyBoolean represents Isis mib rejected adjacency boolean
type IsisMibRejectedAdjacencyBoolean string

const (
    // Disable
    IsisMibRejectedAdjacencyBoolean_false IsisMibRejectedAdjacencyBoolean = "false"

    // Enable
    IsisMibRejectedAdjacencyBoolean_true IsisMibRejectedAdjacencyBoolean = "true"
)

// IsisMibCorruptedLspDetectedBoolean represents Isis mib corrupted lsp detected boolean
type IsisMibCorruptedLspDetectedBoolean string

const (
    // Disable
    IsisMibCorruptedLspDetectedBoolean_false IsisMibCorruptedLspDetectedBoolean = "false"

    // Enable
    IsisMibCorruptedLspDetectedBoolean_true IsisMibCorruptedLspDetectedBoolean = "true"
)

// IsisAdjCheck represents Isis adj check
type IsisAdjCheck string

const (
    // Disabled
    IsisAdjCheck_disabled IsisAdjCheck = "disabled"
)

// IsisispfState represents Isisispf state
type IsisispfState string

const (
    // Enabled
    IsisispfState_enabled IsisispfState = "enabled"
)

// IsisfrrLoadSharing represents Isisfrr load sharing
type IsisfrrLoadSharing string

const (
    // Disable load sharing of prefixes across
    // multiple backups
    IsisfrrLoadSharing_disable IsisfrrLoadSharing = "disable"
)

// IsisMibAuthenticationFailureBoolean represents Isis mib authentication failure boolean
type IsisMibAuthenticationFailureBoolean string

const (
    // Disable
    IsisMibAuthenticationFailureBoolean_false IsisMibAuthenticationFailureBoolean = "false"

    // Enable
    IsisMibAuthenticationFailureBoolean_true IsisMibAuthenticationFailureBoolean = "true"
)

// IsisInterfaceState represents Isis interface state
type IsisInterfaceState string

const (
    // Shutdown
    IsisInterfaceState_shutdown IsisInterfaceState = "shutdown"

    // Suppressed
    IsisInterfaceState_suppressed IsisInterfaceState = "suppressed"

    // Passive
    IsisInterfaceState_passive IsisInterfaceState = "passive"

    // EnabledActive
    IsisInterfaceState_enabled_active IsisInterfaceState = "enabled-active"
)

// IsisTracingMode represents Isis tracing mode
type IsisTracingMode string

const (
    // No tracing
    IsisTracingMode_off IsisTracingMode = "off"

    // Basic tracing (less overhead)
    IsisTracingMode_basic IsisTracingMode = "basic"

    // Enhanced tracing (more overhead)
    IsisTracingMode_enhanced IsisTracingMode = "enhanced"
)

// IsisMetricStyle represents Isis metric style
type IsisMetricStyle string

const (
    // ISO 10589 metric style (old-style)
    IsisMetricStyle_old_metric_style IsisMetricStyle = "old-metric-style"

    // 32-bit metric style (new-style)
    IsisMetricStyle_new_metric_style IsisMetricStyle = "new-metric-style"

    // Both forms of metric style
    IsisMetricStyle_both_metric_style IsisMetricStyle = "both-metric-style"
)

// IsisNsfFlavor represents Isis nsf flavor
type IsisNsfFlavor string

const (
    // Cisco proprietary NSF
    IsisNsfFlavor_cisco_proprietary_nsf IsisNsfFlavor = "cisco-proprietary-nsf"

    // IETF standard NSF
    IsisNsfFlavor_ietf_standard_nsf IsisNsfFlavor = "ietf-standard-nsf"
)

// IsisInterfaceAfState represents Isis interface af state
type IsisInterfaceAfState string

const (
    // Disable
    IsisInterfaceAfState_disable IsisInterfaceAfState = "disable"
)

// IsissidProtected represents Isissid protected
type IsissidProtected string

const (
    // Not protected
    IsissidProtected_disable IsissidProtected = "disable"

    // Protected
    IsissidProtected_enable IsissidProtected = "enable"
)

// IsisApplyWeight represents Isis apply weight
type IsisApplyWeight string

const (
    // Apply weight to ECMP prefixes
    IsisApplyWeight_ecmp_only IsisApplyWeight = "ecmp-only"

    // Apply weight to UCMP prefixes
    IsisApplyWeight_ucmp_only IsisApplyWeight = "ucmp-only"
)

// IsisPrefixPriority represents Isis prefix priority
type IsisPrefixPriority string

const (
    // Critical prefix priority
    IsisPrefixPriority_critical_priority IsisPrefixPriority = "critical-priority"

    // High prefix priority
    IsisPrefixPriority_high_priority IsisPrefixPriority = "high-priority"

    // Medium prefix priority
    IsisPrefixPriority_medium_priority IsisPrefixPriority = "medium-priority"
)

// IsisMibAuthenticationTypeFailureBoolean represents Isis mib authentication type failure boolean
type IsisMibAuthenticationTypeFailureBoolean string

const (
    // Disable
    IsisMibAuthenticationTypeFailureBoolean_false IsisMibAuthenticationTypeFailureBoolean = "false"

    // Enable
    IsisMibAuthenticationTypeFailureBoolean_true IsisMibAuthenticationTypeFailureBoolean = "true"
)

// IsisMicroLoopAvoidance represents Isis micro loop avoidance
type IsisMicroLoopAvoidance string

const (
    // No Avoidance type set
    IsisMicroLoopAvoidance_not_set IsisMicroLoopAvoidance = "not-set"

    // Provide mirco loop avoidance for all prefixes
    IsisMicroLoopAvoidance_micro_loop_avoidance_all IsisMicroLoopAvoidance = "micro-loop-avoidance-all"

    // Provide mirco loop avoidance only for protected
    // prefixes
    IsisMicroLoopAvoidance_micro_loop_avoidance_protected IsisMicroLoopAvoidance = "micro-loop-avoidance-protected"

    // Provide segment-routing mirco loop avoidance
    IsisMicroLoopAvoidance_micro_loop_avoidance_segement_routing IsisMicroLoopAvoidance = "micro-loop-avoidance-segement-routing"
)

// IsisAdvTypeExternal represents Isis adv type external
type IsisAdvTypeExternal string

const (
    // External
    IsisAdvTypeExternal_external IsisAdvTypeExternal = "external"
)

// IsisRemoteLfa represents Isis remote lfa
type IsisRemoteLfa string

const (
    // No remote LFA option set
    IsisRemoteLfa_remote_lfa_none IsisRemoteLfa = "remote-lfa-none"

    // Construct remote LFA tunnel using MPLS LDP
    IsisRemoteLfa_remote_lfa_tunnel_ldp IsisRemoteLfa = "remote-lfa-tunnel-ldp"
)

// IsisMibAreaMismatchBoolean represents Isis mib area mismatch boolean
type IsisMibAreaMismatchBoolean string

const (
    // Disable
    IsisMibAreaMismatchBoolean_false IsisMibAreaMismatchBoolean = "false"

    // Enable
    IsisMibAreaMismatchBoolean_true IsisMibAreaMismatchBoolean = "true"
)

// IsisMibAttemptToExceedMaxSequenceBoolean represents Isis mib attempt to exceed max sequence boolean
type IsisMibAttemptToExceedMaxSequenceBoolean string

const (
    // Disable
    IsisMibAttemptToExceedMaxSequenceBoolean_false IsisMibAttemptToExceedMaxSequenceBoolean = "false"

    // Enable
    IsisMibAttemptToExceedMaxSequenceBoolean_true IsisMibAttemptToExceedMaxSequenceBoolean = "true"
)

// IsisConfigurableLevels represents Isis configurable levels
type IsisConfigurableLevels string

const (
    // Level1
    IsisConfigurableLevels_level1 IsisConfigurableLevels = "level1"

    // Level2
    IsisConfigurableLevels_level2 IsisConfigurableLevels = "level2"

    // Both Levels
    IsisConfigurableLevels_level1_and2 IsisConfigurableLevels = "level1-and2"
)

// IsisfrrTiebreaker represents Isisfrr tiebreaker
type IsisfrrTiebreaker string

const (
    // Prefer backup path via downstream node
    IsisfrrTiebreaker_downstream IsisfrrTiebreaker = "downstream"

    // Prefer line card disjoint backup path
    IsisfrrTiebreaker_lc_disjoint IsisfrrTiebreaker = "lc-disjoint"

    // Prefer backup path with lowest total metric
    IsisfrrTiebreaker_lowest_backup_metric IsisfrrTiebreaker = "lowest-backup-metric"

    // Prefer node protecting backup path
    IsisfrrTiebreaker_node_protecting IsisfrrTiebreaker = "node-protecting"

    // Prefer backup path from ECMP set
    IsisfrrTiebreaker_primary_path IsisfrrTiebreaker = "primary-path"

    // Prefer non-ECMP backup path
    IsisfrrTiebreaker_secondary_path IsisfrrTiebreaker = "secondary-path"

    // Prefer SRLG disjoint backup path
    IsisfrrTiebreaker_srlg_disjoint IsisfrrTiebreaker = "srlg-disjoint"
)

// IsisMibManualAddressDropsBoolean represents Isis mib manual address drops boolean
type IsisMibManualAddressDropsBoolean string

const (
    // Disable
    IsisMibManualAddressDropsBoolean_false IsisMibManualAddressDropsBoolean = "false"

    // Enable
    IsisMibManualAddressDropsBoolean_true IsisMibManualAddressDropsBoolean = "true"
)

// IsisMetricStyleTransition represents Isis metric style transition
type IsisMetricStyleTransition string

const (
    // Disabled
    IsisMetricStyleTransition_disabled IsisMetricStyleTransition = "disabled"

    // Enabled
    IsisMetricStyleTransition_enabled IsisMetricStyleTransition = "enabled"
)

// IsisexplicitNullFlag represents Isisexplicit null flag
type IsisexplicitNullFlag string

const (
    // Disable EXPLICITNULL
    IsisexplicitNullFlag_disable IsisexplicitNullFlag = "disable"

    // Enable EXPLICITNULL
    IsisexplicitNullFlag_enable IsisexplicitNullFlag = "enable"
)

// IsisMetric represents Isis metric
type IsisMetric string

const (
    // Internal metric
    IsisMetric_internal IsisMetric = "internal"

    // External metric
    IsisMetric_external IsisMetric = "external"
)

// IsisHelloPadding represents Isis hello padding
type IsisHelloPadding string

const (
    // Never pad Hellos
    IsisHelloPadding_never IsisHelloPadding = "never"

    // Pad Hellos during adjacency formation only
    IsisHelloPadding_sometimes IsisHelloPadding = "sometimes"
)

// IsisMibDatabaseOverFlowBoolean represents Isis mib database over flow boolean
type IsisMibDatabaseOverFlowBoolean string

const (
    // Disable
    IsisMibDatabaseOverFlowBoolean_false IsisMibDatabaseOverFlowBoolean = "false"

    // Enable
    IsisMibDatabaseOverFlowBoolean_true IsisMibDatabaseOverFlowBoolean = "true"
)

// IsisAdvTypeInterLevel represents Isis adv type inter level
type IsisAdvTypeInterLevel string

const (
    // InterLevel
    IsisAdvTypeInterLevel_inter_level IsisAdvTypeInterLevel = "inter-level"
)

// IsisAuthenticationFailureMode represents Isis authentication failure mode
type IsisAuthenticationFailureMode string

const (
    // Drop non-authenticating PDUs
    IsisAuthenticationFailureMode_drop IsisAuthenticationFailureMode = "drop"

    // Accept non-authenticating PDUs
    IsisAuthenticationFailureMode_send_only IsisAuthenticationFailureMode = "send-only"
)

// IsisMibProtocolsSupportedMismatchBoolean represents Isis mib protocols supported mismatch boolean
type IsisMibProtocolsSupportedMismatchBoolean string

const (
    // Disable
    IsisMibProtocolsSupportedMismatchBoolean_false IsisMibProtocolsSupportedMismatchBoolean = "false"

    // Enable
    IsisMibProtocolsSupportedMismatchBoolean_true IsisMibProtocolsSupportedMismatchBoolean = "true"
)

// IsisRedistProto represents Isis redist proto
type IsisRedistProto string

const (
    // Connected
    IsisRedistProto_connected IsisRedistProto = "connected"

    // Static
    IsisRedistProto_static IsisRedistProto = "static"

    // OSPF
    IsisRedistProto_ospf IsisRedistProto = "ospf"

    // BGP
    IsisRedistProto_bgp IsisRedistProto = "bgp"

    // ISIS
    IsisRedistProto_isis IsisRedistProto = "isis"

    // OSPFv3
    IsisRedistProto_ospfv3 IsisRedistProto = "ospfv3"

    // RIP
    IsisRedistProto_rip IsisRedistProto = "rip"

    // EIGRP
    IsisRedistProto_eigrp IsisRedistProto = "eigrp"

    // Subscriber
    IsisRedistProto_subscriber IsisRedistProto = "subscriber"

    // Application
    IsisRedistProto_application IsisRedistProto = "application"

    // Mobile
    IsisRedistProto_mobile IsisRedistProto = "mobile"
)

// Isissid1 represents Isissid1
type Isissid1 string

const (
    // SID as an index
    Isissid1_index Isissid1 = "index"

    // SID as an absolute label
    Isissid1_absolute Isissid1 = "absolute"
)

// IsisphpFlag represents Isisphp flag
type IsisphpFlag string

const (
    // Enable PHP
    IsisphpFlag_enable IsisphpFlag = "enable"

    // Disable PHP
    IsisphpFlag_disable IsisphpFlag = "disable"
)

// IsisMibIdLengthMismatchBoolean represents Isis mib id length mismatch boolean
type IsisMibIdLengthMismatchBoolean string

const (
    // Disable
    IsisMibIdLengthMismatchBoolean_false IsisMibIdLengthMismatchBoolean = "false"

    // Enable
    IsisMibIdLengthMismatchBoolean_true IsisMibIdLengthMismatchBoolean = "true"
)

// IsisMibAllBoolean represents Isis mib all boolean
type IsisMibAllBoolean string

const (
    // Disable
    IsisMibAllBoolean_false IsisMibAllBoolean = "false"

    // Enable
    IsisMibAllBoolean_true IsisMibAllBoolean = "true"
)

// IsisMibOriginatedLspBufferSizeMismatchBoolean represents boolean
type IsisMibOriginatedLspBufferSizeMismatchBoolean string

const (
    // Disable
    IsisMibOriginatedLspBufferSizeMismatchBoolean_false IsisMibOriginatedLspBufferSizeMismatchBoolean = "false"

    // Enable
    IsisMibOriginatedLspBufferSizeMismatchBoolean_true IsisMibOriginatedLspBufferSizeMismatchBoolean = "true"
)

// Isisfrr represents Isisfrr
type Isisfrr string

const (
    // Prefix independent per-link computation
    Isisfrr_per_link Isisfrr = "per-link"

    // Prefix dependent computation
    Isisfrr_per_prefix Isisfrr = "per-prefix"
)

// IsisAttachedBit represents Isis attached bit
type IsisAttachedBit string

const (
    // Computed from the attached areas
    IsisAttachedBit_area IsisAttachedBit = "area"

    // Forced ON
    IsisAttachedBit_on IsisAttachedBit = "on"

    // Forced OFF
    IsisAttachedBit_off IsisAttachedBit = "off"
)

// NflagClear represents Nflag clear
type NflagClear string

const (
    // Disable N-flag-clear
    NflagClear_disable NflagClear = "disable"

    // Enable N-flag-clear
    NflagClear_enable NflagClear = "enable"
)

// IsisLabelPreference represents Isis label preference
type IsisLabelPreference string

const (
    // Label Distribution Protocol
    IsisLabelPreference_ldp IsisLabelPreference = "ldp"

    // Segment Routing
    IsisLabelPreference_segment_routing IsisLabelPreference = "segment-routing"
)

// IsisMibAdjacencyChangeBoolean represents Isis mib adjacency change boolean
type IsisMibAdjacencyChangeBoolean string

const (
    // Disable
    IsisMibAdjacencyChangeBoolean_false IsisMibAdjacencyChangeBoolean = "false"

    // Enable
    IsisMibAdjacencyChangeBoolean_true IsisMibAdjacencyChangeBoolean = "true"
)

// IsisMibLspErrorDetectedBoolean represents Isis mib lsp error detected boolean
type IsisMibLspErrorDetectedBoolean string

const (
    // Disable
    IsisMibLspErrorDetectedBoolean_false IsisMibLspErrorDetectedBoolean = "false"

    // Enable
    IsisMibLspErrorDetectedBoolean_true IsisMibLspErrorDetectedBoolean = "true"
)

// IsisMibOwnLspPurgeBoolean represents Isis mib own lsp purge boolean
type IsisMibOwnLspPurgeBoolean string

const (
    // Disable
    IsisMibOwnLspPurgeBoolean_false IsisMibOwnLspPurgeBoolean = "false"

    // Enable
    IsisMibOwnLspPurgeBoolean_true IsisMibOwnLspPurgeBoolean = "true"
)

// IsisMibVersionSkewBoolean represents Isis mib version skew boolean
type IsisMibVersionSkewBoolean string

const (
    // Disable
    IsisMibVersionSkewBoolean_false IsisMibVersionSkewBoolean = "false"

    // Enable
    IsisMibVersionSkewBoolean_true IsisMibVersionSkewBoolean = "true"
)

// Isis
// IS-IS configuration for all instances
type Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IS-IS instance configuration.
    Instances Isis_Instances
}

func (isis *Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Isis) GetGoName(yname string) string {
    if yname == "instances" { return "Instances" }
    return ""
}

func (isis *Isis) GetSegmentPath() string {
    return "Cisco-IOS-XR-clns-isis-cfg:isis"
}

func (isis *Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instances" {
        return &isis.Instances
    }
    return nil
}

func (isis *Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instances"] = &isis.Instances
    return children
}

func (isis *Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (isis *Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Isis) GetYangName() string { return "isis" }

func (isis *Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Isis) GetParent() types.Entity { return isis.parent }

func (isis *Isis) GetParentYangName() string { return "Cisco-IOS-XR-clns-isis-cfg" }

// Isis_Instances
// IS-IS instance configuration
type Isis_Instances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a single IS-IS instance. The type is slice of
    // Isis_Instances_Instance.
    Instance []Isis_Instances_Instance
}

func (instances *Isis_Instances) GetFilter() yfilter.YFilter { return instances.YFilter }

func (instances *Isis_Instances) SetFilter(yf yfilter.YFilter) { instances.YFilter = yf }

func (instances *Isis_Instances) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    return ""
}

func (instances *Isis_Instances) GetSegmentPath() string {
    return "instances"
}

func (instances *Isis_Instances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        for _, c := range instances.Instance {
            if instances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance{}
        instances.Instance = append(instances.Instance, child)
        return &instances.Instance[len(instances.Instance)-1]
    }
    return nil
}

func (instances *Isis_Instances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range instances.Instance {
        children[instances.Instance[i].GetSegmentPath()] = &instances.Instance[i]
    }
    return children
}

func (instances *Isis_Instances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (instances *Isis_Instances) GetBundleName() string { return "cisco_ios_xr" }

func (instances *Isis_Instances) GetYangName() string { return "instances" }

func (instances *Isis_Instances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instances *Isis_Instances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instances *Isis_Instances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instances *Isis_Instances) SetParent(parent types.Entity) { instances.parent = parent }

func (instances *Isis_Instances) GetParent() types.Entity { return instances.parent }

func (instances *Isis_Instances) GetParentYangName() string { return "isis" }

// Isis_Instances_Instance
// Configuration for a single IS-IS instance
type Isis_Instances_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Instance identifier. The type is string with
    // length: 1..40.
    InstanceName interface{}

    // Flag to indicate that instance should be running.  This must be the first
    // object created when an IS-IS instance is configured, and the last object
    // deleted when it is deconfigured.  When this object is deleted, the IS-IS
    // instance will exit. The type is interface{}.
    Running interface{}

    // Log changes in adjacency state. The type is interface{}.
    LogAdjacencyChanges interface{}

    // If TRUE, LSPs recieved with bad checksums will result in the purging of
    // that LSP from the LSP DB. If FALSE or not set, the received LSP will just
    // be ignored. The type is bool.
    IgnoreLspErrors interface{}

    // IS type of the IS-IS process. The type is IsisConfigurableLevels. The
    // default value is level1-and2.
    IsType interface{}

    // Tracing mode configuration. The type is IsisTracingMode. The default value
    // is basic.
    TracingMode interface{}

    // Instance ID of the IS-IS process. The type is interface{} with range:
    // 0..65535. The default value is 0.
    InstanceId interface{}

    // If TRUE, dynamic hostname resolution is disabled, and system IDs will
    // always be displayed by show and debug output. The type is bool.
    DynamicHostName interface{}

    // IS-IS NSR configuration. The type is interface{}.
    Nsr interface{}

    // Log PDU drops. The type is interface{}.
    LogPduDrops interface{}

    // Segment Routing Global Block configuration.
    Srgb Isis_Instances_Instance_Srgb

    // LSP generation-interval configuration.
    LspGenerationIntervals Isis_Instances_Instance_LspGenerationIntervals

    // LSP arrival time configuration.
    LspArrivalTimes Isis_Instances_Instance_LspArrivalTimes

    // Trace buffer size configuration.
    TraceBufferSize Isis_Instances_Instance_TraceBufferSize

    // Max Link Metric configuration.
    MaxLinkMetrics Isis_Instances_Instance_MaxLinkMetrics

    // Stagger ISIS adjacency bring up.
    AdjacencyStagger Isis_Instances_Instance_AdjacencyStagger

    // Per-address-family configuration.
    Afs Isis_Instances_Instance_Afs

    // LSP refresh-interval configuration.
    LspRefreshIntervals Isis_Instances_Instance_LspRefreshIntervals

    // IS-IS Distribute BGP-LS configuration.
    Distribute Isis_Instances_Instance_Distribute

    // LSP/SNP accept password configuration.
    LspAcceptPasswords Isis_Instances_Instance_LspAcceptPasswords

    // LSP MTU configuration.
    LspMtus Isis_Instances_Instance_LspMtus

    // IS-IS NSF configuration.
    Nsf Isis_Instances_Instance_Nsf

    // Link Group.
    LinkGroups Isis_Instances_Instance_LinkGroups

    // LSP checksum check interval configuration.
    LspCheckIntervals Isis_Instances_Instance_LspCheckIntervals

    // LSP/SNP password configuration.
    LspPasswords Isis_Instances_Instance_LspPasswords

    // NET configuration.
    Nets Isis_Instances_Instance_Nets

    // LSP lifetime configuration.
    LspLifetimes Isis_Instances_Instance_LspLifetimes

    // LSP overload-bit configuration.
    OverloadBits Isis_Instances_Instance_OverloadBits

    // Per-interface configuration.
    Interfaces Isis_Instances_Instance_Interfaces
}

func (instance *Isis_Instances_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *Isis_Instances_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *Isis_Instances_Instance) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "running" { return "Running" }
    if yname == "log-adjacency-changes" { return "LogAdjacencyChanges" }
    if yname == "ignore-lsp-errors" { return "IgnoreLspErrors" }
    if yname == "is-type" { return "IsType" }
    if yname == "tracing-mode" { return "TracingMode" }
    if yname == "instance-id" { return "InstanceId" }
    if yname == "dynamic-host-name" { return "DynamicHostName" }
    if yname == "nsr" { return "Nsr" }
    if yname == "log-pdu-drops" { return "LogPduDrops" }
    if yname == "srgb" { return "Srgb" }
    if yname == "lsp-generation-intervals" { return "LspGenerationIntervals" }
    if yname == "lsp-arrival-times" { return "LspArrivalTimes" }
    if yname == "trace-buffer-size" { return "TraceBufferSize" }
    if yname == "max-link-metrics" { return "MaxLinkMetrics" }
    if yname == "adjacency-stagger" { return "AdjacencyStagger" }
    if yname == "afs" { return "Afs" }
    if yname == "lsp-refresh-intervals" { return "LspRefreshIntervals" }
    if yname == "distribute" { return "Distribute" }
    if yname == "lsp-accept-passwords" { return "LspAcceptPasswords" }
    if yname == "lsp-mtus" { return "LspMtus" }
    if yname == "nsf" { return "Nsf" }
    if yname == "link-groups" { return "LinkGroups" }
    if yname == "lsp-check-intervals" { return "LspCheckIntervals" }
    if yname == "lsp-passwords" { return "LspPasswords" }
    if yname == "nets" { return "Nets" }
    if yname == "lsp-lifetimes" { return "LspLifetimes" }
    if yname == "overload-bits" { return "OverloadBits" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (instance *Isis_Instances_Instance) GetSegmentPath() string {
    return "instance" + "[instance-name='" + fmt.Sprintf("%v", instance.InstanceName) + "']"
}

func (instance *Isis_Instances_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "srgb" {
        return &instance.Srgb
    }
    if childYangName == "lsp-generation-intervals" {
        return &instance.LspGenerationIntervals
    }
    if childYangName == "lsp-arrival-times" {
        return &instance.LspArrivalTimes
    }
    if childYangName == "trace-buffer-size" {
        return &instance.TraceBufferSize
    }
    if childYangName == "max-link-metrics" {
        return &instance.MaxLinkMetrics
    }
    if childYangName == "adjacency-stagger" {
        return &instance.AdjacencyStagger
    }
    if childYangName == "afs" {
        return &instance.Afs
    }
    if childYangName == "lsp-refresh-intervals" {
        return &instance.LspRefreshIntervals
    }
    if childYangName == "distribute" {
        return &instance.Distribute
    }
    if childYangName == "lsp-accept-passwords" {
        return &instance.LspAcceptPasswords
    }
    if childYangName == "lsp-mtus" {
        return &instance.LspMtus
    }
    if childYangName == "nsf" {
        return &instance.Nsf
    }
    if childYangName == "link-groups" {
        return &instance.LinkGroups
    }
    if childYangName == "lsp-check-intervals" {
        return &instance.LspCheckIntervals
    }
    if childYangName == "lsp-passwords" {
        return &instance.LspPasswords
    }
    if childYangName == "nets" {
        return &instance.Nets
    }
    if childYangName == "lsp-lifetimes" {
        return &instance.LspLifetimes
    }
    if childYangName == "overload-bits" {
        return &instance.OverloadBits
    }
    if childYangName == "interfaces" {
        return &instance.Interfaces
    }
    return nil
}

func (instance *Isis_Instances_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["srgb"] = &instance.Srgb
    children["lsp-generation-intervals"] = &instance.LspGenerationIntervals
    children["lsp-arrival-times"] = &instance.LspArrivalTimes
    children["trace-buffer-size"] = &instance.TraceBufferSize
    children["max-link-metrics"] = &instance.MaxLinkMetrics
    children["adjacency-stagger"] = &instance.AdjacencyStagger
    children["afs"] = &instance.Afs
    children["lsp-refresh-intervals"] = &instance.LspRefreshIntervals
    children["distribute"] = &instance.Distribute
    children["lsp-accept-passwords"] = &instance.LspAcceptPasswords
    children["lsp-mtus"] = &instance.LspMtus
    children["nsf"] = &instance.Nsf
    children["link-groups"] = &instance.LinkGroups
    children["lsp-check-intervals"] = &instance.LspCheckIntervals
    children["lsp-passwords"] = &instance.LspPasswords
    children["nets"] = &instance.Nets
    children["lsp-lifetimes"] = &instance.LspLifetimes
    children["overload-bits"] = &instance.OverloadBits
    children["interfaces"] = &instance.Interfaces
    return children
}

func (instance *Isis_Instances_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = instance.InstanceName
    leafs["running"] = instance.Running
    leafs["log-adjacency-changes"] = instance.LogAdjacencyChanges
    leafs["ignore-lsp-errors"] = instance.IgnoreLspErrors
    leafs["is-type"] = instance.IsType
    leafs["tracing-mode"] = instance.TracingMode
    leafs["instance-id"] = instance.InstanceId
    leafs["dynamic-host-name"] = instance.DynamicHostName
    leafs["nsr"] = instance.Nsr
    leafs["log-pdu-drops"] = instance.LogPduDrops
    return leafs
}

func (instance *Isis_Instances_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *Isis_Instances_Instance) GetYangName() string { return "instance" }

func (instance *Isis_Instances_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *Isis_Instances_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *Isis_Instances_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *Isis_Instances_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *Isis_Instances_Instance) GetParent() types.Entity { return instance.parent }

func (instance *Isis_Instances_Instance) GetParentYangName() string { return "instances" }

// Isis_Instances_Instance_Srgb
// Segment Routing Global Block configuration
// This type is a presence type.
type Isis_Instances_Instance_Srgb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The lower bound of the SRGB. The type is interface{} with range:
    // 16000..1048574. This attribute is mandatory.
    LowerBound interface{}

    // The upper bound of the SRGB. The type is interface{} with range:
    // 16001..1048575. This attribute is mandatory.
    UpperBound interface{}
}

func (srgb *Isis_Instances_Instance_Srgb) GetFilter() yfilter.YFilter { return srgb.YFilter }

func (srgb *Isis_Instances_Instance_Srgb) SetFilter(yf yfilter.YFilter) { srgb.YFilter = yf }

func (srgb *Isis_Instances_Instance_Srgb) GetGoName(yname string) string {
    if yname == "lower-bound" { return "LowerBound" }
    if yname == "upper-bound" { return "UpperBound" }
    return ""
}

func (srgb *Isis_Instances_Instance_Srgb) GetSegmentPath() string {
    return "srgb"
}

func (srgb *Isis_Instances_Instance_Srgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srgb *Isis_Instances_Instance_Srgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srgb *Isis_Instances_Instance_Srgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lower-bound"] = srgb.LowerBound
    leafs["upper-bound"] = srgb.UpperBound
    return leafs
}

func (srgb *Isis_Instances_Instance_Srgb) GetBundleName() string { return "cisco_ios_xr" }

func (srgb *Isis_Instances_Instance_Srgb) GetYangName() string { return "srgb" }

func (srgb *Isis_Instances_Instance_Srgb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgb *Isis_Instances_Instance_Srgb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgb *Isis_Instances_Instance_Srgb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgb *Isis_Instances_Instance_Srgb) SetParent(parent types.Entity) { srgb.parent = parent }

func (srgb *Isis_Instances_Instance_Srgb) GetParent() types.Entity { return srgb.parent }

func (srgb *Isis_Instances_Instance_Srgb) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspGenerationIntervals
// LSP generation-interval configuration
type Isis_Instances_Instance_LspGenerationIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSP generation scheduling parameters. The type is slice of
    // Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval.
    LspGenerationInterval []Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval
}

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetFilter() yfilter.YFilter { return lspGenerationIntervals.YFilter }

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) SetFilter(yf yfilter.YFilter) { lspGenerationIntervals.YFilter = yf }

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetGoName(yname string) string {
    if yname == "lsp-generation-interval" { return "LspGenerationInterval" }
    return ""
}

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetSegmentPath() string {
    return "lsp-generation-intervals"
}

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-generation-interval" {
        for _, c := range lspGenerationIntervals.LspGenerationInterval {
            if lspGenerationIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval{}
        lspGenerationIntervals.LspGenerationInterval = append(lspGenerationIntervals.LspGenerationInterval, child)
        return &lspGenerationIntervals.LspGenerationInterval[len(lspGenerationIntervals.LspGenerationInterval)-1]
    }
    return nil
}

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspGenerationIntervals.LspGenerationInterval {
        children[lspGenerationIntervals.LspGenerationInterval[i].GetSegmentPath()] = &lspGenerationIntervals.LspGenerationInterval[i]
    }
    return children
}

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetYangName() string { return "lsp-generation-intervals" }

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) SetParent(parent types.Entity) { lspGenerationIntervals.parent = parent }

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetParent() types.Entity { return lspGenerationIntervals.parent }

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval
// LSP generation scheduling parameters
type Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Maximum wait before generating local LSP in milliseconds. The type is
    // interface{} with range: 0..120000. Units are millisecond.
    MaximumWait interface{}

    // Initial wait before generating local LSP in milliseconds. The type is
    // interface{} with range: 0..120000. Units are millisecond.
    InitialWait interface{}

    // Secondary wait before generating local LSP in milliseconds. The type is
    // interface{} with range: 0..120000. Units are millisecond.
    SecondaryWait interface{}
}

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetFilter() yfilter.YFilter { return lspGenerationInterval.YFilter }

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) SetFilter(yf yfilter.YFilter) { lspGenerationInterval.YFilter = yf }

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "maximum-wait" { return "MaximumWait" }
    if yname == "initial-wait" { return "InitialWait" }
    if yname == "secondary-wait" { return "SecondaryWait" }
    return ""
}

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetSegmentPath() string {
    return "lsp-generation-interval" + "[level='" + fmt.Sprintf("%v", lspGenerationInterval.Level) + "']"
}

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspGenerationInterval.Level
    leafs["maximum-wait"] = lspGenerationInterval.MaximumWait
    leafs["initial-wait"] = lspGenerationInterval.InitialWait
    leafs["secondary-wait"] = lspGenerationInterval.SecondaryWait
    return leafs
}

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetBundleName() string { return "cisco_ios_xr" }

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetYangName() string { return "lsp-generation-interval" }

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) SetParent(parent types.Entity) { lspGenerationInterval.parent = parent }

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetParent() types.Entity { return lspGenerationInterval.parent }

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetParentYangName() string { return "lsp-generation-intervals" }

// Isis_Instances_Instance_LspArrivalTimes
// LSP arrival time configuration
type Isis_Instances_Instance_LspArrivalTimes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum LSP arrival time. The type is slice of
    // Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime.
    LspArrivalTime []Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime
}

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetFilter() yfilter.YFilter { return lspArrivalTimes.YFilter }

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) SetFilter(yf yfilter.YFilter) { lspArrivalTimes.YFilter = yf }

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetGoName(yname string) string {
    if yname == "lsp-arrival-time" { return "LspArrivalTime" }
    return ""
}

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetSegmentPath() string {
    return "lsp-arrival-times"
}

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-arrival-time" {
        for _, c := range lspArrivalTimes.LspArrivalTime {
            if lspArrivalTimes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime{}
        lspArrivalTimes.LspArrivalTime = append(lspArrivalTimes.LspArrivalTime, child)
        return &lspArrivalTimes.LspArrivalTime[len(lspArrivalTimes.LspArrivalTime)-1]
    }
    return nil
}

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspArrivalTimes.LspArrivalTime {
        children[lspArrivalTimes.LspArrivalTime[i].GetSegmentPath()] = &lspArrivalTimes.LspArrivalTime[i]
    }
    return children
}

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetBundleName() string { return "cisco_ios_xr" }

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetYangName() string { return "lsp-arrival-times" }

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) SetParent(parent types.Entity) { lspArrivalTimes.parent = parent }

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetParent() types.Entity { return lspArrivalTimes.parent }

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime
// Minimum LSP arrival time
type Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Maximum delay expected to take since last LSPin milliseconds. The type is
    // interface{} with range: 0..120000. Units are millisecond.
    MaximumWait interface{}

    // Initial delay expected to take since last LSPin milliseconds. The type is
    // interface{} with range: 0..120000. Units are millisecond.
    InitialWait interface{}

    // Secondary delay expected to take since last LSPin milliseconds. The type is
    // interface{} with range: 0..120000. Units are millisecond.
    SecondaryWait interface{}
}

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetFilter() yfilter.YFilter { return lspArrivalTime.YFilter }

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) SetFilter(yf yfilter.YFilter) { lspArrivalTime.YFilter = yf }

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "maximum-wait" { return "MaximumWait" }
    if yname == "initial-wait" { return "InitialWait" }
    if yname == "secondary-wait" { return "SecondaryWait" }
    return ""
}

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetSegmentPath() string {
    return "lsp-arrival-time" + "[level='" + fmt.Sprintf("%v", lspArrivalTime.Level) + "']"
}

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspArrivalTime.Level
    leafs["maximum-wait"] = lspArrivalTime.MaximumWait
    leafs["initial-wait"] = lspArrivalTime.InitialWait
    leafs["secondary-wait"] = lspArrivalTime.SecondaryWait
    return leafs
}

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetBundleName() string { return "cisco_ios_xr" }

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetYangName() string { return "lsp-arrival-time" }

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) SetParent(parent types.Entity) { lspArrivalTime.parent = parent }

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetParent() types.Entity { return lspArrivalTime.parent }

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetParentYangName() string { return "lsp-arrival-times" }

// Isis_Instances_Instance_TraceBufferSize
// Trace buffer size configuration
type Isis_Instances_Instance_TraceBufferSize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Buffer size for detailed traces. The type is interface{} with range:
    // 1..1000000.
    Detailed interface{}

    // Buffer size for standard traces. The type is interface{} with range:
    // 1..1000000.
    Standard interface{}

    // Buffer size for severe trace. The type is interface{} with range:
    // 1..1000000.
    Severe interface{}

    // Buffer size for hello trace. The type is interface{} with range:
    // 1..1000000.
    Hello interface{}
}

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetFilter() yfilter.YFilter { return traceBufferSize.YFilter }

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) SetFilter(yf yfilter.YFilter) { traceBufferSize.YFilter = yf }

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetGoName(yname string) string {
    if yname == "detailed" { return "Detailed" }
    if yname == "standard" { return "Standard" }
    if yname == "severe" { return "Severe" }
    if yname == "hello" { return "Hello" }
    return ""
}

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetSegmentPath() string {
    return "trace-buffer-size"
}

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["detailed"] = traceBufferSize.Detailed
    leafs["standard"] = traceBufferSize.Standard
    leafs["severe"] = traceBufferSize.Severe
    leafs["hello"] = traceBufferSize.Hello
    return leafs
}

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetBundleName() string { return "cisco_ios_xr" }

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetYangName() string { return "trace-buffer-size" }

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) SetParent(parent types.Entity) { traceBufferSize.parent = parent }

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetParent() types.Entity { return traceBufferSize.parent }

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_MaxLinkMetrics
// Max Link Metric configuration
type Isis_Instances_Instance_MaxLinkMetrics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max Link Metric. The type is slice of
    // Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric.
    MaxLinkMetric []Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric
}

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetFilter() yfilter.YFilter { return maxLinkMetrics.YFilter }

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) SetFilter(yf yfilter.YFilter) { maxLinkMetrics.YFilter = yf }

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetGoName(yname string) string {
    if yname == "max-link-metric" { return "MaxLinkMetric" }
    return ""
}

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetSegmentPath() string {
    return "max-link-metrics"
}

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "max-link-metric" {
        for _, c := range maxLinkMetrics.MaxLinkMetric {
            if maxLinkMetrics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric{}
        maxLinkMetrics.MaxLinkMetric = append(maxLinkMetrics.MaxLinkMetric, child)
        return &maxLinkMetrics.MaxLinkMetric[len(maxLinkMetrics.MaxLinkMetric)-1]
    }
    return nil
}

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range maxLinkMetrics.MaxLinkMetric {
        children[maxLinkMetrics.MaxLinkMetric[i].GetSegmentPath()] = &maxLinkMetrics.MaxLinkMetric[i]
    }
    return children
}

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetBundleName() string { return "cisco_ios_xr" }

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetYangName() string { return "max-link-metrics" }

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) SetParent(parent types.Entity) { maxLinkMetrics.parent = parent }

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetParent() types.Entity { return maxLinkMetrics.parent }

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric
// Max Link Metric
type Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetFilter() yfilter.YFilter { return maxLinkMetric.YFilter }

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) SetFilter(yf yfilter.YFilter) { maxLinkMetric.YFilter = yf }

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    return ""
}

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetSegmentPath() string {
    return "max-link-metric" + "[level='" + fmt.Sprintf("%v", maxLinkMetric.Level) + "']"
}

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = maxLinkMetric.Level
    return leafs
}

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetBundleName() string { return "cisco_ios_xr" }

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetYangName() string { return "max-link-metric" }

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) SetParent(parent types.Entity) { maxLinkMetric.parent = parent }

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetParent() types.Entity { return maxLinkMetric.parent }

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetParentYangName() string { return "max-link-metrics" }

// Isis_Instances_Instance_AdjacencyStagger
// Stagger ISIS adjacency bring up
// This type is a presence type.
type Isis_Instances_Instance_AdjacencyStagger struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Adjacency Stagger: Initial number of neighbors to bring up per area. The
    // type is interface{} with range: 2..65000. The default value is 2.
    InitialNbr interface{}

    // Adjacency Stagger: Subsequent simultaneous number of neighbors to bring up.
    // The type is interface{} with range: 2..65000. The default value is 64.
    MaxNbr interface{}
}

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetFilter() yfilter.YFilter { return adjacencyStagger.YFilter }

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) SetFilter(yf yfilter.YFilter) { adjacencyStagger.YFilter = yf }

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetGoName(yname string) string {
    if yname == "initial-nbr" { return "InitialNbr" }
    if yname == "max-nbr" { return "MaxNbr" }
    return ""
}

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetSegmentPath() string {
    return "adjacency-stagger"
}

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["initial-nbr"] = adjacencyStagger.InitialNbr
    leafs["max-nbr"] = adjacencyStagger.MaxNbr
    return leafs
}

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetBundleName() string { return "cisco_ios_xr" }

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetYangName() string { return "adjacency-stagger" }

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) SetParent(parent types.Entity) { adjacencyStagger.parent = parent }

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetParent() types.Entity { return adjacencyStagger.parent }

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_Afs
// Per-address-family configuration
type Isis_Instances_Instance_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for an IS-IS address-family. If a named (non-default)
    // topology is being created it must be multicast. The type is slice of
    // Isis_Instances_Instance_Afs_Af.
    Af []Isis_Instances_Instance_Afs_Af
}

func (afs *Isis_Instances_Instance_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *Isis_Instances_Instance_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *Isis_Instances_Instance_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *Isis_Instances_Instance_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *Isis_Instances_Instance_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *Isis_Instances_Instance_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *Isis_Instances_Instance_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *Isis_Instances_Instance_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *Isis_Instances_Instance_Afs) GetYangName() string { return "afs" }

func (afs *Isis_Instances_Instance_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *Isis_Instances_Instance_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *Isis_Instances_Instance_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *Isis_Instances_Instance_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *Isis_Instances_Instance_Afs) GetParent() types.Entity { return afs.parent }

func (afs *Isis_Instances_Instance_Afs) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_Afs_Af
// Configuration for an IS-IS address-family. If
// a named (non-default) topology is being
// created it must be multicast.
type Isis_Instances_Instance_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address family. The type is IsisAddressFamily.
    AfName interface{}

    // This attribute is a key. Sub address family. The type is
    // IsisSubAddressFamily.
    SafName interface{}

    // Data container.
    AfData Isis_Instances_Instance_Afs_Af_AfData

    // keys: topology-name. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName.
    TopologyName []Isis_Instances_Instance_Afs_Af_TopologyName
}

func (af *Isis_Instances_Instance_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *Isis_Instances_Instance_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *Isis_Instances_Instance_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "af-data" { return "AfData" }
    if yname == "topology-name" { return "TopologyName" }
    return ""
}

func (af *Isis_Instances_Instance_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']" + "[saf-name='" + fmt.Sprintf("%v", af.SafName) + "']"
}

func (af *Isis_Instances_Instance_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af-data" {
        return &af.AfData
    }
    if childYangName == "topology-name" {
        for _, c := range af.TopologyName {
            if af.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName{}
        af.TopologyName = append(af.TopologyName, child)
        return &af.TopologyName[len(af.TopologyName)-1]
    }
    return nil
}

func (af *Isis_Instances_Instance_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["af-data"] = &af.AfData
    for i := range af.TopologyName {
        children[af.TopologyName[i].GetSegmentPath()] = &af.TopologyName[i]
    }
    return children
}

func (af *Isis_Instances_Instance_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["saf-name"] = af.SafName
    return leafs
}

func (af *Isis_Instances_Instance_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *Isis_Instances_Instance_Afs_Af) GetYangName() string { return "af" }

func (af *Isis_Instances_Instance_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *Isis_Instances_Instance_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *Isis_Instances_Instance_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *Isis_Instances_Instance_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *Isis_Instances_Instance_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *Isis_Instances_Instance_Afs_Af) GetParentYangName() string { return "afs" }

// Isis_Instances_Instance_Afs_Af_AfData
// Data container.
// This type is a presence type.
type Isis_Instances_Instance_Afs_Af_AfData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of active parallel paths per route. The type is interface{}
    // with range: 1..64.
    MaximumPaths interface{}

    // Set the topology ID for a named (non-default) topology. This object must be
    // set before any other configuration is supplied for a named (non-default)
    // topology , and must be the last configuration object to be removed. This
    // item should not be supplied for the non-named default topologies. The type
    // is interface{} with range: 6..4095.
    TopologyId interface{}

    // Run IPv6 Unicast using the standard (IPv4 Unicast) topology. The type is
    // interface{}.
    SingleTopology interface{}

    // Suppress check for consistent AF support on received IIHs. The type is
    // IsisAdjCheck.
    AdjacencyCheck interface{}

    // If TRUE, advertise additional link attributes in our LSP. The type is bool.
    AdvertiseLinkAttributes interface{}

    // Apply weights to UCMP or ECMP only. The type is IsisApplyWeight.
    ApplyWeight interface{}

    // Default IS-IS administrative distance configuration. The type is
    // interface{} with range: 1..255. The default value is 115.
    DefaultAdminDistance interface{}

    // If enabled, advertise prefixes of passive interfaces only. The type is
    // interface{}.
    AdvertisePassiveOnly interface{}

    // If TRUE, Ignore other routers attached bit. The type is bool.
    IgnoreAttachedBit interface{}

    // Set the attached bit in this router's level 1 System LSP. The type is
    // IsisAttachedBit. The default value is area.
    AttachedBit interface{}

    // If TRUE, routes will be installed with the IP address of the first-hop node
    // as the source instead of the originating node. The type is bool.
    RouteSourceFirstHop interface{}

    // Enable Segment Routing configuration.
    SegmentRouting Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting

    // Metric-style configuration.
    MetricStyles Isis_Instances_Instance_Afs_Af_AfData_MetricStyles

    // Fast-ReRoute configuration.
    FrrTable Isis_Instances_Instance_Afs_Af_AfData_FrrTable

    // Stable IP address for system. Will only be applied for the unicast
    // sub-address-family.
    RouterId Isis_Instances_Instance_Afs_Af_AfData_RouterId

    // SPF Prefix Priority configuration.
    SpfPrefixPriorities Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities

    // Summary-prefix configuration.
    SummaryPrefixes Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes

    // Micro Loop Avoidance configuration.
    MicroLoopAvoidance Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance

    // UCMP (UnEqual Cost MultiPath) configuration.
    Ucmp Isis_Instances_Instance_Afs_Af_AfData_Ucmp

    // Maximum number of redistributed prefixesconfiguration.
    MaxRedistPrefixes Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes

    // Route propagation configuration.
    Propagations Isis_Instances_Instance_Afs_Af_AfData_Propagations

    // Protocol redistribution configuration.
    Redistributions Isis_Instances_Instance_Afs_Af_AfData_Redistributions

    // Peoridic SPF configuration.
    SpfPeriodicIntervals Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals

    // SPF-interval configuration.
    SpfIntervals Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals

    // Enable convergence monitoring.
    MonitorConvergence Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence

    // Control origination of a default route with the option of using a policy. 
    // If no policy is specified the default route is advertised with zero cost in
    // level 2 only.
    DefaultInformation Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation

    // Per-route administrative distanceconfiguration.
    AdminDistances Isis_Instances_Instance_Afs_Af_AfData_AdminDistances

    // ISPF configuration.
    Ispf Isis_Instances_Instance_Afs_Af_AfData_Ispf

    // MPLS LDP configuration. MPLS LDP configuration will only be applied for the
    // IPv4-unicast address-family.
    MplsLdpGlobal Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal

    // MPLS configuration. MPLS configuration will only be applied for the
    // IPv4-unicast address-family.
    Mpls Isis_Instances_Instance_Afs_Af_AfData_Mpls

    // Manual Adjacecy SID configuration.
    ManualAdjSids Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids

    // Metric configuration.
    Metrics Isis_Instances_Instance_Afs_Af_AfData_Metrics

    // Weight configuration.
    Weights Isis_Instances_Instance_Afs_Af_AfData_Weights
}

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetFilter() yfilter.YFilter { return afData.YFilter }

func (afData *Isis_Instances_Instance_Afs_Af_AfData) SetFilter(yf yfilter.YFilter) { afData.YFilter = yf }

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetGoName(yname string) string {
    if yname == "maximum-paths" { return "MaximumPaths" }
    if yname == "topology-id" { return "TopologyId" }
    if yname == "single-topology" { return "SingleTopology" }
    if yname == "adjacency-check" { return "AdjacencyCheck" }
    if yname == "advertise-link-attributes" { return "AdvertiseLinkAttributes" }
    if yname == "apply-weight" { return "ApplyWeight" }
    if yname == "default-admin-distance" { return "DefaultAdminDistance" }
    if yname == "advertise-passive-only" { return "AdvertisePassiveOnly" }
    if yname == "ignore-attached-bit" { return "IgnoreAttachedBit" }
    if yname == "attached-bit" { return "AttachedBit" }
    if yname == "route-source-first-hop" { return "RouteSourceFirstHop" }
    if yname == "segment-routing" { return "SegmentRouting" }
    if yname == "metric-styles" { return "MetricStyles" }
    if yname == "frr-table" { return "FrrTable" }
    if yname == "router-id" { return "RouterId" }
    if yname == "spf-prefix-priorities" { return "SpfPrefixPriorities" }
    if yname == "summary-prefixes" { return "SummaryPrefixes" }
    if yname == "micro-loop-avoidance" { return "MicroLoopAvoidance" }
    if yname == "ucmp" { return "Ucmp" }
    if yname == "max-redist-prefixes" { return "MaxRedistPrefixes" }
    if yname == "propagations" { return "Propagations" }
    if yname == "redistributions" { return "Redistributions" }
    if yname == "spf-periodic-intervals" { return "SpfPeriodicIntervals" }
    if yname == "spf-intervals" { return "SpfIntervals" }
    if yname == "monitor-convergence" { return "MonitorConvergence" }
    if yname == "default-information" { return "DefaultInformation" }
    if yname == "admin-distances" { return "AdminDistances" }
    if yname == "ispf" { return "Ispf" }
    if yname == "mpls-ldp-global" { return "MplsLdpGlobal" }
    if yname == "mpls" { return "Mpls" }
    if yname == "manual-adj-sids" { return "ManualAdjSids" }
    if yname == "metrics" { return "Metrics" }
    if yname == "weights" { return "Weights" }
    return ""
}

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetSegmentPath() string {
    return "af-data"
}

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "segment-routing" {
        return &afData.SegmentRouting
    }
    if childYangName == "metric-styles" {
        return &afData.MetricStyles
    }
    if childYangName == "frr-table" {
        return &afData.FrrTable
    }
    if childYangName == "router-id" {
        return &afData.RouterId
    }
    if childYangName == "spf-prefix-priorities" {
        return &afData.SpfPrefixPriorities
    }
    if childYangName == "summary-prefixes" {
        return &afData.SummaryPrefixes
    }
    if childYangName == "micro-loop-avoidance" {
        return &afData.MicroLoopAvoidance
    }
    if childYangName == "ucmp" {
        return &afData.Ucmp
    }
    if childYangName == "max-redist-prefixes" {
        return &afData.MaxRedistPrefixes
    }
    if childYangName == "propagations" {
        return &afData.Propagations
    }
    if childYangName == "redistributions" {
        return &afData.Redistributions
    }
    if childYangName == "spf-periodic-intervals" {
        return &afData.SpfPeriodicIntervals
    }
    if childYangName == "spf-intervals" {
        return &afData.SpfIntervals
    }
    if childYangName == "monitor-convergence" {
        return &afData.MonitorConvergence
    }
    if childYangName == "default-information" {
        return &afData.DefaultInformation
    }
    if childYangName == "admin-distances" {
        return &afData.AdminDistances
    }
    if childYangName == "ispf" {
        return &afData.Ispf
    }
    if childYangName == "mpls-ldp-global" {
        return &afData.MplsLdpGlobal
    }
    if childYangName == "mpls" {
        return &afData.Mpls
    }
    if childYangName == "manual-adj-sids" {
        return &afData.ManualAdjSids
    }
    if childYangName == "metrics" {
        return &afData.Metrics
    }
    if childYangName == "weights" {
        return &afData.Weights
    }
    return nil
}

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["segment-routing"] = &afData.SegmentRouting
    children["metric-styles"] = &afData.MetricStyles
    children["frr-table"] = &afData.FrrTable
    children["router-id"] = &afData.RouterId
    children["spf-prefix-priorities"] = &afData.SpfPrefixPriorities
    children["summary-prefixes"] = &afData.SummaryPrefixes
    children["micro-loop-avoidance"] = &afData.MicroLoopAvoidance
    children["ucmp"] = &afData.Ucmp
    children["max-redist-prefixes"] = &afData.MaxRedistPrefixes
    children["propagations"] = &afData.Propagations
    children["redistributions"] = &afData.Redistributions
    children["spf-periodic-intervals"] = &afData.SpfPeriodicIntervals
    children["spf-intervals"] = &afData.SpfIntervals
    children["monitor-convergence"] = &afData.MonitorConvergence
    children["default-information"] = &afData.DefaultInformation
    children["admin-distances"] = &afData.AdminDistances
    children["ispf"] = &afData.Ispf
    children["mpls-ldp-global"] = &afData.MplsLdpGlobal
    children["mpls"] = &afData.Mpls
    children["manual-adj-sids"] = &afData.ManualAdjSids
    children["metrics"] = &afData.Metrics
    children["weights"] = &afData.Weights
    return children
}

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-paths"] = afData.MaximumPaths
    leafs["topology-id"] = afData.TopologyId
    leafs["single-topology"] = afData.SingleTopology
    leafs["adjacency-check"] = afData.AdjacencyCheck
    leafs["advertise-link-attributes"] = afData.AdvertiseLinkAttributes
    leafs["apply-weight"] = afData.ApplyWeight
    leafs["default-admin-distance"] = afData.DefaultAdminDistance
    leafs["advertise-passive-only"] = afData.AdvertisePassiveOnly
    leafs["ignore-attached-bit"] = afData.IgnoreAttachedBit
    leafs["attached-bit"] = afData.AttachedBit
    leafs["route-source-first-hop"] = afData.RouteSourceFirstHop
    return leafs
}

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetBundleName() string { return "cisco_ios_xr" }

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetYangName() string { return "af-data" }

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afData *Isis_Instances_Instance_Afs_Af_AfData) SetParent(parent types.Entity) { afData.parent = parent }

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetParent() types.Entity { return afData.parent }

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetParentYangName() string { return "af" }

// Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting
// Enable Segment Routing configuration
type Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable per bundle member adjacency SID. The type is interface{}.
    BundleMemberAdjSid interface{}

    // Prefer segment routing labels over LDP labels. The type is
    // IsisLabelPreference.
    Mpls interface{}

    // Enable Segment Routing prefix SID map configuration.
    PrefixSidMap Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetFilter() yfilter.YFilter { return segmentRouting.YFilter }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) SetFilter(yf yfilter.YFilter) { segmentRouting.YFilter = yf }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetGoName(yname string) string {
    if yname == "bundle-member-adj-sid" { return "BundleMemberAdjSid" }
    if yname == "mpls" { return "Mpls" }
    if yname == "prefix-sid-map" { return "PrefixSidMap" }
    return ""
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetSegmentPath() string {
    return "segment-routing"
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-sid-map" {
        return &segmentRouting.PrefixSidMap
    }
    return nil
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-sid-map"] = &segmentRouting.PrefixSidMap
    return children
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bundle-member-adj-sid"] = segmentRouting.BundleMemberAdjSid
    leafs["mpls"] = segmentRouting.Mpls
    return leafs
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetBundleName() string { return "cisco_ios_xr" }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetYangName() string { return "segment-routing" }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) SetParent(parent types.Entity) { segmentRouting.parent = parent }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetParent() types.Entity { return segmentRouting.parent }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap
// Enable Segment Routing prefix SID map
// configuration
type Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Segment Routing prefix SID map advertise local. The type is
    // interface{}.
    AdvertiseLocal interface{}

    // If TRUE, remote prefix SID map advertisements will be used. If FALSE, they
    // will not be used. The type is bool.
    Receive interface{}
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetFilter() yfilter.YFilter { return prefixSidMap.YFilter }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) SetFilter(yf yfilter.YFilter) { prefixSidMap.YFilter = yf }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetGoName(yname string) string {
    if yname == "advertise-local" { return "AdvertiseLocal" }
    if yname == "receive" { return "Receive" }
    return ""
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetSegmentPath() string {
    return "prefix-sid-map"
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["advertise-local"] = prefixSidMap.AdvertiseLocal
    leafs["receive"] = prefixSidMap.Receive
    return leafs
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetYangName() string { return "prefix-sid-map" }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) SetParent(parent types.Entity) { prefixSidMap.parent = parent }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetParent() types.Entity { return prefixSidMap.parent }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetParentYangName() string { return "segment-routing" }

// Isis_Instances_Instance_Afs_Af_AfData_MetricStyles
// Metric-style configuration
type Isis_Instances_Instance_Afs_Af_AfData_MetricStyles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of metric style in LSPs. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle.
    MetricStyle []Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetFilter() yfilter.YFilter { return metricStyles.YFilter }

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) SetFilter(yf yfilter.YFilter) { metricStyles.YFilter = yf }

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetGoName(yname string) string {
    if yname == "metric-style" { return "MetricStyle" }
    return ""
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetSegmentPath() string {
    return "metric-styles"
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metric-style" {
        for _, c := range metricStyles.MetricStyle {
            if metricStyles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle{}
        metricStyles.MetricStyle = append(metricStyles.MetricStyle, child)
        return &metricStyles.MetricStyle[len(metricStyles.MetricStyle)-1]
    }
    return nil
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range metricStyles.MetricStyle {
        children[metricStyles.MetricStyle[i].GetSegmentPath()] = &metricStyles.MetricStyle[i]
    }
    return children
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetBundleName() string { return "cisco_ios_xr" }

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetYangName() string { return "metric-styles" }

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) SetParent(parent types.Entity) { metricStyles.parent = parent }

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetParent() types.Entity { return metricStyles.parent }

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle
// Configuration of metric style in LSPs
type Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Metric Style. The type is IsisMetricStyle. The default value is
    // old-metric-style.
    Style interface{}

    // Transition state. The type is IsisMetricStyleTransition. The default value
    // is disabled.
    TransitionState interface{}
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetFilter() yfilter.YFilter { return metricStyle.YFilter }

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) SetFilter(yf yfilter.YFilter) { metricStyle.YFilter = yf }

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "style" { return "Style" }
    if yname == "transition-state" { return "TransitionState" }
    return ""
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetSegmentPath() string {
    return "metric-style" + "[level='" + fmt.Sprintf("%v", metricStyle.Level) + "']"
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = metricStyle.Level
    leafs["style"] = metricStyle.Style
    leafs["transition-state"] = metricStyle.TransitionState
    return leafs
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetBundleName() string { return "cisco_ios_xr" }

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetYangName() string { return "metric-style" }

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) SetParent(parent types.Entity) { metricStyle.parent = parent }

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetParent() types.Entity { return metricStyle.parent }

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetParentYangName() string { return "metric-styles" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable
// Fast-ReRoute configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Load share prefixes across multiple backups.
    FrrLoadSharings Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings

    // FRR prefix-limit configuration.
    PriorityLimits Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits

    // FRR remote LFA prefix list filter configuration.
    FrrRemoteLfaPrefixes Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes

    // FRR tiebreakers configuration.
    FrrTiebreakers Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers

    // FRR use candidate only configuration.
    FrrUseCandOnlies Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies
}

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetFilter() yfilter.YFilter { return frrTable.YFilter }

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) SetFilter(yf yfilter.YFilter) { frrTable.YFilter = yf }

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetGoName(yname string) string {
    if yname == "frr-load-sharings" { return "FrrLoadSharings" }
    if yname == "priority-limits" { return "PriorityLimits" }
    if yname == "frr-remote-lfa-prefixes" { return "FrrRemoteLfaPrefixes" }
    if yname == "frr-tiebreakers" { return "FrrTiebreakers" }
    if yname == "frr-use-cand-onlies" { return "FrrUseCandOnlies" }
    return ""
}

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetSegmentPath() string {
    return "frr-table"
}

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-load-sharings" {
        return &frrTable.FrrLoadSharings
    }
    if childYangName == "priority-limits" {
        return &frrTable.PriorityLimits
    }
    if childYangName == "frr-remote-lfa-prefixes" {
        return &frrTable.FrrRemoteLfaPrefixes
    }
    if childYangName == "frr-tiebreakers" {
        return &frrTable.FrrTiebreakers
    }
    if childYangName == "frr-use-cand-onlies" {
        return &frrTable.FrrUseCandOnlies
    }
    return nil
}

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frr-load-sharings"] = &frrTable.FrrLoadSharings
    children["priority-limits"] = &frrTable.PriorityLimits
    children["frr-remote-lfa-prefixes"] = &frrTable.FrrRemoteLfaPrefixes
    children["frr-tiebreakers"] = &frrTable.FrrTiebreakers
    children["frr-use-cand-onlies"] = &frrTable.FrrUseCandOnlies
    return children
}

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetBundleName() string { return "cisco_ios_xr" }

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetYangName() string { return "frr-table" }

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) SetParent(parent types.Entity) { frrTable.parent = parent }

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetParent() types.Entity { return frrTable.parent }

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings
// Load share prefixes across multiple
// backups
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable load sharing. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing.
    FrrLoadSharing []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetFilter() yfilter.YFilter { return frrLoadSharings.YFilter }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) SetFilter(yf yfilter.YFilter) { frrLoadSharings.YFilter = yf }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetGoName(yname string) string {
    if yname == "frr-load-sharing" { return "FrrLoadSharing" }
    return ""
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetSegmentPath() string {
    return "frr-load-sharings"
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-load-sharing" {
        for _, c := range frrLoadSharings.FrrLoadSharing {
            if frrLoadSharings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing{}
        frrLoadSharings.FrrLoadSharing = append(frrLoadSharings.FrrLoadSharing, child)
        return &frrLoadSharings.FrrLoadSharing[len(frrLoadSharings.FrrLoadSharing)-1]
    }
    return nil
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrLoadSharings.FrrLoadSharing {
        children[frrLoadSharings.FrrLoadSharing[i].GetSegmentPath()] = &frrLoadSharings.FrrLoadSharing[i]
    }
    return children
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetBundleName() string { return "cisco_ios_xr" }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetYangName() string { return "frr-load-sharings" }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) SetParent(parent types.Entity) { frrLoadSharings.parent = parent }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetParent() types.Entity { return frrLoadSharings.parent }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing
// Disable load sharing
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Load sharing. The type is IsisfrrLoadSharing. This attribute is mandatory.
    LoadSharing interface{}
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetFilter() yfilter.YFilter { return frrLoadSharing.YFilter }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) SetFilter(yf yfilter.YFilter) { frrLoadSharing.YFilter = yf }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "load-sharing" { return "LoadSharing" }
    return ""
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetSegmentPath() string {
    return "frr-load-sharing" + "[level='" + fmt.Sprintf("%v", frrLoadSharing.Level) + "']"
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrLoadSharing.Level
    leafs["load-sharing"] = frrLoadSharing.LoadSharing
    return leafs
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetBundleName() string { return "cisco_ios_xr" }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetYangName() string { return "frr-load-sharing" }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) SetParent(parent types.Entity) { frrLoadSharing.parent = parent }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetParent() types.Entity { return frrLoadSharing.parent }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetParentYangName() string { return "frr-load-sharings" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits
// FRR prefix-limit configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit backup computation upto the prefix priority. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit.
    PriorityLimit []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetFilter() yfilter.YFilter { return priorityLimits.YFilter }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) SetFilter(yf yfilter.YFilter) { priorityLimits.YFilter = yf }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetGoName(yname string) string {
    if yname == "priority-limit" { return "PriorityLimit" }
    return ""
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetSegmentPath() string {
    return "priority-limits"
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "priority-limit" {
        for _, c := range priorityLimits.PriorityLimit {
            if priorityLimits.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit{}
        priorityLimits.PriorityLimit = append(priorityLimits.PriorityLimit, child)
        return &priorityLimits.PriorityLimit[len(priorityLimits.PriorityLimit)-1]
    }
    return nil
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range priorityLimits.PriorityLimit {
        children[priorityLimits.PriorityLimit[i].GetSegmentPath()] = &priorityLimits.PriorityLimit[i]
    }
    return children
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetBundleName() string { return "cisco_ios_xr" }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetYangName() string { return "priority-limits" }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) SetParent(parent types.Entity) { priorityLimits.parent = parent }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetParent() types.Entity { return priorityLimits.parent }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit
// Limit backup computation upto the prefix
// priority
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Compute for all prefixes upto the specified priority. The type is
    // IsisPrefixPriority. This attribute is mandatory.
    Priority interface{}
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetFilter() yfilter.YFilter { return priorityLimit.YFilter }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) SetFilter(yf yfilter.YFilter) { priorityLimit.YFilter = yf }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "frr-type" { return "FrrType" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetSegmentPath() string {
    return "priority-limit" + "[level='" + fmt.Sprintf("%v", priorityLimit.Level) + "']" + "[frr-type='" + fmt.Sprintf("%v", priorityLimit.FrrType) + "']"
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = priorityLimit.Level
    leafs["frr-type"] = priorityLimit.FrrType
    leafs["priority"] = priorityLimit.Priority
    return leafs
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetBundleName() string { return "cisco_ios_xr" }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetYangName() string { return "priority-limit" }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) SetParent(parent types.Entity) { priorityLimit.parent = parent }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetParent() types.Entity { return priorityLimit.parent }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetParentYangName() string { return "priority-limits" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes
// FRR remote LFA prefix list filter
// configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Filter remote LFA router IDs using prefix-list. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix.
    FrrRemoteLfaPrefix []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetFilter() yfilter.YFilter { return frrRemoteLfaPrefixes.YFilter }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) SetFilter(yf yfilter.YFilter) { frrRemoteLfaPrefixes.YFilter = yf }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetGoName(yname string) string {
    if yname == "frr-remote-lfa-prefix" { return "FrrRemoteLfaPrefix" }
    return ""
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetSegmentPath() string {
    return "frr-remote-lfa-prefixes"
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-remote-lfa-prefix" {
        for _, c := range frrRemoteLfaPrefixes.FrrRemoteLfaPrefix {
            if frrRemoteLfaPrefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix{}
        frrRemoteLfaPrefixes.FrrRemoteLfaPrefix = append(frrRemoteLfaPrefixes.FrrRemoteLfaPrefix, child)
        return &frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[len(frrRemoteLfaPrefixes.FrrRemoteLfaPrefix)-1]
    }
    return nil
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrRemoteLfaPrefixes.FrrRemoteLfaPrefix {
        children[frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[i].GetSegmentPath()] = &frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[i]
    }
    return children
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetYangName() string { return "frr-remote-lfa-prefixes" }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) SetParent(parent types.Entity) { frrRemoteLfaPrefixes.parent = parent }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetParent() types.Entity { return frrRemoteLfaPrefixes.parent }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix
// Filter remote LFA router IDs using
// prefix-list
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Name of the prefix list. The type is string with length: 1..32. This
    // attribute is mandatory.
    PrefixListName interface{}
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetFilter() yfilter.YFilter { return frrRemoteLfaPrefix.YFilter }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) SetFilter(yf yfilter.YFilter) { frrRemoteLfaPrefix.YFilter = yf }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "prefix-list-name" { return "PrefixListName" }
    return ""
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetSegmentPath() string {
    return "frr-remote-lfa-prefix" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaPrefix.Level) + "']"
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrRemoteLfaPrefix.Level
    leafs["prefix-list-name"] = frrRemoteLfaPrefix.PrefixListName
    return leafs
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetYangName() string { return "frr-remote-lfa-prefix" }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) SetParent(parent types.Entity) { frrRemoteLfaPrefix.parent = parent }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetParent() types.Entity { return frrRemoteLfaPrefix.parent }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetParentYangName() string { return "frr-remote-lfa-prefixes" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers
// FRR tiebreakers configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure tiebreaker for multiple backups. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker.
    FrrTiebreaker []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetFilter() yfilter.YFilter { return frrTiebreakers.YFilter }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) SetFilter(yf yfilter.YFilter) { frrTiebreakers.YFilter = yf }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetGoName(yname string) string {
    if yname == "frr-tiebreaker" { return "FrrTiebreaker" }
    return ""
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetSegmentPath() string {
    return "frr-tiebreakers"
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-tiebreaker" {
        for _, c := range frrTiebreakers.FrrTiebreaker {
            if frrTiebreakers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker{}
        frrTiebreakers.FrrTiebreaker = append(frrTiebreakers.FrrTiebreaker, child)
        return &frrTiebreakers.FrrTiebreaker[len(frrTiebreakers.FrrTiebreaker)-1]
    }
    return nil
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrTiebreakers.FrrTiebreaker {
        children[frrTiebreakers.FrrTiebreaker[i].GetSegmentPath()] = &frrTiebreakers.FrrTiebreaker[i]
    }
    return children
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetBundleName() string { return "cisco_ios_xr" }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetYangName() string { return "frr-tiebreakers" }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) SetParent(parent types.Entity) { frrTiebreakers.parent = parent }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetParent() types.Entity { return frrTiebreakers.parent }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker
// Configure tiebreaker for multiple backups
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Tiebreaker for which configuration applies. The
    // type is IsisfrrTiebreaker.
    Tiebreaker interface{}

    // Preference order among tiebreakers. The type is interface{} with range:
    // 1..255. This attribute is mandatory.
    Index interface{}
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetFilter() yfilter.YFilter { return frrTiebreaker.YFilter }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) SetFilter(yf yfilter.YFilter) { frrTiebreaker.YFilter = yf }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "tiebreaker" { return "Tiebreaker" }
    if yname == "index" { return "Index" }
    return ""
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetSegmentPath() string {
    return "frr-tiebreaker" + "[level='" + fmt.Sprintf("%v", frrTiebreaker.Level) + "']" + "[tiebreaker='" + fmt.Sprintf("%v", frrTiebreaker.Tiebreaker) + "']"
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrTiebreaker.Level
    leafs["tiebreaker"] = frrTiebreaker.Tiebreaker
    leafs["index"] = frrTiebreaker.Index
    return leafs
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetBundleName() string { return "cisco_ios_xr" }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetYangName() string { return "frr-tiebreaker" }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) SetParent(parent types.Entity) { frrTiebreaker.parent = parent }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetParent() types.Entity { return frrTiebreaker.parent }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetParentYangName() string { return "frr-tiebreakers" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies
// FRR use candidate only configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure use candidate only to exclude interfaces as backup. The type is
    // slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly.
    FrrUseCandOnly []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetFilter() yfilter.YFilter { return frrUseCandOnlies.YFilter }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) SetFilter(yf yfilter.YFilter) { frrUseCandOnlies.YFilter = yf }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetGoName(yname string) string {
    if yname == "frr-use-cand-only" { return "FrrUseCandOnly" }
    return ""
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetSegmentPath() string {
    return "frr-use-cand-onlies"
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-use-cand-only" {
        for _, c := range frrUseCandOnlies.FrrUseCandOnly {
            if frrUseCandOnlies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly{}
        frrUseCandOnlies.FrrUseCandOnly = append(frrUseCandOnlies.FrrUseCandOnly, child)
        return &frrUseCandOnlies.FrrUseCandOnly[len(frrUseCandOnlies.FrrUseCandOnly)-1]
    }
    return nil
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrUseCandOnlies.FrrUseCandOnly {
        children[frrUseCandOnlies.FrrUseCandOnly[i].GetSegmentPath()] = &frrUseCandOnlies.FrrUseCandOnly[i]
    }
    return children
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetBundleName() string { return "cisco_ios_xr" }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetYangName() string { return "frr-use-cand-onlies" }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) SetParent(parent types.Entity) { frrUseCandOnlies.parent = parent }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetParent() types.Entity { return frrUseCandOnlies.parent }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly
// Configure use candidate only to exclude
// interfaces as backup
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetFilter() yfilter.YFilter { return frrUseCandOnly.YFilter }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) SetFilter(yf yfilter.YFilter) { frrUseCandOnly.YFilter = yf }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "frr-type" { return "FrrType" }
    return ""
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetSegmentPath() string {
    return "frr-use-cand-only" + "[level='" + fmt.Sprintf("%v", frrUseCandOnly.Level) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrUseCandOnly.FrrType) + "']"
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrUseCandOnly.Level
    leafs["frr-type"] = frrUseCandOnly.FrrType
    return leafs
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetBundleName() string { return "cisco_ios_xr" }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetYangName() string { return "frr-use-cand-only" }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) SetParent(parent types.Entity) { frrUseCandOnly.parent = parent }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetParent() types.Entity { return frrUseCandOnly.parent }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetParentYangName() string { return "frr-use-cand-onlies" }

// Isis_Instances_Instance_Afs_Af_AfData_RouterId
// Stable IP address for system. Will only be
// applied for the unicast sub-address-family.
type Isis_Instances_Instance_Afs_Af_AfData_RouterId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4/IPv6 address to be used as a router ID. Precisely one of Address and
    // Interface must be specified. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Interface with designated stable IP address to be used as a router ID. This
    // must be a Loopback interface. Precisely one of Address and Interface must
    // be specified. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetFilter() yfilter.YFilter { return routerId.YFilter }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) SetFilter(yf yfilter.YFilter) { routerId.YFilter = yf }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetSegmentPath() string {
    return "router-id"
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = routerId.Address
    leafs["interface-name"] = routerId.InterfaceName
    return leafs
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetBundleName() string { return "cisco_ios_xr" }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetYangName() string { return "router-id" }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) SetParent(parent types.Entity) { routerId.parent = parent }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetParent() types.Entity { return routerId.parent }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities
// SPF Prefix Priority configuration
type Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Determine SPF priority for prefixes. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority.
    SpfPrefixPriority []Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetFilter() yfilter.YFilter { return spfPrefixPriorities.YFilter }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) SetFilter(yf yfilter.YFilter) { spfPrefixPriorities.YFilter = yf }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetGoName(yname string) string {
    if yname == "spf-prefix-priority" { return "SpfPrefixPriority" }
    return ""
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetSegmentPath() string {
    return "spf-prefix-priorities"
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spf-prefix-priority" {
        for _, c := range spfPrefixPriorities.SpfPrefixPriority {
            if spfPrefixPriorities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority{}
        spfPrefixPriorities.SpfPrefixPriority = append(spfPrefixPriorities.SpfPrefixPriority, child)
        return &spfPrefixPriorities.SpfPrefixPriority[len(spfPrefixPriorities.SpfPrefixPriority)-1]
    }
    return nil
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spfPrefixPriorities.SpfPrefixPriority {
        children[spfPrefixPriorities.SpfPrefixPriority[i].GetSegmentPath()] = &spfPrefixPriorities.SpfPrefixPriority[i]
    }
    return children
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetBundleName() string { return "cisco_ios_xr" }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetYangName() string { return "spf-prefix-priorities" }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) SetParent(parent types.Entity) { spfPrefixPriorities.parent = parent }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetParent() types.Entity { return spfPrefixPriorities.parent }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority
// Determine SPF priority for prefixes
type Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SPF Level for prefix prioritization. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. SPF Priority to assign matching prefixes. The type
    // is IsisPrefixPriority.
    PrefixPriorityType interface{}

    // Tag value to determine prefixes for this priority. The type is interface{}
    // with range: 1..4294967295.
    AdminTag interface{}

    // Access List to determine prefixes for this priority. The type is string
    // with length: 1..64.
    AccessListName interface{}
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetFilter() yfilter.YFilter { return spfPrefixPriority.YFilter }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) SetFilter(yf yfilter.YFilter) { spfPrefixPriority.YFilter = yf }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "prefix-priority-type" { return "PrefixPriorityType" }
    if yname == "admin-tag" { return "AdminTag" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetSegmentPath() string {
    return "spf-prefix-priority" + "[level='" + fmt.Sprintf("%v", spfPrefixPriority.Level) + "']" + "[prefix-priority-type='" + fmt.Sprintf("%v", spfPrefixPriority.PrefixPriorityType) + "']"
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = spfPrefixPriority.Level
    leafs["prefix-priority-type"] = spfPrefixPriority.PrefixPriorityType
    leafs["admin-tag"] = spfPrefixPriority.AdminTag
    leafs["access-list-name"] = spfPrefixPriority.AccessListName
    return leafs
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetBundleName() string { return "cisco_ios_xr" }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetYangName() string { return "spf-prefix-priority" }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) SetParent(parent types.Entity) { spfPrefixPriority.parent = parent }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetParent() types.Entity { return spfPrefixPriority.parent }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetParentYangName() string { return "spf-prefix-priorities" }

// Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes
// Summary-prefix configuration
type Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure IP address prefixes to advertise. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix.
    SummaryPrefix []Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetFilter() yfilter.YFilter { return summaryPrefixes.YFilter }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) SetFilter(yf yfilter.YFilter) { summaryPrefixes.YFilter = yf }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetGoName(yname string) string {
    if yname == "summary-prefix" { return "SummaryPrefix" }
    return ""
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetSegmentPath() string {
    return "summary-prefixes"
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary-prefix" {
        for _, c := range summaryPrefixes.SummaryPrefix {
            if summaryPrefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix{}
        summaryPrefixes.SummaryPrefix = append(summaryPrefixes.SummaryPrefix, child)
        return &summaryPrefixes.SummaryPrefix[len(summaryPrefixes.SummaryPrefix)-1]
    }
    return nil
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summaryPrefixes.SummaryPrefix {
        children[summaryPrefixes.SummaryPrefix[i].GetSegmentPath()] = &summaryPrefixes.SummaryPrefix[i]
    }
    return children
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetBundleName() string { return "cisco_ios_xr" }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetYangName() string { return "summary-prefixes" }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) SetParent(parent types.Entity) { summaryPrefixes.parent = parent }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetParent() types.Entity { return summaryPrefixes.parent }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix
// Configure IP address prefixes to advertise
type Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP summary address prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}

    // The tag value. The type is interface{} with range: 1..4294967295.
    Tag interface{}

    // Level in which to summarize routes. The type is interface{} with range:
    // 1..2.
    Level interface{}
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetFilter() yfilter.YFilter { return summaryPrefix.YFilter }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) SetFilter(yf yfilter.YFilter) { summaryPrefix.YFilter = yf }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetGoName(yname string) string {
    if yname == "address-prefix" { return "AddressPrefix" }
    if yname == "tag" { return "Tag" }
    if yname == "level" { return "Level" }
    return ""
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetSegmentPath() string {
    return "summary-prefix" + "[address-prefix='" + fmt.Sprintf("%v", summaryPrefix.AddressPrefix) + "']"
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-prefix"] = summaryPrefix.AddressPrefix
    leafs["tag"] = summaryPrefix.Tag
    leafs["level"] = summaryPrefix.Level
    return leafs
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetYangName() string { return "summary-prefix" }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) SetParent(parent types.Entity) { summaryPrefix.parent = parent }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetParent() types.Entity { return summaryPrefix.parent }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetParentYangName() string { return "summary-prefixes" }

// Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance
// Micro Loop Avoidance configuration
type Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MicroLoop avoidance enable configuration. The type is
    // IsisMicroLoopAvoidance. The default value is micro-loop-avoidance-all.
    Enable interface{}

    // Value of delay in msecs in updating RIB. The type is interface{} with
    // range: 1000..65535. Units are millisecond. The default value is 5000.
    RibUpdateDelay interface{}
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetFilter() yfilter.YFilter { return microLoopAvoidance.YFilter }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) SetFilter(yf yfilter.YFilter) { microLoopAvoidance.YFilter = yf }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "rib-update-delay" { return "RibUpdateDelay" }
    return ""
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetSegmentPath() string {
    return "micro-loop-avoidance"
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = microLoopAvoidance.Enable
    leafs["rib-update-delay"] = microLoopAvoidance.RibUpdateDelay
    return leafs
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetBundleName() string { return "cisco_ios_xr" }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetYangName() string { return "micro-loop-avoidance" }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) SetParent(parent types.Entity) { microLoopAvoidance.parent = parent }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetParent() types.Entity { return microLoopAvoidance.parent }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_Ucmp
// UCMP (UnEqual Cost MultiPath) configuration
type Isis_Instances_Instance_Afs_Af_AfData_Ucmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay in msecs between primary SPF and UCMP computation. The type is
    // interface{} with range: 100..65535. Units are millisecond. The default
    // value is 100.
    DelayInterval interface{}

    // UCMP feature enable configuration.
    Enable Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable

    // Interfaces excluded from UCMP path computation.
    ExcludeInterfaces Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces
}

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetFilter() yfilter.YFilter { return ucmp.YFilter }

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) SetFilter(yf yfilter.YFilter) { ucmp.YFilter = yf }

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetGoName(yname string) string {
    if yname == "delay-interval" { return "DelayInterval" }
    if yname == "enable" { return "Enable" }
    if yname == "exclude-interfaces" { return "ExcludeInterfaces" }
    return ""
}

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetSegmentPath() string {
    return "ucmp"
}

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "enable" {
        return &ucmp.Enable
    }
    if childYangName == "exclude-interfaces" {
        return &ucmp.ExcludeInterfaces
    }
    return nil
}

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["enable"] = &ucmp.Enable
    children["exclude-interfaces"] = &ucmp.ExcludeInterfaces
    return children
}

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["delay-interval"] = ucmp.DelayInterval
    return leafs
}

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetBundleName() string { return "cisco_ios_xr" }

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetYangName() string { return "ucmp" }

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) SetParent(parent types.Entity) { ucmp.parent = parent }

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetParent() types.Entity { return ucmp.parent }

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable
// UCMP feature enable configuration
type Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Value of variance. The type is interface{} with range: 101..10000. The
    // default value is 200.
    Variance interface{}

    // Name of the Prefix List. The type is string with length: 1..32.
    PrefixListName interface{}
}

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetFilter() yfilter.YFilter { return enable.YFilter }

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) SetFilter(yf yfilter.YFilter) { enable.YFilter = yf }

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetGoName(yname string) string {
    if yname == "variance" { return "Variance" }
    if yname == "prefix-list-name" { return "PrefixListName" }
    return ""
}

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetSegmentPath() string {
    return "enable"
}

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["variance"] = enable.Variance
    leafs["prefix-list-name"] = enable.PrefixListName
    return leafs
}

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetBundleName() string { return "cisco_ios_xr" }

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetYangName() string { return "enable" }

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) SetParent(parent types.Entity) { enable.parent = parent }

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetParent() types.Entity { return enable.parent }

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetParentYangName() string { return "ucmp" }

// Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces
// Interfaces excluded from UCMP path
// computation
type Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exclude this interface from UCMP path computation. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetFilter() yfilter.YFilter { return excludeInterfaces.YFilter }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) SetFilter(yf yfilter.YFilter) { excludeInterfaces.YFilter = yf }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetGoName(yname string) string {
    if yname == "exclude-interface" { return "ExcludeInterface" }
    return ""
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetSegmentPath() string {
    return "exclude-interfaces"
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "exclude-interface" {
        for _, c := range excludeInterfaces.ExcludeInterface {
            if excludeInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface{}
        excludeInterfaces.ExcludeInterface = append(excludeInterfaces.ExcludeInterface, child)
        return &excludeInterfaces.ExcludeInterface[len(excludeInterfaces.ExcludeInterface)-1]
    }
    return nil
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range excludeInterfaces.ExcludeInterface {
        children[excludeInterfaces.ExcludeInterface[i].GetSegmentPath()] = &excludeInterfaces.ExcludeInterface[i]
    }
    return children
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetYangName() string { return "exclude-interfaces" }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) SetParent(parent types.Entity) { excludeInterfaces.parent = parent }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetParent() types.Entity { return excludeInterfaces.parent }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetParentYangName() string { return "ucmp" }

// Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface
// Exclude this interface from UCMP path
// computation
type Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface to be excluded. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetFilter() yfilter.YFilter { return excludeInterface.YFilter }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) SetFilter(yf yfilter.YFilter) { excludeInterface.YFilter = yf }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetSegmentPath() string {
    return "exclude-interface" + "[interface-name='" + fmt.Sprintf("%v", excludeInterface.InterfaceName) + "']"
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = excludeInterface.InterfaceName
    return leafs
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetBundleName() string { return "cisco_ios_xr" }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetYangName() string { return "exclude-interface" }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) SetParent(parent types.Entity) { excludeInterface.parent = parent }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetParent() types.Entity { return excludeInterface.parent }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetParentYangName() string { return "exclude-interfaces" }

// Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes
// Maximum number of redistributed
// prefixesconfiguration
type Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An upper limit on the number of redistributed prefixes which may be
    // included in the local system's LSP. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix.
    MaxRedistPrefix []Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetFilter() yfilter.YFilter { return maxRedistPrefixes.YFilter }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) SetFilter(yf yfilter.YFilter) { maxRedistPrefixes.YFilter = yf }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetGoName(yname string) string {
    if yname == "max-redist-prefix" { return "MaxRedistPrefix" }
    return ""
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetSegmentPath() string {
    return "max-redist-prefixes"
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "max-redist-prefix" {
        for _, c := range maxRedistPrefixes.MaxRedistPrefix {
            if maxRedistPrefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix{}
        maxRedistPrefixes.MaxRedistPrefix = append(maxRedistPrefixes.MaxRedistPrefix, child)
        return &maxRedistPrefixes.MaxRedistPrefix[len(maxRedistPrefixes.MaxRedistPrefix)-1]
    }
    return nil
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range maxRedistPrefixes.MaxRedistPrefix {
        children[maxRedistPrefixes.MaxRedistPrefix[i].GetSegmentPath()] = &maxRedistPrefixes.MaxRedistPrefix[i]
    }
    return children
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetBundleName() string { return "cisco_ios_xr" }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetYangName() string { return "max-redist-prefixes" }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) SetParent(parent types.Entity) { maxRedistPrefixes.parent = parent }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetParent() types.Entity { return maxRedistPrefixes.parent }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix
// An upper limit on the number of
// redistributed prefixes which may be
// included in the local system's LSP
type Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Max number of prefixes. The type is interface{} with range: 1..28000. This
    // attribute is mandatory.
    PrefixLimit interface{}
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetFilter() yfilter.YFilter { return maxRedistPrefix.YFilter }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) SetFilter(yf yfilter.YFilter) { maxRedistPrefix.YFilter = yf }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetSegmentPath() string {
    return "max-redist-prefix" + "[level='" + fmt.Sprintf("%v", maxRedistPrefix.Level) + "']"
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = maxRedistPrefix.Level
    leafs["prefix-limit"] = maxRedistPrefix.PrefixLimit
    return leafs
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetYangName() string { return "max-redist-prefix" }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) SetParent(parent types.Entity) { maxRedistPrefix.parent = parent }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetParent() types.Entity { return maxRedistPrefix.parent }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetParentYangName() string { return "max-redist-prefixes" }

// Isis_Instances_Instance_Afs_Af_AfData_Propagations
// Route propagation configuration
type Isis_Instances_Instance_Afs_Af_AfData_Propagations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Propagate routes between IS-IS levels. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation.
    Propagation []Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation
}

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetFilter() yfilter.YFilter { return propagations.YFilter }

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) SetFilter(yf yfilter.YFilter) { propagations.YFilter = yf }

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetGoName(yname string) string {
    if yname == "propagation" { return "Propagation" }
    return ""
}

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetSegmentPath() string {
    return "propagations"
}

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "propagation" {
        for _, c := range propagations.Propagation {
            if propagations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation{}
        propagations.Propagation = append(propagations.Propagation, child)
        return &propagations.Propagation[len(propagations.Propagation)-1]
    }
    return nil
}

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range propagations.Propagation {
        children[propagations.Propagation[i].GetSegmentPath()] = &propagations.Propagation[i]
    }
    return children
}

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetBundleName() string { return "cisco_ios_xr" }

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetYangName() string { return "propagations" }

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) SetParent(parent types.Entity) { propagations.parent = parent }

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetParent() types.Entity { return propagations.parent }

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation
// Propagate routes between IS-IS levels
type Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Source level for routes. The type is
    // IsisInternalLevel.
    SourceLevel interface{}

    // This attribute is a key. Destination level for routes.  Must differ from
    // SourceLevel. The type is IsisInternalLevel.
    DestinationLevel interface{}

    // Route policy limiting routes to be propagated. The type is string with
    // length: 1..64. This attribute is mandatory.
    RoutePolicyName interface{}
}

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetFilter() yfilter.YFilter { return propagation.YFilter }

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) SetFilter(yf yfilter.YFilter) { propagation.YFilter = yf }

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetGoName(yname string) string {
    if yname == "source-level" { return "SourceLevel" }
    if yname == "destination-level" { return "DestinationLevel" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    return ""
}

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetSegmentPath() string {
    return "propagation" + "[source-level='" + fmt.Sprintf("%v", propagation.SourceLevel) + "']" + "[destination-level='" + fmt.Sprintf("%v", propagation.DestinationLevel) + "']"
}

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-level"] = propagation.SourceLevel
    leafs["destination-level"] = propagation.DestinationLevel
    leafs["route-policy-name"] = propagation.RoutePolicyName
    return leafs
}

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetBundleName() string { return "cisco_ios_xr" }

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetYangName() string { return "propagation" }

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) SetParent(parent types.Entity) { propagation.parent = parent }

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetParent() types.Entity { return propagation.parent }

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetParentYangName() string { return "propagations" }

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions
// Protocol redistribution configuration
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribution of other protocols into this IS-IS instance. The type is
    // slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution.
    Redistribution []Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution
}

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetFilter() yfilter.YFilter { return redistributions.YFilter }

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) SetFilter(yf yfilter.YFilter) { redistributions.YFilter = yf }

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetGoName(yname string) string {
    if yname == "redistribution" { return "Redistribution" }
    return ""
}

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetSegmentPath() string {
    return "redistributions"
}

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "redistribution" {
        for _, c := range redistributions.Redistribution {
            if redistributions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution{}
        redistributions.Redistribution = append(redistributions.Redistribution, child)
        return &redistributions.Redistribution[len(redistributions.Redistribution)-1]
    }
    return nil
}

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range redistributions.Redistribution {
        children[redistributions.Redistribution[i].GetSegmentPath()] = &redistributions.Redistribution[i]
    }
    return children
}

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetBundleName() string { return "cisco_ios_xr" }

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetYangName() string { return "redistributions" }

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) SetParent(parent types.Entity) { redistributions.parent = parent }

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetParent() types.Entity { return redistributions.parent }

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution
// Redistribution of other protocols into
// this IS-IS instance
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The protocol to be redistributed.  OSPFv3 may not
    // be specified for an IPv4 topology and OSPF may not be specified for an IPv6
    // topology. The type is IsisRedistProto.
    ProtocolName interface{}

    // connected or static or rip or subscriber or mobile.
    ConnectedOrStaticOrRipOrSubscriberOrMobile Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile

    // ospf or ospfv3 or isis or application. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication.
    OspfOrOspfv3OrIsisOrApplication []Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication

    // bgp. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp.
    Bgp []Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp

    // eigrp. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp.
    Eigrp []Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp
}

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetFilter() yfilter.YFilter { return redistribution.YFilter }

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) SetFilter(yf yfilter.YFilter) { redistribution.YFilter = yf }

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetGoName(yname string) string {
    if yname == "protocol-name" { return "ProtocolName" }
    if yname == "connected-or-static-or-rip-or-subscriber-or-mobile" { return "ConnectedOrStaticOrRipOrSubscriberOrMobile" }
    if yname == "ospf-or-ospfv3-or-isis-or-application" { return "OspfOrOspfv3OrIsisOrApplication" }
    if yname == "bgp" { return "Bgp" }
    if yname == "eigrp" { return "Eigrp" }
    return ""
}

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetSegmentPath() string {
    return "redistribution" + "[protocol-name='" + fmt.Sprintf("%v", redistribution.ProtocolName) + "']"
}

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connected-or-static-or-rip-or-subscriber-or-mobile" {
        return &redistribution.ConnectedOrStaticOrRipOrSubscriberOrMobile
    }
    if childYangName == "ospf-or-ospfv3-or-isis-or-application" {
        for _, c := range redistribution.OspfOrOspfv3OrIsisOrApplication {
            if redistribution.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication{}
        redistribution.OspfOrOspfv3OrIsisOrApplication = append(redistribution.OspfOrOspfv3OrIsisOrApplication, child)
        return &redistribution.OspfOrOspfv3OrIsisOrApplication[len(redistribution.OspfOrOspfv3OrIsisOrApplication)-1]
    }
    if childYangName == "bgp" {
        for _, c := range redistribution.Bgp {
            if redistribution.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp{}
        redistribution.Bgp = append(redistribution.Bgp, child)
        return &redistribution.Bgp[len(redistribution.Bgp)-1]
    }
    if childYangName == "eigrp" {
        for _, c := range redistribution.Eigrp {
            if redistribution.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp{}
        redistribution.Eigrp = append(redistribution.Eigrp, child)
        return &redistribution.Eigrp[len(redistribution.Eigrp)-1]
    }
    return nil
}

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["connected-or-static-or-rip-or-subscriber-or-mobile"] = &redistribution.ConnectedOrStaticOrRipOrSubscriberOrMobile
    for i := range redistribution.OspfOrOspfv3OrIsisOrApplication {
        children[redistribution.OspfOrOspfv3OrIsisOrApplication[i].GetSegmentPath()] = &redistribution.OspfOrOspfv3OrIsisOrApplication[i]
    }
    for i := range redistribution.Bgp {
        children[redistribution.Bgp[i].GetSegmentPath()] = &redistribution.Bgp[i]
    }
    for i := range redistribution.Eigrp {
        children[redistribution.Eigrp[i].GetSegmentPath()] = &redistribution.Eigrp[i]
    }
    return children
}

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-name"] = redistribution.ProtocolName
    return leafs
}

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetBundleName() string { return "cisco_ios_xr" }

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetYangName() string { return "redistribution" }

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) SetParent(parent types.Entity) { redistribution.parent = parent }

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetParent() types.Entity { return redistribution.parent }

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetParentYangName() string { return "redistributions" }

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile
// connected or static or rip or subscriber
// or mobile
// This type is a presence type.
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric for redistributed routes: <0-63> for narrow, <0-16777215> for wide.
    // The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Levels to redistribute routes into. The type is IsisConfigurableLevels.
    Levels interface{}

    // Route policy to control redistribution. The type is string with length:
    // 1..64.
    RoutePolicyName interface{}

    // IS-IS metric type. The type is IsisMetric.
    MetricType interface{}

    // OSPF route types to redistribute.  May only be specified if Protocol is
    // OSPF. The type is interface{} with range: -2147483648..2147483647.
    OspfRouteType interface{}
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetFilter() yfilter.YFilter { return connectedOrStaticOrRipOrSubscriberOrMobile.YFilter }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) SetFilter(yf yfilter.YFilter) { connectedOrStaticOrRipOrSubscriberOrMobile.YFilter = yf }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "levels" { return "Levels" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "metric-type" { return "MetricType" }
    if yname == "ospf-route-type" { return "OspfRouteType" }
    return ""
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetSegmentPath() string {
    return "connected-or-static-or-rip-or-subscriber-or-mobile"
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = connectedOrStaticOrRipOrSubscriberOrMobile.Metric
    leafs["levels"] = connectedOrStaticOrRipOrSubscriberOrMobile.Levels
    leafs["route-policy-name"] = connectedOrStaticOrRipOrSubscriberOrMobile.RoutePolicyName
    leafs["metric-type"] = connectedOrStaticOrRipOrSubscriberOrMobile.MetricType
    leafs["ospf-route-type"] = connectedOrStaticOrRipOrSubscriberOrMobile.OspfRouteType
    return leafs
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetBundleName() string { return "cisco_ios_xr" }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetYangName() string { return "connected-or-static-or-rip-or-subscriber-or-mobile" }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) SetParent(parent types.Entity) { connectedOrStaticOrRipOrSubscriberOrMobile.parent = parent }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetParent() types.Entity { return connectedOrStaticOrRipOrSubscriberOrMobile.parent }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetParentYangName() string { return "redistribution" }

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication
// ospf or ospfv3 or isis or application
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol Instance Identifier.  Mandatory for ISIS,
    // OSPF and application, must not be specified otherwise. The type is string
    // with length: 1..64.
    InstanceName interface{}

    // Metric for redistributed routes: <0-63> for narrow, <0-16777215> for wide.
    // The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Levels to redistribute routes into. The type is IsisConfigurableLevels.
    Levels interface{}

    // Route policy to control redistribution. The type is string with length:
    // 1..64.
    RoutePolicyName interface{}

    // IS-IS metric type. The type is IsisMetric.
    MetricType interface{}

    // OSPF route types to redistribute.  May only be specified if Protocol is
    // OSPF. The type is interface{} with range: -2147483648..2147483647.
    OspfRouteType interface{}
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetFilter() yfilter.YFilter { return ospfOrOspfv3OrIsisOrApplication.YFilter }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) SetFilter(yf yfilter.YFilter) { ospfOrOspfv3OrIsisOrApplication.YFilter = yf }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "metric" { return "Metric" }
    if yname == "levels" { return "Levels" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "metric-type" { return "MetricType" }
    if yname == "ospf-route-type" { return "OspfRouteType" }
    return ""
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetSegmentPath() string {
    return "ospf-or-ospfv3-or-isis-or-application" + "[instance-name='" + fmt.Sprintf("%v", ospfOrOspfv3OrIsisOrApplication.InstanceName) + "']"
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = ospfOrOspfv3OrIsisOrApplication.InstanceName
    leafs["metric"] = ospfOrOspfv3OrIsisOrApplication.Metric
    leafs["levels"] = ospfOrOspfv3OrIsisOrApplication.Levels
    leafs["route-policy-name"] = ospfOrOspfv3OrIsisOrApplication.RoutePolicyName
    leafs["metric-type"] = ospfOrOspfv3OrIsisOrApplication.MetricType
    leafs["ospf-route-type"] = ospfOrOspfv3OrIsisOrApplication.OspfRouteType
    return leafs
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetBundleName() string { return "cisco_ios_xr" }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetYangName() string { return "ospf-or-ospfv3-or-isis-or-application" }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) SetParent(parent types.Entity) { ospfOrOspfv3OrIsisOrApplication.parent = parent }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetParent() types.Entity { return ospfOrOspfv3OrIsisOrApplication.parent }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetParentYangName() string { return "redistribution" }

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp
// bgp
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. First half of BGP AS number in XX.YY format. 
    // Mandatory if Protocol is BGP and must not be specified otherwise. Must be a
    // non-zero value if second half is zero. The type is interface{} with range:
    // 0..65535.
    AsXx interface{}

    // This attribute is a key. Second half of BGP AS number in XX.YY format.
    // Mandatory if Protocol is BGP and must not be specified otherwise. Must be a
    // non-zero value if first half is zero. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}

    // Metric for redistributed routes: <0-63> for narrow, <0-16777215> for wide.
    // The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Levels to redistribute routes into. The type is IsisConfigurableLevels.
    Levels interface{}

    // Route policy to control redistribution. The type is string with length:
    // 1..64.
    RoutePolicyName interface{}

    // IS-IS metric type. The type is IsisMetric.
    MetricType interface{}

    // OSPF route types to redistribute.  May only be specified if Protocol is
    // OSPF. The type is interface{} with range: -2147483648..2147483647.
    OspfRouteType interface{}
}

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetGoName(yname string) string {
    if yname == "as-xx" { return "AsXx" }
    if yname == "as-yy" { return "AsYy" }
    if yname == "metric" { return "Metric" }
    if yname == "levels" { return "Levels" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "metric-type" { return "MetricType" }
    if yname == "ospf-route-type" { return "OspfRouteType" }
    return ""
}

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetSegmentPath() string {
    return "bgp" + "[as-xx='" + fmt.Sprintf("%v", bgp.AsXx) + "']" + "[as-yy='" + fmt.Sprintf("%v", bgp.AsYy) + "']"
}

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-xx"] = bgp.AsXx
    leafs["as-yy"] = bgp.AsYy
    leafs["metric"] = bgp.Metric
    leafs["levels"] = bgp.Levels
    leafs["route-policy-name"] = bgp.RoutePolicyName
    leafs["metric-type"] = bgp.MetricType
    leafs["ospf-route-type"] = bgp.OspfRouteType
    return leafs
}

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetYangName() string { return "bgp" }

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetParentYangName() string { return "redistribution" }

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp
// eigrp
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Eigrp as number. The type is interface{} with
    // range: 1..65535.
    AsZz interface{}

    // Metric for redistributed routes: <0-63> for narrow, <0-16777215> for wide.
    // The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Levels to redistribute routes into. The type is IsisConfigurableLevels.
    Levels interface{}

    // Route policy to control redistribution. The type is string with length:
    // 1..64.
    RoutePolicyName interface{}

    // IS-IS metric type. The type is IsisMetric.
    MetricType interface{}

    // OSPF route types to redistribute.  May only be specified if Protocol is
    // OSPF. The type is interface{} with range: -2147483648..2147483647.
    OspfRouteType interface{}
}

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetFilter() yfilter.YFilter { return eigrp.YFilter }

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) SetFilter(yf yfilter.YFilter) { eigrp.YFilter = yf }

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetGoName(yname string) string {
    if yname == "as-zz" { return "AsZz" }
    if yname == "metric" { return "Metric" }
    if yname == "levels" { return "Levels" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "metric-type" { return "MetricType" }
    if yname == "ospf-route-type" { return "OspfRouteType" }
    return ""
}

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetSegmentPath() string {
    return "eigrp" + "[as-zz='" + fmt.Sprintf("%v", eigrp.AsZz) + "']"
}

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-zz"] = eigrp.AsZz
    leafs["metric"] = eigrp.Metric
    leafs["levels"] = eigrp.Levels
    leafs["route-policy-name"] = eigrp.RoutePolicyName
    leafs["metric-type"] = eigrp.MetricType
    leafs["ospf-route-type"] = eigrp.OspfRouteType
    return leafs
}

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetBundleName() string { return "cisco_ios_xr" }

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetYangName() string { return "eigrp" }

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) SetParent(parent types.Entity) { eigrp.parent = parent }

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetParent() types.Entity { return eigrp.parent }

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetParentYangName() string { return "redistribution" }

// Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals
// Peoridic SPF configuration
type Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum interval between spf runs. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval.
    SpfPeriodicInterval []Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetFilter() yfilter.YFilter { return spfPeriodicIntervals.YFilter }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) SetFilter(yf yfilter.YFilter) { spfPeriodicIntervals.YFilter = yf }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetGoName(yname string) string {
    if yname == "spf-periodic-interval" { return "SpfPeriodicInterval" }
    return ""
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetSegmentPath() string {
    return "spf-periodic-intervals"
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spf-periodic-interval" {
        for _, c := range spfPeriodicIntervals.SpfPeriodicInterval {
            if spfPeriodicIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval{}
        spfPeriodicIntervals.SpfPeriodicInterval = append(spfPeriodicIntervals.SpfPeriodicInterval, child)
        return &spfPeriodicIntervals.SpfPeriodicInterval[len(spfPeriodicIntervals.SpfPeriodicInterval)-1]
    }
    return nil
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spfPeriodicIntervals.SpfPeriodicInterval {
        children[spfPeriodicIntervals.SpfPeriodicInterval[i].GetSegmentPath()] = &spfPeriodicIntervals.SpfPeriodicInterval[i]
    }
    return children
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetYangName() string { return "spf-periodic-intervals" }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) SetParent(parent types.Entity) { spfPeriodicIntervals.parent = parent }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetParent() types.Entity { return spfPeriodicIntervals.parent }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval
// Maximum interval between spf runs
type Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Maximum interval in between SPF runs in seconds. The type is interface{}
    // with range: 0..3600. This attribute is mandatory. Units are second.
    PeriodicInterval interface{}
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetFilter() yfilter.YFilter { return spfPeriodicInterval.YFilter }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) SetFilter(yf yfilter.YFilter) { spfPeriodicInterval.YFilter = yf }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "periodic-interval" { return "PeriodicInterval" }
    return ""
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetSegmentPath() string {
    return "spf-periodic-interval" + "[level='" + fmt.Sprintf("%v", spfPeriodicInterval.Level) + "']"
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = spfPeriodicInterval.Level
    leafs["periodic-interval"] = spfPeriodicInterval.PeriodicInterval
    return leafs
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetBundleName() string { return "cisco_ios_xr" }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetYangName() string { return "spf-periodic-interval" }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) SetParent(parent types.Entity) { spfPeriodicInterval.parent = parent }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetParent() types.Entity { return spfPeriodicInterval.parent }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetParentYangName() string { return "spf-periodic-intervals" }

// Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals
// SPF-interval configuration
type Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route calculation scheduling parameters. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval.
    SpfInterval []Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetFilter() yfilter.YFilter { return spfIntervals.YFilter }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) SetFilter(yf yfilter.YFilter) { spfIntervals.YFilter = yf }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetGoName(yname string) string {
    if yname == "spf-interval" { return "SpfInterval" }
    return ""
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetSegmentPath() string {
    return "spf-intervals"
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spf-interval" {
        for _, c := range spfIntervals.SpfInterval {
            if spfIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval{}
        spfIntervals.SpfInterval = append(spfIntervals.SpfInterval, child)
        return &spfIntervals.SpfInterval[len(spfIntervals.SpfInterval)-1]
    }
    return nil
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spfIntervals.SpfInterval {
        children[spfIntervals.SpfInterval[i].GetSegmentPath()] = &spfIntervals.SpfInterval[i]
    }
    return children
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetYangName() string { return "spf-intervals" }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) SetParent(parent types.Entity) { spfIntervals.parent = parent }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetParent() types.Entity { return spfIntervals.parent }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval
// Route calculation scheduling parameters
type Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Maximum wait before running a route calculation in milliseconds. The type
    // is interface{} with range: 0..120000. Units are millisecond.
    MaximumWait interface{}

    // Initial wait before running a route calculation in milliseconds. The type
    // is interface{} with range: 0..120000. Units are millisecond.
    InitialWait interface{}

    // Secondary wait before running a route calculation in milliseconds. The type
    // is interface{} with range: 0..120000. Units are millisecond.
    SecondaryWait interface{}
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetFilter() yfilter.YFilter { return spfInterval.YFilter }

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) SetFilter(yf yfilter.YFilter) { spfInterval.YFilter = yf }

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "maximum-wait" { return "MaximumWait" }
    if yname == "initial-wait" { return "InitialWait" }
    if yname == "secondary-wait" { return "SecondaryWait" }
    return ""
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetSegmentPath() string {
    return "spf-interval" + "[level='" + fmt.Sprintf("%v", spfInterval.Level) + "']"
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = spfInterval.Level
    leafs["maximum-wait"] = spfInterval.MaximumWait
    leafs["initial-wait"] = spfInterval.InitialWait
    leafs["secondary-wait"] = spfInterval.SecondaryWait
    return leafs
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetBundleName() string { return "cisco_ios_xr" }

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetYangName() string { return "spf-interval" }

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) SetParent(parent types.Entity) { spfInterval.parent = parent }

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetParent() types.Entity { return spfInterval.parent }

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetParentYangName() string { return "spf-intervals" }

// Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence
// Enable convergence monitoring
type Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable convergence monitoring. The type is interface{}.
    Enable interface{}

    // Enable the Tracking of IP-Frr Convergence. The type is interface{}.
    TrackIpFrr interface{}

    // Enable the monitoring of individual prefixes (prefix list name). The type
    // is string with length: 1..32.
    PrefixList interface{}
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetFilter() yfilter.YFilter { return monitorConvergence.YFilter }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) SetFilter(yf yfilter.YFilter) { monitorConvergence.YFilter = yf }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "track-ip-frr" { return "TrackIpFrr" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetSegmentPath() string {
    return "monitor-convergence"
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = monitorConvergence.Enable
    leafs["track-ip-frr"] = monitorConvergence.TrackIpFrr
    leafs["prefix-list"] = monitorConvergence.PrefixList
    return leafs
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetBundleName() string { return "cisco_ios_xr" }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetYangName() string { return "monitor-convergence" }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) SetParent(parent types.Entity) { monitorConvergence.parent = parent }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetParent() types.Entity { return monitorConvergence.parent }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation
// Control origination of a default route with
// the option of using a policy.  If no policy
// is specified the default route is
// advertised with zero cost in level 2 only.
type Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flag to indicate whether default origination is controlled using a policy.
    // The type is bool.
    UsePolicy interface{}

    // Policy name. The type is string with length: 1..64.
    PolicyName interface{}

    // Flag to indicate that the default prefix should be originated as an
    // external route. The type is interface{}.
    External interface{}
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetFilter() yfilter.YFilter { return defaultInformation.YFilter }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) SetFilter(yf yfilter.YFilter) { defaultInformation.YFilter = yf }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetGoName(yname string) string {
    if yname == "use-policy" { return "UsePolicy" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "external" { return "External" }
    return ""
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetSegmentPath() string {
    return "default-information"
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["use-policy"] = defaultInformation.UsePolicy
    leafs["policy-name"] = defaultInformation.PolicyName
    leafs["external"] = defaultInformation.External
    return leafs
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetBundleName() string { return "cisco_ios_xr" }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetYangName() string { return "default-information" }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) SetParent(parent types.Entity) { defaultInformation.parent = parent }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetParent() types.Entity { return defaultInformation.parent }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_AdminDistances
// Per-route administrative
// distanceconfiguration
type Isis_Instances_Instance_Afs_Af_AfData_AdminDistances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative distance configuration. The supplied distance is applied to
    // all routes discovered from the specified source, or only those that match
    // the supplied prefix list if this is specified. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance.
    AdminDistance []Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetFilter() yfilter.YFilter { return adminDistances.YFilter }

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) SetFilter(yf yfilter.YFilter) { adminDistances.YFilter = yf }

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetGoName(yname string) string {
    if yname == "admin-distance" { return "AdminDistance" }
    return ""
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetSegmentPath() string {
    return "admin-distances"
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "admin-distance" {
        for _, c := range adminDistances.AdminDistance {
            if adminDistances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance{}
        adminDistances.AdminDistance = append(adminDistances.AdminDistance, child)
        return &adminDistances.AdminDistance[len(adminDistances.AdminDistance)-1]
    }
    return nil
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range adminDistances.AdminDistance {
        children[adminDistances.AdminDistance[i].GetSegmentPath()] = &adminDistances.AdminDistance[i]
    }
    return children
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetBundleName() string { return "cisco_ios_xr" }

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetYangName() string { return "admin-distances" }

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) SetParent(parent types.Entity) { adminDistances.parent = parent }

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetParent() types.Entity { return adminDistances.parent }

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance
// Administrative distance configuration. The
// supplied distance is applied to all routes
// discovered from the specified source, or
// only those that match the supplied prefix
// list if this is specified
type Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP route source prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}

    // Administrative distance. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    Distance interface{}

    // List of prefixes to which this distance applies. The type is string with
    // length: 1..32.
    PrefixList interface{}
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetFilter() yfilter.YFilter { return adminDistance.YFilter }

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) SetFilter(yf yfilter.YFilter) { adminDistance.YFilter = yf }

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetGoName(yname string) string {
    if yname == "address-prefix" { return "AddressPrefix" }
    if yname == "distance" { return "Distance" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetSegmentPath() string {
    return "admin-distance" + "[address-prefix='" + fmt.Sprintf("%v", adminDistance.AddressPrefix) + "']"
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-prefix"] = adminDistance.AddressPrefix
    leafs["distance"] = adminDistance.Distance
    leafs["prefix-list"] = adminDistance.PrefixList
    return leafs
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetBundleName() string { return "cisco_ios_xr" }

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetYangName() string { return "admin-distance" }

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) SetParent(parent types.Entity) { adminDistance.parent = parent }

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetParent() types.Entity { return adminDistance.parent }

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetParentYangName() string { return "admin-distances" }

// Isis_Instances_Instance_Afs_Af_AfData_Ispf
// ISPF configuration
type Isis_Instances_Instance_Afs_Af_AfData_Ispf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISPF state (enable/disable).
    States Isis_Instances_Instance_Afs_Af_AfData_Ispf_States
}

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetFilter() yfilter.YFilter { return ispf.YFilter }

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) SetFilter(yf yfilter.YFilter) { ispf.YFilter = yf }

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetGoName(yname string) string {
    if yname == "states" { return "States" }
    return ""
}

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetSegmentPath() string {
    return "ispf"
}

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "states" {
        return &ispf.States
    }
    return nil
}

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["states"] = &ispf.States
    return children
}

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetBundleName() string { return "cisco_ios_xr" }

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetYangName() string { return "ispf" }

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) SetParent(parent types.Entity) { ispf.parent = parent }

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetParent() types.Entity { return ispf.parent }

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_Ispf_States
// ISPF state (enable/disable)
type Isis_Instances_Instance_Afs_Af_AfData_Ispf_States struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/disable ISPF. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State.
    State []Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State
}

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetFilter() yfilter.YFilter { return states.YFilter }

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) SetFilter(yf yfilter.YFilter) { states.YFilter = yf }

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    return ""
}

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetSegmentPath() string {
    return "states"
}

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        for _, c := range states.State {
            if states.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State{}
        states.State = append(states.State, child)
        return &states.State[len(states.State)-1]
    }
    return nil
}

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range states.State {
        children[states.State[i].GetSegmentPath()] = &states.State[i]
    }
    return children
}

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetBundleName() string { return "cisco_ios_xr" }

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetYangName() string { return "states" }

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) SetParent(parent types.Entity) { states.parent = parent }

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetParent() types.Entity { return states.parent }

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetParentYangName() string { return "ispf" }

// Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State
// Enable/disable ISPF
type Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // State. The type is IsisispfState. This attribute is mandatory.
    State interface{}
}

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "state" { return "State" }
    return ""
}

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetSegmentPath() string {
    return "state" + "[level='" + fmt.Sprintf("%v", state.Level) + "']"
}

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = state.Level
    leafs["state"] = state.State
    return leafs
}

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetYangName() string { return "state" }

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetParent() types.Entity { return state.parent }

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetParentYangName() string { return "states" }

// Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal
// MPLS LDP configuration. MPLS LDP
// configuration will only be applied for the
// IPv4-unicast address-family.
type Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If TRUE, LDP will be enabled onall IS-IS interfaces enabled for this
    // address-family. The type is bool.
    AutoConfig interface{}
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetFilter() yfilter.YFilter { return mplsLdpGlobal.YFilter }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) SetFilter(yf yfilter.YFilter) { mplsLdpGlobal.YFilter = yf }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetGoName(yname string) string {
    if yname == "auto-config" { return "AutoConfig" }
    return ""
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetSegmentPath() string {
    return "mpls-ldp-global"
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auto-config"] = mplsLdpGlobal.AutoConfig
    return leafs
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetBundleName() string { return "cisco_ios_xr" }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetYangName() string { return "mpls-ldp-global" }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) SetParent(parent types.Entity) { mplsLdpGlobal.parent = parent }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetParent() types.Entity { return mplsLdpGlobal.parent }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_Mpls
// MPLS configuration. MPLS configuration will
// only be applied for the IPv4-unicast
// address-family.
type Isis_Instances_Instance_Afs_Af_AfData_Mpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install TE and non-TE nexthops in the RIB. The type is interface{}.
    IgpIntact interface{}

    // Install non-TE nexthops in the RIB for use by multicast. The type is
    // interface{}.
    MulticastIntact interface{}

    // Traffic Engineering stable IP address for system.
    RouterId Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId

    // Enable MPLS for an IS-IS at the given levels.
    Level Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level
}

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetFilter() yfilter.YFilter { return mpls.YFilter }

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) SetFilter(yf yfilter.YFilter) { mpls.YFilter = yf }

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetGoName(yname string) string {
    if yname == "igp-intact" { return "IgpIntact" }
    if yname == "multicast-intact" { return "MulticastIntact" }
    if yname == "router-id" { return "RouterId" }
    if yname == "level" { return "Level" }
    return ""
}

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetSegmentPath() string {
    return "mpls"
}

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router-id" {
        return &mpls.RouterId
    }
    if childYangName == "level" {
        return &mpls.Level
    }
    return nil
}

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["router-id"] = &mpls.RouterId
    children["level"] = &mpls.Level
    return children
}

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-intact"] = mpls.IgpIntact
    leafs["multicast-intact"] = mpls.MulticastIntact
    return leafs
}

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetBundleName() string { return "cisco_ios_xr" }

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetYangName() string { return "mpls" }

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) SetParent(parent types.Entity) { mpls.parent = parent }

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetParent() types.Entity { return mpls.parent }

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId
// Traffic Engineering stable IP address for
// system
type Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address to be used as a router ID. Precisely one of Address and
    // Interface must be specified. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Interface with designated stable IP address to be used as a router ID. This
    // must be a Loopback interface. Precisely one of Address and Interface must
    // be specified. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetFilter() yfilter.YFilter { return routerId.YFilter }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) SetFilter(yf yfilter.YFilter) { routerId.YFilter = yf }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetSegmentPath() string {
    return "router-id"
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = routerId.Address
    leafs["interface-name"] = routerId.InterfaceName
    return leafs
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetBundleName() string { return "cisco_ios_xr" }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetYangName() string { return "router-id" }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) SetParent(parent types.Entity) { routerId.parent = parent }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetParent() types.Entity { return routerId.parent }

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetParentYangName() string { return "mpls" }

// Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level
// Enable MPLS for an IS-IS at the given
// levels
type Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Level 1 enabled. The type is bool.
    Level1 interface{}

    // Level 2 enabled. The type is bool.
    Level2 interface{}
}

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetFilter() yfilter.YFilter { return level.YFilter }

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) SetFilter(yf yfilter.YFilter) { level.YFilter = yf }

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetGoName(yname string) string {
    if yname == "level1" { return "Level1" }
    if yname == "level2" { return "Level2" }
    return ""
}

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetSegmentPath() string {
    return "level"
}

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level1"] = level.Level1
    leafs["level2"] = level.Level2
    return leafs
}

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetBundleName() string { return "cisco_ios_xr" }

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetYangName() string { return "level" }

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) SetParent(parent types.Entity) { level.parent = parent }

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetParent() types.Entity { return level.parent }

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetParentYangName() string { return "mpls" }

// Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids
// Manual Adjacecy SID configuration
type Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Assign adjancency SID to an interface. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid.
    ManualAdjSid []Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetFilter() yfilter.YFilter { return manualAdjSids.YFilter }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) SetFilter(yf yfilter.YFilter) { manualAdjSids.YFilter = yf }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetGoName(yname string) string {
    if yname == "manual-adj-sid" { return "ManualAdjSid" }
    return ""
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetSegmentPath() string {
    return "manual-adj-sids"
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "manual-adj-sid" {
        for _, c := range manualAdjSids.ManualAdjSid {
            if manualAdjSids.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid{}
        manualAdjSids.ManualAdjSid = append(manualAdjSids.ManualAdjSid, child)
        return &manualAdjSids.ManualAdjSid[len(manualAdjSids.ManualAdjSid)-1]
    }
    return nil
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range manualAdjSids.ManualAdjSid {
        children[manualAdjSids.ManualAdjSid[i].GetSegmentPath()] = &manualAdjSids.ManualAdjSid[i]
    }
    return children
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetBundleName() string { return "cisco_ios_xr" }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetYangName() string { return "manual-adj-sids" }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) SetParent(parent types.Entity) { manualAdjSids.parent = parent }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetParent() types.Entity { return manualAdjSids.parent }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid
// Assign adjancency SID to an interface
type Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Sid type aboslute or index. The type is Isissid1.
    SidType interface{}

    // This attribute is a key. Sid value. The type is interface{} with range:
    // 0..1048575.
    Sid interface{}

    // Enable/Disable SID protection. The type is IsissidProtected. This attribute
    // is mandatory.
    Protected interface{}
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetFilter() yfilter.YFilter { return manualAdjSid.YFilter }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) SetFilter(yf yfilter.YFilter) { manualAdjSid.YFilter = yf }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "sid-type" { return "SidType" }
    if yname == "sid" { return "Sid" }
    if yname == "protected" { return "Protected" }
    return ""
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetSegmentPath() string {
    return "manual-adj-sid" + "[level='" + fmt.Sprintf("%v", manualAdjSid.Level) + "']" + "[sid-type='" + fmt.Sprintf("%v", manualAdjSid.SidType) + "']" + "[sid='" + fmt.Sprintf("%v", manualAdjSid.Sid) + "']"
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = manualAdjSid.Level
    leafs["sid-type"] = manualAdjSid.SidType
    leafs["sid"] = manualAdjSid.Sid
    leafs["protected"] = manualAdjSid.Protected
    return leafs
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetBundleName() string { return "cisco_ios_xr" }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetYangName() string { return "manual-adj-sid" }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) SetParent(parent types.Entity) { manualAdjSid.parent = parent }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetParent() types.Entity { return manualAdjSid.parent }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetParentYangName() string { return "manual-adj-sids" }

// Isis_Instances_Instance_Afs_Af_AfData_Metrics
// Metric configuration
type Isis_Instances_Instance_Afs_Af_AfData_Metrics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric configuration. Legal value depends on the metric-style specified for
    // the topology. If the metric-style defined is narrow, then only a value
    // between <1-63> is allowed and if the metric-style is defined as wide, then
    // a value between <1-16777215> is allowed as the metric value.  All routers
    // exclude links with the maximum wide metric (16777215) from their SPF. The
    // type is slice of Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric.
    Metric []Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric
}

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetFilter() yfilter.YFilter { return metrics.YFilter }

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) SetFilter(yf yfilter.YFilter) { metrics.YFilter = yf }

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    return ""
}

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetSegmentPath() string {
    return "metrics"
}

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metric" {
        for _, c := range metrics.Metric {
            if metrics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric{}
        metrics.Metric = append(metrics.Metric, child)
        return &metrics.Metric[len(metrics.Metric)-1]
    }
    return nil
}

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range metrics.Metric {
        children[metrics.Metric[i].GetSegmentPath()] = &metrics.Metric[i]
    }
    return children
}

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetBundleName() string { return "cisco_ios_xr" }

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetYangName() string { return "metrics" }

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) SetParent(parent types.Entity) { metrics.parent = parent }

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetParent() types.Entity { return metrics.parent }

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric
// Metric configuration. Legal value depends on
// the metric-style specified for the topology. If
// the metric-style defined is narrow, then only a
// value between <1-63> is allowed and if the
// metric-style is defined as wide, then a value
// between <1-16777215> is allowed as the metric
// value.  All routers exclude links with the
// maximum wide metric (16777215) from their SPF
type Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed metric: <1-63> for narrow, <1-16777215> for wide. The type is one
    // of the following types: enumeration
    // Isis.Instances.Instance.Interfaces.Interface.InterfaceAfs.InterfaceAf.TopologyName.Metrics.Metric.Metric
    // This attribute is mandatory., or int with range: 1..16777215 This attribute
    // is mandatory..
    Metric interface{}
}

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetFilter() yfilter.YFilter { return metric.YFilter }

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) SetFilter(yf yfilter.YFilter) { metric.YFilter = yf }

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetSegmentPath() string {
    return "metric" + "[level='" + fmt.Sprintf("%v", metric.Level) + "']"
}

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = metric.Level
    leafs["metric"] = metric.Metric
    return leafs
}

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetBundleName() string { return "cisco_ios_xr" }

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetYangName() string { return "metric" }

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) SetParent(parent types.Entity) { metric.parent = parent }

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetParent() types.Entity { return metric.parent }

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetParentYangName() string { return "metrics" }

// Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric_Metric represents <1-16777215> for wide
type Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric_Metric string

const (
    // Maximum wide metric.  All routers will
    // exclude this link from their SPF
    Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric_Metric_maximum Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric_Metric = "maximum"
)

// Isis_Instances_Instance_Afs_Af_AfData_Weights
// Weight configuration
type Isis_Instances_Instance_Afs_Af_AfData_Weights struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Weight configuration under interface for load balancing. The type is slice
    // of Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight.
    Weight []Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight
}

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetFilter() yfilter.YFilter { return weights.YFilter }

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) SetFilter(yf yfilter.YFilter) { weights.YFilter = yf }

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetGoName(yname string) string {
    if yname == "weight" { return "Weight" }
    return ""
}

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetSegmentPath() string {
    return "weights"
}

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "weight" {
        for _, c := range weights.Weight {
            if weights.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight{}
        weights.Weight = append(weights.Weight, child)
        return &weights.Weight[len(weights.Weight)-1]
    }
    return nil
}

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range weights.Weight {
        children[weights.Weight[i].GetSegmentPath()] = &weights.Weight[i]
    }
    return children
}

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetBundleName() string { return "cisco_ios_xr" }

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetYangName() string { return "weights" }

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) SetParent(parent types.Entity) { weights.parent = parent }

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetParent() types.Entity { return weights.parent }

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetParentYangName() string { return "af-data" }

// Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight
// Weight configuration under interface for load
// balancing
type Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Weight to be configured under interface for Load Balancing. Allowed weight:
    // <1-16777215>. The type is interface{} with range: 1..16777214. This
    // attribute is mandatory.
    Weight interface{}
}

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetFilter() yfilter.YFilter { return weight.YFilter }

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) SetFilter(yf yfilter.YFilter) { weight.YFilter = yf }

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "weight" { return "Weight" }
    return ""
}

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetSegmentPath() string {
    return "weight" + "[level='" + fmt.Sprintf("%v", weight.Level) + "']"
}

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = weight.Level
    leafs["weight"] = weight.Weight
    return leafs
}

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetBundleName() string { return "cisco_ios_xr" }

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetYangName() string { return "weight" }

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) SetParent(parent types.Entity) { weight.parent = parent }

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetParent() types.Entity { return weight.parent }

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetParentYangName() string { return "weights" }

// Isis_Instances_Instance_Afs_Af_TopologyName
// keys: topology-name
type Isis_Instances_Instance_Afs_Af_TopologyName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Topology Name. The type is string with length:
    // 1..32.
    TopologyName interface{}

    // Maximum number of active parallel paths per route. The type is interface{}
    // with range: 1..64.
    MaximumPaths interface{}

    // Set the topology ID for a named (non-default) topology. This object must be
    // set before any other configuration is supplied for a named (non-default)
    // topology , and must be the last configuration object to be removed. This
    // item should not be supplied for the non-named default topologies. The type
    // is interface{} with range: 6..4095.
    TopologyId interface{}

    // Run IPv6 Unicast using the standard (IPv4 Unicast) topology. The type is
    // interface{}.
    SingleTopology interface{}

    // Suppress check for consistent AF support on received IIHs. The type is
    // IsisAdjCheck.
    AdjacencyCheck interface{}

    // If TRUE, advertise additional link attributes in our LSP. The type is bool.
    AdvertiseLinkAttributes interface{}

    // Apply weights to UCMP or ECMP only. The type is IsisApplyWeight.
    ApplyWeight interface{}

    // Default IS-IS administrative distance configuration. The type is
    // interface{} with range: 1..255. The default value is 115.
    DefaultAdminDistance interface{}

    // If enabled, advertise prefixes of passive interfaces only. The type is
    // interface{}.
    AdvertisePassiveOnly interface{}

    // If TRUE, Ignore other routers attached bit. The type is bool.
    IgnoreAttachedBit interface{}

    // Set the attached bit in this router's level 1 System LSP. The type is
    // IsisAttachedBit. The default value is area.
    AttachedBit interface{}

    // If TRUE, routes will be installed with the IP address of the first-hop node
    // as the source instead of the originating node. The type is bool.
    RouteSourceFirstHop interface{}

    // Enable Segment Routing configuration.
    SegmentRouting Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting

    // Metric-style configuration.
    MetricStyles Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles

    // Fast-ReRoute configuration.
    FrrTable Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable

    // Stable IP address for system. Will only be applied for the unicast
    // sub-address-family.
    RouterId Isis_Instances_Instance_Afs_Af_TopologyName_RouterId

    // SPF Prefix Priority configuration.
    SpfPrefixPriorities Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities

    // Summary-prefix configuration.
    SummaryPrefixes Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes

    // Micro Loop Avoidance configuration.
    MicroLoopAvoidance Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance

    // UCMP (UnEqual Cost MultiPath) configuration.
    Ucmp Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp

    // Maximum number of redistributed prefixesconfiguration.
    MaxRedistPrefixes Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes

    // Route propagation configuration.
    Propagations Isis_Instances_Instance_Afs_Af_TopologyName_Propagations

    // Protocol redistribution configuration.
    Redistributions Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions

    // Peoridic SPF configuration.
    SpfPeriodicIntervals Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals

    // SPF-interval configuration.
    SpfIntervals Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals

    // Enable convergence monitoring.
    MonitorConvergence Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence

    // Control origination of a default route with the option of using a policy. 
    // If no policy is specified the default route is advertised with zero cost in
    // level 2 only.
    DefaultInformation Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation

    // Per-route administrative distanceconfiguration.
    AdminDistances Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances

    // ISPF configuration.
    Ispf Isis_Instances_Instance_Afs_Af_TopologyName_Ispf

    // MPLS LDP configuration. MPLS LDP configuration will only be applied for the
    // IPv4-unicast address-family.
    MplsLdpGlobal Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal

    // MPLS configuration. MPLS configuration will only be applied for the
    // IPv4-unicast address-family.
    Mpls Isis_Instances_Instance_Afs_Af_TopologyName_Mpls

    // Manual Adjacecy SID configuration.
    ManualAdjSids Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids

    // Metric configuration.
    Metrics Isis_Instances_Instance_Afs_Af_TopologyName_Metrics

    // Weight configuration.
    Weights Isis_Instances_Instance_Afs_Af_TopologyName_Weights
}

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetFilter() yfilter.YFilter { return topologyName.YFilter }

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) SetFilter(yf yfilter.YFilter) { topologyName.YFilter = yf }

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetGoName(yname string) string {
    if yname == "topology-name" { return "TopologyName" }
    if yname == "maximum-paths" { return "MaximumPaths" }
    if yname == "topology-id" { return "TopologyId" }
    if yname == "single-topology" { return "SingleTopology" }
    if yname == "adjacency-check" { return "AdjacencyCheck" }
    if yname == "advertise-link-attributes" { return "AdvertiseLinkAttributes" }
    if yname == "apply-weight" { return "ApplyWeight" }
    if yname == "default-admin-distance" { return "DefaultAdminDistance" }
    if yname == "advertise-passive-only" { return "AdvertisePassiveOnly" }
    if yname == "ignore-attached-bit" { return "IgnoreAttachedBit" }
    if yname == "attached-bit" { return "AttachedBit" }
    if yname == "route-source-first-hop" { return "RouteSourceFirstHop" }
    if yname == "segment-routing" { return "SegmentRouting" }
    if yname == "metric-styles" { return "MetricStyles" }
    if yname == "frr-table" { return "FrrTable" }
    if yname == "router-id" { return "RouterId" }
    if yname == "spf-prefix-priorities" { return "SpfPrefixPriorities" }
    if yname == "summary-prefixes" { return "SummaryPrefixes" }
    if yname == "micro-loop-avoidance" { return "MicroLoopAvoidance" }
    if yname == "ucmp" { return "Ucmp" }
    if yname == "max-redist-prefixes" { return "MaxRedistPrefixes" }
    if yname == "propagations" { return "Propagations" }
    if yname == "redistributions" { return "Redistributions" }
    if yname == "spf-periodic-intervals" { return "SpfPeriodicIntervals" }
    if yname == "spf-intervals" { return "SpfIntervals" }
    if yname == "monitor-convergence" { return "MonitorConvergence" }
    if yname == "default-information" { return "DefaultInformation" }
    if yname == "admin-distances" { return "AdminDistances" }
    if yname == "ispf" { return "Ispf" }
    if yname == "mpls-ldp-global" { return "MplsLdpGlobal" }
    if yname == "mpls" { return "Mpls" }
    if yname == "manual-adj-sids" { return "ManualAdjSids" }
    if yname == "metrics" { return "Metrics" }
    if yname == "weights" { return "Weights" }
    return ""
}

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetSegmentPath() string {
    return "topology-name" + "[topology-name='" + fmt.Sprintf("%v", topologyName.TopologyName) + "']"
}

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "segment-routing" {
        return &topologyName.SegmentRouting
    }
    if childYangName == "metric-styles" {
        return &topologyName.MetricStyles
    }
    if childYangName == "frr-table" {
        return &topologyName.FrrTable
    }
    if childYangName == "router-id" {
        return &topologyName.RouterId
    }
    if childYangName == "spf-prefix-priorities" {
        return &topologyName.SpfPrefixPriorities
    }
    if childYangName == "summary-prefixes" {
        return &topologyName.SummaryPrefixes
    }
    if childYangName == "micro-loop-avoidance" {
        return &topologyName.MicroLoopAvoidance
    }
    if childYangName == "ucmp" {
        return &topologyName.Ucmp
    }
    if childYangName == "max-redist-prefixes" {
        return &topologyName.MaxRedistPrefixes
    }
    if childYangName == "propagations" {
        return &topologyName.Propagations
    }
    if childYangName == "redistributions" {
        return &topologyName.Redistributions
    }
    if childYangName == "spf-periodic-intervals" {
        return &topologyName.SpfPeriodicIntervals
    }
    if childYangName == "spf-intervals" {
        return &topologyName.SpfIntervals
    }
    if childYangName == "monitor-convergence" {
        return &topologyName.MonitorConvergence
    }
    if childYangName == "default-information" {
        return &topologyName.DefaultInformation
    }
    if childYangName == "admin-distances" {
        return &topologyName.AdminDistances
    }
    if childYangName == "ispf" {
        return &topologyName.Ispf
    }
    if childYangName == "mpls-ldp-global" {
        return &topologyName.MplsLdpGlobal
    }
    if childYangName == "mpls" {
        return &topologyName.Mpls
    }
    if childYangName == "manual-adj-sids" {
        return &topologyName.ManualAdjSids
    }
    if childYangName == "metrics" {
        return &topologyName.Metrics
    }
    if childYangName == "weights" {
        return &topologyName.Weights
    }
    return nil
}

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["segment-routing"] = &topologyName.SegmentRouting
    children["metric-styles"] = &topologyName.MetricStyles
    children["frr-table"] = &topologyName.FrrTable
    children["router-id"] = &topologyName.RouterId
    children["spf-prefix-priorities"] = &topologyName.SpfPrefixPriorities
    children["summary-prefixes"] = &topologyName.SummaryPrefixes
    children["micro-loop-avoidance"] = &topologyName.MicroLoopAvoidance
    children["ucmp"] = &topologyName.Ucmp
    children["max-redist-prefixes"] = &topologyName.MaxRedistPrefixes
    children["propagations"] = &topologyName.Propagations
    children["redistributions"] = &topologyName.Redistributions
    children["spf-periodic-intervals"] = &topologyName.SpfPeriodicIntervals
    children["spf-intervals"] = &topologyName.SpfIntervals
    children["monitor-convergence"] = &topologyName.MonitorConvergence
    children["default-information"] = &topologyName.DefaultInformation
    children["admin-distances"] = &topologyName.AdminDistances
    children["ispf"] = &topologyName.Ispf
    children["mpls-ldp-global"] = &topologyName.MplsLdpGlobal
    children["mpls"] = &topologyName.Mpls
    children["manual-adj-sids"] = &topologyName.ManualAdjSids
    children["metrics"] = &topologyName.Metrics
    children["weights"] = &topologyName.Weights
    return children
}

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-name"] = topologyName.TopologyName
    leafs["maximum-paths"] = topologyName.MaximumPaths
    leafs["topology-id"] = topologyName.TopologyId
    leafs["single-topology"] = topologyName.SingleTopology
    leafs["adjacency-check"] = topologyName.AdjacencyCheck
    leafs["advertise-link-attributes"] = topologyName.AdvertiseLinkAttributes
    leafs["apply-weight"] = topologyName.ApplyWeight
    leafs["default-admin-distance"] = topologyName.DefaultAdminDistance
    leafs["advertise-passive-only"] = topologyName.AdvertisePassiveOnly
    leafs["ignore-attached-bit"] = topologyName.IgnoreAttachedBit
    leafs["attached-bit"] = topologyName.AttachedBit
    leafs["route-source-first-hop"] = topologyName.RouteSourceFirstHop
    return leafs
}

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetBundleName() string { return "cisco_ios_xr" }

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetYangName() string { return "topology-name" }

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) SetParent(parent types.Entity) { topologyName.parent = parent }

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetParent() types.Entity { return topologyName.parent }

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetParentYangName() string { return "af" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting
// Enable Segment Routing configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable per bundle member adjacency SID. The type is interface{}.
    BundleMemberAdjSid interface{}

    // Prefer segment routing labels over LDP labels. The type is
    // IsisLabelPreference.
    Mpls interface{}

    // Enable Segment Routing prefix SID map configuration.
    PrefixSidMap Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetFilter() yfilter.YFilter { return segmentRouting.YFilter }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) SetFilter(yf yfilter.YFilter) { segmentRouting.YFilter = yf }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetGoName(yname string) string {
    if yname == "bundle-member-adj-sid" { return "BundleMemberAdjSid" }
    if yname == "mpls" { return "Mpls" }
    if yname == "prefix-sid-map" { return "PrefixSidMap" }
    return ""
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetSegmentPath() string {
    return "segment-routing"
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-sid-map" {
        return &segmentRouting.PrefixSidMap
    }
    return nil
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-sid-map"] = &segmentRouting.PrefixSidMap
    return children
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bundle-member-adj-sid"] = segmentRouting.BundleMemberAdjSid
    leafs["mpls"] = segmentRouting.Mpls
    return leafs
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetBundleName() string { return "cisco_ios_xr" }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetYangName() string { return "segment-routing" }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) SetParent(parent types.Entity) { segmentRouting.parent = parent }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetParent() types.Entity { return segmentRouting.parent }

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap
// Enable Segment Routing prefix SID map
// configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Segment Routing prefix SID map advertise local. The type is
    // interface{}.
    AdvertiseLocal interface{}

    // If TRUE, remote prefix SID map advertisements will be used. If FALSE, they
    // will not be used. The type is bool.
    Receive interface{}
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetFilter() yfilter.YFilter { return prefixSidMap.YFilter }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) SetFilter(yf yfilter.YFilter) { prefixSidMap.YFilter = yf }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetGoName(yname string) string {
    if yname == "advertise-local" { return "AdvertiseLocal" }
    if yname == "receive" { return "Receive" }
    return ""
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetSegmentPath() string {
    return "prefix-sid-map"
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["advertise-local"] = prefixSidMap.AdvertiseLocal
    leafs["receive"] = prefixSidMap.Receive
    return leafs
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetYangName() string { return "prefix-sid-map" }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) SetParent(parent types.Entity) { prefixSidMap.parent = parent }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetParent() types.Entity { return prefixSidMap.parent }

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetParentYangName() string { return "segment-routing" }

// Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles
// Metric-style configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration of metric style in LSPs. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle.
    MetricStyle []Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetFilter() yfilter.YFilter { return metricStyles.YFilter }

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) SetFilter(yf yfilter.YFilter) { metricStyles.YFilter = yf }

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetGoName(yname string) string {
    if yname == "metric-style" { return "MetricStyle" }
    return ""
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetSegmentPath() string {
    return "metric-styles"
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metric-style" {
        for _, c := range metricStyles.MetricStyle {
            if metricStyles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle{}
        metricStyles.MetricStyle = append(metricStyles.MetricStyle, child)
        return &metricStyles.MetricStyle[len(metricStyles.MetricStyle)-1]
    }
    return nil
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range metricStyles.MetricStyle {
        children[metricStyles.MetricStyle[i].GetSegmentPath()] = &metricStyles.MetricStyle[i]
    }
    return children
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetBundleName() string { return "cisco_ios_xr" }

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetYangName() string { return "metric-styles" }

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) SetParent(parent types.Entity) { metricStyles.parent = parent }

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetParent() types.Entity { return metricStyles.parent }

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle
// Configuration of metric style in LSPs
type Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Metric Style. The type is IsisMetricStyle. The default value is
    // old-metric-style.
    Style interface{}

    // Transition state. The type is IsisMetricStyleTransition. The default value
    // is disabled.
    TransitionState interface{}
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetFilter() yfilter.YFilter { return metricStyle.YFilter }

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) SetFilter(yf yfilter.YFilter) { metricStyle.YFilter = yf }

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "style" { return "Style" }
    if yname == "transition-state" { return "TransitionState" }
    return ""
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetSegmentPath() string {
    return "metric-style" + "[level='" + fmt.Sprintf("%v", metricStyle.Level) + "']"
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = metricStyle.Level
    leafs["style"] = metricStyle.Style
    leafs["transition-state"] = metricStyle.TransitionState
    return leafs
}

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetBundleName() string { return "cisco_ios_xr" }

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetYangName() string { return "metric-style" }

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) SetParent(parent types.Entity) { metricStyle.parent = parent }

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetParent() types.Entity { return metricStyle.parent }

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetParentYangName() string { return "metric-styles" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable
// Fast-ReRoute configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Load share prefixes across multiple backups.
    FrrLoadSharings Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings

    // FRR prefix-limit configuration.
    PriorityLimits Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits

    // FRR remote LFA prefix list filter configuration.
    FrrRemoteLfaPrefixes Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes

    // FRR tiebreakers configuration.
    FrrTiebreakers Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers

    // FRR use candidate only configuration.
    FrrUseCandOnlies Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies
}

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetFilter() yfilter.YFilter { return frrTable.YFilter }

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) SetFilter(yf yfilter.YFilter) { frrTable.YFilter = yf }

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetGoName(yname string) string {
    if yname == "frr-load-sharings" { return "FrrLoadSharings" }
    if yname == "priority-limits" { return "PriorityLimits" }
    if yname == "frr-remote-lfa-prefixes" { return "FrrRemoteLfaPrefixes" }
    if yname == "frr-tiebreakers" { return "FrrTiebreakers" }
    if yname == "frr-use-cand-onlies" { return "FrrUseCandOnlies" }
    return ""
}

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetSegmentPath() string {
    return "frr-table"
}

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-load-sharings" {
        return &frrTable.FrrLoadSharings
    }
    if childYangName == "priority-limits" {
        return &frrTable.PriorityLimits
    }
    if childYangName == "frr-remote-lfa-prefixes" {
        return &frrTable.FrrRemoteLfaPrefixes
    }
    if childYangName == "frr-tiebreakers" {
        return &frrTable.FrrTiebreakers
    }
    if childYangName == "frr-use-cand-onlies" {
        return &frrTable.FrrUseCandOnlies
    }
    return nil
}

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frr-load-sharings"] = &frrTable.FrrLoadSharings
    children["priority-limits"] = &frrTable.PriorityLimits
    children["frr-remote-lfa-prefixes"] = &frrTable.FrrRemoteLfaPrefixes
    children["frr-tiebreakers"] = &frrTable.FrrTiebreakers
    children["frr-use-cand-onlies"] = &frrTable.FrrUseCandOnlies
    return children
}

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetBundleName() string { return "cisco_ios_xr" }

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetYangName() string { return "frr-table" }

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) SetParent(parent types.Entity) { frrTable.parent = parent }

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetParent() types.Entity { return frrTable.parent }

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings
// Load share prefixes across multiple
// backups
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable load sharing. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing.
    FrrLoadSharing []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetFilter() yfilter.YFilter { return frrLoadSharings.YFilter }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) SetFilter(yf yfilter.YFilter) { frrLoadSharings.YFilter = yf }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetGoName(yname string) string {
    if yname == "frr-load-sharing" { return "FrrLoadSharing" }
    return ""
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetSegmentPath() string {
    return "frr-load-sharings"
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-load-sharing" {
        for _, c := range frrLoadSharings.FrrLoadSharing {
            if frrLoadSharings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing{}
        frrLoadSharings.FrrLoadSharing = append(frrLoadSharings.FrrLoadSharing, child)
        return &frrLoadSharings.FrrLoadSharing[len(frrLoadSharings.FrrLoadSharing)-1]
    }
    return nil
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrLoadSharings.FrrLoadSharing {
        children[frrLoadSharings.FrrLoadSharing[i].GetSegmentPath()] = &frrLoadSharings.FrrLoadSharing[i]
    }
    return children
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetBundleName() string { return "cisco_ios_xr" }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetYangName() string { return "frr-load-sharings" }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) SetParent(parent types.Entity) { frrLoadSharings.parent = parent }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetParent() types.Entity { return frrLoadSharings.parent }

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing
// Disable load sharing
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Load sharing. The type is IsisfrrLoadSharing. This attribute is mandatory.
    LoadSharing interface{}
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetFilter() yfilter.YFilter { return frrLoadSharing.YFilter }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) SetFilter(yf yfilter.YFilter) { frrLoadSharing.YFilter = yf }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "load-sharing" { return "LoadSharing" }
    return ""
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetSegmentPath() string {
    return "frr-load-sharing" + "[level='" + fmt.Sprintf("%v", frrLoadSharing.Level) + "']"
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrLoadSharing.Level
    leafs["load-sharing"] = frrLoadSharing.LoadSharing
    return leafs
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetBundleName() string { return "cisco_ios_xr" }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetYangName() string { return "frr-load-sharing" }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) SetParent(parent types.Entity) { frrLoadSharing.parent = parent }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetParent() types.Entity { return frrLoadSharing.parent }

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetParentYangName() string { return "frr-load-sharings" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits
// FRR prefix-limit configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Limit backup computation upto the prefix priority. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit.
    PriorityLimit []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetFilter() yfilter.YFilter { return priorityLimits.YFilter }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) SetFilter(yf yfilter.YFilter) { priorityLimits.YFilter = yf }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetGoName(yname string) string {
    if yname == "priority-limit" { return "PriorityLimit" }
    return ""
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetSegmentPath() string {
    return "priority-limits"
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "priority-limit" {
        for _, c := range priorityLimits.PriorityLimit {
            if priorityLimits.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit{}
        priorityLimits.PriorityLimit = append(priorityLimits.PriorityLimit, child)
        return &priorityLimits.PriorityLimit[len(priorityLimits.PriorityLimit)-1]
    }
    return nil
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range priorityLimits.PriorityLimit {
        children[priorityLimits.PriorityLimit[i].GetSegmentPath()] = &priorityLimits.PriorityLimit[i]
    }
    return children
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetBundleName() string { return "cisco_ios_xr" }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetYangName() string { return "priority-limits" }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) SetParent(parent types.Entity) { priorityLimits.parent = parent }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetParent() types.Entity { return priorityLimits.parent }

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit
// Limit backup computation upto the prefix
// priority
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Compute for all prefixes upto the specified priority. The type is
    // IsisPrefixPriority. This attribute is mandatory.
    Priority interface{}
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetFilter() yfilter.YFilter { return priorityLimit.YFilter }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) SetFilter(yf yfilter.YFilter) { priorityLimit.YFilter = yf }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "frr-type" { return "FrrType" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetSegmentPath() string {
    return "priority-limit" + "[level='" + fmt.Sprintf("%v", priorityLimit.Level) + "']" + "[frr-type='" + fmt.Sprintf("%v", priorityLimit.FrrType) + "']"
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = priorityLimit.Level
    leafs["frr-type"] = priorityLimit.FrrType
    leafs["priority"] = priorityLimit.Priority
    return leafs
}

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetBundleName() string { return "cisco_ios_xr" }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetYangName() string { return "priority-limit" }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) SetParent(parent types.Entity) { priorityLimit.parent = parent }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetParent() types.Entity { return priorityLimit.parent }

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetParentYangName() string { return "priority-limits" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes
// FRR remote LFA prefix list filter
// configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Filter remote LFA router IDs using prefix-list. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix.
    FrrRemoteLfaPrefix []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetFilter() yfilter.YFilter { return frrRemoteLfaPrefixes.YFilter }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) SetFilter(yf yfilter.YFilter) { frrRemoteLfaPrefixes.YFilter = yf }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetGoName(yname string) string {
    if yname == "frr-remote-lfa-prefix" { return "FrrRemoteLfaPrefix" }
    return ""
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetSegmentPath() string {
    return "frr-remote-lfa-prefixes"
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-remote-lfa-prefix" {
        for _, c := range frrRemoteLfaPrefixes.FrrRemoteLfaPrefix {
            if frrRemoteLfaPrefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix{}
        frrRemoteLfaPrefixes.FrrRemoteLfaPrefix = append(frrRemoteLfaPrefixes.FrrRemoteLfaPrefix, child)
        return &frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[len(frrRemoteLfaPrefixes.FrrRemoteLfaPrefix)-1]
    }
    return nil
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrRemoteLfaPrefixes.FrrRemoteLfaPrefix {
        children[frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[i].GetSegmentPath()] = &frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[i]
    }
    return children
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetYangName() string { return "frr-remote-lfa-prefixes" }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) SetParent(parent types.Entity) { frrRemoteLfaPrefixes.parent = parent }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetParent() types.Entity { return frrRemoteLfaPrefixes.parent }

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix
// Filter remote LFA router IDs using
// prefix-list
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Name of the prefix list. The type is string with length: 1..32. This
    // attribute is mandatory.
    PrefixListName interface{}
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetFilter() yfilter.YFilter { return frrRemoteLfaPrefix.YFilter }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) SetFilter(yf yfilter.YFilter) { frrRemoteLfaPrefix.YFilter = yf }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "prefix-list-name" { return "PrefixListName" }
    return ""
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetSegmentPath() string {
    return "frr-remote-lfa-prefix" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaPrefix.Level) + "']"
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrRemoteLfaPrefix.Level
    leafs["prefix-list-name"] = frrRemoteLfaPrefix.PrefixListName
    return leafs
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetYangName() string { return "frr-remote-lfa-prefix" }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) SetParent(parent types.Entity) { frrRemoteLfaPrefix.parent = parent }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetParent() types.Entity { return frrRemoteLfaPrefix.parent }

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetParentYangName() string { return "frr-remote-lfa-prefixes" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers
// FRR tiebreakers configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure tiebreaker for multiple backups. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker.
    FrrTiebreaker []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetFilter() yfilter.YFilter { return frrTiebreakers.YFilter }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) SetFilter(yf yfilter.YFilter) { frrTiebreakers.YFilter = yf }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetGoName(yname string) string {
    if yname == "frr-tiebreaker" { return "FrrTiebreaker" }
    return ""
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetSegmentPath() string {
    return "frr-tiebreakers"
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-tiebreaker" {
        for _, c := range frrTiebreakers.FrrTiebreaker {
            if frrTiebreakers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker{}
        frrTiebreakers.FrrTiebreaker = append(frrTiebreakers.FrrTiebreaker, child)
        return &frrTiebreakers.FrrTiebreaker[len(frrTiebreakers.FrrTiebreaker)-1]
    }
    return nil
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrTiebreakers.FrrTiebreaker {
        children[frrTiebreakers.FrrTiebreaker[i].GetSegmentPath()] = &frrTiebreakers.FrrTiebreaker[i]
    }
    return children
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetBundleName() string { return "cisco_ios_xr" }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetYangName() string { return "frr-tiebreakers" }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) SetParent(parent types.Entity) { frrTiebreakers.parent = parent }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetParent() types.Entity { return frrTiebreakers.parent }

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker
// Configure tiebreaker for multiple backups
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Tiebreaker for which configuration applies. The
    // type is IsisfrrTiebreaker.
    Tiebreaker interface{}

    // Preference order among tiebreakers. The type is interface{} with range:
    // 1..255. This attribute is mandatory.
    Index interface{}
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetFilter() yfilter.YFilter { return frrTiebreaker.YFilter }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) SetFilter(yf yfilter.YFilter) { frrTiebreaker.YFilter = yf }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "tiebreaker" { return "Tiebreaker" }
    if yname == "index" { return "Index" }
    return ""
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetSegmentPath() string {
    return "frr-tiebreaker" + "[level='" + fmt.Sprintf("%v", frrTiebreaker.Level) + "']" + "[tiebreaker='" + fmt.Sprintf("%v", frrTiebreaker.Tiebreaker) + "']"
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrTiebreaker.Level
    leafs["tiebreaker"] = frrTiebreaker.Tiebreaker
    leafs["index"] = frrTiebreaker.Index
    return leafs
}

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetBundleName() string { return "cisco_ios_xr" }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetYangName() string { return "frr-tiebreaker" }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) SetParent(parent types.Entity) { frrTiebreaker.parent = parent }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetParent() types.Entity { return frrTiebreaker.parent }

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetParentYangName() string { return "frr-tiebreakers" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies
// FRR use candidate only configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure use candidate only to exclude interfaces as backup. The type is
    // slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly.
    FrrUseCandOnly []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetFilter() yfilter.YFilter { return frrUseCandOnlies.YFilter }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) SetFilter(yf yfilter.YFilter) { frrUseCandOnlies.YFilter = yf }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetGoName(yname string) string {
    if yname == "frr-use-cand-only" { return "FrrUseCandOnly" }
    return ""
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetSegmentPath() string {
    return "frr-use-cand-onlies"
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-use-cand-only" {
        for _, c := range frrUseCandOnlies.FrrUseCandOnly {
            if frrUseCandOnlies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly{}
        frrUseCandOnlies.FrrUseCandOnly = append(frrUseCandOnlies.FrrUseCandOnly, child)
        return &frrUseCandOnlies.FrrUseCandOnly[len(frrUseCandOnlies.FrrUseCandOnly)-1]
    }
    return nil
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrUseCandOnlies.FrrUseCandOnly {
        children[frrUseCandOnlies.FrrUseCandOnly[i].GetSegmentPath()] = &frrUseCandOnlies.FrrUseCandOnly[i]
    }
    return children
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetBundleName() string { return "cisco_ios_xr" }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetYangName() string { return "frr-use-cand-onlies" }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) SetParent(parent types.Entity) { frrUseCandOnlies.parent = parent }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetParent() types.Entity { return frrUseCandOnlies.parent }

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetParentYangName() string { return "frr-table" }

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly
// Configure use candidate only to exclude
// interfaces as backup
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetFilter() yfilter.YFilter { return frrUseCandOnly.YFilter }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) SetFilter(yf yfilter.YFilter) { frrUseCandOnly.YFilter = yf }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "frr-type" { return "FrrType" }
    return ""
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetSegmentPath() string {
    return "frr-use-cand-only" + "[level='" + fmt.Sprintf("%v", frrUseCandOnly.Level) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrUseCandOnly.FrrType) + "']"
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrUseCandOnly.Level
    leafs["frr-type"] = frrUseCandOnly.FrrType
    return leafs
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetBundleName() string { return "cisco_ios_xr" }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetYangName() string { return "frr-use-cand-only" }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) SetParent(parent types.Entity) { frrUseCandOnly.parent = parent }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetParent() types.Entity { return frrUseCandOnly.parent }

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetParentYangName() string { return "frr-use-cand-onlies" }

// Isis_Instances_Instance_Afs_Af_TopologyName_RouterId
// Stable IP address for system. Will only be
// applied for the unicast sub-address-family.
type Isis_Instances_Instance_Afs_Af_TopologyName_RouterId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4/IPv6 address to be used as a router ID. Precisely one of Address and
    // Interface must be specified. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Interface with designated stable IP address to be used as a router ID. This
    // must be a Loopback interface. Precisely one of Address and Interface must
    // be specified. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetFilter() yfilter.YFilter { return routerId.YFilter }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) SetFilter(yf yfilter.YFilter) { routerId.YFilter = yf }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetSegmentPath() string {
    return "router-id"
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = routerId.Address
    leafs["interface-name"] = routerId.InterfaceName
    return leafs
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetBundleName() string { return "cisco_ios_xr" }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetYangName() string { return "router-id" }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) SetParent(parent types.Entity) { routerId.parent = parent }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetParent() types.Entity { return routerId.parent }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities
// SPF Prefix Priority configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Determine SPF priority for prefixes. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority.
    SpfPrefixPriority []Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetFilter() yfilter.YFilter { return spfPrefixPriorities.YFilter }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) SetFilter(yf yfilter.YFilter) { spfPrefixPriorities.YFilter = yf }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetGoName(yname string) string {
    if yname == "spf-prefix-priority" { return "SpfPrefixPriority" }
    return ""
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetSegmentPath() string {
    return "spf-prefix-priorities"
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spf-prefix-priority" {
        for _, c := range spfPrefixPriorities.SpfPrefixPriority {
            if spfPrefixPriorities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority{}
        spfPrefixPriorities.SpfPrefixPriority = append(spfPrefixPriorities.SpfPrefixPriority, child)
        return &spfPrefixPriorities.SpfPrefixPriority[len(spfPrefixPriorities.SpfPrefixPriority)-1]
    }
    return nil
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spfPrefixPriorities.SpfPrefixPriority {
        children[spfPrefixPriorities.SpfPrefixPriority[i].GetSegmentPath()] = &spfPrefixPriorities.SpfPrefixPriority[i]
    }
    return children
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetBundleName() string { return "cisco_ios_xr" }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetYangName() string { return "spf-prefix-priorities" }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) SetParent(parent types.Entity) { spfPrefixPriorities.parent = parent }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetParent() types.Entity { return spfPrefixPriorities.parent }

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority
// Determine SPF priority for prefixes
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SPF Level for prefix prioritization. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. SPF Priority to assign matching prefixes. The type
    // is IsisPrefixPriority.
    PrefixPriorityType interface{}

    // Tag value to determine prefixes for this priority. The type is interface{}
    // with range: 1..4294967295.
    AdminTag interface{}

    // Access List to determine prefixes for this priority. The type is string
    // with length: 1..64.
    AccessListName interface{}
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetFilter() yfilter.YFilter { return spfPrefixPriority.YFilter }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) SetFilter(yf yfilter.YFilter) { spfPrefixPriority.YFilter = yf }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "prefix-priority-type" { return "PrefixPriorityType" }
    if yname == "admin-tag" { return "AdminTag" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetSegmentPath() string {
    return "spf-prefix-priority" + "[level='" + fmt.Sprintf("%v", spfPrefixPriority.Level) + "']" + "[prefix-priority-type='" + fmt.Sprintf("%v", spfPrefixPriority.PrefixPriorityType) + "']"
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = spfPrefixPriority.Level
    leafs["prefix-priority-type"] = spfPrefixPriority.PrefixPriorityType
    leafs["admin-tag"] = spfPrefixPriority.AdminTag
    leafs["access-list-name"] = spfPrefixPriority.AccessListName
    return leafs
}

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetBundleName() string { return "cisco_ios_xr" }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetYangName() string { return "spf-prefix-priority" }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) SetParent(parent types.Entity) { spfPrefixPriority.parent = parent }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetParent() types.Entity { return spfPrefixPriority.parent }

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetParentYangName() string { return "spf-prefix-priorities" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes
// Summary-prefix configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure IP address prefixes to advertise. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix.
    SummaryPrefix []Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetFilter() yfilter.YFilter { return summaryPrefixes.YFilter }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) SetFilter(yf yfilter.YFilter) { summaryPrefixes.YFilter = yf }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetGoName(yname string) string {
    if yname == "summary-prefix" { return "SummaryPrefix" }
    return ""
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetSegmentPath() string {
    return "summary-prefixes"
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary-prefix" {
        for _, c := range summaryPrefixes.SummaryPrefix {
            if summaryPrefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix{}
        summaryPrefixes.SummaryPrefix = append(summaryPrefixes.SummaryPrefix, child)
        return &summaryPrefixes.SummaryPrefix[len(summaryPrefixes.SummaryPrefix)-1]
    }
    return nil
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summaryPrefixes.SummaryPrefix {
        children[summaryPrefixes.SummaryPrefix[i].GetSegmentPath()] = &summaryPrefixes.SummaryPrefix[i]
    }
    return children
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetBundleName() string { return "cisco_ios_xr" }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetYangName() string { return "summary-prefixes" }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) SetParent(parent types.Entity) { summaryPrefixes.parent = parent }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetParent() types.Entity { return summaryPrefixes.parent }

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix
// Configure IP address prefixes to advertise
type Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP summary address prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}

    // The tag value. The type is interface{} with range: 1..4294967295.
    Tag interface{}

    // Level in which to summarize routes. The type is interface{} with range:
    // 1..2.
    Level interface{}
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetFilter() yfilter.YFilter { return summaryPrefix.YFilter }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) SetFilter(yf yfilter.YFilter) { summaryPrefix.YFilter = yf }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetGoName(yname string) string {
    if yname == "address-prefix" { return "AddressPrefix" }
    if yname == "tag" { return "Tag" }
    if yname == "level" { return "Level" }
    return ""
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetSegmentPath() string {
    return "summary-prefix" + "[address-prefix='" + fmt.Sprintf("%v", summaryPrefix.AddressPrefix) + "']"
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-prefix"] = summaryPrefix.AddressPrefix
    leafs["tag"] = summaryPrefix.Tag
    leafs["level"] = summaryPrefix.Level
    return leafs
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetYangName() string { return "summary-prefix" }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) SetParent(parent types.Entity) { summaryPrefix.parent = parent }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetParent() types.Entity { return summaryPrefix.parent }

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetParentYangName() string { return "summary-prefixes" }

// Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance
// Micro Loop Avoidance configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MicroLoop avoidance enable configuration. The type is
    // IsisMicroLoopAvoidance. The default value is micro-loop-avoidance-all.
    Enable interface{}

    // Value of delay in msecs in updating RIB. The type is interface{} with
    // range: 1000..65535. Units are millisecond. The default value is 5000.
    RibUpdateDelay interface{}
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetFilter() yfilter.YFilter { return microLoopAvoidance.YFilter }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) SetFilter(yf yfilter.YFilter) { microLoopAvoidance.YFilter = yf }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "rib-update-delay" { return "RibUpdateDelay" }
    return ""
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetSegmentPath() string {
    return "micro-loop-avoidance"
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = microLoopAvoidance.Enable
    leafs["rib-update-delay"] = microLoopAvoidance.RibUpdateDelay
    return leafs
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetBundleName() string { return "cisco_ios_xr" }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetYangName() string { return "micro-loop-avoidance" }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) SetParent(parent types.Entity) { microLoopAvoidance.parent = parent }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetParent() types.Entity { return microLoopAvoidance.parent }

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp
// UCMP (UnEqual Cost MultiPath) configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delay in msecs between primary SPF and UCMP computation. The type is
    // interface{} with range: 100..65535. Units are millisecond. The default
    // value is 100.
    DelayInterval interface{}

    // UCMP feature enable configuration.
    Enable Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable

    // Interfaces excluded from UCMP path computation.
    ExcludeInterfaces Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces
}

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetFilter() yfilter.YFilter { return ucmp.YFilter }

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) SetFilter(yf yfilter.YFilter) { ucmp.YFilter = yf }

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetGoName(yname string) string {
    if yname == "delay-interval" { return "DelayInterval" }
    if yname == "enable" { return "Enable" }
    if yname == "exclude-interfaces" { return "ExcludeInterfaces" }
    return ""
}

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetSegmentPath() string {
    return "ucmp"
}

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "enable" {
        return &ucmp.Enable
    }
    if childYangName == "exclude-interfaces" {
        return &ucmp.ExcludeInterfaces
    }
    return nil
}

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["enable"] = &ucmp.Enable
    children["exclude-interfaces"] = &ucmp.ExcludeInterfaces
    return children
}

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["delay-interval"] = ucmp.DelayInterval
    return leafs
}

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetBundleName() string { return "cisco_ios_xr" }

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetYangName() string { return "ucmp" }

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) SetParent(parent types.Entity) { ucmp.parent = parent }

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetParent() types.Entity { return ucmp.parent }

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable
// UCMP feature enable configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Value of variance. The type is interface{} with range: 101..10000. The
    // default value is 200.
    Variance interface{}

    // Name of the Prefix List. The type is string with length: 1..32.
    PrefixListName interface{}
}

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetFilter() yfilter.YFilter { return enable.YFilter }

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) SetFilter(yf yfilter.YFilter) { enable.YFilter = yf }

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetGoName(yname string) string {
    if yname == "variance" { return "Variance" }
    if yname == "prefix-list-name" { return "PrefixListName" }
    return ""
}

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetSegmentPath() string {
    return "enable"
}

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["variance"] = enable.Variance
    leafs["prefix-list-name"] = enable.PrefixListName
    return leafs
}

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetBundleName() string { return "cisco_ios_xr" }

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetYangName() string { return "enable" }

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) SetParent(parent types.Entity) { enable.parent = parent }

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetParent() types.Entity { return enable.parent }

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetParentYangName() string { return "ucmp" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces
// Interfaces excluded from UCMP path
// computation
type Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exclude this interface from UCMP path computation. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetFilter() yfilter.YFilter { return excludeInterfaces.YFilter }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) SetFilter(yf yfilter.YFilter) { excludeInterfaces.YFilter = yf }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetGoName(yname string) string {
    if yname == "exclude-interface" { return "ExcludeInterface" }
    return ""
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetSegmentPath() string {
    return "exclude-interfaces"
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "exclude-interface" {
        for _, c := range excludeInterfaces.ExcludeInterface {
            if excludeInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface{}
        excludeInterfaces.ExcludeInterface = append(excludeInterfaces.ExcludeInterface, child)
        return &excludeInterfaces.ExcludeInterface[len(excludeInterfaces.ExcludeInterface)-1]
    }
    return nil
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range excludeInterfaces.ExcludeInterface {
        children[excludeInterfaces.ExcludeInterface[i].GetSegmentPath()] = &excludeInterfaces.ExcludeInterface[i]
    }
    return children
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetYangName() string { return "exclude-interfaces" }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) SetParent(parent types.Entity) { excludeInterfaces.parent = parent }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetParent() types.Entity { return excludeInterfaces.parent }

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetParentYangName() string { return "ucmp" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface
// Exclude this interface from UCMP path
// computation
type Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface to be excluded. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetFilter() yfilter.YFilter { return excludeInterface.YFilter }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) SetFilter(yf yfilter.YFilter) { excludeInterface.YFilter = yf }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetSegmentPath() string {
    return "exclude-interface" + "[interface-name='" + fmt.Sprintf("%v", excludeInterface.InterfaceName) + "']"
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = excludeInterface.InterfaceName
    return leafs
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetBundleName() string { return "cisco_ios_xr" }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetYangName() string { return "exclude-interface" }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) SetParent(parent types.Entity) { excludeInterface.parent = parent }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetParent() types.Entity { return excludeInterface.parent }

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetParentYangName() string { return "exclude-interfaces" }

// Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes
// Maximum number of redistributed
// prefixesconfiguration
type Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An upper limit on the number of redistributed prefixes which may be
    // included in the local system's LSP. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix.
    MaxRedistPrefix []Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetFilter() yfilter.YFilter { return maxRedistPrefixes.YFilter }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) SetFilter(yf yfilter.YFilter) { maxRedistPrefixes.YFilter = yf }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetGoName(yname string) string {
    if yname == "max-redist-prefix" { return "MaxRedistPrefix" }
    return ""
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetSegmentPath() string {
    return "max-redist-prefixes"
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "max-redist-prefix" {
        for _, c := range maxRedistPrefixes.MaxRedistPrefix {
            if maxRedistPrefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix{}
        maxRedistPrefixes.MaxRedistPrefix = append(maxRedistPrefixes.MaxRedistPrefix, child)
        return &maxRedistPrefixes.MaxRedistPrefix[len(maxRedistPrefixes.MaxRedistPrefix)-1]
    }
    return nil
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range maxRedistPrefixes.MaxRedistPrefix {
        children[maxRedistPrefixes.MaxRedistPrefix[i].GetSegmentPath()] = &maxRedistPrefixes.MaxRedistPrefix[i]
    }
    return children
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetBundleName() string { return "cisco_ios_xr" }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetYangName() string { return "max-redist-prefixes" }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) SetParent(parent types.Entity) { maxRedistPrefixes.parent = parent }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetParent() types.Entity { return maxRedistPrefixes.parent }

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix
// An upper limit on the number of
// redistributed prefixes which may be
// included in the local system's LSP
type Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Max number of prefixes. The type is interface{} with range: 1..28000. This
    // attribute is mandatory.
    PrefixLimit interface{}
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetFilter() yfilter.YFilter { return maxRedistPrefix.YFilter }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) SetFilter(yf yfilter.YFilter) { maxRedistPrefix.YFilter = yf }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "prefix-limit" { return "PrefixLimit" }
    return ""
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetSegmentPath() string {
    return "max-redist-prefix" + "[level='" + fmt.Sprintf("%v", maxRedistPrefix.Level) + "']"
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = maxRedistPrefix.Level
    leafs["prefix-limit"] = maxRedistPrefix.PrefixLimit
    return leafs
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetYangName() string { return "max-redist-prefix" }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) SetParent(parent types.Entity) { maxRedistPrefix.parent = parent }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetParent() types.Entity { return maxRedistPrefix.parent }

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetParentYangName() string { return "max-redist-prefixes" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Propagations
// Route propagation configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Propagations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Propagate routes between IS-IS levels. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation.
    Propagation []Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation
}

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetFilter() yfilter.YFilter { return propagations.YFilter }

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) SetFilter(yf yfilter.YFilter) { propagations.YFilter = yf }

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetGoName(yname string) string {
    if yname == "propagation" { return "Propagation" }
    return ""
}

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetSegmentPath() string {
    return "propagations"
}

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "propagation" {
        for _, c := range propagations.Propagation {
            if propagations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation{}
        propagations.Propagation = append(propagations.Propagation, child)
        return &propagations.Propagation[len(propagations.Propagation)-1]
    }
    return nil
}

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range propagations.Propagation {
        children[propagations.Propagation[i].GetSegmentPath()] = &propagations.Propagation[i]
    }
    return children
}

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetBundleName() string { return "cisco_ios_xr" }

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetYangName() string { return "propagations" }

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) SetParent(parent types.Entity) { propagations.parent = parent }

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetParent() types.Entity { return propagations.parent }

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation
// Propagate routes between IS-IS levels
type Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Source level for routes. The type is
    // IsisInternalLevel.
    SourceLevel interface{}

    // This attribute is a key. Destination level for routes.  Must differ from
    // SourceLevel. The type is IsisInternalLevel.
    DestinationLevel interface{}

    // Route policy limiting routes to be propagated. The type is string with
    // length: 1..64. This attribute is mandatory.
    RoutePolicyName interface{}
}

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetFilter() yfilter.YFilter { return propagation.YFilter }

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) SetFilter(yf yfilter.YFilter) { propagation.YFilter = yf }

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetGoName(yname string) string {
    if yname == "source-level" { return "SourceLevel" }
    if yname == "destination-level" { return "DestinationLevel" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    return ""
}

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetSegmentPath() string {
    return "propagation" + "[source-level='" + fmt.Sprintf("%v", propagation.SourceLevel) + "']" + "[destination-level='" + fmt.Sprintf("%v", propagation.DestinationLevel) + "']"
}

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-level"] = propagation.SourceLevel
    leafs["destination-level"] = propagation.DestinationLevel
    leafs["route-policy-name"] = propagation.RoutePolicyName
    return leafs
}

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetBundleName() string { return "cisco_ios_xr" }

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetYangName() string { return "propagation" }

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) SetParent(parent types.Entity) { propagation.parent = parent }

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetParent() types.Entity { return propagation.parent }

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetParentYangName() string { return "propagations" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions
// Protocol redistribution configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribution of other protocols into this IS-IS instance. The type is
    // slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution.
    Redistribution []Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution
}

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetFilter() yfilter.YFilter { return redistributions.YFilter }

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) SetFilter(yf yfilter.YFilter) { redistributions.YFilter = yf }

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetGoName(yname string) string {
    if yname == "redistribution" { return "Redistribution" }
    return ""
}

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetSegmentPath() string {
    return "redistributions"
}

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "redistribution" {
        for _, c := range redistributions.Redistribution {
            if redistributions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution{}
        redistributions.Redistribution = append(redistributions.Redistribution, child)
        return &redistributions.Redistribution[len(redistributions.Redistribution)-1]
    }
    return nil
}

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range redistributions.Redistribution {
        children[redistributions.Redistribution[i].GetSegmentPath()] = &redistributions.Redistribution[i]
    }
    return children
}

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetBundleName() string { return "cisco_ios_xr" }

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetYangName() string { return "redistributions" }

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) SetParent(parent types.Entity) { redistributions.parent = parent }

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetParent() types.Entity { return redistributions.parent }

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution
// Redistribution of other protocols into
// this IS-IS instance
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The protocol to be redistributed.  OSPFv3 may not
    // be specified for an IPv4 topology and OSPF may not be specified for an IPv6
    // topology. The type is IsisRedistProto.
    ProtocolName interface{}

    // connected or static or rip or subscriber or mobile.
    ConnectedOrStaticOrRipOrSubscriberOrMobile Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile

    // ospf or ospfv3 or isis or application. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication.
    OspfOrOspfv3OrIsisOrApplication []Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication

    // bgp. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp.
    Bgp []Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp

    // eigrp. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp.
    Eigrp []Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp
}

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetFilter() yfilter.YFilter { return redistribution.YFilter }

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) SetFilter(yf yfilter.YFilter) { redistribution.YFilter = yf }

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetGoName(yname string) string {
    if yname == "protocol-name" { return "ProtocolName" }
    if yname == "connected-or-static-or-rip-or-subscriber-or-mobile" { return "ConnectedOrStaticOrRipOrSubscriberOrMobile" }
    if yname == "ospf-or-ospfv3-or-isis-or-application" { return "OspfOrOspfv3OrIsisOrApplication" }
    if yname == "bgp" { return "Bgp" }
    if yname == "eigrp" { return "Eigrp" }
    return ""
}

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetSegmentPath() string {
    return "redistribution" + "[protocol-name='" + fmt.Sprintf("%v", redistribution.ProtocolName) + "']"
}

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connected-or-static-or-rip-or-subscriber-or-mobile" {
        return &redistribution.ConnectedOrStaticOrRipOrSubscriberOrMobile
    }
    if childYangName == "ospf-or-ospfv3-or-isis-or-application" {
        for _, c := range redistribution.OspfOrOspfv3OrIsisOrApplication {
            if redistribution.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication{}
        redistribution.OspfOrOspfv3OrIsisOrApplication = append(redistribution.OspfOrOspfv3OrIsisOrApplication, child)
        return &redistribution.OspfOrOspfv3OrIsisOrApplication[len(redistribution.OspfOrOspfv3OrIsisOrApplication)-1]
    }
    if childYangName == "bgp" {
        for _, c := range redistribution.Bgp {
            if redistribution.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp{}
        redistribution.Bgp = append(redistribution.Bgp, child)
        return &redistribution.Bgp[len(redistribution.Bgp)-1]
    }
    if childYangName == "eigrp" {
        for _, c := range redistribution.Eigrp {
            if redistribution.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp{}
        redistribution.Eigrp = append(redistribution.Eigrp, child)
        return &redistribution.Eigrp[len(redistribution.Eigrp)-1]
    }
    return nil
}

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["connected-or-static-or-rip-or-subscriber-or-mobile"] = &redistribution.ConnectedOrStaticOrRipOrSubscriberOrMobile
    for i := range redistribution.OspfOrOspfv3OrIsisOrApplication {
        children[redistribution.OspfOrOspfv3OrIsisOrApplication[i].GetSegmentPath()] = &redistribution.OspfOrOspfv3OrIsisOrApplication[i]
    }
    for i := range redistribution.Bgp {
        children[redistribution.Bgp[i].GetSegmentPath()] = &redistribution.Bgp[i]
    }
    for i := range redistribution.Eigrp {
        children[redistribution.Eigrp[i].GetSegmentPath()] = &redistribution.Eigrp[i]
    }
    return children
}

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocol-name"] = redistribution.ProtocolName
    return leafs
}

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetBundleName() string { return "cisco_ios_xr" }

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetYangName() string { return "redistribution" }

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) SetParent(parent types.Entity) { redistribution.parent = parent }

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetParent() types.Entity { return redistribution.parent }

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetParentYangName() string { return "redistributions" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile
// connected or static or rip or subscriber
// or mobile
// This type is a presence type.
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric for redistributed routes: <0-63> for narrow, <0-16777215> for wide.
    // The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Levels to redistribute routes into. The type is IsisConfigurableLevels.
    Levels interface{}

    // Route policy to control redistribution. The type is string with length:
    // 1..64.
    RoutePolicyName interface{}

    // IS-IS metric type. The type is IsisMetric.
    MetricType interface{}

    // OSPF route types to redistribute.  May only be specified if Protocol is
    // OSPF. The type is interface{} with range: -2147483648..2147483647.
    OspfRouteType interface{}
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetFilter() yfilter.YFilter { return connectedOrStaticOrRipOrSubscriberOrMobile.YFilter }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) SetFilter(yf yfilter.YFilter) { connectedOrStaticOrRipOrSubscriberOrMobile.YFilter = yf }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    if yname == "levels" { return "Levels" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "metric-type" { return "MetricType" }
    if yname == "ospf-route-type" { return "OspfRouteType" }
    return ""
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetSegmentPath() string {
    return "connected-or-static-or-rip-or-subscriber-or-mobile"
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["metric"] = connectedOrStaticOrRipOrSubscriberOrMobile.Metric
    leafs["levels"] = connectedOrStaticOrRipOrSubscriberOrMobile.Levels
    leafs["route-policy-name"] = connectedOrStaticOrRipOrSubscriberOrMobile.RoutePolicyName
    leafs["metric-type"] = connectedOrStaticOrRipOrSubscriberOrMobile.MetricType
    leafs["ospf-route-type"] = connectedOrStaticOrRipOrSubscriberOrMobile.OspfRouteType
    return leafs
}

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetBundleName() string { return "cisco_ios_xr" }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetYangName() string { return "connected-or-static-or-rip-or-subscriber-or-mobile" }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) SetParent(parent types.Entity) { connectedOrStaticOrRipOrSubscriberOrMobile.parent = parent }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetParent() types.Entity { return connectedOrStaticOrRipOrSubscriberOrMobile.parent }

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetParentYangName() string { return "redistribution" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication
// ospf or ospfv3 or isis or application
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol Instance Identifier.  Mandatory for ISIS,
    // OSPF and application, must not be specified otherwise. The type is string
    // with length: 1..64.
    InstanceName interface{}

    // Metric for redistributed routes: <0-63> for narrow, <0-16777215> for wide.
    // The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Levels to redistribute routes into. The type is IsisConfigurableLevels.
    Levels interface{}

    // Route policy to control redistribution. The type is string with length:
    // 1..64.
    RoutePolicyName interface{}

    // IS-IS metric type. The type is IsisMetric.
    MetricType interface{}

    // OSPF route types to redistribute.  May only be specified if Protocol is
    // OSPF. The type is interface{} with range: -2147483648..2147483647.
    OspfRouteType interface{}
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetFilter() yfilter.YFilter { return ospfOrOspfv3OrIsisOrApplication.YFilter }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) SetFilter(yf yfilter.YFilter) { ospfOrOspfv3OrIsisOrApplication.YFilter = yf }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetGoName(yname string) string {
    if yname == "instance-name" { return "InstanceName" }
    if yname == "metric" { return "Metric" }
    if yname == "levels" { return "Levels" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "metric-type" { return "MetricType" }
    if yname == "ospf-route-type" { return "OspfRouteType" }
    return ""
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetSegmentPath() string {
    return "ospf-or-ospfv3-or-isis-or-application" + "[instance-name='" + fmt.Sprintf("%v", ospfOrOspfv3OrIsisOrApplication.InstanceName) + "']"
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-name"] = ospfOrOspfv3OrIsisOrApplication.InstanceName
    leafs["metric"] = ospfOrOspfv3OrIsisOrApplication.Metric
    leafs["levels"] = ospfOrOspfv3OrIsisOrApplication.Levels
    leafs["route-policy-name"] = ospfOrOspfv3OrIsisOrApplication.RoutePolicyName
    leafs["metric-type"] = ospfOrOspfv3OrIsisOrApplication.MetricType
    leafs["ospf-route-type"] = ospfOrOspfv3OrIsisOrApplication.OspfRouteType
    return leafs
}

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetBundleName() string { return "cisco_ios_xr" }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetYangName() string { return "ospf-or-ospfv3-or-isis-or-application" }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) SetParent(parent types.Entity) { ospfOrOspfv3OrIsisOrApplication.parent = parent }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetParent() types.Entity { return ospfOrOspfv3OrIsisOrApplication.parent }

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetParentYangName() string { return "redistribution" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp
// bgp
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. First half of BGP AS number in XX.YY format. 
    // Mandatory if Protocol is BGP and must not be specified otherwise. Must be a
    // non-zero value if second half is zero. The type is interface{} with range:
    // 0..65535.
    AsXx interface{}

    // This attribute is a key. Second half of BGP AS number in XX.YY format.
    // Mandatory if Protocol is BGP and must not be specified otherwise. Must be a
    // non-zero value if first half is zero. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}

    // Metric for redistributed routes: <0-63> for narrow, <0-16777215> for wide.
    // The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Levels to redistribute routes into. The type is IsisConfigurableLevels.
    Levels interface{}

    // Route policy to control redistribution. The type is string with length:
    // 1..64.
    RoutePolicyName interface{}

    // IS-IS metric type. The type is IsisMetric.
    MetricType interface{}

    // OSPF route types to redistribute.  May only be specified if Protocol is
    // OSPF. The type is interface{} with range: -2147483648..2147483647.
    OspfRouteType interface{}
}

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetGoName(yname string) string {
    if yname == "as-xx" { return "AsXx" }
    if yname == "as-yy" { return "AsYy" }
    if yname == "metric" { return "Metric" }
    if yname == "levels" { return "Levels" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "metric-type" { return "MetricType" }
    if yname == "ospf-route-type" { return "OspfRouteType" }
    return ""
}

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetSegmentPath() string {
    return "bgp" + "[as-xx='" + fmt.Sprintf("%v", bgp.AsXx) + "']" + "[as-yy='" + fmt.Sprintf("%v", bgp.AsYy) + "']"
}

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-xx"] = bgp.AsXx
    leafs["as-yy"] = bgp.AsYy
    leafs["metric"] = bgp.Metric
    leafs["levels"] = bgp.Levels
    leafs["route-policy-name"] = bgp.RoutePolicyName
    leafs["metric-type"] = bgp.MetricType
    leafs["ospf-route-type"] = bgp.OspfRouteType
    return leafs
}

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetYangName() string { return "bgp" }

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetParentYangName() string { return "redistribution" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp
// eigrp
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Eigrp as number. The type is interface{} with
    // range: 1..65535.
    AsZz interface{}

    // Metric for redistributed routes: <0-63> for narrow, <0-16777215> for wide.
    // The type is interface{} with range: 0..16777215.
    Metric interface{}

    // Levels to redistribute routes into. The type is IsisConfigurableLevels.
    Levels interface{}

    // Route policy to control redistribution. The type is string with length:
    // 1..64.
    RoutePolicyName interface{}

    // IS-IS metric type. The type is IsisMetric.
    MetricType interface{}

    // OSPF route types to redistribute.  May only be specified if Protocol is
    // OSPF. The type is interface{} with range: -2147483648..2147483647.
    OspfRouteType interface{}
}

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetFilter() yfilter.YFilter { return eigrp.YFilter }

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) SetFilter(yf yfilter.YFilter) { eigrp.YFilter = yf }

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetGoName(yname string) string {
    if yname == "as-zz" { return "AsZz" }
    if yname == "metric" { return "Metric" }
    if yname == "levels" { return "Levels" }
    if yname == "route-policy-name" { return "RoutePolicyName" }
    if yname == "metric-type" { return "MetricType" }
    if yname == "ospf-route-type" { return "OspfRouteType" }
    return ""
}

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetSegmentPath() string {
    return "eigrp" + "[as-zz='" + fmt.Sprintf("%v", eigrp.AsZz) + "']"
}

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-zz"] = eigrp.AsZz
    leafs["metric"] = eigrp.Metric
    leafs["levels"] = eigrp.Levels
    leafs["route-policy-name"] = eigrp.RoutePolicyName
    leafs["metric-type"] = eigrp.MetricType
    leafs["ospf-route-type"] = eigrp.OspfRouteType
    return leafs
}

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetBundleName() string { return "cisco_ios_xr" }

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetYangName() string { return "eigrp" }

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) SetParent(parent types.Entity) { eigrp.parent = parent }

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetParent() types.Entity { return eigrp.parent }

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetParentYangName() string { return "redistribution" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals
// Peoridic SPF configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum interval between spf runs. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval.
    SpfPeriodicInterval []Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetFilter() yfilter.YFilter { return spfPeriodicIntervals.YFilter }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) SetFilter(yf yfilter.YFilter) { spfPeriodicIntervals.YFilter = yf }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetGoName(yname string) string {
    if yname == "spf-periodic-interval" { return "SpfPeriodicInterval" }
    return ""
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetSegmentPath() string {
    return "spf-periodic-intervals"
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spf-periodic-interval" {
        for _, c := range spfPeriodicIntervals.SpfPeriodicInterval {
            if spfPeriodicIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval{}
        spfPeriodicIntervals.SpfPeriodicInterval = append(spfPeriodicIntervals.SpfPeriodicInterval, child)
        return &spfPeriodicIntervals.SpfPeriodicInterval[len(spfPeriodicIntervals.SpfPeriodicInterval)-1]
    }
    return nil
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spfPeriodicIntervals.SpfPeriodicInterval {
        children[spfPeriodicIntervals.SpfPeriodicInterval[i].GetSegmentPath()] = &spfPeriodicIntervals.SpfPeriodicInterval[i]
    }
    return children
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetYangName() string { return "spf-periodic-intervals" }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) SetParent(parent types.Entity) { spfPeriodicIntervals.parent = parent }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetParent() types.Entity { return spfPeriodicIntervals.parent }

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval
// Maximum interval between spf runs
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Maximum interval in between SPF runs in seconds. The type is interface{}
    // with range: 0..3600. This attribute is mandatory. Units are second.
    PeriodicInterval interface{}
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetFilter() yfilter.YFilter { return spfPeriodicInterval.YFilter }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) SetFilter(yf yfilter.YFilter) { spfPeriodicInterval.YFilter = yf }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "periodic-interval" { return "PeriodicInterval" }
    return ""
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetSegmentPath() string {
    return "spf-periodic-interval" + "[level='" + fmt.Sprintf("%v", spfPeriodicInterval.Level) + "']"
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = spfPeriodicInterval.Level
    leafs["periodic-interval"] = spfPeriodicInterval.PeriodicInterval
    return leafs
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetBundleName() string { return "cisco_ios_xr" }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetYangName() string { return "spf-periodic-interval" }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) SetParent(parent types.Entity) { spfPeriodicInterval.parent = parent }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetParent() types.Entity { return spfPeriodicInterval.parent }

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetParentYangName() string { return "spf-periodic-intervals" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals
// SPF-interval configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route calculation scheduling parameters. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval.
    SpfInterval []Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetFilter() yfilter.YFilter { return spfIntervals.YFilter }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) SetFilter(yf yfilter.YFilter) { spfIntervals.YFilter = yf }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetGoName(yname string) string {
    if yname == "spf-interval" { return "SpfInterval" }
    return ""
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetSegmentPath() string {
    return "spf-intervals"
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "spf-interval" {
        for _, c := range spfIntervals.SpfInterval {
            if spfIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval{}
        spfIntervals.SpfInterval = append(spfIntervals.SpfInterval, child)
        return &spfIntervals.SpfInterval[len(spfIntervals.SpfInterval)-1]
    }
    return nil
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spfIntervals.SpfInterval {
        children[spfIntervals.SpfInterval[i].GetSegmentPath()] = &spfIntervals.SpfInterval[i]
    }
    return children
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetYangName() string { return "spf-intervals" }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) SetParent(parent types.Entity) { spfIntervals.parent = parent }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetParent() types.Entity { return spfIntervals.parent }

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval
// Route calculation scheduling parameters
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Maximum wait before running a route calculation in milliseconds. The type
    // is interface{} with range: 0..120000. Units are millisecond.
    MaximumWait interface{}

    // Initial wait before running a route calculation in milliseconds. The type
    // is interface{} with range: 0..120000. Units are millisecond.
    InitialWait interface{}

    // Secondary wait before running a route calculation in milliseconds. The type
    // is interface{} with range: 0..120000. Units are millisecond.
    SecondaryWait interface{}
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetFilter() yfilter.YFilter { return spfInterval.YFilter }

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) SetFilter(yf yfilter.YFilter) { spfInterval.YFilter = yf }

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "maximum-wait" { return "MaximumWait" }
    if yname == "initial-wait" { return "InitialWait" }
    if yname == "secondary-wait" { return "SecondaryWait" }
    return ""
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetSegmentPath() string {
    return "spf-interval" + "[level='" + fmt.Sprintf("%v", spfInterval.Level) + "']"
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = spfInterval.Level
    leafs["maximum-wait"] = spfInterval.MaximumWait
    leafs["initial-wait"] = spfInterval.InitialWait
    leafs["secondary-wait"] = spfInterval.SecondaryWait
    return leafs
}

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetBundleName() string { return "cisco_ios_xr" }

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetYangName() string { return "spf-interval" }

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) SetParent(parent types.Entity) { spfInterval.parent = parent }

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetParent() types.Entity { return spfInterval.parent }

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetParentYangName() string { return "spf-intervals" }

// Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence
// Enable convergence monitoring
type Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable convergence monitoring. The type is interface{}.
    Enable interface{}

    // Enable the Tracking of IP-Frr Convergence. The type is interface{}.
    TrackIpFrr interface{}

    // Enable the monitoring of individual prefixes (prefix list name). The type
    // is string with length: 1..32.
    PrefixList interface{}
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetFilter() yfilter.YFilter { return monitorConvergence.YFilter }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) SetFilter(yf yfilter.YFilter) { monitorConvergence.YFilter = yf }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "track-ip-frr" { return "TrackIpFrr" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetSegmentPath() string {
    return "monitor-convergence"
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = monitorConvergence.Enable
    leafs["track-ip-frr"] = monitorConvergence.TrackIpFrr
    leafs["prefix-list"] = monitorConvergence.PrefixList
    return leafs
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetBundleName() string { return "cisco_ios_xr" }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetYangName() string { return "monitor-convergence" }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) SetParent(parent types.Entity) { monitorConvergence.parent = parent }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetParent() types.Entity { return monitorConvergence.parent }

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation
// Control origination of a default route with
// the option of using a policy.  If no policy
// is specified the default route is
// advertised with zero cost in level 2 only.
type Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flag to indicate whether default origination is controlled using a policy.
    // The type is bool.
    UsePolicy interface{}

    // Policy name. The type is string with length: 1..64.
    PolicyName interface{}

    // Flag to indicate that the default prefix should be originated as an
    // external route. The type is interface{}.
    External interface{}
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetFilter() yfilter.YFilter { return defaultInformation.YFilter }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) SetFilter(yf yfilter.YFilter) { defaultInformation.YFilter = yf }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetGoName(yname string) string {
    if yname == "use-policy" { return "UsePolicy" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "external" { return "External" }
    return ""
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetSegmentPath() string {
    return "default-information"
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["use-policy"] = defaultInformation.UsePolicy
    leafs["policy-name"] = defaultInformation.PolicyName
    leafs["external"] = defaultInformation.External
    return leafs
}

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetBundleName() string { return "cisco_ios_xr" }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetYangName() string { return "default-information" }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) SetParent(parent types.Entity) { defaultInformation.parent = parent }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetParent() types.Entity { return defaultInformation.parent }

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances
// Per-route administrative
// distanceconfiguration
type Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Administrative distance configuration. The supplied distance is applied to
    // all routes discovered from the specified source, or only those that match
    // the supplied prefix list if this is specified. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance.
    AdminDistance []Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetFilter() yfilter.YFilter { return adminDistances.YFilter }

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) SetFilter(yf yfilter.YFilter) { adminDistances.YFilter = yf }

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetGoName(yname string) string {
    if yname == "admin-distance" { return "AdminDistance" }
    return ""
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetSegmentPath() string {
    return "admin-distances"
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "admin-distance" {
        for _, c := range adminDistances.AdminDistance {
            if adminDistances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance{}
        adminDistances.AdminDistance = append(adminDistances.AdminDistance, child)
        return &adminDistances.AdminDistance[len(adminDistances.AdminDistance)-1]
    }
    return nil
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range adminDistances.AdminDistance {
        children[adminDistances.AdminDistance[i].GetSegmentPath()] = &adminDistances.AdminDistance[i]
    }
    return children
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetBundleName() string { return "cisco_ios_xr" }

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetYangName() string { return "admin-distances" }

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) SetParent(parent types.Entity) { adminDistances.parent = parent }

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetParent() types.Entity { return adminDistances.parent }

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance
// Administrative distance configuration. The
// supplied distance is applied to all routes
// discovered from the specified source, or
// only those that match the supplied prefix
// list if this is specified
type Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP route source prefix. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    AddressPrefix interface{}

    // Administrative distance. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    Distance interface{}

    // List of prefixes to which this distance applies. The type is string with
    // length: 1..32.
    PrefixList interface{}
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetFilter() yfilter.YFilter { return adminDistance.YFilter }

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) SetFilter(yf yfilter.YFilter) { adminDistance.YFilter = yf }

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetGoName(yname string) string {
    if yname == "address-prefix" { return "AddressPrefix" }
    if yname == "distance" { return "Distance" }
    if yname == "prefix-list" { return "PrefixList" }
    return ""
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetSegmentPath() string {
    return "admin-distance" + "[address-prefix='" + fmt.Sprintf("%v", adminDistance.AddressPrefix) + "']"
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-prefix"] = adminDistance.AddressPrefix
    leafs["distance"] = adminDistance.Distance
    leafs["prefix-list"] = adminDistance.PrefixList
    return leafs
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetBundleName() string { return "cisco_ios_xr" }

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetYangName() string { return "admin-distance" }

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) SetParent(parent types.Entity) { adminDistance.parent = parent }

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetParent() types.Entity { return adminDistance.parent }

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetParentYangName() string { return "admin-distances" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Ispf
// ISPF configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Ispf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISPF state (enable/disable).
    States Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States
}

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetFilter() yfilter.YFilter { return ispf.YFilter }

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) SetFilter(yf yfilter.YFilter) { ispf.YFilter = yf }

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetGoName(yname string) string {
    if yname == "states" { return "States" }
    return ""
}

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetSegmentPath() string {
    return "ispf"
}

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "states" {
        return &ispf.States
    }
    return nil
}

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["states"] = &ispf.States
    return children
}

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetBundleName() string { return "cisco_ios_xr" }

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetYangName() string { return "ispf" }

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) SetParent(parent types.Entity) { ispf.parent = parent }

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetParent() types.Entity { return ispf.parent }

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States
// ISPF state (enable/disable)
type Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/disable ISPF. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State.
    State []Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State
}

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetFilter() yfilter.YFilter { return states.YFilter }

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) SetFilter(yf yfilter.YFilter) { states.YFilter = yf }

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    return ""
}

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetSegmentPath() string {
    return "states"
}

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "state" {
        for _, c := range states.State {
            if states.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State{}
        states.State = append(states.State, child)
        return &states.State[len(states.State)-1]
    }
    return nil
}

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range states.State {
        children[states.State[i].GetSegmentPath()] = &states.State[i]
    }
    return children
}

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetBundleName() string { return "cisco_ios_xr" }

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetYangName() string { return "states" }

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) SetParent(parent types.Entity) { states.parent = parent }

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetParent() types.Entity { return states.parent }

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetParentYangName() string { return "ispf" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State
// Enable/disable ISPF
type Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // State. The type is IsisispfState. This attribute is mandatory.
    State interface{}
}

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetFilter() yfilter.YFilter { return state.YFilter }

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) SetFilter(yf yfilter.YFilter) { state.YFilter = yf }

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "state" { return "State" }
    return ""
}

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetSegmentPath() string {
    return "state" + "[level='" + fmt.Sprintf("%v", state.Level) + "']"
}

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = state.Level
    leafs["state"] = state.State
    return leafs
}

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetBundleName() string { return "cisco_ios_xr" }

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetYangName() string { return "state" }

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) SetParent(parent types.Entity) { state.parent = parent }

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetParent() types.Entity { return state.parent }

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetParentYangName() string { return "states" }

// Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal
// MPLS LDP configuration. MPLS LDP
// configuration will only be applied for the
// IPv4-unicast address-family.
type Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If TRUE, LDP will be enabled onall IS-IS interfaces enabled for this
    // address-family. The type is bool.
    AutoConfig interface{}
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetFilter() yfilter.YFilter { return mplsLdpGlobal.YFilter }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) SetFilter(yf yfilter.YFilter) { mplsLdpGlobal.YFilter = yf }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetGoName(yname string) string {
    if yname == "auto-config" { return "AutoConfig" }
    return ""
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetSegmentPath() string {
    return "mpls-ldp-global"
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auto-config"] = mplsLdpGlobal.AutoConfig
    return leafs
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetBundleName() string { return "cisco_ios_xr" }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetYangName() string { return "mpls-ldp-global" }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) SetParent(parent types.Entity) { mplsLdpGlobal.parent = parent }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetParent() types.Entity { return mplsLdpGlobal.parent }

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Mpls
// MPLS configuration. MPLS configuration will
// only be applied for the IPv4-unicast
// address-family.
type Isis_Instances_Instance_Afs_Af_TopologyName_Mpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Install TE and non-TE nexthops in the RIB. The type is interface{}.
    IgpIntact interface{}

    // Install non-TE nexthops in the RIB for use by multicast. The type is
    // interface{}.
    MulticastIntact interface{}

    // Traffic Engineering stable IP address for system.
    RouterId Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId

    // Enable MPLS for an IS-IS at the given levels.
    Level Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level
}

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetFilter() yfilter.YFilter { return mpls.YFilter }

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) SetFilter(yf yfilter.YFilter) { mpls.YFilter = yf }

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetGoName(yname string) string {
    if yname == "igp-intact" { return "IgpIntact" }
    if yname == "multicast-intact" { return "MulticastIntact" }
    if yname == "router-id" { return "RouterId" }
    if yname == "level" { return "Level" }
    return ""
}

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetSegmentPath() string {
    return "mpls"
}

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router-id" {
        return &mpls.RouterId
    }
    if childYangName == "level" {
        return &mpls.Level
    }
    return nil
}

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["router-id"] = &mpls.RouterId
    children["level"] = &mpls.Level
    return children
}

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-intact"] = mpls.IgpIntact
    leafs["multicast-intact"] = mpls.MulticastIntact
    return leafs
}

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetBundleName() string { return "cisco_ios_xr" }

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetYangName() string { return "mpls" }

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) SetParent(parent types.Entity) { mpls.parent = parent }

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetParent() types.Entity { return mpls.parent }

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId
// Traffic Engineering stable IP address for
// system
type Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address to be used as a router ID. Precisely one of Address and
    // Interface must be specified. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Interface with designated stable IP address to be used as a router ID. This
    // must be a Loopback interface. Precisely one of Address and Interface must
    // be specified. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetFilter() yfilter.YFilter { return routerId.YFilter }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) SetFilter(yf yfilter.YFilter) { routerId.YFilter = yf }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetSegmentPath() string {
    return "router-id"
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = routerId.Address
    leafs["interface-name"] = routerId.InterfaceName
    return leafs
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetBundleName() string { return "cisco_ios_xr" }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetYangName() string { return "router-id" }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) SetParent(parent types.Entity) { routerId.parent = parent }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetParent() types.Entity { return routerId.parent }

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetParentYangName() string { return "mpls" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level
// Enable MPLS for an IS-IS at the given
// levels
type Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Level 1 enabled. The type is bool.
    Level1 interface{}

    // Level 2 enabled. The type is bool.
    Level2 interface{}
}

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetFilter() yfilter.YFilter { return level.YFilter }

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) SetFilter(yf yfilter.YFilter) { level.YFilter = yf }

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetGoName(yname string) string {
    if yname == "level1" { return "Level1" }
    if yname == "level2" { return "Level2" }
    return ""
}

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetSegmentPath() string {
    return "level"
}

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level1"] = level.Level1
    leafs["level2"] = level.Level2
    return leafs
}

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetBundleName() string { return "cisco_ios_xr" }

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetYangName() string { return "level" }

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) SetParent(parent types.Entity) { level.parent = parent }

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetParent() types.Entity { return level.parent }

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetParentYangName() string { return "mpls" }

// Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids
// Manual Adjacecy SID configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Assign adjancency SID to an interface. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid.
    ManualAdjSid []Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetFilter() yfilter.YFilter { return manualAdjSids.YFilter }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) SetFilter(yf yfilter.YFilter) { manualAdjSids.YFilter = yf }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetGoName(yname string) string {
    if yname == "manual-adj-sid" { return "ManualAdjSid" }
    return ""
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetSegmentPath() string {
    return "manual-adj-sids"
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "manual-adj-sid" {
        for _, c := range manualAdjSids.ManualAdjSid {
            if manualAdjSids.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid{}
        manualAdjSids.ManualAdjSid = append(manualAdjSids.ManualAdjSid, child)
        return &manualAdjSids.ManualAdjSid[len(manualAdjSids.ManualAdjSid)-1]
    }
    return nil
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range manualAdjSids.ManualAdjSid {
        children[manualAdjSids.ManualAdjSid[i].GetSegmentPath()] = &manualAdjSids.ManualAdjSid[i]
    }
    return children
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetBundleName() string { return "cisco_ios_xr" }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetYangName() string { return "manual-adj-sids" }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) SetParent(parent types.Entity) { manualAdjSids.parent = parent }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetParent() types.Entity { return manualAdjSids.parent }

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid
// Assign adjancency SID to an interface
type Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Sid type aboslute or index. The type is Isissid1.
    SidType interface{}

    // This attribute is a key. Sid value. The type is interface{} with range:
    // 0..1048575.
    Sid interface{}

    // Enable/Disable SID protection. The type is IsissidProtected. This attribute
    // is mandatory.
    Protected interface{}
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetFilter() yfilter.YFilter { return manualAdjSid.YFilter }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) SetFilter(yf yfilter.YFilter) { manualAdjSid.YFilter = yf }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "sid-type" { return "SidType" }
    if yname == "sid" { return "Sid" }
    if yname == "protected" { return "Protected" }
    return ""
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetSegmentPath() string {
    return "manual-adj-sid" + "[level='" + fmt.Sprintf("%v", manualAdjSid.Level) + "']" + "[sid-type='" + fmt.Sprintf("%v", manualAdjSid.SidType) + "']" + "[sid='" + fmt.Sprintf("%v", manualAdjSid.Sid) + "']"
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = manualAdjSid.Level
    leafs["sid-type"] = manualAdjSid.SidType
    leafs["sid"] = manualAdjSid.Sid
    leafs["protected"] = manualAdjSid.Protected
    return leafs
}

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetBundleName() string { return "cisco_ios_xr" }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetYangName() string { return "manual-adj-sid" }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) SetParent(parent types.Entity) { manualAdjSid.parent = parent }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetParent() types.Entity { return manualAdjSid.parent }

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetParentYangName() string { return "manual-adj-sids" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Metrics
// Metric configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Metrics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric configuration. Legal value depends on the metric-style specified for
    // the topology. If the metric-style defined is narrow, then only a value
    // between <1-63> is allowed and if the metric-style is defined as wide, then
    // a value between <1-16777215> is allowed as the metric value.  All routers
    // exclude links with the maximum wide metric (16777215) from their SPF. The
    // type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric.
    Metric []Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric
}

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetFilter() yfilter.YFilter { return metrics.YFilter }

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) SetFilter(yf yfilter.YFilter) { metrics.YFilter = yf }

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    return ""
}

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetSegmentPath() string {
    return "metrics"
}

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metric" {
        for _, c := range metrics.Metric {
            if metrics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric{}
        metrics.Metric = append(metrics.Metric, child)
        return &metrics.Metric[len(metrics.Metric)-1]
    }
    return nil
}

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range metrics.Metric {
        children[metrics.Metric[i].GetSegmentPath()] = &metrics.Metric[i]
    }
    return children
}

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetBundleName() string { return "cisco_ios_xr" }

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetYangName() string { return "metrics" }

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) SetParent(parent types.Entity) { metrics.parent = parent }

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetParent() types.Entity { return metrics.parent }

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric
// Metric configuration. Legal value depends on
// the metric-style specified for the topology. If
// the metric-style defined is narrow, then only a
// value between <1-63> is allowed and if the
// metric-style is defined as wide, then a value
// between <1-16777215> is allowed as the metric
// value.  All routers exclude links with the
// maximum wide metric (16777215) from their SPF
type Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed metric: <1-63> for narrow, <1-16777215> for wide. The type is one
    // of the following types: enumeration
    // Isis.Instances.Instance.Interfaces.Interface.InterfaceAfs.InterfaceAf.TopologyName.Metrics.Metric.Metric
    // This attribute is mandatory., or int with range: 1..16777215 This attribute
    // is mandatory..
    Metric interface{}
}

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetFilter() yfilter.YFilter { return metric.YFilter }

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) SetFilter(yf yfilter.YFilter) { metric.YFilter = yf }

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetSegmentPath() string {
    return "metric" + "[level='" + fmt.Sprintf("%v", metric.Level) + "']"
}

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = metric.Level
    leafs["metric"] = metric.Metric
    return leafs
}

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetBundleName() string { return "cisco_ios_xr" }

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetYangName() string { return "metric" }

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) SetParent(parent types.Entity) { metric.parent = parent }

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetParent() types.Entity { return metric.parent }

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetParentYangName() string { return "metrics" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric_Metric represents <1-16777215> for wide
type Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric_Metric string

const (
    // Maximum wide metric.  All routers will
    // exclude this link from their SPF
    Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric_Metric_maximum Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric_Metric = "maximum"
)

// Isis_Instances_Instance_Afs_Af_TopologyName_Weights
// Weight configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Weights struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Weight configuration under interface for load balancing. The type is slice
    // of Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight.
    Weight []Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight
}

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetFilter() yfilter.YFilter { return weights.YFilter }

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) SetFilter(yf yfilter.YFilter) { weights.YFilter = yf }

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetGoName(yname string) string {
    if yname == "weight" { return "Weight" }
    return ""
}

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetSegmentPath() string {
    return "weights"
}

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "weight" {
        for _, c := range weights.Weight {
            if weights.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight{}
        weights.Weight = append(weights.Weight, child)
        return &weights.Weight[len(weights.Weight)-1]
    }
    return nil
}

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range weights.Weight {
        children[weights.Weight[i].GetSegmentPath()] = &weights.Weight[i]
    }
    return children
}

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetBundleName() string { return "cisco_ios_xr" }

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetYangName() string { return "weights" }

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) SetParent(parent types.Entity) { weights.parent = parent }

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetParent() types.Entity { return weights.parent }

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight
// Weight configuration under interface for load
// balancing
type Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Weight to be configured under interface for Load Balancing. Allowed weight:
    // <1-16777215>. The type is interface{} with range: 1..16777214. This
    // attribute is mandatory.
    Weight interface{}
}

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetFilter() yfilter.YFilter { return weight.YFilter }

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) SetFilter(yf yfilter.YFilter) { weight.YFilter = yf }

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "weight" { return "Weight" }
    return ""
}

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetSegmentPath() string {
    return "weight" + "[level='" + fmt.Sprintf("%v", weight.Level) + "']"
}

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = weight.Level
    leafs["weight"] = weight.Weight
    return leafs
}

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetBundleName() string { return "cisco_ios_xr" }

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetYangName() string { return "weight" }

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) SetParent(parent types.Entity) { weight.parent = parent }

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetParent() types.Entity { return weight.parent }

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetParentYangName() string { return "weights" }

// Isis_Instances_Instance_LspRefreshIntervals
// LSP refresh-interval configuration
type Isis_Instances_Instance_LspRefreshIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval between re-flooding of unchanged LSPs. The type is slice of
    // Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval.
    LspRefreshInterval []Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval
}

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetFilter() yfilter.YFilter { return lspRefreshIntervals.YFilter }

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) SetFilter(yf yfilter.YFilter) { lspRefreshIntervals.YFilter = yf }

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetGoName(yname string) string {
    if yname == "lsp-refresh-interval" { return "LspRefreshInterval" }
    return ""
}

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetSegmentPath() string {
    return "lsp-refresh-intervals"
}

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-refresh-interval" {
        for _, c := range lspRefreshIntervals.LspRefreshInterval {
            if lspRefreshIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval{}
        lspRefreshIntervals.LspRefreshInterval = append(lspRefreshIntervals.LspRefreshInterval, child)
        return &lspRefreshIntervals.LspRefreshInterval[len(lspRefreshIntervals.LspRefreshInterval)-1]
    }
    return nil
}

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspRefreshIntervals.LspRefreshInterval {
        children[lspRefreshIntervals.LspRefreshInterval[i].GetSegmentPath()] = &lspRefreshIntervals.LspRefreshInterval[i]
    }
    return children
}

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetYangName() string { return "lsp-refresh-intervals" }

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) SetParent(parent types.Entity) { lspRefreshIntervals.parent = parent }

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetParent() types.Entity { return lspRefreshIntervals.parent }

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval
// Interval between re-flooding of unchanged
// LSPs
type Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 1..65535. This attribute is
    // mandatory. Units are second.
    Interval interface{}
}

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetFilter() yfilter.YFilter { return lspRefreshInterval.YFilter }

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) SetFilter(yf yfilter.YFilter) { lspRefreshInterval.YFilter = yf }

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetSegmentPath() string {
    return "lsp-refresh-interval" + "[level='" + fmt.Sprintf("%v", lspRefreshInterval.Level) + "']"
}

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspRefreshInterval.Level
    leafs["interval"] = lspRefreshInterval.Interval
    return leafs
}

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetBundleName() string { return "cisco_ios_xr" }

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetYangName() string { return "lsp-refresh-interval" }

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) SetParent(parent types.Entity) { lspRefreshInterval.parent = parent }

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetParent() types.Entity { return lspRefreshInterval.parent }

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetParentYangName() string { return "lsp-refresh-intervals" }

// Isis_Instances_Instance_Distribute
// IS-IS Distribute BGP-LS configuration
// This type is a presence type.
type Isis_Instances_Instance_Distribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance ID. The type is interface{} with range: 1..65535.
    DistInstId interface{}

    // Level. The type is IsisConfigurableLevels.
    Level interface{}

    // Seconds. The type is interface{} with range: 1..20. Units are second.
    DistThrottle interface{}
}

func (distribute *Isis_Instances_Instance_Distribute) GetFilter() yfilter.YFilter { return distribute.YFilter }

func (distribute *Isis_Instances_Instance_Distribute) SetFilter(yf yfilter.YFilter) { distribute.YFilter = yf }

func (distribute *Isis_Instances_Instance_Distribute) GetGoName(yname string) string {
    if yname == "dist-inst-id" { return "DistInstId" }
    if yname == "level" { return "Level" }
    if yname == "dist-throttle" { return "DistThrottle" }
    return ""
}

func (distribute *Isis_Instances_Instance_Distribute) GetSegmentPath() string {
    return "distribute"
}

func (distribute *Isis_Instances_Instance_Distribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (distribute *Isis_Instances_Instance_Distribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (distribute *Isis_Instances_Instance_Distribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dist-inst-id"] = distribute.DistInstId
    leafs["level"] = distribute.Level
    leafs["dist-throttle"] = distribute.DistThrottle
    return leafs
}

func (distribute *Isis_Instances_Instance_Distribute) GetBundleName() string { return "cisco_ios_xr" }

func (distribute *Isis_Instances_Instance_Distribute) GetYangName() string { return "distribute" }

func (distribute *Isis_Instances_Instance_Distribute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (distribute *Isis_Instances_Instance_Distribute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (distribute *Isis_Instances_Instance_Distribute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (distribute *Isis_Instances_Instance_Distribute) SetParent(parent types.Entity) { distribute.parent = parent }

func (distribute *Isis_Instances_Instance_Distribute) GetParent() types.Entity { return distribute.parent }

func (distribute *Isis_Instances_Instance_Distribute) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspAcceptPasswords
// LSP/SNP accept password configuration
type Isis_Instances_Instance_LspAcceptPasswords struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSP/SNP accept passwords. This requires the existence of an LSPPassword of
    // the same level . The type is slice of
    // Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword.
    LspAcceptPassword []Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword
}

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetFilter() yfilter.YFilter { return lspAcceptPasswords.YFilter }

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) SetFilter(yf yfilter.YFilter) { lspAcceptPasswords.YFilter = yf }

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetGoName(yname string) string {
    if yname == "lsp-accept-password" { return "LspAcceptPassword" }
    return ""
}

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetSegmentPath() string {
    return "lsp-accept-passwords"
}

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-accept-password" {
        for _, c := range lspAcceptPasswords.LspAcceptPassword {
            if lspAcceptPasswords.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword{}
        lspAcceptPasswords.LspAcceptPassword = append(lspAcceptPasswords.LspAcceptPassword, child)
        return &lspAcceptPasswords.LspAcceptPassword[len(lspAcceptPasswords.LspAcceptPassword)-1]
    }
    return nil
}

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspAcceptPasswords.LspAcceptPassword {
        children[lspAcceptPasswords.LspAcceptPassword[i].GetSegmentPath()] = &lspAcceptPasswords.LspAcceptPassword[i]
    }
    return children
}

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetBundleName() string { return "cisco_ios_xr" }

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetYangName() string { return "lsp-accept-passwords" }

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) SetParent(parent types.Entity) { lspAcceptPasswords.parent = parent }

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetParent() types.Entity { return lspAcceptPasswords.parent }

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword
// LSP/SNP accept passwords. This requires the
// existence of an LSPPassword of the same level
// .
type Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Password. The type is string with pattern: (!.+)|([^!].+). This attribute
    // is mandatory.
    Password interface{}
}

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetFilter() yfilter.YFilter { return lspAcceptPassword.YFilter }

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) SetFilter(yf yfilter.YFilter) { lspAcceptPassword.YFilter = yf }

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "password" { return "Password" }
    return ""
}

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetSegmentPath() string {
    return "lsp-accept-password" + "[level='" + fmt.Sprintf("%v", lspAcceptPassword.Level) + "']"
}

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspAcceptPassword.Level
    leafs["password"] = lspAcceptPassword.Password
    return leafs
}

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetBundleName() string { return "cisco_ios_xr" }

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetYangName() string { return "lsp-accept-password" }

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) SetParent(parent types.Entity) { lspAcceptPassword.parent = parent }

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetParent() types.Entity { return lspAcceptPassword.parent }

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetParentYangName() string { return "lsp-accept-passwords" }

// Isis_Instances_Instance_LspMtus
// LSP MTU configuration
type Isis_Instances_Instance_LspMtus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSP MTU. The type is slice of Isis_Instances_Instance_LspMtus_LspMtu.
    LspMtu []Isis_Instances_Instance_LspMtus_LspMtu
}

func (lspMtus *Isis_Instances_Instance_LspMtus) GetFilter() yfilter.YFilter { return lspMtus.YFilter }

func (lspMtus *Isis_Instances_Instance_LspMtus) SetFilter(yf yfilter.YFilter) { lspMtus.YFilter = yf }

func (lspMtus *Isis_Instances_Instance_LspMtus) GetGoName(yname string) string {
    if yname == "lsp-mtu" { return "LspMtu" }
    return ""
}

func (lspMtus *Isis_Instances_Instance_LspMtus) GetSegmentPath() string {
    return "lsp-mtus"
}

func (lspMtus *Isis_Instances_Instance_LspMtus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-mtu" {
        for _, c := range lspMtus.LspMtu {
            if lspMtus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_LspMtus_LspMtu{}
        lspMtus.LspMtu = append(lspMtus.LspMtu, child)
        return &lspMtus.LspMtu[len(lspMtus.LspMtu)-1]
    }
    return nil
}

func (lspMtus *Isis_Instances_Instance_LspMtus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspMtus.LspMtu {
        children[lspMtus.LspMtu[i].GetSegmentPath()] = &lspMtus.LspMtu[i]
    }
    return children
}

func (lspMtus *Isis_Instances_Instance_LspMtus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspMtus *Isis_Instances_Instance_LspMtus) GetBundleName() string { return "cisco_ios_xr" }

func (lspMtus *Isis_Instances_Instance_LspMtus) GetYangName() string { return "lsp-mtus" }

func (lspMtus *Isis_Instances_Instance_LspMtus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspMtus *Isis_Instances_Instance_LspMtus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspMtus *Isis_Instances_Instance_LspMtus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspMtus *Isis_Instances_Instance_LspMtus) SetParent(parent types.Entity) { lspMtus.parent = parent }

func (lspMtus *Isis_Instances_Instance_LspMtus) GetParent() types.Entity { return lspMtus.parent }

func (lspMtus *Isis_Instances_Instance_LspMtus) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspMtus_LspMtu
// LSP MTU
type Isis_Instances_Instance_LspMtus_LspMtu struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Bytes. The type is interface{} with range: 128..4352. This attribute is
    // mandatory. Units are byte.
    Mtu interface{}
}

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetFilter() yfilter.YFilter { return lspMtu.YFilter }

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) SetFilter(yf yfilter.YFilter) { lspMtu.YFilter = yf }

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "mtu" { return "Mtu" }
    return ""
}

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetSegmentPath() string {
    return "lsp-mtu" + "[level='" + fmt.Sprintf("%v", lspMtu.Level) + "']"
}

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspMtu.Level
    leafs["mtu"] = lspMtu.Mtu
    return leafs
}

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetBundleName() string { return "cisco_ios_xr" }

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetYangName() string { return "lsp-mtu" }

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) SetParent(parent types.Entity) { lspMtu.parent = parent }

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetParent() types.Entity { return lspMtu.parent }

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetParentYangName() string { return "lsp-mtus" }

// Isis_Instances_Instance_Nsf
// IS-IS NSF configuration
type Isis_Instances_Instance_Nsf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NSF not configured if item is deleted. The type is IsisNsfFlavor.
    Flavor interface{}

    // Per-interface time period to wait for a restart ACK during an IETF-NSF
    // restart. This configuration has no effect if IETF-NSF is not configured.
    // The type is interface{} with range: 1..20. Units are second. The default
    // value is 1.
    InterfaceTimer interface{}

    // Maximum number of times an interface timer may expire during an IETF-NSF
    // restart before the NSF restart is aborted. This configuration has no effect
    // if IETF NSF is not configured. The type is interface{} with range: 1..10.
    // The default value is 10.
    MaxInterfaceTimerExpiry interface{}

    // Maximum route lifetime following restart. When this lifetime expires, old
    // routes will be purged from the RIB. The type is interface{} with range:
    // 5..300. Units are second. The default value is 90.
    Lifetime interface{}
}

func (nsf *Isis_Instances_Instance_Nsf) GetFilter() yfilter.YFilter { return nsf.YFilter }

func (nsf *Isis_Instances_Instance_Nsf) SetFilter(yf yfilter.YFilter) { nsf.YFilter = yf }

func (nsf *Isis_Instances_Instance_Nsf) GetGoName(yname string) string {
    if yname == "flavor" { return "Flavor" }
    if yname == "interface-timer" { return "InterfaceTimer" }
    if yname == "max-interface-timer-expiry" { return "MaxInterfaceTimerExpiry" }
    if yname == "lifetime" { return "Lifetime" }
    return ""
}

func (nsf *Isis_Instances_Instance_Nsf) GetSegmentPath() string {
    return "nsf"
}

func (nsf *Isis_Instances_Instance_Nsf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nsf *Isis_Instances_Instance_Nsf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nsf *Isis_Instances_Instance_Nsf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flavor"] = nsf.Flavor
    leafs["interface-timer"] = nsf.InterfaceTimer
    leafs["max-interface-timer-expiry"] = nsf.MaxInterfaceTimerExpiry
    leafs["lifetime"] = nsf.Lifetime
    return leafs
}

func (nsf *Isis_Instances_Instance_Nsf) GetBundleName() string { return "cisco_ios_xr" }

func (nsf *Isis_Instances_Instance_Nsf) GetYangName() string { return "nsf" }

func (nsf *Isis_Instances_Instance_Nsf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nsf *Isis_Instances_Instance_Nsf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nsf *Isis_Instances_Instance_Nsf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nsf *Isis_Instances_Instance_Nsf) SetParent(parent types.Entity) { nsf.parent = parent }

func (nsf *Isis_Instances_Instance_Nsf) GetParent() types.Entity { return nsf.parent }

func (nsf *Isis_Instances_Instance_Nsf) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LinkGroups
// Link Group
type Isis_Instances_Instance_LinkGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for link group name. The type is slice of
    // Isis_Instances_Instance_LinkGroups_LinkGroup.
    LinkGroup []Isis_Instances_Instance_LinkGroups_LinkGroup
}

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetFilter() yfilter.YFilter { return linkGroups.YFilter }

func (linkGroups *Isis_Instances_Instance_LinkGroups) SetFilter(yf yfilter.YFilter) { linkGroups.YFilter = yf }

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetGoName(yname string) string {
    if yname == "link-group" { return "LinkGroup" }
    return ""
}

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetSegmentPath() string {
    return "link-groups"
}

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-group" {
        for _, c := range linkGroups.LinkGroup {
            if linkGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_LinkGroups_LinkGroup{}
        linkGroups.LinkGroup = append(linkGroups.LinkGroup, child)
        return &linkGroups.LinkGroup[len(linkGroups.LinkGroup)-1]
    }
    return nil
}

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range linkGroups.LinkGroup {
        children[linkGroups.LinkGroup[i].GetSegmentPath()] = &linkGroups.LinkGroup[i]
    }
    return children
}

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetBundleName() string { return "cisco_ios_xr" }

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetYangName() string { return "link-groups" }

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkGroups *Isis_Instances_Instance_LinkGroups) SetParent(parent types.Entity) { linkGroups.parent = parent }

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetParent() types.Entity { return linkGroups.parent }

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LinkGroups_LinkGroup
// Configuration for link group name
type Isis_Instances_Instance_LinkGroups_LinkGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Link Group Name. The type is string with length:
    // 1..40.
    LinkGroupName interface{}

    // Flag to indicate that linkgroup should be running.  This must be the first
    // object created when a linkgroup is configured, and the last object deleted
    // when it is deconfigured.  When this object is deleted, the IS-IS linkgroup
    // will be removed. The type is interface{}.
    Enable interface{}

    // Metric for redistributed routes: <0-63> for narrow, <0-16777215> for wide.
    // The type is interface{} with range: 0..16777215.
    MetricOffset interface{}

    // Revert Members. The type is interface{} with range: 2..64. The default
    // value is 2.
    RevertMembers interface{}

    // Minimum Members. The type is interface{} with range: 2..64. The default
    // value is 2.
    MinimumMembers interface{}
}

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetFilter() yfilter.YFilter { return linkGroup.YFilter }

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) SetFilter(yf yfilter.YFilter) { linkGroup.YFilter = yf }

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetGoName(yname string) string {
    if yname == "link-group-name" { return "LinkGroupName" }
    if yname == "enable" { return "Enable" }
    if yname == "metric-offset" { return "MetricOffset" }
    if yname == "revert-members" { return "RevertMembers" }
    if yname == "minimum-members" { return "MinimumMembers" }
    return ""
}

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetSegmentPath() string {
    return "link-group" + "[link-group-name='" + fmt.Sprintf("%v", linkGroup.LinkGroupName) + "']"
}

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-group-name"] = linkGroup.LinkGroupName
    leafs["enable"] = linkGroup.Enable
    leafs["metric-offset"] = linkGroup.MetricOffset
    leafs["revert-members"] = linkGroup.RevertMembers
    leafs["minimum-members"] = linkGroup.MinimumMembers
    return leafs
}

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetBundleName() string { return "cisco_ios_xr" }

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetYangName() string { return "link-group" }

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) SetParent(parent types.Entity) { linkGroup.parent = parent }

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetParent() types.Entity { return linkGroup.parent }

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetParentYangName() string { return "link-groups" }

// Isis_Instances_Instance_LspCheckIntervals
// LSP checksum check interval configuration
type Isis_Instances_Instance_LspCheckIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSP checksum check interval parameters. The type is slice of
    // Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval.
    LspCheckInterval []Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval
}

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetFilter() yfilter.YFilter { return lspCheckIntervals.YFilter }

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) SetFilter(yf yfilter.YFilter) { lspCheckIntervals.YFilter = yf }

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetGoName(yname string) string {
    if yname == "lsp-check-interval" { return "LspCheckInterval" }
    return ""
}

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetSegmentPath() string {
    return "lsp-check-intervals"
}

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-check-interval" {
        for _, c := range lspCheckIntervals.LspCheckInterval {
            if lspCheckIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval{}
        lspCheckIntervals.LspCheckInterval = append(lspCheckIntervals.LspCheckInterval, child)
        return &lspCheckIntervals.LspCheckInterval[len(lspCheckIntervals.LspCheckInterval)-1]
    }
    return nil
}

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspCheckIntervals.LspCheckInterval {
        children[lspCheckIntervals.LspCheckInterval[i].GetSegmentPath()] = &lspCheckIntervals.LspCheckInterval[i]
    }
    return children
}

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetYangName() string { return "lsp-check-intervals" }

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) SetParent(parent types.Entity) { lspCheckIntervals.parent = parent }

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetParent() types.Entity { return lspCheckIntervals.parent }

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval
// LSP checksum check interval parameters
type Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // LSP checksum check interval time in seconds. The type is interface{} with
    // range: 10..65535. This attribute is mandatory. Units are second.
    Interval interface{}
}

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetFilter() yfilter.YFilter { return lspCheckInterval.YFilter }

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) SetFilter(yf yfilter.YFilter) { lspCheckInterval.YFilter = yf }

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetSegmentPath() string {
    return "lsp-check-interval" + "[level='" + fmt.Sprintf("%v", lspCheckInterval.Level) + "']"
}

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspCheckInterval.Level
    leafs["interval"] = lspCheckInterval.Interval
    return leafs
}

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetBundleName() string { return "cisco_ios_xr" }

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetYangName() string { return "lsp-check-interval" }

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) SetParent(parent types.Entity) { lspCheckInterval.parent = parent }

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetParent() types.Entity { return lspCheckInterval.parent }

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetParentYangName() string { return "lsp-check-intervals" }

// Isis_Instances_Instance_LspPasswords
// LSP/SNP password configuration
type Isis_Instances_Instance_LspPasswords struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSP/SNP passwords. This must exist if an LSPAcceptPassword of the same
    // level exists. The type is slice of
    // Isis_Instances_Instance_LspPasswords_LspPassword.
    LspPassword []Isis_Instances_Instance_LspPasswords_LspPassword
}

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetFilter() yfilter.YFilter { return lspPasswords.YFilter }

func (lspPasswords *Isis_Instances_Instance_LspPasswords) SetFilter(yf yfilter.YFilter) { lspPasswords.YFilter = yf }

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetGoName(yname string) string {
    if yname == "lsp-password" { return "LspPassword" }
    return ""
}

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetSegmentPath() string {
    return "lsp-passwords"
}

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-password" {
        for _, c := range lspPasswords.LspPassword {
            if lspPasswords.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_LspPasswords_LspPassword{}
        lspPasswords.LspPassword = append(lspPasswords.LspPassword, child)
        return &lspPasswords.LspPassword[len(lspPasswords.LspPassword)-1]
    }
    return nil
}

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspPasswords.LspPassword {
        children[lspPasswords.LspPassword[i].GetSegmentPath()] = &lspPasswords.LspPassword[i]
    }
    return children
}

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetBundleName() string { return "cisco_ios_xr" }

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetYangName() string { return "lsp-passwords" }

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspPasswords *Isis_Instances_Instance_LspPasswords) SetParent(parent types.Entity) { lspPasswords.parent = parent }

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetParent() types.Entity { return lspPasswords.parent }

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspPasswords_LspPassword
// LSP/SNP passwords. This must exist if an
// LSPAcceptPassword of the same level exists.
type Isis_Instances_Instance_LspPasswords_LspPassword struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Algorithm. The type is IsisAuthenticationAlgorithm. This attribute is
    // mandatory.
    Algorithm interface{}

    // Failure Mode. The type is IsisAuthenticationFailureMode. This attribute is
    // mandatory.
    FailureMode interface{}

    // SNP packet authentication mode. The type is IsisSnpAuth. This attribute is
    // mandatory.
    AuthenticationType interface{}

    // Password or unencrypted Key Chain name. The type is string with pattern:
    // (!.+)|([^!].+). This attribute is mandatory.
    Password interface{}
}

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetFilter() yfilter.YFilter { return lspPassword.YFilter }

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) SetFilter(yf yfilter.YFilter) { lspPassword.YFilter = yf }

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "algorithm" { return "Algorithm" }
    if yname == "failure-mode" { return "FailureMode" }
    if yname == "authentication-type" { return "AuthenticationType" }
    if yname == "password" { return "Password" }
    return ""
}

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetSegmentPath() string {
    return "lsp-password" + "[level='" + fmt.Sprintf("%v", lspPassword.Level) + "']"
}

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspPassword.Level
    leafs["algorithm"] = lspPassword.Algorithm
    leafs["failure-mode"] = lspPassword.FailureMode
    leafs["authentication-type"] = lspPassword.AuthenticationType
    leafs["password"] = lspPassword.Password
    return leafs
}

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetBundleName() string { return "cisco_ios_xr" }

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetYangName() string { return "lsp-password" }

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) SetParent(parent types.Entity) { lspPassword.parent = parent }

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetParent() types.Entity { return lspPassword.parent }

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetParentYangName() string { return "lsp-passwords" }

// Isis_Instances_Instance_Nets
// NET configuration
type Isis_Instances_Instance_Nets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network Entity Title (NET). The type is slice of
    // Isis_Instances_Instance_Nets_Net.
    Net []Isis_Instances_Instance_Nets_Net
}

func (nets *Isis_Instances_Instance_Nets) GetFilter() yfilter.YFilter { return nets.YFilter }

func (nets *Isis_Instances_Instance_Nets) SetFilter(yf yfilter.YFilter) { nets.YFilter = yf }

func (nets *Isis_Instances_Instance_Nets) GetGoName(yname string) string {
    if yname == "net" { return "Net" }
    return ""
}

func (nets *Isis_Instances_Instance_Nets) GetSegmentPath() string {
    return "nets"
}

func (nets *Isis_Instances_Instance_Nets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "net" {
        for _, c := range nets.Net {
            if nets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Nets_Net{}
        nets.Net = append(nets.Net, child)
        return &nets.Net[len(nets.Net)-1]
    }
    return nil
}

func (nets *Isis_Instances_Instance_Nets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nets.Net {
        children[nets.Net[i].GetSegmentPath()] = &nets.Net[i]
    }
    return children
}

func (nets *Isis_Instances_Instance_Nets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nets *Isis_Instances_Instance_Nets) GetBundleName() string { return "cisco_ios_xr" }

func (nets *Isis_Instances_Instance_Nets) GetYangName() string { return "nets" }

func (nets *Isis_Instances_Instance_Nets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nets *Isis_Instances_Instance_Nets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nets *Isis_Instances_Instance_Nets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nets *Isis_Instances_Instance_Nets) SetParent(parent types.Entity) { nets.parent = parent }

func (nets *Isis_Instances_Instance_Nets) GetParent() types.Entity { return nets.parent }

func (nets *Isis_Instances_Instance_Nets) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_Nets_Net
// Network Entity Title (NET)
type Isis_Instances_Instance_Nets_Net struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Network Entity Title. The type is string with
    // pattern: [a-fA-F0-9]{2}(\.[a-fA-F0-9]{4}){3,9}\.[a-fA-F0-9]{2}.
    NetName interface{}
}

func (net *Isis_Instances_Instance_Nets_Net) GetFilter() yfilter.YFilter { return net.YFilter }

func (net *Isis_Instances_Instance_Nets_Net) SetFilter(yf yfilter.YFilter) { net.YFilter = yf }

func (net *Isis_Instances_Instance_Nets_Net) GetGoName(yname string) string {
    if yname == "net-name" { return "NetName" }
    return ""
}

func (net *Isis_Instances_Instance_Nets_Net) GetSegmentPath() string {
    return "net" + "[net-name='" + fmt.Sprintf("%v", net.NetName) + "']"
}

func (net *Isis_Instances_Instance_Nets_Net) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (net *Isis_Instances_Instance_Nets_Net) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (net *Isis_Instances_Instance_Nets_Net) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["net-name"] = net.NetName
    return leafs
}

func (net *Isis_Instances_Instance_Nets_Net) GetBundleName() string { return "cisco_ios_xr" }

func (net *Isis_Instances_Instance_Nets_Net) GetYangName() string { return "net" }

func (net *Isis_Instances_Instance_Nets_Net) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (net *Isis_Instances_Instance_Nets_Net) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (net *Isis_Instances_Instance_Nets_Net) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (net *Isis_Instances_Instance_Nets_Net) SetParent(parent types.Entity) { net.parent = parent }

func (net *Isis_Instances_Instance_Nets_Net) GetParent() types.Entity { return net.parent }

func (net *Isis_Instances_Instance_Nets_Net) GetParentYangName() string { return "nets" }

// Isis_Instances_Instance_LspLifetimes
// LSP lifetime configuration
type Isis_Instances_Instance_LspLifetimes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum LSP lifetime. The type is slice of
    // Isis_Instances_Instance_LspLifetimes_LspLifetime.
    LspLifetime []Isis_Instances_Instance_LspLifetimes_LspLifetime
}

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetFilter() yfilter.YFilter { return lspLifetimes.YFilter }

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) SetFilter(yf yfilter.YFilter) { lspLifetimes.YFilter = yf }

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetGoName(yname string) string {
    if yname == "lsp-lifetime" { return "LspLifetime" }
    return ""
}

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetSegmentPath() string {
    return "lsp-lifetimes"
}

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-lifetime" {
        for _, c := range lspLifetimes.LspLifetime {
            if lspLifetimes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_LspLifetimes_LspLifetime{}
        lspLifetimes.LspLifetime = append(lspLifetimes.LspLifetime, child)
        return &lspLifetimes.LspLifetime[len(lspLifetimes.LspLifetime)-1]
    }
    return nil
}

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspLifetimes.LspLifetime {
        children[lspLifetimes.LspLifetime[i].GetSegmentPath()] = &lspLifetimes.LspLifetime[i]
    }
    return children
}

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetBundleName() string { return "cisco_ios_xr" }

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetYangName() string { return "lsp-lifetimes" }

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) SetParent(parent types.Entity) { lspLifetimes.parent = parent }

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetParent() types.Entity { return lspLifetimes.parent }

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_LspLifetimes_LspLifetime
// Maximum LSP lifetime
type Isis_Instances_Instance_LspLifetimes_LspLifetime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 1..65535. This attribute is
    // mandatory. Units are second.
    Lifetime interface{}
}

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetFilter() yfilter.YFilter { return lspLifetime.YFilter }

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) SetFilter(yf yfilter.YFilter) { lspLifetime.YFilter = yf }

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "lifetime" { return "Lifetime" }
    return ""
}

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetSegmentPath() string {
    return "lsp-lifetime" + "[level='" + fmt.Sprintf("%v", lspLifetime.Level) + "']"
}

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspLifetime.Level
    leafs["lifetime"] = lspLifetime.Lifetime
    return leafs
}

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetBundleName() string { return "cisco_ios_xr" }

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetYangName() string { return "lsp-lifetime" }

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) SetParent(parent types.Entity) { lspLifetime.parent = parent }

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetParent() types.Entity { return lspLifetime.parent }

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetParentYangName() string { return "lsp-lifetimes" }

// Isis_Instances_Instance_OverloadBits
// LSP overload-bit configuration
type Isis_Instances_Instance_OverloadBits struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the overload bit in the System LSP so that other routers avoid this one
    // in SPF calculations. This may be done either unconditionally, or on startup
    // until either a set time has passed or IS-IS is informed that BGP has
    // converged. This is an Object with a union discriminated on an integer value
    // of the ISISOverloadBitModeType. The type is slice of
    // Isis_Instances_Instance_OverloadBits_OverloadBit.
    OverloadBit []Isis_Instances_Instance_OverloadBits_OverloadBit
}

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetFilter() yfilter.YFilter { return overloadBits.YFilter }

func (overloadBits *Isis_Instances_Instance_OverloadBits) SetFilter(yf yfilter.YFilter) { overloadBits.YFilter = yf }

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetGoName(yname string) string {
    if yname == "overload-bit" { return "OverloadBit" }
    return ""
}

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetSegmentPath() string {
    return "overload-bits"
}

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "overload-bit" {
        for _, c := range overloadBits.OverloadBit {
            if overloadBits.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_OverloadBits_OverloadBit{}
        overloadBits.OverloadBit = append(overloadBits.OverloadBit, child)
        return &overloadBits.OverloadBit[len(overloadBits.OverloadBit)-1]
    }
    return nil
}

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range overloadBits.OverloadBit {
        children[overloadBits.OverloadBit[i].GetSegmentPath()] = &overloadBits.OverloadBit[i]
    }
    return children
}

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetBundleName() string { return "cisco_ios_xr" }

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetYangName() string { return "overload-bits" }

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (overloadBits *Isis_Instances_Instance_OverloadBits) SetParent(parent types.Entity) { overloadBits.parent = parent }

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetParent() types.Entity { return overloadBits.parent }

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_OverloadBits_OverloadBit
// Set the overload bit in the System LSP so
// that other routers avoid this one in SPF
// calculations. This may be done either
// unconditionally, or on startup until either a
// set time has passed or IS-IS is informed that
// BGP has converged. This is an Object with a
// union discriminated on an integer value of
// the ISISOverloadBitModeType.
type Isis_Instances_Instance_OverloadBits_OverloadBit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Circumstances under which the overload bit is set in the system LSP. The
    // type is IsisOverloadBitMode.
    OverloadBitMode interface{}

    // Time in seconds to advertise ourself as overloaded after process startup.
    // The type is interface{} with range: 5..86400. Units are second.
    HippityPeriod interface{}

    // Advertise prefixes from other protocols. The type is IsisAdvTypeExternal.
    ExternalAdvType interface{}

    // Advertise prefixes across ISIS levels. The type is IsisAdvTypeInterLevel.
    InterLevelAdvType interface{}
}

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetFilter() yfilter.YFilter { return overloadBit.YFilter }

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) SetFilter(yf yfilter.YFilter) { overloadBit.YFilter = yf }

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "overload-bit-mode" { return "OverloadBitMode" }
    if yname == "hippity-period" { return "HippityPeriod" }
    if yname == "external-adv-type" { return "ExternalAdvType" }
    if yname == "inter-level-adv-type" { return "InterLevelAdvType" }
    return ""
}

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetSegmentPath() string {
    return "overload-bit" + "[level='" + fmt.Sprintf("%v", overloadBit.Level) + "']"
}

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = overloadBit.Level
    leafs["overload-bit-mode"] = overloadBit.OverloadBitMode
    leafs["hippity-period"] = overloadBit.HippityPeriod
    leafs["external-adv-type"] = overloadBit.ExternalAdvType
    leafs["inter-level-adv-type"] = overloadBit.InterLevelAdvType
    return leafs
}

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetBundleName() string { return "cisco_ios_xr" }

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetYangName() string { return "overload-bit" }

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) SetParent(parent types.Entity) { overloadBit.parent = parent }

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetParent() types.Entity { return overloadBit.parent }

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetParentYangName() string { return "overload-bits" }

// Isis_Instances_Instance_Interfaces
// Per-interface configuration
type Isis_Instances_Instance_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for an IS-IS interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface.
    Interface []Isis_Instances_Instance_Interfaces_Interface
}

func (interfaces *Isis_Instances_Instance_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Isis_Instances_Instance_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Isis_Instances_Instance_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Isis_Instances_Instance_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Isis_Instances_Instance_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Isis_Instances_Instance_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Isis_Instances_Instance_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Isis_Instances_Instance_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Isis_Instances_Instance_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Isis_Instances_Instance_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Isis_Instances_Instance_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Isis_Instances_Instance_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Isis_Instances_Instance_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Isis_Instances_Instance_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Isis_Instances_Instance_Interfaces) GetParentYangName() string { return "instance" }

// Isis_Instances_Instance_Interfaces_Interface
// Configuration for an IS-IS interface
type Isis_Instances_Instance_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This object must be set before any other configuration is supplied for an
    // interface, and must be the last per-interface configuration object to be
    // removed. The type is interface{}.
    Running interface{}

    // Configure circuit type for interface. The type is IsisConfigurableLevels.
    // The default value is level1-and2.
    CircuitType interface{}

    // IS-IS will attempt to form point-to-point over LAN adjacencies over this
    // interface. The type is interface{}.
    PointToPoint interface{}

    // Enable/Disable routing. The type is IsisInterfaceState.
    State interface{}

    // Mesh-group configuration. The type is one of the following types:
    // enumeration Isis.Instances.Instance.Interfaces.Interface.MeshGroup, or int
    // with range: 0..4294967295.
    MeshGroup interface{}

    // Configure high priority detection of interface down event. The type is
    // interface{}.
    LinkDownFastDetect interface{}

    // LSP-retransmission-throttle-interval configuration.
    LspRetransmitThrottleIntervals Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals

    // LSP-retransmission-interval configuration.
    LspRetransmitIntervals Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals

    // BFD configuration.
    Bfd Isis_Instances_Instance_Interfaces_Interface_Bfd

    // DIS-election priority configuration.
    Priorities Isis_Instances_Instance_Interfaces_Interface_Priorities

    // IIH accept password configuration.
    HelloAcceptPasswords Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords

    // IIH password configuration.
    HelloPasswords Isis_Instances_Instance_Interfaces_Interface_HelloPasswords

    // Hello-padding configuration.
    HelloPaddings Isis_Instances_Instance_Interfaces_Interface_HelloPaddings

    // Hello-multiplier configuration.
    HelloMultipliers Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers

    // LSP fast flood threshold configuration.
    LspFastFloodThresholds Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds

    // Prefix attribute N flag clear configuration.
    PrefixAttributeNFlagClears Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears

    // Hello-interval configuration.
    HelloIntervals Isis_Instances_Instance_Interfaces_Interface_HelloIntervals

    // Per-interface address-family configuration.
    InterfaceAfs Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs

    // CSNP-interval configuration.
    CsnpIntervals Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals

    // LSP-interval configuration.
    LspIntervals Isis_Instances_Instance_Interfaces_Interface_LspIntervals
}

func (self *Isis_Instances_Instance_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Isis_Instances_Instance_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Isis_Instances_Instance_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "running" { return "Running" }
    if yname == "circuit-type" { return "CircuitType" }
    if yname == "point-to-point" { return "PointToPoint" }
    if yname == "state" { return "State" }
    if yname == "mesh-group" { return "MeshGroup" }
    if yname == "link-down-fast-detect" { return "LinkDownFastDetect" }
    if yname == "lsp-retransmit-throttle-intervals" { return "LspRetransmitThrottleIntervals" }
    if yname == "lsp-retransmit-intervals" { return "LspRetransmitIntervals" }
    if yname == "bfd" { return "Bfd" }
    if yname == "priorities" { return "Priorities" }
    if yname == "hello-accept-passwords" { return "HelloAcceptPasswords" }
    if yname == "hello-passwords" { return "HelloPasswords" }
    if yname == "hello-paddings" { return "HelloPaddings" }
    if yname == "hello-multipliers" { return "HelloMultipliers" }
    if yname == "lsp-fast-flood-thresholds" { return "LspFastFloodThresholds" }
    if yname == "prefix-attribute-n-flag-clears" { return "PrefixAttributeNFlagClears" }
    if yname == "hello-intervals" { return "HelloIntervals" }
    if yname == "interface-afs" { return "InterfaceAfs" }
    if yname == "csnp-intervals" { return "CsnpIntervals" }
    if yname == "lsp-intervals" { return "LspIntervals" }
    return ""
}

func (self *Isis_Instances_Instance_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Isis_Instances_Instance_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-retransmit-throttle-intervals" {
        return &self.LspRetransmitThrottleIntervals
    }
    if childYangName == "lsp-retransmit-intervals" {
        return &self.LspRetransmitIntervals
    }
    if childYangName == "bfd" {
        return &self.Bfd
    }
    if childYangName == "priorities" {
        return &self.Priorities
    }
    if childYangName == "hello-accept-passwords" {
        return &self.HelloAcceptPasswords
    }
    if childYangName == "hello-passwords" {
        return &self.HelloPasswords
    }
    if childYangName == "hello-paddings" {
        return &self.HelloPaddings
    }
    if childYangName == "hello-multipliers" {
        return &self.HelloMultipliers
    }
    if childYangName == "lsp-fast-flood-thresholds" {
        return &self.LspFastFloodThresholds
    }
    if childYangName == "prefix-attribute-n-flag-clears" {
        return &self.PrefixAttributeNFlagClears
    }
    if childYangName == "hello-intervals" {
        return &self.HelloIntervals
    }
    if childYangName == "interface-afs" {
        return &self.InterfaceAfs
    }
    if childYangName == "csnp-intervals" {
        return &self.CsnpIntervals
    }
    if childYangName == "lsp-intervals" {
        return &self.LspIntervals
    }
    return nil
}

func (self *Isis_Instances_Instance_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lsp-retransmit-throttle-intervals"] = &self.LspRetransmitThrottleIntervals
    children["lsp-retransmit-intervals"] = &self.LspRetransmitIntervals
    children["bfd"] = &self.Bfd
    children["priorities"] = &self.Priorities
    children["hello-accept-passwords"] = &self.HelloAcceptPasswords
    children["hello-passwords"] = &self.HelloPasswords
    children["hello-paddings"] = &self.HelloPaddings
    children["hello-multipliers"] = &self.HelloMultipliers
    children["lsp-fast-flood-thresholds"] = &self.LspFastFloodThresholds
    children["prefix-attribute-n-flag-clears"] = &self.PrefixAttributeNFlagClears
    children["hello-intervals"] = &self.HelloIntervals
    children["interface-afs"] = &self.InterfaceAfs
    children["csnp-intervals"] = &self.CsnpIntervals
    children["lsp-intervals"] = &self.LspIntervals
    return children
}

func (self *Isis_Instances_Instance_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["running"] = self.Running
    leafs["circuit-type"] = self.CircuitType
    leafs["point-to-point"] = self.PointToPoint
    leafs["state"] = self.State
    leafs["mesh-group"] = self.MeshGroup
    leafs["link-down-fast-detect"] = self.LinkDownFastDetect
    return leafs
}

func (self *Isis_Instances_Instance_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Isis_Instances_Instance_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Isis_Instances_Instance_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Isis_Instances_Instance_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Isis_Instances_Instance_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Isis_Instances_Instance_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Isis_Instances_Instance_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Isis_Instances_Instance_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals
// LSP-retransmission-throttle-interval
// configuration
type Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum interval betwen retransissions of different LSPs. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval.
    LspRetransmitThrottleInterval []Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval
}

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetFilter() yfilter.YFilter { return lspRetransmitThrottleIntervals.YFilter }

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) SetFilter(yf yfilter.YFilter) { lspRetransmitThrottleIntervals.YFilter = yf }

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetGoName(yname string) string {
    if yname == "lsp-retransmit-throttle-interval" { return "LspRetransmitThrottleInterval" }
    return ""
}

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetSegmentPath() string {
    return "lsp-retransmit-throttle-intervals"
}

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-retransmit-throttle-interval" {
        for _, c := range lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval {
            if lspRetransmitThrottleIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval{}
        lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval = append(lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval, child)
        return &lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval[len(lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval)-1]
    }
    return nil
}

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval {
        children[lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval[i].GetSegmentPath()] = &lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval[i]
    }
    return children
}

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetYangName() string { return "lsp-retransmit-throttle-intervals" }

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) SetParent(parent types.Entity) { lspRetransmitThrottleIntervals.parent = parent }

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetParent() types.Entity { return lspRetransmitThrottleIntervals.parent }

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval
// Minimum interval betwen retransissions of
// different LSPs
type Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Milliseconds. The type is interface{} with range: 0..65535. This attribute
    // is mandatory. Units are millisecond.
    Interval interface{}
}

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetFilter() yfilter.YFilter { return lspRetransmitThrottleInterval.YFilter }

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) SetFilter(yf yfilter.YFilter) { lspRetransmitThrottleInterval.YFilter = yf }

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetSegmentPath() string {
    return "lsp-retransmit-throttle-interval" + "[level='" + fmt.Sprintf("%v", lspRetransmitThrottleInterval.Level) + "']"
}

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspRetransmitThrottleInterval.Level
    leafs["interval"] = lspRetransmitThrottleInterval.Interval
    return leafs
}

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetBundleName() string { return "cisco_ios_xr" }

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetYangName() string { return "lsp-retransmit-throttle-interval" }

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) SetParent(parent types.Entity) { lspRetransmitThrottleInterval.parent = parent }

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetParent() types.Entity { return lspRetransmitThrottleInterval.parent }

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetParentYangName() string { return "lsp-retransmit-throttle-intervals" }

// Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals
// LSP-retransmission-interval configuration
type Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval between retransmissions of the same LSP. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval.
    LspRetransmitInterval []Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval
}

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetFilter() yfilter.YFilter { return lspRetransmitIntervals.YFilter }

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) SetFilter(yf yfilter.YFilter) { lspRetransmitIntervals.YFilter = yf }

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetGoName(yname string) string {
    if yname == "lsp-retransmit-interval" { return "LspRetransmitInterval" }
    return ""
}

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetSegmentPath() string {
    return "lsp-retransmit-intervals"
}

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-retransmit-interval" {
        for _, c := range lspRetransmitIntervals.LspRetransmitInterval {
            if lspRetransmitIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval{}
        lspRetransmitIntervals.LspRetransmitInterval = append(lspRetransmitIntervals.LspRetransmitInterval, child)
        return &lspRetransmitIntervals.LspRetransmitInterval[len(lspRetransmitIntervals.LspRetransmitInterval)-1]
    }
    return nil
}

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspRetransmitIntervals.LspRetransmitInterval {
        children[lspRetransmitIntervals.LspRetransmitInterval[i].GetSegmentPath()] = &lspRetransmitIntervals.LspRetransmitInterval[i]
    }
    return children
}

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetYangName() string { return "lsp-retransmit-intervals" }

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) SetParent(parent types.Entity) { lspRetransmitIntervals.parent = parent }

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetParent() types.Entity { return lspRetransmitIntervals.parent }

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval
// Interval between retransmissions of the
// same LSP
type Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 0..65535. This attribute is
    // mandatory. Units are second.
    Interval interface{}
}

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetFilter() yfilter.YFilter { return lspRetransmitInterval.YFilter }

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) SetFilter(yf yfilter.YFilter) { lspRetransmitInterval.YFilter = yf }

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetSegmentPath() string {
    return "lsp-retransmit-interval" + "[level='" + fmt.Sprintf("%v", lspRetransmitInterval.Level) + "']"
}

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspRetransmitInterval.Level
    leafs["interval"] = lspRetransmitInterval.Interval
    return leafs
}

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetBundleName() string { return "cisco_ios_xr" }

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetYangName() string { return "lsp-retransmit-interval" }

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) SetParent(parent types.Entity) { lspRetransmitInterval.parent = parent }

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetParent() types.Entity { return lspRetransmitInterval.parent }

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetParentYangName() string { return "lsp-retransmit-intervals" }

// Isis_Instances_Instance_Interfaces_Interface_Bfd
// BFD configuration
type Isis_Instances_Instance_Interfaces_Interface_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    EnableIpv6 interface{}

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    EnableIpv4 interface{}

    // Hello interval for BFD sessions created by isis. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // Detection multiplier for BFD sessions created by isis. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}
}

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetGoName(yname string) string {
    if yname == "enable-ipv6" { return "EnableIpv6" }
    if yname == "enable-ipv4" { return "EnableIpv4" }
    if yname == "interval" { return "Interval" }
    if yname == "detection-multiplier" { return "DetectionMultiplier" }
    return ""
}

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable-ipv6"] = bfd.EnableIpv6
    leafs["enable-ipv4"] = bfd.EnableIpv4
    leafs["interval"] = bfd.Interval
    leafs["detection-multiplier"] = bfd.DetectionMultiplier
    return leafs
}

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetYangName() string { return "bfd" }

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_Priorities
// DIS-election priority configuration
type Isis_Instances_Instance_Interfaces_Interface_Priorities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DIS-election priority. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority.
    Priority []Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority
}

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetFilter() yfilter.YFilter { return priorities.YFilter }

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) SetFilter(yf yfilter.YFilter) { priorities.YFilter = yf }

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetGoName(yname string) string {
    if yname == "priority" { return "Priority" }
    return ""
}

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetSegmentPath() string {
    return "priorities"
}

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "priority" {
        for _, c := range priorities.Priority {
            if priorities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority{}
        priorities.Priority = append(priorities.Priority, child)
        return &priorities.Priority[len(priorities.Priority)-1]
    }
    return nil
}

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range priorities.Priority {
        children[priorities.Priority[i].GetSegmentPath()] = &priorities.Priority[i]
    }
    return children
}

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetBundleName() string { return "cisco_ios_xr" }

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetYangName() string { return "priorities" }

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) SetParent(parent types.Entity) { priorities.parent = parent }

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetParent() types.Entity { return priorities.parent }

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority
// DIS-election priority
type Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Priority. The type is interface{} with range: 0..127. This attribute is
    // mandatory.
    PriorityValue interface{}
}

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetFilter() yfilter.YFilter { return priority.YFilter }

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) SetFilter(yf yfilter.YFilter) { priority.YFilter = yf }

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "priority-value" { return "PriorityValue" }
    return ""
}

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetSegmentPath() string {
    return "priority" + "[level='" + fmt.Sprintf("%v", priority.Level) + "']"
}

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = priority.Level
    leafs["priority-value"] = priority.PriorityValue
    return leafs
}

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetBundleName() string { return "cisco_ios_xr" }

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetYangName() string { return "priority" }

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) SetParent(parent types.Entity) { priority.parent = parent }

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetParent() types.Entity { return priority.parent }

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetParentYangName() string { return "priorities" }

// Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords
// IIH accept password configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IIH accept passwords. This requires the existence of a HelloPassword of the
    // same level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword.
    HelloAcceptPassword []Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword
}

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetFilter() yfilter.YFilter { return helloAcceptPasswords.YFilter }

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) SetFilter(yf yfilter.YFilter) { helloAcceptPasswords.YFilter = yf }

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetGoName(yname string) string {
    if yname == "hello-accept-password" { return "HelloAcceptPassword" }
    return ""
}

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetSegmentPath() string {
    return "hello-accept-passwords"
}

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hello-accept-password" {
        for _, c := range helloAcceptPasswords.HelloAcceptPassword {
            if helloAcceptPasswords.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword{}
        helloAcceptPasswords.HelloAcceptPassword = append(helloAcceptPasswords.HelloAcceptPassword, child)
        return &helloAcceptPasswords.HelloAcceptPassword[len(helloAcceptPasswords.HelloAcceptPassword)-1]
    }
    return nil
}

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helloAcceptPasswords.HelloAcceptPassword {
        children[helloAcceptPasswords.HelloAcceptPassword[i].GetSegmentPath()] = &helloAcceptPasswords.HelloAcceptPassword[i]
    }
    return children
}

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetBundleName() string { return "cisco_ios_xr" }

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetYangName() string { return "hello-accept-passwords" }

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) SetParent(parent types.Entity) { helloAcceptPasswords.parent = parent }

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetParent() types.Entity { return helloAcceptPasswords.parent }

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword
// IIH accept passwords. This requires the
// existence of a HelloPassword of the same
// level.
type Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Password. The type is string with pattern: (!.+)|([^!].+). This attribute
    // is mandatory.
    Password interface{}
}

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetFilter() yfilter.YFilter { return helloAcceptPassword.YFilter }

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) SetFilter(yf yfilter.YFilter) { helloAcceptPassword.YFilter = yf }

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "password" { return "Password" }
    return ""
}

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetSegmentPath() string {
    return "hello-accept-password" + "[level='" + fmt.Sprintf("%v", helloAcceptPassword.Level) + "']"
}

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = helloAcceptPassword.Level
    leafs["password"] = helloAcceptPassword.Password
    return leafs
}

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetBundleName() string { return "cisco_ios_xr" }

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetYangName() string { return "hello-accept-password" }

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) SetParent(parent types.Entity) { helloAcceptPassword.parent = parent }

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetParent() types.Entity { return helloAcceptPassword.parent }

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetParentYangName() string { return "hello-accept-passwords" }

// Isis_Instances_Instance_Interfaces_Interface_HelloPasswords
// IIH password configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloPasswords struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IIH passwords. This must exist if a HelloAcceptPassword of the same level
    // exists. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword.
    HelloPassword []Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword
}

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetFilter() yfilter.YFilter { return helloPasswords.YFilter }

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) SetFilter(yf yfilter.YFilter) { helloPasswords.YFilter = yf }

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetGoName(yname string) string {
    if yname == "hello-password" { return "HelloPassword" }
    return ""
}

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetSegmentPath() string {
    return "hello-passwords"
}

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hello-password" {
        for _, c := range helloPasswords.HelloPassword {
            if helloPasswords.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword{}
        helloPasswords.HelloPassword = append(helloPasswords.HelloPassword, child)
        return &helloPasswords.HelloPassword[len(helloPasswords.HelloPassword)-1]
    }
    return nil
}

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helloPasswords.HelloPassword {
        children[helloPasswords.HelloPassword[i].GetSegmentPath()] = &helloPasswords.HelloPassword[i]
    }
    return children
}

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetBundleName() string { return "cisco_ios_xr" }

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetYangName() string { return "hello-passwords" }

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) SetParent(parent types.Entity) { helloPasswords.parent = parent }

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetParent() types.Entity { return helloPasswords.parent }

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword
// IIH passwords. This must exist if a
// HelloAcceptPassword of the same level
// exists.
type Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Algorithm. The type is IsisAuthenticationAlgorithm. This attribute is
    // mandatory.
    Algorithm interface{}

    // Failure Mode. The type is IsisAuthenticationFailureMode. This attribute is
    // mandatory.
    FailureMode interface{}

    // Password or unencrypted Key Chain name. The type is string with pattern:
    // (!.+)|([^!].+). This attribute is mandatory.
    Password interface{}
}

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetFilter() yfilter.YFilter { return helloPassword.YFilter }

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) SetFilter(yf yfilter.YFilter) { helloPassword.YFilter = yf }

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "algorithm" { return "Algorithm" }
    if yname == "failure-mode" { return "FailureMode" }
    if yname == "password" { return "Password" }
    return ""
}

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetSegmentPath() string {
    return "hello-password" + "[level='" + fmt.Sprintf("%v", helloPassword.Level) + "']"
}

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = helloPassword.Level
    leafs["algorithm"] = helloPassword.Algorithm
    leafs["failure-mode"] = helloPassword.FailureMode
    leafs["password"] = helloPassword.Password
    return leafs
}

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetBundleName() string { return "cisco_ios_xr" }

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetYangName() string { return "hello-password" }

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) SetParent(parent types.Entity) { helloPassword.parent = parent }

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetParent() types.Entity { return helloPassword.parent }

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetParentYangName() string { return "hello-passwords" }

// Isis_Instances_Instance_Interfaces_Interface_HelloPaddings
// Hello-padding configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloPaddings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pad IIHs to the interface MTU. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding.
    HelloPadding []Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding
}

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetFilter() yfilter.YFilter { return helloPaddings.YFilter }

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) SetFilter(yf yfilter.YFilter) { helloPaddings.YFilter = yf }

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetGoName(yname string) string {
    if yname == "hello-padding" { return "HelloPadding" }
    return ""
}

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetSegmentPath() string {
    return "hello-paddings"
}

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hello-padding" {
        for _, c := range helloPaddings.HelloPadding {
            if helloPaddings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding{}
        helloPaddings.HelloPadding = append(helloPaddings.HelloPadding, child)
        return &helloPaddings.HelloPadding[len(helloPaddings.HelloPadding)-1]
    }
    return nil
}

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helloPaddings.HelloPadding {
        children[helloPaddings.HelloPadding[i].GetSegmentPath()] = &helloPaddings.HelloPadding[i]
    }
    return children
}

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetBundleName() string { return "cisco_ios_xr" }

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetYangName() string { return "hello-paddings" }

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) SetParent(parent types.Entity) { helloPaddings.parent = parent }

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetParent() types.Entity { return helloPaddings.parent }

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding
// Pad IIHs to the interface MTU
type Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Hello padding type value. The type is IsisHelloPadding. This attribute is
    // mandatory.
    PaddingType interface{}
}

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetFilter() yfilter.YFilter { return helloPadding.YFilter }

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) SetFilter(yf yfilter.YFilter) { helloPadding.YFilter = yf }

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "padding-type" { return "PaddingType" }
    return ""
}

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetSegmentPath() string {
    return "hello-padding" + "[level='" + fmt.Sprintf("%v", helloPadding.Level) + "']"
}

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = helloPadding.Level
    leafs["padding-type"] = helloPadding.PaddingType
    return leafs
}

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetBundleName() string { return "cisco_ios_xr" }

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetYangName() string { return "hello-padding" }

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) SetParent(parent types.Entity) { helloPadding.parent = parent }

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetParent() types.Entity { return helloPadding.parent }

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetParentYangName() string { return "hello-paddings" }

// Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers
// Hello-multiplier configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hello-multiplier configuration. The number of successive IIHs that may be
    // missed on an adjacency before it is considered down. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier.
    HelloMultiplier []Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier
}

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetFilter() yfilter.YFilter { return helloMultipliers.YFilter }

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) SetFilter(yf yfilter.YFilter) { helloMultipliers.YFilter = yf }

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetGoName(yname string) string {
    if yname == "hello-multiplier" { return "HelloMultiplier" }
    return ""
}

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetSegmentPath() string {
    return "hello-multipliers"
}

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hello-multiplier" {
        for _, c := range helloMultipliers.HelloMultiplier {
            if helloMultipliers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier{}
        helloMultipliers.HelloMultiplier = append(helloMultipliers.HelloMultiplier, child)
        return &helloMultipliers.HelloMultiplier[len(helloMultipliers.HelloMultiplier)-1]
    }
    return nil
}

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helloMultipliers.HelloMultiplier {
        children[helloMultipliers.HelloMultiplier[i].GetSegmentPath()] = &helloMultipliers.HelloMultiplier[i]
    }
    return children
}

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetBundleName() string { return "cisco_ios_xr" }

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetYangName() string { return "hello-multipliers" }

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) SetParent(parent types.Entity) { helloMultipliers.parent = parent }

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetParent() types.Entity { return helloMultipliers.parent }

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier
// Hello-multiplier configuration. The number
// of successive IIHs that may be missed on an
// adjacency before it is considered down.
type Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Hello multiplier value. The type is interface{} with range: 3..1000. This
    // attribute is mandatory.
    Multiplier interface{}
}

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetFilter() yfilter.YFilter { return helloMultiplier.YFilter }

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) SetFilter(yf yfilter.YFilter) { helloMultiplier.YFilter = yf }

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "multiplier" { return "Multiplier" }
    return ""
}

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetSegmentPath() string {
    return "hello-multiplier" + "[level='" + fmt.Sprintf("%v", helloMultiplier.Level) + "']"
}

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = helloMultiplier.Level
    leafs["multiplier"] = helloMultiplier.Multiplier
    return leafs
}

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetBundleName() string { return "cisco_ios_xr" }

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetYangName() string { return "hello-multiplier" }

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) SetParent(parent types.Entity) { helloMultiplier.parent = parent }

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetParent() types.Entity { return helloMultiplier.parent }

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetParentYangName() string { return "hello-multipliers" }

// Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds
// LSP fast flood threshold configuration
type Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of LSPs to send back to back on an interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold.
    LspFastFloodThreshold []Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold
}

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetFilter() yfilter.YFilter { return lspFastFloodThresholds.YFilter }

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) SetFilter(yf yfilter.YFilter) { lspFastFloodThresholds.YFilter = yf }

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetGoName(yname string) string {
    if yname == "lsp-fast-flood-threshold" { return "LspFastFloodThreshold" }
    return ""
}

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetSegmentPath() string {
    return "lsp-fast-flood-thresholds"
}

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-fast-flood-threshold" {
        for _, c := range lspFastFloodThresholds.LspFastFloodThreshold {
            if lspFastFloodThresholds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold{}
        lspFastFloodThresholds.LspFastFloodThreshold = append(lspFastFloodThresholds.LspFastFloodThreshold, child)
        return &lspFastFloodThresholds.LspFastFloodThreshold[len(lspFastFloodThresholds.LspFastFloodThreshold)-1]
    }
    return nil
}

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspFastFloodThresholds.LspFastFloodThreshold {
        children[lspFastFloodThresholds.LspFastFloodThreshold[i].GetSegmentPath()] = &lspFastFloodThresholds.LspFastFloodThreshold[i]
    }
    return children
}

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetBundleName() string { return "cisco_ios_xr" }

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetYangName() string { return "lsp-fast-flood-thresholds" }

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) SetParent(parent types.Entity) { lspFastFloodThresholds.parent = parent }

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetParent() types.Entity { return lspFastFloodThresholds.parent }

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold
// Number of LSPs to send back to back on an
// interface.
type Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Count. The type is interface{} with range: 1..4294967295. This attribute is
    // mandatory.
    Count interface{}
}

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetFilter() yfilter.YFilter { return lspFastFloodThreshold.YFilter }

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) SetFilter(yf yfilter.YFilter) { lspFastFloodThreshold.YFilter = yf }

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "count" { return "Count" }
    return ""
}

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetSegmentPath() string {
    return "lsp-fast-flood-threshold" + "[level='" + fmt.Sprintf("%v", lspFastFloodThreshold.Level) + "']"
}

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspFastFloodThreshold.Level
    leafs["count"] = lspFastFloodThreshold.Count
    return leafs
}

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetYangName() string { return "lsp-fast-flood-threshold" }

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) SetParent(parent types.Entity) { lspFastFloodThreshold.parent = parent }

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetParent() types.Entity { return lspFastFloodThreshold.parent }

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetParentYangName() string { return "lsp-fast-flood-thresholds" }

// Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears
// Prefix attribute N flag clear configuration
type Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear the N flag in prefix attribute flags sub-TLV. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear.
    PrefixAttributeNFlagClear []Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear
}

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetFilter() yfilter.YFilter { return prefixAttributeNFlagClears.YFilter }

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) SetFilter(yf yfilter.YFilter) { prefixAttributeNFlagClears.YFilter = yf }

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetGoName(yname string) string {
    if yname == "prefix-attribute-n-flag-clear" { return "PrefixAttributeNFlagClear" }
    return ""
}

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetSegmentPath() string {
    return "prefix-attribute-n-flag-clears"
}

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-attribute-n-flag-clear" {
        for _, c := range prefixAttributeNFlagClears.PrefixAttributeNFlagClear {
            if prefixAttributeNFlagClears.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear{}
        prefixAttributeNFlagClears.PrefixAttributeNFlagClear = append(prefixAttributeNFlagClears.PrefixAttributeNFlagClear, child)
        return &prefixAttributeNFlagClears.PrefixAttributeNFlagClear[len(prefixAttributeNFlagClears.PrefixAttributeNFlagClear)-1]
    }
    return nil
}

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixAttributeNFlagClears.PrefixAttributeNFlagClear {
        children[prefixAttributeNFlagClears.PrefixAttributeNFlagClear[i].GetSegmentPath()] = &prefixAttributeNFlagClears.PrefixAttributeNFlagClear[i]
    }
    return children
}

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetBundleName() string { return "cisco_ios_xr" }

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetYangName() string { return "prefix-attribute-n-flag-clears" }

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) SetParent(parent types.Entity) { prefixAttributeNFlagClears.parent = parent }

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetParent() types.Entity { return prefixAttributeNFlagClears.parent }

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear
// Clear the N flag in prefix attribute flags
// sub-TLV
type Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetFilter() yfilter.YFilter { return prefixAttributeNFlagClear.YFilter }

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) SetFilter(yf yfilter.YFilter) { prefixAttributeNFlagClear.YFilter = yf }

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    return ""
}

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetSegmentPath() string {
    return "prefix-attribute-n-flag-clear" + "[level='" + fmt.Sprintf("%v", prefixAttributeNFlagClear.Level) + "']"
}

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = prefixAttributeNFlagClear.Level
    return leafs
}

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetBundleName() string { return "cisco_ios_xr" }

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetYangName() string { return "prefix-attribute-n-flag-clear" }

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) SetParent(parent types.Entity) { prefixAttributeNFlagClear.parent = parent }

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetParent() types.Entity { return prefixAttributeNFlagClear.parent }

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetParentYangName() string { return "prefix-attribute-n-flag-clears" }

// Isis_Instances_Instance_Interfaces_Interface_HelloIntervals
// Hello-interval configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hello-interval configuration. The interval at which IIH packets will be
    // sent. This will be three times quicker on a LAN interface which has been
    // electted DIS. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval.
    HelloInterval []Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval
}

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetFilter() yfilter.YFilter { return helloIntervals.YFilter }

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) SetFilter(yf yfilter.YFilter) { helloIntervals.YFilter = yf }

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetGoName(yname string) string {
    if yname == "hello-interval" { return "HelloInterval" }
    return ""
}

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetSegmentPath() string {
    return "hello-intervals"
}

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hello-interval" {
        for _, c := range helloIntervals.HelloInterval {
            if helloIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval{}
        helloIntervals.HelloInterval = append(helloIntervals.HelloInterval, child)
        return &helloIntervals.HelloInterval[len(helloIntervals.HelloInterval)-1]
    }
    return nil
}

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helloIntervals.HelloInterval {
        children[helloIntervals.HelloInterval[i].GetSegmentPath()] = &helloIntervals.HelloInterval[i]
    }
    return children
}

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetYangName() string { return "hello-intervals" }

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) SetParent(parent types.Entity) { helloIntervals.parent = parent }

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetParent() types.Entity { return helloIntervals.parent }

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval
// Hello-interval configuration. The interval
// at which IIH packets will be sent. This
// will be three times quicker on a LAN
// interface which has been electted DIS.
type Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 1..65535. This attribute is
    // mandatory. Units are second.
    Interval interface{}
}

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetFilter() yfilter.YFilter { return helloInterval.YFilter }

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) SetFilter(yf yfilter.YFilter) { helloInterval.YFilter = yf }

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetSegmentPath() string {
    return "hello-interval" + "[level='" + fmt.Sprintf("%v", helloInterval.Level) + "']"
}

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = helloInterval.Level
    leafs["interval"] = helloInterval.Interval
    return leafs
}

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetBundleName() string { return "cisco_ios_xr" }

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetYangName() string { return "hello-interval" }

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) SetParent(parent types.Entity) { helloInterval.parent = parent }

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetParent() types.Entity { return helloInterval.parent }

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetParentYangName() string { return "hello-intervals" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs
// Per-interface address-family configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for an IS-IS address-family on a single interface. If a named
    // (non-default) topology is being created it must be multicast. Also the
    // topology ID mustbe set first and delete last in the router configuration.
    // The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf.
    InterfaceAf []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf
}

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetFilter() yfilter.YFilter { return interfaceAfs.YFilter }

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) SetFilter(yf yfilter.YFilter) { interfaceAfs.YFilter = yf }

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetGoName(yname string) string {
    if yname == "interface-af" { return "InterfaceAf" }
    return ""
}

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetSegmentPath() string {
    return "interface-afs"
}

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-af" {
        for _, c := range interfaceAfs.InterfaceAf {
            if interfaceAfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf{}
        interfaceAfs.InterfaceAf = append(interfaceAfs.InterfaceAf, child)
        return &interfaceAfs.InterfaceAf[len(interfaceAfs.InterfaceAf)-1]
    }
    return nil
}

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceAfs.InterfaceAf {
        children[interfaceAfs.InterfaceAf[i].GetSegmentPath()] = &interfaceAfs.InterfaceAf[i]
    }
    return children
}

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetYangName() string { return "interface-afs" }

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) SetParent(parent types.Entity) { interfaceAfs.parent = parent }

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetParent() types.Entity { return interfaceAfs.parent }

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf
// Configuration for an IS-IS address-family
// on a single interface. If a named
// (non-default) topology is being created it
// must be multicast. Also the topology ID
// mustbe set first and delete last in the
// router configuration.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address family. The type is IsisAddressFamily.
    AfName interface{}

    // This attribute is a key. Sub address family. The type is
    // IsisSubAddressFamily.
    SafName interface{}

    // Data container.
    InterfaceAfData Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData

    // keys: topology-name. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName.
    TopologyName []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName
}

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetFilter() yfilter.YFilter { return interfaceAf.YFilter }

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) SetFilter(yf yfilter.YFilter) { interfaceAf.YFilter = yf }

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "saf-name" { return "SafName" }
    if yname == "interface-af-data" { return "InterfaceAfData" }
    if yname == "topology-name" { return "TopologyName" }
    return ""
}

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetSegmentPath() string {
    return "interface-af" + "[af-name='" + fmt.Sprintf("%v", interfaceAf.AfName) + "']" + "[saf-name='" + fmt.Sprintf("%v", interfaceAf.SafName) + "']"
}

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-af-data" {
        return &interfaceAf.InterfaceAfData
    }
    if childYangName == "topology-name" {
        for _, c := range interfaceAf.TopologyName {
            if interfaceAf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName{}
        interfaceAf.TopologyName = append(interfaceAf.TopologyName, child)
        return &interfaceAf.TopologyName[len(interfaceAf.TopologyName)-1]
    }
    return nil
}

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-af-data"] = &interfaceAf.InterfaceAfData
    for i := range interfaceAf.TopologyName {
        children[interfaceAf.TopologyName[i].GetSegmentPath()] = &interfaceAf.TopologyName[i]
    }
    return children
}

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = interfaceAf.AfName
    leafs["saf-name"] = interfaceAf.SafName
    return leafs
}

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetYangName() string { return "interface-af" }

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) SetParent(parent types.Entity) { interfaceAf.parent = parent }

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetParent() types.Entity { return interfaceAf.parent }

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetParentYangName() string { return "interface-afs" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData
// Data container.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface state. The type is IsisInterfaceAfState.
    InterfaceAfState interface{}

    // The presence of this object allows an address-family to be run over the
    // interface in question.This must be the first object created under the
    // InterfaceAddressFamily container, and the last one deleted. The type is
    // interface{}.
    Running interface{}

    // Assign prefix SID to an interface, ISISPHPFlag will be rejected if set to
    // disable, ISISEXPLICITNULLFlag will override the value of ISISPHPFlag.
    PrefixSid Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid

    // Fast-ReRoute configuration.
    InterfaceFrrTable Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable

    // MPLS LDP configuration.
    MplsLdp Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp

    // Assign prefix SSPF SID to an interface, ISISPHPFlag will be rejected if set
    // to disable, ISISEXPLICITNULLFlag will override the value of ISISPHPFlag.
    PrefixSspfsid Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid

    // AutoMetric configuration.
    AutoMetrics Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics

    // admin-tag configuration.
    AdminTags Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags

    // Provide link group name and level.
    InterfaceLinkGroup Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup

    // Manual Adjacecy SID configuration.
    ManualAdjSids Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids

    // Metric configuration.
    Metrics Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics

    // Weight configuration.
    Weights Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights
}

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetFilter() yfilter.YFilter { return interfaceAfData.YFilter }

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) SetFilter(yf yfilter.YFilter) { interfaceAfData.YFilter = yf }

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetGoName(yname string) string {
    if yname == "interface-af-state" { return "InterfaceAfState" }
    if yname == "running" { return "Running" }
    if yname == "prefix-sid" { return "PrefixSid" }
    if yname == "interface-frr-table" { return "InterfaceFrrTable" }
    if yname == "mpls-ldp" { return "MplsLdp" }
    if yname == "prefix-sspfsid" { return "PrefixSspfsid" }
    if yname == "auto-metrics" { return "AutoMetrics" }
    if yname == "admin-tags" { return "AdminTags" }
    if yname == "interface-link-group" { return "InterfaceLinkGroup" }
    if yname == "manual-adj-sids" { return "ManualAdjSids" }
    if yname == "metrics" { return "Metrics" }
    if yname == "weights" { return "Weights" }
    return ""
}

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetSegmentPath() string {
    return "interface-af-data"
}

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-sid" {
        return &interfaceAfData.PrefixSid
    }
    if childYangName == "interface-frr-table" {
        return &interfaceAfData.InterfaceFrrTable
    }
    if childYangName == "mpls-ldp" {
        return &interfaceAfData.MplsLdp
    }
    if childYangName == "prefix-sspfsid" {
        return &interfaceAfData.PrefixSspfsid
    }
    if childYangName == "auto-metrics" {
        return &interfaceAfData.AutoMetrics
    }
    if childYangName == "admin-tags" {
        return &interfaceAfData.AdminTags
    }
    if childYangName == "interface-link-group" {
        return &interfaceAfData.InterfaceLinkGroup
    }
    if childYangName == "manual-adj-sids" {
        return &interfaceAfData.ManualAdjSids
    }
    if childYangName == "metrics" {
        return &interfaceAfData.Metrics
    }
    if childYangName == "weights" {
        return &interfaceAfData.Weights
    }
    return nil
}

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-sid"] = &interfaceAfData.PrefixSid
    children["interface-frr-table"] = &interfaceAfData.InterfaceFrrTable
    children["mpls-ldp"] = &interfaceAfData.MplsLdp
    children["prefix-sspfsid"] = &interfaceAfData.PrefixSspfsid
    children["auto-metrics"] = &interfaceAfData.AutoMetrics
    children["admin-tags"] = &interfaceAfData.AdminTags
    children["interface-link-group"] = &interfaceAfData.InterfaceLinkGroup
    children["manual-adj-sids"] = &interfaceAfData.ManualAdjSids
    children["metrics"] = &interfaceAfData.Metrics
    children["weights"] = &interfaceAfData.Weights
    return children
}

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-af-state"] = interfaceAfData.InterfaceAfState
    leafs["running"] = interfaceAfData.Running
    return leafs
}

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetYangName() string { return "interface-af-data" }

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) SetParent(parent types.Entity) { interfaceAfData.parent = parent }

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetParent() types.Entity { return interfaceAfData.parent }

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetParentYangName() string { return "interface-af" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid
// Assign prefix SID to an interface,
// ISISPHPFlag will be rejected if set to
// disable, ISISEXPLICITNULLFlag will
// override the value of ISISPHPFlag
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type for the interface. The type is Isissid1. This attribute is
    // mandatory.
    Type interface{}

    // SID value for the interface. The type is interface{} with range:
    // 0..1048575. This attribute is mandatory.
    Value interface{}

    // Enable/Disable Penultimate Hop Popping. The type is IsisphpFlag. This
    // attribute is mandatory.
    Php interface{}

    // Enable/Disable Explicit-NULL flag. The type is IsisexplicitNullFlag. This
    // attribute is mandatory.
    ExplicitNull interface{}

    // Clear N-flag for the prefix-SID. The type is NflagClear. This attribute is
    // mandatory.
    NflagClear interface{}
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetFilter() yfilter.YFilter { return prefixSid.YFilter }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) SetFilter(yf yfilter.YFilter) { prefixSid.YFilter = yf }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "value" { return "Value" }
    if yname == "php" { return "Php" }
    if yname == "explicit-null" { return "ExplicitNull" }
    if yname == "nflag-clear" { return "NflagClear" }
    return ""
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetSegmentPath() string {
    return "prefix-sid"
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = prefixSid.Type
    leafs["value"] = prefixSid.Value
    leafs["php"] = prefixSid.Php
    leafs["explicit-null"] = prefixSid.ExplicitNull
    leafs["nflag-clear"] = prefixSid.NflagClear
    return leafs
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetYangName() string { return "prefix-sid" }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) SetParent(parent types.Entity) { prefixSid.parent = parent }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetParent() types.Entity { return prefixSid.parent }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable
// Fast-ReRoute configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FRR LFA candidate configuration.
    FrrlfaCandidateInterfaces Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces

    // Remote LFA maxmimum metric.
    FrrRemoteLfaMaxMetrics Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics

    // Type of FRR computation per level.
    FrrTypes Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes

    // Remote LFA Enable.
    FrrRemoteLfaTypes Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes

    // Interface FRR Default tiebreaker configuration.
    InterfaceFrrTiebreakerDefaults Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults

    // TI LFA Enable.
    FrrtilfaTypes Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes

    // FRR exclusion configuration.
    FrrExcludeInterfaces Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces

    // Interface FRR tiebreakers configuration.
    InterfaceFrrTiebreakers Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetFilter() yfilter.YFilter { return interfaceFrrTable.YFilter }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) SetFilter(yf yfilter.YFilter) { interfaceFrrTable.YFilter = yf }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetGoName(yname string) string {
    if yname == "frrlfa-candidate-interfaces" { return "FrrlfaCandidateInterfaces" }
    if yname == "frr-remote-lfa-max-metrics" { return "FrrRemoteLfaMaxMetrics" }
    if yname == "frr-types" { return "FrrTypes" }
    if yname == "frr-remote-lfa-types" { return "FrrRemoteLfaTypes" }
    if yname == "interface-frr-tiebreaker-defaults" { return "InterfaceFrrTiebreakerDefaults" }
    if yname == "frrtilfa-types" { return "FrrtilfaTypes" }
    if yname == "frr-exclude-interfaces" { return "FrrExcludeInterfaces" }
    if yname == "interface-frr-tiebreakers" { return "InterfaceFrrTiebreakers" }
    return ""
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetSegmentPath() string {
    return "interface-frr-table"
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frrlfa-candidate-interfaces" {
        return &interfaceFrrTable.FrrlfaCandidateInterfaces
    }
    if childYangName == "frr-remote-lfa-max-metrics" {
        return &interfaceFrrTable.FrrRemoteLfaMaxMetrics
    }
    if childYangName == "frr-types" {
        return &interfaceFrrTable.FrrTypes
    }
    if childYangName == "frr-remote-lfa-types" {
        return &interfaceFrrTable.FrrRemoteLfaTypes
    }
    if childYangName == "interface-frr-tiebreaker-defaults" {
        return &interfaceFrrTable.InterfaceFrrTiebreakerDefaults
    }
    if childYangName == "frrtilfa-types" {
        return &interfaceFrrTable.FrrtilfaTypes
    }
    if childYangName == "frr-exclude-interfaces" {
        return &interfaceFrrTable.FrrExcludeInterfaces
    }
    if childYangName == "interface-frr-tiebreakers" {
        return &interfaceFrrTable.InterfaceFrrTiebreakers
    }
    return nil
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frrlfa-candidate-interfaces"] = &interfaceFrrTable.FrrlfaCandidateInterfaces
    children["frr-remote-lfa-max-metrics"] = &interfaceFrrTable.FrrRemoteLfaMaxMetrics
    children["frr-types"] = &interfaceFrrTable.FrrTypes
    children["frr-remote-lfa-types"] = &interfaceFrrTable.FrrRemoteLfaTypes
    children["interface-frr-tiebreaker-defaults"] = &interfaceFrrTable.InterfaceFrrTiebreakerDefaults
    children["frrtilfa-types"] = &interfaceFrrTable.FrrtilfaTypes
    children["frr-exclude-interfaces"] = &interfaceFrrTable.FrrExcludeInterfaces
    children["interface-frr-tiebreakers"] = &interfaceFrrTable.InterfaceFrrTiebreakers
    return children
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetYangName() string { return "interface-frr-table" }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) SetParent(parent types.Entity) { interfaceFrrTable.parent = parent }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetParent() types.Entity { return interfaceFrrTable.parent }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces
// FRR LFA candidate configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Include an interface to LFA candidate in computation. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface.
    FrrlfaCandidateInterface []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetFilter() yfilter.YFilter { return frrlfaCandidateInterfaces.YFilter }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) SetFilter(yf yfilter.YFilter) { frrlfaCandidateInterfaces.YFilter = yf }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetGoName(yname string) string {
    if yname == "frrlfa-candidate-interface" { return "FrrlfaCandidateInterface" }
    return ""
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetSegmentPath() string {
    return "frrlfa-candidate-interfaces"
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frrlfa-candidate-interface" {
        for _, c := range frrlfaCandidateInterfaces.FrrlfaCandidateInterface {
            if frrlfaCandidateInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface{}
        frrlfaCandidateInterfaces.FrrlfaCandidateInterface = append(frrlfaCandidateInterfaces.FrrlfaCandidateInterface, child)
        return &frrlfaCandidateInterfaces.FrrlfaCandidateInterface[len(frrlfaCandidateInterfaces.FrrlfaCandidateInterface)-1]
    }
    return nil
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrlfaCandidateInterfaces.FrrlfaCandidateInterface {
        children[frrlfaCandidateInterfaces.FrrlfaCandidateInterface[i].GetSegmentPath()] = &frrlfaCandidateInterfaces.FrrlfaCandidateInterface[i]
    }
    return children
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetYangName() string { return "frrlfa-candidate-interfaces" }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) SetParent(parent types.Entity) { frrlfaCandidateInterfaces.parent = parent }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetParent() types.Entity { return frrlfaCandidateInterfaces.parent }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface
// Include an interface to LFA candidate
// in computation
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Level. The type is interface{} with range: 0..2. This attribute is
    // mandatory.
    Level interface{}
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetFilter() yfilter.YFilter { return frrlfaCandidateInterface.YFilter }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) SetFilter(yf yfilter.YFilter) { frrlfaCandidateInterface.YFilter = yf }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "frr-type" { return "FrrType" }
    if yname == "level" { return "Level" }
    return ""
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetSegmentPath() string {
    return "frrlfa-candidate-interface" + "[interface-name='" + fmt.Sprintf("%v", frrlfaCandidateInterface.InterfaceName) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrlfaCandidateInterface.FrrType) + "']"
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = frrlfaCandidateInterface.InterfaceName
    leafs["frr-type"] = frrlfaCandidateInterface.FrrType
    leafs["level"] = frrlfaCandidateInterface.Level
    return leafs
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetBundleName() string { return "cisco_ios_xr" }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetYangName() string { return "frrlfa-candidate-interface" }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) SetParent(parent types.Entity) { frrlfaCandidateInterface.parent = parent }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetParent() types.Entity { return frrlfaCandidateInterface.parent }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetParentYangName() string { return "frrlfa-candidate-interfaces" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics
// Remote LFA maxmimum metric
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum metric for selecting a remote LFA node. The type is
    // slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric.
    FrrRemoteLfaMaxMetric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetFilter() yfilter.YFilter { return frrRemoteLfaMaxMetrics.YFilter }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) SetFilter(yf yfilter.YFilter) { frrRemoteLfaMaxMetrics.YFilter = yf }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetGoName(yname string) string {
    if yname == "frr-remote-lfa-max-metric" { return "FrrRemoteLfaMaxMetric" }
    return ""
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetSegmentPath() string {
    return "frr-remote-lfa-max-metrics"
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-remote-lfa-max-metric" {
        for _, c := range frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric {
            if frrRemoteLfaMaxMetrics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric{}
        frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric = append(frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric, child)
        return &frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[len(frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric)-1]
    }
    return nil
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric {
        children[frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[i].GetSegmentPath()] = &frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[i]
    }
    return children
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetYangName() string { return "frr-remote-lfa-max-metrics" }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) SetParent(parent types.Entity) { frrRemoteLfaMaxMetrics.parent = parent }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetParent() types.Entity { return frrRemoteLfaMaxMetrics.parent }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric
// Configure the maximum metric for
// selecting a remote LFA node
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Value of the metric. The type is interface{} with range: 1..16777215. This
    // attribute is mandatory.
    MaxMetric interface{}
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetFilter() yfilter.YFilter { return frrRemoteLfaMaxMetric.YFilter }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) SetFilter(yf yfilter.YFilter) { frrRemoteLfaMaxMetric.YFilter = yf }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "max-metric" { return "MaxMetric" }
    return ""
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetSegmentPath() string {
    return "frr-remote-lfa-max-metric" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaMaxMetric.Level) + "']"
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrRemoteLfaMaxMetric.Level
    leafs["max-metric"] = frrRemoteLfaMaxMetric.MaxMetric
    return leafs
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetYangName() string { return "frr-remote-lfa-max-metric" }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) SetParent(parent types.Entity) { frrRemoteLfaMaxMetric.parent = parent }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetParent() types.Entity { return frrRemoteLfaMaxMetric.parent }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetParentYangName() string { return "frr-remote-lfa-max-metrics" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes
// Type of FRR computation per level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of computation for prefixes reachable via interface. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType.
    FrrType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetFilter() yfilter.YFilter { return frrTypes.YFilter }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) SetFilter(yf yfilter.YFilter) { frrTypes.YFilter = yf }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetGoName(yname string) string {
    if yname == "frr-type" { return "FrrType" }
    return ""
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetSegmentPath() string {
    return "frr-types"
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-type" {
        for _, c := range frrTypes.FrrType {
            if frrTypes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType{}
        frrTypes.FrrType = append(frrTypes.FrrType, child)
        return &frrTypes.FrrType[len(frrTypes.FrrType)-1]
    }
    return nil
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrTypes.FrrType {
        children[frrTypes.FrrType[i].GetSegmentPath()] = &frrTypes.FrrType[i]
    }
    return children
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetBundleName() string { return "cisco_ios_xr" }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetYangName() string { return "frr-types" }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) SetParent(parent types.Entity) { frrTypes.parent = parent }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetParent() types.Entity { return frrTypes.parent }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType
// Type of computation for prefixes
// reachable via interface
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Computation Type. The type is Isisfrr. This attribute is mandatory.
    Type interface{}
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetFilter() yfilter.YFilter { return frrType.YFilter }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) SetFilter(yf yfilter.YFilter) { frrType.YFilter = yf }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "type" { return "Type" }
    return ""
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetSegmentPath() string {
    return "frr-type" + "[level='" + fmt.Sprintf("%v", frrType.Level) + "']"
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrType.Level
    leafs["type"] = frrType.Type
    return leafs
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetBundleName() string { return "cisco_ios_xr" }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetYangName() string { return "frr-type" }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) SetParent(parent types.Entity) { frrType.parent = parent }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetParent() types.Entity { return frrType.parent }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetParentYangName() string { return "frr-types" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes
// Remote LFA Enable
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable remote lfa for a particular level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType.
    FrrRemoteLfaType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetFilter() yfilter.YFilter { return frrRemoteLfaTypes.YFilter }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) SetFilter(yf yfilter.YFilter) { frrRemoteLfaTypes.YFilter = yf }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetGoName(yname string) string {
    if yname == "frr-remote-lfa-type" { return "FrrRemoteLfaType" }
    return ""
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetSegmentPath() string {
    return "frr-remote-lfa-types"
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-remote-lfa-type" {
        for _, c := range frrRemoteLfaTypes.FrrRemoteLfaType {
            if frrRemoteLfaTypes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType{}
        frrRemoteLfaTypes.FrrRemoteLfaType = append(frrRemoteLfaTypes.FrrRemoteLfaType, child)
        return &frrRemoteLfaTypes.FrrRemoteLfaType[len(frrRemoteLfaTypes.FrrRemoteLfaType)-1]
    }
    return nil
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrRemoteLfaTypes.FrrRemoteLfaType {
        children[frrRemoteLfaTypes.FrrRemoteLfaType[i].GetSegmentPath()] = &frrRemoteLfaTypes.FrrRemoteLfaType[i]
    }
    return children
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetYangName() string { return "frr-remote-lfa-types" }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) SetParent(parent types.Entity) { frrRemoteLfaTypes.parent = parent }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetParent() types.Entity { return frrRemoteLfaTypes.parent }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType
// Enable remote lfa for a particular
// level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Remote LFA Type. The type is IsisRemoteLfa. This attribute is mandatory.
    Type interface{}
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetFilter() yfilter.YFilter { return frrRemoteLfaType.YFilter }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) SetFilter(yf yfilter.YFilter) { frrRemoteLfaType.YFilter = yf }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "type" { return "Type" }
    return ""
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetSegmentPath() string {
    return "frr-remote-lfa-type" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaType.Level) + "']"
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrRemoteLfaType.Level
    leafs["type"] = frrRemoteLfaType.Type
    return leafs
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetYangName() string { return "frr-remote-lfa-type" }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) SetParent(parent types.Entity) { frrRemoteLfaType.parent = parent }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetParent() types.Entity { return frrRemoteLfaType.parent }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetParentYangName() string { return "frr-remote-lfa-types" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults
// Interface FRR Default tiebreaker
// configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure default tiebreaker. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault.
    InterfaceFrrTiebreakerDefault []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetFilter() yfilter.YFilter { return interfaceFrrTiebreakerDefaults.YFilter }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) SetFilter(yf yfilter.YFilter) { interfaceFrrTiebreakerDefaults.YFilter = yf }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetGoName(yname string) string {
    if yname == "interface-frr-tiebreaker-default" { return "InterfaceFrrTiebreakerDefault" }
    return ""
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetSegmentPath() string {
    return "interface-frr-tiebreaker-defaults"
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-frr-tiebreaker-default" {
        for _, c := range interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault {
            if interfaceFrrTiebreakerDefaults.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault{}
        interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault = append(interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault, child)
        return &interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[len(interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault)-1]
    }
    return nil
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault {
        children[interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[i].GetSegmentPath()] = &interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[i]
    }
    return children
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetYangName() string { return "interface-frr-tiebreaker-defaults" }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) SetParent(parent types.Entity) { interfaceFrrTiebreakerDefaults.parent = parent }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetParent() types.Entity { return interfaceFrrTiebreakerDefaults.parent }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault
// Configure default tiebreaker
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetFilter() yfilter.YFilter { return interfaceFrrTiebreakerDefault.YFilter }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) SetFilter(yf yfilter.YFilter) { interfaceFrrTiebreakerDefault.YFilter = yf }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    return ""
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetSegmentPath() string {
    return "interface-frr-tiebreaker-default" + "[level='" + fmt.Sprintf("%v", interfaceFrrTiebreakerDefault.Level) + "']"
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = interfaceFrrTiebreakerDefault.Level
    return leafs
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetYangName() string { return "interface-frr-tiebreaker-default" }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) SetParent(parent types.Entity) { interfaceFrrTiebreakerDefault.parent = parent }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetParent() types.Entity { return interfaceFrrTiebreakerDefault.parent }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetParentYangName() string { return "interface-frr-tiebreaker-defaults" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes
// TI LFA Enable
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable TI lfa for a particular level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType.
    FrrtilfaType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetFilter() yfilter.YFilter { return frrtilfaTypes.YFilter }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) SetFilter(yf yfilter.YFilter) { frrtilfaTypes.YFilter = yf }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetGoName(yname string) string {
    if yname == "frrtilfa-type" { return "FrrtilfaType" }
    return ""
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetSegmentPath() string {
    return "frrtilfa-types"
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frrtilfa-type" {
        for _, c := range frrtilfaTypes.FrrtilfaType {
            if frrtilfaTypes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType{}
        frrtilfaTypes.FrrtilfaType = append(frrtilfaTypes.FrrtilfaType, child)
        return &frrtilfaTypes.FrrtilfaType[len(frrtilfaTypes.FrrtilfaType)-1]
    }
    return nil
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrtilfaTypes.FrrtilfaType {
        children[frrtilfaTypes.FrrtilfaType[i].GetSegmentPath()] = &frrtilfaTypes.FrrtilfaType[i]
    }
    return children
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetBundleName() string { return "cisco_ios_xr" }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetYangName() string { return "frrtilfa-types" }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) SetParent(parent types.Entity) { frrtilfaTypes.parent = parent }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetParent() types.Entity { return frrtilfaTypes.parent }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType
// Enable TI lfa for a particular level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetFilter() yfilter.YFilter { return frrtilfaType.YFilter }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) SetFilter(yf yfilter.YFilter) { frrtilfaType.YFilter = yf }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    return ""
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetSegmentPath() string {
    return "frrtilfa-type" + "[level='" + fmt.Sprintf("%v", frrtilfaType.Level) + "']"
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrtilfaType.Level
    return leafs
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetBundleName() string { return "cisco_ios_xr" }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetYangName() string { return "frrtilfa-type" }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) SetParent(parent types.Entity) { frrtilfaType.parent = parent }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetParent() types.Entity { return frrtilfaType.parent }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetParentYangName() string { return "frrtilfa-types" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces
// FRR exclusion configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exclude an interface from computation. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface.
    FrrExcludeInterface []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetFilter() yfilter.YFilter { return frrExcludeInterfaces.YFilter }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) SetFilter(yf yfilter.YFilter) { frrExcludeInterfaces.YFilter = yf }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetGoName(yname string) string {
    if yname == "frr-exclude-interface" { return "FrrExcludeInterface" }
    return ""
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetSegmentPath() string {
    return "frr-exclude-interfaces"
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-exclude-interface" {
        for _, c := range frrExcludeInterfaces.FrrExcludeInterface {
            if frrExcludeInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface{}
        frrExcludeInterfaces.FrrExcludeInterface = append(frrExcludeInterfaces.FrrExcludeInterface, child)
        return &frrExcludeInterfaces.FrrExcludeInterface[len(frrExcludeInterfaces.FrrExcludeInterface)-1]
    }
    return nil
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrExcludeInterfaces.FrrExcludeInterface {
        children[frrExcludeInterfaces.FrrExcludeInterface[i].GetSegmentPath()] = &frrExcludeInterfaces.FrrExcludeInterface[i]
    }
    return children
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetYangName() string { return "frr-exclude-interfaces" }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) SetParent(parent types.Entity) { frrExcludeInterfaces.parent = parent }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetParent() types.Entity { return frrExcludeInterfaces.parent }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface
// Exclude an interface from computation
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Level. The type is interface{} with range: 0..2. This attribute is
    // mandatory.
    Level interface{}
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetFilter() yfilter.YFilter { return frrExcludeInterface.YFilter }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) SetFilter(yf yfilter.YFilter) { frrExcludeInterface.YFilter = yf }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "frr-type" { return "FrrType" }
    if yname == "level" { return "Level" }
    return ""
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetSegmentPath() string {
    return "frr-exclude-interface" + "[interface-name='" + fmt.Sprintf("%v", frrExcludeInterface.InterfaceName) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrExcludeInterface.FrrType) + "']"
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = frrExcludeInterface.InterfaceName
    leafs["frr-type"] = frrExcludeInterface.FrrType
    leafs["level"] = frrExcludeInterface.Level
    return leafs
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetBundleName() string { return "cisco_ios_xr" }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetYangName() string { return "frr-exclude-interface" }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) SetParent(parent types.Entity) { frrExcludeInterface.parent = parent }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetParent() types.Entity { return frrExcludeInterface.parent }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetParentYangName() string { return "frr-exclude-interfaces" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers
// Interface FRR tiebreakers configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure tiebreaker for multiple backups. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker.
    InterfaceFrrTiebreaker []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetFilter() yfilter.YFilter { return interfaceFrrTiebreakers.YFilter }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) SetFilter(yf yfilter.YFilter) { interfaceFrrTiebreakers.YFilter = yf }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetGoName(yname string) string {
    if yname == "interface-frr-tiebreaker" { return "InterfaceFrrTiebreaker" }
    return ""
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetSegmentPath() string {
    return "interface-frr-tiebreakers"
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-frr-tiebreaker" {
        for _, c := range interfaceFrrTiebreakers.InterfaceFrrTiebreaker {
            if interfaceFrrTiebreakers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker{}
        interfaceFrrTiebreakers.InterfaceFrrTiebreaker = append(interfaceFrrTiebreakers.InterfaceFrrTiebreaker, child)
        return &interfaceFrrTiebreakers.InterfaceFrrTiebreaker[len(interfaceFrrTiebreakers.InterfaceFrrTiebreaker)-1]
    }
    return nil
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceFrrTiebreakers.InterfaceFrrTiebreaker {
        children[interfaceFrrTiebreakers.InterfaceFrrTiebreaker[i].GetSegmentPath()] = &interfaceFrrTiebreakers.InterfaceFrrTiebreaker[i]
    }
    return children
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetYangName() string { return "interface-frr-tiebreakers" }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) SetParent(parent types.Entity) { interfaceFrrTiebreakers.parent = parent }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetParent() types.Entity { return interfaceFrrTiebreakers.parent }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker
// Configure tiebreaker for multiple
// backups
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Tiebreaker for which configuration applies. The
    // type is IsisInterfaceFrrTiebreaker.
    Tiebreaker interface{}

    // Preference order among tiebreakers. The type is interface{} with range:
    // 1..255. This attribute is mandatory.
    Index interface{}
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetFilter() yfilter.YFilter { return interfaceFrrTiebreaker.YFilter }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) SetFilter(yf yfilter.YFilter) { interfaceFrrTiebreaker.YFilter = yf }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "tiebreaker" { return "Tiebreaker" }
    if yname == "index" { return "Index" }
    return ""
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetSegmentPath() string {
    return "interface-frr-tiebreaker" + "[level='" + fmt.Sprintf("%v", interfaceFrrTiebreaker.Level) + "']" + "[tiebreaker='" + fmt.Sprintf("%v", interfaceFrrTiebreaker.Tiebreaker) + "']"
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = interfaceFrrTiebreaker.Level
    leafs["tiebreaker"] = interfaceFrrTiebreaker.Tiebreaker
    leafs["index"] = interfaceFrrTiebreaker.Index
    return leafs
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetYangName() string { return "interface-frr-tiebreaker" }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) SetParent(parent types.Entity) { interfaceFrrTiebreaker.parent = parent }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetParent() types.Entity { return interfaceFrrTiebreaker.parent }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetParentYangName() string { return "interface-frr-tiebreakers" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp
// MPLS LDP configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable MPLS LDP Synchronization for an IS-IS level. The type is interface{}
    // with range: 0..2. The default value is 0.
    SyncLevel interface{}
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetFilter() yfilter.YFilter { return mplsLdp.YFilter }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) SetFilter(yf yfilter.YFilter) { mplsLdp.YFilter = yf }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetGoName(yname string) string {
    if yname == "sync-level" { return "SyncLevel" }
    return ""
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetSegmentPath() string {
    return "mpls-ldp"
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sync-level"] = mplsLdp.SyncLevel
    return leafs
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetBundleName() string { return "cisco_ios_xr" }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetYangName() string { return "mpls-ldp" }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) SetParent(parent types.Entity) { mplsLdp.parent = parent }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetParent() types.Entity { return mplsLdp.parent }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid
// Assign prefix SSPF SID to an interface,
// ISISPHPFlag will be rejected if set to
// disable, ISISEXPLICITNULLFlag will
// override the value of ISISPHPFlag
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type for the interface. The type is Isissid1. This attribute is
    // mandatory.
    Type interface{}

    // SID value for the interface. The type is interface{} with range:
    // 0..1048575. This attribute is mandatory.
    Value interface{}

    // Enable/Disable Penultimate Hop Popping. The type is IsisphpFlag. This
    // attribute is mandatory.
    Php interface{}

    // Enable/Disable Explicit-NULL flag. The type is IsisexplicitNullFlag. This
    // attribute is mandatory.
    ExplicitNull interface{}

    // Clear N-flag for the prefix-SID. The type is NflagClear. This attribute is
    // mandatory.
    NflagClear interface{}
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetFilter() yfilter.YFilter { return prefixSspfsid.YFilter }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) SetFilter(yf yfilter.YFilter) { prefixSspfsid.YFilter = yf }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "value" { return "Value" }
    if yname == "php" { return "Php" }
    if yname == "explicit-null" { return "ExplicitNull" }
    if yname == "nflag-clear" { return "NflagClear" }
    return ""
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetSegmentPath() string {
    return "prefix-sspfsid"
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = prefixSspfsid.Type
    leafs["value"] = prefixSspfsid.Value
    leafs["php"] = prefixSspfsid.Php
    leafs["explicit-null"] = prefixSspfsid.ExplicitNull
    leafs["nflag-clear"] = prefixSspfsid.NflagClear
    return leafs
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetYangName() string { return "prefix-sspfsid" }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) SetParent(parent types.Entity) { prefixSspfsid.parent = parent }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetParent() types.Entity { return prefixSspfsid.parent }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics
// AutoMetric configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoMetric Proactive-Protect configuration. Legal value depends on the
    // metric-style specified for the topology. If the metric-style defined is
    // narrow, then only a value between <1-63> is allowed and if the metric-style
    // is defined as wide, then a value between <1-16777214> is allowed as the
    // auto-metric value. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric.
    AutoMetric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetFilter() yfilter.YFilter { return autoMetrics.YFilter }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) SetFilter(yf yfilter.YFilter) { autoMetrics.YFilter = yf }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetGoName(yname string) string {
    if yname == "auto-metric" { return "AutoMetric" }
    return ""
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetSegmentPath() string {
    return "auto-metrics"
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto-metric" {
        for _, c := range autoMetrics.AutoMetric {
            if autoMetrics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric{}
        autoMetrics.AutoMetric = append(autoMetrics.AutoMetric, child)
        return &autoMetrics.AutoMetric[len(autoMetrics.AutoMetric)-1]
    }
    return nil
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range autoMetrics.AutoMetric {
        children[autoMetrics.AutoMetric[i].GetSegmentPath()] = &autoMetrics.AutoMetric[i]
    }
    return children
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetBundleName() string { return "cisco_ios_xr" }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetYangName() string { return "auto-metrics" }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) SetParent(parent types.Entity) { autoMetrics.parent = parent }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetParent() types.Entity { return autoMetrics.parent }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric
// AutoMetric Proactive-Protect
// configuration. Legal value depends on
// the metric-style specified for the
// topology. If the metric-style defined is
// narrow, then only a value between <1-63>
// is allowed and if the metric-style is
// defined as wide, then a value between
// <1-16777214> is allowed as the
// auto-metric value.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed auto metric:<1-63> for narrow ,<1-16777214> for wide. The type is
    // interface{} with range: 1..16777214. This attribute is mandatory.
    ProactiveProtect interface{}
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetFilter() yfilter.YFilter { return autoMetric.YFilter }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) SetFilter(yf yfilter.YFilter) { autoMetric.YFilter = yf }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "proactive-protect" { return "ProactiveProtect" }
    return ""
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetSegmentPath() string {
    return "auto-metric" + "[level='" + fmt.Sprintf("%v", autoMetric.Level) + "']"
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = autoMetric.Level
    leafs["proactive-protect"] = autoMetric.ProactiveProtect
    return leafs
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetBundleName() string { return "cisco_ios_xr" }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetYangName() string { return "auto-metric" }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) SetParent(parent types.Entity) { autoMetric.parent = parent }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetParent() types.Entity { return autoMetric.parent }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetParentYangName() string { return "auto-metrics" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags
// admin-tag configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Admin tag for advertised interface connected routes. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag.
    AdminTag []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetFilter() yfilter.YFilter { return adminTags.YFilter }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) SetFilter(yf yfilter.YFilter) { adminTags.YFilter = yf }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetGoName(yname string) string {
    if yname == "admin-tag" { return "AdminTag" }
    return ""
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetSegmentPath() string {
    return "admin-tags"
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "admin-tag" {
        for _, c := range adminTags.AdminTag {
            if adminTags.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag{}
        adminTags.AdminTag = append(adminTags.AdminTag, child)
        return &adminTags.AdminTag[len(adminTags.AdminTag)-1]
    }
    return nil
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range adminTags.AdminTag {
        children[adminTags.AdminTag[i].GetSegmentPath()] = &adminTags.AdminTag[i]
    }
    return children
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetBundleName() string { return "cisco_ios_xr" }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetYangName() string { return "admin-tags" }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) SetParent(parent types.Entity) { adminTags.parent = parent }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetParent() types.Entity { return adminTags.parent }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag
// Admin tag for advertised interface
// connected routes
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Tag to associate with connected routes. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    AdminTag interface{}
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetFilter() yfilter.YFilter { return adminTag.YFilter }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) SetFilter(yf yfilter.YFilter) { adminTag.YFilter = yf }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "admin-tag" { return "AdminTag" }
    return ""
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetSegmentPath() string {
    return "admin-tag" + "[level='" + fmt.Sprintf("%v", adminTag.Level) + "']"
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = adminTag.Level
    leafs["admin-tag"] = adminTag.AdminTag
    return leafs
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetBundleName() string { return "cisco_ios_xr" }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetYangName() string { return "admin-tag" }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) SetParent(parent types.Entity) { adminTag.parent = parent }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetParent() types.Entity { return adminTag.parent }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetParentYangName() string { return "admin-tags" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup
// Provide link group name and level
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Group. The type is string with length: 1..40. This attribute is
    // mandatory.
    LinkGroup interface{}

    // Level in which link group will be effective. The type is interface{} with
    // range: 0..2. The default value is 0.
    Level interface{}
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetFilter() yfilter.YFilter { return interfaceLinkGroup.YFilter }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) SetFilter(yf yfilter.YFilter) { interfaceLinkGroup.YFilter = yf }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetGoName(yname string) string {
    if yname == "link-group" { return "LinkGroup" }
    if yname == "level" { return "Level" }
    return ""
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetSegmentPath() string {
    return "interface-link-group"
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-group"] = interfaceLinkGroup.LinkGroup
    leafs["level"] = interfaceLinkGroup.Level
    return leafs
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetYangName() string { return "interface-link-group" }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) SetParent(parent types.Entity) { interfaceLinkGroup.parent = parent }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetParent() types.Entity { return interfaceLinkGroup.parent }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids
// Manual Adjacecy SID configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Assign adjancency SID to an interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid.
    ManualAdjSid []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetFilter() yfilter.YFilter { return manualAdjSids.YFilter }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) SetFilter(yf yfilter.YFilter) { manualAdjSids.YFilter = yf }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetGoName(yname string) string {
    if yname == "manual-adj-sid" { return "ManualAdjSid" }
    return ""
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetSegmentPath() string {
    return "manual-adj-sids"
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "manual-adj-sid" {
        for _, c := range manualAdjSids.ManualAdjSid {
            if manualAdjSids.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid{}
        manualAdjSids.ManualAdjSid = append(manualAdjSids.ManualAdjSid, child)
        return &manualAdjSids.ManualAdjSid[len(manualAdjSids.ManualAdjSid)-1]
    }
    return nil
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range manualAdjSids.ManualAdjSid {
        children[manualAdjSids.ManualAdjSid[i].GetSegmentPath()] = &manualAdjSids.ManualAdjSid[i]
    }
    return children
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetBundleName() string { return "cisco_ios_xr" }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetYangName() string { return "manual-adj-sids" }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) SetParent(parent types.Entity) { manualAdjSids.parent = parent }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetParent() types.Entity { return manualAdjSids.parent }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid
// Assign adjancency SID to an interface
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Sid type aboslute or index. The type is Isissid1.
    SidType interface{}

    // This attribute is a key. Sid value. The type is interface{} with range:
    // 0..1048575.
    Sid interface{}

    // Enable/Disable SID protection. The type is IsissidProtected. This attribute
    // is mandatory.
    Protected interface{}
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetFilter() yfilter.YFilter { return manualAdjSid.YFilter }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) SetFilter(yf yfilter.YFilter) { manualAdjSid.YFilter = yf }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "sid-type" { return "SidType" }
    if yname == "sid" { return "Sid" }
    if yname == "protected" { return "Protected" }
    return ""
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetSegmentPath() string {
    return "manual-adj-sid" + "[level='" + fmt.Sprintf("%v", manualAdjSid.Level) + "']" + "[sid-type='" + fmt.Sprintf("%v", manualAdjSid.SidType) + "']" + "[sid='" + fmt.Sprintf("%v", manualAdjSid.Sid) + "']"
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = manualAdjSid.Level
    leafs["sid-type"] = manualAdjSid.SidType
    leafs["sid"] = manualAdjSid.Sid
    leafs["protected"] = manualAdjSid.Protected
    return leafs
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetBundleName() string { return "cisco_ios_xr" }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetYangName() string { return "manual-adj-sid" }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) SetParent(parent types.Entity) { manualAdjSid.parent = parent }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetParent() types.Entity { return manualAdjSid.parent }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetParentYangName() string { return "manual-adj-sids" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics
// Metric configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric configuration. Legal value depends on the metric-style specified for
    // the topology. If the metric-style defined is narrow, then only a value
    // between <1-63> is allowed and if the metric-style is defined as wide, then
    // a value between <1-16777215> is allowed as the metric value.  All routers
    // exclude links with the maximum wide metric (16777215) from their SPF. The
    // type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric.
    Metric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetFilter() yfilter.YFilter { return metrics.YFilter }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) SetFilter(yf yfilter.YFilter) { metrics.YFilter = yf }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    return ""
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetSegmentPath() string {
    return "metrics"
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metric" {
        for _, c := range metrics.Metric {
            if metrics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric{}
        metrics.Metric = append(metrics.Metric, child)
        return &metrics.Metric[len(metrics.Metric)-1]
    }
    return nil
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range metrics.Metric {
        children[metrics.Metric[i].GetSegmentPath()] = &metrics.Metric[i]
    }
    return children
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetBundleName() string { return "cisco_ios_xr" }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetYangName() string { return "metrics" }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) SetParent(parent types.Entity) { metrics.parent = parent }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetParent() types.Entity { return metrics.parent }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric
// Metric configuration. Legal value depends on
// the metric-style specified for the topology. If
// the metric-style defined is narrow, then only a
// value between <1-63> is allowed and if the
// metric-style is defined as wide, then a value
// between <1-16777215> is allowed as the metric
// value.  All routers exclude links with the
// maximum wide metric (16777215) from their SPF
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed metric: <1-63> for narrow, <1-16777215> for wide. The type is one
    // of the following types: enumeration
    // Isis.Instances.Instance.Interfaces.Interface.InterfaceAfs.InterfaceAf.TopologyName.Metrics.Metric.Metric
    // This attribute is mandatory., or int with range: 1..16777215 This attribute
    // is mandatory..
    Metric interface{}
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetFilter() yfilter.YFilter { return metric.YFilter }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) SetFilter(yf yfilter.YFilter) { metric.YFilter = yf }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetSegmentPath() string {
    return "metric" + "[level='" + fmt.Sprintf("%v", metric.Level) + "']"
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = metric.Level
    leafs["metric"] = metric.Metric
    return leafs
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetBundleName() string { return "cisco_ios_xr" }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetYangName() string { return "metric" }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) SetParent(parent types.Entity) { metric.parent = parent }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetParent() types.Entity { return metric.parent }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetParentYangName() string { return "metrics" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric_Metric represents <1-16777215> for wide
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric_Metric string

const (
    // Maximum wide metric.  All routers will
    // exclude this link from their SPF
    Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric_Metric_maximum Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric_Metric = "maximum"
)

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights
// Weight configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Weight configuration under interface for load balancing. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight.
    Weight []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetFilter() yfilter.YFilter { return weights.YFilter }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) SetFilter(yf yfilter.YFilter) { weights.YFilter = yf }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetGoName(yname string) string {
    if yname == "weight" { return "Weight" }
    return ""
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetSegmentPath() string {
    return "weights"
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "weight" {
        for _, c := range weights.Weight {
            if weights.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight{}
        weights.Weight = append(weights.Weight, child)
        return &weights.Weight[len(weights.Weight)-1]
    }
    return nil
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range weights.Weight {
        children[weights.Weight[i].GetSegmentPath()] = &weights.Weight[i]
    }
    return children
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetBundleName() string { return "cisco_ios_xr" }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetYangName() string { return "weights" }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) SetParent(parent types.Entity) { weights.parent = parent }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetParent() types.Entity { return weights.parent }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetParentYangName() string { return "interface-af-data" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight
// Weight configuration under interface for load
// balancing
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Weight to be configured under interface for Load Balancing. Allowed weight:
    // <1-16777215>. The type is interface{} with range: 1..16777214. This
    // attribute is mandatory.
    Weight interface{}
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetFilter() yfilter.YFilter { return weight.YFilter }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) SetFilter(yf yfilter.YFilter) { weight.YFilter = yf }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "weight" { return "Weight" }
    return ""
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetSegmentPath() string {
    return "weight" + "[level='" + fmt.Sprintf("%v", weight.Level) + "']"
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = weight.Level
    leafs["weight"] = weight.Weight
    return leafs
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetBundleName() string { return "cisco_ios_xr" }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetYangName() string { return "weight" }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) SetParent(parent types.Entity) { weight.parent = parent }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetParent() types.Entity { return weight.parent }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetParentYangName() string { return "weights" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName
// keys: topology-name
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Topology Name. The type is string with length:
    // 1..32.
    TopologyName interface{}

    // Interface state. The type is IsisInterfaceAfState.
    InterfaceAfState interface{}

    // The presence of this object allows an address-family to be run over the
    // interface in question.This must be the first object created under the
    // InterfaceAddressFamily container, and the last one deleted. The type is
    // interface{}.
    Running interface{}

    // Assign prefix SID to an interface, ISISPHPFlag will be rejected if set to
    // disable, ISISEXPLICITNULLFlag will override the value of ISISPHPFlag.
    PrefixSid Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid

    // Fast-ReRoute configuration.
    InterfaceFrrTable Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable

    // MPLS LDP configuration.
    MplsLdp Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp

    // Assign prefix SSPF SID to an interface, ISISPHPFlag will be rejected if set
    // to disable, ISISEXPLICITNULLFlag will override the value of ISISPHPFlag.
    PrefixSspfsid Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid

    // AutoMetric configuration.
    AutoMetrics Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics

    // admin-tag configuration.
    AdminTags Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags

    // Provide link group name and level.
    InterfaceLinkGroup Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup

    // Manual Adjacecy SID configuration.
    ManualAdjSids Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids

    // Metric configuration.
    Metrics Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics

    // Weight configuration.
    Weights Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights
}

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetFilter() yfilter.YFilter { return topologyName.YFilter }

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) SetFilter(yf yfilter.YFilter) { topologyName.YFilter = yf }

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetGoName(yname string) string {
    if yname == "topology-name" { return "TopologyName" }
    if yname == "interface-af-state" { return "InterfaceAfState" }
    if yname == "running" { return "Running" }
    if yname == "prefix-sid" { return "PrefixSid" }
    if yname == "interface-frr-table" { return "InterfaceFrrTable" }
    if yname == "mpls-ldp" { return "MplsLdp" }
    if yname == "prefix-sspfsid" { return "PrefixSspfsid" }
    if yname == "auto-metrics" { return "AutoMetrics" }
    if yname == "admin-tags" { return "AdminTags" }
    if yname == "interface-link-group" { return "InterfaceLinkGroup" }
    if yname == "manual-adj-sids" { return "ManualAdjSids" }
    if yname == "metrics" { return "Metrics" }
    if yname == "weights" { return "Weights" }
    return ""
}

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetSegmentPath() string {
    return "topology-name" + "[topology-name='" + fmt.Sprintf("%v", topologyName.TopologyName) + "']"
}

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-sid" {
        return &topologyName.PrefixSid
    }
    if childYangName == "interface-frr-table" {
        return &topologyName.InterfaceFrrTable
    }
    if childYangName == "mpls-ldp" {
        return &topologyName.MplsLdp
    }
    if childYangName == "prefix-sspfsid" {
        return &topologyName.PrefixSspfsid
    }
    if childYangName == "auto-metrics" {
        return &topologyName.AutoMetrics
    }
    if childYangName == "admin-tags" {
        return &topologyName.AdminTags
    }
    if childYangName == "interface-link-group" {
        return &topologyName.InterfaceLinkGroup
    }
    if childYangName == "manual-adj-sids" {
        return &topologyName.ManualAdjSids
    }
    if childYangName == "metrics" {
        return &topologyName.Metrics
    }
    if childYangName == "weights" {
        return &topologyName.Weights
    }
    return nil
}

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-sid"] = &topologyName.PrefixSid
    children["interface-frr-table"] = &topologyName.InterfaceFrrTable
    children["mpls-ldp"] = &topologyName.MplsLdp
    children["prefix-sspfsid"] = &topologyName.PrefixSspfsid
    children["auto-metrics"] = &topologyName.AutoMetrics
    children["admin-tags"] = &topologyName.AdminTags
    children["interface-link-group"] = &topologyName.InterfaceLinkGroup
    children["manual-adj-sids"] = &topologyName.ManualAdjSids
    children["metrics"] = &topologyName.Metrics
    children["weights"] = &topologyName.Weights
    return children
}

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["topology-name"] = topologyName.TopologyName
    leafs["interface-af-state"] = topologyName.InterfaceAfState
    leafs["running"] = topologyName.Running
    return leafs
}

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetBundleName() string { return "cisco_ios_xr" }

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetYangName() string { return "topology-name" }

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) SetParent(parent types.Entity) { topologyName.parent = parent }

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetParent() types.Entity { return topologyName.parent }

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetParentYangName() string { return "interface-af" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid
// Assign prefix SID to an interface,
// ISISPHPFlag will be rejected if set to
// disable, ISISEXPLICITNULLFlag will
// override the value of ISISPHPFlag
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type for the interface. The type is Isissid1. This attribute is
    // mandatory.
    Type interface{}

    // SID value for the interface. The type is interface{} with range:
    // 0..1048575. This attribute is mandatory.
    Value interface{}

    // Enable/Disable Penultimate Hop Popping. The type is IsisphpFlag. This
    // attribute is mandatory.
    Php interface{}

    // Enable/Disable Explicit-NULL flag. The type is IsisexplicitNullFlag. This
    // attribute is mandatory.
    ExplicitNull interface{}

    // Clear N-flag for the prefix-SID. The type is NflagClear. This attribute is
    // mandatory.
    NflagClear interface{}
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetFilter() yfilter.YFilter { return prefixSid.YFilter }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) SetFilter(yf yfilter.YFilter) { prefixSid.YFilter = yf }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "value" { return "Value" }
    if yname == "php" { return "Php" }
    if yname == "explicit-null" { return "ExplicitNull" }
    if yname == "nflag-clear" { return "NflagClear" }
    return ""
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetSegmentPath() string {
    return "prefix-sid"
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = prefixSid.Type
    leafs["value"] = prefixSid.Value
    leafs["php"] = prefixSid.Php
    leafs["explicit-null"] = prefixSid.ExplicitNull
    leafs["nflag-clear"] = prefixSid.NflagClear
    return leafs
}

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetYangName() string { return "prefix-sid" }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) SetParent(parent types.Entity) { prefixSid.parent = parent }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetParent() types.Entity { return prefixSid.parent }

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable
// Fast-ReRoute configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FRR LFA candidate configuration.
    FrrlfaCandidateInterfaces Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces

    // Remote LFA maxmimum metric.
    FrrRemoteLfaMaxMetrics Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics

    // Type of FRR computation per level.
    FrrTypes Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes

    // Remote LFA Enable.
    FrrRemoteLfaTypes Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes

    // Interface FRR Default tiebreaker configuration.
    InterfaceFrrTiebreakerDefaults Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults

    // TI LFA Enable.
    FrrtilfaTypes Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes

    // FRR exclusion configuration.
    FrrExcludeInterfaces Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces

    // Interface FRR tiebreakers configuration.
    InterfaceFrrTiebreakers Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetFilter() yfilter.YFilter { return interfaceFrrTable.YFilter }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) SetFilter(yf yfilter.YFilter) { interfaceFrrTable.YFilter = yf }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetGoName(yname string) string {
    if yname == "frrlfa-candidate-interfaces" { return "FrrlfaCandidateInterfaces" }
    if yname == "frr-remote-lfa-max-metrics" { return "FrrRemoteLfaMaxMetrics" }
    if yname == "frr-types" { return "FrrTypes" }
    if yname == "frr-remote-lfa-types" { return "FrrRemoteLfaTypes" }
    if yname == "interface-frr-tiebreaker-defaults" { return "InterfaceFrrTiebreakerDefaults" }
    if yname == "frrtilfa-types" { return "FrrtilfaTypes" }
    if yname == "frr-exclude-interfaces" { return "FrrExcludeInterfaces" }
    if yname == "interface-frr-tiebreakers" { return "InterfaceFrrTiebreakers" }
    return ""
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetSegmentPath() string {
    return "interface-frr-table"
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frrlfa-candidate-interfaces" {
        return &interfaceFrrTable.FrrlfaCandidateInterfaces
    }
    if childYangName == "frr-remote-lfa-max-metrics" {
        return &interfaceFrrTable.FrrRemoteLfaMaxMetrics
    }
    if childYangName == "frr-types" {
        return &interfaceFrrTable.FrrTypes
    }
    if childYangName == "frr-remote-lfa-types" {
        return &interfaceFrrTable.FrrRemoteLfaTypes
    }
    if childYangName == "interface-frr-tiebreaker-defaults" {
        return &interfaceFrrTable.InterfaceFrrTiebreakerDefaults
    }
    if childYangName == "frrtilfa-types" {
        return &interfaceFrrTable.FrrtilfaTypes
    }
    if childYangName == "frr-exclude-interfaces" {
        return &interfaceFrrTable.FrrExcludeInterfaces
    }
    if childYangName == "interface-frr-tiebreakers" {
        return &interfaceFrrTable.InterfaceFrrTiebreakers
    }
    return nil
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["frrlfa-candidate-interfaces"] = &interfaceFrrTable.FrrlfaCandidateInterfaces
    children["frr-remote-lfa-max-metrics"] = &interfaceFrrTable.FrrRemoteLfaMaxMetrics
    children["frr-types"] = &interfaceFrrTable.FrrTypes
    children["frr-remote-lfa-types"] = &interfaceFrrTable.FrrRemoteLfaTypes
    children["interface-frr-tiebreaker-defaults"] = &interfaceFrrTable.InterfaceFrrTiebreakerDefaults
    children["frrtilfa-types"] = &interfaceFrrTable.FrrtilfaTypes
    children["frr-exclude-interfaces"] = &interfaceFrrTable.FrrExcludeInterfaces
    children["interface-frr-tiebreakers"] = &interfaceFrrTable.InterfaceFrrTiebreakers
    return children
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetYangName() string { return "interface-frr-table" }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) SetParent(parent types.Entity) { interfaceFrrTable.parent = parent }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetParent() types.Entity { return interfaceFrrTable.parent }

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces
// FRR LFA candidate configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Include an interface to LFA candidate in computation. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface.
    FrrlfaCandidateInterface []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetFilter() yfilter.YFilter { return frrlfaCandidateInterfaces.YFilter }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) SetFilter(yf yfilter.YFilter) { frrlfaCandidateInterfaces.YFilter = yf }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetGoName(yname string) string {
    if yname == "frrlfa-candidate-interface" { return "FrrlfaCandidateInterface" }
    return ""
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetSegmentPath() string {
    return "frrlfa-candidate-interfaces"
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frrlfa-candidate-interface" {
        for _, c := range frrlfaCandidateInterfaces.FrrlfaCandidateInterface {
            if frrlfaCandidateInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface{}
        frrlfaCandidateInterfaces.FrrlfaCandidateInterface = append(frrlfaCandidateInterfaces.FrrlfaCandidateInterface, child)
        return &frrlfaCandidateInterfaces.FrrlfaCandidateInterface[len(frrlfaCandidateInterfaces.FrrlfaCandidateInterface)-1]
    }
    return nil
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrlfaCandidateInterfaces.FrrlfaCandidateInterface {
        children[frrlfaCandidateInterfaces.FrrlfaCandidateInterface[i].GetSegmentPath()] = &frrlfaCandidateInterfaces.FrrlfaCandidateInterface[i]
    }
    return children
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetYangName() string { return "frrlfa-candidate-interfaces" }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) SetParent(parent types.Entity) { frrlfaCandidateInterfaces.parent = parent }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetParent() types.Entity { return frrlfaCandidateInterfaces.parent }

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface
// Include an interface to LFA candidate
// in computation
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Level. The type is interface{} with range: 0..2. This attribute is
    // mandatory.
    Level interface{}
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetFilter() yfilter.YFilter { return frrlfaCandidateInterface.YFilter }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) SetFilter(yf yfilter.YFilter) { frrlfaCandidateInterface.YFilter = yf }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "frr-type" { return "FrrType" }
    if yname == "level" { return "Level" }
    return ""
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetSegmentPath() string {
    return "frrlfa-candidate-interface" + "[interface-name='" + fmt.Sprintf("%v", frrlfaCandidateInterface.InterfaceName) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrlfaCandidateInterface.FrrType) + "']"
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = frrlfaCandidateInterface.InterfaceName
    leafs["frr-type"] = frrlfaCandidateInterface.FrrType
    leafs["level"] = frrlfaCandidateInterface.Level
    return leafs
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetBundleName() string { return "cisco_ios_xr" }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetYangName() string { return "frrlfa-candidate-interface" }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) SetParent(parent types.Entity) { frrlfaCandidateInterface.parent = parent }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetParent() types.Entity { return frrlfaCandidateInterface.parent }

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetParentYangName() string { return "frrlfa-candidate-interfaces" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics
// Remote LFA maxmimum metric
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the maximum metric for selecting a remote LFA node. The type is
    // slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric.
    FrrRemoteLfaMaxMetric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetFilter() yfilter.YFilter { return frrRemoteLfaMaxMetrics.YFilter }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) SetFilter(yf yfilter.YFilter) { frrRemoteLfaMaxMetrics.YFilter = yf }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetGoName(yname string) string {
    if yname == "frr-remote-lfa-max-metric" { return "FrrRemoteLfaMaxMetric" }
    return ""
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetSegmentPath() string {
    return "frr-remote-lfa-max-metrics"
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-remote-lfa-max-metric" {
        for _, c := range frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric {
            if frrRemoteLfaMaxMetrics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric{}
        frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric = append(frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric, child)
        return &frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[len(frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric)-1]
    }
    return nil
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric {
        children[frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[i].GetSegmentPath()] = &frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[i]
    }
    return children
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetYangName() string { return "frr-remote-lfa-max-metrics" }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) SetParent(parent types.Entity) { frrRemoteLfaMaxMetrics.parent = parent }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetParent() types.Entity { return frrRemoteLfaMaxMetrics.parent }

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric
// Configure the maximum metric for
// selecting a remote LFA node
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Value of the metric. The type is interface{} with range: 1..16777215. This
    // attribute is mandatory.
    MaxMetric interface{}
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetFilter() yfilter.YFilter { return frrRemoteLfaMaxMetric.YFilter }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) SetFilter(yf yfilter.YFilter) { frrRemoteLfaMaxMetric.YFilter = yf }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "max-metric" { return "MaxMetric" }
    return ""
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetSegmentPath() string {
    return "frr-remote-lfa-max-metric" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaMaxMetric.Level) + "']"
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrRemoteLfaMaxMetric.Level
    leafs["max-metric"] = frrRemoteLfaMaxMetric.MaxMetric
    return leafs
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetYangName() string { return "frr-remote-lfa-max-metric" }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) SetParent(parent types.Entity) { frrRemoteLfaMaxMetric.parent = parent }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetParent() types.Entity { return frrRemoteLfaMaxMetric.parent }

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetParentYangName() string { return "frr-remote-lfa-max-metrics" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes
// Type of FRR computation per level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of computation for prefixes reachable via interface. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType.
    FrrType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetFilter() yfilter.YFilter { return frrTypes.YFilter }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) SetFilter(yf yfilter.YFilter) { frrTypes.YFilter = yf }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetGoName(yname string) string {
    if yname == "frr-type" { return "FrrType" }
    return ""
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetSegmentPath() string {
    return "frr-types"
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-type" {
        for _, c := range frrTypes.FrrType {
            if frrTypes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType{}
        frrTypes.FrrType = append(frrTypes.FrrType, child)
        return &frrTypes.FrrType[len(frrTypes.FrrType)-1]
    }
    return nil
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrTypes.FrrType {
        children[frrTypes.FrrType[i].GetSegmentPath()] = &frrTypes.FrrType[i]
    }
    return children
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetBundleName() string { return "cisco_ios_xr" }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetYangName() string { return "frr-types" }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) SetParent(parent types.Entity) { frrTypes.parent = parent }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetParent() types.Entity { return frrTypes.parent }

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType
// Type of computation for prefixes
// reachable via interface
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Computation Type. The type is Isisfrr. This attribute is mandatory.
    Type interface{}
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetFilter() yfilter.YFilter { return frrType.YFilter }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) SetFilter(yf yfilter.YFilter) { frrType.YFilter = yf }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "type" { return "Type" }
    return ""
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetSegmentPath() string {
    return "frr-type" + "[level='" + fmt.Sprintf("%v", frrType.Level) + "']"
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrType.Level
    leafs["type"] = frrType.Type
    return leafs
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetBundleName() string { return "cisco_ios_xr" }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetYangName() string { return "frr-type" }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) SetParent(parent types.Entity) { frrType.parent = parent }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetParent() types.Entity { return frrType.parent }

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetParentYangName() string { return "frr-types" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes
// Remote LFA Enable
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable remote lfa for a particular level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType.
    FrrRemoteLfaType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetFilter() yfilter.YFilter { return frrRemoteLfaTypes.YFilter }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) SetFilter(yf yfilter.YFilter) { frrRemoteLfaTypes.YFilter = yf }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetGoName(yname string) string {
    if yname == "frr-remote-lfa-type" { return "FrrRemoteLfaType" }
    return ""
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetSegmentPath() string {
    return "frr-remote-lfa-types"
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-remote-lfa-type" {
        for _, c := range frrRemoteLfaTypes.FrrRemoteLfaType {
            if frrRemoteLfaTypes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType{}
        frrRemoteLfaTypes.FrrRemoteLfaType = append(frrRemoteLfaTypes.FrrRemoteLfaType, child)
        return &frrRemoteLfaTypes.FrrRemoteLfaType[len(frrRemoteLfaTypes.FrrRemoteLfaType)-1]
    }
    return nil
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrRemoteLfaTypes.FrrRemoteLfaType {
        children[frrRemoteLfaTypes.FrrRemoteLfaType[i].GetSegmentPath()] = &frrRemoteLfaTypes.FrrRemoteLfaType[i]
    }
    return children
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetYangName() string { return "frr-remote-lfa-types" }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) SetParent(parent types.Entity) { frrRemoteLfaTypes.parent = parent }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetParent() types.Entity { return frrRemoteLfaTypes.parent }

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType
// Enable remote lfa for a particular
// level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Remote LFA Type. The type is IsisRemoteLfa. This attribute is mandatory.
    Type interface{}
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetFilter() yfilter.YFilter { return frrRemoteLfaType.YFilter }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) SetFilter(yf yfilter.YFilter) { frrRemoteLfaType.YFilter = yf }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "type" { return "Type" }
    return ""
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetSegmentPath() string {
    return "frr-remote-lfa-type" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaType.Level) + "']"
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrRemoteLfaType.Level
    leafs["type"] = frrRemoteLfaType.Type
    return leafs
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetBundleName() string { return "cisco_ios_xr" }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetYangName() string { return "frr-remote-lfa-type" }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) SetParent(parent types.Entity) { frrRemoteLfaType.parent = parent }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetParent() types.Entity { return frrRemoteLfaType.parent }

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetParentYangName() string { return "frr-remote-lfa-types" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults
// Interface FRR Default tiebreaker
// configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure default tiebreaker. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault.
    InterfaceFrrTiebreakerDefault []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetFilter() yfilter.YFilter { return interfaceFrrTiebreakerDefaults.YFilter }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) SetFilter(yf yfilter.YFilter) { interfaceFrrTiebreakerDefaults.YFilter = yf }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetGoName(yname string) string {
    if yname == "interface-frr-tiebreaker-default" { return "InterfaceFrrTiebreakerDefault" }
    return ""
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetSegmentPath() string {
    return "interface-frr-tiebreaker-defaults"
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-frr-tiebreaker-default" {
        for _, c := range interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault {
            if interfaceFrrTiebreakerDefaults.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault{}
        interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault = append(interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault, child)
        return &interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[len(interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault)-1]
    }
    return nil
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault {
        children[interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[i].GetSegmentPath()] = &interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[i]
    }
    return children
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetYangName() string { return "interface-frr-tiebreaker-defaults" }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) SetParent(parent types.Entity) { interfaceFrrTiebreakerDefaults.parent = parent }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetParent() types.Entity { return interfaceFrrTiebreakerDefaults.parent }

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault
// Configure default tiebreaker
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetFilter() yfilter.YFilter { return interfaceFrrTiebreakerDefault.YFilter }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) SetFilter(yf yfilter.YFilter) { interfaceFrrTiebreakerDefault.YFilter = yf }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    return ""
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetSegmentPath() string {
    return "interface-frr-tiebreaker-default" + "[level='" + fmt.Sprintf("%v", interfaceFrrTiebreakerDefault.Level) + "']"
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = interfaceFrrTiebreakerDefault.Level
    return leafs
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetYangName() string { return "interface-frr-tiebreaker-default" }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) SetParent(parent types.Entity) { interfaceFrrTiebreakerDefault.parent = parent }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetParent() types.Entity { return interfaceFrrTiebreakerDefault.parent }

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetParentYangName() string { return "interface-frr-tiebreaker-defaults" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes
// TI LFA Enable
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable TI lfa for a particular level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType.
    FrrtilfaType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetFilter() yfilter.YFilter { return frrtilfaTypes.YFilter }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) SetFilter(yf yfilter.YFilter) { frrtilfaTypes.YFilter = yf }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetGoName(yname string) string {
    if yname == "frrtilfa-type" { return "FrrtilfaType" }
    return ""
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetSegmentPath() string {
    return "frrtilfa-types"
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frrtilfa-type" {
        for _, c := range frrtilfaTypes.FrrtilfaType {
            if frrtilfaTypes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType{}
        frrtilfaTypes.FrrtilfaType = append(frrtilfaTypes.FrrtilfaType, child)
        return &frrtilfaTypes.FrrtilfaType[len(frrtilfaTypes.FrrtilfaType)-1]
    }
    return nil
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrtilfaTypes.FrrtilfaType {
        children[frrtilfaTypes.FrrtilfaType[i].GetSegmentPath()] = &frrtilfaTypes.FrrtilfaType[i]
    }
    return children
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetBundleName() string { return "cisco_ios_xr" }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetYangName() string { return "frrtilfa-types" }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) SetParent(parent types.Entity) { frrtilfaTypes.parent = parent }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetParent() types.Entity { return frrtilfaTypes.parent }

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType
// Enable TI lfa for a particular level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetFilter() yfilter.YFilter { return frrtilfaType.YFilter }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) SetFilter(yf yfilter.YFilter) { frrtilfaType.YFilter = yf }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    return ""
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetSegmentPath() string {
    return "frrtilfa-type" + "[level='" + fmt.Sprintf("%v", frrtilfaType.Level) + "']"
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = frrtilfaType.Level
    return leafs
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetBundleName() string { return "cisco_ios_xr" }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetYangName() string { return "frrtilfa-type" }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) SetParent(parent types.Entity) { frrtilfaType.parent = parent }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetParent() types.Entity { return frrtilfaType.parent }

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetParentYangName() string { return "frrtilfa-types" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces
// FRR exclusion configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Exclude an interface from computation. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface.
    FrrExcludeInterface []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetFilter() yfilter.YFilter { return frrExcludeInterfaces.YFilter }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) SetFilter(yf yfilter.YFilter) { frrExcludeInterfaces.YFilter = yf }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetGoName(yname string) string {
    if yname == "frr-exclude-interface" { return "FrrExcludeInterface" }
    return ""
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetSegmentPath() string {
    return "frr-exclude-interfaces"
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frr-exclude-interface" {
        for _, c := range frrExcludeInterfaces.FrrExcludeInterface {
            if frrExcludeInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface{}
        frrExcludeInterfaces.FrrExcludeInterface = append(frrExcludeInterfaces.FrrExcludeInterface, child)
        return &frrExcludeInterfaces.FrrExcludeInterface[len(frrExcludeInterfaces.FrrExcludeInterface)-1]
    }
    return nil
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range frrExcludeInterfaces.FrrExcludeInterface {
        children[frrExcludeInterfaces.FrrExcludeInterface[i].GetSegmentPath()] = &frrExcludeInterfaces.FrrExcludeInterface[i]
    }
    return children
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetYangName() string { return "frr-exclude-interfaces" }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) SetParent(parent types.Entity) { frrExcludeInterfaces.parent = parent }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetParent() types.Entity { return frrExcludeInterfaces.parent }

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface
// Exclude an interface from computation
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Level. The type is interface{} with range: 0..2. This attribute is
    // mandatory.
    Level interface{}
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetFilter() yfilter.YFilter { return frrExcludeInterface.YFilter }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) SetFilter(yf yfilter.YFilter) { frrExcludeInterface.YFilter = yf }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "frr-type" { return "FrrType" }
    if yname == "level" { return "Level" }
    return ""
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetSegmentPath() string {
    return "frr-exclude-interface" + "[interface-name='" + fmt.Sprintf("%v", frrExcludeInterface.InterfaceName) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrExcludeInterface.FrrType) + "']"
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = frrExcludeInterface.InterfaceName
    leafs["frr-type"] = frrExcludeInterface.FrrType
    leafs["level"] = frrExcludeInterface.Level
    return leafs
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetBundleName() string { return "cisco_ios_xr" }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetYangName() string { return "frr-exclude-interface" }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) SetParent(parent types.Entity) { frrExcludeInterface.parent = parent }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetParent() types.Entity { return frrExcludeInterface.parent }

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetParentYangName() string { return "frr-exclude-interfaces" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers
// Interface FRR tiebreakers configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure tiebreaker for multiple backups. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker.
    InterfaceFrrTiebreaker []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetFilter() yfilter.YFilter { return interfaceFrrTiebreakers.YFilter }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) SetFilter(yf yfilter.YFilter) { interfaceFrrTiebreakers.YFilter = yf }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetGoName(yname string) string {
    if yname == "interface-frr-tiebreaker" { return "InterfaceFrrTiebreaker" }
    return ""
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetSegmentPath() string {
    return "interface-frr-tiebreakers"
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-frr-tiebreaker" {
        for _, c := range interfaceFrrTiebreakers.InterfaceFrrTiebreaker {
            if interfaceFrrTiebreakers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker{}
        interfaceFrrTiebreakers.InterfaceFrrTiebreaker = append(interfaceFrrTiebreakers.InterfaceFrrTiebreaker, child)
        return &interfaceFrrTiebreakers.InterfaceFrrTiebreaker[len(interfaceFrrTiebreakers.InterfaceFrrTiebreaker)-1]
    }
    return nil
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceFrrTiebreakers.InterfaceFrrTiebreaker {
        children[interfaceFrrTiebreakers.InterfaceFrrTiebreaker[i].GetSegmentPath()] = &interfaceFrrTiebreakers.InterfaceFrrTiebreaker[i]
    }
    return children
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetYangName() string { return "interface-frr-tiebreakers" }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) SetParent(parent types.Entity) { interfaceFrrTiebreakers.parent = parent }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetParent() types.Entity { return interfaceFrrTiebreakers.parent }

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetParentYangName() string { return "interface-frr-table" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker
// Configure tiebreaker for multiple
// backups
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Tiebreaker for which configuration applies. The
    // type is IsisInterfaceFrrTiebreaker.
    Tiebreaker interface{}

    // Preference order among tiebreakers. The type is interface{} with range:
    // 1..255. This attribute is mandatory.
    Index interface{}
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetFilter() yfilter.YFilter { return interfaceFrrTiebreaker.YFilter }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) SetFilter(yf yfilter.YFilter) { interfaceFrrTiebreaker.YFilter = yf }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "tiebreaker" { return "Tiebreaker" }
    if yname == "index" { return "Index" }
    return ""
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetSegmentPath() string {
    return "interface-frr-tiebreaker" + "[level='" + fmt.Sprintf("%v", interfaceFrrTiebreaker.Level) + "']" + "[tiebreaker='" + fmt.Sprintf("%v", interfaceFrrTiebreaker.Tiebreaker) + "']"
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = interfaceFrrTiebreaker.Level
    leafs["tiebreaker"] = interfaceFrrTiebreaker.Tiebreaker
    leafs["index"] = interfaceFrrTiebreaker.Index
    return leafs
}

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetYangName() string { return "interface-frr-tiebreaker" }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) SetParent(parent types.Entity) { interfaceFrrTiebreaker.parent = parent }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetParent() types.Entity { return interfaceFrrTiebreaker.parent }

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetParentYangName() string { return "interface-frr-tiebreakers" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp
// MPLS LDP configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable MPLS LDP Synchronization for an IS-IS level. The type is interface{}
    // with range: 0..2. The default value is 0.
    SyncLevel interface{}
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetFilter() yfilter.YFilter { return mplsLdp.YFilter }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) SetFilter(yf yfilter.YFilter) { mplsLdp.YFilter = yf }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetGoName(yname string) string {
    if yname == "sync-level" { return "SyncLevel" }
    return ""
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetSegmentPath() string {
    return "mpls-ldp"
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sync-level"] = mplsLdp.SyncLevel
    return leafs
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetBundleName() string { return "cisco_ios_xr" }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetYangName() string { return "mpls-ldp" }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) SetParent(parent types.Entity) { mplsLdp.parent = parent }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetParent() types.Entity { return mplsLdp.parent }

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid
// Assign prefix SSPF SID to an interface,
// ISISPHPFlag will be rejected if set to
// disable, ISISEXPLICITNULLFlag will
// override the value of ISISPHPFlag
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type for the interface. The type is Isissid1. This attribute is
    // mandatory.
    Type interface{}

    // SID value for the interface. The type is interface{} with range:
    // 0..1048575. This attribute is mandatory.
    Value interface{}

    // Enable/Disable Penultimate Hop Popping. The type is IsisphpFlag. This
    // attribute is mandatory.
    Php interface{}

    // Enable/Disable Explicit-NULL flag. The type is IsisexplicitNullFlag. This
    // attribute is mandatory.
    ExplicitNull interface{}

    // Clear N-flag for the prefix-SID. The type is NflagClear. This attribute is
    // mandatory.
    NflagClear interface{}
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetFilter() yfilter.YFilter { return prefixSspfsid.YFilter }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) SetFilter(yf yfilter.YFilter) { prefixSspfsid.YFilter = yf }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "value" { return "Value" }
    if yname == "php" { return "Php" }
    if yname == "explicit-null" { return "ExplicitNull" }
    if yname == "nflag-clear" { return "NflagClear" }
    return ""
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetSegmentPath() string {
    return "prefix-sspfsid"
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = prefixSspfsid.Type
    leafs["value"] = prefixSspfsid.Value
    leafs["php"] = prefixSspfsid.Php
    leafs["explicit-null"] = prefixSspfsid.ExplicitNull
    leafs["nflag-clear"] = prefixSspfsid.NflagClear
    return leafs
}

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetYangName() string { return "prefix-sspfsid" }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) SetParent(parent types.Entity) { prefixSspfsid.parent = parent }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetParent() types.Entity { return prefixSspfsid.parent }

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics
// AutoMetric configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoMetric Proactive-Protect configuration. Legal value depends on the
    // metric-style specified for the topology. If the metric-style defined is
    // narrow, then only a value between <1-63> is allowed and if the metric-style
    // is defined as wide, then a value between <1-16777214> is allowed as the
    // auto-metric value. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric.
    AutoMetric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetFilter() yfilter.YFilter { return autoMetrics.YFilter }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) SetFilter(yf yfilter.YFilter) { autoMetrics.YFilter = yf }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetGoName(yname string) string {
    if yname == "auto-metric" { return "AutoMetric" }
    return ""
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetSegmentPath() string {
    return "auto-metrics"
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto-metric" {
        for _, c := range autoMetrics.AutoMetric {
            if autoMetrics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric{}
        autoMetrics.AutoMetric = append(autoMetrics.AutoMetric, child)
        return &autoMetrics.AutoMetric[len(autoMetrics.AutoMetric)-1]
    }
    return nil
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range autoMetrics.AutoMetric {
        children[autoMetrics.AutoMetric[i].GetSegmentPath()] = &autoMetrics.AutoMetric[i]
    }
    return children
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetBundleName() string { return "cisco_ios_xr" }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetYangName() string { return "auto-metrics" }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) SetParent(parent types.Entity) { autoMetrics.parent = parent }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetParent() types.Entity { return autoMetrics.parent }

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric
// AutoMetric Proactive-Protect
// configuration. Legal value depends on
// the metric-style specified for the
// topology. If the metric-style defined is
// narrow, then only a value between <1-63>
// is allowed and if the metric-style is
// defined as wide, then a value between
// <1-16777214> is allowed as the
// auto-metric value.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed auto metric:<1-63> for narrow ,<1-16777214> for wide. The type is
    // interface{} with range: 1..16777214. This attribute is mandatory.
    ProactiveProtect interface{}
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetFilter() yfilter.YFilter { return autoMetric.YFilter }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) SetFilter(yf yfilter.YFilter) { autoMetric.YFilter = yf }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "proactive-protect" { return "ProactiveProtect" }
    return ""
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetSegmentPath() string {
    return "auto-metric" + "[level='" + fmt.Sprintf("%v", autoMetric.Level) + "']"
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = autoMetric.Level
    leafs["proactive-protect"] = autoMetric.ProactiveProtect
    return leafs
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetBundleName() string { return "cisco_ios_xr" }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetYangName() string { return "auto-metric" }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) SetParent(parent types.Entity) { autoMetric.parent = parent }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetParent() types.Entity { return autoMetric.parent }

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetParentYangName() string { return "auto-metrics" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags
// admin-tag configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Admin tag for advertised interface connected routes. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag.
    AdminTag []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetFilter() yfilter.YFilter { return adminTags.YFilter }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) SetFilter(yf yfilter.YFilter) { adminTags.YFilter = yf }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetGoName(yname string) string {
    if yname == "admin-tag" { return "AdminTag" }
    return ""
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetSegmentPath() string {
    return "admin-tags"
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "admin-tag" {
        for _, c := range adminTags.AdminTag {
            if adminTags.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag{}
        adminTags.AdminTag = append(adminTags.AdminTag, child)
        return &adminTags.AdminTag[len(adminTags.AdminTag)-1]
    }
    return nil
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range adminTags.AdminTag {
        children[adminTags.AdminTag[i].GetSegmentPath()] = &adminTags.AdminTag[i]
    }
    return children
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetBundleName() string { return "cisco_ios_xr" }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetYangName() string { return "admin-tags" }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) SetParent(parent types.Entity) { adminTags.parent = parent }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetParent() types.Entity { return adminTags.parent }

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag
// Admin tag for advertised interface
// connected routes
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Tag to associate with connected routes. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    AdminTag interface{}
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetFilter() yfilter.YFilter { return adminTag.YFilter }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) SetFilter(yf yfilter.YFilter) { adminTag.YFilter = yf }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "admin-tag" { return "AdminTag" }
    return ""
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetSegmentPath() string {
    return "admin-tag" + "[level='" + fmt.Sprintf("%v", adminTag.Level) + "']"
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = adminTag.Level
    leafs["admin-tag"] = adminTag.AdminTag
    return leafs
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetBundleName() string { return "cisco_ios_xr" }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetYangName() string { return "admin-tag" }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) SetParent(parent types.Entity) { adminTag.parent = parent }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetParent() types.Entity { return adminTag.parent }

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetParentYangName() string { return "admin-tags" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup
// Provide link group name and level
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Group. The type is string with length: 1..40. This attribute is
    // mandatory.
    LinkGroup interface{}

    // Level in which link group will be effective. The type is interface{} with
    // range: 0..2. The default value is 0.
    Level interface{}
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetFilter() yfilter.YFilter { return interfaceLinkGroup.YFilter }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) SetFilter(yf yfilter.YFilter) { interfaceLinkGroup.YFilter = yf }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetGoName(yname string) string {
    if yname == "link-group" { return "LinkGroup" }
    if yname == "level" { return "Level" }
    return ""
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetSegmentPath() string {
    return "interface-link-group"
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-group"] = interfaceLinkGroup.LinkGroup
    leafs["level"] = interfaceLinkGroup.Level
    return leafs
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetYangName() string { return "interface-link-group" }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) SetParent(parent types.Entity) { interfaceLinkGroup.parent = parent }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetParent() types.Entity { return interfaceLinkGroup.parent }

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids
// Manual Adjacecy SID configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Assign adjancency SID to an interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid.
    ManualAdjSid []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetFilter() yfilter.YFilter { return manualAdjSids.YFilter }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) SetFilter(yf yfilter.YFilter) { manualAdjSids.YFilter = yf }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetGoName(yname string) string {
    if yname == "manual-adj-sid" { return "ManualAdjSid" }
    return ""
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetSegmentPath() string {
    return "manual-adj-sids"
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "manual-adj-sid" {
        for _, c := range manualAdjSids.ManualAdjSid {
            if manualAdjSids.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid{}
        manualAdjSids.ManualAdjSid = append(manualAdjSids.ManualAdjSid, child)
        return &manualAdjSids.ManualAdjSid[len(manualAdjSids.ManualAdjSid)-1]
    }
    return nil
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range manualAdjSids.ManualAdjSid {
        children[manualAdjSids.ManualAdjSid[i].GetSegmentPath()] = &manualAdjSids.ManualAdjSid[i]
    }
    return children
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetBundleName() string { return "cisco_ios_xr" }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetYangName() string { return "manual-adj-sids" }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) SetParent(parent types.Entity) { manualAdjSids.parent = parent }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetParent() types.Entity { return manualAdjSids.parent }

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid
// Assign adjancency SID to an interface
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Sid type aboslute or index. The type is Isissid1.
    SidType interface{}

    // This attribute is a key. Sid value. The type is interface{} with range:
    // 0..1048575.
    Sid interface{}

    // Enable/Disable SID protection. The type is IsissidProtected. This attribute
    // is mandatory.
    Protected interface{}
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetFilter() yfilter.YFilter { return manualAdjSid.YFilter }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) SetFilter(yf yfilter.YFilter) { manualAdjSid.YFilter = yf }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "sid-type" { return "SidType" }
    if yname == "sid" { return "Sid" }
    if yname == "protected" { return "Protected" }
    return ""
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetSegmentPath() string {
    return "manual-adj-sid" + "[level='" + fmt.Sprintf("%v", manualAdjSid.Level) + "']" + "[sid-type='" + fmt.Sprintf("%v", manualAdjSid.SidType) + "']" + "[sid='" + fmt.Sprintf("%v", manualAdjSid.Sid) + "']"
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = manualAdjSid.Level
    leafs["sid-type"] = manualAdjSid.SidType
    leafs["sid"] = manualAdjSid.Sid
    leafs["protected"] = manualAdjSid.Protected
    return leafs
}

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetBundleName() string { return "cisco_ios_xr" }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetYangName() string { return "manual-adj-sid" }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) SetParent(parent types.Entity) { manualAdjSid.parent = parent }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetParent() types.Entity { return manualAdjSid.parent }

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetParentYangName() string { return "manual-adj-sids" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics
// Metric configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Metric configuration. Legal value depends on the metric-style specified for
    // the topology. If the metric-style defined is narrow, then only a value
    // between <1-63> is allowed and if the metric-style is defined as wide, then
    // a value between <1-16777215> is allowed as the metric value.  All routers
    // exclude links with the maximum wide metric (16777215) from their SPF. The
    // type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric.
    Metric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetFilter() yfilter.YFilter { return metrics.YFilter }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) SetFilter(yf yfilter.YFilter) { metrics.YFilter = yf }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetGoName(yname string) string {
    if yname == "metric" { return "Metric" }
    return ""
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetSegmentPath() string {
    return "metrics"
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "metric" {
        for _, c := range metrics.Metric {
            if metrics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric{}
        metrics.Metric = append(metrics.Metric, child)
        return &metrics.Metric[len(metrics.Metric)-1]
    }
    return nil
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range metrics.Metric {
        children[metrics.Metric[i].GetSegmentPath()] = &metrics.Metric[i]
    }
    return children
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetBundleName() string { return "cisco_ios_xr" }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetYangName() string { return "metrics" }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) SetParent(parent types.Entity) { metrics.parent = parent }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetParent() types.Entity { return metrics.parent }

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric
// Metric configuration. Legal value depends on
// the metric-style specified for the topology. If
// the metric-style defined is narrow, then only a
// value between <1-63> is allowed and if the
// metric-style is defined as wide, then a value
// between <1-16777215> is allowed as the metric
// value.  All routers exclude links with the
// maximum wide metric (16777215) from their SPF
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed metric: <1-63> for narrow, <1-16777215> for wide. The type is one
    // of the following types: enumeration
    // Isis.Instances.Instance.Interfaces.Interface.InterfaceAfs.InterfaceAf.TopologyName.Metrics.Metric.Metric
    // This attribute is mandatory., or int with range: 1..16777215 This attribute
    // is mandatory..
    Metric interface{}
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetFilter() yfilter.YFilter { return metric.YFilter }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) SetFilter(yf yfilter.YFilter) { metric.YFilter = yf }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "metric" { return "Metric" }
    return ""
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetSegmentPath() string {
    return "metric" + "[level='" + fmt.Sprintf("%v", metric.Level) + "']"
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = metric.Level
    leafs["metric"] = metric.Metric
    return leafs
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetBundleName() string { return "cisco_ios_xr" }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetYangName() string { return "metric" }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) SetParent(parent types.Entity) { metric.parent = parent }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetParent() types.Entity { return metric.parent }

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetParentYangName() string { return "metrics" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric_Metric represents <1-16777215> for wide
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric_Metric string

const (
    // Maximum wide metric.  All routers will
    // exclude this link from their SPF
    Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric_Metric_maximum Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric_Metric = "maximum"
)

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights
// Weight configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Weight configuration under interface for load balancing. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight.
    Weight []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetFilter() yfilter.YFilter { return weights.YFilter }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) SetFilter(yf yfilter.YFilter) { weights.YFilter = yf }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetGoName(yname string) string {
    if yname == "weight" { return "Weight" }
    return ""
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetSegmentPath() string {
    return "weights"
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "weight" {
        for _, c := range weights.Weight {
            if weights.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight{}
        weights.Weight = append(weights.Weight, child)
        return &weights.Weight[len(weights.Weight)-1]
    }
    return nil
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range weights.Weight {
        children[weights.Weight[i].GetSegmentPath()] = &weights.Weight[i]
    }
    return children
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetBundleName() string { return "cisco_ios_xr" }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetYangName() string { return "weights" }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) SetParent(parent types.Entity) { weights.parent = parent }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetParent() types.Entity { return weights.parent }

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetParentYangName() string { return "topology-name" }

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight
// Weight configuration under interface for load
// balancing
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Weight to be configured under interface for Load Balancing. Allowed weight:
    // <1-16777215>. The type is interface{} with range: 1..16777214. This
    // attribute is mandatory.
    Weight interface{}
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetFilter() yfilter.YFilter { return weight.YFilter }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) SetFilter(yf yfilter.YFilter) { weight.YFilter = yf }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "weight" { return "Weight" }
    return ""
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetSegmentPath() string {
    return "weight" + "[level='" + fmt.Sprintf("%v", weight.Level) + "']"
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = weight.Level
    leafs["weight"] = weight.Weight
    return leafs
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetBundleName() string { return "cisco_ios_xr" }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetYangName() string { return "weight" }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) SetParent(parent types.Entity) { weight.parent = parent }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetParent() types.Entity { return weight.parent }

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetParentYangName() string { return "weights" }

// Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals
// CSNP-interval configuration
type Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CSNP-interval configuration. No fixed default value as this depends on the
    // media type of the interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval.
    CsnpInterval []Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval
}

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetFilter() yfilter.YFilter { return csnpIntervals.YFilter }

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) SetFilter(yf yfilter.YFilter) { csnpIntervals.YFilter = yf }

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetGoName(yname string) string {
    if yname == "csnp-interval" { return "CsnpInterval" }
    return ""
}

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetSegmentPath() string {
    return "csnp-intervals"
}

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "csnp-interval" {
        for _, c := range csnpIntervals.CsnpInterval {
            if csnpIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval{}
        csnpIntervals.CsnpInterval = append(csnpIntervals.CsnpInterval, child)
        return &csnpIntervals.CsnpInterval[len(csnpIntervals.CsnpInterval)-1]
    }
    return nil
}

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range csnpIntervals.CsnpInterval {
        children[csnpIntervals.CsnpInterval[i].GetSegmentPath()] = &csnpIntervals.CsnpInterval[i]
    }
    return children
}

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetYangName() string { return "csnp-intervals" }

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) SetParent(parent types.Entity) { csnpIntervals.parent = parent }

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetParent() types.Entity { return csnpIntervals.parent }

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval
// CSNP-interval configuration. No fixed
// default value as this depends on the media
// type of the interface.
type Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 0..65535. This attribute is
    // mandatory. Units are second.
    Interval interface{}
}

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetFilter() yfilter.YFilter { return csnpInterval.YFilter }

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) SetFilter(yf yfilter.YFilter) { csnpInterval.YFilter = yf }

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetSegmentPath() string {
    return "csnp-interval" + "[level='" + fmt.Sprintf("%v", csnpInterval.Level) + "']"
}

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = csnpInterval.Level
    leafs["interval"] = csnpInterval.Interval
    return leafs
}

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetBundleName() string { return "cisco_ios_xr" }

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetYangName() string { return "csnp-interval" }

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) SetParent(parent types.Entity) { csnpInterval.parent = parent }

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetParent() types.Entity { return csnpInterval.parent }

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetParentYangName() string { return "csnp-intervals" }

// Isis_Instances_Instance_Interfaces_Interface_LspIntervals
// LSP-interval configuration
type Isis_Instances_Instance_Interfaces_Interface_LspIntervals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval between transmission of LSPs on interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval.
    LspInterval []Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval
}

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetFilter() yfilter.YFilter { return lspIntervals.YFilter }

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) SetFilter(yf yfilter.YFilter) { lspIntervals.YFilter = yf }

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetGoName(yname string) string {
    if yname == "lsp-interval" { return "LspInterval" }
    return ""
}

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetSegmentPath() string {
    return "lsp-intervals"
}

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-interval" {
        for _, c := range lspIntervals.LspInterval {
            if lspIntervals.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval{}
        lspIntervals.LspInterval = append(lspIntervals.LspInterval, child)
        return &lspIntervals.LspInterval[len(lspIntervals.LspInterval)-1]
    }
    return nil
}

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lspIntervals.LspInterval {
        children[lspIntervals.LspInterval[i].GetSegmentPath()] = &lspIntervals.LspInterval[i]
    }
    return children
}

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetBundleName() string { return "cisco_ios_xr" }

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetYangName() string { return "lsp-intervals" }

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) SetParent(parent types.Entity) { lspIntervals.parent = parent }

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetParent() types.Entity { return lspIntervals.parent }

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetParentYangName() string { return "interface" }

// Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval
// Interval between transmission of LSPs on
// interface.
type Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Milliseconds. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory. Units are millisecond.
    Interval interface{}
}

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetFilter() yfilter.YFilter { return lspInterval.YFilter }

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) SetFilter(yf yfilter.YFilter) { lspInterval.YFilter = yf }

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetSegmentPath() string {
    return "lsp-interval" + "[level='" + fmt.Sprintf("%v", lspInterval.Level) + "']"
}

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = lspInterval.Level
    leafs["interval"] = lspInterval.Interval
    return leafs
}

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetBundleName() string { return "cisco_ios_xr" }

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetYangName() string { return "lsp-interval" }

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) SetParent(parent types.Entity) { lspInterval.parent = parent }

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetParent() types.Entity { return lspInterval.parent }

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetParentYangName() string { return "lsp-intervals" }

// Isis_Instances_Instance_Interfaces_Interface_MeshGroup represents Mesh-group configuration
type Isis_Instances_Instance_Interfaces_Interface_MeshGroup string

const (
    // Blocked mesh group.  Changed LSPs are not
    // flooded over blocked interfaces
    Isis_Instances_Instance_Interfaces_Interface_MeshGroup_blocked Isis_Instances_Instance_Interfaces_Interface_MeshGroup = "blocked"
)

