// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs4k-freqsync package operational data.
// 
// This module contains definitions
// for the following management objects:
//   frequency-synchronization: Frequency Synchronization
//     operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs4k_freqsync_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs4k_freqsync_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs4k-freqsync-oper frequency-synchronization}", reflect.TypeOf(FrequencySynchronization{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs4k-freqsync-oper:frequency-synchronization", reflect.TypeOf(FrequencySynchronization{}))
}

// FsyncBagDampingState represents Damping state
type FsyncBagDampingState string

const (
    // Down
    FsyncBagDampingState_damping_state_down FsyncBagDampingState = "damping-state-down"

    // Coming up
    FsyncBagDampingState_damping_state_coming_up FsyncBagDampingState = "damping-state-coming-up"

    // Up
    FsyncBagDampingState_damping_state_up FsyncBagDampingState = "damping-state-up"

    // Going down
    FsyncBagDampingState_damping_state_going_down FsyncBagDampingState = "damping-state-going-down"
)

// FsyncBagStreamState represents Platform stream status
type FsyncBagStreamState string

const (
    // Invalid stream
    FsyncBagStreamState_stream_invalid FsyncBagStreamState = "stream-invalid"

    // Stream available
    FsyncBagStreamState_stream_available FsyncBagStreamState = "stream-available"

    // Stream locked
    FsyncBagStreamState_stream_locked FsyncBagStreamState = "stream-locked"

    // Stream in holdover
    FsyncBagStreamState_stream_holdover FsyncBagStreamState = "stream-holdover"

    // Stream free running
    FsyncBagStreamState_stream_freerun FsyncBagStreamState = "stream-freerun"

    // Stream failed
    FsyncBagStreamState_stream_failed FsyncBagStreamState = "stream-failed"

    // Unmonitored stream
    FsyncBagStreamState_stream_unmonitored FsyncBagStreamState = "stream-unmonitored"

    // Stream error
    FsyncBagStreamState_stream_error FsyncBagStreamState = "stream-error"
)

// FsyncBagEsmcPeerState represents ESMC peer state
type FsyncBagEsmcPeerState string

const (
    // Peer state down
    FsyncBagEsmcPeerState_peer_down FsyncBagEsmcPeerState = "peer-down"

    // Peer state up
    FsyncBagEsmcPeerState_peer_up FsyncBagEsmcPeerState = "peer-up"

    // Peer state timed out
    FsyncBagEsmcPeerState_peer_timed_out FsyncBagEsmcPeerState = "peer-timed-out"

    // Peer state unknown
    FsyncBagEsmcPeerState_peer_unknown FsyncBagEsmcPeerState = "peer-unknown"
)

// ImStateEnum represents Im state enum
type ImStateEnum string

const (
    // im state not ready
    ImStateEnum_im_state_not_ready ImStateEnum = "im-state-not-ready"

    // im state admin down
    ImStateEnum_im_state_admin_down ImStateEnum = "im-state-admin-down"

    // im state down
    ImStateEnum_im_state_down ImStateEnum = "im-state-down"

    // im state up
    ImStateEnum_im_state_up ImStateEnum = "im-state-up"

    // im state shutdown
    ImStateEnum_im_state_shutdown ImStateEnum = "im-state-shutdown"

    // im state err disable
    ImStateEnum_im_state_err_disable ImStateEnum = "im-state-err-disable"

    // im state down immediate
    ImStateEnum_im_state_down_immediate ImStateEnum = "im-state-down-immediate"

    // im state down immediate admin
    ImStateEnum_im_state_down_immediate_admin ImStateEnum = "im-state-down-immediate-admin"

    // im state down graceful
    ImStateEnum_im_state_down_graceful ImStateEnum = "im-state-down-graceful"

    // im state begin shutdown
    ImStateEnum_im_state_begin_shutdown ImStateEnum = "im-state-begin-shutdown"

    // im state end shutdown
    ImStateEnum_im_state_end_shutdown ImStateEnum = "im-state-end-shutdown"

    // im state begin error disable
    ImStateEnum_im_state_begin_error_disable ImStateEnum = "im-state-begin-error-disable"

    // im state end error disable
    ImStateEnum_im_state_end_error_disable ImStateEnum = "im-state-end-error-disable"

    // im state begin down graceful
    ImStateEnum_im_state_begin_down_graceful ImStateEnum = "im-state-begin-down-graceful"

    // im state reset
    ImStateEnum_im_state_reset ImStateEnum = "im-state-reset"

    // im state operational
    ImStateEnum_im_state_operational ImStateEnum = "im-state-operational"

    // im state not operational
    ImStateEnum_im_state_not_operational ImStateEnum = "im-state-not-operational"

    // im state unknown
    ImStateEnum_im_state_unknown ImStateEnum = "im-state-unknown"

    // im state last
    ImStateEnum_im_state_last ImStateEnum = "im-state-last"
)

// FsyncBagStreamInput represents Stream input type
type FsyncBagStreamInput string

const (
    // Invalid stream input
    FsyncBagStreamInput_invalid_input FsyncBagStreamInput = "invalid-input"

    // Source stream input
    FsyncBagStreamInput_source_input FsyncBagStreamInput = "source-input"

    // Selection point stream input
    FsyncBagStreamInput_selection_point_input FsyncBagStreamInput = "selection-point-input"
)

// FsyncBagQlO2G2Value represents Quality level option 2, generation 2 values
type FsyncBagQlO2G2Value string

const (
    // Invalid
    FsyncBagQlO2G2Value_option2_generation2_invalid FsyncBagQlO2G2Value = "option2-generation2-invalid"

    // Do not use
    FsyncBagQlO2G2Value_option2_generation2_do_not_use FsyncBagQlO2G2Value = "option2-generation2-do-not-use"

    // Failed
    FsyncBagQlO2G2Value_option2_generation2_failed FsyncBagQlO2G2Value = "option2-generation2-failed"

    // None
    FsyncBagQlO2G2Value_option2_generation2_none FsyncBagQlO2G2Value = "option2-generation2-none"

    // Primary reference source
    FsyncBagQlO2G2Value_option2_generation2prs FsyncBagQlO2G2Value = "option2-generation2prs"

    // Synchronized - traceability unknown
    FsyncBagQlO2G2Value_option2_generation2stu FsyncBagQlO2G2Value = "option2-generation2stu"

    // Stratum 2
    FsyncBagQlO2G2Value_option2_generation2_stratum2 FsyncBagQlO2G2Value = "option2-generation2-stratum2"

    // Stratum 3
    FsyncBagQlO2G2Value_option2_generation2_stratum3 FsyncBagQlO2G2Value = "option2-generation2-stratum3"

    // Transit node clock
    FsyncBagQlO2G2Value_option2_generation2tnc FsyncBagQlO2G2Value = "option2-generation2tnc"

    // Stratum 3E
    FsyncBagQlO2G2Value_option2_generation2_stratum3e FsyncBagQlO2G2Value = "option2-generation2-stratum3e"

    // SONET clock self timed
    FsyncBagQlO2G2Value_option2_generation2smc FsyncBagQlO2G2Value = "option2-generation2smc"

    // Stratum 4 freerun
    FsyncBagQlO2G2Value_option2_generation2_stratum4 FsyncBagQlO2G2Value = "option2-generation2-stratum4"
)

// FsyncBagQlOption represents Quality level option
type FsyncBagQlOption string

const (
    // No quality level option
    FsyncBagQlOption_no_quality_level_option FsyncBagQlOption = "no-quality-level-option"

    // ITU-T Quality level option 1
    FsyncBagQlOption_option1 FsyncBagQlOption = "option1"

    // ITU-T Quality level option 2, generation 1
    FsyncBagQlOption_option2_generation1 FsyncBagQlOption = "option2-generation1"

    // ITU-T Quality level option 2, generation 2
    FsyncBagQlOption_option2_generation2 FsyncBagQlOption = "option2-generation2"

    // Invalid quality level option
    FsyncBagQlOption_invalid_quality_level_option FsyncBagQlOption = "invalid-quality-level-option"
)

// FsyncBagSourceState represents Source state
type FsyncBagSourceState string

const (
    // Unknown
    FsyncBagSourceState_source_state_unknown FsyncBagSourceState = "source-state-unknown"

    // Up
    FsyncBagSourceState_source_state_up FsyncBagSourceState = "source-state-up"

    // Down
    FsyncBagSourceState_source_state_down FsyncBagSourceState = "source-state-down"
)

// FsyncBagForwardtraceNode represents Selection forwardtrace node information
type FsyncBagForwardtraceNode string

const (
    // A selection point forwardtrace node
    FsyncBagForwardtraceNode_forward_trace_node_selection_point FsyncBagForwardtraceNode = "forward-trace-node-selection-point"

    // A timing source forwardtrace node
    FsyncBagForwardtraceNode_forward_trace_node_source FsyncBagForwardtraceNode = "forward-trace-node-source"
)

// FsyncStream represents Fsync stream
type FsyncStream string

const (
    // Stream input from a local source
    FsyncStream_local FsyncStream = "local"

    // Stream input from a selection point on a remote
    // node
    FsyncStream_selection_point FsyncStream = "selection-point"
)

// FsyncBagClockIntfClass represents Clock-interface class
type FsyncBagClockIntfClass string

const (
    // BITS T1
    FsyncBagClockIntfClass_clock_class_bitst1 FsyncBagClockIntfClass = "clock-class-bitst1"

    // BITS E1
    FsyncBagClockIntfClass_clock_class_bitse1 FsyncBagClockIntfClass = "clock-class-bitse1"

    // BITS 2M
    FsyncBagClockIntfClass_clock_class_bits2m FsyncBagClockIntfClass = "clock-class-bits2m"

    // BITS 6M
    FsyncBagClockIntfClass_clock_class_bits6m FsyncBagClockIntfClass = "clock-class-bits6m"

    // BITS 64K
    FsyncBagClockIntfClass_clock_class_bits64k FsyncBagClockIntfClass = "clock-class-bits64k"

    // DTI
    FsyncBagClockIntfClass_clock_class_dti FsyncBagClockIntfClass = "clock-class-dti"

    // GPS
    FsyncBagClockIntfClass_clock_class_gps FsyncBagClockIntfClass = "clock-class-gps"

    // Inter-Chassis Sync
    FsyncBagClockIntfClass_clock_class_chassis_sync FsyncBagClockIntfClass = "clock-class-chassis-sync"

    // Bits J1
    FsyncBagClockIntfClass_clock_class_bitsj1 FsyncBagClockIntfClass = "clock-class-bitsj1"
)

// FsyncBagQlO1Value represents Quality level option 1 values
type FsyncBagQlO1Value string

const (
    // Invalid
    FsyncBagQlO1Value_option1_invalid FsyncBagQlO1Value = "option1-invalid"

    // Do not use
    FsyncBagQlO1Value_option1_do_not_use FsyncBagQlO1Value = "option1-do-not-use"

    // Failed
    FsyncBagQlO1Value_option1_failed FsyncBagQlO1Value = "option1-failed"

    // None
    FsyncBagQlO1Value_option1_none FsyncBagQlO1Value = "option1-none"

    // Primary reference clock
    FsyncBagQlO1Value_option1prc FsyncBagQlO1Value = "option1prc"

    // Type I or V slave clock
    FsyncBagQlO1Value_option1ssu_a FsyncBagQlO1Value = "option1ssu-a"

    // Type VI slave clock
    FsyncBagQlO1Value_option1ssu_b FsyncBagQlO1Value = "option1ssu-b"

    // SONET equipment clock
    FsyncBagQlO1Value_option1sec FsyncBagQlO1Value = "option1sec"
)

// FsyncBagSourceClass represents Source class
type FsyncBagSourceClass string

const (
    // Invalid source class
    FsyncBagSourceClass_invalid_source FsyncBagSourceClass = "invalid-source"

    // Ethernet interface
    FsyncBagSourceClass_ethernet_interface_source FsyncBagSourceClass = "ethernet-interface-source"

    // SONET interface
    FsyncBagSourceClass_sonet_interface_source FsyncBagSourceClass = "sonet-interface-source"

    // Clock interface
    FsyncBagSourceClass_clock_interface_source FsyncBagSourceClass = "clock-interface-source"

    // Internal clock
    FsyncBagSourceClass_internal_clock_source FsyncBagSourceClass = "internal-clock-source"

    // PTP clock
    FsyncBagSourceClass_ptp_source FsyncBagSourceClass = "ptp-source"

    // Satellite Access Interface
    FsyncBagSourceClass_satellite_access_interface_source FsyncBagSourceClass = "satellite-access-interface-source"

    // NTP clock
    FsyncBagSourceClass_ntp_source FsyncBagSourceClass = "ntp-source"
)

// FsyncBagQlO2G1Value represents Quality level option 2, generation 1 values
type FsyncBagQlO2G1Value string

const (
    // Invalid
    FsyncBagQlO2G1Value_option2_generation1_invalid FsyncBagQlO2G1Value = "option2-generation1-invalid"

    // Do not use
    FsyncBagQlO2G1Value_option2_generation1_do_not_use FsyncBagQlO2G1Value = "option2-generation1-do-not-use"

    // Failed
    FsyncBagQlO2G1Value_option2_generation1_failed FsyncBagQlO2G1Value = "option2-generation1-failed"

    // None
    FsyncBagQlO2G1Value_option2_generation1_none FsyncBagQlO2G1Value = "option2-generation1-none"

    // Primary reference source
    FsyncBagQlO2G1Value_option2_generation1prs FsyncBagQlO2G1Value = "option2-generation1prs"

    // Synchronized - traceability unknown
    FsyncBagQlO2G1Value_option2_generation1stu FsyncBagQlO2G1Value = "option2-generation1stu"

    // Stratum 2
    FsyncBagQlO2G1Value_option2_generation1_stratum2 FsyncBagQlO2G1Value = "option2-generation1-stratum2"

    // Stratum 3
    FsyncBagQlO2G1Value_option2_generation1_stratum3 FsyncBagQlO2G1Value = "option2-generation1-stratum3"

    // SONET clock self timed
    FsyncBagQlO2G1Value_option2_generation1smc FsyncBagQlO2G1Value = "option2-generation1smc"

    // Stratum 4 freerun
    FsyncBagQlO2G1Value_option2_generation1_stratum4 FsyncBagQlO2G1Value = "option2-generation1-stratum4"
)

// FsyncSource represents Fsync source
type FsyncSource string

const (
    // An ethernet interface
    FsyncSource_ethernet FsyncSource = "ethernet"

    // A SONET interface
    FsyncSource_sonet FsyncSource = "sonet"

    // A clock interface
    FsyncSource_clock FsyncSource = "clock"

    // An internal clock
    FsyncSource_internal FsyncSource = "internal"

    // A PTP clock
    FsyncSource_ptp FsyncSource = "ptp"

    // A satellite access interface clock
    FsyncSource_satellite_access FsyncSource = "satellite-access"

    // An NTP clock
    FsyncSource_ntp FsyncSource = "ntp"
)

// FrequencySynchronization
// Frequency Synchronization operational data
type FrequencySynchronization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table for node-specific operational data.
    Nodes FrequencySynchronization_Nodes

    // Summary operational data.
    Summary FrequencySynchronization_Summary

    // Table for global node-specific operational data.
    GlobalNodes FrequencySynchronization_GlobalNodes

    // Table for interface operational data.
    Interfaces FrequencySynchronization_Interfaces

    // Selection forwardtrace operational data for clock-interfaces.
    ClockInterfaceSelectionForwardTraces FrequencySynchronization_ClockInterfaceSelectionForwardTraces

    // Selection backtrace operational data for clock-interfaces.
    ClockInterfaceSelectionBackTraces FrequencySynchronization_ClockInterfaceSelectionBackTraces

    // Table for global interface operational data.
    GlobalInterfaces FrequencySynchronization_GlobalInterfaces

    // Table for clock operational data.
    Clocks FrequencySynchronization_Clocks
}

func (frequencySynchronization *FrequencySynchronization) GetFilter() yfilter.YFilter { return frequencySynchronization.YFilter }

func (frequencySynchronization *FrequencySynchronization) SetFilter(yf yfilter.YFilter) { frequencySynchronization.YFilter = yf }

func (frequencySynchronization *FrequencySynchronization) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "summary" { return "Summary" }
    if yname == "global-nodes" { return "GlobalNodes" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "clock-interface-selection-forward-traces" { return "ClockInterfaceSelectionForwardTraces" }
    if yname == "clock-interface-selection-back-traces" { return "ClockInterfaceSelectionBackTraces" }
    if yname == "global-interfaces" { return "GlobalInterfaces" }
    if yname == "clocks" { return "Clocks" }
    return ""
}

func (frequencySynchronization *FrequencySynchronization) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs4k-freqsync-oper:frequency-synchronization"
}

func (frequencySynchronization *FrequencySynchronization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &frequencySynchronization.Nodes
    }
    if childYangName == "summary" {
        return &frequencySynchronization.Summary
    }
    if childYangName == "global-nodes" {
        return &frequencySynchronization.GlobalNodes
    }
    if childYangName == "interfaces" {
        return &frequencySynchronization.Interfaces
    }
    if childYangName == "clock-interface-selection-forward-traces" {
        return &frequencySynchronization.ClockInterfaceSelectionForwardTraces
    }
    if childYangName == "clock-interface-selection-back-traces" {
        return &frequencySynchronization.ClockInterfaceSelectionBackTraces
    }
    if childYangName == "global-interfaces" {
        return &frequencySynchronization.GlobalInterfaces
    }
    if childYangName == "clocks" {
        return &frequencySynchronization.Clocks
    }
    return nil
}

func (frequencySynchronization *FrequencySynchronization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &frequencySynchronization.Nodes
    children["summary"] = &frequencySynchronization.Summary
    children["global-nodes"] = &frequencySynchronization.GlobalNodes
    children["interfaces"] = &frequencySynchronization.Interfaces
    children["clock-interface-selection-forward-traces"] = &frequencySynchronization.ClockInterfaceSelectionForwardTraces
    children["clock-interface-selection-back-traces"] = &frequencySynchronization.ClockInterfaceSelectionBackTraces
    children["global-interfaces"] = &frequencySynchronization.GlobalInterfaces
    children["clocks"] = &frequencySynchronization.Clocks
    return children
}

func (frequencySynchronization *FrequencySynchronization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (frequencySynchronization *FrequencySynchronization) GetBundleName() string { return "cisco_ios_xr" }

func (frequencySynchronization *FrequencySynchronization) GetYangName() string { return "frequency-synchronization" }

func (frequencySynchronization *FrequencySynchronization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frequencySynchronization *FrequencySynchronization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frequencySynchronization *FrequencySynchronization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frequencySynchronization *FrequencySynchronization) SetParent(parent types.Entity) { frequencySynchronization.parent = parent }

func (frequencySynchronization *FrequencySynchronization) GetParent() types.Entity { return frequencySynchronization.parent }

func (frequencySynchronization *FrequencySynchronization) GetParentYangName() string { return "Cisco-IOS-XR-ncs4k-freqsync-oper" }

// FrequencySynchronization_Nodes
// Table for node-specific operational data
type FrequencySynchronization_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // FrequencySynchronization_Nodes_Node.
    Node []FrequencySynchronization_Nodes_Node
}

func (nodes *FrequencySynchronization_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *FrequencySynchronization_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *FrequencySynchronization_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *FrequencySynchronization_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *FrequencySynchronization_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *FrequencySynchronization_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *FrequencySynchronization_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *FrequencySynchronization_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *FrequencySynchronization_Nodes) GetYangName() string { return "nodes" }

func (nodes *FrequencySynchronization_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *FrequencySynchronization_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *FrequencySynchronization_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *FrequencySynchronization_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *FrequencySynchronization_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *FrequencySynchronization_Nodes) GetParentYangName() string { return "frequency-synchronization" }

