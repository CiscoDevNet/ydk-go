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

// SpanMonitorSession
// Monitor Session operational data
type SpanMonitorSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global operational data.
    Global SpanMonitorSession_Global

    // Node table for node-specific operational data.
    Nodes SpanMonitorSession_Nodes
}

func (spanMonitorSession *SpanMonitorSession) GetFilter() yfilter.YFilter { return spanMonitorSession.YFilter }

func (spanMonitorSession *SpanMonitorSession) SetFilter(yf yfilter.YFilter) { spanMonitorSession.YFilter = yf }

func (spanMonitorSession *SpanMonitorSession) GetGoName(yname string) string {
    if yname == "global" { return "Global" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (spanMonitorSession *SpanMonitorSession) GetSegmentPath() string {
    return "Cisco-IOS-XR-Ethernet-SPAN-oper:span-monitor-session"
}

func (spanMonitorSession *SpanMonitorSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global" {
        return &spanMonitorSession.Global
    }
    if childYangName == "nodes" {
        return &spanMonitorSession.Nodes
    }
    return nil
}

func (spanMonitorSession *SpanMonitorSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global"] = &spanMonitorSession.Global
    children["nodes"] = &spanMonitorSession.Nodes
    return children
}

func (spanMonitorSession *SpanMonitorSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spanMonitorSession *SpanMonitorSession) GetBundleName() string { return "cisco_ios_xr" }

func (spanMonitorSession *SpanMonitorSession) GetYangName() string { return "span-monitor-session" }

func (spanMonitorSession *SpanMonitorSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanMonitorSession *SpanMonitorSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanMonitorSession *SpanMonitorSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanMonitorSession *SpanMonitorSession) SetParent(parent types.Entity) { spanMonitorSession.parent = parent }

func (spanMonitorSession *SpanMonitorSession) GetParent() types.Entity { return spanMonitorSession.parent }

func (spanMonitorSession *SpanMonitorSession) GetParentYangName() string { return "Cisco-IOS-XR-Ethernet-SPAN-oper" }

// SpanMonitorSession_Global
// Global operational data
type SpanMonitorSession_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of statistics for source interfaces.
    Statistics SpanMonitorSession_Global_Statistics

    // Global Monitor Sessions table.
    GlobalSessions SpanMonitorSession_Global_GlobalSessions
}

func (global *SpanMonitorSession_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *SpanMonitorSession_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *SpanMonitorSession_Global) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    if yname == "global-sessions" { return "GlobalSessions" }
    return ""
}

func (global *SpanMonitorSession_Global) GetSegmentPath() string {
    return "global"
}

func (global *SpanMonitorSession_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &global.Statistics
    }
    if childYangName == "global-sessions" {
        return &global.GlobalSessions
    }
    return nil
}

func (global *SpanMonitorSession_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &global.Statistics
    children["global-sessions"] = &global.GlobalSessions
    return children
}

func (global *SpanMonitorSession_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *SpanMonitorSession_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *SpanMonitorSession_Global) GetYangName() string { return "global" }

func (global *SpanMonitorSession_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *SpanMonitorSession_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *SpanMonitorSession_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *SpanMonitorSession_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *SpanMonitorSession_Global) GetParent() types.Entity { return global.parent }

func (global *SpanMonitorSession_Global) GetParentYangName() string { return "span-monitor-session" }

// SpanMonitorSession_Global_Statistics
// Table of statistics for source interfaces
type SpanMonitorSession_Global_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Statistics for a particular source interface. The type is slice of
    // SpanMonitorSession_Global_Statistics_Statistic.
    Statistic []SpanMonitorSession_Global_Statistics_Statistic
}

func (statistics *SpanMonitorSession_Global_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *SpanMonitorSession_Global_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *SpanMonitorSession_Global_Statistics) GetGoName(yname string) string {
    if yname == "statistic" { return "Statistic" }
    return ""
}

func (statistics *SpanMonitorSession_Global_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *SpanMonitorSession_Global_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistic" {
        for _, c := range statistics.Statistic {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SpanMonitorSession_Global_Statistics_Statistic{}
        statistics.Statistic = append(statistics.Statistic, child)
        return &statistics.Statistic[len(statistics.Statistic)-1]
    }
    return nil
}

func (statistics *SpanMonitorSession_Global_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statistics.Statistic {
        children[statistics.Statistic[i].GetSegmentPath()] = &statistics.Statistic[i]
    }
    return children
}

func (statistics *SpanMonitorSession_Global_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *SpanMonitorSession_Global_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *SpanMonitorSession_Global_Statistics) GetYangName() string { return "statistics" }

func (statistics *SpanMonitorSession_Global_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *SpanMonitorSession_Global_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *SpanMonitorSession_Global_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *SpanMonitorSession_Global_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *SpanMonitorSession_Global_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *SpanMonitorSession_Global_Statistics) GetParentYangName() string { return "global" }

// SpanMonitorSession_Global_Statistics_Statistic
// Statistics for a particular source interface
type SpanMonitorSession_Global_Statistics_Statistic struct {
    parent types.Entity
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

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetFilter() yfilter.YFilter { return statistic.YFilter }

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) SetFilter(yf yfilter.YFilter) { statistic.YFilter = yf }

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    if yname == "interface" { return "Interface" }
    if yname == "rx-packets-mirrored" { return "RxPacketsMirrored" }
    if yname == "rx-octets-mirrored" { return "RxOctetsMirrored" }
    if yname == "tx-packets-mirrored" { return "TxPacketsMirrored" }
    if yname == "tx-octets-mirrored" { return "TxOctetsMirrored" }
    if yname == "packets-not-mirrored" { return "PacketsNotMirrored" }
    if yname == "octets-not-mirrored" { return "OctetsNotMirrored" }
    return ""
}

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetSegmentPath() string {
    return "statistic" + "[session='" + fmt.Sprintf("%v", statistic.Session) + "']" + "[interface='" + fmt.Sprintf("%v", statistic.Interface) + "']"
}

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session"] = statistic.Session
    leafs["interface"] = statistic.Interface
    leafs["rx-packets-mirrored"] = statistic.RxPacketsMirrored
    leafs["rx-octets-mirrored"] = statistic.RxOctetsMirrored
    leafs["tx-packets-mirrored"] = statistic.TxPacketsMirrored
    leafs["tx-octets-mirrored"] = statistic.TxOctetsMirrored
    leafs["packets-not-mirrored"] = statistic.PacketsNotMirrored
    leafs["octets-not-mirrored"] = statistic.OctetsNotMirrored
    return leafs
}

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetBundleName() string { return "cisco_ios_xr" }

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetYangName() string { return "statistic" }

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) SetParent(parent types.Entity) { statistic.parent = parent }

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetParent() types.Entity { return statistic.parent }

