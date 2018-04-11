// This module contains a collection of YANG definitions
// for Cisco IOS-XR Ethernet-SPAN package operational data.
// 
// This module contains definitions
// for the following management objects:
//   span-monitor-session: Monitor Session operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_span_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_span_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-Ethernet-SPAN-oper span-monitor-session}", reflect.TypeOf(SpanMonitorSession{}))
    ydk.RegisterEntity("Cisco-IOS-XR-Ethernet-SPAN-oper:span-monitor-session", reflect.TypeOf(SpanMonitorSession{}))
}

// MirrorInterval represents Monitor-session mirror intervals
type MirrorInterval string

const (
    // Mirror all packets
    MirrorInterval_mirror_interval_all MirrorInterval = "mirror-interval-all"

    // Mirror Interval 512
    MirrorInterval_mirror_interval512 MirrorInterval = "mirror-interval512"

    // Mirror Interval 1K
    MirrorInterval_mirror_interval1k MirrorInterval = "mirror-interval1k"

    // Mirror Interval 2K
    MirrorInterval_mirror_interval2k MirrorInterval = "mirror-interval2k"

    // Mirror Interval 4K
    MirrorInterval_mirror_interval4k MirrorInterval = "mirror-interval4k"

    // Mirror Interval 8K
    MirrorInterval_mirror_interval8k MirrorInterval = "mirror-interval8k"

    // Mirror Interval 16K
    MirrorInterval_mirror_interval16k MirrorInterval = "mirror-interval16k"
)

// TrafficDirection represents Monitor-session traffic directions
type TrafficDirection string

const (
    // Invalid
    TrafficDirection_invalid TrafficDirection = "invalid"

    // Received
    TrafficDirection_rx_only TrafficDirection = "rx-only"

    // Transmitted
    TrafficDirection_tx_only TrafficDirection = "tx-only"

    // Both
    TrafficDirection_both TrafficDirection = "both"
)

// DestinationClass represents Destination class
type DestinationClass string

const (
    // Destination is an interface
    DestinationClass_interface_class DestinationClass = "interface-class"

    // Destination is a pseudowire
    DestinationClass_pseudowire_class DestinationClass = "pseudowire-class"

    // Destination is a next-hop IPv4 address
    DestinationClass_next_hop_ipv4_class DestinationClass = "next-hop-ipv4-class"

    // Destination is a next-hop IPv6 address
    DestinationClass_next_hop_ipv6_class DestinationClass = "next-hop-ipv6-class"

    // Destination is not specified
    DestinationClass_invalid_class DestinationClass = "invalid-class"
)

// SessionClass represents Session class
type SessionClass string

const (
    // Ethernet mirroring session
    SessionClass_ethernet_class SessionClass = "ethernet-class"

    // IPv4 mirroring session
    SessionClass_ipv4_class SessionClass = "ipv4-class"

    // IPv6 mirroring session
    SessionClass_ipv6_class SessionClass = "ipv6-class"

    // MPLS-IPv4 mirroring session
    SessionClass_mplsipv4_class SessionClass = "mplsipv4-class"

    // MPLS-IPv6 mirroring session
    SessionClass_mplsipv6_class SessionClass = "mplsipv6-class"

    // Invalid session class
    SessionClass_invalid_class SessionClass = "invalid-class"
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

// SpanMonitorSession
// Monitor Session operational data
type SpanMonitorSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global operational data.
    Global SpanMonitorSession_Global

    // Node table for node-specific operational data.
    Nodes SpanMonitorSession_Nodes
}

func (spanMonitorSession *SpanMonitorSession) GetEntityData() *types.CommonEntityData {
    spanMonitorSession.EntityData.YFilter = spanMonitorSession.YFilter
    spanMonitorSession.EntityData.YangName = "span-monitor-session"
    spanMonitorSession.EntityData.BundleName = "cisco_ios_xr"
    spanMonitorSession.EntityData.ParentYangName = "Cisco-IOS-XR-Ethernet-SPAN-oper"
    spanMonitorSession.EntityData.SegmentPath = "Cisco-IOS-XR-Ethernet-SPAN-oper:span-monitor-session"
    spanMonitorSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMonitorSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMonitorSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMonitorSession.EntityData.Children = make(map[string]types.YChild)
    spanMonitorSession.EntityData.Children["global"] = types.YChild{"Global", &spanMonitorSession.Global}
    spanMonitorSession.EntityData.Children["nodes"] = types.YChild{"Nodes", &spanMonitorSession.Nodes}
    spanMonitorSession.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spanMonitorSession.EntityData)
}

// SpanMonitorSession_Global
// Global operational data
type SpanMonitorSession_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of statistics for source interfaces.
    Statistics SpanMonitorSession_Global_Statistics

    // Global Monitor Sessions table.
    GlobalSessions SpanMonitorSession_Global_GlobalSessions
}

func (global *SpanMonitorSession_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "span-monitor-session"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["statistics"] = types.YChild{"Statistics", &global.Statistics}
    global.EntityData.Children["global-sessions"] = types.YChild{"GlobalSessions", &global.GlobalSessions}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(global.EntityData)
}

// SpanMonitorSession_Global_Statistics
// Table of statistics for source interfaces
type SpanMonitorSession_Global_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a particular source interface. The type is slice of
    // SpanMonitorSession_Global_Statistics_Statistic.
    Statistic []SpanMonitorSession_Global_Statistics_Statistic
}