// FrequencySynchronization_Nodes_Node
// Node-specific data for a particular node
type FrequencySynchronization_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Selection point table.
    SelectionPoints FrequencySynchronization_Nodes_Node_SelectionPoints

    // PTP operational data.
    Ptp FrequencySynchronization_Nodes_Node_Ptp

    // Table for selection point input operational data.
    SelectionPointInputs FrequencySynchronization_Nodes_Node_SelectionPointInputs

    // NTP operational data.
    Ntp FrequencySynchronization_Nodes_Node_Ntp

    // Configuration error operational data.
    ConfigurationErrors FrequencySynchronization_Nodes_Node_ConfigurationErrors

    // SSM operational data.
    SsmSummary FrequencySynchronization_Nodes_Node_SsmSummary
}

func (node *FrequencySynchronization_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *FrequencySynchronization_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *FrequencySynchronization_Nodes_Node) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "selection-points" { return "SelectionPoints" }
    if yname == "ptp" { return "Ptp" }
    if yname == "selection-point-inputs" { return "SelectionPointInputs" }
    if yname == "ntp" { return "Ntp" }
    if yname == "configuration-errors" { return "ConfigurationErrors" }
    if yname == "ssm-summary" { return "SsmSummary" }
    return ""
}

func (node *FrequencySynchronization_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
}

func (node *FrequencySynchronization_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selection-points" {
        return &node.SelectionPoints
    }
    if childYangName == "ptp" {
        return &node.Ptp
    }
    if childYangName == "selection-point-inputs" {
        return &node.SelectionPointInputs
    }
    if childYangName == "ntp" {
        return &node.Ntp
    }
    if childYangName == "configuration-errors" {
        return &node.ConfigurationErrors
    }
    if childYangName == "ssm-summary" {
        return &node.SsmSummary
    }
    return nil
}

func (node *FrequencySynchronization_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["selection-points"] = &node.SelectionPoints
    children["ptp"] = &node.Ptp
    children["selection-point-inputs"] = &node.SelectionPointInputs
    children["ntp"] = &node.Ntp
    children["configuration-errors"] = &node.ConfigurationErrors
    children["ssm-summary"] = &node.SsmSummary
    return children
}

func (node *FrequencySynchronization_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = node.Node
    return leafs
}

func (node *FrequencySynchronization_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *FrequencySynchronization_Nodes_Node) GetYangName() string { return "node" }

func (node *FrequencySynchronization_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *FrequencySynchronization_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *FrequencySynchronization_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *FrequencySynchronization_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *FrequencySynchronization_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *FrequencySynchronization_Nodes_Node) GetParentYangName() string { return "nodes" }

// FrequencySynchronization_Nodes_Node_SelectionPoints
// Selection point table
type FrequencySynchronization_Nodes_Node_SelectionPoints struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for a given selection point. The type is slice of
    // FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint.
    SelectionPoint []FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint
}

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetFilter() yfilter.YFilter { return selectionPoints.YFilter }

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) SetFilter(yf yfilter.YFilter) { selectionPoints.YFilter = yf }

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetGoName(yname string) string {
    if yname == "selection-point" { return "SelectionPoint" }
    return ""
}

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetSegmentPath() string {
    return "selection-points"
}

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selection-point" {
        for _, c := range selectionPoints.SelectionPoint {
            if selectionPoints.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint{}
        selectionPoints.SelectionPoint = append(selectionPoints.SelectionPoint, child)
        return &selectionPoints.SelectionPoint[len(selectionPoints.SelectionPoint)-1]
    }
    return nil
}

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range selectionPoints.SelectionPoint {
        children[selectionPoints.SelectionPoint[i].GetSegmentPath()] = &selectionPoints.SelectionPoint[i]
    }
    return children
}

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetYangName() string { return "selection-points" }

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) SetParent(parent types.Entity) { selectionPoints.parent = parent }

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetParent() types.Entity { return selectionPoints.parent }

func (selectionPoints *FrequencySynchronization_Nodes_Node_SelectionPoints) GetParentYangName() string { return "node" }

// FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint
// Operational data for a given selection point
type FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Selection point ID. The type is interface{} with
    // range: -2147483648..2147483647.
    SelectionPoint interface{}

    // Selection Point Type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Description. The type is string.
    Description interface{}

    // Number of inputs. The type is interface{} with range: 0..4294967295.
    Inputs interface{}

    // Number of inputs that are selected. The type is interface{} with range:
    // 0..4294967295.
    InputsSelected interface{}

    // The selection point is a time-of-day selection point. The type is bool.
    TimeOfDaySelection interface{}

    // Information about the output of the selection point.
    Output FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output

    // Time the SP was last programmed.
    LastProgrammed FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed

    // Time the last selection was made.
    LastSelection FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetFilter() yfilter.YFilter { return selectionPoint.YFilter }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) SetFilter(yf yfilter.YFilter) { selectionPoint.YFilter = yf }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point" { return "SelectionPoint" }
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "description" { return "Description" }
    if yname == "inputs" { return "Inputs" }
    if yname == "inputs-selected" { return "InputsSelected" }
    if yname == "time-of-day-selection" { return "TimeOfDaySelection" }
    if yname == "output" { return "Output" }
    if yname == "last-programmed" { return "LastProgrammed" }
    if yname == "last-selection" { return "LastSelection" }
    return ""
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetSegmentPath() string {
    return "selection-point" + "[selection-point='" + fmt.Sprintf("%v", selectionPoint.SelectionPoint) + "']"
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output" {
        return &selectionPoint.Output
    }
    if childYangName == "last-programmed" {
        return &selectionPoint.LastProgrammed
    }
    if childYangName == "last-selection" {
        return &selectionPoint.LastSelection
    }
    return nil
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output"] = &selectionPoint.Output
    children["last-programmed"] = &selectionPoint.LastProgrammed
    children["last-selection"] = &selectionPoint.LastSelection
    return children
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point"] = selectionPoint.SelectionPoint
    leafs["selection-point-type"] = selectionPoint.SelectionPointType
    leafs["description"] = selectionPoint.Description
    leafs["inputs"] = selectionPoint.Inputs
    leafs["inputs-selected"] = selectionPoint.InputsSelected
    leafs["time-of-day-selection"] = selectionPoint.TimeOfDaySelection
    return leafs
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetYangName() string { return "selection-point" }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) SetParent(parent types.Entity) { selectionPoint.parent = parent }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetParent() types.Entity { return selectionPoint.parent }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint) GetParentYangName() string { return "selection-points" }

// FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output
// Information about the output of the selection
// point
type FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Used for local clock output. The type is bool.
    LocalClockOuput interface{}

    // Used for local line interface output. The type is bool.
    LocalLineOutput interface{}

    // Used for local time-of-day output. The type is bool.
    LocalTimeOfDayOutput interface{}

    // SPA selection points. The type is slice of interface{} with range: 0..255.
    SpaSelectionPoint []interface{}

    // SPA selection points descrption. The type is slice of string.
    SpaSelectionPointsDescription []interface{}

    // Node selection points. The type is slice of interface{} with range: 0..255.
    NodeSelectionPoint []interface{}

    // Node selection points descrption. The type is slice of string.
    NodeSelectionPointsDescription []interface{}

    // Chassis selection points. The type is slice of interface{} with range:
    // 0..255.
    ChassisSelectionPoint []interface{}

    // Chassis selection points descrption. The type is slice of string.
    ChassisSelectionPointsDescription []interface{}

    // Router selection points. The type is slice of interface{} with range:
    // 0..255.
    RouterSelectionPoint []interface{}

    // Router selection points descrption. The type is slice of string.
    RouterSelectionPointsDescription []interface{}
}

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetGoName(yname string) string {
    if yname == "local-clock-ouput" { return "LocalClockOuput" }
    if yname == "local-line-output" { return "LocalLineOutput" }
    if yname == "local-time-of-day-output" { return "LocalTimeOfDayOutput" }
    if yname == "spa-selection-point" { return "SpaSelectionPoint" }
    if yname == "spa-selection-points-description" { return "SpaSelectionPointsDescription" }
    if yname == "node-selection-point" { return "NodeSelectionPoint" }
    if yname == "node-selection-points-description" { return "NodeSelectionPointsDescription" }
    if yname == "chassis-selection-point" { return "ChassisSelectionPoint" }
    if yname == "chassis-selection-points-description" { return "ChassisSelectionPointsDescription" }
    if yname == "router-selection-point" { return "RouterSelectionPoint" }
    if yname == "router-selection-points-description" { return "RouterSelectionPointsDescription" }
    return ""
}

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetSegmentPath() string {
    return "output"
}

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-clock-ouput"] = output.LocalClockOuput
    leafs["local-line-output"] = output.LocalLineOutput
    leafs["local-time-of-day-output"] = output.LocalTimeOfDayOutput
    leafs["spa-selection-point"] = output.SpaSelectionPoint
    leafs["spa-selection-points-description"] = output.SpaSelectionPointsDescription
    leafs["node-selection-point"] = output.NodeSelectionPoint
    leafs["node-selection-points-description"] = output.NodeSelectionPointsDescription
    leafs["chassis-selection-point"] = output.ChassisSelectionPoint
    leafs["chassis-selection-points-description"] = output.ChassisSelectionPointsDescription
    leafs["router-selection-point"] = output.RouterSelectionPoint
    leafs["router-selection-points-description"] = output.RouterSelectionPointsDescription
    return leafs
}

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetYangName() string { return "output" }

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetParent() types.Entity { return output.parent }

func (output *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_Output) GetParentYangName() string { return "selection-point" }

// FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed
// Time the SP was last programmed
type FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetFilter() yfilter.YFilter { return lastProgrammed.YFilter }

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) SetFilter(yf yfilter.YFilter) { lastProgrammed.YFilter = yf }

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetSegmentPath() string {
    return "last-programmed"
}

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = lastProgrammed.Seconds
    leafs["nanoseconds"] = lastProgrammed.Nanoseconds
    return leafs
}

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetBundleName() string { return "cisco_ios_xr" }

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetYangName() string { return "last-programmed" }

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) SetParent(parent types.Entity) { lastProgrammed.parent = parent }

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetParent() types.Entity { return lastProgrammed.parent }

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastProgrammed) GetParentYangName() string { return "selection-point" }

// FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection
// Time the last selection was made
type FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetFilter() yfilter.YFilter { return lastSelection.YFilter }

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) SetFilter(yf yfilter.YFilter) { lastSelection.YFilter = yf }

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetSegmentPath() string {
    return "last-selection"
}

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = lastSelection.Seconds
    leafs["nanoseconds"] = lastSelection.Nanoseconds
    return leafs
}

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetBundleName() string { return "cisco_ios_xr" }

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetYangName() string { return "last-selection" }

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) SetParent(parent types.Entity) { lastSelection.parent = parent }

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetParent() types.Entity { return lastSelection.parent }

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPoints_SelectionPoint_LastSelection) GetParentYangName() string { return "selection-point" }

// FrequencySynchronization_Nodes_Node_Ptp
// PTP operational data
type FrequencySynchronization_Nodes_Node_Ptp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PTP state. The type is FsyncBagSourceState.
    State interface{}

    // The PTP clock supports frequency. The type is bool.
    SupportsFrequency interface{}

    // The PTP clock supports time. The type is bool.
    SupportsTimeOfDay interface{}

    // The priority of the PTP clock when selected between frequency sources. The
    // type is interface{} with range: 0..255.
    FrequencyPriority interface{}

    // The priority of the PTP clock when selecting between time-of-day sources.
    // The type is interface{} with range: 0..255.
    TimeOfDayPriority interface{}

    // Spa selection points. The type is slice of interface{} with range: 0..255.
    SpaSelectionPoint []interface{}

    // Spa selection points descrption. The type is slice of string.
    SpaSelectionPointsDescription []interface{}

    // Node selection points. The type is slice of interface{} with range: 0..255.
    NodeSelectionPoint []interface{}

    // Node selection points descrption. The type is slice of string.
    NodeSelectionPointsDescription []interface{}

    // The effective input quality level.
    QualityLevelEffectiveInput FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput
}

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetFilter() yfilter.YFilter { return ptp.YFilter }

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) SetFilter(yf yfilter.YFilter) { ptp.YFilter = yf }

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    if yname == "supports-frequency" { return "SupportsFrequency" }
    if yname == "supports-time-of-day" { return "SupportsTimeOfDay" }
    if yname == "frequency-priority" { return "FrequencyPriority" }
    if yname == "time-of-day-priority" { return "TimeOfDayPriority" }
    if yname == "spa-selection-point" { return "SpaSelectionPoint" }
    if yname == "spa-selection-points-description" { return "SpaSelectionPointsDescription" }
    if yname == "node-selection-point" { return "NodeSelectionPoint" }
    if yname == "node-selection-points-description" { return "NodeSelectionPointsDescription" }
    if yname == "quality-level-effective-input" { return "QualityLevelEffectiveInput" }
    return ""
}

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetSegmentPath() string {
    return "ptp"
}

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "quality-level-effective-input" {
        return &ptp.QualityLevelEffectiveInput
    }
    return nil
}

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["quality-level-effective-input"] = &ptp.QualityLevelEffectiveInput
    return children
}

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state"] = ptp.State
    leafs["supports-frequency"] = ptp.SupportsFrequency
    leafs["supports-time-of-day"] = ptp.SupportsTimeOfDay
    leafs["frequency-priority"] = ptp.FrequencyPriority
    leafs["time-of-day-priority"] = ptp.TimeOfDayPriority
    leafs["spa-selection-point"] = ptp.SpaSelectionPoint
    leafs["spa-selection-points-description"] = ptp.SpaSelectionPointsDescription
    leafs["node-selection-point"] = ptp.NodeSelectionPoint
    leafs["node-selection-points-description"] = ptp.NodeSelectionPointsDescription
    return leafs
}

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetBundleName() string { return "cisco_ios_xr" }

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetYangName() string { return "ptp" }

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) SetParent(parent types.Entity) { ptp.parent = parent }

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetParent() types.Entity { return ptp.parent }

func (ptp *FrequencySynchronization_Nodes_Node_Ptp) GetParentYangName() string { return "node" }

// FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput
// The effective input quality level
type FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetFilter() yfilter.YFilter { return qualityLevelEffectiveInput.YFilter }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) SetFilter(yf yfilter.YFilter) { qualityLevelEffectiveInput.YFilter = yf }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetSegmentPath() string {
    return "quality-level-effective-input"
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelEffectiveInput.QualityLevelOption
    leafs["option1-value"] = qualityLevelEffectiveInput.Option1Value
    leafs["option2-generation1-value"] = qualityLevelEffectiveInput.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelEffectiveInput.Option2Generation2Value
    return leafs
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetYangName() string { return "quality-level-effective-input" }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) SetParent(parent types.Entity) { qualityLevelEffectiveInput.parent = parent }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetParent() types.Entity { return qualityLevelEffectiveInput.parent }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ptp_QualityLevelEffectiveInput) GetParentYangName() string { return "ptp" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs
// Table for selection point input operational
// data
type FrequencySynchronization_Nodes_Node_SelectionPointInputs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for a particular selection point input. The type is slice
    // of
    // FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput.
    SelectionPointInput []FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput
}

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetFilter() yfilter.YFilter { return selectionPointInputs.YFilter }

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) SetFilter(yf yfilter.YFilter) { selectionPointInputs.YFilter = yf }

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetGoName(yname string) string {
    if yname == "selection-point-input" { return "SelectionPointInput" }
    return ""
}

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetSegmentPath() string {
    return "selection-point-inputs"
}

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selection-point-input" {
        for _, c := range selectionPointInputs.SelectionPointInput {
            if selectionPointInputs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput{}
        selectionPointInputs.SelectionPointInput = append(selectionPointInputs.SelectionPointInput, child)
        return &selectionPointInputs.SelectionPointInput[len(selectionPointInputs.SelectionPointInput)-1]
    }
    return nil
}

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range selectionPointInputs.SelectionPointInput {
        children[selectionPointInputs.SelectionPointInput[i].GetSegmentPath()] = &selectionPointInputs.SelectionPointInput[i]
    }
    return children
}

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetYangName() string { return "selection-point-inputs" }

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) SetParent(parent types.Entity) { selectionPointInputs.parent = parent }

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetParent() types.Entity { return selectionPointInputs.parent }

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetParentYangName() string { return "node" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput
// Operational data for a particular selection
// point input
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SelectionPoint interface{}

    // Type of stream. The type is FsyncStream.
    StreamType interface{}

    // Type of source. The type is FsyncSource.
    SourceType interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Clock port. The type is interface{} with range: -2147483648..2147483647.
    ClockPort interface{}

    // Last node for a selection point stream. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    LastNode interface{}

    // Last selection point for a selection point stream. The type is interface{}
    // with range: -2147483648..2147483647.
    LastSelectionPoint interface{}

    // Output ID for a selection point stream. The type is interface{} with range:
    // -2147483648..2147483647.
    OutputId interface{}

    // The selection point input supports frequency. The type is bool.
    SupportsFrequency interface{}

    // The selection point input supports time-of-day. The type is bool.
    SupportsTimeOfDay interface{}

    // Priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Time-of-day priority. The type is interface{} with range: 0..255.
    TimeOfDayPriority interface{}

    // The selection point input is selected. The type is bool.
    Selected interface{}

    // Platform output ID, if the input is selected. The type is interface{} with
    // range: 0..255.
    OutputIdXr interface{}

    // Platform status. The type is FsyncBagStreamState.
    PlatformStatus interface{}

    // Why the stream has failed. The type is string.
    PlatformFailedReason interface{}

    // The selection point the input is for.
    InputSelectionPoint FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint

    // Stream.
    Stream FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream

    // Original source.
    OriginalSource FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource

    // Quality Level.
    QualityLevel FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel
}

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetFilter() yfilter.YFilter { return selectionPointInput.YFilter }

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) SetFilter(yf yfilter.YFilter) { selectionPointInput.YFilter = yf }

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetGoName(yname string) string {
    if yname == "selection-point" { return "SelectionPoint" }
    if yname == "stream-type" { return "StreamType" }
    if yname == "source-type" { return "SourceType" }
    if yname == "interface" { return "Interface" }
    if yname == "clock-port" { return "ClockPort" }
    if yname == "last-node" { return "LastNode" }
    if yname == "last-selection-point" { return "LastSelectionPoint" }
    if yname == "output-id" { return "OutputId" }
    if yname == "supports-frequency" { return "SupportsFrequency" }
    if yname == "supports-time-of-day" { return "SupportsTimeOfDay" }
    if yname == "priority" { return "Priority" }
    if yname == "time-of-day-priority" { return "TimeOfDayPriority" }
    if yname == "selected" { return "Selected" }
    if yname == "output-id-xr" { return "OutputIdXr" }
    if yname == "platform-status" { return "PlatformStatus" }
    if yname == "platform-failed-reason" { return "PlatformFailedReason" }
    if yname == "input-selection-point" { return "InputSelectionPoint" }
    if yname == "stream" { return "Stream" }
    if yname == "original-source" { return "OriginalSource" }
    if yname == "quality-level" { return "QualityLevel" }
    return ""
}

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetSegmentPath() string {
    return "selection-point-input"
}

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input-selection-point" {
        return &selectionPointInput.InputSelectionPoint
    }
    if childYangName == "stream" {
        return &selectionPointInput.Stream
    }
    if childYangName == "original-source" {
        return &selectionPointInput.OriginalSource
    }
    if childYangName == "quality-level" {
        return &selectionPointInput.QualityLevel
    }
    return nil
}

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input-selection-point"] = &selectionPointInput.InputSelectionPoint
    children["stream"] = &selectionPointInput.Stream
    children["original-source"] = &selectionPointInput.OriginalSource
    children["quality-level"] = &selectionPointInput.QualityLevel
    return children
}

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point"] = selectionPointInput.SelectionPoint
    leafs["stream-type"] = selectionPointInput.StreamType
    leafs["source-type"] = selectionPointInput.SourceType
    leafs["interface"] = selectionPointInput.Interface
    leafs["clock-port"] = selectionPointInput.ClockPort
    leafs["last-node"] = selectionPointInput.LastNode
    leafs["last-selection-point"] = selectionPointInput.LastSelectionPoint
    leafs["output-id"] = selectionPointInput.OutputId
    leafs["supports-frequency"] = selectionPointInput.SupportsFrequency
    leafs["supports-time-of-day"] = selectionPointInput.SupportsTimeOfDay
    leafs["priority"] = selectionPointInput.Priority
    leafs["time-of-day-priority"] = selectionPointInput.TimeOfDayPriority
    leafs["selected"] = selectionPointInput.Selected
    leafs["output-id-xr"] = selectionPointInput.OutputIdXr
    leafs["platform-status"] = selectionPointInput.PlatformStatus
    leafs["platform-failed-reason"] = selectionPointInput.PlatformFailedReason
    return leafs
}

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetYangName() string { return "selection-point-input" }

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) SetParent(parent types.Entity) { selectionPointInput.parent = parent }

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetParent() types.Entity { return selectionPointInput.parent }

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetParentYangName() string { return "selection-point-inputs" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint
// The selection point the input is for
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}
}

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetFilter() yfilter.YFilter { return inputSelectionPoint.YFilter }

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) SetFilter(yf yfilter.YFilter) { inputSelectionPoint.YFilter = yf }

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "selection-point-description" { return "SelectionPointDescription" }
    if yname == "node" { return "Node" }
    return ""
}

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetSegmentPath() string {
    return "input-selection-point"
}

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point-type"] = inputSelectionPoint.SelectionPointType
    leafs["selection-point-description"] = inputSelectionPoint.SelectionPointDescription
    leafs["node"] = inputSelectionPoint.Node
    return leafs
}

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetYangName() string { return "input-selection-point" }

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) SetParent(parent types.Entity) { inputSelectionPoint.parent = parent }

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetParent() types.Entity { return inputSelectionPoint.parent }

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetParentYangName() string { return "selection-point-input" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream
// Stream
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // StreamInput. The type is FsyncBagStreamInput.
    StreamInput interface{}

    // Source ID.
    SourceId FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId

    // Selection point ID.
    SelectionPointId FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId
}

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetFilter() yfilter.YFilter { return stream.YFilter }

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) SetFilter(yf yfilter.YFilter) { stream.YFilter = yf }

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetGoName(yname string) string {
    if yname == "stream-input" { return "StreamInput" }
    if yname == "source-id" { return "SourceId" }
    if yname == "selection-point-id" { return "SelectionPointId" }
    return ""
}

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetSegmentPath() string {
    return "stream"
}

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source-id" {
        return &stream.SourceId
    }
    if childYangName == "selection-point-id" {
        return &stream.SelectionPointId
    }
    return nil
}

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source-id"] = &stream.SourceId
    children["selection-point-id"] = &stream.SelectionPointId
    return children
}

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stream-input"] = stream.StreamInput
    return leafs
}

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetBundleName() string { return "cisco_ios_xr" }

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetYangName() string { return "stream" }

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) SetParent(parent types.Entity) { stream.parent = parent }

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetParent() types.Entity { return stream.parent }

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetParentYangName() string { return "selection-point-input" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId
// Source ID
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId
}

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetFilter() yfilter.YFilter { return sourceId.YFilter }

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) SetFilter(yf yfilter.YFilter) { sourceId.YFilter = yf }

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetSegmentPath() string {
    return "source-id"
}

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &sourceId.ClockId
    }
    return nil
}

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &sourceId.ClockId
    return children
}

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = sourceId.SourceClass
    leafs["ethernet-interface"] = sourceId.EthernetInterface
    leafs["sonet-interface"] = sourceId.SonetInterface
    leafs["node"] = sourceId.Node
    leafs["ptp-node"] = sourceId.PtpNode
    leafs["satellite-access-interface"] = sourceId.SatelliteAccessInterface
    leafs["ntp-node"] = sourceId.NtpNode
    return leafs
}

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetBundleName() string { return "cisco_ios_xr" }

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetYangName() string { return "source-id" }

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) SetParent(parent types.Entity) { sourceId.parent = parent }

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetParent() types.Entity { return sourceId.parent }

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetParentYangName() string { return "stream" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId
// Clock ID
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetParentYangName() string { return "source-id" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId
// Selection point ID
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Output ID. The type is interface{} with range: 0..255.
    OutputId interface{}

    // Last selection point.
    SelectionPoint FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint
}

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetFilter() yfilter.YFilter { return selectionPointId.YFilter }

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) SetFilter(yf yfilter.YFilter) { selectionPointId.YFilter = yf }

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetGoName(yname string) string {
    if yname == "output-id" { return "OutputId" }
    if yname == "selection-point" { return "SelectionPoint" }
    return ""
}

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetSegmentPath() string {
    return "selection-point-id"
}

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selection-point" {
        return &selectionPointId.SelectionPoint
    }
    return nil
}

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["selection-point"] = &selectionPointId.SelectionPoint
    return children
}

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["output-id"] = selectionPointId.OutputId
    return leafs
}

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetYangName() string { return "selection-point-id" }

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) SetParent(parent types.Entity) { selectionPointId.parent = parent }

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetParent() types.Entity { return selectionPointId.parent }

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetParentYangName() string { return "stream" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint
// Last selection point
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetFilter() yfilter.YFilter { return selectionPoint.YFilter }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) SetFilter(yf yfilter.YFilter) { selectionPoint.YFilter = yf }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "selection-point-description" { return "SelectionPointDescription" }
    if yname == "node" { return "Node" }
    return ""
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetSegmentPath() string {
    return "selection-point"
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point-type"] = selectionPoint.SelectionPointType
    leafs["selection-point-description"] = selectionPoint.SelectionPointDescription
    leafs["node"] = selectionPoint.Node
    return leafs
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetYangName() string { return "selection-point" }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) SetParent(parent types.Entity) { selectionPoint.parent = parent }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetParent() types.Entity { return selectionPoint.parent }

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetParentYangName() string { return "selection-point-id" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource
// Original source
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId
}

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetFilter() yfilter.YFilter { return originalSource.YFilter }

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) SetFilter(yf yfilter.YFilter) { originalSource.YFilter = yf }

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetSegmentPath() string {
    return "original-source"
}

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &originalSource.ClockId
    }
    return nil
}

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &originalSource.ClockId
    return children
}

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = originalSource.SourceClass
    leafs["ethernet-interface"] = originalSource.EthernetInterface
    leafs["sonet-interface"] = originalSource.SonetInterface
    leafs["node"] = originalSource.Node
    leafs["ptp-node"] = originalSource.PtpNode
    leafs["satellite-access-interface"] = originalSource.SatelliteAccessInterface
    leafs["ntp-node"] = originalSource.NtpNode
    return leafs
}

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetBundleName() string { return "cisco_ios_xr" }

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetYangName() string { return "original-source" }

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) SetParent(parent types.Entity) { originalSource.parent = parent }

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetParent() types.Entity { return originalSource.parent }

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetParentYangName() string { return "selection-point-input" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId
// Clock ID
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetParentYangName() string { return "original-source" }

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel
// Quality Level
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetFilter() yfilter.YFilter { return qualityLevel.YFilter }

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) SetFilter(yf yfilter.YFilter) { qualityLevel.YFilter = yf }

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetSegmentPath() string {
    return "quality-level"
}

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevel.QualityLevelOption
    leafs["option1-value"] = qualityLevel.Option1Value
    leafs["option2-generation1-value"] = qualityLevel.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevel.Option2Generation2Value
    return leafs
}

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetYangName() string { return "quality-level" }

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) SetParent(parent types.Entity) { qualityLevel.parent = parent }

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetParent() types.Entity { return qualityLevel.parent }

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetParentYangName() string { return "selection-point-input" }

// FrequencySynchronization_Nodes_Node_Ntp
// NTP operational data
type FrequencySynchronization_Nodes_Node_Ntp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NTP state. The type is FsyncBagSourceState.
    State interface{}

    // The NTP clock supports frequency. The type is bool.
    SupportsFrequency interface{}

    // The NTP clock supports time. The type is bool.
    SupportsTimeOfDay interface{}

    // The priority of the NTP clock when selected between frequency sources. The
    // type is interface{} with range: 0..255.
    FrequencyPriority interface{}

    // The priority of the NTP clock when selecting between time-of-day sources.
    // The type is interface{} with range: 0..255.
    TimeOfDayPriority interface{}

    // Spa selection points. The type is slice of interface{} with range: 0..255.
    SpaSelectionPoint []interface{}

    // Spa selection points descrption. The type is slice of string.
    SpaSelectionPointsDescription []interface{}

    // Node selection points. The type is slice of interface{} with range: 0..255.
    NodeSelectionPoint []interface{}

    // Node selection points descrption. The type is slice of string.
    NodeSelectionPointsDescription []interface{}

    // The effective input quality level.
    QualityLevelEffectiveInput FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput
}

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetFilter() yfilter.YFilter { return ntp.YFilter }

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) SetFilter(yf yfilter.YFilter) { ntp.YFilter = yf }

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    if yname == "supports-frequency" { return "SupportsFrequency" }
    if yname == "supports-time-of-day" { return "SupportsTimeOfDay" }
    if yname == "frequency-priority" { return "FrequencyPriority" }
    if yname == "time-of-day-priority" { return "TimeOfDayPriority" }
    if yname == "spa-selection-point" { return "SpaSelectionPoint" }
    if yname == "spa-selection-points-description" { return "SpaSelectionPointsDescription" }
    if yname == "node-selection-point" { return "NodeSelectionPoint" }
    if yname == "node-selection-points-description" { return "NodeSelectionPointsDescription" }
    if yname == "quality-level-effective-input" { return "QualityLevelEffectiveInput" }
    return ""
}

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetSegmentPath() string {
    return "ntp"
}

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "quality-level-effective-input" {
        return &ntp.QualityLevelEffectiveInput
    }
    return nil
}

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["quality-level-effective-input"] = &ntp.QualityLevelEffectiveInput
    return children
}

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state"] = ntp.State
    leafs["supports-frequency"] = ntp.SupportsFrequency
    leafs["supports-time-of-day"] = ntp.SupportsTimeOfDay
    leafs["frequency-priority"] = ntp.FrequencyPriority
    leafs["time-of-day-priority"] = ntp.TimeOfDayPriority
    leafs["spa-selection-point"] = ntp.SpaSelectionPoint
    leafs["spa-selection-points-description"] = ntp.SpaSelectionPointsDescription
    leafs["node-selection-point"] = ntp.NodeSelectionPoint
    leafs["node-selection-points-description"] = ntp.NodeSelectionPointsDescription
    return leafs
}

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetBundleName() string { return "cisco_ios_xr" }

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetYangName() string { return "ntp" }

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) SetParent(parent types.Entity) { ntp.parent = parent }

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetParent() types.Entity { return ntp.parent }

func (ntp *FrequencySynchronization_Nodes_Node_Ntp) GetParentYangName() string { return "node" }

// FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput
// The effective input quality level
type FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetFilter() yfilter.YFilter { return qualityLevelEffectiveInput.YFilter }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) SetFilter(yf yfilter.YFilter) { qualityLevelEffectiveInput.YFilter = yf }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetSegmentPath() string {
    return "quality-level-effective-input"
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelEffectiveInput.QualityLevelOption
    leafs["option1-value"] = qualityLevelEffectiveInput.Option1Value
    leafs["option2-generation1-value"] = qualityLevelEffectiveInput.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelEffectiveInput.Option2Generation2Value
    return leafs
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetYangName() string { return "quality-level-effective-input" }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) SetParent(parent types.Entity) { qualityLevelEffectiveInput.parent = parent }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetParent() types.Entity { return qualityLevelEffectiveInput.parent }

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_Ntp_QualityLevelEffectiveInput) GetParentYangName() string { return "ntp" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors
// Configuration error operational data
type FrequencySynchronization_Nodes_Node_ConfigurationErrors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration errors. The type is slice of
    // FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource.
    ErrorSource []FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource
}

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetFilter() yfilter.YFilter { return configurationErrors.YFilter }

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) SetFilter(yf yfilter.YFilter) { configurationErrors.YFilter = yf }

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetGoName(yname string) string {
    if yname == "error-source" { return "ErrorSource" }
    return ""
}

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetSegmentPath() string {
    return "configuration-errors"
}

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error-source" {
        for _, c := range configurationErrors.ErrorSource {
            if configurationErrors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource{}
        configurationErrors.ErrorSource = append(configurationErrors.ErrorSource, child)
        return &configurationErrors.ErrorSource[len(configurationErrors.ErrorSource)-1]
    }
    return nil
}

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range configurationErrors.ErrorSource {
        children[configurationErrors.ErrorSource[i].GetSegmentPath()] = &configurationErrors.ErrorSource[i]
    }
    return children
}

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetBundleName() string { return "cisco_ios_xr" }

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetYangName() string { return "configuration-errors" }

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) SetParent(parent types.Entity) { configurationErrors.parent = parent }

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetParent() types.Entity { return configurationErrors.parent }

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetParentYangName() string { return "node" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource
// Configuration errors
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Frequency Synchronization enable error. The type is bool.
    EnableError interface{}

    // Minimum input QL config error. The type is bool.
    InputMinError interface{}

    // Exact input QL config error. The type is bool.
    InputExactError interface{}

    // Maximum input Ql config error. The type is bool.
    InputMaxError interface{}

    // Minimum output QL config error. The type is bool.
    OuputMinError interface{}

    // Exact output QL config error. The type is bool.
    OutputExactError interface{}

    // Maximum output QL config error. The type is bool.
    OutputMaxError interface{}

    // Input/Output mismatch error. The type is bool.
    InputOutputMismatch interface{}

    // Source.
    Source FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source

    // Configured minimum input QL.
    InputMinQl FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl

    // Configured exact input QL.
    InputExactQl FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl

    // Configured maximum input QL.
    InputMaxQl FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl

    // Configured mininum output QL.
    OutputMinQl FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl

    // Configured exact output QL.
    OutputExactQl FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl

    // Configured exact maximum QL.
    OutputMaxQl FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl
}

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetFilter() yfilter.YFilter { return errorSource.YFilter }

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) SetFilter(yf yfilter.YFilter) { errorSource.YFilter = yf }

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetGoName(yname string) string {
    if yname == "enable-error" { return "EnableError" }
    if yname == "input-min-error" { return "InputMinError" }
    if yname == "input-exact-error" { return "InputExactError" }
    if yname == "input-max-error" { return "InputMaxError" }
    if yname == "ouput-min-error" { return "OuputMinError" }
    if yname == "output-exact-error" { return "OutputExactError" }
    if yname == "output-max-error" { return "OutputMaxError" }
    if yname == "input-output-mismatch" { return "InputOutputMismatch" }
    if yname == "source" { return "Source" }
    if yname == "input-min-ql" { return "InputMinQl" }
    if yname == "input-exact-ql" { return "InputExactQl" }
    if yname == "input-max-ql" { return "InputMaxQl" }
    if yname == "output-min-ql" { return "OutputMinQl" }
    if yname == "output-exact-ql" { return "OutputExactQl" }
    if yname == "output-max-ql" { return "OutputMaxQl" }
    return ""
}

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetSegmentPath() string {
    return "error-source"
}

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        return &errorSource.Source
    }
    if childYangName == "input-min-ql" {
        return &errorSource.InputMinQl
    }
    if childYangName == "input-exact-ql" {
        return &errorSource.InputExactQl
    }
    if childYangName == "input-max-ql" {
        return &errorSource.InputMaxQl
    }
    if childYangName == "output-min-ql" {
        return &errorSource.OutputMinQl
    }
    if childYangName == "output-exact-ql" {
        return &errorSource.OutputExactQl
    }
    if childYangName == "output-max-ql" {
        return &errorSource.OutputMaxQl
    }
    return nil
}

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source"] = &errorSource.Source
    children["input-min-ql"] = &errorSource.InputMinQl
    children["input-exact-ql"] = &errorSource.InputExactQl
    children["input-max-ql"] = &errorSource.InputMaxQl
    children["output-min-ql"] = &errorSource.OutputMinQl
    children["output-exact-ql"] = &errorSource.OutputExactQl
    children["output-max-ql"] = &errorSource.OutputMaxQl
    return children
}

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable-error"] = errorSource.EnableError
    leafs["input-min-error"] = errorSource.InputMinError
    leafs["input-exact-error"] = errorSource.InputExactError
    leafs["input-max-error"] = errorSource.InputMaxError
    leafs["ouput-min-error"] = errorSource.OuputMinError
    leafs["output-exact-error"] = errorSource.OutputExactError
    leafs["output-max-error"] = errorSource.OutputMaxError
    leafs["input-output-mismatch"] = errorSource.InputOutputMismatch
    return leafs
}

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetBundleName() string { return "cisco_ios_xr" }

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetYangName() string { return "error-source" }

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) SetParent(parent types.Entity) { errorSource.parent = parent }

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetParent() types.Entity { return errorSource.parent }

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetParentYangName() string { return "configuration-errors" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source
// Source
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId
}

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetSegmentPath() string {
    return "source"
}

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &source.ClockId
    }
    return nil
}

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &source.ClockId
    return children
}

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = source.SourceClass
    leafs["ethernet-interface"] = source.EthernetInterface
    leafs["sonet-interface"] = source.SonetInterface
    leafs["node"] = source.Node
    leafs["ptp-node"] = source.PtpNode
    leafs["satellite-access-interface"] = source.SatelliteAccessInterface
    leafs["ntp-node"] = source.NtpNode
    return leafs
}

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetYangName() string { return "source" }

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetParent() types.Entity { return source.parent }

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetParentYangName() string { return "error-source" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId
// Clock ID
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetParentYangName() string { return "source" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl
// Configured minimum input QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetFilter() yfilter.YFilter { return inputMinQl.YFilter }

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) SetFilter(yf yfilter.YFilter) { inputMinQl.YFilter = yf }

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetSegmentPath() string {
    return "input-min-ql"
}

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = inputMinQl.QualityLevelOption
    leafs["option1-value"] = inputMinQl.Option1Value
    leafs["option2-generation1-value"] = inputMinQl.Option2Generation1Value
    leafs["option2-generation2-value"] = inputMinQl.Option2Generation2Value
    return leafs
}

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetBundleName() string { return "cisco_ios_xr" }

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetYangName() string { return "input-min-ql" }

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) SetParent(parent types.Entity) { inputMinQl.parent = parent }

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetParent() types.Entity { return inputMinQl.parent }

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetParentYangName() string { return "error-source" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl
// Configured exact input QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetFilter() yfilter.YFilter { return inputExactQl.YFilter }

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) SetFilter(yf yfilter.YFilter) { inputExactQl.YFilter = yf }

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetSegmentPath() string {
    return "input-exact-ql"
}

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = inputExactQl.QualityLevelOption
    leafs["option1-value"] = inputExactQl.Option1Value
    leafs["option2-generation1-value"] = inputExactQl.Option2Generation1Value
    leafs["option2-generation2-value"] = inputExactQl.Option2Generation2Value
    return leafs
}

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetBundleName() string { return "cisco_ios_xr" }

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetYangName() string { return "input-exact-ql" }

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) SetParent(parent types.Entity) { inputExactQl.parent = parent }

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetParent() types.Entity { return inputExactQl.parent }

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetParentYangName() string { return "error-source" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl
// Configured maximum input QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetFilter() yfilter.YFilter { return inputMaxQl.YFilter }

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) SetFilter(yf yfilter.YFilter) { inputMaxQl.YFilter = yf }

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetSegmentPath() string {
    return "input-max-ql"
}

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = inputMaxQl.QualityLevelOption
    leafs["option1-value"] = inputMaxQl.Option1Value
    leafs["option2-generation1-value"] = inputMaxQl.Option2Generation1Value
    leafs["option2-generation2-value"] = inputMaxQl.Option2Generation2Value
    return leafs
}

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetBundleName() string { return "cisco_ios_xr" }

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetYangName() string { return "input-max-ql" }

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) SetParent(parent types.Entity) { inputMaxQl.parent = parent }

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetParent() types.Entity { return inputMaxQl.parent }

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetParentYangName() string { return "error-source" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl
// Configured mininum output QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetFilter() yfilter.YFilter { return outputMinQl.YFilter }

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) SetFilter(yf yfilter.YFilter) { outputMinQl.YFilter = yf }

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetSegmentPath() string {
    return "output-min-ql"
}

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = outputMinQl.QualityLevelOption
    leafs["option1-value"] = outputMinQl.Option1Value
    leafs["option2-generation1-value"] = outputMinQl.Option2Generation1Value
    leafs["option2-generation2-value"] = outputMinQl.Option2Generation2Value
    return leafs
}

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetBundleName() string { return "cisco_ios_xr" }

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetYangName() string { return "output-min-ql" }

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) SetParent(parent types.Entity) { outputMinQl.parent = parent }

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetParent() types.Entity { return outputMinQl.parent }

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetParentYangName() string { return "error-source" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl
// Configured exact output QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetFilter() yfilter.YFilter { return outputExactQl.YFilter }

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) SetFilter(yf yfilter.YFilter) { outputExactQl.YFilter = yf }

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetSegmentPath() string {
    return "output-exact-ql"
}

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = outputExactQl.QualityLevelOption
    leafs["option1-value"] = outputExactQl.Option1Value
    leafs["option2-generation1-value"] = outputExactQl.Option2Generation1Value
    leafs["option2-generation2-value"] = outputExactQl.Option2Generation2Value
    return leafs
}

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetBundleName() string { return "cisco_ios_xr" }

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetYangName() string { return "output-exact-ql" }

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) SetParent(parent types.Entity) { outputExactQl.parent = parent }

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetParent() types.Entity { return outputExactQl.parent }

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetParentYangName() string { return "error-source" }

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl
// Configured exact maximum QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetFilter() yfilter.YFilter { return outputMaxQl.YFilter }

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) SetFilter(yf yfilter.YFilter) { outputMaxQl.YFilter = yf }

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetSegmentPath() string {
    return "output-max-ql"
}

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = outputMaxQl.QualityLevelOption
    leafs["option1-value"] = outputMaxQl.Option1Value
    leafs["option2-generation1-value"] = outputMaxQl.Option2Generation1Value
    leafs["option2-generation2-value"] = outputMaxQl.Option2Generation2Value
    return leafs
}

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetBundleName() string { return "cisco_ios_xr" }

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetYangName() string { return "output-max-ql" }

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) SetParent(parent types.Entity) { outputMaxQl.parent = parent }

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetParent() types.Entity { return outputMaxQl.parent }

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetParentYangName() string { return "error-source" }