func (statistic *SpanMonitorSession_Global_Statistics_Statistic) GetParentYangName() string { return "statistics" }

// SpanMonitorSession_Global_GlobalSessions
// Global Monitor Sessions table
type SpanMonitorSession_Global_GlobalSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a globally-configured monitor session. The type is slice
    // of SpanMonitorSession_Global_GlobalSessions_GlobalSession.
    GlobalSession []SpanMonitorSession_Global_GlobalSessions_GlobalSession
}

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetFilter() yfilter.YFilter { return globalSessions.YFilter }

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) SetFilter(yf yfilter.YFilter) { globalSessions.YFilter = yf }

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetGoName(yname string) string {
    if yname == "global-session" { return "GlobalSession" }
    return ""
}

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetSegmentPath() string {
    return "global-sessions"
}

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-session" {
        for _, c := range globalSessions.GlobalSession {
            if globalSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SpanMonitorSession_Global_GlobalSessions_GlobalSession{}
        globalSessions.GlobalSession = append(globalSessions.GlobalSession, child)
        return &globalSessions.GlobalSession[len(globalSessions.GlobalSession)-1]
    }
    return nil
}

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalSessions.GlobalSession {
        children[globalSessions.GlobalSession[i].GetSegmentPath()] = &globalSessions.GlobalSession[i]
    }
    return children
}

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetBundleName() string { return "cisco_ios_xr" }

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetYangName() string { return "global-sessions" }

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) SetParent(parent types.Entity) { globalSessions.parent = parent }

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetParent() types.Entity { return globalSessions.parent }

func (globalSessions *SpanMonitorSession_Global_GlobalSessions) GetParentYangName() string { return "global" }

// SpanMonitorSession_Global_GlobalSessions_GlobalSession
// Information about a globally-configured
// monitor session
type SpanMonitorSession_Global_GlobalSessions_GlobalSession struct {
    parent types.Entity
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
}

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetFilter() yfilter.YFilter { return globalSession.YFilter }

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) SetFilter(yf yfilter.YFilter) { globalSession.YFilter = yf }

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    if yname == "name" { return "Name" }
    if yname == "session-class" { return "SessionClass" }
    if yname == "id" { return "Id" }
    if yname == "destination-error" { return "DestinationError" }
    if yname == "destination-interface-name" { return "DestinationInterfaceName" }
    if yname == "destination-interface-handle" { return "DestinationInterfaceHandle" }
    if yname == "interface-error" { return "InterfaceError" }
    if yname == "destination-data" { return "DestinationData" }
    if yname == "destination-id" { return "DestinationId" }
    return ""
}

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetSegmentPath() string {
    return "global-session" + "[session='" + fmt.Sprintf("%v", globalSession.Session) + "']"
}

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-data" {
        return &globalSession.DestinationData
    }
    if childYangName == "destination-id" {
        return &globalSession.DestinationId
    }
    return nil
}

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destination-data"] = &globalSession.DestinationData
    children["destination-id"] = &globalSession.DestinationId
    return children
}

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session"] = globalSession.Session
    leafs["name"] = globalSession.Name
    leafs["session-class"] = globalSession.SessionClass
    leafs["id"] = globalSession.Id
    leafs["destination-error"] = globalSession.DestinationError
    leafs["destination-interface-name"] = globalSession.DestinationInterfaceName
    leafs["destination-interface-handle"] = globalSession.DestinationInterfaceHandle
    leafs["interface-error"] = globalSession.InterfaceError
    return leafs
}

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetBundleName() string { return "cisco_ios_xr" }

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetYangName() string { return "global-session" }

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) SetParent(parent types.Entity) { globalSession.parent = parent }

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetParent() types.Entity { return globalSession.parent }

func (globalSession *SpanMonitorSession_Global_GlobalSessions_GlobalSession) GetParentYangName() string { return "global-sessions" }

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData
// Destination data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData struct {
    parent types.Entity
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

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetFilter() yfilter.YFilter { return destinationData.YFilter }

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) SetFilter(yf yfilter.YFilter) { destinationData.YFilter = yf }

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetGoName(yname string) string {
    if yname == "destination-class" { return "DestinationClass" }
    if yname == "invalid-value" { return "InvalidValue" }
    if yname == "interface-data" { return "InterfaceData" }
    if yname == "pseudowire-data" { return "PseudowireData" }
    if yname == "next-hop-ipv4-data" { return "NextHopIpv4Data" }
    if yname == "next-hop-ipv6-data" { return "NextHopIpv6Data" }
    return ""
}

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetSegmentPath() string {
    return "destination-data"
}

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-data" {
        return &destinationData.InterfaceData
    }
    if childYangName == "pseudowire-data" {
        return &destinationData.PseudowireData
    }
    if childYangName == "next-hop-ipv4-data" {
        return &destinationData.NextHopIpv4Data
    }
    if childYangName == "next-hop-ipv6-data" {
        return &destinationData.NextHopIpv6Data
    }
    return nil
}

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-data"] = &destinationData.InterfaceData
    children["pseudowire-data"] = &destinationData.PseudowireData
    children["next-hop-ipv4-data"] = &destinationData.NextHopIpv4Data
    children["next-hop-ipv6-data"] = &destinationData.NextHopIpv6Data
    return children
}

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-class"] = destinationData.DestinationClass
    leafs["invalid-value"] = destinationData.InvalidValue
    return leafs
}

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetBundleName() string { return "cisco_ios_xr" }

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetYangName() string { return "destination-data" }

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) SetParent(parent types.Entity) { destinationData.parent = parent }

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetParent() types.Entity { return destinationData.parent }

