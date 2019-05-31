// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-infra-tmplmgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   dynamic-template: All dynamic template configurations
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_infra_tmplmgr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_infra_tmplmgr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg dynamic-template}", reflect.TypeOf(DynamicTemplate{}))
    ydk.RegisterEntity("Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template", reflect.TypeOf(DynamicTemplate{}))
}

// DynamicTemplate
// All dynamic template configurations
type DynamicTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Templates of the PPP Type.
    Ppps DynamicTemplate_Ppps

    // The IP Subscriber Template Table.
    IpSubscribers DynamicTemplate_IpSubscribers

    // The Service Type Template Table.
    SubscriberServices DynamicTemplate_SubscriberServices
}

func (dynamicTemplate *DynamicTemplate) GetEntityData() *types.CommonEntityData {
    dynamicTemplate.EntityData.YFilter = dynamicTemplate.YFilter
    dynamicTemplate.EntityData.YangName = "dynamic-template"
    dynamicTemplate.EntityData.BundleName = "cisco_ios_xr"
    dynamicTemplate.EntityData.ParentYangName = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg"
    dynamicTemplate.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template"
    dynamicTemplate.EntityData.AbsolutePath = dynamicTemplate.EntityData.SegmentPath
    dynamicTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicTemplate.EntityData.Children = types.NewOrderedMap()
    dynamicTemplate.EntityData.Children.Append("ppps", types.YChild{"Ppps", &dynamicTemplate.Ppps})
    dynamicTemplate.EntityData.Children.Append("ip-subscribers", types.YChild{"IpSubscribers", &dynamicTemplate.IpSubscribers})
    dynamicTemplate.EntityData.Children.Append("subscriber-services", types.YChild{"SubscriberServices", &dynamicTemplate.SubscriberServices})
    dynamicTemplate.EntityData.Leafs = types.NewOrderedMap()

    dynamicTemplate.EntityData.YListKeys = []string {}

    return &(dynamicTemplate.EntityData)
}

// DynamicTemplate_Ppps
// Templates of the PPP Type
type DynamicTemplate_Ppps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Template of the PPP Type. The type is slice of DynamicTemplate_Ppps_Ppp.
    Ppp []*DynamicTemplate_Ppps_Ppp
}

func (ppps *DynamicTemplate_Ppps) GetEntityData() *types.CommonEntityData {
    ppps.EntityData.YFilter = ppps.YFilter
    ppps.EntityData.YangName = "ppps"
    ppps.EntityData.BundleName = "cisco_ios_xr"
    ppps.EntityData.ParentYangName = "dynamic-template"
    ppps.EntityData.SegmentPath = "ppps"
    ppps.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/" + ppps.EntityData.SegmentPath
    ppps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ppps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ppps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ppps.EntityData.Children = types.NewOrderedMap()
    ppps.EntityData.Children.Append("ppp", types.YChild{"Ppp", nil})
    for i := range ppps.Ppp {
        ppps.EntityData.Children.Append(types.GetSegmentPath(ppps.Ppp[i]), types.YChild{"Ppp", ppps.Ppp[i]})
    }
    ppps.EntityData.Leafs = types.NewOrderedMap()

    ppps.EntityData.YListKeys = []string {}

    return &(ppps.EntityData)
}

// DynamicTemplate_Ppps_Ppp
// A Template of the PPP Type
type DynamicTemplate_Ppps_Ppp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of the template. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TemplateName interface{}

    // Assign the interface to a VRF . The type is string with length: 1..32.
    Vrf interface{}

    // Monitor Session container for this template.
    SpanMonitorSessions DynamicTemplate_Ppps_Ppp_SpanMonitorSessions

    // IPv4 Packet Filtering configuration for the template.
    Ipv4PacketFilter DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter

    // IPv6 Packet Filtering configuration for the interface.
    Ipv6PacketFilter DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter

    // IGMPconfiguration.
    Igmp DynamicTemplate_Ppps_Ppp_Igmp

    // Interface IPv4 Network configuration data.
    Ipv4Network DynamicTemplate_Ppps_Ppp_Ipv4Network

    // Interface IPv6 Network configuration data.
    Ipv6Network DynamicTemplate_Ppps_Ppp_Ipv6Network

    // Interface IPv6 Network configuration data.
    Ipv6Neighbor DynamicTemplate_Ppps_Ppp_Ipv6Neighbor

    // Interface dhcpv6 configuration data.
    Dhcpv6 DynamicTemplate_Ppps_Ppp_Dhcpv6

    // Dynamic Template PBR configuration.
    Pbr DynamicTemplate_Ppps_Ppp_Pbr

    // PPP template configuration data.
    PppTemplate DynamicTemplate_Ppps_Ppp_PppTemplate

    // QoS dynamically applied configuration template.
    Qos DynamicTemplate_Ppps_Ppp_Qos

    // Subscriber accounting dynamic-template commands.
    Accounting DynamicTemplate_Ppps_Ppp_Accounting

    // PPPoE template configuration data.
    PppoeTemplate DynamicTemplate_Ppps_Ppp_PppoeTemplate
}

func (ppp *DynamicTemplate_Ppps_Ppp) GetEntityData() *types.CommonEntityData {
    ppp.EntityData.YFilter = ppp.YFilter
    ppp.EntityData.YangName = "ppp"
    ppp.EntityData.BundleName = "cisco_ios_xr"
    ppp.EntityData.ParentYangName = "ppps"
    ppp.EntityData.SegmentPath = "ppp" + types.AddKeyToken(ppp.TemplateName, "template-name")
    ppp.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/" + ppp.EntityData.SegmentPath
    ppp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ppp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ppp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ppp.EntityData.Children = types.NewOrderedMap()
    ppp.EntityData.Children.Append("Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions", types.YChild{"SpanMonitorSessions", &ppp.SpanMonitorSessions})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter", types.YChild{"Ipv4PacketFilter", &ppp.Ipv4PacketFilter})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter", types.YChild{"Ipv6PacketFilter", &ppp.Ipv6PacketFilter})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp", types.YChild{"Igmp", &ppp.Igmp})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network", types.YChild{"Ipv4Network", &ppp.Ipv4Network})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network", types.YChild{"Ipv6Network", &ppp.Ipv6Network})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor", types.YChild{"Ipv6Neighbor", &ppp.Ipv6Neighbor})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6", types.YChild{"Dhcpv6", &ppp.Dhcpv6})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-pbr-subscriber-cfg:pbr", types.YChild{"Pbr", &ppp.Pbr})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template", types.YChild{"PppTemplate", &ppp.PppTemplate})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-qos-ma-bng-cfg:qos", types.YChild{"Qos", &ppp.Qos})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-subscriber-accounting-cfg:accounting", types.YChild{"Accounting", &ppp.Accounting})
    ppp.EntityData.Children.Append("Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-template", types.YChild{"PppoeTemplate", &ppp.PppoeTemplate})
    ppp.EntityData.Leafs = types.NewOrderedMap()
    ppp.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ppp.TemplateName})
    ppp.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", ppp.Vrf})

    ppp.EntityData.YListKeys = []string {"TemplateName"}

    return &(ppp.EntityData)
}

// DynamicTemplate_Ppps_Ppp_SpanMonitorSessions
// Monitor Session container for this template
type DynamicTemplate_Ppps_Ppp_SpanMonitorSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular class of Monitor Session. The type is slice
    // of DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession.
    SpanMonitorSession []*DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession
}

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetEntityData() *types.CommonEntityData {
    spanMonitorSessions.EntityData.YFilter = spanMonitorSessions.YFilter
    spanMonitorSessions.EntityData.YangName = "span-monitor-sessions"
    spanMonitorSessions.EntityData.BundleName = "cisco_ios_xr"
    spanMonitorSessions.EntityData.ParentYangName = "ppp"
    spanMonitorSessions.EntityData.SegmentPath = "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions"
    spanMonitorSessions.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + spanMonitorSessions.EntityData.SegmentPath
    spanMonitorSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMonitorSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMonitorSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMonitorSessions.EntityData.Children = types.NewOrderedMap()
    spanMonitorSessions.EntityData.Children.Append("span-monitor-session", types.YChild{"SpanMonitorSession", nil})
    for i := range spanMonitorSessions.SpanMonitorSession {
        spanMonitorSessions.EntityData.Children.Append(types.GetSegmentPath(spanMonitorSessions.SpanMonitorSession[i]), types.YChild{"SpanMonitorSession", spanMonitorSessions.SpanMonitorSession[i]})
    }
    spanMonitorSessions.EntityData.Leafs = types.NewOrderedMap()

    spanMonitorSessions.EntityData.YListKeys = []string {}

    return &(spanMonitorSessions.EntityData)
}

// DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession
// Configuration for a particular class of Monitor
// Session
type DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Session Class. The type is SpanSessionClass.
    SessionClass interface{}

    // Mirror a specified number of bytes from start of packet. The type is
    // interface{} with range: 1..10000. Units are byte.
    MirrorFirst interface{}

    // Specify the mirror interval. The type is SpanMirrorInterval.
    MirrorInterval interface{}

    // Attach the interface to a Monitor Session.
    Attachment DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment

    // Enable ACL matching for traffic mirroring.
    Acl DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl
}

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetEntityData() *types.CommonEntityData {
    spanMonitorSession.EntityData.YFilter = spanMonitorSession.YFilter
    spanMonitorSession.EntityData.YangName = "span-monitor-session"
    spanMonitorSession.EntityData.BundleName = "cisco_ios_xr"
    spanMonitorSession.EntityData.ParentYangName = "span-monitor-sessions"
    spanMonitorSession.EntityData.SegmentPath = "span-monitor-session" + types.AddKeyToken(spanMonitorSession.SessionClass, "session-class")
    spanMonitorSession.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions/" + spanMonitorSession.EntityData.SegmentPath
    spanMonitorSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMonitorSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMonitorSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMonitorSession.EntityData.Children = types.NewOrderedMap()
    spanMonitorSession.EntityData.Children.Append("attachment", types.YChild{"Attachment", &spanMonitorSession.Attachment})
    spanMonitorSession.EntityData.Children.Append("acl", types.YChild{"Acl", &spanMonitorSession.Acl})
    spanMonitorSession.EntityData.Leafs = types.NewOrderedMap()
    spanMonitorSession.EntityData.Leafs.Append("session-class", types.YLeaf{"SessionClass", spanMonitorSession.SessionClass})
    spanMonitorSession.EntityData.Leafs.Append("mirror-first", types.YLeaf{"MirrorFirst", spanMonitorSession.MirrorFirst})
    spanMonitorSession.EntityData.Leafs.Append("mirror-interval", types.YLeaf{"MirrorInterval", spanMonitorSession.MirrorInterval})

    spanMonitorSession.EntityData.YListKeys = []string {"SessionClass"}

    return &(spanMonitorSession.EntityData)
}

// DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment
// Attach the interface to a Monitor Session
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Session Name. The type is string with length: 1..79. This attribute is
    // mandatory.
    SessionName interface{}

    // Specify the direction of traffic to replicate (optional). The type is
    // SpanTrafficDirection.
    Direction interface{}

    // Enable port level traffic mirroring. The type is interface{}.
    PortLevelEnable interface{}
}

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetEntityData() *types.CommonEntityData {
    attachment.EntityData.YFilter = attachment.YFilter
    attachment.EntityData.YangName = "attachment"
    attachment.EntityData.BundleName = "cisco_ios_xr"
    attachment.EntityData.ParentYangName = "span-monitor-session"
    attachment.EntityData.SegmentPath = "attachment"
    attachment.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions/span-monitor-session/" + attachment.EntityData.SegmentPath
    attachment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachment.EntityData.Children = types.NewOrderedMap()
    attachment.EntityData.Leafs = types.NewOrderedMap()
    attachment.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", attachment.SessionName})
    attachment.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", attachment.Direction})
    attachment.EntityData.Leafs.Append("port-level-enable", types.YLeaf{"PortLevelEnable", attachment.PortLevelEnable})

    attachment.EntityData.YListKeys = []string {}

    return &(attachment.EntityData)
}

// DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl
// Enable ACL matching for traffic mirroring
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable ACL. The type is interface{}. This attribute is mandatory.
    AclEnable interface{}

    // ACL Name. The type is string with length: 1..80.
    AclName interface{}
}

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetEntityData() *types.CommonEntityData {
    acl.EntityData.YFilter = acl.YFilter
    acl.EntityData.YangName = "acl"
    acl.EntityData.BundleName = "cisco_ios_xr"
    acl.EntityData.ParentYangName = "span-monitor-session"
    acl.EntityData.SegmentPath = "acl"
    acl.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions/span-monitor-session/" + acl.EntityData.SegmentPath
    acl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acl.EntityData.Children = types.NewOrderedMap()
    acl.EntityData.Leafs = types.NewOrderedMap()
    acl.EntityData.Leafs.Append("acl-enable", types.YLeaf{"AclEnable", acl.AclEnable})
    acl.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", acl.AclName})

    acl.EntityData.YListKeys = []string {}

    return &(acl.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter
// IPv4 Packet Filtering configuration for the
// template
type DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound

    // IPv4 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound
}

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetEntityData() *types.CommonEntityData {
    ipv4PacketFilter.EntityData.YFilter = ipv4PacketFilter.YFilter
    ipv4PacketFilter.EntityData.YangName = "ipv4-packet-filter"
    ipv4PacketFilter.EntityData.BundleName = "cisco_ios_xr"
    ipv4PacketFilter.EntityData.ParentYangName = "ppp"
    ipv4PacketFilter.EntityData.SegmentPath = "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter"
    ipv4PacketFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + ipv4PacketFilter.EntityData.SegmentPath
    ipv4PacketFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PacketFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PacketFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PacketFilter.EntityData.Children = types.NewOrderedMap()
    ipv4PacketFilter.EntityData.Children.Append("outbound", types.YChild{"Outbound", &ipv4PacketFilter.Outbound})
    ipv4PacketFilter.EntityData.Children.Append("inbound", types.YChild{"Inbound", &ipv4PacketFilter.Inbound})
    ipv4PacketFilter.EntityData.Leafs = types.NewOrderedMap()

    ipv4PacketFilter.EntityData.YListKeys = []string {}

    return &(ipv4PacketFilter.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound
// IPv4 Packet filter to be applied to outbound
// packets
type DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "ipv4-packet-filter"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter/" + outbound.EntityData.SegmentPath
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = types.NewOrderedMap()
    outbound.EntityData.Leafs = types.NewOrderedMap()
    outbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", outbound.CommonAclName})
    outbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", outbound.Name})
    outbound.EntityData.Leafs.Append("hardware-count", types.YLeaf{"HardwareCount", outbound.HardwareCount})
    outbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", outbound.InterfaceStatistics})

    outbound.EntityData.YListKeys = []string {}

    return &(outbound.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound
// IPv4 Packet filter to be applied to inbound
// packets
type DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Inbound packets NOTE: This
    // parameter is mandatory if 'CommonACLName' is not specified. The type is
    // string with length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "ipv4-packet-filter"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter/" + inbound.EntityData.SegmentPath
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = types.NewOrderedMap()
    inbound.EntityData.Leafs = types.NewOrderedMap()
    inbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", inbound.CommonAclName})
    inbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", inbound.Name})
    inbound.EntityData.Leafs.Append("hardware-count", types.YLeaf{"HardwareCount", inbound.HardwareCount})
    inbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", inbound.InterfaceStatistics})

    inbound.EntityData.YListKeys = []string {}

    return &(inbound.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter
// IPv6 Packet Filtering configuration for the
// interface
type DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound

    // IPv6 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound
}

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetEntityData() *types.CommonEntityData {
    ipv6PacketFilter.EntityData.YFilter = ipv6PacketFilter.YFilter
    ipv6PacketFilter.EntityData.YangName = "ipv6-packet-filter"
    ipv6PacketFilter.EntityData.BundleName = "cisco_ios_xr"
    ipv6PacketFilter.EntityData.ParentYangName = "ppp"
    ipv6PacketFilter.EntityData.SegmentPath = "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter"
    ipv6PacketFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + ipv6PacketFilter.EntityData.SegmentPath
    ipv6PacketFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6PacketFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6PacketFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6PacketFilter.EntityData.Children = types.NewOrderedMap()
    ipv6PacketFilter.EntityData.Children.Append("inbound", types.YChild{"Inbound", &ipv6PacketFilter.Inbound})
    ipv6PacketFilter.EntityData.Children.Append("outbound", types.YChild{"Outbound", &ipv6PacketFilter.Outbound})
    ipv6PacketFilter.EntityData.Leafs = types.NewOrderedMap()

    ipv6PacketFilter.EntityData.YListKeys = []string {}

    return &(ipv6PacketFilter.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound
// IPv6 Packet filter to be applied to inbound
// packets
type DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Inbound  NOTE: This parameter is
    // mandatory if 'CommonACLName' is not specified. The type is string with
    // length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "ipv6-packet-filter"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter/" + inbound.EntityData.SegmentPath
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = types.NewOrderedMap()
    inbound.EntityData.Leafs = types.NewOrderedMap()
    inbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", inbound.CommonAclName})
    inbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", inbound.Name})
    inbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", inbound.InterfaceStatistics})

    inbound.EntityData.YListKeys = []string {}

    return &(inbound.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound
// IPv6 Packet filter to be applied to outbound
// packets
type DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "ipv6-packet-filter"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter/" + outbound.EntityData.SegmentPath
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = types.NewOrderedMap()
    outbound.EntityData.Leafs = types.NewOrderedMap()
    outbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", outbound.CommonAclName})
    outbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", outbound.Name})
    outbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", outbound.InterfaceStatistics})

    outbound.EntityData.YListKeys = []string {}

    return &(outbound.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Igmp
// IGMPconfiguration
type DynamicTemplate_Ppps_Ppp_Igmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default VRF.
    DefaultVrf DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf
}

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetEntityData() *types.CommonEntityData {
    igmp.EntityData.YFilter = igmp.YFilter
    igmp.EntityData.YangName = "igmp"
    igmp.EntityData.BundleName = "cisco_ios_xr"
    igmp.EntityData.ParentYangName = "ppp"
    igmp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp"
    igmp.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + igmp.EntityData.SegmentPath
    igmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igmp.EntityData.Children = types.NewOrderedMap()
    igmp.EntityData.Children.Append("default-vrf", types.YChild{"DefaultVrf", &igmp.DefaultVrf})
    igmp.EntityData.Leafs = types.NewOrderedMap()

    igmp.EntityData.YListKeys = []string {}

    return &(igmp.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf
// Default VRF
type DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGMP Max Groups. The type is interface{} with range: 1..40000. The default
    // value is 25000.
    MaxGroups interface{}

    // Access list specifying access-list group range. The type is string with
    // length: 1..64.
    AccessGroup interface{}

    // IGMP Version. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // Query response value in seconds. The type is interface{} with range: 1..25.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Configure Multicast mode variable. The type is DynTmplMulticastMode.
    MulticastMode interface{}

    // IGMPv3 explicit host tracking.
    ExplicitTracking DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking
}

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "igmp"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp/" + defaultVrf.EntityData.SegmentPath
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = types.NewOrderedMap()
    defaultVrf.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &defaultVrf.ExplicitTracking})
    defaultVrf.EntityData.Leafs = types.NewOrderedMap()
    defaultVrf.EntityData.Leafs.Append("max-groups", types.YLeaf{"MaxGroups", defaultVrf.MaxGroups})
    defaultVrf.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", defaultVrf.AccessGroup})
    defaultVrf.EntityData.Leafs.Append("version", types.YLeaf{"Version", defaultVrf.Version})
    defaultVrf.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", defaultVrf.QueryInterval})
    defaultVrf.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", defaultVrf.QueryMaxResponseTime})
    defaultVrf.EntityData.Leafs.Append("multicast-mode", types.YLeaf{"MulticastMode", defaultVrf.MulticastMode})

    defaultVrf.EntityData.YListKeys = []string {}

    return &(defaultVrf.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable or disable, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "default-vrf"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp/default-vrf/" + explicitTracking.EntityData.SegmentPath
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv4Network
// Interface IPv4 Network configuration data
type DynamicTemplate_Ppps_Ppp_Ipv4Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IP processing without an explicit address. The type is string.
    Unnumbered interface{}

    // The IP Maximum Transmission Unit. The type is interface{} with range:
    // 68..65535. Units are byte.
    Mtu interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool. The default value is
    // false.
    Unreachables interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool. The default value is
    // true.
    Rpf interface{}
}

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetEntityData() *types.CommonEntityData {
    ipv4Network.EntityData.YFilter = ipv4Network.YFilter
    ipv4Network.EntityData.YangName = "ipv4-network"
    ipv4Network.EntityData.BundleName = "cisco_ios_xr"
    ipv4Network.EntityData.ParentYangName = "ppp"
    ipv4Network.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network"
    ipv4Network.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + ipv4Network.EntityData.SegmentPath
    ipv4Network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Network.EntityData.Children = types.NewOrderedMap()
    ipv4Network.EntityData.Leafs = types.NewOrderedMap()
    ipv4Network.EntityData.Leafs.Append("unnumbered", types.YLeaf{"Unnumbered", ipv4Network.Unnumbered})
    ipv4Network.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", ipv4Network.Mtu})
    ipv4Network.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", ipv4Network.Unreachables})
    ipv4Network.EntityData.Leafs.Append("rpf", types.YLeaf{"Rpf", ipv4Network.Rpf})

    ipv4Network.EntityData.YListKeys = []string {}

    return &(ipv4Network.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6Network
// Interface IPv6 Network configuration data
type DynamicTemplate_Ppps_Ppp_Ipv6Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MTU Setting of Interface. The type is interface{} with range: 1280..65535.
    // Units are byte.
    Mtu interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool.
    Rpf interface{}

    // Override Sending of ICMP Unreachable Messages. The type is interface{}.
    Unreachables interface{}

    // Set the IPv6 address of an interface.
    Addresses DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses
}

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetEntityData() *types.CommonEntityData {
    ipv6Network.EntityData.YFilter = ipv6Network.YFilter
    ipv6Network.EntityData.YangName = "ipv6-network"
    ipv6Network.EntityData.BundleName = "cisco_ios_xr"
    ipv6Network.EntityData.ParentYangName = "ppp"
    ipv6Network.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network"
    ipv6Network.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + ipv6Network.EntityData.SegmentPath
    ipv6Network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Network.EntityData.Children = types.NewOrderedMap()
    ipv6Network.EntityData.Children.Append("addresses", types.YChild{"Addresses", &ipv6Network.Addresses})
    ipv6Network.EntityData.Leafs = types.NewOrderedMap()
    ipv6Network.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", ipv6Network.Mtu})
    ipv6Network.EntityData.Leafs.Append("rpf", types.YLeaf{"Rpf", ipv6Network.Rpf})
    ipv6Network.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", ipv6Network.Unreachables})

    ipv6Network.EntityData.YListKeys = []string {}

    return &(ipv6Network.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses
// Set the IPv6 address of an interface
type DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto IPv6 Interface Configuration.
    AutoConfiguration DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration
}

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "ipv6-network"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network/" + addresses.EntityData.SegmentPath
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("auto-configuration", types.YChild{"AutoConfiguration", &addresses.AutoConfiguration})
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration
// Auto IPv6 Interface Configuration
type DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The flag to enable auto ipv6 interface configuration. The type is
    // interface{}.
    Enable interface{}
}

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetEntityData() *types.CommonEntityData {
    autoConfiguration.EntityData.YFilter = autoConfiguration.YFilter
    autoConfiguration.EntityData.YangName = "auto-configuration"
    autoConfiguration.EntityData.BundleName = "cisco_ios_xr"
    autoConfiguration.EntityData.ParentYangName = "addresses"
    autoConfiguration.EntityData.SegmentPath = "auto-configuration"
    autoConfiguration.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network/addresses/" + autoConfiguration.EntityData.SegmentPath
    autoConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoConfiguration.EntityData.Children = types.NewOrderedMap()
    autoConfiguration.EntityData.Leafs = types.NewOrderedMap()
    autoConfiguration.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", autoConfiguration.Enable})

    autoConfiguration.EntityData.YListKeys = []string {}

    return &(autoConfiguration.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor
// Interface IPv6 Network configuration data
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the IPv6 framed ipv6 prefix pool for a subscriber interface . The type
    // is string.
    FramedPrefixPool interface{}

    // Host to use stateful protocol for address configuration. The type is
    // interface{}.
    ManagedConfig interface{}

    // Host to use stateful protocol for non-address configuration. The type is
    // interface{}.
    OtherConfig interface{}

    // Start RA on ipv6-enable config. The type is interface{}.
    StartRaOnIpv6Enable interface{}

    // NUD enable. The type is interface{}.
    NudEnable interface{}

    // Set IPv6 Router Advertisement (RA) lifetime in seconds. The type is
    // interface{} with range: 0..9000. Units are second.
    RaLifetime interface{}

    // RA Router Preference. The type is Ipv6NdRouterPrefTemplate.
    RouterPreference interface{}

    // Enable suppress IPv6 router advertisement. The type is interface{}.
    RaSuppress interface{}

    // Enable RA unicast Flag. The type is interface{}.
    RaUnicast interface{}

    // Unspecify IPv6 Router Advertisement (RA) hop-limit. The type is
    // interface{}.
    RaUnspecifyHoplimit interface{}

    // RA suppress MTU flag. The type is interface{}.
    RaSuppressMtu interface{}

    // Suppress cache learning flag. The type is interface{}.
    SuppressCacheLearning interface{}

    // Set advertised reachability time in milliseconds. The type is interface{}
    // with range: 0..3600000. Units are millisecond.
    ReachableTime interface{}

    // Set advertised NS retransmission interval in milliseconds. The type is
    // interface{} with range: 1000..4294967295. Units are millisecond.
    NsInterval interface{}

    // Set IPv6 Router Advertisement (RA) interval in seconds.
    RaInterval DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval

    // Set the IPv6 framed ipv6 prefix for a subscriber interface .
    FramedPrefix DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix

    // Duplicate Address Detection (DAD).
    DuplicateAddressDetection DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection

    // IPv6 ND RA Initial.
    RaInitial DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial
}

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetEntityData() *types.CommonEntityData {
    ipv6Neighbor.EntityData.YFilter = ipv6Neighbor.YFilter
    ipv6Neighbor.EntityData.YangName = "ipv6-neighbor"
    ipv6Neighbor.EntityData.BundleName = "cisco_ios_xr"
    ipv6Neighbor.EntityData.ParentYangName = "ppp"
    ipv6Neighbor.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor"
    ipv6Neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + ipv6Neighbor.EntityData.SegmentPath
    ipv6Neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Neighbor.EntityData.Children = types.NewOrderedMap()
    ipv6Neighbor.EntityData.Children.Append("ra-interval", types.YChild{"RaInterval", &ipv6Neighbor.RaInterval})
    ipv6Neighbor.EntityData.Children.Append("framed-prefix", types.YChild{"FramedPrefix", &ipv6Neighbor.FramedPrefix})
    ipv6Neighbor.EntityData.Children.Append("duplicate-address-detection", types.YChild{"DuplicateAddressDetection", &ipv6Neighbor.DuplicateAddressDetection})
    ipv6Neighbor.EntityData.Children.Append("ra-initial", types.YChild{"RaInitial", &ipv6Neighbor.RaInitial})
    ipv6Neighbor.EntityData.Leafs = types.NewOrderedMap()
    ipv6Neighbor.EntityData.Leafs.Append("framed-prefix-pool", types.YLeaf{"FramedPrefixPool", ipv6Neighbor.FramedPrefixPool})
    ipv6Neighbor.EntityData.Leafs.Append("managed-config", types.YLeaf{"ManagedConfig", ipv6Neighbor.ManagedConfig})
    ipv6Neighbor.EntityData.Leafs.Append("other-config", types.YLeaf{"OtherConfig", ipv6Neighbor.OtherConfig})
    ipv6Neighbor.EntityData.Leafs.Append("start-ra-on-ipv6-enable", types.YLeaf{"StartRaOnIpv6Enable", ipv6Neighbor.StartRaOnIpv6Enable})
    ipv6Neighbor.EntityData.Leafs.Append("nud-enable", types.YLeaf{"NudEnable", ipv6Neighbor.NudEnable})
    ipv6Neighbor.EntityData.Leafs.Append("ra-lifetime", types.YLeaf{"RaLifetime", ipv6Neighbor.RaLifetime})
    ipv6Neighbor.EntityData.Leafs.Append("router-preference", types.YLeaf{"RouterPreference", ipv6Neighbor.RouterPreference})
    ipv6Neighbor.EntityData.Leafs.Append("ra-suppress", types.YLeaf{"RaSuppress", ipv6Neighbor.RaSuppress})
    ipv6Neighbor.EntityData.Leafs.Append("ra-unicast", types.YLeaf{"RaUnicast", ipv6Neighbor.RaUnicast})
    ipv6Neighbor.EntityData.Leafs.Append("ra-unspecify-hoplimit", types.YLeaf{"RaUnspecifyHoplimit", ipv6Neighbor.RaUnspecifyHoplimit})
    ipv6Neighbor.EntityData.Leafs.Append("ra-suppress-mtu", types.YLeaf{"RaSuppressMtu", ipv6Neighbor.RaSuppressMtu})
    ipv6Neighbor.EntityData.Leafs.Append("suppress-cache-learning", types.YLeaf{"SuppressCacheLearning", ipv6Neighbor.SuppressCacheLearning})
    ipv6Neighbor.EntityData.Leafs.Append("reachable-time", types.YLeaf{"ReachableTime", ipv6Neighbor.ReachableTime})
    ipv6Neighbor.EntityData.Leafs.Append("ns-interval", types.YLeaf{"NsInterval", ipv6Neighbor.NsInterval})

    ipv6Neighbor.EntityData.YListKeys = []string {}

    return &(ipv6Neighbor.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval
// Set IPv6 Router Advertisement (RA) interval in
// seconds
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Maximum interface{}

    // Minimum RA interval in seconds. Must be less than 0.75 * maximum interval.
    // The type is interface{} with range: 3..1800. Units are second.
    Minimum interface{}
}

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetEntityData() *types.CommonEntityData {
    raInterval.EntityData.YFilter = raInterval.YFilter
    raInterval.EntityData.YangName = "ra-interval"
    raInterval.EntityData.BundleName = "cisco_ios_xr"
    raInterval.EntityData.ParentYangName = "ipv6-neighbor"
    raInterval.EntityData.SegmentPath = "ra-interval"
    raInterval.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + raInterval.EntityData.SegmentPath
    raInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    raInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    raInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    raInterval.EntityData.Children = types.NewOrderedMap()
    raInterval.EntityData.Leafs = types.NewOrderedMap()
    raInterval.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", raInterval.Maximum})
    raInterval.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", raInterval.Minimum})

    raInterval.EntityData.YListKeys = []string {}

    return &(raInterval.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix
// Set the IPv6 framed ipv6 prefix for a
// subscriber interface 
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IPv6 framed prefix length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}

    // IPV6 framed prefix address. The type is string. This attribute is
    // mandatory.
    Prefix interface{}
}

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetEntityData() *types.CommonEntityData {
    framedPrefix.EntityData.YFilter = framedPrefix.YFilter
    framedPrefix.EntityData.YangName = "framed-prefix"
    framedPrefix.EntityData.BundleName = "cisco_ios_xr"
    framedPrefix.EntityData.ParentYangName = "ipv6-neighbor"
    framedPrefix.EntityData.SegmentPath = "framed-prefix"
    framedPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + framedPrefix.EntityData.SegmentPath
    framedPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    framedPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    framedPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    framedPrefix.EntityData.Children = types.NewOrderedMap()
    framedPrefix.EntityData.Leafs = types.NewOrderedMap()
    framedPrefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", framedPrefix.PrefixLength})
    framedPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", framedPrefix.Prefix})

    framedPrefix.EntityData.YListKeys = []string {}

    return &(framedPrefix.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection
// Duplicate Address Detection (DAD)
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set IPv6 duplicate address detection transmits. The type is interface{}
    // with range: 0..600.
    Attempts interface{}
}

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetEntityData() *types.CommonEntityData {
    duplicateAddressDetection.EntityData.YFilter = duplicateAddressDetection.YFilter
    duplicateAddressDetection.EntityData.YangName = "duplicate-address-detection"
    duplicateAddressDetection.EntityData.BundleName = "cisco_ios_xr"
    duplicateAddressDetection.EntityData.ParentYangName = "ipv6-neighbor"
    duplicateAddressDetection.EntityData.SegmentPath = "duplicate-address-detection"
    duplicateAddressDetection.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + duplicateAddressDetection.EntityData.SegmentPath
    duplicateAddressDetection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duplicateAddressDetection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duplicateAddressDetection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duplicateAddressDetection.EntityData.Children = types.NewOrderedMap()
    duplicateAddressDetection.EntityData.Leafs = types.NewOrderedMap()
    duplicateAddressDetection.EntityData.Leafs.Append("attempts", types.YLeaf{"Attempts", duplicateAddressDetection.Attempts})

    duplicateAddressDetection.EntityData.YListKeys = []string {}

    return &(duplicateAddressDetection.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial
// IPv6 ND RA Initial
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Initial RA count. The type is interface{} with range: 0..32. This attribute
    // is mandatory.
    Count interface{}

    // Initial RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Interval interface{}
}

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetEntityData() *types.CommonEntityData {
    raInitial.EntityData.YFilter = raInitial.YFilter
    raInitial.EntityData.YangName = "ra-initial"
    raInitial.EntityData.BundleName = "cisco_ios_xr"
    raInitial.EntityData.ParentYangName = "ipv6-neighbor"
    raInitial.EntityData.SegmentPath = "ra-initial"
    raInitial.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + raInitial.EntityData.SegmentPath
    raInitial.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    raInitial.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    raInitial.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    raInitial.EntityData.Children = types.NewOrderedMap()
    raInitial.EntityData.Leafs = types.NewOrderedMap()
    raInitial.EntityData.Leafs.Append("count", types.YLeaf{"Count", raInitial.Count})
    raInitial.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", raInitial.Interval})

    raInitial.EntityData.YListKeys = []string {}

    return &(raInitial.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Dhcpv6
// Interface dhcpv6 configuration data
type DynamicTemplate_Ppps_Ppp_Dhcpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dns IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DnsIpv6address interface{}

    // Select proxy/server profile based on mode class name. The type is string.
    ModeClass interface{}

    // Cisco VSA to configure any dhcpv6 ip lease per subscriber. The type is
    // string.
    Dhcpv6Iplease interface{}

    // Cisco VSA to configure any dhcpv6 option per subscriber. The type is
    // string.
    Dhcpv6Option interface{}

    // The pool to be used for Address assignment. The type is string.
    AddressPool interface{}

    // The pool to be used for Prefix Delegation. The type is string.
    DelegatedPrefixPool interface{}

    // The class to be used for proxy/server profile. The type is string.
    Class interface{}

    // Stateful IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    StatefulAddress interface{}

    // The prefix to be used for Prefix Delegation.
    DelegatedPrefix DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix
}

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetEntityData() *types.CommonEntityData {
    dhcpv6.EntityData.YFilter = dhcpv6.YFilter
    dhcpv6.EntityData.YangName = "dhcpv6"
    dhcpv6.EntityData.BundleName = "cisco_ios_xr"
    dhcpv6.EntityData.ParentYangName = "ppp"
    dhcpv6.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6"
    dhcpv6.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + dhcpv6.EntityData.SegmentPath
    dhcpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpv6.EntityData.Children = types.NewOrderedMap()
    dhcpv6.EntityData.Children.Append("delegated-prefix", types.YChild{"DelegatedPrefix", &dhcpv6.DelegatedPrefix})
    dhcpv6.EntityData.Leafs = types.NewOrderedMap()
    dhcpv6.EntityData.Leafs.Append("dns-ipv6address", types.YLeaf{"DnsIpv6address", dhcpv6.DnsIpv6address})
    dhcpv6.EntityData.Leafs.Append("mode-class", types.YLeaf{"ModeClass", dhcpv6.ModeClass})
    dhcpv6.EntityData.Leafs.Append("dhcpv6-iplease", types.YLeaf{"Dhcpv6Iplease", dhcpv6.Dhcpv6Iplease})
    dhcpv6.EntityData.Leafs.Append("dhcpv6-option", types.YLeaf{"Dhcpv6Option", dhcpv6.Dhcpv6Option})
    dhcpv6.EntityData.Leafs.Append("address-pool", types.YLeaf{"AddressPool", dhcpv6.AddressPool})
    dhcpv6.EntityData.Leafs.Append("delegated-prefix-pool", types.YLeaf{"DelegatedPrefixPool", dhcpv6.DelegatedPrefixPool})
    dhcpv6.EntityData.Leafs.Append("class", types.YLeaf{"Class", dhcpv6.Class})
    dhcpv6.EntityData.Leafs.Append("stateful-address", types.YLeaf{"StatefulAddress", dhcpv6.StatefulAddress})

    dhcpv6.EntityData.YListKeys = []string {}

    return &(dhcpv6.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix
// The prefix to be used for Prefix Delegation
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IPv6 Prefix. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Prefix interface{}

    // PD Prefix Length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}
}

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetEntityData() *types.CommonEntityData {
    delegatedPrefix.EntityData.YFilter = delegatedPrefix.YFilter
    delegatedPrefix.EntityData.YangName = "delegated-prefix"
    delegatedPrefix.EntityData.BundleName = "cisco_ios_xr"
    delegatedPrefix.EntityData.ParentYangName = "dhcpv6"
    delegatedPrefix.EntityData.SegmentPath = "delegated-prefix"
    delegatedPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6/" + delegatedPrefix.EntityData.SegmentPath
    delegatedPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delegatedPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delegatedPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delegatedPrefix.EntityData.Children = types.NewOrderedMap()
    delegatedPrefix.EntityData.Leafs = types.NewOrderedMap()
    delegatedPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", delegatedPrefix.Prefix})
    delegatedPrefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", delegatedPrefix.PrefixLength})

    delegatedPrefix.EntityData.YListKeys = []string {}

    return &(delegatedPrefix.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Pbr
// Dynamic Template PBR configuration
type DynamicTemplate_Ppps_Ppp_Pbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class for subscriber ingress policy. The type is string.
    ServicePolicyIn interface{}

    // Ingress service policy.
    ServicePolicies DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicies
}

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetEntityData() *types.CommonEntityData {
    pbr.EntityData.YFilter = pbr.YFilter
    pbr.EntityData.YangName = "pbr"
    pbr.EntityData.BundleName = "cisco_ios_xr"
    pbr.EntityData.ParentYangName = "ppp"
    pbr.EntityData.SegmentPath = "Cisco-IOS-XR-pbr-subscriber-cfg:pbr"
    pbr.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + pbr.EntityData.SegmentPath
    pbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbr.EntityData.Children = types.NewOrderedMap()
    pbr.EntityData.Children.Append("service-policies", types.YChild{"ServicePolicies", &pbr.ServicePolicies})
    pbr.EntityData.Leafs = types.NewOrderedMap()
    pbr.EntityData.Leafs.Append("service-policy-in", types.YLeaf{"ServicePolicyIn", pbr.ServicePolicyIn})

    pbr.EntityData.YListKeys = []string {}

    return &(pbr.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicies
// Ingress service policy
type DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service policy details. The type is slice of
    // DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicies_ServicePolicy.
    ServicePolicy []*DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicies_ServicePolicy
}

func (servicePolicies *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicies) GetEntityData() *types.CommonEntityData {
    servicePolicies.EntityData.YFilter = servicePolicies.YFilter
    servicePolicies.EntityData.YangName = "service-policies"
    servicePolicies.EntityData.BundleName = "cisco_ios_xr"
    servicePolicies.EntityData.ParentYangName = "pbr"
    servicePolicies.EntityData.SegmentPath = "service-policies"
    servicePolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-pbr-subscriber-cfg:pbr/" + servicePolicies.EntityData.SegmentPath
    servicePolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicies.EntityData.Children = types.NewOrderedMap()
    servicePolicies.EntityData.Children.Append("service-policy", types.YChild{"ServicePolicy", nil})
    for i := range servicePolicies.ServicePolicy {
        servicePolicies.EntityData.Children.Append(types.GetSegmentPath(servicePolicies.ServicePolicy[i]), types.YChild{"ServicePolicy", servicePolicies.ServicePolicy[i]})
    }
    servicePolicies.EntityData.Leafs = types.NewOrderedMap()

    servicePolicies.EntityData.YListKeys = []string {}

    return &(servicePolicies.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicies_ServicePolicy
// Service policy details
type DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicies_ServicePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of policy-map. The type is string with
    // length: 1..64.
    ServicePolicy interface{}
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicies_ServicePolicy) GetEntityData() *types.CommonEntityData {
    servicePolicy.EntityData.YFilter = servicePolicy.YFilter
    servicePolicy.EntityData.YangName = "service-policy"
    servicePolicy.EntityData.BundleName = "cisco_ios_xr"
    servicePolicy.EntityData.ParentYangName = "service-policies"
    servicePolicy.EntityData.SegmentPath = "service-policy" + types.AddKeyToken(servicePolicy.ServicePolicy, "service-policy")
    servicePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-pbr-subscriber-cfg:pbr/service-policies/" + servicePolicy.EntityData.SegmentPath
    servicePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicy.EntityData.Children = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs.Append("service-policy", types.YLeaf{"ServicePolicy", servicePolicy.ServicePolicy})

    servicePolicy.EntityData.YListKeys = []string {"ServicePolicy"}

    return &(servicePolicy.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate
// PPP template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPP FSM global template configuration data.
    Fsm DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm

    // PPP LCP global template configuration data.
    Lcp DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp

    // PPP IPv6CP global template configuration data.
    Ipv6cp DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6cp

    // PPP IPCP global template configuration data.
    Ipcp DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp
}

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetEntityData() *types.CommonEntityData {
    pppTemplate.EntityData.YFilter = pppTemplate.YFilter
    pppTemplate.EntityData.YangName = "ppp-template"
    pppTemplate.EntityData.BundleName = "cisco_ios_xr"
    pppTemplate.EntityData.ParentYangName = "ppp"
    pppTemplate.EntityData.SegmentPath = "Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template"
    pppTemplate.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + pppTemplate.EntityData.SegmentPath
    pppTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppTemplate.EntityData.Children = types.NewOrderedMap()
    pppTemplate.EntityData.Children.Append("fsm", types.YChild{"Fsm", &pppTemplate.Fsm})
    pppTemplate.EntityData.Children.Append("lcp", types.YChild{"Lcp", &pppTemplate.Lcp})
    pppTemplate.EntityData.Children.Append("ipv6cp", types.YChild{"Ipv6cp", &pppTemplate.Ipv6cp})
    pppTemplate.EntityData.Children.Append("ipcp", types.YChild{"Ipcp", &pppTemplate.Ipcp})
    pppTemplate.EntityData.Leafs = types.NewOrderedMap()

    pppTemplate.EntityData.YListKeys = []string {}

    return &(pppTemplate.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm
// PPP FSM global template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This specifies the maximum number of consecutive Conf-Naks. The type is
    // interface{} with range: 2..10. The default value is 5.
    MaxConsecutiveConfNaks interface{}

    // This specifies the maximum number of unacknowledged Conf-Requests. The type
    // is interface{} with range: 4..20. The default value is 10.
    MaxUnacknowledgedConfRequests interface{}

    // This specifies the maximum time to wait for a response during PPP
    // negotiation. The type is interface{} with range: 1..10. The default value
    // is 3.
    RetryTimeout interface{}

    // This specifies the maximum time to wait before sending Protocol Reject. The
    // type is interface{} with range: 1..60. The default value is 60.
    ProtocolRejectTimeout interface{}
}

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetEntityData() *types.CommonEntityData {
    fsm.EntityData.YFilter = fsm.YFilter
    fsm.EntityData.YangName = "fsm"
    fsm.EntityData.BundleName = "cisco_ios_xr"
    fsm.EntityData.ParentYangName = "ppp-template"
    fsm.EntityData.SegmentPath = "fsm"
    fsm.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/" + fsm.EntityData.SegmentPath
    fsm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fsm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fsm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fsm.EntityData.Children = types.NewOrderedMap()
    fsm.EntityData.Leafs = types.NewOrderedMap()
    fsm.EntityData.Leafs.Append("max-consecutive-conf-naks", types.YLeaf{"MaxConsecutiveConfNaks", fsm.MaxConsecutiveConfNaks})
    fsm.EntityData.Leafs.Append("max-unacknowledged-conf-requests", types.YLeaf{"MaxUnacknowledgedConfRequests", fsm.MaxUnacknowledgedConfRequests})
    fsm.EntityData.Leafs.Append("retry-timeout", types.YLeaf{"RetryTimeout", fsm.RetryTimeout})
    fsm.EntityData.Leafs.Append("protocol-reject-timeout", types.YLeaf{"ProtocolRejectTimeout", fsm.ProtocolRejectTimeout})

    fsm.EntityData.YListKeys = []string {}

    return &(fsm.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp
// PPP LCP global template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether to ignore attempts to renegotiate LCP. The type is
    // interface{}.
    Renegotiation interface{}

    // This is the Service-Type. The type is interface{} with range: 0..15. The
    // default value is 0.
    ServiceType interface{}

    // Enable Sending LCP Terminate request on shutdown. The type is interface{}.
    SendTermRequestOnShutDown interface{}

    // Ignore MRU negotiated with peer while setting interface BW. The type is
    // interface{}.
    MruIgnore interface{}

    // This specifies the session absolute timeout value.
    AbsoluteTimeout DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout

    // This specifies the time to delay before starting active LCPnegotiations.
    Delay DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay

    // PPP authentication parameters.
    Authentication DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication

    // This specifies the rate at which EchoReq packets are sent.
    Keepalive DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive
}

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetEntityData() *types.CommonEntityData {
    lcp.EntityData.YFilter = lcp.YFilter
    lcp.EntityData.YangName = "lcp"
    lcp.EntityData.BundleName = "cisco_ios_xr"
    lcp.EntityData.ParentYangName = "ppp-template"
    lcp.EntityData.SegmentPath = "lcp"
    lcp.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/" + lcp.EntityData.SegmentPath
    lcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcp.EntityData.Children = types.NewOrderedMap()
    lcp.EntityData.Children.Append("absolute-timeout", types.YChild{"AbsoluteTimeout", &lcp.AbsoluteTimeout})
    lcp.EntityData.Children.Append("delay", types.YChild{"Delay", &lcp.Delay})
    lcp.EntityData.Children.Append("authentication", types.YChild{"Authentication", &lcp.Authentication})
    lcp.EntityData.Children.Append("keepalive", types.YChild{"Keepalive", &lcp.Keepalive})
    lcp.EntityData.Leafs = types.NewOrderedMap()
    lcp.EntityData.Leafs.Append("renegotiation", types.YLeaf{"Renegotiation", lcp.Renegotiation})
    lcp.EntityData.Leafs.Append("service-type", types.YLeaf{"ServiceType", lcp.ServiceType})
    lcp.EntityData.Leafs.Append("send-term-request-on-shut-down", types.YLeaf{"SendTermRequestOnShutDown", lcp.SendTermRequestOnShutDown})
    lcp.EntityData.Leafs.Append("mru-ignore", types.YLeaf{"MruIgnore", lcp.MruIgnore})

    lcp.EntityData.YListKeys = []string {}

    return &(lcp.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout
// This specifies the session absolute timeout
// value
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minutes. The type is interface{} with range: 0..35000000.
    Minutes interface{}

    // Seconds. The type is interface{} with range: 0..59.
    Seconds interface{}
}

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetEntityData() *types.CommonEntityData {
    absoluteTimeout.EntityData.YFilter = absoluteTimeout.YFilter
    absoluteTimeout.EntityData.YangName = "absolute-timeout"
    absoluteTimeout.EntityData.BundleName = "cisco_ios_xr"
    absoluteTimeout.EntityData.ParentYangName = "lcp"
    absoluteTimeout.EntityData.SegmentPath = "absolute-timeout"
    absoluteTimeout.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/lcp/" + absoluteTimeout.EntityData.SegmentPath
    absoluteTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    absoluteTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    absoluteTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    absoluteTimeout.EntityData.Children = types.NewOrderedMap()
    absoluteTimeout.EntityData.Leafs = types.NewOrderedMap()
    absoluteTimeout.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", absoluteTimeout.Minutes})
    absoluteTimeout.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", absoluteTimeout.Seconds})

    absoluteTimeout.EntityData.YListKeys = []string {}

    return &(absoluteTimeout.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay
// This specifies the time to delay before
// starting active LCPnegotiations
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..255.
    Seconds interface{}

    // Milliseconds. The type is interface{} with range: 0..999.
    Milliseconds interface{}
}

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetEntityData() *types.CommonEntityData {
    delay.EntityData.YFilter = delay.YFilter
    delay.EntityData.YangName = "delay"
    delay.EntityData.BundleName = "cisco_ios_xr"
    delay.EntityData.ParentYangName = "lcp"
    delay.EntityData.SegmentPath = "delay"
    delay.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/lcp/" + delay.EntityData.SegmentPath
    delay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delay.EntityData.Children = types.NewOrderedMap()
    delay.EntityData.Leafs = types.NewOrderedMap()
    delay.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", delay.Seconds})
    delay.EntityData.Leafs.Append("milliseconds", types.YLeaf{"Milliseconds", delay.Milliseconds})

    delay.EntityData.YListKeys = []string {}

    return &(delay.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication
// PPP authentication parameters
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This specifies the CHAP hostname. The type is string.
    ChapHostName interface{}

    // <1> for accepting null-passwordduring authentication. The type is
    // interface{} with range: 0..4294967295.
    Pap interface{}

    // This specifies the MS-CHAP hostname. The type is string.
    MschapHostName interface{}

    // This specifies whether to allow multiple authentication failures and, if
    // so, how many. The type is interface{} with range: 0..10.
    MaxAuthenticationFailures interface{}

    // Maximum time to wait for an authentication response. The type is
    // interface{} with range: 3..30. The default value is 10.
    Timeout interface{}

    // This specifies the PPP link authentication method.
    Methods DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods
}

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "lcp"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/lcp/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Children.Append("methods", types.YChild{"Methods", &authentication.Methods})
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("chap-host-name", types.YLeaf{"ChapHostName", authentication.ChapHostName})
    authentication.EntityData.Leafs.Append("pap", types.YLeaf{"Pap", authentication.Pap})
    authentication.EntityData.Leafs.Append("mschap-host-name", types.YLeaf{"MschapHostName", authentication.MschapHostName})
    authentication.EntityData.Leafs.Append("max-authentication-failures", types.YLeaf{"MaxAuthenticationFailures", authentication.MaxAuthenticationFailures})
    authentication.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", authentication.Timeout})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods
// This specifies the PPP link authentication
// method
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select between one and three authentication methods in order of preference.
    // The type is slice of PppAuthenticationMethodGbl.
    Method []interface{}
}

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetEntityData() *types.CommonEntityData {
    methods.EntityData.YFilter = methods.YFilter
    methods.EntityData.YangName = "methods"
    methods.EntityData.BundleName = "cisco_ios_xr"
    methods.EntityData.ParentYangName = "authentication"
    methods.EntityData.SegmentPath = "methods"
    methods.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/lcp/authentication/" + methods.EntityData.SegmentPath
    methods.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    methods.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    methods.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    methods.EntityData.Children = types.NewOrderedMap()
    methods.EntityData.Leafs = types.NewOrderedMap()
    methods.EntityData.Leafs.Append("method", types.YLeaf{"Method", methods.Method})

    methods.EntityData.YListKeys = []string {}

    return &(methods.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive
// This specifies the rate at which EchoReq
// packets are sent
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE to disable keepalives, FALSE to specify a new keepalive interval. The
    // type is bool.
    KeepaliveDisable interface{}

    // The keepalive interval. Leave unspecified when disabling keepalives. The
    // type is interface{} with range: 10..180.
    Interval interface{}

    // The keepalive retry count. Leave unspecified when disabling keepalives. The
    // type is interface{} with range: 1..255.
    RetryCount interface{}
}

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetEntityData() *types.CommonEntityData {
    keepalive.EntityData.YFilter = keepalive.YFilter
    keepalive.EntityData.YangName = "keepalive"
    keepalive.EntityData.BundleName = "cisco_ios_xr"
    keepalive.EntityData.ParentYangName = "lcp"
    keepalive.EntityData.SegmentPath = "keepalive"
    keepalive.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/lcp/" + keepalive.EntityData.SegmentPath
    keepalive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keepalive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keepalive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keepalive.EntityData.Children = types.NewOrderedMap()
    keepalive.EntityData.Leafs = types.NewOrderedMap()
    keepalive.EntityData.Leafs.Append("keepalive-disable", types.YLeaf{"KeepaliveDisable", keepalive.KeepaliveDisable})
    keepalive.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", keepalive.Interval})
    keepalive.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", keepalive.RetryCount})

    keepalive.EntityData.YListKeys = []string {}

    return &(keepalive.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6cp
// PPP IPv6CP global template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6cp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether to run IPv6CP in Passive mode. The type is interface{}.
    Passive interface{}

    // Specify whether to ignore attempts to renegotiate IPv6CP. The type is
    // interface{}.
    Renegotiation interface{}

    // Specify the Interface-Id to impose on the peer. The type is string.
    PeerInterfaceId interface{}

    // Specify whether to protocol reject IPv6CP. The type is interface{}.
    ProtocolReject interface{}
}

func (ipv6cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6cp) GetEntityData() *types.CommonEntityData {
    ipv6cp.EntityData.YFilter = ipv6cp.YFilter
    ipv6cp.EntityData.YangName = "ipv6cp"
    ipv6cp.EntityData.BundleName = "cisco_ios_xr"
    ipv6cp.EntityData.ParentYangName = "ppp-template"
    ipv6cp.EntityData.SegmentPath = "ipv6cp"
    ipv6cp.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/" + ipv6cp.EntityData.SegmentPath
    ipv6cp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6cp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6cp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6cp.EntityData.Children = types.NewOrderedMap()
    ipv6cp.EntityData.Leafs = types.NewOrderedMap()
    ipv6cp.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", ipv6cp.Passive})
    ipv6cp.EntityData.Leafs.Append("renegotiation", types.YLeaf{"Renegotiation", ipv6cp.Renegotiation})
    ipv6cp.EntityData.Leafs.Append("peer-interface-id", types.YLeaf{"PeerInterfaceId", ipv6cp.PeerInterfaceId})
    ipv6cp.EntityData.Leafs.Append("protocol-reject", types.YLeaf{"ProtocolReject", ipv6cp.ProtocolReject})

    ipv6cp.EntityData.YListKeys = []string {}

    return &(ipv6cp.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp
// PPP IPCP global template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether to ignore attempts to renegotiate IPCP. The type is
    // interface{}.
    Renegotiation interface{}

    // Specify whether to run IPCP in Passive mode. The type is interface{}.
    Passive interface{}

    // Specify whether to protocol reject IPCP. The type is interface{}.
    ProtocolReject interface{}

    // Specify the IPv4 netmask to negotiate for the peer. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerNetmask interface{}

    // IPCP WINS parameters.
    Wins DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins

    // IPCP DNS parameters.
    Dns DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns

    // IPCP address parameters.
    PeerAddress DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress
}

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetEntityData() *types.CommonEntityData {
    ipcp.EntityData.YFilter = ipcp.YFilter
    ipcp.EntityData.YangName = "ipcp"
    ipcp.EntityData.BundleName = "cisco_ios_xr"
    ipcp.EntityData.ParentYangName = "ppp-template"
    ipcp.EntityData.SegmentPath = "ipcp"
    ipcp.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/" + ipcp.EntityData.SegmentPath
    ipcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipcp.EntityData.Children = types.NewOrderedMap()
    ipcp.EntityData.Children.Append("wins", types.YChild{"Wins", &ipcp.Wins})
    ipcp.EntityData.Children.Append("dns", types.YChild{"Dns", &ipcp.Dns})
    ipcp.EntityData.Children.Append("peer-address", types.YChild{"PeerAddress", &ipcp.PeerAddress})
    ipcp.EntityData.Leafs = types.NewOrderedMap()
    ipcp.EntityData.Leafs.Append("renegotiation", types.YLeaf{"Renegotiation", ipcp.Renegotiation})
    ipcp.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", ipcp.Passive})
    ipcp.EntityData.Leafs.Append("protocol-reject", types.YLeaf{"ProtocolReject", ipcp.ProtocolReject})
    ipcp.EntityData.Leafs.Append("peer-netmask", types.YLeaf{"PeerNetmask", ipcp.PeerNetmask})

    ipcp.EntityData.YListKeys = []string {}

    return &(ipcp.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins
// IPCP WINS parameters
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify WINS address(es) to provide.
    WinsAddresses DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses
}

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetEntityData() *types.CommonEntityData {
    wins.EntityData.YFilter = wins.YFilter
    wins.EntityData.YangName = "wins"
    wins.EntityData.BundleName = "cisco_ios_xr"
    wins.EntityData.ParentYangName = "ipcp"
    wins.EntityData.SegmentPath = "wins"
    wins.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/ipcp/" + wins.EntityData.SegmentPath
    wins.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wins.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wins.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wins.EntityData.Children = types.NewOrderedMap()
    wins.EntityData.Children.Append("wins-addresses", types.YChild{"WinsAddresses", &wins.WinsAddresses})
    wins.EntityData.Leafs = types.NewOrderedMap()

    wins.EntityData.YListKeys = []string {}

    return &(wins.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses
// Specify WINS address(es) to provide
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Primary WINS IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Primary interface{}

    // Secondary WINS IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Secondary interface{}
}

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetEntityData() *types.CommonEntityData {
    winsAddresses.EntityData.YFilter = winsAddresses.YFilter
    winsAddresses.EntityData.YangName = "wins-addresses"
    winsAddresses.EntityData.BundleName = "cisco_ios_xr"
    winsAddresses.EntityData.ParentYangName = "wins"
    winsAddresses.EntityData.SegmentPath = "wins-addresses"
    winsAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/ipcp/wins/" + winsAddresses.EntityData.SegmentPath
    winsAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    winsAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    winsAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    winsAddresses.EntityData.Children = types.NewOrderedMap()
    winsAddresses.EntityData.Leafs = types.NewOrderedMap()
    winsAddresses.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", winsAddresses.Primary})
    winsAddresses.EntityData.Leafs.Append("secondary", types.YLeaf{"Secondary", winsAddresses.Secondary})

    winsAddresses.EntityData.YListKeys = []string {}

    return &(winsAddresses.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns
// IPCP DNS parameters
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify DNS address(es) to provide.
    DnsAddresses DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses
}

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetEntityData() *types.CommonEntityData {
    dns.EntityData.YFilter = dns.YFilter
    dns.EntityData.YangName = "dns"
    dns.EntityData.BundleName = "cisco_ios_xr"
    dns.EntityData.ParentYangName = "ipcp"
    dns.EntityData.SegmentPath = "dns"
    dns.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/ipcp/" + dns.EntityData.SegmentPath
    dns.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dns.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dns.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dns.EntityData.Children = types.NewOrderedMap()
    dns.EntityData.Children.Append("dns-addresses", types.YChild{"DnsAddresses", &dns.DnsAddresses})
    dns.EntityData.Leafs = types.NewOrderedMap()

    dns.EntityData.YListKeys = []string {}

    return &(dns.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses
// Specify DNS address(es) to provide
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Primary DNS IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Primary interface{}

    // Secondary DNS IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Secondary interface{}
}

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetEntityData() *types.CommonEntityData {
    dnsAddresses.EntityData.YFilter = dnsAddresses.YFilter
    dnsAddresses.EntityData.YangName = "dns-addresses"
    dnsAddresses.EntityData.BundleName = "cisco_ios_xr"
    dnsAddresses.EntityData.ParentYangName = "dns"
    dnsAddresses.EntityData.SegmentPath = "dns-addresses"
    dnsAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/ipcp/dns/" + dnsAddresses.EntityData.SegmentPath
    dnsAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnsAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnsAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnsAddresses.EntityData.Children = types.NewOrderedMap()
    dnsAddresses.EntityData.Leafs = types.NewOrderedMap()
    dnsAddresses.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", dnsAddresses.Primary})
    dnsAddresses.EntityData.Leafs.Append("secondary", types.YLeaf{"Secondary", dnsAddresses.Secondary})

    dnsAddresses.EntityData.YListKeys = []string {}

    return &(dnsAddresses.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress
// IPCP address parameters
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify an IP address to assign to peers through IPCP. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Default interface{}

    // Accepts an IP address from the peer if in the pool, else allocates one from
    // the pool. The type is string.
    Pool interface{}
}

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetEntityData() *types.CommonEntityData {
    peerAddress.EntityData.YFilter = peerAddress.YFilter
    peerAddress.EntityData.YangName = "peer-address"
    peerAddress.EntityData.BundleName = "cisco_ios_xr"
    peerAddress.EntityData.ParentYangName = "ipcp"
    peerAddress.EntityData.SegmentPath = "peer-address"
    peerAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template/ipcp/" + peerAddress.EntityData.SegmentPath
    peerAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAddress.EntityData.Children = types.NewOrderedMap()
    peerAddress.EntityData.Leafs = types.NewOrderedMap()
    peerAddress.EntityData.Leafs.Append("default", types.YLeaf{"Default", peerAddress.Default})
    peerAddress.EntityData.Leafs.Append("pool", types.YLeaf{"Pool", peerAddress.Pool})

    peerAddress.EntityData.YListKeys = []string {}

    return &(peerAddress.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Qos
// QoS dynamically applied configuration template
type DynamicTemplate_Ppps_Ppp_Qos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service policy to be applied in ingress/egress direction.
    ServicePolicy DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy

    // QoS L2 overhead accounting.
    Account DynamicTemplate_Ppps_Ppp_Qos_Account

    // QoS to be applied in egress direction.
    Output DynamicTemplate_Ppps_Ppp_Qos_Output
}

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetEntityData() *types.CommonEntityData {
    qos.EntityData.YFilter = qos.YFilter
    qos.EntityData.YangName = "qos"
    qos.EntityData.BundleName = "cisco_ios_xr"
    qos.EntityData.ParentYangName = "ppp"
    qos.EntityData.SegmentPath = "Cisco-IOS-XR-qos-ma-bng-cfg:qos"
    qos.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + qos.EntityData.SegmentPath
    qos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qos.EntityData.Children = types.NewOrderedMap()
    qos.EntityData.Children.Append("service-policy", types.YChild{"ServicePolicy", &qos.ServicePolicy})
    qos.EntityData.Children.Append("account", types.YChild{"Account", &qos.Account})
    qos.EntityData.Children.Append("output", types.YChild{"Output", &qos.Output})
    qos.EntityData.Leafs = types.NewOrderedMap()

    qos.EntityData.YListKeys = []string {}

    return &(qos.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy
// Service policy to be applied in ingress/egress
// direction
type DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber ingress policy.
    Input DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input

    // Subscriber egress policy.
    Output DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetEntityData() *types.CommonEntityData {
    servicePolicy.EntityData.YFilter = servicePolicy.YFilter
    servicePolicy.EntityData.YangName = "service-policy"
    servicePolicy.EntityData.BundleName = "cisco_ios_xr"
    servicePolicy.EntityData.ParentYangName = "qos"
    servicePolicy.EntityData.SegmentPath = "service-policy"
    servicePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-qos-ma-bng-cfg:qos/" + servicePolicy.EntityData.SegmentPath
    servicePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicy.EntityData.Children = types.NewOrderedMap()
    servicePolicy.EntityData.Children.Append("input", types.YChild{"Input", &servicePolicy.Input})
    servicePolicy.EntityData.Children.Append("output", types.YChild{"Output", &servicePolicy.Output})
    servicePolicy.EntityData.Leafs = types.NewOrderedMap()

    servicePolicy.EntityData.YListKeys = []string {}

    return &(servicePolicy.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input
// Subscriber ingress policy
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Name of policy-map. The type is string. This attribute is mandatory.
    PolicyName interface{}

    // Name of the SPI. The type is string.
    SpiName interface{}

    // TRUE for merge enabled for service-policy applied on dynamic template. The
    // type is bool.
    Merge interface{}

    // Merge ID value. The type is interface{} with range: 0..255.
    MergeId interface{}

    // TRUE for account stats enabled for service-policy applied on dynamic
    // template. Note: account stats not supported for subscriber type 'ppp' and
    // 'ipsubscriber'. The type is bool.
    AccountStats interface{}
}

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "service-policy"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-qos-ma-bng-cfg:qos/service-policy/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", input.PolicyName})
    input.EntityData.Leafs.Append("spi-name", types.YLeaf{"SpiName", input.SpiName})
    input.EntityData.Leafs.Append("merge", types.YLeaf{"Merge", input.Merge})
    input.EntityData.Leafs.Append("merge-id", types.YLeaf{"MergeId", input.MergeId})
    input.EntityData.Leafs.Append("account-stats", types.YLeaf{"AccountStats", input.AccountStats})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output
// Subscriber egress policy
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Name of policy-map. The type is string. This attribute is mandatory.
    PolicyName interface{}

    // Name of the SPI. The type is string.
    SpiName interface{}

    // TRUE for merge enabled for service-policy applied on dynamic template. The
    // type is bool.
    Merge interface{}

    // Merge ID value. The type is interface{} with range: 0..255.
    MergeId interface{}

    // TRUE for account stats enabled for service-policy applied on dynamic
    // template. Note: account stats not supported for subscriber type 'ppp' and
    // 'ipsubscriber'. The type is bool.
    AccountStats interface{}
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "service-policy"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-qos-ma-bng-cfg:qos/service-policy/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", output.PolicyName})
    output.EntityData.Leafs.Append("spi-name", types.YLeaf{"SpiName", output.SpiName})
    output.EntityData.Leafs.Append("merge", types.YLeaf{"Merge", output.Merge})
    output.EntityData.Leafs.Append("merge-id", types.YLeaf{"MergeId", output.MergeId})
    output.EntityData.Leafs.Append("account-stats", types.YLeaf{"AccountStats", output.AccountStats})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Qos_Account
// QoS L2 overhead accounting
type DynamicTemplate_Ppps_Ppp_Qos_Account struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ATM adaptation layer AAL. The type is Qosl2DataLink.
    Aal interface{}

    // Specify encapsulation type. The type is Qosl2Encap.
    Encapsulation interface{}

    // ATM cell tax to L2 overhead. The type is interface{}.
    AtmCellTax interface{}

    // Numeric L2 overhead offset. The type is interface{} with range: -63..63.
    UserDefined interface{}
}

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetEntityData() *types.CommonEntityData {
    account.EntityData.YFilter = account.YFilter
    account.EntityData.YangName = "account"
    account.EntityData.BundleName = "cisco_ios_xr"
    account.EntityData.ParentYangName = "qos"
    account.EntityData.SegmentPath = "account"
    account.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-qos-ma-bng-cfg:qos/" + account.EntityData.SegmentPath
    account.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    account.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    account.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    account.EntityData.Children = types.NewOrderedMap()
    account.EntityData.Leafs = types.NewOrderedMap()
    account.EntityData.Leafs.Append("aal", types.YLeaf{"Aal", account.Aal})
    account.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", account.Encapsulation})
    account.EntityData.Leafs.Append("atm-cell-tax", types.YLeaf{"AtmCellTax", account.AtmCellTax})
    account.EntityData.Leafs.Append("user-defined", types.YLeaf{"UserDefined", account.UserDefined})

    account.EntityData.YListKeys = []string {}

    return &(account.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Qos_Output
// QoS to be applied in egress direction
type DynamicTemplate_Ppps_Ppp_Qos_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth value for the subscriber (in kbps). The type is
    // interface{} with range: 1..4294967295. Units are kbit/s.
    MinimumBandwidth interface{}
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "qos"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-qos-ma-bng-cfg:qos/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("minimum-bandwidth", types.YLeaf{"MinimumBandwidth", output.MinimumBandwidth})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Accounting
// Subscriber accounting dynamic-template commands
type DynamicTemplate_Ppps_Ppp_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber monitor feature. The type is string.
    MonitorFeature interface{}

    // Subscriber accounting prepaid feature. The type is string.
    PrepaidFeature interface{}

    // Subscriber accounting idle timeout.
    IdleTimeout DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout

    // Subscriber accounting session accounting.
    Session DynamicTemplate_Ppps_Ppp_Accounting_Session

    // Subscriber accounting service accounting.
    ServiceAccounting DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting
}

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "ppp"
    accounting.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-accounting-cfg:accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("idle-timeout", types.YChild{"IdleTimeout", &accounting.IdleTimeout})
    accounting.EntityData.Children.Append("session", types.YChild{"Session", &accounting.Session})
    accounting.EntityData.Children.Append("service-accounting", types.YChild{"ServiceAccounting", &accounting.ServiceAccounting})
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("monitor-feature", types.YLeaf{"MonitorFeature", accounting.MonitorFeature})
    accounting.EntityData.Leafs.Append("prepaid-feature", types.YLeaf{"PrepaidFeature", accounting.PrepaidFeature})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout
// Subscriber accounting idle timeout
type DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Idle timeout value in seconds. The type is interface{} with range:
    // 60..4320000.
    TimeoutValue interface{}

    // Threshold in minute(s) per packet. The type is interface{} with range:
    // 1..10000.
    Threshold interface{}

    // Idle timeout traffic direction. The type is string.
    Direction interface{}
}

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetEntityData() *types.CommonEntityData {
    idleTimeout.EntityData.YFilter = idleTimeout.YFilter
    idleTimeout.EntityData.YangName = "idle-timeout"
    idleTimeout.EntityData.BundleName = "cisco_ios_xr"
    idleTimeout.EntityData.ParentYangName = "accounting"
    idleTimeout.EntityData.SegmentPath = "idle-timeout"
    idleTimeout.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-subscriber-accounting-cfg:accounting/" + idleTimeout.EntityData.SegmentPath
    idleTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idleTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idleTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idleTimeout.EntityData.Children = types.NewOrderedMap()
    idleTimeout.EntityData.Leafs = types.NewOrderedMap()
    idleTimeout.EntityData.Leafs.Append("timeout-value", types.YLeaf{"TimeoutValue", idleTimeout.TimeoutValue})
    idleTimeout.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", idleTimeout.Threshold})
    idleTimeout.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", idleTimeout.Direction})

    idleTimeout.EntityData.YListKeys = []string {}

    return &(idleTimeout.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Accounting_Session
// Subscriber accounting session accounting
type DynamicTemplate_Ppps_Ppp_Accounting_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session accounting method list name. The type is string.
    MethodListName interface{}

    // Interim accounting interval in minutes. The type is interface{} with range:
    // 0..4294967295.
    PeriodicInterval interface{}

    // Dual stack wait delay in seconds. The type is interface{} with range:
    // 0..4294967295.
    DualStackDelay interface{}

    // Hold Accounting start based on IA_PD. The type is HoldAcctStart.
    HoldAcctStart interface{}
}

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "accounting"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-subscriber-accounting-cfg:accounting/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("method-list-name", types.YLeaf{"MethodListName", session.MethodListName})
    session.EntityData.Leafs.Append("periodic-interval", types.YLeaf{"PeriodicInterval", session.PeriodicInterval})
    session.EntityData.Leafs.Append("dual-stack-delay", types.YLeaf{"DualStackDelay", session.DualStackDelay})
    session.EntityData.Leafs.Append("hold-acct-start", types.YLeaf{"HoldAcctStart", session.HoldAcctStart})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// DynamicTemplate_Ppps_Ppp_Accounting_Session_HoldAcctStart represents Hold Accounting start based on IA_PD
type DynamicTemplate_Ppps_Ppp_Accounting_Session_HoldAcctStart string

const (
    // Based on ipv6 delegated prefix
    DynamicTemplate_Ppps_Ppp_Accounting_Session_HoldAcctStart_ipv6_prefix_delegation DynamicTemplate_Ppps_Ppp_Accounting_Session_HoldAcctStart = "ipv6-prefix-delegation"
)

// DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting
// Subscriber accounting service accounting
type DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service accounting method list name. The type is string.
    MethodListName interface{}

    // Accounting interim interval in minutes. The type is interface{} with range:
    // 0..4294967295.
    AccountingInterimInterval interface{}
}

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetEntityData() *types.CommonEntityData {
    serviceAccounting.EntityData.YFilter = serviceAccounting.YFilter
    serviceAccounting.EntityData.YangName = "service-accounting"
    serviceAccounting.EntityData.BundleName = "cisco_ios_xr"
    serviceAccounting.EntityData.ParentYangName = "accounting"
    serviceAccounting.EntityData.SegmentPath = "service-accounting"
    serviceAccounting.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/Cisco-IOS-XR-subscriber-accounting-cfg:accounting/" + serviceAccounting.EntityData.SegmentPath
    serviceAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceAccounting.EntityData.Children = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs.Append("method-list-name", types.YLeaf{"MethodListName", serviceAccounting.MethodListName})
    serviceAccounting.EntityData.Leafs.Append("accounting-interim-interval", types.YLeaf{"AccountingInterimInterval", serviceAccounting.AccountingInterimInterval})

    serviceAccounting.EntityData.YListKeys = []string {}

    return &(serviceAccounting.EntityData)
}

// DynamicTemplate_Ppps_Ppp_PppoeTemplate
// PPPoE template configuration data
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_PppoeTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Specify the Port limit (attr 62) to apply to the subscriber. The type is
    // interface{} with range: 1..65535. This attribute is mandatory.
    PortLimit interface{}
}

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetEntityData() *types.CommonEntityData {
    pppoeTemplate.EntityData.YFilter = pppoeTemplate.YFilter
    pppoeTemplate.EntityData.YangName = "pppoe-template"
    pppoeTemplate.EntityData.BundleName = "cisco_ios_xr"
    pppoeTemplate.EntityData.ParentYangName = "ppp"
    pppoeTemplate.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-template"
    pppoeTemplate.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ppps/ppp/" + pppoeTemplate.EntityData.SegmentPath
    pppoeTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoeTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoeTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoeTemplate.EntityData.Children = types.NewOrderedMap()
    pppoeTemplate.EntityData.Leafs = types.NewOrderedMap()
    pppoeTemplate.EntityData.Leafs.Append("port-limit", types.YLeaf{"PortLimit", pppoeTemplate.PortLimit})

    pppoeTemplate.EntityData.YListKeys = []string {}

    return &(pppoeTemplate.EntityData)
}

// DynamicTemplate_IpSubscribers
// The IP Subscriber Template Table
type DynamicTemplate_IpSubscribers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A IP Subscriber Type Template . The type is slice of
    // DynamicTemplate_IpSubscribers_IpSubscriber.
    IpSubscriber []*DynamicTemplate_IpSubscribers_IpSubscriber
}

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetEntityData() *types.CommonEntityData {
    ipSubscribers.EntityData.YFilter = ipSubscribers.YFilter
    ipSubscribers.EntityData.YangName = "ip-subscribers"
    ipSubscribers.EntityData.BundleName = "cisco_ios_xr"
    ipSubscribers.EntityData.ParentYangName = "dynamic-template"
    ipSubscribers.EntityData.SegmentPath = "ip-subscribers"
    ipSubscribers.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/" + ipSubscribers.EntityData.SegmentPath
    ipSubscribers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscribers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscribers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscribers.EntityData.Children = types.NewOrderedMap()
    ipSubscribers.EntityData.Children.Append("ip-subscriber", types.YChild{"IpSubscriber", nil})
    for i := range ipSubscribers.IpSubscriber {
        ipSubscribers.EntityData.Children.Append(types.GetSegmentPath(ipSubscribers.IpSubscriber[i]), types.YChild{"IpSubscriber", ipSubscribers.IpSubscriber[i]})
    }
    ipSubscribers.EntityData.Leafs = types.NewOrderedMap()

    ipSubscribers.EntityData.YListKeys = []string {}

    return &(ipSubscribers.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber
// A IP Subscriber Type Template 
type DynamicTemplate_IpSubscribers_IpSubscriber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of the template. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TemplateName interface{}

    // Assign the interface to a VRF . The type is string with length: 1..32.
    Vrf interface{}

    // Monitor Session container for this template.
    SpanMonitorSessions DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions

    // IPv4 Packet Filtering configuration for the template.
    Ipv4PacketFilter DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter

    // IPv6 Packet Filtering configuration for the interface.
    Ipv6PacketFilter DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter

    // Interface dhcpv4 configuration data.
    Dhcpd DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd

    // IGMPconfiguration.
    Igmp DynamicTemplate_IpSubscribers_IpSubscriber_Igmp

    // Interface IPv4 Network configuration data.
    Ipv4Network DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network

    // Interface IPv6 Network configuration data.
    Ipv6Network DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network

    // Interface IPv6 Network configuration data.
    Ipv6Neighbor DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor

    // Interface dhcpv6 configuration data.
    Dhcpv6 DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6

    // Dynamic Template PBR configuration.
    Pbr DynamicTemplate_IpSubscribers_IpSubscriber_Pbr

    // QoS dynamically applied configuration template.
    Qos DynamicTemplate_IpSubscribers_IpSubscriber_Qos

    // Subscriber accounting dynamic-template commands.
    Accounting DynamicTemplate_IpSubscribers_IpSubscriber_Accounting
}

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetEntityData() *types.CommonEntityData {
    ipSubscriber.EntityData.YFilter = ipSubscriber.YFilter
    ipSubscriber.EntityData.YangName = "ip-subscriber"
    ipSubscriber.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriber.EntityData.ParentYangName = "ip-subscribers"
    ipSubscriber.EntityData.SegmentPath = "ip-subscriber" + types.AddKeyToken(ipSubscriber.TemplateName, "template-name")
    ipSubscriber.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/" + ipSubscriber.EntityData.SegmentPath
    ipSubscriber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriber.EntityData.Children = types.NewOrderedMap()
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions", types.YChild{"SpanMonitorSessions", &ipSubscriber.SpanMonitorSessions})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter", types.YChild{"Ipv4PacketFilter", &ipSubscriber.Ipv4PacketFilter})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter", types.YChild{"Ipv6PacketFilter", &ipSubscriber.Ipv6PacketFilter})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-ipv4-dhcpd-subscriber-cfg:dhcpd", types.YChild{"Dhcpd", &ipSubscriber.Dhcpd})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp", types.YChild{"Igmp", &ipSubscriber.Igmp})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network", types.YChild{"Ipv4Network", &ipSubscriber.Ipv4Network})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network", types.YChild{"Ipv6Network", &ipSubscriber.Ipv6Network})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor", types.YChild{"Ipv6Neighbor", &ipSubscriber.Ipv6Neighbor})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6", types.YChild{"Dhcpv6", &ipSubscriber.Dhcpv6})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-pbr-subscriber-cfg:pbr", types.YChild{"Pbr", &ipSubscriber.Pbr})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-qos-ma-bng-cfg:qos", types.YChild{"Qos", &ipSubscriber.Qos})
    ipSubscriber.EntityData.Children.Append("Cisco-IOS-XR-subscriber-accounting-cfg:accounting", types.YChild{"Accounting", &ipSubscriber.Accounting})
    ipSubscriber.EntityData.Leafs = types.NewOrderedMap()
    ipSubscriber.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", ipSubscriber.TemplateName})
    ipSubscriber.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", ipSubscriber.Vrf})

    ipSubscriber.EntityData.YListKeys = []string {"TemplateName"}

    return &(ipSubscriber.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions
// Monitor Session container for this template
type DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular class of Monitor Session. The type is slice
    // of
    // DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession.
    SpanMonitorSession []*DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession
}

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetEntityData() *types.CommonEntityData {
    spanMonitorSessions.EntityData.YFilter = spanMonitorSessions.YFilter
    spanMonitorSessions.EntityData.YangName = "span-monitor-sessions"
    spanMonitorSessions.EntityData.BundleName = "cisco_ios_xr"
    spanMonitorSessions.EntityData.ParentYangName = "ip-subscriber"
    spanMonitorSessions.EntityData.SegmentPath = "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions"
    spanMonitorSessions.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + spanMonitorSessions.EntityData.SegmentPath
    spanMonitorSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMonitorSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMonitorSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMonitorSessions.EntityData.Children = types.NewOrderedMap()
    spanMonitorSessions.EntityData.Children.Append("span-monitor-session", types.YChild{"SpanMonitorSession", nil})
    for i := range spanMonitorSessions.SpanMonitorSession {
        spanMonitorSessions.EntityData.Children.Append(types.GetSegmentPath(spanMonitorSessions.SpanMonitorSession[i]), types.YChild{"SpanMonitorSession", spanMonitorSessions.SpanMonitorSession[i]})
    }
    spanMonitorSessions.EntityData.Leafs = types.NewOrderedMap()

    spanMonitorSessions.EntityData.YListKeys = []string {}

    return &(spanMonitorSessions.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession
// Configuration for a particular class of Monitor
// Session
type DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Session Class. The type is SpanSessionClass.
    SessionClass interface{}

    // Mirror a specified number of bytes from start of packet. The type is
    // interface{} with range: 1..10000. Units are byte.
    MirrorFirst interface{}

    // Specify the mirror interval. The type is SpanMirrorInterval.
    MirrorInterval interface{}

    // Attach the interface to a Monitor Session.
    Attachment DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment

    // Enable ACL matching for traffic mirroring.
    Acl DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl
}

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetEntityData() *types.CommonEntityData {
    spanMonitorSession.EntityData.YFilter = spanMonitorSession.YFilter
    spanMonitorSession.EntityData.YangName = "span-monitor-session"
    spanMonitorSession.EntityData.BundleName = "cisco_ios_xr"
    spanMonitorSession.EntityData.ParentYangName = "span-monitor-sessions"
    spanMonitorSession.EntityData.SegmentPath = "span-monitor-session" + types.AddKeyToken(spanMonitorSession.SessionClass, "session-class")
    spanMonitorSession.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions/" + spanMonitorSession.EntityData.SegmentPath
    spanMonitorSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMonitorSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMonitorSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMonitorSession.EntityData.Children = types.NewOrderedMap()
    spanMonitorSession.EntityData.Children.Append("attachment", types.YChild{"Attachment", &spanMonitorSession.Attachment})
    spanMonitorSession.EntityData.Children.Append("acl", types.YChild{"Acl", &spanMonitorSession.Acl})
    spanMonitorSession.EntityData.Leafs = types.NewOrderedMap()
    spanMonitorSession.EntityData.Leafs.Append("session-class", types.YLeaf{"SessionClass", spanMonitorSession.SessionClass})
    spanMonitorSession.EntityData.Leafs.Append("mirror-first", types.YLeaf{"MirrorFirst", spanMonitorSession.MirrorFirst})
    spanMonitorSession.EntityData.Leafs.Append("mirror-interval", types.YLeaf{"MirrorInterval", spanMonitorSession.MirrorInterval})

    spanMonitorSession.EntityData.YListKeys = []string {"SessionClass"}

    return &(spanMonitorSession.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment
// Attach the interface to a Monitor Session
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Session Name. The type is string with length: 1..79. This attribute is
    // mandatory.
    SessionName interface{}

    // Specify the direction of traffic to replicate (optional). The type is
    // SpanTrafficDirection.
    Direction interface{}

    // Enable port level traffic mirroring. The type is interface{}.
    PortLevelEnable interface{}
}

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetEntityData() *types.CommonEntityData {
    attachment.EntityData.YFilter = attachment.YFilter
    attachment.EntityData.YangName = "attachment"
    attachment.EntityData.BundleName = "cisco_ios_xr"
    attachment.EntityData.ParentYangName = "span-monitor-session"
    attachment.EntityData.SegmentPath = "attachment"
    attachment.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions/span-monitor-session/" + attachment.EntityData.SegmentPath
    attachment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachment.EntityData.Children = types.NewOrderedMap()
    attachment.EntityData.Leafs = types.NewOrderedMap()
    attachment.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", attachment.SessionName})
    attachment.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", attachment.Direction})
    attachment.EntityData.Leafs.Append("port-level-enable", types.YLeaf{"PortLevelEnable", attachment.PortLevelEnable})

    attachment.EntityData.YListKeys = []string {}

    return &(attachment.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl
// Enable ACL matching for traffic mirroring
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable ACL. The type is interface{}. This attribute is mandatory.
    AclEnable interface{}

    // ACL Name. The type is string with length: 1..80.
    AclName interface{}
}

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetEntityData() *types.CommonEntityData {
    acl.EntityData.YFilter = acl.YFilter
    acl.EntityData.YangName = "acl"
    acl.EntityData.BundleName = "cisco_ios_xr"
    acl.EntityData.ParentYangName = "span-monitor-session"
    acl.EntityData.SegmentPath = "acl"
    acl.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions/span-monitor-session/" + acl.EntityData.SegmentPath
    acl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acl.EntityData.Children = types.NewOrderedMap()
    acl.EntityData.Leafs = types.NewOrderedMap()
    acl.EntityData.Leafs.Append("acl-enable", types.YLeaf{"AclEnable", acl.AclEnable})
    acl.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", acl.AclName})

    acl.EntityData.YListKeys = []string {}

    return &(acl.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter
// IPv4 Packet Filtering configuration for the
// template
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound

    // IPv4 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound
}

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetEntityData() *types.CommonEntityData {
    ipv4PacketFilter.EntityData.YFilter = ipv4PacketFilter.YFilter
    ipv4PacketFilter.EntityData.YangName = "ipv4-packet-filter"
    ipv4PacketFilter.EntityData.BundleName = "cisco_ios_xr"
    ipv4PacketFilter.EntityData.ParentYangName = "ip-subscriber"
    ipv4PacketFilter.EntityData.SegmentPath = "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter"
    ipv4PacketFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + ipv4PacketFilter.EntityData.SegmentPath
    ipv4PacketFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PacketFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PacketFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PacketFilter.EntityData.Children = types.NewOrderedMap()
    ipv4PacketFilter.EntityData.Children.Append("outbound", types.YChild{"Outbound", &ipv4PacketFilter.Outbound})
    ipv4PacketFilter.EntityData.Children.Append("inbound", types.YChild{"Inbound", &ipv4PacketFilter.Inbound})
    ipv4PacketFilter.EntityData.Leafs = types.NewOrderedMap()

    ipv4PacketFilter.EntityData.YListKeys = []string {}

    return &(ipv4PacketFilter.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound
// IPv4 Packet filter to be applied to outbound
// packets
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "ipv4-packet-filter"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter/" + outbound.EntityData.SegmentPath
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = types.NewOrderedMap()
    outbound.EntityData.Leafs = types.NewOrderedMap()
    outbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", outbound.CommonAclName})
    outbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", outbound.Name})
    outbound.EntityData.Leafs.Append("hardware-count", types.YLeaf{"HardwareCount", outbound.HardwareCount})
    outbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", outbound.InterfaceStatistics})

    outbound.EntityData.YListKeys = []string {}

    return &(outbound.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound
// IPv4 Packet filter to be applied to inbound
// packets
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Inbound packets NOTE: This
    // parameter is mandatory if 'CommonACLName' is not specified. The type is
    // string with length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "ipv4-packet-filter"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter/" + inbound.EntityData.SegmentPath
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = types.NewOrderedMap()
    inbound.EntityData.Leafs = types.NewOrderedMap()
    inbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", inbound.CommonAclName})
    inbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", inbound.Name})
    inbound.EntityData.Leafs.Append("hardware-count", types.YLeaf{"HardwareCount", inbound.HardwareCount})
    inbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", inbound.InterfaceStatistics})

    inbound.EntityData.YListKeys = []string {}

    return &(inbound.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter
// IPv6 Packet Filtering configuration for the
// interface
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound

    // IPv6 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound
}

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetEntityData() *types.CommonEntityData {
    ipv6PacketFilter.EntityData.YFilter = ipv6PacketFilter.YFilter
    ipv6PacketFilter.EntityData.YangName = "ipv6-packet-filter"
    ipv6PacketFilter.EntityData.BundleName = "cisco_ios_xr"
    ipv6PacketFilter.EntityData.ParentYangName = "ip-subscriber"
    ipv6PacketFilter.EntityData.SegmentPath = "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter"
    ipv6PacketFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + ipv6PacketFilter.EntityData.SegmentPath
    ipv6PacketFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6PacketFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6PacketFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6PacketFilter.EntityData.Children = types.NewOrderedMap()
    ipv6PacketFilter.EntityData.Children.Append("inbound", types.YChild{"Inbound", &ipv6PacketFilter.Inbound})
    ipv6PacketFilter.EntityData.Children.Append("outbound", types.YChild{"Outbound", &ipv6PacketFilter.Outbound})
    ipv6PacketFilter.EntityData.Leafs = types.NewOrderedMap()

    ipv6PacketFilter.EntityData.YListKeys = []string {}

    return &(ipv6PacketFilter.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound
// IPv6 Packet filter to be applied to inbound
// packets
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Inbound  NOTE: This parameter is
    // mandatory if 'CommonACLName' is not specified. The type is string with
    // length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "ipv6-packet-filter"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter/" + inbound.EntityData.SegmentPath
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = types.NewOrderedMap()
    inbound.EntityData.Leafs = types.NewOrderedMap()
    inbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", inbound.CommonAclName})
    inbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", inbound.Name})
    inbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", inbound.InterfaceStatistics})

    inbound.EntityData.YListKeys = []string {}

    return &(inbound.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound
// IPv6 Packet filter to be applied to outbound
// packets
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "ipv6-packet-filter"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter/" + outbound.EntityData.SegmentPath
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = types.NewOrderedMap()
    outbound.EntityData.Leafs = types.NewOrderedMap()
    outbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", outbound.CommonAclName})
    outbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", outbound.Name})
    outbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", outbound.InterfaceStatistics})

    outbound.EntityData.YListKeys = []string {}

    return &(outbound.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd
// Interface dhcpv4 configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Cisco VSA to configure any dhcp4 ip lease per subscriber. The type is
    // string.
    Dhcpv4Iplease interface{}

    // The class to be used for proxy/server profile. The type is string.
    Class interface{}

    // Select proxy/server profile based on mode class name. The type is string.
    ModeClass interface{}

    // The Default Gateway IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DefaultGateway interface{}

    // The pool to be used for Prefix Delegation. The type is interface{} with
    // range: 0..4294967295.
    SessionLimit interface{}

    // Cisco VSA to configure any dhcp4 option per subscriber. The type is string.
    Dhcpv4Option interface{}
}

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetEntityData() *types.CommonEntityData {
    dhcpd.EntityData.YFilter = dhcpd.YFilter
    dhcpd.EntityData.YangName = "dhcpd"
    dhcpd.EntityData.BundleName = "cisco_ios_xr"
    dhcpd.EntityData.ParentYangName = "ip-subscriber"
    dhcpd.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-dhcpd-subscriber-cfg:dhcpd"
    dhcpd.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + dhcpd.EntityData.SegmentPath
    dhcpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpd.EntityData.Children = types.NewOrderedMap()
    dhcpd.EntityData.Leafs = types.NewOrderedMap()
    dhcpd.EntityData.Leafs.Append("dhcpv4-iplease", types.YLeaf{"Dhcpv4Iplease", dhcpd.Dhcpv4Iplease})
    dhcpd.EntityData.Leafs.Append("class", types.YLeaf{"Class", dhcpd.Class})
    dhcpd.EntityData.Leafs.Append("mode-class", types.YLeaf{"ModeClass", dhcpd.ModeClass})
    dhcpd.EntityData.Leafs.Append("default-gateway", types.YLeaf{"DefaultGateway", dhcpd.DefaultGateway})
    dhcpd.EntityData.Leafs.Append("session-limit", types.YLeaf{"SessionLimit", dhcpd.SessionLimit})
    dhcpd.EntityData.Leafs.Append("dhcpv4-option", types.YLeaf{"Dhcpv4Option", dhcpd.Dhcpv4Option})

    dhcpd.EntityData.YListKeys = []string {}

    return &(dhcpd.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Igmp
// IGMPconfiguration
type DynamicTemplate_IpSubscribers_IpSubscriber_Igmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default VRF.
    DefaultVrf DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf
}

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetEntityData() *types.CommonEntityData {
    igmp.EntityData.YFilter = igmp.YFilter
    igmp.EntityData.YangName = "igmp"
    igmp.EntityData.BundleName = "cisco_ios_xr"
    igmp.EntityData.ParentYangName = "ip-subscriber"
    igmp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp"
    igmp.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + igmp.EntityData.SegmentPath
    igmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igmp.EntityData.Children = types.NewOrderedMap()
    igmp.EntityData.Children.Append("default-vrf", types.YChild{"DefaultVrf", &igmp.DefaultVrf})
    igmp.EntityData.Leafs = types.NewOrderedMap()

    igmp.EntityData.YListKeys = []string {}

    return &(igmp.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf
// Default VRF
type DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGMP Max Groups. The type is interface{} with range: 1..40000. The default
    // value is 25000.
    MaxGroups interface{}

    // Access list specifying access-list group range. The type is string with
    // length: 1..64.
    AccessGroup interface{}

    // IGMP Version. The type is interface{} with range: 1..3. The default value
    // is 3.
    Version interface{}

    // Query interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second. The default value is 60.
    QueryInterval interface{}

    // Query response value in seconds. The type is interface{} with range: 1..25.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Configure Multicast mode variable. The type is DynTmplMulticastMode.
    MulticastMode interface{}

    // IGMPv3 explicit host tracking.
    ExplicitTracking DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking
}

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "igmp"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp/" + defaultVrf.EntityData.SegmentPath
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = types.NewOrderedMap()
    defaultVrf.EntityData.Children.Append("explicit-tracking", types.YChild{"ExplicitTracking", &defaultVrf.ExplicitTracking})
    defaultVrf.EntityData.Leafs = types.NewOrderedMap()
    defaultVrf.EntityData.Leafs.Append("max-groups", types.YLeaf{"MaxGroups", defaultVrf.MaxGroups})
    defaultVrf.EntityData.Leafs.Append("access-group", types.YLeaf{"AccessGroup", defaultVrf.AccessGroup})
    defaultVrf.EntityData.Leafs.Append("version", types.YLeaf{"Version", defaultVrf.Version})
    defaultVrf.EntityData.Leafs.Append("query-interval", types.YLeaf{"QueryInterval", defaultVrf.QueryInterval})
    defaultVrf.EntityData.Leafs.Append("query-max-response-time", types.YLeaf{"QueryMaxResponseTime", defaultVrf.QueryMaxResponseTime})
    defaultVrf.EntityData.Leafs.Append("multicast-mode", types.YLeaf{"MulticastMode", defaultVrf.MulticastMode})

    defaultVrf.EntityData.YListKeys = []string {}

    return &(defaultVrf.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable or disable, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetEntityData() *types.CommonEntityData {
    explicitTracking.EntityData.YFilter = explicitTracking.YFilter
    explicitTracking.EntityData.YangName = "explicit-tracking"
    explicitTracking.EntityData.BundleName = "cisco_ios_xr"
    explicitTracking.EntityData.ParentYangName = "default-vrf"
    explicitTracking.EntityData.SegmentPath = "explicit-tracking"
    explicitTracking.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp/default-vrf/" + explicitTracking.EntityData.SegmentPath
    explicitTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicitTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicitTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicitTracking.EntityData.Children = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs = types.NewOrderedMap()
    explicitTracking.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", explicitTracking.Enable})
    explicitTracking.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", explicitTracking.AccessListName})

    explicitTracking.EntityData.YListKeys = []string {}

    return &(explicitTracking.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network
// Interface IPv4 Network configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IP processing without an explicit address. The type is string.
    Unnumbered interface{}

    // The IP Maximum Transmission Unit. The type is interface{} with range:
    // 68..65535. Units are byte.
    Mtu interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool. The default value is
    // false.
    Unreachables interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool. The default value is
    // true.
    Rpf interface{}
}

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetEntityData() *types.CommonEntityData {
    ipv4Network.EntityData.YFilter = ipv4Network.YFilter
    ipv4Network.EntityData.YangName = "ipv4-network"
    ipv4Network.EntityData.BundleName = "cisco_ios_xr"
    ipv4Network.EntityData.ParentYangName = "ip-subscriber"
    ipv4Network.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network"
    ipv4Network.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + ipv4Network.EntityData.SegmentPath
    ipv4Network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Network.EntityData.Children = types.NewOrderedMap()
    ipv4Network.EntityData.Leafs = types.NewOrderedMap()
    ipv4Network.EntityData.Leafs.Append("unnumbered", types.YLeaf{"Unnumbered", ipv4Network.Unnumbered})
    ipv4Network.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", ipv4Network.Mtu})
    ipv4Network.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", ipv4Network.Unreachables})
    ipv4Network.EntityData.Leafs.Append("rpf", types.YLeaf{"Rpf", ipv4Network.Rpf})

    ipv4Network.EntityData.YListKeys = []string {}

    return &(ipv4Network.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network
// Interface IPv6 Network configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MTU Setting of Interface. The type is interface{} with range: 1280..65535.
    // Units are byte.
    Mtu interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool.
    Rpf interface{}

    // Override Sending of ICMP Unreachable Messages. The type is interface{}.
    Unreachables interface{}

    // Set the IPv6 address of an interface.
    Addresses DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses
}

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetEntityData() *types.CommonEntityData {
    ipv6Network.EntityData.YFilter = ipv6Network.YFilter
    ipv6Network.EntityData.YangName = "ipv6-network"
    ipv6Network.EntityData.BundleName = "cisco_ios_xr"
    ipv6Network.EntityData.ParentYangName = "ip-subscriber"
    ipv6Network.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network"
    ipv6Network.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + ipv6Network.EntityData.SegmentPath
    ipv6Network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Network.EntityData.Children = types.NewOrderedMap()
    ipv6Network.EntityData.Children.Append("addresses", types.YChild{"Addresses", &ipv6Network.Addresses})
    ipv6Network.EntityData.Leafs = types.NewOrderedMap()
    ipv6Network.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", ipv6Network.Mtu})
    ipv6Network.EntityData.Leafs.Append("rpf", types.YLeaf{"Rpf", ipv6Network.Rpf})
    ipv6Network.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", ipv6Network.Unreachables})

    ipv6Network.EntityData.YListKeys = []string {}

    return &(ipv6Network.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses
// Set the IPv6 address of an interface
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto IPv6 Interface Configuration.
    AutoConfiguration DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration
}

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "ipv6-network"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network/" + addresses.EntityData.SegmentPath
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("auto-configuration", types.YChild{"AutoConfiguration", &addresses.AutoConfiguration})
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration
// Auto IPv6 Interface Configuration
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The flag to enable auto ipv6 interface configuration. The type is
    // interface{}.
    Enable interface{}
}

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetEntityData() *types.CommonEntityData {
    autoConfiguration.EntityData.YFilter = autoConfiguration.YFilter
    autoConfiguration.EntityData.YangName = "auto-configuration"
    autoConfiguration.EntityData.BundleName = "cisco_ios_xr"
    autoConfiguration.EntityData.ParentYangName = "addresses"
    autoConfiguration.EntityData.SegmentPath = "auto-configuration"
    autoConfiguration.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network/addresses/" + autoConfiguration.EntityData.SegmentPath
    autoConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoConfiguration.EntityData.Children = types.NewOrderedMap()
    autoConfiguration.EntityData.Leafs = types.NewOrderedMap()
    autoConfiguration.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", autoConfiguration.Enable})

    autoConfiguration.EntityData.YListKeys = []string {}

    return &(autoConfiguration.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor
// Interface IPv6 Network configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the IPv6 framed ipv6 prefix pool for a subscriber interface . The type
    // is string.
    FramedPrefixPool interface{}

    // Host to use stateful protocol for address configuration. The type is
    // interface{}.
    ManagedConfig interface{}

    // Host to use stateful protocol for non-address configuration. The type is
    // interface{}.
    OtherConfig interface{}

    // Start RA on ipv6-enable config. The type is interface{}.
    StartRaOnIpv6Enable interface{}

    // NUD enable. The type is interface{}.
    NudEnable interface{}

    // Set IPv6 Router Advertisement (RA) lifetime in seconds. The type is
    // interface{} with range: 0..9000. Units are second.
    RaLifetime interface{}

    // RA Router Preference. The type is Ipv6NdRouterPrefTemplate.
    RouterPreference interface{}

    // Enable suppress IPv6 router advertisement. The type is interface{}.
    RaSuppress interface{}

    // Enable RA unicast Flag. The type is interface{}.
    RaUnicast interface{}

    // Unspecify IPv6 Router Advertisement (RA) hop-limit. The type is
    // interface{}.
    RaUnspecifyHoplimit interface{}

    // RA suppress MTU flag. The type is interface{}.
    RaSuppressMtu interface{}

    // Suppress cache learning flag. The type is interface{}.
    SuppressCacheLearning interface{}

    // Set advertised reachability time in milliseconds. The type is interface{}
    // with range: 0..3600000. Units are millisecond.
    ReachableTime interface{}

    // Set advertised NS retransmission interval in milliseconds. The type is
    // interface{} with range: 1000..4294967295. Units are millisecond.
    NsInterval interface{}

    // Set IPv6 Router Advertisement (RA) interval in seconds.
    RaInterval DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval

    // Set the IPv6 framed ipv6 prefix for a subscriber interface .
    FramedPrefix DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix

    // Duplicate Address Detection (DAD).
    DuplicateAddressDetection DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection

    // IPv6 ND RA Initial.
    RaInitial DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial
}

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetEntityData() *types.CommonEntityData {
    ipv6Neighbor.EntityData.YFilter = ipv6Neighbor.YFilter
    ipv6Neighbor.EntityData.YangName = "ipv6-neighbor"
    ipv6Neighbor.EntityData.BundleName = "cisco_ios_xr"
    ipv6Neighbor.EntityData.ParentYangName = "ip-subscriber"
    ipv6Neighbor.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor"
    ipv6Neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + ipv6Neighbor.EntityData.SegmentPath
    ipv6Neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Neighbor.EntityData.Children = types.NewOrderedMap()
    ipv6Neighbor.EntityData.Children.Append("ra-interval", types.YChild{"RaInterval", &ipv6Neighbor.RaInterval})
    ipv6Neighbor.EntityData.Children.Append("framed-prefix", types.YChild{"FramedPrefix", &ipv6Neighbor.FramedPrefix})
    ipv6Neighbor.EntityData.Children.Append("duplicate-address-detection", types.YChild{"DuplicateAddressDetection", &ipv6Neighbor.DuplicateAddressDetection})
    ipv6Neighbor.EntityData.Children.Append("ra-initial", types.YChild{"RaInitial", &ipv6Neighbor.RaInitial})
    ipv6Neighbor.EntityData.Leafs = types.NewOrderedMap()
    ipv6Neighbor.EntityData.Leafs.Append("framed-prefix-pool", types.YLeaf{"FramedPrefixPool", ipv6Neighbor.FramedPrefixPool})
    ipv6Neighbor.EntityData.Leafs.Append("managed-config", types.YLeaf{"ManagedConfig", ipv6Neighbor.ManagedConfig})
    ipv6Neighbor.EntityData.Leafs.Append("other-config", types.YLeaf{"OtherConfig", ipv6Neighbor.OtherConfig})
    ipv6Neighbor.EntityData.Leafs.Append("start-ra-on-ipv6-enable", types.YLeaf{"StartRaOnIpv6Enable", ipv6Neighbor.StartRaOnIpv6Enable})
    ipv6Neighbor.EntityData.Leafs.Append("nud-enable", types.YLeaf{"NudEnable", ipv6Neighbor.NudEnable})
    ipv6Neighbor.EntityData.Leafs.Append("ra-lifetime", types.YLeaf{"RaLifetime", ipv6Neighbor.RaLifetime})
    ipv6Neighbor.EntityData.Leafs.Append("router-preference", types.YLeaf{"RouterPreference", ipv6Neighbor.RouterPreference})
    ipv6Neighbor.EntityData.Leafs.Append("ra-suppress", types.YLeaf{"RaSuppress", ipv6Neighbor.RaSuppress})
    ipv6Neighbor.EntityData.Leafs.Append("ra-unicast", types.YLeaf{"RaUnicast", ipv6Neighbor.RaUnicast})
    ipv6Neighbor.EntityData.Leafs.Append("ra-unspecify-hoplimit", types.YLeaf{"RaUnspecifyHoplimit", ipv6Neighbor.RaUnspecifyHoplimit})
    ipv6Neighbor.EntityData.Leafs.Append("ra-suppress-mtu", types.YLeaf{"RaSuppressMtu", ipv6Neighbor.RaSuppressMtu})
    ipv6Neighbor.EntityData.Leafs.Append("suppress-cache-learning", types.YLeaf{"SuppressCacheLearning", ipv6Neighbor.SuppressCacheLearning})
    ipv6Neighbor.EntityData.Leafs.Append("reachable-time", types.YLeaf{"ReachableTime", ipv6Neighbor.ReachableTime})
    ipv6Neighbor.EntityData.Leafs.Append("ns-interval", types.YLeaf{"NsInterval", ipv6Neighbor.NsInterval})

    ipv6Neighbor.EntityData.YListKeys = []string {}

    return &(ipv6Neighbor.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval
// Set IPv6 Router Advertisement (RA) interval in
// seconds
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Maximum interface{}

    // Minimum RA interval in seconds. Must be less than 0.75 * maximum interval.
    // The type is interface{} with range: 3..1800. Units are second.
    Minimum interface{}
}

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetEntityData() *types.CommonEntityData {
    raInterval.EntityData.YFilter = raInterval.YFilter
    raInterval.EntityData.YangName = "ra-interval"
    raInterval.EntityData.BundleName = "cisco_ios_xr"
    raInterval.EntityData.ParentYangName = "ipv6-neighbor"
    raInterval.EntityData.SegmentPath = "ra-interval"
    raInterval.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + raInterval.EntityData.SegmentPath
    raInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    raInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    raInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    raInterval.EntityData.Children = types.NewOrderedMap()
    raInterval.EntityData.Leafs = types.NewOrderedMap()
    raInterval.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", raInterval.Maximum})
    raInterval.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", raInterval.Minimum})

    raInterval.EntityData.YListKeys = []string {}

    return &(raInterval.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix
// Set the IPv6 framed ipv6 prefix for a
// subscriber interface 
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IPv6 framed prefix length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}

    // IPV6 framed prefix address. The type is string. This attribute is
    // mandatory.
    Prefix interface{}
}

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetEntityData() *types.CommonEntityData {
    framedPrefix.EntityData.YFilter = framedPrefix.YFilter
    framedPrefix.EntityData.YangName = "framed-prefix"
    framedPrefix.EntityData.BundleName = "cisco_ios_xr"
    framedPrefix.EntityData.ParentYangName = "ipv6-neighbor"
    framedPrefix.EntityData.SegmentPath = "framed-prefix"
    framedPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + framedPrefix.EntityData.SegmentPath
    framedPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    framedPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    framedPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    framedPrefix.EntityData.Children = types.NewOrderedMap()
    framedPrefix.EntityData.Leafs = types.NewOrderedMap()
    framedPrefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", framedPrefix.PrefixLength})
    framedPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", framedPrefix.Prefix})

    framedPrefix.EntityData.YListKeys = []string {}

    return &(framedPrefix.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection
// Duplicate Address Detection (DAD)
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set IPv6 duplicate address detection transmits. The type is interface{}
    // with range: 0..600.
    Attempts interface{}
}

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetEntityData() *types.CommonEntityData {
    duplicateAddressDetection.EntityData.YFilter = duplicateAddressDetection.YFilter
    duplicateAddressDetection.EntityData.YangName = "duplicate-address-detection"
    duplicateAddressDetection.EntityData.BundleName = "cisco_ios_xr"
    duplicateAddressDetection.EntityData.ParentYangName = "ipv6-neighbor"
    duplicateAddressDetection.EntityData.SegmentPath = "duplicate-address-detection"
    duplicateAddressDetection.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + duplicateAddressDetection.EntityData.SegmentPath
    duplicateAddressDetection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duplicateAddressDetection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duplicateAddressDetection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duplicateAddressDetection.EntityData.Children = types.NewOrderedMap()
    duplicateAddressDetection.EntityData.Leafs = types.NewOrderedMap()
    duplicateAddressDetection.EntityData.Leafs.Append("attempts", types.YLeaf{"Attempts", duplicateAddressDetection.Attempts})

    duplicateAddressDetection.EntityData.YListKeys = []string {}

    return &(duplicateAddressDetection.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial
// IPv6 ND RA Initial
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Initial RA count. The type is interface{} with range: 0..32. This attribute
    // is mandatory.
    Count interface{}

    // Initial RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Interval interface{}
}

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetEntityData() *types.CommonEntityData {
    raInitial.EntityData.YFilter = raInitial.YFilter
    raInitial.EntityData.YangName = "ra-initial"
    raInitial.EntityData.BundleName = "cisco_ios_xr"
    raInitial.EntityData.ParentYangName = "ipv6-neighbor"
    raInitial.EntityData.SegmentPath = "ra-initial"
    raInitial.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + raInitial.EntityData.SegmentPath
    raInitial.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    raInitial.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    raInitial.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    raInitial.EntityData.Children = types.NewOrderedMap()
    raInitial.EntityData.Leafs = types.NewOrderedMap()
    raInitial.EntityData.Leafs.Append("count", types.YLeaf{"Count", raInitial.Count})
    raInitial.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", raInitial.Interval})

    raInitial.EntityData.YListKeys = []string {}

    return &(raInitial.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6
// Interface dhcpv6 configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dns IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DnsIpv6address interface{}

    // Select proxy/server profile based on mode class name. The type is string.
    ModeClass interface{}

    // Cisco VSA to configure any dhcpv6 ip lease per subscriber. The type is
    // string.
    Dhcpv6Iplease interface{}

    // Cisco VSA to configure any dhcpv6 option per subscriber. The type is
    // string.
    Dhcpv6Option interface{}

    // The pool to be used for Address assignment. The type is string.
    AddressPool interface{}

    // The pool to be used for Prefix Delegation. The type is string.
    DelegatedPrefixPool interface{}

    // The class to be used for proxy/server profile. The type is string.
    Class interface{}

    // Stateful IPv6 Address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    StatefulAddress interface{}

    // The prefix to be used for Prefix Delegation.
    DelegatedPrefix DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix
}

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetEntityData() *types.CommonEntityData {
    dhcpv6.EntityData.YFilter = dhcpv6.YFilter
    dhcpv6.EntityData.YangName = "dhcpv6"
    dhcpv6.EntityData.BundleName = "cisco_ios_xr"
    dhcpv6.EntityData.ParentYangName = "ip-subscriber"
    dhcpv6.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6"
    dhcpv6.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + dhcpv6.EntityData.SegmentPath
    dhcpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpv6.EntityData.Children = types.NewOrderedMap()
    dhcpv6.EntityData.Children.Append("delegated-prefix", types.YChild{"DelegatedPrefix", &dhcpv6.DelegatedPrefix})
    dhcpv6.EntityData.Leafs = types.NewOrderedMap()
    dhcpv6.EntityData.Leafs.Append("dns-ipv6address", types.YLeaf{"DnsIpv6address", dhcpv6.DnsIpv6address})
    dhcpv6.EntityData.Leafs.Append("mode-class", types.YLeaf{"ModeClass", dhcpv6.ModeClass})
    dhcpv6.EntityData.Leafs.Append("dhcpv6-iplease", types.YLeaf{"Dhcpv6Iplease", dhcpv6.Dhcpv6Iplease})
    dhcpv6.EntityData.Leafs.Append("dhcpv6-option", types.YLeaf{"Dhcpv6Option", dhcpv6.Dhcpv6Option})
    dhcpv6.EntityData.Leafs.Append("address-pool", types.YLeaf{"AddressPool", dhcpv6.AddressPool})
    dhcpv6.EntityData.Leafs.Append("delegated-prefix-pool", types.YLeaf{"DelegatedPrefixPool", dhcpv6.DelegatedPrefixPool})
    dhcpv6.EntityData.Leafs.Append("class", types.YLeaf{"Class", dhcpv6.Class})
    dhcpv6.EntityData.Leafs.Append("stateful-address", types.YLeaf{"StatefulAddress", dhcpv6.StatefulAddress})

    dhcpv6.EntityData.YListKeys = []string {}

    return &(dhcpv6.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix
// The prefix to be used for Prefix Delegation
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IPv6 Prefix. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    // This attribute is mandatory.
    Prefix interface{}

    // PD Prefix Length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}
}

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetEntityData() *types.CommonEntityData {
    delegatedPrefix.EntityData.YFilter = delegatedPrefix.YFilter
    delegatedPrefix.EntityData.YangName = "delegated-prefix"
    delegatedPrefix.EntityData.BundleName = "cisco_ios_xr"
    delegatedPrefix.EntityData.ParentYangName = "dhcpv6"
    delegatedPrefix.EntityData.SegmentPath = "delegated-prefix"
    delegatedPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6/" + delegatedPrefix.EntityData.SegmentPath
    delegatedPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delegatedPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delegatedPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delegatedPrefix.EntityData.Children = types.NewOrderedMap()
    delegatedPrefix.EntityData.Leafs = types.NewOrderedMap()
    delegatedPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", delegatedPrefix.Prefix})
    delegatedPrefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", delegatedPrefix.PrefixLength})

    delegatedPrefix.EntityData.YListKeys = []string {}

    return &(delegatedPrefix.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Pbr
// Dynamic Template PBR configuration
type DynamicTemplate_IpSubscribers_IpSubscriber_Pbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class for subscriber ingress policy. The type is string.
    ServicePolicyIn interface{}

    // Ingress service policy.
    ServicePolicies DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicies
}

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetEntityData() *types.CommonEntityData {
    pbr.EntityData.YFilter = pbr.YFilter
    pbr.EntityData.YangName = "pbr"
    pbr.EntityData.BundleName = "cisco_ios_xr"
    pbr.EntityData.ParentYangName = "ip-subscriber"
    pbr.EntityData.SegmentPath = "Cisco-IOS-XR-pbr-subscriber-cfg:pbr"
    pbr.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + pbr.EntityData.SegmentPath
    pbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbr.EntityData.Children = types.NewOrderedMap()
    pbr.EntityData.Children.Append("service-policies", types.YChild{"ServicePolicies", &pbr.ServicePolicies})
    pbr.EntityData.Leafs = types.NewOrderedMap()
    pbr.EntityData.Leafs.Append("service-policy-in", types.YLeaf{"ServicePolicyIn", pbr.ServicePolicyIn})

    pbr.EntityData.YListKeys = []string {}

    return &(pbr.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicies
// Ingress service policy
type DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service policy details. The type is slice of
    // DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicies_ServicePolicy.
    ServicePolicy []*DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicies_ServicePolicy
}

func (servicePolicies *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicies) GetEntityData() *types.CommonEntityData {
    servicePolicies.EntityData.YFilter = servicePolicies.YFilter
    servicePolicies.EntityData.YangName = "service-policies"
    servicePolicies.EntityData.BundleName = "cisco_ios_xr"
    servicePolicies.EntityData.ParentYangName = "pbr"
    servicePolicies.EntityData.SegmentPath = "service-policies"
    servicePolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-pbr-subscriber-cfg:pbr/" + servicePolicies.EntityData.SegmentPath
    servicePolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicies.EntityData.Children = types.NewOrderedMap()
    servicePolicies.EntityData.Children.Append("service-policy", types.YChild{"ServicePolicy", nil})
    for i := range servicePolicies.ServicePolicy {
        servicePolicies.EntityData.Children.Append(types.GetSegmentPath(servicePolicies.ServicePolicy[i]), types.YChild{"ServicePolicy", servicePolicies.ServicePolicy[i]})
    }
    servicePolicies.EntityData.Leafs = types.NewOrderedMap()

    servicePolicies.EntityData.YListKeys = []string {}

    return &(servicePolicies.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicies_ServicePolicy
// Service policy details
type DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicies_ServicePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of policy-map. The type is string with
    // length: 1..64.
    ServicePolicy interface{}
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicies_ServicePolicy) GetEntityData() *types.CommonEntityData {
    servicePolicy.EntityData.YFilter = servicePolicy.YFilter
    servicePolicy.EntityData.YangName = "service-policy"
    servicePolicy.EntityData.BundleName = "cisco_ios_xr"
    servicePolicy.EntityData.ParentYangName = "service-policies"
    servicePolicy.EntityData.SegmentPath = "service-policy" + types.AddKeyToken(servicePolicy.ServicePolicy, "service-policy")
    servicePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-pbr-subscriber-cfg:pbr/service-policies/" + servicePolicy.EntityData.SegmentPath
    servicePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicy.EntityData.Children = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs.Append("service-policy", types.YLeaf{"ServicePolicy", servicePolicy.ServicePolicy})

    servicePolicy.EntityData.YListKeys = []string {"ServicePolicy"}

    return &(servicePolicy.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos
// QoS dynamically applied configuration template
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service policy to be applied in ingress/egress direction.
    ServicePolicy DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy

    // QoS L2 overhead accounting.
    Account DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account

    // QoS to be applied in egress direction.
    Output DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output
}

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetEntityData() *types.CommonEntityData {
    qos.EntityData.YFilter = qos.YFilter
    qos.EntityData.YangName = "qos"
    qos.EntityData.BundleName = "cisco_ios_xr"
    qos.EntityData.ParentYangName = "ip-subscriber"
    qos.EntityData.SegmentPath = "Cisco-IOS-XR-qos-ma-bng-cfg:qos"
    qos.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + qos.EntityData.SegmentPath
    qos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qos.EntityData.Children = types.NewOrderedMap()
    qos.EntityData.Children.Append("service-policy", types.YChild{"ServicePolicy", &qos.ServicePolicy})
    qos.EntityData.Children.Append("account", types.YChild{"Account", &qos.Account})
    qos.EntityData.Children.Append("output", types.YChild{"Output", &qos.Output})
    qos.EntityData.Leafs = types.NewOrderedMap()

    qos.EntityData.YListKeys = []string {}

    return &(qos.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy
// Service policy to be applied in ingress/egress
// direction
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber ingress policy.
    Input DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input

    // Subscriber egress policy.
    Output DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetEntityData() *types.CommonEntityData {
    servicePolicy.EntityData.YFilter = servicePolicy.YFilter
    servicePolicy.EntityData.YangName = "service-policy"
    servicePolicy.EntityData.BundleName = "cisco_ios_xr"
    servicePolicy.EntityData.ParentYangName = "qos"
    servicePolicy.EntityData.SegmentPath = "service-policy"
    servicePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-qos-ma-bng-cfg:qos/" + servicePolicy.EntityData.SegmentPath
    servicePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicy.EntityData.Children = types.NewOrderedMap()
    servicePolicy.EntityData.Children.Append("input", types.YChild{"Input", &servicePolicy.Input})
    servicePolicy.EntityData.Children.Append("output", types.YChild{"Output", &servicePolicy.Output})
    servicePolicy.EntityData.Leafs = types.NewOrderedMap()

    servicePolicy.EntityData.YListKeys = []string {}

    return &(servicePolicy.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input
// Subscriber ingress policy
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Name of policy-map. The type is string. This attribute is mandatory.
    PolicyName interface{}

    // Name of the SPI. The type is string.
    SpiName interface{}

    // TRUE for merge enabled for service-policy applied on dynamic template. The
    // type is bool.
    Merge interface{}

    // Merge ID value. The type is interface{} with range: 0..255.
    MergeId interface{}

    // TRUE for account stats enabled for service-policy applied on dynamic
    // template. Note: account stats not supported for subscriber type 'ppp' and
    // 'ipsubscriber'. The type is bool.
    AccountStats interface{}
}

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "service-policy"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-qos-ma-bng-cfg:qos/service-policy/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", input.PolicyName})
    input.EntityData.Leafs.Append("spi-name", types.YLeaf{"SpiName", input.SpiName})
    input.EntityData.Leafs.Append("merge", types.YLeaf{"Merge", input.Merge})
    input.EntityData.Leafs.Append("merge-id", types.YLeaf{"MergeId", input.MergeId})
    input.EntityData.Leafs.Append("account-stats", types.YLeaf{"AccountStats", input.AccountStats})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output
// Subscriber egress policy
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Name of policy-map. The type is string. This attribute is mandatory.
    PolicyName interface{}

    // Name of the SPI. The type is string.
    SpiName interface{}

    // TRUE for merge enabled for service-policy applied on dynamic template. The
    // type is bool.
    Merge interface{}

    // Merge ID value. The type is interface{} with range: 0..255.
    MergeId interface{}

    // TRUE for account stats enabled for service-policy applied on dynamic
    // template. Note: account stats not supported for subscriber type 'ppp' and
    // 'ipsubscriber'. The type is bool.
    AccountStats interface{}
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "service-policy"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-qos-ma-bng-cfg:qos/service-policy/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", output.PolicyName})
    output.EntityData.Leafs.Append("spi-name", types.YLeaf{"SpiName", output.SpiName})
    output.EntityData.Leafs.Append("merge", types.YLeaf{"Merge", output.Merge})
    output.EntityData.Leafs.Append("merge-id", types.YLeaf{"MergeId", output.MergeId})
    output.EntityData.Leafs.Append("account-stats", types.YLeaf{"AccountStats", output.AccountStats})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account
// QoS L2 overhead accounting
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ATM adaptation layer AAL. The type is Qosl2DataLink.
    Aal interface{}

    // Specify encapsulation type. The type is Qosl2Encap.
    Encapsulation interface{}

    // ATM cell tax to L2 overhead. The type is interface{}.
    AtmCellTax interface{}

    // Numeric L2 overhead offset. The type is interface{} with range: -63..63.
    UserDefined interface{}
}

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetEntityData() *types.CommonEntityData {
    account.EntityData.YFilter = account.YFilter
    account.EntityData.YangName = "account"
    account.EntityData.BundleName = "cisco_ios_xr"
    account.EntityData.ParentYangName = "qos"
    account.EntityData.SegmentPath = "account"
    account.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-qos-ma-bng-cfg:qos/" + account.EntityData.SegmentPath
    account.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    account.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    account.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    account.EntityData.Children = types.NewOrderedMap()
    account.EntityData.Leafs = types.NewOrderedMap()
    account.EntityData.Leafs.Append("aal", types.YLeaf{"Aal", account.Aal})
    account.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", account.Encapsulation})
    account.EntityData.Leafs.Append("atm-cell-tax", types.YLeaf{"AtmCellTax", account.AtmCellTax})
    account.EntityData.Leafs.Append("user-defined", types.YLeaf{"UserDefined", account.UserDefined})

    account.EntityData.YListKeys = []string {}

    return &(account.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output
// QoS to be applied in egress direction
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth value for the subscriber (in kbps). The type is
    // interface{} with range: 1..4294967295. Units are kbit/s.
    MinimumBandwidth interface{}
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "qos"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-qos-ma-bng-cfg:qos/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("minimum-bandwidth", types.YLeaf{"MinimumBandwidth", output.MinimumBandwidth})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Accounting
// Subscriber accounting dynamic-template commands
type DynamicTemplate_IpSubscribers_IpSubscriber_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber monitor feature. The type is string.
    MonitorFeature interface{}

    // Subscriber accounting prepaid feature. The type is string.
    PrepaidFeature interface{}

    // Subscriber accounting service accounting.
    ServiceAccounting DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting

    // Subscriber accounting session accounting.
    Session DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session

    // Subscriber accounting idle timeout.
    IdleTimeout DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout
}

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "ip-subscriber"
    accounting.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-accounting-cfg:accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("service-accounting", types.YChild{"ServiceAccounting", &accounting.ServiceAccounting})
    accounting.EntityData.Children.Append("session", types.YChild{"Session", &accounting.Session})
    accounting.EntityData.Children.Append("idle-timeout", types.YChild{"IdleTimeout", &accounting.IdleTimeout})
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("monitor-feature", types.YLeaf{"MonitorFeature", accounting.MonitorFeature})
    accounting.EntityData.Leafs.Append("prepaid-feature", types.YLeaf{"PrepaidFeature", accounting.PrepaidFeature})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting
// Subscriber accounting service accounting
type DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service accounting method list name. The type is string.
    MethodListName interface{}

    // Accounting interim interval in minutes. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    AccountingInterimInterval interface{}
}

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetEntityData() *types.CommonEntityData {
    serviceAccounting.EntityData.YFilter = serviceAccounting.YFilter
    serviceAccounting.EntityData.YangName = "service-accounting"
    serviceAccounting.EntityData.BundleName = "cisco_ios_xr"
    serviceAccounting.EntityData.ParentYangName = "accounting"
    serviceAccounting.EntityData.SegmentPath = "service-accounting"
    serviceAccounting.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-subscriber-accounting-cfg:accounting/" + serviceAccounting.EntityData.SegmentPath
    serviceAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceAccounting.EntityData.Children = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs.Append("method-list-name", types.YLeaf{"MethodListName", serviceAccounting.MethodListName})
    serviceAccounting.EntityData.Leafs.Append("accounting-interim-interval", types.YLeaf{"AccountingInterimInterval", serviceAccounting.AccountingInterimInterval})

    serviceAccounting.EntityData.YListKeys = []string {}

    return &(serviceAccounting.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session
// Subscriber accounting session accounting
type DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session accounting method list name. The type is string.
    MethodListName interface{}

    // Interim accounting interval in minutes. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    PeriodicInterval interface{}

    // Dual stack wait delay in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    DualStackDelay interface{}

    // Hold Accounting start based on IA_PD. The type is HoldAcctStart.
    HoldAcctStart interface{}
}

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "accounting"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-subscriber-accounting-cfg:accounting/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("method-list-name", types.YLeaf{"MethodListName", session.MethodListName})
    session.EntityData.Leafs.Append("periodic-interval", types.YLeaf{"PeriodicInterval", session.PeriodicInterval})
    session.EntityData.Leafs.Append("dual-stack-delay", types.YLeaf{"DualStackDelay", session.DualStackDelay})
    session.EntityData.Leafs.Append("hold-acct-start", types.YLeaf{"HoldAcctStart", session.HoldAcctStart})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session_HoldAcctStart represents Hold Accounting start based on IA_PD
type DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session_HoldAcctStart string

const (
    // Based on ipv6 delegated prefix
    DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session_HoldAcctStart_ipv6_prefix_delegation DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session_HoldAcctStart = "ipv6-prefix-delegation"
)

// DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout
// Subscriber accounting idle timeout
type DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Idle timeout value in seconds. The type is interface{} with range:
    // 60..4320000. Units are second.
    TimeoutValue interface{}

    // Threshold in minute(s) per packet. The type is interface{} with range:
    // 1..10000.
    Threshold interface{}

    // Idle timeout traffic direction. The type is string.
    Direction interface{}
}

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetEntityData() *types.CommonEntityData {
    idleTimeout.EntityData.YFilter = idleTimeout.YFilter
    idleTimeout.EntityData.YangName = "idle-timeout"
    idleTimeout.EntityData.BundleName = "cisco_ios_xr"
    idleTimeout.EntityData.ParentYangName = "accounting"
    idleTimeout.EntityData.SegmentPath = "idle-timeout"
    idleTimeout.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/ip-subscribers/ip-subscriber/Cisco-IOS-XR-subscriber-accounting-cfg:accounting/" + idleTimeout.EntityData.SegmentPath
    idleTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idleTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idleTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idleTimeout.EntityData.Children = types.NewOrderedMap()
    idleTimeout.EntityData.Leafs = types.NewOrderedMap()
    idleTimeout.EntityData.Leafs.Append("timeout-value", types.YLeaf{"TimeoutValue", idleTimeout.TimeoutValue})
    idleTimeout.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", idleTimeout.Threshold})
    idleTimeout.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", idleTimeout.Direction})

    idleTimeout.EntityData.YListKeys = []string {}

    return &(idleTimeout.EntityData)
}

// DynamicTemplate_SubscriberServices
// The Service Type Template Table
type DynamicTemplate_SubscriberServices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Service Type Template . The type is slice of
    // DynamicTemplate_SubscriberServices_SubscriberService.
    SubscriberService []*DynamicTemplate_SubscriberServices_SubscriberService
}

func (subscriberServices *DynamicTemplate_SubscriberServices) GetEntityData() *types.CommonEntityData {
    subscriberServices.EntityData.YFilter = subscriberServices.YFilter
    subscriberServices.EntityData.YangName = "subscriber-services"
    subscriberServices.EntityData.BundleName = "cisco_ios_xr"
    subscriberServices.EntityData.ParentYangName = "dynamic-template"
    subscriberServices.EntityData.SegmentPath = "subscriber-services"
    subscriberServices.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/" + subscriberServices.EntityData.SegmentPath
    subscriberServices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberServices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberServices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberServices.EntityData.Children = types.NewOrderedMap()
    subscriberServices.EntityData.Children.Append("subscriber-service", types.YChild{"SubscriberService", nil})
    for i := range subscriberServices.SubscriberService {
        subscriberServices.EntityData.Children.Append(types.GetSegmentPath(subscriberServices.SubscriberService[i]), types.YChild{"SubscriberService", subscriberServices.SubscriberService[i]})
    }
    subscriberServices.EntityData.Leafs = types.NewOrderedMap()

    subscriberServices.EntityData.YListKeys = []string {}

    return &(subscriberServices.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService
// A Service Type Template 
type DynamicTemplate_SubscriberServices_SubscriberService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of the template. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TemplateName interface{}

    // Assign the interface to a VRF . The type is string with length: 1..32.
    Vrf interface{}

    // Monitor Session container for this template.
    SpanMonitorSessions DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions

    // IPv4 Packet Filtering configuration for the template.
    Ipv4PacketFilter DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter

    // IPv6 Packet Filtering configuration for the interface.
    Ipv6PacketFilter DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter

    // Interface IPv4 Network configuration data.
    Ipv4Network DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network

    // Interface IPv6 Network configuration data.
    Ipv6Network DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network

    // Interface IPv6 Network configuration data.
    Ipv6Neighbor DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor

    // Subscriber attribute configuration data.
    SubscriberAttribute DynamicTemplate_SubscriberServices_SubscriberService_SubscriberAttribute

    // Dynamic Template PBR configuration.
    Pbr DynamicTemplate_SubscriberServices_SubscriberService_Pbr

    // QoS dynamically applied configuration template.
    Qos DynamicTemplate_SubscriberServices_SubscriberService_Qos

    // Subscriber accounting dynamic-template commands.
    Accounting DynamicTemplate_SubscriberServices_SubscriberService_Accounting
}

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetEntityData() *types.CommonEntityData {
    subscriberService.EntityData.YFilter = subscriberService.YFilter
    subscriberService.EntityData.YangName = "subscriber-service"
    subscriberService.EntityData.BundleName = "cisco_ios_xr"
    subscriberService.EntityData.ParentYangName = "subscriber-services"
    subscriberService.EntityData.SegmentPath = "subscriber-service" + types.AddKeyToken(subscriberService.TemplateName, "template-name")
    subscriberService.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/" + subscriberService.EntityData.SegmentPath
    subscriberService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberService.EntityData.Children = types.NewOrderedMap()
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions", types.YChild{"SpanMonitorSessions", &subscriberService.SpanMonitorSessions})
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter", types.YChild{"Ipv4PacketFilter", &subscriberService.Ipv4PacketFilter})
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter", types.YChild{"Ipv6PacketFilter", &subscriberService.Ipv6PacketFilter})
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network", types.YChild{"Ipv4Network", &subscriberService.Ipv4Network})
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network", types.YChild{"Ipv6Network", &subscriberService.Ipv6Network})
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor", types.YChild{"Ipv6Neighbor", &subscriberService.Ipv6Neighbor})
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-opendns-deviceid-cfg:subscriber-attribute", types.YChild{"SubscriberAttribute", &subscriberService.SubscriberAttribute})
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-pbr-subscriber-cfg:pbr", types.YChild{"Pbr", &subscriberService.Pbr})
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-qos-ma-bng-cfg:qos", types.YChild{"Qos", &subscriberService.Qos})
    subscriberService.EntityData.Children.Append("Cisco-IOS-XR-subscriber-accounting-cfg:accounting", types.YChild{"Accounting", &subscriberService.Accounting})
    subscriberService.EntityData.Leafs = types.NewOrderedMap()
    subscriberService.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", subscriberService.TemplateName})
    subscriberService.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", subscriberService.Vrf})

    subscriberService.EntityData.YListKeys = []string {"TemplateName"}

    return &(subscriberService.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions
// Monitor Session container for this template
type DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular class of Monitor Session. The type is slice
    // of
    // DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession.
    SpanMonitorSession []*DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession
}

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetEntityData() *types.CommonEntityData {
    spanMonitorSessions.EntityData.YFilter = spanMonitorSessions.YFilter
    spanMonitorSessions.EntityData.YangName = "span-monitor-sessions"
    spanMonitorSessions.EntityData.BundleName = "cisco_ios_xr"
    spanMonitorSessions.EntityData.ParentYangName = "subscriber-service"
    spanMonitorSessions.EntityData.SegmentPath = "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions"
    spanMonitorSessions.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + spanMonitorSessions.EntityData.SegmentPath
    spanMonitorSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMonitorSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMonitorSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMonitorSessions.EntityData.Children = types.NewOrderedMap()
    spanMonitorSessions.EntityData.Children.Append("span-monitor-session", types.YChild{"SpanMonitorSession", nil})
    for i := range spanMonitorSessions.SpanMonitorSession {
        spanMonitorSessions.EntityData.Children.Append(types.GetSegmentPath(spanMonitorSessions.SpanMonitorSession[i]), types.YChild{"SpanMonitorSession", spanMonitorSessions.SpanMonitorSession[i]})
    }
    spanMonitorSessions.EntityData.Leafs = types.NewOrderedMap()

    spanMonitorSessions.EntityData.YListKeys = []string {}

    return &(spanMonitorSessions.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession
// Configuration for a particular class of Monitor
// Session
type DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Session Class. The type is SpanSessionClass.
    SessionClass interface{}

    // Mirror a specified number of bytes from start of packet. The type is
    // interface{} with range: 1..10000. Units are byte.
    MirrorFirst interface{}

    // Specify the mirror interval. The type is SpanMirrorInterval.
    MirrorInterval interface{}

    // Attach the interface to a Monitor Session.
    Attachment DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment

    // Enable ACL matching for traffic mirroring.
    Acl DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl
}

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetEntityData() *types.CommonEntityData {
    spanMonitorSession.EntityData.YFilter = spanMonitorSession.YFilter
    spanMonitorSession.EntityData.YangName = "span-monitor-session"
    spanMonitorSession.EntityData.BundleName = "cisco_ios_xr"
    spanMonitorSession.EntityData.ParentYangName = "span-monitor-sessions"
    spanMonitorSession.EntityData.SegmentPath = "span-monitor-session" + types.AddKeyToken(spanMonitorSession.SessionClass, "session-class")
    spanMonitorSession.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions/" + spanMonitorSession.EntityData.SegmentPath
    spanMonitorSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spanMonitorSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spanMonitorSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spanMonitorSession.EntityData.Children = types.NewOrderedMap()
    spanMonitorSession.EntityData.Children.Append("attachment", types.YChild{"Attachment", &spanMonitorSession.Attachment})
    spanMonitorSession.EntityData.Children.Append("acl", types.YChild{"Acl", &spanMonitorSession.Acl})
    spanMonitorSession.EntityData.Leafs = types.NewOrderedMap()
    spanMonitorSession.EntityData.Leafs.Append("session-class", types.YLeaf{"SessionClass", spanMonitorSession.SessionClass})
    spanMonitorSession.EntityData.Leafs.Append("mirror-first", types.YLeaf{"MirrorFirst", spanMonitorSession.MirrorFirst})
    spanMonitorSession.EntityData.Leafs.Append("mirror-interval", types.YLeaf{"MirrorInterval", spanMonitorSession.MirrorInterval})

    spanMonitorSession.EntityData.YListKeys = []string {"SessionClass"}

    return &(spanMonitorSession.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment
// Attach the interface to a Monitor Session
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Session Name. The type is string with length: 1..79. This attribute is
    // mandatory.
    SessionName interface{}

    // Specify the direction of traffic to replicate (optional). The type is
    // SpanTrafficDirection.
    Direction interface{}

    // Enable port level traffic mirroring. The type is interface{}.
    PortLevelEnable interface{}
}

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetEntityData() *types.CommonEntityData {
    attachment.EntityData.YFilter = attachment.YFilter
    attachment.EntityData.YangName = "attachment"
    attachment.EntityData.BundleName = "cisco_ios_xr"
    attachment.EntityData.ParentYangName = "span-monitor-session"
    attachment.EntityData.SegmentPath = "attachment"
    attachment.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions/span-monitor-session/" + attachment.EntityData.SegmentPath
    attachment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attachment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attachment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attachment.EntityData.Children = types.NewOrderedMap()
    attachment.EntityData.Leafs = types.NewOrderedMap()
    attachment.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", attachment.SessionName})
    attachment.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", attachment.Direction})
    attachment.EntityData.Leafs.Append("port-level-enable", types.YLeaf{"PortLevelEnable", attachment.PortLevelEnable})

    attachment.EntityData.YListKeys = []string {}

    return &(attachment.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl
// Enable ACL matching for traffic mirroring
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable ACL. The type is interface{}. This attribute is mandatory.
    AclEnable interface{}

    // ACL Name. The type is string with length: 1..80.
    AclName interface{}
}

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetEntityData() *types.CommonEntityData {
    acl.EntityData.YFilter = acl.YFilter
    acl.EntityData.YangName = "acl"
    acl.EntityData.BundleName = "cisco_ios_xr"
    acl.EntityData.ParentYangName = "span-monitor-session"
    acl.EntityData.SegmentPath = "acl"
    acl.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions/span-monitor-session/" + acl.EntityData.SegmentPath
    acl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acl.EntityData.Children = types.NewOrderedMap()
    acl.EntityData.Leafs = types.NewOrderedMap()
    acl.EntityData.Leafs.Append("acl-enable", types.YLeaf{"AclEnable", acl.AclEnable})
    acl.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", acl.AclName})

    acl.EntityData.YListKeys = []string {}

    return &(acl.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter
// IPv4 Packet Filtering configuration for the
// template
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound

    // IPv4 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound
}

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetEntityData() *types.CommonEntityData {
    ipv4PacketFilter.EntityData.YFilter = ipv4PacketFilter.YFilter
    ipv4PacketFilter.EntityData.YangName = "ipv4-packet-filter"
    ipv4PacketFilter.EntityData.BundleName = "cisco_ios_xr"
    ipv4PacketFilter.EntityData.ParentYangName = "subscriber-service"
    ipv4PacketFilter.EntityData.SegmentPath = "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter"
    ipv4PacketFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + ipv4PacketFilter.EntityData.SegmentPath
    ipv4PacketFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4PacketFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4PacketFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4PacketFilter.EntityData.Children = types.NewOrderedMap()
    ipv4PacketFilter.EntityData.Children.Append("outbound", types.YChild{"Outbound", &ipv4PacketFilter.Outbound})
    ipv4PacketFilter.EntityData.Children.Append("inbound", types.YChild{"Inbound", &ipv4PacketFilter.Inbound})
    ipv4PacketFilter.EntityData.Leafs = types.NewOrderedMap()

    ipv4PacketFilter.EntityData.YListKeys = []string {}

    return &(ipv4PacketFilter.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound
// IPv4 Packet filter to be applied to outbound
// packets
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "ipv4-packet-filter"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter/" + outbound.EntityData.SegmentPath
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = types.NewOrderedMap()
    outbound.EntityData.Leafs = types.NewOrderedMap()
    outbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", outbound.CommonAclName})
    outbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", outbound.Name})
    outbound.EntityData.Leafs.Append("hardware-count", types.YLeaf{"HardwareCount", outbound.HardwareCount})
    outbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", outbound.InterfaceStatistics})

    outbound.EntityData.YListKeys = []string {}

    return &(outbound.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound
// IPv4 Packet filter to be applied to inbound
// packets
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Inbound packets NOTE: This
    // parameter is mandatory if 'CommonACLName' is not specified. The type is
    // string with length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "ipv4-packet-filter"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter/" + inbound.EntityData.SegmentPath
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = types.NewOrderedMap()
    inbound.EntityData.Leafs = types.NewOrderedMap()
    inbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", inbound.CommonAclName})
    inbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", inbound.Name})
    inbound.EntityData.Leafs.Append("hardware-count", types.YLeaf{"HardwareCount", inbound.HardwareCount})
    inbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", inbound.InterfaceStatistics})

    inbound.EntityData.YListKeys = []string {}

    return &(inbound.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter
// IPv6 Packet Filtering configuration for the
// interface
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound

    // IPv6 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound
}

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetEntityData() *types.CommonEntityData {
    ipv6PacketFilter.EntityData.YFilter = ipv6PacketFilter.YFilter
    ipv6PacketFilter.EntityData.YangName = "ipv6-packet-filter"
    ipv6PacketFilter.EntityData.BundleName = "cisco_ios_xr"
    ipv6PacketFilter.EntityData.ParentYangName = "subscriber-service"
    ipv6PacketFilter.EntityData.SegmentPath = "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter"
    ipv6PacketFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + ipv6PacketFilter.EntityData.SegmentPath
    ipv6PacketFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6PacketFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6PacketFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6PacketFilter.EntityData.Children = types.NewOrderedMap()
    ipv6PacketFilter.EntityData.Children.Append("inbound", types.YChild{"Inbound", &ipv6PacketFilter.Inbound})
    ipv6PacketFilter.EntityData.Children.Append("outbound", types.YChild{"Outbound", &ipv6PacketFilter.Outbound})
    ipv6PacketFilter.EntityData.Leafs = types.NewOrderedMap()

    ipv6PacketFilter.EntityData.YListKeys = []string {}

    return &(ipv6PacketFilter.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound
// IPv6 Packet filter to be applied to inbound
// packets
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Inbound  NOTE: This parameter is
    // mandatory if 'CommonACLName' is not specified. The type is string with
    // length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetEntityData() *types.CommonEntityData {
    inbound.EntityData.YFilter = inbound.YFilter
    inbound.EntityData.YangName = "inbound"
    inbound.EntityData.BundleName = "cisco_ios_xr"
    inbound.EntityData.ParentYangName = "ipv6-packet-filter"
    inbound.EntityData.SegmentPath = "inbound"
    inbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter/" + inbound.EntityData.SegmentPath
    inbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inbound.EntityData.Children = types.NewOrderedMap()
    inbound.EntityData.Leafs = types.NewOrderedMap()
    inbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", inbound.CommonAclName})
    inbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", inbound.Name})
    inbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", inbound.InterfaceStatistics})

    inbound.EntityData.YListKeys = []string {}

    return &(inbound.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound
// IPv6 Packet filter to be applied to outbound
// packets
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string with length: 1..64.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..64.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetEntityData() *types.CommonEntityData {
    outbound.EntityData.YFilter = outbound.YFilter
    outbound.EntityData.YangName = "outbound"
    outbound.EntityData.BundleName = "cisco_ios_xr"
    outbound.EntityData.ParentYangName = "ipv6-packet-filter"
    outbound.EntityData.SegmentPath = "outbound"
    outbound.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter/" + outbound.EntityData.SegmentPath
    outbound.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outbound.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outbound.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outbound.EntityData.Children = types.NewOrderedMap()
    outbound.EntityData.Leafs = types.NewOrderedMap()
    outbound.EntityData.Leafs.Append("common-acl-name", types.YLeaf{"CommonAclName", outbound.CommonAclName})
    outbound.EntityData.Leafs.Append("name", types.YLeaf{"Name", outbound.Name})
    outbound.EntityData.Leafs.Append("interface-statistics", types.YLeaf{"InterfaceStatistics", outbound.InterfaceStatistics})

    outbound.EntityData.YListKeys = []string {}

    return &(outbound.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network
// Interface IPv4 Network configuration data
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IP processing without an explicit address. The type is string.
    Unnumbered interface{}

    // The IP Maximum Transmission Unit. The type is interface{} with range:
    // 68..65535. Units are byte.
    Mtu interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool. The default value is
    // false.
    Unreachables interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool. The default value is
    // true.
    Rpf interface{}
}

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetEntityData() *types.CommonEntityData {
    ipv4Network.EntityData.YFilter = ipv4Network.YFilter
    ipv4Network.EntityData.YangName = "ipv4-network"
    ipv4Network.EntityData.BundleName = "cisco_ios_xr"
    ipv4Network.EntityData.ParentYangName = "subscriber-service"
    ipv4Network.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network"
    ipv4Network.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + ipv4Network.EntityData.SegmentPath
    ipv4Network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Network.EntityData.Children = types.NewOrderedMap()
    ipv4Network.EntityData.Leafs = types.NewOrderedMap()
    ipv4Network.EntityData.Leafs.Append("unnumbered", types.YLeaf{"Unnumbered", ipv4Network.Unnumbered})
    ipv4Network.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", ipv4Network.Mtu})
    ipv4Network.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", ipv4Network.Unreachables})
    ipv4Network.EntityData.Leafs.Append("rpf", types.YLeaf{"Rpf", ipv4Network.Rpf})

    ipv4Network.EntityData.YListKeys = []string {}

    return &(ipv4Network.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network
// Interface IPv6 Network configuration data
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MTU Setting of Interface. The type is interface{} with range: 1280..65535.
    // Units are byte.
    Mtu interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool.
    Rpf interface{}

    // Override Sending of ICMP Unreachable Messages. The type is interface{}.
    Unreachables interface{}

    // Set the IPv6 address of an interface.
    Addresses DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses
}

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetEntityData() *types.CommonEntityData {
    ipv6Network.EntityData.YFilter = ipv6Network.YFilter
    ipv6Network.EntityData.YangName = "ipv6-network"
    ipv6Network.EntityData.BundleName = "cisco_ios_xr"
    ipv6Network.EntityData.ParentYangName = "subscriber-service"
    ipv6Network.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network"
    ipv6Network.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + ipv6Network.EntityData.SegmentPath
    ipv6Network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Network.EntityData.Children = types.NewOrderedMap()
    ipv6Network.EntityData.Children.Append("addresses", types.YChild{"Addresses", &ipv6Network.Addresses})
    ipv6Network.EntityData.Leafs = types.NewOrderedMap()
    ipv6Network.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", ipv6Network.Mtu})
    ipv6Network.EntityData.Leafs.Append("rpf", types.YLeaf{"Rpf", ipv6Network.Rpf})
    ipv6Network.EntityData.Leafs.Append("unreachables", types.YLeaf{"Unreachables", ipv6Network.Unreachables})

    ipv6Network.EntityData.YListKeys = []string {}

    return &(ipv6Network.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses
// Set the IPv6 address of an interface
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto IPv6 Interface Configuration.
    AutoConfiguration DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration
}

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetEntityData() *types.CommonEntityData {
    addresses.EntityData.YFilter = addresses.YFilter
    addresses.EntityData.YangName = "addresses"
    addresses.EntityData.BundleName = "cisco_ios_xr"
    addresses.EntityData.ParentYangName = "ipv6-network"
    addresses.EntityData.SegmentPath = "addresses"
    addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network/" + addresses.EntityData.SegmentPath
    addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresses.EntityData.Children = types.NewOrderedMap()
    addresses.EntityData.Children.Append("auto-configuration", types.YChild{"AutoConfiguration", &addresses.AutoConfiguration})
    addresses.EntityData.Leafs = types.NewOrderedMap()

    addresses.EntityData.YListKeys = []string {}

    return &(addresses.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration
// Auto IPv6 Interface Configuration
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The flag to enable auto ipv6 interface configuration. The type is
    // interface{}.
    Enable interface{}
}

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetEntityData() *types.CommonEntityData {
    autoConfiguration.EntityData.YFilter = autoConfiguration.YFilter
    autoConfiguration.EntityData.YangName = "auto-configuration"
    autoConfiguration.EntityData.BundleName = "cisco_ios_xr"
    autoConfiguration.EntityData.ParentYangName = "addresses"
    autoConfiguration.EntityData.SegmentPath = "auto-configuration"
    autoConfiguration.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network/addresses/" + autoConfiguration.EntityData.SegmentPath
    autoConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoConfiguration.EntityData.Children = types.NewOrderedMap()
    autoConfiguration.EntityData.Leafs = types.NewOrderedMap()
    autoConfiguration.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", autoConfiguration.Enable})

    autoConfiguration.EntityData.YListKeys = []string {}

    return &(autoConfiguration.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor
// Interface IPv6 Network configuration data
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the IPv6 framed ipv6 prefix pool for a subscriber interface . The type
    // is string.
    FramedPrefixPool interface{}

    // Host to use stateful protocol for address configuration. The type is
    // interface{}.
    ManagedConfig interface{}

    // Host to use stateful protocol for non-address configuration. The type is
    // interface{}.
    OtherConfig interface{}

    // Start RA on ipv6-enable config. The type is interface{}.
    StartRaOnIpv6Enable interface{}

    // NUD enable. The type is interface{}.
    NudEnable interface{}

    // Set IPv6 Router Advertisement (RA) lifetime in seconds. The type is
    // interface{} with range: 0..9000. Units are second.
    RaLifetime interface{}

    // RA Router Preference. The type is Ipv6NdRouterPrefTemplate.
    RouterPreference interface{}

    // Enable suppress IPv6 router advertisement. The type is interface{}.
    RaSuppress interface{}

    // Enable RA unicast Flag. The type is interface{}.
    RaUnicast interface{}

    // Unspecify IPv6 Router Advertisement (RA) hop-limit. The type is
    // interface{}.
    RaUnspecifyHoplimit interface{}

    // RA suppress MTU flag. The type is interface{}.
    RaSuppressMtu interface{}

    // Suppress cache learning flag. The type is interface{}.
    SuppressCacheLearning interface{}

    // Set advertised reachability time in milliseconds. The type is interface{}
    // with range: 0..3600000. Units are millisecond.
    ReachableTime interface{}

    // Set advertised NS retransmission interval in milliseconds. The type is
    // interface{} with range: 1000..4294967295. Units are millisecond.
    NsInterval interface{}

    // Set IPv6 Router Advertisement (RA) interval in seconds.
    RaInterval DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval

    // Set the IPv6 framed ipv6 prefix for a subscriber interface .
    FramedPrefix DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix

    // Duplicate Address Detection (DAD).
    DuplicateAddressDetection DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection

    // IPv6 ND RA Initial.
    RaInitial DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial
}

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetEntityData() *types.CommonEntityData {
    ipv6Neighbor.EntityData.YFilter = ipv6Neighbor.YFilter
    ipv6Neighbor.EntityData.YangName = "ipv6-neighbor"
    ipv6Neighbor.EntityData.BundleName = "cisco_ios_xr"
    ipv6Neighbor.EntityData.ParentYangName = "subscriber-service"
    ipv6Neighbor.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor"
    ipv6Neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + ipv6Neighbor.EntityData.SegmentPath
    ipv6Neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Neighbor.EntityData.Children = types.NewOrderedMap()
    ipv6Neighbor.EntityData.Children.Append("ra-interval", types.YChild{"RaInterval", &ipv6Neighbor.RaInterval})
    ipv6Neighbor.EntityData.Children.Append("framed-prefix", types.YChild{"FramedPrefix", &ipv6Neighbor.FramedPrefix})
    ipv6Neighbor.EntityData.Children.Append("duplicate-address-detection", types.YChild{"DuplicateAddressDetection", &ipv6Neighbor.DuplicateAddressDetection})
    ipv6Neighbor.EntityData.Children.Append("ra-initial", types.YChild{"RaInitial", &ipv6Neighbor.RaInitial})
    ipv6Neighbor.EntityData.Leafs = types.NewOrderedMap()
    ipv6Neighbor.EntityData.Leafs.Append("framed-prefix-pool", types.YLeaf{"FramedPrefixPool", ipv6Neighbor.FramedPrefixPool})
    ipv6Neighbor.EntityData.Leafs.Append("managed-config", types.YLeaf{"ManagedConfig", ipv6Neighbor.ManagedConfig})
    ipv6Neighbor.EntityData.Leafs.Append("other-config", types.YLeaf{"OtherConfig", ipv6Neighbor.OtherConfig})
    ipv6Neighbor.EntityData.Leafs.Append("start-ra-on-ipv6-enable", types.YLeaf{"StartRaOnIpv6Enable", ipv6Neighbor.StartRaOnIpv6Enable})
    ipv6Neighbor.EntityData.Leafs.Append("nud-enable", types.YLeaf{"NudEnable", ipv6Neighbor.NudEnable})
    ipv6Neighbor.EntityData.Leafs.Append("ra-lifetime", types.YLeaf{"RaLifetime", ipv6Neighbor.RaLifetime})
    ipv6Neighbor.EntityData.Leafs.Append("router-preference", types.YLeaf{"RouterPreference", ipv6Neighbor.RouterPreference})
    ipv6Neighbor.EntityData.Leafs.Append("ra-suppress", types.YLeaf{"RaSuppress", ipv6Neighbor.RaSuppress})
    ipv6Neighbor.EntityData.Leafs.Append("ra-unicast", types.YLeaf{"RaUnicast", ipv6Neighbor.RaUnicast})
    ipv6Neighbor.EntityData.Leafs.Append("ra-unspecify-hoplimit", types.YLeaf{"RaUnspecifyHoplimit", ipv6Neighbor.RaUnspecifyHoplimit})
    ipv6Neighbor.EntityData.Leafs.Append("ra-suppress-mtu", types.YLeaf{"RaSuppressMtu", ipv6Neighbor.RaSuppressMtu})
    ipv6Neighbor.EntityData.Leafs.Append("suppress-cache-learning", types.YLeaf{"SuppressCacheLearning", ipv6Neighbor.SuppressCacheLearning})
    ipv6Neighbor.EntityData.Leafs.Append("reachable-time", types.YLeaf{"ReachableTime", ipv6Neighbor.ReachableTime})
    ipv6Neighbor.EntityData.Leafs.Append("ns-interval", types.YLeaf{"NsInterval", ipv6Neighbor.NsInterval})

    ipv6Neighbor.EntityData.YListKeys = []string {}

    return &(ipv6Neighbor.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval
// Set IPv6 Router Advertisement (RA) interval in
// seconds
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Maximum RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Maximum interface{}

    // Minimum RA interval in seconds. Must be less than 0.75 * maximum interval.
    // The type is interface{} with range: 3..1800. Units are second.
    Minimum interface{}
}

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetEntityData() *types.CommonEntityData {
    raInterval.EntityData.YFilter = raInterval.YFilter
    raInterval.EntityData.YangName = "ra-interval"
    raInterval.EntityData.BundleName = "cisco_ios_xr"
    raInterval.EntityData.ParentYangName = "ipv6-neighbor"
    raInterval.EntityData.SegmentPath = "ra-interval"
    raInterval.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + raInterval.EntityData.SegmentPath
    raInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    raInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    raInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    raInterval.EntityData.Children = types.NewOrderedMap()
    raInterval.EntityData.Leafs = types.NewOrderedMap()
    raInterval.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", raInterval.Maximum})
    raInterval.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", raInterval.Minimum})

    raInterval.EntityData.YListKeys = []string {}

    return &(raInterval.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix
// Set the IPv6 framed ipv6 prefix for a
// subscriber interface 
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // IPv6 framed prefix length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}

    // IPV6 framed prefix address. The type is string. This attribute is
    // mandatory.
    Prefix interface{}
}

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetEntityData() *types.CommonEntityData {
    framedPrefix.EntityData.YFilter = framedPrefix.YFilter
    framedPrefix.EntityData.YangName = "framed-prefix"
    framedPrefix.EntityData.BundleName = "cisco_ios_xr"
    framedPrefix.EntityData.ParentYangName = "ipv6-neighbor"
    framedPrefix.EntityData.SegmentPath = "framed-prefix"
    framedPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + framedPrefix.EntityData.SegmentPath
    framedPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    framedPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    framedPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    framedPrefix.EntityData.Children = types.NewOrderedMap()
    framedPrefix.EntityData.Leafs = types.NewOrderedMap()
    framedPrefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", framedPrefix.PrefixLength})
    framedPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", framedPrefix.Prefix})

    framedPrefix.EntityData.YListKeys = []string {}

    return &(framedPrefix.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection
// Duplicate Address Detection (DAD)
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set IPv6 duplicate address detection transmits. The type is interface{}
    // with range: 0..600.
    Attempts interface{}
}

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetEntityData() *types.CommonEntityData {
    duplicateAddressDetection.EntityData.YFilter = duplicateAddressDetection.YFilter
    duplicateAddressDetection.EntityData.YangName = "duplicate-address-detection"
    duplicateAddressDetection.EntityData.BundleName = "cisco_ios_xr"
    duplicateAddressDetection.EntityData.ParentYangName = "ipv6-neighbor"
    duplicateAddressDetection.EntityData.SegmentPath = "duplicate-address-detection"
    duplicateAddressDetection.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + duplicateAddressDetection.EntityData.SegmentPath
    duplicateAddressDetection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duplicateAddressDetection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duplicateAddressDetection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duplicateAddressDetection.EntityData.Children = types.NewOrderedMap()
    duplicateAddressDetection.EntityData.Leafs = types.NewOrderedMap()
    duplicateAddressDetection.EntityData.Leafs.Append("attempts", types.YLeaf{"Attempts", duplicateAddressDetection.Attempts})

    duplicateAddressDetection.EntityData.YListKeys = []string {}

    return &(duplicateAddressDetection.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial
// IPv6 ND RA Initial
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Initial RA count. The type is interface{} with range: 0..32. This attribute
    // is mandatory.
    Count interface{}

    // Initial RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Interval interface{}
}

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetEntityData() *types.CommonEntityData {
    raInitial.EntityData.YFilter = raInitial.YFilter
    raInitial.EntityData.YangName = "ra-initial"
    raInitial.EntityData.BundleName = "cisco_ios_xr"
    raInitial.EntityData.ParentYangName = "ipv6-neighbor"
    raInitial.EntityData.SegmentPath = "ra-initial"
    raInitial.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor/" + raInitial.EntityData.SegmentPath
    raInitial.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    raInitial.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    raInitial.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    raInitial.EntityData.Children = types.NewOrderedMap()
    raInitial.EntityData.Leafs = types.NewOrderedMap()
    raInitial.EntityData.Leafs.Append("count", types.YLeaf{"Count", raInitial.Count})
    raInitial.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", raInitial.Interval})

    raInitial.EntityData.YListKeys = []string {}

    return &(raInitial.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_SubscriberAttribute
// Subscriber attribute configuration data
type DynamicTemplate_SubscriberServices_SubscriberService_SubscriberAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OpenDNS configuration data.
    OpenDns DynamicTemplate_SubscriberServices_SubscriberService_SubscriberAttribute_OpenDns
}

func (subscriberAttribute *DynamicTemplate_SubscriberServices_SubscriberService_SubscriberAttribute) GetEntityData() *types.CommonEntityData {
    subscriberAttribute.EntityData.YFilter = subscriberAttribute.YFilter
    subscriberAttribute.EntityData.YangName = "subscriber-attribute"
    subscriberAttribute.EntityData.BundleName = "cisco_ios_xr"
    subscriberAttribute.EntityData.ParentYangName = "subscriber-service"
    subscriberAttribute.EntityData.SegmentPath = "Cisco-IOS-XR-opendns-deviceid-cfg:subscriber-attribute"
    subscriberAttribute.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + subscriberAttribute.EntityData.SegmentPath
    subscriberAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberAttribute.EntityData.Children = types.NewOrderedMap()
    subscriberAttribute.EntityData.Children.Append("open-dns", types.YChild{"OpenDns", &subscriberAttribute.OpenDns})
    subscriberAttribute.EntityData.Leafs = types.NewOrderedMap()

    subscriberAttribute.EntityData.YListKeys = []string {}

    return &(subscriberAttribute.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_SubscriberAttribute_OpenDns
// OpenDNS configuration data
type DynamicTemplate_SubscriberServices_SubscriberService_SubscriberAttribute_OpenDns struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify deviceID to be used for applying OpenDNS policies. The type is
    // string.
    DeviceId interface{}
}

func (openDns *DynamicTemplate_SubscriberServices_SubscriberService_SubscriberAttribute_OpenDns) GetEntityData() *types.CommonEntityData {
    openDns.EntityData.YFilter = openDns.YFilter
    openDns.EntityData.YangName = "open-dns"
    openDns.EntityData.BundleName = "cisco_ios_xr"
    openDns.EntityData.ParentYangName = "subscriber-attribute"
    openDns.EntityData.SegmentPath = "open-dns"
    openDns.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-opendns-deviceid-cfg:subscriber-attribute/" + openDns.EntityData.SegmentPath
    openDns.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    openDns.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    openDns.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    openDns.EntityData.Children = types.NewOrderedMap()
    openDns.EntityData.Leafs = types.NewOrderedMap()
    openDns.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", openDns.DeviceId})

    openDns.EntityData.YListKeys = []string {}

    return &(openDns.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Pbr
// Dynamic Template PBR configuration
type DynamicTemplate_SubscriberServices_SubscriberService_Pbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class for subscriber ingress policy. The type is string.
    ServicePolicyIn interface{}

    // Ingress service policy.
    ServicePolicies DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicies
}

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetEntityData() *types.CommonEntityData {
    pbr.EntityData.YFilter = pbr.YFilter
    pbr.EntityData.YangName = "pbr"
    pbr.EntityData.BundleName = "cisco_ios_xr"
    pbr.EntityData.ParentYangName = "subscriber-service"
    pbr.EntityData.SegmentPath = "Cisco-IOS-XR-pbr-subscriber-cfg:pbr"
    pbr.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + pbr.EntityData.SegmentPath
    pbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbr.EntityData.Children = types.NewOrderedMap()
    pbr.EntityData.Children.Append("service-policies", types.YChild{"ServicePolicies", &pbr.ServicePolicies})
    pbr.EntityData.Leafs = types.NewOrderedMap()
    pbr.EntityData.Leafs.Append("service-policy-in", types.YLeaf{"ServicePolicyIn", pbr.ServicePolicyIn})

    pbr.EntityData.YListKeys = []string {}

    return &(pbr.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicies
// Ingress service policy
type DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service policy details. The type is slice of
    // DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicies_ServicePolicy.
    ServicePolicy []*DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicies_ServicePolicy
}

func (servicePolicies *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicies) GetEntityData() *types.CommonEntityData {
    servicePolicies.EntityData.YFilter = servicePolicies.YFilter
    servicePolicies.EntityData.YangName = "service-policies"
    servicePolicies.EntityData.BundleName = "cisco_ios_xr"
    servicePolicies.EntityData.ParentYangName = "pbr"
    servicePolicies.EntityData.SegmentPath = "service-policies"
    servicePolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-pbr-subscriber-cfg:pbr/" + servicePolicies.EntityData.SegmentPath
    servicePolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicies.EntityData.Children = types.NewOrderedMap()
    servicePolicies.EntityData.Children.Append("service-policy", types.YChild{"ServicePolicy", nil})
    for i := range servicePolicies.ServicePolicy {
        servicePolicies.EntityData.Children.Append(types.GetSegmentPath(servicePolicies.ServicePolicy[i]), types.YChild{"ServicePolicy", servicePolicies.ServicePolicy[i]})
    }
    servicePolicies.EntityData.Leafs = types.NewOrderedMap()

    servicePolicies.EntityData.YListKeys = []string {}

    return &(servicePolicies.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicies_ServicePolicy
// Service policy details
type DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicies_ServicePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of policy-map. The type is string with
    // length: 1..64.
    ServicePolicy interface{}
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicies_ServicePolicy) GetEntityData() *types.CommonEntityData {
    servicePolicy.EntityData.YFilter = servicePolicy.YFilter
    servicePolicy.EntityData.YangName = "service-policy"
    servicePolicy.EntityData.BundleName = "cisco_ios_xr"
    servicePolicy.EntityData.ParentYangName = "service-policies"
    servicePolicy.EntityData.SegmentPath = "service-policy" + types.AddKeyToken(servicePolicy.ServicePolicy, "service-policy")
    servicePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-pbr-subscriber-cfg:pbr/service-policies/" + servicePolicy.EntityData.SegmentPath
    servicePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicy.EntityData.Children = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs.Append("service-policy", types.YLeaf{"ServicePolicy", servicePolicy.ServicePolicy})

    servicePolicy.EntityData.YListKeys = []string {"ServicePolicy"}

    return &(servicePolicy.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Qos
// QoS dynamically applied configuration template
type DynamicTemplate_SubscriberServices_SubscriberService_Qos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service policy to be applied in ingress/egress direction.
    ServicePolicy DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy

    // QoS L2 overhead accounting.
    Account DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account

    // QoS to be applied in egress direction.
    Output DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output
}

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetEntityData() *types.CommonEntityData {
    qos.EntityData.YFilter = qos.YFilter
    qos.EntityData.YangName = "qos"
    qos.EntityData.BundleName = "cisco_ios_xr"
    qos.EntityData.ParentYangName = "subscriber-service"
    qos.EntityData.SegmentPath = "Cisco-IOS-XR-qos-ma-bng-cfg:qos"
    qos.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + qos.EntityData.SegmentPath
    qos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qos.EntityData.Children = types.NewOrderedMap()
    qos.EntityData.Children.Append("service-policy", types.YChild{"ServicePolicy", &qos.ServicePolicy})
    qos.EntityData.Children.Append("account", types.YChild{"Account", &qos.Account})
    qos.EntityData.Children.Append("output", types.YChild{"Output", &qos.Output})
    qos.EntityData.Leafs = types.NewOrderedMap()

    qos.EntityData.YListKeys = []string {}

    return &(qos.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy
// Service policy to be applied in ingress/egress
// direction
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber ingress policy.
    Input DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input

    // Subscriber egress policy.
    Output DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetEntityData() *types.CommonEntityData {
    servicePolicy.EntityData.YFilter = servicePolicy.YFilter
    servicePolicy.EntityData.YangName = "service-policy"
    servicePolicy.EntityData.BundleName = "cisco_ios_xr"
    servicePolicy.EntityData.ParentYangName = "qos"
    servicePolicy.EntityData.SegmentPath = "service-policy"
    servicePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-qos-ma-bng-cfg:qos/" + servicePolicy.EntityData.SegmentPath
    servicePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicy.EntityData.Children = types.NewOrderedMap()
    servicePolicy.EntityData.Children.Append("input", types.YChild{"Input", &servicePolicy.Input})
    servicePolicy.EntityData.Children.Append("output", types.YChild{"Output", &servicePolicy.Output})
    servicePolicy.EntityData.Leafs = types.NewOrderedMap()

    servicePolicy.EntityData.YListKeys = []string {}

    return &(servicePolicy.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input
// Subscriber ingress policy
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Name of policy-map. The type is string. This attribute is mandatory.
    PolicyName interface{}

    // Name of the SPI. The type is string.
    SpiName interface{}

    // TRUE for merge enabled for service-policy applied on dynamic template. The
    // type is bool.
    Merge interface{}

    // Merge ID value. The type is interface{} with range: 0..255.
    MergeId interface{}

    // TRUE for account stats enabled for service-policy applied on dynamic
    // template. Note: account stats not supported for subscriber type 'ppp' and
    // 'ipsubscriber'. The type is bool.
    AccountStats interface{}
}

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "service-policy"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-qos-ma-bng-cfg:qos/service-policy/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", input.PolicyName})
    input.EntityData.Leafs.Append("spi-name", types.YLeaf{"SpiName", input.SpiName})
    input.EntityData.Leafs.Append("merge", types.YLeaf{"Merge", input.Merge})
    input.EntityData.Leafs.Append("merge-id", types.YLeaf{"MergeId", input.MergeId})
    input.EntityData.Leafs.Append("account-stats", types.YLeaf{"AccountStats", input.AccountStats})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output
// Subscriber egress policy
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Name of policy-map. The type is string. This attribute is mandatory.
    PolicyName interface{}

    // Name of the SPI. The type is string.
    SpiName interface{}

    // TRUE for merge enabled for service-policy applied on dynamic template. The
    // type is bool.
    Merge interface{}

    // Merge ID value. The type is interface{} with range: 0..255.
    MergeId interface{}

    // TRUE for account stats enabled for service-policy applied on dynamic
    // template. Note: account stats not supported for subscriber type 'ppp' and
    // 'ipsubscriber'. The type is bool.
    AccountStats interface{}
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "service-policy"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-qos-ma-bng-cfg:qos/service-policy/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", output.PolicyName})
    output.EntityData.Leafs.Append("spi-name", types.YLeaf{"SpiName", output.SpiName})
    output.EntityData.Leafs.Append("merge", types.YLeaf{"Merge", output.Merge})
    output.EntityData.Leafs.Append("merge-id", types.YLeaf{"MergeId", output.MergeId})
    output.EntityData.Leafs.Append("account-stats", types.YLeaf{"AccountStats", output.AccountStats})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account
// QoS L2 overhead accounting
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ATM adaptation layer AAL. The type is Qosl2DataLink.
    Aal interface{}

    // Specify encapsulation type. The type is Qosl2Encap.
    Encapsulation interface{}

    // ATM cell tax to L2 overhead. The type is interface{}.
    AtmCellTax interface{}

    // Numeric L2 overhead offset. The type is interface{} with range: -63..63.
    UserDefined interface{}
}

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetEntityData() *types.CommonEntityData {
    account.EntityData.YFilter = account.YFilter
    account.EntityData.YangName = "account"
    account.EntityData.BundleName = "cisco_ios_xr"
    account.EntityData.ParentYangName = "qos"
    account.EntityData.SegmentPath = "account"
    account.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-qos-ma-bng-cfg:qos/" + account.EntityData.SegmentPath
    account.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    account.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    account.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    account.EntityData.Children = types.NewOrderedMap()
    account.EntityData.Leafs = types.NewOrderedMap()
    account.EntityData.Leafs.Append("aal", types.YLeaf{"Aal", account.Aal})
    account.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", account.Encapsulation})
    account.EntityData.Leafs.Append("atm-cell-tax", types.YLeaf{"AtmCellTax", account.AtmCellTax})
    account.EntityData.Leafs.Append("user-defined", types.YLeaf{"UserDefined", account.UserDefined})

    account.EntityData.YListKeys = []string {}

    return &(account.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output
// QoS to be applied in egress direction
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth value for the subscriber (in kbps). The type is
    // interface{} with range: 1..4294967295. Units are kbit/s.
    MinimumBandwidth interface{}
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "qos"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-qos-ma-bng-cfg:qos/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("minimum-bandwidth", types.YLeaf{"MinimumBandwidth", output.MinimumBandwidth})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Accounting
// Subscriber accounting dynamic-template commands
type DynamicTemplate_SubscriberServices_SubscriberService_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber monitor feature. The type is string.
    MonitorFeature interface{}

    // Subscriber accounting prepaid feature. The type is string.
    PrepaidFeature interface{}

    // Subscriber accounting service accounting.
    ServiceAccounting DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting

    // Subscriber accounting session accounting.
    Session DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session

    // Subscriber accounting idle timeout.
    IdleTimeout DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout
}

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "subscriber-service"
    accounting.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-accounting-cfg:accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("service-accounting", types.YChild{"ServiceAccounting", &accounting.ServiceAccounting})
    accounting.EntityData.Children.Append("session", types.YChild{"Session", &accounting.Session})
    accounting.EntityData.Children.Append("idle-timeout", types.YChild{"IdleTimeout", &accounting.IdleTimeout})
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("monitor-feature", types.YLeaf{"MonitorFeature", accounting.MonitorFeature})
    accounting.EntityData.Leafs.Append("prepaid-feature", types.YLeaf{"PrepaidFeature", accounting.PrepaidFeature})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting
// Subscriber accounting service accounting
type DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service accounting method list name. The type is string.
    MethodListName interface{}

    // Accounting interim interval in minutes. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    AccountingInterimInterval interface{}
}

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetEntityData() *types.CommonEntityData {
    serviceAccounting.EntityData.YFilter = serviceAccounting.YFilter
    serviceAccounting.EntityData.YangName = "service-accounting"
    serviceAccounting.EntityData.BundleName = "cisco_ios_xr"
    serviceAccounting.EntityData.ParentYangName = "accounting"
    serviceAccounting.EntityData.SegmentPath = "service-accounting"
    serviceAccounting.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-subscriber-accounting-cfg:accounting/" + serviceAccounting.EntityData.SegmentPath
    serviceAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceAccounting.EntityData.Children = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs.Append("method-list-name", types.YLeaf{"MethodListName", serviceAccounting.MethodListName})
    serviceAccounting.EntityData.Leafs.Append("accounting-interim-interval", types.YLeaf{"AccountingInterimInterval", serviceAccounting.AccountingInterimInterval})

    serviceAccounting.EntityData.YListKeys = []string {}

    return &(serviceAccounting.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session
// Subscriber accounting session accounting
type DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session accounting method list name. The type is string.
    MethodListName interface{}

    // Interim accounting interval in minutes. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    PeriodicInterval interface{}

    // Dual stack wait delay in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    DualStackDelay interface{}

    // Hold Accounting start based on IA_PD. The type is HoldAcctStart.
    HoldAcctStart interface{}
}

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "accounting"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-subscriber-accounting-cfg:accounting/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("method-list-name", types.YLeaf{"MethodListName", session.MethodListName})
    session.EntityData.Leafs.Append("periodic-interval", types.YLeaf{"PeriodicInterval", session.PeriodicInterval})
    session.EntityData.Leafs.Append("dual-stack-delay", types.YLeaf{"DualStackDelay", session.DualStackDelay})
    session.EntityData.Leafs.Append("hold-acct-start", types.YLeaf{"HoldAcctStart", session.HoldAcctStart})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session_HoldAcctStart represents Hold Accounting start based on IA_PD
type DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session_HoldAcctStart string

const (
    // Based on ipv6 delegated prefix
    DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session_HoldAcctStart_ipv6_prefix_delegation DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session_HoldAcctStart = "ipv6-prefix-delegation"
)

// DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout
// Subscriber accounting idle timeout
type DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Idle timeout value in seconds. The type is interface{} with range:
    // 60..4320000. Units are second.
    TimeoutValue interface{}

    // Threshold in minute(s) per packet. The type is interface{} with range:
    // 1..10000.
    Threshold interface{}

    // Idle timeout traffic direction. The type is string.
    Direction interface{}
}

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetEntityData() *types.CommonEntityData {
    idleTimeout.EntityData.YFilter = idleTimeout.YFilter
    idleTimeout.EntityData.YangName = "idle-timeout"
    idleTimeout.EntityData.BundleName = "cisco_ios_xr"
    idleTimeout.EntityData.ParentYangName = "accounting"
    idleTimeout.EntityData.SegmentPath = "idle-timeout"
    idleTimeout.EntityData.AbsolutePath = "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template/subscriber-services/subscriber-service/Cisco-IOS-XR-subscriber-accounting-cfg:accounting/" + idleTimeout.EntityData.SegmentPath
    idleTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    idleTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    idleTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    idleTimeout.EntityData.Children = types.NewOrderedMap()
    idleTimeout.EntityData.Leafs = types.NewOrderedMap()
    idleTimeout.EntityData.Leafs.Append("timeout-value", types.YLeaf{"TimeoutValue", idleTimeout.TimeoutValue})
    idleTimeout.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", idleTimeout.Threshold})
    idleTimeout.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", idleTimeout.Direction})

    idleTimeout.EntityData.YListKeys = []string {}

    return &(idleTimeout.EntityData)
}

