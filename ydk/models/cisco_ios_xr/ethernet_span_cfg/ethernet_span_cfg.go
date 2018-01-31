// This module contains a collection of YANG definitions
// for Cisco IOS-XR Ethernet-SPAN package configuration.
// 
// This module contains definitions
// for the following management objects:
//   span-monitor-session: none
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-l2vpn-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_span_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_span_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-Ethernet-SPAN-cfg span-monitor-session}", reflect.TypeOf(SpanMonitorSession{}))
    ydk.RegisterEntity("Cisco-IOS-XR-Ethernet-SPAN-cfg:span-monitor-session", reflect.TypeOf(SpanMonitorSession{}))
}

// SpanTrafficDirection represents Span traffic direction
type SpanTrafficDirection string

const (
    // Replicate only received (ingress) traffic
    SpanTrafficDirection_rx_only SpanTrafficDirection = "rx-only"

    // Replicate only transmitted (egress) traffic
    SpanTrafficDirection_tx_only SpanTrafficDirection = "tx-only"
)

// SpanMirrorInterval represents Span mirror interval
type SpanMirrorInterval string

const (
    // Mirror 1 in every 512 packets
    SpanMirrorInterval_Y_512 SpanMirrorInterval = "512"

    // Mirror 1 in every 1024 packets
    SpanMirrorInterval_Y_1k SpanMirrorInterval = "1k"

    // Mirror 1 in every 2048 packets
    SpanMirrorInterval_Y_2k SpanMirrorInterval = "2k"

    // Mirror 1 in every 4096 packets
    SpanMirrorInterval_Y_4k SpanMirrorInterval = "4k"

    // Mirror 1 in every 8192 packets
    SpanMirrorInterval_Y_8k SpanMirrorInterval = "8k"

    // Mirror 1 in every 16384 packets
    SpanMirrorInterval_Y_16k SpanMirrorInterval = "16k"
)

// SpanDestination represents Span destination
type SpanDestination string

const (
    // Destination Interface
    SpanDestination_interface_ SpanDestination = "interface"

    // Destination Pseudowire
    SpanDestination_pseudowire SpanDestination = "pseudowire"

    // Destination next-hop IPv4 address
    SpanDestination_ipv4_address SpanDestination = "ipv4-address"

    // Destination next-hop IPv6 address
    SpanDestination_ipv6_address SpanDestination = "ipv6-address"
)

// SpanMonitorSession
// none
type SpanMonitorSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Monitor-session configuration commands.
    Sessions SpanMonitorSession_Sessions
}

func (spanMonitorSession *SpanMonitorSession) GetFilter() yfilter.YFilter { return spanMonitorSession.YFilter }

func (spanMonitorSession *SpanMonitorSession) SetFilter(yf yfilter.YFilter) { spanMonitorSession.YFilter = yf }

func (spanMonitorSession *SpanMonitorSession) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (spanMonitorSession *SpanMonitorSession) GetSegmentPath() string {
    return "Cisco-IOS-XR-Ethernet-SPAN-cfg:span-monitor-session"
}

func (spanMonitorSession *SpanMonitorSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &spanMonitorSession.Sessions
    }
    return nil
}

func (spanMonitorSession *SpanMonitorSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &spanMonitorSession.Sessions
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

func (spanMonitorSession *SpanMonitorSession) GetParentYangName() string { return "Cisco-IOS-XR-Ethernet-SPAN-cfg" }

// SpanMonitorSession_Sessions
// Monitor-session configuration commands
type SpanMonitorSession_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a particular Monitor Session. The type is slice of
    // SpanMonitorSession_Sessions_Session.
    Session []SpanMonitorSession_Sessions_Session
}

func (sessions *SpanMonitorSession_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *SpanMonitorSession_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *SpanMonitorSession_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *SpanMonitorSession_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *SpanMonitorSession_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := SpanMonitorSession_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *SpanMonitorSession_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *SpanMonitorSession_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *SpanMonitorSession_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *SpanMonitorSession_Sessions) GetYangName() string { return "sessions" }

func (sessions *SpanMonitorSession_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *SpanMonitorSession_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *SpanMonitorSession_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *SpanMonitorSession_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *SpanMonitorSession_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *SpanMonitorSession_Sessions) GetParentYangName() string { return "span-monitor-session" }

// SpanMonitorSession_Sessions_Session
// Configuration for a particular Monitor Session
type SpanMonitorSession_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session Name. The type is string with length:
    // 1..79.
    Session interface{}

    // Enable a Monitor Session.  Setting this item causes the Monitor Session to
    // be created. The type is SpanSessionClass. The default value is ethernet.
    Class interface{}

    // Specify a destination.
    Destination SpanMonitorSession_Sessions_Session_Destination
}

func (session *SpanMonitorSession_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *SpanMonitorSession_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *SpanMonitorSession_Sessions_Session) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    if yname == "class" { return "Class" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (session *SpanMonitorSession_Sessions_Session) GetSegmentPath() string {
    return "session" + "[session='" + fmt.Sprintf("%v", session.Session) + "']"
}

func (session *SpanMonitorSession_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination" {
        return &session.Destination
    }
    return nil
}

func (session *SpanMonitorSession_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["destination"] = &session.Destination
    return children
}

func (session *SpanMonitorSession_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session"] = session.Session
    leafs["class"] = session.Class
    return leafs
}

func (session *SpanMonitorSession_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *SpanMonitorSession_Sessions_Session) GetYangName() string { return "session" }

func (session *SpanMonitorSession_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *SpanMonitorSession_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *SpanMonitorSession_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *SpanMonitorSession_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *SpanMonitorSession_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *SpanMonitorSession_Sessions_Session) GetParentYangName() string { return "sessions" }

// SpanMonitorSession_Sessions_Session_Destination
// Specify a destination
type SpanMonitorSession_Sessions_Session_Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the type of destination. The type is SpanDestination.
    DestinationType interface{}

    // Specify the destination interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    DestinationInterfaceName interface{}

    // Specify the destination next-hop IPv4 address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationIpv4Address interface{}

    // Specify the destination next-hop IPv6 address. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationIpv6Address interface{}
}

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *SpanMonitorSession_Sessions_Session_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetGoName(yname string) string {
    if yname == "destination-type" { return "DestinationType" }
    if yname == "destination-interface-name" { return "DestinationInterfaceName" }
    if yname == "destination-ipv4-address" { return "DestinationIpv4Address" }
    if yname == "destination-ipv6-address" { return "DestinationIpv6Address" }
    return ""
}

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetSegmentPath() string {
    return "destination"
}

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-type"] = destination.DestinationType
    leafs["destination-interface-name"] = destination.DestinationInterfaceName
    leafs["destination-ipv4-address"] = destination.DestinationIpv4Address
    leafs["destination-ipv6-address"] = destination.DestinationIpv6Address
    return leafs
}

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetBundleName() string { return "cisco_ios_xr" }

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetYangName() string { return "destination" }

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destination *SpanMonitorSession_Sessions_Session_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetParent() types.Entity { return destination.parent }

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetParentYangName() string { return "session" }