func (destinationData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData) GetParentYangName() string { return "global-session" }

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData
// Interface data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string.
    InterfaceName interface{}

    // Interface State. The type is ImStateEnum.
    InterfaceState interface{}
}

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetFilter() yfilter.YFilter { return interfaceData.YFilter }

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) SetFilter(yf yfilter.YFilter) { interfaceData.YFilter = yf }

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-state" { return "InterfaceState" }
    return ""
}

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetSegmentPath() string {
    return "interface-data"
}

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = interfaceData.InterfaceName
    leafs["interface-state"] = interfaceData.InterfaceState
    return leafs
}

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetYangName() string { return "interface-data" }

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) SetParent(parent types.Entity) { interfaceData.parent = parent }

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetParent() types.Entity { return interfaceData.parent }

func (interfaceData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_InterfaceData) GetParentYangName() string { return "destination-data" }

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData
// Pseudowire data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pseudowire Name. The type is string.
    PseudowireName interface{}

    // Pseudowire State. The type is bool.
    PseudowireIsUp interface{}
}

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetFilter() yfilter.YFilter { return pseudowireData.YFilter }

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) SetFilter(yf yfilter.YFilter) { pseudowireData.YFilter = yf }

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetGoName(yname string) string {
    if yname == "pseudowire-name" { return "PseudowireName" }
    if yname == "pseudowire-is-up" { return "PseudowireIsUp" }
    return ""
}

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetSegmentPath() string {
    return "pseudowire-data"
}

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pseudowire-name"] = pseudowireData.PseudowireName
    leafs["pseudowire-is-up"] = pseudowireData.PseudowireIsUp
    return leafs
}

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetBundleName() string { return "cisco_ios_xr" }

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetYangName() string { return "pseudowire-data" }

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) SetParent(parent types.Entity) { pseudowireData.parent = parent }

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetParent() types.Entity { return pseudowireData.parent }

func (pseudowireData *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_PseudowireData) GetParentYangName() string { return "destination-data" }

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data
// Next-hop IPv4 data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Address is reachable. The type is bool.
    AddressIsReachable interface{}
}

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetFilter() yfilter.YFilter { return nextHopIpv4Data.YFilter }

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) SetFilter(yf yfilter.YFilter) { nextHopIpv4Data.YFilter = yf }

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "address-is-reachable" { return "AddressIsReachable" }
    return ""
}

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetSegmentPath() string {
    return "next-hop-ipv4-data"
}

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = nextHopIpv4Data.Ipv4Address
    leafs["vrf-name"] = nextHopIpv4Data.VrfName
    leafs["address-is-reachable"] = nextHopIpv4Data.AddressIsReachable
    return leafs
}

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetBundleName() string { return "cisco_ios_xr" }

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetYangName() string { return "next-hop-ipv4-data" }

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) SetParent(parent types.Entity) { nextHopIpv4Data.parent = parent }

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetParent() types.Entity { return nextHopIpv4Data.parent }

func (nextHopIpv4Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv4Data) GetParentYangName() string { return "destination-data" }

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data
// Next-hop IPv6 data
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Address is reachable. The type is bool.
    AddressIsReachable interface{}
}

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetFilter() yfilter.YFilter { return nextHopIpv6Data.YFilter }

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) SetFilter(yf yfilter.YFilter) { nextHopIpv6Data.YFilter = yf }

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "address-is-reachable" { return "AddressIsReachable" }
    return ""
}

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetSegmentPath() string {
    return "next-hop-ipv6-data"
}

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = nextHopIpv6Data.Ipv6Address
    leafs["vrf-name"] = nextHopIpv6Data.VrfName
    leafs["address-is-reachable"] = nextHopIpv6Data.AddressIsReachable
    return leafs
}

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetBundleName() string { return "cisco_ios_xr" }

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetYangName() string { return "next-hop-ipv6-data" }

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) SetParent(parent types.Entity) { nextHopIpv6Data.parent = parent }

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetParent() types.Entity { return nextHopIpv6Data.parent }

func (nextHopIpv6Data *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationData_NextHopIpv6Data) GetParentYangName() string { return "destination-data" }

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId
// Destination ID
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId struct {
    parent types.Entity
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

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetFilter() yfilter.YFilter { return destinationId.YFilter }

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) SetFilter(yf yfilter.YFilter) { destinationId.YFilter = yf }

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetGoName(yname string) string {
    if yname == "destination-class" { return "DestinationClass" }
    if yname == "interface" { return "Interface" }
    if yname == "pseudowire-id" { return "PseudowireId" }
    if yname == "invalid-value" { return "InvalidValue" }
    if yname == "ipv4-address-and-vrf" { return "Ipv4AddressAndVrf" }
    if yname == "ipv6-address-and-vrf" { return "Ipv6AddressAndVrf" }
    return ""
}

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetSegmentPath() string {
    return "destination-id"
}

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-address-and-vrf" {
        return &destinationId.Ipv4AddressAndVrf
    }
    if childYangName == "ipv6-address-and-vrf" {
        return &destinationId.Ipv6AddressAndVrf
    }
    return nil
}

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-address-and-vrf"] = &destinationId.Ipv4AddressAndVrf
    children["ipv6-address-and-vrf"] = &destinationId.Ipv6AddressAndVrf
    return children
}

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-class"] = destinationId.DestinationClass
    leafs["interface"] = destinationId.Interface
    leafs["pseudowire-id"] = destinationId.PseudowireId
    leafs["invalid-value"] = destinationId.InvalidValue
    return leafs
}

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetBundleName() string { return "cisco_ios_xr" }

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetYangName() string { return "destination-id" }

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) SetParent(parent types.Entity) { destinationId.parent = parent }

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetParent() types.Entity { return destinationId.parent }