func (statistics *SpanMonitorSession_Global_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "global"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["statistic"] = types.YChild{"Statistic", nil}
    for i := range statistics.Statistic {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Statistic[i])] = types.YChild{"Statistic", &statistics.Statistic[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// SpanMonitorSession_Global_Statistics_Statistic
// Statistics for a particular source interface
type SpanMonitorSession_Global_Statistics_Statistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Name. The type is string with length:
    // 1..79.
    Session interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // RX Packets Mirrored. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPacketsMirrored interface{}

    // RX Octets Mirrored. The type is interface{} with range:
    // 0..18446744073709551615.
    RxOctetsMirrored interface{}

    // TX Packets Mirrored. The type is interface{} with range:
    // 0..18446744073709551615.
    TxPacketsMirrored interface{}

    // TX Octets Mirrored. The type is interface{} with range:
    // 0..18446744073709551615.
    TxOctetsMirrored interface{}

    // Packets Not Mirrored. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsNotMirrored interface{}

    // Octets Not Mirrored. The type is interface{} with range:
    // 0..18446744073709551615.
    OctetsNotMirrored interface{}
}

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetEntityData() *types.CommonEntityData {
    statistic.EntityData.YFilter = statistic.YFilter
    statistic.EntityData.YangName = "statistic"
    statistic.EntityData.BundleName = "cisco_ios_xr"
    statistic.EntityData.ParentYangName = "statistics"
    statistic.EntityData.SegmentPath = "statistic" + "[session='" + fmt.Sprintf("%v", statistic.Session) + "']" + "[interface='" + fmt.Sprintf("%v", statistic.Interface_) + "']"
    statistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistic.EntityData.Children = make(map[string]types.YChild)
    statistic.EntityData.Leafs = make(map[string]types.YLeaf)
    statistic.EntityData.Leafs["session"] = types.YLeaf{"Session", statistic.Session}
    statistic.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", statistic.Interface_}
    statistic.EntityData.Leafs["rx-packets-mirrored"] = types.YLeaf{"RxPacketsMirrored", statistic.RxPacketsMirrored}
    statistic.EntityData.Leafs["rx-octets-mirrored"] = types.YLeaf{"RxOctetsMirrored", statistic.RxOctetsMirrored}
    statistic.EntityData.Leafs["tx-packets-mirrored"] = types.YLeaf{"TxPacketsMirrored", statistic.TxPacketsMirrored}
    statistic.EntityData.Leafs["tx-octets-mirrored"] = types.YLeaf{"TxOctetsMirrored", statistic.TxOctetsMirrored}
    statistic.EntityData.Leafs["packets-not-mirrored"] = types.YLeaf{"PacketsNotMirrored", statistic.PacketsNotMirrored}
    statistic.EntityData.Leafs["octets-not-mirrored"] = types.YLeaf{"OctetsNotMirrored", statistic.OctetsNotMirrored}
    return &(statistic.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions
// Global Monitor Sessions table
type SpanMonitorSession_Global_GlobalSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a globally-configured monitor session. The type is slice
    // of SpanMonitorSession_Global_GlobalSessions_GlobalSession.
    GlobalSession []SpanMonitorSession_Global_GlobalSessions_GlobalSession
}

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetEntityData() *types.CommonEntityData {
    globalSessions.EntityData.YFilter = globalSessions.YFilter
    globalSessions.EntityData.YangName = "global-sessions"
    globalSessions.EntityData.BundleName = "cisco_ios_xr"
    globalSessions.EntityData.ParentYangName = "global"
    globalSessions.EntityData.SegmentPath = "global-sessions"
    globalSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalSessions.EntityData.Children = make(map[string]types.YChild)
    globalSessions.EntityData.Children["global-session"] = types.YChild{"GlobalSession", nil}
    for i := range globalSessions.GlobalSession {
        globalSessions.EntityData.Children[types.GetSegmentPath(&globalSessions.GlobalSession[i])] = types.YChild{"GlobalSession", &globalSessions.GlobalSession[i]}
    }
    globalSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalSessions.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession
// Information about a globally-configured
// monitor session
type SpanMonitorSession_Global_GlobalSessions_GlobalSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Name. The type is string with length:
    // 1..79.
    Session interface{}

    // Session Name. The type is string.
    Name interface{}

    // Session class. The type is SessionClass.
    SessionClass interface{}

    // Numerical ID assigned to session. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // Last error observed for the destination . The type is interface{} with
    // range: 0..4294967295.
    DestinationError interface{}

    // Destination interface name (deprecated by DestinationData, invalid for
    // pseudowires). The type is string.
    DestinationInterfaceName interface{}

    // Destination interface handle (deprecated by DestinationID, invalid for
    // pseudowires). The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    DestinationInterfaceHandle interface{}

    // Last error observed for the destination interface (deprecated by
    // DestinationError). The type is interface{} with range: 0..4294967295.
    InterfaceError interface{}

    // Destination data.
    DestinationData SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData

    // Destination ID.
    DestinationId SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId
}

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetEntityData() *types.CommonEntityData {
    globalSession.EntityData.YFilter = globalSession.YFilter
    globalSession.EntityData.YangName = "global-session"
    globalSession.EntityData.BundleName = "cisco_ios_xr"
    globalSession.EntityData.ParentYangName = "global-sessions"
    globalSession.EntityData.SegmentPath = "global-session" + "[session='" + fmt.Sprintf("%v", globalSession.Session) + "']"
    globalSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalSession.EntityData.Children = make(map[string]types.YChild)
    globalSession.EntityData.Children["destination-data"] = types.YChild{"DestinationData", &globalSession.DestinationData}
    globalSession.EntityData.Children["destination-id"] = types.YChild{"DestinationId", &globalSession.DestinationId}
    globalSession.EntityData.Leafs = make(map[string]types.YLeaf)
    globalSession.EntityData.Leafs["session"] = types.YLeaf{"Session", globalSession.Session}
    globalSession.EntityData.Leafs["name"] = types.YLeaf{"Name", globalSession.Name}
    globalSession.EntityData.Leafs["session-class"] = types.YLeaf{"SessionClass", globalSession.SessionClass}
    globalSession.EntityData.Leafs["id"] = types.YLeaf{"Id", globalSession.Id}
    globalSession.EntityData.Leafs["destination-error"] = types.YLeaf{"DestinationError", globalSession.DestinationError}
    globalSession.EntityData.Leafs["destination-interface-name"] = types.YLeaf{"DestinationInterfaceName", globalSession.DestinationInterfaceName}
    globalSession.EntityData.Leafs["destination-interface-handle"] = types.YLeaf{"DestinationInterfaceHandle", globalSession.DestinationInterfaceHandle}
    globalSession.EntityData.Leafs["interface-error"] = types.YLeaf{"InterfaceError", globalSession.InterfaceError}
    return &(globalSession.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData
// Destination data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Invalid Parameter. The type is interface{} with range: 0..4294967295.
    InvalidValue interface{}

    // Interface data.
    InterfaceData SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData

    // Pseudowire data.
    PseudowireData SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData

    // Next-hop IPv4 data.
    NextHopIpv4Data SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data

    // Next-hop IPv6 data.
    NextHopIpv6Data SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data
}

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetEntityData() *types.CommonEntityData {
    destinationData.EntityData.YFilter = destinationData.YFilter
    destinationData.EntityData.YangName = "destination-data"
    destinationData.EntityData.BundleName = "cisco_ios_xr"
    destinationData.EntityData.ParentYangName = "global-session"
    destinationData.EntityData.SegmentPath = "destination-data"
    destinationData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationData.EntityData.Children = make(map[string]types.YChild)
    destinationData.EntityData.Children["interface-data"] = types.YChild{"InterfaceData", &destinationData.InterfaceData}
    destinationData.EntityData.Children["pseudowire-data"] = types.YChild{"PseudowireData", &destinationData.PseudowireData}
    destinationData.EntityData.Children["next-hop-ipv4-data"] = types.YChild{"NextHopIpv4Data", &destinationData.NextHopIpv4Data}
    destinationData.EntityData.Children["next-hop-ipv6-data"] = types.YChild{"NextHopIpv6Data", &destinationData.NextHopIpv6Data}
    destinationData.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationData.EntityData.Leafs["destination-class"] = types.YLeaf{"DestinationClass", destinationData.DestinationClass}
    destinationData.EntityData.Leafs["invalid-value"] = types.YLeaf{"InvalidValue", destinationData.InvalidValue}
    return &(destinationData.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData
// Interface data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Interface State. The type is ImStateEnum.
    InterfaceState interface{}
}

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetEntityData() *types.CommonEntityData {
    interfaceData.EntityData.YFilter = interfaceData.YFilter
    interfaceData.EntityData.YangName = "interface-data"
    interfaceData.EntityData.BundleName = "cisco_ios_xr"
    interfaceData.EntityData.ParentYangName = "destination-data"
    interfaceData.EntityData.SegmentPath = "interface-data"
    interfaceData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceData.EntityData.Children = make(map[string]types.YChild)
    interfaceData.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceData.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceData.InterfaceName}
    interfaceData.EntityData.Leafs["interface-state"] = types.YLeaf{"InterfaceState", interfaceData.InterfaceState}
    return &(interfaceData.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData
// Pseudowire data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire Name. The type is string.
    PseudowireName interface{}

    // Pseudowire State. The type is bool.
    PseudowireIsUp interface{}
}

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetEntityData() *types.CommonEntityData {
    pseudowireData.EntityData.YFilter = pseudowireData.YFilter
    pseudowireData.EntityData.YangName = "pseudowire-data"
    pseudowireData.EntityData.BundleName = "cisco_ios_xr"
    pseudowireData.EntityData.ParentYangName = "destination-data"
    pseudowireData.EntityData.SegmentPath = "pseudowire-data"
    pseudowireData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pseudowireData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pseudowireData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pseudowireData.EntityData.Children = make(map[string]types.YChild)
    pseudowireData.EntityData.Leafs = make(map[string]types.YLeaf)
    pseudowireData.EntityData.Leafs["pseudowire-name"] = types.YLeaf{"PseudowireName", pseudowireData.PseudowireName}
    pseudowireData.EntityData.Leafs["pseudowire-is-up"] = types.YLeaf{"PseudowireIsUp", pseudowireData.PseudowireIsUp}
    return &(pseudowireData.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data
// Next-hop IPv4 data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Address is reachable. The type is bool.
    AddressIsReachable interface{}
}

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetEntityData() *types.CommonEntityData {
    nextHopIpv4Data.EntityData.YFilter = nextHopIpv4Data.YFilter
    nextHopIpv4Data.EntityData.YangName = "next-hop-ipv4-data"
    nextHopIpv4Data.EntityData.BundleName = "cisco_ios_xr"
    nextHopIpv4Data.EntityData.ParentYangName = "destination-data"
    nextHopIpv4Data.EntityData.SegmentPath = "next-hop-ipv4-data"
    nextHopIpv4Data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHopIpv4Data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHopIpv4Data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHopIpv4Data.EntityData.Children = make(map[string]types.YChild)
    nextHopIpv4Data.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHopIpv4Data.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", nextHopIpv4Data.Ipv4Address}
    nextHopIpv4Data.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", nextHopIpv4Data.VrfName}
    nextHopIpv4Data.EntityData.Leafs["address-is-reachable"] = types.YLeaf{"AddressIsReachable", nextHopIpv4Data.AddressIsReachable}
    return &(nextHopIpv4Data.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data
// Next-hop IPv6 data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Address is reachable. The type is bool.
    AddressIsReachable interface{}
}

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetEntityData() *types.CommonEntityData {
    nextHopIpv6Data.EntityData.YFilter = nextHopIpv6Data.YFilter
    nextHopIpv6Data.EntityData.YangName = "next-hop-ipv6-data"
    nextHopIpv6Data.EntityData.BundleName = "cisco_ios_xr"
    nextHopIpv6Data.EntityData.ParentYangName = "destination-data"
    nextHopIpv6Data.EntityData.SegmentPath = "next-hop-ipv6-data"
    nextHopIpv6Data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHopIpv6Data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHopIpv6Data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHopIpv6Data.EntityData.Children = make(map[string]types.YChild)
    nextHopIpv6Data.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHopIpv6Data.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", nextHopIpv6Data.Ipv6Address}
    nextHopIpv6Data.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", nextHopIpv6Data.VrfName}
    nextHopIpv6Data.EntityData.Leafs["address-is-reachable"] = types.YLeaf{"AddressIsReachable", nextHopIpv6Data.AddressIsReachable}
    return &(nextHopIpv6Data.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId
// Destination ID
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Pseudowire XCID. The type is interface{} with range: 0..4294967295.
    PseudowireId interface{}

    // Invalid Parameter. The type is interface{} with range: 0..4294967295.
    InvalidValue interface{}

    // IPv4 address.
    Ipv4AddressAndVrf SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf

    // IPv6 address.
    Ipv6AddressAndVrf SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf
}

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetEntityData() *types.CommonEntityData {
    destinationId.EntityData.YFilter = destinationId.YFilter
    destinationId.EntityData.YangName = "destination-id"
    destinationId.EntityData.BundleName = "cisco_ios_xr"
    destinationId.EntityData.ParentYangName = "global-session"
    destinationId.EntityData.SegmentPath = "destination-id"
    destinationId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationId.EntityData.Children = make(map[string]types.YChild)
    destinationId.EntityData.Children["ipv4-address-and-vrf"] = types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf}
    destinationId.EntityData.Children["ipv6-address-and-vrf"] = types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf}
    destinationId.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationId.EntityData.Leafs["destination-class"] = types.YLeaf{"DestinationClass", destinationId.DestinationClass}
    destinationId.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", destinationId.Interface_}
    destinationId.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", destinationId.PseudowireId}
    destinationId.EntityData.Leafs["invalid-value"] = types.YLeaf{"InvalidValue", destinationId.InvalidValue}
    return &(destinationId.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv4AddressAndVrf.EntityData.YFilter = ipv4AddressAndVrf.YFilter
    ipv4AddressAndVrf.EntityData.YangName = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv4AddressAndVrf.EntityData.SegmentPath = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv4AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4AddressAndVrf.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address}
    ipv4AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName}
    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv6AddressAndVrf.EntityData.YFilter = ipv6AddressAndVrf.YFilter
    ipv6AddressAndVrf.EntityData.YangName = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv6AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv6AddressAndVrf.EntityData.SegmentPath = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv6AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6AddressAndVrf.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address}
    ipv6AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName}
    return &(ipv6AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes
// Node table for node-specific operational data
type SpanMonitorSession_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // SpanMonitorSession_Nodes_Node.
    Node []SpanMonitorSession_Nodes_Node
}

func (nodes *SpanMonitorSession_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "span-monitor-session"
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

// SpanMonitorSession_Nodes_Node
// Node-specific data for a particular node
type SpanMonitorSession_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Node interface{}

    // Table of source interfaces configured as attached to a session.
    Attachments SpanMonitorSession_Nodes_Node_Attachments

    // Table of sessions set up in the hardware.  When all sessions are operating
    // correctly the entries in this table should match those entries in
    // GlobalSessionTable that have a destination configured.
    HardwareSessions SpanMonitorSession_Nodes_Node_HardwareSessions

    // Table of source interfaces set up in the hardware.  The entries in this
    // table should match the entries in AttachmentTable when all sessions are
    // operating correctly.
    Interfaces SpanMonitorSession_Nodes_Node_Interfaces
}

func (node *SpanMonitorSession_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["attachments"] = types.YChild{"Attachments", &node.Attachments}
    node.EntityData.Children["hardware-sessions"] = types.YChild{"HardwareSessions", &node.HardwareSessions}
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node"] = types.YLeaf{"Node", node.Node}
    return &(node.EntityData)
}

// SpanMonitorSession_Nodes_Node_Attachments
// Table of source interfaces configured as
// attached to a session
type SpanMonitorSession_Nodes_Node_Attachments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular source interface configured as attached to
    // monitor session. The type is slice of
    // SpanMonitorSession_Nodes_Node_Attachments_Attachment.
    Attachment []SpanMonitorSession_Nodes_Node_Attachments_Attachment
}

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetEntityData() *types.CommonEntityData {
    attachments.EntityData.YFilter = attachments.YFilter
    attachments.EntityData.YangName = "attachments"
    attachments.EntityData.BundleName = "cisco_ios_xr"
    attachments.EntityData.ParentYangName = "node"
    attachments.EntityData.SegmentPath = "attachments"
    attachments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachments.EntityData.Children = make(map[string]types.YChild)
    attachments.EntityData.Children["attachment"] = types.YChild{"Attachment", nil}
    for i := range attachments.Attachment {
        attachments.EntityData.Children[types.GetSegmentPath(&attachments.Attachment[i])] = types.YChild{"Attachment", &attachments.Attachment[i]}
    }
    attachments.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attachments.EntityData)
}

// SpanMonitorSession_Nodes_Node_Attachments_Attachment
// Information about a particular source
// interface configured as attached to monitor
// session
type SpanMonitorSession_Nodes_Node_Attachments_Attachment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Name. The type is string with length:
    // 1..79.
    Session interface{}

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Session Name. The type is string.
    Name interface{}

    // Local attachment class. The type is SessionClass.
    LocalClass interface{}

    // Numerical ID assigned to session. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // Global session class. The type is SessionClass.
    GlobalClass interface{}

    // The Session is configured globally. The type is bool.
    SessionIsConfigured interface{}

    // Source interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}

    // Source interface state. The type is ImStateEnum.
    SourceInterfaceState interface{}

    // Last error returned from PFI for this interface. The type is interface{}
    // with range: 0..4294967295.
    PfiError interface{}

    // The destination PW type is not supported. The type is bool.
    DestPwTypeNotSupported interface{}

    // This source interface is a destination for another monitor-session. The
    // type is bool.
    SourceInterfaceIsADestination interface{}

    // Destination interface (deprecated by DestinationID, invalid for
    // pseudowires). The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    DestinationInterface interface{}

    // Traffic mirroring direction (deprecated by TrafficParameters). The type is
    // TrafficDirection.
    TrafficDirection interface{}

    // Traffic mirroring parameters.
    TrafficParameters SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters

    // Destination ID.
    DestinationId SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId
}

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetEntityData() *types.CommonEntityData {
    attachment.EntityData.YFilter = attachment.YFilter
    attachment.EntityData.YangName = "attachment"
    attachment.EntityData.BundleName = "cisco_ios_xr"
    attachment.EntityData.ParentYangName = "attachments"
    attachment.EntityData.SegmentPath = "attachment" + "[session='" + fmt.Sprintf("%v", attachment.Session) + "']" + "[interface='" + fmt.Sprintf("%v", attachment.Interface_) + "']"
    attachment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachment.EntityData.Children = make(map[string]types.YChild)
    attachment.EntityData.Children["traffic-parameters"] = types.YChild{"TrafficParameters", &attachment.TrafficParameters}
    attachment.EntityData.Children["destination-id"] = types.YChild{"DestinationId", &attachment.DestinationId}
    attachment.EntityData.Leafs = make(map[string]types.YLeaf)
    attachment.EntityData.Leafs["session"] = types.YLeaf{"Session", attachment.Session}
    attachment.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", attachment.Interface_}
    attachment.EntityData.Leafs["name"] = types.YLeaf{"Name", attachment.Name}
    attachment.EntityData.Leafs["local-class"] = types.YLeaf{"LocalClass", attachment.LocalClass}
    attachment.EntityData.Leafs["id"] = types.YLeaf{"Id", attachment.Id}
    attachment.EntityData.Leafs["global-class"] = types.YLeaf{"GlobalClass", attachment.GlobalClass}
    attachment.EntityData.Leafs["session-is-configured"] = types.YLeaf{"SessionIsConfigured", attachment.SessionIsConfigured}
    attachment.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", attachment.SourceInterface}
    attachment.EntityData.Leafs["source-interface-state"] = types.YLeaf{"SourceInterfaceState", attachment.SourceInterfaceState}
    attachment.EntityData.Leafs["pfi-error"] = types.YLeaf{"PfiError", attachment.PfiError}
    attachment.EntityData.Leafs["dest-pw-type-not-supported"] = types.YLeaf{"DestPwTypeNotSupported", attachment.DestPwTypeNotSupported}
    attachment.EntityData.Leafs["source-interface-is-a-destination"] = types.YLeaf{"SourceInterfaceIsADestination", attachment.SourceInterfaceIsADestination}
    attachment.EntityData.Leafs["destination-interface"] = types.YLeaf{"DestinationInterface", attachment.DestinationInterface}
    attachment.EntityData.Leafs["traffic-direction"] = types.YLeaf{"TrafficDirection", attachment.TrafficDirection}
    return &(attachment.EntityData)
}

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters
// Traffic mirroring parameters
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Direction. The type is TrafficDirection.
    TrafficDirection interface{}

    // Port level mirroring. The type is bool.
    PortLevel interface{}

    // ACL enabled. The type is bool.
    IsAclEnabled interface{}

    // Number of bytes to mirror. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    MirrorBytes interface{}

    // Interval between mirrored packets. The type is MirrorInterval.
    MirrorInterval interface{}

    // ACL name. The type is string.
    AclName interface{}
}

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetEntityData() *types.CommonEntityData {
    trafficParameters.EntityData.YFilter = trafficParameters.YFilter
    trafficParameters.EntityData.YangName = "traffic-parameters"
    trafficParameters.EntityData.BundleName = "cisco_ios_xr"
    trafficParameters.EntityData.ParentYangName = "attachment"
    trafficParameters.EntityData.SegmentPath = "traffic-parameters"
    trafficParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficParameters.EntityData.Children = make(map[string]types.YChild)
    trafficParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficParameters.EntityData.Leafs["traffic-direction"] = types.YLeaf{"TrafficDirection", trafficParameters.TrafficDirection}
    trafficParameters.EntityData.Leafs["port-level"] = types.YLeaf{"PortLevel", trafficParameters.PortLevel}
    trafficParameters.EntityData.Leafs["is-acl-enabled"] = types.YLeaf{"IsAclEnabled", trafficParameters.IsAclEnabled}
    trafficParameters.EntityData.Leafs["mirror-bytes"] = types.YLeaf{"MirrorBytes", trafficParameters.MirrorBytes}
    trafficParameters.EntityData.Leafs["mirror-interval"] = types.YLeaf{"MirrorInterval", trafficParameters.MirrorInterval}
    trafficParameters.EntityData.Leafs["acl-name"] = types.YLeaf{"AclName", trafficParameters.AclName}
    return &(trafficParameters.EntityData)
}

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId
// Destination ID
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Pseudowire XCID. The type is interface{} with range: 0..4294967295.
    PseudowireId interface{}

    // Invalid Parameter. The type is interface{} with range: 0..4294967295.
    InvalidValue interface{}

    // IPv4 address.
    Ipv4AddressAndVrf SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf

    // IPv6 address.
    Ipv6AddressAndVrf SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf
}

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetEntityData() *types.CommonEntityData {
    destinationId.EntityData.YFilter = destinationId.YFilter
    destinationId.EntityData.YangName = "destination-id"
    destinationId.EntityData.BundleName = "cisco_ios_xr"
    destinationId.EntityData.ParentYangName = "attachment"
    destinationId.EntityData.SegmentPath = "destination-id"
    destinationId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationId.EntityData.Children = make(map[string]types.YChild)
    destinationId.EntityData.Children["ipv4-address-and-vrf"] = types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf}
    destinationId.EntityData.Children["ipv6-address-and-vrf"] = types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf}
    destinationId.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationId.EntityData.Leafs["destination-class"] = types.YLeaf{"DestinationClass", destinationId.DestinationClass}
    destinationId.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", destinationId.Interface_}
    destinationId.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", destinationId.PseudowireId}
    destinationId.EntityData.Leafs["invalid-value"] = types.YLeaf{"InvalidValue", destinationId.InvalidValue}
    return &(destinationId.EntityData)
}

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv4AddressAndVrf.EntityData.YFilter = ipv4AddressAndVrf.YFilter
    ipv4AddressAndVrf.EntityData.YangName = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv4AddressAndVrf.EntityData.SegmentPath = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv4AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4AddressAndVrf.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address}
    ipv4AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName}
    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv6AddressAndVrf.EntityData.YFilter = ipv6AddressAndVrf.YFilter
    ipv6AddressAndVrf.EntityData.YangName = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv6AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv6AddressAndVrf.EntityData.SegmentPath = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv6AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6AddressAndVrf.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address}
    ipv6AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName}
    return &(ipv6AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_HardwareSessions
// Table of sessions set up in the hardware. 
// When all sessions are operating correctly the
// entries in this table should match those
// entries in GlobalSessionTable that have a
// destination configured
type SpanMonitorSession_Nodes_Node_HardwareSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular session that is set up in the hardware. The
    // type is slice of
    // SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession.
    HardwareSession []SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession
}

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetEntityData() *types.CommonEntityData {
    hardwareSessions.EntityData.YFilter = hardwareSessions.YFilter
    hardwareSessions.EntityData.YangName = "hardware-sessions"
    hardwareSessions.EntityData.BundleName = "cisco_ios_xr"
    hardwareSessions.EntityData.ParentYangName = "node"
    hardwareSessions.EntityData.SegmentPath = "hardware-sessions"
    hardwareSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareSessions.EntityData.Children = make(map[string]types.YChild)
    hardwareSessions.EntityData.Children["hardware-session"] = types.YChild{"HardwareSession", nil}
    for i := range hardwareSessions.HardwareSession {
        hardwareSessions.EntityData.Children[types.GetSegmentPath(&hardwareSessions.HardwareSession[i])] = types.YChild{"HardwareSession", &hardwareSessions.HardwareSession[i]}
    }
    hardwareSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hardwareSessions.EntityData)
}

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession
// Information about a particular session that
// is set up in the hardware
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sesssion class. The type is SpanSessionClass.
    SessionClass interface{}

    // Session ID. The type is interface{} with range: -2147483648..2147483647.
    SessionId interface{}

    // Assigned numerical ID for this session. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // Configured Session Name. The type is string.
    Name interface{}

    // Session class. The type is SessionClass.
    SessionClassXr interface{}

    // Destination interface (deprecated by DestinationID, invalid for
    // pseudowires). The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    DestinationInterface interface{}

    // Last error observed for this session while programming the hardware. The
    // type is interface{} with range: 0..4294967295.
    PlatformError interface{}

    // Destination ID.
    DestinationId SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId
}

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetEntityData() *types.CommonEntityData {
    hardwareSession.EntityData.YFilter = hardwareSession.YFilter
    hardwareSession.EntityData.YangName = "hardware-session"
    hardwareSession.EntityData.BundleName = "cisco_ios_xr"
    hardwareSession.EntityData.ParentYangName = "hardware-sessions"
    hardwareSession.EntityData.SegmentPath = "hardware-session"
    hardwareSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareSession.EntityData.Children = make(map[string]types.YChild)
    hardwareSession.EntityData.Children["destination-id"] = types.YChild{"DestinationId", &hardwareSession.DestinationId}
    hardwareSession.EntityData.Leafs = make(map[string]types.YLeaf)
    hardwareSession.EntityData.Leafs["session-class"] = types.YLeaf{"SessionClass", hardwareSession.SessionClass}
    hardwareSession.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", hardwareSession.SessionId}
    hardwareSession.EntityData.Leafs["id"] = types.YLeaf{"Id", hardwareSession.Id}
    hardwareSession.EntityData.Leafs["name"] = types.YLeaf{"Name", hardwareSession.Name}
    hardwareSession.EntityData.Leafs["session-class-xr"] = types.YLeaf{"SessionClassXr", hardwareSession.SessionClassXr}
    hardwareSession.EntityData.Leafs["destination-interface"] = types.YLeaf{"DestinationInterface", hardwareSession.DestinationInterface}
    hardwareSession.EntityData.Leafs["platform-error"] = types.YLeaf{"PlatformError", hardwareSession.PlatformError}
    return &(hardwareSession.EntityData)
}

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId
// Destination ID
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Pseudowire XCID. The type is interface{} with range: 0..4294967295.
    PseudowireId interface{}

    // Invalid Parameter. The type is interface{} with range: 0..4294967295.
    InvalidValue interface{}

    // IPv4 address.
    Ipv4AddressAndVrf SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf

    // IPv6 address.
    Ipv6AddressAndVrf SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf
}

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetEntityData() *types.CommonEntityData {
    destinationId.EntityData.YFilter = destinationId.YFilter
    destinationId.EntityData.YangName = "destination-id"
    destinationId.EntityData.BundleName = "cisco_ios_xr"
    destinationId.EntityData.ParentYangName = "hardware-session"
    destinationId.EntityData.SegmentPath = "destination-id"
    destinationId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationId.EntityData.Children = make(map[string]types.YChild)
    destinationId.EntityData.Children["ipv4-address-and-vrf"] = types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf}
    destinationId.EntityData.Children["ipv6-address-and-vrf"] = types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf}
    destinationId.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationId.EntityData.Leafs["destination-class"] = types.YLeaf{"DestinationClass", destinationId.DestinationClass}
    destinationId.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", destinationId.Interface_}
    destinationId.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", destinationId.PseudowireId}
    destinationId.EntityData.Leafs["invalid-value"] = types.YLeaf{"InvalidValue", destinationId.InvalidValue}
    return &(destinationId.EntityData)
}

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv4AddressAndVrf.EntityData.YFilter = ipv4AddressAndVrf.YFilter
    ipv4AddressAndVrf.EntityData.YangName = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv4AddressAndVrf.EntityData.SegmentPath = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv4AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4AddressAndVrf.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address}
    ipv4AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName}
    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv6AddressAndVrf.EntityData.YFilter = ipv6AddressAndVrf.YFilter
    ipv6AddressAndVrf.EntityData.YangName = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv6AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv6AddressAndVrf.EntityData.SegmentPath = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv6AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6AddressAndVrf.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address}
    ipv6AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName}
    return &(ipv6AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces
// Table of source interfaces set up in the
// hardware.  The entries in this table should
// match the entries in AttachmentTable when all
// sessions are operating correctly
type SpanMonitorSession_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular interface that is set up in the hardware.
    // The type is slice of SpanMonitorSession_Nodes_Node_Interfaces_Interface_.
    Interface_ []SpanMonitorSession_Nodes_Node_Interfaces_Interface
}

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
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

