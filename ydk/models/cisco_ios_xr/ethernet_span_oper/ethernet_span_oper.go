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

    spanMonitorSession.EntityData.Children = types.NewOrderedMap()
    spanMonitorSession.EntityData.Children.Append("global", types.YChild{"Global", &spanMonitorSession.Global})
    spanMonitorSession.EntityData.Children.Append("nodes", types.YChild{"Nodes", &spanMonitorSession.Nodes})
    spanMonitorSession.EntityData.Leafs = types.NewOrderedMap()

    spanMonitorSession.EntityData.YListKeys = []string {}

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

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("statistics", types.YChild{"Statistics", &global.Statistics})
    global.EntityData.Children.Append("global-sessions", types.YChild{"GlobalSessions", &global.GlobalSessions})
    global.EntityData.Leafs = types.NewOrderedMap()

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// SpanMonitorSession_Global_Statistics
// Table of statistics for source interfaces
type SpanMonitorSession_Global_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics for a particular source interface. The type is slice of
    // SpanMonitorSession_Global_Statistics_Statistic.
    Statistic []*SpanMonitorSession_Global_Statistics_Statistic
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

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("statistic", types.YChild{"Statistic", nil})
    for i := range statistics.Statistic {
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.Statistic[i]), types.YChild{"Statistic", statistics.Statistic[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

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
    // [a-zA-Z0-9./-]+.
    Interface interface{}

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
    statistic.EntityData.SegmentPath = "statistic" + types.AddKeyToken(statistic.Session, "session") + types.AddKeyToken(statistic.Interface, "interface")
    statistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistic.EntityData.Children = types.NewOrderedMap()
    statistic.EntityData.Leafs = types.NewOrderedMap()
    statistic.EntityData.Leafs.Append("session", types.YLeaf{"Session", statistic.Session})
    statistic.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", statistic.Interface})
    statistic.EntityData.Leafs.Append("rx-packets-mirrored", types.YLeaf{"RxPacketsMirrored", statistic.RxPacketsMirrored})
    statistic.EntityData.Leafs.Append("rx-octets-mirrored", types.YLeaf{"RxOctetsMirrored", statistic.RxOctetsMirrored})
    statistic.EntityData.Leafs.Append("tx-packets-mirrored", types.YLeaf{"TxPacketsMirrored", statistic.TxPacketsMirrored})
    statistic.EntityData.Leafs.Append("tx-octets-mirrored", types.YLeaf{"TxOctetsMirrored", statistic.TxOctetsMirrored})
    statistic.EntityData.Leafs.Append("packets-not-mirrored", types.YLeaf{"PacketsNotMirrored", statistic.PacketsNotMirrored})
    statistic.EntityData.Leafs.Append("octets-not-mirrored", types.YLeaf{"OctetsNotMirrored", statistic.OctetsNotMirrored})

    statistic.EntityData.YListKeys = []string {"Session", "Interface"}

    return &(statistic.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions
// Global Monitor Sessions table
type SpanMonitorSession_Global_GlobalSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a globally-configured monitor session. The type is slice
    // of SpanMonitorSession_Global_GlobalSessions_GlobalSession.
    GlobalSession []*SpanMonitorSession_Global_GlobalSessions_GlobalSession
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

    globalSessions.EntityData.Children = types.NewOrderedMap()
    globalSessions.EntityData.Children.Append("global-session", types.YChild{"GlobalSession", nil})
    for i := range globalSessions.GlobalSession {
        globalSessions.EntityData.Children.Append(types.GetSegmentPath(globalSessions.GlobalSession[i]), types.YChild{"GlobalSession", globalSessions.GlobalSession[i]})
    }
    globalSessions.EntityData.Leafs = types.NewOrderedMap()

    globalSessions.EntityData.YListKeys = []string {}

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
    // pseudowires). The type is string with pattern: [a-zA-Z0-9./-]+.
    DestinationInterfaceHandle interface{}

    // Last error observed for the destination interface (deprecated by
    // DestinationError). The type is interface{} with range: 0..4294967295.
    InterfaceError interface{}

    // Destination data.
    DestinationData SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData

    // Destination ID.
    DestinationId SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId

    // Inject interface data.
    InjectInterface SpanMonitorSession_Global_GlobalSessions_GlobalSession_InjectInterface
}

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetEntityData() *types.CommonEntityData {
    globalSession.EntityData.YFilter = globalSession.YFilter
    globalSession.EntityData.YangName = "global-session"
    globalSession.EntityData.BundleName = "cisco_ios_xr"
    globalSession.EntityData.ParentYangName = "global-sessions"
    globalSession.EntityData.SegmentPath = "global-session" + types.AddKeyToken(globalSession.Session, "session")
    globalSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalSession.EntityData.Children = types.NewOrderedMap()
    globalSession.EntityData.Children.Append("destination-data", types.YChild{"DestinationData", &globalSession.DestinationData})
    globalSession.EntityData.Children.Append("destination-id", types.YChild{"DestinationId", &globalSession.DestinationId})
    globalSession.EntityData.Children.Append("inject-interface", types.YChild{"InjectInterface", &globalSession.InjectInterface})
    globalSession.EntityData.Leafs = types.NewOrderedMap()
    globalSession.EntityData.Leafs.Append("session", types.YLeaf{"Session", globalSession.Session})
    globalSession.EntityData.Leafs.Append("name", types.YLeaf{"Name", globalSession.Name})
    globalSession.EntityData.Leafs.Append("session-class", types.YLeaf{"SessionClass", globalSession.SessionClass})
    globalSession.EntityData.Leafs.Append("id", types.YLeaf{"Id", globalSession.Id})
    globalSession.EntityData.Leafs.Append("destination-error", types.YLeaf{"DestinationError", globalSession.DestinationError})
    globalSession.EntityData.Leafs.Append("destination-interface-name", types.YLeaf{"DestinationInterfaceName", globalSession.DestinationInterfaceName})
    globalSession.EntityData.Leafs.Append("destination-interface-handle", types.YLeaf{"DestinationInterfaceHandle", globalSession.DestinationInterfaceHandle})
    globalSession.EntityData.Leafs.Append("interface-error", types.YLeaf{"InterfaceError", globalSession.InterfaceError})

    globalSession.EntityData.YListKeys = []string {"Session"}

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

    destinationData.EntityData.Children = types.NewOrderedMap()
    destinationData.EntityData.Children.Append("interface-data", types.YChild{"InterfaceData", &destinationData.InterfaceData})
    destinationData.EntityData.Children.Append("pseudowire-data", types.YChild{"PseudowireData", &destinationData.PseudowireData})
    destinationData.EntityData.Children.Append("next-hop-ipv4-data", types.YChild{"NextHopIpv4Data", &destinationData.NextHopIpv4Data})
    destinationData.EntityData.Children.Append("next-hop-ipv6-data", types.YChild{"NextHopIpv6Data", &destinationData.NextHopIpv6Data})
    destinationData.EntityData.Leafs = types.NewOrderedMap()
    destinationData.EntityData.Leafs.Append("destination-class", types.YLeaf{"DestinationClass", destinationData.DestinationClass})
    destinationData.EntityData.Leafs.Append("invalid-value", types.YLeaf{"InvalidValue", destinationData.InvalidValue})

    destinationData.EntityData.YListKeys = []string {}

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

    interfaceData.EntityData.Children = types.NewOrderedMap()
    interfaceData.EntityData.Leafs = types.NewOrderedMap()
    interfaceData.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceData.InterfaceName})
    interfaceData.EntityData.Leafs.Append("interface-state", types.YLeaf{"InterfaceState", interfaceData.InterfaceState})

    interfaceData.EntityData.YListKeys = []string {}

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

    pseudowireData.EntityData.Children = types.NewOrderedMap()
    pseudowireData.EntityData.Leafs = types.NewOrderedMap()
    pseudowireData.EntityData.Leafs.Append("pseudowire-name", types.YLeaf{"PseudowireName", pseudowireData.PseudowireName})
    pseudowireData.EntityData.Leafs.Append("pseudowire-is-up", types.YLeaf{"PseudowireIsUp", pseudowireData.PseudowireIsUp})

    pseudowireData.EntityData.YListKeys = []string {}

    return &(pseudowireData.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data
// Next-hop IPv4 data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    nextHopIpv4Data.EntityData.Children = types.NewOrderedMap()
    nextHopIpv4Data.EntityData.Leafs = types.NewOrderedMap()
    nextHopIpv4Data.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHopIpv4Data.Ipv4Address})
    nextHopIpv4Data.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", nextHopIpv4Data.VrfName})
    nextHopIpv4Data.EntityData.Leafs.Append("address-is-reachable", types.YLeaf{"AddressIsReachable", nextHopIpv4Data.AddressIsReachable})

    nextHopIpv4Data.EntityData.YListKeys = []string {}

    return &(nextHopIpv4Data.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data
// Next-hop IPv6 data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    nextHopIpv6Data.EntityData.Children = types.NewOrderedMap()
    nextHopIpv6Data.EntityData.Leafs = types.NewOrderedMap()
    nextHopIpv6Data.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHopIpv6Data.Ipv6Address})
    nextHopIpv6Data.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", nextHopIpv6Data.VrfName})
    nextHopIpv6Data.EntityData.Leafs.Append("address-is-reachable", types.YLeaf{"AddressIsReachable", nextHopIpv6Data.AddressIsReachable})

    nextHopIpv6Data.EntityData.YListKeys = []string {}

    return &(nextHopIpv6Data.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId
// Destination ID
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

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

    destinationId.EntityData.Children = types.NewOrderedMap()
    destinationId.EntityData.Children.Append("ipv4-address-and-vrf", types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf})
    destinationId.EntityData.Children.Append("ipv6-address-and-vrf", types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf})
    destinationId.EntityData.Leafs = types.NewOrderedMap()
    destinationId.EntityData.Leafs.Append("destination-class", types.YLeaf{"DestinationClass", destinationId.DestinationClass})
    destinationId.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", destinationId.Interface})
    destinationId.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", destinationId.PseudowireId})
    destinationId.EntityData.Leafs.Append("invalid-value", types.YLeaf{"InvalidValue", destinationId.InvalidValue})

    destinationId.EntityData.YListKeys = []string {}

    return &(destinationId.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    ipv4AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address})
    ipv4AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName})

    ipv4AddressAndVrf.EntityData.YListKeys = []string {}

    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    ipv6AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address})
    ipv6AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName})

    ipv6AddressAndVrf.EntityData.YListKeys = []string {}

    return &(ipv6AddressAndVrf.EntityData)
}

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_InjectInterface
// Inject interface data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_InjectInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    Name interface{}
}

