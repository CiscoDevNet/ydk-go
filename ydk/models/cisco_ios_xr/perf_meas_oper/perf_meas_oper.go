// This module contains a collection of YANG definitions
// for Cisco IOS-XR perf-meas package operational data.
// 
// This module contains definitions
// for the following management objects:
//   performance-measurement: Performance Measurement operational
//     data
//   performance-measurement-responder: performance measurement
//     responder
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package perf_meas_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package perf_meas_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-perf-meas-oper performance-measurement}", reflect.TypeOf(PerformanceMeasurement{}))
    ydk.RegisterEntity("Cisco-IOS-XR-perf-meas-oper:performance-measurement", reflect.TypeOf(PerformanceMeasurement{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-perf-meas-oper performance-measurement-responder}", reflect.TypeOf(PerformanceMeasurementResponder{}))
    ydk.RegisterEntity("Cisco-IOS-XR-perf-meas-oper:performance-measurement-responder", reflect.TypeOf(PerformanceMeasurementResponder{}))
}

// PmAdvertReason represents PM advertisement reason
type PmAdvertReason string

const (
    // Periodic timer expired. No advertisements have
    // occured
    PmAdvertReason_periodic_timer_expired_no_advertisements PmAdvertReason = "periodic-timer-expired-no-advertisements"

    // Periodic timer expired. Average value threshold
    // crossed
    PmAdvertReason_periodic_advertisement_threshold_average PmAdvertReason = "periodic-advertisement-threshold-average"

    // Periodic timer expired. Minimum value threshold
    // crossed
    PmAdvertReason_periodic_advertisement_threshold_minimum PmAdvertReason = "periodic-advertisement-threshold-minimum"

    // Periodic timer expired. Maximum value threshold
    // crossed
    PmAdvertReason_periodic_advertisement_threshold_maximum PmAdvertReason = "periodic-advertisement-threshold-maximum"

    // Periodic timer expired. Variance value
    // threshold crossed
    PmAdvertReason_periodic_advertisement_threshold_variance PmAdvertReason = "periodic-advertisement-threshold-variance"

    // Accelerated minimum value threshold crossed
    PmAdvertReason_accelerated_advertisement_threshold_minimum PmAdvertReason = "accelerated-advertisement-threshold-minimum"

    // Accelerated minimum value upper bound crossed
    PmAdvertReason_accelerated_advertisement_upper_bound_minimum PmAdvertReason = "accelerated-advertisement-upper-bound-minimum"

    // Advertisement enabled
    PmAdvertReason_advertisement_enabled PmAdvertReason = "advertisement-enabled"

    // Advertisement disabled
    PmAdvertReason_advertisement_disabled PmAdvertReason = "advertisement-disabled"

    // Session unconfigured
    PmAdvertReason_session_unconfigured PmAdvertReason = "session-unconfigured"

    // Session cleared via CLI
    PmAdvertReason_clear_cli_command PmAdvertReason = "clear-cli-command"

    // Advertise delay config
    PmAdvertReason_advertise_delay_config PmAdvertReason = "advertise-delay-config"

    // Advertise delay unconfig
    PmAdvertReason_advertise_delay_unconfig PmAdvertReason = "advertise-delay-unconfig"

    // Recevied control code error, as per RFC 6374,
    // from the responder
    PmAdvertReason_received_control_code_error PmAdvertReason = "received-control-code-error"

    // Link is in down state
    PmAdvertReason_link_is_down PmAdvertReason = "link-is-down"
)

// PmProbeNotRunningReason represents PM probe not running reason
type PmProbeNotRunningReason string

const (
    // Probe is running
    PmProbeNotRunningReason_probe_is_running PmProbeNotRunningReason = "probe-is-running"

    // Measurement is not supported on this plaftorm
    PmProbeNotRunningReason_platform_not_supported PmProbeNotRunningReason = "platform-not-supported"

    // Node is not V1 active
    PmProbeNotRunningReason_nonv1_active_node PmProbeNotRunningReason = "nonv1-active-node"

    // An uncleared control code error was received
    PmProbeNotRunningReason_control_code_error PmProbeNotRunningReason = "control-code-error"

    // Interface admin down
    PmProbeNotRunningReason_interface_admin_down PmProbeNotRunningReason = "interface-admin-down"

    // Interface down
    PmProbeNotRunningReason_interface_down PmProbeNotRunningReason = "interface-down"

    // MPLS is not enabled on interface
    PmProbeNotRunningReason_mpls_capability_not_present PmProbeNotRunningReason = "mpls-capability-not-present"

    // Interface not present or preconfigured
    PmProbeNotRunningReason_interface_not_present PmProbeNotRunningReason = "interface-not-present"

    // IP address is not present on interface
    PmProbeNotRunningReason_ip_address_not_configured PmProbeNotRunningReason = "ip-address-not-configured"
)

// PmMeasurement represents PM Measurement Type
type PmMeasurement string

const (
    // Delay Measurement Type
    PmMeasurement_delay_measurement_type PmMeasurement = "delay-measurement-type"
)

// PmTransport represents PM Transport Type
type PmTransport string

const (
    // Interface transport type
    PmTransport_interface_transport_type PmTransport = "interface-transport-type"

    // RSVP-TE tunnel transport type
    PmTransport_rsvpte_transport_type PmTransport = "rsvpte-transport-type"

    // SR Policy transport type
    PmTransport_sr_policy_transport_type PmTransport = "sr-policy-transport-type"
)

// PmDelayMode represents PM Delay Mode Type
type PmDelayMode string

const (
    // One-way delay-measurement mode
    PmDelayMode_delay_mode_one_way PmDelayMode = "delay-mode-one-way"

    // Two-way delay-measurement mode
    PmDelayMode_delay_mode_two_way PmDelayMode = "delay-mode-two-way"
)

// PerformanceMeasurement
// Performance Measurement operational data
type PerformanceMeasurement struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node table for node-specific operational data.
    Nodes PerformanceMeasurement_Nodes
}

func (performanceMeasurement *PerformanceMeasurement) GetEntityData() *types.CommonEntityData {
    performanceMeasurement.EntityData.YFilter = performanceMeasurement.YFilter
    performanceMeasurement.EntityData.YangName = "performance-measurement"
    performanceMeasurement.EntityData.BundleName = "cisco_ios_xr"
    performanceMeasurement.EntityData.ParentYangName = "Cisco-IOS-XR-perf-meas-oper"
    performanceMeasurement.EntityData.SegmentPath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement"
    performanceMeasurement.EntityData.AbsolutePath = performanceMeasurement.EntityData.SegmentPath
    performanceMeasurement.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    performanceMeasurement.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    performanceMeasurement.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    performanceMeasurement.EntityData.Children = types.NewOrderedMap()
    performanceMeasurement.EntityData.Children.Append("nodes", types.YChild{"Nodes", &performanceMeasurement.Nodes})
    performanceMeasurement.EntityData.Leafs = types.NewOrderedMap()

    performanceMeasurement.EntityData.YListKeys = []string {}

    return &(performanceMeasurement.EntityData)
}

// PerformanceMeasurement_Nodes
// Node table for node-specific operational data
type PerformanceMeasurement_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // PerformanceMeasurement_Nodes_Node.
    Node []*PerformanceMeasurement_Nodes_Node
}

func (nodes *PerformanceMeasurement_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "performance-measurement"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerformanceMeasurement_Nodes_Node
// Node-specific data for a particular node
type PerformanceMeasurement_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Summary.
    Summary PerformanceMeasurement_Nodes_Node_Summary

    // Interfaces.
    Interfaces PerformanceMeasurement_Nodes_Node_Interfaces
}

func (node *PerformanceMeasurement_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Node, "node")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node", types.YLeaf{"Node", node.Node})

    node.EntityData.YListKeys = []string {"Node"}

    return &(node.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Summary
// Summary
type PerformanceMeasurement_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of delay measurement interfaces enabled. The type is interface{}
    // with range: 0..4294967295.
    TotalInterfaces interface{}

    // Delay summary.
    DelaySummary PerformanceMeasurement_Nodes_Node_Summary_DelaySummary
}

func (summary *PerformanceMeasurement_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("delay-summary", types.YChild{"DelaySummary", &summary.DelaySummary})
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("total-interfaces", types.YLeaf{"TotalInterfaces", summary.TotalInterfaces})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Summary_DelaySummary
// Delay summary
type PerformanceMeasurement_Nodes_Node_Summary_DelaySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface delay summary.
    InterfaceDelaySummary PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary

    // PM delay global counters.
    DelayGlobalCounters PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_DelayGlobalCounters
}

