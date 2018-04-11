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

// NflagClear represents Nflag clear
type NflagClear string

const (
    // Disable N-flag-clear
    NflagClear_disable NflagClear = "disable"

    // Enable N-flag-clear
    NflagClear_enable NflagClear = "enable"
)

// IsisAdvTypeInterLevel represents Isis adv type inter level
type IsisAdvTypeInterLevel string

const (
    // InterLevel
    IsisAdvTypeInterLevel_inter_level IsisAdvTypeInterLevel = "inter-level"
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

// Isisfrr represents Isisfrr
type Isisfrr string

const (
    // Prefix independent per-link computation
    Isisfrr_per_link Isisfrr = "per-link"

    // Prefix dependent computation
    Isisfrr_per_prefix Isisfrr = "per-prefix"
)

// IsisfrrLoadSharing represents Isisfrr load sharing
type IsisfrrLoadSharing string

const (
    // Disable load sharing of prefixes across
    // multiple backups
    IsisfrrLoadSharing_disable IsisfrrLoadSharing = "disable"
)

// IsisAuthenticationFailureMode represents Isis authentication failure mode
type IsisAuthenticationFailureMode string

const (
    // Drop non-authenticating PDUs
    IsisAuthenticationFailureMode_drop IsisAuthenticationFailureMode = "drop"

    // Accept non-authenticating PDUs
    IsisAuthenticationFailureMode_send_only IsisAuthenticationFailureMode = "send-only"
)

// IsisApplyWeight represents Isis apply weight
type IsisApplyWeight string

const (
    // Apply weight to ECMP prefixes
    IsisApplyWeight_ecmp_only IsisApplyWeight = "ecmp-only"

    // Apply weight to UCMP prefixes
    IsisApplyWeight_ucmp_only IsisApplyWeight = "ucmp-only"

    // Apply weight to ECMP prefixes
    IsisApplyWeight_ecmp_only_bandwidth IsisApplyWeight = "ecmp-only-bandwidth"
)

// IsisLabelPreference represents Isis label preference
type IsisLabelPreference string

const (
    // Label Distribution Protocol
    IsisLabelPreference_ldp IsisLabelPreference = "ldp"

    // Segment Routing
    IsisLabelPreference_segment_routing IsisLabelPreference = "segment-routing"
)

// IsisMetricStyleTransition represents Isis metric style transition
type IsisMetricStyleTransition string

const (
    // Disabled
    IsisMetricStyleTransition_disabled IsisMetricStyleTransition = "disabled"

    // Enabled
    IsisMetricStyleTransition_enabled IsisMetricStyleTransition = "enabled"
)

// Isissid1 represents Isissid1
type Isissid1 string

const (
    // SID as an index
    Isissid1_index Isissid1 = "index"

    // SID as an absolute label
    Isissid1_absolute Isissid1 = "absolute"
)

// IsisMetric represents Isis metric
type IsisMetric string

const (
    // Internal metric
    IsisMetric_internal IsisMetric = "internal"

    // External metric
    IsisMetric_external IsisMetric = "external"

    // RIB Internal metric
    IsisMetric_rib_internal IsisMetric = "rib-internal"

    // RIB External metric
    IsisMetric_rib_external IsisMetric = "rib-external"
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

// IsisHelloPadding represents Isis hello padding
type IsisHelloPadding string

const (
    // Never pad Hellos
    IsisHelloPadding_never IsisHelloPadding = "never"

    // Pad Hellos during adjacency formation only
    IsisHelloPadding_sometimes IsisHelloPadding = "sometimes"
)

// IsisSnpAuth represents Isis snp auth
type IsisSnpAuth string

const (
    // Authenticate SNP send only
    IsisSnpAuth_send_only IsisSnpAuth = "send-only"

    // Authenticate SNP send and recv
    IsisSnpAuth_full IsisSnpAuth = "full"
)

// IsisInterfaceAfState represents Isis interface af state
type IsisInterfaceAfState string

const (
    // Disable
    IsisInterfaceAfState_disable IsisInterfaceAfState = "disable"
)

// IsisexplicitNullFlag represents Isisexplicit null flag
type IsisexplicitNullFlag string

const (
    // Disable EXPLICITNULL
    IsisexplicitNullFlag_disable IsisexplicitNullFlag = "disable"

    // Enable EXPLICITNULL
    IsisexplicitNullFlag_enable IsisexplicitNullFlag = "enable"
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

// IsisispfState represents Isisispf state
type IsisispfState string

const (
    // Enabled
    IsisispfState_enabled IsisispfState = "enabled"
)

// IsissidProtected represents Isissid protected
type IsissidProtected string

const (
    // Not protected
    IsissidProtected_disable IsissidProtected = "disable"

    // Protected
    IsissidProtected_enable IsissidProtected = "enable"
)

// IsisAdvTypeExternal represents Isis adv type external
type IsisAdvTypeExternal string

const (
    // External
    IsisAdvTypeExternal_external IsisAdvTypeExternal = "external"
)

// IsisphpFlag represents Isisphp flag
type IsisphpFlag string

const (
    // Enable PHP
    IsisphpFlag_enable IsisphpFlag = "enable"

    // Disable PHP
    IsisphpFlag_disable IsisphpFlag = "disable"
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

// IsisRemoteLfa represents Isis remote lfa
type IsisRemoteLfa string

const (
    // No remote LFA option set
    IsisRemoteLfa_remote_lfa_none IsisRemoteLfa = "remote-lfa-none"

    // Construct remote LFA tunnel using MPLS LDP
    IsisRemoteLfa_remote_lfa_tunnel_ldp IsisRemoteLfa = "remote-lfa-tunnel-ldp"
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

// IsisAdjCheck represents Isis adj check
type IsisAdjCheck string

const (
    // Disabled
    IsisAdjCheck_disabled IsisAdjCheck = "disabled"
)

// IsisInterfaceFrrTiebreaker represents Isis interface frr tiebreaker
type IsisInterfaceFrrTiebreaker string

const (
    // Prefer node protecting backup path
    IsisInterfaceFrrTiebreaker_node_protecting IsisInterfaceFrrTiebreaker = "node-protecting"

    // Prefer SRLG disjoint backup path
    IsisInterfaceFrrTiebreaker_srlg_disjoint IsisInterfaceFrrTiebreaker = "srlg-disjoint"
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

// IsisNsfFlavor represents Isis nsf flavor
type IsisNsfFlavor string

const (
    // Cisco proprietary NSF
    IsisNsfFlavor_cisco_proprietary_nsf IsisNsfFlavor = "cisco-proprietary-nsf"

    // IETF standard NSF
    IsisNsfFlavor_ietf_standard_nsf IsisNsfFlavor = "ietf-standard-nsf"
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

// IsisMibAuthenticationTypeFailureBoolean represents Isis mib authentication type failure boolean
type IsisMibAuthenticationTypeFailureBoolean string

const (
    // Disable
    IsisMibAuthenticationTypeFailureBoolean_false IsisMibAuthenticationTypeFailureBoolean = "false"

    // Enable
    IsisMibAuthenticationTypeFailureBoolean_true IsisMibAuthenticationTypeFailureBoolean = "true"
)

// IsisMibMaxAreaAddressMismatchBoolean represents Isis mib max area address mismatch boolean
type IsisMibMaxAreaAddressMismatchBoolean string

const (
    // Disable
    IsisMibMaxAreaAddressMismatchBoolean_false IsisMibMaxAreaAddressMismatchBoolean = "false"

    // Enable
    IsisMibMaxAreaAddressMismatchBoolean_true IsisMibMaxAreaAddressMismatchBoolean = "true"
)

// IsisMibSequenceNumberSkipBoolean represents Isis mib sequence number skip boolean
type IsisMibSequenceNumberSkipBoolean string

const (
    // Disable
    IsisMibSequenceNumberSkipBoolean_false IsisMibSequenceNumberSkipBoolean = "false"

    // Enable
    IsisMibSequenceNumberSkipBoolean_true IsisMibSequenceNumberSkipBoolean = "true"
)

// IsisMibDatabaseOverFlowBoolean represents Isis mib database over flow boolean
type IsisMibDatabaseOverFlowBoolean string

const (
    // Disable
    IsisMibDatabaseOverFlowBoolean_false IsisMibDatabaseOverFlowBoolean = "false"

    // Enable
    IsisMibDatabaseOverFlowBoolean_true IsisMibDatabaseOverFlowBoolean = "true"
)

// IsisMibAllBoolean represents Isis mib all boolean
type IsisMibAllBoolean string

const (
    // Disable
    IsisMibAllBoolean_false IsisMibAllBoolean = "false"

    // Enable
    IsisMibAllBoolean_true IsisMibAllBoolean = "true"
)

// IsisMibLspTooLargeToPropagateBoolean represents Isis mib lsp too large to propagate boolean
type IsisMibLspTooLargeToPropagateBoolean string

const (
    // Disable
    IsisMibLspTooLargeToPropagateBoolean_false IsisMibLspTooLargeToPropagateBoolean = "false"

    // Enable
    IsisMibLspTooLargeToPropagateBoolean_true IsisMibLspTooLargeToPropagateBoolean = "true"
)

// IsisMibOwnLspPurgeBoolean represents Isis mib own lsp purge boolean
type IsisMibOwnLspPurgeBoolean string

const (
    // Disable
    IsisMibOwnLspPurgeBoolean_false IsisMibOwnLspPurgeBoolean = "false"

    // Enable
    IsisMibOwnLspPurgeBoolean_true IsisMibOwnLspPurgeBoolean = "true"
)

// IsisMibAdjacencyChangeBoolean represents Isis mib adjacency change boolean
type IsisMibAdjacencyChangeBoolean string

const (
    // Disable
    IsisMibAdjacencyChangeBoolean_false IsisMibAdjacencyChangeBoolean = "false"

    // Enable
    IsisMibAdjacencyChangeBoolean_true IsisMibAdjacencyChangeBoolean = "true"
)

// IsisMibProtocolsSupportedMismatchBoolean represents Isis mib protocols supported mismatch boolean
type IsisMibProtocolsSupportedMismatchBoolean string

const (
    // Disable
    IsisMibProtocolsSupportedMismatchBoolean_false IsisMibProtocolsSupportedMismatchBoolean = "false"

    // Enable
    IsisMibProtocolsSupportedMismatchBoolean_true IsisMibProtocolsSupportedMismatchBoolean = "true"
)

// IsisMibAttemptToExceedMaxSequenceBoolean represents Isis mib attempt to exceed max sequence boolean
type IsisMibAttemptToExceedMaxSequenceBoolean string

const (
    // Disable
    IsisMibAttemptToExceedMaxSequenceBoolean_false IsisMibAttemptToExceedMaxSequenceBoolean = "false"

    // Enable
    IsisMibAttemptToExceedMaxSequenceBoolean_true IsisMibAttemptToExceedMaxSequenceBoolean = "true"
)

// IsisMibIdLengthMismatchBoolean represents Isis mib id length mismatch boolean
type IsisMibIdLengthMismatchBoolean string

const (
    // Disable
    IsisMibIdLengthMismatchBoolean_false IsisMibIdLengthMismatchBoolean = "false"

    // Enable
    IsisMibIdLengthMismatchBoolean_true IsisMibIdLengthMismatchBoolean = "true"
)

// IsisMibOriginatedLspBufferSizeMismatchBoolean represents boolean
type IsisMibOriginatedLspBufferSizeMismatchBoolean string

const (
    // Disable
    IsisMibOriginatedLspBufferSizeMismatchBoolean_false IsisMibOriginatedLspBufferSizeMismatchBoolean = "false"

    // Enable
    IsisMibOriginatedLspBufferSizeMismatchBoolean_true IsisMibOriginatedLspBufferSizeMismatchBoolean = "true"
)

// IsisMibAreaMismatchBoolean represents Isis mib area mismatch boolean
type IsisMibAreaMismatchBoolean string

const (
    // Disable
    IsisMibAreaMismatchBoolean_false IsisMibAreaMismatchBoolean = "false"

    // Enable
    IsisMibAreaMismatchBoolean_true IsisMibAreaMismatchBoolean = "true"
)

// IsisMibCorruptedLspDetectedBoolean represents Isis mib corrupted lsp detected boolean
type IsisMibCorruptedLspDetectedBoolean string

const (
    // Disable
    IsisMibCorruptedLspDetectedBoolean_false IsisMibCorruptedLspDetectedBoolean = "false"

    // Enable
    IsisMibCorruptedLspDetectedBoolean_true IsisMibCorruptedLspDetectedBoolean = "true"
)

// IsisMibLspErrorDetectedBoolean represents Isis mib lsp error detected boolean
type IsisMibLspErrorDetectedBoolean string

const (
    // Disable
    IsisMibLspErrorDetectedBoolean_false IsisMibLspErrorDetectedBoolean = "false"

    // Enable
    IsisMibLspErrorDetectedBoolean_true IsisMibLspErrorDetectedBoolean = "true"
)

// IsisMibRejectedAdjacencyBoolean represents Isis mib rejected adjacency boolean
type IsisMibRejectedAdjacencyBoolean string

const (
    // Disable
    IsisMibRejectedAdjacencyBoolean_false IsisMibRejectedAdjacencyBoolean = "false"

    // Enable
    IsisMibRejectedAdjacencyBoolean_true IsisMibRejectedAdjacencyBoolean = "true"
)

// IsisMibVersionSkewBoolean represents Isis mib version skew boolean
type IsisMibVersionSkewBoolean string

const (
    // Disable
    IsisMibVersionSkewBoolean_false IsisMibVersionSkewBoolean = "false"

    // Enable
    IsisMibVersionSkewBoolean_true IsisMibVersionSkewBoolean = "true"
)

// IsisMibAuthenticationFailureBoolean represents Isis mib authentication failure boolean
type IsisMibAuthenticationFailureBoolean string

const (
    // Disable
    IsisMibAuthenticationFailureBoolean_false IsisMibAuthenticationFailureBoolean = "false"

    // Enable
    IsisMibAuthenticationFailureBoolean_true IsisMibAuthenticationFailureBoolean = "true"
)

// Isis
// IS-IS configuration for all instances
type Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IS-IS instance configuration.
    Instances Isis_Instances
}

func (isis *Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "Cisco-IOS-XR-clns-isis-cfg"
    isis.EntityData.SegmentPath = "Cisco-IOS-XR-clns-isis-cfg:isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Children["instances"] = types.YChild{"Instances", &isis.Instances}
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(isis.EntityData)
}

// Isis_Instances
// IS-IS instance configuration
type Isis_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a single IS-IS instance. The type is slice of
    // Isis_Instances_Instance.
    Instance []Isis_Instances_Instance
}

func (instances *Isis_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xr"
    instances.EntityData.ParentYangName = "isis"
    instances.EntityData.SegmentPath = "instances"
    instances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instances.EntityData.Children = make(map[string]types.YChild)
    instances.EntityData.Children["instance"] = types.YChild{"Instance", nil}
    for i := range instances.Instance {
        instances.EntityData.Children[types.GetSegmentPath(&instances.Instance[i])] = types.YChild{"Instance", &instances.Instance[i]}
    }
    instances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(instances.EntityData)
}

// Isis_Instances_Instance
// Configuration for a single IS-IS instance
type Isis_Instances_Instance struct {
    EntityData types.CommonEntityData
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

    // VRF context for ISIS process. The type is string with length: 1..32.
    VrfContext interface{}

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

func (instance *Isis_Instances_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instances"
    instance.EntityData.SegmentPath = "instance" + "[instance-name='" + fmt.Sprintf("%v", instance.InstanceName) + "']"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Children["srgb"] = types.YChild{"Srgb", &instance.Srgb}
    instance.EntityData.Children["lsp-generation-intervals"] = types.YChild{"LspGenerationIntervals", &instance.LspGenerationIntervals}
    instance.EntityData.Children["lsp-arrival-times"] = types.YChild{"LspArrivalTimes", &instance.LspArrivalTimes}
    instance.EntityData.Children["trace-buffer-size"] = types.YChild{"TraceBufferSize", &instance.TraceBufferSize}
    instance.EntityData.Children["max-link-metrics"] = types.YChild{"MaxLinkMetrics", &instance.MaxLinkMetrics}
    instance.EntityData.Children["adjacency-stagger"] = types.YChild{"AdjacencyStagger", &instance.AdjacencyStagger}
    instance.EntityData.Children["afs"] = types.YChild{"Afs", &instance.Afs}
    instance.EntityData.Children["lsp-refresh-intervals"] = types.YChild{"LspRefreshIntervals", &instance.LspRefreshIntervals}
    instance.EntityData.Children["distribute"] = types.YChild{"Distribute", &instance.Distribute}
    instance.EntityData.Children["lsp-accept-passwords"] = types.YChild{"LspAcceptPasswords", &instance.LspAcceptPasswords}
    instance.EntityData.Children["lsp-mtus"] = types.YChild{"LspMtus", &instance.LspMtus}
    instance.EntityData.Children["nsf"] = types.YChild{"Nsf", &instance.Nsf}
    instance.EntityData.Children["link-groups"] = types.YChild{"LinkGroups", &instance.LinkGroups}
    instance.EntityData.Children["lsp-check-intervals"] = types.YChild{"LspCheckIntervals", &instance.LspCheckIntervals}
    instance.EntityData.Children["lsp-passwords"] = types.YChild{"LspPasswords", &instance.LspPasswords}
    instance.EntityData.Children["nets"] = types.YChild{"Nets", &instance.Nets}
    instance.EntityData.Children["lsp-lifetimes"] = types.YChild{"LspLifetimes", &instance.LspLifetimes}
    instance.EntityData.Children["overload-bits"] = types.YChild{"OverloadBits", &instance.OverloadBits}
    instance.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &instance.Interfaces}
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["instance-name"] = types.YLeaf{"InstanceName", instance.InstanceName}
    instance.EntityData.Leafs["running"] = types.YLeaf{"Running", instance.Running}
    instance.EntityData.Leafs["log-adjacency-changes"] = types.YLeaf{"LogAdjacencyChanges", instance.LogAdjacencyChanges}
    instance.EntityData.Leafs["ignore-lsp-errors"] = types.YLeaf{"IgnoreLspErrors", instance.IgnoreLspErrors}
    instance.EntityData.Leafs["is-type"] = types.YLeaf{"IsType", instance.IsType}
    instance.EntityData.Leafs["tracing-mode"] = types.YLeaf{"TracingMode", instance.TracingMode}
    instance.EntityData.Leafs["vrf-context"] = types.YLeaf{"VrfContext", instance.VrfContext}
    instance.EntityData.Leafs["instance-id"] = types.YLeaf{"InstanceId", instance.InstanceId}
    instance.EntityData.Leafs["dynamic-host-name"] = types.YLeaf{"DynamicHostName", instance.DynamicHostName}
    instance.EntityData.Leafs["nsr"] = types.YLeaf{"Nsr", instance.Nsr}
    instance.EntityData.Leafs["log-pdu-drops"] = types.YLeaf{"LogPduDrops", instance.LogPduDrops}
    return &(instance.EntityData)
}

// Isis_Instances_Instance_Srgb
// Segment Routing Global Block configuration
// This type is a presence type.
type Isis_Instances_Instance_Srgb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The lower bound of the SRGB. The type is interface{} with range:
    // 16000..1048574. This attribute is mandatory.
    LowerBound interface{}

    // The upper bound of the SRGB. The type is interface{} with range:
    // 16001..1048575. This attribute is mandatory.
    UpperBound interface{}
}

func (srgb *Isis_Instances_Instance_Srgb) GetEntityData() *types.CommonEntityData {
    srgb.EntityData.YFilter = srgb.YFilter
    srgb.EntityData.YangName = "srgb"
    srgb.EntityData.BundleName = "cisco_ios_xr"
    srgb.EntityData.ParentYangName = "instance"
    srgb.EntityData.SegmentPath = "srgb"
    srgb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgb.EntityData.Children = make(map[string]types.YChild)
    srgb.EntityData.Leafs = make(map[string]types.YLeaf)
    srgb.EntityData.Leafs["lower-bound"] = types.YLeaf{"LowerBound", srgb.LowerBound}
    srgb.EntityData.Leafs["upper-bound"] = types.YLeaf{"UpperBound", srgb.UpperBound}
    return &(srgb.EntityData)
}

// Isis_Instances_Instance_LspGenerationIntervals
// LSP generation-interval configuration
type Isis_Instances_Instance_LspGenerationIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP generation scheduling parameters. The type is slice of
    // Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval.
    LspGenerationInterval []Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval
}

func (lspGenerationIntervals *Isis_Instances_Instance_LspGenerationIntervals) GetEntityData() *types.CommonEntityData {
    lspGenerationIntervals.EntityData.YFilter = lspGenerationIntervals.YFilter
    lspGenerationIntervals.EntityData.YangName = "lsp-generation-intervals"
    lspGenerationIntervals.EntityData.BundleName = "cisco_ios_xr"
    lspGenerationIntervals.EntityData.ParentYangName = "instance"
    lspGenerationIntervals.EntityData.SegmentPath = "lsp-generation-intervals"
    lspGenerationIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspGenerationIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspGenerationIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspGenerationIntervals.EntityData.Children = make(map[string]types.YChild)
    lspGenerationIntervals.EntityData.Children["lsp-generation-interval"] = types.YChild{"LspGenerationInterval", nil}
    for i := range lspGenerationIntervals.LspGenerationInterval {
        lspGenerationIntervals.EntityData.Children[types.GetSegmentPath(&lspGenerationIntervals.LspGenerationInterval[i])] = types.YChild{"LspGenerationInterval", &lspGenerationIntervals.LspGenerationInterval[i]}
    }
    lspGenerationIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspGenerationIntervals.EntityData)
}

// Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval
// LSP generation scheduling parameters
type Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval struct {
    EntityData types.CommonEntityData
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

func (lspGenerationInterval *Isis_Instances_Instance_LspGenerationIntervals_LspGenerationInterval) GetEntityData() *types.CommonEntityData {
    lspGenerationInterval.EntityData.YFilter = lspGenerationInterval.YFilter
    lspGenerationInterval.EntityData.YangName = "lsp-generation-interval"
    lspGenerationInterval.EntityData.BundleName = "cisco_ios_xr"
    lspGenerationInterval.EntityData.ParentYangName = "lsp-generation-intervals"
    lspGenerationInterval.EntityData.SegmentPath = "lsp-generation-interval" + "[level='" + fmt.Sprintf("%v", lspGenerationInterval.Level) + "']"
    lspGenerationInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspGenerationInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspGenerationInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspGenerationInterval.EntityData.Children = make(map[string]types.YChild)
    lspGenerationInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    lspGenerationInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", lspGenerationInterval.Level}
    lspGenerationInterval.EntityData.Leafs["maximum-wait"] = types.YLeaf{"MaximumWait", lspGenerationInterval.MaximumWait}
    lspGenerationInterval.EntityData.Leafs["initial-wait"] = types.YLeaf{"InitialWait", lspGenerationInterval.InitialWait}
    lspGenerationInterval.EntityData.Leafs["secondary-wait"] = types.YLeaf{"SecondaryWait", lspGenerationInterval.SecondaryWait}
    return &(lspGenerationInterval.EntityData)
}

// Isis_Instances_Instance_LspArrivalTimes
// LSP arrival time configuration
type Isis_Instances_Instance_LspArrivalTimes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum LSP arrival time. The type is slice of
    // Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime.
    LspArrivalTime []Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime
}

func (lspArrivalTimes *Isis_Instances_Instance_LspArrivalTimes) GetEntityData() *types.CommonEntityData {
    lspArrivalTimes.EntityData.YFilter = lspArrivalTimes.YFilter
    lspArrivalTimes.EntityData.YangName = "lsp-arrival-times"
    lspArrivalTimes.EntityData.BundleName = "cisco_ios_xr"
    lspArrivalTimes.EntityData.ParentYangName = "instance"
    lspArrivalTimes.EntityData.SegmentPath = "lsp-arrival-times"
    lspArrivalTimes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspArrivalTimes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspArrivalTimes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspArrivalTimes.EntityData.Children = make(map[string]types.YChild)
    lspArrivalTimes.EntityData.Children["lsp-arrival-time"] = types.YChild{"LspArrivalTime", nil}
    for i := range lspArrivalTimes.LspArrivalTime {
        lspArrivalTimes.EntityData.Children[types.GetSegmentPath(&lspArrivalTimes.LspArrivalTime[i])] = types.YChild{"LspArrivalTime", &lspArrivalTimes.LspArrivalTime[i]}
    }
    lspArrivalTimes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspArrivalTimes.EntityData)
}

// Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime
// Minimum LSP arrival time
type Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime struct {
    EntityData types.CommonEntityData
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

func (lspArrivalTime *Isis_Instances_Instance_LspArrivalTimes_LspArrivalTime) GetEntityData() *types.CommonEntityData {
    lspArrivalTime.EntityData.YFilter = lspArrivalTime.YFilter
    lspArrivalTime.EntityData.YangName = "lsp-arrival-time"
    lspArrivalTime.EntityData.BundleName = "cisco_ios_xr"
    lspArrivalTime.EntityData.ParentYangName = "lsp-arrival-times"
    lspArrivalTime.EntityData.SegmentPath = "lsp-arrival-time" + "[level='" + fmt.Sprintf("%v", lspArrivalTime.Level) + "']"
    lspArrivalTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspArrivalTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspArrivalTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspArrivalTime.EntityData.Children = make(map[string]types.YChild)
    lspArrivalTime.EntityData.Leafs = make(map[string]types.YLeaf)
    lspArrivalTime.EntityData.Leafs["level"] = types.YLeaf{"Level", lspArrivalTime.Level}
    lspArrivalTime.EntityData.Leafs["maximum-wait"] = types.YLeaf{"MaximumWait", lspArrivalTime.MaximumWait}
    lspArrivalTime.EntityData.Leafs["initial-wait"] = types.YLeaf{"InitialWait", lspArrivalTime.InitialWait}
    lspArrivalTime.EntityData.Leafs["secondary-wait"] = types.YLeaf{"SecondaryWait", lspArrivalTime.SecondaryWait}
    return &(lspArrivalTime.EntityData)
}

// Isis_Instances_Instance_TraceBufferSize
// Trace buffer size configuration
type Isis_Instances_Instance_TraceBufferSize struct {
    EntityData types.CommonEntityData
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

func (traceBufferSize *Isis_Instances_Instance_TraceBufferSize) GetEntityData() *types.CommonEntityData {
    traceBufferSize.EntityData.YFilter = traceBufferSize.YFilter
    traceBufferSize.EntityData.YangName = "trace-buffer-size"
    traceBufferSize.EntityData.BundleName = "cisco_ios_xr"
    traceBufferSize.EntityData.ParentYangName = "instance"
    traceBufferSize.EntityData.SegmentPath = "trace-buffer-size"
    traceBufferSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBufferSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBufferSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBufferSize.EntityData.Children = make(map[string]types.YChild)
    traceBufferSize.EntityData.Leafs = make(map[string]types.YLeaf)
    traceBufferSize.EntityData.Leafs["detailed"] = types.YLeaf{"Detailed", traceBufferSize.Detailed}
    traceBufferSize.EntityData.Leafs["standard"] = types.YLeaf{"Standard", traceBufferSize.Standard}
    traceBufferSize.EntityData.Leafs["severe"] = types.YLeaf{"Severe", traceBufferSize.Severe}
    traceBufferSize.EntityData.Leafs["hello"] = types.YLeaf{"Hello", traceBufferSize.Hello}
    return &(traceBufferSize.EntityData)
}

// Isis_Instances_Instance_MaxLinkMetrics
// Max Link Metric configuration
type Isis_Instances_Instance_MaxLinkMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Max Link Metric. The type is slice of
    // Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric.
    MaxLinkMetric []Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric
}

func (maxLinkMetrics *Isis_Instances_Instance_MaxLinkMetrics) GetEntityData() *types.CommonEntityData {
    maxLinkMetrics.EntityData.YFilter = maxLinkMetrics.YFilter
    maxLinkMetrics.EntityData.YangName = "max-link-metrics"
    maxLinkMetrics.EntityData.BundleName = "cisco_ios_xr"
    maxLinkMetrics.EntityData.ParentYangName = "instance"
    maxLinkMetrics.EntityData.SegmentPath = "max-link-metrics"
    maxLinkMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxLinkMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxLinkMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxLinkMetrics.EntityData.Children = make(map[string]types.YChild)
    maxLinkMetrics.EntityData.Children["max-link-metric"] = types.YChild{"MaxLinkMetric", nil}
    for i := range maxLinkMetrics.MaxLinkMetric {
        maxLinkMetrics.EntityData.Children[types.GetSegmentPath(&maxLinkMetrics.MaxLinkMetric[i])] = types.YChild{"MaxLinkMetric", &maxLinkMetrics.MaxLinkMetric[i]}
    }
    maxLinkMetrics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(maxLinkMetrics.EntityData)
}

// Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric
// Max Link Metric
type Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (maxLinkMetric *Isis_Instances_Instance_MaxLinkMetrics_MaxLinkMetric) GetEntityData() *types.CommonEntityData {
    maxLinkMetric.EntityData.YFilter = maxLinkMetric.YFilter
    maxLinkMetric.EntityData.YangName = "max-link-metric"
    maxLinkMetric.EntityData.BundleName = "cisco_ios_xr"
    maxLinkMetric.EntityData.ParentYangName = "max-link-metrics"
    maxLinkMetric.EntityData.SegmentPath = "max-link-metric" + "[level='" + fmt.Sprintf("%v", maxLinkMetric.Level) + "']"
    maxLinkMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxLinkMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxLinkMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxLinkMetric.EntityData.Children = make(map[string]types.YChild)
    maxLinkMetric.EntityData.Leafs = make(map[string]types.YLeaf)
    maxLinkMetric.EntityData.Leafs["level"] = types.YLeaf{"Level", maxLinkMetric.Level}
    return &(maxLinkMetric.EntityData)
}

// Isis_Instances_Instance_AdjacencyStagger
// Stagger ISIS adjacency bring up
// This type is a presence type.
type Isis_Instances_Instance_AdjacencyStagger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Adjacency Stagger: Initial number of neighbors to bring up per area. The
    // type is interface{} with range: 2..65000. The default value is 2.
    InitialNbr interface{}

    // Adjacency Stagger: Subsequent simultaneous number of neighbors to bring up.
    // The type is interface{} with range: 2..65000. The default value is 64.
    MaxNbr interface{}
}

func (adjacencyStagger *Isis_Instances_Instance_AdjacencyStagger) GetEntityData() *types.CommonEntityData {
    adjacencyStagger.EntityData.YFilter = adjacencyStagger.YFilter
    adjacencyStagger.EntityData.YangName = "adjacency-stagger"
    adjacencyStagger.EntityData.BundleName = "cisco_ios_xr"
    adjacencyStagger.EntityData.ParentYangName = "instance"
    adjacencyStagger.EntityData.SegmentPath = "adjacency-stagger"
    adjacencyStagger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencyStagger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencyStagger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencyStagger.EntityData.Children = make(map[string]types.YChild)
    adjacencyStagger.EntityData.Leafs = make(map[string]types.YLeaf)
    adjacencyStagger.EntityData.Leafs["initial-nbr"] = types.YLeaf{"InitialNbr", adjacencyStagger.InitialNbr}
    adjacencyStagger.EntityData.Leafs["max-nbr"] = types.YLeaf{"MaxNbr", adjacencyStagger.MaxNbr}
    return &(adjacencyStagger.EntityData)
}

// Isis_Instances_Instance_Afs
// Per-address-family configuration
type Isis_Instances_Instance_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for an IS-IS address-family. If a named (non-default)
    // topology is being created it must be multicast. The type is slice of
    // Isis_Instances_Instance_Afs_Af.
    Af []Isis_Instances_Instance_Afs_Af
}

func (afs *Isis_Instances_Instance_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "instance"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = make(map[string]types.YChild)
    afs.EntityData.Children["af"] = types.YChild{"Af", nil}
    for i := range afs.Af {
        afs.EntityData.Children[types.GetSegmentPath(&afs.Af[i])] = types.YChild{"Af", &afs.Af[i]}
    }
    afs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(afs.EntityData)
}