// SpanMonitorSession_Nodes_Node_Interfaces_Interface
// Information about a particular interface that
// is set up in the hardware
type SpanMonitorSession_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Source interface. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}

    // Last error observed for this interface while programming the hardware. The
    // type is interface{} with range: 0..4294967295.
    PlatformError interface{}

    // Destination interface (deprecated by Attachment). The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    DestinationInterface interface{}

    // Traffic mirroring direction (deprecated by Attachment). The type is
    // TrafficDirection.
    TrafficDirection interface{}

    // Destination ID (deprecated by Attachment).
    DestinationId SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId

    // Traffic mirroring parameters (deprecated by Attachment).
    TrafficMirroringParameters SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters

    // Attachment information. The type is slice of
    // SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment.
    Attachment []SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment
}

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface_) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["destination-id"] = types.YChild{"DestinationId", &self.DestinationId}
    self.EntityData.Children["traffic-mirroring-parameters"] = types.YChild{"TrafficMirroringParameters", &self.TrafficMirroringParameters}
    self.EntityData.Children["attachment"] = types.YChild{"Attachment", nil}
    for i := range self.Attachment {
        self.EntityData.Children[types.GetSegmentPath(&self.Attachment[i])] = types.YChild{"Attachment", &self.Attachment[i]}
    }
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
    self.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", self.SourceInterface}
    self.EntityData.Leafs["platform-error"] = types.YLeaf{"PlatformError", self.PlatformError}
    self.EntityData.Leafs["destination-interface"] = types.YLeaf{"DestinationInterface", self.DestinationInterface}
    self.EntityData.Leafs["traffic-direction"] = types.YLeaf{"TrafficDirection", self.TrafficDirection}
    return &(self.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId
// Destination ID (deprecated by Attachment)
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Pseudowire XCID. The type is interface{} with range: 0..4294967295.
    PseudowireId interface{}

    // Invalid Parameter. The type is interface{} with range: 0..4294967295.
    InvalidValue interface{}

    // IPv4 address.
    Ipv4AddressAndVrf SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf

    // IPv6 address.
    Ipv6AddressAndVrf SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetEntityData() *types.CommonEntityData {
    destinationId.EntityData.YFilter = destinationId.YFilter
    destinationId.EntityData.YangName = "destination-id"
    destinationId.EntityData.BundleName = "cisco_ios_xr"
    destinationId.EntityData.ParentYangName = "interface"
    destinationId.EntityData.SegmentPath = "destination-id"
    destinationId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationId.EntityData.Children = make(map[string]types.YChild)
    destinationId.EntityData.Children["ipv4-address-and-vrf"] = types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf}
    destinationId.EntityData.Children["ipv6-address-and-vrf"] = types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf}
    destinationId.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationId.EntityData.Leafs["destination-class"] = types.YLeaf{"DestinationClass", destinationId.DestinationClass}
    destinationId.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", destinationId.Interface_}
    destinationId.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", destinationId.PseudowireId}
    destinationId.EntityData.Leafs["invalid-value"] = types.YLeaf{"InvalidValue", destinationId.InvalidValue}
    return &(destinationId.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv4AddressAndVrf.EntityData.YFilter = ipv4AddressAndVrf.YFilter
    ipv4AddressAndVrf.EntityData.YangName = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv4AddressAndVrf.EntityData.SegmentPath = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv4AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4AddressAndVrf.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address}
    ipv4AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName}
    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv6AddressAndVrf.EntityData.YFilter = ipv6AddressAndVrf.YFilter
    ipv6AddressAndVrf.EntityData.YangName = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv6AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv6AddressAndVrf.EntityData.SegmentPath = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv6AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6AddressAndVrf.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address}
    ipv6AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName}
    return &(ipv6AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters
// Traffic mirroring parameters (deprecated by
// Attachment)
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Direction. The type is TrafficDirection.
    TrafficDirection interface{}

    // Port level mirroring. The type is bool.
    PortLevel interface{}

    // ACL enabled. The type is bool.
    IsAclEnabled interface{}

    // Number of bytes to mirror. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    MirrorBytes interface{}

    // Interval between mirrored packets. The type is MirrorInterval.
    MirrorInterval interface{}

    // ACL name. The type is string.
    AclName interface{}
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetEntityData() *types.CommonEntityData {
    trafficMirroringParameters.EntityData.YFilter = trafficMirroringParameters.YFilter
    trafficMirroringParameters.EntityData.YangName = "traffic-mirroring-parameters"
    trafficMirroringParameters.EntityData.BundleName = "cisco_ios_xr"
    trafficMirroringParameters.EntityData.ParentYangName = "interface"
    trafficMirroringParameters.EntityData.SegmentPath = "traffic-mirroring-parameters"
    trafficMirroringParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficMirroringParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficMirroringParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficMirroringParameters.EntityData.Children = make(map[string]types.YChild)
    trafficMirroringParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficMirroringParameters.EntityData.Leafs["traffic-direction"] = types.YLeaf{"TrafficDirection", trafficMirroringParameters.TrafficDirection}
    trafficMirroringParameters.EntityData.Leafs["port-level"] = types.YLeaf{"PortLevel", trafficMirroringParameters.PortLevel}
    trafficMirroringParameters.EntityData.Leafs["is-acl-enabled"] = types.YLeaf{"IsAclEnabled", trafficMirroringParameters.IsAclEnabled}
    trafficMirroringParameters.EntityData.Leafs["mirror-bytes"] = types.YLeaf{"MirrorBytes", trafficMirroringParameters.MirrorBytes}
    trafficMirroringParameters.EntityData.Leafs["mirror-interval"] = types.YLeaf{"MirrorInterval", trafficMirroringParameters.MirrorInterval}
    trafficMirroringParameters.EntityData.Leafs["acl-name"] = types.YLeaf{"AclName", trafficMirroringParameters.AclName}
    return &(trafficMirroringParameters.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment
// Attachment information
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attachment class. The type is SessionClass.
    Class interface{}

    // Destination ID.
    DestinationId SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId

    // Traffic mirroring parameters.
    TrafficMirroringParameters SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters
}

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetEntityData() *types.CommonEntityData {
    attachment.EntityData.YFilter = attachment.YFilter
    attachment.EntityData.YangName = "attachment"
    attachment.EntityData.BundleName = "cisco_ios_xr"
    attachment.EntityData.ParentYangName = "interface"
    attachment.EntityData.SegmentPath = "attachment"
    attachment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachment.EntityData.Children = make(map[string]types.YChild)
    attachment.EntityData.Children["destination-id"] = types.YChild{"DestinationId", &attachment.DestinationId}
    attachment.EntityData.Children["traffic-mirroring-parameters"] = types.YChild{"TrafficMirroringParameters", &attachment.TrafficMirroringParameters}
    attachment.EntityData.Leafs = make(map[string]types.YLeaf)
    attachment.EntityData.Leafs["class"] = types.YLeaf{"Class", attachment.Class}
    return &(attachment.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId
// Destination ID
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Pseudowire XCID. The type is interface{} with range: 0..4294967295.
    PseudowireId interface{}

    // Invalid Parameter. The type is interface{} with range: 0..4294967295.
    InvalidValue interface{}

    // IPv4 address.
    Ipv4AddressAndVrf SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf

    // IPv6 address.
    Ipv6AddressAndVrf SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetEntityData() *types.CommonEntityData {
    destinationId.EntityData.YFilter = destinationId.YFilter
    destinationId.EntityData.YangName = "destination-id"
    destinationId.EntityData.BundleName = "cisco_ios_xr"
    destinationId.EntityData.ParentYangName = "attachment"
    destinationId.EntityData.SegmentPath = "destination-id"
    destinationId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationId.EntityData.Children = make(map[string]types.YChild)
    destinationId.EntityData.Children["ipv4-address-and-vrf"] = types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf}
    destinationId.EntityData.Children["ipv6-address-and-vrf"] = types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf}
    destinationId.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationId.EntityData.Leafs["destination-class"] = types.YLeaf{"DestinationClass", destinationId.DestinationClass}
    destinationId.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", destinationId.Interface_}
    destinationId.EntityData.Leafs["pseudowire-id"] = types.YLeaf{"PseudowireId", destinationId.PseudowireId}
    destinationId.EntityData.Leafs["invalid-value"] = types.YLeaf{"InvalidValue", destinationId.InvalidValue}
    return &(destinationId.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv4AddressAndVrf.EntityData.YFilter = ipv4AddressAndVrf.YFilter
    ipv4AddressAndVrf.EntityData.YangName = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv4AddressAndVrf.EntityData.SegmentPath = "ipv4-address-and-vrf"
    ipv4AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv4AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4AddressAndVrf.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address}
    ipv4AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName}
    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetEntityData() *types.CommonEntityData {
    ipv6AddressAndVrf.EntityData.YFilter = ipv6AddressAndVrf.YFilter
    ipv6AddressAndVrf.EntityData.YangName = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.BundleName = "cisco_ios_xr"
    ipv6AddressAndVrf.EntityData.ParentYangName = "destination-id"
    ipv6AddressAndVrf.EntityData.SegmentPath = "ipv6-address-and-vrf"
    ipv6AddressAndVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AddressAndVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AddressAndVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AddressAndVrf.EntityData.Children = make(map[string]types.YChild)
    ipv6AddressAndVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6AddressAndVrf.EntityData.Leafs["ipv6-address"] = types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address}
    ipv6AddressAndVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName}
    return &(ipv6AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters
// Traffic mirroring parameters
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Direction. The type is TrafficDirection.
    TrafficDirection interface{}

    // Port level mirroring. The type is bool.
    PortLevel interface{}

    // ACL enabled. The type is bool.
    IsAclEnabled interface{}

    // Number of bytes to mirror. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    MirrorBytes interface{}

    // Interval between mirrored packets. The type is MirrorInterval.
    MirrorInterval interface{}

    // ACL name. The type is string.
    AclName interface{}
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetEntityData() *types.CommonEntityData {
    trafficMirroringParameters.EntityData.YFilter = trafficMirroringParameters.YFilter
    trafficMirroringParameters.EntityData.YangName = "traffic-mirroring-parameters"
    trafficMirroringParameters.EntityData.BundleName = "cisco_ios_xr"
    trafficMirroringParameters.EntityData.ParentYangName = "attachment"
    trafficMirroringParameters.EntityData.SegmentPath = "traffic-mirroring-parameters"
    trafficMirroringParameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficMirroringParameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficMirroringParameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficMirroringParameters.EntityData.Children = make(map[string]types.YChild)
    trafficMirroringParameters.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficMirroringParameters.EntityData.Leafs["traffic-direction"] = types.YLeaf{"TrafficDirection", trafficMirroringParameters.TrafficDirection}
    trafficMirroringParameters.EntityData.Leafs["port-level"] = types.YLeaf{"PortLevel", trafficMirroringParameters.PortLevel}
    trafficMirroringParameters.EntityData.Leafs["is-acl-enabled"] = types.YLeaf{"IsAclEnabled", trafficMirroringParameters.IsAclEnabled}
    trafficMirroringParameters.EntityData.Leafs["mirror-bytes"] = types.YLeaf{"MirrorBytes", trafficMirroringParameters.MirrorBytes}
    trafficMirroringParameters.EntityData.Leafs["mirror-interval"] = types.YLeaf{"MirrorInterval", trafficMirroringParameters.MirrorInterval}
    trafficMirroringParameters.EntityData.Leafs["acl-name"] = types.YLeaf{"AclName", trafficMirroringParameters.AclName}
    return &(trafficMirroringParameters.EntityData)
}