func (delaySummary *PerformanceMeasurement_Nodes_Node_Summary_DelaySummary) GetEntityData() *types.CommonEntityData {
    delaySummary.EntityData.YFilter = delaySummary.YFilter
    delaySummary.EntityData.YangName = "delay-summary"
    delaySummary.EntityData.BundleName = "cisco_ios_xr"
    delaySummary.EntityData.ParentYangName = "summary"
    delaySummary.EntityData.SegmentPath = "delay-summary"
    delaySummary.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/summary/" + delaySummary.EntityData.SegmentPath
    delaySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delaySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delaySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delaySummary.EntityData.Children = types.NewOrderedMap()
    delaySummary.EntityData.Children.Append("interface-delay-summary", types.YChild{"InterfaceDelaySummary", &delaySummary.InterfaceDelaySummary})
    delaySummary.EntityData.Children.Append("delay-global-counters", types.YChild{"DelayGlobalCounters", &delaySummary.DelayGlobalCounters})
    delaySummary.EntityData.Leafs = types.NewOrderedMap()

    delaySummary.EntityData.YListKeys = []string {}

    return &(delaySummary.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary
// Interface delay summary
type PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of delay measurement sessions enabled. The type is interface{} with
    // range: 0..4294967295.
    TotalDelaySessions interface{}

    // Delay profile.
    DelayProfile PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayProfile

    // PM delay counters for a transport types.
    DelayTransportCounters PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters
}

func (interfaceDelaySummary *PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary) GetEntityData() *types.CommonEntityData {
    interfaceDelaySummary.EntityData.YFilter = interfaceDelaySummary.YFilter
    interfaceDelaySummary.EntityData.YangName = "interface-delay-summary"
    interfaceDelaySummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceDelaySummary.EntityData.ParentYangName = "delay-summary"
    interfaceDelaySummary.EntityData.SegmentPath = "interface-delay-summary"
    interfaceDelaySummary.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/summary/delay-summary/" + interfaceDelaySummary.EntityData.SegmentPath
    interfaceDelaySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDelaySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDelaySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDelaySummary.EntityData.Children = types.NewOrderedMap()
    interfaceDelaySummary.EntityData.Children.Append("delay-profile", types.YChild{"DelayProfile", &interfaceDelaySummary.DelayProfile})
    interfaceDelaySummary.EntityData.Children.Append("delay-transport-counters", types.YChild{"DelayTransportCounters", &interfaceDelaySummary.DelayTransportCounters})
    interfaceDelaySummary.EntityData.Leafs = types.NewOrderedMap()
    interfaceDelaySummary.EntityData.Leafs.Append("total-delay-sessions", types.YLeaf{"TotalDelaySessions", interfaceDelaySummary.TotalDelaySessions})

    interfaceDelaySummary.EntityData.YListKeys = []string {}

    return &(interfaceDelaySummary.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayProfile
// Delay profile
type PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Probe Interval (sec). The type is interface{} with range: 0..4294967295.
    ProbeInterval interface{}

    // Burst Interval (msec). The type is interface{} with range: 0..4294967295.
    BurstInterval interface{}

    // Burst Count. The type is interface{} with range: 0..4294967295.
    BurstCount interface{}

    // Delay Measurement Mode Type. The type is PmDelayMode.
    DelayMeasurementMode interface{}

    // Advertisement Periodic Enabled. The type is bool.
    PeriodicAdvertisementEnabled interface{}

    // Advertisement Periodic Interval (sec). The type is interface{} with range:
    // 0..4294967295.
    PeriodicAdvertisementInterval interface{}

    // Advertisement Periodic Effective (sec). The type is interface{} with range:
    // 0..4294967295.
    EffectivePeriodicAdvertisementInterval interface{}

    // Advertisement Periodic Threshold (%). The type is interface{} with range:
    // 0..4294967295.
    PeriodicAdvertisementThreshold interface{}

    // Advertisement Periodic Minimum Change (uSec). The type is interface{} with
    // range: 0..4294967295.
    PeriodicAdvertisementMinimumChange interface{}

    // Advertisement Accelerated Threshold (%). The type is interface{} with
    // range: 0..4294967295.
    AcceleratedAdvertisementThreshold interface{}

    // Advertisement Accelerated Minimum Change (uSec). The type is interface{}
    // with range: 0..4294967295.
    AcceleratedAdvertisementMinimumChange interface{}

    // Advertisement Accelerated Enabled. The type is bool.
    AcceleratedAdvertisementEnabled interface{}
}

func (delayProfile *PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayProfile) GetEntityData() *types.CommonEntityData {
    delayProfile.EntityData.YFilter = delayProfile.YFilter
    delayProfile.EntityData.YangName = "delay-profile"
    delayProfile.EntityData.BundleName = "cisco_ios_xr"
    delayProfile.EntityData.ParentYangName = "interface-delay-summary"
    delayProfile.EntityData.SegmentPath = "delay-profile"
    delayProfile.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/summary/delay-summary/interface-delay-summary/" + delayProfile.EntityData.SegmentPath
    delayProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayProfile.EntityData.Children = types.NewOrderedMap()
    delayProfile.EntityData.Leafs = types.NewOrderedMap()
    delayProfile.EntityData.Leafs.Append("probe-interval", types.YLeaf{"ProbeInterval", delayProfile.ProbeInterval})
    delayProfile.EntityData.Leafs.Append("burst-interval", types.YLeaf{"BurstInterval", delayProfile.BurstInterval})
    delayProfile.EntityData.Leafs.Append("burst-count", types.YLeaf{"BurstCount", delayProfile.BurstCount})
    delayProfile.EntityData.Leafs.Append("delay-measurement-mode", types.YLeaf{"DelayMeasurementMode", delayProfile.DelayMeasurementMode})
    delayProfile.EntityData.Leafs.Append("periodic-advertisement-enabled", types.YLeaf{"PeriodicAdvertisementEnabled", delayProfile.PeriodicAdvertisementEnabled})
    delayProfile.EntityData.Leafs.Append("periodic-advertisement-interval", types.YLeaf{"PeriodicAdvertisementInterval", delayProfile.PeriodicAdvertisementInterval})
    delayProfile.EntityData.Leafs.Append("effective-periodic-advertisement-interval", types.YLeaf{"EffectivePeriodicAdvertisementInterval", delayProfile.EffectivePeriodicAdvertisementInterval})
    delayProfile.EntityData.Leafs.Append("periodic-advertisement-threshold", types.YLeaf{"PeriodicAdvertisementThreshold", delayProfile.PeriodicAdvertisementThreshold})
    delayProfile.EntityData.Leafs.Append("periodic-advertisement-minimum-change", types.YLeaf{"PeriodicAdvertisementMinimumChange", delayProfile.PeriodicAdvertisementMinimumChange})
    delayProfile.EntityData.Leafs.Append("accelerated-advertisement-threshold", types.YLeaf{"AcceleratedAdvertisementThreshold", delayProfile.AcceleratedAdvertisementThreshold})
    delayProfile.EntityData.Leafs.Append("accelerated-advertisement-minimum-change", types.YLeaf{"AcceleratedAdvertisementMinimumChange", delayProfile.AcceleratedAdvertisementMinimumChange})
    delayProfile.EntityData.Leafs.Append("accelerated-advertisement-enabled", types.YLeaf{"AcceleratedAdvertisementEnabled", delayProfile.AcceleratedAdvertisementEnabled})

    delayProfile.EntityData.YListKeys = []string {}

    return &(delayProfile.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters
// PM delay counters for a transport types
type PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic counters for a PM interface instance.
    GenericCounters PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_GenericCounters

    // Exclusive counters for a PM interface instance.
    ExclusiveCounters PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_ExclusiveCounters
}

func (delayTransportCounters *PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters) GetEntityData() *types.CommonEntityData {
    delayTransportCounters.EntityData.YFilter = delayTransportCounters.YFilter
    delayTransportCounters.EntityData.YangName = "delay-transport-counters"
    delayTransportCounters.EntityData.BundleName = "cisco_ios_xr"
    delayTransportCounters.EntityData.ParentYangName = "interface-delay-summary"
    delayTransportCounters.EntityData.SegmentPath = "delay-transport-counters"
    delayTransportCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/summary/delay-summary/interface-delay-summary/" + delayTransportCounters.EntityData.SegmentPath
    delayTransportCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayTransportCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayTransportCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayTransportCounters.EntityData.Children = types.NewOrderedMap()
    delayTransportCounters.EntityData.Children.Append("generic-counters", types.YChild{"GenericCounters", &delayTransportCounters.GenericCounters})
    delayTransportCounters.EntityData.Children.Append("exclusive-counters", types.YChild{"ExclusiveCounters", &delayTransportCounters.ExclusiveCounters})
    delayTransportCounters.EntityData.Leafs = types.NewOrderedMap()

    delayTransportCounters.EntityData.YListKeys = []string {}

    return &(delayTransportCounters.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_GenericCounters
// Generic counters for a PM interface instance
type PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_GenericCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Query packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketsSent interface{}

    // Query packets sent error. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketSentErrors interface{}

    // Query packet sent error, no IP address. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketSentErrorNoIpAddress interface{}

    // Query packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketsReceived interface{}

    // Received packet error, negative delay. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketErrorNegativeDelay interface{}

    // Received packet error, delay exceeds threshold. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketErrorDelayExceedsThreshold interface{}

    // Received packet error, missing Tx timestamp. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorMissingTxTimestamp interface{}

    // Received packet error, missing Rx timestamp. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorMissingRxTimestamp interface{}

    // Received packet error, probe full. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketErrorProbeFull interface{}

    // Received packet error, probe not started. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorProbeNotStarted interface{}

    // Received packet with a control code error. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketControlCodeError interface{}

    // Received packet with a control code notification. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketControlCodeNotification interface{}

    // Probes started. The type is interface{} with range:
    // 0..18446744073709551615.
    ProbesStarted interface{}

    // Probes completed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProbesComplete interface{}

    // Probes incomplete. The type is interface{} with range:
    // 0..18446744073709551615.
    ProbesIncomplete interface{}

    // Advertisements. The type is interface{} with range:
    // 0..18446744073709551615.
    Advertisement interface{}
}

func (genericCounters *PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_GenericCounters) GetEntityData() *types.CommonEntityData {
    genericCounters.EntityData.YFilter = genericCounters.YFilter
    genericCounters.EntityData.YangName = "generic-counters"
    genericCounters.EntityData.BundleName = "cisco_ios_xr"
    genericCounters.EntityData.ParentYangName = "delay-transport-counters"
    genericCounters.EntityData.SegmentPath = "generic-counters"
    genericCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/summary/delay-summary/interface-delay-summary/delay-transport-counters/" + genericCounters.EntityData.SegmentPath
    genericCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounters.EntityData.Children = types.NewOrderedMap()
    genericCounters.EntityData.Leafs = types.NewOrderedMap()
    genericCounters.EntityData.Leafs.Append("query-packets-sent", types.YLeaf{"QueryPacketsSent", genericCounters.QueryPacketsSent})
    genericCounters.EntityData.Leafs.Append("query-packet-sent-errors", types.YLeaf{"QueryPacketSentErrors", genericCounters.QueryPacketSentErrors})
    genericCounters.EntityData.Leafs.Append("query-packet-sent-error-no-ip-address", types.YLeaf{"QueryPacketSentErrorNoIpAddress", genericCounters.QueryPacketSentErrorNoIpAddress})
    genericCounters.EntityData.Leafs.Append("query-packets-received", types.YLeaf{"QueryPacketsReceived", genericCounters.QueryPacketsReceived})
    genericCounters.EntityData.Leafs.Append("received-packet-error-negative-delay", types.YLeaf{"ReceivedPacketErrorNegativeDelay", genericCounters.ReceivedPacketErrorNegativeDelay})
    genericCounters.EntityData.Leafs.Append("received-packet-error-delay-exceeds-threshold", types.YLeaf{"ReceivedPacketErrorDelayExceedsThreshold", genericCounters.ReceivedPacketErrorDelayExceedsThreshold})
    genericCounters.EntityData.Leafs.Append("received-packet-error-missing-tx-timestamp", types.YLeaf{"ReceivedPacketErrorMissingTxTimestamp", genericCounters.ReceivedPacketErrorMissingTxTimestamp})
    genericCounters.EntityData.Leafs.Append("received-packet-error-missing-rx-timestamp", types.YLeaf{"ReceivedPacketErrorMissingRxTimestamp", genericCounters.ReceivedPacketErrorMissingRxTimestamp})
    genericCounters.EntityData.Leafs.Append("received-packet-error-probe-full", types.YLeaf{"ReceivedPacketErrorProbeFull", genericCounters.ReceivedPacketErrorProbeFull})
    genericCounters.EntityData.Leafs.Append("received-packet-error-probe-not-started", types.YLeaf{"ReceivedPacketErrorProbeNotStarted", genericCounters.ReceivedPacketErrorProbeNotStarted})
    genericCounters.EntityData.Leafs.Append("received-packet-control-code-error", types.YLeaf{"ReceivedPacketControlCodeError", genericCounters.ReceivedPacketControlCodeError})
    genericCounters.EntityData.Leafs.Append("received-packet-control-code-notification", types.YLeaf{"ReceivedPacketControlCodeNotification", genericCounters.ReceivedPacketControlCodeNotification})
    genericCounters.EntityData.Leafs.Append("probes-started", types.YLeaf{"ProbesStarted", genericCounters.ProbesStarted})
    genericCounters.EntityData.Leafs.Append("probes-complete", types.YLeaf{"ProbesComplete", genericCounters.ProbesComplete})
    genericCounters.EntityData.Leafs.Append("probes-incomplete", types.YLeaf{"ProbesIncomplete", genericCounters.ProbesIncomplete})
    genericCounters.EntityData.Leafs.Append("advertisement", types.YLeaf{"Advertisement", genericCounters.Advertisement})

    genericCounters.EntityData.YListKeys = []string {}

    return &(genericCounters.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_ExclusiveCounters
// Exclusive counters for a PM interface instance
type PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_ExclusiveCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is PmTransport.
    Type interface{}

    // Counters Exclusive for interface.
    InterfaceExclusiveCounters PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_ExclusiveCounters_InterfaceExclusiveCounters
}

func (exclusiveCounters *PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_ExclusiveCounters) GetEntityData() *types.CommonEntityData {
    exclusiveCounters.EntityData.YFilter = exclusiveCounters.YFilter
    exclusiveCounters.EntityData.YangName = "exclusive-counters"
    exclusiveCounters.EntityData.BundleName = "cisco_ios_xr"
    exclusiveCounters.EntityData.ParentYangName = "delay-transport-counters"
    exclusiveCounters.EntityData.SegmentPath = "exclusive-counters"
    exclusiveCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/summary/delay-summary/interface-delay-summary/delay-transport-counters/" + exclusiveCounters.EntityData.SegmentPath
    exclusiveCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exclusiveCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exclusiveCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exclusiveCounters.EntityData.Children = types.NewOrderedMap()
    exclusiveCounters.EntityData.Children.Append("interface-exclusive-counters", types.YChild{"InterfaceExclusiveCounters", &exclusiveCounters.InterfaceExclusiveCounters})
    exclusiveCounters.EntityData.Leafs = types.NewOrderedMap()
    exclusiveCounters.EntityData.Leafs.Append("type", types.YLeaf{"Type", exclusiveCounters.Type})

    exclusiveCounters.EntityData.YListKeys = []string {}

    return &(exclusiveCounters.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_ExclusiveCounters_InterfaceExclusiveCounters
// Counters Exclusive for interface
type PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_ExclusiveCounters_InterfaceExclusiveCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Query packet sent error, interface down. The type is interface{} with
    // range: 0..18446744073709551615.
    QueryPacketSentErrorInterfaceDown interface{}

    // Query packet sent error, no MPLS caps. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketSentErrorNoMplsCaps interface{}
}

func (interfaceExclusiveCounters *PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_InterfaceDelaySummary_DelayTransportCounters_ExclusiveCounters_InterfaceExclusiveCounters) GetEntityData() *types.CommonEntityData {
    interfaceExclusiveCounters.EntityData.YFilter = interfaceExclusiveCounters.YFilter
    interfaceExclusiveCounters.EntityData.YangName = "interface-exclusive-counters"
    interfaceExclusiveCounters.EntityData.BundleName = "cisco_ios_xr"
    interfaceExclusiveCounters.EntityData.ParentYangName = "exclusive-counters"
    interfaceExclusiveCounters.EntityData.SegmentPath = "interface-exclusive-counters"
    interfaceExclusiveCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/summary/delay-summary/interface-delay-summary/delay-transport-counters/exclusive-counters/" + interfaceExclusiveCounters.EntityData.SegmentPath
    interfaceExclusiveCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceExclusiveCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceExclusiveCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceExclusiveCounters.EntityData.Children = types.NewOrderedMap()
    interfaceExclusiveCounters.EntityData.Leafs = types.NewOrderedMap()
    interfaceExclusiveCounters.EntityData.Leafs.Append("query-packet-sent-error-interface-down", types.YLeaf{"QueryPacketSentErrorInterfaceDown", interfaceExclusiveCounters.QueryPacketSentErrorInterfaceDown})
    interfaceExclusiveCounters.EntityData.Leafs.Append("query-packet-sent-error-no-mpls-caps", types.YLeaf{"QueryPacketSentErrorNoMplsCaps", interfaceExclusiveCounters.QueryPacketSentErrorNoMplsCaps})

    interfaceExclusiveCounters.EntityData.YListKeys = []string {}

    return &(interfaceExclusiveCounters.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_DelayGlobalCounters
// PM delay global counters
type PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_DelayGlobalCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Query packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketsSent interface{}

    // Query packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketsReceived interface{}

    // Received packet error, invalid session ID. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorInvalidSessionId interface{}

    // Received packet error, no session. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketErrorNoSession interface{}
}

func (delayGlobalCounters *PerformanceMeasurement_Nodes_Node_Summary_DelaySummary_DelayGlobalCounters) GetEntityData() *types.CommonEntityData {
    delayGlobalCounters.EntityData.YFilter = delayGlobalCounters.YFilter
    delayGlobalCounters.EntityData.YangName = "delay-global-counters"
    delayGlobalCounters.EntityData.BundleName = "cisco_ios_xr"
    delayGlobalCounters.EntityData.ParentYangName = "delay-summary"
    delayGlobalCounters.EntityData.SegmentPath = "delay-global-counters"
    delayGlobalCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/summary/delay-summary/" + delayGlobalCounters.EntityData.SegmentPath
    delayGlobalCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayGlobalCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayGlobalCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayGlobalCounters.EntityData.Children = types.NewOrderedMap()
    delayGlobalCounters.EntityData.Leafs = types.NewOrderedMap()
    delayGlobalCounters.EntityData.Leafs.Append("query-packets-sent", types.YLeaf{"QueryPacketsSent", delayGlobalCounters.QueryPacketsSent})
    delayGlobalCounters.EntityData.Leafs.Append("query-packets-received", types.YLeaf{"QueryPacketsReceived", delayGlobalCounters.QueryPacketsReceived})
    delayGlobalCounters.EntityData.Leafs.Append("received-packet-error-invalid-session-id", types.YLeaf{"ReceivedPacketErrorInvalidSessionId", delayGlobalCounters.ReceivedPacketErrorInvalidSessionId})
    delayGlobalCounters.EntityData.Leafs.Append("received-packet-error-no-session", types.YLeaf{"ReceivedPacketErrorNoSession", delayGlobalCounters.ReceivedPacketErrorNoSession})

    delayGlobalCounters.EntityData.YListKeys = []string {}

    return &(delayGlobalCounters.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces
// Interfaces
type PerformanceMeasurement_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface detailed table.
    InterfaceDetails PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails

    // Delay-measurement intformation.
    InterfaceDelay PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay
}

func (interfaces *PerformanceMeasurement_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface-details", types.YChild{"InterfaceDetails", &interfaces.InterfaceDetails})
    interfaces.EntityData.Children.Append("interface-delay", types.YChild{"InterfaceDelay", &interfaces.InterfaceDelay})
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails
// Interface detailed table
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed interface information. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail.
    InterfaceDetail []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail
}

func (interfaceDetails *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails) GetEntityData() *types.CommonEntityData {
    interfaceDetails.EntityData.YFilter = interfaceDetails.YFilter
    interfaceDetails.EntityData.YangName = "interface-details"
    interfaceDetails.EntityData.BundleName = "cisco_ios_xr"
    interfaceDetails.EntityData.ParentYangName = "interfaces"
    interfaceDetails.EntityData.SegmentPath = "interface-details"
    interfaceDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/" + interfaceDetails.EntityData.SegmentPath
    interfaceDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDetails.EntityData.Children = types.NewOrderedMap()
    interfaceDetails.EntityData.Children.Append("interface-detail", types.YChild{"InterfaceDetail", nil})
    for i := range interfaceDetails.InterfaceDetail {
        interfaceDetails.EntityData.Children.Append(types.GetSegmentPath(interfaceDetails.InterfaceDetail[i]), types.YChild{"InterfaceDetail", interfaceDetails.InterfaceDetail[i]})
    }
    interfaceDetails.EntityData.Leafs = types.NewOrderedMap()

    interfaceDetails.EntityData.YListKeys = []string {}

    return &(interfaceDetails.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail
// Detailed interface information
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface state. The type is bool.
    InterfaceState interface{}

    // Source Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Source IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceV6Address interface{}

    // Source Mac address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SourceMacAddress interface{}

    // Primary VLAN Tag. The type is interface{} with range: 0..65535.
    PrimaryVlanTag interface{}

    // Secondary VLAN Tag. The type is interface{} with range: 0..65535.
    SecondaryVlanTag interface{}

    // Delay-measurement sessions. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession.
    DelayMeasurementSession []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession
}

func (interfaceDetail *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail) GetEntityData() *types.CommonEntityData {
    interfaceDetail.EntityData.YFilter = interfaceDetail.YFilter
    interfaceDetail.EntityData.YangName = "interface-detail"
    interfaceDetail.EntityData.BundleName = "cisco_ios_xr"
    interfaceDetail.EntityData.ParentYangName = "interface-details"
    interfaceDetail.EntityData.SegmentPath = "interface-detail" + types.AddKeyToken(interfaceDetail.InterfaceName, "interface-name")
    interfaceDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/" + interfaceDetail.EntityData.SegmentPath
    interfaceDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDetail.EntityData.Children = types.NewOrderedMap()
    interfaceDetail.EntityData.Children.Append("delay-measurement-session", types.YChild{"DelayMeasurementSession", nil})
    for i := range interfaceDetail.DelayMeasurementSession {
        types.SetYListKey(interfaceDetail.DelayMeasurementSession[i], i)
        interfaceDetail.EntityData.Children.Append(types.GetSegmentPath(interfaceDetail.DelayMeasurementSession[i]), types.YChild{"DelayMeasurementSession", interfaceDetail.DelayMeasurementSession[i]})
    }
    interfaceDetail.EntityData.Leafs = types.NewOrderedMap()
    interfaceDetail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceDetail.InterfaceName})
    interfaceDetail.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", interfaceDetail.InterfaceNameXr})
    interfaceDetail.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", interfaceDetail.InterfaceHandle})
    interfaceDetail.EntityData.Leafs.Append("interface-state", types.YLeaf{"InterfaceState", interfaceDetail.InterfaceState})
    interfaceDetail.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", interfaceDetail.SourceAddress})
    interfaceDetail.EntityData.Leafs.Append("source-v6-address", types.YLeaf{"SourceV6Address", interfaceDetail.SourceV6Address})
    interfaceDetail.EntityData.Leafs.Append("source-mac-address", types.YLeaf{"SourceMacAddress", interfaceDetail.SourceMacAddress})
    interfaceDetail.EntityData.Leafs.Append("primary-vlan-tag", types.YLeaf{"PrimaryVlanTag", interfaceDetail.PrimaryVlanTag})
    interfaceDetail.EntityData.Leafs.Append("secondary-vlan-tag", types.YLeaf{"SecondaryVlanTag", interfaceDetail.SecondaryVlanTag})

    interfaceDetail.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceDetail.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession
// Delay-measurement sessions
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Transport Type. The type is PmTransport.
    TransportType interface{}

    // Measurement Type. The type is PmMeasurement.
    MeasurementType interface{}

    // Periodic advertisement interval in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    PeriodicAdvertisementIntervalInSec interface{}

    // Information for the current probe.
    CurrentProbe PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_CurrentProbe

    // Session counters.
    SessionCounters PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters

    // Last advertisement information.
    LastAdvertisementInformation PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastAdvertisementInformation

    // Next advertisement information.
    NextAdvertisementInformation PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_NextAdvertisementInformation

    // Last notifcation control code received.
    LastNotificationControlCode PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastNotificationControlCode

    // Last error control code received.
    LastErrorControlCode PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastErrorControlCode

    // Current probe history. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_ProbeHistory.
    ProbeHistory []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_ProbeHistory
}

func (delayMeasurementSession *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession) GetEntityData() *types.CommonEntityData {
    delayMeasurementSession.EntityData.YFilter = delayMeasurementSession.YFilter
    delayMeasurementSession.EntityData.YangName = "delay-measurement-session"
    delayMeasurementSession.EntityData.BundleName = "cisco_ios_xr"
    delayMeasurementSession.EntityData.ParentYangName = "interface-detail"
    delayMeasurementSession.EntityData.SegmentPath = "delay-measurement-session" + types.AddNoKeyToken(delayMeasurementSession)
    delayMeasurementSession.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/" + delayMeasurementSession.EntityData.SegmentPath
    delayMeasurementSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delayMeasurementSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delayMeasurementSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delayMeasurementSession.EntityData.Children = types.NewOrderedMap()
    delayMeasurementSession.EntityData.Children.Append("current-probe", types.YChild{"CurrentProbe", &delayMeasurementSession.CurrentProbe})
    delayMeasurementSession.EntityData.Children.Append("session-counters", types.YChild{"SessionCounters", &delayMeasurementSession.SessionCounters})
    delayMeasurementSession.EntityData.Children.Append("last-advertisement-information", types.YChild{"LastAdvertisementInformation", &delayMeasurementSession.LastAdvertisementInformation})
    delayMeasurementSession.EntityData.Children.Append("next-advertisement-information", types.YChild{"NextAdvertisementInformation", &delayMeasurementSession.NextAdvertisementInformation})
    delayMeasurementSession.EntityData.Children.Append("last-notification-control-code", types.YChild{"LastNotificationControlCode", &delayMeasurementSession.LastNotificationControlCode})
    delayMeasurementSession.EntityData.Children.Append("last-error-control-code", types.YChild{"LastErrorControlCode", &delayMeasurementSession.LastErrorControlCode})
    delayMeasurementSession.EntityData.Children.Append("probe-history", types.YChild{"ProbeHistory", nil})
    for i := range delayMeasurementSession.ProbeHistory {
        types.SetYListKey(delayMeasurementSession.ProbeHistory[i], i)
        delayMeasurementSession.EntityData.Children.Append(types.GetSegmentPath(delayMeasurementSession.ProbeHistory[i]), types.YChild{"ProbeHistory", delayMeasurementSession.ProbeHistory[i]})
    }
    delayMeasurementSession.EntityData.Leafs = types.NewOrderedMap()
    delayMeasurementSession.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", delayMeasurementSession.SessionId})
    delayMeasurementSession.EntityData.Leafs.Append("transport-type", types.YLeaf{"TransportType", delayMeasurementSession.TransportType})
    delayMeasurementSession.EntityData.Leafs.Append("measurement-type", types.YLeaf{"MeasurementType", delayMeasurementSession.MeasurementType})
    delayMeasurementSession.EntityData.Leafs.Append("periodic-advertisement-interval-in-sec", types.YLeaf{"PeriodicAdvertisementIntervalInSec", delayMeasurementSession.PeriodicAdvertisementIntervalInSec})

    delayMeasurementSession.EntityData.YListKeys = []string {}

    return &(delayMeasurementSession.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_CurrentProbe
// Information for the current probe
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_CurrentProbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp in milliseconds of the current probe start (milliseconds since
    // Jan. 1, 1970). The type is interface{} with range: 0..18446744073709551615.
    // Units are millisecond.
    ProbeStartTimeStamp interface{}

    // Time in milliseconds until the next probe starts. The type is interface{}
    // with range: 0..4294967295. Units are millisecond.
    NextProbeStartTimeRemainingInMillisec interface{}

    // Time in milliseconds until the next packet is sent. The type is interface{}
    // with range: 0..4294967295. Units are millisecond.
    NextPacketSentTimeRemainingInMillisec interface{}

    // Number of packets sent in the current probe. The type is interface{} with
    // range: 0..4294967295.
    NumberOfPacketsSent interface{}

    // Number of packets received in the current probe. The type is interface{}
    // with range: 0..4294967295.
    NumberOfPacketsReceived interface{}

    // Reason why probe is not running. The type is PmProbeNotRunningReason.
    ProbeNotRunningReason interface{}

    // Summarized  results of the current probe.
    ProbeResults PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_CurrentProbe_ProbeResults
}

func (currentProbe *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_CurrentProbe) GetEntityData() *types.CommonEntityData {
    currentProbe.EntityData.YFilter = currentProbe.YFilter
    currentProbe.EntityData.YangName = "current-probe"
    currentProbe.EntityData.BundleName = "cisco_ios_xr"
    currentProbe.EntityData.ParentYangName = "delay-measurement-session"
    currentProbe.EntityData.SegmentPath = "current-probe"
    currentProbe.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/" + currentProbe.EntityData.SegmentPath
    currentProbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentProbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentProbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentProbe.EntityData.Children = types.NewOrderedMap()
    currentProbe.EntityData.Children.Append("probe-results", types.YChild{"ProbeResults", &currentProbe.ProbeResults})
    currentProbe.EntityData.Leafs = types.NewOrderedMap()
    currentProbe.EntityData.Leafs.Append("probe-start-time-stamp", types.YLeaf{"ProbeStartTimeStamp", currentProbe.ProbeStartTimeStamp})
    currentProbe.EntityData.Leafs.Append("next-probe-start-time-remaining-in-millisec", types.YLeaf{"NextProbeStartTimeRemainingInMillisec", currentProbe.NextProbeStartTimeRemainingInMillisec})
    currentProbe.EntityData.Leafs.Append("next-packet-sent-time-remaining-in-millisec", types.YLeaf{"NextPacketSentTimeRemainingInMillisec", currentProbe.NextPacketSentTimeRemainingInMillisec})
    currentProbe.EntityData.Leafs.Append("number-of-packets-sent", types.YLeaf{"NumberOfPacketsSent", currentProbe.NumberOfPacketsSent})
    currentProbe.EntityData.Leafs.Append("number-of-packets-received", types.YLeaf{"NumberOfPacketsReceived", currentProbe.NumberOfPacketsReceived})
    currentProbe.EntityData.Leafs.Append("probe-not-running-reason", types.YLeaf{"ProbeNotRunningReason", currentProbe.ProbeNotRunningReason})

    currentProbe.EntityData.YListKeys = []string {}

    return &(currentProbe.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_CurrentProbe_ProbeResults
// Summarized  results of the current probe
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_CurrentProbe_ProbeResults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average value (uSec). The type is interface{} with range: 0..4294967295.
    Average interface{}

    // Minimum value (uSec). The type is interface{} with range: 0..4294967295.
    Minimum interface{}

    // Maximum value (uSec). The type is interface{} with range: 0..4294967295.
    Maximum interface{}

    // Variance value (uSec). The type is interface{} with range: 0..4294967295.
    Variance interface{}
}

func (probeResults *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_CurrentProbe_ProbeResults) GetEntityData() *types.CommonEntityData {
    probeResults.EntityData.YFilter = probeResults.YFilter
    probeResults.EntityData.YangName = "probe-results"
    probeResults.EntityData.BundleName = "cisco_ios_xr"
    probeResults.EntityData.ParentYangName = "current-probe"
    probeResults.EntityData.SegmentPath = "probe-results"
    probeResults.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/current-probe/" + probeResults.EntityData.SegmentPath
    probeResults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probeResults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probeResults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probeResults.EntityData.Children = types.NewOrderedMap()
    probeResults.EntityData.Leafs = types.NewOrderedMap()
    probeResults.EntityData.Leafs.Append("average", types.YLeaf{"Average", probeResults.Average})
    probeResults.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", probeResults.Minimum})
    probeResults.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", probeResults.Maximum})
    probeResults.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", probeResults.Variance})

    probeResults.EntityData.YListKeys = []string {}

    return &(probeResults.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters
// Session counters
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic counters for a PM interface instance.
    GenericCounters PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_GenericCounters

    // Exclusive counters for a PM interface instance.
    ExclusiveCounters PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_ExclusiveCounters
}

func (sessionCounters *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters) GetEntityData() *types.CommonEntityData {
    sessionCounters.EntityData.YFilter = sessionCounters.YFilter
    sessionCounters.EntityData.YangName = "session-counters"
    sessionCounters.EntityData.BundleName = "cisco_ios_xr"
    sessionCounters.EntityData.ParentYangName = "delay-measurement-session"
    sessionCounters.EntityData.SegmentPath = "session-counters"
    sessionCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/" + sessionCounters.EntityData.SegmentPath
    sessionCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionCounters.EntityData.Children = types.NewOrderedMap()
    sessionCounters.EntityData.Children.Append("generic-counters", types.YChild{"GenericCounters", &sessionCounters.GenericCounters})
    sessionCounters.EntityData.Children.Append("exclusive-counters", types.YChild{"ExclusiveCounters", &sessionCounters.ExclusiveCounters})
    sessionCounters.EntityData.Leafs = types.NewOrderedMap()

    sessionCounters.EntityData.YListKeys = []string {}

    return &(sessionCounters.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_GenericCounters
// Generic counters for a PM interface instance
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_GenericCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Query packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketsSent interface{}

    // Query packets sent error. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketSentErrors interface{}

    // Query packet sent error, no IP address. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketSentErrorNoIpAddress interface{}

    // Query packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketsReceived interface{}

    // Received packet error, negative delay. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketErrorNegativeDelay interface{}

    // Received packet error, delay exceeds threshold. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketErrorDelayExceedsThreshold interface{}

    // Received packet error, missing Tx timestamp. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorMissingTxTimestamp interface{}

    // Received packet error, missing Rx timestamp. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorMissingRxTimestamp interface{}

    // Received packet error, probe full. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketErrorProbeFull interface{}

    // Received packet error, probe not started. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorProbeNotStarted interface{}

    // Received packet with a control code error. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketControlCodeError interface{}

    // Received packet with a control code notification. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketControlCodeNotification interface{}

    // Probes started. The type is interface{} with range:
    // 0..18446744073709551615.
    ProbesStarted interface{}

    // Probes completed. The type is interface{} with range:
    // 0..18446744073709551615.
    ProbesComplete interface{}

    // Probes incomplete. The type is interface{} with range:
    // 0..18446744073709551615.
    ProbesIncomplete interface{}

    // Advertisements. The type is interface{} with range:
    // 0..18446744073709551615.
    Advertisement interface{}
}

func (genericCounters *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_GenericCounters) GetEntityData() *types.CommonEntityData {
    genericCounters.EntityData.YFilter = genericCounters.YFilter
    genericCounters.EntityData.YangName = "generic-counters"
    genericCounters.EntityData.BundleName = "cisco_ios_xr"
    genericCounters.EntityData.ParentYangName = "session-counters"
    genericCounters.EntityData.SegmentPath = "generic-counters"
    genericCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/session-counters/" + genericCounters.EntityData.SegmentPath
    genericCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericCounters.EntityData.Children = types.NewOrderedMap()
    genericCounters.EntityData.Leafs = types.NewOrderedMap()
    genericCounters.EntityData.Leafs.Append("query-packets-sent", types.YLeaf{"QueryPacketsSent", genericCounters.QueryPacketsSent})
    genericCounters.EntityData.Leafs.Append("query-packet-sent-errors", types.YLeaf{"QueryPacketSentErrors", genericCounters.QueryPacketSentErrors})
    genericCounters.EntityData.Leafs.Append("query-packet-sent-error-no-ip-address", types.YLeaf{"QueryPacketSentErrorNoIpAddress", genericCounters.QueryPacketSentErrorNoIpAddress})
    genericCounters.EntityData.Leafs.Append("query-packets-received", types.YLeaf{"QueryPacketsReceived", genericCounters.QueryPacketsReceived})
    genericCounters.EntityData.Leafs.Append("received-packet-error-negative-delay", types.YLeaf{"ReceivedPacketErrorNegativeDelay", genericCounters.ReceivedPacketErrorNegativeDelay})
    genericCounters.EntityData.Leafs.Append("received-packet-error-delay-exceeds-threshold", types.YLeaf{"ReceivedPacketErrorDelayExceedsThreshold", genericCounters.ReceivedPacketErrorDelayExceedsThreshold})
    genericCounters.EntityData.Leafs.Append("received-packet-error-missing-tx-timestamp", types.YLeaf{"ReceivedPacketErrorMissingTxTimestamp", genericCounters.ReceivedPacketErrorMissingTxTimestamp})
    genericCounters.EntityData.Leafs.Append("received-packet-error-missing-rx-timestamp", types.YLeaf{"ReceivedPacketErrorMissingRxTimestamp", genericCounters.ReceivedPacketErrorMissingRxTimestamp})
    genericCounters.EntityData.Leafs.Append("received-packet-error-probe-full", types.YLeaf{"ReceivedPacketErrorProbeFull", genericCounters.ReceivedPacketErrorProbeFull})
    genericCounters.EntityData.Leafs.Append("received-packet-error-probe-not-started", types.YLeaf{"ReceivedPacketErrorProbeNotStarted", genericCounters.ReceivedPacketErrorProbeNotStarted})
    genericCounters.EntityData.Leafs.Append("received-packet-control-code-error", types.YLeaf{"ReceivedPacketControlCodeError", genericCounters.ReceivedPacketControlCodeError})
    genericCounters.EntityData.Leafs.Append("received-packet-control-code-notification", types.YLeaf{"ReceivedPacketControlCodeNotification", genericCounters.ReceivedPacketControlCodeNotification})
    genericCounters.EntityData.Leafs.Append("probes-started", types.YLeaf{"ProbesStarted", genericCounters.ProbesStarted})
    genericCounters.EntityData.Leafs.Append("probes-complete", types.YLeaf{"ProbesComplete", genericCounters.ProbesComplete})
    genericCounters.EntityData.Leafs.Append("probes-incomplete", types.YLeaf{"ProbesIncomplete", genericCounters.ProbesIncomplete})
    genericCounters.EntityData.Leafs.Append("advertisement", types.YLeaf{"Advertisement", genericCounters.Advertisement})

    genericCounters.EntityData.YListKeys = []string {}

    return &(genericCounters.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_ExclusiveCounters
// Exclusive counters for a PM interface instance
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_ExclusiveCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is PmTransport.
    Type interface{}

    // Counters Exclusive for interface.
    InterfaceExclusiveCounters PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_ExclusiveCounters_InterfaceExclusiveCounters
}

func (exclusiveCounters *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_ExclusiveCounters) GetEntityData() *types.CommonEntityData {
    exclusiveCounters.EntityData.YFilter = exclusiveCounters.YFilter
    exclusiveCounters.EntityData.YangName = "exclusive-counters"
    exclusiveCounters.EntityData.BundleName = "cisco_ios_xr"
    exclusiveCounters.EntityData.ParentYangName = "session-counters"
    exclusiveCounters.EntityData.SegmentPath = "exclusive-counters"
    exclusiveCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/session-counters/" + exclusiveCounters.EntityData.SegmentPath
    exclusiveCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exclusiveCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exclusiveCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exclusiveCounters.EntityData.Children = types.NewOrderedMap()
    exclusiveCounters.EntityData.Children.Append("interface-exclusive-counters", types.YChild{"InterfaceExclusiveCounters", &exclusiveCounters.InterfaceExclusiveCounters})
    exclusiveCounters.EntityData.Leafs = types.NewOrderedMap()
    exclusiveCounters.EntityData.Leafs.Append("type", types.YLeaf{"Type", exclusiveCounters.Type})

    exclusiveCounters.EntityData.YListKeys = []string {}

    return &(exclusiveCounters.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_ExclusiveCounters_InterfaceExclusiveCounters
// Counters Exclusive for interface
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_ExclusiveCounters_InterfaceExclusiveCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Query packet sent error, interface down. The type is interface{} with
    // range: 0..18446744073709551615.
    QueryPacketSentErrorInterfaceDown interface{}

    // Query packet sent error, no MPLS caps. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketSentErrorNoMplsCaps interface{}
}

func (interfaceExclusiveCounters *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_SessionCounters_ExclusiveCounters_InterfaceExclusiveCounters) GetEntityData() *types.CommonEntityData {
    interfaceExclusiveCounters.EntityData.YFilter = interfaceExclusiveCounters.YFilter
    interfaceExclusiveCounters.EntityData.YangName = "interface-exclusive-counters"
    interfaceExclusiveCounters.EntityData.BundleName = "cisco_ios_xr"
    interfaceExclusiveCounters.EntityData.ParentYangName = "exclusive-counters"
    interfaceExclusiveCounters.EntityData.SegmentPath = "interface-exclusive-counters"
    interfaceExclusiveCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/session-counters/exclusive-counters/" + interfaceExclusiveCounters.EntityData.SegmentPath
    interfaceExclusiveCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceExclusiveCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceExclusiveCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceExclusiveCounters.EntityData.Children = types.NewOrderedMap()
    interfaceExclusiveCounters.EntityData.Leafs = types.NewOrderedMap()
    interfaceExclusiveCounters.EntityData.Leafs.Append("query-packet-sent-error-interface-down", types.YLeaf{"QueryPacketSentErrorInterfaceDown", interfaceExclusiveCounters.QueryPacketSentErrorInterfaceDown})
    interfaceExclusiveCounters.EntityData.Leafs.Append("query-packet-sent-error-no-mpls-caps", types.YLeaf{"QueryPacketSentErrorNoMplsCaps", interfaceExclusiveCounters.QueryPacketSentErrorNoMplsCaps})

    interfaceExclusiveCounters.EntityData.YListKeys = []string {}

    return &(interfaceExclusiveCounters.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastAdvertisementInformation
// Last advertisement information
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastAdvertisementInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time of the advertisement (milliseconds since Jan. 1, 1970). The type is
    // interface{} with range: 0..18446744073709551615. Units are millisecond.
    TimeOfAdvertisement interface{}

    // Reason for advertisement. The type is PmAdvertReason.
    AdvertisementReason interface{}

    // Advertised values.
    AdvertisedValues PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastAdvertisementInformation_AdvertisedValues
}

func (lastAdvertisementInformation *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastAdvertisementInformation) GetEntityData() *types.CommonEntityData {
    lastAdvertisementInformation.EntityData.YFilter = lastAdvertisementInformation.YFilter
    lastAdvertisementInformation.EntityData.YangName = "last-advertisement-information"
    lastAdvertisementInformation.EntityData.BundleName = "cisco_ios_xr"
    lastAdvertisementInformation.EntityData.ParentYangName = "delay-measurement-session"
    lastAdvertisementInformation.EntityData.SegmentPath = "last-advertisement-information"
    lastAdvertisementInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/" + lastAdvertisementInformation.EntityData.SegmentPath
    lastAdvertisementInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastAdvertisementInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastAdvertisementInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastAdvertisementInformation.EntityData.Children = types.NewOrderedMap()
    lastAdvertisementInformation.EntityData.Children.Append("advertised-values", types.YChild{"AdvertisedValues", &lastAdvertisementInformation.AdvertisedValues})
    lastAdvertisementInformation.EntityData.Leafs = types.NewOrderedMap()
    lastAdvertisementInformation.EntityData.Leafs.Append("time-of-advertisement", types.YLeaf{"TimeOfAdvertisement", lastAdvertisementInformation.TimeOfAdvertisement})
    lastAdvertisementInformation.EntityData.Leafs.Append("advertisement-reason", types.YLeaf{"AdvertisementReason", lastAdvertisementInformation.AdvertisementReason})

    lastAdvertisementInformation.EntityData.YListKeys = []string {}

    return &(lastAdvertisementInformation.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastAdvertisementInformation_AdvertisedValues
// Advertised values
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastAdvertisementInformation_AdvertisedValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average value (uSec). The type is interface{} with range: 0..4294967295.
    Average interface{}

    // Minimum value (uSec). The type is interface{} with range: 0..4294967295.
    Minimum interface{}

    // Maximum value (uSec). The type is interface{} with range: 0..4294967295.
    Maximum interface{}

    // Variance value (uSec). The type is interface{} with range: 0..4294967295.
    Variance interface{}
}

func (advertisedValues *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastAdvertisementInformation_AdvertisedValues) GetEntityData() *types.CommonEntityData {
    advertisedValues.EntityData.YFilter = advertisedValues.YFilter
    advertisedValues.EntityData.YangName = "advertised-values"
    advertisedValues.EntityData.BundleName = "cisco_ios_xr"
    advertisedValues.EntityData.ParentYangName = "last-advertisement-information"
    advertisedValues.EntityData.SegmentPath = "advertised-values"
    advertisedValues.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/last-advertisement-information/" + advertisedValues.EntityData.SegmentPath
    advertisedValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertisedValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertisedValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertisedValues.EntityData.Children = types.NewOrderedMap()
    advertisedValues.EntityData.Leafs = types.NewOrderedMap()
    advertisedValues.EntityData.Leafs.Append("average", types.YLeaf{"Average", advertisedValues.Average})
    advertisedValues.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", advertisedValues.Minimum})
    advertisedValues.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", advertisedValues.Maximum})
    advertisedValues.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", advertisedValues.Variance})

    advertisedValues.EntityData.YListKeys = []string {}

    return &(advertisedValues.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_NextAdvertisementInformation
// Next advertisement information
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_NextAdvertisementInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Probes remaining until next periodic advertisement check. The type is
    // interface{} with range: 0..4294967295.
    AdvertisementIntervalProbesRemaining interface{}

    // Rolling average value (uSec). The type is interface{} with range:
    // 0..4294967295.
    RollingAverage interface{}

    // Advertisement interval values.
    AdvertisementIntervalValues PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_NextAdvertisementInformation_AdvertisementIntervalValues
}

func (nextAdvertisementInformation *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_NextAdvertisementInformation) GetEntityData() *types.CommonEntityData {
    nextAdvertisementInformation.EntityData.YFilter = nextAdvertisementInformation.YFilter
    nextAdvertisementInformation.EntityData.YangName = "next-advertisement-information"
    nextAdvertisementInformation.EntityData.BundleName = "cisco_ios_xr"
    nextAdvertisementInformation.EntityData.ParentYangName = "delay-measurement-session"
    nextAdvertisementInformation.EntityData.SegmentPath = "next-advertisement-information"
    nextAdvertisementInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/" + nextAdvertisementInformation.EntityData.SegmentPath
    nextAdvertisementInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextAdvertisementInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextAdvertisementInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextAdvertisementInformation.EntityData.Children = types.NewOrderedMap()
    nextAdvertisementInformation.EntityData.Children.Append("advertisement-interval-values", types.YChild{"AdvertisementIntervalValues", &nextAdvertisementInformation.AdvertisementIntervalValues})
    nextAdvertisementInformation.EntityData.Leafs = types.NewOrderedMap()
    nextAdvertisementInformation.EntityData.Leafs.Append("advertisement-interval-probes-remaining", types.YLeaf{"AdvertisementIntervalProbesRemaining", nextAdvertisementInformation.AdvertisementIntervalProbesRemaining})
    nextAdvertisementInformation.EntityData.Leafs.Append("rolling-average", types.YLeaf{"RollingAverage", nextAdvertisementInformation.RollingAverage})

    nextAdvertisementInformation.EntityData.YListKeys = []string {}

    return &(nextAdvertisementInformation.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_NextAdvertisementInformation_AdvertisementIntervalValues
// Advertisement interval values
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_NextAdvertisementInformation_AdvertisementIntervalValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average value (uSec). The type is interface{} with range: 0..4294967295.
    Average interface{}

    // Minimum value (uSec). The type is interface{} with range: 0..4294967295.
    Minimum interface{}

    // Maximum value (uSec). The type is interface{} with range: 0..4294967295.
    Maximum interface{}

    // Variance value (uSec). The type is interface{} with range: 0..4294967295.
    Variance interface{}
}

func (advertisementIntervalValues *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_NextAdvertisementInformation_AdvertisementIntervalValues) GetEntityData() *types.CommonEntityData {
    advertisementIntervalValues.EntityData.YFilter = advertisementIntervalValues.YFilter
    advertisementIntervalValues.EntityData.YangName = "advertisement-interval-values"
    advertisementIntervalValues.EntityData.BundleName = "cisco_ios_xr"
    advertisementIntervalValues.EntityData.ParentYangName = "next-advertisement-information"
    advertisementIntervalValues.EntityData.SegmentPath = "advertisement-interval-values"
    advertisementIntervalValues.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/next-advertisement-information/" + advertisementIntervalValues.EntityData.SegmentPath
    advertisementIntervalValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertisementIntervalValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertisementIntervalValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertisementIntervalValues.EntityData.Children = types.NewOrderedMap()
    advertisementIntervalValues.EntityData.Leafs = types.NewOrderedMap()
    advertisementIntervalValues.EntityData.Leafs.Append("average", types.YLeaf{"Average", advertisementIntervalValues.Average})
    advertisementIntervalValues.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", advertisementIntervalValues.Minimum})
    advertisementIntervalValues.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", advertisementIntervalValues.Maximum})
    advertisementIntervalValues.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", advertisementIntervalValues.Variance})

    advertisementIntervalValues.EntityData.YListKeys = []string {}

    return &(advertisementIntervalValues.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastNotificationControlCode
// Last notifcation control code received
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastNotificationControlCode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS PM RFC 6374 control code. The type is interface{} with range: 0..255.
    ControlCode interface{}

    // Received timestamp of the control code (milliseconds since Jan. 1, 1970).
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // millisecond.
    Timestamp interface{}
}

func (lastNotificationControlCode *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastNotificationControlCode) GetEntityData() *types.CommonEntityData {
    lastNotificationControlCode.EntityData.YFilter = lastNotificationControlCode.YFilter
    lastNotificationControlCode.EntityData.YangName = "last-notification-control-code"
    lastNotificationControlCode.EntityData.BundleName = "cisco_ios_xr"
    lastNotificationControlCode.EntityData.ParentYangName = "delay-measurement-session"
    lastNotificationControlCode.EntityData.SegmentPath = "last-notification-control-code"
    lastNotificationControlCode.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/" + lastNotificationControlCode.EntityData.SegmentPath
    lastNotificationControlCode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastNotificationControlCode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastNotificationControlCode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastNotificationControlCode.EntityData.Children = types.NewOrderedMap()
    lastNotificationControlCode.EntityData.Leafs = types.NewOrderedMap()
    lastNotificationControlCode.EntityData.Leafs.Append("control-code", types.YLeaf{"ControlCode", lastNotificationControlCode.ControlCode})
    lastNotificationControlCode.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", lastNotificationControlCode.Timestamp})

    lastNotificationControlCode.EntityData.YListKeys = []string {}

    return &(lastNotificationControlCode.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastErrorControlCode
// Last error control code received
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastErrorControlCode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS PM RFC 6374 control code. The type is interface{} with range: 0..255.
    ControlCode interface{}

    // Received timestamp of the control code (milliseconds since Jan. 1, 1970).
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // millisecond.
    Timestamp interface{}
}

func (lastErrorControlCode *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_LastErrorControlCode) GetEntityData() *types.CommonEntityData {
    lastErrorControlCode.EntityData.YFilter = lastErrorControlCode.YFilter
    lastErrorControlCode.EntityData.YangName = "last-error-control-code"
    lastErrorControlCode.EntityData.BundleName = "cisco_ios_xr"
    lastErrorControlCode.EntityData.ParentYangName = "delay-measurement-session"
    lastErrorControlCode.EntityData.SegmentPath = "last-error-control-code"
    lastErrorControlCode.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/" + lastErrorControlCode.EntityData.SegmentPath
    lastErrorControlCode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErrorControlCode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErrorControlCode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErrorControlCode.EntityData.Children = types.NewOrderedMap()
    lastErrorControlCode.EntityData.Leafs = types.NewOrderedMap()
    lastErrorControlCode.EntityData.Leafs.Append("control-code", types.YLeaf{"ControlCode", lastErrorControlCode.ControlCode})
    lastErrorControlCode.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", lastErrorControlCode.Timestamp})

    lastErrorControlCode.EntityData.YListKeys = []string {}

    return &(lastErrorControlCode.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_ProbeHistory
// Current probe history
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_ProbeHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Measurement value (nsec). The type is interface{} with range:
    // 0..18446744073709551615.
    MeasurementValue interface{}

    // Timestamp when the measurement was taken (milliseconds since Jan. 1, 1970).
    // The type is interface{} with range: 0..18446744073709551615. Units are
    // millisecond.
    QueryTimestamp interface{}
}

func (probeHistory *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDetails_InterfaceDetail_DelayMeasurementSession_ProbeHistory) GetEntityData() *types.CommonEntityData {
    probeHistory.EntityData.YFilter = probeHistory.YFilter
    probeHistory.EntityData.YangName = "probe-history"
    probeHistory.EntityData.BundleName = "cisco_ios_xr"
    probeHistory.EntityData.ParentYangName = "delay-measurement-session"
    probeHistory.EntityData.SegmentPath = "probe-history" + types.AddNoKeyToken(probeHistory)
    probeHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-details/interface-detail/delay-measurement-session/" + probeHistory.EntityData.SegmentPath
    probeHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probeHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probeHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probeHistory.EntityData.Children = types.NewOrderedMap()
    probeHistory.EntityData.Leafs = types.NewOrderedMap()
    probeHistory.EntityData.Leafs.Append("measurement-value", types.YLeaf{"MeasurementValue", probeHistory.MeasurementValue})
    probeHistory.EntityData.Leafs.Append("query-timestamp", types.YLeaf{"QueryTimestamp", probeHistory.QueryTimestamp})

    probeHistory.EntityData.YListKeys = []string {}

    return &(probeHistory.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay
// Delay-measurement intformation
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of last probe aggregation.
    InterfaceLastAggregations PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations

    // Table of probe histories.
    InterfaceProbeHistories PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories

    // Table of aggregated probe histories.
    InterfaceAggregatedHistories PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories

    // Table of last probes.
    InterfaceLastProbes PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes

    // Table of last advertisements.
    InterfaceLastAdvertisements PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements
}

func (interfaceDelay *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay) GetEntityData() *types.CommonEntityData {
    interfaceDelay.EntityData.YFilter = interfaceDelay.YFilter
    interfaceDelay.EntityData.YangName = "interface-delay"
    interfaceDelay.EntityData.BundleName = "cisco_ios_xr"
    interfaceDelay.EntityData.ParentYangName = "interfaces"
    interfaceDelay.EntityData.SegmentPath = "interface-delay"
    interfaceDelay.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/" + interfaceDelay.EntityData.SegmentPath
    interfaceDelay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDelay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDelay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDelay.EntityData.Children = types.NewOrderedMap()
    interfaceDelay.EntityData.Children.Append("interface-last-aggregations", types.YChild{"InterfaceLastAggregations", &interfaceDelay.InterfaceLastAggregations})
    interfaceDelay.EntityData.Children.Append("interface-probe-histories", types.YChild{"InterfaceProbeHistories", &interfaceDelay.InterfaceProbeHistories})
    interfaceDelay.EntityData.Children.Append("interface-aggregated-histories", types.YChild{"InterfaceAggregatedHistories", &interfaceDelay.InterfaceAggregatedHistories})
    interfaceDelay.EntityData.Children.Append("interface-last-probes", types.YChild{"InterfaceLastProbes", &interfaceDelay.InterfaceLastProbes})
    interfaceDelay.EntityData.Children.Append("interface-last-advertisements", types.YChild{"InterfaceLastAdvertisements", &interfaceDelay.InterfaceLastAdvertisements})
    interfaceDelay.EntityData.Leafs = types.NewOrderedMap()

    interfaceDelay.EntityData.YListKeys = []string {}

    return &(interfaceDelay.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations
// Table of last probe aggregation
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Last probe aggregation information. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation.
    InterfaceLastAggregation []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation
}

func (interfaceLastAggregations *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations) GetEntityData() *types.CommonEntityData {
    interfaceLastAggregations.EntityData.YFilter = interfaceLastAggregations.YFilter
    interfaceLastAggregations.EntityData.YangName = "interface-last-aggregations"
    interfaceLastAggregations.EntityData.BundleName = "cisco_ios_xr"
    interfaceLastAggregations.EntityData.ParentYangName = "interface-delay"
    interfaceLastAggregations.EntityData.SegmentPath = "interface-last-aggregations"
    interfaceLastAggregations.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/" + interfaceLastAggregations.EntityData.SegmentPath
    interfaceLastAggregations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLastAggregations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLastAggregations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLastAggregations.EntityData.Children = types.NewOrderedMap()
    interfaceLastAggregations.EntityData.Children.Append("interface-last-aggregation", types.YChild{"InterfaceLastAggregation", nil})
    for i := range interfaceLastAggregations.InterfaceLastAggregation {
        interfaceLastAggregations.EntityData.Children.Append(types.GetSegmentPath(interfaceLastAggregations.InterfaceLastAggregation[i]), types.YChild{"InterfaceLastAggregation", interfaceLastAggregations.InterfaceLastAggregation[i]})
    }
    interfaceLastAggregations.EntityData.Leafs = types.NewOrderedMap()

    interfaceLastAggregations.EntityData.YListKeys = []string {}

    return &(interfaceLastAggregations.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation
// Last probe aggregation information
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Last probe aggregation.
    LastAggregation PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation_LastAggregation
}

func (interfaceLastAggregation *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation) GetEntityData() *types.CommonEntityData {
    interfaceLastAggregation.EntityData.YFilter = interfaceLastAggregation.YFilter
    interfaceLastAggregation.EntityData.YangName = "interface-last-aggregation"
    interfaceLastAggregation.EntityData.BundleName = "cisco_ios_xr"
    interfaceLastAggregation.EntityData.ParentYangName = "interface-last-aggregations"
    interfaceLastAggregation.EntityData.SegmentPath = "interface-last-aggregation" + types.AddKeyToken(interfaceLastAggregation.InterfaceName, "interface-name")
    interfaceLastAggregation.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-last-aggregations/" + interfaceLastAggregation.EntityData.SegmentPath
    interfaceLastAggregation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLastAggregation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLastAggregation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLastAggregation.EntityData.Children = types.NewOrderedMap()
    interfaceLastAggregation.EntityData.Children.Append("last-aggregation", types.YChild{"LastAggregation", &interfaceLastAggregation.LastAggregation})
    interfaceLastAggregation.EntityData.Leafs = types.NewOrderedMap()
    interfaceLastAggregation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceLastAggregation.InterfaceName})

    interfaceLastAggregation.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceLastAggregation.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation_LastAggregation
// Last probe aggregation
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation_LastAggregation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time probe aggregation was done (milliseconds since Jan. 1, 1970). The type
    // is interface{} with range: 0..18446744073709551615. Units are millisecond.
    AggregationTimestamp interface{}

    // Action performed with the aggregated values. The type is PmAdvertReason.
    AggregationAction interface{}

    // Aggregated probe values.
    AggregatedProbeValues PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation_LastAggregation_AggregatedProbeValues
}

func (lastAggregation *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation_LastAggregation) GetEntityData() *types.CommonEntityData {
    lastAggregation.EntityData.YFilter = lastAggregation.YFilter
    lastAggregation.EntityData.YangName = "last-aggregation"
    lastAggregation.EntityData.BundleName = "cisco_ios_xr"
    lastAggregation.EntityData.ParentYangName = "interface-last-aggregation"
    lastAggregation.EntityData.SegmentPath = "last-aggregation"
    lastAggregation.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-last-aggregations/interface-last-aggregation/" + lastAggregation.EntityData.SegmentPath
    lastAggregation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastAggregation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastAggregation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastAggregation.EntityData.Children = types.NewOrderedMap()
    lastAggregation.EntityData.Children.Append("aggregated-probe-values", types.YChild{"AggregatedProbeValues", &lastAggregation.AggregatedProbeValues})
    lastAggregation.EntityData.Leafs = types.NewOrderedMap()
    lastAggregation.EntityData.Leafs.Append("aggregation-timestamp", types.YLeaf{"AggregationTimestamp", lastAggregation.AggregationTimestamp})
    lastAggregation.EntityData.Leafs.Append("aggregation-action", types.YLeaf{"AggregationAction", lastAggregation.AggregationAction})

    lastAggregation.EntityData.YListKeys = []string {}

    return &(lastAggregation.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation_LastAggregation_AggregatedProbeValues
// Aggregated probe values
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation_LastAggregation_AggregatedProbeValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average value (uSec). The type is interface{} with range: 0..4294967295.
    Average interface{}

    // Minimum value (uSec). The type is interface{} with range: 0..4294967295.
    Minimum interface{}

    // Maximum value (uSec). The type is interface{} with range: 0..4294967295.
    Maximum interface{}

    // Variance value (uSec). The type is interface{} with range: 0..4294967295.
    Variance interface{}
}

func (aggregatedProbeValues *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAggregations_InterfaceLastAggregation_LastAggregation_AggregatedProbeValues) GetEntityData() *types.CommonEntityData {
    aggregatedProbeValues.EntityData.YFilter = aggregatedProbeValues.YFilter
    aggregatedProbeValues.EntityData.YangName = "aggregated-probe-values"
    aggregatedProbeValues.EntityData.BundleName = "cisco_ios_xr"
    aggregatedProbeValues.EntityData.ParentYangName = "last-aggregation"
    aggregatedProbeValues.EntityData.SegmentPath = "aggregated-probe-values"
    aggregatedProbeValues.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-last-aggregations/interface-last-aggregation/last-aggregation/" + aggregatedProbeValues.EntityData.SegmentPath
    aggregatedProbeValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregatedProbeValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregatedProbeValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregatedProbeValues.EntityData.Children = types.NewOrderedMap()
    aggregatedProbeValues.EntityData.Leafs = types.NewOrderedMap()
    aggregatedProbeValues.EntityData.Leafs.Append("average", types.YLeaf{"Average", aggregatedProbeValues.Average})
    aggregatedProbeValues.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", aggregatedProbeValues.Minimum})
    aggregatedProbeValues.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", aggregatedProbeValues.Maximum})
    aggregatedProbeValues.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", aggregatedProbeValues.Variance})

    aggregatedProbeValues.EntityData.YListKeys = []string {}

    return &(aggregatedProbeValues.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories
// Table of probe histories
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Probe history information. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory.
    InterfaceProbeHistory []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory
}

func (interfaceProbeHistories *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories) GetEntityData() *types.CommonEntityData {
    interfaceProbeHistories.EntityData.YFilter = interfaceProbeHistories.YFilter
    interfaceProbeHistories.EntityData.YangName = "interface-probe-histories"
    interfaceProbeHistories.EntityData.BundleName = "cisco_ios_xr"
    interfaceProbeHistories.EntityData.ParentYangName = "interface-delay"
    interfaceProbeHistories.EntityData.SegmentPath = "interface-probe-histories"
    interfaceProbeHistories.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/" + interfaceProbeHistories.EntityData.SegmentPath
    interfaceProbeHistories.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProbeHistories.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProbeHistories.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProbeHistories.EntityData.Children = types.NewOrderedMap()
    interfaceProbeHistories.EntityData.Children.Append("interface-probe-history", types.YChild{"InterfaceProbeHistory", nil})
    for i := range interfaceProbeHistories.InterfaceProbeHistory {
        interfaceProbeHistories.EntityData.Children.Append(types.GetSegmentPath(interfaceProbeHistories.InterfaceProbeHistory[i]), types.YChild{"InterfaceProbeHistory", interfaceProbeHistories.InterfaceProbeHistory[i]})
    }
    interfaceProbeHistories.EntityData.Leafs = types.NewOrderedMap()

    interfaceProbeHistories.EntityData.YListKeys = []string {}

    return &(interfaceProbeHistories.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory
// Probe history information
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // History of previous probes. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory_History.
    History []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory_History
}

func (interfaceProbeHistory *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory) GetEntityData() *types.CommonEntityData {
    interfaceProbeHistory.EntityData.YFilter = interfaceProbeHistory.YFilter
    interfaceProbeHistory.EntityData.YangName = "interface-probe-history"
    interfaceProbeHistory.EntityData.BundleName = "cisco_ios_xr"
    interfaceProbeHistory.EntityData.ParentYangName = "interface-probe-histories"
    interfaceProbeHistory.EntityData.SegmentPath = "interface-probe-history" + types.AddKeyToken(interfaceProbeHistory.InterfaceName, "interface-name")
    interfaceProbeHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-probe-histories/" + interfaceProbeHistory.EntityData.SegmentPath
    interfaceProbeHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceProbeHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceProbeHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceProbeHistory.EntityData.Children = types.NewOrderedMap()
    interfaceProbeHistory.EntityData.Children.Append("history", types.YChild{"History", nil})
    for i := range interfaceProbeHistory.History {
        types.SetYListKey(interfaceProbeHistory.History[i], i)
        interfaceProbeHistory.EntityData.Children.Append(types.GetSegmentPath(interfaceProbeHistory.History[i]), types.YChild{"History", interfaceProbeHistory.History[i]})
    }
    interfaceProbeHistory.EntityData.Leafs = types.NewOrderedMap()
    interfaceProbeHistory.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceProbeHistory.InterfaceName})
    interfaceProbeHistory.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", interfaceProbeHistory.InterfaceNameXr})
    interfaceProbeHistory.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", interfaceProbeHistory.InterfaceHandle})

    interfaceProbeHistory.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceProbeHistory.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory_History
// History of previous probes
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Time last probe started (milliseconds since Jan. 1, 1970). The type is
    // interface{} with range: 0..18446744073709551615. Units are millisecond.
    ProbeStartTimestamp interface{}

    // Packets sent. The type is interface{} with range: 0..4294967295.
    PacketsSent interface{}

    // Packets received. The type is interface{} with range: 0..4294967295.
    PacketsReceived interface{}

    // Probe values.
    ProbeValues PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory_History_ProbeValues
}

func (history *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "interface-probe-history"
    history.EntityData.SegmentPath = "history" + types.AddNoKeyToken(history)
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-probe-histories/interface-probe-history/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("probe-values", types.YChild{"ProbeValues", &history.ProbeValues})
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("probe-start-timestamp", types.YLeaf{"ProbeStartTimestamp", history.ProbeStartTimestamp})
    history.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", history.PacketsSent})
    history.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", history.PacketsReceived})

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory_History_ProbeValues
// Probe values
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory_History_ProbeValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average value (uSec). The type is interface{} with range: 0..4294967295.
    Average interface{}

    // Minimum value (uSec). The type is interface{} with range: 0..4294967295.
    Minimum interface{}

    // Maximum value (uSec). The type is interface{} with range: 0..4294967295.
    Maximum interface{}

    // Variance value (uSec). The type is interface{} with range: 0..4294967295.
    Variance interface{}
}

func (probeValues *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceProbeHistories_InterfaceProbeHistory_History_ProbeValues) GetEntityData() *types.CommonEntityData {
    probeValues.EntityData.YFilter = probeValues.YFilter
    probeValues.EntityData.YangName = "probe-values"
    probeValues.EntityData.BundleName = "cisco_ios_xr"
    probeValues.EntityData.ParentYangName = "history"
    probeValues.EntityData.SegmentPath = "probe-values"
    probeValues.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-probe-histories/interface-probe-history/history/" + probeValues.EntityData.SegmentPath
    probeValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probeValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probeValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probeValues.EntityData.Children = types.NewOrderedMap()
    probeValues.EntityData.Leafs = types.NewOrderedMap()
    probeValues.EntityData.Leafs.Append("average", types.YLeaf{"Average", probeValues.Average})
    probeValues.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", probeValues.Minimum})
    probeValues.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", probeValues.Maximum})
    probeValues.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", probeValues.Variance})

    probeValues.EntityData.YListKeys = []string {}

    return &(probeValues.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories
// Table of aggregated probe histories
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Aggregated probe history information. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory.
    InterfaceAggregatedHistory []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory
}

func (interfaceAggregatedHistories *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories) GetEntityData() *types.CommonEntityData {
    interfaceAggregatedHistories.EntityData.YFilter = interfaceAggregatedHistories.YFilter
    interfaceAggregatedHistories.EntityData.YangName = "interface-aggregated-histories"
    interfaceAggregatedHistories.EntityData.BundleName = "cisco_ios_xr"
    interfaceAggregatedHistories.EntityData.ParentYangName = "interface-delay"
    interfaceAggregatedHistories.EntityData.SegmentPath = "interface-aggregated-histories"
    interfaceAggregatedHistories.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/" + interfaceAggregatedHistories.EntityData.SegmentPath
    interfaceAggregatedHistories.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAggregatedHistories.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAggregatedHistories.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAggregatedHistories.EntityData.Children = types.NewOrderedMap()
    interfaceAggregatedHistories.EntityData.Children.Append("interface-aggregated-history", types.YChild{"InterfaceAggregatedHistory", nil})
    for i := range interfaceAggregatedHistories.InterfaceAggregatedHistory {
        interfaceAggregatedHistories.EntityData.Children.Append(types.GetSegmentPath(interfaceAggregatedHistories.InterfaceAggregatedHistory[i]), types.YChild{"InterfaceAggregatedHistory", interfaceAggregatedHistories.InterfaceAggregatedHistory[i]})
    }
    interfaceAggregatedHistories.EntityData.Leafs = types.NewOrderedMap()

    interfaceAggregatedHistories.EntityData.YListKeys = []string {}

    return &(interfaceAggregatedHistories.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory
// Aggregated probe history information
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // History of previous probe aggregations. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory_History.
    History []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory_History
}

func (interfaceAggregatedHistory *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory) GetEntityData() *types.CommonEntityData {
    interfaceAggregatedHistory.EntityData.YFilter = interfaceAggregatedHistory.YFilter
    interfaceAggregatedHistory.EntityData.YangName = "interface-aggregated-history"
    interfaceAggregatedHistory.EntityData.BundleName = "cisco_ios_xr"
    interfaceAggregatedHistory.EntityData.ParentYangName = "interface-aggregated-histories"
    interfaceAggregatedHistory.EntityData.SegmentPath = "interface-aggregated-history" + types.AddKeyToken(interfaceAggregatedHistory.InterfaceName, "interface-name")
    interfaceAggregatedHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-aggregated-histories/" + interfaceAggregatedHistory.EntityData.SegmentPath
    interfaceAggregatedHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAggregatedHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAggregatedHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAggregatedHistory.EntityData.Children = types.NewOrderedMap()
    interfaceAggregatedHistory.EntityData.Children.Append("history", types.YChild{"History", nil})
    for i := range interfaceAggregatedHistory.History {
        types.SetYListKey(interfaceAggregatedHistory.History[i], i)
        interfaceAggregatedHistory.EntityData.Children.Append(types.GetSegmentPath(interfaceAggregatedHistory.History[i]), types.YChild{"History", interfaceAggregatedHistory.History[i]})
    }
    interfaceAggregatedHistory.EntityData.Leafs = types.NewOrderedMap()
    interfaceAggregatedHistory.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceAggregatedHistory.InterfaceName})
    interfaceAggregatedHistory.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", interfaceAggregatedHistory.InterfaceNameXr})
    interfaceAggregatedHistory.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", interfaceAggregatedHistory.InterfaceHandle})

    interfaceAggregatedHistory.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceAggregatedHistory.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory_History
// History of previous probe aggregations
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Time probe aggregation was done (milliseconds since Jan. 1, 1970). The type
    // is interface{} with range: 0..18446744073709551615. Units are millisecond.
    AggregationTimestamp interface{}

    // Action performed with the aggregated values. The type is PmAdvertReason.
    AggregationAction interface{}

    // Aggregated probe values.
    AggregatedProbeValues PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory_History_AggregatedProbeValues
}

func (history *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "interface-aggregated-history"
    history.EntityData.SegmentPath = "history" + types.AddNoKeyToken(history)
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-aggregated-histories/interface-aggregated-history/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("aggregated-probe-values", types.YChild{"AggregatedProbeValues", &history.AggregatedProbeValues})
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("aggregation-timestamp", types.YLeaf{"AggregationTimestamp", history.AggregationTimestamp})
    history.EntityData.Leafs.Append("aggregation-action", types.YLeaf{"AggregationAction", history.AggregationAction})

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory_History_AggregatedProbeValues
// Aggregated probe values
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory_History_AggregatedProbeValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average value (uSec). The type is interface{} with range: 0..4294967295.
    Average interface{}

    // Minimum value (uSec). The type is interface{} with range: 0..4294967295.
    Minimum interface{}

    // Maximum value (uSec). The type is interface{} with range: 0..4294967295.
    Maximum interface{}

    // Variance value (uSec). The type is interface{} with range: 0..4294967295.
    Variance interface{}
}

func (aggregatedProbeValues *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceAggregatedHistories_InterfaceAggregatedHistory_History_AggregatedProbeValues) GetEntityData() *types.CommonEntityData {
    aggregatedProbeValues.EntityData.YFilter = aggregatedProbeValues.YFilter
    aggregatedProbeValues.EntityData.YangName = "aggregated-probe-values"
    aggregatedProbeValues.EntityData.BundleName = "cisco_ios_xr"
    aggregatedProbeValues.EntityData.ParentYangName = "history"
    aggregatedProbeValues.EntityData.SegmentPath = "aggregated-probe-values"
    aggregatedProbeValues.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-aggregated-histories/interface-aggregated-history/history/" + aggregatedProbeValues.EntityData.SegmentPath
    aggregatedProbeValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregatedProbeValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregatedProbeValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregatedProbeValues.EntityData.Children = types.NewOrderedMap()
    aggregatedProbeValues.EntityData.Leafs = types.NewOrderedMap()
    aggregatedProbeValues.EntityData.Leafs.Append("average", types.YLeaf{"Average", aggregatedProbeValues.Average})
    aggregatedProbeValues.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", aggregatedProbeValues.Minimum})
    aggregatedProbeValues.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", aggregatedProbeValues.Maximum})
    aggregatedProbeValues.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", aggregatedProbeValues.Variance})

    aggregatedProbeValues.EntityData.YListKeys = []string {}

    return &(aggregatedProbeValues.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes
// Table of last probes
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Last measurement information. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe.
    InterfaceLastProbe []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe
}

func (interfaceLastProbes *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes) GetEntityData() *types.CommonEntityData {
    interfaceLastProbes.EntityData.YFilter = interfaceLastProbes.YFilter
    interfaceLastProbes.EntityData.YangName = "interface-last-probes"
    interfaceLastProbes.EntityData.BundleName = "cisco_ios_xr"
    interfaceLastProbes.EntityData.ParentYangName = "interface-delay"
    interfaceLastProbes.EntityData.SegmentPath = "interface-last-probes"
    interfaceLastProbes.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/" + interfaceLastProbes.EntityData.SegmentPath
    interfaceLastProbes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLastProbes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLastProbes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLastProbes.EntityData.Children = types.NewOrderedMap()
    interfaceLastProbes.EntityData.Children.Append("interface-last-probe", types.YChild{"InterfaceLastProbe", nil})
    for i := range interfaceLastProbes.InterfaceLastProbe {
        interfaceLastProbes.EntityData.Children.Append(types.GetSegmentPath(interfaceLastProbes.InterfaceLastProbe[i]), types.YChild{"InterfaceLastProbe", interfaceLastProbes.InterfaceLastProbe[i]})
    }
    interfaceLastProbes.EntityData.Leafs = types.NewOrderedMap()

    interfaceLastProbes.EntityData.YListKeys = []string {}

    return &(interfaceLastProbes.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe
// Last measurement information
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Last probe.
    LastProbe PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe_LastProbe
}

func (interfaceLastProbe *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe) GetEntityData() *types.CommonEntityData {
    interfaceLastProbe.EntityData.YFilter = interfaceLastProbe.YFilter
    interfaceLastProbe.EntityData.YangName = "interface-last-probe"
    interfaceLastProbe.EntityData.BundleName = "cisco_ios_xr"
    interfaceLastProbe.EntityData.ParentYangName = "interface-last-probes"
    interfaceLastProbe.EntityData.SegmentPath = "interface-last-probe" + types.AddKeyToken(interfaceLastProbe.InterfaceName, "interface-name")
    interfaceLastProbe.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-last-probes/" + interfaceLastProbe.EntityData.SegmentPath
    interfaceLastProbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLastProbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLastProbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLastProbe.EntityData.Children = types.NewOrderedMap()
    interfaceLastProbe.EntityData.Children.Append("last-probe", types.YChild{"LastProbe", &interfaceLastProbe.LastProbe})
    interfaceLastProbe.EntityData.Leafs = types.NewOrderedMap()
    interfaceLastProbe.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceLastProbe.InterfaceName})

    interfaceLastProbe.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceLastProbe.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe_LastProbe
// Last probe
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe_LastProbe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time last probe started (milliseconds since Jan. 1, 1970). The type is
    // interface{} with range: 0..18446744073709551615. Units are millisecond.
    ProbeStartTimestamp interface{}

    // Packets sent. The type is interface{} with range: 0..4294967295.
    PacketsSent interface{}

    // Packets received. The type is interface{} with range: 0..4294967295.
    PacketsReceived interface{}

    // Probe values.
    ProbeValues PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe_LastProbe_ProbeValues
}

func (lastProbe *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe_LastProbe) GetEntityData() *types.CommonEntityData {
    lastProbe.EntityData.YFilter = lastProbe.YFilter
    lastProbe.EntityData.YangName = "last-probe"
    lastProbe.EntityData.BundleName = "cisco_ios_xr"
    lastProbe.EntityData.ParentYangName = "interface-last-probe"
    lastProbe.EntityData.SegmentPath = "last-probe"
    lastProbe.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-last-probes/interface-last-probe/" + lastProbe.EntityData.SegmentPath
    lastProbe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastProbe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastProbe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastProbe.EntityData.Children = types.NewOrderedMap()
    lastProbe.EntityData.Children.Append("probe-values", types.YChild{"ProbeValues", &lastProbe.ProbeValues})
    lastProbe.EntityData.Leafs = types.NewOrderedMap()
    lastProbe.EntityData.Leafs.Append("probe-start-timestamp", types.YLeaf{"ProbeStartTimestamp", lastProbe.ProbeStartTimestamp})
    lastProbe.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", lastProbe.PacketsSent})
    lastProbe.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", lastProbe.PacketsReceived})

    lastProbe.EntityData.YListKeys = []string {}

    return &(lastProbe.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe_LastProbe_ProbeValues
// Probe values
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe_LastProbe_ProbeValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average value (uSec). The type is interface{} with range: 0..4294967295.
    Average interface{}

    // Minimum value (uSec). The type is interface{} with range: 0..4294967295.
    Minimum interface{}

    // Maximum value (uSec). The type is interface{} with range: 0..4294967295.
    Maximum interface{}

    // Variance value (uSec). The type is interface{} with range: 0..4294967295.
    Variance interface{}
}

func (probeValues *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastProbes_InterfaceLastProbe_LastProbe_ProbeValues) GetEntityData() *types.CommonEntityData {
    probeValues.EntityData.YFilter = probeValues.YFilter
    probeValues.EntityData.YangName = "probe-values"
    probeValues.EntityData.BundleName = "cisco_ios_xr"
    probeValues.EntityData.ParentYangName = "last-probe"
    probeValues.EntityData.SegmentPath = "probe-values"
    probeValues.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-last-probes/interface-last-probe/last-probe/" + probeValues.EntityData.SegmentPath
    probeValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    probeValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    probeValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    probeValues.EntityData.Children = types.NewOrderedMap()
    probeValues.EntityData.Leafs = types.NewOrderedMap()
    probeValues.EntityData.Leafs.Append("average", types.YLeaf{"Average", probeValues.Average})
    probeValues.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", probeValues.Minimum})
    probeValues.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", probeValues.Maximum})
    probeValues.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", probeValues.Variance})

    probeValues.EntityData.YListKeys = []string {}

    return &(probeValues.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements
// Table of last advertisements
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Last advertisement information. The type is slice of
    // PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement.
    InterfaceLastAdvertisement []*PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement
}

func (interfaceLastAdvertisements *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements) GetEntityData() *types.CommonEntityData {
    interfaceLastAdvertisements.EntityData.YFilter = interfaceLastAdvertisements.YFilter
    interfaceLastAdvertisements.EntityData.YangName = "interface-last-advertisements"
    interfaceLastAdvertisements.EntityData.BundleName = "cisco_ios_xr"
    interfaceLastAdvertisements.EntityData.ParentYangName = "interface-delay"
    interfaceLastAdvertisements.EntityData.SegmentPath = "interface-last-advertisements"
    interfaceLastAdvertisements.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/" + interfaceLastAdvertisements.EntityData.SegmentPath
    interfaceLastAdvertisements.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLastAdvertisements.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLastAdvertisements.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLastAdvertisements.EntityData.Children = types.NewOrderedMap()
    interfaceLastAdvertisements.EntityData.Children.Append("interface-last-advertisement", types.YChild{"InterfaceLastAdvertisement", nil})
    for i := range interfaceLastAdvertisements.InterfaceLastAdvertisement {
        interfaceLastAdvertisements.EntityData.Children.Append(types.GetSegmentPath(interfaceLastAdvertisements.InterfaceLastAdvertisement[i]), types.YChild{"InterfaceLastAdvertisement", interfaceLastAdvertisements.InterfaceLastAdvertisement[i]})
    }
    interfaceLastAdvertisements.EntityData.Leafs = types.NewOrderedMap()

    interfaceLastAdvertisements.EntityData.YListKeys = []string {}

    return &(interfaceLastAdvertisements.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement
// Last advertisement information
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Last advertisement.
    LastAdvertisement PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement_LastAdvertisement
}

func (interfaceLastAdvertisement *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement) GetEntityData() *types.CommonEntityData {
    interfaceLastAdvertisement.EntityData.YFilter = interfaceLastAdvertisement.YFilter
    interfaceLastAdvertisement.EntityData.YangName = "interface-last-advertisement"
    interfaceLastAdvertisement.EntityData.BundleName = "cisco_ios_xr"
    interfaceLastAdvertisement.EntityData.ParentYangName = "interface-last-advertisements"
    interfaceLastAdvertisement.EntityData.SegmentPath = "interface-last-advertisement" + types.AddKeyToken(interfaceLastAdvertisement.InterfaceName, "interface-name")
    interfaceLastAdvertisement.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-last-advertisements/" + interfaceLastAdvertisement.EntityData.SegmentPath
    interfaceLastAdvertisement.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceLastAdvertisement.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceLastAdvertisement.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceLastAdvertisement.EntityData.Children = types.NewOrderedMap()
    interfaceLastAdvertisement.EntityData.Children.Append("last-advertisement", types.YChild{"LastAdvertisement", &interfaceLastAdvertisement.LastAdvertisement})
    interfaceLastAdvertisement.EntityData.Leafs = types.NewOrderedMap()
    interfaceLastAdvertisement.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceLastAdvertisement.InterfaceName})

    interfaceLastAdvertisement.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceLastAdvertisement.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement_LastAdvertisement
// Last advertisement
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement_LastAdvertisement struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time of the advertisement (milliseconds since Jan. 1, 1970). The type is
    // interface{} with range: 0..18446744073709551615. Units are millisecond.
    TimeOfAdvertisement interface{}

    // Reason for advertisement. The type is PmAdvertReason.
    AdvertisementReason interface{}

    // Advertised values.
    AdvertisedValues PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement_LastAdvertisement_AdvertisedValues
}

func (lastAdvertisement *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement_LastAdvertisement) GetEntityData() *types.CommonEntityData {
    lastAdvertisement.EntityData.YFilter = lastAdvertisement.YFilter
    lastAdvertisement.EntityData.YangName = "last-advertisement"
    lastAdvertisement.EntityData.BundleName = "cisco_ios_xr"
    lastAdvertisement.EntityData.ParentYangName = "interface-last-advertisement"
    lastAdvertisement.EntityData.SegmentPath = "last-advertisement"
    lastAdvertisement.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-last-advertisements/interface-last-advertisement/" + lastAdvertisement.EntityData.SegmentPath
    lastAdvertisement.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastAdvertisement.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastAdvertisement.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastAdvertisement.EntityData.Children = types.NewOrderedMap()
    lastAdvertisement.EntityData.Children.Append("advertised-values", types.YChild{"AdvertisedValues", &lastAdvertisement.AdvertisedValues})
    lastAdvertisement.EntityData.Leafs = types.NewOrderedMap()
    lastAdvertisement.EntityData.Leafs.Append("time-of-advertisement", types.YLeaf{"TimeOfAdvertisement", lastAdvertisement.TimeOfAdvertisement})
    lastAdvertisement.EntityData.Leafs.Append("advertisement-reason", types.YLeaf{"AdvertisementReason", lastAdvertisement.AdvertisementReason})

    lastAdvertisement.EntityData.YListKeys = []string {}

    return &(lastAdvertisement.EntityData)
}

// PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement_LastAdvertisement_AdvertisedValues
// Advertised values
type PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement_LastAdvertisement_AdvertisedValues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Average value (uSec). The type is interface{} with range: 0..4294967295.
    Average interface{}

    // Minimum value (uSec). The type is interface{} with range: 0..4294967295.
    Minimum interface{}

    // Maximum value (uSec). The type is interface{} with range: 0..4294967295.
    Maximum interface{}

    // Variance value (uSec). The type is interface{} with range: 0..4294967295.
    Variance interface{}
}

func (advertisedValues *PerformanceMeasurement_Nodes_Node_Interfaces_InterfaceDelay_InterfaceLastAdvertisements_InterfaceLastAdvertisement_LastAdvertisement_AdvertisedValues) GetEntityData() *types.CommonEntityData {
    advertisedValues.EntityData.YFilter = advertisedValues.YFilter
    advertisedValues.EntityData.YangName = "advertised-values"
    advertisedValues.EntityData.BundleName = "cisco_ios_xr"
    advertisedValues.EntityData.ParentYangName = "last-advertisement"
    advertisedValues.EntityData.SegmentPath = "advertised-values"
    advertisedValues.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement/nodes/node/interfaces/interface-delay/interface-last-advertisements/interface-last-advertisement/last-advertisement/" + advertisedValues.EntityData.SegmentPath
    advertisedValues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    advertisedValues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    advertisedValues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    advertisedValues.EntityData.Children = types.NewOrderedMap()
    advertisedValues.EntityData.Leafs = types.NewOrderedMap()
    advertisedValues.EntityData.Leafs.Append("average", types.YLeaf{"Average", advertisedValues.Average})
    advertisedValues.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", advertisedValues.Minimum})
    advertisedValues.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", advertisedValues.Maximum})
    advertisedValues.EntityData.Leafs.Append("variance", types.YLeaf{"Variance", advertisedValues.Variance})

    advertisedValues.EntityData.YListKeys = []string {}

    return &(advertisedValues.EntityData)
}

// PerformanceMeasurementResponder
// performance measurement responder
type PerformanceMeasurementResponder struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node table for node-specific operational data.
    Nodes PerformanceMeasurementResponder_Nodes
}

func (performanceMeasurementResponder *PerformanceMeasurementResponder) GetEntityData() *types.CommonEntityData {
    performanceMeasurementResponder.EntityData.YFilter = performanceMeasurementResponder.YFilter
    performanceMeasurementResponder.EntityData.YangName = "performance-measurement-responder"
    performanceMeasurementResponder.EntityData.BundleName = "cisco_ios_xr"
    performanceMeasurementResponder.EntityData.ParentYangName = "Cisco-IOS-XR-perf-meas-oper"
    performanceMeasurementResponder.EntityData.SegmentPath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement-responder"
    performanceMeasurementResponder.EntityData.AbsolutePath = performanceMeasurementResponder.EntityData.SegmentPath
    performanceMeasurementResponder.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    performanceMeasurementResponder.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    performanceMeasurementResponder.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    performanceMeasurementResponder.EntityData.Children = types.NewOrderedMap()
    performanceMeasurementResponder.EntityData.Children.Append("nodes", types.YChild{"Nodes", &performanceMeasurementResponder.Nodes})
    performanceMeasurementResponder.EntityData.Leafs = types.NewOrderedMap()

    performanceMeasurementResponder.EntityData.YListKeys = []string {}

    return &(performanceMeasurementResponder.EntityData)
}

// PerformanceMeasurementResponder_Nodes
// Node table for node-specific operational data
type PerformanceMeasurementResponder_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // PerformanceMeasurementResponder_Nodes_Node.
    Node []*PerformanceMeasurementResponder_Nodes_Node
}

func (nodes *PerformanceMeasurementResponder_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "performance-measurement-responder"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement-responder/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// PerformanceMeasurementResponder_Nodes_Node
// Node-specific data for a particular node
type PerformanceMeasurementResponder_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Summary.
    Summary PerformanceMeasurementResponder_Nodes_Node_Summary

    // Table of interfaces.
    Interfaces PerformanceMeasurementResponder_Nodes_Node_Interfaces
}

func (node *PerformanceMeasurementResponder_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Node, "node")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement-responder/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node", types.YLeaf{"Node", node.Node})

    node.EntityData.YListKeys = []string {"Node"}

    return &(node.EntityData)
}

// PerformanceMeasurementResponder_Nodes_Node_Summary
// Summary
type PerformanceMeasurementResponder_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of interfaces in the responder cache. The type is interface{} with
    // range: 0..4294967295.
    TotalInterfaces interface{}

    // Global incoming packet rate. The type is interface{} with range:
    // 0..4294967295.
    PacketRate interface{}

    // Global incoming packet rate high water mark. The type is interface{} with
    // range: 0..4294967295.
    PacketRateHighWaterMark interface{}

    // Global responder counters.
    ResponderCounters PerformanceMeasurementResponder_Nodes_Node_Summary_ResponderCounters
}

func (summary *PerformanceMeasurementResponder_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement-responder/nodes/node/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("responder-counters", types.YChild{"ResponderCounters", &summary.ResponderCounters})
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("total-interfaces", types.YLeaf{"TotalInterfaces", summary.TotalInterfaces})
    summary.EntityData.Leafs.Append("packet-rate", types.YLeaf{"PacketRate", summary.PacketRate})
    summary.EntityData.Leafs.Append("packet-rate-high-water-mark", types.YLeaf{"PacketRateHighWaterMark", summary.PacketRateHighWaterMark})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// PerformanceMeasurementResponder_Nodes_Node_Summary_ResponderCounters
// Global responder counters
type PerformanceMeasurementResponder_Nodes_Node_Summary_ResponderCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Response packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ReplyPacketSent interface{}

    // Response packets sent error. The type is interface{} with range:
    // 0..18446744073709551615.
    ReplyPacketSentError interface{}

    // Response packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketReceived interface{}

    // Received packet error, URO TLV not present. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorUrotlvNotPresent interface{}

    // Received packet error, invalid source port number. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketErrorInvalidSourcePortNumber interface{}

    // Received packet error, no source address. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorNoSourceAddress interface{}

    // Received packet error, no return path. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketErrorNoReturnPath interface{}

    // Received packet error, invalid querier control code. The type is
    // interface{} with range: 0..18446744073709551615.
    ReceivedPacketErrorInvalidQuerierControlCode interface{}

    // Received packet error, unsupported timestamp format. The type is
    // interface{} with range: 0..18446744073709551615.
    ReceivedPacketErrorUnsupportedTimestampFormat interface{}

    // Received packet error, timestamp not available. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketErrorTimestampNotAvailable interface{}

    // Received packet error, unsupported mandatory TLV. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketErrorUnsupportedMandatoryTlv interface{}

    // Received packet error, invalid packet. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketErrorInvalidPacket interface{}
}

func (responderCounters *PerformanceMeasurementResponder_Nodes_Node_Summary_ResponderCounters) GetEntityData() *types.CommonEntityData {
    responderCounters.EntityData.YFilter = responderCounters.YFilter
    responderCounters.EntityData.YangName = "responder-counters"
    responderCounters.EntityData.BundleName = "cisco_ios_xr"
    responderCounters.EntityData.ParentYangName = "summary"
    responderCounters.EntityData.SegmentPath = "responder-counters"
    responderCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement-responder/nodes/node/summary/" + responderCounters.EntityData.SegmentPath
    responderCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    responderCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    responderCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    responderCounters.EntityData.Children = types.NewOrderedMap()
    responderCounters.EntityData.Leafs = types.NewOrderedMap()
    responderCounters.EntityData.Leafs.Append("reply-packet-sent", types.YLeaf{"ReplyPacketSent", responderCounters.ReplyPacketSent})
    responderCounters.EntityData.Leafs.Append("reply-packet-sent-error", types.YLeaf{"ReplyPacketSentError", responderCounters.ReplyPacketSentError})
    responderCounters.EntityData.Leafs.Append("query-packet-received", types.YLeaf{"QueryPacketReceived", responderCounters.QueryPacketReceived})
    responderCounters.EntityData.Leafs.Append("received-packet-error-urotlv-not-present", types.YLeaf{"ReceivedPacketErrorUrotlvNotPresent", responderCounters.ReceivedPacketErrorUrotlvNotPresent})
    responderCounters.EntityData.Leafs.Append("received-packet-error-invalid-source-port-number", types.YLeaf{"ReceivedPacketErrorInvalidSourcePortNumber", responderCounters.ReceivedPacketErrorInvalidSourcePortNumber})
    responderCounters.EntityData.Leafs.Append("received-packet-error-no-source-address", types.YLeaf{"ReceivedPacketErrorNoSourceAddress", responderCounters.ReceivedPacketErrorNoSourceAddress})
    responderCounters.EntityData.Leafs.Append("received-packet-error-no-return-path", types.YLeaf{"ReceivedPacketErrorNoReturnPath", responderCounters.ReceivedPacketErrorNoReturnPath})
    responderCounters.EntityData.Leafs.Append("received-packet-error-invalid-querier-control-code", types.YLeaf{"ReceivedPacketErrorInvalidQuerierControlCode", responderCounters.ReceivedPacketErrorInvalidQuerierControlCode})
    responderCounters.EntityData.Leafs.Append("received-packet-error-unsupported-timestamp-format", types.YLeaf{"ReceivedPacketErrorUnsupportedTimestampFormat", responderCounters.ReceivedPacketErrorUnsupportedTimestampFormat})
    responderCounters.EntityData.Leafs.Append("received-packet-error-timestamp-not-available", types.YLeaf{"ReceivedPacketErrorTimestampNotAvailable", responderCounters.ReceivedPacketErrorTimestampNotAvailable})
    responderCounters.EntityData.Leafs.Append("received-packet-error-unsupported-mandatory-tlv", types.YLeaf{"ReceivedPacketErrorUnsupportedMandatoryTlv", responderCounters.ReceivedPacketErrorUnsupportedMandatoryTlv})
    responderCounters.EntityData.Leafs.Append("received-packet-error-invalid-packet", types.YLeaf{"ReceivedPacketErrorInvalidPacket", responderCounters.ReceivedPacketErrorInvalidPacket})

    responderCounters.EntityData.YListKeys = []string {}

    return &(responderCounters.EntityData)
}

// PerformanceMeasurementResponder_Nodes_Node_Interfaces
// Table of interfaces
type PerformanceMeasurementResponder_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface information. The type is slice of
    // PerformanceMeasurementResponder_Nodes_Node_Interfaces_Interface.
    Interface []*PerformanceMeasurementResponder_Nodes_Node_Interfaces_Interface
}

func (interfaces *PerformanceMeasurementResponder_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement-responder/nodes/node/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// PerformanceMeasurementResponder_Nodes_Node_Interfaces_Interface
// Interface information
type PerformanceMeasurementResponder_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    InterfaceName interface{}

    // Interface name. The type is string.
    InterfaceNameXr interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Source Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Source V6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceV6Address interface{}

    // Incoming packet rate. The type is interface{} with range: 0..4294967295.
    PacketRate interface{}

    // Incoming packet rate high water mark. The type is interface{} with range:
    // 0..4294967295.
    PacketRateHighWaterMark interface{}

    // Seconds until an inactive interface is cleaned up. The type is interface{}
    // with range: 0..4294967295. Units are second.
    CleanupTimeRemaining interface{}

    // Per interface responder counters.
    InterfaceCounters PerformanceMeasurementResponder_Nodes_Node_Interfaces_Interface_InterfaceCounters
}

func (self *PerformanceMeasurementResponder_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement-responder/nodes/node/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("interface-counters", types.YChild{"InterfaceCounters", &self.InterfaceCounters})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", self.InterfaceNameXr})
    self.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", self.InterfaceHandle})
    self.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", self.SourceAddress})
    self.EntityData.Leafs.Append("source-v6-address", types.YLeaf{"SourceV6Address", self.SourceV6Address})
    self.EntityData.Leafs.Append("packet-rate", types.YLeaf{"PacketRate", self.PacketRate})
    self.EntityData.Leafs.Append("packet-rate-high-water-mark", types.YLeaf{"PacketRateHighWaterMark", self.PacketRateHighWaterMark})
    self.EntityData.Leafs.Append("cleanup-time-remaining", types.YLeaf{"CleanupTimeRemaining", self.CleanupTimeRemaining})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// PerformanceMeasurementResponder_Nodes_Node_Interfaces_Interface_InterfaceCounters
// Per interface responder counters
type PerformanceMeasurementResponder_Nodes_Node_Interfaces_Interface_InterfaceCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Response packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ReplyPacketSent interface{}

    // Response packets sent error. The type is interface{} with range:
    // 0..18446744073709551615.
    ReplyPacketSentError interface{}

    // Response packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    QueryPacketReceived interface{}

    // Received packet error, URO TLV not present. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorUrotlvNotPresent interface{}

    // Received packet error, invalid source port number. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketErrorInvalidSourcePortNumber interface{}

    // Received packet error, no source address. The type is interface{} with
    // range: 0..18446744073709551615.
    ReceivedPacketErrorNoSourceAddress interface{}

    // Received packet error, no return path. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketErrorNoReturnPath interface{}

    // Received packet error, invalid querier control code. The type is
    // interface{} with range: 0..18446744073709551615.
    ReceivedPacketErrorInvalidQuerierControlCode interface{}

    // Received packet error, unsupported timestamp format. The type is
    // interface{} with range: 0..18446744073709551615.
    ReceivedPacketErrorUnsupportedTimestampFormat interface{}

    // Received packet error, timestamp not available. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketErrorTimestampNotAvailable interface{}

    // Received packet error, unsupported mandatory TLV. The type is interface{}
    // with range: 0..18446744073709551615.
    ReceivedPacketErrorUnsupportedMandatoryTlv interface{}

    // Received packet error, invalid packet. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketErrorInvalidPacket interface{}
}

func (interfaceCounters *PerformanceMeasurementResponder_Nodes_Node_Interfaces_Interface_InterfaceCounters) GetEntityData() *types.CommonEntityData {
    interfaceCounters.EntityData.YFilter = interfaceCounters.YFilter
    interfaceCounters.EntityData.YangName = "interface-counters"
    interfaceCounters.EntityData.BundleName = "cisco_ios_xr"
    interfaceCounters.EntityData.ParentYangName = "interface"
    interfaceCounters.EntityData.SegmentPath = "interface-counters"
    interfaceCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-perf-meas-oper:performance-measurement-responder/nodes/node/interfaces/interface/" + interfaceCounters.EntityData.SegmentPath
    interfaceCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceCounters.EntityData.Children = types.NewOrderedMap()
    interfaceCounters.EntityData.Leafs = types.NewOrderedMap()
    interfaceCounters.EntityData.Leafs.Append("reply-packet-sent", types.YLeaf{"ReplyPacketSent", interfaceCounters.ReplyPacketSent})
    interfaceCounters.EntityData.Leafs.Append("reply-packet-sent-error", types.YLeaf{"ReplyPacketSentError", interfaceCounters.ReplyPacketSentError})
    interfaceCounters.EntityData.Leafs.Append("query-packet-received", types.YLeaf{"QueryPacketReceived", interfaceCounters.QueryPacketReceived})
    interfaceCounters.EntityData.Leafs.Append("received-packet-error-urotlv-not-present", types.YLeaf{"ReceivedPacketErrorUrotlvNotPresent", interfaceCounters.ReceivedPacketErrorUrotlvNotPresent})
    interfaceCounters.EntityData.Leafs.Append("received-packet-error-invalid-source-port-number", types.YLeaf{"ReceivedPacketErrorInvalidSourcePortNumber", interfaceCounters.ReceivedPacketErrorInvalidSourcePortNumber})
    interfaceCounters.EntityData.Leafs.Append("received-packet-error-no-source-address", types.YLeaf{"ReceivedPacketErrorNoSourceAddress", interfaceCounters.ReceivedPacketErrorNoSourceAddress})
    interfaceCounters.EntityData.Leafs.Append("received-packet-error-no-return-path", types.YLeaf{"ReceivedPacketErrorNoReturnPath", interfaceCounters.ReceivedPacketErrorNoReturnPath})
    interfaceCounters.EntityData.Leafs.Append("received-packet-error-invalid-querier-control-code", types.YLeaf{"ReceivedPacketErrorInvalidQuerierControlCode", interfaceCounters.ReceivedPacketErrorInvalidQuerierControlCode})
    interfaceCounters.EntityData.Leafs.Append("received-packet-error-unsupported-timestamp-format", types.YLeaf{"ReceivedPacketErrorUnsupportedTimestampFormat", interfaceCounters.ReceivedPacketErrorUnsupportedTimestampFormat})
    interfaceCounters.EntityData.Leafs.Append("received-packet-error-timestamp-not-available", types.YLeaf{"ReceivedPacketErrorTimestampNotAvailable", interfaceCounters.ReceivedPacketErrorTimestampNotAvailable})
    interfaceCounters.EntityData.Leafs.Append("received-packet-error-unsupported-mandatory-tlv", types.YLeaf{"ReceivedPacketErrorUnsupportedMandatoryTlv", interfaceCounters.ReceivedPacketErrorUnsupportedMandatoryTlv})
    interfaceCounters.EntityData.Leafs.Append("received-packet-error-invalid-packet", types.YLeaf{"ReceivedPacketErrorInvalidPacket", interfaceCounters.ReceivedPacketErrorInvalidPacket})

    interfaceCounters.EntityData.YListKeys = []string {}

    return &(interfaceCounters.EntityData)
}