func (destinationId *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId) GetParentYangName() string { return "global-session" }

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetFilter() yfilter.YFilter { return ipv4AddressAndVrf.YFilter }

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv4AddressAndVrf.YFilter = yf }

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetSegmentPath() string {
    return "ipv4-address-and-vrf"
}

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = ipv4AddressAndVrf.Ipv4Address
    leafs["vrf-name"] = ipv4AddressAndVrf.VrfName
    return leafs
}

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetYangName() string { return "ipv4-address-and-vrf" }

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) SetParent(parent types.Entity) { ipv4AddressAndVrf.parent = parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetParent() types.Entity { return ipv4AddressAndVrf.parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv4AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetFilter() yfilter.YFilter { return ipv6AddressAndVrf.YFilter }

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv6AddressAndVrf.YFilter = yf }

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetSegmentPath() string {
    return "ipv6-address-and-vrf"
}

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6AddressAndVrf.Ipv6Address
    leafs["vrf-name"] = ipv6AddressAndVrf.VrfName
    return leafs
}

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetYangName() string { return "ipv6-address-and-vrf" }

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) SetParent(parent types.Entity) { ipv6AddressAndVrf.parent = parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetParent() types.Entity { return ipv6AddressAndVrf.parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Global_GlobalSessions_GlobalSession_DestinationId_Ipv6AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Nodes
// Node table for node-specific operational data
type SpanMonitorSession_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific data for a particular node. The type is slice of
    // SpanMonitorSession_Nodes_Node.
    Node []SpanMonitorSession_Nodes_Node
}

func (nodes *SpanMonitorSession_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *SpanMonitorSession_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *SpanMonitorSession_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *SpanMonitorSession_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *SpanMonitorSession_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SpanMonitorSession_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *SpanMonitorSession_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *SpanMonitorSession_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *SpanMonitorSession_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *SpanMonitorSession_Nodes) GetYangName() string { return "nodes" }

func (nodes *SpanMonitorSession_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *SpanMonitorSession_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *SpanMonitorSession_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *SpanMonitorSession_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *SpanMonitorSession_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *SpanMonitorSession_Nodes) GetParentYangName() string { return "span-monitor-session" }

// SpanMonitorSession_Nodes_Node
// Node-specific data for a particular node
type SpanMonitorSession_Nodes_Node struct {
    parent types.Entity
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

func (node *SpanMonitorSession_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *SpanMonitorSession_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *SpanMonitorSession_Nodes_Node) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "attachments" { return "Attachments" }
    if yname == "hardware-sessions" { return "HardwareSessions" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (node *SpanMonitorSession_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
}

func (node *SpanMonitorSession_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attachments" {
        return &node.Attachments
    }
    if childYangName == "hardware-sessions" {
        return &node.HardwareSessions
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    return nil
}

func (node *SpanMonitorSession_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attachments"] = &node.Attachments
    children["hardware-sessions"] = &node.HardwareSessions
    children["interfaces"] = &node.Interfaces
    return children
}

func (node *SpanMonitorSession_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = node.Node
    return leafs
}

func (node *SpanMonitorSession_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *SpanMonitorSession_Nodes_Node) GetYangName() string { return "node" }

func (node *SpanMonitorSession_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *SpanMonitorSession_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *SpanMonitorSession_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *SpanMonitorSession_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *SpanMonitorSession_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *SpanMonitorSession_Nodes_Node) GetParentYangName() string { return "nodes" }

// SpanMonitorSession_Nodes_Node_Attachments
// Table of source interfaces configured as
// attached to a session
type SpanMonitorSession_Nodes_Node_Attachments struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular source interface configured as attached to
    // monitor session. The type is slice of
    // SpanMonitorSession_Nodes_Node_Attachments_Attachment.
    Attachment []SpanMonitorSession_Nodes_Node_Attachments_Attachment
}

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetFilter() yfilter.YFilter { return attachments.YFilter }

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) SetFilter(yf yfilter.YFilter) { attachments.YFilter = yf }

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetGoName(yname string) string {
    if yname == "attachment" { return "Attachment" }
    return ""
}

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetSegmentPath() string {
    return "attachments"
}

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attachment" {
        for _, c := range attachments.Attachment {
            if attachments.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SpanMonitorSession_Nodes_Node_Attachments_Attachment{}
        attachments.Attachment = append(attachments.Attachment, child)
        return &attachments.Attachment[len(attachments.Attachment)-1]
    }
    return nil
}

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attachments.Attachment {
        children[attachments.Attachment[i].GetSegmentPath()] = &attachments.Attachment[i]
    }
    return children
}

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetBundleName() string { return "cisco_ios_xr" }

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetYangName() string { return "attachments" }

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) SetParent(parent types.Entity) { attachments.parent = parent }

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetParent() types.Entity { return attachments.parent }

func (attachments *SpanMonitorSession_Nodes_Node_Attachments) GetParentYangName() string { return "node" }

// SpanMonitorSession_Nodes_Node_Attachments_Attachment
// Information about a particular source
// interface configured as attached to monitor
// session
type SpanMonitorSession_Nodes_Node_Attachments_Attachment struct {
    parent types.Entity
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

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetFilter() yfilter.YFilter { return attachment.YFilter }

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) SetFilter(yf yfilter.YFilter) { attachment.YFilter = yf }

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    if yname == "interface" { return "Interface" }
    if yname == "name" { return "Name" }
    if yname == "local-class" { return "LocalClass" }
    if yname == "id" { return "Id" }
    if yname == "global-class" { return "GlobalClass" }
    if yname == "session-is-configured" { return "SessionIsConfigured" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "source-interface-state" { return "SourceInterfaceState" }
    if yname == "pfi-error" { return "PfiError" }
    if yname == "dest-pw-type-not-supported" { return "DestPwTypeNotSupported" }
    if yname == "source-interface-is-a-destination" { return "SourceInterfaceIsADestination" }
    if yname == "destination-interface" { return "DestinationInterface" }
    if yname == "traffic-direction" { return "TrafficDirection" }
    if yname == "traffic-parameters" { return "TrafficParameters" }
    if yname == "destination-id" { return "DestinationId" }
    return ""
}

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetSegmentPath() string {
    return "attachment" + "[session='" + fmt.Sprintf("%v", attachment.Session) + "']" + "[interface='" + fmt.Sprintf("%v", attachment.Interface) + "']"
}

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic-parameters" {
        return &attachment.TrafficParameters
    }
    if childYangName == "destination-id" {
        return &attachment.DestinationId
    }
    return nil
}

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traffic-parameters"] = &attachment.TrafficParameters
    children["destination-id"] = &attachment.DestinationId
    return children
}

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session"] = attachment.Session
    leafs["interface"] = attachment.Interface
    leafs["name"] = attachment.Name
    leafs["local-class"] = attachment.LocalClass
    leafs["id"] = attachment.Id
    leafs["global-class"] = attachment.GlobalClass
    leafs["session-is-configured"] = attachment.SessionIsConfigured
    leafs["source-interface"] = attachment.SourceInterface
    leafs["source-interface-state"] = attachment.SourceInterfaceState
    leafs["pfi-error"] = attachment.PfiError
    leafs["dest-pw-type-not-supported"] = attachment.DestPwTypeNotSupported
    leafs["source-interface-is-a-destination"] = attachment.SourceInterfaceIsADestination
    leafs["destination-interface"] = attachment.DestinationInterface
    leafs["traffic-direction"] = attachment.TrafficDirection
    return leafs
}

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetBundleName() string { return "cisco_ios_xr" }

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetYangName() string { return "attachment" }

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) SetParent(parent types.Entity) { attachment.parent = parent }

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetParent() types.Entity { return attachment.parent }