// FrequencySynchronization_Nodes_Node_SsmSummary
// SSM operational data
type FrequencySynchronization_Nodes_Node_SsmSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of ethernet interfaces in synchronous mode. The type is interface{}
    // with range: 0..4294967295.
    EthernetSources interface{}

    // Number of ethernet interfaces assigned for selection. The type is
    // interface{} with range: 0..4294967295.
    EthernetSourcesSelect interface{}

    // Number of ethernet interfaces with SSM enabled. The type is interface{}
    // with range: 0..4294967295.
    EthernetSourcesEnabled interface{}

    // Number of SONET interfaces in synchronous mode. The type is interface{}
    // with range: 0..4294967295.
    SonetSources interface{}

    // Number of SONET interfaces assigned for selection. The type is interface{}
    // with range: 0..4294967295.
    SonetSourcesSelect interface{}

    // Number of SONET interfaces with SSM enabled. The type is interface{} with
    // range: 0..4294967295.
    SonetSourcesEnabled interface{}

    // Total event SSMs sent. The type is interface{} with range: 0..4294967295.
    EventsSent interface{}

    // Total event SSMs received. The type is interface{} with range:
    // 0..4294967295.
    EventsReceived interface{}

    // Total information SSMs sent. The type is interface{} with range:
    // 0..4294967295.
    InfosSent interface{}

    // Total information SSMs received. The type is interface{} with range:
    // 0..4294967295.
    InfosReceived interface{}

    // Total DNU SSMs sent. The type is interface{} with range: 0..4294967295.
    DnUsSent interface{}

    // Total DNU SSMs received. The type is interface{} with range: 0..4294967295.
    DnUsReceived interface{}
}

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetFilter() yfilter.YFilter { return ssmSummary.YFilter }

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) SetFilter(yf yfilter.YFilter) { ssmSummary.YFilter = yf }

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetGoName(yname string) string {
    if yname == "ethernet-sources" { return "EthernetSources" }
    if yname == "ethernet-sources-select" { return "EthernetSourcesSelect" }
    if yname == "ethernet-sources-enabled" { return "EthernetSourcesEnabled" }
    if yname == "sonet-sources" { return "SonetSources" }
    if yname == "sonet-sources-select" { return "SonetSourcesSelect" }
    if yname == "sonet-sources-enabled" { return "SonetSourcesEnabled" }
    if yname == "events-sent" { return "EventsSent" }
    if yname == "events-received" { return "EventsReceived" }
    if yname == "infos-sent" { return "InfosSent" }
    if yname == "infos-received" { return "InfosReceived" }
    if yname == "dn-us-sent" { return "DnUsSent" }
    if yname == "dn-us-received" { return "DnUsReceived" }
    return ""
}

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetSegmentPath() string {
    return "ssm-summary"
}

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethernet-sources"] = ssmSummary.EthernetSources
    leafs["ethernet-sources-select"] = ssmSummary.EthernetSourcesSelect
    leafs["ethernet-sources-enabled"] = ssmSummary.EthernetSourcesEnabled
    leafs["sonet-sources"] = ssmSummary.SonetSources
    leafs["sonet-sources-select"] = ssmSummary.SonetSourcesSelect
    leafs["sonet-sources-enabled"] = ssmSummary.SonetSourcesEnabled
    leafs["events-sent"] = ssmSummary.EventsSent
    leafs["events-received"] = ssmSummary.EventsReceived
    leafs["infos-sent"] = ssmSummary.InfosSent
    leafs["infos-received"] = ssmSummary.InfosReceived
    leafs["dn-us-sent"] = ssmSummary.DnUsSent
    leafs["dn-us-received"] = ssmSummary.DnUsReceived
    return leafs
}

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetYangName() string { return "ssm-summary" }

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) SetParent(parent types.Entity) { ssmSummary.parent = parent }

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetParent() types.Entity { return ssmSummary.parent }

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetParentYangName() string { return "node" }

// FrequencySynchronization_Summary
// Summary operational data
type FrequencySynchronization_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary of sources selected for frequency. The type is slice of
    // FrequencySynchronization_Summary_FrequencySummary.
    FrequencySummary []FrequencySynchronization_Summary_FrequencySummary

    // Summary of sources selected for time-of-day. The type is slice of
    // FrequencySynchronization_Summary_TimeOfDaySummary.
    TimeOfDaySummary []FrequencySynchronization_Summary_TimeOfDaySummary
}

func (summary *FrequencySynchronization_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *FrequencySynchronization_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *FrequencySynchronization_Summary) GetGoName(yname string) string {
    if yname == "frequency-summary" { return "FrequencySummary" }
    if yname == "time-of-day-summary" { return "TimeOfDaySummary" }
    return ""
}

func (summary *FrequencySynchronization_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *FrequencySynchronization_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "frequency-summary" {
        for _, c := range summary.FrequencySummary {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_Summary_FrequencySummary{}
        summary.FrequencySummary = append(summary.FrequencySummary, child)
        return &summary.FrequencySummary[len(summary.FrequencySummary)-1]
    }
    if childYangName == "time-of-day-summary" {
        for _, c := range summary.TimeOfDaySummary {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_Summary_TimeOfDaySummary{}
        summary.TimeOfDaySummary = append(summary.TimeOfDaySummary, child)
        return &summary.TimeOfDaySummary[len(summary.TimeOfDaySummary)-1]
    }
    return nil
}

func (summary *FrequencySynchronization_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range summary.FrequencySummary {
        children[summary.FrequencySummary[i].GetSegmentPath()] = &summary.FrequencySummary[i]
    }
    for i := range summary.TimeOfDaySummary {
        children[summary.TimeOfDaySummary[i].GetSegmentPath()] = &summary.TimeOfDaySummary[i]
    }
    return children
}

func (summary *FrequencySynchronization_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summary *FrequencySynchronization_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *FrequencySynchronization_Summary) GetYangName() string { return "summary" }

func (summary *FrequencySynchronization_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *FrequencySynchronization_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *FrequencySynchronization_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *FrequencySynchronization_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *FrequencySynchronization_Summary) GetParent() types.Entity { return summary.parent }

func (summary *FrequencySynchronization_Summary) GetParentYangName() string { return "frequency-synchronization" }

// FrequencySynchronization_Summary_FrequencySummary
// Summary of sources selected for frequency
type FrequencySynchronization_Summary_FrequencySummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of clock-interfaces being driven by the source. The type is
    // interface{} with range: 0..4294967295.
    ClockCount interface{}

    // The number of Ethernet interfaces being driven by the source. The type is
    // interface{} with range: 0..4294967295.
    EthernetCount interface{}

    // The number of SONET/SDH interfaces being driven by the source. The type is
    // interface{} with range: 0..4294967295.
    SonetCount interface{}

    // The source associated with this summary information.
    Source FrequencySynchronization_Summary_FrequencySummary_Source
}

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetFilter() yfilter.YFilter { return frequencySummary.YFilter }

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) SetFilter(yf yfilter.YFilter) { frequencySummary.YFilter = yf }

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetGoName(yname string) string {
    if yname == "clock-count" { return "ClockCount" }
    if yname == "ethernet-count" { return "EthernetCount" }
    if yname == "sonet-count" { return "SonetCount" }
    if yname == "source" { return "Source" }
    return ""
}

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetSegmentPath() string {
    return "frequency-summary"
}

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        return &frequencySummary.Source
    }
    return nil
}

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source"] = &frequencySummary.Source
    return children
}

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-count"] = frequencySummary.ClockCount
    leafs["ethernet-count"] = frequencySummary.EthernetCount
    leafs["sonet-count"] = frequencySummary.SonetCount
    return leafs
}

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetBundleName() string { return "cisco_ios_xr" }

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetYangName() string { return "frequency-summary" }

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) SetParent(parent types.Entity) { frequencySummary.parent = parent }

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetParent() types.Entity { return frequencySummary.parent }

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetParentYangName() string { return "summary" }

// FrequencySynchronization_Summary_FrequencySummary_Source
// The source associated with this summary
// information
type FrequencySynchronization_Summary_FrequencySummary_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Summary_FrequencySummary_Source_ClockId
}

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetSegmentPath() string {
    return "source"
}

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &source.ClockId
    }
    return nil
}

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &source.ClockId
    return children
}

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = source.SourceClass
    leafs["ethernet-interface"] = source.EthernetInterface
    leafs["sonet-interface"] = source.SonetInterface
    leafs["node"] = source.Node
    leafs["ptp-node"] = source.PtpNode
    leafs["satellite-access-interface"] = source.SatelliteAccessInterface
    leafs["ntp-node"] = source.NtpNode
    return leafs
}

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetYangName() string { return "source" }

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetParent() types.Entity { return source.parent }

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetParentYangName() string { return "frequency-summary" }

// FrequencySynchronization_Summary_FrequencySummary_Source_ClockId
// Clock ID
type FrequencySynchronization_Summary_FrequencySummary_Source_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetParentYangName() string { return "source" }

// FrequencySynchronization_Summary_TimeOfDaySummary
// Summary of sources selected for time-of-day
type FrequencySynchronization_Summary_TimeOfDaySummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of cards having their time-of-day set by the source. The type is
    // interface{} with range: 0..4294967295.
    NodeCount interface{}

    // The source associated with this summary information.
    Source FrequencySynchronization_Summary_TimeOfDaySummary_Source
}

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetFilter() yfilter.YFilter { return timeOfDaySummary.YFilter }

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) SetFilter(yf yfilter.YFilter) { timeOfDaySummary.YFilter = yf }

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetGoName(yname string) string {
    if yname == "node-count" { return "NodeCount" }
    if yname == "source" { return "Source" }
    return ""
}

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetSegmentPath() string {
    return "time-of-day-summary"
}

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        return &timeOfDaySummary.Source
    }
    return nil
}

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source"] = &timeOfDaySummary.Source
    return children
}

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-count"] = timeOfDaySummary.NodeCount
    return leafs
}

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetBundleName() string { return "cisco_ios_xr" }

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetYangName() string { return "time-of-day-summary" }

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) SetParent(parent types.Entity) { timeOfDaySummary.parent = parent }

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetParent() types.Entity { return timeOfDaySummary.parent }

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetParentYangName() string { return "summary" }

// FrequencySynchronization_Summary_TimeOfDaySummary_Source
// The source associated with this summary
// information
type FrequencySynchronization_Summary_TimeOfDaySummary_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId
}

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetSegmentPath() string {
    return "source"
}

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &source.ClockId
    }
    return nil
}

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &source.ClockId
    return children
}

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = source.SourceClass
    leafs["ethernet-interface"] = source.EthernetInterface
    leafs["sonet-interface"] = source.SonetInterface
    leafs["node"] = source.Node
    leafs["ptp-node"] = source.PtpNode
    leafs["satellite-access-interface"] = source.SatelliteAccessInterface
    leafs["ntp-node"] = source.NtpNode
    return leafs
}

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetYangName() string { return "source" }

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetParent() types.Entity { return source.parent }

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetParentYangName() string { return "time-of-day-summary" }

// FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId
// Clock ID
type FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetParentYangName() string { return "source" }

// FrequencySynchronization_GlobalNodes
// Table for global node-specific operational data
type FrequencySynchronization_GlobalNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global node-specific data for a particular node. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode.
    GlobalNode []FrequencySynchronization_GlobalNodes_GlobalNode
}

func (globalNodes *FrequencySynchronization_GlobalNodes) GetFilter() yfilter.YFilter { return globalNodes.YFilter }

func (globalNodes *FrequencySynchronization_GlobalNodes) SetFilter(yf yfilter.YFilter) { globalNodes.YFilter = yf }

func (globalNodes *FrequencySynchronization_GlobalNodes) GetGoName(yname string) string {
    if yname == "global-node" { return "GlobalNode" }
    return ""
}

func (globalNodes *FrequencySynchronization_GlobalNodes) GetSegmentPath() string {
    return "global-nodes"
}

func (globalNodes *FrequencySynchronization_GlobalNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-node" {
        for _, c := range globalNodes.GlobalNode {
            if globalNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_GlobalNodes_GlobalNode{}
        globalNodes.GlobalNode = append(globalNodes.GlobalNode, child)
        return &globalNodes.GlobalNode[len(globalNodes.GlobalNode)-1]
    }
    return nil
}

func (globalNodes *FrequencySynchronization_GlobalNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalNodes.GlobalNode {
        children[globalNodes.GlobalNode[i].GetSegmentPath()] = &globalNodes.GlobalNode[i]
    }
    return children
}

func (globalNodes *FrequencySynchronization_GlobalNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalNodes *FrequencySynchronization_GlobalNodes) GetBundleName() string { return "cisco_ios_xr" }

func (globalNodes *FrequencySynchronization_GlobalNodes) GetYangName() string { return "global-nodes" }

func (globalNodes *FrequencySynchronization_GlobalNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalNodes *FrequencySynchronization_GlobalNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalNodes *FrequencySynchronization_GlobalNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalNodes *FrequencySynchronization_GlobalNodes) SetParent(parent types.Entity) { globalNodes.parent = parent }

func (globalNodes *FrequencySynchronization_GlobalNodes) GetParent() types.Entity { return globalNodes.parent }

func (globalNodes *FrequencySynchronization_GlobalNodes) GetParentYangName() string { return "frequency-synchronization" }

// FrequencySynchronization_GlobalNodes_GlobalNode
// Global node-specific data for a particular node
type FrequencySynchronization_GlobalNodes_GlobalNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Selection backtrace operational data for time-of-day on a particular node.
    TimeOfDayBackTrace FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace

    // Selection forwardtrace operational data for a NTP clock.
    NtpSelectionForwardTrace FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace

    // Selection forwardtrace operational data for a PTP clock.
    PtpSelectionForwardTrace FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace
}

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetFilter() yfilter.YFilter { return globalNode.YFilter }

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) SetFilter(yf yfilter.YFilter) { globalNode.YFilter = yf }

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "time-of-day-back-trace" { return "TimeOfDayBackTrace" }
    if yname == "ntp-selection-forward-trace" { return "NtpSelectionForwardTrace" }
    if yname == "ptp-selection-forward-trace" { return "PtpSelectionForwardTrace" }
    return ""
}

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetSegmentPath() string {
    return "global-node" + "[node='" + fmt.Sprintf("%v", globalNode.Node) + "']"
}

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "time-of-day-back-trace" {
        return &globalNode.TimeOfDayBackTrace
    }
    if childYangName == "ntp-selection-forward-trace" {
        return &globalNode.NtpSelectionForwardTrace
    }
    if childYangName == "ptp-selection-forward-trace" {
        return &globalNode.PtpSelectionForwardTrace
    }
    return nil
}

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["time-of-day-back-trace"] = &globalNode.TimeOfDayBackTrace
    children["ntp-selection-forward-trace"] = &globalNode.NtpSelectionForwardTrace
    children["ptp-selection-forward-trace"] = &globalNode.PtpSelectionForwardTrace
    return children
}

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = globalNode.Node
    return leafs
}

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetBundleName() string { return "cisco_ios_xr" }

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetYangName() string { return "global-node" }

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) SetParent(parent types.Entity) { globalNode.parent = parent }

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetParent() types.Entity { return globalNode.parent }

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetParentYangName() string { return "global-nodes" }

// FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace
// Selection backtrace operational data for
// time-of-day on a particular node
type FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source which has been selected for output.
    SelectedSource FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource

    // List of selection points in the backtrace. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint.
    SelectionPoint []FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint
}

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetFilter() yfilter.YFilter { return timeOfDayBackTrace.YFilter }

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) SetFilter(yf yfilter.YFilter) { timeOfDayBackTrace.YFilter = yf }

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetGoName(yname string) string {
    if yname == "selected-source" { return "SelectedSource" }
    if yname == "selection-point" { return "SelectionPoint" }
    return ""
}

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetSegmentPath() string {
    return "time-of-day-back-trace"
}

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selected-source" {
        return &timeOfDayBackTrace.SelectedSource
    }
    if childYangName == "selection-point" {
        for _, c := range timeOfDayBackTrace.SelectionPoint {
            if timeOfDayBackTrace.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint{}
        timeOfDayBackTrace.SelectionPoint = append(timeOfDayBackTrace.SelectionPoint, child)
        return &timeOfDayBackTrace.SelectionPoint[len(timeOfDayBackTrace.SelectionPoint)-1]
    }
    return nil
}

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["selected-source"] = &timeOfDayBackTrace.SelectedSource
    for i := range timeOfDayBackTrace.SelectionPoint {
        children[timeOfDayBackTrace.SelectionPoint[i].GetSegmentPath()] = &timeOfDayBackTrace.SelectionPoint[i]
    }
    return children
}

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetBundleName() string { return "cisco_ios_xr" }

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetYangName() string { return "time-of-day-back-trace" }

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) SetParent(parent types.Entity) { timeOfDayBackTrace.parent = parent }

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetParent() types.Entity { return timeOfDayBackTrace.parent }

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetParentYangName() string { return "global-node" }

// FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource
// Source which has been selected for output
type FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetFilter() yfilter.YFilter { return selectedSource.YFilter }

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) SetFilter(yf yfilter.YFilter) { selectedSource.YFilter = yf }

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetSegmentPath() string {
    return "selected-source"
}

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &selectedSource.ClockId
    }
    return nil
}

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &selectedSource.ClockId
    return children
}

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = selectedSource.SourceClass
    leafs["ethernet-interface"] = selectedSource.EthernetInterface
    leafs["sonet-interface"] = selectedSource.SonetInterface
    leafs["node"] = selectedSource.Node
    leafs["ptp-node"] = selectedSource.PtpNode
    leafs["satellite-access-interface"] = selectedSource.SatelliteAccessInterface
    leafs["ntp-node"] = selectedSource.NtpNode
    return leafs
}

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetBundleName() string { return "cisco_ios_xr" }

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetYangName() string { return "selected-source" }

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) SetParent(parent types.Entity) { selectedSource.parent = parent }

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetParent() types.Entity { return selectedSource.parent }

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetParentYangName() string { return "time-of-day-back-trace" }

// FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetParentYangName() string { return "selected-source" }

// FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint
// List of selection points in the backtrace
type FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetFilter() yfilter.YFilter { return selectionPoint.YFilter }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) SetFilter(yf yfilter.YFilter) { selectionPoint.YFilter = yf }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "selection-point-description" { return "SelectionPointDescription" }
    if yname == "node" { return "Node" }
    return ""
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetSegmentPath() string {
    return "selection-point"
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point-type"] = selectionPoint.SelectionPointType
    leafs["selection-point-description"] = selectionPoint.SelectionPointDescription
    leafs["node"] = selectionPoint.Node
    return leafs
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetYangName() string { return "selection-point" }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) SetParent(parent types.Entity) { selectionPoint.parent = parent }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetParent() types.Entity { return selectionPoint.parent }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetParentYangName() string { return "time-of-day-back-trace" }

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace
// Selection forwardtrace operational data for a
// NTP clock
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection ForwardTrace. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace.
    ForwardTrace []FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace
}

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetFilter() yfilter.YFilter { return ntpSelectionForwardTrace.YFilter }

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) SetFilter(yf yfilter.YFilter) { ntpSelectionForwardTrace.YFilter = yf }

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetGoName(yname string) string {
    if yname == "forward-trace" { return "ForwardTrace" }
    return ""
}

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetSegmentPath() string {
    return "ntp-selection-forward-trace"
}

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forward-trace" {
        for _, c := range ntpSelectionForwardTrace.ForwardTrace {
            if ntpSelectionForwardTrace.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace{}
        ntpSelectionForwardTrace.ForwardTrace = append(ntpSelectionForwardTrace.ForwardTrace, child)
        return &ntpSelectionForwardTrace.ForwardTrace[len(ntpSelectionForwardTrace.ForwardTrace)-1]
    }
    return nil
}

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ntpSelectionForwardTrace.ForwardTrace {
        children[ntpSelectionForwardTrace.ForwardTrace[i].GetSegmentPath()] = &ntpSelectionForwardTrace.ForwardTrace[i]
    }
    return children
}

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetBundleName() string { return "cisco_ios_xr" }

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetYangName() string { return "ntp-selection-forward-trace" }

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) SetParent(parent types.Entity) { ntpSelectionForwardTrace.parent = parent }

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetParent() types.Entity { return ntpSelectionForwardTrace.parent }

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetParentYangName() string { return "global-node" }

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace
// Selection ForwardTrace
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The source or selection point at this point in the forwardtrace.
    ForwardTraceNode FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetFilter() yfilter.YFilter { return forwardTrace.YFilter }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) SetFilter(yf yfilter.YFilter) { forwardTrace.YFilter = yf }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetGoName(yname string) string {
    if yname == "forward-trace-node" { return "ForwardTraceNode" }
    return ""
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetSegmentPath() string {
    return "forward-trace"
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forward-trace-node" {
        return &forwardTrace.ForwardTraceNode
    }
    return nil
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["forward-trace-node"] = &forwardTrace.ForwardTraceNode
    return children
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetBundleName() string { return "cisco_ios_xr" }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetYangName() string { return "forward-trace" }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) SetParent(parent types.Entity) { forwardTrace.parent = parent }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetParent() types.Entity { return forwardTrace.parent }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetParentYangName() string { return "ntp-selection-forward-trace" }

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode
// The source or selection point at this point in
// the forwardtrace
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NodeType. The type is FsyncBagForwardtraceNode.
    NodeType interface{}

    // Selection Point.
    SelectionPoint FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint

    // Timing Source.
    Source FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetFilter() yfilter.YFilter { return forwardTraceNode.YFilter }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) SetFilter(yf yfilter.YFilter) { forwardTraceNode.YFilter = yf }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetGoName(yname string) string {
    if yname == "node-type" { return "NodeType" }
    if yname == "selection-point" { return "SelectionPoint" }
    if yname == "source" { return "Source" }
    return ""
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetSegmentPath() string {
    return "forward-trace-node"
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selection-point" {
        return &forwardTraceNode.SelectionPoint
    }
    if childYangName == "source" {
        return &forwardTraceNode.Source
    }
    return nil
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["selection-point"] = &forwardTraceNode.SelectionPoint
    children["source"] = &forwardTraceNode.Source
    return children
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-type"] = forwardTraceNode.NodeType
    return leafs
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetBundleName() string { return "cisco_ios_xr" }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetYangName() string { return "forward-trace-node" }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) SetParent(parent types.Entity) { forwardTraceNode.parent = parent }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetParent() types.Entity { return forwardTraceNode.parent }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetParentYangName() string { return "forward-trace" }

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint
// Selection Point
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetFilter() yfilter.YFilter { return selectionPoint.YFilter }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) SetFilter(yf yfilter.YFilter) { selectionPoint.YFilter = yf }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "selection-point-description" { return "SelectionPointDescription" }
    if yname == "node" { return "Node" }
    return ""
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetSegmentPath() string {
    return "selection-point"
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point-type"] = selectionPoint.SelectionPointType
    leafs["selection-point-description"] = selectionPoint.SelectionPointDescription
    leafs["node"] = selectionPoint.Node
    return leafs
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetYangName() string { return "selection-point" }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) SetParent(parent types.Entity) { selectionPoint.parent = parent }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetParent() types.Entity { return selectionPoint.parent }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetParentYangName() string { return "forward-trace-node" }

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
// Timing Source
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetSegmentPath() string {
    return "source"
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &source.ClockId
    }
    return nil
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &source.ClockId
    return children
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = source.SourceClass
    leafs["ethernet-interface"] = source.EthernetInterface
    leafs["sonet-interface"] = source.SonetInterface
    leafs["node"] = source.Node
    leafs["ptp-node"] = source.PtpNode
    leafs["satellite-access-interface"] = source.SatelliteAccessInterface
    leafs["ntp-node"] = source.NtpNode
    return leafs
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetYangName() string { return "source" }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetParent() types.Entity { return source.parent }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetParentYangName() string { return "forward-trace-node" }

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
// Clock ID
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetParentYangName() string { return "source" }

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace
// Selection forwardtrace operational data for a
// PTP clock
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection ForwardTrace. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace.
    ForwardTrace []FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace
}

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetFilter() yfilter.YFilter { return ptpSelectionForwardTrace.YFilter }

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) SetFilter(yf yfilter.YFilter) { ptpSelectionForwardTrace.YFilter = yf }

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetGoName(yname string) string {
    if yname == "forward-trace" { return "ForwardTrace" }
    return ""
}

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetSegmentPath() string {
    return "ptp-selection-forward-trace"
}

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forward-trace" {
        for _, c := range ptpSelectionForwardTrace.ForwardTrace {
            if ptpSelectionForwardTrace.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace{}
        ptpSelectionForwardTrace.ForwardTrace = append(ptpSelectionForwardTrace.ForwardTrace, child)
        return &ptpSelectionForwardTrace.ForwardTrace[len(ptpSelectionForwardTrace.ForwardTrace)-1]
    }
    return nil
}

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ptpSelectionForwardTrace.ForwardTrace {
        children[ptpSelectionForwardTrace.ForwardTrace[i].GetSegmentPath()] = &ptpSelectionForwardTrace.ForwardTrace[i]
    }
    return children
}

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetBundleName() string { return "cisco_ios_xr" }

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetYangName() string { return "ptp-selection-forward-trace" }

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) SetParent(parent types.Entity) { ptpSelectionForwardTrace.parent = parent }

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetParent() types.Entity { return ptpSelectionForwardTrace.parent }

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetParentYangName() string { return "global-node" }

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace
// Selection ForwardTrace
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The source or selection point at this point in the forwardtrace.
    ForwardTraceNode FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetFilter() yfilter.YFilter { return forwardTrace.YFilter }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) SetFilter(yf yfilter.YFilter) { forwardTrace.YFilter = yf }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetGoName(yname string) string {
    if yname == "forward-trace-node" { return "ForwardTraceNode" }
    return ""
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetSegmentPath() string {
    return "forward-trace"
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forward-trace-node" {
        return &forwardTrace.ForwardTraceNode
    }
    return nil
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["forward-trace-node"] = &forwardTrace.ForwardTraceNode
    return children
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetBundleName() string { return "cisco_ios_xr" }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetYangName() string { return "forward-trace" }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) SetParent(parent types.Entity) { forwardTrace.parent = parent }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetParent() types.Entity { return forwardTrace.parent }

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetParentYangName() string { return "ptp-selection-forward-trace" }

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode
// The source or selection point at this point in
// the forwardtrace
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NodeType. The type is FsyncBagForwardtraceNode.
    NodeType interface{}

    // Selection Point.
    SelectionPoint FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint

    // Timing Source.
    Source FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetFilter() yfilter.YFilter { return forwardTraceNode.YFilter }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) SetFilter(yf yfilter.YFilter) { forwardTraceNode.YFilter = yf }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetGoName(yname string) string {
    if yname == "node-type" { return "NodeType" }
    if yname == "selection-point" { return "SelectionPoint" }
    if yname == "source" { return "Source" }
    return ""
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetSegmentPath() string {
    return "forward-trace-node"
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selection-point" {
        return &forwardTraceNode.SelectionPoint
    }
    if childYangName == "source" {
        return &forwardTraceNode.Source
    }
    return nil
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["selection-point"] = &forwardTraceNode.SelectionPoint
    children["source"] = &forwardTraceNode.Source
    return children
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-type"] = forwardTraceNode.NodeType
    return leafs
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetBundleName() string { return "cisco_ios_xr" }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetYangName() string { return "forward-trace-node" }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) SetParent(parent types.Entity) { forwardTraceNode.parent = parent }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetParent() types.Entity { return forwardTraceNode.parent }

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetParentYangName() string { return "forward-trace" }

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint
// Selection Point
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetFilter() yfilter.YFilter { return selectionPoint.YFilter }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) SetFilter(yf yfilter.YFilter) { selectionPoint.YFilter = yf }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "selection-point-description" { return "SelectionPointDescription" }
    if yname == "node" { return "Node" }
    return ""
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetSegmentPath() string {
    return "selection-point"
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point-type"] = selectionPoint.SelectionPointType
    leafs["selection-point-description"] = selectionPoint.SelectionPointDescription
    leafs["node"] = selectionPoint.Node
    return leafs
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetYangName() string { return "selection-point" }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) SetParent(parent types.Entity) { selectionPoint.parent = parent }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetParent() types.Entity { return selectionPoint.parent }

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetParentYangName() string { return "forward-trace-node" }

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
// Timing Source
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetSegmentPath() string {
    return "source"
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &source.ClockId
    }
    return nil
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &source.ClockId
    return children
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = source.SourceClass
    leafs["ethernet-interface"] = source.EthernetInterface
    leafs["sonet-interface"] = source.SonetInterface
    leafs["node"] = source.Node
    leafs["ptp-node"] = source.PtpNode
    leafs["satellite-access-interface"] = source.SatelliteAccessInterface
    leafs["ntp-node"] = source.NtpNode
    return leafs
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetYangName() string { return "source" }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetParent() types.Entity { return source.parent }

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetParentYangName() string { return "forward-trace-node" }

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
// Clock ID
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetParentYangName() string { return "source" }

// FrequencySynchronization_Interfaces
// Table for interface operational data
type FrequencySynchronization_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for a particular interface. The type is slice of
    // FrequencySynchronization_Interfaces_Interface.
    Interface []FrequencySynchronization_Interfaces_Interface
}

func (interfaces *FrequencySynchronization_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *FrequencySynchronization_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *FrequencySynchronization_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *FrequencySynchronization_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *FrequencySynchronization_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *FrequencySynchronization_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *FrequencySynchronization_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *FrequencySynchronization_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *FrequencySynchronization_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *FrequencySynchronization_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *FrequencySynchronization_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *FrequencySynchronization_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *FrequencySynchronization_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *FrequencySynchronization_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *FrequencySynchronization_Interfaces) GetParentYangName() string { return "frequency-synchronization" }

// FrequencySynchronization_Interfaces_Interface
// Operational data for a particular interface
type FrequencySynchronization_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface name. The type is string.
    Name interface{}

    // Interface state. The type is ImStateEnum.
    State interface{}

    // SSM is enabled on the interface. The type is bool.
    SsmEnabled interface{}

    // The interface output is squelched. The type is bool.
    Squelched interface{}

    // The interface is an input for selection. The type is bool.
    SelectionInput interface{}

    // Priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Time-of-day priority. The type is interface{} with range: 0..255.
    TimeOfDayPriority interface{}

    // Damping state. The type is FsyncBagDampingState.
    DampingState interface{}

    // Time until damping state changes in ms. The type is interface{} with range:
    // 0..4294967295.
    DampingTime interface{}

    // Wait-to-restore time for the interface. The type is interface{} with range:
    // 0..65535.
    WaitToRestoreTime interface{}

    // The PTP clock supports frequency. The type is bool.
    SupportsFrequency interface{}

    // The PTP clock supports time. The type is bool.
    SupportsTimeOfDay interface{}

    // Spa selection points. The type is slice of interface{} with range: 0..255.
    SpaSelectionPoint []interface{}

    // Spa selection points descrption. The type is slice of string.
    SpaSelectionPointsDescription []interface{}

    // Node selection points. The type is slice of interface{} with range: 0..255.
    NodeSelectionPoint []interface{}

    // Node selection points descrption. The type is slice of string.
    NodeSelectionPointsDescription []interface{}

    // The source ID for the interface.
    Source FrequencySynchronization_Interfaces_Interface_Source

    // Timing source selected for interface output.
    SelectedSource FrequencySynchronization_Interfaces_Interface_SelectedSource

    // Received quality level.
    QualityLevelReceived FrequencySynchronization_Interfaces_Interface_QualityLevelReceived

    // Quality level after damping has been applied.
    QualityLevelDamped FrequencySynchronization_Interfaces_Interface_QualityLevelDamped

    // The effective input quality level.
    QualityLevelEffectiveInput FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput

    // The effective output quality level.
    QualityLevelEffectiveOutput FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput

    // The quality level of the source driving this interface.
    QualityLevelSelectedSource FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource

    // Ethernet peer information.
    EthernetPeerInformation FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation

    // ESMC Statistics.
    EsmcStatistics FrequencySynchronization_Interfaces_Interface_EsmcStatistics
}

func (self *FrequencySynchronization_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *FrequencySynchronization_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *FrequencySynchronization_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "name" { return "Name" }
    if yname == "state" { return "State" }
    if yname == "ssm-enabled" { return "SsmEnabled" }
    if yname == "squelched" { return "Squelched" }
    if yname == "selection-input" { return "SelectionInput" }
    if yname == "priority" { return "Priority" }
    if yname == "time-of-day-priority" { return "TimeOfDayPriority" }
    if yname == "damping-state" { return "DampingState" }
    if yname == "damping-time" { return "DampingTime" }
    if yname == "wait-to-restore-time" { return "WaitToRestoreTime" }
    if yname == "supports-frequency" { return "SupportsFrequency" }
    if yname == "supports-time-of-day" { return "SupportsTimeOfDay" }
    if yname == "spa-selection-point" { return "SpaSelectionPoint" }
    if yname == "spa-selection-points-description" { return "SpaSelectionPointsDescription" }
    if yname == "node-selection-point" { return "NodeSelectionPoint" }
    if yname == "node-selection-points-description" { return "NodeSelectionPointsDescription" }
    if yname == "source" { return "Source" }
    if yname == "selected-source" { return "SelectedSource" }
    if yname == "quality-level-received" { return "QualityLevelReceived" }
    if yname == "quality-level-damped" { return "QualityLevelDamped" }
    if yname == "quality-level-effective-input" { return "QualityLevelEffectiveInput" }
    if yname == "quality-level-effective-output" { return "QualityLevelEffectiveOutput" }
    if yname == "quality-level-selected-source" { return "QualityLevelSelectedSource" }
    if yname == "ethernet-peer-information" { return "EthernetPeerInformation" }
    if yname == "esmc-statistics" { return "EsmcStatistics" }
    return ""
}

func (self *FrequencySynchronization_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *FrequencySynchronization_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        return &self.Source
    }
    if childYangName == "selected-source" {
        return &self.SelectedSource
    }
    if childYangName == "quality-level-received" {
        return &self.QualityLevelReceived
    }
    if childYangName == "quality-level-damped" {
        return &self.QualityLevelDamped
    }
    if childYangName == "quality-level-effective-input" {
        return &self.QualityLevelEffectiveInput
    }
    if childYangName == "quality-level-effective-output" {
        return &self.QualityLevelEffectiveOutput
    }
    if childYangName == "quality-level-selected-source" {
        return &self.QualityLevelSelectedSource
    }
    if childYangName == "ethernet-peer-information" {
        return &self.EthernetPeerInformation
    }
    if childYangName == "esmc-statistics" {
        return &self.EsmcStatistics
    }
    return nil
}

func (self *FrequencySynchronization_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source"] = &self.Source
    children["selected-source"] = &self.SelectedSource
    children["quality-level-received"] = &self.QualityLevelReceived
    children["quality-level-damped"] = &self.QualityLevelDamped
    children["quality-level-effective-input"] = &self.QualityLevelEffectiveInput
    children["quality-level-effective-output"] = &self.QualityLevelEffectiveOutput
    children["quality-level-selected-source"] = &self.QualityLevelSelectedSource
    children["ethernet-peer-information"] = &self.EthernetPeerInformation
    children["esmc-statistics"] = &self.EsmcStatistics
    return children
}

func (self *FrequencySynchronization_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["name"] = self.Name
    leafs["state"] = self.State
    leafs["ssm-enabled"] = self.SsmEnabled
    leafs["squelched"] = self.Squelched
    leafs["selection-input"] = self.SelectionInput
    leafs["priority"] = self.Priority
    leafs["time-of-day-priority"] = self.TimeOfDayPriority
    leafs["damping-state"] = self.DampingState
    leafs["damping-time"] = self.DampingTime
    leafs["wait-to-restore-time"] = self.WaitToRestoreTime
    leafs["supports-frequency"] = self.SupportsFrequency
    leafs["supports-time-of-day"] = self.SupportsTimeOfDay
    leafs["spa-selection-point"] = self.SpaSelectionPoint
    leafs["spa-selection-points-description"] = self.SpaSelectionPointsDescription
    leafs["node-selection-point"] = self.NodeSelectionPoint
    leafs["node-selection-points-description"] = self.NodeSelectionPointsDescription
    return leafs
}