func (injectInterface *SpanMonitorSession_Global_GlobalSessions_GlobalSession_InjectInterface) GetEntityData() *types.CommonEntityData {
    injectInterface.EntityData.YFilter = injectInterface.YFilter
    injectInterface.EntityData.YangName = "inject-interface"
    injectInterface.EntityData.BundleName = "cisco_ios_xr"
    injectInterface.EntityData.ParentYangName = "global-session"
    injectInterface.EntityData.SegmentPath = "inject-interface"
    injectInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    injectInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    injectInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    injectInterface.EntityData.Children = types.NewOrderedMap()
    injectInterface.EntityData.Leafs = types.NewOrderedMap()
    injectInterface.EntityData.Leafs.Append("name", types.YLeaf{"Name", injectInterface.Name})

    injectInterface.EntityData.YListKeys = []string {}

    return &(injectInterface.EntityData)
}

// SpanMonitorSession_Nodes
// Node table for node-specific operational data
type SpanMonitorSession_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // SpanMonitorSession_Nodes_Node.
    Node []*SpanMonitorSession_Nodes_Node
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

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// SpanMonitorSession_Nodes_Node
// Node-specific data for a particular node
type SpanMonitorSession_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
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
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Node, "node")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("attachments", types.YChild{"Attachments", &node.Attachments})
    node.EntityData.Children.Append("hardware-sessions", types.YChild{"HardwareSessions", &node.HardwareSessions})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node", types.YLeaf{"Node", node.Node})

    node.EntityData.YListKeys = []string {"Node"}

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
    Attachment []*SpanMonitorSession_Nodes_Node_Attachments_Attachment
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

    attachments.EntityData.Children = types.NewOrderedMap()
    attachments.EntityData.Children.Append("attachment", types.YChild{"Attachment", nil})
    for i := range attachments.Attachment {
        attachments.EntityData.Children.Append(types.GetSegmentPath(attachments.Attachment[i]), types.YChild{"Attachment", attachments.Attachment[i]})
    }
    attachments.EntityData.Leafs = types.NewOrderedMap()

    attachments.EntityData.YListKeys = []string {}

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
    // [a-zA-Z0-9./-]+.
    Interface interface{}

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

    // Source interface. The type is string with pattern: [a-zA-Z0-9./-]+.
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
    // pseudowires). The type is string with pattern: [a-zA-Z0-9./-]+.
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
    attachment.EntityData.SegmentPath = "attachment" + types.AddKeyToken(attachment.Session, "session") + types.AddKeyToken(attachment.Interface, "interface")
    attachment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachment.EntityData.Children = types.NewOrderedMap()
    attachment.EntityData.Children.Append("traffic-parameters", types.YChild{"TrafficParameters", &attachment.TrafficParameters})
    attachment.EntityData.Children.Append("destination-id", types.YChild{"DestinationId", &attachment.DestinationId})
    attachment.EntityData.Leafs = types.NewOrderedMap()
    attachment.EntityData.Leafs.Append("session", types.YLeaf{"Session", attachment.Session})
    attachment.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", attachment.Interface})
    attachment.EntityData.Leafs.Append("name", types.YLeaf{"Name", attachment.Name})
    attachment.EntityData.Leafs.Append("local-class", types.YLeaf{"LocalClass", attachment.LocalClass})
    attachment.EntityData.Leafs.Append("id", types.YLeaf{"Id", attachment.Id})
    attachment.EntityData.Leafs.Append("global-class", types.YLeaf{"GlobalClass", attachment.GlobalClass})
    attachment.EntityData.Leafs.Append("session-is-configured", types.YLeaf{"SessionIsConfigured", attachment.SessionIsConfigured})
    attachment.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", attachment.SourceInterface})
    attachment.EntityData.Leafs.Append("source-interface-state", types.YLeaf{"SourceInterfaceState", attachment.SourceInterfaceState})
    attachment.EntityData.Leafs.Append("pfi-error", types.YLeaf{"PfiError", attachment.PfiError})
    attachment.EntityData.Leafs.Append("dest-pw-type-not-supported", types.YLeaf{"DestPwTypeNotSupported", attachment.DestPwTypeNotSupported})
    attachment.EntityData.Leafs.Append("source-interface-is-a-destination", types.YLeaf{"SourceInterfaceIsADestination", attachment.SourceInterfaceIsADestination})
    attachment.EntityData.Leafs.Append("destination-interface", types.YLeaf{"DestinationInterface", attachment.DestinationInterface})
    attachment.EntityData.Leafs.Append("traffic-direction", types.YLeaf{"TrafficDirection", attachment.TrafficDirection})

    attachment.EntityData.YListKeys = []string {"Session", "Interface"}

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

    trafficParameters.EntityData.Children = types.NewOrderedMap()
    trafficParameters.EntityData.Leafs = types.NewOrderedMap()
    trafficParameters.EntityData.Leafs.Append("traffic-direction", types.YLeaf{"TrafficDirection", trafficParameters.TrafficDirection})
    trafficParameters.EntityData.Leafs.Append("port-level", types.YLeaf{"PortLevel", trafficParameters.PortLevel})
    trafficParameters.EntityData.Leafs.Append("is-acl-enabled", types.YLeaf{"IsAclEnabled", trafficParameters.IsAclEnabled})
    trafficParameters.EntityData.Leafs.Append("mirror-bytes", types.YLeaf{"MirrorBytes", trafficParameters.MirrorBytes})
    trafficParameters.EntityData.Leafs.Append("mirror-interval", types.YLeaf{"MirrorInterval", trafficParameters.MirrorInterval})
    trafficParameters.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", trafficParameters.AclName})

    trafficParameters.EntityData.YListKeys = []string {}

    return &(trafficParameters.EntityData)
}

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId
// Destination ID
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

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

    destinationId.EntityData.Children = types.NewOrderedMap()
    destinationId.EntityData.Children.Append("ipv4-address-and-vrf", types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf})
    destinationId.EntityData.Children.Append("ipv6-address-and-vrf", types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf})
    destinationId.EntityData.Leafs = types.NewOrderedMap()
    destinationId.EntityData.Leafs.Append("destination-class", types.YLeaf{"DestinationClass", destinationId.DestinationClass})
    destinationId.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", destinationId.Interface})
    destinationId.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", destinationId.PseudowireId})
    destinationId.EntityData.Leafs.Append("invalid-value", types.YLeaf{"InvalidValue", destinationId.InvalidValue})

    destinationId.EntityData.YListKeys = []string {}

    return &(destinationId.EntityData)
}

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    ipv4AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address})
    ipv4AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName})

    ipv4AddressAndVrf.EntityData.YListKeys = []string {}

    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    ipv6AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address})
    ipv6AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName})

    ipv6AddressAndVrf.EntityData.YListKeys = []string {}

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
    HardwareSession []*SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession
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

    hardwareSessions.EntityData.Children = types.NewOrderedMap()
    hardwareSessions.EntityData.Children.Append("hardware-session", types.YChild{"HardwareSession", nil})
    for i := range hardwareSessions.HardwareSession {
        hardwareSessions.EntityData.Children.Append(types.GetSegmentPath(hardwareSessions.HardwareSession[i]), types.YChild{"HardwareSession", hardwareSessions.HardwareSession[i]})
    }
    hardwareSessions.EntityData.Leafs = types.NewOrderedMap()

    hardwareSessions.EntityData.YListKeys = []string {}

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

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Assigned numerical ID for this session. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // Configured Session Name. The type is string.
    Name interface{}

    // Session class. The type is SessionClass.
    SessionClassXr interface{}

    // Destination interface (deprecated by DestinationID, invalid for
    // pseudowires). The type is string with pattern: [a-zA-Z0-9./-]+.
    DestinationInterface interface{}

    // Last error observed for this session while programming the hardware. The
    // type is interface{} with range: 0..4294967295.
    PlatformError interface{}

    // Inject Interface ifhandle. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InjectInterfaceIfh interface{}

    // Inject Interface MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    InjectInterfaceMac interface{}

    // An inject interface is flagged as invalid on a particular node if the
    // interface exists on that node, and there is no attachment interface config
    // for it. The type is bool.
    InjectInterfaceInvalid interface{}

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

    hardwareSession.EntityData.Children = types.NewOrderedMap()
    hardwareSession.EntityData.Children.Append("destination-id", types.YChild{"DestinationId", &hardwareSession.DestinationId})
    hardwareSession.EntityData.Leafs = types.NewOrderedMap()
    hardwareSession.EntityData.Leafs.Append("session-class", types.YLeaf{"SessionClass", hardwareSession.SessionClass})
    hardwareSession.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", hardwareSession.SessionId})
    hardwareSession.EntityData.Leafs.Append("id", types.YLeaf{"Id", hardwareSession.Id})
    hardwareSession.EntityData.Leafs.Append("name", types.YLeaf{"Name", hardwareSession.Name})
    hardwareSession.EntityData.Leafs.Append("session-class-xr", types.YLeaf{"SessionClassXr", hardwareSession.SessionClassXr})
    hardwareSession.EntityData.Leafs.Append("destination-interface", types.YLeaf{"DestinationInterface", hardwareSession.DestinationInterface})
    hardwareSession.EntityData.Leafs.Append("platform-error", types.YLeaf{"PlatformError", hardwareSession.PlatformError})
    hardwareSession.EntityData.Leafs.Append("inject-interface-ifh", types.YLeaf{"InjectInterfaceIfh", hardwareSession.InjectInterfaceIfh})
    hardwareSession.EntityData.Leafs.Append("inject-interface-mac", types.YLeaf{"InjectInterfaceMac", hardwareSession.InjectInterfaceMac})
    hardwareSession.EntityData.Leafs.Append("inject-interface-invalid", types.YLeaf{"InjectInterfaceInvalid", hardwareSession.InjectInterfaceInvalid})

    hardwareSession.EntityData.YListKeys = []string {}

    return &(hardwareSession.EntityData)
}

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId
// Destination ID
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

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

    destinationId.EntityData.Children = types.NewOrderedMap()
    destinationId.EntityData.Children.Append("ipv4-address-and-vrf", types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf})
    destinationId.EntityData.Children.Append("ipv6-address-and-vrf", types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf})
    destinationId.EntityData.Leafs = types.NewOrderedMap()
    destinationId.EntityData.Leafs.Append("destination-class", types.YLeaf{"DestinationClass", destinationId.DestinationClass})
    destinationId.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", destinationId.Interface})
    destinationId.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", destinationId.PseudowireId})
    destinationId.EntityData.Leafs.Append("invalid-value", types.YLeaf{"InvalidValue", destinationId.InvalidValue})

    destinationId.EntityData.YListKeys = []string {}

    return &(destinationId.EntityData)
}

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    ipv4AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address})
    ipv4AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName})

    ipv4AddressAndVrf.EntityData.YListKeys = []string {}

    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    ipv6AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address})
    ipv6AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName})

    ipv6AddressAndVrf.EntityData.YListKeys = []string {}

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
    // The type is slice of SpanMonitorSession_Nodes_Node_Interfaces_Interface.
    Interface []*SpanMonitorSession_Nodes_Node_Interfaces_Interface
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

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface
// Information about a particular interface that
// is set up in the hardware
type SpanMonitorSession_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Interface interface{}

    // Source interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Last error observed for this interface while programming the hardware. The
    // type is interface{} with range: 0..4294967295.
    PlatformError interface{}

    // Destination interface (deprecated by Attachment). The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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
    Attachment []*SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment
}

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Interface, "interface")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("destination-id", types.YChild{"DestinationId", &self.DestinationId})
    self.EntityData.Children.Append("traffic-mirroring-parameters", types.YChild{"TrafficMirroringParameters", &self.TrafficMirroringParameters})
    self.EntityData.Children.Append("attachment", types.YChild{"Attachment", nil})
    for i := range self.Attachment {
        self.EntityData.Children.Append(types.GetSegmentPath(self.Attachment[i]), types.YChild{"Attachment", self.Attachment[i]})
    }
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", self.SourceInterface})
    self.EntityData.Leafs.Append("platform-error", types.YLeaf{"PlatformError", self.PlatformError})
    self.EntityData.Leafs.Append("destination-interface", types.YLeaf{"DestinationInterface", self.DestinationInterface})
    self.EntityData.Leafs.Append("traffic-direction", types.YLeaf{"TrafficDirection", self.TrafficDirection})

    self.EntityData.YListKeys = []string {"Interface"}

    return &(self.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId
// Destination ID (deprecated by Attachment)
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

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

    destinationId.EntityData.Children = types.NewOrderedMap()
    destinationId.EntityData.Children.Append("ipv4-address-and-vrf", types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf})
    destinationId.EntityData.Children.Append("ipv6-address-and-vrf", types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf})
    destinationId.EntityData.Leafs = types.NewOrderedMap()
    destinationId.EntityData.Leafs.Append("destination-class", types.YLeaf{"DestinationClass", destinationId.DestinationClass})
    destinationId.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", destinationId.Interface})
    destinationId.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", destinationId.PseudowireId})
    destinationId.EntityData.Leafs.Append("invalid-value", types.YLeaf{"InvalidValue", destinationId.InvalidValue})

    destinationId.EntityData.YListKeys = []string {}

    return &(destinationId.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    ipv4AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address})
    ipv4AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName})

    ipv4AddressAndVrf.EntityData.YListKeys = []string {}

    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    ipv6AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address})
    ipv6AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName})

    ipv6AddressAndVrf.EntityData.YListKeys = []string {}

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

    trafficMirroringParameters.EntityData.Children = types.NewOrderedMap()
    trafficMirroringParameters.EntityData.Leafs = types.NewOrderedMap()
    trafficMirroringParameters.EntityData.Leafs.Append("traffic-direction", types.YLeaf{"TrafficDirection", trafficMirroringParameters.TrafficDirection})
    trafficMirroringParameters.EntityData.Leafs.Append("port-level", types.YLeaf{"PortLevel", trafficMirroringParameters.PortLevel})
    trafficMirroringParameters.EntityData.Leafs.Append("is-acl-enabled", types.YLeaf{"IsAclEnabled", trafficMirroringParameters.IsAclEnabled})
    trafficMirroringParameters.EntityData.Leafs.Append("mirror-bytes", types.YLeaf{"MirrorBytes", trafficMirroringParameters.MirrorBytes})
    trafficMirroringParameters.EntityData.Leafs.Append("mirror-interval", types.YLeaf{"MirrorInterval", trafficMirroringParameters.MirrorInterval})
    trafficMirroringParameters.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", trafficMirroringParameters.AclName})

    trafficMirroringParameters.EntityData.YListKeys = []string {}

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

    attachment.EntityData.Children = types.NewOrderedMap()
    attachment.EntityData.Children.Append("destination-id", types.YChild{"DestinationId", &attachment.DestinationId})
    attachment.EntityData.Children.Append("traffic-mirroring-parameters", types.YChild{"TrafficMirroringParameters", &attachment.TrafficMirroringParameters})
    attachment.EntityData.Leafs = types.NewOrderedMap()
    attachment.EntityData.Leafs.Append("class", types.YLeaf{"Class", attachment.Class})

    attachment.EntityData.YListKeys = []string {}

    return &(attachment.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId
// Destination ID
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DestinationClass. The type is DestinationClass.
    DestinationClass interface{}

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

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

    destinationId.EntityData.Children = types.NewOrderedMap()
    destinationId.EntityData.Children.Append("ipv4-address-and-vrf", types.YChild{"Ipv4AddressAndVrf", &destinationId.Ipv4AddressAndVrf})
    destinationId.EntityData.Children.Append("ipv6-address-and-vrf", types.YChild{"Ipv6AddressAndVrf", &destinationId.Ipv6AddressAndVrf})
    destinationId.EntityData.Leafs = types.NewOrderedMap()
    destinationId.EntityData.Leafs.Append("destination-class", types.YLeaf{"DestinationClass", destinationId.DestinationClass})
    destinationId.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", destinationId.Interface})
    destinationId.EntityData.Leafs.Append("pseudowire-id", types.YLeaf{"PseudowireId", destinationId.PseudowireId})
    destinationId.EntityData.Leafs.Append("invalid-value", types.YLeaf{"InvalidValue", destinationId.InvalidValue})

    destinationId.EntityData.YListKeys = []string {}

    return &(destinationId.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    ipv4AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv4AddressAndVrf.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", ipv4AddressAndVrf.Ipv4Address})
    ipv4AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4AddressAndVrf.VrfName})

    ipv4AddressAndVrf.EntityData.YListKeys = []string {}

    return &(ipv4AddressAndVrf.EntityData)
}

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    ipv6AddressAndVrf.EntityData.Children = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs = types.NewOrderedMap()
    ipv6AddressAndVrf.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", ipv6AddressAndVrf.Ipv6Address})
    ipv6AddressAndVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6AddressAndVrf.VrfName})

    ipv6AddressAndVrf.EntityData.YListKeys = []string {}

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

    trafficMirroringParameters.EntityData.Children = types.NewOrderedMap()
    trafficMirroringParameters.EntityData.Leafs = types.NewOrderedMap()
    trafficMirroringParameters.EntityData.Leafs.Append("traffic-direction", types.YLeaf{"TrafficDirection", trafficMirroringParameters.TrafficDirection})
    trafficMirroringParameters.EntityData.Leafs.Append("port-level", types.YLeaf{"PortLevel", trafficMirroringParameters.PortLevel})
    trafficMirroringParameters.EntityData.Leafs.Append("is-acl-enabled", types.YLeaf{"IsAclEnabled", trafficMirroringParameters.IsAclEnabled})
    trafficMirroringParameters.EntityData.Leafs.Append("mirror-bytes", types.YLeaf{"MirrorBytes", trafficMirroringParameters.MirrorBytes})
    trafficMirroringParameters.EntityData.Leafs.Append("mirror-interval", types.YLeaf{"MirrorInterval", trafficMirroringParameters.MirrorInterval})
    trafficMirroringParameters.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", trafficMirroringParameters.AclName})

    trafficMirroringParameters.EntityData.YListKeys = []string {}

    return &(trafficMirroringParameters.EntityData)
}

