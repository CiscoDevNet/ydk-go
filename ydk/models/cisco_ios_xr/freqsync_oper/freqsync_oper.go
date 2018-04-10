// This module contains a collection of YANG definitions
// for Cisco IOS-XR freqsync package operational data.
// 
// This module contains definitions
// for the following management objects:
//   frequency-synchronization: Frequency Synchronization
//     operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package freqsync_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package freqsync_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-freqsync-oper frequency-synchronization}", reflect.TypeOf(FrequencySynchronization{}))
    ydk.RegisterEntity("Cisco-IOS-XR-freqsync-oper:frequency-synchronization", reflect.TypeOf(FrequencySynchronization{}))
}

// FsyncStream represents Fsync stream
type FsyncStream string

const (
    // Stream input from a local source
    FsyncStream_local FsyncStream = "local"

    // Stream input from a selection point on a remote
    // node
    FsyncStream_selection_point FsyncStream = "selection-point"
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

// FsyncBagStreamState represents Platform stream status
type FsyncBagStreamState string

const (
    // Invalid stream
    FsyncBagStreamState_stream_invalid FsyncBagStreamState = "stream-invalid"

    // Unqualified stream
    FsyncBagStreamState_stream_unqualified FsyncBagStreamState = "stream-unqualified"

    // Stream available
    FsyncBagStreamState_stream_available FsyncBagStreamState = "stream-available"

    // Stream available acquiring
    FsyncBagStreamState_stream_available_acquiring FsyncBagStreamState = "stream-available-acquiring"

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

    // Unknown
    FsyncBagClockIntfClass_clock_class_unknown FsyncBagClockIntfClass = "clock-class-unknown"
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

// FsyncBagForwardtraceNode represents Selection forwardtrace node information
type FsyncBagForwardtraceNode string

const (
    // A selection point forwardtrace node
    FsyncBagForwardtraceNode_forward_trace_node_selection_point FsyncBagForwardtraceNode = "forward-trace-node-selection-point"

    // A timing source forwardtrace node
    FsyncBagForwardtraceNode_forward_trace_node_source FsyncBagForwardtraceNode = "forward-trace-node-source"
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

// FrequencySynchronization
// Frequency Synchronization operational data
type FrequencySynchronization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table for global node-specific operational data.
    GlobalNodes FrequencySynchronization_GlobalNodes

    // Table for global interface operational data.
    GlobalInterfaces FrequencySynchronization_GlobalInterfaces

    // Summary operational data.
    Summary FrequencySynchronization_Summary

    // Table for interface operational data.
    InterfaceDatas FrequencySynchronization_InterfaceDatas

    // Table for node-specific operational data.
    Nodes FrequencySynchronization_Nodes
}

func (frequencySynchronization *FrequencySynchronization) GetEntityData() *types.CommonEntityData {
    frequencySynchronization.EntityData.YFilter = frequencySynchronization.YFilter
    frequencySynchronization.EntityData.YangName = "frequency-synchronization"
    frequencySynchronization.EntityData.BundleName = "cisco_ios_xr"
    frequencySynchronization.EntityData.ParentYangName = "Cisco-IOS-XR-freqsync-oper"
    frequencySynchronization.EntityData.SegmentPath = "Cisco-IOS-XR-freqsync-oper:frequency-synchronization"
    frequencySynchronization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frequencySynchronization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frequencySynchronization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frequencySynchronization.EntityData.Children = make(map[string]types.YChild)
    frequencySynchronization.EntityData.Children["global-nodes"] = types.YChild{"GlobalNodes", &frequencySynchronization.GlobalNodes}
    frequencySynchronization.EntityData.Children["global-interfaces"] = types.YChild{"GlobalInterfaces", &frequencySynchronization.GlobalInterfaces}
    frequencySynchronization.EntityData.Children["summary"] = types.YChild{"Summary", &frequencySynchronization.Summary}
    frequencySynchronization.EntityData.Children["interface-datas"] = types.YChild{"InterfaceDatas", &frequencySynchronization.InterfaceDatas}
    frequencySynchronization.EntityData.Children["nodes"] = types.YChild{"Nodes", &frequencySynchronization.Nodes}
    frequencySynchronization.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(frequencySynchronization.EntityData)
}

// FrequencySynchronization_GlobalNodes
// Table for global node-specific operational data
type FrequencySynchronization_GlobalNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global node-specific data for a particular node. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode.
    GlobalNode []FrequencySynchronization_GlobalNodes_GlobalNode
}

func (globalNodes *FrequencySynchronization_GlobalNodes) GetEntityData() *types.CommonEntityData {
    globalNodes.EntityData.YFilter = globalNodes.YFilter
    globalNodes.EntityData.YangName = "global-nodes"
    globalNodes.EntityData.BundleName = "cisco_ios_xr"
    globalNodes.EntityData.ParentYangName = "frequency-synchronization"
    globalNodes.EntityData.SegmentPath = "global-nodes"
    globalNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalNodes.EntityData.Children = make(map[string]types.YChild)
    globalNodes.EntityData.Children["global-node"] = types.YChild{"GlobalNode", nil}
    for i := range globalNodes.GlobalNode {
        globalNodes.EntityData.Children[types.GetSegmentPath(&globalNodes.GlobalNode[i])] = types.YChild{"GlobalNode", &globalNodes.GlobalNode[i]}
    }
    globalNodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalNodes.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode
// Global node-specific data for a particular node
type FrequencySynchronization_GlobalNodes_GlobalNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Selection backtrace operational data for clock-interfaces.
    ClockInterfaceSelectionBackTraces FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces

    // Selection forwardtrace operational data for clock-interfaces.
    ClockInterfaceSelectionForwardTraces FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces

    // Selection backtrace operational data for time-of-day on a particular node.
    TimeOfDayBackTrace FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace

    // Selection forwardtrace operational data for a NTP clock.
    NtpSelectionForwardTrace FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace

    // Selection forwardtrace operational data for a PTP clock.
    PtpSelectionForwardTrace FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace
}

func (globalNode *FrequencySynchronization_GlobalNodes_GlobalNode) GetEntityData() *types.CommonEntityData {
    globalNode.EntityData.YFilter = globalNode.YFilter
    globalNode.EntityData.YangName = "global-node"
    globalNode.EntityData.BundleName = "cisco_ios_xr"
    globalNode.EntityData.ParentYangName = "global-nodes"
    globalNode.EntityData.SegmentPath = "global-node" + "[node='" + fmt.Sprintf("%v", globalNode.Node) + "']"
    globalNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalNode.EntityData.Children = make(map[string]types.YChild)
    globalNode.EntityData.Children["clock-interface-selection-back-traces"] = types.YChild{"ClockInterfaceSelectionBackTraces", &globalNode.ClockInterfaceSelectionBackTraces}
    globalNode.EntityData.Children["clock-interface-selection-forward-traces"] = types.YChild{"ClockInterfaceSelectionForwardTraces", &globalNode.ClockInterfaceSelectionForwardTraces}
    globalNode.EntityData.Children["time-of-day-back-trace"] = types.YChild{"TimeOfDayBackTrace", &globalNode.TimeOfDayBackTrace}
    globalNode.EntityData.Children["ntp-selection-forward-trace"] = types.YChild{"NtpSelectionForwardTrace", &globalNode.NtpSelectionForwardTrace}
    globalNode.EntityData.Children["ptp-selection-forward-trace"] = types.YChild{"PtpSelectionForwardTrace", &globalNode.PtpSelectionForwardTrace}
    globalNode.EntityData.Leafs = make(map[string]types.YLeaf)
    globalNode.EntityData.Leafs["node"] = types.YLeaf{"Node", globalNode.Node}
    return &(globalNode.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces
// Selection backtrace operational data for
// clock-interfaces
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection backtrace operational data for a particular clock-interface. The
    // type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace.
    ClockInterfaceSelectionBackTrace []FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace
}

func (clockInterfaceSelectionBackTraces *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces) GetEntityData() *types.CommonEntityData {
    clockInterfaceSelectionBackTraces.EntityData.YFilter = clockInterfaceSelectionBackTraces.YFilter
    clockInterfaceSelectionBackTraces.EntityData.YangName = "clock-interface-selection-back-traces"
    clockInterfaceSelectionBackTraces.EntityData.BundleName = "cisco_ios_xr"
    clockInterfaceSelectionBackTraces.EntityData.ParentYangName = "global-node"
    clockInterfaceSelectionBackTraces.EntityData.SegmentPath = "clock-interface-selection-back-traces"
    clockInterfaceSelectionBackTraces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockInterfaceSelectionBackTraces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockInterfaceSelectionBackTraces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockInterfaceSelectionBackTraces.EntityData.Children = make(map[string]types.YChild)
    clockInterfaceSelectionBackTraces.EntityData.Children["clock-interface-selection-back-trace"] = types.YChild{"ClockInterfaceSelectionBackTrace", nil}
    for i := range clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace {
        clockInterfaceSelectionBackTraces.EntityData.Children[types.GetSegmentPath(&clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace[i])] = types.YChild{"ClockInterfaceSelectionBackTrace", &clockInterfaceSelectionBackTraces.ClockInterfaceSelectionBackTrace[i]}
    }
    clockInterfaceSelectionBackTraces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clockInterfaceSelectionBackTraces.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace
// Selection backtrace operational data for a
// particular clock-interface
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Clock type. The type is FsyncClock.
    ClockType interface{}

    // This attribute is a key. Clock port. The type is interface{} with range:
    // -2147483648..2147483647.
    Port interface{}

    // Source which has been selected for output.
    SelectedSource FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource

    // List of selection points in the backtrace. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint.
    SelectionPoint []FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint
}

func (clockInterfaceSelectionBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace) GetEntityData() *types.CommonEntityData {
    clockInterfaceSelectionBackTrace.EntityData.YFilter = clockInterfaceSelectionBackTrace.YFilter
    clockInterfaceSelectionBackTrace.EntityData.YangName = "clock-interface-selection-back-trace"
    clockInterfaceSelectionBackTrace.EntityData.BundleName = "cisco_ios_xr"
    clockInterfaceSelectionBackTrace.EntityData.ParentYangName = "clock-interface-selection-back-traces"
    clockInterfaceSelectionBackTrace.EntityData.SegmentPath = "clock-interface-selection-back-trace" + "[clock-type='" + fmt.Sprintf("%v", clockInterfaceSelectionBackTrace.ClockType) + "']" + "[port='" + fmt.Sprintf("%v", clockInterfaceSelectionBackTrace.Port) + "']"
    clockInterfaceSelectionBackTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockInterfaceSelectionBackTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockInterfaceSelectionBackTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockInterfaceSelectionBackTrace.EntityData.Children = make(map[string]types.YChild)
    clockInterfaceSelectionBackTrace.EntityData.Children["selected-source"] = types.YChild{"SelectedSource", &clockInterfaceSelectionBackTrace.SelectedSource}
    clockInterfaceSelectionBackTrace.EntityData.Children["selection-point"] = types.YChild{"SelectionPoint", nil}
    for i := range clockInterfaceSelectionBackTrace.SelectionPoint {
        clockInterfaceSelectionBackTrace.EntityData.Children[types.GetSegmentPath(&clockInterfaceSelectionBackTrace.SelectionPoint[i])] = types.YChild{"SelectionPoint", &clockInterfaceSelectionBackTrace.SelectionPoint[i]}
    }
    clockInterfaceSelectionBackTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    clockInterfaceSelectionBackTrace.EntityData.Leafs["clock-type"] = types.YLeaf{"ClockType", clockInterfaceSelectionBackTrace.ClockType}
    clockInterfaceSelectionBackTrace.EntityData.Leafs["port"] = types.YLeaf{"Port", clockInterfaceSelectionBackTrace.Port}
    return &(clockInterfaceSelectionBackTrace.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource
// Source which has been selected for output
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource) GetEntityData() *types.CommonEntityData {
    selectedSource.EntityData.YFilter = selectedSource.YFilter
    selectedSource.EntityData.YangName = "selected-source"
    selectedSource.EntityData.BundleName = "cisco_ios_xr"
    selectedSource.EntityData.ParentYangName = "clock-interface-selection-back-trace"
    selectedSource.EntityData.SegmentPath = "selected-source"
    selectedSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectedSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectedSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectedSource.EntityData.Children = make(map[string]types.YChild)
    selectedSource.EntityData.Children["clock-id"] = types.YChild{"ClockId", &selectedSource.ClockId}
    selectedSource.EntityData.Leafs = make(map[string]types.YLeaf)
    selectedSource.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", selectedSource.SourceClass}
    selectedSource.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", selectedSource.EthernetInterface}
    selectedSource.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", selectedSource.SonetInterface}
    selectedSource.EntityData.Leafs["node"] = types.YLeaf{"Node", selectedSource.Node}
    selectedSource.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", selectedSource.PtpNode}
    selectedSource.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", selectedSource.SatelliteAccessInterface}
    selectedSource.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", selectedSource.NtpNode}
    return &(selectedSource.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectedSource_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "selected-source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint
// List of selection points in the backtrace
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionBackTraces_ClockInterfaceSelectionBackTrace_SelectionPoint) GetEntityData() *types.CommonEntityData {
    selectionPoint.EntityData.YFilter = selectionPoint.YFilter
    selectionPoint.EntityData.YangName = "selection-point"
    selectionPoint.EntityData.BundleName = "cisco_ios_xr"
    selectionPoint.EntityData.ParentYangName = "clock-interface-selection-back-trace"
    selectionPoint.EntityData.SegmentPath = "selection-point"
    selectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPoint.EntityData.Children = make(map[string]types.YChild)
    selectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPoint.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", selectionPoint.SelectionPointType}
    selectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", selectionPoint.SelectionPointDescription}
    selectionPoint.EntityData.Leafs["node"] = types.YLeaf{"Node", selectionPoint.Node}
    return &(selectionPoint.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces
// Selection forwardtrace operational data for
// clock-interfaces
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection forwardtrace operational data for a particular clock-interface.
    // The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace.
    ClockInterfaceSelectionForwardTrace []FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace
}

func (clockInterfaceSelectionForwardTraces *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces) GetEntityData() *types.CommonEntityData {
    clockInterfaceSelectionForwardTraces.EntityData.YFilter = clockInterfaceSelectionForwardTraces.YFilter
    clockInterfaceSelectionForwardTraces.EntityData.YangName = "clock-interface-selection-forward-traces"
    clockInterfaceSelectionForwardTraces.EntityData.BundleName = "cisco_ios_xr"
    clockInterfaceSelectionForwardTraces.EntityData.ParentYangName = "global-node"
    clockInterfaceSelectionForwardTraces.EntityData.SegmentPath = "clock-interface-selection-forward-traces"
    clockInterfaceSelectionForwardTraces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockInterfaceSelectionForwardTraces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockInterfaceSelectionForwardTraces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockInterfaceSelectionForwardTraces.EntityData.Children = make(map[string]types.YChild)
    clockInterfaceSelectionForwardTraces.EntityData.Children["clock-interface-selection-forward-trace"] = types.YChild{"ClockInterfaceSelectionForwardTrace", nil}
    for i := range clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace {
        clockInterfaceSelectionForwardTraces.EntityData.Children[types.GetSegmentPath(&clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace[i])] = types.YChild{"ClockInterfaceSelectionForwardTrace", &clockInterfaceSelectionForwardTraces.ClockInterfaceSelectionForwardTrace[i]}
    }
    clockInterfaceSelectionForwardTraces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clockInterfaceSelectionForwardTraces.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace
// Selection forwardtrace operational data for a
// particular clock-interface
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Clock type. The type is FsyncClock.
    ClockType interface{}

    // This attribute is a key. Clock port. The type is interface{} with range:
    // -2147483648..2147483647.
    Port interface{}

    // Selection ForwardTrace. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace.
    ForwardTrace []FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace
}

func (clockInterfaceSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace) GetEntityData() *types.CommonEntityData {
    clockInterfaceSelectionForwardTrace.EntityData.YFilter = clockInterfaceSelectionForwardTrace.YFilter
    clockInterfaceSelectionForwardTrace.EntityData.YangName = "clock-interface-selection-forward-trace"
    clockInterfaceSelectionForwardTrace.EntityData.BundleName = "cisco_ios_xr"
    clockInterfaceSelectionForwardTrace.EntityData.ParentYangName = "clock-interface-selection-forward-traces"
    clockInterfaceSelectionForwardTrace.EntityData.SegmentPath = "clock-interface-selection-forward-trace" + "[clock-type='" + fmt.Sprintf("%v", clockInterfaceSelectionForwardTrace.ClockType) + "']" + "[port='" + fmt.Sprintf("%v", clockInterfaceSelectionForwardTrace.Port) + "']"
    clockInterfaceSelectionForwardTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockInterfaceSelectionForwardTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockInterfaceSelectionForwardTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockInterfaceSelectionForwardTrace.EntityData.Children = make(map[string]types.YChild)
    clockInterfaceSelectionForwardTrace.EntityData.Children["forward-trace"] = types.YChild{"ForwardTrace", nil}
    for i := range clockInterfaceSelectionForwardTrace.ForwardTrace {
        clockInterfaceSelectionForwardTrace.EntityData.Children[types.GetSegmentPath(&clockInterfaceSelectionForwardTrace.ForwardTrace[i])] = types.YChild{"ForwardTrace", &clockInterfaceSelectionForwardTrace.ForwardTrace[i]}
    }
    clockInterfaceSelectionForwardTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    clockInterfaceSelectionForwardTrace.EntityData.Leafs["clock-type"] = types.YLeaf{"ClockType", clockInterfaceSelectionForwardTrace.ClockType}
    clockInterfaceSelectionForwardTrace.EntityData.Leafs["port"] = types.YLeaf{"Port", clockInterfaceSelectionForwardTrace.Port}
    return &(clockInterfaceSelectionForwardTrace.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace
// Selection ForwardTrace
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The source or selection point at this point in the forwardtrace.
    ForwardTraceNode FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace) GetEntityData() *types.CommonEntityData {
    forwardTrace.EntityData.YFilter = forwardTrace.YFilter
    forwardTrace.EntityData.YangName = "forward-trace"
    forwardTrace.EntityData.BundleName = "cisco_ios_xr"
    forwardTrace.EntityData.ParentYangName = "clock-interface-selection-forward-trace"
    forwardTrace.EntityData.SegmentPath = "forward-trace"
    forwardTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardTrace.EntityData.Children = make(map[string]types.YChild)
    forwardTrace.EntityData.Children["forward-trace-node"] = types.YChild{"ForwardTraceNode", &forwardTrace.ForwardTraceNode}
    forwardTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(forwardTrace.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode
// The source or selection point at this point in
// the forwardtrace
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NodeType. The type is FsyncBagForwardtraceNode.
    NodeType interface{}

    // Selection Point.
    SelectionPoint FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint

    // Timing Source.
    Source FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetEntityData() *types.CommonEntityData {
    forwardTraceNode.EntityData.YFilter = forwardTraceNode.YFilter
    forwardTraceNode.EntityData.YangName = "forward-trace-node"
    forwardTraceNode.EntityData.BundleName = "cisco_ios_xr"
    forwardTraceNode.EntityData.ParentYangName = "forward-trace"
    forwardTraceNode.EntityData.SegmentPath = "forward-trace-node"
    forwardTraceNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardTraceNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardTraceNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardTraceNode.EntityData.Children = make(map[string]types.YChild)
    forwardTraceNode.EntityData.Children["selection-point"] = types.YChild{"SelectionPoint", &forwardTraceNode.SelectionPoint}
    forwardTraceNode.EntityData.Children["source"] = types.YChild{"Source", &forwardTraceNode.Source}
    forwardTraceNode.EntityData.Leafs = make(map[string]types.YLeaf)
    forwardTraceNode.EntityData.Leafs["node-type"] = types.YLeaf{"NodeType", forwardTraceNode.NodeType}
    return &(forwardTraceNode.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint
// Selection Point
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetEntityData() *types.CommonEntityData {
    selectionPoint.EntityData.YFilter = selectionPoint.YFilter
    selectionPoint.EntityData.YangName = "selection-point"
    selectionPoint.EntityData.BundleName = "cisco_ios_xr"
    selectionPoint.EntityData.ParentYangName = "forward-trace-node"
    selectionPoint.EntityData.SegmentPath = "selection-point"
    selectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPoint.EntityData.Children = make(map[string]types.YChild)
    selectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPoint.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", selectionPoint.SelectionPointType}
    selectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", selectionPoint.SelectionPointDescription}
    selectionPoint.EntityData.Leafs["node"] = types.YLeaf{"Node", selectionPoint.Node}
    return &(selectionPoint.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
// Timing Source
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "forward-trace-node"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Children["clock-id"] = types.YChild{"ClockId", &source.ClockId}
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", source.SourceClass}
    source.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", source.EthernetInterface}
    source.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", source.SonetInterface}
    source.EntityData.Leafs["node"] = types.YLeaf{"Node", source.Node}
    source.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", source.PtpNode}
    source.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", source.SatelliteAccessInterface}
    source.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", source.NtpNode}
    return &(source.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
// Clock ID
type FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_ClockInterfaceSelectionForwardTraces_ClockInterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace
// Selection backtrace operational data for
// time-of-day on a particular node
type FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source which has been selected for output.
    SelectedSource FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource

    // List of selection points in the backtrace. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint.
    SelectionPoint []FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint
}

func (timeOfDayBackTrace *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace) GetEntityData() *types.CommonEntityData {
    timeOfDayBackTrace.EntityData.YFilter = timeOfDayBackTrace.YFilter
    timeOfDayBackTrace.EntityData.YangName = "time-of-day-back-trace"
    timeOfDayBackTrace.EntityData.BundleName = "cisco_ios_xr"
    timeOfDayBackTrace.EntityData.ParentYangName = "global-node"
    timeOfDayBackTrace.EntityData.SegmentPath = "time-of-day-back-trace"
    timeOfDayBackTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeOfDayBackTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeOfDayBackTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeOfDayBackTrace.EntityData.Children = make(map[string]types.YChild)
    timeOfDayBackTrace.EntityData.Children["selected-source"] = types.YChild{"SelectedSource", &timeOfDayBackTrace.SelectedSource}
    timeOfDayBackTrace.EntityData.Children["selection-point"] = types.YChild{"SelectionPoint", nil}
    for i := range timeOfDayBackTrace.SelectionPoint {
        timeOfDayBackTrace.EntityData.Children[types.GetSegmentPath(&timeOfDayBackTrace.SelectionPoint[i])] = types.YChild{"SelectionPoint", &timeOfDayBackTrace.SelectionPoint[i]}
    }
    timeOfDayBackTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(timeOfDayBackTrace.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource
// Source which has been selected for output
type FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource) GetEntityData() *types.CommonEntityData {
    selectedSource.EntityData.YFilter = selectedSource.YFilter
    selectedSource.EntityData.YangName = "selected-source"
    selectedSource.EntityData.BundleName = "cisco_ios_xr"
    selectedSource.EntityData.ParentYangName = "time-of-day-back-trace"
    selectedSource.EntityData.SegmentPath = "selected-source"
    selectedSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectedSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectedSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectedSource.EntityData.Children = make(map[string]types.YChild)
    selectedSource.EntityData.Children["clock-id"] = types.YChild{"ClockId", &selectedSource.ClockId}
    selectedSource.EntityData.Leafs = make(map[string]types.YLeaf)
    selectedSource.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", selectedSource.SourceClass}
    selectedSource.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", selectedSource.EthernetInterface}
    selectedSource.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", selectedSource.SonetInterface}
    selectedSource.EntityData.Leafs["node"] = types.YLeaf{"Node", selectedSource.Node}
    selectedSource.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", selectedSource.PtpNode}
    selectedSource.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", selectedSource.SatelliteAccessInterface}
    selectedSource.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", selectedSource.NtpNode}
    return &(selectedSource.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectedSource_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "selected-source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint
// List of selection points in the backtrace
type FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_TimeOfDayBackTrace_SelectionPoint) GetEntityData() *types.CommonEntityData {
    selectionPoint.EntityData.YFilter = selectionPoint.YFilter
    selectionPoint.EntityData.YangName = "selection-point"
    selectionPoint.EntityData.BundleName = "cisco_ios_xr"
    selectionPoint.EntityData.ParentYangName = "time-of-day-back-trace"
    selectionPoint.EntityData.SegmentPath = "selection-point"
    selectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPoint.EntityData.Children = make(map[string]types.YChild)
    selectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPoint.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", selectionPoint.SelectionPointType}
    selectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", selectionPoint.SelectionPointDescription}
    selectionPoint.EntityData.Leafs["node"] = types.YLeaf{"Node", selectionPoint.Node}
    return &(selectionPoint.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace
// Selection forwardtrace operational data for a
// NTP clock
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection ForwardTrace. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace.
    ForwardTrace []FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace
}

func (ntpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace) GetEntityData() *types.CommonEntityData {
    ntpSelectionForwardTrace.EntityData.YFilter = ntpSelectionForwardTrace.YFilter
    ntpSelectionForwardTrace.EntityData.YangName = "ntp-selection-forward-trace"
    ntpSelectionForwardTrace.EntityData.BundleName = "cisco_ios_xr"
    ntpSelectionForwardTrace.EntityData.ParentYangName = "global-node"
    ntpSelectionForwardTrace.EntityData.SegmentPath = "ntp-selection-forward-trace"
    ntpSelectionForwardTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ntpSelectionForwardTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ntpSelectionForwardTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ntpSelectionForwardTrace.EntityData.Children = make(map[string]types.YChild)
    ntpSelectionForwardTrace.EntityData.Children["forward-trace"] = types.YChild{"ForwardTrace", nil}
    for i := range ntpSelectionForwardTrace.ForwardTrace {
        ntpSelectionForwardTrace.EntityData.Children[types.GetSegmentPath(&ntpSelectionForwardTrace.ForwardTrace[i])] = types.YChild{"ForwardTrace", &ntpSelectionForwardTrace.ForwardTrace[i]}
    }
    ntpSelectionForwardTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ntpSelectionForwardTrace.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace
// Selection ForwardTrace
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The source or selection point at this point in the forwardtrace.
    ForwardTraceNode FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace) GetEntityData() *types.CommonEntityData {
    forwardTrace.EntityData.YFilter = forwardTrace.YFilter
    forwardTrace.EntityData.YangName = "forward-trace"
    forwardTrace.EntityData.BundleName = "cisco_ios_xr"
    forwardTrace.EntityData.ParentYangName = "ntp-selection-forward-trace"
    forwardTrace.EntityData.SegmentPath = "forward-trace"
    forwardTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardTrace.EntityData.Children = make(map[string]types.YChild)
    forwardTrace.EntityData.Children["forward-trace-node"] = types.YChild{"ForwardTraceNode", &forwardTrace.ForwardTraceNode}
    forwardTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(forwardTrace.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode
// The source or selection point at this point in
// the forwardtrace
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NodeType. The type is FsyncBagForwardtraceNode.
    NodeType interface{}

    // Selection Point.
    SelectionPoint FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint

    // Timing Source.
    Source FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetEntityData() *types.CommonEntityData {
    forwardTraceNode.EntityData.YFilter = forwardTraceNode.YFilter
    forwardTraceNode.EntityData.YangName = "forward-trace-node"
    forwardTraceNode.EntityData.BundleName = "cisco_ios_xr"
    forwardTraceNode.EntityData.ParentYangName = "forward-trace"
    forwardTraceNode.EntityData.SegmentPath = "forward-trace-node"
    forwardTraceNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardTraceNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardTraceNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardTraceNode.EntityData.Children = make(map[string]types.YChild)
    forwardTraceNode.EntityData.Children["selection-point"] = types.YChild{"SelectionPoint", &forwardTraceNode.SelectionPoint}
    forwardTraceNode.EntityData.Children["source"] = types.YChild{"Source", &forwardTraceNode.Source}
    forwardTraceNode.EntityData.Leafs = make(map[string]types.YLeaf)
    forwardTraceNode.EntityData.Leafs["node-type"] = types.YLeaf{"NodeType", forwardTraceNode.NodeType}
    return &(forwardTraceNode.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint
// Selection Point
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetEntityData() *types.CommonEntityData {
    selectionPoint.EntityData.YFilter = selectionPoint.YFilter
    selectionPoint.EntityData.YangName = "selection-point"
    selectionPoint.EntityData.BundleName = "cisco_ios_xr"
    selectionPoint.EntityData.ParentYangName = "forward-trace-node"
    selectionPoint.EntityData.SegmentPath = "selection-point"
    selectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPoint.EntityData.Children = make(map[string]types.YChild)
    selectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPoint.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", selectionPoint.SelectionPointType}
    selectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", selectionPoint.SelectionPointDescription}
    selectionPoint.EntityData.Leafs["node"] = types.YLeaf{"Node", selectionPoint.Node}
    return &(selectionPoint.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
// Timing Source
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "forward-trace-node"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Children["clock-id"] = types.YChild{"ClockId", &source.ClockId}
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", source.SourceClass}
    source.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", source.EthernetInterface}
    source.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", source.SonetInterface}
    source.EntityData.Leafs["node"] = types.YLeaf{"Node", source.Node}
    source.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", source.PtpNode}
    source.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", source.SatelliteAccessInterface}
    source.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", source.NtpNode}
    return &(source.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
// Clock ID
type FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_NtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace
// Selection forwardtrace operational data for a
// PTP clock
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection ForwardTrace. The type is slice of
    // FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace.
    ForwardTrace []FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace
}

func (ptpSelectionForwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace) GetEntityData() *types.CommonEntityData {
    ptpSelectionForwardTrace.EntityData.YFilter = ptpSelectionForwardTrace.YFilter
    ptpSelectionForwardTrace.EntityData.YangName = "ptp-selection-forward-trace"
    ptpSelectionForwardTrace.EntityData.BundleName = "cisco_ios_xr"
    ptpSelectionForwardTrace.EntityData.ParentYangName = "global-node"
    ptpSelectionForwardTrace.EntityData.SegmentPath = "ptp-selection-forward-trace"
    ptpSelectionForwardTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ptpSelectionForwardTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ptpSelectionForwardTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ptpSelectionForwardTrace.EntityData.Children = make(map[string]types.YChild)
    ptpSelectionForwardTrace.EntityData.Children["forward-trace"] = types.YChild{"ForwardTrace", nil}
    for i := range ptpSelectionForwardTrace.ForwardTrace {
        ptpSelectionForwardTrace.EntityData.Children[types.GetSegmentPath(&ptpSelectionForwardTrace.ForwardTrace[i])] = types.YChild{"ForwardTrace", &ptpSelectionForwardTrace.ForwardTrace[i]}
    }
    ptpSelectionForwardTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ptpSelectionForwardTrace.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace
// Selection ForwardTrace
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The source or selection point at this point in the forwardtrace.
    ForwardTraceNode FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode
}

func (forwardTrace *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace) GetEntityData() *types.CommonEntityData {
    forwardTrace.EntityData.YFilter = forwardTrace.YFilter
    forwardTrace.EntityData.YangName = "forward-trace"
    forwardTrace.EntityData.BundleName = "cisco_ios_xr"
    forwardTrace.EntityData.ParentYangName = "ptp-selection-forward-trace"
    forwardTrace.EntityData.SegmentPath = "forward-trace"
    forwardTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardTrace.EntityData.Children = make(map[string]types.YChild)
    forwardTrace.EntityData.Children["forward-trace-node"] = types.YChild{"ForwardTraceNode", &forwardTrace.ForwardTraceNode}
    forwardTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(forwardTrace.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode
// The source or selection point at this point in
// the forwardtrace
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NodeType. The type is FsyncBagForwardtraceNode.
    NodeType interface{}

    // Selection Point.
    SelectionPoint FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint

    // Timing Source.
    Source FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
}

func (forwardTraceNode *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetEntityData() *types.CommonEntityData {
    forwardTraceNode.EntityData.YFilter = forwardTraceNode.YFilter
    forwardTraceNode.EntityData.YangName = "forward-trace-node"
    forwardTraceNode.EntityData.BundleName = "cisco_ios_xr"
    forwardTraceNode.EntityData.ParentYangName = "forward-trace"
    forwardTraceNode.EntityData.SegmentPath = "forward-trace-node"
    forwardTraceNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardTraceNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardTraceNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardTraceNode.EntityData.Children = make(map[string]types.YChild)
    forwardTraceNode.EntityData.Children["selection-point"] = types.YChild{"SelectionPoint", &forwardTraceNode.SelectionPoint}
    forwardTraceNode.EntityData.Children["source"] = types.YChild{"Source", &forwardTraceNode.Source}
    forwardTraceNode.EntityData.Leafs = make(map[string]types.YLeaf)
    forwardTraceNode.EntityData.Leafs["node-type"] = types.YLeaf{"NodeType", forwardTraceNode.NodeType}
    return &(forwardTraceNode.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint
// Selection Point
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetEntityData() *types.CommonEntityData {
    selectionPoint.EntityData.YFilter = selectionPoint.YFilter
    selectionPoint.EntityData.YangName = "selection-point"
    selectionPoint.EntityData.BundleName = "cisco_ios_xr"
    selectionPoint.EntityData.ParentYangName = "forward-trace-node"
    selectionPoint.EntityData.SegmentPath = "selection-point"
    selectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPoint.EntityData.Children = make(map[string]types.YChild)
    selectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPoint.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", selectionPoint.SelectionPointType}
    selectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", selectionPoint.SelectionPointDescription}
    selectionPoint.EntityData.Leafs["node"] = types.YLeaf{"Node", selectionPoint.Node}
    return &(selectionPoint.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
// Timing Source
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
}

func (source *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "forward-trace-node"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Children["clock-id"] = types.YChild{"ClockId", &source.ClockId}
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", source.SourceClass}
    source.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", source.EthernetInterface}
    source.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", source.SonetInterface}
    source.EntityData.Leafs["node"] = types.YLeaf{"Node", source.Node}
    source.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", source.PtpNode}
    source.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", source.SatelliteAccessInterface}
    source.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", source.NtpNode}
    return &(source.EntityData)
}

// FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
// Clock ID
type FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalNodes_GlobalNode_PtpSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_GlobalInterfaces
// Table for global interface operational data
type FrequencySynchronization_GlobalInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global interface information for a particular interface. The type is slice
    // of FrequencySynchronization_GlobalInterfaces_GlobalInterface.
    GlobalInterface []FrequencySynchronization_GlobalInterfaces_GlobalInterface
}

func (globalInterfaces *FrequencySynchronization_GlobalInterfaces) GetEntityData() *types.CommonEntityData {
    globalInterfaces.EntityData.YFilter = globalInterfaces.YFilter
    globalInterfaces.EntityData.YangName = "global-interfaces"
    globalInterfaces.EntityData.BundleName = "cisco_ios_xr"
    globalInterfaces.EntityData.ParentYangName = "frequency-synchronization"
    globalInterfaces.EntityData.SegmentPath = "global-interfaces"
    globalInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalInterfaces.EntityData.Children = make(map[string]types.YChild)
    globalInterfaces.EntityData.Children["global-interface"] = types.YChild{"GlobalInterface", nil}
    for i := range globalInterfaces.GlobalInterface {
        globalInterfaces.EntityData.Children[types.GetSegmentPath(&globalInterfaces.GlobalInterface[i])] = types.YChild{"GlobalInterface", &globalInterfaces.GlobalInterface[i]}
    }
    globalInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalInterfaces.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface
// Global interface information for a particular
// interface
type FrequencySynchronization_GlobalInterfaces_GlobalInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Selection forwardtrace operational data for a particular interface.
    InterfaceSelectionForwardTrace FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace

    // Selection backtrace operational data for a particular interface.
    InterfaceSelectionBackTrace FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace
}

func (globalInterface *FrequencySynchronization_GlobalInterfaces_GlobalInterface) GetEntityData() *types.CommonEntityData {
    globalInterface.EntityData.YFilter = globalInterface.YFilter
    globalInterface.EntityData.YangName = "global-interface"
    globalInterface.EntityData.BundleName = "cisco_ios_xr"
    globalInterface.EntityData.ParentYangName = "global-interfaces"
    globalInterface.EntityData.SegmentPath = "global-interface" + "[interface-name='" + fmt.Sprintf("%v", globalInterface.InterfaceName) + "']"
    globalInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalInterface.EntityData.Children = make(map[string]types.YChild)
    globalInterface.EntityData.Children["interface-selection-forward-trace"] = types.YChild{"InterfaceSelectionForwardTrace", &globalInterface.InterfaceSelectionForwardTrace}
    globalInterface.EntityData.Children["interface-selection-back-trace"] = types.YChild{"InterfaceSelectionBackTrace", &globalInterface.InterfaceSelectionBackTrace}
    globalInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    globalInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", globalInterface.InterfaceName}
    return &(globalInterface.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace
// Selection forwardtrace operational data for a
// particular interface
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection ForwardTrace. The type is slice of
    // FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace.
    ForwardTrace []FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace
}

func (interfaceSelectionForwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace) GetEntityData() *types.CommonEntityData {
    interfaceSelectionForwardTrace.EntityData.YFilter = interfaceSelectionForwardTrace.YFilter
    interfaceSelectionForwardTrace.EntityData.YangName = "interface-selection-forward-trace"
    interfaceSelectionForwardTrace.EntityData.BundleName = "cisco_ios_xr"
    interfaceSelectionForwardTrace.EntityData.ParentYangName = "global-interface"
    interfaceSelectionForwardTrace.EntityData.SegmentPath = "interface-selection-forward-trace"
    interfaceSelectionForwardTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSelectionForwardTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSelectionForwardTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSelectionForwardTrace.EntityData.Children = make(map[string]types.YChild)
    interfaceSelectionForwardTrace.EntityData.Children["forward-trace"] = types.YChild{"ForwardTrace", nil}
    for i := range interfaceSelectionForwardTrace.ForwardTrace {
        interfaceSelectionForwardTrace.EntityData.Children[types.GetSegmentPath(&interfaceSelectionForwardTrace.ForwardTrace[i])] = types.YChild{"ForwardTrace", &interfaceSelectionForwardTrace.ForwardTrace[i]}
    }
    interfaceSelectionForwardTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceSelectionForwardTrace.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace
// Selection ForwardTrace
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The source or selection point at this point in the forwardtrace.
    ForwardTraceNode FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode
}

func (forwardTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace) GetEntityData() *types.CommonEntityData {
    forwardTrace.EntityData.YFilter = forwardTrace.YFilter
    forwardTrace.EntityData.YangName = "forward-trace"
    forwardTrace.EntityData.BundleName = "cisco_ios_xr"
    forwardTrace.EntityData.ParentYangName = "interface-selection-forward-trace"
    forwardTrace.EntityData.SegmentPath = "forward-trace"
    forwardTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardTrace.EntityData.Children = make(map[string]types.YChild)
    forwardTrace.EntityData.Children["forward-trace-node"] = types.YChild{"ForwardTraceNode", &forwardTrace.ForwardTraceNode}
    forwardTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(forwardTrace.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode
// The source or selection point at this point in
// the forwardtrace
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NodeType. The type is FsyncBagForwardtraceNode.
    NodeType interface{}

    // Selection Point.
    SelectionPoint FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint

    // Timing Source.
    Source FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
}

func (forwardTraceNode *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode) GetEntityData() *types.CommonEntityData {
    forwardTraceNode.EntityData.YFilter = forwardTraceNode.YFilter
    forwardTraceNode.EntityData.YangName = "forward-trace-node"
    forwardTraceNode.EntityData.BundleName = "cisco_ios_xr"
    forwardTraceNode.EntityData.ParentYangName = "forward-trace"
    forwardTraceNode.EntityData.SegmentPath = "forward-trace-node"
    forwardTraceNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardTraceNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardTraceNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardTraceNode.EntityData.Children = make(map[string]types.YChild)
    forwardTraceNode.EntityData.Children["selection-point"] = types.YChild{"SelectionPoint", &forwardTraceNode.SelectionPoint}
    forwardTraceNode.EntityData.Children["source"] = types.YChild{"Source", &forwardTraceNode.Source}
    forwardTraceNode.EntityData.Leafs = make(map[string]types.YLeaf)
    forwardTraceNode.EntityData.Leafs["node-type"] = types.YLeaf{"NodeType", forwardTraceNode.NodeType}
    return &(forwardTraceNode.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint
// Selection Point
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_SelectionPoint) GetEntityData() *types.CommonEntityData {
    selectionPoint.EntityData.YFilter = selectionPoint.YFilter
    selectionPoint.EntityData.YangName = "selection-point"
    selectionPoint.EntityData.BundleName = "cisco_ios_xr"
    selectionPoint.EntityData.ParentYangName = "forward-trace-node"
    selectionPoint.EntityData.SegmentPath = "selection-point"
    selectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPoint.EntityData.Children = make(map[string]types.YChild)
    selectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPoint.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", selectionPoint.SelectionPointType}
    selectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", selectionPoint.SelectionPointDescription}
    selectionPoint.EntityData.Leafs["node"] = types.YLeaf{"Node", selectionPoint.Node}
    return &(selectionPoint.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source
// Timing Source
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
}

func (source *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "forward-trace-node"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Children["clock-id"] = types.YChild{"ClockId", &source.ClockId}
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", source.SourceClass}
    source.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", source.EthernetInterface}
    source.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", source.SonetInterface}
    source.EntityData.Leafs["node"] = types.YLeaf{"Node", source.Node}
    source.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", source.PtpNode}
    source.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", source.SatelliteAccessInterface}
    source.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", source.NtpNode}
    return &(source.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId
// Clock ID
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionForwardTrace_ForwardTrace_ForwardTraceNode_Source_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace
// Selection backtrace operational data for a
// particular interface
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source which has been selected for output.
    SelectedSource FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource

    // List of selection points in the backtrace. The type is slice of
    // FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint.
    SelectionPoint []FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint
}

func (interfaceSelectionBackTrace *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace) GetEntityData() *types.CommonEntityData {
    interfaceSelectionBackTrace.EntityData.YFilter = interfaceSelectionBackTrace.YFilter
    interfaceSelectionBackTrace.EntityData.YangName = "interface-selection-back-trace"
    interfaceSelectionBackTrace.EntityData.BundleName = "cisco_ios_xr"
    interfaceSelectionBackTrace.EntityData.ParentYangName = "global-interface"
    interfaceSelectionBackTrace.EntityData.SegmentPath = "interface-selection-back-trace"
    interfaceSelectionBackTrace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSelectionBackTrace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSelectionBackTrace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSelectionBackTrace.EntityData.Children = make(map[string]types.YChild)
    interfaceSelectionBackTrace.EntityData.Children["selected-source"] = types.YChild{"SelectedSource", &interfaceSelectionBackTrace.SelectedSource}
    interfaceSelectionBackTrace.EntityData.Children["selection-point"] = types.YChild{"SelectionPoint", nil}
    for i := range interfaceSelectionBackTrace.SelectionPoint {
        interfaceSelectionBackTrace.EntityData.Children[types.GetSegmentPath(&interfaceSelectionBackTrace.SelectionPoint[i])] = types.YChild{"SelectionPoint", &interfaceSelectionBackTrace.SelectionPoint[i]}
    }
    interfaceSelectionBackTrace.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceSelectionBackTrace.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource
// Source which has been selected for output
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource) GetEntityData() *types.CommonEntityData {
    selectedSource.EntityData.YFilter = selectedSource.YFilter
    selectedSource.EntityData.YangName = "selected-source"
    selectedSource.EntityData.BundleName = "cisco_ios_xr"
    selectedSource.EntityData.ParentYangName = "interface-selection-back-trace"
    selectedSource.EntityData.SegmentPath = "selected-source"
    selectedSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectedSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectedSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectedSource.EntityData.Children = make(map[string]types.YChild)
    selectedSource.EntityData.Children["clock-id"] = types.YChild{"ClockId", &selectedSource.ClockId}
    selectedSource.EntityData.Leafs = make(map[string]types.YLeaf)
    selectedSource.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", selectedSource.SourceClass}
    selectedSource.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", selectedSource.EthernetInterface}
    selectedSource.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", selectedSource.SonetInterface}
    selectedSource.EntityData.Leafs["node"] = types.YLeaf{"Node", selectedSource.Node}
    selectedSource.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", selectedSource.PtpNode}
    selectedSource.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", selectedSource.SatelliteAccessInterface}
    selectedSource.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", selectedSource.NtpNode}
    return &(selectedSource.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectedSource_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "selected-source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint
// List of selection points in the backtrace
type FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_GlobalInterfaces_GlobalInterface_InterfaceSelectionBackTrace_SelectionPoint) GetEntityData() *types.CommonEntityData {
    selectionPoint.EntityData.YFilter = selectionPoint.YFilter
    selectionPoint.EntityData.YangName = "selection-point"
    selectionPoint.EntityData.BundleName = "cisco_ios_xr"
    selectionPoint.EntityData.ParentYangName = "interface-selection-back-trace"
    selectionPoint.EntityData.SegmentPath = "selection-point"
    selectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPoint.EntityData.Children = make(map[string]types.YChild)
    selectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPoint.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", selectionPoint.SelectionPointType}
    selectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", selectionPoint.SelectionPointDescription}
    selectionPoint.EntityData.Leafs["node"] = types.YLeaf{"Node", selectionPoint.Node}
    return &(selectionPoint.EntityData)
}

// FrequencySynchronization_Summary
// Summary operational data
type FrequencySynchronization_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of sources selected for frequency. The type is slice of
    // FrequencySynchronization_Summary_FrequencySummary.
    FrequencySummary []FrequencySynchronization_Summary_FrequencySummary

    // Summary of sources selected for time-of-day. The type is slice of
    // FrequencySynchronization_Summary_TimeOfDaySummary.
    TimeOfDaySummary []FrequencySynchronization_Summary_TimeOfDaySummary
}

func (summary *FrequencySynchronization_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "frequency-synchronization"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["frequency-summary"] = types.YChild{"FrequencySummary", nil}
    for i := range summary.FrequencySummary {
        summary.EntityData.Children[types.GetSegmentPath(&summary.FrequencySummary[i])] = types.YChild{"FrequencySummary", &summary.FrequencySummary[i]}
    }
    summary.EntityData.Children["time-of-day-summary"] = types.YChild{"TimeOfDaySummary", nil}
    for i := range summary.TimeOfDaySummary {
        summary.EntityData.Children[types.GetSegmentPath(&summary.TimeOfDaySummary[i])] = types.YChild{"TimeOfDaySummary", &summary.TimeOfDaySummary[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summary.EntityData)
}

// FrequencySynchronization_Summary_FrequencySummary
// Summary of sources selected for frequency
type FrequencySynchronization_Summary_FrequencySummary struct {
    EntityData types.CommonEntityData
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

func (frequencySummary *FrequencySynchronization_Summary_FrequencySummary) GetEntityData() *types.CommonEntityData {
    frequencySummary.EntityData.YFilter = frequencySummary.YFilter
    frequencySummary.EntityData.YangName = "frequency-summary"
    frequencySummary.EntityData.BundleName = "cisco_ios_xr"
    frequencySummary.EntityData.ParentYangName = "summary"
    frequencySummary.EntityData.SegmentPath = "frequency-summary"
    frequencySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frequencySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frequencySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frequencySummary.EntityData.Children = make(map[string]types.YChild)
    frequencySummary.EntityData.Children["source"] = types.YChild{"Source", &frequencySummary.Source}
    frequencySummary.EntityData.Leafs = make(map[string]types.YLeaf)
    frequencySummary.EntityData.Leafs["clock-count"] = types.YLeaf{"ClockCount", frequencySummary.ClockCount}
    frequencySummary.EntityData.Leafs["ethernet-count"] = types.YLeaf{"EthernetCount", frequencySummary.EthernetCount}
    frequencySummary.EntityData.Leafs["sonet-count"] = types.YLeaf{"SonetCount", frequencySummary.SonetCount}
    return &(frequencySummary.EntityData)
}

// FrequencySynchronization_Summary_FrequencySummary_Source
// The source associated with this summary
// information
type FrequencySynchronization_Summary_FrequencySummary_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Summary_FrequencySummary_Source_ClockId
}

func (source *FrequencySynchronization_Summary_FrequencySummary_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "frequency-summary"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Children["clock-id"] = types.YChild{"ClockId", &source.ClockId}
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", source.SourceClass}
    source.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", source.EthernetInterface}
    source.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", source.SonetInterface}
    source.EntityData.Leafs["node"] = types.YLeaf{"Node", source.Node}
    source.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", source.PtpNode}
    source.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", source.SatelliteAccessInterface}
    source.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", source.NtpNode}
    return &(source.EntityData)
}

// FrequencySynchronization_Summary_FrequencySummary_Source_ClockId
// Clock ID
type FrequencySynchronization_Summary_FrequencySummary_Source_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Summary_FrequencySummary_Source_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_Summary_TimeOfDaySummary
// Summary of sources selected for time-of-day
type FrequencySynchronization_Summary_TimeOfDaySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of cards having their time-of-day set by the source. The type is
    // interface{} with range: 0..4294967295.
    NodeCount interface{}

    // The source associated with this summary information.
    Source FrequencySynchronization_Summary_TimeOfDaySummary_Source
}

func (timeOfDaySummary *FrequencySynchronization_Summary_TimeOfDaySummary) GetEntityData() *types.CommonEntityData {
    timeOfDaySummary.EntityData.YFilter = timeOfDaySummary.YFilter
    timeOfDaySummary.EntityData.YangName = "time-of-day-summary"
    timeOfDaySummary.EntityData.BundleName = "cisco_ios_xr"
    timeOfDaySummary.EntityData.ParentYangName = "summary"
    timeOfDaySummary.EntityData.SegmentPath = "time-of-day-summary"
    timeOfDaySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeOfDaySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeOfDaySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeOfDaySummary.EntityData.Children = make(map[string]types.YChild)
    timeOfDaySummary.EntityData.Children["source"] = types.YChild{"Source", &timeOfDaySummary.Source}
    timeOfDaySummary.EntityData.Leafs = make(map[string]types.YLeaf)
    timeOfDaySummary.EntityData.Leafs["node-count"] = types.YLeaf{"NodeCount", timeOfDaySummary.NodeCount}
    return &(timeOfDaySummary.EntityData)
}

// FrequencySynchronization_Summary_TimeOfDaySummary_Source
// The source associated with this summary
// information
type FrequencySynchronization_Summary_TimeOfDaySummary_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId
}

func (source *FrequencySynchronization_Summary_TimeOfDaySummary_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "time-of-day-summary"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Children["clock-id"] = types.YChild{"ClockId", &source.ClockId}
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", source.SourceClass}
    source.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", source.EthernetInterface}
    source.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", source.SonetInterface}
    source.EntityData.Leafs["node"] = types.YLeaf{"Node", source.Node}
    source.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", source.PtpNode}
    source.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", source.SatelliteAccessInterface}
    source.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", source.NtpNode}
    return &(source.EntityData)
}

// FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId
// Clock ID
type FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Summary_TimeOfDaySummary_Source_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_InterfaceDatas
// Table for interface operational data
type FrequencySynchronization_InterfaceDatas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular interface. The type is slice of
    // FrequencySynchronization_InterfaceDatas_InterfaceData.
    InterfaceData []FrequencySynchronization_InterfaceDatas_InterfaceData
}

func (interfaceDatas *FrequencySynchronization_InterfaceDatas) GetEntityData() *types.CommonEntityData {
    interfaceDatas.EntityData.YFilter = interfaceDatas.YFilter
    interfaceDatas.EntityData.YangName = "interface-datas"
    interfaceDatas.EntityData.BundleName = "cisco_ios_xr"
    interfaceDatas.EntityData.ParentYangName = "frequency-synchronization"
    interfaceDatas.EntityData.SegmentPath = "interface-datas"
    interfaceDatas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceDatas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceDatas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceDatas.EntityData.Children = make(map[string]types.YChild)
    interfaceDatas.EntityData.Children["interface-data"] = types.YChild{"InterfaceData", nil}
    for i := range interfaceDatas.InterfaceData {
        interfaceDatas.EntityData.Children[types.GetSegmentPath(&interfaceDatas.InterfaceData[i])] = types.YChild{"InterfaceData", &interfaceDatas.InterfaceData[i]}
    }
    interfaceDatas.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceDatas.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData
// Operational data for a particular interface
type FrequencySynchronization_InterfaceDatas_InterfaceData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

    // The source ID for the interface.
    Source FrequencySynchronization_InterfaceDatas_InterfaceData_Source

    // Timing source selected for interface output.
    SelectedSource FrequencySynchronization_InterfaceDatas_InterfaceData_SelectedSource

    // Received quality level.
    QualityLevelReceived FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelReceived

    // Quality level after damping has been applied.
    QualityLevelDamped FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelDamped

    // The effective input quality level.
    QualityLevelEffectiveInput FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelEffectiveInput

    // The effective output quality level.
    QualityLevelEffectiveOutput FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelEffectiveOutput

    // The quality level of the source driving this interface.
    QualityLevelSelectedSource FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelSelectedSource

    // Ethernet peer information.
    EthernetPeerInformation FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation

    // ESMC Statistics.
    EsmcStatistics FrequencySynchronization_InterfaceDatas_InterfaceData_EsmcStatistics

    // Spa selection points. The type is slice of
    // FrequencySynchronization_InterfaceDatas_InterfaceData_SpaSelectionPoint.
    SpaSelectionPoint []FrequencySynchronization_InterfaceDatas_InterfaceData_SpaSelectionPoint

    // Node selection points. The type is slice of
    // FrequencySynchronization_InterfaceDatas_InterfaceData_NodeSelectionPoint.
    NodeSelectionPoint []FrequencySynchronization_InterfaceDatas_InterfaceData_NodeSelectionPoint
}

func (interfaceData *FrequencySynchronization_InterfaceDatas_InterfaceData) GetEntityData() *types.CommonEntityData {
    interfaceData.EntityData.YFilter = interfaceData.YFilter
    interfaceData.EntityData.YangName = "interface-data"
    interfaceData.EntityData.BundleName = "cisco_ios_xr"
    interfaceData.EntityData.ParentYangName = "interface-datas"
    interfaceData.EntityData.SegmentPath = "interface-data" + "[interface-name='" + fmt.Sprintf("%v", interfaceData.InterfaceName) + "']"
    interfaceData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceData.EntityData.Children = make(map[string]types.YChild)
    interfaceData.EntityData.Children["source"] = types.YChild{"Source", &interfaceData.Source}
    interfaceData.EntityData.Children["selected-source"] = types.YChild{"SelectedSource", &interfaceData.SelectedSource}
    interfaceData.EntityData.Children["quality-level-received"] = types.YChild{"QualityLevelReceived", &interfaceData.QualityLevelReceived}
    interfaceData.EntityData.Children["quality-level-damped"] = types.YChild{"QualityLevelDamped", &interfaceData.QualityLevelDamped}
    interfaceData.EntityData.Children["quality-level-effective-input"] = types.YChild{"QualityLevelEffectiveInput", &interfaceData.QualityLevelEffectiveInput}
    interfaceData.EntityData.Children["quality-level-effective-output"] = types.YChild{"QualityLevelEffectiveOutput", &interfaceData.QualityLevelEffectiveOutput}
    interfaceData.EntityData.Children["quality-level-selected-source"] = types.YChild{"QualityLevelSelectedSource", &interfaceData.QualityLevelSelectedSource}
    interfaceData.EntityData.Children["ethernet-peer-information"] = types.YChild{"EthernetPeerInformation", &interfaceData.EthernetPeerInformation}
    interfaceData.EntityData.Children["esmc-statistics"] = types.YChild{"EsmcStatistics", &interfaceData.EsmcStatistics}
    interfaceData.EntityData.Children["spa-selection-point"] = types.YChild{"SpaSelectionPoint", nil}
    for i := range interfaceData.SpaSelectionPoint {
        interfaceData.EntityData.Children[types.GetSegmentPath(&interfaceData.SpaSelectionPoint[i])] = types.YChild{"SpaSelectionPoint", &interfaceData.SpaSelectionPoint[i]}
    }
    interfaceData.EntityData.Children["node-selection-point"] = types.YChild{"NodeSelectionPoint", nil}
    for i := range interfaceData.NodeSelectionPoint {
        interfaceData.EntityData.Children[types.GetSegmentPath(&interfaceData.NodeSelectionPoint[i])] = types.YChild{"NodeSelectionPoint", &interfaceData.NodeSelectionPoint[i]}
    }
    interfaceData.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceData.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceData.InterfaceName}
    interfaceData.EntityData.Leafs["name"] = types.YLeaf{"Name", interfaceData.Name}
    interfaceData.EntityData.Leafs["state"] = types.YLeaf{"State", interfaceData.State}
    interfaceData.EntityData.Leafs["ssm-enabled"] = types.YLeaf{"SsmEnabled", interfaceData.SsmEnabled}
    interfaceData.EntityData.Leafs["squelched"] = types.YLeaf{"Squelched", interfaceData.Squelched}
    interfaceData.EntityData.Leafs["selection-input"] = types.YLeaf{"SelectionInput", interfaceData.SelectionInput}
    interfaceData.EntityData.Leafs["priority"] = types.YLeaf{"Priority", interfaceData.Priority}
    interfaceData.EntityData.Leafs["time-of-day-priority"] = types.YLeaf{"TimeOfDayPriority", interfaceData.TimeOfDayPriority}
    interfaceData.EntityData.Leafs["damping-state"] = types.YLeaf{"DampingState", interfaceData.DampingState}
    interfaceData.EntityData.Leafs["damping-time"] = types.YLeaf{"DampingTime", interfaceData.DampingTime}
    interfaceData.EntityData.Leafs["wait-to-restore-time"] = types.YLeaf{"WaitToRestoreTime", interfaceData.WaitToRestoreTime}
    interfaceData.EntityData.Leafs["supports-frequency"] = types.YLeaf{"SupportsFrequency", interfaceData.SupportsFrequency}
    interfaceData.EntityData.Leafs["supports-time-of-day"] = types.YLeaf{"SupportsTimeOfDay", interfaceData.SupportsTimeOfDay}
    return &(interfaceData.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_Source
// The source ID for the interface
type FrequencySynchronization_InterfaceDatas_InterfaceData_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_InterfaceDatas_InterfaceData_Source_ClockId
}

func (source *FrequencySynchronization_InterfaceDatas_InterfaceData_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "interface-data"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Children["clock-id"] = types.YChild{"ClockId", &source.ClockId}
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", source.SourceClass}
    source.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", source.EthernetInterface}
    source.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", source.SonetInterface}
    source.EntityData.Leafs["node"] = types.YLeaf{"Node", source.Node}
    source.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", source.PtpNode}
    source.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", source.SatelliteAccessInterface}
    source.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", source.NtpNode}
    return &(source.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_Source_ClockId
// Clock ID
type FrequencySynchronization_InterfaceDatas_InterfaceData_Source_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_InterfaceDatas_InterfaceData_Source_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_SelectedSource
// Timing source selected for interface output
type FrequencySynchronization_InterfaceDatas_InterfaceData_SelectedSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_InterfaceDatas_InterfaceData_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_InterfaceDatas_InterfaceData_SelectedSource) GetEntityData() *types.CommonEntityData {
    selectedSource.EntityData.YFilter = selectedSource.YFilter
    selectedSource.EntityData.YangName = "selected-source"
    selectedSource.EntityData.BundleName = "cisco_ios_xr"
    selectedSource.EntityData.ParentYangName = "interface-data"
    selectedSource.EntityData.SegmentPath = "selected-source"
    selectedSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectedSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectedSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectedSource.EntityData.Children = make(map[string]types.YChild)
    selectedSource.EntityData.Children["clock-id"] = types.YChild{"ClockId", &selectedSource.ClockId}
    selectedSource.EntityData.Leafs = make(map[string]types.YLeaf)
    selectedSource.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", selectedSource.SourceClass}
    selectedSource.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", selectedSource.EthernetInterface}
    selectedSource.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", selectedSource.SonetInterface}
    selectedSource.EntityData.Leafs["node"] = types.YLeaf{"Node", selectedSource.Node}
    selectedSource.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", selectedSource.PtpNode}
    selectedSource.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", selectedSource.SatelliteAccessInterface}
    selectedSource.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", selectedSource.NtpNode}
    return &(selectedSource.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_InterfaceDatas_InterfaceData_SelectedSource_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_InterfaceDatas_InterfaceData_SelectedSource_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "selected-source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelReceived
// Received quality level
type FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelReceived struct {
    EntityData types.CommonEntityData
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

func (qualityLevelReceived *FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelReceived) GetEntityData() *types.CommonEntityData {
    qualityLevelReceived.EntityData.YFilter = qualityLevelReceived.YFilter
    qualityLevelReceived.EntityData.YangName = "quality-level-received"
    qualityLevelReceived.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelReceived.EntityData.ParentYangName = "interface-data"
    qualityLevelReceived.EntityData.SegmentPath = "quality-level-received"
    qualityLevelReceived.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelReceived.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelReceived.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelReceived.EntityData.Children = make(map[string]types.YChild)
    qualityLevelReceived.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelReceived.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelReceived.QualityLevelOption}
    qualityLevelReceived.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelReceived.Option1Value}
    qualityLevelReceived.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelReceived.Option2Generation1Value}
    qualityLevelReceived.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelReceived.Option2Generation2Value}
    return &(qualityLevelReceived.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelDamped
// Quality level after damping has been applied
type FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelDamped struct {
    EntityData types.CommonEntityData
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

func (qualityLevelDamped *FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelDamped) GetEntityData() *types.CommonEntityData {
    qualityLevelDamped.EntityData.YFilter = qualityLevelDamped.YFilter
    qualityLevelDamped.EntityData.YangName = "quality-level-damped"
    qualityLevelDamped.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelDamped.EntityData.ParentYangName = "interface-data"
    qualityLevelDamped.EntityData.SegmentPath = "quality-level-damped"
    qualityLevelDamped.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelDamped.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelDamped.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelDamped.EntityData.Children = make(map[string]types.YChild)
    qualityLevelDamped.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelDamped.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelDamped.QualityLevelOption}
    qualityLevelDamped.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelDamped.Option1Value}
    qualityLevelDamped.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelDamped.Option2Generation1Value}
    qualityLevelDamped.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelDamped.Option2Generation2Value}
    return &(qualityLevelDamped.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelEffectiveInput
// The effective input quality level
type FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelEffectiveInput struct {
    EntityData types.CommonEntityData
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

func (qualityLevelEffectiveInput *FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelEffectiveInput) GetEntityData() *types.CommonEntityData {
    qualityLevelEffectiveInput.EntityData.YFilter = qualityLevelEffectiveInput.YFilter
    qualityLevelEffectiveInput.EntityData.YangName = "quality-level-effective-input"
    qualityLevelEffectiveInput.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelEffectiveInput.EntityData.ParentYangName = "interface-data"
    qualityLevelEffectiveInput.EntityData.SegmentPath = "quality-level-effective-input"
    qualityLevelEffectiveInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelEffectiveInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelEffectiveInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelEffectiveInput.EntityData.Children = make(map[string]types.YChild)
    qualityLevelEffectiveInput.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelEffectiveInput.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelEffectiveInput.QualityLevelOption}
    qualityLevelEffectiveInput.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelEffectiveInput.Option1Value}
    qualityLevelEffectiveInput.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelEffectiveInput.Option2Generation1Value}
    qualityLevelEffectiveInput.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelEffectiveInput.Option2Generation2Value}
    return &(qualityLevelEffectiveInput.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelEffectiveOutput
// The effective output quality level
type FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelEffectiveOutput struct {
    EntityData types.CommonEntityData
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

func (qualityLevelEffectiveOutput *FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelEffectiveOutput) GetEntityData() *types.CommonEntityData {
    qualityLevelEffectiveOutput.EntityData.YFilter = qualityLevelEffectiveOutput.YFilter
    qualityLevelEffectiveOutput.EntityData.YangName = "quality-level-effective-output"
    qualityLevelEffectiveOutput.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelEffectiveOutput.EntityData.ParentYangName = "interface-data"
    qualityLevelEffectiveOutput.EntityData.SegmentPath = "quality-level-effective-output"
    qualityLevelEffectiveOutput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelEffectiveOutput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelEffectiveOutput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelEffectiveOutput.EntityData.Children = make(map[string]types.YChild)
    qualityLevelEffectiveOutput.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelEffectiveOutput.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelEffectiveOutput.QualityLevelOption}
    qualityLevelEffectiveOutput.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelEffectiveOutput.Option1Value}
    qualityLevelEffectiveOutput.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelEffectiveOutput.Option2Generation1Value}
    qualityLevelEffectiveOutput.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelEffectiveOutput.Option2Generation2Value}
    return &(qualityLevelEffectiveOutput.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelSelectedSource
// The quality level of the source driving this
// interface
type FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelSelectedSource struct {
    EntityData types.CommonEntityData
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

func (qualityLevelSelectedSource *FrequencySynchronization_InterfaceDatas_InterfaceData_QualityLevelSelectedSource) GetEntityData() *types.CommonEntityData {
    qualityLevelSelectedSource.EntityData.YFilter = qualityLevelSelectedSource.YFilter
    qualityLevelSelectedSource.EntityData.YangName = "quality-level-selected-source"
    qualityLevelSelectedSource.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelSelectedSource.EntityData.ParentYangName = "interface-data"
    qualityLevelSelectedSource.EntityData.SegmentPath = "quality-level-selected-source"
    qualityLevelSelectedSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelSelectedSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelSelectedSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelSelectedSource.EntityData.Children = make(map[string]types.YChild)
    qualityLevelSelectedSource.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelSelectedSource.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelSelectedSource.QualityLevelOption}
    qualityLevelSelectedSource.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelSelectedSource.Option1Value}
    qualityLevelSelectedSource.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelSelectedSource.Option2Generation1Value}
    qualityLevelSelectedSource.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelSelectedSource.Option2Generation2Value}
    return &(qualityLevelSelectedSource.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation
// Ethernet peer information
type FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation struct {
    EntityData types.CommonEntityData
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
    PeerStateTime FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation_PeerStateTime

    // Time of last SSM received.
    LastSsm FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation_LastSsm
}

func (ethernetPeerInformation *FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation) GetEntityData() *types.CommonEntityData {
    ethernetPeerInformation.EntityData.YFilter = ethernetPeerInformation.YFilter
    ethernetPeerInformation.EntityData.YangName = "ethernet-peer-information"
    ethernetPeerInformation.EntityData.BundleName = "cisco_ios_xr"
    ethernetPeerInformation.EntityData.ParentYangName = "interface-data"
    ethernetPeerInformation.EntityData.SegmentPath = "ethernet-peer-information"
    ethernetPeerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetPeerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetPeerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetPeerInformation.EntityData.Children = make(map[string]types.YChild)
    ethernetPeerInformation.EntityData.Children["peer-state-time"] = types.YChild{"PeerStateTime", &ethernetPeerInformation.PeerStateTime}
    ethernetPeerInformation.EntityData.Children["last-ssm"] = types.YChild{"LastSsm", &ethernetPeerInformation.LastSsm}
    ethernetPeerInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetPeerInformation.EntityData.Leafs["peer-state"] = types.YLeaf{"PeerState", ethernetPeerInformation.PeerState}
    ethernetPeerInformation.EntityData.Leafs["peer-up-count"] = types.YLeaf{"PeerUpCount", ethernetPeerInformation.PeerUpCount}
    ethernetPeerInformation.EntityData.Leafs["peer-timeout-count"] = types.YLeaf{"PeerTimeoutCount", ethernetPeerInformation.PeerTimeoutCount}
    return &(ethernetPeerInformation.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation_PeerStateTime
// Time of last peer state transition
type FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation_PeerStateTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (peerStateTime *FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation_PeerStateTime) GetEntityData() *types.CommonEntityData {
    peerStateTime.EntityData.YFilter = peerStateTime.YFilter
    peerStateTime.EntityData.YangName = "peer-state-time"
    peerStateTime.EntityData.BundleName = "cisco_ios_xr"
    peerStateTime.EntityData.ParentYangName = "ethernet-peer-information"
    peerStateTime.EntityData.SegmentPath = "peer-state-time"
    peerStateTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerStateTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerStateTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerStateTime.EntityData.Children = make(map[string]types.YChild)
    peerStateTime.EntityData.Leafs = make(map[string]types.YLeaf)
    peerStateTime.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", peerStateTime.Seconds}
    peerStateTime.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", peerStateTime.Nanoseconds}
    return &(peerStateTime.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation_LastSsm
// Time of last SSM received
type FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation_LastSsm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastSsm *FrequencySynchronization_InterfaceDatas_InterfaceData_EthernetPeerInformation_LastSsm) GetEntityData() *types.CommonEntityData {
    lastSsm.EntityData.YFilter = lastSsm.YFilter
    lastSsm.EntityData.YangName = "last-ssm"
    lastSsm.EntityData.BundleName = "cisco_ios_xr"
    lastSsm.EntityData.ParentYangName = "ethernet-peer-information"
    lastSsm.EntityData.SegmentPath = "last-ssm"
    lastSsm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastSsm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastSsm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastSsm.EntityData.Children = make(map[string]types.YChild)
    lastSsm.EntityData.Leafs = make(map[string]types.YLeaf)
    lastSsm.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", lastSsm.Seconds}
    lastSsm.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", lastSsm.Nanoseconds}
    return &(lastSsm.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_EsmcStatistics
// ESMC Statistics
type FrequencySynchronization_InterfaceDatas_InterfaceData_EsmcStatistics struct {
    EntityData types.CommonEntityData
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

func (esmcStatistics *FrequencySynchronization_InterfaceDatas_InterfaceData_EsmcStatistics) GetEntityData() *types.CommonEntityData {
    esmcStatistics.EntityData.YFilter = esmcStatistics.YFilter
    esmcStatistics.EntityData.YangName = "esmc-statistics"
    esmcStatistics.EntityData.BundleName = "cisco_ios_xr"
    esmcStatistics.EntityData.ParentYangName = "interface-data"
    esmcStatistics.EntityData.SegmentPath = "esmc-statistics"
    esmcStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esmcStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esmcStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esmcStatistics.EntityData.Children = make(map[string]types.YChild)
    esmcStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    esmcStatistics.EntityData.Leafs["esmc-events-sent"] = types.YLeaf{"EsmcEventsSent", esmcStatistics.EsmcEventsSent}
    esmcStatistics.EntityData.Leafs["esmc-events-received"] = types.YLeaf{"EsmcEventsReceived", esmcStatistics.EsmcEventsReceived}
    esmcStatistics.EntityData.Leafs["esmc-infos-sent"] = types.YLeaf{"EsmcInfosSent", esmcStatistics.EsmcInfosSent}
    esmcStatistics.EntityData.Leafs["esmc-infos-received"] = types.YLeaf{"EsmcInfosReceived", esmcStatistics.EsmcInfosReceived}
    esmcStatistics.EntityData.Leafs["esmc-dn-us-sent"] = types.YLeaf{"EsmcDnUsSent", esmcStatistics.EsmcDnUsSent}
    esmcStatistics.EntityData.Leafs["esmc-dn-us-received"] = types.YLeaf{"EsmcDnUsReceived", esmcStatistics.EsmcDnUsReceived}
    esmcStatistics.EntityData.Leafs["esmc-malformed-received"] = types.YLeaf{"EsmcMalformedReceived", esmcStatistics.EsmcMalformedReceived}
    esmcStatistics.EntityData.Leafs["esmc-received-error"] = types.YLeaf{"EsmcReceivedError", esmcStatistics.EsmcReceivedError}
    return &(esmcStatistics.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_SpaSelectionPoint
// Spa selection points
type FrequencySynchronization_InterfaceDatas_InterfaceData_SpaSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (spaSelectionPoint *FrequencySynchronization_InterfaceDatas_InterfaceData_SpaSelectionPoint) GetEntityData() *types.CommonEntityData {
    spaSelectionPoint.EntityData.YFilter = spaSelectionPoint.YFilter
    spaSelectionPoint.EntityData.YangName = "spa-selection-point"
    spaSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    spaSelectionPoint.EntityData.ParentYangName = "interface-data"
    spaSelectionPoint.EntityData.SegmentPath = "spa-selection-point"
    spaSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spaSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spaSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spaSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    spaSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    spaSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", spaSelectionPoint.SelectionPoint}
    spaSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", spaSelectionPoint.SelectionPointDescription}
    return &(spaSelectionPoint.EntityData)
}

// FrequencySynchronization_InterfaceDatas_InterfaceData_NodeSelectionPoint
// Node selection points
type FrequencySynchronization_InterfaceDatas_InterfaceData_NodeSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (nodeSelectionPoint *FrequencySynchronization_InterfaceDatas_InterfaceData_NodeSelectionPoint) GetEntityData() *types.CommonEntityData {
    nodeSelectionPoint.EntityData.YFilter = nodeSelectionPoint.YFilter
    nodeSelectionPoint.EntityData.YangName = "node-selection-point"
    nodeSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    nodeSelectionPoint.EntityData.ParentYangName = "interface-data"
    nodeSelectionPoint.EntityData.SegmentPath = "node-selection-point"
    nodeSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    nodeSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", nodeSelectionPoint.SelectionPoint}
    nodeSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", nodeSelectionPoint.SelectionPointDescription}
    return &(nodeSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes
// Table for node-specific operational data
type FrequencySynchronization_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // FrequencySynchronization_Nodes_Node.
    Node []FrequencySynchronization_Nodes_Node
}

func (nodes *FrequencySynchronization_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "frequency-synchronization"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// FrequencySynchronization_Nodes_Node
// Node-specific data for a particular node
type FrequencySynchronization_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // NTP operational data.
    NtpData FrequencySynchronization_Nodes_Node_NtpData

    // Selection point data table.
    SelectionPointDatas FrequencySynchronization_Nodes_Node_SelectionPointDatas

    // Configuration error operational data.
    ConfigurationErrors FrequencySynchronization_Nodes_Node_ConfigurationErrors

    // PTP operational data.
    PtpData FrequencySynchronization_Nodes_Node_PtpData

    // SSM operational data.
    SsmSummary FrequencySynchronization_Nodes_Node_SsmSummary

    // Table for clock operational data.
    ClockDatas FrequencySynchronization_Nodes_Node_ClockDatas

    // Table for selection point input operational data.
    SelectionPointInputs FrequencySynchronization_Nodes_Node_SelectionPointInputs
}

func (node *FrequencySynchronization_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["ntp-data"] = types.YChild{"NtpData", &node.NtpData}
    node.EntityData.Children["selection-point-datas"] = types.YChild{"SelectionPointDatas", &node.SelectionPointDatas}
    node.EntityData.Children["configuration-errors"] = types.YChild{"ConfigurationErrors", &node.ConfigurationErrors}
    node.EntityData.Children["ptp-data"] = types.YChild{"PtpData", &node.PtpData}
    node.EntityData.Children["ssm-summary"] = types.YChild{"SsmSummary", &node.SsmSummary}
    node.EntityData.Children["clock-datas"] = types.YChild{"ClockDatas", &node.ClockDatas}
    node.EntityData.Children["selection-point-inputs"] = types.YChild{"SelectionPointInputs", &node.SelectionPointInputs}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node"] = types.YLeaf{"Node", node.Node}
    return &(node.EntityData)
}

// FrequencySynchronization_Nodes_Node_NtpData
// NTP operational data
type FrequencySynchronization_Nodes_Node_NtpData struct {
    EntityData types.CommonEntityData
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

    // The effective input quality level.
    QualityLevelEffectiveInput FrequencySynchronization_Nodes_Node_NtpData_QualityLevelEffectiveInput

    // Spa selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_NtpData_SpaSelectionPoint.
    SpaSelectionPoint []FrequencySynchronization_Nodes_Node_NtpData_SpaSelectionPoint

    // Node selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_NtpData_NodeSelectionPoint.
    NodeSelectionPoint []FrequencySynchronization_Nodes_Node_NtpData_NodeSelectionPoint
}

func (ntpData *FrequencySynchronization_Nodes_Node_NtpData) GetEntityData() *types.CommonEntityData {
    ntpData.EntityData.YFilter = ntpData.YFilter
    ntpData.EntityData.YangName = "ntp-data"
    ntpData.EntityData.BundleName = "cisco_ios_xr"
    ntpData.EntityData.ParentYangName = "node"
    ntpData.EntityData.SegmentPath = "ntp-data"
    ntpData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ntpData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ntpData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ntpData.EntityData.Children = make(map[string]types.YChild)
    ntpData.EntityData.Children["quality-level-effective-input"] = types.YChild{"QualityLevelEffectiveInput", &ntpData.QualityLevelEffectiveInput}
    ntpData.EntityData.Children["spa-selection-point"] = types.YChild{"SpaSelectionPoint", nil}
    for i := range ntpData.SpaSelectionPoint {
        ntpData.EntityData.Children[types.GetSegmentPath(&ntpData.SpaSelectionPoint[i])] = types.YChild{"SpaSelectionPoint", &ntpData.SpaSelectionPoint[i]}
    }
    ntpData.EntityData.Children["node-selection-point"] = types.YChild{"NodeSelectionPoint", nil}
    for i := range ntpData.NodeSelectionPoint {
        ntpData.EntityData.Children[types.GetSegmentPath(&ntpData.NodeSelectionPoint[i])] = types.YChild{"NodeSelectionPoint", &ntpData.NodeSelectionPoint[i]}
    }
    ntpData.EntityData.Leafs = make(map[string]types.YLeaf)
    ntpData.EntityData.Leafs["state"] = types.YLeaf{"State", ntpData.State}
    ntpData.EntityData.Leafs["supports-frequency"] = types.YLeaf{"SupportsFrequency", ntpData.SupportsFrequency}
    ntpData.EntityData.Leafs["supports-time-of-day"] = types.YLeaf{"SupportsTimeOfDay", ntpData.SupportsTimeOfDay}
    ntpData.EntityData.Leafs["frequency-priority"] = types.YLeaf{"FrequencyPriority", ntpData.FrequencyPriority}
    ntpData.EntityData.Leafs["time-of-day-priority"] = types.YLeaf{"TimeOfDayPriority", ntpData.TimeOfDayPriority}
    return &(ntpData.EntityData)
}

// FrequencySynchronization_Nodes_Node_NtpData_QualityLevelEffectiveInput
// The effective input quality level
type FrequencySynchronization_Nodes_Node_NtpData_QualityLevelEffectiveInput struct {
    EntityData types.CommonEntityData
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

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_NtpData_QualityLevelEffectiveInput) GetEntityData() *types.CommonEntityData {
    qualityLevelEffectiveInput.EntityData.YFilter = qualityLevelEffectiveInput.YFilter
    qualityLevelEffectiveInput.EntityData.YangName = "quality-level-effective-input"
    qualityLevelEffectiveInput.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelEffectiveInput.EntityData.ParentYangName = "ntp-data"
    qualityLevelEffectiveInput.EntityData.SegmentPath = "quality-level-effective-input"
    qualityLevelEffectiveInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelEffectiveInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelEffectiveInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelEffectiveInput.EntityData.Children = make(map[string]types.YChild)
    qualityLevelEffectiveInput.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelEffectiveInput.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelEffectiveInput.QualityLevelOption}
    qualityLevelEffectiveInput.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelEffectiveInput.Option1Value}
    qualityLevelEffectiveInput.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelEffectiveInput.Option2Generation1Value}
    qualityLevelEffectiveInput.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelEffectiveInput.Option2Generation2Value}
    return &(qualityLevelEffectiveInput.EntityData)
}

// FrequencySynchronization_Nodes_Node_NtpData_SpaSelectionPoint
// Spa selection points
type FrequencySynchronization_Nodes_Node_NtpData_SpaSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (spaSelectionPoint *FrequencySynchronization_Nodes_Node_NtpData_SpaSelectionPoint) GetEntityData() *types.CommonEntityData {
    spaSelectionPoint.EntityData.YFilter = spaSelectionPoint.YFilter
    spaSelectionPoint.EntityData.YangName = "spa-selection-point"
    spaSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    spaSelectionPoint.EntityData.ParentYangName = "ntp-data"
    spaSelectionPoint.EntityData.SegmentPath = "spa-selection-point"
    spaSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spaSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spaSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spaSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    spaSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    spaSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", spaSelectionPoint.SelectionPoint}
    spaSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", spaSelectionPoint.SelectionPointDescription}
    return &(spaSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_NtpData_NodeSelectionPoint
// Node selection points
type FrequencySynchronization_Nodes_Node_NtpData_NodeSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (nodeSelectionPoint *FrequencySynchronization_Nodes_Node_NtpData_NodeSelectionPoint) GetEntityData() *types.CommonEntityData {
    nodeSelectionPoint.EntityData.YFilter = nodeSelectionPoint.YFilter
    nodeSelectionPoint.EntityData.YangName = "node-selection-point"
    nodeSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    nodeSelectionPoint.EntityData.ParentYangName = "ntp-data"
    nodeSelectionPoint.EntityData.SegmentPath = "node-selection-point"
    nodeSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    nodeSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", nodeSelectionPoint.SelectionPoint}
    nodeSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", nodeSelectionPoint.SelectionPointDescription}
    return &(nodeSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointDatas
// Selection point data table
type FrequencySynchronization_Nodes_Node_SelectionPointDatas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a given selection point. The type is slice of
    // FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData.
    SelectionPointData []FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData
}

func (selectionPointDatas *FrequencySynchronization_Nodes_Node_SelectionPointDatas) GetEntityData() *types.CommonEntityData {
    selectionPointDatas.EntityData.YFilter = selectionPointDatas.YFilter
    selectionPointDatas.EntityData.YangName = "selection-point-datas"
    selectionPointDatas.EntityData.BundleName = "cisco_ios_xr"
    selectionPointDatas.EntityData.ParentYangName = "node"
    selectionPointDatas.EntityData.SegmentPath = "selection-point-datas"
    selectionPointDatas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPointDatas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPointDatas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPointDatas.EntityData.Children = make(map[string]types.YChild)
    selectionPointDatas.EntityData.Children["selection-point-data"] = types.YChild{"SelectionPointData", nil}
    for i := range selectionPointDatas.SelectionPointData {
        selectionPointDatas.EntityData.Children[types.GetSegmentPath(&selectionPointDatas.SelectionPointData[i])] = types.YChild{"SelectionPointData", &selectionPointDatas.SelectionPointData[i]}
    }
    selectionPointDatas.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(selectionPointDatas.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData
// Operational data for a given selection point
type FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData struct {
    EntityData types.CommonEntityData
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
    Output FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output

    // Time the SP was last programmed.
    LastProgrammed FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_LastProgrammed

    // Time the last selection was made.
    LastSelection FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_LastSelection
}

func (selectionPointData *FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData) GetEntityData() *types.CommonEntityData {
    selectionPointData.EntityData.YFilter = selectionPointData.YFilter
    selectionPointData.EntityData.YangName = "selection-point-data"
    selectionPointData.EntityData.BundleName = "cisco_ios_xr"
    selectionPointData.EntityData.ParentYangName = "selection-point-datas"
    selectionPointData.EntityData.SegmentPath = "selection-point-data" + "[selection-point='" + fmt.Sprintf("%v", selectionPointData.SelectionPoint) + "']"
    selectionPointData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPointData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPointData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPointData.EntityData.Children = make(map[string]types.YChild)
    selectionPointData.EntityData.Children["output"] = types.YChild{"Output", &selectionPointData.Output}
    selectionPointData.EntityData.Children["last-programmed"] = types.YChild{"LastProgrammed", &selectionPointData.LastProgrammed}
    selectionPointData.EntityData.Children["last-selection"] = types.YChild{"LastSelection", &selectionPointData.LastSelection}
    selectionPointData.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPointData.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", selectionPointData.SelectionPoint}
    selectionPointData.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", selectionPointData.SelectionPointType}
    selectionPointData.EntityData.Leafs["description"] = types.YLeaf{"Description", selectionPointData.Description}
    selectionPointData.EntityData.Leafs["inputs"] = types.YLeaf{"Inputs", selectionPointData.Inputs}
    selectionPointData.EntityData.Leafs["inputs-selected"] = types.YLeaf{"InputsSelected", selectionPointData.InputsSelected}
    selectionPointData.EntityData.Leafs["time-of-day-selection"] = types.YLeaf{"TimeOfDaySelection", selectionPointData.TimeOfDaySelection}
    return &(selectionPointData.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output
// Information about the output of the selection
// point
type FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Used for local clock output. The type is bool.
    LocalClockOuput interface{}

    // Used for local line interface output. The type is bool.
    LocalLineOutput interface{}

    // Used for local time-of-day output. The type is bool.
    LocalTimeOfDayOutput interface{}

    // SPA selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_SpaSelectionPoint.
    SpaSelectionPoint []FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_SpaSelectionPoint

    // Node selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_NodeSelectionPoint.
    NodeSelectionPoint []FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_NodeSelectionPoint

    // Chassis selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_ChassisSelectionPoint.
    ChassisSelectionPoint []FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_ChassisSelectionPoint

    // Router selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_RouterSelectionPoint.
    RouterSelectionPoint []FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_RouterSelectionPoint
}

func (output *FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "selection-point-data"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Children["spa-selection-point"] = types.YChild{"SpaSelectionPoint", nil}
    for i := range output.SpaSelectionPoint {
        output.EntityData.Children[types.GetSegmentPath(&output.SpaSelectionPoint[i])] = types.YChild{"SpaSelectionPoint", &output.SpaSelectionPoint[i]}
    }
    output.EntityData.Children["node-selection-point"] = types.YChild{"NodeSelectionPoint", nil}
    for i := range output.NodeSelectionPoint {
        output.EntityData.Children[types.GetSegmentPath(&output.NodeSelectionPoint[i])] = types.YChild{"NodeSelectionPoint", &output.NodeSelectionPoint[i]}
    }
    output.EntityData.Children["chassis-selection-point"] = types.YChild{"ChassisSelectionPoint", nil}
    for i := range output.ChassisSelectionPoint {
        output.EntityData.Children[types.GetSegmentPath(&output.ChassisSelectionPoint[i])] = types.YChild{"ChassisSelectionPoint", &output.ChassisSelectionPoint[i]}
    }
    output.EntityData.Children["router-selection-point"] = types.YChild{"RouterSelectionPoint", nil}
    for i := range output.RouterSelectionPoint {
        output.EntityData.Children[types.GetSegmentPath(&output.RouterSelectionPoint[i])] = types.YChild{"RouterSelectionPoint", &output.RouterSelectionPoint[i]}
    }
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["local-clock-ouput"] = types.YLeaf{"LocalClockOuput", output.LocalClockOuput}
    output.EntityData.Leafs["local-line-output"] = types.YLeaf{"LocalLineOutput", output.LocalLineOutput}
    output.EntityData.Leafs["local-time-of-day-output"] = types.YLeaf{"LocalTimeOfDayOutput", output.LocalTimeOfDayOutput}
    return &(output.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_SpaSelectionPoint
// SPA selection points
type FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_SpaSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (spaSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_SpaSelectionPoint) GetEntityData() *types.CommonEntityData {
    spaSelectionPoint.EntityData.YFilter = spaSelectionPoint.YFilter
    spaSelectionPoint.EntityData.YangName = "spa-selection-point"
    spaSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    spaSelectionPoint.EntityData.ParentYangName = "output"
    spaSelectionPoint.EntityData.SegmentPath = "spa-selection-point"
    spaSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spaSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spaSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spaSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    spaSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    spaSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", spaSelectionPoint.SelectionPoint}
    spaSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", spaSelectionPoint.SelectionPointDescription}
    return &(spaSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_NodeSelectionPoint
// Node selection points
type FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_NodeSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (nodeSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_NodeSelectionPoint) GetEntityData() *types.CommonEntityData {
    nodeSelectionPoint.EntityData.YFilter = nodeSelectionPoint.YFilter
    nodeSelectionPoint.EntityData.YangName = "node-selection-point"
    nodeSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    nodeSelectionPoint.EntityData.ParentYangName = "output"
    nodeSelectionPoint.EntityData.SegmentPath = "node-selection-point"
    nodeSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    nodeSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", nodeSelectionPoint.SelectionPoint}
    nodeSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", nodeSelectionPoint.SelectionPointDescription}
    return &(nodeSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_ChassisSelectionPoint
// Chassis selection points
type FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_ChassisSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (chassisSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_ChassisSelectionPoint) GetEntityData() *types.CommonEntityData {
    chassisSelectionPoint.EntityData.YFilter = chassisSelectionPoint.YFilter
    chassisSelectionPoint.EntityData.YangName = "chassis-selection-point"
    chassisSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    chassisSelectionPoint.EntityData.ParentYangName = "output"
    chassisSelectionPoint.EntityData.SegmentPath = "chassis-selection-point"
    chassisSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    chassisSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    chassisSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", chassisSelectionPoint.SelectionPoint}
    chassisSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", chassisSelectionPoint.SelectionPointDescription}
    return &(chassisSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_RouterSelectionPoint
// Router selection points
type FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_RouterSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (routerSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_Output_RouterSelectionPoint) GetEntityData() *types.CommonEntityData {
    routerSelectionPoint.EntityData.YFilter = routerSelectionPoint.YFilter
    routerSelectionPoint.EntityData.YangName = "router-selection-point"
    routerSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    routerSelectionPoint.EntityData.ParentYangName = "output"
    routerSelectionPoint.EntityData.SegmentPath = "router-selection-point"
    routerSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routerSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routerSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routerSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    routerSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    routerSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", routerSelectionPoint.SelectionPoint}
    routerSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", routerSelectionPoint.SelectionPointDescription}
    return &(routerSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_LastProgrammed
// Time the SP was last programmed
type FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_LastProgrammed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastProgrammed *FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_LastProgrammed) GetEntityData() *types.CommonEntityData {
    lastProgrammed.EntityData.YFilter = lastProgrammed.YFilter
    lastProgrammed.EntityData.YangName = "last-programmed"
    lastProgrammed.EntityData.BundleName = "cisco_ios_xr"
    lastProgrammed.EntityData.ParentYangName = "selection-point-data"
    lastProgrammed.EntityData.SegmentPath = "last-programmed"
    lastProgrammed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastProgrammed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastProgrammed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastProgrammed.EntityData.Children = make(map[string]types.YChild)
    lastProgrammed.EntityData.Leafs = make(map[string]types.YLeaf)
    lastProgrammed.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", lastProgrammed.Seconds}
    lastProgrammed.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", lastProgrammed.Nanoseconds}
    return &(lastProgrammed.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_LastSelection
// Time the last selection was made
type FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_LastSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..4294967295. Units are
    // second.
    Seconds interface{}

    // Nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lastSelection *FrequencySynchronization_Nodes_Node_SelectionPointDatas_SelectionPointData_LastSelection) GetEntityData() *types.CommonEntityData {
    lastSelection.EntityData.YFilter = lastSelection.YFilter
    lastSelection.EntityData.YangName = "last-selection"
    lastSelection.EntityData.BundleName = "cisco_ios_xr"
    lastSelection.EntityData.ParentYangName = "selection-point-data"
    lastSelection.EntityData.SegmentPath = "last-selection"
    lastSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastSelection.EntityData.Children = make(map[string]types.YChild)
    lastSelection.EntityData.Leafs = make(map[string]types.YLeaf)
    lastSelection.EntityData.Leafs["seconds"] = types.YLeaf{"Seconds", lastSelection.Seconds}
    lastSelection.EntityData.Leafs["nanoseconds"] = types.YLeaf{"Nanoseconds", lastSelection.Nanoseconds}
    return &(lastSelection.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors
// Configuration error operational data
type FrequencySynchronization_Nodes_Node_ConfigurationErrors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration errors. The type is slice of
    // FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource.
    ErrorSource []FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource
}

func (configurationErrors *FrequencySynchronization_Nodes_Node_ConfigurationErrors) GetEntityData() *types.CommonEntityData {
    configurationErrors.EntityData.YFilter = configurationErrors.YFilter
    configurationErrors.EntityData.YangName = "configuration-errors"
    configurationErrors.EntityData.BundleName = "cisco_ios_xr"
    configurationErrors.EntityData.ParentYangName = "node"
    configurationErrors.EntityData.SegmentPath = "configuration-errors"
    configurationErrors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationErrors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationErrors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationErrors.EntityData.Children = make(map[string]types.YChild)
    configurationErrors.EntityData.Children["error-source"] = types.YChild{"ErrorSource", nil}
    for i := range configurationErrors.ErrorSource {
        configurationErrors.EntityData.Children[types.GetSegmentPath(&configurationErrors.ErrorSource[i])] = types.YChild{"ErrorSource", &configurationErrors.ErrorSource[i]}
    }
    configurationErrors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(configurationErrors.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource
// Configuration errors
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource struct {
    EntityData types.CommonEntityData
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

func (errorSource *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource) GetEntityData() *types.CommonEntityData {
    errorSource.EntityData.YFilter = errorSource.YFilter
    errorSource.EntityData.YangName = "error-source"
    errorSource.EntityData.BundleName = "cisco_ios_xr"
    errorSource.EntityData.ParentYangName = "configuration-errors"
    errorSource.EntityData.SegmentPath = "error-source"
    errorSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    errorSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    errorSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    errorSource.EntityData.Children = make(map[string]types.YChild)
    errorSource.EntityData.Children["source"] = types.YChild{"Source", &errorSource.Source}
    errorSource.EntityData.Children["input-min-ql"] = types.YChild{"InputMinQl", &errorSource.InputMinQl}
    errorSource.EntityData.Children["input-exact-ql"] = types.YChild{"InputExactQl", &errorSource.InputExactQl}
    errorSource.EntityData.Children["input-max-ql"] = types.YChild{"InputMaxQl", &errorSource.InputMaxQl}
    errorSource.EntityData.Children["output-min-ql"] = types.YChild{"OutputMinQl", &errorSource.OutputMinQl}
    errorSource.EntityData.Children["output-exact-ql"] = types.YChild{"OutputExactQl", &errorSource.OutputExactQl}
    errorSource.EntityData.Children["output-max-ql"] = types.YChild{"OutputMaxQl", &errorSource.OutputMaxQl}
    errorSource.EntityData.Leafs = make(map[string]types.YLeaf)
    errorSource.EntityData.Leafs["enable-error"] = types.YLeaf{"EnableError", errorSource.EnableError}
    errorSource.EntityData.Leafs["input-min-error"] = types.YLeaf{"InputMinError", errorSource.InputMinError}
    errorSource.EntityData.Leafs["input-exact-error"] = types.YLeaf{"InputExactError", errorSource.InputExactError}
    errorSource.EntityData.Leafs["input-max-error"] = types.YLeaf{"InputMaxError", errorSource.InputMaxError}
    errorSource.EntityData.Leafs["ouput-min-error"] = types.YLeaf{"OuputMinError", errorSource.OuputMinError}
    errorSource.EntityData.Leafs["output-exact-error"] = types.YLeaf{"OutputExactError", errorSource.OutputExactError}
    errorSource.EntityData.Leafs["output-max-error"] = types.YLeaf{"OutputMaxError", errorSource.OutputMaxError}
    errorSource.EntityData.Leafs["input-output-mismatch"] = types.YLeaf{"InputOutputMismatch", errorSource.InputOutputMismatch}
    return &(errorSource.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source
// Source
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId
}

func (source *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "error-source"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Children["clock-id"] = types.YChild{"ClockId", &source.ClockId}
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", source.SourceClass}
    source.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", source.EthernetInterface}
    source.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", source.SonetInterface}
    source.EntityData.Leafs["node"] = types.YLeaf{"Node", source.Node}
    source.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", source.PtpNode}
    source.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", source.SatelliteAccessInterface}
    source.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", source.NtpNode}
    return &(source.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId
// Clock ID
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_Source_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl
// Configured minimum input QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl struct {
    EntityData types.CommonEntityData
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

func (inputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMinQl) GetEntityData() *types.CommonEntityData {
    inputMinQl.EntityData.YFilter = inputMinQl.YFilter
    inputMinQl.EntityData.YangName = "input-min-ql"
    inputMinQl.EntityData.BundleName = "cisco_ios_xr"
    inputMinQl.EntityData.ParentYangName = "error-source"
    inputMinQl.EntityData.SegmentPath = "input-min-ql"
    inputMinQl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputMinQl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputMinQl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputMinQl.EntityData.Children = make(map[string]types.YChild)
    inputMinQl.EntityData.Leafs = make(map[string]types.YLeaf)
    inputMinQl.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", inputMinQl.QualityLevelOption}
    inputMinQl.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", inputMinQl.Option1Value}
    inputMinQl.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", inputMinQl.Option2Generation1Value}
    inputMinQl.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", inputMinQl.Option2Generation2Value}
    return &(inputMinQl.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl
// Configured exact input QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl struct {
    EntityData types.CommonEntityData
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

func (inputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputExactQl) GetEntityData() *types.CommonEntityData {
    inputExactQl.EntityData.YFilter = inputExactQl.YFilter
    inputExactQl.EntityData.YangName = "input-exact-ql"
    inputExactQl.EntityData.BundleName = "cisco_ios_xr"
    inputExactQl.EntityData.ParentYangName = "error-source"
    inputExactQl.EntityData.SegmentPath = "input-exact-ql"
    inputExactQl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputExactQl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputExactQl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputExactQl.EntityData.Children = make(map[string]types.YChild)
    inputExactQl.EntityData.Leafs = make(map[string]types.YLeaf)
    inputExactQl.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", inputExactQl.QualityLevelOption}
    inputExactQl.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", inputExactQl.Option1Value}
    inputExactQl.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", inputExactQl.Option2Generation1Value}
    inputExactQl.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", inputExactQl.Option2Generation2Value}
    return &(inputExactQl.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl
// Configured maximum input QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl struct {
    EntityData types.CommonEntityData
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

func (inputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_InputMaxQl) GetEntityData() *types.CommonEntityData {
    inputMaxQl.EntityData.YFilter = inputMaxQl.YFilter
    inputMaxQl.EntityData.YangName = "input-max-ql"
    inputMaxQl.EntityData.BundleName = "cisco_ios_xr"
    inputMaxQl.EntityData.ParentYangName = "error-source"
    inputMaxQl.EntityData.SegmentPath = "input-max-ql"
    inputMaxQl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputMaxQl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputMaxQl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputMaxQl.EntityData.Children = make(map[string]types.YChild)
    inputMaxQl.EntityData.Leafs = make(map[string]types.YLeaf)
    inputMaxQl.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", inputMaxQl.QualityLevelOption}
    inputMaxQl.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", inputMaxQl.Option1Value}
    inputMaxQl.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", inputMaxQl.Option2Generation1Value}
    inputMaxQl.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", inputMaxQl.Option2Generation2Value}
    return &(inputMaxQl.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl
// Configured mininum output QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl struct {
    EntityData types.CommonEntityData
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

func (outputMinQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMinQl) GetEntityData() *types.CommonEntityData {
    outputMinQl.EntityData.YFilter = outputMinQl.YFilter
    outputMinQl.EntityData.YangName = "output-min-ql"
    outputMinQl.EntityData.BundleName = "cisco_ios_xr"
    outputMinQl.EntityData.ParentYangName = "error-source"
    outputMinQl.EntityData.SegmentPath = "output-min-ql"
    outputMinQl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputMinQl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputMinQl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputMinQl.EntityData.Children = make(map[string]types.YChild)
    outputMinQl.EntityData.Leafs = make(map[string]types.YLeaf)
    outputMinQl.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", outputMinQl.QualityLevelOption}
    outputMinQl.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", outputMinQl.Option1Value}
    outputMinQl.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", outputMinQl.Option2Generation1Value}
    outputMinQl.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", outputMinQl.Option2Generation2Value}
    return &(outputMinQl.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl
// Configured exact output QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl struct {
    EntityData types.CommonEntityData
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

func (outputExactQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputExactQl) GetEntityData() *types.CommonEntityData {
    outputExactQl.EntityData.YFilter = outputExactQl.YFilter
    outputExactQl.EntityData.YangName = "output-exact-ql"
    outputExactQl.EntityData.BundleName = "cisco_ios_xr"
    outputExactQl.EntityData.ParentYangName = "error-source"
    outputExactQl.EntityData.SegmentPath = "output-exact-ql"
    outputExactQl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputExactQl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputExactQl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputExactQl.EntityData.Children = make(map[string]types.YChild)
    outputExactQl.EntityData.Leafs = make(map[string]types.YLeaf)
    outputExactQl.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", outputExactQl.QualityLevelOption}
    outputExactQl.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", outputExactQl.Option1Value}
    outputExactQl.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", outputExactQl.Option2Generation1Value}
    outputExactQl.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", outputExactQl.Option2Generation2Value}
    return &(outputExactQl.EntityData)
}

// FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl
// Configured exact maximum QL
type FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl struct {
    EntityData types.CommonEntityData
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

func (outputMaxQl *FrequencySynchronization_Nodes_Node_ConfigurationErrors_ErrorSource_OutputMaxQl) GetEntityData() *types.CommonEntityData {
    outputMaxQl.EntityData.YFilter = outputMaxQl.YFilter
    outputMaxQl.EntityData.YangName = "output-max-ql"
    outputMaxQl.EntityData.BundleName = "cisco_ios_xr"
    outputMaxQl.EntityData.ParentYangName = "error-source"
    outputMaxQl.EntityData.SegmentPath = "output-max-ql"
    outputMaxQl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputMaxQl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputMaxQl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputMaxQl.EntityData.Children = make(map[string]types.YChild)
    outputMaxQl.EntityData.Leafs = make(map[string]types.YLeaf)
    outputMaxQl.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", outputMaxQl.QualityLevelOption}
    outputMaxQl.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", outputMaxQl.Option1Value}
    outputMaxQl.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", outputMaxQl.Option2Generation1Value}
    outputMaxQl.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", outputMaxQl.Option2Generation2Value}
    return &(outputMaxQl.EntityData)
}

// FrequencySynchronization_Nodes_Node_PtpData
// PTP operational data
type FrequencySynchronization_Nodes_Node_PtpData struct {
    EntityData types.CommonEntityData
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

    // The effective input quality level.
    QualityLevelEffectiveInput FrequencySynchronization_Nodes_Node_PtpData_QualityLevelEffectiveInput

    // Spa selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_PtpData_SpaSelectionPoint.
    SpaSelectionPoint []FrequencySynchronization_Nodes_Node_PtpData_SpaSelectionPoint

    // Node selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_PtpData_NodeSelectionPoint.
    NodeSelectionPoint []FrequencySynchronization_Nodes_Node_PtpData_NodeSelectionPoint
}

func (ptpData *FrequencySynchronization_Nodes_Node_PtpData) GetEntityData() *types.CommonEntityData {
    ptpData.EntityData.YFilter = ptpData.YFilter
    ptpData.EntityData.YangName = "ptp-data"
    ptpData.EntityData.BundleName = "cisco_ios_xr"
    ptpData.EntityData.ParentYangName = "node"
    ptpData.EntityData.SegmentPath = "ptp-data"
    ptpData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ptpData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ptpData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ptpData.EntityData.Children = make(map[string]types.YChild)
    ptpData.EntityData.Children["quality-level-effective-input"] = types.YChild{"QualityLevelEffectiveInput", &ptpData.QualityLevelEffectiveInput}
    ptpData.EntityData.Children["spa-selection-point"] = types.YChild{"SpaSelectionPoint", nil}
    for i := range ptpData.SpaSelectionPoint {
        ptpData.EntityData.Children[types.GetSegmentPath(&ptpData.SpaSelectionPoint[i])] = types.YChild{"SpaSelectionPoint", &ptpData.SpaSelectionPoint[i]}
    }
    ptpData.EntityData.Children["node-selection-point"] = types.YChild{"NodeSelectionPoint", nil}
    for i := range ptpData.NodeSelectionPoint {
        ptpData.EntityData.Children[types.GetSegmentPath(&ptpData.NodeSelectionPoint[i])] = types.YChild{"NodeSelectionPoint", &ptpData.NodeSelectionPoint[i]}
    }
    ptpData.EntityData.Leafs = make(map[string]types.YLeaf)
    ptpData.EntityData.Leafs["state"] = types.YLeaf{"State", ptpData.State}
    ptpData.EntityData.Leafs["supports-frequency"] = types.YLeaf{"SupportsFrequency", ptpData.SupportsFrequency}
    ptpData.EntityData.Leafs["supports-time-of-day"] = types.YLeaf{"SupportsTimeOfDay", ptpData.SupportsTimeOfDay}
    ptpData.EntityData.Leafs["frequency-priority"] = types.YLeaf{"FrequencyPriority", ptpData.FrequencyPriority}
    ptpData.EntityData.Leafs["time-of-day-priority"] = types.YLeaf{"TimeOfDayPriority", ptpData.TimeOfDayPriority}
    return &(ptpData.EntityData)
}

// FrequencySynchronization_Nodes_Node_PtpData_QualityLevelEffectiveInput
// The effective input quality level
type FrequencySynchronization_Nodes_Node_PtpData_QualityLevelEffectiveInput struct {
    EntityData types.CommonEntityData
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

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_PtpData_QualityLevelEffectiveInput) GetEntityData() *types.CommonEntityData {
    qualityLevelEffectiveInput.EntityData.YFilter = qualityLevelEffectiveInput.YFilter
    qualityLevelEffectiveInput.EntityData.YangName = "quality-level-effective-input"
    qualityLevelEffectiveInput.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelEffectiveInput.EntityData.ParentYangName = "ptp-data"
    qualityLevelEffectiveInput.EntityData.SegmentPath = "quality-level-effective-input"
    qualityLevelEffectiveInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelEffectiveInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelEffectiveInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelEffectiveInput.EntityData.Children = make(map[string]types.YChild)
    qualityLevelEffectiveInput.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelEffectiveInput.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelEffectiveInput.QualityLevelOption}
    qualityLevelEffectiveInput.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelEffectiveInput.Option1Value}
    qualityLevelEffectiveInput.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelEffectiveInput.Option2Generation1Value}
    qualityLevelEffectiveInput.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelEffectiveInput.Option2Generation2Value}
    return &(qualityLevelEffectiveInput.EntityData)
}

// FrequencySynchronization_Nodes_Node_PtpData_SpaSelectionPoint
// Spa selection points
type FrequencySynchronization_Nodes_Node_PtpData_SpaSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (spaSelectionPoint *FrequencySynchronization_Nodes_Node_PtpData_SpaSelectionPoint) GetEntityData() *types.CommonEntityData {
    spaSelectionPoint.EntityData.YFilter = spaSelectionPoint.YFilter
    spaSelectionPoint.EntityData.YangName = "spa-selection-point"
    spaSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    spaSelectionPoint.EntityData.ParentYangName = "ptp-data"
    spaSelectionPoint.EntityData.SegmentPath = "spa-selection-point"
    spaSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spaSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spaSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spaSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    spaSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    spaSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", spaSelectionPoint.SelectionPoint}
    spaSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", spaSelectionPoint.SelectionPointDescription}
    return &(spaSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_PtpData_NodeSelectionPoint
// Node selection points
type FrequencySynchronization_Nodes_Node_PtpData_NodeSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (nodeSelectionPoint *FrequencySynchronization_Nodes_Node_PtpData_NodeSelectionPoint) GetEntityData() *types.CommonEntityData {
    nodeSelectionPoint.EntityData.YFilter = nodeSelectionPoint.YFilter
    nodeSelectionPoint.EntityData.YangName = "node-selection-point"
    nodeSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    nodeSelectionPoint.EntityData.ParentYangName = "ptp-data"
    nodeSelectionPoint.EntityData.SegmentPath = "node-selection-point"
    nodeSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    nodeSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", nodeSelectionPoint.SelectionPoint}
    nodeSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", nodeSelectionPoint.SelectionPointDescription}
    return &(nodeSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_SsmSummary
// SSM operational data
type FrequencySynchronization_Nodes_Node_SsmSummary struct {
    EntityData types.CommonEntityData
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

func (ssmSummary *FrequencySynchronization_Nodes_Node_SsmSummary) GetEntityData() *types.CommonEntityData {
    ssmSummary.EntityData.YFilter = ssmSummary.YFilter
    ssmSummary.EntityData.YangName = "ssm-summary"
    ssmSummary.EntityData.BundleName = "cisco_ios_xr"
    ssmSummary.EntityData.ParentYangName = "node"
    ssmSummary.EntityData.SegmentPath = "ssm-summary"
    ssmSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssmSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssmSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssmSummary.EntityData.Children = make(map[string]types.YChild)
    ssmSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    ssmSummary.EntityData.Leafs["ethernet-sources"] = types.YLeaf{"EthernetSources", ssmSummary.EthernetSources}
    ssmSummary.EntityData.Leafs["ethernet-sources-select"] = types.YLeaf{"EthernetSourcesSelect", ssmSummary.EthernetSourcesSelect}
    ssmSummary.EntityData.Leafs["ethernet-sources-enabled"] = types.YLeaf{"EthernetSourcesEnabled", ssmSummary.EthernetSourcesEnabled}
    ssmSummary.EntityData.Leafs["sonet-sources"] = types.YLeaf{"SonetSources", ssmSummary.SonetSources}
    ssmSummary.EntityData.Leafs["sonet-sources-select"] = types.YLeaf{"SonetSourcesSelect", ssmSummary.SonetSourcesSelect}
    ssmSummary.EntityData.Leafs["sonet-sources-enabled"] = types.YLeaf{"SonetSourcesEnabled", ssmSummary.SonetSourcesEnabled}
    ssmSummary.EntityData.Leafs["events-sent"] = types.YLeaf{"EventsSent", ssmSummary.EventsSent}
    ssmSummary.EntityData.Leafs["events-received"] = types.YLeaf{"EventsReceived", ssmSummary.EventsReceived}
    ssmSummary.EntityData.Leafs["infos-sent"] = types.YLeaf{"InfosSent", ssmSummary.InfosSent}
    ssmSummary.EntityData.Leafs["infos-received"] = types.YLeaf{"InfosReceived", ssmSummary.InfosReceived}
    ssmSummary.EntityData.Leafs["dn-us-sent"] = types.YLeaf{"DnUsSent", ssmSummary.DnUsSent}
    ssmSummary.EntityData.Leafs["dn-us-received"] = types.YLeaf{"DnUsReceived", ssmSummary.DnUsReceived}
    return &(ssmSummary.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas
// Table for clock operational data
type FrequencySynchronization_Nodes_Node_ClockDatas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular clock. The type is slice of
    // FrequencySynchronization_Nodes_Node_ClockDatas_ClockData.
    ClockData []FrequencySynchronization_Nodes_Node_ClockDatas_ClockData
}

func (clockDatas *FrequencySynchronization_Nodes_Node_ClockDatas) GetEntityData() *types.CommonEntityData {
    clockDatas.EntityData.YFilter = clockDatas.YFilter
    clockDatas.EntityData.YangName = "clock-datas"
    clockDatas.EntityData.BundleName = "cisco_ios_xr"
    clockDatas.EntityData.ParentYangName = "node"
    clockDatas.EntityData.SegmentPath = "clock-datas"
    clockDatas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockDatas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockDatas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockDatas.EntityData.Children = make(map[string]types.YChild)
    clockDatas.EntityData.Children["clock-data"] = types.YChild{"ClockData", nil}
    for i := range clockDatas.ClockData {
        clockDatas.EntityData.Children[types.GetSegmentPath(&clockDatas.ClockData[i])] = types.YChild{"ClockData", &clockDatas.ClockData[i]}
    }
    clockDatas.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clockDatas.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData
// Operational data for a particular clock
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Clock type. The type is FsyncClock.
    ClockType interface{}

    // This attribute is a key. Clock port. The type is interface{} with range:
    // -2147483648..2147483647.
    Port interface{}

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
    ClockTypeXr interface{}

    // The PTP clock supports frequency. The type is bool.
    SupportsFrequency interface{}

    // The PTP clock supports time. The type is bool.
    SupportsTimeOfDay interface{}

    // The source ID for the clock.
    Source FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_Source

    // Timing source selected for clock output.
    SelectedSource FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SelectedSource

    // Received quality level.
    QualityLevelReceived FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelReceived

    // Quality level after damping has been applied.
    QualityLevelDamped FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelDamped

    // The effective input quality level.
    QualityLevelEffectiveInput FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelEffectiveInput

    // The effective output quality level.
    QualityLevelEffectiveOutput FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelEffectiveOutput

    // The quality level of the source driving this interface.
    QualityLevelSelectedSource FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelSelectedSource

    // Spa selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SpaSelectionPoint.
    SpaSelectionPoint []FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SpaSelectionPoint

    // Node selection points. The type is slice of
    // FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_NodeSelectionPoint.
    NodeSelectionPoint []FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_NodeSelectionPoint
}

func (clockData *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData) GetEntityData() *types.CommonEntityData {
    clockData.EntityData.YFilter = clockData.YFilter
    clockData.EntityData.YangName = "clock-data"
    clockData.EntityData.BundleName = "cisco_ios_xr"
    clockData.EntityData.ParentYangName = "clock-datas"
    clockData.EntityData.SegmentPath = "clock-data" + "[clock-type='" + fmt.Sprintf("%v", clockData.ClockType) + "']" + "[port='" + fmt.Sprintf("%v", clockData.Port) + "']"
    clockData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockData.EntityData.Children = make(map[string]types.YChild)
    clockData.EntityData.Children["source"] = types.YChild{"Source", &clockData.Source}
    clockData.EntityData.Children["selected-source"] = types.YChild{"SelectedSource", &clockData.SelectedSource}
    clockData.EntityData.Children["quality-level-received"] = types.YChild{"QualityLevelReceived", &clockData.QualityLevelReceived}
    clockData.EntityData.Children["quality-level-damped"] = types.YChild{"QualityLevelDamped", &clockData.QualityLevelDamped}
    clockData.EntityData.Children["quality-level-effective-input"] = types.YChild{"QualityLevelEffectiveInput", &clockData.QualityLevelEffectiveInput}
    clockData.EntityData.Children["quality-level-effective-output"] = types.YChild{"QualityLevelEffectiveOutput", &clockData.QualityLevelEffectiveOutput}
    clockData.EntityData.Children["quality-level-selected-source"] = types.YChild{"QualityLevelSelectedSource", &clockData.QualityLevelSelectedSource}
    clockData.EntityData.Children["spa-selection-point"] = types.YChild{"SpaSelectionPoint", nil}
    for i := range clockData.SpaSelectionPoint {
        clockData.EntityData.Children[types.GetSegmentPath(&clockData.SpaSelectionPoint[i])] = types.YChild{"SpaSelectionPoint", &clockData.SpaSelectionPoint[i]}
    }
    clockData.EntityData.Children["node-selection-point"] = types.YChild{"NodeSelectionPoint", nil}
    for i := range clockData.NodeSelectionPoint {
        clockData.EntityData.Children[types.GetSegmentPath(&clockData.NodeSelectionPoint[i])] = types.YChild{"NodeSelectionPoint", &clockData.NodeSelectionPoint[i]}
    }
    clockData.EntityData.Leafs = make(map[string]types.YLeaf)
    clockData.EntityData.Leafs["clock-type"] = types.YLeaf{"ClockType", clockData.ClockType}
    clockData.EntityData.Leafs["port"] = types.YLeaf{"Port", clockData.Port}
    clockData.EntityData.Leafs["state"] = types.YLeaf{"State", clockData.State}
    clockData.EntityData.Leafs["down-reason"] = types.YLeaf{"DownReason", clockData.DownReason}
    clockData.EntityData.Leafs["description"] = types.YLeaf{"Description", clockData.Description}
    clockData.EntityData.Leafs["priority"] = types.YLeaf{"Priority", clockData.Priority}
    clockData.EntityData.Leafs["time-of-day-priority"] = types.YLeaf{"TimeOfDayPriority", clockData.TimeOfDayPriority}
    clockData.EntityData.Leafs["ssm-support"] = types.YLeaf{"SsmSupport", clockData.SsmSupport}
    clockData.EntityData.Leafs["ssm-enabled"] = types.YLeaf{"SsmEnabled", clockData.SsmEnabled}
    clockData.EntityData.Leafs["loopback"] = types.YLeaf{"Loopback", clockData.Loopback}
    clockData.EntityData.Leafs["selection-input"] = types.YLeaf{"SelectionInput", clockData.SelectionInput}
    clockData.EntityData.Leafs["squelched"] = types.YLeaf{"Squelched", clockData.Squelched}
    clockData.EntityData.Leafs["damping-state"] = types.YLeaf{"DampingState", clockData.DampingState}
    clockData.EntityData.Leafs["damping-time"] = types.YLeaf{"DampingTime", clockData.DampingTime}
    clockData.EntityData.Leafs["input-disabled"] = types.YLeaf{"InputDisabled", clockData.InputDisabled}
    clockData.EntityData.Leafs["output-disabled"] = types.YLeaf{"OutputDisabled", clockData.OutputDisabled}
    clockData.EntityData.Leafs["wait-to-restore-time"] = types.YLeaf{"WaitToRestoreTime", clockData.WaitToRestoreTime}
    clockData.EntityData.Leafs["clock-type-xr"] = types.YLeaf{"ClockTypeXr", clockData.ClockTypeXr}
    clockData.EntityData.Leafs["supports-frequency"] = types.YLeaf{"SupportsFrequency", clockData.SupportsFrequency}
    clockData.EntityData.Leafs["supports-time-of-day"] = types.YLeaf{"SupportsTimeOfDay", clockData.SupportsTimeOfDay}
    return &(clockData.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_Source
// The source ID for the clock
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_Source_ClockId
}

func (source *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "clock-data"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Children["clock-id"] = types.YChild{"ClockId", &source.ClockId}
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", source.SourceClass}
    source.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", source.EthernetInterface}
    source.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", source.SonetInterface}
    source.EntityData.Leafs["node"] = types.YLeaf{"Node", source.Node}
    source.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", source.PtpNode}
    source.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", source.SatelliteAccessInterface}
    source.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", source.NtpNode}
    return &(source.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_Source_ClockId
// Clock ID
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_Source_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_Source_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SelectedSource
// Timing source selected for clock output
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SelectedSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SelectedSource_ClockId
}

func (selectedSource *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SelectedSource) GetEntityData() *types.CommonEntityData {
    selectedSource.EntityData.YFilter = selectedSource.YFilter
    selectedSource.EntityData.YangName = "selected-source"
    selectedSource.EntityData.BundleName = "cisco_ios_xr"
    selectedSource.EntityData.ParentYangName = "clock-data"
    selectedSource.EntityData.SegmentPath = "selected-source"
    selectedSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectedSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectedSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectedSource.EntityData.Children = make(map[string]types.YChild)
    selectedSource.EntityData.Children["clock-id"] = types.YChild{"ClockId", &selectedSource.ClockId}
    selectedSource.EntityData.Leafs = make(map[string]types.YLeaf)
    selectedSource.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", selectedSource.SourceClass}
    selectedSource.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", selectedSource.EthernetInterface}
    selectedSource.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", selectedSource.SonetInterface}
    selectedSource.EntityData.Leafs["node"] = types.YLeaf{"Node", selectedSource.Node}
    selectedSource.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", selectedSource.PtpNode}
    selectedSource.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", selectedSource.SatelliteAccessInterface}
    selectedSource.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", selectedSource.NtpNode}
    return &(selectedSource.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SelectedSource_ClockId
// Clock ID
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SelectedSource_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SelectedSource_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "selected-source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelReceived
// Received quality level
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelReceived struct {
    EntityData types.CommonEntityData
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

func (qualityLevelReceived *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelReceived) GetEntityData() *types.CommonEntityData {
    qualityLevelReceived.EntityData.YFilter = qualityLevelReceived.YFilter
    qualityLevelReceived.EntityData.YangName = "quality-level-received"
    qualityLevelReceived.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelReceived.EntityData.ParentYangName = "clock-data"
    qualityLevelReceived.EntityData.SegmentPath = "quality-level-received"
    qualityLevelReceived.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelReceived.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelReceived.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelReceived.EntityData.Children = make(map[string]types.YChild)
    qualityLevelReceived.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelReceived.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelReceived.QualityLevelOption}
    qualityLevelReceived.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelReceived.Option1Value}
    qualityLevelReceived.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelReceived.Option2Generation1Value}
    qualityLevelReceived.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelReceived.Option2Generation2Value}
    return &(qualityLevelReceived.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelDamped
// Quality level after damping has been applied
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelDamped struct {
    EntityData types.CommonEntityData
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

func (qualityLevelDamped *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelDamped) GetEntityData() *types.CommonEntityData {
    qualityLevelDamped.EntityData.YFilter = qualityLevelDamped.YFilter
    qualityLevelDamped.EntityData.YangName = "quality-level-damped"
    qualityLevelDamped.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelDamped.EntityData.ParentYangName = "clock-data"
    qualityLevelDamped.EntityData.SegmentPath = "quality-level-damped"
    qualityLevelDamped.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelDamped.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelDamped.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelDamped.EntityData.Children = make(map[string]types.YChild)
    qualityLevelDamped.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelDamped.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelDamped.QualityLevelOption}
    qualityLevelDamped.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelDamped.Option1Value}
    qualityLevelDamped.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelDamped.Option2Generation1Value}
    qualityLevelDamped.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelDamped.Option2Generation2Value}
    return &(qualityLevelDamped.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelEffectiveInput
// The effective input quality level
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelEffectiveInput struct {
    EntityData types.CommonEntityData
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

func (qualityLevelEffectiveInput *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelEffectiveInput) GetEntityData() *types.CommonEntityData {
    qualityLevelEffectiveInput.EntityData.YFilter = qualityLevelEffectiveInput.YFilter
    qualityLevelEffectiveInput.EntityData.YangName = "quality-level-effective-input"
    qualityLevelEffectiveInput.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelEffectiveInput.EntityData.ParentYangName = "clock-data"
    qualityLevelEffectiveInput.EntityData.SegmentPath = "quality-level-effective-input"
    qualityLevelEffectiveInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelEffectiveInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelEffectiveInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelEffectiveInput.EntityData.Children = make(map[string]types.YChild)
    qualityLevelEffectiveInput.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelEffectiveInput.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelEffectiveInput.QualityLevelOption}
    qualityLevelEffectiveInput.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelEffectiveInput.Option1Value}
    qualityLevelEffectiveInput.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelEffectiveInput.Option2Generation1Value}
    qualityLevelEffectiveInput.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelEffectiveInput.Option2Generation2Value}
    return &(qualityLevelEffectiveInput.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelEffectiveOutput
// The effective output quality level
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelEffectiveOutput struct {
    EntityData types.CommonEntityData
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

func (qualityLevelEffectiveOutput *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelEffectiveOutput) GetEntityData() *types.CommonEntityData {
    qualityLevelEffectiveOutput.EntityData.YFilter = qualityLevelEffectiveOutput.YFilter
    qualityLevelEffectiveOutput.EntityData.YangName = "quality-level-effective-output"
    qualityLevelEffectiveOutput.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelEffectiveOutput.EntityData.ParentYangName = "clock-data"
    qualityLevelEffectiveOutput.EntityData.SegmentPath = "quality-level-effective-output"
    qualityLevelEffectiveOutput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelEffectiveOutput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelEffectiveOutput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelEffectiveOutput.EntityData.Children = make(map[string]types.YChild)
    qualityLevelEffectiveOutput.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelEffectiveOutput.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelEffectiveOutput.QualityLevelOption}
    qualityLevelEffectiveOutput.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelEffectiveOutput.Option1Value}
    qualityLevelEffectiveOutput.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelEffectiveOutput.Option2Generation1Value}
    qualityLevelEffectiveOutput.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelEffectiveOutput.Option2Generation2Value}
    return &(qualityLevelEffectiveOutput.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelSelectedSource
// The quality level of the source driving this
// interface
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelSelectedSource struct {
    EntityData types.CommonEntityData
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

func (qualityLevelSelectedSource *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_QualityLevelSelectedSource) GetEntityData() *types.CommonEntityData {
    qualityLevelSelectedSource.EntityData.YFilter = qualityLevelSelectedSource.YFilter
    qualityLevelSelectedSource.EntityData.YangName = "quality-level-selected-source"
    qualityLevelSelectedSource.EntityData.BundleName = "cisco_ios_xr"
    qualityLevelSelectedSource.EntityData.ParentYangName = "clock-data"
    qualityLevelSelectedSource.EntityData.SegmentPath = "quality-level-selected-source"
    qualityLevelSelectedSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevelSelectedSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevelSelectedSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevelSelectedSource.EntityData.Children = make(map[string]types.YChild)
    qualityLevelSelectedSource.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevelSelectedSource.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevelSelectedSource.QualityLevelOption}
    qualityLevelSelectedSource.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevelSelectedSource.Option1Value}
    qualityLevelSelectedSource.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevelSelectedSource.Option2Generation1Value}
    qualityLevelSelectedSource.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevelSelectedSource.Option2Generation2Value}
    return &(qualityLevelSelectedSource.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SpaSelectionPoint
// Spa selection points
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SpaSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (spaSelectionPoint *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_SpaSelectionPoint) GetEntityData() *types.CommonEntityData {
    spaSelectionPoint.EntityData.YFilter = spaSelectionPoint.YFilter
    spaSelectionPoint.EntityData.YangName = "spa-selection-point"
    spaSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    spaSelectionPoint.EntityData.ParentYangName = "clock-data"
    spaSelectionPoint.EntityData.SegmentPath = "spa-selection-point"
    spaSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spaSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spaSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spaSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    spaSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    spaSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", spaSelectionPoint.SelectionPoint}
    spaSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", spaSelectionPoint.SelectionPointDescription}
    return &(spaSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_NodeSelectionPoint
// Node selection points
type FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_NodeSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range: 0..255.
    SelectionPoint interface{}

    // Selection point description. The type is string.
    SelectionPointDescription interface{}
}

func (nodeSelectionPoint *FrequencySynchronization_Nodes_Node_ClockDatas_ClockData_NodeSelectionPoint) GetEntityData() *types.CommonEntityData {
    nodeSelectionPoint.EntityData.YFilter = nodeSelectionPoint.YFilter
    nodeSelectionPoint.EntityData.YangName = "node-selection-point"
    nodeSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    nodeSelectionPoint.EntityData.ParentYangName = "clock-data"
    nodeSelectionPoint.EntityData.SegmentPath = "node-selection-point"
    nodeSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    nodeSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeSelectionPoint.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", nodeSelectionPoint.SelectionPoint}
    nodeSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", nodeSelectionPoint.SelectionPointDescription}
    return &(nodeSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs
// Table for selection point input operational
// data
type FrequencySynchronization_Nodes_Node_SelectionPointInputs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular selection point input. The type is slice
    // of
    // FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput.
    SelectionPointInput []FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput
}

func (selectionPointInputs *FrequencySynchronization_Nodes_Node_SelectionPointInputs) GetEntityData() *types.CommonEntityData {
    selectionPointInputs.EntityData.YFilter = selectionPointInputs.YFilter
    selectionPointInputs.EntityData.YangName = "selection-point-inputs"
    selectionPointInputs.EntityData.BundleName = "cisco_ios_xr"
    selectionPointInputs.EntityData.ParentYangName = "node"
    selectionPointInputs.EntityData.SegmentPath = "selection-point-inputs"
    selectionPointInputs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPointInputs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPointInputs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPointInputs.EntityData.Children = make(map[string]types.YChild)
    selectionPointInputs.EntityData.Children["selection-point-input"] = types.YChild{"SelectionPointInput", nil}
    for i := range selectionPointInputs.SelectionPointInput {
        selectionPointInputs.EntityData.Children[types.GetSegmentPath(&selectionPointInputs.SelectionPointInput[i])] = types.YChild{"SelectionPointInput", &selectionPointInputs.SelectionPointInput[i]}
    }
    selectionPointInputs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(selectionPointInputs.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput
// Operational data for a particular selection
// point input
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SelectionPoint interface{}

    // Type of stream. The type is FsyncStream.
    StreamType interface{}

    // Type of source. The type is FsyncSource.
    SourceType interface{}

    // Interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Clock port. The type is interface{} with range: -2147483648..2147483647.
    ClockPort interface{}

    // Last node for a selection point stream. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (selectionPointInput *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput) GetEntityData() *types.CommonEntityData {
    selectionPointInput.EntityData.YFilter = selectionPointInput.YFilter
    selectionPointInput.EntityData.YangName = "selection-point-input"
    selectionPointInput.EntityData.BundleName = "cisco_ios_xr"
    selectionPointInput.EntityData.ParentYangName = "selection-point-inputs"
    selectionPointInput.EntityData.SegmentPath = "selection-point-input"
    selectionPointInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPointInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPointInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPointInput.EntityData.Children = make(map[string]types.YChild)
    selectionPointInput.EntityData.Children["input-selection-point"] = types.YChild{"InputSelectionPoint", &selectionPointInput.InputSelectionPoint}
    selectionPointInput.EntityData.Children["stream"] = types.YChild{"Stream", &selectionPointInput.Stream}
    selectionPointInput.EntityData.Children["original-source"] = types.YChild{"OriginalSource", &selectionPointInput.OriginalSource}
    selectionPointInput.EntityData.Children["quality-level"] = types.YChild{"QualityLevel", &selectionPointInput.QualityLevel}
    selectionPointInput.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPointInput.EntityData.Leafs["selection-point"] = types.YLeaf{"SelectionPoint", selectionPointInput.SelectionPoint}
    selectionPointInput.EntityData.Leafs["stream-type"] = types.YLeaf{"StreamType", selectionPointInput.StreamType}
    selectionPointInput.EntityData.Leafs["source-type"] = types.YLeaf{"SourceType", selectionPointInput.SourceType}
    selectionPointInput.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", selectionPointInput.Interface_}
    selectionPointInput.EntityData.Leafs["clock-port"] = types.YLeaf{"ClockPort", selectionPointInput.ClockPort}
    selectionPointInput.EntityData.Leafs["last-node"] = types.YLeaf{"LastNode", selectionPointInput.LastNode}
    selectionPointInput.EntityData.Leafs["last-selection-point"] = types.YLeaf{"LastSelectionPoint", selectionPointInput.LastSelectionPoint}
    selectionPointInput.EntityData.Leafs["output-id"] = types.YLeaf{"OutputId", selectionPointInput.OutputId}
    selectionPointInput.EntityData.Leafs["supports-frequency"] = types.YLeaf{"SupportsFrequency", selectionPointInput.SupportsFrequency}
    selectionPointInput.EntityData.Leafs["supports-time-of-day"] = types.YLeaf{"SupportsTimeOfDay", selectionPointInput.SupportsTimeOfDay}
    selectionPointInput.EntityData.Leafs["priority"] = types.YLeaf{"Priority", selectionPointInput.Priority}
    selectionPointInput.EntityData.Leafs["time-of-day-priority"] = types.YLeaf{"TimeOfDayPriority", selectionPointInput.TimeOfDayPriority}
    selectionPointInput.EntityData.Leafs["selected"] = types.YLeaf{"Selected", selectionPointInput.Selected}
    selectionPointInput.EntityData.Leafs["output-id-xr"] = types.YLeaf{"OutputIdXr", selectionPointInput.OutputIdXr}
    selectionPointInput.EntityData.Leafs["platform-status"] = types.YLeaf{"PlatformStatus", selectionPointInput.PlatformStatus}
    selectionPointInput.EntityData.Leafs["platform-failed-reason"] = types.YLeaf{"PlatformFailedReason", selectionPointInput.PlatformFailedReason}
    return &(selectionPointInput.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint
// The selection point the input is for
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}
}

func (inputSelectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_InputSelectionPoint) GetEntityData() *types.CommonEntityData {
    inputSelectionPoint.EntityData.YFilter = inputSelectionPoint.YFilter
    inputSelectionPoint.EntityData.YangName = "input-selection-point"
    inputSelectionPoint.EntityData.BundleName = "cisco_ios_xr"
    inputSelectionPoint.EntityData.ParentYangName = "selection-point-input"
    inputSelectionPoint.EntityData.SegmentPath = "input-selection-point"
    inputSelectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inputSelectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inputSelectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inputSelectionPoint.EntityData.Children = make(map[string]types.YChild)
    inputSelectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    inputSelectionPoint.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", inputSelectionPoint.SelectionPointType}
    inputSelectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", inputSelectionPoint.SelectionPointDescription}
    inputSelectionPoint.EntityData.Leafs["node"] = types.YLeaf{"Node", inputSelectionPoint.Node}
    return &(inputSelectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream
// Stream
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // StreamInput. The type is FsyncBagStreamInput.
    StreamInput interface{}

    // Source ID.
    SourceId FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId

    // Selection point ID.
    SelectionPointId FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId
}

func (stream *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream) GetEntityData() *types.CommonEntityData {
    stream.EntityData.YFilter = stream.YFilter
    stream.EntityData.YangName = "stream"
    stream.EntityData.BundleName = "cisco_ios_xr"
    stream.EntityData.ParentYangName = "selection-point-input"
    stream.EntityData.SegmentPath = "stream"
    stream.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stream.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stream.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stream.EntityData.Children = make(map[string]types.YChild)
    stream.EntityData.Children["source-id"] = types.YChild{"SourceId", &stream.SourceId}
    stream.EntityData.Children["selection-point-id"] = types.YChild{"SelectionPointId", &stream.SelectionPointId}
    stream.EntityData.Leafs = make(map[string]types.YLeaf)
    stream.EntityData.Leafs["stream-input"] = types.YLeaf{"StreamInput", stream.StreamInput}
    return &(stream.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId
// Source ID
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId
}

func (sourceId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId) GetEntityData() *types.CommonEntityData {
    sourceId.EntityData.YFilter = sourceId.YFilter
    sourceId.EntityData.YangName = "source-id"
    sourceId.EntityData.BundleName = "cisco_ios_xr"
    sourceId.EntityData.ParentYangName = "stream"
    sourceId.EntityData.SegmentPath = "source-id"
    sourceId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceId.EntityData.Children = make(map[string]types.YChild)
    sourceId.EntityData.Children["clock-id"] = types.YChild{"ClockId", &sourceId.ClockId}
    sourceId.EntityData.Leafs = make(map[string]types.YLeaf)
    sourceId.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", sourceId.SourceClass}
    sourceId.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", sourceId.EthernetInterface}
    sourceId.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", sourceId.SonetInterface}
    sourceId.EntityData.Leafs["node"] = types.YLeaf{"Node", sourceId.Node}
    sourceId.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", sourceId.PtpNode}
    sourceId.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", sourceId.SatelliteAccessInterface}
    sourceId.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", sourceId.NtpNode}
    return &(sourceId.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId
// Clock ID
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SourceId_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "source-id"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId
// Selection point ID
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Output ID. The type is interface{} with range: 0..255.
    OutputId interface{}

    // Last selection point.
    SelectionPoint FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint
}

func (selectionPointId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId) GetEntityData() *types.CommonEntityData {
    selectionPointId.EntityData.YFilter = selectionPointId.YFilter
    selectionPointId.EntityData.YangName = "selection-point-id"
    selectionPointId.EntityData.BundleName = "cisco_ios_xr"
    selectionPointId.EntityData.ParentYangName = "stream"
    selectionPointId.EntityData.SegmentPath = "selection-point-id"
    selectionPointId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPointId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPointId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPointId.EntityData.Children = make(map[string]types.YChild)
    selectionPointId.EntityData.Children["selection-point"] = types.YChild{"SelectionPoint", &selectionPointId.SelectionPoint}
    selectionPointId.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPointId.EntityData.Leafs["output-id"] = types.YLeaf{"OutputId", selectionPointId.OutputId}
    return &(selectionPointId.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint
// Last selection point
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Selection point type. The type is interface{} with range: 0..255.
    SelectionPointType interface{}

    // Selection point descrption. The type is string.
    SelectionPointDescription interface{}

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}
}

func (selectionPoint *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_Stream_SelectionPointId_SelectionPoint) GetEntityData() *types.CommonEntityData {
    selectionPoint.EntityData.YFilter = selectionPoint.YFilter
    selectionPoint.EntityData.YangName = "selection-point"
    selectionPoint.EntityData.BundleName = "cisco_ios_xr"
    selectionPoint.EntityData.ParentYangName = "selection-point-id"
    selectionPoint.EntityData.SegmentPath = "selection-point"
    selectionPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectionPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectionPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectionPoint.EntityData.Children = make(map[string]types.YChild)
    selectionPoint.EntityData.Leafs = make(map[string]types.YLeaf)
    selectionPoint.EntityData.Leafs["selection-point-type"] = types.YLeaf{"SelectionPointType", selectionPoint.SelectionPointType}
    selectionPoint.EntityData.Leafs["selection-point-description"] = types.YLeaf{"SelectionPointDescription", selectionPoint.SelectionPointDescription}
    selectionPoint.EntityData.Leafs["node"] = types.YLeaf{"Node", selectionPoint.Node}
    return &(selectionPoint.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource
// Original source
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SourceClass. The type is FsyncBagSourceClass.
    SourceClass interface{}

    // Ethernet interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    EthernetInterface interface{}

    // SONET interfaces. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SonetInterface interface{}

    // Internal Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // PTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    PtpNode interface{}

    // Satellite Access Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SatelliteAccessInterface interface{}

    // NTP Clock Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NtpNode interface{}

    // Clock ID.
    ClockId FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId
}

func (originalSource *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource) GetEntityData() *types.CommonEntityData {
    originalSource.EntityData.YFilter = originalSource.YFilter
    originalSource.EntityData.YangName = "original-source"
    originalSource.EntityData.BundleName = "cisco_ios_xr"
    originalSource.EntityData.ParentYangName = "selection-point-input"
    originalSource.EntityData.SegmentPath = "original-source"
    originalSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    originalSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    originalSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    originalSource.EntityData.Children = make(map[string]types.YChild)
    originalSource.EntityData.Children["clock-id"] = types.YChild{"ClockId", &originalSource.ClockId}
    originalSource.EntityData.Leafs = make(map[string]types.YLeaf)
    originalSource.EntityData.Leafs["source-class"] = types.YLeaf{"SourceClass", originalSource.SourceClass}
    originalSource.EntityData.Leafs["ethernet-interface"] = types.YLeaf{"EthernetInterface", originalSource.EthernetInterface}
    originalSource.EntityData.Leafs["sonet-interface"] = types.YLeaf{"SonetInterface", originalSource.SonetInterface}
    originalSource.EntityData.Leafs["node"] = types.YLeaf{"Node", originalSource.Node}
    originalSource.EntityData.Leafs["ptp-node"] = types.YLeaf{"PtpNode", originalSource.PtpNode}
    originalSource.EntityData.Leafs["satellite-access-interface"] = types.YLeaf{"SatelliteAccessInterface", originalSource.SatelliteAccessInterface}
    originalSource.EntityData.Leafs["ntp-node"] = types.YLeaf{"NtpNode", originalSource.NtpNode}
    return &(originalSource.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId
// Clock ID
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Port number. The type is interface{} with range: 0..4294967295.
    Port interface{}
}

func (clockId *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_OriginalSource_ClockId) GetEntityData() *types.CommonEntityData {
    clockId.EntityData.YFilter = clockId.YFilter
    clockId.EntityData.YangName = "clock-id"
    clockId.EntityData.BundleName = "cisco_ios_xr"
    clockId.EntityData.ParentYangName = "original-source"
    clockId.EntityData.SegmentPath = "clock-id"
    clockId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clockId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clockId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clockId.EntityData.Children = make(map[string]types.YChild)
    clockId.EntityData.Leafs = make(map[string]types.YLeaf)
    clockId.EntityData.Leafs["node"] = types.YLeaf{"Node", clockId.Node}
    clockId.EntityData.Leafs["port"] = types.YLeaf{"Port", clockId.Port}
    return &(clockId.EntityData)
}

// FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel
// Quality Level
type FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel struct {
    EntityData types.CommonEntityData
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

func (qualityLevel *FrequencySynchronization_Nodes_Node_SelectionPointInputs_SelectionPointInput_QualityLevel) GetEntityData() *types.CommonEntityData {
    qualityLevel.EntityData.YFilter = qualityLevel.YFilter
    qualityLevel.EntityData.YangName = "quality-level"
    qualityLevel.EntityData.BundleName = "cisco_ios_xr"
    qualityLevel.EntityData.ParentYangName = "selection-point-input"
    qualityLevel.EntityData.SegmentPath = "quality-level"
    qualityLevel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qualityLevel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qualityLevel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qualityLevel.EntityData.Children = make(map[string]types.YChild)
    qualityLevel.EntityData.Leafs = make(map[string]types.YLeaf)
    qualityLevel.EntityData.Leafs["quality-level-option"] = types.YLeaf{"QualityLevelOption", qualityLevel.QualityLevelOption}
    qualityLevel.EntityData.Leafs["option1-value"] = types.YLeaf{"Option1Value", qualityLevel.Option1Value}
    qualityLevel.EntityData.Leafs["option2-generation1-value"] = types.YLeaf{"Option2Generation1Value", qualityLevel.Option2Generation1Value}
    qualityLevel.EntityData.Leafs["option2-generation2-value"] = types.YLeaf{"Option2Generation2Value", qualityLevel.Option2Generation2Value}
    return &(qualityLevel.EntityData)
}