func (self *FrequencySynchronization_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *FrequencySynchronization_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *FrequencySynchronization_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *FrequencySynchronization_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *FrequencySynchronization_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *FrequencySynchronization_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *FrequencySynchronization_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *FrequencySynchronization_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// FrequencySynchronization_Interfaces_Interface_Source
// The source ID for the interface
type FrequencySynchronization_Interfaces_Interface_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Interfaces_Interface_Source_ClockId
}

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *FrequencySynchronization_Interfaces_Interface_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetSegmentPath() string {
    return "source"
}

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &source.ClockId
    }
    return nil
}

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &source.ClockId
    return children
}

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = source.SourceClass
    leafs["ethernet-interface"] = source.EthernetInterface
    leafs["sonet-interface"] = source.SonetInterface
    leafs["node"] = source.Node
    leafs["ptp-node"] = source.PtpNode
    leafs["satellite-access-interface"] = source.SatelliteAccessInterface
    leafs["ntp-node"] = source.NtpNode
    return leafs
}

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetYangName() string { return "source" }

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *FrequencySynchronization_Interfaces_Interface_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetParent() types.Entity { return source.parent }

func (source *FrequencySynchronization_Interfaces_Interface_Source) GetParentYangName() string { return "interface" }

// FrequencySynchronization_Interfaces_Interface_Source_ClockId
// Clock ID
type FrequencySynchronization_Interfaces_Interface_Source_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_Interfaces_Interface_Source_ClockId) GetParentYangName() string { return "source" }

// FrequencySynchronization_Interfaces_Interface_SelectedSource
// Timing source selected for interface output
type FrequencySynchronization_Interfaces_Interface_SelectedSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetFilter() yfilter.YFilter { return selectedSource.YFilter }

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) SetFilter(yf yfilter.YFilter) { selectedSource.YFilter = yf }

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetSegmentPath() string {
    return "selected-source"
}

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &selectedSource.ClockId
    }
    return nil
}

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &selectedSource.ClockId
    return children
}

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = selectedSource.SourceClass
    leafs["ethernet-interface"] = selectedSource.EthernetInterface
    leafs["sonet-interface"] = selectedSource.SonetInterface
    leafs["node"] = selectedSource.Node
    leafs["ptp-node"] = selectedSource.PtpNode
    leafs["satellite-access-interface"] = selectedSource.SatelliteAccessInterface
    leafs["ntp-node"] = selectedSource.NtpNode
    return leafs
}

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetBundleName() string { return "cisco_ios_xr" }

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetYangName() string { return "selected-source" }

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) SetParent(parent types.Entity) { selectedSource.parent = parent }

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetParent() types.Entity { return selectedSource.parent }

func (selectedSource *FrequencySynchronization_Interfaces_Interface_SelectedSource) GetParentYangName() string { return "interface" }

// FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_Interfaces_Interface_SelectedSource_ClockId) GetParentYangName() string { return "selected-source" }

// FrequencySynchronization_Interfaces_Interface_QualityLevelReceived
// Received quality level
type FrequencySynchronization_Interfaces_Interface_QualityLevelReceived struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetFilter() yfilter.YFilter { return qualityLevelReceived.YFilter }

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) SetFilter(yf yfilter.YFilter) { qualityLevelReceived.YFilter = yf }

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetSegmentPath() string {
    return "quality-level-received"
}

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelReceived.QualityLevelOption
    leafs["option1-value"] = qualityLevelReceived.Option1Value
    leafs["option2-generation1-value"] = qualityLevelReceived.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelReceived.Option2Generation2Value
    return leafs
}

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetYangName() string { return "quality-level-received" }

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) SetParent(parent types.Entity) { qualityLevelReceived.parent = parent }

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetParent() types.Entity { return qualityLevelReceived.parent }

func (qualityLevelReceived *FrequencySynchronization_Interfaces_Interface_QualityLevelReceived) GetParentYangName() string { return "interface" }

// FrequencySynchronization_Interfaces_Interface_QualityLevelDamped
// Quality level after damping has been applied
type FrequencySynchronization_Interfaces_Interface_QualityLevelDamped struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetFilter() yfilter.YFilter { return qualityLevelDamped.YFilter }

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) SetFilter(yf yfilter.YFilter) { qualityLevelDamped.YFilter = yf }

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetSegmentPath() string {
    return "quality-level-damped"
}

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelDamped.QualityLevelOption
    leafs["option1-value"] = qualityLevelDamped.Option1Value
    leafs["option2-generation1-value"] = qualityLevelDamped.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelDamped.Option2Generation2Value
    return leafs
}

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetYangName() string { return "quality-level-damped" }

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) SetParent(parent types.Entity) { qualityLevelDamped.parent = parent }

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetParent() types.Entity { return qualityLevelDamped.parent }

func (qualityLevelDamped *FrequencySynchronization_Interfaces_Interface_QualityLevelDamped) GetParentYangName() string { return "interface" }

// FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput
// The effective input quality level
type FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetFilter() yfilter.YFilter { return qualityLevelEffectiveInput.YFilter }

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) SetFilter(yf yfilter.YFilter) { qualityLevelEffectiveInput.YFilter = yf }

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetSegmentPath() string {
    return "quality-level-effective-input"
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelEffectiveInput.QualityLevelOption
    leafs["option1-value"] = qualityLevelEffectiveInput.Option1Value
    leafs["option2-generation1-value"] = qualityLevelEffectiveInput.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelEffectiveInput.Option2Generation2Value
    return leafs
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetYangName() string { return "quality-level-effective-input" }

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) SetParent(parent types.Entity) { qualityLevelEffectiveInput.parent = parent }

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetParent() types.Entity { return qualityLevelEffectiveInput.parent }

func (qualityLevelEffectiveInput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveInput) GetParentYangName() string { return "interface" }

// FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput
// The effective output quality level
type FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetFilter() yfilter.YFilter { return qualityLevelEffectiveOutput.YFilter }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) SetFilter(yf yfilter.YFilter) { qualityLevelEffectiveOutput.YFilter = yf }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetSegmentPath() string {
    return "quality-level-effective-output"
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelEffectiveOutput.QualityLevelOption
    leafs["option1-value"] = qualityLevelEffectiveOutput.Option1Value
    leafs["option2-generation1-value"] = qualityLevelEffectiveOutput.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelEffectiveOutput.Option2Generation2Value
    return leafs
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetYangName() string { return "quality-level-effective-output" }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) SetParent(parent types.Entity) { qualityLevelEffectiveOutput.parent = parent }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetParent() types.Entity { return qualityLevelEffectiveOutput.parent }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Interfaces_Interface_QualityLevelEffectiveOutput) GetParentYangName() string { return "interface" }

// FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource
// The quality level of the source driving this
// interface
type FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetFilter() yfilter.YFilter { return qualityLevelSelectedSource.YFilter }

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) SetFilter(yf yfilter.YFilter) { qualityLevelSelectedSource.YFilter = yf }

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetSegmentPath() string {
    return "quality-level-selected-source"
}

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelSelectedSource.QualityLevelOption
    leafs["option1-value"] = qualityLevelSelectedSource.Option1Value
    leafs["option2-generation1-value"] = qualityLevelSelectedSource.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelSelectedSource.Option2Generation2Value
    return leafs
}

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetYangName() string { return "quality-level-selected-source" }

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) SetParent(parent types.Entity) { qualityLevelSelectedSource.parent = parent }

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetParent() types.Entity { return qualityLevelSelectedSource.parent }

func (qualityLevelSelectedSource *FrequencySynchronization_Interfaces_Interface_QualityLevelSelectedSource) GetParentYangName() string { return "interface" }

// FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation
// Ethernet peer information
type FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer state. The type is FsyncBagEsmcPeerState.
    PeerState interface{}

    // Number of times the peer has come up. The type is interface{} with range:
    // 0..65535.
    PeerUpCount interface{}

    // Number of times the peer has timed out. The type is interface{} with range:
    // 0..65535.
    PeerTimeoutCount interface{}

    // Time of last peer state transition.
    PeerStateTime FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime

    // Time of last SSM received.
    LastSsm FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm
}

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetFilter() yfilter.YFilter { return ethernetPeerInformation.YFilter }

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) SetFilter(yf yfilter.YFilter) { ethernetPeerInformation.YFilter = yf }

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetGoName(yname string) string {
    if yname == "peer-state" { return "PeerState" }
    if yname == "peer-up-count" { return "PeerUpCount" }
    if yname == "peer-timeout-count" { return "PeerTimeoutCount" }
    if yname == "peer-state-time" { return "PeerStateTime" }
    if yname == "last-ssm" { return "LastSsm" }
    return ""
}

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetSegmentPath() string {
    return "ethernet-peer-information"
}

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-state-time" {
        return &ethernetPeerInformation.PeerStateTime
    }
    if childYangName == "last-ssm" {
        return &ethernetPeerInformation.LastSsm
    }
    return nil
}

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-state-time"] = &ethernetPeerInformation.PeerStateTime
    children["last-ssm"] = &ethernetPeerInformation.LastSsm
    return children
}

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-state"] = ethernetPeerInformation.PeerState
    leafs["peer-up-count"] = ethernetPeerInformation.PeerUpCount
    leafs["peer-timeout-count"] = ethernetPeerInformation.PeerTimeoutCount
    return leafs
}

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetYangName() string { return "ethernet-peer-information" }

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) SetParent(parent types.Entity) { ethernetPeerInformation.parent = parent }

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetParent() types.Entity { return ethernetPeerInformation.parent }

func (ethernetPeerInformation *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation) GetParentYangName() string { return "interface" }

// FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime
// Time of last peer state transition
type FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetFilter() yfilter.YFilter { return peerStateTime.YFilter }

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) SetFilter(yf yfilter.YFilter) { peerStateTime.YFilter = yf }

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetSegmentPath() string {
    return "peer-state-time"
}

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = peerStateTime.Seconds
    leafs["nanoseconds"] = peerStateTime.Nanoseconds
    return leafs
}

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetBundleName() string { return "cisco_ios_xr" }

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetYangName() string { return "peer-state-time" }

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) SetParent(parent types.Entity) { peerStateTime.parent = parent }

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetParent() types.Entity { return peerStateTime.parent }

func (peerStateTime *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_PeerStateTime) GetParentYangName() string { return "ethernet-peer-information" }

// FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm
// Time of last SSM received
type FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetFilter() yfilter.YFilter { return lastSsm.YFilter }

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) SetFilter(yf yfilter.YFilter) { lastSsm.YFilter = yf }

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "nanoseconds" { return "Nanoseconds" }
    return ""
}

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetSegmentPath() string {
    return "last-ssm"
}

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = lastSsm.Seconds
    leafs["nanoseconds"] = lastSsm.Nanoseconds
    return leafs
}

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetBundleName() string { return "cisco_ios_xr" }

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetYangName() string { return "last-ssm" }

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) SetParent(parent types.Entity) { lastSsm.parent = parent }

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetParent() types.Entity { return lastSsm.parent }

func (lastSsm *FrequencySynchronization_Interfaces_Interface_EthernetPeerInformation_LastSsm) GetParentYangName() string { return "ethernet-peer-information" }

// FrequencySynchronization_Interfaces_Interface_EsmcStatistics
// ESMC Statistics
type FrequencySynchronization_Interfaces_Interface_EsmcStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of event SSMs sent. The type is interface{} with range: 0..65535.
    EsmcEventsSent interface{}

    // Number of event SSMs received. The type is interface{} with range:
    // 0..65535.
    EsmcEventsReceived interface{}

    // Number of info SSMs sent. The type is interface{} with range:
    // 0..4294967295.
    EsmcInfosSent interface{}

    // Number of info SSms received. The type is interface{} with range:
    // 0..4294967295.
    EsmcInfosReceived interface{}

    // Number of SSMs with DNU QL sent. The type is interface{} with range:
    // 0..4294967295.
    EsmcDnUsSent interface{}

    // Number of SSMs with DNU QL received. The type is interface{} with range:
    // 0..4294967295.
    EsmcDnUsReceived interface{}

    // Number of malformed packets received. The type is interface{} with range:
    // 0..65535.
    EsmcMalformedReceived interface{}

    // Number of received packets that failed to be handled. The type is
    // interface{} with range: 0..65535.
    EsmcReceivedError interface{}
}

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetFilter() yfilter.YFilter { return esmcStatistics.YFilter }

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) SetFilter(yf yfilter.YFilter) { esmcStatistics.YFilter = yf }

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetGoName(yname string) string {
    if yname == "esmc-events-sent" { return "EsmcEventsSent" }
    if yname == "esmc-events-received" { return "EsmcEventsReceived" }
    if yname == "esmc-infos-sent" { return "EsmcInfosSent" }
    if yname == "esmc-infos-received" { return "EsmcInfosReceived" }
    if yname == "esmc-dn-us-sent" { return "EsmcDnUsSent" }
    if yname == "esmc-dn-us-received" { return "EsmcDnUsReceived" }
    if yname == "esmc-malformed-received" { return "EsmcMalformedReceived" }
    if yname == "esmc-received-error" { return "EsmcReceivedError" }
    return ""
}

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetSegmentPath() string {
    return "esmc-statistics"
}

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["esmc-events-sent"] = esmcStatistics.EsmcEventsSent
    leafs["esmc-events-received"] = esmcStatistics.EsmcEventsReceived
    leafs["esmc-infos-sent"] = esmcStatistics.EsmcInfosSent
    leafs["esmc-infos-received"] = esmcStatistics.EsmcInfosReceived
    leafs["esmc-dn-us-sent"] = esmcStatistics.EsmcDnUsSent
    leafs["esmc-dn-us-received"] = esmcStatistics.EsmcDnUsReceived
    leafs["esmc-malformed-received"] = esmcStatistics.EsmcMalformedReceived
    leafs["esmc-received-error"] = esmcStatistics.EsmcReceivedError
    return leafs
}

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetYangName() string { return "esmc-statistics" }

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) SetParent(parent types.Entity) { esmcStatistics.parent = parent }

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetParent() types.Entity { return esmcStatistics.parent }

func (esmcStatistics *FrequencySynchronization_Interfaces_Interface_EsmcStatistics) GetParentYangName() string { return "interface" }

// FrequencySynchronization_ClockInterfaceSelectionForwardTraces
// Selection forwardtrace operational data for
// clock-interfaces
type FrequencySynchronization_ClockInterfaceSelectionForwardTraces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection forwardtrace operational data for a particular clock-interface.
    // The type is slice of
    // FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace.
    ClockInterfaceSelectionForwardTrace []FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace
}

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetFilter() yfilter.YFilter { return clockInterfaceSelectionForwardTraces.YFilter }

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) SetFilter(yf yfilter.YFilter) { clockInterfaceSelectionForwardTraces.YFilter = yf }

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetGoName(yname string) string {
    if yname == "clock-interface-selection-forward-trace" { return "ClockInterfaceSelectionForwardTrace" }
    return ""
}

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetSegmentPath() string {
    return "clock-interface-selection-forward-traces"
}

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-interface-selection-forward-trace" {
        for _, c := range clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace {
            if clockInterfaceSelectionForwardTraces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace{}
        clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace = append(clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace, child)
        return &clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace[len(clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace)-1]
    }
    return nil
}

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace {
        children[clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace[i].GetSegmentPath()] = &clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace[i]
    }
    return children
}

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetBundleName() string { return "cisco_ios_xr" }

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetYangName() string { return "clock-interface-selection-forward-traces" }

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) SetParent(parent types.Entity) { clockInterfaceSelectionForwardTraces.parent = parent }

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetParent() types.Entity { return clockInterfaceSelectionForwardTraces.parent }

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_ClockInterfaceSelectionForwardTraces) GetParentYangName() string { return "frequency-synchronization" }

// FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace
// Selection forwardtrace operational data for a
// particular clock-interface
type FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Clock Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClockName interface{}

    // Selection ForwardTrace. The type is slice of
    // FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace.
    ForwardTrace []FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace
}

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetFilter() yfilter.YFilter { return clockInterfaceSelectionForwardTrace.YFilter }

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) SetFilter(yf yfilter.YFilter) { clockInterfaceSelectionForwardTrace.YFilter = yf }

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetGoName(yname string) string {
    if yname == "clock-name" { return "ClockName" }
    if yname == "forward-trace" { return "ForwardTrace" }
    return ""
}

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetSegmentPath() string {
    return "clock-interface-selection-forward-trace" + "[clock-name='" + fmt.Sprintf("%v", clockInterfaceSelectionForwardTrace.ClockName) + "']"
}

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forward-trace" {
        for _, c := range clockInterfaceSelectionForwardTrace.ForwardTrace {
            if clockInterfaceSelectionForwardTrace.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace{}
        clockInterfaceSelectionForwardTrace.ForwardTrace = append(clockInterfaceSelectionForwardTrace.ForwardTrace, child)
        return &clockInterfaceSelectionForwardTrace.ForwardTrace[len(clockInterfaceSelectionForwardTrace.ForwardTrace)-1]
    }
    return nil
}

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clockInterfaceSelectionForwardTrace.ForwardTrace {
        children[clockInterfaceSelectionForwardTrace.ForwardTrace[i].GetSegmentPath()] = &clockInterfaceSelectionForwardTrace.ForwardTrace[i]
    }
    return children
}

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-name"] = clockInterfaceSelectionForwardTrace.ClockName
    return leafs
}

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetBundleName() string { return "cisco_ios_xr" }

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetYangName() string { return "clock-interface-selection-forward-trace" }

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) SetParent(parent types.Entity) { clockInterfaceSelectionForwardTrace.parent = parent }

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetParent() types.Entity { return clockInterfaceSelectionForwardTrace.parent }

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetParentYangName() string { return "clock-interface-selection-forward-traces" }

// FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace
// Selection ForwardTrace
type FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The source or selection point at this point in the forwardtrace.
    ForwardTraceNode FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode
}

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetFilter() yfilter.YFilter { return forwardTrace.YFilter }

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) SetFilter(yf yfilter.YFilter) { forwardTrace.YFilter = yf }

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetGoName(yname string) string {
    if yname == "forward-trace-node" { return "ForwardTraceNode" }
    return ""
}

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetSegmentPath() string {
    return "forward-trace"
}

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forward-trace-node" {
        return &forwardTrace.ForwardTraceNode
    }
    return nil
}

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["forward-trace-node"] = &forwardTrace.ForwardTraceNode
    return children
}

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetBundleName() string { return "cisco_ios_xr" }

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetYangName() string { return "forward-trace" }

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) SetParent(parent types.Entity) { forwardTrace.parent = parent }

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetParent() types.Entity { return forwardTrace.parent }

func (forwardTrace *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetParentYangName() string { return "clock-interface-selection-forward-trace" }

// FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode
// The source or selection point at this point in
// the forwardtrace
type FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NodeType. The type is FsyncBagForwardtraceNode.
    NodeType interface{}

    // Selection Point.
    SelectionPoint FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint

    // Timing Source.
    Source FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
}

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetFilter() yfilter.YFilter { return forwardTraceNode.YFilter }

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) SetFilter(yf yfilter.YFilter) { forwardTraceNode.YFilter = yf }

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetGoName(yname string) string {
    if yname == "node-type" { return "NodeType" }
    if yname == "selection-point" { return "SelectionPoint" }
    if yname == "source" { return "Source" }
    return ""
}

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetSegmentPath() string {
    return "forward-trace-node"
}

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selection-point" {
        return &forwardTraceNode.SelectionPoint
    }
    if childYangName == "source" {
        return &forwardTraceNode.Source
    }
    return nil
}

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["selection-point"] = &forwardTraceNode.SelectionPoint
    children["source"] = &forwardTraceNode.Source
    return children
}

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-type"] = forwardTraceNode.NodeType
    return leafs
}

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetBundleName() string { return "cisco_ios_xr" }

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetYangName() string { return "forward-trace-node" }

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) SetParent(parent types.Entity) { forwardTraceNode.parent = parent }

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetParent() types.Entity { return forwardTraceNode.parent }