func (attachment *SpanMonitorSession_Nodes_Node_Attachments_Attachment) GetParentYangName() string { return "attachments" }

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters
// Traffic mirroring parameters
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters struct {
    parent types.Entity
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

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetFilter() yfilter.YFilter { return trafficParameters.YFilter }

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) SetFilter(yf yfilter.YFilter) { trafficParameters.YFilter = yf }

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetGoName(yname string) string {
    if yname == "traffic-direction" { return "TrafficDirection" }
    if yname == "port-level" { return "PortLevel" }
    if yname == "is-acl-enabled" { return "IsAclEnabled" }
    if yname == "mirror-bytes" { return "MirrorBytes" }
    if yname == "mirror-interval" { return "MirrorInterval" }
    if yname == "acl-name" { return "AclName" }
    return ""
}

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetSegmentPath() string {
    return "traffic-parameters"
}

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["traffic-direction"] = trafficParameters.TrafficDirection
    leafs["port-level"] = trafficParameters.PortLevel
    leafs["is-acl-enabled"] = trafficParameters.IsAclEnabled
    leafs["mirror-bytes"] = trafficParameters.MirrorBytes
    leafs["mirror-interval"] = trafficParameters.MirrorInterval
    leafs["acl-name"] = trafficParameters.AclName
    return leafs
}

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetBundleName() string { return "cisco_ios_xr" }

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetYangName() string { return "traffic-parameters" }

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) SetParent(parent types.Entity) { trafficParameters.parent = parent }

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetParent() types.Entity { return trafficParameters.parent }

func (trafficParameters *SpanMonitorSession_Nodes_Node_Attachments_Attachment_TrafficParameters) GetParentYangName() string { return "attachment" }

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId
// Destination ID
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId struct {
    parent types.Entity
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

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetFilter() yfilter.YFilter { return destinationId.YFilter }

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) SetFilter(yf yfilter.YFilter) { destinationId.YFilter = yf }

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetGoName(yname string) string {
    if yname == "destination-class" { return "DestinationClass" }
    if yname == "interface" { return "Interface" }
    if yname == "pseudowire-id" { return "PseudowireId" }
    if yname == "invalid-value" { return "InvalidValue" }
    if yname == "ipv4-address-and-vrf" { return "Ipv4AddressAndVrf" }
    if yname == "ipv6-address-and-vrf" { return "Ipv6AddressAndVrf" }
    return ""
}

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetSegmentPath() string {
    return "destination-id"
}

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-address-and-vrf" {
        return &destinationId.Ipv4AddressAndVrf
    }
    if childYangName == "ipv6-address-and-vrf" {
        return &destinationId.Ipv6AddressAndVrf
    }
    return nil
}

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-address-and-vrf"] = &destinationId.Ipv4AddressAndVrf
    children["ipv6-address-and-vrf"] = &destinationId.Ipv6AddressAndVrf
    return children
}

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-class"] = destinationId.DestinationClass
    leafs["interface"] = destinationId.Interface
    leafs["pseudowire-id"] = destinationId.PseudowireId
    leafs["invalid-value"] = destinationId.InvalidValue
    return leafs
}

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetBundleName() string { return "cisco_ios_xr" }

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetYangName() string { return "destination-id" }

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) SetParent(parent types.Entity) { destinationId.parent = parent }

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetParent() types.Entity { return destinationId.parent }

func (destinationId *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId) GetParentYangName() string { return "attachment" }

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetFilter() yfilter.YFilter { return ipv4AddressAndVrf.YFilter }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv4AddressAndVrf.YFilter = yf }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetSegmentPath() string {
    return "ipv4-address-and-vrf"
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = ipv4AddressAndVrf.Ipv4Address
    leafs["vrf-name"] = ipv4AddressAndVrf.VrfName
    return leafs
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetYangName() string { return "ipv4-address-and-vrf" }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) SetParent(parent types.Entity) { ipv4AddressAndVrf.parent = parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetParent() types.Entity { return ipv4AddressAndVrf.parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv4AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetFilter() yfilter.YFilter { return ipv6AddressAndVrf.YFilter }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv6AddressAndVrf.YFilter = yf }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetSegmentPath() string {
    return "ipv6-address-and-vrf"
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6AddressAndVrf.Ipv6Address
    leafs["vrf-name"] = ipv6AddressAndVrf.VrfName
    return leafs
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetYangName() string { return "ipv6-address-and-vrf" }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) SetParent(parent types.Entity) { ipv6AddressAndVrf.parent = parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetParent() types.Entity { return ipv6AddressAndVrf.parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Attachments_Attachment_DestinationId_Ipv6AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Nodes_Node_HardwareSessions
// Table of sessions set up in the hardware. 
// When all sessions are operating correctly the
// entries in this table should match those
// entries in GlobalSessionTable that have a
// destination configured
type SpanMonitorSession_Nodes_Node_HardwareSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular session that is set up in the hardware. The
    // type is slice of
    // SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession.
    HardwareSession []SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession
}

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetFilter() yfilter.YFilter { return hardwareSessions.YFilter }

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) SetFilter(yf yfilter.YFilter) { hardwareSessions.YFilter = yf }

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetGoName(yname string) string {
    if yname == "hardware-session" { return "HardwareSession" }
    return ""
}

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetSegmentPath() string {
    return "hardware-sessions"
}

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hardware-session" {
        for _, c := range hardwareSessions.HardwareSession {
            if hardwareSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession{}
        hardwareSessions.HardwareSession = append(hardwareSessions.HardwareSession, child)
        return &hardwareSessions.HardwareSession[len(hardwareSessions.HardwareSession)-1]
    }
    return nil
}

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hardwareSessions.HardwareSession {
        children[hardwareSessions.HardwareSession[i].GetSegmentPath()] = &hardwareSessions.HardwareSession[i]
    }
    return children
}

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetYangName() string { return "hardware-sessions" }

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) SetParent(parent types.Entity) { hardwareSessions.parent = parent }

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetParent() types.Entity { return hardwareSessions.parent }