// Isis_Instances_Instance_Afs_Af
// Configuration for an IS-IS address-family. If
// a named (non-default) topology is being
// created it must be multicast.
type Isis_Instances_Instance_Afs_Af struct {
    EntityData types.CommonEntityData
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

func (af *Isis_Instances_Instance_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']" + "[saf-name='" + fmt.Sprintf("%v", af.SafName) + "']"
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = make(map[string]types.YChild)
    af.EntityData.Children["af-data"] = types.YChild{"AfData", &af.AfData}
    af.EntityData.Children["topology-name"] = types.YChild{"TopologyName", nil}
    for i := range af.TopologyName {
        af.EntityData.Children[types.GetSegmentPath(&af.TopologyName[i])] = types.YChild{"TopologyName", &af.TopologyName[i]}
    }
    af.EntityData.Leafs = make(map[string]types.YLeaf)
    af.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", af.AfName}
    af.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", af.SafName}
    return &(af.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData
// Data container.
// This type is a presence type.
type Isis_Instances_Instance_Afs_Af_AfData struct {
    EntityData types.CommonEntityData
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

func (afData *Isis_Instances_Instance_Afs_Af_AfData) GetEntityData() *types.CommonEntityData {
    afData.EntityData.YFilter = afData.YFilter
    afData.EntityData.YangName = "af-data"
    afData.EntityData.BundleName = "cisco_ios_xr"
    afData.EntityData.ParentYangName = "af"
    afData.EntityData.SegmentPath = "af-data"
    afData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afData.EntityData.Children = make(map[string]types.YChild)
    afData.EntityData.Children["segment-routing"] = types.YChild{"SegmentRouting", &afData.SegmentRouting}
    afData.EntityData.Children["metric-styles"] = types.YChild{"MetricStyles", &afData.MetricStyles}
    afData.EntityData.Children["frr-table"] = types.YChild{"FrrTable", &afData.FrrTable}
    afData.EntityData.Children["router-id"] = types.YChild{"RouterId", &afData.RouterId}
    afData.EntityData.Children["spf-prefix-priorities"] = types.YChild{"SpfPrefixPriorities", &afData.SpfPrefixPriorities}
    afData.EntityData.Children["summary-prefixes"] = types.YChild{"SummaryPrefixes", &afData.SummaryPrefixes}
    afData.EntityData.Children["micro-loop-avoidance"] = types.YChild{"MicroLoopAvoidance", &afData.MicroLoopAvoidance}
    afData.EntityData.Children["ucmp"] = types.YChild{"Ucmp", &afData.Ucmp}
    afData.EntityData.Children["max-redist-prefixes"] = types.YChild{"MaxRedistPrefixes", &afData.MaxRedistPrefixes}
    afData.EntityData.Children["propagations"] = types.YChild{"Propagations", &afData.Propagations}
    afData.EntityData.Children["redistributions"] = types.YChild{"Redistributions", &afData.Redistributions}
    afData.EntityData.Children["spf-periodic-intervals"] = types.YChild{"SpfPeriodicIntervals", &afData.SpfPeriodicIntervals}
    afData.EntityData.Children["spf-intervals"] = types.YChild{"SpfIntervals", &afData.SpfIntervals}
    afData.EntityData.Children["monitor-convergence"] = types.YChild{"MonitorConvergence", &afData.MonitorConvergence}
    afData.EntityData.Children["default-information"] = types.YChild{"DefaultInformation", &afData.DefaultInformation}
    afData.EntityData.Children["admin-distances"] = types.YChild{"AdminDistances", &afData.AdminDistances}
    afData.EntityData.Children["ispf"] = types.YChild{"Ispf", &afData.Ispf}
    afData.EntityData.Children["mpls-ldp-global"] = types.YChild{"MplsLdpGlobal", &afData.MplsLdpGlobal}
    afData.EntityData.Children["mpls"] = types.YChild{"Mpls", &afData.Mpls}
    afData.EntityData.Children["manual-adj-sids"] = types.YChild{"ManualAdjSids", &afData.ManualAdjSids}
    afData.EntityData.Children["metrics"] = types.YChild{"Metrics", &afData.Metrics}
    afData.EntityData.Children["weights"] = types.YChild{"Weights", &afData.Weights}
    afData.EntityData.Leafs = make(map[string]types.YLeaf)
    afData.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", afData.MaximumPaths}
    afData.EntityData.Leafs["topology-id"] = types.YLeaf{"TopologyId", afData.TopologyId}
    afData.EntityData.Leafs["single-topology"] = types.YLeaf{"SingleTopology", afData.SingleTopology}
    afData.EntityData.Leafs["adjacency-check"] = types.YLeaf{"AdjacencyCheck", afData.AdjacencyCheck}
    afData.EntityData.Leafs["advertise-link-attributes"] = types.YLeaf{"AdvertiseLinkAttributes", afData.AdvertiseLinkAttributes}
    afData.EntityData.Leafs["apply-weight"] = types.YLeaf{"ApplyWeight", afData.ApplyWeight}
    afData.EntityData.Leafs["default-admin-distance"] = types.YLeaf{"DefaultAdminDistance", afData.DefaultAdminDistance}
    afData.EntityData.Leafs["advertise-passive-only"] = types.YLeaf{"AdvertisePassiveOnly", afData.AdvertisePassiveOnly}
    afData.EntityData.Leafs["ignore-attached-bit"] = types.YLeaf{"IgnoreAttachedBit", afData.IgnoreAttachedBit}
    afData.EntityData.Leafs["attached-bit"] = types.YLeaf{"AttachedBit", afData.AttachedBit}
    afData.EntityData.Leafs["route-source-first-hop"] = types.YLeaf{"RouteSourceFirstHop", afData.RouteSourceFirstHop}
    return &(afData.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting
// Enable Segment Routing configuration
type Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable per bundle member adjacency SID. The type is interface{}.
    BundleMemberAdjSid interface{}

    // Prefer segment routing labels over LDP labels. The type is
    // IsisLabelPreference.
    Mpls interface{}

    // Enable Segment Routing prefix SID map configuration.
    PrefixSidMap Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting) GetEntityData() *types.CommonEntityData {
    segmentRouting.EntityData.YFilter = segmentRouting.YFilter
    segmentRouting.EntityData.YangName = "segment-routing"
    segmentRouting.EntityData.BundleName = "cisco_ios_xr"
    segmentRouting.EntityData.ParentYangName = "af-data"
    segmentRouting.EntityData.SegmentPath = "segment-routing"
    segmentRouting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRouting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRouting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRouting.EntityData.Children = make(map[string]types.YChild)
    segmentRouting.EntityData.Children["prefix-sid-map"] = types.YChild{"PrefixSidMap", &segmentRouting.PrefixSidMap}
    segmentRouting.EntityData.Leafs = make(map[string]types.YLeaf)
    segmentRouting.EntityData.Leafs["bundle-member-adj-sid"] = types.YLeaf{"BundleMemberAdjSid", segmentRouting.BundleMemberAdjSid}
    segmentRouting.EntityData.Leafs["mpls"] = types.YLeaf{"Mpls", segmentRouting.Mpls}
    return &(segmentRouting.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap
// Enable Segment Routing prefix SID map
// configuration
type Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Segment Routing prefix SID map advertise local. The type is
    // interface{}.
    AdvertiseLocal interface{}

    // If TRUE, remote prefix SID map advertisements will be used. If FALSE, they
    // will not be used. The type is bool.
    Receive interface{}
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_AfData_SegmentRouting_PrefixSidMap) GetEntityData() *types.CommonEntityData {
    prefixSidMap.EntityData.YFilter = prefixSidMap.YFilter
    prefixSidMap.EntityData.YangName = "prefix-sid-map"
    prefixSidMap.EntityData.BundleName = "cisco_ios_xr"
    prefixSidMap.EntityData.ParentYangName = "segment-routing"
    prefixSidMap.EntityData.SegmentPath = "prefix-sid-map"
    prefixSidMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSidMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSidMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSidMap.EntityData.Children = make(map[string]types.YChild)
    prefixSidMap.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixSidMap.EntityData.Leafs["advertise-local"] = types.YLeaf{"AdvertiseLocal", prefixSidMap.AdvertiseLocal}
    prefixSidMap.EntityData.Leafs["receive"] = types.YLeaf{"Receive", prefixSidMap.Receive}
    return &(prefixSidMap.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_MetricStyles
// Metric-style configuration
type Isis_Instances_Instance_Afs_Af_AfData_MetricStyles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of metric style in LSPs. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle.
    MetricStyle []Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles) GetEntityData() *types.CommonEntityData {
    metricStyles.EntityData.YFilter = metricStyles.YFilter
    metricStyles.EntityData.YangName = "metric-styles"
    metricStyles.EntityData.BundleName = "cisco_ios_xr"
    metricStyles.EntityData.ParentYangName = "af-data"
    metricStyles.EntityData.SegmentPath = "metric-styles"
    metricStyles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metricStyles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metricStyles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metricStyles.EntityData.Children = make(map[string]types.YChild)
    metricStyles.EntityData.Children["metric-style"] = types.YChild{"MetricStyle", nil}
    for i := range metricStyles.MetricStyle {
        metricStyles.EntityData.Children[types.GetSegmentPath(&metricStyles.MetricStyle[i])] = types.YChild{"MetricStyle", &metricStyles.MetricStyle[i]}
    }
    metricStyles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(metricStyles.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle
// Configuration of metric style in LSPs
type Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle struct {
    EntityData types.CommonEntityData
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

func (metricStyle *Isis_Instances_Instance_Afs_Af_AfData_MetricStyles_MetricStyle) GetEntityData() *types.CommonEntityData {
    metricStyle.EntityData.YFilter = metricStyle.YFilter
    metricStyle.EntityData.YangName = "metric-style"
    metricStyle.EntityData.BundleName = "cisco_ios_xr"
    metricStyle.EntityData.ParentYangName = "metric-styles"
    metricStyle.EntityData.SegmentPath = "metric-style" + "[level='" + fmt.Sprintf("%v", metricStyle.Level) + "']"
    metricStyle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metricStyle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metricStyle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metricStyle.EntityData.Children = make(map[string]types.YChild)
    metricStyle.EntityData.Leafs = make(map[string]types.YLeaf)
    metricStyle.EntityData.Leafs["level"] = types.YLeaf{"Level", metricStyle.Level}
    metricStyle.EntityData.Leafs["style"] = types.YLeaf{"Style", metricStyle.Style}
    metricStyle.EntityData.Leafs["transition-state"] = types.YLeaf{"TransitionState", metricStyle.TransitionState}
    return &(metricStyle.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable
// Fast-ReRoute configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable struct {
    EntityData types.CommonEntityData
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

func (frrTable *Isis_Instances_Instance_Afs_Af_AfData_FrrTable) GetEntityData() *types.CommonEntityData {
    frrTable.EntityData.YFilter = frrTable.YFilter
    frrTable.EntityData.YangName = "frr-table"
    frrTable.EntityData.BundleName = "cisco_ios_xr"
    frrTable.EntityData.ParentYangName = "af-data"
    frrTable.EntityData.SegmentPath = "frr-table"
    frrTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrTable.EntityData.Children = make(map[string]types.YChild)
    frrTable.EntityData.Children["frr-load-sharings"] = types.YChild{"FrrLoadSharings", &frrTable.FrrLoadSharings}
    frrTable.EntityData.Children["priority-limits"] = types.YChild{"PriorityLimits", &frrTable.PriorityLimits}
    frrTable.EntityData.Children["frr-remote-lfa-prefixes"] = types.YChild{"FrrRemoteLfaPrefixes", &frrTable.FrrRemoteLfaPrefixes}
    frrTable.EntityData.Children["frr-tiebreakers"] = types.YChild{"FrrTiebreakers", &frrTable.FrrTiebreakers}
    frrTable.EntityData.Children["frr-use-cand-onlies"] = types.YChild{"FrrUseCandOnlies", &frrTable.FrrUseCandOnlies}
    frrTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrTable.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings
// Load share prefixes across multiple
// backups
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable load sharing. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing.
    FrrLoadSharing []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings) GetEntityData() *types.CommonEntityData {
    frrLoadSharings.EntityData.YFilter = frrLoadSharings.YFilter
    frrLoadSharings.EntityData.YangName = "frr-load-sharings"
    frrLoadSharings.EntityData.BundleName = "cisco_ios_xr"
    frrLoadSharings.EntityData.ParentYangName = "frr-table"
    frrLoadSharings.EntityData.SegmentPath = "frr-load-sharings"
    frrLoadSharings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrLoadSharings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrLoadSharings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrLoadSharings.EntityData.Children = make(map[string]types.YChild)
    frrLoadSharings.EntityData.Children["frr-load-sharing"] = types.YChild{"FrrLoadSharing", nil}
    for i := range frrLoadSharings.FrrLoadSharing {
        frrLoadSharings.EntityData.Children[types.GetSegmentPath(&frrLoadSharings.FrrLoadSharing[i])] = types.YChild{"FrrLoadSharing", &frrLoadSharings.FrrLoadSharing[i]}
    }
    frrLoadSharings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrLoadSharings.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing
// Disable load sharing
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Load sharing. The type is IsisfrrLoadSharing. This attribute is mandatory.
    LoadSharing interface{}
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrLoadSharings_FrrLoadSharing) GetEntityData() *types.CommonEntityData {
    frrLoadSharing.EntityData.YFilter = frrLoadSharing.YFilter
    frrLoadSharing.EntityData.YangName = "frr-load-sharing"
    frrLoadSharing.EntityData.BundleName = "cisco_ios_xr"
    frrLoadSharing.EntityData.ParentYangName = "frr-load-sharings"
    frrLoadSharing.EntityData.SegmentPath = "frr-load-sharing" + "[level='" + fmt.Sprintf("%v", frrLoadSharing.Level) + "']"
    frrLoadSharing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrLoadSharing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrLoadSharing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrLoadSharing.EntityData.Children = make(map[string]types.YChild)
    frrLoadSharing.EntityData.Leafs = make(map[string]types.YLeaf)
    frrLoadSharing.EntityData.Leafs["level"] = types.YLeaf{"Level", frrLoadSharing.Level}
    frrLoadSharing.EntityData.Leafs["load-sharing"] = types.YLeaf{"LoadSharing", frrLoadSharing.LoadSharing}
    return &(frrLoadSharing.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits
// FRR prefix-limit configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit backup computation upto the prefix priority. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit.
    PriorityLimit []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits) GetEntityData() *types.CommonEntityData {
    priorityLimits.EntityData.YFilter = priorityLimits.YFilter
    priorityLimits.EntityData.YangName = "priority-limits"
    priorityLimits.EntityData.BundleName = "cisco_ios_xr"
    priorityLimits.EntityData.ParentYangName = "frr-table"
    priorityLimits.EntityData.SegmentPath = "priority-limits"
    priorityLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priorityLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priorityLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priorityLimits.EntityData.Children = make(map[string]types.YChild)
    priorityLimits.EntityData.Children["priority-limit"] = types.YChild{"PriorityLimit", nil}
    for i := range priorityLimits.PriorityLimit {
        priorityLimits.EntityData.Children[types.GetSegmentPath(&priorityLimits.PriorityLimit[i])] = types.YChild{"PriorityLimit", &priorityLimits.PriorityLimit[i]}
    }
    priorityLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(priorityLimits.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit
// Limit backup computation upto the prefix
// priority
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit struct {
    EntityData types.CommonEntityData
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

func (priorityLimit *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_PriorityLimits_PriorityLimit) GetEntityData() *types.CommonEntityData {
    priorityLimit.EntityData.YFilter = priorityLimit.YFilter
    priorityLimit.EntityData.YangName = "priority-limit"
    priorityLimit.EntityData.BundleName = "cisco_ios_xr"
    priorityLimit.EntityData.ParentYangName = "priority-limits"
    priorityLimit.EntityData.SegmentPath = "priority-limit" + "[level='" + fmt.Sprintf("%v", priorityLimit.Level) + "']" + "[frr-type='" + fmt.Sprintf("%v", priorityLimit.FrrType) + "']"
    priorityLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priorityLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priorityLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priorityLimit.EntityData.Children = make(map[string]types.YChild)
    priorityLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    priorityLimit.EntityData.Leafs["level"] = types.YLeaf{"Level", priorityLimit.Level}
    priorityLimit.EntityData.Leafs["frr-type"] = types.YLeaf{"FrrType", priorityLimit.FrrType}
    priorityLimit.EntityData.Leafs["priority"] = types.YLeaf{"Priority", priorityLimit.Priority}
    return &(priorityLimit.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes
// FRR remote LFA prefix list filter
// configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter remote LFA router IDs using prefix-list. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix.
    FrrRemoteLfaPrefix []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaPrefixes.EntityData.YFilter = frrRemoteLfaPrefixes.YFilter
    frrRemoteLfaPrefixes.EntityData.YangName = "frr-remote-lfa-prefixes"
    frrRemoteLfaPrefixes.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaPrefixes.EntityData.ParentYangName = "frr-table"
    frrRemoteLfaPrefixes.EntityData.SegmentPath = "frr-remote-lfa-prefixes"
    frrRemoteLfaPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaPrefixes.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaPrefixes.EntityData.Children["frr-remote-lfa-prefix"] = types.YChild{"FrrRemoteLfaPrefix", nil}
    for i := range frrRemoteLfaPrefixes.FrrRemoteLfaPrefix {
        frrRemoteLfaPrefixes.EntityData.Children[types.GetSegmentPath(&frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[i])] = types.YChild{"FrrRemoteLfaPrefix", &frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[i]}
    }
    frrRemoteLfaPrefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrRemoteLfaPrefixes.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix
// Filter remote LFA router IDs using
// prefix-list
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Name of the prefix list. The type is string with length: 1..32. This
    // attribute is mandatory.
    PrefixListName interface{}
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaPrefix.EntityData.YFilter = frrRemoteLfaPrefix.YFilter
    frrRemoteLfaPrefix.EntityData.YangName = "frr-remote-lfa-prefix"
    frrRemoteLfaPrefix.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaPrefix.EntityData.ParentYangName = "frr-remote-lfa-prefixes"
    frrRemoteLfaPrefix.EntityData.SegmentPath = "frr-remote-lfa-prefix" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaPrefix.Level) + "']"
    frrRemoteLfaPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaPrefix.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    frrRemoteLfaPrefix.EntityData.Leafs["level"] = types.YLeaf{"Level", frrRemoteLfaPrefix.Level}
    frrRemoteLfaPrefix.EntityData.Leafs["prefix-list-name"] = types.YLeaf{"PrefixListName", frrRemoteLfaPrefix.PrefixListName}
    return &(frrRemoteLfaPrefix.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers
// FRR tiebreakers configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure tiebreaker for multiple backups. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker.
    FrrTiebreaker []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers) GetEntityData() *types.CommonEntityData {
    frrTiebreakers.EntityData.YFilter = frrTiebreakers.YFilter
    frrTiebreakers.EntityData.YangName = "frr-tiebreakers"
    frrTiebreakers.EntityData.BundleName = "cisco_ios_xr"
    frrTiebreakers.EntityData.ParentYangName = "frr-table"
    frrTiebreakers.EntityData.SegmentPath = "frr-tiebreakers"
    frrTiebreakers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrTiebreakers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrTiebreakers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrTiebreakers.EntityData.Children = make(map[string]types.YChild)
    frrTiebreakers.EntityData.Children["frr-tiebreaker"] = types.YChild{"FrrTiebreaker", nil}
    for i := range frrTiebreakers.FrrTiebreaker {
        frrTiebreakers.EntityData.Children[types.GetSegmentPath(&frrTiebreakers.FrrTiebreaker[i])] = types.YChild{"FrrTiebreaker", &frrTiebreakers.FrrTiebreaker[i]}
    }
    frrTiebreakers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrTiebreakers.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker
// Configure tiebreaker for multiple backups
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker struct {
    EntityData types.CommonEntityData
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

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrTiebreakers_FrrTiebreaker) GetEntityData() *types.CommonEntityData {
    frrTiebreaker.EntityData.YFilter = frrTiebreaker.YFilter
    frrTiebreaker.EntityData.YangName = "frr-tiebreaker"
    frrTiebreaker.EntityData.BundleName = "cisco_ios_xr"
    frrTiebreaker.EntityData.ParentYangName = "frr-tiebreakers"
    frrTiebreaker.EntityData.SegmentPath = "frr-tiebreaker" + "[level='" + fmt.Sprintf("%v", frrTiebreaker.Level) + "']" + "[tiebreaker='" + fmt.Sprintf("%v", frrTiebreaker.Tiebreaker) + "']"
    frrTiebreaker.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrTiebreaker.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrTiebreaker.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrTiebreaker.EntityData.Children = make(map[string]types.YChild)
    frrTiebreaker.EntityData.Leafs = make(map[string]types.YLeaf)
    frrTiebreaker.EntityData.Leafs["level"] = types.YLeaf{"Level", frrTiebreaker.Level}
    frrTiebreaker.EntityData.Leafs["tiebreaker"] = types.YLeaf{"Tiebreaker", frrTiebreaker.Tiebreaker}
    frrTiebreaker.EntityData.Leafs["index"] = types.YLeaf{"Index", frrTiebreaker.Index}
    return &(frrTiebreaker.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies
// FRR use candidate only configuration
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure use candidate only to exclude interfaces as backup. The type is
    // slice of
    // Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly.
    FrrUseCandOnly []Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies) GetEntityData() *types.CommonEntityData {
    frrUseCandOnlies.EntityData.YFilter = frrUseCandOnlies.YFilter
    frrUseCandOnlies.EntityData.YangName = "frr-use-cand-onlies"
    frrUseCandOnlies.EntityData.BundleName = "cisco_ios_xr"
    frrUseCandOnlies.EntityData.ParentYangName = "frr-table"
    frrUseCandOnlies.EntityData.SegmentPath = "frr-use-cand-onlies"
    frrUseCandOnlies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrUseCandOnlies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrUseCandOnlies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrUseCandOnlies.EntityData.Children = make(map[string]types.YChild)
    frrUseCandOnlies.EntityData.Children["frr-use-cand-only"] = types.YChild{"FrrUseCandOnly", nil}
    for i := range frrUseCandOnlies.FrrUseCandOnly {
        frrUseCandOnlies.EntityData.Children[types.GetSegmentPath(&frrUseCandOnlies.FrrUseCandOnly[i])] = types.YChild{"FrrUseCandOnly", &frrUseCandOnlies.FrrUseCandOnly[i]}
    }
    frrUseCandOnlies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrUseCandOnlies.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly
// Configure use candidate only to exclude
// interfaces as backup
type Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_AfData_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetEntityData() *types.CommonEntityData {
    frrUseCandOnly.EntityData.YFilter = frrUseCandOnly.YFilter
    frrUseCandOnly.EntityData.YangName = "frr-use-cand-only"
    frrUseCandOnly.EntityData.BundleName = "cisco_ios_xr"
    frrUseCandOnly.EntityData.ParentYangName = "frr-use-cand-onlies"
    frrUseCandOnly.EntityData.SegmentPath = "frr-use-cand-only" + "[level='" + fmt.Sprintf("%v", frrUseCandOnly.Level) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrUseCandOnly.FrrType) + "']"
    frrUseCandOnly.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrUseCandOnly.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrUseCandOnly.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrUseCandOnly.EntityData.Children = make(map[string]types.YChild)
    frrUseCandOnly.EntityData.Leafs = make(map[string]types.YLeaf)
    frrUseCandOnly.EntityData.Leafs["level"] = types.YLeaf{"Level", frrUseCandOnly.Level}
    frrUseCandOnly.EntityData.Leafs["frr-type"] = types.YLeaf{"FrrType", frrUseCandOnly.FrrType}
    return &(frrUseCandOnly.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_RouterId
// Stable IP address for system. Will only be
// applied for the unicast sub-address-family.
type Isis_Instances_Instance_Afs_Af_AfData_RouterId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4/IPv6 address to be used as a router ID. Precisely one of Address and
    // Interface must be specified. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Interface with designated stable IP address to be used as a router ID. This
    // must be a Loopback interface. Precisely one of Address and Interface must
    // be specified. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_RouterId) GetEntityData() *types.CommonEntityData {
    routerId.EntityData.YFilter = routerId.YFilter
    routerId.EntityData.YangName = "router-id"
    routerId.EntityData.BundleName = "cisco_ios_xr"
    routerId.EntityData.ParentYangName = "af-data"
    routerId.EntityData.SegmentPath = "router-id"
    routerId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routerId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routerId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routerId.EntityData.Children = make(map[string]types.YChild)
    routerId.EntityData.Leafs = make(map[string]types.YLeaf)
    routerId.EntityData.Leafs["address"] = types.YLeaf{"Address", routerId.Address}
    routerId.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", routerId.InterfaceName}
    return &(routerId.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities
// SPF Prefix Priority configuration
type Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Determine SPF priority for prefixes. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority.
    SpfPrefixPriority []Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities) GetEntityData() *types.CommonEntityData {
    spfPrefixPriorities.EntityData.YFilter = spfPrefixPriorities.YFilter
    spfPrefixPriorities.EntityData.YangName = "spf-prefix-priorities"
    spfPrefixPriorities.EntityData.BundleName = "cisco_ios_xr"
    spfPrefixPriorities.EntityData.ParentYangName = "af-data"
    spfPrefixPriorities.EntityData.SegmentPath = "spf-prefix-priorities"
    spfPrefixPriorities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfPrefixPriorities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfPrefixPriorities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfPrefixPriorities.EntityData.Children = make(map[string]types.YChild)
    spfPrefixPriorities.EntityData.Children["spf-prefix-priority"] = types.YChild{"SpfPrefixPriority", nil}
    for i := range spfPrefixPriorities.SpfPrefixPriority {
        spfPrefixPriorities.EntityData.Children[types.GetSegmentPath(&spfPrefixPriorities.SpfPrefixPriority[i])] = types.YChild{"SpfPrefixPriority", &spfPrefixPriorities.SpfPrefixPriority[i]}
    }
    spfPrefixPriorities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spfPrefixPriorities.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority
// Determine SPF priority for prefixes
type Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority struct {
    EntityData types.CommonEntityData
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

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_AfData_SpfPrefixPriorities_SpfPrefixPriority) GetEntityData() *types.CommonEntityData {
    spfPrefixPriority.EntityData.YFilter = spfPrefixPriority.YFilter
    spfPrefixPriority.EntityData.YangName = "spf-prefix-priority"
    spfPrefixPriority.EntityData.BundleName = "cisco_ios_xr"
    spfPrefixPriority.EntityData.ParentYangName = "spf-prefix-priorities"
    spfPrefixPriority.EntityData.SegmentPath = "spf-prefix-priority" + "[level='" + fmt.Sprintf("%v", spfPrefixPriority.Level) + "']" + "[prefix-priority-type='" + fmt.Sprintf("%v", spfPrefixPriority.PrefixPriorityType) + "']"
    spfPrefixPriority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfPrefixPriority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfPrefixPriority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfPrefixPriority.EntityData.Children = make(map[string]types.YChild)
    spfPrefixPriority.EntityData.Leafs = make(map[string]types.YLeaf)
    spfPrefixPriority.EntityData.Leafs["level"] = types.YLeaf{"Level", spfPrefixPriority.Level}
    spfPrefixPriority.EntityData.Leafs["prefix-priority-type"] = types.YLeaf{"PrefixPriorityType", spfPrefixPriority.PrefixPriorityType}
    spfPrefixPriority.EntityData.Leafs["admin-tag"] = types.YLeaf{"AdminTag", spfPrefixPriority.AdminTag}
    spfPrefixPriority.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", spfPrefixPriority.AccessListName}
    return &(spfPrefixPriority.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes
// Summary-prefix configuration
type Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IP address prefixes to advertise. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix.
    SummaryPrefix []Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes) GetEntityData() *types.CommonEntityData {
    summaryPrefixes.EntityData.YFilter = summaryPrefixes.YFilter
    summaryPrefixes.EntityData.YangName = "summary-prefixes"
    summaryPrefixes.EntityData.BundleName = "cisco_ios_xr"
    summaryPrefixes.EntityData.ParentYangName = "af-data"
    summaryPrefixes.EntityData.SegmentPath = "summary-prefixes"
    summaryPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPrefixes.EntityData.Children = make(map[string]types.YChild)
    summaryPrefixes.EntityData.Children["summary-prefix"] = types.YChild{"SummaryPrefix", nil}
    for i := range summaryPrefixes.SummaryPrefix {
        summaryPrefixes.EntityData.Children[types.GetSegmentPath(&summaryPrefixes.SummaryPrefix[i])] = types.YChild{"SummaryPrefix", &summaryPrefixes.SummaryPrefix[i]}
    }
    summaryPrefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summaryPrefixes.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix
// Configure IP address prefixes to advertise
type Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP summary address prefix. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    AddressPrefix interface{}

    // The tag value. The type is interface{} with range: 1..4294967295.
    Tag interface{}

    // Level in which to summarize routes. The type is interface{} with range:
    // 1..2.
    Level interface{}
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_AfData_SummaryPrefixes_SummaryPrefix) GetEntityData() *types.CommonEntityData {
    summaryPrefix.EntityData.YFilter = summaryPrefix.YFilter
    summaryPrefix.EntityData.YangName = "summary-prefix"
    summaryPrefix.EntityData.BundleName = "cisco_ios_xr"
    summaryPrefix.EntityData.ParentYangName = "summary-prefixes"
    summaryPrefix.EntityData.SegmentPath = "summary-prefix" + "[address-prefix='" + fmt.Sprintf("%v", summaryPrefix.AddressPrefix) + "']"
    summaryPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPrefix.EntityData.Children = make(map[string]types.YChild)
    summaryPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    summaryPrefix.EntityData.Leafs["address-prefix"] = types.YLeaf{"AddressPrefix", summaryPrefix.AddressPrefix}
    summaryPrefix.EntityData.Leafs["tag"] = types.YLeaf{"Tag", summaryPrefix.Tag}
    summaryPrefix.EntityData.Leafs["level"] = types.YLeaf{"Level", summaryPrefix.Level}
    return &(summaryPrefix.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance
// Micro Loop Avoidance configuration
type Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MicroLoop avoidance enable configuration. The type is
    // IsisMicroLoopAvoidance.
    Enable interface{}

    // Value of delay in msecs in updating RIB. The type is interface{} with
    // range: 1000..65535. Units are millisecond. The default value is 5000.
    RibUpdateDelay interface{}
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_AfData_MicroLoopAvoidance) GetEntityData() *types.CommonEntityData {
    microLoopAvoidance.EntityData.YFilter = microLoopAvoidance.YFilter
    microLoopAvoidance.EntityData.YangName = "micro-loop-avoidance"
    microLoopAvoidance.EntityData.BundleName = "cisco_ios_xr"
    microLoopAvoidance.EntityData.ParentYangName = "af-data"
    microLoopAvoidance.EntityData.SegmentPath = "micro-loop-avoidance"
    microLoopAvoidance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    microLoopAvoidance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    microLoopAvoidance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    microLoopAvoidance.EntityData.Children = make(map[string]types.YChild)
    microLoopAvoidance.EntityData.Leafs = make(map[string]types.YLeaf)
    microLoopAvoidance.EntityData.Leafs["enable"] = types.YLeaf{"Enable", microLoopAvoidance.Enable}
    microLoopAvoidance.EntityData.Leafs["rib-update-delay"] = types.YLeaf{"RibUpdateDelay", microLoopAvoidance.RibUpdateDelay}
    return &(microLoopAvoidance.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Ucmp
// UCMP (UnEqual Cost MultiPath) configuration
type Isis_Instances_Instance_Afs_Af_AfData_Ucmp struct {
    EntityData types.CommonEntityData
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

func (ucmp *Isis_Instances_Instance_Afs_Af_AfData_Ucmp) GetEntityData() *types.CommonEntityData {
    ucmp.EntityData.YFilter = ucmp.YFilter
    ucmp.EntityData.YangName = "ucmp"
    ucmp.EntityData.BundleName = "cisco_ios_xr"
    ucmp.EntityData.ParentYangName = "af-data"
    ucmp.EntityData.SegmentPath = "ucmp"
    ucmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ucmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ucmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ucmp.EntityData.Children = make(map[string]types.YChild)
    ucmp.EntityData.Children["enable"] = types.YChild{"Enable", &ucmp.Enable}
    ucmp.EntityData.Children["exclude-interfaces"] = types.YChild{"ExcludeInterfaces", &ucmp.ExcludeInterfaces}
    ucmp.EntityData.Leafs = make(map[string]types.YLeaf)
    ucmp.EntityData.Leafs["delay-interval"] = types.YLeaf{"DelayInterval", ucmp.DelayInterval}
    return &(ucmp.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable
// UCMP feature enable configuration
type Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Value of variance. The type is interface{} with range: 101..10000. The
    // default value is 200.
    Variance interface{}

    // Name of the Prefix List. The type is string with length: 1..32.
    PrefixListName interface{}
}

func (enable *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_Enable) GetEntityData() *types.CommonEntityData {
    enable.EntityData.YFilter = enable.YFilter
    enable.EntityData.YangName = "enable"
    enable.EntityData.BundleName = "cisco_ios_xr"
    enable.EntityData.ParentYangName = "ucmp"
    enable.EntityData.SegmentPath = "enable"
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = make(map[string]types.YChild)
    enable.EntityData.Leafs = make(map[string]types.YLeaf)
    enable.EntityData.Leafs["variance"] = types.YLeaf{"Variance", enable.Variance}
    enable.EntityData.Leafs["prefix-list-name"] = types.YLeaf{"PrefixListName", enable.PrefixListName}
    return &(enable.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces
// Interfaces excluded from UCMP path
// computation
type Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude this interface from UCMP path computation. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "ucmp"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = make(map[string]types.YChild)
    excludeInterfaces.EntityData.Children["exclude-interface"] = types.YChild{"ExcludeInterface", nil}
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children[types.GetSegmentPath(&excludeInterfaces.ExcludeInterface[i])] = types.YChild{"ExcludeInterface", &excludeInterfaces.ExcludeInterface[i]}
    }
    excludeInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(excludeInterfaces.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface
// Exclude this interface from UCMP path
// computation
type Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface to be excluded. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_AfData_Ucmp_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + "[interface-name='" + fmt.Sprintf("%v", excludeInterface.InterfaceName) + "']"
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = make(map[string]types.YChild)
    excludeInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    excludeInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", excludeInterface.InterfaceName}
    return &(excludeInterface.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes
// Maximum number of redistributed
// prefixesconfiguration
type Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An upper limit on the number of redistributed prefixes which may be
    // included in the local system's LSP. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix.
    MaxRedistPrefix []Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes) GetEntityData() *types.CommonEntityData {
    maxRedistPrefixes.EntityData.YFilter = maxRedistPrefixes.YFilter
    maxRedistPrefixes.EntityData.YangName = "max-redist-prefixes"
    maxRedistPrefixes.EntityData.BundleName = "cisco_ios_xr"
    maxRedistPrefixes.EntityData.ParentYangName = "af-data"
    maxRedistPrefixes.EntityData.SegmentPath = "max-redist-prefixes"
    maxRedistPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxRedistPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxRedistPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxRedistPrefixes.EntityData.Children = make(map[string]types.YChild)
    maxRedistPrefixes.EntityData.Children["max-redist-prefix"] = types.YChild{"MaxRedistPrefix", nil}
    for i := range maxRedistPrefixes.MaxRedistPrefix {
        maxRedistPrefixes.EntityData.Children[types.GetSegmentPath(&maxRedistPrefixes.MaxRedistPrefix[i])] = types.YChild{"MaxRedistPrefix", &maxRedistPrefixes.MaxRedistPrefix[i]}
    }
    maxRedistPrefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(maxRedistPrefixes.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix
// An upper limit on the number of
// redistributed prefixes which may be
// included in the local system's LSP
type Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Max number of prefixes. The type is interface{} with range: 1..28000. This
    // attribute is mandatory.
    PrefixLimit interface{}
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_AfData_MaxRedistPrefixes_MaxRedistPrefix) GetEntityData() *types.CommonEntityData {
    maxRedistPrefix.EntityData.YFilter = maxRedistPrefix.YFilter
    maxRedistPrefix.EntityData.YangName = "max-redist-prefix"
    maxRedistPrefix.EntityData.BundleName = "cisco_ios_xr"
    maxRedistPrefix.EntityData.ParentYangName = "max-redist-prefixes"
    maxRedistPrefix.EntityData.SegmentPath = "max-redist-prefix" + "[level='" + fmt.Sprintf("%v", maxRedistPrefix.Level) + "']"
    maxRedistPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxRedistPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxRedistPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxRedistPrefix.EntityData.Children = make(map[string]types.YChild)
    maxRedistPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    maxRedistPrefix.EntityData.Leafs["level"] = types.YLeaf{"Level", maxRedistPrefix.Level}
    maxRedistPrefix.EntityData.Leafs["prefix-limit"] = types.YLeaf{"PrefixLimit", maxRedistPrefix.PrefixLimit}
    return &(maxRedistPrefix.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Propagations
// Route propagation configuration
type Isis_Instances_Instance_Afs_Af_AfData_Propagations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Propagate routes between IS-IS levels. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation.
    Propagation []Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation
}

func (propagations *Isis_Instances_Instance_Afs_Af_AfData_Propagations) GetEntityData() *types.CommonEntityData {
    propagations.EntityData.YFilter = propagations.YFilter
    propagations.EntityData.YangName = "propagations"
    propagations.EntityData.BundleName = "cisco_ios_xr"
    propagations.EntityData.ParentYangName = "af-data"
    propagations.EntityData.SegmentPath = "propagations"
    propagations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    propagations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    propagations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    propagations.EntityData.Children = make(map[string]types.YChild)
    propagations.EntityData.Children["propagation"] = types.YChild{"Propagation", nil}
    for i := range propagations.Propagation {
        propagations.EntityData.Children[types.GetSegmentPath(&propagations.Propagation[i])] = types.YChild{"Propagation", &propagations.Propagation[i]}
    }
    propagations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(propagations.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation
// Propagate routes between IS-IS levels
type Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation struct {
    EntityData types.CommonEntityData
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

func (propagation *Isis_Instances_Instance_Afs_Af_AfData_Propagations_Propagation) GetEntityData() *types.CommonEntityData {
    propagation.EntityData.YFilter = propagation.YFilter
    propagation.EntityData.YangName = "propagation"
    propagation.EntityData.BundleName = "cisco_ios_xr"
    propagation.EntityData.ParentYangName = "propagations"
    propagation.EntityData.SegmentPath = "propagation" + "[source-level='" + fmt.Sprintf("%v", propagation.SourceLevel) + "']" + "[destination-level='" + fmt.Sprintf("%v", propagation.DestinationLevel) + "']"
    propagation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    propagation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    propagation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    propagation.EntityData.Children = make(map[string]types.YChild)
    propagation.EntityData.Leafs = make(map[string]types.YLeaf)
    propagation.EntityData.Leafs["source-level"] = types.YLeaf{"SourceLevel", propagation.SourceLevel}
    propagation.EntityData.Leafs["destination-level"] = types.YLeaf{"DestinationLevel", propagation.DestinationLevel}
    propagation.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", propagation.RoutePolicyName}
    return &(propagation.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions
// Protocol redistribution configuration
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribution of other protocols into this IS-IS instance. The type is
    // slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution.
    Redistribution []Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution
}

func (redistributions *Isis_Instances_Instance_Afs_Af_AfData_Redistributions) GetEntityData() *types.CommonEntityData {
    redistributions.EntityData.YFilter = redistributions.YFilter
    redistributions.EntityData.YangName = "redistributions"
    redistributions.EntityData.BundleName = "cisco_ios_xr"
    redistributions.EntityData.ParentYangName = "af-data"
    redistributions.EntityData.SegmentPath = "redistributions"
    redistributions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributions.EntityData.Children = make(map[string]types.YChild)
    redistributions.EntityData.Children["redistribution"] = types.YChild{"Redistribution", nil}
    for i := range redistributions.Redistribution {
        redistributions.EntityData.Children[types.GetSegmentPath(&redistributions.Redistribution[i])] = types.YChild{"Redistribution", &redistributions.Redistribution[i]}
    }
    redistributions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(redistributions.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution
// Redistribution of other protocols into
// this IS-IS instance
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution struct {
    EntityData types.CommonEntityData
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

func (redistribution *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution) GetEntityData() *types.CommonEntityData {
    redistribution.EntityData.YFilter = redistribution.YFilter
    redistribution.EntityData.YangName = "redistribution"
    redistribution.EntityData.BundleName = "cisco_ios_xr"
    redistribution.EntityData.ParentYangName = "redistributions"
    redistribution.EntityData.SegmentPath = "redistribution" + "[protocol-name='" + fmt.Sprintf("%v", redistribution.ProtocolName) + "']"
    redistribution.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistribution.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistribution.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistribution.EntityData.Children = make(map[string]types.YChild)
    redistribution.EntityData.Children["connected-or-static-or-rip-or-subscriber-or-mobile"] = types.YChild{"ConnectedOrStaticOrRipOrSubscriberOrMobile", &redistribution.ConnectedOrStaticOrRipOrSubscriberOrMobile}
    redistribution.EntityData.Children["ospf-or-ospfv3-or-isis-or-application"] = types.YChild{"OspfOrOspfv3OrIsisOrApplication", nil}
    for i := range redistribution.OspfOrOspfv3OrIsisOrApplication {
        redistribution.EntityData.Children[types.GetSegmentPath(&redistribution.OspfOrOspfv3OrIsisOrApplication[i])] = types.YChild{"OspfOrOspfv3OrIsisOrApplication", &redistribution.OspfOrOspfv3OrIsisOrApplication[i]}
    }
    redistribution.EntityData.Children["bgp"] = types.YChild{"Bgp", nil}
    for i := range redistribution.Bgp {
        redistribution.EntityData.Children[types.GetSegmentPath(&redistribution.Bgp[i])] = types.YChild{"Bgp", &redistribution.Bgp[i]}
    }
    redistribution.EntityData.Children["eigrp"] = types.YChild{"Eigrp", nil}
    for i := range redistribution.Eigrp {
        redistribution.EntityData.Children[types.GetSegmentPath(&redistribution.Eigrp[i])] = types.YChild{"Eigrp", &redistribution.Eigrp[i]}
    }
    redistribution.EntityData.Leafs = make(map[string]types.YLeaf)
    redistribution.EntityData.Leafs["protocol-name"] = types.YLeaf{"ProtocolName", redistribution.ProtocolName}
    return &(redistribution.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile
// connected or static or rip or subscriber
// or mobile
// This type is a presence type.
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile struct {
    EntityData types.CommonEntityData
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

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetEntityData() *types.CommonEntityData {
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.YFilter = connectedOrStaticOrRipOrSubscriberOrMobile.YFilter
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.YangName = "connected-or-static-or-rip-or-subscriber-or-mobile"
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.BundleName = "cisco_ios_xr"
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.ParentYangName = "redistribution"
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.SegmentPath = "connected-or-static-or-rip-or-subscriber-or-mobile"
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Children = make(map[string]types.YChild)
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs = make(map[string]types.YLeaf)
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["metric"] = types.YLeaf{"Metric", connectedOrStaticOrRipOrSubscriberOrMobile.Metric}
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["levels"] = types.YLeaf{"Levels", connectedOrStaticOrRipOrSubscriberOrMobile.Levels}
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", connectedOrStaticOrRipOrSubscriberOrMobile.RoutePolicyName}
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", connectedOrStaticOrRipOrSubscriberOrMobile.MetricType}
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["ospf-route-type"] = types.YLeaf{"OspfRouteType", connectedOrStaticOrRipOrSubscriberOrMobile.OspfRouteType}
    return &(connectedOrStaticOrRipOrSubscriberOrMobile.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication
// ospf or ospfv3 or isis or application
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication struct {
    EntityData types.CommonEntityData
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

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetEntityData() *types.CommonEntityData {
    ospfOrOspfv3OrIsisOrApplication.EntityData.YFilter = ospfOrOspfv3OrIsisOrApplication.YFilter
    ospfOrOspfv3OrIsisOrApplication.EntityData.YangName = "ospf-or-ospfv3-or-isis-or-application"
    ospfOrOspfv3OrIsisOrApplication.EntityData.BundleName = "cisco_ios_xr"
    ospfOrOspfv3OrIsisOrApplication.EntityData.ParentYangName = "redistribution"
    ospfOrOspfv3OrIsisOrApplication.EntityData.SegmentPath = "ospf-or-ospfv3-or-isis-or-application" + "[instance-name='" + fmt.Sprintf("%v", ospfOrOspfv3OrIsisOrApplication.InstanceName) + "']"
    ospfOrOspfv3OrIsisOrApplication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfOrOspfv3OrIsisOrApplication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfOrOspfv3OrIsisOrApplication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfOrOspfv3OrIsisOrApplication.EntityData.Children = make(map[string]types.YChild)
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["instance-name"] = types.YLeaf{"InstanceName", ospfOrOspfv3OrIsisOrApplication.InstanceName}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfOrOspfv3OrIsisOrApplication.Metric}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["levels"] = types.YLeaf{"Levels", ospfOrOspfv3OrIsisOrApplication.Levels}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", ospfOrOspfv3OrIsisOrApplication.RoutePolicyName}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", ospfOrOspfv3OrIsisOrApplication.MetricType}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["ospf-route-type"] = types.YLeaf{"OspfRouteType", ospfOrOspfv3OrIsisOrApplication.OspfRouteType}
    return &(ospfOrOspfv3OrIsisOrApplication.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp
// bgp
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp struct {
    EntityData types.CommonEntityData
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

func (bgp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "redistribution"
    bgp.EntityData.SegmentPath = "bgp" + "[as-xx='" + fmt.Sprintf("%v", bgp.AsXx) + "']" + "[as-yy='" + fmt.Sprintf("%v", bgp.AsYy) + "']"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["as-xx"] = types.YLeaf{"AsXx", bgp.AsXx}
    bgp.EntityData.Leafs["as-yy"] = types.YLeaf{"AsYy", bgp.AsYy}
    bgp.EntityData.Leafs["metric"] = types.YLeaf{"Metric", bgp.Metric}
    bgp.EntityData.Leafs["levels"] = types.YLeaf{"Levels", bgp.Levels}
    bgp.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", bgp.RoutePolicyName}
    bgp.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", bgp.MetricType}
    bgp.EntityData.Leafs["ospf-route-type"] = types.YLeaf{"OspfRouteType", bgp.OspfRouteType}
    return &(bgp.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp
// eigrp
type Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp struct {
    EntityData types.CommonEntityData
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

func (eigrp *Isis_Instances_Instance_Afs_Af_AfData_Redistributions_Redistribution_Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "redistribution"
    eigrp.EntityData.SegmentPath = "eigrp" + "[as-zz='" + fmt.Sprintf("%v", eigrp.AsZz) + "']"
    eigrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrp.EntityData.Children = make(map[string]types.YChild)
    eigrp.EntityData.Leafs = make(map[string]types.YLeaf)
    eigrp.EntityData.Leafs["as-zz"] = types.YLeaf{"AsZz", eigrp.AsZz}
    eigrp.EntityData.Leafs["metric"] = types.YLeaf{"Metric", eigrp.Metric}
    eigrp.EntityData.Leafs["levels"] = types.YLeaf{"Levels", eigrp.Levels}
    eigrp.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", eigrp.RoutePolicyName}
    eigrp.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", eigrp.MetricType}
    eigrp.EntityData.Leafs["ospf-route-type"] = types.YLeaf{"OspfRouteType", eigrp.OspfRouteType}
    return &(eigrp.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals
// Peoridic SPF configuration
type Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum interval between spf runs. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval.
    SpfPeriodicInterval []Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals) GetEntityData() *types.CommonEntityData {
    spfPeriodicIntervals.EntityData.YFilter = spfPeriodicIntervals.YFilter
    spfPeriodicIntervals.EntityData.YangName = "spf-periodic-intervals"
    spfPeriodicIntervals.EntityData.BundleName = "cisco_ios_xr"
    spfPeriodicIntervals.EntityData.ParentYangName = "af-data"
    spfPeriodicIntervals.EntityData.SegmentPath = "spf-periodic-intervals"
    spfPeriodicIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfPeriodicIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfPeriodicIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfPeriodicIntervals.EntityData.Children = make(map[string]types.YChild)
    spfPeriodicIntervals.EntityData.Children["spf-periodic-interval"] = types.YChild{"SpfPeriodicInterval", nil}
    for i := range spfPeriodicIntervals.SpfPeriodicInterval {
        spfPeriodicIntervals.EntityData.Children[types.GetSegmentPath(&spfPeriodicIntervals.SpfPeriodicInterval[i])] = types.YChild{"SpfPeriodicInterval", &spfPeriodicIntervals.SpfPeriodicInterval[i]}
    }
    spfPeriodicIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spfPeriodicIntervals.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval
// Maximum interval between spf runs
type Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Maximum interval in between SPF runs in seconds. The type is interface{}
    // with range: 0..3600. This attribute is mandatory. Units are second.
    PeriodicInterval interface{}
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfPeriodicIntervals_SpfPeriodicInterval) GetEntityData() *types.CommonEntityData {
    spfPeriodicInterval.EntityData.YFilter = spfPeriodicInterval.YFilter
    spfPeriodicInterval.EntityData.YangName = "spf-periodic-interval"
    spfPeriodicInterval.EntityData.BundleName = "cisco_ios_xr"
    spfPeriodicInterval.EntityData.ParentYangName = "spf-periodic-intervals"
    spfPeriodicInterval.EntityData.SegmentPath = "spf-periodic-interval" + "[level='" + fmt.Sprintf("%v", spfPeriodicInterval.Level) + "']"
    spfPeriodicInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfPeriodicInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfPeriodicInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfPeriodicInterval.EntityData.Children = make(map[string]types.YChild)
    spfPeriodicInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    spfPeriodicInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", spfPeriodicInterval.Level}
    spfPeriodicInterval.EntityData.Leafs["periodic-interval"] = types.YLeaf{"PeriodicInterval", spfPeriodicInterval.PeriodicInterval}
    return &(spfPeriodicInterval.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals
// SPF-interval configuration
type Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route calculation scheduling parameters. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval.
    SpfInterval []Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals) GetEntityData() *types.CommonEntityData {
    spfIntervals.EntityData.YFilter = spfIntervals.YFilter
    spfIntervals.EntityData.YangName = "spf-intervals"
    spfIntervals.EntityData.BundleName = "cisco_ios_xr"
    spfIntervals.EntityData.ParentYangName = "af-data"
    spfIntervals.EntityData.SegmentPath = "spf-intervals"
    spfIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfIntervals.EntityData.Children = make(map[string]types.YChild)
    spfIntervals.EntityData.Children["spf-interval"] = types.YChild{"SpfInterval", nil}
    for i := range spfIntervals.SpfInterval {
        spfIntervals.EntityData.Children[types.GetSegmentPath(&spfIntervals.SpfInterval[i])] = types.YChild{"SpfInterval", &spfIntervals.SpfInterval[i]}
    }
    spfIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spfIntervals.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval
// Route calculation scheduling parameters
type Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval struct {
    EntityData types.CommonEntityData
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

func (spfInterval *Isis_Instances_Instance_Afs_Af_AfData_SpfIntervals_SpfInterval) GetEntityData() *types.CommonEntityData {
    spfInterval.EntityData.YFilter = spfInterval.YFilter
    spfInterval.EntityData.YangName = "spf-interval"
    spfInterval.EntityData.BundleName = "cisco_ios_xr"
    spfInterval.EntityData.ParentYangName = "spf-intervals"
    spfInterval.EntityData.SegmentPath = "spf-interval" + "[level='" + fmt.Sprintf("%v", spfInterval.Level) + "']"
    spfInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfInterval.EntityData.Children = make(map[string]types.YChild)
    spfInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    spfInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", spfInterval.Level}
    spfInterval.EntityData.Leafs["maximum-wait"] = types.YLeaf{"MaximumWait", spfInterval.MaximumWait}
    spfInterval.EntityData.Leafs["initial-wait"] = types.YLeaf{"InitialWait", spfInterval.InitialWait}
    spfInterval.EntityData.Leafs["secondary-wait"] = types.YLeaf{"SecondaryWait", spfInterval.SecondaryWait}
    return &(spfInterval.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence
// Enable convergence monitoring
type Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable convergence monitoring. The type is interface{}.
    Enable interface{}

    // Enable the Tracking of IP-Frr Convergence. The type is interface{}.
    TrackIpFrr interface{}

    // Enable the monitoring of individual prefixes (prefix list name). The type
    // is string with length: 1..32.
    PrefixList interface{}
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_AfData_MonitorConvergence) GetEntityData() *types.CommonEntityData {
    monitorConvergence.EntityData.YFilter = monitorConvergence.YFilter
    monitorConvergence.EntityData.YangName = "monitor-convergence"
    monitorConvergence.EntityData.BundleName = "cisco_ios_xr"
    monitorConvergence.EntityData.ParentYangName = "af-data"
    monitorConvergence.EntityData.SegmentPath = "monitor-convergence"
    monitorConvergence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorConvergence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorConvergence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorConvergence.EntityData.Children = make(map[string]types.YChild)
    monitorConvergence.EntityData.Leafs = make(map[string]types.YLeaf)
    monitorConvergence.EntityData.Leafs["enable"] = types.YLeaf{"Enable", monitorConvergence.Enable}
    monitorConvergence.EntityData.Leafs["track-ip-frr"] = types.YLeaf{"TrackIpFrr", monitorConvergence.TrackIpFrr}
    monitorConvergence.EntityData.Leafs["prefix-list"] = types.YLeaf{"PrefixList", monitorConvergence.PrefixList}
    return &(monitorConvergence.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation
// Control origination of a default route with
// the option of using a policy.  If no policy
// is specified the default route is
// advertised with zero cost in level 2 only.
type Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation struct {
    EntityData types.CommonEntityData
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

func (defaultInformation *Isis_Instances_Instance_Afs_Af_AfData_DefaultInformation) GetEntityData() *types.CommonEntityData {
    defaultInformation.EntityData.YFilter = defaultInformation.YFilter
    defaultInformation.EntityData.YangName = "default-information"
    defaultInformation.EntityData.BundleName = "cisco_ios_xr"
    defaultInformation.EntityData.ParentYangName = "af-data"
    defaultInformation.EntityData.SegmentPath = "default-information"
    defaultInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultInformation.EntityData.Children = make(map[string]types.YChild)
    defaultInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultInformation.EntityData.Leafs["use-policy"] = types.YLeaf{"UsePolicy", defaultInformation.UsePolicy}
    defaultInformation.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", defaultInformation.PolicyName}
    defaultInformation.EntityData.Leafs["external"] = types.YLeaf{"External", defaultInformation.External}
    return &(defaultInformation.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_AdminDistances
// Per-route administrative
// distanceconfiguration
type Isis_Instances_Instance_Afs_Af_AfData_AdminDistances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative distance configuration. The supplied distance is applied to
    // all routes discovered from the specified source, or only those that match
    // the supplied prefix list if this is specified. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance.
    AdminDistance []Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances) GetEntityData() *types.CommonEntityData {
    adminDistances.EntityData.YFilter = adminDistances.YFilter
    adminDistances.EntityData.YangName = "admin-distances"
    adminDistances.EntityData.BundleName = "cisco_ios_xr"
    adminDistances.EntityData.ParentYangName = "af-data"
    adminDistances.EntityData.SegmentPath = "admin-distances"
    adminDistances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminDistances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminDistances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminDistances.EntityData.Children = make(map[string]types.YChild)
    adminDistances.EntityData.Children["admin-distance"] = types.YChild{"AdminDistance", nil}
    for i := range adminDistances.AdminDistance {
        adminDistances.EntityData.Children[types.GetSegmentPath(&adminDistances.AdminDistance[i])] = types.YChild{"AdminDistance", &adminDistances.AdminDistance[i]}
    }
    adminDistances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adminDistances.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance
// Administrative distance configuration. The
// supplied distance is applied to all routes
// discovered from the specified source, or
// only those that match the supplied prefix
// list if this is specified
type Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP route source prefix. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    AddressPrefix interface{}

    // Administrative distance. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    Distance interface{}

    // List of prefixes to which this distance applies. The type is string with
    // length: 1..32.
    PrefixList interface{}
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_AfData_AdminDistances_AdminDistance) GetEntityData() *types.CommonEntityData {
    adminDistance.EntityData.YFilter = adminDistance.YFilter
    adminDistance.EntityData.YangName = "admin-distance"
    adminDistance.EntityData.BundleName = "cisco_ios_xr"
    adminDistance.EntityData.ParentYangName = "admin-distances"
    adminDistance.EntityData.SegmentPath = "admin-distance" + "[address-prefix='" + fmt.Sprintf("%v", adminDistance.AddressPrefix) + "']"
    adminDistance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminDistance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminDistance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminDistance.EntityData.Children = make(map[string]types.YChild)
    adminDistance.EntityData.Leafs = make(map[string]types.YLeaf)
    adminDistance.EntityData.Leafs["address-prefix"] = types.YLeaf{"AddressPrefix", adminDistance.AddressPrefix}
    adminDistance.EntityData.Leafs["distance"] = types.YLeaf{"Distance", adminDistance.Distance}
    adminDistance.EntityData.Leafs["prefix-list"] = types.YLeaf{"PrefixList", adminDistance.PrefixList}
    return &(adminDistance.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Ispf
// ISPF configuration
type Isis_Instances_Instance_Afs_Af_AfData_Ispf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISPF state (enable/disable).
    States Isis_Instances_Instance_Afs_Af_AfData_Ispf_States
}

func (ispf *Isis_Instances_Instance_Afs_Af_AfData_Ispf) GetEntityData() *types.CommonEntityData {
    ispf.EntityData.YFilter = ispf.YFilter
    ispf.EntityData.YangName = "ispf"
    ispf.EntityData.BundleName = "cisco_ios_xr"
    ispf.EntityData.ParentYangName = "af-data"
    ispf.EntityData.SegmentPath = "ispf"
    ispf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ispf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ispf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ispf.EntityData.Children = make(map[string]types.YChild)
    ispf.EntityData.Children["states"] = types.YChild{"States", &ispf.States}
    ispf.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ispf.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Ispf_States
// ISPF state (enable/disable)
type Isis_Instances_Instance_Afs_Af_AfData_Ispf_States struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/disable ISPF. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State.
    State []Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State
}

func (states *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States) GetEntityData() *types.CommonEntityData {
    states.EntityData.YFilter = states.YFilter
    states.EntityData.YangName = "states"
    states.EntityData.BundleName = "cisco_ios_xr"
    states.EntityData.ParentYangName = "ispf"
    states.EntityData.SegmentPath = "states"
    states.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    states.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    states.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    states.EntityData.Children = make(map[string]types.YChild)
    states.EntityData.Children["state"] = types.YChild{"State", nil}
    for i := range states.State {
        states.EntityData.Children[types.GetSegmentPath(&states.State[i])] = types.YChild{"State", &states.State[i]}
    }
    states.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(states.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State
// Enable/disable ISPF
type Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // State. The type is IsisispfState. This attribute is mandatory.
    State interface{}
}

func (state *Isis_Instances_Instance_Afs_Af_AfData_Ispf_States_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "states"
    state.EntityData.SegmentPath = "state" + "[level='" + fmt.Sprintf("%v", state.Level) + "']"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["level"] = types.YLeaf{"Level", state.Level}
    state.EntityData.Leafs["state"] = types.YLeaf{"State", state.State}
    return &(state.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal
// MPLS LDP configuration. MPLS LDP
// configuration will only be applied for the
// IPv4-unicast address-family.
type Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If TRUE, LDP will be enabled onall IS-IS interfaces enabled for this
    // address-family. The type is bool.
    AutoConfig interface{}
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_AfData_MplsLdpGlobal) GetEntityData() *types.CommonEntityData {
    mplsLdpGlobal.EntityData.YFilter = mplsLdpGlobal.YFilter
    mplsLdpGlobal.EntityData.YangName = "mpls-ldp-global"
    mplsLdpGlobal.EntityData.BundleName = "cisco_ios_xr"
    mplsLdpGlobal.EntityData.ParentYangName = "af-data"
    mplsLdpGlobal.EntityData.SegmentPath = "mpls-ldp-global"
    mplsLdpGlobal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLdpGlobal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLdpGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLdpGlobal.EntityData.Children = make(map[string]types.YChild)
    mplsLdpGlobal.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsLdpGlobal.EntityData.Leafs["auto-config"] = types.YLeaf{"AutoConfig", mplsLdpGlobal.AutoConfig}
    return &(mplsLdpGlobal.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Mpls
// MPLS configuration. MPLS configuration will
// only be applied for the IPv4-unicast
// address-family.
type Isis_Instances_Instance_Afs_Af_AfData_Mpls struct {
    EntityData types.CommonEntityData
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

func (mpls *Isis_Instances_Instance_Afs_Af_AfData_Mpls) GetEntityData() *types.CommonEntityData {
    mpls.EntityData.YFilter = mpls.YFilter
    mpls.EntityData.YangName = "mpls"
    mpls.EntityData.BundleName = "cisco_ios_xr"
    mpls.EntityData.ParentYangName = "af-data"
    mpls.EntityData.SegmentPath = "mpls"
    mpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpls.EntityData.Children = make(map[string]types.YChild)
    mpls.EntityData.Children["router-id"] = types.YChild{"RouterId", &mpls.RouterId}
    mpls.EntityData.Children["level"] = types.YChild{"Level", &mpls.Level}
    mpls.EntityData.Leafs = make(map[string]types.YLeaf)
    mpls.EntityData.Leafs["igp-intact"] = types.YLeaf{"IgpIntact", mpls.IgpIntact}
    mpls.EntityData.Leafs["multicast-intact"] = types.YLeaf{"MulticastIntact", mpls.MulticastIntact}
    return &(mpls.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId
// Traffic Engineering stable IP address for
// system
type Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address to be used as a router ID. Precisely one of Address and
    // Interface must be specified. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Interface with designated stable IP address to be used as a router ID. This
    // must be a Loopback interface. Precisely one of Address and Interface must
    // be specified. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (routerId *Isis_Instances_Instance_Afs_Af_AfData_Mpls_RouterId) GetEntityData() *types.CommonEntityData {
    routerId.EntityData.YFilter = routerId.YFilter
    routerId.EntityData.YangName = "router-id"
    routerId.EntityData.BundleName = "cisco_ios_xr"
    routerId.EntityData.ParentYangName = "mpls"
    routerId.EntityData.SegmentPath = "router-id"
    routerId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routerId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routerId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routerId.EntityData.Children = make(map[string]types.YChild)
    routerId.EntityData.Leafs = make(map[string]types.YLeaf)
    routerId.EntityData.Leafs["address"] = types.YLeaf{"Address", routerId.Address}
    routerId.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", routerId.InterfaceName}
    return &(routerId.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level
// Enable MPLS for an IS-IS at the given
// levels
type Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Level 1 enabled. The type is bool.
    Level1 interface{}

    // Level 2 enabled. The type is bool.
    Level2 interface{}
}

func (level *Isis_Instances_Instance_Afs_Af_AfData_Mpls_Level) GetEntityData() *types.CommonEntityData {
    level.EntityData.YFilter = level.YFilter
    level.EntityData.YangName = "level"
    level.EntityData.BundleName = "cisco_ios_xr"
    level.EntityData.ParentYangName = "mpls"
    level.EntityData.SegmentPath = "level"
    level.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    level.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    level.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    level.EntityData.Children = make(map[string]types.YChild)
    level.EntityData.Leafs = make(map[string]types.YLeaf)
    level.EntityData.Leafs["level1"] = types.YLeaf{"Level1", level.Level1}
    level.EntityData.Leafs["level2"] = types.YLeaf{"Level2", level.Level2}
    return &(level.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids
// Manual Adjacecy SID configuration
type Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Assign adjancency SID to an interface. The type is slice of
    // Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid.
    ManualAdjSid []Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids) GetEntityData() *types.CommonEntityData {
    manualAdjSids.EntityData.YFilter = manualAdjSids.YFilter
    manualAdjSids.EntityData.YangName = "manual-adj-sids"
    manualAdjSids.EntityData.BundleName = "cisco_ios_xr"
    manualAdjSids.EntityData.ParentYangName = "af-data"
    manualAdjSids.EntityData.SegmentPath = "manual-adj-sids"
    manualAdjSids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualAdjSids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualAdjSids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualAdjSids.EntityData.Children = make(map[string]types.YChild)
    manualAdjSids.EntityData.Children["manual-adj-sid"] = types.YChild{"ManualAdjSid", nil}
    for i := range manualAdjSids.ManualAdjSid {
        manualAdjSids.EntityData.Children[types.GetSegmentPath(&manualAdjSids.ManualAdjSid[i])] = types.YChild{"ManualAdjSid", &manualAdjSids.ManualAdjSid[i]}
    }
    manualAdjSids.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(manualAdjSids.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid
// Assign adjancency SID to an interface
type Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid struct {
    EntityData types.CommonEntityData
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

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_AfData_ManualAdjSids_ManualAdjSid) GetEntityData() *types.CommonEntityData {
    manualAdjSid.EntityData.YFilter = manualAdjSid.YFilter
    manualAdjSid.EntityData.YangName = "manual-adj-sid"
    manualAdjSid.EntityData.BundleName = "cisco_ios_xr"
    manualAdjSid.EntityData.ParentYangName = "manual-adj-sids"
    manualAdjSid.EntityData.SegmentPath = "manual-adj-sid" + "[level='" + fmt.Sprintf("%v", manualAdjSid.Level) + "']" + "[sid-type='" + fmt.Sprintf("%v", manualAdjSid.SidType) + "']" + "[sid='" + fmt.Sprintf("%v", manualAdjSid.Sid) + "']"
    manualAdjSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualAdjSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualAdjSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualAdjSid.EntityData.Children = make(map[string]types.YChild)
    manualAdjSid.EntityData.Leafs = make(map[string]types.YLeaf)
    manualAdjSid.EntityData.Leafs["level"] = types.YLeaf{"Level", manualAdjSid.Level}
    manualAdjSid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", manualAdjSid.SidType}
    manualAdjSid.EntityData.Leafs["sid"] = types.YLeaf{"Sid", manualAdjSid.Sid}
    manualAdjSid.EntityData.Leafs["protected"] = types.YLeaf{"Protected", manualAdjSid.Protected}
    return &(manualAdjSid.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Metrics
// Metric configuration
type Isis_Instances_Instance_Afs_Af_AfData_Metrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric configuration. Legal value depends on the metric-style specified for
    // the topology. If the metric-style defined is narrow, then only a value
    // between <1-63> is allowed and if the metric-style is defined as wide, then
    // a value between <1-16777215> is allowed as the metric value.  All routers
    // exclude links with the maximum wide metric (16777215) from their SPF. The
    // type is slice of Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric.
    Metric []Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric
}

func (metrics *Isis_Instances_Instance_Afs_Af_AfData_Metrics) GetEntityData() *types.CommonEntityData {
    metrics.EntityData.YFilter = metrics.YFilter
    metrics.EntityData.YangName = "metrics"
    metrics.EntityData.BundleName = "cisco_ios_xr"
    metrics.EntityData.ParentYangName = "af-data"
    metrics.EntityData.SegmentPath = "metrics"
    metrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metrics.EntityData.Children = make(map[string]types.YChild)
    metrics.EntityData.Children["metric"] = types.YChild{"Metric", nil}
    for i := range metrics.Metric {
        metrics.EntityData.Children[types.GetSegmentPath(&metrics.Metric[i])] = types.YChild{"Metric", &metrics.Metric[i]}
    }
    metrics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(metrics.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed metric: <1-63> for narrow, <1-16777215> for wide. The type is one
    // of the following types: enumeration
    // Isis.Instances.Instance.Interfaces.Interface.InterfaceAfs.InterfaceAf.TopologyName.Metrics.Metric.Metric_
    // This attribute is mandatory., or int with range: 1..16777215 This attribute
    // is mandatory..
    Metric interface{}
}

func (metric *Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "metrics"
    metric.EntityData.SegmentPath = "metric" + "[level='" + fmt.Sprintf("%v", metric.Level) + "']"
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = make(map[string]types.YChild)
    metric.EntityData.Leafs = make(map[string]types.YLeaf)
    metric.EntityData.Leafs["level"] = types.YLeaf{"Level", metric.Level}
    metric.EntityData.Leafs["metric"] = types.YLeaf{"Metric", metric.Metric}
    return &(metric.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric_Metric_ represents <1-16777215> for wide
type Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric_Metric_ string

const (
    // Maximum wide metric.  All routers will
    // exclude this link from their SPF
    Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric_Metric__maximum Isis_Instances_Instance_Afs_Af_AfData_Metrics_Metric_Metric_ = "maximum"
)

// Isis_Instances_Instance_Afs_Af_AfData_Weights
// Weight configuration
type Isis_Instances_Instance_Afs_Af_AfData_Weights struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Weight configuration under interface for load balancing. The type is slice
    // of Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight.
    Weight []Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight
}

func (weights *Isis_Instances_Instance_Afs_Af_AfData_Weights) GetEntityData() *types.CommonEntityData {
    weights.EntityData.YFilter = weights.YFilter
    weights.EntityData.YangName = "weights"
    weights.EntityData.BundleName = "cisco_ios_xr"
    weights.EntityData.ParentYangName = "af-data"
    weights.EntityData.SegmentPath = "weights"
    weights.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    weights.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    weights.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    weights.EntityData.Children = make(map[string]types.YChild)
    weights.EntityData.Children["weight"] = types.YChild{"Weight", nil}
    for i := range weights.Weight {
        weights.EntityData.Children[types.GetSegmentPath(&weights.Weight[i])] = types.YChild{"Weight", &weights.Weight[i]}
    }
    weights.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(weights.EntityData)
}

// Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight
// Weight configuration under interface for load
// balancing
type Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Weight to be configured under interface for Load Balancing. Allowed weight:
    // <1-16777215>. The type is interface{} with range: 1..16777214. This
    // attribute is mandatory.
    Weight interface{}
}

func (weight *Isis_Instances_Instance_Afs_Af_AfData_Weights_Weight) GetEntityData() *types.CommonEntityData {
    weight.EntityData.YFilter = weight.YFilter
    weight.EntityData.YangName = "weight"
    weight.EntityData.BundleName = "cisco_ios_xr"
    weight.EntityData.ParentYangName = "weights"
    weight.EntityData.SegmentPath = "weight" + "[level='" + fmt.Sprintf("%v", weight.Level) + "']"
    weight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    weight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    weight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    weight.EntityData.Children = make(map[string]types.YChild)
    weight.EntityData.Leafs = make(map[string]types.YLeaf)
    weight.EntityData.Leafs["level"] = types.YLeaf{"Level", weight.Level}
    weight.EntityData.Leafs["weight"] = types.YLeaf{"Weight", weight.Weight}
    return &(weight.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName
// keys: topology-name
type Isis_Instances_Instance_Afs_Af_TopologyName struct {
    EntityData types.CommonEntityData
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

func (topologyName *Isis_Instances_Instance_Afs_Af_TopologyName) GetEntityData() *types.CommonEntityData {
    topologyName.EntityData.YFilter = topologyName.YFilter
    topologyName.EntityData.YangName = "topology-name"
    topologyName.EntityData.BundleName = "cisco_ios_xr"
    topologyName.EntityData.ParentYangName = "af"
    topologyName.EntityData.SegmentPath = "topology-name" + "[topology-name='" + fmt.Sprintf("%v", topologyName.TopologyName) + "']"
    topologyName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyName.EntityData.Children = make(map[string]types.YChild)
    topologyName.EntityData.Children["segment-routing"] = types.YChild{"SegmentRouting", &topologyName.SegmentRouting}
    topologyName.EntityData.Children["metric-styles"] = types.YChild{"MetricStyles", &topologyName.MetricStyles}
    topologyName.EntityData.Children["frr-table"] = types.YChild{"FrrTable", &topologyName.FrrTable}
    topologyName.EntityData.Children["router-id"] = types.YChild{"RouterId", &topologyName.RouterId}
    topologyName.EntityData.Children["spf-prefix-priorities"] = types.YChild{"SpfPrefixPriorities", &topologyName.SpfPrefixPriorities}
    topologyName.EntityData.Children["summary-prefixes"] = types.YChild{"SummaryPrefixes", &topologyName.SummaryPrefixes}
    topologyName.EntityData.Children["micro-loop-avoidance"] = types.YChild{"MicroLoopAvoidance", &topologyName.MicroLoopAvoidance}
    topologyName.EntityData.Children["ucmp"] = types.YChild{"Ucmp", &topologyName.Ucmp}
    topologyName.EntityData.Children["max-redist-prefixes"] = types.YChild{"MaxRedistPrefixes", &topologyName.MaxRedistPrefixes}
    topologyName.EntityData.Children["propagations"] = types.YChild{"Propagations", &topologyName.Propagations}
    topologyName.EntityData.Children["redistributions"] = types.YChild{"Redistributions", &topologyName.Redistributions}
    topologyName.EntityData.Children["spf-periodic-intervals"] = types.YChild{"SpfPeriodicIntervals", &topologyName.SpfPeriodicIntervals}
    topologyName.EntityData.Children["spf-intervals"] = types.YChild{"SpfIntervals", &topologyName.SpfIntervals}
    topologyName.EntityData.Children["monitor-convergence"] = types.YChild{"MonitorConvergence", &topologyName.MonitorConvergence}
    topologyName.EntityData.Children["default-information"] = types.YChild{"DefaultInformation", &topologyName.DefaultInformation}
    topologyName.EntityData.Children["admin-distances"] = types.YChild{"AdminDistances", &topologyName.AdminDistances}
    topologyName.EntityData.Children["ispf"] = types.YChild{"Ispf", &topologyName.Ispf}
    topologyName.EntityData.Children["mpls-ldp-global"] = types.YChild{"MplsLdpGlobal", &topologyName.MplsLdpGlobal}
    topologyName.EntityData.Children["mpls"] = types.YChild{"Mpls", &topologyName.Mpls}
    topologyName.EntityData.Children["manual-adj-sids"] = types.YChild{"ManualAdjSids", &topologyName.ManualAdjSids}
    topologyName.EntityData.Children["metrics"] = types.YChild{"Metrics", &topologyName.Metrics}
    topologyName.EntityData.Children["weights"] = types.YChild{"Weights", &topologyName.Weights}
    topologyName.EntityData.Leafs = make(map[string]types.YLeaf)
    topologyName.EntityData.Leafs["topology-name"] = types.YLeaf{"TopologyName", topologyName.TopologyName}
    topologyName.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", topologyName.MaximumPaths}
    topologyName.EntityData.Leafs["topology-id"] = types.YLeaf{"TopologyId", topologyName.TopologyId}
    topologyName.EntityData.Leafs["single-topology"] = types.YLeaf{"SingleTopology", topologyName.SingleTopology}
    topologyName.EntityData.Leafs["adjacency-check"] = types.YLeaf{"AdjacencyCheck", topologyName.AdjacencyCheck}
    topologyName.EntityData.Leafs["advertise-link-attributes"] = types.YLeaf{"AdvertiseLinkAttributes", topologyName.AdvertiseLinkAttributes}
    topologyName.EntityData.Leafs["apply-weight"] = types.YLeaf{"ApplyWeight", topologyName.ApplyWeight}
    topologyName.EntityData.Leafs["default-admin-distance"] = types.YLeaf{"DefaultAdminDistance", topologyName.DefaultAdminDistance}
    topologyName.EntityData.Leafs["advertise-passive-only"] = types.YLeaf{"AdvertisePassiveOnly", topologyName.AdvertisePassiveOnly}
    topologyName.EntityData.Leafs["ignore-attached-bit"] = types.YLeaf{"IgnoreAttachedBit", topologyName.IgnoreAttachedBit}
    topologyName.EntityData.Leafs["attached-bit"] = types.YLeaf{"AttachedBit", topologyName.AttachedBit}
    topologyName.EntityData.Leafs["route-source-first-hop"] = types.YLeaf{"RouteSourceFirstHop", topologyName.RouteSourceFirstHop}
    return &(topologyName.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting
// Enable Segment Routing configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable per bundle member adjacency SID. The type is interface{}.
    BundleMemberAdjSid interface{}

    // Prefer segment routing labels over LDP labels. The type is
    // IsisLabelPreference.
    Mpls interface{}

    // Enable Segment Routing prefix SID map configuration.
    PrefixSidMap Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap
}

func (segmentRouting *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting) GetEntityData() *types.CommonEntityData {
    segmentRouting.EntityData.YFilter = segmentRouting.YFilter
    segmentRouting.EntityData.YangName = "segment-routing"
    segmentRouting.EntityData.BundleName = "cisco_ios_xr"
    segmentRouting.EntityData.ParentYangName = "topology-name"
    segmentRouting.EntityData.SegmentPath = "segment-routing"
    segmentRouting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRouting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRouting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRouting.EntityData.Children = make(map[string]types.YChild)
    segmentRouting.EntityData.Children["prefix-sid-map"] = types.YChild{"PrefixSidMap", &segmentRouting.PrefixSidMap}
    segmentRouting.EntityData.Leafs = make(map[string]types.YLeaf)
    segmentRouting.EntityData.Leafs["bundle-member-adj-sid"] = types.YLeaf{"BundleMemberAdjSid", segmentRouting.BundleMemberAdjSid}
    segmentRouting.EntityData.Leafs["mpls"] = types.YLeaf{"Mpls", segmentRouting.Mpls}
    return &(segmentRouting.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap
// Enable Segment Routing prefix SID map
// configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Segment Routing prefix SID map advertise local. The type is
    // interface{}.
    AdvertiseLocal interface{}

    // If TRUE, remote prefix SID map advertisements will be used. If FALSE, they
    // will not be used. The type is bool.
    Receive interface{}
}

func (prefixSidMap *Isis_Instances_Instance_Afs_Af_TopologyName_SegmentRouting_PrefixSidMap) GetEntityData() *types.CommonEntityData {
    prefixSidMap.EntityData.YFilter = prefixSidMap.YFilter
    prefixSidMap.EntityData.YangName = "prefix-sid-map"
    prefixSidMap.EntityData.BundleName = "cisco_ios_xr"
    prefixSidMap.EntityData.ParentYangName = "segment-routing"
    prefixSidMap.EntityData.SegmentPath = "prefix-sid-map"
    prefixSidMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSidMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSidMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSidMap.EntityData.Children = make(map[string]types.YChild)
    prefixSidMap.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixSidMap.EntityData.Leafs["advertise-local"] = types.YLeaf{"AdvertiseLocal", prefixSidMap.AdvertiseLocal}
    prefixSidMap.EntityData.Leafs["receive"] = types.YLeaf{"Receive", prefixSidMap.Receive}
    return &(prefixSidMap.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles
// Metric-style configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration of metric style in LSPs. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle.
    MetricStyle []Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle
}

func (metricStyles *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles) GetEntityData() *types.CommonEntityData {
    metricStyles.EntityData.YFilter = metricStyles.YFilter
    metricStyles.EntityData.YangName = "metric-styles"
    metricStyles.EntityData.BundleName = "cisco_ios_xr"
    metricStyles.EntityData.ParentYangName = "topology-name"
    metricStyles.EntityData.SegmentPath = "metric-styles"
    metricStyles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metricStyles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metricStyles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metricStyles.EntityData.Children = make(map[string]types.YChild)
    metricStyles.EntityData.Children["metric-style"] = types.YChild{"MetricStyle", nil}
    for i := range metricStyles.MetricStyle {
        metricStyles.EntityData.Children[types.GetSegmentPath(&metricStyles.MetricStyle[i])] = types.YChild{"MetricStyle", &metricStyles.MetricStyle[i]}
    }
    metricStyles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(metricStyles.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle
// Configuration of metric style in LSPs
type Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle struct {
    EntityData types.CommonEntityData
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

func (metricStyle *Isis_Instances_Instance_Afs_Af_TopologyName_MetricStyles_MetricStyle) GetEntityData() *types.CommonEntityData {
    metricStyle.EntityData.YFilter = metricStyle.YFilter
    metricStyle.EntityData.YangName = "metric-style"
    metricStyle.EntityData.BundleName = "cisco_ios_xr"
    metricStyle.EntityData.ParentYangName = "metric-styles"
    metricStyle.EntityData.SegmentPath = "metric-style" + "[level='" + fmt.Sprintf("%v", metricStyle.Level) + "']"
    metricStyle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metricStyle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metricStyle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metricStyle.EntityData.Children = make(map[string]types.YChild)
    metricStyle.EntityData.Leafs = make(map[string]types.YLeaf)
    metricStyle.EntityData.Leafs["level"] = types.YLeaf{"Level", metricStyle.Level}
    metricStyle.EntityData.Leafs["style"] = types.YLeaf{"Style", metricStyle.Style}
    metricStyle.EntityData.Leafs["transition-state"] = types.YLeaf{"TransitionState", metricStyle.TransitionState}
    return &(metricStyle.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable
// Fast-ReRoute configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable struct {
    EntityData types.CommonEntityData
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

func (frrTable *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable) GetEntityData() *types.CommonEntityData {
    frrTable.EntityData.YFilter = frrTable.YFilter
    frrTable.EntityData.YangName = "frr-table"
    frrTable.EntityData.BundleName = "cisco_ios_xr"
    frrTable.EntityData.ParentYangName = "topology-name"
    frrTable.EntityData.SegmentPath = "frr-table"
    frrTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrTable.EntityData.Children = make(map[string]types.YChild)
    frrTable.EntityData.Children["frr-load-sharings"] = types.YChild{"FrrLoadSharings", &frrTable.FrrLoadSharings}
    frrTable.EntityData.Children["priority-limits"] = types.YChild{"PriorityLimits", &frrTable.PriorityLimits}
    frrTable.EntityData.Children["frr-remote-lfa-prefixes"] = types.YChild{"FrrRemoteLfaPrefixes", &frrTable.FrrRemoteLfaPrefixes}
    frrTable.EntityData.Children["frr-tiebreakers"] = types.YChild{"FrrTiebreakers", &frrTable.FrrTiebreakers}
    frrTable.EntityData.Children["frr-use-cand-onlies"] = types.YChild{"FrrUseCandOnlies", &frrTable.FrrUseCandOnlies}
    frrTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrTable.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings
// Load share prefixes across multiple
// backups
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable load sharing. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing.
    FrrLoadSharing []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing
}

func (frrLoadSharings *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings) GetEntityData() *types.CommonEntityData {
    frrLoadSharings.EntityData.YFilter = frrLoadSharings.YFilter
    frrLoadSharings.EntityData.YangName = "frr-load-sharings"
    frrLoadSharings.EntityData.BundleName = "cisco_ios_xr"
    frrLoadSharings.EntityData.ParentYangName = "frr-table"
    frrLoadSharings.EntityData.SegmentPath = "frr-load-sharings"
    frrLoadSharings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrLoadSharings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrLoadSharings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrLoadSharings.EntityData.Children = make(map[string]types.YChild)
    frrLoadSharings.EntityData.Children["frr-load-sharing"] = types.YChild{"FrrLoadSharing", nil}
    for i := range frrLoadSharings.FrrLoadSharing {
        frrLoadSharings.EntityData.Children[types.GetSegmentPath(&frrLoadSharings.FrrLoadSharing[i])] = types.YChild{"FrrLoadSharing", &frrLoadSharings.FrrLoadSharing[i]}
    }
    frrLoadSharings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrLoadSharings.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing
// Disable load sharing
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Load sharing. The type is IsisfrrLoadSharing. This attribute is mandatory.
    LoadSharing interface{}
}

func (frrLoadSharing *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrLoadSharings_FrrLoadSharing) GetEntityData() *types.CommonEntityData {
    frrLoadSharing.EntityData.YFilter = frrLoadSharing.YFilter
    frrLoadSharing.EntityData.YangName = "frr-load-sharing"
    frrLoadSharing.EntityData.BundleName = "cisco_ios_xr"
    frrLoadSharing.EntityData.ParentYangName = "frr-load-sharings"
    frrLoadSharing.EntityData.SegmentPath = "frr-load-sharing" + "[level='" + fmt.Sprintf("%v", frrLoadSharing.Level) + "']"
    frrLoadSharing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrLoadSharing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrLoadSharing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrLoadSharing.EntityData.Children = make(map[string]types.YChild)
    frrLoadSharing.EntityData.Leafs = make(map[string]types.YLeaf)
    frrLoadSharing.EntityData.Leafs["level"] = types.YLeaf{"Level", frrLoadSharing.Level}
    frrLoadSharing.EntityData.Leafs["load-sharing"] = types.YLeaf{"LoadSharing", frrLoadSharing.LoadSharing}
    return &(frrLoadSharing.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits
// FRR prefix-limit configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Limit backup computation upto the prefix priority. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit.
    PriorityLimit []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit
}

func (priorityLimits *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits) GetEntityData() *types.CommonEntityData {
    priorityLimits.EntityData.YFilter = priorityLimits.YFilter
    priorityLimits.EntityData.YangName = "priority-limits"
    priorityLimits.EntityData.BundleName = "cisco_ios_xr"
    priorityLimits.EntityData.ParentYangName = "frr-table"
    priorityLimits.EntityData.SegmentPath = "priority-limits"
    priorityLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priorityLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priorityLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priorityLimits.EntityData.Children = make(map[string]types.YChild)
    priorityLimits.EntityData.Children["priority-limit"] = types.YChild{"PriorityLimit", nil}
    for i := range priorityLimits.PriorityLimit {
        priorityLimits.EntityData.Children[types.GetSegmentPath(&priorityLimits.PriorityLimit[i])] = types.YChild{"PriorityLimit", &priorityLimits.PriorityLimit[i]}
    }
    priorityLimits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(priorityLimits.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit
// Limit backup computation upto the prefix
// priority
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit struct {
    EntityData types.CommonEntityData
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

func (priorityLimit *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_PriorityLimits_PriorityLimit) GetEntityData() *types.CommonEntityData {
    priorityLimit.EntityData.YFilter = priorityLimit.YFilter
    priorityLimit.EntityData.YangName = "priority-limit"
    priorityLimit.EntityData.BundleName = "cisco_ios_xr"
    priorityLimit.EntityData.ParentYangName = "priority-limits"
    priorityLimit.EntityData.SegmentPath = "priority-limit" + "[level='" + fmt.Sprintf("%v", priorityLimit.Level) + "']" + "[frr-type='" + fmt.Sprintf("%v", priorityLimit.FrrType) + "']"
    priorityLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priorityLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priorityLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priorityLimit.EntityData.Children = make(map[string]types.YChild)
    priorityLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    priorityLimit.EntityData.Leafs["level"] = types.YLeaf{"Level", priorityLimit.Level}
    priorityLimit.EntityData.Leafs["frr-type"] = types.YLeaf{"FrrType", priorityLimit.FrrType}
    priorityLimit.EntityData.Leafs["priority"] = types.YLeaf{"Priority", priorityLimit.Priority}
    return &(priorityLimit.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes
// FRR remote LFA prefix list filter
// configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter remote LFA router IDs using prefix-list. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix.
    FrrRemoteLfaPrefix []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix
}

func (frrRemoteLfaPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaPrefixes.EntityData.YFilter = frrRemoteLfaPrefixes.YFilter
    frrRemoteLfaPrefixes.EntityData.YangName = "frr-remote-lfa-prefixes"
    frrRemoteLfaPrefixes.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaPrefixes.EntityData.ParentYangName = "frr-table"
    frrRemoteLfaPrefixes.EntityData.SegmentPath = "frr-remote-lfa-prefixes"
    frrRemoteLfaPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaPrefixes.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaPrefixes.EntityData.Children["frr-remote-lfa-prefix"] = types.YChild{"FrrRemoteLfaPrefix", nil}
    for i := range frrRemoteLfaPrefixes.FrrRemoteLfaPrefix {
        frrRemoteLfaPrefixes.EntityData.Children[types.GetSegmentPath(&frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[i])] = types.YChild{"FrrRemoteLfaPrefix", &frrRemoteLfaPrefixes.FrrRemoteLfaPrefix[i]}
    }
    frrRemoteLfaPrefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrRemoteLfaPrefixes.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix
// Filter remote LFA router IDs using
// prefix-list
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Name of the prefix list. The type is string with length: 1..32. This
    // attribute is mandatory.
    PrefixListName interface{}
}

func (frrRemoteLfaPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrRemoteLfaPrefixes_FrrRemoteLfaPrefix) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaPrefix.EntityData.YFilter = frrRemoteLfaPrefix.YFilter
    frrRemoteLfaPrefix.EntityData.YangName = "frr-remote-lfa-prefix"
    frrRemoteLfaPrefix.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaPrefix.EntityData.ParentYangName = "frr-remote-lfa-prefixes"
    frrRemoteLfaPrefix.EntityData.SegmentPath = "frr-remote-lfa-prefix" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaPrefix.Level) + "']"
    frrRemoteLfaPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaPrefix.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    frrRemoteLfaPrefix.EntityData.Leafs["level"] = types.YLeaf{"Level", frrRemoteLfaPrefix.Level}
    frrRemoteLfaPrefix.EntityData.Leafs["prefix-list-name"] = types.YLeaf{"PrefixListName", frrRemoteLfaPrefix.PrefixListName}
    return &(frrRemoteLfaPrefix.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers
// FRR tiebreakers configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure tiebreaker for multiple backups. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker.
    FrrTiebreaker []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker
}

func (frrTiebreakers *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers) GetEntityData() *types.CommonEntityData {
    frrTiebreakers.EntityData.YFilter = frrTiebreakers.YFilter
    frrTiebreakers.EntityData.YangName = "frr-tiebreakers"
    frrTiebreakers.EntityData.BundleName = "cisco_ios_xr"
    frrTiebreakers.EntityData.ParentYangName = "frr-table"
    frrTiebreakers.EntityData.SegmentPath = "frr-tiebreakers"
    frrTiebreakers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrTiebreakers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrTiebreakers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrTiebreakers.EntityData.Children = make(map[string]types.YChild)
    frrTiebreakers.EntityData.Children["frr-tiebreaker"] = types.YChild{"FrrTiebreaker", nil}
    for i := range frrTiebreakers.FrrTiebreaker {
        frrTiebreakers.EntityData.Children[types.GetSegmentPath(&frrTiebreakers.FrrTiebreaker[i])] = types.YChild{"FrrTiebreaker", &frrTiebreakers.FrrTiebreaker[i]}
    }
    frrTiebreakers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrTiebreakers.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker
// Configure tiebreaker for multiple backups
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker struct {
    EntityData types.CommonEntityData
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

func (frrTiebreaker *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrTiebreakers_FrrTiebreaker) GetEntityData() *types.CommonEntityData {
    frrTiebreaker.EntityData.YFilter = frrTiebreaker.YFilter
    frrTiebreaker.EntityData.YangName = "frr-tiebreaker"
    frrTiebreaker.EntityData.BundleName = "cisco_ios_xr"
    frrTiebreaker.EntityData.ParentYangName = "frr-tiebreakers"
    frrTiebreaker.EntityData.SegmentPath = "frr-tiebreaker" + "[level='" + fmt.Sprintf("%v", frrTiebreaker.Level) + "']" + "[tiebreaker='" + fmt.Sprintf("%v", frrTiebreaker.Tiebreaker) + "']"
    frrTiebreaker.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrTiebreaker.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrTiebreaker.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrTiebreaker.EntityData.Children = make(map[string]types.YChild)
    frrTiebreaker.EntityData.Leafs = make(map[string]types.YLeaf)
    frrTiebreaker.EntityData.Leafs["level"] = types.YLeaf{"Level", frrTiebreaker.Level}
    frrTiebreaker.EntityData.Leafs["tiebreaker"] = types.YLeaf{"Tiebreaker", frrTiebreaker.Tiebreaker}
    frrTiebreaker.EntityData.Leafs["index"] = types.YLeaf{"Index", frrTiebreaker.Index}
    return &(frrTiebreaker.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies
// FRR use candidate only configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure use candidate only to exclude interfaces as backup. The type is
    // slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly.
    FrrUseCandOnly []Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly
}

func (frrUseCandOnlies *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies) GetEntityData() *types.CommonEntityData {
    frrUseCandOnlies.EntityData.YFilter = frrUseCandOnlies.YFilter
    frrUseCandOnlies.EntityData.YangName = "frr-use-cand-onlies"
    frrUseCandOnlies.EntityData.BundleName = "cisco_ios_xr"
    frrUseCandOnlies.EntityData.ParentYangName = "frr-table"
    frrUseCandOnlies.EntityData.SegmentPath = "frr-use-cand-onlies"
    frrUseCandOnlies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrUseCandOnlies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrUseCandOnlies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrUseCandOnlies.EntityData.Children = make(map[string]types.YChild)
    frrUseCandOnlies.EntityData.Children["frr-use-cand-only"] = types.YChild{"FrrUseCandOnly", nil}
    for i := range frrUseCandOnlies.FrrUseCandOnly {
        frrUseCandOnlies.EntityData.Children[types.GetSegmentPath(&frrUseCandOnlies.FrrUseCandOnly[i])] = types.YChild{"FrrUseCandOnly", &frrUseCandOnlies.FrrUseCandOnly[i]}
    }
    frrUseCandOnlies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrUseCandOnlies.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly
// Configure use candidate only to exclude
// interfaces as backup
type Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}
}

func (frrUseCandOnly *Isis_Instances_Instance_Afs_Af_TopologyName_FrrTable_FrrUseCandOnlies_FrrUseCandOnly) GetEntityData() *types.CommonEntityData {
    frrUseCandOnly.EntityData.YFilter = frrUseCandOnly.YFilter
    frrUseCandOnly.EntityData.YangName = "frr-use-cand-only"
    frrUseCandOnly.EntityData.BundleName = "cisco_ios_xr"
    frrUseCandOnly.EntityData.ParentYangName = "frr-use-cand-onlies"
    frrUseCandOnly.EntityData.SegmentPath = "frr-use-cand-only" + "[level='" + fmt.Sprintf("%v", frrUseCandOnly.Level) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrUseCandOnly.FrrType) + "']"
    frrUseCandOnly.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrUseCandOnly.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrUseCandOnly.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrUseCandOnly.EntityData.Children = make(map[string]types.YChild)
    frrUseCandOnly.EntityData.Leafs = make(map[string]types.YLeaf)
    frrUseCandOnly.EntityData.Leafs["level"] = types.YLeaf{"Level", frrUseCandOnly.Level}
    frrUseCandOnly.EntityData.Leafs["frr-type"] = types.YLeaf{"FrrType", frrUseCandOnly.FrrType}
    return &(frrUseCandOnly.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_RouterId
// Stable IP address for system. Will only be
// applied for the unicast sub-address-family.
type Isis_Instances_Instance_Afs_Af_TopologyName_RouterId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4/IPv6 address to be used as a router ID. Precisely one of Address and
    // Interface must be specified. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Interface with designated stable IP address to be used as a router ID. This
    // must be a Loopback interface. Precisely one of Address and Interface must
    // be specified. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_RouterId) GetEntityData() *types.CommonEntityData {
    routerId.EntityData.YFilter = routerId.YFilter
    routerId.EntityData.YangName = "router-id"
    routerId.EntityData.BundleName = "cisco_ios_xr"
    routerId.EntityData.ParentYangName = "topology-name"
    routerId.EntityData.SegmentPath = "router-id"
    routerId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routerId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routerId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routerId.EntityData.Children = make(map[string]types.YChild)
    routerId.EntityData.Leafs = make(map[string]types.YLeaf)
    routerId.EntityData.Leafs["address"] = types.YLeaf{"Address", routerId.Address}
    routerId.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", routerId.InterfaceName}
    return &(routerId.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities
// SPF Prefix Priority configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Determine SPF priority for prefixes. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority.
    SpfPrefixPriority []Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority
}

func (spfPrefixPriorities *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities) GetEntityData() *types.CommonEntityData {
    spfPrefixPriorities.EntityData.YFilter = spfPrefixPriorities.YFilter
    spfPrefixPriorities.EntityData.YangName = "spf-prefix-priorities"
    spfPrefixPriorities.EntityData.BundleName = "cisco_ios_xr"
    spfPrefixPriorities.EntityData.ParentYangName = "topology-name"
    spfPrefixPriorities.EntityData.SegmentPath = "spf-prefix-priorities"
    spfPrefixPriorities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfPrefixPriorities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfPrefixPriorities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfPrefixPriorities.EntityData.Children = make(map[string]types.YChild)
    spfPrefixPriorities.EntityData.Children["spf-prefix-priority"] = types.YChild{"SpfPrefixPriority", nil}
    for i := range spfPrefixPriorities.SpfPrefixPriority {
        spfPrefixPriorities.EntityData.Children[types.GetSegmentPath(&spfPrefixPriorities.SpfPrefixPriority[i])] = types.YChild{"SpfPrefixPriority", &spfPrefixPriorities.SpfPrefixPriority[i]}
    }
    spfPrefixPriorities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spfPrefixPriorities.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority
// Determine SPF priority for prefixes
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority struct {
    EntityData types.CommonEntityData
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

func (spfPrefixPriority *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPrefixPriorities_SpfPrefixPriority) GetEntityData() *types.CommonEntityData {
    spfPrefixPriority.EntityData.YFilter = spfPrefixPriority.YFilter
    spfPrefixPriority.EntityData.YangName = "spf-prefix-priority"
    spfPrefixPriority.EntityData.BundleName = "cisco_ios_xr"
    spfPrefixPriority.EntityData.ParentYangName = "spf-prefix-priorities"
    spfPrefixPriority.EntityData.SegmentPath = "spf-prefix-priority" + "[level='" + fmt.Sprintf("%v", spfPrefixPriority.Level) + "']" + "[prefix-priority-type='" + fmt.Sprintf("%v", spfPrefixPriority.PrefixPriorityType) + "']"
    spfPrefixPriority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfPrefixPriority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfPrefixPriority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfPrefixPriority.EntityData.Children = make(map[string]types.YChild)
    spfPrefixPriority.EntityData.Leafs = make(map[string]types.YLeaf)
    spfPrefixPriority.EntityData.Leafs["level"] = types.YLeaf{"Level", spfPrefixPriority.Level}
    spfPrefixPriority.EntityData.Leafs["prefix-priority-type"] = types.YLeaf{"PrefixPriorityType", spfPrefixPriority.PrefixPriorityType}
    spfPrefixPriority.EntityData.Leafs["admin-tag"] = types.YLeaf{"AdminTag", spfPrefixPriority.AdminTag}
    spfPrefixPriority.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", spfPrefixPriority.AccessListName}
    return &(spfPrefixPriority.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes
// Summary-prefix configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IP address prefixes to advertise. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix.
    SummaryPrefix []Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix
}

func (summaryPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes) GetEntityData() *types.CommonEntityData {
    summaryPrefixes.EntityData.YFilter = summaryPrefixes.YFilter
    summaryPrefixes.EntityData.YangName = "summary-prefixes"
    summaryPrefixes.EntityData.BundleName = "cisco_ios_xr"
    summaryPrefixes.EntityData.ParentYangName = "topology-name"
    summaryPrefixes.EntityData.SegmentPath = "summary-prefixes"
    summaryPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPrefixes.EntityData.Children = make(map[string]types.YChild)
    summaryPrefixes.EntityData.Children["summary-prefix"] = types.YChild{"SummaryPrefix", nil}
    for i := range summaryPrefixes.SummaryPrefix {
        summaryPrefixes.EntityData.Children[types.GetSegmentPath(&summaryPrefixes.SummaryPrefix[i])] = types.YChild{"SummaryPrefix", &summaryPrefixes.SummaryPrefix[i]}
    }
    summaryPrefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summaryPrefixes.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix
// Configure IP address prefixes to advertise
type Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP summary address prefix. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    AddressPrefix interface{}

    // The tag value. The type is interface{} with range: 1..4294967295.
    Tag interface{}

    // Level in which to summarize routes. The type is interface{} with range:
    // 1..2.
    Level interface{}
}

func (summaryPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_SummaryPrefixes_SummaryPrefix) GetEntityData() *types.CommonEntityData {
    summaryPrefix.EntityData.YFilter = summaryPrefix.YFilter
    summaryPrefix.EntityData.YangName = "summary-prefix"
    summaryPrefix.EntityData.BundleName = "cisco_ios_xr"
    summaryPrefix.EntityData.ParentYangName = "summary-prefixes"
    summaryPrefix.EntityData.SegmentPath = "summary-prefix" + "[address-prefix='" + fmt.Sprintf("%v", summaryPrefix.AddressPrefix) + "']"
    summaryPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPrefix.EntityData.Children = make(map[string]types.YChild)
    summaryPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    summaryPrefix.EntityData.Leafs["address-prefix"] = types.YLeaf{"AddressPrefix", summaryPrefix.AddressPrefix}
    summaryPrefix.EntityData.Leafs["tag"] = types.YLeaf{"Tag", summaryPrefix.Tag}
    summaryPrefix.EntityData.Leafs["level"] = types.YLeaf{"Level", summaryPrefix.Level}
    return &(summaryPrefix.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance
// Micro Loop Avoidance configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MicroLoop avoidance enable configuration. The type is
    // IsisMicroLoopAvoidance.
    Enable interface{}

    // Value of delay in msecs in updating RIB. The type is interface{} with
    // range: 1000..65535. Units are millisecond. The default value is 5000.
    RibUpdateDelay interface{}
}

func (microLoopAvoidance *Isis_Instances_Instance_Afs_Af_TopologyName_MicroLoopAvoidance) GetEntityData() *types.CommonEntityData {
    microLoopAvoidance.EntityData.YFilter = microLoopAvoidance.YFilter
    microLoopAvoidance.EntityData.YangName = "micro-loop-avoidance"
    microLoopAvoidance.EntityData.BundleName = "cisco_ios_xr"
    microLoopAvoidance.EntityData.ParentYangName = "topology-name"
    microLoopAvoidance.EntityData.SegmentPath = "micro-loop-avoidance"
    microLoopAvoidance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    microLoopAvoidance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    microLoopAvoidance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    microLoopAvoidance.EntityData.Children = make(map[string]types.YChild)
    microLoopAvoidance.EntityData.Leafs = make(map[string]types.YLeaf)
    microLoopAvoidance.EntityData.Leafs["enable"] = types.YLeaf{"Enable", microLoopAvoidance.Enable}
    microLoopAvoidance.EntityData.Leafs["rib-update-delay"] = types.YLeaf{"RibUpdateDelay", microLoopAvoidance.RibUpdateDelay}
    return &(microLoopAvoidance.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp
// UCMP (UnEqual Cost MultiPath) configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp struct {
    EntityData types.CommonEntityData
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

func (ucmp *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp) GetEntityData() *types.CommonEntityData {
    ucmp.EntityData.YFilter = ucmp.YFilter
    ucmp.EntityData.YangName = "ucmp"
    ucmp.EntityData.BundleName = "cisco_ios_xr"
    ucmp.EntityData.ParentYangName = "topology-name"
    ucmp.EntityData.SegmentPath = "ucmp"
    ucmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ucmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ucmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ucmp.EntityData.Children = make(map[string]types.YChild)
    ucmp.EntityData.Children["enable"] = types.YChild{"Enable", &ucmp.Enable}
    ucmp.EntityData.Children["exclude-interfaces"] = types.YChild{"ExcludeInterfaces", &ucmp.ExcludeInterfaces}
    ucmp.EntityData.Leafs = make(map[string]types.YLeaf)
    ucmp.EntityData.Leafs["delay-interval"] = types.YLeaf{"DelayInterval", ucmp.DelayInterval}
    return &(ucmp.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable
// UCMP feature enable configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Value of variance. The type is interface{} with range: 101..10000. The
    // default value is 200.
    Variance interface{}

    // Name of the Prefix List. The type is string with length: 1..32.
    PrefixListName interface{}
}

func (enable *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_Enable) GetEntityData() *types.CommonEntityData {
    enable.EntityData.YFilter = enable.YFilter
    enable.EntityData.YangName = "enable"
    enable.EntityData.BundleName = "cisco_ios_xr"
    enable.EntityData.ParentYangName = "ucmp"
    enable.EntityData.SegmentPath = "enable"
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = make(map[string]types.YChild)
    enable.EntityData.Leafs = make(map[string]types.YLeaf)
    enable.EntityData.Leafs["variance"] = types.YLeaf{"Variance", enable.Variance}
    enable.EntityData.Leafs["prefix-list-name"] = types.YLeaf{"PrefixListName", enable.PrefixListName}
    return &(enable.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces
// Interfaces excluded from UCMP path
// computation
type Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude this interface from UCMP path computation. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "ucmp"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = make(map[string]types.YChild)
    excludeInterfaces.EntityData.Children["exclude-interface"] = types.YChild{"ExcludeInterface", nil}
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children[types.GetSegmentPath(&excludeInterfaces.ExcludeInterface[i])] = types.YChild{"ExcludeInterface", &excludeInterfaces.ExcludeInterface[i]}
    }
    excludeInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(excludeInterfaces.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface
// Exclude this interface from UCMP path
// computation
type Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the interface to be excluded. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Isis_Instances_Instance_Afs_Af_TopologyName_Ucmp_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + "[interface-name='" + fmt.Sprintf("%v", excludeInterface.InterfaceName) + "']"
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = make(map[string]types.YChild)
    excludeInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    excludeInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", excludeInterface.InterfaceName}
    return &(excludeInterface.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes
// Maximum number of redistributed
// prefixesconfiguration
type Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An upper limit on the number of redistributed prefixes which may be
    // included in the local system's LSP. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix.
    MaxRedistPrefix []Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix
}

func (maxRedistPrefixes *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes) GetEntityData() *types.CommonEntityData {
    maxRedistPrefixes.EntityData.YFilter = maxRedistPrefixes.YFilter
    maxRedistPrefixes.EntityData.YangName = "max-redist-prefixes"
    maxRedistPrefixes.EntityData.BundleName = "cisco_ios_xr"
    maxRedistPrefixes.EntityData.ParentYangName = "topology-name"
    maxRedistPrefixes.EntityData.SegmentPath = "max-redist-prefixes"
    maxRedistPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxRedistPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxRedistPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxRedistPrefixes.EntityData.Children = make(map[string]types.YChild)
    maxRedistPrefixes.EntityData.Children["max-redist-prefix"] = types.YChild{"MaxRedistPrefix", nil}
    for i := range maxRedistPrefixes.MaxRedistPrefix {
        maxRedistPrefixes.EntityData.Children[types.GetSegmentPath(&maxRedistPrefixes.MaxRedistPrefix[i])] = types.YChild{"MaxRedistPrefix", &maxRedistPrefixes.MaxRedistPrefix[i]}
    }
    maxRedistPrefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(maxRedistPrefixes.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix
// An upper limit on the number of
// redistributed prefixes which may be
// included in the local system's LSP
type Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Max number of prefixes. The type is interface{} with range: 1..28000. This
    // attribute is mandatory.
    PrefixLimit interface{}
}

func (maxRedistPrefix *Isis_Instances_Instance_Afs_Af_TopologyName_MaxRedistPrefixes_MaxRedistPrefix) GetEntityData() *types.CommonEntityData {
    maxRedistPrefix.EntityData.YFilter = maxRedistPrefix.YFilter
    maxRedistPrefix.EntityData.YangName = "max-redist-prefix"
    maxRedistPrefix.EntityData.BundleName = "cisco_ios_xr"
    maxRedistPrefix.EntityData.ParentYangName = "max-redist-prefixes"
    maxRedistPrefix.EntityData.SegmentPath = "max-redist-prefix" + "[level='" + fmt.Sprintf("%v", maxRedistPrefix.Level) + "']"
    maxRedistPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxRedistPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxRedistPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxRedistPrefix.EntityData.Children = make(map[string]types.YChild)
    maxRedistPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    maxRedistPrefix.EntityData.Leafs["level"] = types.YLeaf{"Level", maxRedistPrefix.Level}
    maxRedistPrefix.EntityData.Leafs["prefix-limit"] = types.YLeaf{"PrefixLimit", maxRedistPrefix.PrefixLimit}
    return &(maxRedistPrefix.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Propagations
// Route propagation configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Propagations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Propagate routes between IS-IS levels. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation.
    Propagation []Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation
}

func (propagations *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations) GetEntityData() *types.CommonEntityData {
    propagations.EntityData.YFilter = propagations.YFilter
    propagations.EntityData.YangName = "propagations"
    propagations.EntityData.BundleName = "cisco_ios_xr"
    propagations.EntityData.ParentYangName = "topology-name"
    propagations.EntityData.SegmentPath = "propagations"
    propagations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    propagations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    propagations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    propagations.EntityData.Children = make(map[string]types.YChild)
    propagations.EntityData.Children["propagation"] = types.YChild{"Propagation", nil}
    for i := range propagations.Propagation {
        propagations.EntityData.Children[types.GetSegmentPath(&propagations.Propagation[i])] = types.YChild{"Propagation", &propagations.Propagation[i]}
    }
    propagations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(propagations.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation
// Propagate routes between IS-IS levels
type Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation struct {
    EntityData types.CommonEntityData
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

func (propagation *Isis_Instances_Instance_Afs_Af_TopologyName_Propagations_Propagation) GetEntityData() *types.CommonEntityData {
    propagation.EntityData.YFilter = propagation.YFilter
    propagation.EntityData.YangName = "propagation"
    propagation.EntityData.BundleName = "cisco_ios_xr"
    propagation.EntityData.ParentYangName = "propagations"
    propagation.EntityData.SegmentPath = "propagation" + "[source-level='" + fmt.Sprintf("%v", propagation.SourceLevel) + "']" + "[destination-level='" + fmt.Sprintf("%v", propagation.DestinationLevel) + "']"
    propagation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    propagation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    propagation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    propagation.EntityData.Children = make(map[string]types.YChild)
    propagation.EntityData.Leafs = make(map[string]types.YLeaf)
    propagation.EntityData.Leafs["source-level"] = types.YLeaf{"SourceLevel", propagation.SourceLevel}
    propagation.EntityData.Leafs["destination-level"] = types.YLeaf{"DestinationLevel", propagation.DestinationLevel}
    propagation.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", propagation.RoutePolicyName}
    return &(propagation.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions
// Protocol redistribution configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribution of other protocols into this IS-IS instance. The type is
    // slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution.
    Redistribution []Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution
}

func (redistributions *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions) GetEntityData() *types.CommonEntityData {
    redistributions.EntityData.YFilter = redistributions.YFilter
    redistributions.EntityData.YangName = "redistributions"
    redistributions.EntityData.BundleName = "cisco_ios_xr"
    redistributions.EntityData.ParentYangName = "topology-name"
    redistributions.EntityData.SegmentPath = "redistributions"
    redistributions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributions.EntityData.Children = make(map[string]types.YChild)
    redistributions.EntityData.Children["redistribution"] = types.YChild{"Redistribution", nil}
    for i := range redistributions.Redistribution {
        redistributions.EntityData.Children[types.GetSegmentPath(&redistributions.Redistribution[i])] = types.YChild{"Redistribution", &redistributions.Redistribution[i]}
    }
    redistributions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(redistributions.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution
// Redistribution of other protocols into
// this IS-IS instance
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution struct {
    EntityData types.CommonEntityData
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

func (redistribution *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution) GetEntityData() *types.CommonEntityData {
    redistribution.EntityData.YFilter = redistribution.YFilter
    redistribution.EntityData.YangName = "redistribution"
    redistribution.EntityData.BundleName = "cisco_ios_xr"
    redistribution.EntityData.ParentYangName = "redistributions"
    redistribution.EntityData.SegmentPath = "redistribution" + "[protocol-name='" + fmt.Sprintf("%v", redistribution.ProtocolName) + "']"
    redistribution.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistribution.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistribution.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistribution.EntityData.Children = make(map[string]types.YChild)
    redistribution.EntityData.Children["connected-or-static-or-rip-or-subscriber-or-mobile"] = types.YChild{"ConnectedOrStaticOrRipOrSubscriberOrMobile", &redistribution.ConnectedOrStaticOrRipOrSubscriberOrMobile}
    redistribution.EntityData.Children["ospf-or-ospfv3-or-isis-or-application"] = types.YChild{"OspfOrOspfv3OrIsisOrApplication", nil}
    for i := range redistribution.OspfOrOspfv3OrIsisOrApplication {
        redistribution.EntityData.Children[types.GetSegmentPath(&redistribution.OspfOrOspfv3OrIsisOrApplication[i])] = types.YChild{"OspfOrOspfv3OrIsisOrApplication", &redistribution.OspfOrOspfv3OrIsisOrApplication[i]}
    }
    redistribution.EntityData.Children["bgp"] = types.YChild{"Bgp", nil}
    for i := range redistribution.Bgp {
        redistribution.EntityData.Children[types.GetSegmentPath(&redistribution.Bgp[i])] = types.YChild{"Bgp", &redistribution.Bgp[i]}
    }
    redistribution.EntityData.Children["eigrp"] = types.YChild{"Eigrp", nil}
    for i := range redistribution.Eigrp {
        redistribution.EntityData.Children[types.GetSegmentPath(&redistribution.Eigrp[i])] = types.YChild{"Eigrp", &redistribution.Eigrp[i]}
    }
    redistribution.EntityData.Leafs = make(map[string]types.YLeaf)
    redistribution.EntityData.Leafs["protocol-name"] = types.YLeaf{"ProtocolName", redistribution.ProtocolName}
    return &(redistribution.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile
// connected or static or rip or subscriber
// or mobile
// This type is a presence type.
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile struct {
    EntityData types.CommonEntityData
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

func (connectedOrStaticOrRipOrSubscriberOrMobile *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_ConnectedOrStaticOrRipOrSubscriberOrMobile) GetEntityData() *types.CommonEntityData {
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.YFilter = connectedOrStaticOrRipOrSubscriberOrMobile.YFilter
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.YangName = "connected-or-static-or-rip-or-subscriber-or-mobile"
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.BundleName = "cisco_ios_xr"
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.ParentYangName = "redistribution"
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.SegmentPath = "connected-or-static-or-rip-or-subscriber-or-mobile"
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Children = make(map[string]types.YChild)
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs = make(map[string]types.YLeaf)
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["metric"] = types.YLeaf{"Metric", connectedOrStaticOrRipOrSubscriberOrMobile.Metric}
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["levels"] = types.YLeaf{"Levels", connectedOrStaticOrRipOrSubscriberOrMobile.Levels}
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", connectedOrStaticOrRipOrSubscriberOrMobile.RoutePolicyName}
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", connectedOrStaticOrRipOrSubscriberOrMobile.MetricType}
    connectedOrStaticOrRipOrSubscriberOrMobile.EntityData.Leafs["ospf-route-type"] = types.YLeaf{"OspfRouteType", connectedOrStaticOrRipOrSubscriberOrMobile.OspfRouteType}
    return &(connectedOrStaticOrRipOrSubscriberOrMobile.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication
// ospf or ospfv3 or isis or application
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication struct {
    EntityData types.CommonEntityData
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

func (ospfOrOspfv3OrIsisOrApplication *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_OspfOrOspfv3OrIsisOrApplication) GetEntityData() *types.CommonEntityData {
    ospfOrOspfv3OrIsisOrApplication.EntityData.YFilter = ospfOrOspfv3OrIsisOrApplication.YFilter
    ospfOrOspfv3OrIsisOrApplication.EntityData.YangName = "ospf-or-ospfv3-or-isis-or-application"
    ospfOrOspfv3OrIsisOrApplication.EntityData.BundleName = "cisco_ios_xr"
    ospfOrOspfv3OrIsisOrApplication.EntityData.ParentYangName = "redistribution"
    ospfOrOspfv3OrIsisOrApplication.EntityData.SegmentPath = "ospf-or-ospfv3-or-isis-or-application" + "[instance-name='" + fmt.Sprintf("%v", ospfOrOspfv3OrIsisOrApplication.InstanceName) + "']"
    ospfOrOspfv3OrIsisOrApplication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfOrOspfv3OrIsisOrApplication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfOrOspfv3OrIsisOrApplication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfOrOspfv3OrIsisOrApplication.EntityData.Children = make(map[string]types.YChild)
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs = make(map[string]types.YLeaf)
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["instance-name"] = types.YLeaf{"InstanceName", ospfOrOspfv3OrIsisOrApplication.InstanceName}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["metric"] = types.YLeaf{"Metric", ospfOrOspfv3OrIsisOrApplication.Metric}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["levels"] = types.YLeaf{"Levels", ospfOrOspfv3OrIsisOrApplication.Levels}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", ospfOrOspfv3OrIsisOrApplication.RoutePolicyName}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", ospfOrOspfv3OrIsisOrApplication.MetricType}
    ospfOrOspfv3OrIsisOrApplication.EntityData.Leafs["ospf-route-type"] = types.YLeaf{"OspfRouteType", ospfOrOspfv3OrIsisOrApplication.OspfRouteType}
    return &(ospfOrOspfv3OrIsisOrApplication.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp
// bgp
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp struct {
    EntityData types.CommonEntityData
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

func (bgp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "redistribution"
    bgp.EntityData.SegmentPath = "bgp" + "[as-xx='" + fmt.Sprintf("%v", bgp.AsXx) + "']" + "[as-yy='" + fmt.Sprintf("%v", bgp.AsYy) + "']"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["as-xx"] = types.YLeaf{"AsXx", bgp.AsXx}
    bgp.EntityData.Leafs["as-yy"] = types.YLeaf{"AsYy", bgp.AsYy}
    bgp.EntityData.Leafs["metric"] = types.YLeaf{"Metric", bgp.Metric}
    bgp.EntityData.Leafs["levels"] = types.YLeaf{"Levels", bgp.Levels}
    bgp.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", bgp.RoutePolicyName}
    bgp.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", bgp.MetricType}
    bgp.EntityData.Leafs["ospf-route-type"] = types.YLeaf{"OspfRouteType", bgp.OspfRouteType}
    return &(bgp.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp
// eigrp
type Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp struct {
    EntityData types.CommonEntityData
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

func (eigrp *Isis_Instances_Instance_Afs_Af_TopologyName_Redistributions_Redistribution_Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "redistribution"
    eigrp.EntityData.SegmentPath = "eigrp" + "[as-zz='" + fmt.Sprintf("%v", eigrp.AsZz) + "']"
    eigrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrp.EntityData.Children = make(map[string]types.YChild)
    eigrp.EntityData.Leafs = make(map[string]types.YLeaf)
    eigrp.EntityData.Leafs["as-zz"] = types.YLeaf{"AsZz", eigrp.AsZz}
    eigrp.EntityData.Leafs["metric"] = types.YLeaf{"Metric", eigrp.Metric}
    eigrp.EntityData.Leafs["levels"] = types.YLeaf{"Levels", eigrp.Levels}
    eigrp.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", eigrp.RoutePolicyName}
    eigrp.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", eigrp.MetricType}
    eigrp.EntityData.Leafs["ospf-route-type"] = types.YLeaf{"OspfRouteType", eigrp.OspfRouteType}
    return &(eigrp.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals
// Peoridic SPF configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum interval between spf runs. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval.
    SpfPeriodicInterval []Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval
}

func (spfPeriodicIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals) GetEntityData() *types.CommonEntityData {
    spfPeriodicIntervals.EntityData.YFilter = spfPeriodicIntervals.YFilter
    spfPeriodicIntervals.EntityData.YangName = "spf-periodic-intervals"
    spfPeriodicIntervals.EntityData.BundleName = "cisco_ios_xr"
    spfPeriodicIntervals.EntityData.ParentYangName = "topology-name"
    spfPeriodicIntervals.EntityData.SegmentPath = "spf-periodic-intervals"
    spfPeriodicIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfPeriodicIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfPeriodicIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfPeriodicIntervals.EntityData.Children = make(map[string]types.YChild)
    spfPeriodicIntervals.EntityData.Children["spf-periodic-interval"] = types.YChild{"SpfPeriodicInterval", nil}
    for i := range spfPeriodicIntervals.SpfPeriodicInterval {
        spfPeriodicIntervals.EntityData.Children[types.GetSegmentPath(&spfPeriodicIntervals.SpfPeriodicInterval[i])] = types.YChild{"SpfPeriodicInterval", &spfPeriodicIntervals.SpfPeriodicInterval[i]}
    }
    spfPeriodicIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spfPeriodicIntervals.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval
// Maximum interval between spf runs
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Maximum interval in between SPF runs in seconds. The type is interface{}
    // with range: 0..3600. This attribute is mandatory. Units are second.
    PeriodicInterval interface{}
}

func (spfPeriodicInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfPeriodicIntervals_SpfPeriodicInterval) GetEntityData() *types.CommonEntityData {
    spfPeriodicInterval.EntityData.YFilter = spfPeriodicInterval.YFilter
    spfPeriodicInterval.EntityData.YangName = "spf-periodic-interval"
    spfPeriodicInterval.EntityData.BundleName = "cisco_ios_xr"
    spfPeriodicInterval.EntityData.ParentYangName = "spf-periodic-intervals"
    spfPeriodicInterval.EntityData.SegmentPath = "spf-periodic-interval" + "[level='" + fmt.Sprintf("%v", spfPeriodicInterval.Level) + "']"
    spfPeriodicInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfPeriodicInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfPeriodicInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfPeriodicInterval.EntityData.Children = make(map[string]types.YChild)
    spfPeriodicInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    spfPeriodicInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", spfPeriodicInterval.Level}
    spfPeriodicInterval.EntityData.Leafs["periodic-interval"] = types.YLeaf{"PeriodicInterval", spfPeriodicInterval.PeriodicInterval}
    return &(spfPeriodicInterval.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals
// SPF-interval configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route calculation scheduling parameters. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval.
    SpfInterval []Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval
}

func (spfIntervals *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals) GetEntityData() *types.CommonEntityData {
    spfIntervals.EntityData.YFilter = spfIntervals.YFilter
    spfIntervals.EntityData.YangName = "spf-intervals"
    spfIntervals.EntityData.BundleName = "cisco_ios_xr"
    spfIntervals.EntityData.ParentYangName = "topology-name"
    spfIntervals.EntityData.SegmentPath = "spf-intervals"
    spfIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfIntervals.EntityData.Children = make(map[string]types.YChild)
    spfIntervals.EntityData.Children["spf-interval"] = types.YChild{"SpfInterval", nil}
    for i := range spfIntervals.SpfInterval {
        spfIntervals.EntityData.Children[types.GetSegmentPath(&spfIntervals.SpfInterval[i])] = types.YChild{"SpfInterval", &spfIntervals.SpfInterval[i]}
    }
    spfIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spfIntervals.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval
// Route calculation scheduling parameters
type Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval struct {
    EntityData types.CommonEntityData
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

func (spfInterval *Isis_Instances_Instance_Afs_Af_TopologyName_SpfIntervals_SpfInterval) GetEntityData() *types.CommonEntityData {
    spfInterval.EntityData.YFilter = spfInterval.YFilter
    spfInterval.EntityData.YangName = "spf-interval"
    spfInterval.EntityData.BundleName = "cisco_ios_xr"
    spfInterval.EntityData.ParentYangName = "spf-intervals"
    spfInterval.EntityData.SegmentPath = "spf-interval" + "[level='" + fmt.Sprintf("%v", spfInterval.Level) + "']"
    spfInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spfInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spfInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spfInterval.EntityData.Children = make(map[string]types.YChild)
    spfInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    spfInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", spfInterval.Level}
    spfInterval.EntityData.Leafs["maximum-wait"] = types.YLeaf{"MaximumWait", spfInterval.MaximumWait}
    spfInterval.EntityData.Leafs["initial-wait"] = types.YLeaf{"InitialWait", spfInterval.InitialWait}
    spfInterval.EntityData.Leafs["secondary-wait"] = types.YLeaf{"SecondaryWait", spfInterval.SecondaryWait}
    return &(spfInterval.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence
// Enable convergence monitoring
type Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable convergence monitoring. The type is interface{}.
    Enable interface{}

    // Enable the Tracking of IP-Frr Convergence. The type is interface{}.
    TrackIpFrr interface{}

    // Enable the monitoring of individual prefixes (prefix list name). The type
    // is string with length: 1..32.
    PrefixList interface{}
}

func (monitorConvergence *Isis_Instances_Instance_Afs_Af_TopologyName_MonitorConvergence) GetEntityData() *types.CommonEntityData {
    monitorConvergence.EntityData.YFilter = monitorConvergence.YFilter
    monitorConvergence.EntityData.YangName = "monitor-convergence"
    monitorConvergence.EntityData.BundleName = "cisco_ios_xr"
    monitorConvergence.EntityData.ParentYangName = "topology-name"
    monitorConvergence.EntityData.SegmentPath = "monitor-convergence"
    monitorConvergence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    monitorConvergence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    monitorConvergence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    monitorConvergence.EntityData.Children = make(map[string]types.YChild)
    monitorConvergence.EntityData.Leafs = make(map[string]types.YLeaf)
    monitorConvergence.EntityData.Leafs["enable"] = types.YLeaf{"Enable", monitorConvergence.Enable}
    monitorConvergence.EntityData.Leafs["track-ip-frr"] = types.YLeaf{"TrackIpFrr", monitorConvergence.TrackIpFrr}
    monitorConvergence.EntityData.Leafs["prefix-list"] = types.YLeaf{"PrefixList", monitorConvergence.PrefixList}
    return &(monitorConvergence.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation
// Control origination of a default route with
// the option of using a policy.  If no policy
// is specified the default route is
// advertised with zero cost in level 2 only.
type Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation struct {
    EntityData types.CommonEntityData
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

func (defaultInformation *Isis_Instances_Instance_Afs_Af_TopologyName_DefaultInformation) GetEntityData() *types.CommonEntityData {
    defaultInformation.EntityData.YFilter = defaultInformation.YFilter
    defaultInformation.EntityData.YangName = "default-information"
    defaultInformation.EntityData.BundleName = "cisco_ios_xr"
    defaultInformation.EntityData.ParentYangName = "topology-name"
    defaultInformation.EntityData.SegmentPath = "default-information"
    defaultInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultInformation.EntityData.Children = make(map[string]types.YChild)
    defaultInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultInformation.EntityData.Leafs["use-policy"] = types.YLeaf{"UsePolicy", defaultInformation.UsePolicy}
    defaultInformation.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", defaultInformation.PolicyName}
    defaultInformation.EntityData.Leafs["external"] = types.YLeaf{"External", defaultInformation.External}
    return &(defaultInformation.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances
// Per-route administrative
// distanceconfiguration
type Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Administrative distance configuration. The supplied distance is applied to
    // all routes discovered from the specified source, or only those that match
    // the supplied prefix list if this is specified. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance.
    AdminDistance []Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance
}

func (adminDistances *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances) GetEntityData() *types.CommonEntityData {
    adminDistances.EntityData.YFilter = adminDistances.YFilter
    adminDistances.EntityData.YangName = "admin-distances"
    adminDistances.EntityData.BundleName = "cisco_ios_xr"
    adminDistances.EntityData.ParentYangName = "topology-name"
    adminDistances.EntityData.SegmentPath = "admin-distances"
    adminDistances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminDistances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminDistances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminDistances.EntityData.Children = make(map[string]types.YChild)
    adminDistances.EntityData.Children["admin-distance"] = types.YChild{"AdminDistance", nil}
    for i := range adminDistances.AdminDistance {
        adminDistances.EntityData.Children[types.GetSegmentPath(&adminDistances.AdminDistance[i])] = types.YChild{"AdminDistance", &adminDistances.AdminDistance[i]}
    }
    adminDistances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adminDistances.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance
// Administrative distance configuration. The
// supplied distance is applied to all routes
// discovered from the specified source, or
// only those that match the supplied prefix
// list if this is specified
type Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP route source prefix. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))'.
    AddressPrefix interface{}

    // Administrative distance. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    Distance interface{}

    // List of prefixes to which this distance applies. The type is string with
    // length: 1..32.
    PrefixList interface{}
}

func (adminDistance *Isis_Instances_Instance_Afs_Af_TopologyName_AdminDistances_AdminDistance) GetEntityData() *types.CommonEntityData {
    adminDistance.EntityData.YFilter = adminDistance.YFilter
    adminDistance.EntityData.YangName = "admin-distance"
    adminDistance.EntityData.BundleName = "cisco_ios_xr"
    adminDistance.EntityData.ParentYangName = "admin-distances"
    adminDistance.EntityData.SegmentPath = "admin-distance" + "[address-prefix='" + fmt.Sprintf("%v", adminDistance.AddressPrefix) + "']"
    adminDistance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminDistance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminDistance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminDistance.EntityData.Children = make(map[string]types.YChild)
    adminDistance.EntityData.Leafs = make(map[string]types.YLeaf)
    adminDistance.EntityData.Leafs["address-prefix"] = types.YLeaf{"AddressPrefix", adminDistance.AddressPrefix}
    adminDistance.EntityData.Leafs["distance"] = types.YLeaf{"Distance", adminDistance.Distance}
    adminDistance.EntityData.Leafs["prefix-list"] = types.YLeaf{"PrefixList", adminDistance.PrefixList}
    return &(adminDistance.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Ispf
// ISPF configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Ispf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISPF state (enable/disable).
    States Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States
}

func (ispf *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf) GetEntityData() *types.CommonEntityData {
    ispf.EntityData.YFilter = ispf.YFilter
    ispf.EntityData.YangName = "ispf"
    ispf.EntityData.BundleName = "cisco_ios_xr"
    ispf.EntityData.ParentYangName = "topology-name"
    ispf.EntityData.SegmentPath = "ispf"
    ispf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ispf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ispf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ispf.EntityData.Children = make(map[string]types.YChild)
    ispf.EntityData.Children["states"] = types.YChild{"States", &ispf.States}
    ispf.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ispf.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States
// ISPF state (enable/disable)
type Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/disable ISPF. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State.
    State []Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State
}

func (states *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States) GetEntityData() *types.CommonEntityData {
    states.EntityData.YFilter = states.YFilter
    states.EntityData.YangName = "states"
    states.EntityData.BundleName = "cisco_ios_xr"
    states.EntityData.ParentYangName = "ispf"
    states.EntityData.SegmentPath = "states"
    states.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    states.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    states.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    states.EntityData.Children = make(map[string]types.YChild)
    states.EntityData.Children["state"] = types.YChild{"State", nil}
    for i := range states.State {
        states.EntityData.Children[types.GetSegmentPath(&states.State[i])] = types.YChild{"State", &states.State[i]}
    }
    states.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(states.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State
// Enable/disable ISPF
type Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // State. The type is IsisispfState. This attribute is mandatory.
    State interface{}
}

func (state *Isis_Instances_Instance_Afs_Af_TopologyName_Ispf_States_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xr"
    state.EntityData.ParentYangName = "states"
    state.EntityData.SegmentPath = "state" + "[level='" + fmt.Sprintf("%v", state.Level) + "']"
    state.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    state.EntityData.Children = make(map[string]types.YChild)
    state.EntityData.Leafs = make(map[string]types.YLeaf)
    state.EntityData.Leafs["level"] = types.YLeaf{"Level", state.Level}
    state.EntityData.Leafs["state"] = types.YLeaf{"State", state.State}
    return &(state.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal
// MPLS LDP configuration. MPLS LDP
// configuration will only be applied for the
// IPv4-unicast address-family.
type Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If TRUE, LDP will be enabled onall IS-IS interfaces enabled for this
    // address-family. The type is bool.
    AutoConfig interface{}
}

func (mplsLdpGlobal *Isis_Instances_Instance_Afs_Af_TopologyName_MplsLdpGlobal) GetEntityData() *types.CommonEntityData {
    mplsLdpGlobal.EntityData.YFilter = mplsLdpGlobal.YFilter
    mplsLdpGlobal.EntityData.YangName = "mpls-ldp-global"
    mplsLdpGlobal.EntityData.BundleName = "cisco_ios_xr"
    mplsLdpGlobal.EntityData.ParentYangName = "topology-name"
    mplsLdpGlobal.EntityData.SegmentPath = "mpls-ldp-global"
    mplsLdpGlobal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLdpGlobal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLdpGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLdpGlobal.EntityData.Children = make(map[string]types.YChild)
    mplsLdpGlobal.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsLdpGlobal.EntityData.Leafs["auto-config"] = types.YLeaf{"AutoConfig", mplsLdpGlobal.AutoConfig}
    return &(mplsLdpGlobal.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Mpls
// MPLS configuration. MPLS configuration will
// only be applied for the IPv4-unicast
// address-family.
type Isis_Instances_Instance_Afs_Af_TopologyName_Mpls struct {
    EntityData types.CommonEntityData
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

func (mpls *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls) GetEntityData() *types.CommonEntityData {
    mpls.EntityData.YFilter = mpls.YFilter
    mpls.EntityData.YangName = "mpls"
    mpls.EntityData.BundleName = "cisco_ios_xr"
    mpls.EntityData.ParentYangName = "topology-name"
    mpls.EntityData.SegmentPath = "mpls"
    mpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpls.EntityData.Children = make(map[string]types.YChild)
    mpls.EntityData.Children["router-id"] = types.YChild{"RouterId", &mpls.RouterId}
    mpls.EntityData.Children["level"] = types.YChild{"Level", &mpls.Level}
    mpls.EntityData.Leafs = make(map[string]types.YLeaf)
    mpls.EntityData.Leafs["igp-intact"] = types.YLeaf{"IgpIntact", mpls.IgpIntact}
    mpls.EntityData.Leafs["multicast-intact"] = types.YLeaf{"MulticastIntact", mpls.MulticastIntact}
    return &(mpls.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId
// Traffic Engineering stable IP address for
// system
type Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address to be used as a router ID. Precisely one of Address and
    // Interface must be specified. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Interface with designated stable IP address to be used as a router ID. This
    // must be a Loopback interface. Precisely one of Address and Interface must
    // be specified. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (routerId *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_RouterId) GetEntityData() *types.CommonEntityData {
    routerId.EntityData.YFilter = routerId.YFilter
    routerId.EntityData.YangName = "router-id"
    routerId.EntityData.BundleName = "cisco_ios_xr"
    routerId.EntityData.ParentYangName = "mpls"
    routerId.EntityData.SegmentPath = "router-id"
    routerId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routerId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routerId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routerId.EntityData.Children = make(map[string]types.YChild)
    routerId.EntityData.Leafs = make(map[string]types.YLeaf)
    routerId.EntityData.Leafs["address"] = types.YLeaf{"Address", routerId.Address}
    routerId.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", routerId.InterfaceName}
    return &(routerId.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level
// Enable MPLS for an IS-IS at the given
// levels
type Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Level 1 enabled. The type is bool.
    Level1 interface{}

    // Level 2 enabled. The type is bool.
    Level2 interface{}
}

func (level *Isis_Instances_Instance_Afs_Af_TopologyName_Mpls_Level) GetEntityData() *types.CommonEntityData {
    level.EntityData.YFilter = level.YFilter
    level.EntityData.YangName = "level"
    level.EntityData.BundleName = "cisco_ios_xr"
    level.EntityData.ParentYangName = "mpls"
    level.EntityData.SegmentPath = "level"
    level.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    level.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    level.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    level.EntityData.Children = make(map[string]types.YChild)
    level.EntityData.Leafs = make(map[string]types.YLeaf)
    level.EntityData.Leafs["level1"] = types.YLeaf{"Level1", level.Level1}
    level.EntityData.Leafs["level2"] = types.YLeaf{"Level2", level.Level2}
    return &(level.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids
// Manual Adjacecy SID configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Assign adjancency SID to an interface. The type is slice of
    // Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid.
    ManualAdjSid []Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid
}

func (manualAdjSids *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids) GetEntityData() *types.CommonEntityData {
    manualAdjSids.EntityData.YFilter = manualAdjSids.YFilter
    manualAdjSids.EntityData.YangName = "manual-adj-sids"
    manualAdjSids.EntityData.BundleName = "cisco_ios_xr"
    manualAdjSids.EntityData.ParentYangName = "topology-name"
    manualAdjSids.EntityData.SegmentPath = "manual-adj-sids"
    manualAdjSids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualAdjSids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualAdjSids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualAdjSids.EntityData.Children = make(map[string]types.YChild)
    manualAdjSids.EntityData.Children["manual-adj-sid"] = types.YChild{"ManualAdjSid", nil}
    for i := range manualAdjSids.ManualAdjSid {
        manualAdjSids.EntityData.Children[types.GetSegmentPath(&manualAdjSids.ManualAdjSid[i])] = types.YChild{"ManualAdjSid", &manualAdjSids.ManualAdjSid[i]}
    }
    manualAdjSids.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(manualAdjSids.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid
// Assign adjancency SID to an interface
type Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid struct {
    EntityData types.CommonEntityData
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

func (manualAdjSid *Isis_Instances_Instance_Afs_Af_TopologyName_ManualAdjSids_ManualAdjSid) GetEntityData() *types.CommonEntityData {
    manualAdjSid.EntityData.YFilter = manualAdjSid.YFilter
    manualAdjSid.EntityData.YangName = "manual-adj-sid"
    manualAdjSid.EntityData.BundleName = "cisco_ios_xr"
    manualAdjSid.EntityData.ParentYangName = "manual-adj-sids"
    manualAdjSid.EntityData.SegmentPath = "manual-adj-sid" + "[level='" + fmt.Sprintf("%v", manualAdjSid.Level) + "']" + "[sid-type='" + fmt.Sprintf("%v", manualAdjSid.SidType) + "']" + "[sid='" + fmt.Sprintf("%v", manualAdjSid.Sid) + "']"
    manualAdjSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualAdjSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualAdjSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualAdjSid.EntityData.Children = make(map[string]types.YChild)
    manualAdjSid.EntityData.Leafs = make(map[string]types.YLeaf)
    manualAdjSid.EntityData.Leafs["level"] = types.YLeaf{"Level", manualAdjSid.Level}
    manualAdjSid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", manualAdjSid.SidType}
    manualAdjSid.EntityData.Leafs["sid"] = types.YLeaf{"Sid", manualAdjSid.Sid}
    manualAdjSid.EntityData.Leafs["protected"] = types.YLeaf{"Protected", manualAdjSid.Protected}
    return &(manualAdjSid.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Metrics
// Metric configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Metrics struct {
    EntityData types.CommonEntityData
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

func (metrics *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics) GetEntityData() *types.CommonEntityData {
    metrics.EntityData.YFilter = metrics.YFilter
    metrics.EntityData.YangName = "metrics"
    metrics.EntityData.BundleName = "cisco_ios_xr"
    metrics.EntityData.ParentYangName = "topology-name"
    metrics.EntityData.SegmentPath = "metrics"
    metrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metrics.EntityData.Children = make(map[string]types.YChild)
    metrics.EntityData.Children["metric"] = types.YChild{"Metric", nil}
    for i := range metrics.Metric {
        metrics.EntityData.Children[types.GetSegmentPath(&metrics.Metric[i])] = types.YChild{"Metric", &metrics.Metric[i]}
    }
    metrics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(metrics.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed metric: <1-63> for narrow, <1-16777215> for wide. The type is one
    // of the following types: enumeration
    // Isis.Instances.Instance.Interfaces.Interface.InterfaceAfs.InterfaceAf.TopologyName.Metrics.Metric.Metric_
    // This attribute is mandatory., or int with range: 1..16777215 This attribute
    // is mandatory..
    Metric interface{}
}

func (metric *Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "metrics"
    metric.EntityData.SegmentPath = "metric" + "[level='" + fmt.Sprintf("%v", metric.Level) + "']"
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = make(map[string]types.YChild)
    metric.EntityData.Leafs = make(map[string]types.YLeaf)
    metric.EntityData.Leafs["level"] = types.YLeaf{"Level", metric.Level}
    metric.EntityData.Leafs["metric"] = types.YLeaf{"Metric", metric.Metric}
    return &(metric.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric_Metric_ represents <1-16777215> for wide
type Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric_Metric_ string

const (
    // Maximum wide metric.  All routers will
    // exclude this link from their SPF
    Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric_Metric__maximum Isis_Instances_Instance_Afs_Af_TopologyName_Metrics_Metric_Metric_ = "maximum"
)

// Isis_Instances_Instance_Afs_Af_TopologyName_Weights
// Weight configuration
type Isis_Instances_Instance_Afs_Af_TopologyName_Weights struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Weight configuration under interface for load balancing. The type is slice
    // of Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight.
    Weight []Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight
}

func (weights *Isis_Instances_Instance_Afs_Af_TopologyName_Weights) GetEntityData() *types.CommonEntityData {
    weights.EntityData.YFilter = weights.YFilter
    weights.EntityData.YangName = "weights"
    weights.EntityData.BundleName = "cisco_ios_xr"
    weights.EntityData.ParentYangName = "topology-name"
    weights.EntityData.SegmentPath = "weights"
    weights.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    weights.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    weights.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    weights.EntityData.Children = make(map[string]types.YChild)
    weights.EntityData.Children["weight"] = types.YChild{"Weight", nil}
    for i := range weights.Weight {
        weights.EntityData.Children[types.GetSegmentPath(&weights.Weight[i])] = types.YChild{"Weight", &weights.Weight[i]}
    }
    weights.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(weights.EntityData)
}

// Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight
// Weight configuration under interface for load
// balancing
type Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Weight to be configured under interface for Load Balancing. Allowed weight:
    // <1-16777215>. The type is interface{} with range: 1..16777214. This
    // attribute is mandatory.
    Weight interface{}
}

func (weight *Isis_Instances_Instance_Afs_Af_TopologyName_Weights_Weight) GetEntityData() *types.CommonEntityData {
    weight.EntityData.YFilter = weight.YFilter
    weight.EntityData.YangName = "weight"
    weight.EntityData.BundleName = "cisco_ios_xr"
    weight.EntityData.ParentYangName = "weights"
    weight.EntityData.SegmentPath = "weight" + "[level='" + fmt.Sprintf("%v", weight.Level) + "']"
    weight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    weight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    weight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    weight.EntityData.Children = make(map[string]types.YChild)
    weight.EntityData.Leafs = make(map[string]types.YLeaf)
    weight.EntityData.Leafs["level"] = types.YLeaf{"Level", weight.Level}
    weight.EntityData.Leafs["weight"] = types.YLeaf{"Weight", weight.Weight}
    return &(weight.EntityData)
}

// Isis_Instances_Instance_LspRefreshIntervals
// LSP refresh-interval configuration
type Isis_Instances_Instance_LspRefreshIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval between re-flooding of unchanged LSPs. The type is slice of
    // Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval.
    LspRefreshInterval []Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval
}

func (lspRefreshIntervals *Isis_Instances_Instance_LspRefreshIntervals) GetEntityData() *types.CommonEntityData {
    lspRefreshIntervals.EntityData.YFilter = lspRefreshIntervals.YFilter
    lspRefreshIntervals.EntityData.YangName = "lsp-refresh-intervals"
    lspRefreshIntervals.EntityData.BundleName = "cisco_ios_xr"
    lspRefreshIntervals.EntityData.ParentYangName = "instance"
    lspRefreshIntervals.EntityData.SegmentPath = "lsp-refresh-intervals"
    lspRefreshIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRefreshIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRefreshIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRefreshIntervals.EntityData.Children = make(map[string]types.YChild)
    lspRefreshIntervals.EntityData.Children["lsp-refresh-interval"] = types.YChild{"LspRefreshInterval", nil}
    for i := range lspRefreshIntervals.LspRefreshInterval {
        lspRefreshIntervals.EntityData.Children[types.GetSegmentPath(&lspRefreshIntervals.LspRefreshInterval[i])] = types.YChild{"LspRefreshInterval", &lspRefreshIntervals.LspRefreshInterval[i]}
    }
    lspRefreshIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspRefreshIntervals.EntityData)
}

// Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval
// Interval between re-flooding of unchanged
// LSPs
type Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 1..65535. This attribute is
    // mandatory. Units are second.
    Interval interface{}
}

func (lspRefreshInterval *Isis_Instances_Instance_LspRefreshIntervals_LspRefreshInterval) GetEntityData() *types.CommonEntityData {
    lspRefreshInterval.EntityData.YFilter = lspRefreshInterval.YFilter
    lspRefreshInterval.EntityData.YangName = "lsp-refresh-interval"
    lspRefreshInterval.EntityData.BundleName = "cisco_ios_xr"
    lspRefreshInterval.EntityData.ParentYangName = "lsp-refresh-intervals"
    lspRefreshInterval.EntityData.SegmentPath = "lsp-refresh-interval" + "[level='" + fmt.Sprintf("%v", lspRefreshInterval.Level) + "']"
    lspRefreshInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRefreshInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRefreshInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRefreshInterval.EntityData.Children = make(map[string]types.YChild)
    lspRefreshInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    lspRefreshInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", lspRefreshInterval.Level}
    lspRefreshInterval.EntityData.Leafs["interval"] = types.YLeaf{"Interval", lspRefreshInterval.Interval}
    return &(lspRefreshInterval.EntityData)
}

// Isis_Instances_Instance_Distribute
// IS-IS Distribute BGP-LS configuration
// This type is a presence type.
type Isis_Instances_Instance_Distribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance ID. The type is interface{} with range: 1..65535.
    DistInstId interface{}

    // Level. The type is IsisConfigurableLevels.
    Level interface{}

    // Seconds. The type is interface{} with range: 1..20. Units are second.
    DistThrottle interface{}
}

func (distribute *Isis_Instances_Instance_Distribute) GetEntityData() *types.CommonEntityData {
    distribute.EntityData.YFilter = distribute.YFilter
    distribute.EntityData.YangName = "distribute"
    distribute.EntityData.BundleName = "cisco_ios_xr"
    distribute.EntityData.ParentYangName = "instance"
    distribute.EntityData.SegmentPath = "distribute"
    distribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distribute.EntityData.Children = make(map[string]types.YChild)
    distribute.EntityData.Leafs = make(map[string]types.YLeaf)
    distribute.EntityData.Leafs["dist-inst-id"] = types.YLeaf{"DistInstId", distribute.DistInstId}
    distribute.EntityData.Leafs["level"] = types.YLeaf{"Level", distribute.Level}
    distribute.EntityData.Leafs["dist-throttle"] = types.YLeaf{"DistThrottle", distribute.DistThrottle}
    return &(distribute.EntityData)
}

// Isis_Instances_Instance_LspAcceptPasswords
// LSP/SNP accept password configuration
type Isis_Instances_Instance_LspAcceptPasswords struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP/SNP accept passwords. This requires the existence of an LSPPassword of
    // the same level . The type is slice of
    // Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword.
    LspAcceptPassword []Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword
}

func (lspAcceptPasswords *Isis_Instances_Instance_LspAcceptPasswords) GetEntityData() *types.CommonEntityData {
    lspAcceptPasswords.EntityData.YFilter = lspAcceptPasswords.YFilter
    lspAcceptPasswords.EntityData.YangName = "lsp-accept-passwords"
    lspAcceptPasswords.EntityData.BundleName = "cisco_ios_xr"
    lspAcceptPasswords.EntityData.ParentYangName = "instance"
    lspAcceptPasswords.EntityData.SegmentPath = "lsp-accept-passwords"
    lspAcceptPasswords.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspAcceptPasswords.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspAcceptPasswords.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspAcceptPasswords.EntityData.Children = make(map[string]types.YChild)
    lspAcceptPasswords.EntityData.Children["lsp-accept-password"] = types.YChild{"LspAcceptPassword", nil}
    for i := range lspAcceptPasswords.LspAcceptPassword {
        lspAcceptPasswords.EntityData.Children[types.GetSegmentPath(&lspAcceptPasswords.LspAcceptPassword[i])] = types.YChild{"LspAcceptPassword", &lspAcceptPasswords.LspAcceptPassword[i]}
    }
    lspAcceptPasswords.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspAcceptPasswords.EntityData)
}

// Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword
// LSP/SNP accept passwords. This requires the
// existence of an LSPPassword of the same level
// .
type Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Password. The type is string with pattern: b'(!.+)|([^!].+)'. This
    // attribute is mandatory.
    Password interface{}
}

func (lspAcceptPassword *Isis_Instances_Instance_LspAcceptPasswords_LspAcceptPassword) GetEntityData() *types.CommonEntityData {
    lspAcceptPassword.EntityData.YFilter = lspAcceptPassword.YFilter
    lspAcceptPassword.EntityData.YangName = "lsp-accept-password"
    lspAcceptPassword.EntityData.BundleName = "cisco_ios_xr"
    lspAcceptPassword.EntityData.ParentYangName = "lsp-accept-passwords"
    lspAcceptPassword.EntityData.SegmentPath = "lsp-accept-password" + "[level='" + fmt.Sprintf("%v", lspAcceptPassword.Level) + "']"
    lspAcceptPassword.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspAcceptPassword.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspAcceptPassword.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspAcceptPassword.EntityData.Children = make(map[string]types.YChild)
    lspAcceptPassword.EntityData.Leafs = make(map[string]types.YLeaf)
    lspAcceptPassword.EntityData.Leafs["level"] = types.YLeaf{"Level", lspAcceptPassword.Level}
    lspAcceptPassword.EntityData.Leafs["password"] = types.YLeaf{"Password", lspAcceptPassword.Password}
    return &(lspAcceptPassword.EntityData)
}

// Isis_Instances_Instance_LspMtus
// LSP MTU configuration
type Isis_Instances_Instance_LspMtus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP MTU. The type is slice of Isis_Instances_Instance_LspMtus_LspMtu.
    LspMtu []Isis_Instances_Instance_LspMtus_LspMtu
}

func (lspMtus *Isis_Instances_Instance_LspMtus) GetEntityData() *types.CommonEntityData {
    lspMtus.EntityData.YFilter = lspMtus.YFilter
    lspMtus.EntityData.YangName = "lsp-mtus"
    lspMtus.EntityData.BundleName = "cisco_ios_xr"
    lspMtus.EntityData.ParentYangName = "instance"
    lspMtus.EntityData.SegmentPath = "lsp-mtus"
    lspMtus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspMtus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspMtus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspMtus.EntityData.Children = make(map[string]types.YChild)
    lspMtus.EntityData.Children["lsp-mtu"] = types.YChild{"LspMtu", nil}
    for i := range lspMtus.LspMtu {
        lspMtus.EntityData.Children[types.GetSegmentPath(&lspMtus.LspMtu[i])] = types.YChild{"LspMtu", &lspMtus.LspMtu[i]}
    }
    lspMtus.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspMtus.EntityData)
}

// Isis_Instances_Instance_LspMtus_LspMtu
// LSP MTU
type Isis_Instances_Instance_LspMtus_LspMtu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Bytes. The type is interface{} with range: 128..4352. This attribute is
    // mandatory. Units are byte.
    Mtu interface{}
}

func (lspMtu *Isis_Instances_Instance_LspMtus_LspMtu) GetEntityData() *types.CommonEntityData {
    lspMtu.EntityData.YFilter = lspMtu.YFilter
    lspMtu.EntityData.YangName = "lsp-mtu"
    lspMtu.EntityData.BundleName = "cisco_ios_xr"
    lspMtu.EntityData.ParentYangName = "lsp-mtus"
    lspMtu.EntityData.SegmentPath = "lsp-mtu" + "[level='" + fmt.Sprintf("%v", lspMtu.Level) + "']"
    lspMtu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspMtu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspMtu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspMtu.EntityData.Children = make(map[string]types.YChild)
    lspMtu.EntityData.Leafs = make(map[string]types.YLeaf)
    lspMtu.EntityData.Leafs["level"] = types.YLeaf{"Level", lspMtu.Level}
    lspMtu.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", lspMtu.Mtu}
    return &(lspMtu.EntityData)
}

// Isis_Instances_Instance_Nsf
// IS-IS NSF configuration
type Isis_Instances_Instance_Nsf struct {
    EntityData types.CommonEntityData
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

func (nsf *Isis_Instances_Instance_Nsf) GetEntityData() *types.CommonEntityData {
    nsf.EntityData.YFilter = nsf.YFilter
    nsf.EntityData.YangName = "nsf"
    nsf.EntityData.BundleName = "cisco_ios_xr"
    nsf.EntityData.ParentYangName = "instance"
    nsf.EntityData.SegmentPath = "nsf"
    nsf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nsf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nsf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nsf.EntityData.Children = make(map[string]types.YChild)
    nsf.EntityData.Leafs = make(map[string]types.YLeaf)
    nsf.EntityData.Leafs["flavor"] = types.YLeaf{"Flavor", nsf.Flavor}
    nsf.EntityData.Leafs["interface-timer"] = types.YLeaf{"InterfaceTimer", nsf.InterfaceTimer}
    nsf.EntityData.Leafs["max-interface-timer-expiry"] = types.YLeaf{"MaxInterfaceTimerExpiry", nsf.MaxInterfaceTimerExpiry}
    nsf.EntityData.Leafs["lifetime"] = types.YLeaf{"Lifetime", nsf.Lifetime}
    return &(nsf.EntityData)
}

// Isis_Instances_Instance_LinkGroups
// Link Group
type Isis_Instances_Instance_LinkGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for link group name. The type is slice of
    // Isis_Instances_Instance_LinkGroups_LinkGroup.
    LinkGroup []Isis_Instances_Instance_LinkGroups_LinkGroup
}

func (linkGroups *Isis_Instances_Instance_LinkGroups) GetEntityData() *types.CommonEntityData {
    linkGroups.EntityData.YFilter = linkGroups.YFilter
    linkGroups.EntityData.YangName = "link-groups"
    linkGroups.EntityData.BundleName = "cisco_ios_xr"
    linkGroups.EntityData.ParentYangName = "instance"
    linkGroups.EntityData.SegmentPath = "link-groups"
    linkGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkGroups.EntityData.Children = make(map[string]types.YChild)
    linkGroups.EntityData.Children["link-group"] = types.YChild{"LinkGroup", nil}
    for i := range linkGroups.LinkGroup {
        linkGroups.EntityData.Children[types.GetSegmentPath(&linkGroups.LinkGroup[i])] = types.YChild{"LinkGroup", &linkGroups.LinkGroup[i]}
    }
    linkGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(linkGroups.EntityData)
}

// Isis_Instances_Instance_LinkGroups_LinkGroup
// Configuration for link group name
type Isis_Instances_Instance_LinkGroups_LinkGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Link Group Name. The type is string with length:
    // 1..40.
    LinkGroupName interface{}

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

func (linkGroup *Isis_Instances_Instance_LinkGroups_LinkGroup) GetEntityData() *types.CommonEntityData {
    linkGroup.EntityData.YFilter = linkGroup.YFilter
    linkGroup.EntityData.YangName = "link-group"
    linkGroup.EntityData.BundleName = "cisco_ios_xr"
    linkGroup.EntityData.ParentYangName = "link-groups"
    linkGroup.EntityData.SegmentPath = "link-group" + "[link-group-name='" + fmt.Sprintf("%v", linkGroup.LinkGroupName) + "']"
    linkGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkGroup.EntityData.Children = make(map[string]types.YChild)
    linkGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    linkGroup.EntityData.Leafs["link-group-name"] = types.YLeaf{"LinkGroupName", linkGroup.LinkGroupName}
    linkGroup.EntityData.Leafs["metric-offset"] = types.YLeaf{"MetricOffset", linkGroup.MetricOffset}
    linkGroup.EntityData.Leafs["revert-members"] = types.YLeaf{"RevertMembers", linkGroup.RevertMembers}
    linkGroup.EntityData.Leafs["minimum-members"] = types.YLeaf{"MinimumMembers", linkGroup.MinimumMembers}
    return &(linkGroup.EntityData)
}

// Isis_Instances_Instance_LspCheckIntervals
// LSP checksum check interval configuration
type Isis_Instances_Instance_LspCheckIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP checksum check interval parameters. The type is slice of
    // Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval.
    LspCheckInterval []Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval
}

func (lspCheckIntervals *Isis_Instances_Instance_LspCheckIntervals) GetEntityData() *types.CommonEntityData {
    lspCheckIntervals.EntityData.YFilter = lspCheckIntervals.YFilter
    lspCheckIntervals.EntityData.YangName = "lsp-check-intervals"
    lspCheckIntervals.EntityData.BundleName = "cisco_ios_xr"
    lspCheckIntervals.EntityData.ParentYangName = "instance"
    lspCheckIntervals.EntityData.SegmentPath = "lsp-check-intervals"
    lspCheckIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspCheckIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspCheckIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspCheckIntervals.EntityData.Children = make(map[string]types.YChild)
    lspCheckIntervals.EntityData.Children["lsp-check-interval"] = types.YChild{"LspCheckInterval", nil}
    for i := range lspCheckIntervals.LspCheckInterval {
        lspCheckIntervals.EntityData.Children[types.GetSegmentPath(&lspCheckIntervals.LspCheckInterval[i])] = types.YChild{"LspCheckInterval", &lspCheckIntervals.LspCheckInterval[i]}
    }
    lspCheckIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspCheckIntervals.EntityData)
}

// Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval
// LSP checksum check interval parameters
type Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // LSP checksum check interval time in seconds. The type is interface{} with
    // range: 10..65535. This attribute is mandatory. Units are second.
    Interval interface{}
}

func (lspCheckInterval *Isis_Instances_Instance_LspCheckIntervals_LspCheckInterval) GetEntityData() *types.CommonEntityData {
    lspCheckInterval.EntityData.YFilter = lspCheckInterval.YFilter
    lspCheckInterval.EntityData.YangName = "lsp-check-interval"
    lspCheckInterval.EntityData.BundleName = "cisco_ios_xr"
    lspCheckInterval.EntityData.ParentYangName = "lsp-check-intervals"
    lspCheckInterval.EntityData.SegmentPath = "lsp-check-interval" + "[level='" + fmt.Sprintf("%v", lspCheckInterval.Level) + "']"
    lspCheckInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspCheckInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspCheckInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspCheckInterval.EntityData.Children = make(map[string]types.YChild)
    lspCheckInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    lspCheckInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", lspCheckInterval.Level}
    lspCheckInterval.EntityData.Leafs["interval"] = types.YLeaf{"Interval", lspCheckInterval.Interval}
    return &(lspCheckInterval.EntityData)
}

// Isis_Instances_Instance_LspPasswords
// LSP/SNP password configuration
type Isis_Instances_Instance_LspPasswords struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP/SNP passwords. This must exist if an LSPAcceptPassword of the same
    // level exists. The type is slice of
    // Isis_Instances_Instance_LspPasswords_LspPassword.
    LspPassword []Isis_Instances_Instance_LspPasswords_LspPassword
}

func (lspPasswords *Isis_Instances_Instance_LspPasswords) GetEntityData() *types.CommonEntityData {
    lspPasswords.EntityData.YFilter = lspPasswords.YFilter
    lspPasswords.EntityData.YangName = "lsp-passwords"
    lspPasswords.EntityData.BundleName = "cisco_ios_xr"
    lspPasswords.EntityData.ParentYangName = "instance"
    lspPasswords.EntityData.SegmentPath = "lsp-passwords"
    lspPasswords.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPasswords.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPasswords.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPasswords.EntityData.Children = make(map[string]types.YChild)
    lspPasswords.EntityData.Children["lsp-password"] = types.YChild{"LspPassword", nil}
    for i := range lspPasswords.LspPassword {
        lspPasswords.EntityData.Children[types.GetSegmentPath(&lspPasswords.LspPassword[i])] = types.YChild{"LspPassword", &lspPasswords.LspPassword[i]}
    }
    lspPasswords.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspPasswords.EntityData)
}

// Isis_Instances_Instance_LspPasswords_LspPassword
// LSP/SNP passwords. This must exist if an
// LSPAcceptPassword of the same level exists.
type Isis_Instances_Instance_LspPasswords_LspPassword struct {
    EntityData types.CommonEntityData
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
    // b'(!.+)|([^!].+)'. This attribute is mandatory.
    Password interface{}
}

func (lspPassword *Isis_Instances_Instance_LspPasswords_LspPassword) GetEntityData() *types.CommonEntityData {
    lspPassword.EntityData.YFilter = lspPassword.YFilter
    lspPassword.EntityData.YangName = "lsp-password"
    lspPassword.EntityData.BundleName = "cisco_ios_xr"
    lspPassword.EntityData.ParentYangName = "lsp-passwords"
    lspPassword.EntityData.SegmentPath = "lsp-password" + "[level='" + fmt.Sprintf("%v", lspPassword.Level) + "']"
    lspPassword.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPassword.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPassword.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPassword.EntityData.Children = make(map[string]types.YChild)
    lspPassword.EntityData.Leafs = make(map[string]types.YLeaf)
    lspPassword.EntityData.Leafs["level"] = types.YLeaf{"Level", lspPassword.Level}
    lspPassword.EntityData.Leafs["algorithm"] = types.YLeaf{"Algorithm", lspPassword.Algorithm}
    lspPassword.EntityData.Leafs["failure-mode"] = types.YLeaf{"FailureMode", lspPassword.FailureMode}
    lspPassword.EntityData.Leafs["authentication-type"] = types.YLeaf{"AuthenticationType", lspPassword.AuthenticationType}
    lspPassword.EntityData.Leafs["password"] = types.YLeaf{"Password", lspPassword.Password}
    return &(lspPassword.EntityData)
}

// Isis_Instances_Instance_Nets
// NET configuration
type Isis_Instances_Instance_Nets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network Entity Title (NET). The type is slice of
    // Isis_Instances_Instance_Nets_Net.
    Net []Isis_Instances_Instance_Nets_Net
}

func (nets *Isis_Instances_Instance_Nets) GetEntityData() *types.CommonEntityData {
    nets.EntityData.YFilter = nets.YFilter
    nets.EntityData.YangName = "nets"
    nets.EntityData.BundleName = "cisco_ios_xr"
    nets.EntityData.ParentYangName = "instance"
    nets.EntityData.SegmentPath = "nets"
    nets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nets.EntityData.Children = make(map[string]types.YChild)
    nets.EntityData.Children["net"] = types.YChild{"Net", nil}
    for i := range nets.Net {
        nets.EntityData.Children[types.GetSegmentPath(&nets.Net[i])] = types.YChild{"Net", &nets.Net[i]}
    }
    nets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nets.EntityData)
}

// Isis_Instances_Instance_Nets_Net
// Network Entity Title (NET)
type Isis_Instances_Instance_Nets_Net struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Network Entity Title. The type is string with
    // pattern: b'[a-fA-F0-9]{2}(\\.[a-fA-F0-9]{4}){3,9}\\.[a-fA-F0-9]{2}'.
    NetName interface{}
}

func (net *Isis_Instances_Instance_Nets_Net) GetEntityData() *types.CommonEntityData {
    net.EntityData.YFilter = net.YFilter
    net.EntityData.YangName = "net"
    net.EntityData.BundleName = "cisco_ios_xr"
    net.EntityData.ParentYangName = "nets"
    net.EntityData.SegmentPath = "net" + "[net-name='" + fmt.Sprintf("%v", net.NetName) + "']"
    net.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    net.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    net.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    net.EntityData.Children = make(map[string]types.YChild)
    net.EntityData.Leafs = make(map[string]types.YLeaf)
    net.EntityData.Leafs["net-name"] = types.YLeaf{"NetName", net.NetName}
    return &(net.EntityData)
}

// Isis_Instances_Instance_LspLifetimes
// LSP lifetime configuration
type Isis_Instances_Instance_LspLifetimes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum LSP lifetime. The type is slice of
    // Isis_Instances_Instance_LspLifetimes_LspLifetime.
    LspLifetime []Isis_Instances_Instance_LspLifetimes_LspLifetime
}

func (lspLifetimes *Isis_Instances_Instance_LspLifetimes) GetEntityData() *types.CommonEntityData {
    lspLifetimes.EntityData.YFilter = lspLifetimes.YFilter
    lspLifetimes.EntityData.YangName = "lsp-lifetimes"
    lspLifetimes.EntityData.BundleName = "cisco_ios_xr"
    lspLifetimes.EntityData.ParentYangName = "instance"
    lspLifetimes.EntityData.SegmentPath = "lsp-lifetimes"
    lspLifetimes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspLifetimes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspLifetimes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspLifetimes.EntityData.Children = make(map[string]types.YChild)
    lspLifetimes.EntityData.Children["lsp-lifetime"] = types.YChild{"LspLifetime", nil}
    for i := range lspLifetimes.LspLifetime {
        lspLifetimes.EntityData.Children[types.GetSegmentPath(&lspLifetimes.LspLifetime[i])] = types.YChild{"LspLifetime", &lspLifetimes.LspLifetime[i]}
    }
    lspLifetimes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspLifetimes.EntityData)
}

// Isis_Instances_Instance_LspLifetimes_LspLifetime
// Maximum LSP lifetime
type Isis_Instances_Instance_LspLifetimes_LspLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 1..65535. This attribute is
    // mandatory. Units are second.
    Lifetime interface{}
}

func (lspLifetime *Isis_Instances_Instance_LspLifetimes_LspLifetime) GetEntityData() *types.CommonEntityData {
    lspLifetime.EntityData.YFilter = lspLifetime.YFilter
    lspLifetime.EntityData.YangName = "lsp-lifetime"
    lspLifetime.EntityData.BundleName = "cisco_ios_xr"
    lspLifetime.EntityData.ParentYangName = "lsp-lifetimes"
    lspLifetime.EntityData.SegmentPath = "lsp-lifetime" + "[level='" + fmt.Sprintf("%v", lspLifetime.Level) + "']"
    lspLifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspLifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspLifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspLifetime.EntityData.Children = make(map[string]types.YChild)
    lspLifetime.EntityData.Leafs = make(map[string]types.YLeaf)
    lspLifetime.EntityData.Leafs["level"] = types.YLeaf{"Level", lspLifetime.Level}
    lspLifetime.EntityData.Leafs["lifetime"] = types.YLeaf{"Lifetime", lspLifetime.Lifetime}
    return &(lspLifetime.EntityData)
}

// Isis_Instances_Instance_OverloadBits
// LSP overload-bit configuration
type Isis_Instances_Instance_OverloadBits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the overload bit in the System LSP so that other routers avoid this one
    // in SPF calculations. This may be done either unconditionally, or on startup
    // until either a set time has passed or IS-IS is informed that BGP has
    // converged. This is an Object with a union discriminated on an integer value
    // of the ISISOverloadBitModeType. The type is slice of
    // Isis_Instances_Instance_OverloadBits_OverloadBit.
    OverloadBit []Isis_Instances_Instance_OverloadBits_OverloadBit
}

func (overloadBits *Isis_Instances_Instance_OverloadBits) GetEntityData() *types.CommonEntityData {
    overloadBits.EntityData.YFilter = overloadBits.YFilter
    overloadBits.EntityData.YangName = "overload-bits"
    overloadBits.EntityData.BundleName = "cisco_ios_xr"
    overloadBits.EntityData.ParentYangName = "instance"
    overloadBits.EntityData.SegmentPath = "overload-bits"
    overloadBits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overloadBits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overloadBits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overloadBits.EntityData.Children = make(map[string]types.YChild)
    overloadBits.EntityData.Children["overload-bit"] = types.YChild{"OverloadBit", nil}
    for i := range overloadBits.OverloadBit {
        overloadBits.EntityData.Children[types.GetSegmentPath(&overloadBits.OverloadBit[i])] = types.YChild{"OverloadBit", &overloadBits.OverloadBit[i]}
    }
    overloadBits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(overloadBits.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (overloadBit *Isis_Instances_Instance_OverloadBits_OverloadBit) GetEntityData() *types.CommonEntityData {
    overloadBit.EntityData.YFilter = overloadBit.YFilter
    overloadBit.EntityData.YangName = "overload-bit"
    overloadBit.EntityData.BundleName = "cisco_ios_xr"
    overloadBit.EntityData.ParentYangName = "overload-bits"
    overloadBit.EntityData.SegmentPath = "overload-bit" + "[level='" + fmt.Sprintf("%v", overloadBit.Level) + "']"
    overloadBit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overloadBit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overloadBit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overloadBit.EntityData.Children = make(map[string]types.YChild)
    overloadBit.EntityData.Leafs = make(map[string]types.YLeaf)
    overloadBit.EntityData.Leafs["level"] = types.YLeaf{"Level", overloadBit.Level}
    overloadBit.EntityData.Leafs["overload-bit-mode"] = types.YLeaf{"OverloadBitMode", overloadBit.OverloadBitMode}
    overloadBit.EntityData.Leafs["hippity-period"] = types.YLeaf{"HippityPeriod", overloadBit.HippityPeriod}
    overloadBit.EntityData.Leafs["external-adv-type"] = types.YLeaf{"ExternalAdvType", overloadBit.ExternalAdvType}
    overloadBit.EntityData.Leafs["inter-level-adv-type"] = types.YLeaf{"InterLevelAdvType", overloadBit.InterLevelAdvType}
    return &(overloadBit.EntityData)
}

// Isis_Instances_Instance_Interfaces
// Per-interface configuration
type Isis_Instances_Instance_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for an IS-IS interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_.
    Interface_ []Isis_Instances_Instance_Interfaces_Interface
}

func (interfaces *Isis_Instances_Instance_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "instance"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface
// Configuration for an IS-IS interface
type Isis_Instances_Instance_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (self *Isis_Instances_Instance_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["lsp-retransmit-throttle-intervals"] = types.YChild{"LspRetransmitThrottleIntervals", &self.LspRetransmitThrottleIntervals}
    self.EntityData.Children["lsp-retransmit-intervals"] = types.YChild{"LspRetransmitIntervals", &self.LspRetransmitIntervals}
    self.EntityData.Children["bfd"] = types.YChild{"Bfd", &self.Bfd}
    self.EntityData.Children["priorities"] = types.YChild{"Priorities", &self.Priorities}
    self.EntityData.Children["hello-accept-passwords"] = types.YChild{"HelloAcceptPasswords", &self.HelloAcceptPasswords}
    self.EntityData.Children["hello-passwords"] = types.YChild{"HelloPasswords", &self.HelloPasswords}
    self.EntityData.Children["hello-paddings"] = types.YChild{"HelloPaddings", &self.HelloPaddings}
    self.EntityData.Children["hello-multipliers"] = types.YChild{"HelloMultipliers", &self.HelloMultipliers}
    self.EntityData.Children["lsp-fast-flood-thresholds"] = types.YChild{"LspFastFloodThresholds", &self.LspFastFloodThresholds}
    self.EntityData.Children["prefix-attribute-n-flag-clears"] = types.YChild{"PrefixAttributeNFlagClears", &self.PrefixAttributeNFlagClears}
    self.EntityData.Children["hello-intervals"] = types.YChild{"HelloIntervals", &self.HelloIntervals}
    self.EntityData.Children["interface-afs"] = types.YChild{"InterfaceAfs", &self.InterfaceAfs}
    self.EntityData.Children["csnp-intervals"] = types.YChild{"CsnpIntervals", &self.CsnpIntervals}
    self.EntityData.Children["lsp-intervals"] = types.YChild{"LspIntervals", &self.LspIntervals}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["running"] = types.YLeaf{"Running", self.Running}
    self.EntityData.Leafs["circuit-type"] = types.YLeaf{"CircuitType", self.CircuitType}
    self.EntityData.Leafs["point-to-point"] = types.YLeaf{"PointToPoint", self.PointToPoint}
    self.EntityData.Leafs["state"] = types.YLeaf{"State", self.State}
    self.EntityData.Leafs["mesh-group"] = types.YLeaf{"MeshGroup", self.MeshGroup}
    self.EntityData.Leafs["link-down-fast-detect"] = types.YLeaf{"LinkDownFastDetect", self.LinkDownFastDetect}
    return &(self.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals
// LSP-retransmission-throttle-interval
// configuration
type Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum interval betwen retransissions of different LSPs. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval.
    LspRetransmitThrottleInterval []Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval
}

func (lspRetransmitThrottleIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals) GetEntityData() *types.CommonEntityData {
    lspRetransmitThrottleIntervals.EntityData.YFilter = lspRetransmitThrottleIntervals.YFilter
    lspRetransmitThrottleIntervals.EntityData.YangName = "lsp-retransmit-throttle-intervals"
    lspRetransmitThrottleIntervals.EntityData.BundleName = "cisco_ios_xr"
    lspRetransmitThrottleIntervals.EntityData.ParentYangName = "interface"
    lspRetransmitThrottleIntervals.EntityData.SegmentPath = "lsp-retransmit-throttle-intervals"
    lspRetransmitThrottleIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRetransmitThrottleIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRetransmitThrottleIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRetransmitThrottleIntervals.EntityData.Children = make(map[string]types.YChild)
    lspRetransmitThrottleIntervals.EntityData.Children["lsp-retransmit-throttle-interval"] = types.YChild{"LspRetransmitThrottleInterval", nil}
    for i := range lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval {
        lspRetransmitThrottleIntervals.EntityData.Children[types.GetSegmentPath(&lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval[i])] = types.YChild{"LspRetransmitThrottleInterval", &lspRetransmitThrottleIntervals.LspRetransmitThrottleInterval[i]}
    }
    lspRetransmitThrottleIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspRetransmitThrottleIntervals.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval
// Minimum interval betwen retransissions of
// different LSPs
type Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Milliseconds. The type is interface{} with range: 0..65535. This attribute
    // is mandatory. Units are millisecond.
    Interval interface{}
}

func (lspRetransmitThrottleInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitThrottleIntervals_LspRetransmitThrottleInterval) GetEntityData() *types.CommonEntityData {
    lspRetransmitThrottleInterval.EntityData.YFilter = lspRetransmitThrottleInterval.YFilter
    lspRetransmitThrottleInterval.EntityData.YangName = "lsp-retransmit-throttle-interval"
    lspRetransmitThrottleInterval.EntityData.BundleName = "cisco_ios_xr"
    lspRetransmitThrottleInterval.EntityData.ParentYangName = "lsp-retransmit-throttle-intervals"
    lspRetransmitThrottleInterval.EntityData.SegmentPath = "lsp-retransmit-throttle-interval" + "[level='" + fmt.Sprintf("%v", lspRetransmitThrottleInterval.Level) + "']"
    lspRetransmitThrottleInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRetransmitThrottleInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRetransmitThrottleInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRetransmitThrottleInterval.EntityData.Children = make(map[string]types.YChild)
    lspRetransmitThrottleInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    lspRetransmitThrottleInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", lspRetransmitThrottleInterval.Level}
    lspRetransmitThrottleInterval.EntityData.Leafs["interval"] = types.YLeaf{"Interval", lspRetransmitThrottleInterval.Interval}
    return &(lspRetransmitThrottleInterval.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals
// LSP-retransmission-interval configuration
type Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval between retransmissions of the same LSP. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval.
    LspRetransmitInterval []Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval
}

func (lspRetransmitIntervals *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals) GetEntityData() *types.CommonEntityData {
    lspRetransmitIntervals.EntityData.YFilter = lspRetransmitIntervals.YFilter
    lspRetransmitIntervals.EntityData.YangName = "lsp-retransmit-intervals"
    lspRetransmitIntervals.EntityData.BundleName = "cisco_ios_xr"
    lspRetransmitIntervals.EntityData.ParentYangName = "interface"
    lspRetransmitIntervals.EntityData.SegmentPath = "lsp-retransmit-intervals"
    lspRetransmitIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRetransmitIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRetransmitIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRetransmitIntervals.EntityData.Children = make(map[string]types.YChild)
    lspRetransmitIntervals.EntityData.Children["lsp-retransmit-interval"] = types.YChild{"LspRetransmitInterval", nil}
    for i := range lspRetransmitIntervals.LspRetransmitInterval {
        lspRetransmitIntervals.EntityData.Children[types.GetSegmentPath(&lspRetransmitIntervals.LspRetransmitInterval[i])] = types.YChild{"LspRetransmitInterval", &lspRetransmitIntervals.LspRetransmitInterval[i]}
    }
    lspRetransmitIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspRetransmitIntervals.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval
// Interval between retransmissions of the
// same LSP
type Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 0..65535. This attribute is
    // mandatory. Units are second.
    Interval interface{}
}

func (lspRetransmitInterval *Isis_Instances_Instance_Interfaces_Interface_LspRetransmitIntervals_LspRetransmitInterval) GetEntityData() *types.CommonEntityData {
    lspRetransmitInterval.EntityData.YFilter = lspRetransmitInterval.YFilter
    lspRetransmitInterval.EntityData.YangName = "lsp-retransmit-interval"
    lspRetransmitInterval.EntityData.BundleName = "cisco_ios_xr"
    lspRetransmitInterval.EntityData.ParentYangName = "lsp-retransmit-intervals"
    lspRetransmitInterval.EntityData.SegmentPath = "lsp-retransmit-interval" + "[level='" + fmt.Sprintf("%v", lspRetransmitInterval.Level) + "']"
    lspRetransmitInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspRetransmitInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspRetransmitInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspRetransmitInterval.EntityData.Children = make(map[string]types.YChild)
    lspRetransmitInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    lspRetransmitInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", lspRetransmitInterval.Level}
    lspRetransmitInterval.EntityData.Leafs["interval"] = types.YLeaf{"Interval", lspRetransmitInterval.Interval}
    return &(lspRetransmitInterval.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_Bfd
// BFD configuration
type Isis_Instances_Instance_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
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

func (bfd *Isis_Instances_Instance_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = make(map[string]types.YChild)
    bfd.EntityData.Leafs = make(map[string]types.YLeaf)
    bfd.EntityData.Leafs["enable-ipv6"] = types.YLeaf{"EnableIpv6", bfd.EnableIpv6}
    bfd.EntityData.Leafs["enable-ipv4"] = types.YLeaf{"EnableIpv4", bfd.EnableIpv4}
    bfd.EntityData.Leafs["interval"] = types.YLeaf{"Interval", bfd.Interval}
    bfd.EntityData.Leafs["detection-multiplier"] = types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier}
    return &(bfd.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_Priorities
// DIS-election priority configuration
type Isis_Instances_Instance_Interfaces_Interface_Priorities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DIS-election priority. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority.
    Priority []Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority
}

func (priorities *Isis_Instances_Instance_Interfaces_Interface_Priorities) GetEntityData() *types.CommonEntityData {
    priorities.EntityData.YFilter = priorities.YFilter
    priorities.EntityData.YangName = "priorities"
    priorities.EntityData.BundleName = "cisco_ios_xr"
    priorities.EntityData.ParentYangName = "interface"
    priorities.EntityData.SegmentPath = "priorities"
    priorities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priorities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priorities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priorities.EntityData.Children = make(map[string]types.YChild)
    priorities.EntityData.Children["priority"] = types.YChild{"Priority", nil}
    for i := range priorities.Priority {
        priorities.EntityData.Children[types.GetSegmentPath(&priorities.Priority[i])] = types.YChild{"Priority", &priorities.Priority[i]}
    }
    priorities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(priorities.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority
// DIS-election priority
type Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Priority. The type is interface{} with range: 0..127. This attribute is
    // mandatory.
    PriorityValue interface{}
}

func (priority *Isis_Instances_Instance_Interfaces_Interface_Priorities_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "priorities"
    priority.EntityData.SegmentPath = "priority" + "[level='" + fmt.Sprintf("%v", priority.Level) + "']"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = make(map[string]types.YChild)
    priority.EntityData.Leafs = make(map[string]types.YLeaf)
    priority.EntityData.Leafs["level"] = types.YLeaf{"Level", priority.Level}
    priority.EntityData.Leafs["priority-value"] = types.YLeaf{"PriorityValue", priority.PriorityValue}
    return &(priority.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords
// IIH accept password configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IIH accept passwords. This requires the existence of a HelloPassword of the
    // same level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword.
    HelloAcceptPassword []Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword
}

func (helloAcceptPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords) GetEntityData() *types.CommonEntityData {
    helloAcceptPasswords.EntityData.YFilter = helloAcceptPasswords.YFilter
    helloAcceptPasswords.EntityData.YangName = "hello-accept-passwords"
    helloAcceptPasswords.EntityData.BundleName = "cisco_ios_xr"
    helloAcceptPasswords.EntityData.ParentYangName = "interface"
    helloAcceptPasswords.EntityData.SegmentPath = "hello-accept-passwords"
    helloAcceptPasswords.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloAcceptPasswords.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloAcceptPasswords.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloAcceptPasswords.EntityData.Children = make(map[string]types.YChild)
    helloAcceptPasswords.EntityData.Children["hello-accept-password"] = types.YChild{"HelloAcceptPassword", nil}
    for i := range helloAcceptPasswords.HelloAcceptPassword {
        helloAcceptPasswords.EntityData.Children[types.GetSegmentPath(&helloAcceptPasswords.HelloAcceptPassword[i])] = types.YChild{"HelloAcceptPassword", &helloAcceptPasswords.HelloAcceptPassword[i]}
    }
    helloAcceptPasswords.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helloAcceptPasswords.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword
// IIH accept passwords. This requires the
// existence of a HelloPassword of the same
// level.
type Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Password. The type is string with pattern: b'(!.+)|([^!].+)'. This
    // attribute is mandatory.
    Password interface{}
}

func (helloAcceptPassword *Isis_Instances_Instance_Interfaces_Interface_HelloAcceptPasswords_HelloAcceptPassword) GetEntityData() *types.CommonEntityData {
    helloAcceptPassword.EntityData.YFilter = helloAcceptPassword.YFilter
    helloAcceptPassword.EntityData.YangName = "hello-accept-password"
    helloAcceptPassword.EntityData.BundleName = "cisco_ios_xr"
    helloAcceptPassword.EntityData.ParentYangName = "hello-accept-passwords"
    helloAcceptPassword.EntityData.SegmentPath = "hello-accept-password" + "[level='" + fmt.Sprintf("%v", helloAcceptPassword.Level) + "']"
    helloAcceptPassword.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloAcceptPassword.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloAcceptPassword.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloAcceptPassword.EntityData.Children = make(map[string]types.YChild)
    helloAcceptPassword.EntityData.Leafs = make(map[string]types.YLeaf)
    helloAcceptPassword.EntityData.Leafs["level"] = types.YLeaf{"Level", helloAcceptPassword.Level}
    helloAcceptPassword.EntityData.Leafs["password"] = types.YLeaf{"Password", helloAcceptPassword.Password}
    return &(helloAcceptPassword.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloPasswords
// IIH password configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloPasswords struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IIH passwords. This must exist if a HelloAcceptPassword of the same level
    // exists. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword.
    HelloPassword []Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword
}

func (helloPasswords *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords) GetEntityData() *types.CommonEntityData {
    helloPasswords.EntityData.YFilter = helloPasswords.YFilter
    helloPasswords.EntityData.YangName = "hello-passwords"
    helloPasswords.EntityData.BundleName = "cisco_ios_xr"
    helloPasswords.EntityData.ParentYangName = "interface"
    helloPasswords.EntityData.SegmentPath = "hello-passwords"
    helloPasswords.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloPasswords.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloPasswords.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloPasswords.EntityData.Children = make(map[string]types.YChild)
    helloPasswords.EntityData.Children["hello-password"] = types.YChild{"HelloPassword", nil}
    for i := range helloPasswords.HelloPassword {
        helloPasswords.EntityData.Children[types.GetSegmentPath(&helloPasswords.HelloPassword[i])] = types.YChild{"HelloPassword", &helloPasswords.HelloPassword[i]}
    }
    helloPasswords.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helloPasswords.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword
// IIH passwords. This must exist if a
// HelloAcceptPassword of the same level
// exists.
type Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword struct {
    EntityData types.CommonEntityData
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
    // b'(!.+)|([^!].+)'. This attribute is mandatory.
    Password interface{}
}

func (helloPassword *Isis_Instances_Instance_Interfaces_Interface_HelloPasswords_HelloPassword) GetEntityData() *types.CommonEntityData {
    helloPassword.EntityData.YFilter = helloPassword.YFilter
    helloPassword.EntityData.YangName = "hello-password"
    helloPassword.EntityData.BundleName = "cisco_ios_xr"
    helloPassword.EntityData.ParentYangName = "hello-passwords"
    helloPassword.EntityData.SegmentPath = "hello-password" + "[level='" + fmt.Sprintf("%v", helloPassword.Level) + "']"
    helloPassword.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloPassword.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloPassword.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloPassword.EntityData.Children = make(map[string]types.YChild)
    helloPassword.EntityData.Leafs = make(map[string]types.YLeaf)
    helloPassword.EntityData.Leafs["level"] = types.YLeaf{"Level", helloPassword.Level}
    helloPassword.EntityData.Leafs["algorithm"] = types.YLeaf{"Algorithm", helloPassword.Algorithm}
    helloPassword.EntityData.Leafs["failure-mode"] = types.YLeaf{"FailureMode", helloPassword.FailureMode}
    helloPassword.EntityData.Leafs["password"] = types.YLeaf{"Password", helloPassword.Password}
    return &(helloPassword.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloPaddings
// Hello-padding configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloPaddings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pad IIHs to the interface MTU. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding.
    HelloPadding []Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding
}

func (helloPaddings *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings) GetEntityData() *types.CommonEntityData {
    helloPaddings.EntityData.YFilter = helloPaddings.YFilter
    helloPaddings.EntityData.YangName = "hello-paddings"
    helloPaddings.EntityData.BundleName = "cisco_ios_xr"
    helloPaddings.EntityData.ParentYangName = "interface"
    helloPaddings.EntityData.SegmentPath = "hello-paddings"
    helloPaddings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloPaddings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloPaddings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloPaddings.EntityData.Children = make(map[string]types.YChild)
    helloPaddings.EntityData.Children["hello-padding"] = types.YChild{"HelloPadding", nil}
    for i := range helloPaddings.HelloPadding {
        helloPaddings.EntityData.Children[types.GetSegmentPath(&helloPaddings.HelloPadding[i])] = types.YChild{"HelloPadding", &helloPaddings.HelloPadding[i]}
    }
    helloPaddings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helloPaddings.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding
// Pad IIHs to the interface MTU
type Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Hello padding type value. The type is IsisHelloPadding. This attribute is
    // mandatory.
    PaddingType interface{}
}

func (helloPadding *Isis_Instances_Instance_Interfaces_Interface_HelloPaddings_HelloPadding) GetEntityData() *types.CommonEntityData {
    helloPadding.EntityData.YFilter = helloPadding.YFilter
    helloPadding.EntityData.YangName = "hello-padding"
    helloPadding.EntityData.BundleName = "cisco_ios_xr"
    helloPadding.EntityData.ParentYangName = "hello-paddings"
    helloPadding.EntityData.SegmentPath = "hello-padding" + "[level='" + fmt.Sprintf("%v", helloPadding.Level) + "']"
    helloPadding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloPadding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloPadding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloPadding.EntityData.Children = make(map[string]types.YChild)
    helloPadding.EntityData.Leafs = make(map[string]types.YLeaf)
    helloPadding.EntityData.Leafs["level"] = types.YLeaf{"Level", helloPadding.Level}
    helloPadding.EntityData.Leafs["padding-type"] = types.YLeaf{"PaddingType", helloPadding.PaddingType}
    return &(helloPadding.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers
// Hello-multiplier configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello-multiplier configuration. The number of successive IIHs that may be
    // missed on an adjacency before it is considered down. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier.
    HelloMultiplier []Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier
}

func (helloMultipliers *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers) GetEntityData() *types.CommonEntityData {
    helloMultipliers.EntityData.YFilter = helloMultipliers.YFilter
    helloMultipliers.EntityData.YangName = "hello-multipliers"
    helloMultipliers.EntityData.BundleName = "cisco_ios_xr"
    helloMultipliers.EntityData.ParentYangName = "interface"
    helloMultipliers.EntityData.SegmentPath = "hello-multipliers"
    helloMultipliers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloMultipliers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloMultipliers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloMultipliers.EntityData.Children = make(map[string]types.YChild)
    helloMultipliers.EntityData.Children["hello-multiplier"] = types.YChild{"HelloMultiplier", nil}
    for i := range helloMultipliers.HelloMultiplier {
        helloMultipliers.EntityData.Children[types.GetSegmentPath(&helloMultipliers.HelloMultiplier[i])] = types.YChild{"HelloMultiplier", &helloMultipliers.HelloMultiplier[i]}
    }
    helloMultipliers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helloMultipliers.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier
// Hello-multiplier configuration. The number
// of successive IIHs that may be missed on an
// adjacency before it is considered down.
type Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Hello multiplier value. The type is interface{} with range: 3..1000. This
    // attribute is mandatory.
    Multiplier interface{}
}

func (helloMultiplier *Isis_Instances_Instance_Interfaces_Interface_HelloMultipliers_HelloMultiplier) GetEntityData() *types.CommonEntityData {
    helloMultiplier.EntityData.YFilter = helloMultiplier.YFilter
    helloMultiplier.EntityData.YangName = "hello-multiplier"
    helloMultiplier.EntityData.BundleName = "cisco_ios_xr"
    helloMultiplier.EntityData.ParentYangName = "hello-multipliers"
    helloMultiplier.EntityData.SegmentPath = "hello-multiplier" + "[level='" + fmt.Sprintf("%v", helloMultiplier.Level) + "']"
    helloMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloMultiplier.EntityData.Children = make(map[string]types.YChild)
    helloMultiplier.EntityData.Leafs = make(map[string]types.YLeaf)
    helloMultiplier.EntityData.Leafs["level"] = types.YLeaf{"Level", helloMultiplier.Level}
    helloMultiplier.EntityData.Leafs["multiplier"] = types.YLeaf{"Multiplier", helloMultiplier.Multiplier}
    return &(helloMultiplier.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds
// LSP fast flood threshold configuration
type Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of LSPs to send back to back on an interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold.
    LspFastFloodThreshold []Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold
}

func (lspFastFloodThresholds *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds) GetEntityData() *types.CommonEntityData {
    lspFastFloodThresholds.EntityData.YFilter = lspFastFloodThresholds.YFilter
    lspFastFloodThresholds.EntityData.YangName = "lsp-fast-flood-thresholds"
    lspFastFloodThresholds.EntityData.BundleName = "cisco_ios_xr"
    lspFastFloodThresholds.EntityData.ParentYangName = "interface"
    lspFastFloodThresholds.EntityData.SegmentPath = "lsp-fast-flood-thresholds"
    lspFastFloodThresholds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspFastFloodThresholds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspFastFloodThresholds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspFastFloodThresholds.EntityData.Children = make(map[string]types.YChild)
    lspFastFloodThresholds.EntityData.Children["lsp-fast-flood-threshold"] = types.YChild{"LspFastFloodThreshold", nil}
    for i := range lspFastFloodThresholds.LspFastFloodThreshold {
        lspFastFloodThresholds.EntityData.Children[types.GetSegmentPath(&lspFastFloodThresholds.LspFastFloodThreshold[i])] = types.YChild{"LspFastFloodThreshold", &lspFastFloodThresholds.LspFastFloodThreshold[i]}
    }
    lspFastFloodThresholds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspFastFloodThresholds.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold
// Number of LSPs to send back to back on an
// interface.
type Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Count. The type is interface{} with range: 1..4294967295. This attribute is
    // mandatory.
    Count interface{}
}

func (lspFastFloodThreshold *Isis_Instances_Instance_Interfaces_Interface_LspFastFloodThresholds_LspFastFloodThreshold) GetEntityData() *types.CommonEntityData {
    lspFastFloodThreshold.EntityData.YFilter = lspFastFloodThreshold.YFilter
    lspFastFloodThreshold.EntityData.YangName = "lsp-fast-flood-threshold"
    lspFastFloodThreshold.EntityData.BundleName = "cisco_ios_xr"
    lspFastFloodThreshold.EntityData.ParentYangName = "lsp-fast-flood-thresholds"
    lspFastFloodThreshold.EntityData.SegmentPath = "lsp-fast-flood-threshold" + "[level='" + fmt.Sprintf("%v", lspFastFloodThreshold.Level) + "']"
    lspFastFloodThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspFastFloodThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspFastFloodThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspFastFloodThreshold.EntityData.Children = make(map[string]types.YChild)
    lspFastFloodThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    lspFastFloodThreshold.EntityData.Leafs["level"] = types.YLeaf{"Level", lspFastFloodThreshold.Level}
    lspFastFloodThreshold.EntityData.Leafs["count"] = types.YLeaf{"Count", lspFastFloodThreshold.Count}
    return &(lspFastFloodThreshold.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears
// Prefix attribute N flag clear configuration
type Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear the N flag in prefix attribute flags sub-TLV. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear.
    PrefixAttributeNFlagClear []Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear
}

func (prefixAttributeNFlagClears *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears) GetEntityData() *types.CommonEntityData {
    prefixAttributeNFlagClears.EntityData.YFilter = prefixAttributeNFlagClears.YFilter
    prefixAttributeNFlagClears.EntityData.YangName = "prefix-attribute-n-flag-clears"
    prefixAttributeNFlagClears.EntityData.BundleName = "cisco_ios_xr"
    prefixAttributeNFlagClears.EntityData.ParentYangName = "interface"
    prefixAttributeNFlagClears.EntityData.SegmentPath = "prefix-attribute-n-flag-clears"
    prefixAttributeNFlagClears.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixAttributeNFlagClears.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixAttributeNFlagClears.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixAttributeNFlagClears.EntityData.Children = make(map[string]types.YChild)
    prefixAttributeNFlagClears.EntityData.Children["prefix-attribute-n-flag-clear"] = types.YChild{"PrefixAttributeNFlagClear", nil}
    for i := range prefixAttributeNFlagClears.PrefixAttributeNFlagClear {
        prefixAttributeNFlagClears.EntityData.Children[types.GetSegmentPath(&prefixAttributeNFlagClears.PrefixAttributeNFlagClear[i])] = types.YChild{"PrefixAttributeNFlagClear", &prefixAttributeNFlagClears.PrefixAttributeNFlagClear[i]}
    }
    prefixAttributeNFlagClears.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixAttributeNFlagClears.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear
// Clear the N flag in prefix attribute flags
// sub-TLV
type Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (prefixAttributeNFlagClear *Isis_Instances_Instance_Interfaces_Interface_PrefixAttributeNFlagClears_PrefixAttributeNFlagClear) GetEntityData() *types.CommonEntityData {
    prefixAttributeNFlagClear.EntityData.YFilter = prefixAttributeNFlagClear.YFilter
    prefixAttributeNFlagClear.EntityData.YangName = "prefix-attribute-n-flag-clear"
    prefixAttributeNFlagClear.EntityData.BundleName = "cisco_ios_xr"
    prefixAttributeNFlagClear.EntityData.ParentYangName = "prefix-attribute-n-flag-clears"
    prefixAttributeNFlagClear.EntityData.SegmentPath = "prefix-attribute-n-flag-clear" + "[level='" + fmt.Sprintf("%v", prefixAttributeNFlagClear.Level) + "']"
    prefixAttributeNFlagClear.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixAttributeNFlagClear.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixAttributeNFlagClear.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixAttributeNFlagClear.EntityData.Children = make(map[string]types.YChild)
    prefixAttributeNFlagClear.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixAttributeNFlagClear.EntityData.Leafs["level"] = types.YLeaf{"Level", prefixAttributeNFlagClear.Level}
    return &(prefixAttributeNFlagClear.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloIntervals
// Hello-interval configuration
type Isis_Instances_Instance_Interfaces_Interface_HelloIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello-interval configuration. The interval at which IIH packets will be
    // sent. This will be three times quicker on a LAN interface which has been
    // electted DIS. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval.
    HelloInterval []Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval
}

func (helloIntervals *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals) GetEntityData() *types.CommonEntityData {
    helloIntervals.EntityData.YFilter = helloIntervals.YFilter
    helloIntervals.EntityData.YangName = "hello-intervals"
    helloIntervals.EntityData.BundleName = "cisco_ios_xr"
    helloIntervals.EntityData.ParentYangName = "interface"
    helloIntervals.EntityData.SegmentPath = "hello-intervals"
    helloIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloIntervals.EntityData.Children = make(map[string]types.YChild)
    helloIntervals.EntityData.Children["hello-interval"] = types.YChild{"HelloInterval", nil}
    for i := range helloIntervals.HelloInterval {
        helloIntervals.EntityData.Children[types.GetSegmentPath(&helloIntervals.HelloInterval[i])] = types.YChild{"HelloInterval", &helloIntervals.HelloInterval[i]}
    }
    helloIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(helloIntervals.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval
// Hello-interval configuration. The interval
// at which IIH packets will be sent. This
// will be three times quicker on a LAN
// interface which has been electted DIS.
type Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 1..65535. This attribute is
    // mandatory. Units are second.
    Interval interface{}
}

func (helloInterval *Isis_Instances_Instance_Interfaces_Interface_HelloIntervals_HelloInterval) GetEntityData() *types.CommonEntityData {
    helloInterval.EntityData.YFilter = helloInterval.YFilter
    helloInterval.EntityData.YangName = "hello-interval"
    helloInterval.EntityData.BundleName = "cisco_ios_xr"
    helloInterval.EntityData.ParentYangName = "hello-intervals"
    helloInterval.EntityData.SegmentPath = "hello-interval" + "[level='" + fmt.Sprintf("%v", helloInterval.Level) + "']"
    helloInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helloInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helloInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helloInterval.EntityData.Children = make(map[string]types.YChild)
    helloInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    helloInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", helloInterval.Level}
    helloInterval.EntityData.Leafs["interval"] = types.YLeaf{"Interval", helloInterval.Interval}
    return &(helloInterval.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs
// Per-interface address-family configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for an IS-IS address-family on a single interface. If a named
    // (non-default) topology is being created it must be multicast. Also the
    // topology ID mustbe set first and delete last in the router configuration.
    // The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf.
    InterfaceAf []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf
}

func (interfaceAfs *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs) GetEntityData() *types.CommonEntityData {
    interfaceAfs.EntityData.YFilter = interfaceAfs.YFilter
    interfaceAfs.EntityData.YangName = "interface-afs"
    interfaceAfs.EntityData.BundleName = "cisco_ios_xr"
    interfaceAfs.EntityData.ParentYangName = "interface"
    interfaceAfs.EntityData.SegmentPath = "interface-afs"
    interfaceAfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAfs.EntityData.Children = make(map[string]types.YChild)
    interfaceAfs.EntityData.Children["interface-af"] = types.YChild{"InterfaceAf", nil}
    for i := range interfaceAfs.InterfaceAf {
        interfaceAfs.EntityData.Children[types.GetSegmentPath(&interfaceAfs.InterfaceAf[i])] = types.YChild{"InterfaceAf", &interfaceAfs.InterfaceAf[i]}
    }
    interfaceAfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceAfs.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf
// Configuration for an IS-IS address-family
// on a single interface. If a named
// (non-default) topology is being created it
// must be multicast. Also the topology ID
// mustbe set first and delete last in the
// router configuration.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf struct {
    EntityData types.CommonEntityData
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

func (interfaceAf *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf) GetEntityData() *types.CommonEntityData {
    interfaceAf.EntityData.YFilter = interfaceAf.YFilter
    interfaceAf.EntityData.YangName = "interface-af"
    interfaceAf.EntityData.BundleName = "cisco_ios_xr"
    interfaceAf.EntityData.ParentYangName = "interface-afs"
    interfaceAf.EntityData.SegmentPath = "interface-af" + "[af-name='" + fmt.Sprintf("%v", interfaceAf.AfName) + "']" + "[saf-name='" + fmt.Sprintf("%v", interfaceAf.SafName) + "']"
    interfaceAf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAf.EntityData.Children = make(map[string]types.YChild)
    interfaceAf.EntityData.Children["interface-af-data"] = types.YChild{"InterfaceAfData", &interfaceAf.InterfaceAfData}
    interfaceAf.EntityData.Children["topology-name"] = types.YChild{"TopologyName", nil}
    for i := range interfaceAf.TopologyName {
        interfaceAf.EntityData.Children[types.GetSegmentPath(&interfaceAf.TopologyName[i])] = types.YChild{"TopologyName", &interfaceAf.TopologyName[i]}
    }
    interfaceAf.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceAf.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", interfaceAf.AfName}
    interfaceAf.EntityData.Leafs["saf-name"] = types.YLeaf{"SafName", interfaceAf.SafName}
    return &(interfaceAf.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData
// Data container.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData struct {
    EntityData types.CommonEntityData
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

func (interfaceAfData *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData) GetEntityData() *types.CommonEntityData {
    interfaceAfData.EntityData.YFilter = interfaceAfData.YFilter
    interfaceAfData.EntityData.YangName = "interface-af-data"
    interfaceAfData.EntityData.BundleName = "cisco_ios_xr"
    interfaceAfData.EntityData.ParentYangName = "interface-af"
    interfaceAfData.EntityData.SegmentPath = "interface-af-data"
    interfaceAfData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAfData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAfData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAfData.EntityData.Children = make(map[string]types.YChild)
    interfaceAfData.EntityData.Children["prefix-sid"] = types.YChild{"PrefixSid", &interfaceAfData.PrefixSid}
    interfaceAfData.EntityData.Children["interface-frr-table"] = types.YChild{"InterfaceFrrTable", &interfaceAfData.InterfaceFrrTable}
    interfaceAfData.EntityData.Children["mpls-ldp"] = types.YChild{"MplsLdp", &interfaceAfData.MplsLdp}
    interfaceAfData.EntityData.Children["prefix-sspfsid"] = types.YChild{"PrefixSspfsid", &interfaceAfData.PrefixSspfsid}
    interfaceAfData.EntityData.Children["auto-metrics"] = types.YChild{"AutoMetrics", &interfaceAfData.AutoMetrics}
    interfaceAfData.EntityData.Children["admin-tags"] = types.YChild{"AdminTags", &interfaceAfData.AdminTags}
    interfaceAfData.EntityData.Children["interface-link-group"] = types.YChild{"InterfaceLinkGroup", &interfaceAfData.InterfaceLinkGroup}
    interfaceAfData.EntityData.Children["manual-adj-sids"] = types.YChild{"ManualAdjSids", &interfaceAfData.ManualAdjSids}
    interfaceAfData.EntityData.Children["metrics"] = types.YChild{"Metrics", &interfaceAfData.Metrics}
    interfaceAfData.EntityData.Children["weights"] = types.YChild{"Weights", &interfaceAfData.Weights}
    interfaceAfData.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceAfData.EntityData.Leafs["interface-af-state"] = types.YLeaf{"InterfaceAfState", interfaceAfData.InterfaceAfState}
    interfaceAfData.EntityData.Leafs["running"] = types.YLeaf{"Running", interfaceAfData.Running}
    return &(interfaceAfData.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid
// Assign prefix SID to an interface,
// ISISPHPFlag will be rejected if set to
// disable, ISISEXPLICITNULLFlag will
// override the value of ISISPHPFlag
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type for the interface. The type is Isissid1. This attribute is
    // mandatory.
    Type_ interface{}

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

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSid) GetEntityData() *types.CommonEntityData {
    prefixSid.EntityData.YFilter = prefixSid.YFilter
    prefixSid.EntityData.YangName = "prefix-sid"
    prefixSid.EntityData.BundleName = "cisco_ios_xr"
    prefixSid.EntityData.ParentYangName = "interface-af-data"
    prefixSid.EntityData.SegmentPath = "prefix-sid"
    prefixSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSid.EntityData.Children = make(map[string]types.YChild)
    prefixSid.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixSid.EntityData.Leafs["type"] = types.YLeaf{"Type_", prefixSid.Type_}
    prefixSid.EntityData.Leafs["value"] = types.YLeaf{"Value", prefixSid.Value}
    prefixSid.EntityData.Leafs["php"] = types.YLeaf{"Php", prefixSid.Php}
    prefixSid.EntityData.Leafs["explicit-null"] = types.YLeaf{"ExplicitNull", prefixSid.ExplicitNull}
    prefixSid.EntityData.Leafs["nflag-clear"] = types.YLeaf{"NflagClear", prefixSid.NflagClear}
    return &(prefixSid.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable
// Fast-ReRoute configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable struct {
    EntityData types.CommonEntityData
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

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable) GetEntityData() *types.CommonEntityData {
    interfaceFrrTable.EntityData.YFilter = interfaceFrrTable.YFilter
    interfaceFrrTable.EntityData.YangName = "interface-frr-table"
    interfaceFrrTable.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTable.EntityData.ParentYangName = "interface-af-data"
    interfaceFrrTable.EntityData.SegmentPath = "interface-frr-table"
    interfaceFrrTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTable.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTable.EntityData.Children["frrlfa-candidate-interfaces"] = types.YChild{"FrrlfaCandidateInterfaces", &interfaceFrrTable.FrrlfaCandidateInterfaces}
    interfaceFrrTable.EntityData.Children["frr-remote-lfa-max-metrics"] = types.YChild{"FrrRemoteLfaMaxMetrics", &interfaceFrrTable.FrrRemoteLfaMaxMetrics}
    interfaceFrrTable.EntityData.Children["frr-types"] = types.YChild{"FrrTypes", &interfaceFrrTable.FrrTypes}
    interfaceFrrTable.EntityData.Children["frr-remote-lfa-types"] = types.YChild{"FrrRemoteLfaTypes", &interfaceFrrTable.FrrRemoteLfaTypes}
    interfaceFrrTable.EntityData.Children["interface-frr-tiebreaker-defaults"] = types.YChild{"InterfaceFrrTiebreakerDefaults", &interfaceFrrTable.InterfaceFrrTiebreakerDefaults}
    interfaceFrrTable.EntityData.Children["frrtilfa-types"] = types.YChild{"FrrtilfaTypes", &interfaceFrrTable.FrrtilfaTypes}
    interfaceFrrTable.EntityData.Children["frr-exclude-interfaces"] = types.YChild{"FrrExcludeInterfaces", &interfaceFrrTable.FrrExcludeInterfaces}
    interfaceFrrTable.EntityData.Children["interface-frr-tiebreakers"] = types.YChild{"InterfaceFrrTiebreakers", &interfaceFrrTable.InterfaceFrrTiebreakers}
    interfaceFrrTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceFrrTable.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces
// FRR LFA candidate configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Include an interface to LFA candidate in computation. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface.
    FrrlfaCandidateInterface []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetEntityData() *types.CommonEntityData {
    frrlfaCandidateInterfaces.EntityData.YFilter = frrlfaCandidateInterfaces.YFilter
    frrlfaCandidateInterfaces.EntityData.YangName = "frrlfa-candidate-interfaces"
    frrlfaCandidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    frrlfaCandidateInterfaces.EntityData.ParentYangName = "interface-frr-table"
    frrlfaCandidateInterfaces.EntityData.SegmentPath = "frrlfa-candidate-interfaces"
    frrlfaCandidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrlfaCandidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrlfaCandidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrlfaCandidateInterfaces.EntityData.Children = make(map[string]types.YChild)
    frrlfaCandidateInterfaces.EntityData.Children["frrlfa-candidate-interface"] = types.YChild{"FrrlfaCandidateInterface", nil}
    for i := range frrlfaCandidateInterfaces.FrrlfaCandidateInterface {
        frrlfaCandidateInterfaces.EntityData.Children[types.GetSegmentPath(&frrlfaCandidateInterfaces.FrrlfaCandidateInterface[i])] = types.YChild{"FrrlfaCandidateInterface", &frrlfaCandidateInterfaces.FrrlfaCandidateInterface[i]}
    }
    frrlfaCandidateInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrlfaCandidateInterfaces.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface
// Include an interface to LFA candidate
// in computation
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Level. The type is interface{} with range: 0..2. This attribute is
    // mandatory.
    Level interface{}
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetEntityData() *types.CommonEntityData {
    frrlfaCandidateInterface.EntityData.YFilter = frrlfaCandidateInterface.YFilter
    frrlfaCandidateInterface.EntityData.YangName = "frrlfa-candidate-interface"
    frrlfaCandidateInterface.EntityData.BundleName = "cisco_ios_xr"
    frrlfaCandidateInterface.EntityData.ParentYangName = "frrlfa-candidate-interfaces"
    frrlfaCandidateInterface.EntityData.SegmentPath = "frrlfa-candidate-interface" + "[interface-name='" + fmt.Sprintf("%v", frrlfaCandidateInterface.InterfaceName) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrlfaCandidateInterface.FrrType) + "']"
    frrlfaCandidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrlfaCandidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrlfaCandidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrlfaCandidateInterface.EntityData.Children = make(map[string]types.YChild)
    frrlfaCandidateInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    frrlfaCandidateInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", frrlfaCandidateInterface.InterfaceName}
    frrlfaCandidateInterface.EntityData.Leafs["frr-type"] = types.YLeaf{"FrrType", frrlfaCandidateInterface.FrrType}
    frrlfaCandidateInterface.EntityData.Leafs["level"] = types.YLeaf{"Level", frrlfaCandidateInterface.Level}
    return &(frrlfaCandidateInterface.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics
// Remote LFA maxmimum metric
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum metric for selecting a remote LFA node. The type is
    // slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric.
    FrrRemoteLfaMaxMetric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaMaxMetrics.EntityData.YFilter = frrRemoteLfaMaxMetrics.YFilter
    frrRemoteLfaMaxMetrics.EntityData.YangName = "frr-remote-lfa-max-metrics"
    frrRemoteLfaMaxMetrics.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaMaxMetrics.EntityData.ParentYangName = "interface-frr-table"
    frrRemoteLfaMaxMetrics.EntityData.SegmentPath = "frr-remote-lfa-max-metrics"
    frrRemoteLfaMaxMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaMaxMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaMaxMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaMaxMetrics.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaMaxMetrics.EntityData.Children["frr-remote-lfa-max-metric"] = types.YChild{"FrrRemoteLfaMaxMetric", nil}
    for i := range frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric {
        frrRemoteLfaMaxMetrics.EntityData.Children[types.GetSegmentPath(&frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[i])] = types.YChild{"FrrRemoteLfaMaxMetric", &frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[i]}
    }
    frrRemoteLfaMaxMetrics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrRemoteLfaMaxMetrics.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric
// Configure the maximum metric for
// selecting a remote LFA node
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Value of the metric. The type is interface{} with range: 1..16777215. This
    // attribute is mandatory.
    MaxMetric interface{}
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaMaxMetric.EntityData.YFilter = frrRemoteLfaMaxMetric.YFilter
    frrRemoteLfaMaxMetric.EntityData.YangName = "frr-remote-lfa-max-metric"
    frrRemoteLfaMaxMetric.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaMaxMetric.EntityData.ParentYangName = "frr-remote-lfa-max-metrics"
    frrRemoteLfaMaxMetric.EntityData.SegmentPath = "frr-remote-lfa-max-metric" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaMaxMetric.Level) + "']"
    frrRemoteLfaMaxMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaMaxMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaMaxMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaMaxMetric.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaMaxMetric.EntityData.Leafs = make(map[string]types.YLeaf)
    frrRemoteLfaMaxMetric.EntityData.Leafs["level"] = types.YLeaf{"Level", frrRemoteLfaMaxMetric.Level}
    frrRemoteLfaMaxMetric.EntityData.Leafs["max-metric"] = types.YLeaf{"MaxMetric", frrRemoteLfaMaxMetric.MaxMetric}
    return &(frrRemoteLfaMaxMetric.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes
// Type of FRR computation per level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of computation for prefixes reachable via interface. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType.
    FrrType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes) GetEntityData() *types.CommonEntityData {
    frrTypes.EntityData.YFilter = frrTypes.YFilter
    frrTypes.EntityData.YangName = "frr-types"
    frrTypes.EntityData.BundleName = "cisco_ios_xr"
    frrTypes.EntityData.ParentYangName = "interface-frr-table"
    frrTypes.EntityData.SegmentPath = "frr-types"
    frrTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrTypes.EntityData.Children = make(map[string]types.YChild)
    frrTypes.EntityData.Children["frr-type"] = types.YChild{"FrrType", nil}
    for i := range frrTypes.FrrType {
        frrTypes.EntityData.Children[types.GetSegmentPath(&frrTypes.FrrType[i])] = types.YChild{"FrrType", &frrTypes.FrrType[i]}
    }
    frrTypes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrTypes.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType
// Type of computation for prefixes
// reachable via interface
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Computation Type. The type is Isisfrr. This attribute is mandatory.
    Type_ interface{}
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrTypes_FrrType) GetEntityData() *types.CommonEntityData {
    frrType.EntityData.YFilter = frrType.YFilter
    frrType.EntityData.YangName = "frr-type"
    frrType.EntityData.BundleName = "cisco_ios_xr"
    frrType.EntityData.ParentYangName = "frr-types"
    frrType.EntityData.SegmentPath = "frr-type" + "[level='" + fmt.Sprintf("%v", frrType.Level) + "']"
    frrType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrType.EntityData.Children = make(map[string]types.YChild)
    frrType.EntityData.Leafs = make(map[string]types.YLeaf)
    frrType.EntityData.Leafs["level"] = types.YLeaf{"Level", frrType.Level}
    frrType.EntityData.Leafs["type"] = types.YLeaf{"Type_", frrType.Type_}
    return &(frrType.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes
// Remote LFA Enable
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable remote lfa for a particular level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType.
    FrrRemoteLfaType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaTypes.EntityData.YFilter = frrRemoteLfaTypes.YFilter
    frrRemoteLfaTypes.EntityData.YangName = "frr-remote-lfa-types"
    frrRemoteLfaTypes.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaTypes.EntityData.ParentYangName = "interface-frr-table"
    frrRemoteLfaTypes.EntityData.SegmentPath = "frr-remote-lfa-types"
    frrRemoteLfaTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaTypes.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaTypes.EntityData.Children["frr-remote-lfa-type"] = types.YChild{"FrrRemoteLfaType", nil}
    for i := range frrRemoteLfaTypes.FrrRemoteLfaType {
        frrRemoteLfaTypes.EntityData.Children[types.GetSegmentPath(&frrRemoteLfaTypes.FrrRemoteLfaType[i])] = types.YChild{"FrrRemoteLfaType", &frrRemoteLfaTypes.FrrRemoteLfaType[i]}
    }
    frrRemoteLfaTypes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrRemoteLfaTypes.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType
// Enable remote lfa for a particular
// level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Remote LFA Type. The type is IsisRemoteLfa. This attribute is mandatory.
    Type_ interface{}
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaType.EntityData.YFilter = frrRemoteLfaType.YFilter
    frrRemoteLfaType.EntityData.YangName = "frr-remote-lfa-type"
    frrRemoteLfaType.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaType.EntityData.ParentYangName = "frr-remote-lfa-types"
    frrRemoteLfaType.EntityData.SegmentPath = "frr-remote-lfa-type" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaType.Level) + "']"
    frrRemoteLfaType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaType.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaType.EntityData.Leafs = make(map[string]types.YLeaf)
    frrRemoteLfaType.EntityData.Leafs["level"] = types.YLeaf{"Level", frrRemoteLfaType.Level}
    frrRemoteLfaType.EntityData.Leafs["type"] = types.YLeaf{"Type_", frrRemoteLfaType.Type_}
    return &(frrRemoteLfaType.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults
// Interface FRR Default tiebreaker
// configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure default tiebreaker. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault.
    InterfaceFrrTiebreakerDefault []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetEntityData() *types.CommonEntityData {
    interfaceFrrTiebreakerDefaults.EntityData.YFilter = interfaceFrrTiebreakerDefaults.YFilter
    interfaceFrrTiebreakerDefaults.EntityData.YangName = "interface-frr-tiebreaker-defaults"
    interfaceFrrTiebreakerDefaults.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTiebreakerDefaults.EntityData.ParentYangName = "interface-frr-table"
    interfaceFrrTiebreakerDefaults.EntityData.SegmentPath = "interface-frr-tiebreaker-defaults"
    interfaceFrrTiebreakerDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTiebreakerDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTiebreakerDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTiebreakerDefaults.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTiebreakerDefaults.EntityData.Children["interface-frr-tiebreaker-default"] = types.YChild{"InterfaceFrrTiebreakerDefault", nil}
    for i := range interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault {
        interfaceFrrTiebreakerDefaults.EntityData.Children[types.GetSegmentPath(&interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[i])] = types.YChild{"InterfaceFrrTiebreakerDefault", &interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[i]}
    }
    interfaceFrrTiebreakerDefaults.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceFrrTiebreakerDefaults.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault
// Configure default tiebreaker
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetEntityData() *types.CommonEntityData {
    interfaceFrrTiebreakerDefault.EntityData.YFilter = interfaceFrrTiebreakerDefault.YFilter
    interfaceFrrTiebreakerDefault.EntityData.YangName = "interface-frr-tiebreaker-default"
    interfaceFrrTiebreakerDefault.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTiebreakerDefault.EntityData.ParentYangName = "interface-frr-tiebreaker-defaults"
    interfaceFrrTiebreakerDefault.EntityData.SegmentPath = "interface-frr-tiebreaker-default" + "[level='" + fmt.Sprintf("%v", interfaceFrrTiebreakerDefault.Level) + "']"
    interfaceFrrTiebreakerDefault.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTiebreakerDefault.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTiebreakerDefault.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTiebreakerDefault.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTiebreakerDefault.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceFrrTiebreakerDefault.EntityData.Leafs["level"] = types.YLeaf{"Level", interfaceFrrTiebreakerDefault.Level}
    return &(interfaceFrrTiebreakerDefault.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes
// TI LFA Enable
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable TI lfa for a particular level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType.
    FrrtilfaType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes) GetEntityData() *types.CommonEntityData {
    frrtilfaTypes.EntityData.YFilter = frrtilfaTypes.YFilter
    frrtilfaTypes.EntityData.YangName = "frrtilfa-types"
    frrtilfaTypes.EntityData.BundleName = "cisco_ios_xr"
    frrtilfaTypes.EntityData.ParentYangName = "interface-frr-table"
    frrtilfaTypes.EntityData.SegmentPath = "frrtilfa-types"
    frrtilfaTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrtilfaTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrtilfaTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrtilfaTypes.EntityData.Children = make(map[string]types.YChild)
    frrtilfaTypes.EntityData.Children["frrtilfa-type"] = types.YChild{"FrrtilfaType", nil}
    for i := range frrtilfaTypes.FrrtilfaType {
        frrtilfaTypes.EntityData.Children[types.GetSegmentPath(&frrtilfaTypes.FrrtilfaType[i])] = types.YChild{"FrrtilfaType", &frrtilfaTypes.FrrtilfaType[i]}
    }
    frrtilfaTypes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrtilfaTypes.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType
// Enable TI lfa for a particular level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetEntityData() *types.CommonEntityData {
    frrtilfaType.EntityData.YFilter = frrtilfaType.YFilter
    frrtilfaType.EntityData.YangName = "frrtilfa-type"
    frrtilfaType.EntityData.BundleName = "cisco_ios_xr"
    frrtilfaType.EntityData.ParentYangName = "frrtilfa-types"
    frrtilfaType.EntityData.SegmentPath = "frrtilfa-type" + "[level='" + fmt.Sprintf("%v", frrtilfaType.Level) + "']"
    frrtilfaType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrtilfaType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrtilfaType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrtilfaType.EntityData.Children = make(map[string]types.YChild)
    frrtilfaType.EntityData.Leafs = make(map[string]types.YLeaf)
    frrtilfaType.EntityData.Leafs["level"] = types.YLeaf{"Level", frrtilfaType.Level}
    return &(frrtilfaType.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces
// FRR exclusion configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from computation. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface.
    FrrExcludeInterface []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    frrExcludeInterfaces.EntityData.YFilter = frrExcludeInterfaces.YFilter
    frrExcludeInterfaces.EntityData.YangName = "frr-exclude-interfaces"
    frrExcludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    frrExcludeInterfaces.EntityData.ParentYangName = "interface-frr-table"
    frrExcludeInterfaces.EntityData.SegmentPath = "frr-exclude-interfaces"
    frrExcludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrExcludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrExcludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrExcludeInterfaces.EntityData.Children = make(map[string]types.YChild)
    frrExcludeInterfaces.EntityData.Children["frr-exclude-interface"] = types.YChild{"FrrExcludeInterface", nil}
    for i := range frrExcludeInterfaces.FrrExcludeInterface {
        frrExcludeInterfaces.EntityData.Children[types.GetSegmentPath(&frrExcludeInterfaces.FrrExcludeInterface[i])] = types.YChild{"FrrExcludeInterface", &frrExcludeInterfaces.FrrExcludeInterface[i]}
    }
    frrExcludeInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrExcludeInterfaces.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface
// Exclude an interface from computation
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Level. The type is interface{} with range: 0..2. This attribute is
    // mandatory.
    Level interface{}
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetEntityData() *types.CommonEntityData {
    frrExcludeInterface.EntityData.YFilter = frrExcludeInterface.YFilter
    frrExcludeInterface.EntityData.YangName = "frr-exclude-interface"
    frrExcludeInterface.EntityData.BundleName = "cisco_ios_xr"
    frrExcludeInterface.EntityData.ParentYangName = "frr-exclude-interfaces"
    frrExcludeInterface.EntityData.SegmentPath = "frr-exclude-interface" + "[interface-name='" + fmt.Sprintf("%v", frrExcludeInterface.InterfaceName) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrExcludeInterface.FrrType) + "']"
    frrExcludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrExcludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrExcludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrExcludeInterface.EntityData.Children = make(map[string]types.YChild)
    frrExcludeInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    frrExcludeInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", frrExcludeInterface.InterfaceName}
    frrExcludeInterface.EntityData.Leafs["frr-type"] = types.YLeaf{"FrrType", frrExcludeInterface.FrrType}
    frrExcludeInterface.EntityData.Leafs["level"] = types.YLeaf{"Level", frrExcludeInterface.Level}
    return &(frrExcludeInterface.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers
// Interface FRR tiebreakers configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure tiebreaker for multiple backups. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker.
    InterfaceFrrTiebreaker []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers) GetEntityData() *types.CommonEntityData {
    interfaceFrrTiebreakers.EntityData.YFilter = interfaceFrrTiebreakers.YFilter
    interfaceFrrTiebreakers.EntityData.YangName = "interface-frr-tiebreakers"
    interfaceFrrTiebreakers.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTiebreakers.EntityData.ParentYangName = "interface-frr-table"
    interfaceFrrTiebreakers.EntityData.SegmentPath = "interface-frr-tiebreakers"
    interfaceFrrTiebreakers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTiebreakers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTiebreakers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTiebreakers.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTiebreakers.EntityData.Children["interface-frr-tiebreaker"] = types.YChild{"InterfaceFrrTiebreaker", nil}
    for i := range interfaceFrrTiebreakers.InterfaceFrrTiebreaker {
        interfaceFrrTiebreakers.EntityData.Children[types.GetSegmentPath(&interfaceFrrTiebreakers.InterfaceFrrTiebreaker[i])] = types.YChild{"InterfaceFrrTiebreaker", &interfaceFrrTiebreakers.InterfaceFrrTiebreaker[i]}
    }
    interfaceFrrTiebreakers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceFrrTiebreakers.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker
// Configure tiebreaker for multiple
// backups
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker struct {
    EntityData types.CommonEntityData
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

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetEntityData() *types.CommonEntityData {
    interfaceFrrTiebreaker.EntityData.YFilter = interfaceFrrTiebreaker.YFilter
    interfaceFrrTiebreaker.EntityData.YangName = "interface-frr-tiebreaker"
    interfaceFrrTiebreaker.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTiebreaker.EntityData.ParentYangName = "interface-frr-tiebreakers"
    interfaceFrrTiebreaker.EntityData.SegmentPath = "interface-frr-tiebreaker" + "[level='" + fmt.Sprintf("%v", interfaceFrrTiebreaker.Level) + "']" + "[tiebreaker='" + fmt.Sprintf("%v", interfaceFrrTiebreaker.Tiebreaker) + "']"
    interfaceFrrTiebreaker.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTiebreaker.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTiebreaker.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTiebreaker.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTiebreaker.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceFrrTiebreaker.EntityData.Leafs["level"] = types.YLeaf{"Level", interfaceFrrTiebreaker.Level}
    interfaceFrrTiebreaker.EntityData.Leafs["tiebreaker"] = types.YLeaf{"Tiebreaker", interfaceFrrTiebreaker.Tiebreaker}
    interfaceFrrTiebreaker.EntityData.Leafs["index"] = types.YLeaf{"Index", interfaceFrrTiebreaker.Index}
    return &(interfaceFrrTiebreaker.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp
// MPLS LDP configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable MPLS LDP Synchronization for an IS-IS level. The type is interface{}
    // with range: 0..2. The default value is 0.
    SyncLevel interface{}
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_MplsLdp) GetEntityData() *types.CommonEntityData {
    mplsLdp.EntityData.YFilter = mplsLdp.YFilter
    mplsLdp.EntityData.YangName = "mpls-ldp"
    mplsLdp.EntityData.BundleName = "cisco_ios_xr"
    mplsLdp.EntityData.ParentYangName = "interface-af-data"
    mplsLdp.EntityData.SegmentPath = "mpls-ldp"
    mplsLdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLdp.EntityData.Children = make(map[string]types.YChild)
    mplsLdp.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsLdp.EntityData.Leafs["sync-level"] = types.YLeaf{"SyncLevel", mplsLdp.SyncLevel}
    return &(mplsLdp.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid
// Assign prefix SSPF SID to an interface,
// ISISPHPFlag will be rejected if set to
// disable, ISISEXPLICITNULLFlag will
// override the value of ISISPHPFlag
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type for the interface. The type is Isissid1. This attribute is
    // mandatory.
    Type_ interface{}

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

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_PrefixSspfsid) GetEntityData() *types.CommonEntityData {
    prefixSspfsid.EntityData.YFilter = prefixSspfsid.YFilter
    prefixSspfsid.EntityData.YangName = "prefix-sspfsid"
    prefixSspfsid.EntityData.BundleName = "cisco_ios_xr"
    prefixSspfsid.EntityData.ParentYangName = "interface-af-data"
    prefixSspfsid.EntityData.SegmentPath = "prefix-sspfsid"
    prefixSspfsid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSspfsid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSspfsid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSspfsid.EntityData.Children = make(map[string]types.YChild)
    prefixSspfsid.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixSspfsid.EntityData.Leafs["type"] = types.YLeaf{"Type_", prefixSspfsid.Type_}
    prefixSspfsid.EntityData.Leafs["value"] = types.YLeaf{"Value", prefixSspfsid.Value}
    prefixSspfsid.EntityData.Leafs["php"] = types.YLeaf{"Php", prefixSspfsid.Php}
    prefixSspfsid.EntityData.Leafs["explicit-null"] = types.YLeaf{"ExplicitNull", prefixSspfsid.ExplicitNull}
    prefixSspfsid.EntityData.Leafs["nflag-clear"] = types.YLeaf{"NflagClear", prefixSspfsid.NflagClear}
    return &(prefixSspfsid.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics
// AutoMetric configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoMetric Proactive-Protect configuration. Legal value depends on the
    // metric-style specified for the topology. If the metric-style defined is
    // narrow, then only a value between <1-63> is allowed and if the metric-style
    // is defined as wide, then a value between <1-16777214> is allowed as the
    // auto-metric value. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric.
    AutoMetric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics) GetEntityData() *types.CommonEntityData {
    autoMetrics.EntityData.YFilter = autoMetrics.YFilter
    autoMetrics.EntityData.YangName = "auto-metrics"
    autoMetrics.EntityData.BundleName = "cisco_ios_xr"
    autoMetrics.EntityData.ParentYangName = "interface-af-data"
    autoMetrics.EntityData.SegmentPath = "auto-metrics"
    autoMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoMetrics.EntityData.Children = make(map[string]types.YChild)
    autoMetrics.EntityData.Children["auto-metric"] = types.YChild{"AutoMetric", nil}
    for i := range autoMetrics.AutoMetric {
        autoMetrics.EntityData.Children[types.GetSegmentPath(&autoMetrics.AutoMetric[i])] = types.YChild{"AutoMetric", &autoMetrics.AutoMetric[i]}
    }
    autoMetrics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(autoMetrics.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed auto metric:<1-63> for narrow ,<1-16777214> for wide. The type is
    // interface{} with range: 1..16777214. This attribute is mandatory.
    ProactiveProtect interface{}
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AutoMetrics_AutoMetric) GetEntityData() *types.CommonEntityData {
    autoMetric.EntityData.YFilter = autoMetric.YFilter
    autoMetric.EntityData.YangName = "auto-metric"
    autoMetric.EntityData.BundleName = "cisco_ios_xr"
    autoMetric.EntityData.ParentYangName = "auto-metrics"
    autoMetric.EntityData.SegmentPath = "auto-metric" + "[level='" + fmt.Sprintf("%v", autoMetric.Level) + "']"
    autoMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoMetric.EntityData.Children = make(map[string]types.YChild)
    autoMetric.EntityData.Leafs = make(map[string]types.YLeaf)
    autoMetric.EntityData.Leafs["level"] = types.YLeaf{"Level", autoMetric.Level}
    autoMetric.EntityData.Leafs["proactive-protect"] = types.YLeaf{"ProactiveProtect", autoMetric.ProactiveProtect}
    return &(autoMetric.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags
// admin-tag configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Admin tag for advertised interface connected routes. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag.
    AdminTag []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags) GetEntityData() *types.CommonEntityData {
    adminTags.EntityData.YFilter = adminTags.YFilter
    adminTags.EntityData.YangName = "admin-tags"
    adminTags.EntityData.BundleName = "cisco_ios_xr"
    adminTags.EntityData.ParentYangName = "interface-af-data"
    adminTags.EntityData.SegmentPath = "admin-tags"
    adminTags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminTags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminTags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminTags.EntityData.Children = make(map[string]types.YChild)
    adminTags.EntityData.Children["admin-tag"] = types.YChild{"AdminTag", nil}
    for i := range adminTags.AdminTag {
        adminTags.EntityData.Children[types.GetSegmentPath(&adminTags.AdminTag[i])] = types.YChild{"AdminTag", &adminTags.AdminTag[i]}
    }
    adminTags.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adminTags.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag
// Admin tag for advertised interface
// connected routes
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Tag to associate with connected routes. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    AdminTag interface{}
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_AdminTags_AdminTag) GetEntityData() *types.CommonEntityData {
    adminTag.EntityData.YFilter = adminTag.YFilter
    adminTag.EntityData.YangName = "admin-tag"
    adminTag.EntityData.BundleName = "cisco_ios_xr"
    adminTag.EntityData.ParentYangName = "admin-tags"
    adminTag.EntityData.SegmentPath = "admin-tag" + "[level='" + fmt.Sprintf("%v", adminTag.Level) + "']"
    adminTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminTag.EntityData.Children = make(map[string]types.YChild)
    adminTag.EntityData.Leafs = make(map[string]types.YLeaf)
    adminTag.EntityData.Leafs["level"] = types.YLeaf{"Level", adminTag.Level}
    adminTag.EntityData.Leafs["admin-tag"] = types.YLeaf{"AdminTag", adminTag.AdminTag}
    return &(adminTag.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup
// Provide link group name and level
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link Group. The type is string with length: 1..40. This attribute is
    // mandatory.
    LinkGroup interface{}

    // Level in which link group will be effective. The type is interface{} with
    // range: 0..2. The default value is 0.
    Level interface{}
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_InterfaceLinkGroup) GetEntityData() *types.CommonEntityData {
    interfaceLinkGroup.EntityData.YFilter = interfaceLinkGroup.YFilter
    interfaceLinkGroup.EntityData.YangName = "interface-link-group"
    interfaceLinkGroup.EntityData.BundleName = "cisco_ios_xr"
    interfaceLinkGroup.EntityData.ParentYangName = "interface-af-data"
    interfaceLinkGroup.EntityData.SegmentPath = "interface-link-group"
    interfaceLinkGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLinkGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLinkGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLinkGroup.EntityData.Children = make(map[string]types.YChild)
    interfaceLinkGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceLinkGroup.EntityData.Leafs["link-group"] = types.YLeaf{"LinkGroup", interfaceLinkGroup.LinkGroup}
    interfaceLinkGroup.EntityData.Leafs["level"] = types.YLeaf{"Level", interfaceLinkGroup.Level}
    return &(interfaceLinkGroup.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids
// Manual Adjacecy SID configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Assign adjancency SID to an interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid.
    ManualAdjSid []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids) GetEntityData() *types.CommonEntityData {
    manualAdjSids.EntityData.YFilter = manualAdjSids.YFilter
    manualAdjSids.EntityData.YangName = "manual-adj-sids"
    manualAdjSids.EntityData.BundleName = "cisco_ios_xr"
    manualAdjSids.EntityData.ParentYangName = "interface-af-data"
    manualAdjSids.EntityData.SegmentPath = "manual-adj-sids"
    manualAdjSids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualAdjSids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualAdjSids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualAdjSids.EntityData.Children = make(map[string]types.YChild)
    manualAdjSids.EntityData.Children["manual-adj-sid"] = types.YChild{"ManualAdjSid", nil}
    for i := range manualAdjSids.ManualAdjSid {
        manualAdjSids.EntityData.Children[types.GetSegmentPath(&manualAdjSids.ManualAdjSid[i])] = types.YChild{"ManualAdjSid", &manualAdjSids.ManualAdjSid[i]}
    }
    manualAdjSids.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(manualAdjSids.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid
// Assign adjancency SID to an interface
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid struct {
    EntityData types.CommonEntityData
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

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_ManualAdjSids_ManualAdjSid) GetEntityData() *types.CommonEntityData {
    manualAdjSid.EntityData.YFilter = manualAdjSid.YFilter
    manualAdjSid.EntityData.YangName = "manual-adj-sid"
    manualAdjSid.EntityData.BundleName = "cisco_ios_xr"
    manualAdjSid.EntityData.ParentYangName = "manual-adj-sids"
    manualAdjSid.EntityData.SegmentPath = "manual-adj-sid" + "[level='" + fmt.Sprintf("%v", manualAdjSid.Level) + "']" + "[sid-type='" + fmt.Sprintf("%v", manualAdjSid.SidType) + "']" + "[sid='" + fmt.Sprintf("%v", manualAdjSid.Sid) + "']"
    manualAdjSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualAdjSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualAdjSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualAdjSid.EntityData.Children = make(map[string]types.YChild)
    manualAdjSid.EntityData.Leafs = make(map[string]types.YLeaf)
    manualAdjSid.EntityData.Leafs["level"] = types.YLeaf{"Level", manualAdjSid.Level}
    manualAdjSid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", manualAdjSid.SidType}
    manualAdjSid.EntityData.Leafs["sid"] = types.YLeaf{"Sid", manualAdjSid.Sid}
    manualAdjSid.EntityData.Leafs["protected"] = types.YLeaf{"Protected", manualAdjSid.Protected}
    return &(manualAdjSid.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics
// Metric configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics struct {
    EntityData types.CommonEntityData
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

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics) GetEntityData() *types.CommonEntityData {
    metrics.EntityData.YFilter = metrics.YFilter
    metrics.EntityData.YangName = "metrics"
    metrics.EntityData.BundleName = "cisco_ios_xr"
    metrics.EntityData.ParentYangName = "interface-af-data"
    metrics.EntityData.SegmentPath = "metrics"
    metrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metrics.EntityData.Children = make(map[string]types.YChild)
    metrics.EntityData.Children["metric"] = types.YChild{"Metric", nil}
    for i := range metrics.Metric {
        metrics.EntityData.Children[types.GetSegmentPath(&metrics.Metric[i])] = types.YChild{"Metric", &metrics.Metric[i]}
    }
    metrics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(metrics.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed metric: <1-63> for narrow, <1-16777215> for wide. The type is one
    // of the following types: enumeration
    // Isis.Instances.Instance.Interfaces.Interface.InterfaceAfs.InterfaceAf.TopologyName.Metrics.Metric.Metric_
    // This attribute is mandatory., or int with range: 1..16777215 This attribute
    // is mandatory..
    Metric interface{}
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "metrics"
    metric.EntityData.SegmentPath = "metric" + "[level='" + fmt.Sprintf("%v", metric.Level) + "']"
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = make(map[string]types.YChild)
    metric.EntityData.Leafs = make(map[string]types.YLeaf)
    metric.EntityData.Leafs["level"] = types.YLeaf{"Level", metric.Level}
    metric.EntityData.Leafs["metric"] = types.YLeaf{"Metric", metric.Metric}
    return &(metric.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric_Metric_ represents <1-16777215> for wide
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric_Metric_ string

const (
    // Maximum wide metric.  All routers will
    // exclude this link from their SPF
    Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric_Metric__maximum Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Metrics_Metric_Metric_ = "maximum"
)

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights
// Weight configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Weight configuration under interface for load balancing. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight.
    Weight []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights) GetEntityData() *types.CommonEntityData {
    weights.EntityData.YFilter = weights.YFilter
    weights.EntityData.YangName = "weights"
    weights.EntityData.BundleName = "cisco_ios_xr"
    weights.EntityData.ParentYangName = "interface-af-data"
    weights.EntityData.SegmentPath = "weights"
    weights.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    weights.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    weights.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    weights.EntityData.Children = make(map[string]types.YChild)
    weights.EntityData.Children["weight"] = types.YChild{"Weight", nil}
    for i := range weights.Weight {
        weights.EntityData.Children[types.GetSegmentPath(&weights.Weight[i])] = types.YChild{"Weight", &weights.Weight[i]}
    }
    weights.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(weights.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight
// Weight configuration under interface for load
// balancing
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Weight to be configured under interface for Load Balancing. Allowed weight:
    // <1-16777215>. The type is interface{} with range: 1..16777214. This
    // attribute is mandatory.
    Weight interface{}
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_InterfaceAfData_Weights_Weight) GetEntityData() *types.CommonEntityData {
    weight.EntityData.YFilter = weight.YFilter
    weight.EntityData.YangName = "weight"
    weight.EntityData.BundleName = "cisco_ios_xr"
    weight.EntityData.ParentYangName = "weights"
    weight.EntityData.SegmentPath = "weight" + "[level='" + fmt.Sprintf("%v", weight.Level) + "']"
    weight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    weight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    weight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    weight.EntityData.Children = make(map[string]types.YChild)
    weight.EntityData.Leafs = make(map[string]types.YLeaf)
    weight.EntityData.Leafs["level"] = types.YLeaf{"Level", weight.Level}
    weight.EntityData.Leafs["weight"] = types.YLeaf{"Weight", weight.Weight}
    return &(weight.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName
// keys: topology-name
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName struct {
    EntityData types.CommonEntityData
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

func (topologyName *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName) GetEntityData() *types.CommonEntityData {
    topologyName.EntityData.YFilter = topologyName.YFilter
    topologyName.EntityData.YangName = "topology-name"
    topologyName.EntityData.BundleName = "cisco_ios_xr"
    topologyName.EntityData.ParentYangName = "interface-af"
    topologyName.EntityData.SegmentPath = "topology-name" + "[topology-name='" + fmt.Sprintf("%v", topologyName.TopologyName) + "']"
    topologyName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyName.EntityData.Children = make(map[string]types.YChild)
    topologyName.EntityData.Children["prefix-sid"] = types.YChild{"PrefixSid", &topologyName.PrefixSid}
    topologyName.EntityData.Children["interface-frr-table"] = types.YChild{"InterfaceFrrTable", &topologyName.InterfaceFrrTable}
    topologyName.EntityData.Children["mpls-ldp"] = types.YChild{"MplsLdp", &topologyName.MplsLdp}
    topologyName.EntityData.Children["prefix-sspfsid"] = types.YChild{"PrefixSspfsid", &topologyName.PrefixSspfsid}
    topologyName.EntityData.Children["auto-metrics"] = types.YChild{"AutoMetrics", &topologyName.AutoMetrics}
    topologyName.EntityData.Children["admin-tags"] = types.YChild{"AdminTags", &topologyName.AdminTags}
    topologyName.EntityData.Children["interface-link-group"] = types.YChild{"InterfaceLinkGroup", &topologyName.InterfaceLinkGroup}
    topologyName.EntityData.Children["manual-adj-sids"] = types.YChild{"ManualAdjSids", &topologyName.ManualAdjSids}
    topologyName.EntityData.Children["metrics"] = types.YChild{"Metrics", &topologyName.Metrics}
    topologyName.EntityData.Children["weights"] = types.YChild{"Weights", &topologyName.Weights}
    topologyName.EntityData.Leafs = make(map[string]types.YLeaf)
    topologyName.EntityData.Leafs["topology-name"] = types.YLeaf{"TopologyName", topologyName.TopologyName}
    topologyName.EntityData.Leafs["interface-af-state"] = types.YLeaf{"InterfaceAfState", topologyName.InterfaceAfState}
    topologyName.EntityData.Leafs["running"] = types.YLeaf{"Running", topologyName.Running}
    return &(topologyName.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid
// Assign prefix SID to an interface,
// ISISPHPFlag will be rejected if set to
// disable, ISISEXPLICITNULLFlag will
// override the value of ISISPHPFlag
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type for the interface. The type is Isissid1. This attribute is
    // mandatory.
    Type_ interface{}

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

func (prefixSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSid) GetEntityData() *types.CommonEntityData {
    prefixSid.EntityData.YFilter = prefixSid.YFilter
    prefixSid.EntityData.YangName = "prefix-sid"
    prefixSid.EntityData.BundleName = "cisco_ios_xr"
    prefixSid.EntityData.ParentYangName = "topology-name"
    prefixSid.EntityData.SegmentPath = "prefix-sid"
    prefixSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSid.EntityData.Children = make(map[string]types.YChild)
    prefixSid.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixSid.EntityData.Leafs["type"] = types.YLeaf{"Type_", prefixSid.Type_}
    prefixSid.EntityData.Leafs["value"] = types.YLeaf{"Value", prefixSid.Value}
    prefixSid.EntityData.Leafs["php"] = types.YLeaf{"Php", prefixSid.Php}
    prefixSid.EntityData.Leafs["explicit-null"] = types.YLeaf{"ExplicitNull", prefixSid.ExplicitNull}
    prefixSid.EntityData.Leafs["nflag-clear"] = types.YLeaf{"NflagClear", prefixSid.NflagClear}
    return &(prefixSid.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable
// Fast-ReRoute configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable struct {
    EntityData types.CommonEntityData
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

func (interfaceFrrTable *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable) GetEntityData() *types.CommonEntityData {
    interfaceFrrTable.EntityData.YFilter = interfaceFrrTable.YFilter
    interfaceFrrTable.EntityData.YangName = "interface-frr-table"
    interfaceFrrTable.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTable.EntityData.ParentYangName = "topology-name"
    interfaceFrrTable.EntityData.SegmentPath = "interface-frr-table"
    interfaceFrrTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTable.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTable.EntityData.Children["frrlfa-candidate-interfaces"] = types.YChild{"FrrlfaCandidateInterfaces", &interfaceFrrTable.FrrlfaCandidateInterfaces}
    interfaceFrrTable.EntityData.Children["frr-remote-lfa-max-metrics"] = types.YChild{"FrrRemoteLfaMaxMetrics", &interfaceFrrTable.FrrRemoteLfaMaxMetrics}
    interfaceFrrTable.EntityData.Children["frr-types"] = types.YChild{"FrrTypes", &interfaceFrrTable.FrrTypes}
    interfaceFrrTable.EntityData.Children["frr-remote-lfa-types"] = types.YChild{"FrrRemoteLfaTypes", &interfaceFrrTable.FrrRemoteLfaTypes}
    interfaceFrrTable.EntityData.Children["interface-frr-tiebreaker-defaults"] = types.YChild{"InterfaceFrrTiebreakerDefaults", &interfaceFrrTable.InterfaceFrrTiebreakerDefaults}
    interfaceFrrTable.EntityData.Children["frrtilfa-types"] = types.YChild{"FrrtilfaTypes", &interfaceFrrTable.FrrtilfaTypes}
    interfaceFrrTable.EntityData.Children["frr-exclude-interfaces"] = types.YChild{"FrrExcludeInterfaces", &interfaceFrrTable.FrrExcludeInterfaces}
    interfaceFrrTable.EntityData.Children["interface-frr-tiebreakers"] = types.YChild{"InterfaceFrrTiebreakers", &interfaceFrrTable.InterfaceFrrTiebreakers}
    interfaceFrrTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceFrrTable.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces
// FRR LFA candidate configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Include an interface to LFA candidate in computation. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface.
    FrrlfaCandidateInterface []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface
}

func (frrlfaCandidateInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces) GetEntityData() *types.CommonEntityData {
    frrlfaCandidateInterfaces.EntityData.YFilter = frrlfaCandidateInterfaces.YFilter
    frrlfaCandidateInterfaces.EntityData.YangName = "frrlfa-candidate-interfaces"
    frrlfaCandidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    frrlfaCandidateInterfaces.EntityData.ParentYangName = "interface-frr-table"
    frrlfaCandidateInterfaces.EntityData.SegmentPath = "frrlfa-candidate-interfaces"
    frrlfaCandidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrlfaCandidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrlfaCandidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrlfaCandidateInterfaces.EntityData.Children = make(map[string]types.YChild)
    frrlfaCandidateInterfaces.EntityData.Children["frrlfa-candidate-interface"] = types.YChild{"FrrlfaCandidateInterface", nil}
    for i := range frrlfaCandidateInterfaces.FrrlfaCandidateInterface {
        frrlfaCandidateInterfaces.EntityData.Children[types.GetSegmentPath(&frrlfaCandidateInterfaces.FrrlfaCandidateInterface[i])] = types.YChild{"FrrlfaCandidateInterface", &frrlfaCandidateInterfaces.FrrlfaCandidateInterface[i]}
    }
    frrlfaCandidateInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrlfaCandidateInterfaces.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface
// Include an interface to LFA candidate
// in computation
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Level. The type is interface{} with range: 0..2. This attribute is
    // mandatory.
    Level interface{}
}

func (frrlfaCandidateInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrlfaCandidateInterfaces_FrrlfaCandidateInterface) GetEntityData() *types.CommonEntityData {
    frrlfaCandidateInterface.EntityData.YFilter = frrlfaCandidateInterface.YFilter
    frrlfaCandidateInterface.EntityData.YangName = "frrlfa-candidate-interface"
    frrlfaCandidateInterface.EntityData.BundleName = "cisco_ios_xr"
    frrlfaCandidateInterface.EntityData.ParentYangName = "frrlfa-candidate-interfaces"
    frrlfaCandidateInterface.EntityData.SegmentPath = "frrlfa-candidate-interface" + "[interface-name='" + fmt.Sprintf("%v", frrlfaCandidateInterface.InterfaceName) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrlfaCandidateInterface.FrrType) + "']"
    frrlfaCandidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrlfaCandidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrlfaCandidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrlfaCandidateInterface.EntityData.Children = make(map[string]types.YChild)
    frrlfaCandidateInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    frrlfaCandidateInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", frrlfaCandidateInterface.InterfaceName}
    frrlfaCandidateInterface.EntityData.Leafs["frr-type"] = types.YLeaf{"FrrType", frrlfaCandidateInterface.FrrType}
    frrlfaCandidateInterface.EntityData.Leafs["level"] = types.YLeaf{"Level", frrlfaCandidateInterface.Level}
    return &(frrlfaCandidateInterface.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics
// Remote LFA maxmimum metric
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the maximum metric for selecting a remote LFA node. The type is
    // slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric.
    FrrRemoteLfaMaxMetric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric
}

func (frrRemoteLfaMaxMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaMaxMetrics.EntityData.YFilter = frrRemoteLfaMaxMetrics.YFilter
    frrRemoteLfaMaxMetrics.EntityData.YangName = "frr-remote-lfa-max-metrics"
    frrRemoteLfaMaxMetrics.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaMaxMetrics.EntityData.ParentYangName = "interface-frr-table"
    frrRemoteLfaMaxMetrics.EntityData.SegmentPath = "frr-remote-lfa-max-metrics"
    frrRemoteLfaMaxMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaMaxMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaMaxMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaMaxMetrics.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaMaxMetrics.EntityData.Children["frr-remote-lfa-max-metric"] = types.YChild{"FrrRemoteLfaMaxMetric", nil}
    for i := range frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric {
        frrRemoteLfaMaxMetrics.EntityData.Children[types.GetSegmentPath(&frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[i])] = types.YChild{"FrrRemoteLfaMaxMetric", &frrRemoteLfaMaxMetrics.FrrRemoteLfaMaxMetric[i]}
    }
    frrRemoteLfaMaxMetrics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrRemoteLfaMaxMetrics.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric
// Configure the maximum metric for
// selecting a remote LFA node
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Value of the metric. The type is interface{} with range: 1..16777215. This
    // attribute is mandatory.
    MaxMetric interface{}
}

func (frrRemoteLfaMaxMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaMaxMetrics_FrrRemoteLfaMaxMetric) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaMaxMetric.EntityData.YFilter = frrRemoteLfaMaxMetric.YFilter
    frrRemoteLfaMaxMetric.EntityData.YangName = "frr-remote-lfa-max-metric"
    frrRemoteLfaMaxMetric.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaMaxMetric.EntityData.ParentYangName = "frr-remote-lfa-max-metrics"
    frrRemoteLfaMaxMetric.EntityData.SegmentPath = "frr-remote-lfa-max-metric" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaMaxMetric.Level) + "']"
    frrRemoteLfaMaxMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaMaxMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaMaxMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaMaxMetric.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaMaxMetric.EntityData.Leafs = make(map[string]types.YLeaf)
    frrRemoteLfaMaxMetric.EntityData.Leafs["level"] = types.YLeaf{"Level", frrRemoteLfaMaxMetric.Level}
    frrRemoteLfaMaxMetric.EntityData.Leafs["max-metric"] = types.YLeaf{"MaxMetric", frrRemoteLfaMaxMetric.MaxMetric}
    return &(frrRemoteLfaMaxMetric.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes
// Type of FRR computation per level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of computation for prefixes reachable via interface. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType.
    FrrType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType
}

func (frrTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes) GetEntityData() *types.CommonEntityData {
    frrTypes.EntityData.YFilter = frrTypes.YFilter
    frrTypes.EntityData.YangName = "frr-types"
    frrTypes.EntityData.BundleName = "cisco_ios_xr"
    frrTypes.EntityData.ParentYangName = "interface-frr-table"
    frrTypes.EntityData.SegmentPath = "frr-types"
    frrTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrTypes.EntityData.Children = make(map[string]types.YChild)
    frrTypes.EntityData.Children["frr-type"] = types.YChild{"FrrType", nil}
    for i := range frrTypes.FrrType {
        frrTypes.EntityData.Children[types.GetSegmentPath(&frrTypes.FrrType[i])] = types.YChild{"FrrType", &frrTypes.FrrType[i]}
    }
    frrTypes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrTypes.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType
// Type of computation for prefixes
// reachable via interface
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Computation Type. The type is Isisfrr. This attribute is mandatory.
    Type_ interface{}
}

func (frrType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrTypes_FrrType) GetEntityData() *types.CommonEntityData {
    frrType.EntityData.YFilter = frrType.YFilter
    frrType.EntityData.YangName = "frr-type"
    frrType.EntityData.BundleName = "cisco_ios_xr"
    frrType.EntityData.ParentYangName = "frr-types"
    frrType.EntityData.SegmentPath = "frr-type" + "[level='" + fmt.Sprintf("%v", frrType.Level) + "']"
    frrType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrType.EntityData.Children = make(map[string]types.YChild)
    frrType.EntityData.Leafs = make(map[string]types.YLeaf)
    frrType.EntityData.Leafs["level"] = types.YLeaf{"Level", frrType.Level}
    frrType.EntityData.Leafs["type"] = types.YLeaf{"Type_", frrType.Type_}
    return &(frrType.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes
// Remote LFA Enable
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable remote lfa for a particular level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType.
    FrrRemoteLfaType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType
}

func (frrRemoteLfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaTypes.EntityData.YFilter = frrRemoteLfaTypes.YFilter
    frrRemoteLfaTypes.EntityData.YangName = "frr-remote-lfa-types"
    frrRemoteLfaTypes.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaTypes.EntityData.ParentYangName = "interface-frr-table"
    frrRemoteLfaTypes.EntityData.SegmentPath = "frr-remote-lfa-types"
    frrRemoteLfaTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaTypes.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaTypes.EntityData.Children["frr-remote-lfa-type"] = types.YChild{"FrrRemoteLfaType", nil}
    for i := range frrRemoteLfaTypes.FrrRemoteLfaType {
        frrRemoteLfaTypes.EntityData.Children[types.GetSegmentPath(&frrRemoteLfaTypes.FrrRemoteLfaType[i])] = types.YChild{"FrrRemoteLfaType", &frrRemoteLfaTypes.FrrRemoteLfaType[i]}
    }
    frrRemoteLfaTypes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrRemoteLfaTypes.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType
// Enable remote lfa for a particular
// level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Remote LFA Type. The type is IsisRemoteLfa. This attribute is mandatory.
    Type_ interface{}
}

func (frrRemoteLfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrRemoteLfaTypes_FrrRemoteLfaType) GetEntityData() *types.CommonEntityData {
    frrRemoteLfaType.EntityData.YFilter = frrRemoteLfaType.YFilter
    frrRemoteLfaType.EntityData.YangName = "frr-remote-lfa-type"
    frrRemoteLfaType.EntityData.BundleName = "cisco_ios_xr"
    frrRemoteLfaType.EntityData.ParentYangName = "frr-remote-lfa-types"
    frrRemoteLfaType.EntityData.SegmentPath = "frr-remote-lfa-type" + "[level='" + fmt.Sprintf("%v", frrRemoteLfaType.Level) + "']"
    frrRemoteLfaType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrRemoteLfaType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrRemoteLfaType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrRemoteLfaType.EntityData.Children = make(map[string]types.YChild)
    frrRemoteLfaType.EntityData.Leafs = make(map[string]types.YLeaf)
    frrRemoteLfaType.EntityData.Leafs["level"] = types.YLeaf{"Level", frrRemoteLfaType.Level}
    frrRemoteLfaType.EntityData.Leafs["type"] = types.YLeaf{"Type_", frrRemoteLfaType.Type_}
    return &(frrRemoteLfaType.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults
// Interface FRR Default tiebreaker
// configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure default tiebreaker. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault.
    InterfaceFrrTiebreakerDefault []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault
}

func (interfaceFrrTiebreakerDefaults *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults) GetEntityData() *types.CommonEntityData {
    interfaceFrrTiebreakerDefaults.EntityData.YFilter = interfaceFrrTiebreakerDefaults.YFilter
    interfaceFrrTiebreakerDefaults.EntityData.YangName = "interface-frr-tiebreaker-defaults"
    interfaceFrrTiebreakerDefaults.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTiebreakerDefaults.EntityData.ParentYangName = "interface-frr-table"
    interfaceFrrTiebreakerDefaults.EntityData.SegmentPath = "interface-frr-tiebreaker-defaults"
    interfaceFrrTiebreakerDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTiebreakerDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTiebreakerDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTiebreakerDefaults.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTiebreakerDefaults.EntityData.Children["interface-frr-tiebreaker-default"] = types.YChild{"InterfaceFrrTiebreakerDefault", nil}
    for i := range interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault {
        interfaceFrrTiebreakerDefaults.EntityData.Children[types.GetSegmentPath(&interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[i])] = types.YChild{"InterfaceFrrTiebreakerDefault", &interfaceFrrTiebreakerDefaults.InterfaceFrrTiebreakerDefault[i]}
    }
    interfaceFrrTiebreakerDefaults.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceFrrTiebreakerDefaults.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault
// Configure default tiebreaker
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (interfaceFrrTiebreakerDefault *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakerDefaults_InterfaceFrrTiebreakerDefault) GetEntityData() *types.CommonEntityData {
    interfaceFrrTiebreakerDefault.EntityData.YFilter = interfaceFrrTiebreakerDefault.YFilter
    interfaceFrrTiebreakerDefault.EntityData.YangName = "interface-frr-tiebreaker-default"
    interfaceFrrTiebreakerDefault.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTiebreakerDefault.EntityData.ParentYangName = "interface-frr-tiebreaker-defaults"
    interfaceFrrTiebreakerDefault.EntityData.SegmentPath = "interface-frr-tiebreaker-default" + "[level='" + fmt.Sprintf("%v", interfaceFrrTiebreakerDefault.Level) + "']"
    interfaceFrrTiebreakerDefault.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTiebreakerDefault.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTiebreakerDefault.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTiebreakerDefault.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTiebreakerDefault.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceFrrTiebreakerDefault.EntityData.Leafs["level"] = types.YLeaf{"Level", interfaceFrrTiebreakerDefault.Level}
    return &(interfaceFrrTiebreakerDefault.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes
// TI LFA Enable
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable TI lfa for a particular level. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType.
    FrrtilfaType []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType
}

func (frrtilfaTypes *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes) GetEntityData() *types.CommonEntityData {
    frrtilfaTypes.EntityData.YFilter = frrtilfaTypes.YFilter
    frrtilfaTypes.EntityData.YangName = "frrtilfa-types"
    frrtilfaTypes.EntityData.BundleName = "cisco_ios_xr"
    frrtilfaTypes.EntityData.ParentYangName = "interface-frr-table"
    frrtilfaTypes.EntityData.SegmentPath = "frrtilfa-types"
    frrtilfaTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrtilfaTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrtilfaTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrtilfaTypes.EntityData.Children = make(map[string]types.YChild)
    frrtilfaTypes.EntityData.Children["frrtilfa-type"] = types.YChild{"FrrtilfaType", nil}
    for i := range frrtilfaTypes.FrrtilfaType {
        frrtilfaTypes.EntityData.Children[types.GetSegmentPath(&frrtilfaTypes.FrrtilfaType[i])] = types.YChild{"FrrtilfaType", &frrtilfaTypes.FrrtilfaType[i]}
    }
    frrtilfaTypes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrtilfaTypes.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType
// Enable TI lfa for a particular level
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}
}

func (frrtilfaType *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrtilfaTypes_FrrtilfaType) GetEntityData() *types.CommonEntityData {
    frrtilfaType.EntityData.YFilter = frrtilfaType.YFilter
    frrtilfaType.EntityData.YangName = "frrtilfa-type"
    frrtilfaType.EntityData.BundleName = "cisco_ios_xr"
    frrtilfaType.EntityData.ParentYangName = "frrtilfa-types"
    frrtilfaType.EntityData.SegmentPath = "frrtilfa-type" + "[level='" + fmt.Sprintf("%v", frrtilfaType.Level) + "']"
    frrtilfaType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrtilfaType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrtilfaType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrtilfaType.EntityData.Children = make(map[string]types.YChild)
    frrtilfaType.EntityData.Leafs = make(map[string]types.YLeaf)
    frrtilfaType.EntityData.Leafs["level"] = types.YLeaf{"Level", frrtilfaType.Level}
    return &(frrtilfaType.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces
// FRR exclusion configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from computation. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface.
    FrrExcludeInterface []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface
}

func (frrExcludeInterfaces *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    frrExcludeInterfaces.EntityData.YFilter = frrExcludeInterfaces.YFilter
    frrExcludeInterfaces.EntityData.YangName = "frr-exclude-interfaces"
    frrExcludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    frrExcludeInterfaces.EntityData.ParentYangName = "interface-frr-table"
    frrExcludeInterfaces.EntityData.SegmentPath = "frr-exclude-interfaces"
    frrExcludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrExcludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrExcludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrExcludeInterfaces.EntityData.Children = make(map[string]types.YChild)
    frrExcludeInterfaces.EntityData.Children["frr-exclude-interface"] = types.YChild{"FrrExcludeInterface", nil}
    for i := range frrExcludeInterfaces.FrrExcludeInterface {
        frrExcludeInterfaces.EntityData.Children[types.GetSegmentPath(&frrExcludeInterfaces.FrrExcludeInterface[i])] = types.YChild{"FrrExcludeInterface", &frrExcludeInterfaces.FrrExcludeInterface[i]}
    }
    frrExcludeInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frrExcludeInterfaces.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface
// Exclude an interface from computation
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // This attribute is a key. Computation Type. The type is Isisfrr.
    FrrType interface{}

    // Level. The type is interface{} with range: 0..2. This attribute is
    // mandatory.
    Level interface{}
}

func (frrExcludeInterface *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_FrrExcludeInterfaces_FrrExcludeInterface) GetEntityData() *types.CommonEntityData {
    frrExcludeInterface.EntityData.YFilter = frrExcludeInterface.YFilter
    frrExcludeInterface.EntityData.YangName = "frr-exclude-interface"
    frrExcludeInterface.EntityData.BundleName = "cisco_ios_xr"
    frrExcludeInterface.EntityData.ParentYangName = "frr-exclude-interfaces"
    frrExcludeInterface.EntityData.SegmentPath = "frr-exclude-interface" + "[interface-name='" + fmt.Sprintf("%v", frrExcludeInterface.InterfaceName) + "']" + "[frr-type='" + fmt.Sprintf("%v", frrExcludeInterface.FrrType) + "']"
    frrExcludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frrExcludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frrExcludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frrExcludeInterface.EntityData.Children = make(map[string]types.YChild)
    frrExcludeInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    frrExcludeInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", frrExcludeInterface.InterfaceName}
    frrExcludeInterface.EntityData.Leafs["frr-type"] = types.YLeaf{"FrrType", frrExcludeInterface.FrrType}
    frrExcludeInterface.EntityData.Leafs["level"] = types.YLeaf{"Level", frrExcludeInterface.Level}
    return &(frrExcludeInterface.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers
// Interface FRR tiebreakers configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure tiebreaker for multiple backups. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker.
    InterfaceFrrTiebreaker []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker
}

func (interfaceFrrTiebreakers *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers) GetEntityData() *types.CommonEntityData {
    interfaceFrrTiebreakers.EntityData.YFilter = interfaceFrrTiebreakers.YFilter
    interfaceFrrTiebreakers.EntityData.YangName = "interface-frr-tiebreakers"
    interfaceFrrTiebreakers.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTiebreakers.EntityData.ParentYangName = "interface-frr-table"
    interfaceFrrTiebreakers.EntityData.SegmentPath = "interface-frr-tiebreakers"
    interfaceFrrTiebreakers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTiebreakers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTiebreakers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTiebreakers.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTiebreakers.EntityData.Children["interface-frr-tiebreaker"] = types.YChild{"InterfaceFrrTiebreaker", nil}
    for i := range interfaceFrrTiebreakers.InterfaceFrrTiebreaker {
        interfaceFrrTiebreakers.EntityData.Children[types.GetSegmentPath(&interfaceFrrTiebreakers.InterfaceFrrTiebreaker[i])] = types.YChild{"InterfaceFrrTiebreaker", &interfaceFrrTiebreakers.InterfaceFrrTiebreaker[i]}
    }
    interfaceFrrTiebreakers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceFrrTiebreakers.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker
// Configure tiebreaker for multiple
// backups
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker struct {
    EntityData types.CommonEntityData
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

func (interfaceFrrTiebreaker *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceFrrTable_InterfaceFrrTiebreakers_InterfaceFrrTiebreaker) GetEntityData() *types.CommonEntityData {
    interfaceFrrTiebreaker.EntityData.YFilter = interfaceFrrTiebreaker.YFilter
    interfaceFrrTiebreaker.EntityData.YangName = "interface-frr-tiebreaker"
    interfaceFrrTiebreaker.EntityData.BundleName = "cisco_ios_xr"
    interfaceFrrTiebreaker.EntityData.ParentYangName = "interface-frr-tiebreakers"
    interfaceFrrTiebreaker.EntityData.SegmentPath = "interface-frr-tiebreaker" + "[level='" + fmt.Sprintf("%v", interfaceFrrTiebreaker.Level) + "']" + "[tiebreaker='" + fmt.Sprintf("%v", interfaceFrrTiebreaker.Tiebreaker) + "']"
    interfaceFrrTiebreaker.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceFrrTiebreaker.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceFrrTiebreaker.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceFrrTiebreaker.EntityData.Children = make(map[string]types.YChild)
    interfaceFrrTiebreaker.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceFrrTiebreaker.EntityData.Leafs["level"] = types.YLeaf{"Level", interfaceFrrTiebreaker.Level}
    interfaceFrrTiebreaker.EntityData.Leafs["tiebreaker"] = types.YLeaf{"Tiebreaker", interfaceFrrTiebreaker.Tiebreaker}
    interfaceFrrTiebreaker.EntityData.Leafs["index"] = types.YLeaf{"Index", interfaceFrrTiebreaker.Index}
    return &(interfaceFrrTiebreaker.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp
// MPLS LDP configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable MPLS LDP Synchronization for an IS-IS level. The type is interface{}
    // with range: 0..2. The default value is 0.
    SyncLevel interface{}
}

func (mplsLdp *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_MplsLdp) GetEntityData() *types.CommonEntityData {
    mplsLdp.EntityData.YFilter = mplsLdp.YFilter
    mplsLdp.EntityData.YangName = "mpls-ldp"
    mplsLdp.EntityData.BundleName = "cisco_ios_xr"
    mplsLdp.EntityData.ParentYangName = "topology-name"
    mplsLdp.EntityData.SegmentPath = "mpls-ldp"
    mplsLdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLdp.EntityData.Children = make(map[string]types.YChild)
    mplsLdp.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsLdp.EntityData.Leafs["sync-level"] = types.YLeaf{"SyncLevel", mplsLdp.SyncLevel}
    return &(mplsLdp.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid
// Assign prefix SSPF SID to an interface,
// ISISPHPFlag will be rejected if set to
// disable, ISISEXPLICITNULLFlag will
// override the value of ISISPHPFlag
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type for the interface. The type is Isissid1. This attribute is
    // mandatory.
    Type_ interface{}

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

func (prefixSspfsid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_PrefixSspfsid) GetEntityData() *types.CommonEntityData {
    prefixSspfsid.EntityData.YFilter = prefixSspfsid.YFilter
    prefixSspfsid.EntityData.YangName = "prefix-sspfsid"
    prefixSspfsid.EntityData.BundleName = "cisco_ios_xr"
    prefixSspfsid.EntityData.ParentYangName = "topology-name"
    prefixSspfsid.EntityData.SegmentPath = "prefix-sspfsid"
    prefixSspfsid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSspfsid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSspfsid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSspfsid.EntityData.Children = make(map[string]types.YChild)
    prefixSspfsid.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixSspfsid.EntityData.Leafs["type"] = types.YLeaf{"Type_", prefixSspfsid.Type_}
    prefixSspfsid.EntityData.Leafs["value"] = types.YLeaf{"Value", prefixSspfsid.Value}
    prefixSspfsid.EntityData.Leafs["php"] = types.YLeaf{"Php", prefixSspfsid.Php}
    prefixSspfsid.EntityData.Leafs["explicit-null"] = types.YLeaf{"ExplicitNull", prefixSspfsid.ExplicitNull}
    prefixSspfsid.EntityData.Leafs["nflag-clear"] = types.YLeaf{"NflagClear", prefixSspfsid.NflagClear}
    return &(prefixSspfsid.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics
// AutoMetric configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoMetric Proactive-Protect configuration. Legal value depends on the
    // metric-style specified for the topology. If the metric-style defined is
    // narrow, then only a value between <1-63> is allowed and if the metric-style
    // is defined as wide, then a value between <1-16777214> is allowed as the
    // auto-metric value. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric.
    AutoMetric []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric
}

func (autoMetrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics) GetEntityData() *types.CommonEntityData {
    autoMetrics.EntityData.YFilter = autoMetrics.YFilter
    autoMetrics.EntityData.YangName = "auto-metrics"
    autoMetrics.EntityData.BundleName = "cisco_ios_xr"
    autoMetrics.EntityData.ParentYangName = "topology-name"
    autoMetrics.EntityData.SegmentPath = "auto-metrics"
    autoMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoMetrics.EntityData.Children = make(map[string]types.YChild)
    autoMetrics.EntityData.Children["auto-metric"] = types.YChild{"AutoMetric", nil}
    for i := range autoMetrics.AutoMetric {
        autoMetrics.EntityData.Children[types.GetSegmentPath(&autoMetrics.AutoMetric[i])] = types.YChild{"AutoMetric", &autoMetrics.AutoMetric[i]}
    }
    autoMetrics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(autoMetrics.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed auto metric:<1-63> for narrow ,<1-16777214> for wide. The type is
    // interface{} with range: 1..16777214. This attribute is mandatory.
    ProactiveProtect interface{}
}

func (autoMetric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AutoMetrics_AutoMetric) GetEntityData() *types.CommonEntityData {
    autoMetric.EntityData.YFilter = autoMetric.YFilter
    autoMetric.EntityData.YangName = "auto-metric"
    autoMetric.EntityData.BundleName = "cisco_ios_xr"
    autoMetric.EntityData.ParentYangName = "auto-metrics"
    autoMetric.EntityData.SegmentPath = "auto-metric" + "[level='" + fmt.Sprintf("%v", autoMetric.Level) + "']"
    autoMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoMetric.EntityData.Children = make(map[string]types.YChild)
    autoMetric.EntityData.Leafs = make(map[string]types.YLeaf)
    autoMetric.EntityData.Leafs["level"] = types.YLeaf{"Level", autoMetric.Level}
    autoMetric.EntityData.Leafs["proactive-protect"] = types.YLeaf{"ProactiveProtect", autoMetric.ProactiveProtect}
    return &(autoMetric.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags
// admin-tag configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Admin tag for advertised interface connected routes. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag.
    AdminTag []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag
}

func (adminTags *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags) GetEntityData() *types.CommonEntityData {
    adminTags.EntityData.YFilter = adminTags.YFilter
    adminTags.EntityData.YangName = "admin-tags"
    adminTags.EntityData.BundleName = "cisco_ios_xr"
    adminTags.EntityData.ParentYangName = "topology-name"
    adminTags.EntityData.SegmentPath = "admin-tags"
    adminTags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminTags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminTags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminTags.EntityData.Children = make(map[string]types.YChild)
    adminTags.EntityData.Children["admin-tag"] = types.YChild{"AdminTag", nil}
    for i := range adminTags.AdminTag {
        adminTags.EntityData.Children[types.GetSegmentPath(&adminTags.AdminTag[i])] = types.YChild{"AdminTag", &adminTags.AdminTag[i]}
    }
    adminTags.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adminTags.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag
// Admin tag for advertised interface
// connected routes
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Tag to associate with connected routes. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    AdminTag interface{}
}

func (adminTag *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_AdminTags_AdminTag) GetEntityData() *types.CommonEntityData {
    adminTag.EntityData.YFilter = adminTag.YFilter
    adminTag.EntityData.YangName = "admin-tag"
    adminTag.EntityData.BundleName = "cisco_ios_xr"
    adminTag.EntityData.ParentYangName = "admin-tags"
    adminTag.EntityData.SegmentPath = "admin-tag" + "[level='" + fmt.Sprintf("%v", adminTag.Level) + "']"
    adminTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adminTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adminTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adminTag.EntityData.Children = make(map[string]types.YChild)
    adminTag.EntityData.Leafs = make(map[string]types.YLeaf)
    adminTag.EntityData.Leafs["level"] = types.YLeaf{"Level", adminTag.Level}
    adminTag.EntityData.Leafs["admin-tag"] = types.YLeaf{"AdminTag", adminTag.AdminTag}
    return &(adminTag.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup
// Provide link group name and level
// This type is a presence type.
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link Group. The type is string with length: 1..40. This attribute is
    // mandatory.
    LinkGroup interface{}

    // Level in which link group will be effective. The type is interface{} with
    // range: 0..2. The default value is 0.
    Level interface{}
}

func (interfaceLinkGroup *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_InterfaceLinkGroup) GetEntityData() *types.CommonEntityData {
    interfaceLinkGroup.EntityData.YFilter = interfaceLinkGroup.YFilter
    interfaceLinkGroup.EntityData.YangName = "interface-link-group"
    interfaceLinkGroup.EntityData.BundleName = "cisco_ios_xr"
    interfaceLinkGroup.EntityData.ParentYangName = "topology-name"
    interfaceLinkGroup.EntityData.SegmentPath = "interface-link-group"
    interfaceLinkGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLinkGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLinkGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLinkGroup.EntityData.Children = make(map[string]types.YChild)
    interfaceLinkGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceLinkGroup.EntityData.Leafs["link-group"] = types.YLeaf{"LinkGroup", interfaceLinkGroup.LinkGroup}
    interfaceLinkGroup.EntityData.Leafs["level"] = types.YLeaf{"Level", interfaceLinkGroup.Level}
    return &(interfaceLinkGroup.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids
// Manual Adjacecy SID configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Assign adjancency SID to an interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid.
    ManualAdjSid []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid
}

func (manualAdjSids *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids) GetEntityData() *types.CommonEntityData {
    manualAdjSids.EntityData.YFilter = manualAdjSids.YFilter
    manualAdjSids.EntityData.YangName = "manual-adj-sids"
    manualAdjSids.EntityData.BundleName = "cisco_ios_xr"
    manualAdjSids.EntityData.ParentYangName = "topology-name"
    manualAdjSids.EntityData.SegmentPath = "manual-adj-sids"
    manualAdjSids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualAdjSids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualAdjSids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualAdjSids.EntityData.Children = make(map[string]types.YChild)
    manualAdjSids.EntityData.Children["manual-adj-sid"] = types.YChild{"ManualAdjSid", nil}
    for i := range manualAdjSids.ManualAdjSid {
        manualAdjSids.EntityData.Children[types.GetSegmentPath(&manualAdjSids.ManualAdjSid[i])] = types.YChild{"ManualAdjSid", &manualAdjSids.ManualAdjSid[i]}
    }
    manualAdjSids.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(manualAdjSids.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid
// Assign adjancency SID to an interface
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid struct {
    EntityData types.CommonEntityData
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

func (manualAdjSid *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_ManualAdjSids_ManualAdjSid) GetEntityData() *types.CommonEntityData {
    manualAdjSid.EntityData.YFilter = manualAdjSid.YFilter
    manualAdjSid.EntityData.YangName = "manual-adj-sid"
    manualAdjSid.EntityData.BundleName = "cisco_ios_xr"
    manualAdjSid.EntityData.ParentYangName = "manual-adj-sids"
    manualAdjSid.EntityData.SegmentPath = "manual-adj-sid" + "[level='" + fmt.Sprintf("%v", manualAdjSid.Level) + "']" + "[sid-type='" + fmt.Sprintf("%v", manualAdjSid.SidType) + "']" + "[sid='" + fmt.Sprintf("%v", manualAdjSid.Sid) + "']"
    manualAdjSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manualAdjSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manualAdjSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manualAdjSid.EntityData.Children = make(map[string]types.YChild)
    manualAdjSid.EntityData.Leafs = make(map[string]types.YLeaf)
    manualAdjSid.EntityData.Leafs["level"] = types.YLeaf{"Level", manualAdjSid.Level}
    manualAdjSid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", manualAdjSid.SidType}
    manualAdjSid.EntityData.Leafs["sid"] = types.YLeaf{"Sid", manualAdjSid.Sid}
    manualAdjSid.EntityData.Leafs["protected"] = types.YLeaf{"Protected", manualAdjSid.Protected}
    return &(manualAdjSid.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics
// Metric configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics struct {
    EntityData types.CommonEntityData
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

func (metrics *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics) GetEntityData() *types.CommonEntityData {
    metrics.EntityData.YFilter = metrics.YFilter
    metrics.EntityData.YangName = "metrics"
    metrics.EntityData.BundleName = "cisco_ios_xr"
    metrics.EntityData.ParentYangName = "topology-name"
    metrics.EntityData.SegmentPath = "metrics"
    metrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metrics.EntityData.Children = make(map[string]types.YChild)
    metrics.EntityData.Children["metric"] = types.YChild{"Metric", nil}
    for i := range metrics.Metric {
        metrics.EntityData.Children[types.GetSegmentPath(&metrics.Metric[i])] = types.YChild{"Metric", &metrics.Metric[i]}
    }
    metrics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(metrics.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Allowed metric: <1-63> for narrow, <1-16777215> for wide. The type is one
    // of the following types: enumeration
    // Isis.Instances.Instance.Interfaces.Interface.InterfaceAfs.InterfaceAf.TopologyName.Metrics.Metric.Metric_
    // This attribute is mandatory., or int with range: 1..16777215 This attribute
    // is mandatory..
    Metric interface{}
}

func (metric *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "metrics"
    metric.EntityData.SegmentPath = "metric" + "[level='" + fmt.Sprintf("%v", metric.Level) + "']"
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = make(map[string]types.YChild)
    metric.EntityData.Leafs = make(map[string]types.YLeaf)
    metric.EntityData.Leafs["level"] = types.YLeaf{"Level", metric.Level}
    metric.EntityData.Leafs["metric"] = types.YLeaf{"Metric", metric.Metric}
    return &(metric.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric_Metric_ represents <1-16777215> for wide
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric_Metric_ string

const (
    // Maximum wide metric.  All routers will
    // exclude this link from their SPF
    Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric_Metric__maximum Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Metrics_Metric_Metric_ = "maximum"
)

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights
// Weight configuration
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Weight configuration under interface for load balancing. The type is slice
    // of
    // Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight.
    Weight []Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight
}

func (weights *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights) GetEntityData() *types.CommonEntityData {
    weights.EntityData.YFilter = weights.YFilter
    weights.EntityData.YangName = "weights"
    weights.EntityData.BundleName = "cisco_ios_xr"
    weights.EntityData.ParentYangName = "topology-name"
    weights.EntityData.SegmentPath = "weights"
    weights.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    weights.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    weights.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    weights.EntityData.Children = make(map[string]types.YChild)
    weights.EntityData.Children["weight"] = types.YChild{"Weight", nil}
    for i := range weights.Weight {
        weights.EntityData.Children[types.GetSegmentPath(&weights.Weight[i])] = types.YChild{"Weight", &weights.Weight[i]}
    }
    weights.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(weights.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight
// Weight configuration under interface for load
// balancing
type Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Weight to be configured under interface for Load Balancing. Allowed weight:
    // <1-16777215>. The type is interface{} with range: 1..16777214. This
    // attribute is mandatory.
    Weight interface{}
}

func (weight *Isis_Instances_Instance_Interfaces_Interface_InterfaceAfs_InterfaceAf_TopologyName_Weights_Weight) GetEntityData() *types.CommonEntityData {
    weight.EntityData.YFilter = weight.YFilter
    weight.EntityData.YangName = "weight"
    weight.EntityData.BundleName = "cisco_ios_xr"
    weight.EntityData.ParentYangName = "weights"
    weight.EntityData.SegmentPath = "weight" + "[level='" + fmt.Sprintf("%v", weight.Level) + "']"
    weight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    weight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    weight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    weight.EntityData.Children = make(map[string]types.YChild)
    weight.EntityData.Leafs = make(map[string]types.YLeaf)
    weight.EntityData.Leafs["level"] = types.YLeaf{"Level", weight.Level}
    weight.EntityData.Leafs["weight"] = types.YLeaf{"Weight", weight.Weight}
    return &(weight.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals
// CSNP-interval configuration
type Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CSNP-interval configuration. No fixed default value as this depends on the
    // media type of the interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval.
    CsnpInterval []Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval
}

func (csnpIntervals *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals) GetEntityData() *types.CommonEntityData {
    csnpIntervals.EntityData.YFilter = csnpIntervals.YFilter
    csnpIntervals.EntityData.YangName = "csnp-intervals"
    csnpIntervals.EntityData.BundleName = "cisco_ios_xr"
    csnpIntervals.EntityData.ParentYangName = "interface"
    csnpIntervals.EntityData.SegmentPath = "csnp-intervals"
    csnpIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csnpIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csnpIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csnpIntervals.EntityData.Children = make(map[string]types.YChild)
    csnpIntervals.EntityData.Children["csnp-interval"] = types.YChild{"CsnpInterval", nil}
    for i := range csnpIntervals.CsnpInterval {
        csnpIntervals.EntityData.Children[types.GetSegmentPath(&csnpIntervals.CsnpInterval[i])] = types.YChild{"CsnpInterval", &csnpIntervals.CsnpInterval[i]}
    }
    csnpIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(csnpIntervals.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval
// CSNP-interval configuration. No fixed
// default value as this depends on the media
// type of the interface.
type Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Seconds. The type is interface{} with range: 0..65535. This attribute is
    // mandatory. Units are second.
    Interval interface{}
}

func (csnpInterval *Isis_Instances_Instance_Interfaces_Interface_CsnpIntervals_CsnpInterval) GetEntityData() *types.CommonEntityData {
    csnpInterval.EntityData.YFilter = csnpInterval.YFilter
    csnpInterval.EntityData.YangName = "csnp-interval"
    csnpInterval.EntityData.BundleName = "cisco_ios_xr"
    csnpInterval.EntityData.ParentYangName = "csnp-intervals"
    csnpInterval.EntityData.SegmentPath = "csnp-interval" + "[level='" + fmt.Sprintf("%v", csnpInterval.Level) + "']"
    csnpInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csnpInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csnpInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csnpInterval.EntityData.Children = make(map[string]types.YChild)
    csnpInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    csnpInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", csnpInterval.Level}
    csnpInterval.EntityData.Leafs["interval"] = types.YLeaf{"Interval", csnpInterval.Interval}
    return &(csnpInterval.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_LspIntervals
// LSP-interval configuration
type Isis_Instances_Instance_Interfaces_Interface_LspIntervals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval between transmission of LSPs on interface. The type is slice of
    // Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval.
    LspInterval []Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval
}

func (lspIntervals *Isis_Instances_Instance_Interfaces_Interface_LspIntervals) GetEntityData() *types.CommonEntityData {
    lspIntervals.EntityData.YFilter = lspIntervals.YFilter
    lspIntervals.EntityData.YangName = "lsp-intervals"
    lspIntervals.EntityData.BundleName = "cisco_ios_xr"
    lspIntervals.EntityData.ParentYangName = "interface"
    lspIntervals.EntityData.SegmentPath = "lsp-intervals"
    lspIntervals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspIntervals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspIntervals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspIntervals.EntityData.Children = make(map[string]types.YChild)
    lspIntervals.EntityData.Children["lsp-interval"] = types.YChild{"LspInterval", nil}
    for i := range lspIntervals.LspInterval {
        lspIntervals.EntityData.Children[types.GetSegmentPath(&lspIntervals.LspInterval[i])] = types.YChild{"LspInterval", &lspIntervals.LspInterval[i]}
    }
    lspIntervals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspIntervals.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval
// Interval between transmission of LSPs on
// interface.
type Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Level to which configuration applies. The type is
    // IsisInternalLevel.
    Level interface{}

    // Milliseconds. The type is interface{} with range: 1..4294967295. This
    // attribute is mandatory. Units are millisecond.
    Interval interface{}
}

func (lspInterval *Isis_Instances_Instance_Interfaces_Interface_LspIntervals_LspInterval) GetEntityData() *types.CommonEntityData {
    lspInterval.EntityData.YFilter = lspInterval.YFilter
    lspInterval.EntityData.YangName = "lsp-interval"
    lspInterval.EntityData.BundleName = "cisco_ios_xr"
    lspInterval.EntityData.ParentYangName = "lsp-intervals"
    lspInterval.EntityData.SegmentPath = "lsp-interval" + "[level='" + fmt.Sprintf("%v", lspInterval.Level) + "']"
    lspInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspInterval.EntityData.Children = make(map[string]types.YChild)
    lspInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    lspInterval.EntityData.Leafs["level"] = types.YLeaf{"Level", lspInterval.Level}
    lspInterval.EntityData.Leafs["interval"] = types.YLeaf{"Interval", lspInterval.Interval}
    return &(lspInterval.EntityData)
}

// Isis_Instances_Instance_Interfaces_Interface_MeshGroup represents Mesh-group configuration
type Isis_Instances_Instance_Interfaces_Interface_MeshGroup string

const (
    // Blocked mesh group.  Changed LSPs are not
    // flooded over blocked interfaces
    Isis_Instances_Instance_Interfaces_Interface_MeshGroup_blocked Isis_Instances_Instance_Interfaces_Interface_MeshGroup = "blocked"
)