func (forwardTraceNode *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetParentYangName() string { return "forward-trace" }

// FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint
// Selection Point
type FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetFilter() yfilter.YFilter { return selectionPoint.YFilter }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) SetFilter(yf yfilter.YFilter) { selectionPoint.YFilter = yf }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "selection-point-description" { return "SelectionPointDescription" }
    if yname == "node" { return "Node" }
    return ""
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetSegmentPath() string {
    return "selection-point"
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point-type"] = selectionPoint.SelectionPointType
    leafs["selection-point-description"] = selectionPoint.SelectionPointDescription
    leafs["node"] = selectionPoint.Node
    return leafs
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetYangName() string { return "selection-point" }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) SetParent(parent types.Entity) { selectionPoint.parent = parent }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetParent() types.Entity { return selectionPoint.parent }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetParentYangName() string { return "forward-trace-node" }

// FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
// Timing Source
type FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
}

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetSegmentPath() string {
    return "source"
}

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &source.ClockId
    }
    return nil
}

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &source.ClockId
    return children
}

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = source.SourceClass
    leafs["ethernet-interface"] = source.EthernetInterface
    leafs["sonet-interface"] = source.SonetInterface
    leafs["node"] = source.Node
    leafs["ptp-node"] = source.PtpNode
    leafs["satellite-access-interface"] = source.SatelliteAccessInterface
    leafs["ntp-node"] = source.NtpNode
    return leafs
}

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetYangName() string { return "source" }

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetParent() types.Entity { return source.parent }

func (source *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetParentYangName() string { return "forward-trace-node" }

// FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
// Clock ID
type FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetParentYangName() string { return "source" }

// FrequencySynchronization_ClockInterfaceSelectionBackTraces
// Selection backtrace operational data for
// clock-interfaces
type FrequencySynchronization_ClockInterfaceSelectionBackTraces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection backtrace operational data for a particular clock-interface. The
    // type is slice of
    // FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace.
    ClockInterfaceSelectionBackTrace []FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace
}

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetFilter() yfilter.YFilter { return clockInterfaceSelectionBackTraces.YFilter }

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) SetFilter(yf yfilter.YFilter) { clockInterfaceSelectionBackTraces.YFilter = yf }

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetGoName(yname string) string {
    if yname == "clock-interface-selection-back-trace" { return "ClockInterfaceSelectionBackTrace" }
    return ""
}

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetSegmentPath() string {
    return "clock-interface-selection-back-traces"
}

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-interface-selection-back-trace" {
        for _, c := range clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace {
            if clockInterfaceSelectionBackTraces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace{}
        clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace = append(clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace, child)
        return &clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace[len(clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace)-1]
    }
    return nil
}

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace {
        children[clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace[i].GetSegmentPath()] = &clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace[i]
    }
    return children
}

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetBundleName() string { return "cisco_ios_xr" }

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetYangName() string { return "clock-interface-selection-back-traces" }

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) SetParent(parent types.Entity) { clockInterfaceSelectionBackTraces.parent = parent }

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetParent() types.Entity { return clockInterfaceSelectionBackTraces.parent }

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_ClockInterfaceSelectionBackTraces) GetParentYangName() string { return "frequency-synchronization" }

// FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace
// Selection backtrace operational data for a
// particular clock-interface
type FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Clock Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClockName interface{}

    // Source which has been selected for output.
    SelectedSource FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource

    // List of selection points in the backtrace. The type is slice of
    // FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint.
    SelectionPoint []FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint
}

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetFilter() yfilter.YFilter { return clockInterfaceSelectionBackTrace.YFilter }

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) SetFilter(yf yfilter.YFilter) { clockInterfaceSelectionBackTrace.YFilter = yf }

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetGoName(yname string) string {
    if yname == "clock-name" { return "ClockName" }
    if yname == "selected-source" { return "SelectedSource" }
    if yname == "selection-point" { return "SelectionPoint" }
    return ""
}

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetSegmentPath() string {
    return "clock-interface-selection-back-trace" + "[clock-name='" + fmt.Sprintf("%v", clockInterfaceSelectionBackTrace.ClockName) + "']"
}

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selected-source" {
        return &clockInterfaceSelectionBackTrace.SelectedSource
    }
    if childYangName == "selection-point" {
        for _, c := range clockInterfaceSelectionBackTrace.SelectionPoint {
            if clockInterfaceSelectionBackTrace.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint{}
        clockInterfaceSelectionBackTrace.SelectionPoint = append(clockInterfaceSelectionBackTrace.SelectionPoint, child)
        return &clockInterfaceSelectionBackTrace.SelectionPoint[len(clockInterfaceSelectionBackTrace.SelectionPoint)-1]
    }
    return nil
}

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["selected-source"] = &clockInterfaceSelectionBackTrace.SelectedSource
    for i := range clockInterfaceSelectionBackTrace.SelectionPoint {
        children[clockInterfaceSelectionBackTrace.SelectionPoint[i].GetSegmentPath()] = &clockInterfaceSelectionBackTrace.SelectionPoint[i]
    }
    return children
}

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-name"] = clockInterfaceSelectionBackTrace.ClockName
    return leafs
}

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetBundleName() string { return "cisco_ios_xr" }

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetYangName() string { return "clock-interface-selection-back-trace" }

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) SetParent(parent types.Entity) { clockInterfaceSelectionBackTrace.parent = parent }

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetParent() types.Entity { return clockInterfaceSelectionBackTrace.parent }

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetParentYangName() string { return "clock-interface-selection-back-traces" }

// FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource
// Source which has been selected for output
type FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetFilter() yfilter.YFilter { return selectedSource.YFilter }

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) SetFilter(yf yfilter.YFilter) { selectedSource.YFilter = yf }

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetSegmentPath() string {
    return "selected-source"
}

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &selectedSource.ClockId
    }
    return nil
}

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &selectedSource.ClockId
    return children
}

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = selectedSource.SourceClass
    leafs["ethernet-interface"] = selectedSource.EthernetInterface
    leafs["sonet-interface"] = selectedSource.SonetInterface
    leafs["node"] = selectedSource.Node
    leafs["ptp-node"] = selectedSource.PtpNode
    leafs["satellite-access-interface"] = selectedSource.SatelliteAccessInterface
    leafs["ntp-node"] = selectedSource.NtpNode
    return leafs
}

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetBundleName() string { return "cisco_ios_xr" }

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetYangName() string { return "selected-source" }

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) SetParent(parent types.Entity) { selectedSource.parent = parent }

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetParent() types.Entity { return selectedSource.parent }

func (selectedSource *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetParentYangName() string { return "clock-interface-selection-back-trace" }

// FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetParentYangName() string { return "selected-source" }

// FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint
// List of selection points in the backtrace
type FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetFilter() yfilter.YFilter { return selectionPoint.YFilter }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) SetFilter(yf yfilter.YFilter) { selectionPoint.YFilter = yf }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "selection-point-description" { return "SelectionPointDescription" }
    if yname == "node" { return "Node" }
    return ""
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetSegmentPath() string {
    return "selection-point"
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point-type"] = selectionPoint.SelectionPointType
    leafs["selection-point-description"] = selectionPoint.SelectionPointDescription
    leafs["node"] = selectionPoint.Node
    return leafs
}

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetYangName() string { return "selection-point" }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) SetParent(parent types.Entity) { selectionPoint.parent = parent }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetParent() types.Entity { return selectionPoint.parent }

func (selectionPoint *FrequencySynchronization_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetParentYangName() string { return "clock-interface-selection-back-trace" }

// FrequencySynchronization_GlobalInterfaces
// Table for global interface operational data
type FrequencySynchronization_GlobalInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global interface information for a particular interface. The type is slice
    // of FrequencySynchronization_GlobalInterfaces_GlobalInterface.
    GlobalInterface []FrequencySynchronization_GlobalInterfaces_GlobalInterface
}

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetFilter() yfilter.YFilter { return globalInterfaces.YFilter }

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) SetFilter(yf yfilter.YFilter) { globalInterfaces.YFilter = yf }

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetGoName(yname string) string {
    if yname == "global-interface" { return "GlobalInterface" }
    return ""
}

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetSegmentPath() string {
    return "global-interfaces"
}

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-interface" {
        for _, c := range globalInterfaces.GlobalInterface {
            if globalInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_GlobalInterfaces_GlobalInterface{}
        globalInterfaces.GlobalInterface = append(globalInterfaces.GlobalInterface, child)
        return &globalInterfaces.GlobalInterface[len(globalInterfaces.GlobalInterface)-1]
    }
    return nil
}

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalInterfaces.GlobalInterface {
        children[globalInterfaces.GlobalInterface[i].GetSegmentPath()] = &globalInterfaces.GlobalInterface[i]
    }
    return children
}

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetYangName() string { return "global-interfaces" }

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) SetParent(parent types.Entity) { globalInterfaces.parent = parent }

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetParent() types.Entity { return globalInterfaces.parent }

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetParentYangName() string { return "frequency-synchronization" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface
// Global interface information for a particular
// interface
type FrequencySynchronization_GlobalInterfaces_GlobalInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Selection backtrace operational data for a particular interface.
    InterfaceSelectionBackTrace FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace

    // Selection forwardtrace operational data for a particular interface.
    InterfaceSelectionForwardTrace FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace
}

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetFilter() yfilter.YFilter { return globalInterface.YFilter }

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) SetFilter(yf yfilter.YFilter) { globalInterface.YFilter = yf }

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-selection-back-trace" { return "InterfaceSelectionBackTrace" }
    if yname == "interface-selection-forward-trace" { return "InterfaceSelectionForwardTrace" }
    return ""
}

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetSegmentPath() string {
    return "global-interface" + "[interface-name='" + fmt.Sprintf("%v", globalInterface.InterfaceName) + "']"
}

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-selection-back-trace" {
        return &globalInterface.InterfaceSelectionBackTrace
    }
    if childYangName == "interface-selection-forward-trace" {
        return &globalInterface.InterfaceSelectionForwardTrace
    }
    return nil
}

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-selection-back-trace"] = &globalInterface.InterfaceSelectionBackTrace
    children["interface-selection-forward-trace"] = &globalInterface.InterfaceSelectionForwardTrace
    return children
}

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = globalInterface.InterfaceName
    return leafs
}

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetBundleName() string { return "cisco_ios_xr" }

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetYangName() string { return "global-interface" }

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) SetParent(parent types.Entity) { globalInterface.parent = parent }

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetParent() types.Entity { return globalInterface.parent }

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetParentYangName() string { return "global-interfaces" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace
// Selection backtrace operational data for a
// particular interface
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source which has been selected for output.
    SelectedSource FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource

    // List of selection points in the backtrace. The type is slice of
    // FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint.
    SelectionPoint []FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint
}

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetFilter() yfilter.YFilter { return interfaceSelectionBackTrace.YFilter }

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) SetFilter(yf yfilter.YFilter) { interfaceSelectionBackTrace.YFilter = yf }

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetGoName(yname string) string {
    if yname == "selected-source" { return "SelectedSource" }
    if yname == "selection-point" { return "SelectionPoint" }
    return ""
}

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetSegmentPath() string {
    return "interface-selection-back-trace"
}

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selected-source" {
        return &interfaceSelectionBackTrace.SelectedSource
    }
    if childYangName == "selection-point" {
        for _, c := range interfaceSelectionBackTrace.SelectionPoint {
            if interfaceSelectionBackTrace.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint{}
        interfaceSelectionBackTrace.SelectionPoint = append(interfaceSelectionBackTrace.SelectionPoint, child)
        return &interfaceSelectionBackTrace.SelectionPoint[len(interfaceSelectionBackTrace.SelectionPoint)-1]
    }
    return nil
}

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["selected-source"] = &interfaceSelectionBackTrace.SelectedSource
    for i := range interfaceSelectionBackTrace.SelectionPoint {
        children[interfaceSelectionBackTrace.SelectionPoint[i].GetSegmentPath()] = &interfaceSelectionBackTrace.SelectionPoint[i]
    }
    return children
}

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetYangName() string { return "interface-selection-back-trace" }

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) SetParent(parent types.Entity) { interfaceSelectionBackTrace.parent = parent }

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetParent() types.Entity { return interfaceSelectionBackTrace.parent }

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetParentYangName() string { return "global-interface" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource
// Source which has been selected for output
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetFilter() yfilter.YFilter { return selectedSource.YFilter }

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) SetFilter(yf yfilter.YFilter) { selectedSource.YFilter = yf }

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetSegmentPath() string {
    return "selected-source"
}

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &selectedSource.ClockId
    }
    return nil
}

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &selectedSource.ClockId
    return children
}

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = selectedSource.SourceClass
    leafs["ethernet-interface"] = selectedSource.EthernetInterface
    leafs["sonet-interface"] = selectedSource.SonetInterface
    leafs["node"] = selectedSource.Node
    leafs["ptp-node"] = selectedSource.PtpNode
    leafs["satellite-access-interface"] = selectedSource.SatelliteAccessInterface
    leafs["ntp-node"] = selectedSource.NtpNode
    return leafs
}

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetBundleName() string { return "cisco_ios_xr" }

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetYangName() string { return "selected-source" }

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) SetParent(parent types.Entity) { selectedSource.parent = parent }

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetParent() types.Entity { return selectedSource.parent }

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetParentYangName() string { return "interface-selection-back-trace" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetParentYangName() string { return "selected-source" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint
// List of selection points in the backtrace
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetFilter() yfilter.YFilter { return selectionPoint.YFilter }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) SetFilter(yf yfilter.YFilter) { selectionPoint.YFilter = yf }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "selection-point-description" { return "SelectionPointDescription" }
    if yname == "node" { return "Node" }
    return ""
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetSegmentPath() string {
    return "selection-point"
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point-type"] = selectionPoint.SelectionPointType
    leafs["selection-point-description"] = selectionPoint.SelectionPointDescription
    leafs["node"] = selectionPoint.Node
    return leafs
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetYangName() string { return "selection-point" }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) SetParent(parent types.Entity) { selectionPoint.parent = parent }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetParent() types.Entity { return selectionPoint.parent }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetParentYangName() string { return "interface-selection-back-trace" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace
// Selection forwardtrace operational data for a
// particular interface
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection ForwardTrace. The type is slice of
    // FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace.
    ForwardTrace []FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace
}

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetFilter() yfilter.YFilter { return interfaceSelectionForwardTrace.YFilter }

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) SetFilter(yf yfilter.YFilter) { interfaceSelectionForwardTrace.YFilter = yf }

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetGoName(yname string) string {
    if yname == "forward-trace" { return "ForwardTrace" }
    return ""
}

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetSegmentPath() string {
    return "interface-selection-forward-trace"
}

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forward-trace" {
        for _, c := range interfaceSelectionForwardTrace.ForwardTrace {
            if interfaceSelectionForwardTrace.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace{}
        interfaceSelectionForwardTrace.ForwardTrace = append(interfaceSelectionForwardTrace.ForwardTrace, child)
        return &interfaceSelectionForwardTrace.ForwardTrace[len(interfaceSelectionForwardTrace.ForwardTrace)-1]
    }
    return nil
}

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaceSelectionForwardTrace.ForwardTrace {
        children[interfaceSelectionForwardTrace.ForwardTrace[i].GetSegmentPath()] = &interfaceSelectionForwardTrace.ForwardTrace[i]
    }
    return children
}

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetYangName() string { return "interface-selection-forward-trace" }

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) SetParent(parent types.Entity) { interfaceSelectionForwardTrace.parent = parent }

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetParent() types.Entity { return interfaceSelectionForwardTrace.parent }

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetParentYangName() string { return "global-interface" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace
// Selection ForwardTrace
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The source or selection point at this point in the forwardtrace.
    ForwardTraceNode FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode
}

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetFilter() yfilter.YFilter { return forwardTrace.YFilter }

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) SetFilter(yf yfilter.YFilter) { forwardTrace.YFilter = yf }

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetGoName(yname string) string {
    if yname == "forward-trace-node" { return "ForwardTraceNode" }
    return ""
}

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetSegmentPath() string {
    return "forward-trace"
}

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forward-trace-node" {
        return &forwardTrace.ForwardTraceNode
    }
    return nil
}

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["forward-trace-node"] = &forwardTrace.ForwardTraceNode
    return children
}

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetBundleName() string { return "cisco_ios_xr" }

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetYangName() string { return "forward-trace" }

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) SetParent(parent types.Entity) { forwardTrace.parent = parent }

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetParent() types.Entity { return forwardTrace.parent }

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetParentYangName() string { return "interface-selection-forward-trace" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode
// The source or selection point at this point in
// the forwardtrace
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NodeType. The type is FsyncBagForwardtraceNode.
    NodeType interface{}

    // Selection Point.
    SelectionPoint FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint

    // Timing Source.
    Source FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
}

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetFilter() yfilter.YFilter { return forwardTraceNode.YFilter }

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) SetFilter(yf yfilter.YFilter) { forwardTraceNode.YFilter = yf }

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetGoName(yname string) string {
    if yname == "node-type" { return "NodeType" }
    if yname == "selection-point" { return "SelectionPoint" }
    if yname == "source" { return "Source" }
    return ""
}

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetSegmentPath() string {
    return "forward-trace-node"
}

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "selection-point" {
        return &forwardTraceNode.SelectionPoint
    }
    if childYangName == "source" {
        return &forwardTraceNode.Source
    }
    return nil
}

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["selection-point"] = &forwardTraceNode.SelectionPoint
    children["source"] = &forwardTraceNode.Source
    return children
}

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-type"] = forwardTraceNode.NodeType
    return leafs
}

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetBundleName() string { return "cisco_ios_xr" }

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetYangName() string { return "forward-trace-node" }

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) SetParent(parent types.Entity) { forwardTraceNode.parent = parent }

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetParent() types.Entity { return forwardTraceNode.parent }

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetParentYangName() string { return "forward-trace" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint
// Selection Point
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetFilter() yfilter.YFilter { return selectionPoint.YFilter }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) SetFilter(yf yfilter.YFilter) { selectionPoint.YFilter = yf }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetGoName(yname string) string {
    if yname == "selection-point-type" { return "SelectionPointType" }
    if yname == "selection-point-description" { return "SelectionPointDescription" }
    if yname == "node" { return "Node" }
    return ""
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetSegmentPath() string {
    return "selection-point"
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selection-point-type"] = selectionPoint.SelectionPointType
    leafs["selection-point-description"] = selectionPoint.SelectionPointDescription
    leafs["node"] = selectionPoint.Node
    return leafs
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetBundleName() string { return "cisco_ios_xr" }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetYangName() string { return "selection-point" }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) SetParent(parent types.Entity) { selectionPoint.parent = parent }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetParent() types.Entity { return selectionPoint.parent }

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetParentYangName() string { return "forward-trace-node" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
// Timing Source
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
}

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetSegmentPath() string {
    return "source"
}

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &source.ClockId
    }
    return nil
}

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &source.ClockId
    return children
}

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = source.SourceClass
    leafs["ethernet-interface"] = source.EthernetInterface
    leafs["sonet-interface"] = source.SonetInterface
    leafs["node"] = source.Node
    leafs["ptp-node"] = source.PtpNode
    leafs["satellite-access-interface"] = source.SatelliteAccessInterface
    leafs["ntp-node"] = source.NtpNode
    return leafs
}

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetYangName() string { return "source" }

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetParent() types.Entity { return source.parent }

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetParentYangName() string { return "forward-trace-node" }

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
// Clock ID
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetParentYangName() string { return "source" }

// FrequencySynchronization_Clocks
// Table for clock operational data
type FrequencySynchronization_Clocks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Operational data for a particular clock. The type is slice of
    // FrequencySynchronization_Clocks_Clock.
    Clock []FrequencySynchronization_Clocks_Clock
}

func (clocks *FrequencySynchronization_Clocks) GetFilter() yfilter.YFilter { return clocks.YFilter }

func (clocks *FrequencySynchronization_Clocks) SetFilter(yf yfilter.YFilter) { clocks.YFilter = yf }