func (hardwareSessions *SpanMonitorSession_Nodes_Node_HardwareSessions) GetParentYangName() string { return "node" }

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession
// Information about a particular session that
// is set up in the hardware
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession struct {
    parent types.Entity
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
    // pseudowires). The type is string with pattern: [a-zA-Z0-9./-]+.
    DestinationInterface interface{}

    // Last error observed for this session while programming the hardware. The
    // type is interface{} with range: 0..4294967295.
    PlatformError interface{}

    // Destination ID.
    DestinationId SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId
}

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetFilter() yfilter.YFilter { return hardwareSession.YFilter }

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) SetFilter(yf yfilter.YFilter) { hardwareSession.YFilter = yf }

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetGoName(yname string) string {
    if yname == "session-class" { return "SessionClass" }
    if yname == "session-id" { return "SessionId" }
    if yname == "id" { return "Id" }
    if yname == "name" { return "Name" }
    if yname == "session-class-xr" { return "SessionClassXr" }
    if yname == "destination-interface" { return "DestinationInterface" }
    if yname == "platform-error" { return "PlatformError" }
    if yname == "destination-id" { return "DestinationId" }
    return ""
}

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetSegmentPath() string {
    return "hardware-session"
}

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-id" {
        return &hardwareSession.DestinationId
    }
    return nil
}

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destination-id"] = &hardwareSession.DestinationId
    return children
}

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-class"] = hardwareSession.SessionClass
    leafs["session-id"] = hardwareSession.SessionId
    leafs["id"] = hardwareSession.Id
    leafs["name"] = hardwareSession.Name
    leafs["session-class-xr"] = hardwareSession.SessionClassXr
    leafs["destination-interface"] = hardwareSession.DestinationInterface
    leafs["platform-error"] = hardwareSession.PlatformError
    return leafs
}

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetYangName() string { return "hardware-session" }

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) SetParent(parent types.Entity) { hardwareSession.parent = parent }

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetParent() types.Entity { return hardwareSession.parent }

func (hardwareSession *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession) GetParentYangName() string { return "hardware-sessions" }

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId
// Destination ID
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId struct {
    parent types.Entity
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

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetFilter() yfilter.YFilter { return destinationId.YFilter }

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) SetFilter(yf yfilter.YFilter) { destinationId.YFilter = yf }

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetGoName(yname string) string {
    if yname == "destination-class" { return "DestinationClass" }
    if yname == "interface" { return "Interface" }
    if yname == "pseudowire-id" { return "PseudowireId" }
    if yname == "invalid-value" { return "InvalidValue" }
    if yname == "ipv4-address-and-vrf" { return "Ipv4AddressAndVrf" }
    if yname == "ipv6-address-and-vrf" { return "Ipv6AddressAndVrf" }
    return ""
}

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetSegmentPath() string {
    return "destination-id"
}

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-address-and-vrf" {
        return &destinationId.Ipv4AddressAndVrf
    }
    if childYangName == "ipv6-address-and-vrf" {
        return &destinationId.Ipv6AddressAndVrf
    }
    return nil
}

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-address-and-vrf"] = &destinationId.Ipv4AddressAndVrf
    children["ipv6-address-and-vrf"] = &destinationId.Ipv6AddressAndVrf
    return children
}

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-class"] = destinationId.DestinationClass
    leafs["interface"] = destinationId.Interface
    leafs["pseudowire-id"] = destinationId.PseudowireId
    leafs["invalid-value"] = destinationId.InvalidValue
    return leafs
}

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetBundleName() string { return "cisco_ios_xr" }

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetYangName() string { return "destination-id" }

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) SetParent(parent types.Entity) { destinationId.parent = parent }

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetParent() types.Entity { return destinationId.parent }

