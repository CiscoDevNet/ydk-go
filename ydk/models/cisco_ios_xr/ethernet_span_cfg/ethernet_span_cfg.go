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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Monitor-session configuration commands.
    Sessions SpanMonitorSession_Sessions
}

func (spanMonitorSession *SpanMonitorSession) GetEntityData() *types.CommonEntityData {
    spanMonitorSession.EntityData.YFilter = spanMonitorSession.YFilter
    spanMonitorSession.EntityData.YangName = "span-monitor-session"
    spanMonitorSession.EntityData.BundleName = "cisco_ios_xr"
    spanMonitorSession.EntityData.ParentYangName = "Cisco-IOS-XR-Ethernet-SPAN-cfg"
    spanMonitorSession.EntityData.SegmentPath = "Cisco-IOS-XR-Ethernet-SPAN-cfg:span-monitor-session"
    spanMonitorSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMonitorSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMonitorSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMonitorSession.EntityData.Children = make(map[string]types.YChild)
    spanMonitorSession.EntityData.Children["sessions"] = types.YChild{"Sessions", &spanMonitorSession.Sessions}
    spanMonitorSession.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(spanMonitorSession.EntityData)
}

// SpanMonitorSession_Sessions
// Monitor-session configuration commands
type SpanMonitorSession_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular Monitor Session. The type is slice of
    // SpanMonitorSession_Sessions_Session.
    Session []SpanMonitorSession_Sessions_Session
}

func (sessions *SpanMonitorSession_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "span-monitor-session"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["session"] = types.YChild{"Session", nil}
    for i := range sessions.Session {
        sessions.EntityData.Children[types.GetSegmentPath(&sessions.Session[i])] = types.YChild{"Session", &sessions.Session[i]}
    }
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// SpanMonitorSession_Sessions_Session
// Configuration for a particular Monitor Session
type SpanMonitorSession_Sessions_Session struct {
    EntityData types.CommonEntityData
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

func (session *SpanMonitorSession_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + "[session='" + fmt.Sprintf("%v", session.Session) + "']"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["destination"] = types.YChild{"Destination", &session.Destination}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["session"] = types.YLeaf{"Session", session.Session}
    session.EntityData.Leafs["class"] = types.YLeaf{"Class", session.Class}
    return &(session.EntityData)
}

// SpanMonitorSession_Sessions_Session_Destination
// Specify a destination
type SpanMonitorSession_Sessions_Session_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the type of destination. The type is SpanDestination.
    DestinationType interface{}

    // Specify the destination interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    DestinationInterfaceName interface{}

    // Specify the destination next-hop IPv4 address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationIpv4Address interface{}

    // Specify the destination next-hop IPv6 address. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DestinationIpv6Address interface{}
}

func (destination *SpanMonitorSession_Sessions_Session_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "session"
    destination.EntityData.SegmentPath = "destination"
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = make(map[string]types.YChild)
    destination.EntityData.Leafs = make(map[string]types.YLeaf)
    destination.EntityData.Leafs["destination-type"] = types.YLeaf{"DestinationType", destination.DestinationType}
    destination.EntityData.Leafs["destination-interface-name"] = types.YLeaf{"DestinationInterfaceName", destination.DestinationInterfaceName}
    destination.EntityData.Leafs["destination-ipv4-address"] = types.YLeaf{"DestinationIpv4Address", destination.DestinationIpv4Address}
    destination.EntityData.Leafs["destination-ipv6-address"] = types.YLeaf{"DestinationIpv6Address", destination.DestinationIpv6Address}
    return &(destination.EntityData)
}