func (clocks *FrequencySynchronization_Clocks) GetGoName(yname string) string {
    if yname == "clock" { return "Clock" }
    return ""
}

func (clocks *FrequencySynchronization_Clocks) GetSegmentPath() string {
    return "clocks"
}

func (clocks *FrequencySynchronization_Clocks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock" {
        for _, c := range clocks.Clock {
            if clocks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FrequencySynchronization_Clocks_Clock{}
        clocks.Clock = append(clocks.Clock, child)
        return &clocks.Clock[len(clocks.Clock)-1]
    }
    return nil
}

func (clocks *FrequencySynchronization_Clocks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clocks.Clock {
        children[clocks.Clock[i].GetSegmentPath()] = &clocks.Clock[i]
    }
    return children
}

func (clocks *FrequencySynchronization_Clocks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clocks *FrequencySynchronization_Clocks) GetBundleName() string { return "cisco_ios_xr" }

func (clocks *FrequencySynchronization_Clocks) GetYangName() string { return "clocks" }

func (clocks *FrequencySynchronization_Clocks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clocks *FrequencySynchronization_Clocks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clocks *FrequencySynchronization_Clocks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clocks *FrequencySynchronization_Clocks) SetParent(parent types.Entity) { clocks.parent = parent }

func (clocks *FrequencySynchronization_Clocks) GetParent() types.Entity { return clocks.parent }

func (clocks *FrequencySynchronization_Clocks) GetParentYangName() string { return "frequency-synchronization" }

// FrequencySynchronization_Clocks_Clock
// Operational data for a particular clock
type FrequencySynchronization_Clocks_Clock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Clock Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClockName interface{}

    // Clock state. The type is FsyncBagSourceState.
    State interface{}

    // Why the clock is down. The type is string.
    DownReason interface{}

    // Clock description. The type is string.
    Description interface{}

    // Priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Time-of-day priority. The type is interface{} with range: 0..255.
    TimeOfDayPriority interface{}

    // The clock supports SSMs. The type is bool.
    SsmSupport interface{}

    // The clock output is squelched. The type is bool.
    SsmEnabled interface{}

    // The clock is looped back. The type is bool.
    Loopback interface{}

    // The clock is an input for selection. The type is bool.
    SelectionInput interface{}

    // The clock output is squelched. The type is bool.
    Squelched interface{}

    // Damping state. The type is FsyncBagDampingState.
    DampingState interface{}

    // Time until damping state changes in ms. The type is interface{} with range:
    // 0..4294967295.
    DampingTime interface{}

    // Timing input is disabled. The type is bool.
    InputDisabled interface{}

    // Timing output is disabled. The type is bool.
    OutputDisabled interface{}

    // Wait-to-restore time for the clock. The type is interface{} with range:
    // 0..65535.
    WaitToRestoreTime interface{}

    // The type of clock. The type is FsyncBagClockIntfClass.
    ClockType interface{}

    // The PTP clock supports frequency. The type is bool.
    SupportsFrequency interface{}

    // The PTP clock supports time. The type is bool.
    SupportsTimeOfDay interface{}

    // Spa selection points. The type is slice of interface{} with range: 0..255.
    SpaSelectionPoint []interface{}

    // Spa selection points descrption. The type is slice of string.
    SpaSelectionPointsDescription []interface{}

    // Node selection points. The type is slice of interface{} with range: 0..255.
    NodeSelectionPoint []interface{}

    // Node selection points descrption. The type is slice of string.
    NodeSelectionPointsDescription []interface{}

    // The source ID for the clock.
    Source FrequencySynchronization_Clocks_Clock_Source

    // Timing source selected for clock output.
    SelectedSource FrequencySynchronization_Clocks_Clock_SelectedSource

    // Received quality level.
    QualityLevelReceived FrequencySynchronization_Clocks_Clock_QualityLevelReceived

    // Quality level after damping has been applied.
    QualityLevelDamped FrequencySynchronization_Clocks_Clock_QualityLevelDamped

    // The effective input quality level.
    QualityLevelEffectiveInput FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput

    // The effective output quality level.
    QualityLevelEffectiveOutput FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput

    // The quality level of the source driving this interface.
    QualityLevelSelectedSource FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource
}

func (clock *FrequencySynchronization_Clocks_Clock) GetFilter() yfilter.YFilter { return clock.YFilter }

func (clock *FrequencySynchronization_Clocks_Clock) SetFilter(yf yfilter.YFilter) { clock.YFilter = yf }

func (clock *FrequencySynchronization_Clocks_Clock) GetGoName(yname string) string {
    if yname == "clock-name" { return "ClockName" }
    if yname == "state" { return "State" }
    if yname == "down-reason" { return "DownReason" }
    if yname == "description" { return "Description" }
    if yname == "priority" { return "Priority" }
    if yname == "time-of-day-priority" { return "TimeOfDayPriority" }
    if yname == "ssm-support" { return "SsmSupport" }
    if yname == "ssm-enabled" { return "SsmEnabled" }
    if yname == "loopback" { return "Loopback" }
    if yname == "selection-input" { return "SelectionInput" }
    if yname == "squelched" { return "Squelched" }
    if yname == "damping-state" { return "DampingState" }
    if yname == "damping-time" { return "DampingTime" }
    if yname == "input-disabled" { return "InputDisabled" }
    if yname == "output-disabled" { return "OutputDisabled" }
    if yname == "wait-to-restore-time" { return "WaitToRestoreTime" }
    if yname == "clock-type" { return "ClockType" }
    if yname == "supports-frequency" { return "SupportsFrequency" }
    if yname == "supports-time-of-day" { return "SupportsTimeOfDay" }
    if yname == "spa-selection-point" { return "SpaSelectionPoint" }
    if yname == "spa-selection-points-description" { return "SpaSelectionPointsDescription" }
    if yname == "node-selection-point" { return "NodeSelectionPoint" }
    if yname == "node-selection-points-description" { return "NodeSelectionPointsDescription" }
    if yname == "source" { return "Source" }
    if yname == "selected-source" { return "SelectedSource" }
    if yname == "quality-level-received" { return "QualityLevelReceived" }
    if yname == "quality-level-damped" { return "QualityLevelDamped" }
    if yname == "quality-level-effective-input" { return "QualityLevelEffectiveInput" }
    if yname == "quality-level-effective-output" { return "QualityLevelEffectiveOutput" }
    if yname == "quality-level-selected-source" { return "QualityLevelSelectedSource" }
    return ""
}

func (clock *FrequencySynchronization_Clocks_Clock) GetSegmentPath() string {
    return "clock" + "[clock-name='" + fmt.Sprintf("%v", clock.ClockName) + "']"
}

func (clock *FrequencySynchronization_Clocks_Clock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        return &clock.Source
    }
    if childYangName == "selected-source" {
        return &clock.SelectedSource
    }
    if childYangName == "quality-level-received" {
        return &clock.QualityLevelReceived
    }
    if childYangName == "quality-level-damped" {
        return &clock.QualityLevelDamped
    }
    if childYangName == "quality-level-effective-input" {
        return &clock.QualityLevelEffectiveInput
    }
    if childYangName == "quality-level-effective-output" {
        return &clock.QualityLevelEffectiveOutput
    }
    if childYangName == "quality-level-selected-source" {
        return &clock.QualityLevelSelectedSource
    }
    return nil
}

func (clock *FrequencySynchronization_Clocks_Clock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source"] = &clock.Source
    children["selected-source"] = &clock.SelectedSource
    children["quality-level-received"] = &clock.QualityLevelReceived
    children["quality-level-damped"] = &clock.QualityLevelDamped
    children["quality-level-effective-input"] = &clock.QualityLevelEffectiveInput
    children["quality-level-effective-output"] = &clock.QualityLevelEffectiveOutput
    children["quality-level-selected-source"] = &clock.QualityLevelSelectedSource
    return children
}

func (clock *FrequencySynchronization_Clocks_Clock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["clock-name"] = clock.ClockName
    leafs["state"] = clock.State
    leafs["down-reason"] = clock.DownReason
    leafs["description"] = clock.Description
    leafs["priority"] = clock.Priority
    leafs["time-of-day-priority"] = clock.TimeOfDayPriority
    leafs["ssm-support"] = clock.SsmSupport
    leafs["ssm-enabled"] = clock.SsmEnabled
    leafs["loopback"] = clock.Loopback
    leafs["selection-input"] = clock.SelectionInput
    leafs["squelched"] = clock.Squelched
    leafs["damping-state"] = clock.DampingState
    leafs["damping-time"] = clock.DampingTime
    leafs["input-disabled"] = clock.InputDisabled
    leafs["output-disabled"] = clock.OutputDisabled
    leafs["wait-to-restore-time"] = clock.WaitToRestoreTime
    leafs["clock-type"] = clock.ClockType
    leafs["supports-frequency"] = clock.SupportsFrequency
    leafs["supports-time-of-day"] = clock.SupportsTimeOfDay
    leafs["spa-selection-point"] = clock.SpaSelectionPoint
    leafs["spa-selection-points-description"] = clock.SpaSelectionPointsDescription
    leafs["node-selection-point"] = clock.NodeSelectionPoint
    leafs["node-selection-points-description"] = clock.NodeSelectionPointsDescription
    return leafs
}

func (clock *FrequencySynchronization_Clocks_Clock) GetBundleName() string { return "cisco_ios_xr" }

func (clock *FrequencySynchronization_Clocks_Clock) GetYangName() string { return "clock" }

func (clock *FrequencySynchronization_Clocks_Clock) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clock *FrequencySynchronization_Clocks_Clock) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clock *FrequencySynchronization_Clocks_Clock) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clock *FrequencySynchronization_Clocks_Clock) SetParent(parent types.Entity) { clock.parent = parent }

func (clock *FrequencySynchronization_Clocks_Clock) GetParent() types.Entity { return clock.parent }

func (clock *FrequencySynchronization_Clocks_Clock) GetParentYangName() string { return "clocks" }

// FrequencySynchronization_Clocks_Clock_Source
// The source ID for the clock
type FrequencySynchronization_Clocks_Clock_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Clocks_Clock_Source_ClockId
}

func (source *FrequencySynchronization_Clocks_Clock_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *FrequencySynchronization_Clocks_Clock_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *FrequencySynchronization_Clocks_Clock_Source) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (source *FrequencySynchronization_Clocks_Clock_Source) GetSegmentPath() string {
    return "source"
}

func (source *FrequencySynchronization_Clocks_Clock_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &source.ClockId
    }
    return nil
}

func (source *FrequencySynchronization_Clocks_Clock_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &source.ClockId
    return children
}

func (source *FrequencySynchronization_Clocks_Clock_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = source.SourceClass
    leafs["ethernet-interface"] = source.EthernetInterface
    leafs["sonet-interface"] = source.SonetInterface
    leafs["node"] = source.Node
    leafs["ptp-node"] = source.PtpNode
    leafs["satellite-access-interface"] = source.SatelliteAccessInterface
    leafs["ntp-node"] = source.NtpNode
    return leafs
}

func (source *FrequencySynchronization_Clocks_Clock_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *FrequencySynchronization_Clocks_Clock_Source) GetYangName() string { return "source" }

func (source *FrequencySynchronization_Clocks_Clock_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *FrequencySynchronization_Clocks_Clock_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *FrequencySynchronization_Clocks_Clock_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *FrequencySynchronization_Clocks_Clock_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *FrequencySynchronization_Clocks_Clock_Source) GetParent() types.Entity { return source.parent }

func (source *FrequencySynchronization_Clocks_Clock_Source) GetParentYangName() string { return "clock" }

// FrequencySynchronization_Clocks_Clock_Source_ClockId
// Clock ID
type FrequencySynchronization_Clocks_Clock_Source_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_Clocks_Clock_Source_ClockId) GetParentYangName() string { return "source" }

// FrequencySynchronization_Clocks_Clock_SelectedSource
// Timing source selected for clock output
type FrequencySynchronization_Clocks_Clock_SelectedSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetFilter() yfilter.YFilter { return selectedSource.YFilter }

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) SetFilter(yf yfilter.YFilter) { selectedSource.YFilter = yf }

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetGoName(yname string) string {
    if yname == "source-class" { return "SourceClass" }
    if yname == "ethernet-interface" { return "EthernetInterface" }
    if yname == "sonet-interface" { return "SonetInterface" }
    if yname == "node" { return "Node" }
    if yname == "ptp-node" { return "PtpNode" }
    if yname == "satellite-access-interface" { return "SatelliteAccessInterface" }
    if yname == "ntp-node" { return "NtpNode" }
    if yname == "clock-id" { return "ClockId" }
    return ""
}

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetSegmentPath() string {
    return "selected-source"
}

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clock-id" {
        return &selectedSource.ClockId
    }
    return nil
}

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clock-id"] = &selectedSource.ClockId
    return children
}

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-class"] = selectedSource.SourceClass
    leafs["ethernet-interface"] = selectedSource.EthernetInterface
    leafs["sonet-interface"] = selectedSource.SonetInterface
    leafs["node"] = selectedSource.Node
    leafs["ptp-node"] = selectedSource.PtpNode
    leafs["satellite-access-interface"] = selectedSource.SatelliteAccessInterface
    leafs["ntp-node"] = selectedSource.NtpNode
    return leafs
}

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetBundleName() string { return "cisco_ios_xr" }

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetYangName() string { return "selected-source" }

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) SetParent(parent types.Entity) { selectedSource.parent = parent }

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetParent() types.Entity { return selectedSource.parent }

func (selectedSource *FrequencySynchronization_Clocks_Clock_SelectedSource) GetParentYangName() string { return "clock" }

// FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetFilter() yfilter.YFilter { return clockId.YFilter }

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) SetFilter(yf yfilter.YFilter) { clockId.YFilter = yf }

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "port" { return "Port" }
    return ""
}

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetSegmentPath() string {
    return "clock-id"
}

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = clockId.Node
    leafs["port"] = clockId.Port
    return leafs
}

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetBundleName() string { return "cisco_ios_xr" }

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetYangName() string { return "clock-id" }

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) SetParent(parent types.Entity) { clockId.parent = parent }

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetParent() types.Entity { return clockId.parent }

func (clockId *FrequencySynchronization_Clocks_Clock_SelectedSource_ClockId) GetParentYangName() string { return "selected-source" }

// FrequencySynchronization_Clocks_Clock_QualityLevelReceived
// Received quality level
type FrequencySynchronization_Clocks_Clock_QualityLevelReceived struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetFilter() yfilter.YFilter { return qualityLevelReceived.YFilter }

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) SetFilter(yf yfilter.YFilter) { qualityLevelReceived.YFilter = yf }

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetSegmentPath() string {
    return "quality-level-received"
}

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelReceived.QualityLevelOption
    leafs["option1-value"] = qualityLevelReceived.Option1Value
    leafs["option2-generation1-value"] = qualityLevelReceived.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelReceived.Option2Generation2Value
    return leafs
}

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetYangName() string { return "quality-level-received" }

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) SetParent(parent types.Entity) { qualityLevelReceived.parent = parent }

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetParent() types.Entity { return qualityLevelReceived.parent }

func (qualityLevelReceived *FrequencySynchronization_Clocks_Clock_QualityLevelReceived) GetParentYangName() string { return "clock" }

// FrequencySynchronization_Clocks_Clock_QualityLevelDamped
// Quality level after damping has been applied
type FrequencySynchronization_Clocks_Clock_QualityLevelDamped struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetFilter() yfilter.YFilter { return qualityLevelDamped.YFilter }

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) SetFilter(yf yfilter.YFilter) { qualityLevelDamped.YFilter = yf }

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetSegmentPath() string {
    return "quality-level-damped"
}

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelDamped.QualityLevelOption
    leafs["option1-value"] = qualityLevelDamped.Option1Value
    leafs["option2-generation1-value"] = qualityLevelDamped.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelDamped.Option2Generation2Value
    return leafs
}

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetYangName() string { return "quality-level-damped" }

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) SetParent(parent types.Entity) { qualityLevelDamped.parent = parent }

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetParent() types.Entity { return qualityLevelDamped.parent }

func (qualityLevelDamped *FrequencySynchronization_Clocks_Clock_QualityLevelDamped) GetParentYangName() string { return "clock" }

// FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput
// The effective input quality level
type FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetFilter() yfilter.YFilter { return qualityLevelEffectiveInput.YFilter }

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) SetFilter(yf yfilter.YFilter) { qualityLevelEffectiveInput.YFilter = yf }

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetSegmentPath() string {
    return "quality-level-effective-input"
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelEffectiveInput.QualityLevelOption
    leafs["option1-value"] = qualityLevelEffectiveInput.Option1Value
    leafs["option2-generation1-value"] = qualityLevelEffectiveInput.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelEffectiveInput.Option2Generation2Value
    return leafs
}

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetYangName() string { return "quality-level-effective-input" }

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) SetParent(parent types.Entity) { qualityLevelEffectiveInput.parent = parent }

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetParent() types.Entity { return qualityLevelEffectiveInput.parent }

func (qualityLevelEffectiveInput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveInput) GetParentYangName() string { return "clock" }

// FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput
// The effective output quality level
type FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetFilter() yfilter.YFilter { return qualityLevelEffectiveOutput.YFilter }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) SetFilter(yf yfilter.YFilter) { qualityLevelEffectiveOutput.YFilter = yf }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetSegmentPath() string {
    return "quality-level-effective-output"
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelEffectiveOutput.QualityLevelOption
    leafs["option1-value"] = qualityLevelEffectiveOutput.Option1Value
    leafs["option2-generation1-value"] = qualityLevelEffectiveOutput.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelEffectiveOutput.Option2Generation2Value
    return leafs
}

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetYangName() string { return "quality-level-effective-output" }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) SetParent(parent types.Entity) { qualityLevelEffectiveOutput.parent = parent }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetParent() types.Entity { return qualityLevelEffectiveOutput.parent }

func (qualityLevelEffectiveOutput *FrequencySynchronization_Clocks_Clock_QualityLevelEffectiveOutput) GetParentYangName() string { return "clock" }

// FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource
// The quality level of the source driving this
// interface
type FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QualityLevelOption. The type is FsyncBagQlOption.
    QualityLevelOption interface{}

    // ITU-T Option 1 QL value. The type is FsyncBagQlO1Value.
    Option1Value interface{}

    // ITU-T Option 2, generation 1 value. The type is FsyncBagQlO2G1Value.
    Option2Generation1Value interface{}

    // ITU-T Option 2, generation 2 value. The type is FsyncBagQlO2G2Value.
    Option2Generation2Value interface{}
}

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetFilter() yfilter.YFilter { return qualityLevelSelectedSource.YFilter }

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) SetFilter(yf yfilter.YFilter) { qualityLevelSelectedSource.YFilter = yf }

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetGoName(yname string) string {
    if yname == "quality-level-option" { return "QualityLevelOption" }
    if yname == "option1-value" { return "Option1Value" }
    if yname == "option2-generation1-value" { return "Option2Generation1Value" }
    if yname == "option2-generation2-value" { return "Option2Generation2Value" }
    return ""
}

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetSegmentPath() string {
    return "quality-level-selected-source"
}

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["quality-level-option"] = qualityLevelSelectedSource.QualityLevelOption
    leafs["option1-value"] = qualityLevelSelectedSource.Option1Value
    leafs["option2-generation1-value"] = qualityLevelSelectedSource.Option2Generation1Value
    leafs["option2-generation2-value"] = qualityLevelSelectedSource.Option2Generation2Value
    return leafs
}

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetBundleName() string { return "cisco_ios_xr" }

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetYangName() string { return "quality-level-selected-source" }

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) SetParent(parent types.Entity) { qualityLevelSelectedSource.parent = parent }

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetParent() types.Entity { return qualityLevelSelectedSource.parent }

func (qualityLevelSelectedSource *FrequencySynchronization_Clocks_Clock_QualityLevelSelectedSource) GetParentYangName() string { return "clock" }