func (destinationId *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId) GetParentYangName() string { return "hardware-session" }

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetFilter() yfilter.YFilter { return ipv4AddressAndVrf.YFilter }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv4AddressAndVrf.YFilter = yf }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetSegmentPath() string {
    return "ipv4-address-and-vrf"
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = ipv4AddressAndVrf.Ipv4Address
    leafs["vrf-name"] = ipv4AddressAndVrf.VrfName
    return leafs
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetYangName() string { return "ipv4-address-and-vrf" }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) SetParent(parent types.Entity) { ipv4AddressAndVrf.parent = parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetParent() types.Entity { return ipv4AddressAndVrf.parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv4AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetFilter() yfilter.YFilter { return ipv6AddressAndVrf.YFilter }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv6AddressAndVrf.YFilter = yf }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetSegmentPath() string {
    return "ipv6-address-and-vrf"
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6AddressAndVrf.Ipv6Address
    leafs["vrf-name"] = ipv6AddressAndVrf.VrfName
    return leafs
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetYangName() string { return "ipv6-address-and-vrf" }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) SetParent(parent types.Entity) { ipv6AddressAndVrf.parent = parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetParent() types.Entity { return ipv6AddressAndVrf.parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_HardwareSessions_HardwareSession_DestinationId_Ipv6AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Nodes_Node_Interfaces
// Table of source interfaces set up in the
// hardware.  The entries in this table should
// match the entries in AttachmentTable when all
// sessions are operating correctly
type SpanMonitorSession_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular interface that is set up in the hardware.
    // The type is slice of SpanMonitorSession_Nodes_Node_Interfaces_Interface.
    Interface []SpanMonitorSession_Nodes_Node_Interfaces_Interface
}

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SpanMonitorSession_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *SpanMonitorSession_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface
// Information about a particular interface that
// is set up in the hardware
type SpanMonitorSession_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
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
    Attachment []SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment
}

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "platform-error" { return "PlatformError" }
    if yname == "destination-interface" { return "DestinationInterface" }
    if yname == "traffic-direction" { return "TrafficDirection" }
    if yname == "destination-id" { return "DestinationId" }
    if yname == "traffic-mirroring-parameters" { return "TrafficMirroringParameters" }
    if yname == "attachment" { return "Attachment" }
    return ""
}

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-id" {
        return &self.DestinationId
    }
    if childYangName == "traffic-mirroring-parameters" {
        return &self.TrafficMirroringParameters
    }
    if childYangName == "attachment" {
        for _, c := range self.Attachment {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment{}
        self.Attachment = append(self.Attachment, child)
        return &self.Attachment[len(self.Attachment)-1]
    }
    return nil
}

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destination-id"] = &self.DestinationId
    children["traffic-mirroring-parameters"] = &self.TrafficMirroringParameters
    for i := range self.Attachment {
        children[self.Attachment[i].GetSegmentPath()] = &self.Attachment[i]
    }
    return children
}

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = self.Interface
    leafs["source-interface"] = self.SourceInterface
    leafs["platform-error"] = self.PlatformError
    leafs["destination-interface"] = self.DestinationInterface
    leafs["traffic-direction"] = self.TrafficDirection
    return leafs
}

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *SpanMonitorSession_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId
// Destination ID (deprecated by Attachment)
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId struct {
    parent types.Entity
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

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetFilter() yfilter.YFilter { return destinationId.YFilter }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) SetFilter(yf yfilter.YFilter) { destinationId.YFilter = yf }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetGoName(yname string) string {
    if yname == "destination-class" { return "DestinationClass" }
    if yname == "interface" { return "Interface" }
    if yname == "pseudowire-id" { return "PseudowireId" }
    if yname == "invalid-value" { return "InvalidValue" }
    if yname == "ipv4-address-and-vrf" { return "Ipv4AddressAndVrf" }
    if yname == "ipv6-address-and-vrf" { return "Ipv6AddressAndVrf" }
    return ""
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetSegmentPath() string {
    return "destination-id"
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-address-and-vrf" {
        return &destinationId.Ipv4AddressAndVrf
    }
    if childYangName == "ipv6-address-and-vrf" {
        return &destinationId.Ipv6AddressAndVrf
    }
    return nil
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-address-and-vrf"] = &destinationId.Ipv4AddressAndVrf
    children["ipv6-address-and-vrf"] = &destinationId.Ipv6AddressAndVrf
    return children
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-class"] = destinationId.DestinationClass
    leafs["interface"] = destinationId.Interface
    leafs["pseudowire-id"] = destinationId.PseudowireId
    leafs["invalid-value"] = destinationId.InvalidValue
    return leafs
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetBundleName() string { return "cisco_ios_xr" }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetYangName() string { return "destination-id" }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) SetParent(parent types.Entity) { destinationId.parent = parent }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetParent() types.Entity { return destinationId.parent }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId) GetParentYangName() string { return "interface" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetFilter() yfilter.YFilter { return ipv4AddressAndVrf.YFilter }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv4AddressAndVrf.YFilter = yf }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetSegmentPath() string {
    return "ipv4-address-and-vrf"
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = ipv4AddressAndVrf.Ipv4Address
    leafs["vrf-name"] = ipv4AddressAndVrf.VrfName
    return leafs
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetYangName() string { return "ipv4-address-and-vrf" }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) SetParent(parent types.Entity) { ipv4AddressAndVrf.parent = parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetParent() types.Entity { return ipv4AddressAndVrf.parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv4AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetFilter() yfilter.YFilter { return ipv6AddressAndVrf.YFilter }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv6AddressAndVrf.YFilter = yf }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetSegmentPath() string {
    return "ipv6-address-and-vrf"
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6AddressAndVrf.Ipv6Address
    leafs["vrf-name"] = ipv6AddressAndVrf.VrfName
    return leafs
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetYangName() string { return "ipv6-address-and-vrf" }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) SetParent(parent types.Entity) { ipv6AddressAndVrf.parent = parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetParent() types.Entity { return ipv6AddressAndVrf.parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_DestinationId_Ipv6AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters
// Traffic mirroring parameters (deprecated by
// Attachment)
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters struct {
    parent types.Entity
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

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetFilter() yfilter.YFilter { return trafficMirroringParameters.YFilter }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) SetFilter(yf yfilter.YFilter) { trafficMirroringParameters.YFilter = yf }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetGoName(yname string) string {
    if yname == "traffic-direction" { return "TrafficDirection" }
    if yname == "port-level" { return "PortLevel" }
    if yname == "is-acl-enabled" { return "IsAclEnabled" }
    if yname == "mirror-bytes" { return "MirrorBytes" }
    if yname == "mirror-interval" { return "MirrorInterval" }
    if yname == "acl-name" { return "AclName" }
    return ""
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetSegmentPath() string {
    return "traffic-mirroring-parameters"
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["traffic-direction"] = trafficMirroringParameters.TrafficDirection
    leafs["port-level"] = trafficMirroringParameters.PortLevel
    leafs["is-acl-enabled"] = trafficMirroringParameters.IsAclEnabled
    leafs["mirror-bytes"] = trafficMirroringParameters.MirrorBytes
    leafs["mirror-interval"] = trafficMirroringParameters.MirrorInterval
    leafs["acl-name"] = trafficMirroringParameters.AclName
    return leafs
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetBundleName() string { return "cisco_ios_xr" }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetYangName() string { return "traffic-mirroring-parameters" }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) SetParent(parent types.Entity) { trafficMirroringParameters.parent = parent }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetParent() types.Entity { return trafficMirroringParameters.parent }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_TrafficMirroringParameters) GetParentYangName() string { return "interface" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment
// Attachment information
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Attachment class. The type is SessionClass.
    Class interface{}

    // Destination ID.
    DestinationId SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId

    // Traffic mirroring parameters.
    TrafficMirroringParameters SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters
}

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetFilter() yfilter.YFilter { return attachment.YFilter }

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) SetFilter(yf yfilter.YFilter) { attachment.YFilter = yf }

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    if yname == "destination-id" { return "DestinationId" }
    if yname == "traffic-mirroring-parameters" { return "TrafficMirroringParameters" }
    return ""
}

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetSegmentPath() string {
    return "attachment"
}

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-id" {
        return &attachment.DestinationId
    }
    if childYangName == "traffic-mirroring-parameters" {
        return &attachment.TrafficMirroringParameters
    }
    return nil
}

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destination-id"] = &attachment.DestinationId
    children["traffic-mirroring-parameters"] = &attachment.TrafficMirroringParameters
    return children
}

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class"] = attachment.Class
    return leafs
}

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetBundleName() string { return "cisco_ios_xr" }

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetYangName() string { return "attachment" }

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) SetParent(parent types.Entity) { attachment.parent = parent }

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetParent() types.Entity { return attachment.parent }

func (attachment *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment) GetParentYangName() string { return "interface" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId
// Destination ID
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId struct {
    parent types.Entity
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

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetFilter() yfilter.YFilter { return destinationId.YFilter }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) SetFilter(yf yfilter.YFilter) { destinationId.YFilter = yf }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetGoName(yname string) string {
    if yname == "destination-class" { return "DestinationClass" }
    if yname == "interface" { return "Interface" }
    if yname == "pseudowire-id" { return "PseudowireId" }
    if yname == "invalid-value" { return "InvalidValue" }
    if yname == "ipv4-address-and-vrf" { return "Ipv4AddressAndVrf" }
    if yname == "ipv6-address-and-vrf" { return "Ipv6AddressAndVrf" }
    return ""
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetSegmentPath() string {
    return "destination-id"
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-address-and-vrf" {
        return &destinationId.Ipv4AddressAndVrf
    }
    if childYangName == "ipv6-address-and-vrf" {
        return &destinationId.Ipv6AddressAndVrf
    }
    return nil
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-address-and-vrf"] = &destinationId.Ipv4AddressAndVrf
    children["ipv6-address-and-vrf"] = &destinationId.Ipv6AddressAndVrf
    return children
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-class"] = destinationId.DestinationClass
    leafs["interface"] = destinationId.Interface
    leafs["pseudowire-id"] = destinationId.PseudowireId
    leafs["invalid-value"] = destinationId.InvalidValue
    return leafs
}

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetBundleName() string { return "cisco_ios_xr" }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetYangName() string { return "destination-id" }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) SetParent(parent types.Entity) { destinationId.parent = parent }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetParent() types.Entity { return destinationId.parent }

func (destinationId *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId) GetParentYangName() string { return "attachment" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf
// IPv4 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetFilter() yfilter.YFilter { return ipv4AddressAndVrf.YFilter }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv4AddressAndVrf.YFilter = yf }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetSegmentPath() string {
    return "ipv4-address-and-vrf"
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-address"] = ipv4AddressAndVrf.Ipv4Address
    leafs["vrf-name"] = ipv4AddressAndVrf.VrfName
    return leafs
}

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetYangName() string { return "ipv4-address-and-vrf" }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) SetParent(parent types.Entity) { ipv4AddressAndVrf.parent = parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetParent() types.Entity { return ipv4AddressAndVrf.parent }

func (ipv4AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv4AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf
// IPv6 address
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // VRF. The type is string.
    VrfName interface{}
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetFilter() yfilter.YFilter { return ipv6AddressAndVrf.YFilter }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) SetFilter(yf yfilter.YFilter) { ipv6AddressAndVrf.YFilter = yf }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetSegmentPath() string {
    return "ipv6-address-and-vrf"
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-address"] = ipv6AddressAndVrf.Ipv6Address
    leafs["vrf-name"] = ipv6AddressAndVrf.VrfName
    return leafs
}

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetYangName() string { return "ipv6-address-and-vrf" }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) SetParent(parent types.Entity) { ipv6AddressAndVrf.parent = parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetParent() types.Entity { return ipv6AddressAndVrf.parent }

func (ipv6AddressAndVrf *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_DestinationId_Ipv6AddressAndVrf) GetParentYangName() string { return "destination-id" }

// SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters
// Traffic mirroring parameters
type SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters struct {
    parent types.Entity
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

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetFilter() yfilter.YFilter { return trafficMirroringParameters.YFilter }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) SetFilter(yf yfilter.YFilter) { trafficMirroringParameters.YFilter = yf }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetGoName(yname string) string {
    if yname == "traffic-direction" { return "TrafficDirection" }
    if yname == "port-level" { return "PortLevel" }
    if yname == "is-acl-enabled" { return "IsAclEnabled" }
    if yname == "mirror-bytes" { return "MirrorBytes" }
    if yname == "mirror-interval" { return "MirrorInterval" }
    if yname == "acl-name" { return "AclName" }
    return ""
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetSegmentPath() string {
    return "traffic-mirroring-parameters"
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["traffic-direction"] = trafficMirroringParameters.TrafficDirection
    leafs["port-level"] = trafficMirroringParameters.PortLevel
    leafs["is-acl-enabled"] = trafficMirroringParameters.IsAclEnabled
    leafs["mirror-bytes"] = trafficMirroringParameters.MirrorBytes
    leafs["mirror-interval"] = trafficMirroringParameters.MirrorInterval
    leafs["acl-name"] = trafficMirroringParameters.AclName
    return leafs
}

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetBundleName() string { return "cisco_ios_xr" }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetYangName() string { return "traffic-mirroring-parameters" }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) SetParent(parent types.Entity) { trafficMirroringParameters.parent = parent }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetParent() types.Entity { return trafficMirroringParameters.parent }

func (trafficMirroringParameters *SpanMonitorSession_Nodes_Node_Interfaces_Interface_Attachment_TrafficMirroringParameters) GetParentYangName() string { return "attachment" }

