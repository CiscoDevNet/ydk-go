// This module contains a collection of YANG definitions
// for Cisco IOS-XR subscriber-infra-tmplmgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   dynamic-template: All dynamic template configurations
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Templates of the PPP Type.
    Ppps DynamicTemplate_Ppps

    // The IP Subscriber Template Table.
    IpSubscribers DynamicTemplate_IpSubscribers

    // The Service Type Template Table.
    SubscriberServices DynamicTemplate_SubscriberServices
}

func (dynamicTemplate *DynamicTemplate) GetFilter() yfilter.YFilter { return dynamicTemplate.YFilter }

func (dynamicTemplate *DynamicTemplate) SetFilter(yf yfilter.YFilter) { dynamicTemplate.YFilter = yf }

func (dynamicTemplate *DynamicTemplate) GetGoName(yname string) string {
    if yname == "ppps" { return "Ppps" }
    if yname == "ip-subscribers" { return "IpSubscribers" }
    if yname == "subscriber-services" { return "SubscriberServices" }
    return ""
}

func (dynamicTemplate *DynamicTemplate) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg:dynamic-template"
}

func (dynamicTemplate *DynamicTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ppps" {
        return &dynamicTemplate.Ppps
    }
    if childYangName == "ip-subscribers" {
        return &dynamicTemplate.IpSubscribers
    }
    if childYangName == "subscriber-services" {
        return &dynamicTemplate.SubscriberServices
    }
    return nil
}

func (dynamicTemplate *DynamicTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ppps"] = &dynamicTemplate.Ppps
    children["ip-subscribers"] = &dynamicTemplate.IpSubscribers
    children["subscriber-services"] = &dynamicTemplate.SubscriberServices
    return children
}

func (dynamicTemplate *DynamicTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dynamicTemplate *DynamicTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicTemplate *DynamicTemplate) GetYangName() string { return "dynamic-template" }

func (dynamicTemplate *DynamicTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicTemplate *DynamicTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicTemplate *DynamicTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicTemplate *DynamicTemplate) SetParent(parent types.Entity) { dynamicTemplate.parent = parent }

func (dynamicTemplate *DynamicTemplate) GetParent() types.Entity { return dynamicTemplate.parent }

func (dynamicTemplate *DynamicTemplate) GetParentYangName() string { return "Cisco-IOS-XR-subscriber-infra-tmplmgr-cfg" }

// DynamicTemplate_Ppps
// Templates of the PPP Type
type DynamicTemplate_Ppps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Template of the PPP Type. The type is slice of DynamicTemplate_Ppps_Ppp.
    Ppp []DynamicTemplate_Ppps_Ppp
}

func (ppps *DynamicTemplate_Ppps) GetFilter() yfilter.YFilter { return ppps.YFilter }

func (ppps *DynamicTemplate_Ppps) SetFilter(yf yfilter.YFilter) { ppps.YFilter = yf }

func (ppps *DynamicTemplate_Ppps) GetGoName(yname string) string {
    if yname == "ppp" { return "Ppp" }
    return ""
}

func (ppps *DynamicTemplate_Ppps) GetSegmentPath() string {
    return "ppps"
}

func (ppps *DynamicTemplate_Ppps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ppp" {
        for _, c := range ppps.Ppp {
            if ppps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DynamicTemplate_Ppps_Ppp{}
        ppps.Ppp = append(ppps.Ppp, child)
        return &ppps.Ppp[len(ppps.Ppp)-1]
    }
    return nil
}

func (ppps *DynamicTemplate_Ppps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ppps.Ppp {
        children[ppps.Ppp[i].GetSegmentPath()] = &ppps.Ppp[i]
    }
    return children
}

func (ppps *DynamicTemplate_Ppps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ppps *DynamicTemplate_Ppps) GetBundleName() string { return "cisco_ios_xr" }

func (ppps *DynamicTemplate_Ppps) GetYangName() string { return "ppps" }

func (ppps *DynamicTemplate_Ppps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ppps *DynamicTemplate_Ppps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ppps *DynamicTemplate_Ppps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ppps *DynamicTemplate_Ppps) SetParent(parent types.Entity) { ppps.parent = parent }

func (ppps *DynamicTemplate_Ppps) GetParent() types.Entity { return ppps.parent }

func (ppps *DynamicTemplate_Ppps) GetParentYangName() string { return "dynamic-template" }

// DynamicTemplate_Ppps_Ppp
// A Template of the PPP Type
type DynamicTemplate_Ppps_Ppp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the template. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (ppp *DynamicTemplate_Ppps_Ppp) GetFilter() yfilter.YFilter { return ppp.YFilter }

func (ppp *DynamicTemplate_Ppps_Ppp) SetFilter(yf yfilter.YFilter) { ppp.YFilter = yf }

func (ppp *DynamicTemplate_Ppps_Ppp) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "vrf" { return "Vrf" }
    if yname == "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions" { return "SpanMonitorSessions" }
    if yname == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter" { return "Ipv4PacketFilter" }
    if yname == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter" { return "Ipv6PacketFilter" }
    if yname == "Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp" { return "Igmp" }
    if yname == "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network" { return "Ipv4Network" }
    if yname == "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network" { return "Ipv6Network" }
    if yname == "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor" { return "Ipv6Neighbor" }
    if yname == "Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6" { return "Dhcpv6" }
    if yname == "Cisco-IOS-XR-pbr-subscriber-cfg:pbr" { return "Pbr" }
    if yname == "Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template" { return "PppTemplate" }
    if yname == "Cisco-IOS-XR-qos-ma-bng-cfg:qos" { return "Qos" }
    if yname == "Cisco-IOS-XR-subscriber-accounting-cfg:accounting" { return "Accounting" }
    if yname == "Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-template" { return "PppoeTemplate" }
    return ""
}

func (ppp *DynamicTemplate_Ppps_Ppp) GetSegmentPath() string {
    return "ppp" + "[template-name='" + fmt.Sprintf("%v", ppp.TemplateName) + "']"
}

func (ppp *DynamicTemplate_Ppps_Ppp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions" {
        return &ppp.SpanMonitorSessions
    }
    if childYangName == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter" {
        return &ppp.Ipv4PacketFilter
    }
    if childYangName == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter" {
        return &ppp.Ipv6PacketFilter
    }
    if childYangName == "Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp" {
        return &ppp.Igmp
    }
    if childYangName == "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network" {
        return &ppp.Ipv4Network
    }
    if childYangName == "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network" {
        return &ppp.Ipv6Network
    }
    if childYangName == "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor" {
        return &ppp.Ipv6Neighbor
    }
    if childYangName == "Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6" {
        return &ppp.Dhcpv6
    }
    if childYangName == "Cisco-IOS-XR-pbr-subscriber-cfg:pbr" {
        return &ppp.Pbr
    }
    if childYangName == "Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template" {
        return &ppp.PppTemplate
    }
    if childYangName == "Cisco-IOS-XR-qos-ma-bng-cfg:qos" {
        return &ppp.Qos
    }
    if childYangName == "Cisco-IOS-XR-subscriber-accounting-cfg:accounting" {
        return &ppp.Accounting
    }
    if childYangName == "Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-template" {
        return &ppp.PppoeTemplate
    }
    return nil
}

func (ppp *DynamicTemplate_Ppps_Ppp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions"] = &ppp.SpanMonitorSessions
    children["Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter"] = &ppp.Ipv4PacketFilter
    children["Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter"] = &ppp.Ipv6PacketFilter
    children["Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp"] = &ppp.Igmp
    children["Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network"] = &ppp.Ipv4Network
    children["Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network"] = &ppp.Ipv6Network
    children["Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor"] = &ppp.Ipv6Neighbor
    children["Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6"] = &ppp.Dhcpv6
    children["Cisco-IOS-XR-pbr-subscriber-cfg:pbr"] = &ppp.Pbr
    children["Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template"] = &ppp.PppTemplate
    children["Cisco-IOS-XR-qos-ma-bng-cfg:qos"] = &ppp.Qos
    children["Cisco-IOS-XR-subscriber-accounting-cfg:accounting"] = &ppp.Accounting
    children["Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-template"] = &ppp.PppoeTemplate
    return children
}

func (ppp *DynamicTemplate_Ppps_Ppp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ppp.TemplateName
    leafs["vrf"] = ppp.Vrf
    return leafs
}

func (ppp *DynamicTemplate_Ppps_Ppp) GetBundleName() string { return "cisco_ios_xr" }

func (ppp *DynamicTemplate_Ppps_Ppp) GetYangName() string { return "ppp" }

func (ppp *DynamicTemplate_Ppps_Ppp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ppp *DynamicTemplate_Ppps_Ppp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ppp *DynamicTemplate_Ppps_Ppp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ppp *DynamicTemplate_Ppps_Ppp) SetParent(parent types.Entity) { ppp.parent = parent }

func (ppp *DynamicTemplate_Ppps_Ppp) GetParent() types.Entity { return ppp.parent }

func (ppp *DynamicTemplate_Ppps_Ppp) GetParentYangName() string { return "ppps" }

// DynamicTemplate_Ppps_Ppp_SpanMonitorSessions
// Monitor Session container for this template
type DynamicTemplate_Ppps_Ppp_SpanMonitorSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a particular class of Monitor Session. The type is slice
    // of DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession.
    SpanMonitorSession []DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession
}

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetFilter() yfilter.YFilter { return spanMonitorSessions.YFilter }

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) SetFilter(yf yfilter.YFilter) { spanMonitorSessions.YFilter = yf }

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetGoName(yname string) string {
    if yname == "span-monitor-session" { return "SpanMonitorSession" }
    return ""
}

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetSegmentPath() string {
    return "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions"
}

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "span-monitor-session" {
        for _, c := range spanMonitorSessions.SpanMonitorSession {
            if spanMonitorSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession{}
        spanMonitorSessions.SpanMonitorSession = append(spanMonitorSessions.SpanMonitorSession, child)
        return &spanMonitorSessions.SpanMonitorSession[len(spanMonitorSessions.SpanMonitorSession)-1]
    }
    return nil
}

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spanMonitorSessions.SpanMonitorSession {
        children[spanMonitorSessions.SpanMonitorSession[i].GetSegmentPath()] = &spanMonitorSessions.SpanMonitorSession[i]
    }
    return children
}

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetBundleName() string { return "cisco_ios_xr" }

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetYangName() string { return "span-monitor-sessions" }

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) SetParent(parent types.Entity) { spanMonitorSessions.parent = parent }

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetParent() types.Entity { return spanMonitorSessions.parent }

func (spanMonitorSessions *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession
// Configuration for a particular class of Monitor
// Session
type DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetFilter() yfilter.YFilter { return spanMonitorSession.YFilter }

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) SetFilter(yf yfilter.YFilter) { spanMonitorSession.YFilter = yf }

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetGoName(yname string) string {
    if yname == "session-class" { return "SessionClass" }
    if yname == "mirror-first" { return "MirrorFirst" }
    if yname == "mirror-interval" { return "MirrorInterval" }
    if yname == "attachment" { return "Attachment" }
    if yname == "acl" { return "Acl" }
    return ""
}

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetSegmentPath() string {
    return "span-monitor-session" + "[session-class='" + fmt.Sprintf("%v", spanMonitorSession.SessionClass) + "']"
}

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attachment" {
        return &spanMonitorSession.Attachment
    }
    if childYangName == "acl" {
        return &spanMonitorSession.Acl
    }
    return nil
}

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attachment"] = &spanMonitorSession.Attachment
    children["acl"] = &spanMonitorSession.Acl
    return children
}

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-class"] = spanMonitorSession.SessionClass
    leafs["mirror-first"] = spanMonitorSession.MirrorFirst
    leafs["mirror-interval"] = spanMonitorSession.MirrorInterval
    return leafs
}

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetBundleName() string { return "cisco_ios_xr" }

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetYangName() string { return "span-monitor-session" }

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) SetParent(parent types.Entity) { spanMonitorSession.parent = parent }

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetParent() types.Entity { return spanMonitorSession.parent }

func (spanMonitorSession *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession) GetParentYangName() string { return "span-monitor-sessions" }

// DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment
// Attach the interface to a Monitor Session
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session Name. The type is string with length: 1..79. This attribute is
    // mandatory.
    SessionName interface{}

    // Specify the direction of traffic to replicate (optional). The type is
    // SpanTrafficDirection.
    Direction interface{}

    // Enable port level traffic mirroring. The type is interface{}.
    PortLevelEnable interface{}
}

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetFilter() yfilter.YFilter { return attachment.YFilter }

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) SetFilter(yf yfilter.YFilter) { attachment.YFilter = yf }

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetGoName(yname string) string {
    if yname == "session-name" { return "SessionName" }
    if yname == "direction" { return "Direction" }
    if yname == "port-level-enable" { return "PortLevelEnable" }
    return ""
}

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetSegmentPath() string {
    return "attachment"
}

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-name"] = attachment.SessionName
    leafs["direction"] = attachment.Direction
    leafs["port-level-enable"] = attachment.PortLevelEnable
    return leafs
}

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetBundleName() string { return "cisco_ios_xr" }

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetYangName() string { return "attachment" }

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) SetParent(parent types.Entity) { attachment.parent = parent }

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetParent() types.Entity { return attachment.parent }

func (attachment *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Attachment) GetParentYangName() string { return "span-monitor-session" }

// DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl
// Enable ACL matching for traffic mirroring
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ACL. The type is interface{}. This attribute is mandatory.
    AclEnable interface{}

    // ACL Name. The type is string with length: 1..80.
    AclName interface{}
}

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetFilter() yfilter.YFilter { return acl.YFilter }

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) SetFilter(yf yfilter.YFilter) { acl.YFilter = yf }

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetGoName(yname string) string {
    if yname == "acl-enable" { return "AclEnable" }
    if yname == "acl-name" { return "AclName" }
    return ""
}

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetSegmentPath() string {
    return "acl"
}

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["acl-enable"] = acl.AclEnable
    leafs["acl-name"] = acl.AclName
    return leafs
}

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetBundleName() string { return "cisco_ios_xr" }

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetYangName() string { return "acl" }

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) SetParent(parent types.Entity) { acl.parent = parent }

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetParent() types.Entity { return acl.parent }

func (acl *DynamicTemplate_Ppps_Ppp_SpanMonitorSessions_SpanMonitorSession_Acl) GetParentYangName() string { return "span-monitor-session" }

// DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter
// IPv4 Packet Filtering configuration for the
// template
type DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound

    // IPv4 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound
}

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetFilter() yfilter.YFilter { return ipv4PacketFilter.YFilter }

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) SetFilter(yf yfilter.YFilter) { ipv4PacketFilter.YFilter = yf }

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetGoName(yname string) string {
    if yname == "outbound" { return "Outbound" }
    if yname == "inbound" { return "Inbound" }
    return ""
}

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter"
}

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "outbound" {
        return &ipv4PacketFilter.Outbound
    }
    if childYangName == "inbound" {
        return &ipv4PacketFilter.Inbound
    }
    return nil
}

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["outbound"] = &ipv4PacketFilter.Outbound
    children["inbound"] = &ipv4PacketFilter.Inbound
    return children
}

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetYangName() string { return "ipv4-packet-filter" }

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) SetParent(parent types.Entity) { ipv4PacketFilter.parent = parent }

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetParent() types.Entity { return ipv4PacketFilter.parent }

func (ipv4PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound
// IPv4 Packet filter to be applied to outbound
// packets
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..65. This attribute is mandatory.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetFilter() yfilter.YFilter { return outbound.YFilter }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) SetFilter(yf yfilter.YFilter) { outbound.YFilter = yf }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "hardware-count" { return "HardwareCount" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetSegmentPath() string {
    return "outbound"
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = outbound.CommonAclName
    leafs["name"] = outbound.Name
    leafs["hardware-count"] = outbound.HardwareCount
    leafs["interface-statistics"] = outbound.InterfaceStatistics
    return leafs
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetBundleName() string { return "cisco_ios_xr" }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetYangName() string { return "outbound" }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) SetParent(parent types.Entity) { outbound.parent = parent }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetParent() types.Entity { return outbound.parent }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Outbound) GetParentYangName() string { return "ipv4-packet-filter" }

// DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound
// IPv4 Packet filter to be applied to inbound
// packets
type DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Inbound packets NOTE: This
    // parameter is mandatory if 'CommonACLName' is not specified. The type is
    // string with length: 1..65.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetFilter() yfilter.YFilter { return inbound.YFilter }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) SetFilter(yf yfilter.YFilter) { inbound.YFilter = yf }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "hardware-count" { return "HardwareCount" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetSegmentPath() string {
    return "inbound"
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = inbound.CommonAclName
    leafs["name"] = inbound.Name
    leafs["hardware-count"] = inbound.HardwareCount
    leafs["interface-statistics"] = inbound.InterfaceStatistics
    return leafs
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetBundleName() string { return "cisco_ios_xr" }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetYangName() string { return "inbound" }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) SetParent(parent types.Entity) { inbound.parent = parent }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetParent() types.Entity { return inbound.parent }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv4PacketFilter_Inbound) GetParentYangName() string { return "ipv4-packet-filter" }

// DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter
// IPv6 Packet Filtering configuration for the
// interface
type DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound

    // IPv6 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound
}

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetFilter() yfilter.YFilter { return ipv6PacketFilter.YFilter }

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) SetFilter(yf yfilter.YFilter) { ipv6PacketFilter.YFilter = yf }

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetGoName(yname string) string {
    if yname == "inbound" { return "Inbound" }
    if yname == "outbound" { return "Outbound" }
    return ""
}

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter"
}

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inbound" {
        return &ipv6PacketFilter.Inbound
    }
    if childYangName == "outbound" {
        return &ipv6PacketFilter.Outbound
    }
    return nil
}

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["inbound"] = &ipv6PacketFilter.Inbound
    children["outbound"] = &ipv6PacketFilter.Outbound
    return children
}

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetYangName() string { return "ipv6-packet-filter" }

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) SetParent(parent types.Entity) { ipv6PacketFilter.parent = parent }

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetParent() types.Entity { return ipv6PacketFilter.parent }

func (ipv6PacketFilter *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound
// IPv6 Packet filter to be applied to inbound
// packets
type DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Inbound  NOTE: This parameter is
    // mandatory if 'CommonACLName' is not specified. The type is string with
    // length: 1..65.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetFilter() yfilter.YFilter { return inbound.YFilter }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) SetFilter(yf yfilter.YFilter) { inbound.YFilter = yf }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetSegmentPath() string {
    return "inbound"
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = inbound.CommonAclName
    leafs["name"] = inbound.Name
    leafs["interface-statistics"] = inbound.InterfaceStatistics
    return leafs
}

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetBundleName() string { return "cisco_ios_xr" }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetYangName() string { return "inbound" }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) SetParent(parent types.Entity) { inbound.parent = parent }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetParent() types.Entity { return inbound.parent }

func (inbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Inbound) GetParentYangName() string { return "ipv6-packet-filter" }

// DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound
// IPv6 Packet filter to be applied to outbound
// packets
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..65. This attribute is mandatory.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetFilter() yfilter.YFilter { return outbound.YFilter }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) SetFilter(yf yfilter.YFilter) { outbound.YFilter = yf }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetSegmentPath() string {
    return "outbound"
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = outbound.CommonAclName
    leafs["name"] = outbound.Name
    leafs["interface-statistics"] = outbound.InterfaceStatistics
    return leafs
}

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetBundleName() string { return "cisco_ios_xr" }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetYangName() string { return "outbound" }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) SetParent(parent types.Entity) { outbound.parent = parent }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetParent() types.Entity { return outbound.parent }

func (outbound *DynamicTemplate_Ppps_Ppp_Ipv6PacketFilter_Outbound) GetParentYangName() string { return "ipv6-packet-filter" }

// DynamicTemplate_Ppps_Ppp_Igmp
// IGMPconfiguration
type DynamicTemplate_Ppps_Ppp_Igmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default VRF.
    DefaultVrf DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf
}

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetFilter() yfilter.YFilter { return igmp.YFilter }

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) SetFilter(yf yfilter.YFilter) { igmp.YFilter = yf }

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetGoName(yname string) string {
    if yname == "default-vrf" { return "DefaultVrf" }
    return ""
}

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp"
}

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-vrf" {
        return &igmp.DefaultVrf
    }
    return nil
}

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-vrf"] = &igmp.DefaultVrf
    return children
}

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetBundleName() string { return "cisco_ios_xr" }

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetYangName() string { return "igmp" }

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) SetParent(parent types.Entity) { igmp.parent = parent }

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetParent() types.Entity { return igmp.parent }

func (igmp *DynamicTemplate_Ppps_Ppp_Igmp) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf
// Default VRF
type DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf struct {
    parent types.Entity
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

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Configure Multicast mode variable. The type is DynTmplMulticastMode.
    MulticastMode interface{}

    // IGMPv3 explicit host tracking.
    ExplicitTracking DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking
}

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetFilter() yfilter.YFilter { return defaultVrf.YFilter }

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) SetFilter(yf yfilter.YFilter) { defaultVrf.YFilter = yf }

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetGoName(yname string) string {
    if yname == "max-groups" { return "MaxGroups" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "version" { return "Version" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "multicast-mode" { return "MulticastMode" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetSegmentPath() string {
    return "default-vrf"
}

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "explicit-tracking" {
        return &defaultVrf.ExplicitTracking
    }
    return nil
}

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["explicit-tracking"] = &defaultVrf.ExplicitTracking
    return children
}

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-groups"] = defaultVrf.MaxGroups
    leafs["access-group"] = defaultVrf.AccessGroup
    leafs["version"] = defaultVrf.Version
    leafs["query-interval"] = defaultVrf.QueryInterval
    leafs["query-max-response-time"] = defaultVrf.QueryMaxResponseTime
    leafs["multicast-mode"] = defaultVrf.MulticastMode
    return leafs
}

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetBundleName() string { return "cisco_ios_xr" }

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetYangName() string { return "default-vrf" }

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) SetParent(parent types.Entity) { defaultVrf.parent = parent }

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetParent() types.Entity { return defaultVrf.parent }

func (defaultVrf *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf) GetParentYangName() string { return "igmp" }

// DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable or disable, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *DynamicTemplate_Ppps_Ppp_Igmp_DefaultVrf_ExplicitTracking) GetParentYangName() string { return "default-vrf" }

// DynamicTemplate_Ppps_Ppp_Ipv4Network
// Interface IPv4 Network configuration data
type DynamicTemplate_Ppps_Ppp_Ipv4Network struct {
    parent types.Entity
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

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetFilter() yfilter.YFilter { return ipv4Network.YFilter }

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) SetFilter(yf yfilter.YFilter) { ipv4Network.YFilter = yf }

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetGoName(yname string) string {
    if yname == "unnumbered" { return "Unnumbered" }
    if yname == "mtu" { return "Mtu" }
    if yname == "unreachables" { return "Unreachables" }
    if yname == "rpf" { return "Rpf" }
    return ""
}

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network"
}

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unnumbered"] = ipv4Network.Unnumbered
    leafs["mtu"] = ipv4Network.Mtu
    leafs["unreachables"] = ipv4Network.Unreachables
    leafs["rpf"] = ipv4Network.Rpf
    return leafs
}

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetYangName() string { return "ipv4-network" }

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) SetParent(parent types.Entity) { ipv4Network.parent = parent }

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetParent() types.Entity { return ipv4Network.parent }

func (ipv4Network *DynamicTemplate_Ppps_Ppp_Ipv4Network) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Ipv6Network
// Interface IPv6 Network configuration data
type DynamicTemplate_Ppps_Ppp_Ipv6Network struct {
    parent types.Entity
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

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetFilter() yfilter.YFilter { return ipv6Network.YFilter }

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) SetFilter(yf yfilter.YFilter) { ipv6Network.YFilter = yf }

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetGoName(yname string) string {
    if yname == "mtu" { return "Mtu" }
    if yname == "rpf" { return "Rpf" }
    if yname == "unreachables" { return "Unreachables" }
    if yname == "addresses" { return "Addresses" }
    return ""
}

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network"
}

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &ipv6Network.Addresses
    }
    return nil
}

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &ipv6Network.Addresses
    return children
}

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mtu"] = ipv6Network.Mtu
    leafs["rpf"] = ipv6Network.Rpf
    leafs["unreachables"] = ipv6Network.Unreachables
    return leafs
}

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetYangName() string { return "ipv6-network" }

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) SetParent(parent types.Entity) { ipv6Network.parent = parent }

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetParent() types.Entity { return ipv6Network.parent }

func (ipv6Network *DynamicTemplate_Ppps_Ppp_Ipv6Network) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses
// Set the IPv6 address of an interface
type DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Auto IPv6 Interface Configuration.
    AutoConfiguration DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration
}

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetGoName(yname string) string {
    if yname == "auto-configuration" { return "AutoConfiguration" }
    return ""
}

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto-configuration" {
        return &addresses.AutoConfiguration
    }
    return nil
}

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto-configuration"] = &addresses.AutoConfiguration
    return children
}

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetYangName() string { return "addresses" }

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses) GetParentYangName() string { return "ipv6-network" }

// DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration
// Auto IPv6 Interface Configuration
type DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The flag to enable auto ipv6 interface configuration. The type is
    // interface{}.
    Enable interface{}
}

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetFilter() yfilter.YFilter { return autoConfiguration.YFilter }

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) SetFilter(yf yfilter.YFilter) { autoConfiguration.YFilter = yf }

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetSegmentPath() string {
    return "auto-configuration"
}

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = autoConfiguration.Enable
    return leafs
}

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetYangName() string { return "auto-configuration" }

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) SetParent(parent types.Entity) { autoConfiguration.parent = parent }

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetParent() types.Entity { return autoConfiguration.parent }

func (autoConfiguration *DynamicTemplate_Ppps_Ppp_Ipv6Network_Addresses_AutoConfiguration) GetParentYangName() string { return "addresses" }

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor
// Interface IPv6 Network configuration data
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor struct {
    parent types.Entity
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

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetFilter() yfilter.YFilter { return ipv6Neighbor.YFilter }

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) SetFilter(yf yfilter.YFilter) { ipv6Neighbor.YFilter = yf }

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetGoName(yname string) string {
    if yname == "framed-prefix-pool" { return "FramedPrefixPool" }
    if yname == "managed-config" { return "ManagedConfig" }
    if yname == "other-config" { return "OtherConfig" }
    if yname == "start-ra-on-ipv6-enable" { return "StartRaOnIpv6Enable" }
    if yname == "nud-enable" { return "NudEnable" }
    if yname == "ra-lifetime" { return "RaLifetime" }
    if yname == "router-preference" { return "RouterPreference" }
    if yname == "ra-suppress" { return "RaSuppress" }
    if yname == "ra-unicast" { return "RaUnicast" }
    if yname == "ra-unspecify-hoplimit" { return "RaUnspecifyHoplimit" }
    if yname == "ra-suppress-mtu" { return "RaSuppressMtu" }
    if yname == "suppress-cache-learning" { return "SuppressCacheLearning" }
    if yname == "reachable-time" { return "ReachableTime" }
    if yname == "ns-interval" { return "NsInterval" }
    if yname == "ra-interval" { return "RaInterval" }
    if yname == "framed-prefix" { return "FramedPrefix" }
    if yname == "duplicate-address-detection" { return "DuplicateAddressDetection" }
    if yname == "ra-initial" { return "RaInitial" }
    return ""
}

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor"
}

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ra-interval" {
        return &ipv6Neighbor.RaInterval
    }
    if childYangName == "framed-prefix" {
        return &ipv6Neighbor.FramedPrefix
    }
    if childYangName == "duplicate-address-detection" {
        return &ipv6Neighbor.DuplicateAddressDetection
    }
    if childYangName == "ra-initial" {
        return &ipv6Neighbor.RaInitial
    }
    return nil
}

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ra-interval"] = &ipv6Neighbor.RaInterval
    children["framed-prefix"] = &ipv6Neighbor.FramedPrefix
    children["duplicate-address-detection"] = &ipv6Neighbor.DuplicateAddressDetection
    children["ra-initial"] = &ipv6Neighbor.RaInitial
    return children
}

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["framed-prefix-pool"] = ipv6Neighbor.FramedPrefixPool
    leafs["managed-config"] = ipv6Neighbor.ManagedConfig
    leafs["other-config"] = ipv6Neighbor.OtherConfig
    leafs["start-ra-on-ipv6-enable"] = ipv6Neighbor.StartRaOnIpv6Enable
    leafs["nud-enable"] = ipv6Neighbor.NudEnable
    leafs["ra-lifetime"] = ipv6Neighbor.RaLifetime
    leafs["router-preference"] = ipv6Neighbor.RouterPreference
    leafs["ra-suppress"] = ipv6Neighbor.RaSuppress
    leafs["ra-unicast"] = ipv6Neighbor.RaUnicast
    leafs["ra-unspecify-hoplimit"] = ipv6Neighbor.RaUnspecifyHoplimit
    leafs["ra-suppress-mtu"] = ipv6Neighbor.RaSuppressMtu
    leafs["suppress-cache-learning"] = ipv6Neighbor.SuppressCacheLearning
    leafs["reachable-time"] = ipv6Neighbor.ReachableTime
    leafs["ns-interval"] = ipv6Neighbor.NsInterval
    return leafs
}

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetYangName() string { return "ipv6-neighbor" }

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) SetParent(parent types.Entity) { ipv6Neighbor.parent = parent }

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetParent() types.Entity { return ipv6Neighbor.parent }

func (ipv6Neighbor *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval
// Set IPv6 Router Advertisement (RA) interval in
// seconds
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Maximum interface{}

    // Minimum RA interval in seconds. Must be less than 0.75 * maximum interval.
    // The type is interface{} with range: 3..1800. Units are second.
    Minimum interface{}
}

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetFilter() yfilter.YFilter { return raInterval.YFilter }

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) SetFilter(yf yfilter.YFilter) { raInterval.YFilter = yf }

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "minimum" { return "Minimum" }
    return ""
}

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetSegmentPath() string {
    return "ra-interval"
}

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = raInterval.Maximum
    leafs["minimum"] = raInterval.Minimum
    return leafs
}

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetBundleName() string { return "cisco_ios_xr" }

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetYangName() string { return "ra-interval" }

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) SetParent(parent types.Entity) { raInterval.parent = parent }

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetParent() types.Entity { return raInterval.parent }

func (raInterval *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInterval) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix
// Set the IPv6 framed ipv6 prefix for a
// subscriber interface 
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 framed prefix length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}

    // IPV6 framed prefix address. The type is string. This attribute is
    // mandatory.
    Prefix interface{}
}

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetFilter() yfilter.YFilter { return framedPrefix.YFilter }

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) SetFilter(yf yfilter.YFilter) { framedPrefix.YFilter = yf }

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetSegmentPath() string {
    return "framed-prefix"
}

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = framedPrefix.PrefixLength
    leafs["prefix"] = framedPrefix.Prefix
    return leafs
}

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetYangName() string { return "framed-prefix" }

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) SetParent(parent types.Entity) { framedPrefix.parent = parent }

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetParent() types.Entity { return framedPrefix.parent }

func (framedPrefix *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_FramedPrefix) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection
// Duplicate Address Detection (DAD)
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set IPv6 duplicate address detection transmits. The type is interface{}
    // with range: 0..600.
    Attempts interface{}
}

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetFilter() yfilter.YFilter { return duplicateAddressDetection.YFilter }

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) SetFilter(yf yfilter.YFilter) { duplicateAddressDetection.YFilter = yf }

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetGoName(yname string) string {
    if yname == "attempts" { return "Attempts" }
    return ""
}

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetSegmentPath() string {
    return "duplicate-address-detection"
}

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attempts"] = duplicateAddressDetection.Attempts
    return leafs
}

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetBundleName() string { return "cisco_ios_xr" }

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetYangName() string { return "duplicate-address-detection" }

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) SetParent(parent types.Entity) { duplicateAddressDetection.parent = parent }

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetParent() types.Entity { return duplicateAddressDetection.parent }

func (duplicateAddressDetection *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_DuplicateAddressDetection) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial
// IPv6 ND RA Initial
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Initial RA count. The type is interface{} with range: 0..32. This attribute
    // is mandatory.
    Count interface{}

    // Initial RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Interval interface{}
}

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetFilter() yfilter.YFilter { return raInitial.YFilter }

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) SetFilter(yf yfilter.YFilter) { raInitial.YFilter = yf }

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetGoName(yname string) string {
    if yname == "count" { return "Count" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetSegmentPath() string {
    return "ra-initial"
}

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["count"] = raInitial.Count
    leafs["interval"] = raInitial.Interval
    return leafs
}

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetBundleName() string { return "cisco_ios_xr" }

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetYangName() string { return "ra-initial" }

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) SetParent(parent types.Entity) { raInitial.parent = parent }

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetParent() types.Entity { return raInitial.parent }

func (raInitial *DynamicTemplate_Ppps_Ppp_Ipv6Neighbor_RaInitial) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_Ppps_Ppp_Dhcpv6
// Interface dhcpv6 configuration data
type DynamicTemplate_Ppps_Ppp_Dhcpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dns IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DnsIpv6Address interface{}

    // Select proxy/server profile based on mode class name. The type is string.
    ModeClass interface{}

    // The pool to be used for Address assignment. The type is string.
    AddressPool interface{}

    // The pool to be used for Prefix Delegation. The type is string.
    DelegatedPrefixPool interface{}

    // The class to be used for proxy/server profile. The type is string.
    Class interface{}

    // Stateful IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StatefulAddress interface{}

    // The prefix to be used for Prefix Delegation.
    DelegatedPrefix DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix
}

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetFilter() yfilter.YFilter { return dhcpv6.YFilter }

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) SetFilter(yf yfilter.YFilter) { dhcpv6.YFilter = yf }

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetGoName(yname string) string {
    if yname == "dns-ipv6address" { return "DnsIpv6Address" }
    if yname == "mode-class" { return "ModeClass" }
    if yname == "address-pool" { return "AddressPool" }
    if yname == "delegated-prefix-pool" { return "DelegatedPrefixPool" }
    if yname == "class" { return "Class" }
    if yname == "stateful-address" { return "StatefulAddress" }
    if yname == "delegated-prefix" { return "DelegatedPrefix" }
    return ""
}

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6"
}

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "delegated-prefix" {
        return &dhcpv6.DelegatedPrefix
    }
    return nil
}

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["delegated-prefix"] = &dhcpv6.DelegatedPrefix
    return children
}

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dns-ipv6address"] = dhcpv6.DnsIpv6Address
    leafs["mode-class"] = dhcpv6.ModeClass
    leafs["address-pool"] = dhcpv6.AddressPool
    leafs["delegated-prefix-pool"] = dhcpv6.DelegatedPrefixPool
    leafs["class"] = dhcpv6.Class
    leafs["stateful-address"] = dhcpv6.StatefulAddress
    return leafs
}

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetYangName() string { return "dhcpv6" }

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) SetParent(parent types.Entity) { dhcpv6.parent = parent }

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetParent() types.Entity { return dhcpv6.parent }

func (dhcpv6 *DynamicTemplate_Ppps_Ppp_Dhcpv6) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix
// The prefix to be used for Prefix Delegation
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Prefix interface{}

    // PD Prefix Length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}
}

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetFilter() yfilter.YFilter { return delegatedPrefix.YFilter }

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) SetFilter(yf yfilter.YFilter) { delegatedPrefix.YFilter = yf }

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetSegmentPath() string {
    return "delegated-prefix"
}

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = delegatedPrefix.Prefix
    leafs["prefix-length"] = delegatedPrefix.PrefixLength
    return leafs
}

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetYangName() string { return "delegated-prefix" }

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) SetParent(parent types.Entity) { delegatedPrefix.parent = parent }

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetParent() types.Entity { return delegatedPrefix.parent }

func (delegatedPrefix *DynamicTemplate_Ppps_Ppp_Dhcpv6_DelegatedPrefix) GetParentYangName() string { return "dhcpv6" }

// DynamicTemplate_Ppps_Ppp_Pbr
// Dynamic Template PBR configuration
type DynamicTemplate_Ppps_Ppp_Pbr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Class for subscriber ingress policy. The type is string.
    ServicePolicyIn interface{}

    // PBR service policy configuration.
    ServicePolicy DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy
}

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetFilter() yfilter.YFilter { return pbr.YFilter }

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) SetFilter(yf yfilter.YFilter) { pbr.YFilter = yf }

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetGoName(yname string) string {
    if yname == "service-policy-in" { return "ServicePolicyIn" }
    if yname == "service-policy" { return "ServicePolicy" }
    return ""
}

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetSegmentPath() string {
    return "Cisco-IOS-XR-pbr-subscriber-cfg:pbr"
}

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-policy" {
        return &pbr.ServicePolicy
    }
    return nil
}

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-policy"] = &pbr.ServicePolicy
    return children
}

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["service-policy-in"] = pbr.ServicePolicyIn
    return leafs
}

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetBundleName() string { return "cisco_ios_xr" }

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetYangName() string { return "pbr" }

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) SetParent(parent types.Entity) { pbr.parent = parent }

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetParent() types.Entity { return pbr.parent }

func (pbr *DynamicTemplate_Ppps_Ppp_Pbr) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy
// PBR service policy configuration
type DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ingress service policy. The type is string.
    Input interface{}
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetFilter() yfilter.YFilter { return servicePolicy.YFilter }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) SetFilter(yf yfilter.YFilter) { servicePolicy.YFilter = yf }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetSegmentPath() string {
    return "service-policy"
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input"] = servicePolicy.Input
    return leafs
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetYangName() string { return "service-policy" }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) SetParent(parent types.Entity) { servicePolicy.parent = parent }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetParent() types.Entity { return servicePolicy.parent }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Pbr_ServicePolicy) GetParentYangName() string { return "pbr" }

// DynamicTemplate_Ppps_Ppp_PppTemplate
// PPP template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPP FSM global template configuration data.
    Fsm DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm

    // PPP LCP global template configuration data.
    Lcp DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp

    // PPP IPv6CP global template configuration data.
    Ipv6Cp DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp

    // PPP IPCP global template configuration data.
    Ipcp DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp
}

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetFilter() yfilter.YFilter { return pppTemplate.YFilter }

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) SetFilter(yf yfilter.YFilter) { pppTemplate.YFilter = yf }

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetGoName(yname string) string {
    if yname == "fsm" { return "Fsm" }
    if yname == "lcp" { return "Lcp" }
    if yname == "ipv6cp" { return "Ipv6Cp" }
    if yname == "ipcp" { return "Ipcp" }
    return ""
}

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetSegmentPath() string {
    return "Cisco-IOS-XR-ppp-ma-gbl-cfg:ppp-template"
}

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fsm" {
        return &pppTemplate.Fsm
    }
    if childYangName == "lcp" {
        return &pppTemplate.Lcp
    }
    if childYangName == "ipv6cp" {
        return &pppTemplate.Ipv6Cp
    }
    if childYangName == "ipcp" {
        return &pppTemplate.Ipcp
    }
    return nil
}

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fsm"] = &pppTemplate.Fsm
    children["lcp"] = &pppTemplate.Lcp
    children["ipv6cp"] = &pppTemplate.Ipv6Cp
    children["ipcp"] = &pppTemplate.Ipcp
    return children
}

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetYangName() string { return "ppp-template" }

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) SetParent(parent types.Entity) { pppTemplate.parent = parent }

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetParent() types.Entity { return pppTemplate.parent }

func (pppTemplate *DynamicTemplate_Ppps_Ppp_PppTemplate) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm
// PPP FSM global template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm struct {
    parent types.Entity
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

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetFilter() yfilter.YFilter { return fsm.YFilter }

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) SetFilter(yf yfilter.YFilter) { fsm.YFilter = yf }

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetGoName(yname string) string {
    if yname == "max-consecutive-conf-naks" { return "MaxConsecutiveConfNaks" }
    if yname == "max-unacknowledged-conf-requests" { return "MaxUnacknowledgedConfRequests" }
    if yname == "retry-timeout" { return "RetryTimeout" }
    if yname == "protocol-reject-timeout" { return "ProtocolRejectTimeout" }
    return ""
}

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetSegmentPath() string {
    return "fsm"
}

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-consecutive-conf-naks"] = fsm.MaxConsecutiveConfNaks
    leafs["max-unacknowledged-conf-requests"] = fsm.MaxUnacknowledgedConfRequests
    leafs["retry-timeout"] = fsm.RetryTimeout
    leafs["protocol-reject-timeout"] = fsm.ProtocolRejectTimeout
    return leafs
}

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetBundleName() string { return "cisco_ios_xr" }

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetYangName() string { return "fsm" }

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) SetParent(parent types.Entity) { fsm.parent = parent }

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetParent() types.Entity { return fsm.parent }

func (fsm *DynamicTemplate_Ppps_Ppp_PppTemplate_Fsm) GetParentYangName() string { return "ppp-template" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp
// PPP LCP global template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp struct {
    parent types.Entity
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

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetFilter() yfilter.YFilter { return lcp.YFilter }

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) SetFilter(yf yfilter.YFilter) { lcp.YFilter = yf }

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetGoName(yname string) string {
    if yname == "renegotiation" { return "Renegotiation" }
    if yname == "service-type" { return "ServiceType" }
    if yname == "send-term-request-on-shut-down" { return "SendTermRequestOnShutDown" }
    if yname == "mru-ignore" { return "MruIgnore" }
    if yname == "absolute-timeout" { return "AbsoluteTimeout" }
    if yname == "delay" { return "Delay" }
    if yname == "authentication" { return "Authentication" }
    if yname == "keepalive" { return "Keepalive" }
    return ""
}

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetSegmentPath() string {
    return "lcp"
}

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "absolute-timeout" {
        return &lcp.AbsoluteTimeout
    }
    if childYangName == "delay" {
        return &lcp.Delay
    }
    if childYangName == "authentication" {
        return &lcp.Authentication
    }
    if childYangName == "keepalive" {
        return &lcp.Keepalive
    }
    return nil
}

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["absolute-timeout"] = &lcp.AbsoluteTimeout
    children["delay"] = &lcp.Delay
    children["authentication"] = &lcp.Authentication
    children["keepalive"] = &lcp.Keepalive
    return children
}

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["renegotiation"] = lcp.Renegotiation
    leafs["service-type"] = lcp.ServiceType
    leafs["send-term-request-on-shut-down"] = lcp.SendTermRequestOnShutDown
    leafs["mru-ignore"] = lcp.MruIgnore
    return leafs
}

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetBundleName() string { return "cisco_ios_xr" }

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetYangName() string { return "lcp" }

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) SetParent(parent types.Entity) { lcp.parent = parent }

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetParent() types.Entity { return lcp.parent }

func (lcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp) GetParentYangName() string { return "ppp-template" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout
// This specifies the session absolute timeout
// value
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minutes. The type is interface{} with range: 0..35000000.
    Minutes interface{}

    // Seconds. The type is interface{} with range: 0..59.
    Seconds interface{}
}

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetFilter() yfilter.YFilter { return absoluteTimeout.YFilter }

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) SetFilter(yf yfilter.YFilter) { absoluteTimeout.YFilter = yf }

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetGoName(yname string) string {
    if yname == "minutes" { return "Minutes" }
    if yname == "seconds" { return "Seconds" }
    return ""
}

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetSegmentPath() string {
    return "absolute-timeout"
}

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minutes"] = absoluteTimeout.Minutes
    leafs["seconds"] = absoluteTimeout.Seconds
    return leafs
}

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetBundleName() string { return "cisco_ios_xr" }

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetYangName() string { return "absolute-timeout" }

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) SetParent(parent types.Entity) { absoluteTimeout.parent = parent }

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetParent() types.Entity { return absoluteTimeout.parent }

func (absoluteTimeout *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_AbsoluteTimeout) GetParentYangName() string { return "lcp" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay
// This specifies the time to delay before
// starting active LCPnegotiations
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Seconds. The type is interface{} with range: 0..255.
    Seconds interface{}

    // Milliseconds. The type is interface{} with range: 0..999.
    Milliseconds interface{}
}

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetFilter() yfilter.YFilter { return delay.YFilter }

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) SetFilter(yf yfilter.YFilter) { delay.YFilter = yf }

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetGoName(yname string) string {
    if yname == "seconds" { return "Seconds" }
    if yname == "milliseconds" { return "Milliseconds" }
    return ""
}

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetSegmentPath() string {
    return "delay"
}

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["seconds"] = delay.Seconds
    leafs["milliseconds"] = delay.Milliseconds
    return leafs
}

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetBundleName() string { return "cisco_ios_xr" }

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetYangName() string { return "delay" }

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) SetParent(parent types.Entity) { delay.parent = parent }

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetParent() types.Entity { return delay.parent }

func (delay *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Delay) GetParentYangName() string { return "lcp" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication
// PPP authentication parameters
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This specifies the CHAP hostname. The type is string.
    ChapHostName interface{}

    // <1> for accepting null-passwordduring authentication. The type is
    // interface{} with range: -2147483648..2147483647.
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

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetGoName(yname string) string {
    if yname == "chap-host-name" { return "ChapHostName" }
    if yname == "pap" { return "Pap" }
    if yname == "mschap-host-name" { return "MschapHostName" }
    if yname == "max-authentication-failures" { return "MaxAuthenticationFailures" }
    if yname == "timeout" { return "Timeout" }
    if yname == "methods" { return "Methods" }
    return ""
}

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "methods" {
        return &authentication.Methods
    }
    return nil
}

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["methods"] = &authentication.Methods
    return children
}

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chap-host-name"] = authentication.ChapHostName
    leafs["pap"] = authentication.Pap
    leafs["mschap-host-name"] = authentication.MschapHostName
    leafs["max-authentication-failures"] = authentication.MaxAuthenticationFailures
    leafs["timeout"] = authentication.Timeout
    return leafs
}

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetYangName() string { return "authentication" }

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication) GetParentYangName() string { return "lcp" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods
// This specifies the PPP link authentication
// method
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Select between one and three authentication methods in order of preference.
    // The type is slice of PppAuthenticationMethodGbl.
    Method []interface{}
}

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetFilter() yfilter.YFilter { return methods.YFilter }

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) SetFilter(yf yfilter.YFilter) { methods.YFilter = yf }

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetGoName(yname string) string {
    if yname == "method" { return "Method" }
    return ""
}

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetSegmentPath() string {
    return "methods"
}

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method"] = methods.Method
    return leafs
}

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetBundleName() string { return "cisco_ios_xr" }

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetYangName() string { return "methods" }

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) SetParent(parent types.Entity) { methods.parent = parent }

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetParent() types.Entity { return methods.parent }

func (methods *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Authentication_Methods) GetParentYangName() string { return "authentication" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive
// This specifies the rate at which EchoReq
// packets are sent
type DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive struct {
    parent types.Entity
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

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetFilter() yfilter.YFilter { return keepalive.YFilter }

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) SetFilter(yf yfilter.YFilter) { keepalive.YFilter = yf }

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetGoName(yname string) string {
    if yname == "keepalive-disable" { return "KeepaliveDisable" }
    if yname == "interval" { return "Interval" }
    if yname == "retry-count" { return "RetryCount" }
    return ""
}

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetSegmentPath() string {
    return "keepalive"
}

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["keepalive-disable"] = keepalive.KeepaliveDisable
    leafs["interval"] = keepalive.Interval
    leafs["retry-count"] = keepalive.RetryCount
    return leafs
}

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetBundleName() string { return "cisco_ios_xr" }

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetYangName() string { return "keepalive" }

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) SetParent(parent types.Entity) { keepalive.parent = parent }

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetParent() types.Entity { return keepalive.parent }

func (keepalive *DynamicTemplate_Ppps_Ppp_PppTemplate_Lcp_Keepalive) GetParentYangName() string { return "lcp" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp
// PPP IPv6CP global template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp struct {
    parent types.Entity
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

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetFilter() yfilter.YFilter { return ipv6Cp.YFilter }

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) SetFilter(yf yfilter.YFilter) { ipv6Cp.YFilter = yf }

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetGoName(yname string) string {
    if yname == "passive" { return "Passive" }
    if yname == "renegotiation" { return "Renegotiation" }
    if yname == "peer-interface-id" { return "PeerInterfaceId" }
    if yname == "protocol-reject" { return "ProtocolReject" }
    return ""
}

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetSegmentPath() string {
    return "ipv6cp"
}

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["passive"] = ipv6Cp.Passive
    leafs["renegotiation"] = ipv6Cp.Renegotiation
    leafs["peer-interface-id"] = ipv6Cp.PeerInterfaceId
    leafs["protocol-reject"] = ipv6Cp.ProtocolReject
    return leafs
}

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetYangName() string { return "ipv6cp" }

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) SetParent(parent types.Entity) { ipv6Cp.parent = parent }

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetParent() types.Entity { return ipv6Cp.parent }

func (ipv6Cp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipv6Cp) GetParentYangName() string { return "ppp-template" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp
// PPP IPCP global template configuration data
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerNetmask interface{}

    // IPCP WINS parameters.
    Wins DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins

    // IPCP DNS parameters.
    Dns DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns

    // IPCP address parameters.
    PeerAddress DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress
}

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetFilter() yfilter.YFilter { return ipcp.YFilter }

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) SetFilter(yf yfilter.YFilter) { ipcp.YFilter = yf }

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetGoName(yname string) string {
    if yname == "renegotiation" { return "Renegotiation" }
    if yname == "passive" { return "Passive" }
    if yname == "protocol-reject" { return "ProtocolReject" }
    if yname == "peer-netmask" { return "PeerNetmask" }
    if yname == "wins" { return "Wins" }
    if yname == "dns" { return "Dns" }
    if yname == "peer-address" { return "PeerAddress" }
    return ""
}

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetSegmentPath() string {
    return "ipcp"
}

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "wins" {
        return &ipcp.Wins
    }
    if childYangName == "dns" {
        return &ipcp.Dns
    }
    if childYangName == "peer-address" {
        return &ipcp.PeerAddress
    }
    return nil
}

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["wins"] = &ipcp.Wins
    children["dns"] = &ipcp.Dns
    children["peer-address"] = &ipcp.PeerAddress
    return children
}

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["renegotiation"] = ipcp.Renegotiation
    leafs["passive"] = ipcp.Passive
    leafs["protocol-reject"] = ipcp.ProtocolReject
    leafs["peer-netmask"] = ipcp.PeerNetmask
    return leafs
}

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetBundleName() string { return "cisco_ios_xr" }

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetYangName() string { return "ipcp" }

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) SetParent(parent types.Entity) { ipcp.parent = parent }

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetParent() types.Entity { return ipcp.parent }

func (ipcp *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp) GetParentYangName() string { return "ppp-template" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins
// IPCP WINS parameters
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify WINS address(es) to provide.
    WinsAddresses DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses
}

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetFilter() yfilter.YFilter { return wins.YFilter }

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) SetFilter(yf yfilter.YFilter) { wins.YFilter = yf }

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetGoName(yname string) string {
    if yname == "wins-addresses" { return "WinsAddresses" }
    return ""
}

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetSegmentPath() string {
    return "wins"
}

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "wins-addresses" {
        return &wins.WinsAddresses
    }
    return nil
}

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["wins-addresses"] = &wins.WinsAddresses
    return children
}

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetBundleName() string { return "cisco_ios_xr" }

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetYangName() string { return "wins" }

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) SetParent(parent types.Entity) { wins.parent = parent }

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetParent() types.Entity { return wins.parent }

func (wins *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins) GetParentYangName() string { return "ipcp" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses
// Specify WINS address(es) to provide
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Primary WINS IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // Secondary WINS IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Secondary interface{}
}

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetFilter() yfilter.YFilter { return winsAddresses.YFilter }

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) SetFilter(yf yfilter.YFilter) { winsAddresses.YFilter = yf }

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetGoName(yname string) string {
    if yname == "primary" { return "Primary" }
    if yname == "secondary" { return "Secondary" }
    return ""
}

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetSegmentPath() string {
    return "wins-addresses"
}

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["primary"] = winsAddresses.Primary
    leafs["secondary"] = winsAddresses.Secondary
    return leafs
}

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetYangName() string { return "wins-addresses" }

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) SetParent(parent types.Entity) { winsAddresses.parent = parent }

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetParent() types.Entity { return winsAddresses.parent }

func (winsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Wins_WinsAddresses) GetParentYangName() string { return "wins" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns
// IPCP DNS parameters
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify DNS address(es) to provide.
    DnsAddresses DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses
}

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetFilter() yfilter.YFilter { return dns.YFilter }

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) SetFilter(yf yfilter.YFilter) { dns.YFilter = yf }

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetGoName(yname string) string {
    if yname == "dns-addresses" { return "DnsAddresses" }
    return ""
}

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetSegmentPath() string {
    return "dns"
}

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dns-addresses" {
        return &dns.DnsAddresses
    }
    return nil
}

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dns-addresses"] = &dns.DnsAddresses
    return children
}

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetBundleName() string { return "cisco_ios_xr" }

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetYangName() string { return "dns" }

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) SetParent(parent types.Entity) { dns.parent = parent }

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetParent() types.Entity { return dns.parent }

func (dns *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns) GetParentYangName() string { return "ipcp" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses
// Specify DNS address(es) to provide
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Primary DNS IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // Secondary DNS IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Secondary interface{}
}

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetFilter() yfilter.YFilter { return dnsAddresses.YFilter }

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) SetFilter(yf yfilter.YFilter) { dnsAddresses.YFilter = yf }

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetGoName(yname string) string {
    if yname == "primary" { return "Primary" }
    if yname == "secondary" { return "Secondary" }
    return ""
}

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetSegmentPath() string {
    return "dns-addresses"
}

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["primary"] = dnsAddresses.Primary
    leafs["secondary"] = dnsAddresses.Secondary
    return leafs
}

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetYangName() string { return "dns-addresses" }

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) SetParent(parent types.Entity) { dnsAddresses.parent = parent }

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetParent() types.Entity { return dnsAddresses.parent }

func (dnsAddresses *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_Dns_DnsAddresses) GetParentYangName() string { return "dns" }

// DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress
// IPCP address parameters
type DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify an IP address to assign to peers through IPCP. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Default interface{}

    // Accepts an IP address from the peer if in the pool, else allocates one from
    // the pool. The type is string.
    Pool interface{}
}

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetFilter() yfilter.YFilter { return peerAddress.YFilter }

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) SetFilter(yf yfilter.YFilter) { peerAddress.YFilter = yf }

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetGoName(yname string) string {
    if yname == "default" { return "Default" }
    if yname == "pool" { return "Pool" }
    return ""
}

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetSegmentPath() string {
    return "peer-address"
}

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default"] = peerAddress.Default
    leafs["pool"] = peerAddress.Pool
    return leafs
}

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetBundleName() string { return "cisco_ios_xr" }

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetYangName() string { return "peer-address" }

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) SetParent(parent types.Entity) { peerAddress.parent = parent }

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetParent() types.Entity { return peerAddress.parent }

func (peerAddress *DynamicTemplate_Ppps_Ppp_PppTemplate_Ipcp_PeerAddress) GetParentYangName() string { return "ipcp" }

// DynamicTemplate_Ppps_Ppp_Qos
// QoS dynamically applied configuration template
type DynamicTemplate_Ppps_Ppp_Qos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service policy to be applied in ingress/egress direction.
    ServicePolicy DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy

    // QoS L2 overhead accounting.
    Account DynamicTemplate_Ppps_Ppp_Qos_Account

    // QoS to be applied in egress direction.
    Output DynamicTemplate_Ppps_Ppp_Qos_Output
}

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetFilter() yfilter.YFilter { return qos.YFilter }

func (qos *DynamicTemplate_Ppps_Ppp_Qos) SetFilter(yf yfilter.YFilter) { qos.YFilter = yf }

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetGoName(yname string) string {
    if yname == "service-policy" { return "ServicePolicy" }
    if yname == "account" { return "Account" }
    if yname == "output" { return "Output" }
    return ""
}

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetSegmentPath() string {
    return "Cisco-IOS-XR-qos-ma-bng-cfg:qos"
}

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-policy" {
        return &qos.ServicePolicy
    }
    if childYangName == "account" {
        return &qos.Account
    }
    if childYangName == "output" {
        return &qos.Output
    }
    return nil
}

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-policy"] = &qos.ServicePolicy
    children["account"] = &qos.Account
    children["output"] = &qos.Output
    return children
}

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetBundleName() string { return "cisco_ios_xr" }

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetYangName() string { return "qos" }

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qos *DynamicTemplate_Ppps_Ppp_Qos) SetParent(parent types.Entity) { qos.parent = parent }

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetParent() types.Entity { return qos.parent }

func (qos *DynamicTemplate_Ppps_Ppp_Qos) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy
// Service policy to be applied in ingress/egress
// direction
type DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber ingress policy.
    Input DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input

    // Subscriber egress policy.
    Output DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetFilter() yfilter.YFilter { return servicePolicy.YFilter }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) SetFilter(yf yfilter.YFilter) { servicePolicy.YFilter = yf }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetSegmentPath() string {
    return "service-policy"
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &servicePolicy.Input
    }
    if childYangName == "output" {
        return &servicePolicy.Output
    }
    return nil
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &servicePolicy.Input
    children["output"] = &servicePolicy.Output
    return children
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetYangName() string { return "service-policy" }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) SetParent(parent types.Entity) { servicePolicy.parent = parent }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetParent() types.Entity { return servicePolicy.parent }

func (servicePolicy *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy) GetParentYangName() string { return "qos" }

// DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input
// Subscriber ingress policy
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "spi-name" { return "SpiName" }
    if yname == "merge" { return "Merge" }
    if yname == "merge-id" { return "MergeId" }
    if yname == "account-stats" { return "AccountStats" }
    return ""
}

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetSegmentPath() string {
    return "input"
}

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = input.PolicyName
    leafs["spi-name"] = input.SpiName
    leafs["merge"] = input.Merge
    leafs["merge-id"] = input.MergeId
    leafs["account-stats"] = input.AccountStats
    return leafs
}

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetYangName() string { return "input" }

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetParent() types.Entity { return input.parent }

func (input *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Input) GetParentYangName() string { return "service-policy" }

// DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output
// Subscriber egress policy
// This type is a presence type.
type DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "spi-name" { return "SpiName" }
    if yname == "merge" { return "Merge" }
    if yname == "merge-id" { return "MergeId" }
    if yname == "account-stats" { return "AccountStats" }
    return ""
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetSegmentPath() string {
    return "output"
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = output.PolicyName
    leafs["spi-name"] = output.SpiName
    leafs["merge"] = output.Merge
    leafs["merge-id"] = output.MergeId
    leafs["account-stats"] = output.AccountStats
    return leafs
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetYangName() string { return "output" }

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetParent() types.Entity { return output.parent }

func (output *DynamicTemplate_Ppps_Ppp_Qos_ServicePolicy_Output) GetParentYangName() string { return "service-policy" }

// DynamicTemplate_Ppps_Ppp_Qos_Account
// QoS L2 overhead accounting
type DynamicTemplate_Ppps_Ppp_Qos_Account struct {
    parent types.Entity
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

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetFilter() yfilter.YFilter { return account.YFilter }

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) SetFilter(yf yfilter.YFilter) { account.YFilter = yf }

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetGoName(yname string) string {
    if yname == "aal" { return "Aal" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "atm-cell-tax" { return "AtmCellTax" }
    if yname == "user-defined" { return "UserDefined" }
    return ""
}

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetSegmentPath() string {
    return "account"
}

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aal"] = account.Aal
    leafs["encapsulation"] = account.Encapsulation
    leafs["atm-cell-tax"] = account.AtmCellTax
    leafs["user-defined"] = account.UserDefined
    return leafs
}

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetBundleName() string { return "cisco_ios_xr" }

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetYangName() string { return "account" }

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) SetParent(parent types.Entity) { account.parent = parent }

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetParent() types.Entity { return account.parent }

func (account *DynamicTemplate_Ppps_Ppp_Qos_Account) GetParentYangName() string { return "qos" }

// DynamicTemplate_Ppps_Ppp_Qos_Output
// QoS to be applied in egress direction
type DynamicTemplate_Ppps_Ppp_Qos_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum bandwidth value for the subscriber (in kbps). The type is
    // interface{} with range: 1..4294967295. Units are kbit/s.
    MinimumBandwidth interface{}
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetGoName(yname string) string {
    if yname == "minimum-bandwidth" { return "MinimumBandwidth" }
    return ""
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetSegmentPath() string {
    return "output"
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minimum-bandwidth"] = output.MinimumBandwidth
    return leafs
}

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetYangName() string { return "output" }

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetParent() types.Entity { return output.parent }

func (output *DynamicTemplate_Ppps_Ppp_Qos_Output) GetParentYangName() string { return "qos" }

// DynamicTemplate_Ppps_Ppp_Accounting
// Subscriber accounting dynamic-template commands
type DynamicTemplate_Ppps_Ppp_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber accounting prepaid feature. The type is string.
    PrepaidFeature interface{}

    // Subscriber accounting idle timeout.
    IdleTimeout DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout

    // Subscriber accounting session accounting.
    Session DynamicTemplate_Ppps_Ppp_Accounting_Session

    // Subscriber accounting service accounting.
    ServiceAccounting DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting
}

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetGoName(yname string) string {
    if yname == "prepaid-feature" { return "PrepaidFeature" }
    if yname == "idle-timeout" { return "IdleTimeout" }
    if yname == "session" { return "Session" }
    if yname == "service-accounting" { return "ServiceAccounting" }
    return ""
}

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-accounting-cfg:accounting"
}

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "idle-timeout" {
        return &accounting.IdleTimeout
    }
    if childYangName == "session" {
        return &accounting.Session
    }
    if childYangName == "service-accounting" {
        return &accounting.ServiceAccounting
    }
    return nil
}

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["idle-timeout"] = &accounting.IdleTimeout
    children["session"] = &accounting.Session
    children["service-accounting"] = &accounting.ServiceAccounting
    return children
}

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prepaid-feature"] = accounting.PrepaidFeature
    return leafs
}

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetYangName() string { return "accounting" }

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *DynamicTemplate_Ppps_Ppp_Accounting) GetParentYangName() string { return "ppp" }

// DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout
// Subscriber accounting idle timeout
type DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout struct {
    parent types.Entity
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

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetFilter() yfilter.YFilter { return idleTimeout.YFilter }

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) SetFilter(yf yfilter.YFilter) { idleTimeout.YFilter = yf }

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetGoName(yname string) string {
    if yname == "timeout-value" { return "TimeoutValue" }
    if yname == "threshold" { return "Threshold" }
    if yname == "direction" { return "Direction" }
    return ""
}

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetSegmentPath() string {
    return "idle-timeout"
}

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout-value"] = idleTimeout.TimeoutValue
    leafs["threshold"] = idleTimeout.Threshold
    leafs["direction"] = idleTimeout.Direction
    return leafs
}

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetBundleName() string { return "cisco_ios_xr" }

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetYangName() string { return "idle-timeout" }

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) SetParent(parent types.Entity) { idleTimeout.parent = parent }

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetParent() types.Entity { return idleTimeout.parent }

func (idleTimeout *DynamicTemplate_Ppps_Ppp_Accounting_IdleTimeout) GetParentYangName() string { return "accounting" }

// DynamicTemplate_Ppps_Ppp_Accounting_Session
// Subscriber accounting session accounting
type DynamicTemplate_Ppps_Ppp_Accounting_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session accounting method list name. The type is string.
    MethodListName interface{}

    // Interim accounting interval in minutes. The type is interface{} with range:
    // -2147483648..2147483647.
    PeriodicInterval interface{}

    // Dual stack wait delay in seconds. The type is interface{} with range:
    // -2147483648..2147483647.
    DualStackDelay interface{}

    // Hold Accounting start based on IA_PD. The type is interface{} with range:
    // -2147483648..2147483647.
    HoldAcctStart interface{}
}

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetGoName(yname string) string {
    if yname == "method-list-name" { return "MethodListName" }
    if yname == "periodic-interval" { return "PeriodicInterval" }
    if yname == "dual-stack-delay" { return "DualStackDelay" }
    if yname == "hold-acct-start" { return "HoldAcctStart" }
    return ""
}

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetSegmentPath() string {
    return "session"
}

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method-list-name"] = session.MethodListName
    leafs["periodic-interval"] = session.PeriodicInterval
    leafs["dual-stack-delay"] = session.DualStackDelay
    leafs["hold-acct-start"] = session.HoldAcctStart
    return leafs
}

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetYangName() string { return "session" }

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetParent() types.Entity { return session.parent }

func (session *DynamicTemplate_Ppps_Ppp_Accounting_Session) GetParentYangName() string { return "accounting" }

// DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting
// Subscriber accounting service accounting
type DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service accounting method list name. The type is string.
    MethodListName interface{}

    // Accounting interim interval in minutes. The type is interface{} with range:
    // -2147483648..2147483647.
    AccountingInterimInterval interface{}
}

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetFilter() yfilter.YFilter { return serviceAccounting.YFilter }

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) SetFilter(yf yfilter.YFilter) { serviceAccounting.YFilter = yf }

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetGoName(yname string) string {
    if yname == "method-list-name" { return "MethodListName" }
    if yname == "accounting-interim-interval" { return "AccountingInterimInterval" }
    return ""
}

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetSegmentPath() string {
    return "service-accounting"
}

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method-list-name"] = serviceAccounting.MethodListName
    leafs["accounting-interim-interval"] = serviceAccounting.AccountingInterimInterval
    return leafs
}

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetBundleName() string { return "cisco_ios_xr" }

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetYangName() string { return "service-accounting" }

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) SetParent(parent types.Entity) { serviceAccounting.parent = parent }

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetParent() types.Entity { return serviceAccounting.parent }

func (serviceAccounting *DynamicTemplate_Ppps_Ppp_Accounting_ServiceAccounting) GetParentYangName() string { return "accounting" }

// DynamicTemplate_Ppps_Ppp_PppoeTemplate
// PPPoE template configuration data
type DynamicTemplate_Ppps_Ppp_PppoeTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the Port limit (attr 62) to apply to the subscriber. The type is
    // interface{} with range: 1..65535.
    PortLimit interface{}
}

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetFilter() yfilter.YFilter { return pppoeTemplate.YFilter }

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) SetFilter(yf yfilter.YFilter) { pppoeTemplate.YFilter = yf }

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetGoName(yname string) string {
    if yname == "port-limit" { return "PortLimit" }
    return ""
}

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-pppoe-ma-gbl-cfg:pppoe-template"
}

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["port-limit"] = pppoeTemplate.PortLimit
    return leafs
}

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetYangName() string { return "pppoe-template" }

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) SetParent(parent types.Entity) { pppoeTemplate.parent = parent }

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetParent() types.Entity { return pppoeTemplate.parent }

func (pppoeTemplate *DynamicTemplate_Ppps_Ppp_PppoeTemplate) GetParentYangName() string { return "ppp" }

// DynamicTemplate_IpSubscribers
// The IP Subscriber Template Table
type DynamicTemplate_IpSubscribers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A IP Subscriber Type Template . The type is slice of
    // DynamicTemplate_IpSubscribers_IpSubscriber.
    IpSubscriber []DynamicTemplate_IpSubscribers_IpSubscriber
}

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetFilter() yfilter.YFilter { return ipSubscribers.YFilter }

func (ipSubscribers *DynamicTemplate_IpSubscribers) SetFilter(yf yfilter.YFilter) { ipSubscribers.YFilter = yf }

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetGoName(yname string) string {
    if yname == "ip-subscriber" { return "IpSubscriber" }
    return ""
}

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetSegmentPath() string {
    return "ip-subscribers"
}

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ip-subscriber" {
        for _, c := range ipSubscribers.IpSubscriber {
            if ipSubscribers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DynamicTemplate_IpSubscribers_IpSubscriber{}
        ipSubscribers.IpSubscriber = append(ipSubscribers.IpSubscriber, child)
        return &ipSubscribers.IpSubscriber[len(ipSubscribers.IpSubscriber)-1]
    }
    return nil
}

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipSubscribers.IpSubscriber {
        children[ipSubscribers.IpSubscriber[i].GetSegmentPath()] = &ipSubscribers.IpSubscriber[i]
    }
    return children
}

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetBundleName() string { return "cisco_ios_xr" }

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetYangName() string { return "ip-subscribers" }

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipSubscribers *DynamicTemplate_IpSubscribers) SetParent(parent types.Entity) { ipSubscribers.parent = parent }

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetParent() types.Entity { return ipSubscribers.parent }

func (ipSubscribers *DynamicTemplate_IpSubscribers) GetParentYangName() string { return "dynamic-template" }

// DynamicTemplate_IpSubscribers_IpSubscriber
// A IP Subscriber Type Template 
type DynamicTemplate_IpSubscribers_IpSubscriber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the template. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetFilter() yfilter.YFilter { return ipSubscriber.YFilter }

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) SetFilter(yf yfilter.YFilter) { ipSubscriber.YFilter = yf }

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "vrf" { return "Vrf" }
    if yname == "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions" { return "SpanMonitorSessions" }
    if yname == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter" { return "Ipv4PacketFilter" }
    if yname == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter" { return "Ipv6PacketFilter" }
    if yname == "Cisco-IOS-XR-ipv4-dhcpd-subscriber-cfg:dhcpd" { return "Dhcpd" }
    if yname == "Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp" { return "Igmp" }
    if yname == "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network" { return "Ipv4Network" }
    if yname == "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network" { return "Ipv6Network" }
    if yname == "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor" { return "Ipv6Neighbor" }
    if yname == "Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6" { return "Dhcpv6" }
    if yname == "Cisco-IOS-XR-pbr-subscriber-cfg:pbr" { return "Pbr" }
    if yname == "Cisco-IOS-XR-qos-ma-bng-cfg:qos" { return "Qos" }
    if yname == "Cisco-IOS-XR-subscriber-accounting-cfg:accounting" { return "Accounting" }
    return ""
}

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetSegmentPath() string {
    return "ip-subscriber" + "[template-name='" + fmt.Sprintf("%v", ipSubscriber.TemplateName) + "']"
}

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions" {
        return &ipSubscriber.SpanMonitorSessions
    }
    if childYangName == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter" {
        return &ipSubscriber.Ipv4PacketFilter
    }
    if childYangName == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter" {
        return &ipSubscriber.Ipv6PacketFilter
    }
    if childYangName == "Cisco-IOS-XR-ipv4-dhcpd-subscriber-cfg:dhcpd" {
        return &ipSubscriber.Dhcpd
    }
    if childYangName == "Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp" {
        return &ipSubscriber.Igmp
    }
    if childYangName == "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network" {
        return &ipSubscriber.Ipv4Network
    }
    if childYangName == "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network" {
        return &ipSubscriber.Ipv6Network
    }
    if childYangName == "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor" {
        return &ipSubscriber.Ipv6Neighbor
    }
    if childYangName == "Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6" {
        return &ipSubscriber.Dhcpv6
    }
    if childYangName == "Cisco-IOS-XR-pbr-subscriber-cfg:pbr" {
        return &ipSubscriber.Pbr
    }
    if childYangName == "Cisco-IOS-XR-qos-ma-bng-cfg:qos" {
        return &ipSubscriber.Qos
    }
    if childYangName == "Cisco-IOS-XR-subscriber-accounting-cfg:accounting" {
        return &ipSubscriber.Accounting
    }
    return nil
}

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions"] = &ipSubscriber.SpanMonitorSessions
    children["Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter"] = &ipSubscriber.Ipv4PacketFilter
    children["Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter"] = &ipSubscriber.Ipv6PacketFilter
    children["Cisco-IOS-XR-ipv4-dhcpd-subscriber-cfg:dhcpd"] = &ipSubscriber.Dhcpd
    children["Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp"] = &ipSubscriber.Igmp
    children["Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network"] = &ipSubscriber.Ipv4Network
    children["Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network"] = &ipSubscriber.Ipv6Network
    children["Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor"] = &ipSubscriber.Ipv6Neighbor
    children["Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6"] = &ipSubscriber.Dhcpv6
    children["Cisco-IOS-XR-pbr-subscriber-cfg:pbr"] = &ipSubscriber.Pbr
    children["Cisco-IOS-XR-qos-ma-bng-cfg:qos"] = &ipSubscriber.Qos
    children["Cisco-IOS-XR-subscriber-accounting-cfg:accounting"] = &ipSubscriber.Accounting
    return children
}

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = ipSubscriber.TemplateName
    leafs["vrf"] = ipSubscriber.Vrf
    return leafs
}

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetBundleName() string { return "cisco_ios_xr" }

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetYangName() string { return "ip-subscriber" }

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) SetParent(parent types.Entity) { ipSubscriber.parent = parent }

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetParent() types.Entity { return ipSubscriber.parent }

func (ipSubscriber *DynamicTemplate_IpSubscribers_IpSubscriber) GetParentYangName() string { return "ip-subscribers" }

// DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions
// Monitor Session container for this template
type DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a particular class of Monitor Session. The type is slice
    // of
    // DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession.
    SpanMonitorSession []DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession
}

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetFilter() yfilter.YFilter { return spanMonitorSessions.YFilter }

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) SetFilter(yf yfilter.YFilter) { spanMonitorSessions.YFilter = yf }

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetGoName(yname string) string {
    if yname == "span-monitor-session" { return "SpanMonitorSession" }
    return ""
}

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetSegmentPath() string {
    return "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions"
}

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "span-monitor-session" {
        for _, c := range spanMonitorSessions.SpanMonitorSession {
            if spanMonitorSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession{}
        spanMonitorSessions.SpanMonitorSession = append(spanMonitorSessions.SpanMonitorSession, child)
        return &spanMonitorSessions.SpanMonitorSession[len(spanMonitorSessions.SpanMonitorSession)-1]
    }
    return nil
}

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spanMonitorSessions.SpanMonitorSession {
        children[spanMonitorSessions.SpanMonitorSession[i].GetSegmentPath()] = &spanMonitorSessions.SpanMonitorSession[i]
    }
    return children
}

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetBundleName() string { return "cisco_ios_xr" }

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetYangName() string { return "span-monitor-sessions" }

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) SetParent(parent types.Entity) { spanMonitorSessions.parent = parent }

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetParent() types.Entity { return spanMonitorSessions.parent }

func (spanMonitorSessions *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession
// Configuration for a particular class of Monitor
// Session
type DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetFilter() yfilter.YFilter { return spanMonitorSession.YFilter }

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) SetFilter(yf yfilter.YFilter) { spanMonitorSession.YFilter = yf }

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetGoName(yname string) string {
    if yname == "session-class" { return "SessionClass" }
    if yname == "mirror-first" { return "MirrorFirst" }
    if yname == "mirror-interval" { return "MirrorInterval" }
    if yname == "attachment" { return "Attachment" }
    if yname == "acl" { return "Acl" }
    return ""
}

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetSegmentPath() string {
    return "span-monitor-session" + "[session-class='" + fmt.Sprintf("%v", spanMonitorSession.SessionClass) + "']"
}

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attachment" {
        return &spanMonitorSession.Attachment
    }
    if childYangName == "acl" {
        return &spanMonitorSession.Acl
    }
    return nil
}

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attachment"] = &spanMonitorSession.Attachment
    children["acl"] = &spanMonitorSession.Acl
    return children
}

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-class"] = spanMonitorSession.SessionClass
    leafs["mirror-first"] = spanMonitorSession.MirrorFirst
    leafs["mirror-interval"] = spanMonitorSession.MirrorInterval
    return leafs
}

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetBundleName() string { return "cisco_ios_xr" }

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetYangName() string { return "span-monitor-session" }

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) SetParent(parent types.Entity) { spanMonitorSession.parent = parent }

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetParent() types.Entity { return spanMonitorSession.parent }

func (spanMonitorSession *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession) GetParentYangName() string { return "span-monitor-sessions" }

// DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment
// Attach the interface to a Monitor Session
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session Name. The type is string with length: 1..79. This attribute is
    // mandatory.
    SessionName interface{}

    // Specify the direction of traffic to replicate (optional). The type is
    // SpanTrafficDirection.
    Direction interface{}

    // Enable port level traffic mirroring. The type is interface{}.
    PortLevelEnable interface{}
}

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetFilter() yfilter.YFilter { return attachment.YFilter }

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) SetFilter(yf yfilter.YFilter) { attachment.YFilter = yf }

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetGoName(yname string) string {
    if yname == "session-name" { return "SessionName" }
    if yname == "direction" { return "Direction" }
    if yname == "port-level-enable" { return "PortLevelEnable" }
    return ""
}

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetSegmentPath() string {
    return "attachment"
}

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-name"] = attachment.SessionName
    leafs["direction"] = attachment.Direction
    leafs["port-level-enable"] = attachment.PortLevelEnable
    return leafs
}

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetBundleName() string { return "cisco_ios_xr" }

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetYangName() string { return "attachment" }

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) SetParent(parent types.Entity) { attachment.parent = parent }

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetParent() types.Entity { return attachment.parent }

func (attachment *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Attachment) GetParentYangName() string { return "span-monitor-session" }

// DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl
// Enable ACL matching for traffic mirroring
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ACL. The type is interface{}. This attribute is mandatory.
    AclEnable interface{}

    // ACL Name. The type is string with length: 1..80.
    AclName interface{}
}

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetFilter() yfilter.YFilter { return acl.YFilter }

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) SetFilter(yf yfilter.YFilter) { acl.YFilter = yf }

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetGoName(yname string) string {
    if yname == "acl-enable" { return "AclEnable" }
    if yname == "acl-name" { return "AclName" }
    return ""
}

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetSegmentPath() string {
    return "acl"
}

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["acl-enable"] = acl.AclEnable
    leafs["acl-name"] = acl.AclName
    return leafs
}

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetBundleName() string { return "cisco_ios_xr" }

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetYangName() string { return "acl" }

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) SetParent(parent types.Entity) { acl.parent = parent }

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetParent() types.Entity { return acl.parent }

func (acl *DynamicTemplate_IpSubscribers_IpSubscriber_SpanMonitorSessions_SpanMonitorSession_Acl) GetParentYangName() string { return "span-monitor-session" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter
// IPv4 Packet Filtering configuration for the
// template
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound

    // IPv4 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound
}

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetFilter() yfilter.YFilter { return ipv4PacketFilter.YFilter }

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) SetFilter(yf yfilter.YFilter) { ipv4PacketFilter.YFilter = yf }

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetGoName(yname string) string {
    if yname == "outbound" { return "Outbound" }
    if yname == "inbound" { return "Inbound" }
    return ""
}

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter"
}

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "outbound" {
        return &ipv4PacketFilter.Outbound
    }
    if childYangName == "inbound" {
        return &ipv4PacketFilter.Inbound
    }
    return nil
}

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["outbound"] = &ipv4PacketFilter.Outbound
    children["inbound"] = &ipv4PacketFilter.Inbound
    return children
}

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetYangName() string { return "ipv4-packet-filter" }

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) SetParent(parent types.Entity) { ipv4PacketFilter.parent = parent }

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetParent() types.Entity { return ipv4PacketFilter.parent }

func (ipv4PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound
// IPv4 Packet filter to be applied to outbound
// packets
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..65. This attribute is mandatory.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetFilter() yfilter.YFilter { return outbound.YFilter }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) SetFilter(yf yfilter.YFilter) { outbound.YFilter = yf }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "hardware-count" { return "HardwareCount" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetSegmentPath() string {
    return "outbound"
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = outbound.CommonAclName
    leafs["name"] = outbound.Name
    leafs["hardware-count"] = outbound.HardwareCount
    leafs["interface-statistics"] = outbound.InterfaceStatistics
    return leafs
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetBundleName() string { return "cisco_ios_xr" }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetYangName() string { return "outbound" }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) SetParent(parent types.Entity) { outbound.parent = parent }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetParent() types.Entity { return outbound.parent }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Outbound) GetParentYangName() string { return "ipv4-packet-filter" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound
// IPv4 Packet filter to be applied to inbound
// packets
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Inbound packets NOTE: This
    // parameter is mandatory if 'CommonACLName' is not specified. The type is
    // string with length: 1..65.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetFilter() yfilter.YFilter { return inbound.YFilter }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) SetFilter(yf yfilter.YFilter) { inbound.YFilter = yf }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "hardware-count" { return "HardwareCount" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetSegmentPath() string {
    return "inbound"
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = inbound.CommonAclName
    leafs["name"] = inbound.Name
    leafs["hardware-count"] = inbound.HardwareCount
    leafs["interface-statistics"] = inbound.InterfaceStatistics
    return leafs
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetBundleName() string { return "cisco_ios_xr" }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetYangName() string { return "inbound" }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) SetParent(parent types.Entity) { inbound.parent = parent }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetParent() types.Entity { return inbound.parent }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4PacketFilter_Inbound) GetParentYangName() string { return "ipv4-packet-filter" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter
// IPv6 Packet Filtering configuration for the
// interface
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound

    // IPv6 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound
}

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetFilter() yfilter.YFilter { return ipv6PacketFilter.YFilter }

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) SetFilter(yf yfilter.YFilter) { ipv6PacketFilter.YFilter = yf }

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetGoName(yname string) string {
    if yname == "inbound" { return "Inbound" }
    if yname == "outbound" { return "Outbound" }
    return ""
}

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter"
}

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inbound" {
        return &ipv6PacketFilter.Inbound
    }
    if childYangName == "outbound" {
        return &ipv6PacketFilter.Outbound
    }
    return nil
}

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["inbound"] = &ipv6PacketFilter.Inbound
    children["outbound"] = &ipv6PacketFilter.Outbound
    return children
}

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetYangName() string { return "ipv6-packet-filter" }

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) SetParent(parent types.Entity) { ipv6PacketFilter.parent = parent }

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetParent() types.Entity { return ipv6PacketFilter.parent }

func (ipv6PacketFilter *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound
// IPv6 Packet filter to be applied to inbound
// packets
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Inbound  NOTE: This parameter is
    // mandatory if 'CommonACLName' is not specified. The type is string with
    // length: 1..65.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetFilter() yfilter.YFilter { return inbound.YFilter }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) SetFilter(yf yfilter.YFilter) { inbound.YFilter = yf }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetSegmentPath() string {
    return "inbound"
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = inbound.CommonAclName
    leafs["name"] = inbound.Name
    leafs["interface-statistics"] = inbound.InterfaceStatistics
    return leafs
}

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetBundleName() string { return "cisco_ios_xr" }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetYangName() string { return "inbound" }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) SetParent(parent types.Entity) { inbound.parent = parent }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetParent() types.Entity { return inbound.parent }

func (inbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Inbound) GetParentYangName() string { return "ipv6-packet-filter" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound
// IPv6 Packet filter to be applied to outbound
// packets
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..65. This attribute is mandatory.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetFilter() yfilter.YFilter { return outbound.YFilter }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) SetFilter(yf yfilter.YFilter) { outbound.YFilter = yf }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetSegmentPath() string {
    return "outbound"
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = outbound.CommonAclName
    leafs["name"] = outbound.Name
    leafs["interface-statistics"] = outbound.InterfaceStatistics
    return leafs
}

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetBundleName() string { return "cisco_ios_xr" }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetYangName() string { return "outbound" }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) SetParent(parent types.Entity) { outbound.parent = parent }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetParent() types.Entity { return outbound.parent }

func (outbound *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6PacketFilter_Outbound) GetParentYangName() string { return "ipv6-packet-filter" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd
// Interface dhcpv4 configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The class to be used for proxy/server profile. The type is string.
    Class interface{}

    // The Default Gateway IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DefaultGateway interface{}

    // The pool to be used for Prefix Delegation. The type is interface{} with
    // range: -2147483648..2147483647.
    SessionLimit interface{}

    // Cisco VSA to configure any dhcp4 option per subscriber. The type is string.
    Dhcpv4Option interface{}
}

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetFilter() yfilter.YFilter { return dhcpd.YFilter }

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) SetFilter(yf yfilter.YFilter) { dhcpd.YFilter = yf }

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    if yname == "default-gateway" { return "DefaultGateway" }
    if yname == "session-limit" { return "SessionLimit" }
    if yname == "dhcpv4-option" { return "Dhcpv4Option" }
    return ""
}

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-dhcpd-subscriber-cfg:dhcpd"
}

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class"] = dhcpd.Class
    leafs["default-gateway"] = dhcpd.DefaultGateway
    leafs["session-limit"] = dhcpd.SessionLimit
    leafs["dhcpv4-option"] = dhcpd.Dhcpv4Option
    return leafs
}

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetYangName() string { return "dhcpd" }

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) SetParent(parent types.Entity) { dhcpd.parent = parent }

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetParent() types.Entity { return dhcpd.parent }

func (dhcpd *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpd) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Igmp
// IGMPconfiguration
type DynamicTemplate_IpSubscribers_IpSubscriber_Igmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default VRF.
    DefaultVrf DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf
}

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetFilter() yfilter.YFilter { return igmp.YFilter }

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) SetFilter(yf yfilter.YFilter) { igmp.YFilter = yf }

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetGoName(yname string) string {
    if yname == "default-vrf" { return "DefaultVrf" }
    return ""
}

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-igmp-dyn-tmpl-cfg:igmp"
}

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-vrf" {
        return &igmp.DefaultVrf
    }
    return nil
}

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-vrf"] = &igmp.DefaultVrf
    return children
}

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetBundleName() string { return "cisco_ios_xr" }

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetYangName() string { return "igmp" }

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) SetParent(parent types.Entity) { igmp.parent = parent }

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetParent() types.Entity { return igmp.parent }

func (igmp *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf
// Default VRF
type DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf struct {
    parent types.Entity
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

    // Query response value in seconds. The type is interface{} with range: 1..12.
    // Units are second. The default value is 10.
    QueryMaxResponseTime interface{}

    // Configure Multicast mode variable. The type is DynTmplMulticastMode.
    MulticastMode interface{}

    // IGMPv3 explicit host tracking.
    ExplicitTracking DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking
}

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetFilter() yfilter.YFilter { return defaultVrf.YFilter }

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) SetFilter(yf yfilter.YFilter) { defaultVrf.YFilter = yf }

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetGoName(yname string) string {
    if yname == "max-groups" { return "MaxGroups" }
    if yname == "access-group" { return "AccessGroup" }
    if yname == "version" { return "Version" }
    if yname == "query-interval" { return "QueryInterval" }
    if yname == "query-max-response-time" { return "QueryMaxResponseTime" }
    if yname == "multicast-mode" { return "MulticastMode" }
    if yname == "explicit-tracking" { return "ExplicitTracking" }
    return ""
}

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetSegmentPath() string {
    return "default-vrf"
}

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "explicit-tracking" {
        return &defaultVrf.ExplicitTracking
    }
    return nil
}

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["explicit-tracking"] = &defaultVrf.ExplicitTracking
    return children
}

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-groups"] = defaultVrf.MaxGroups
    leafs["access-group"] = defaultVrf.AccessGroup
    leafs["version"] = defaultVrf.Version
    leafs["query-interval"] = defaultVrf.QueryInterval
    leafs["query-max-response-time"] = defaultVrf.QueryMaxResponseTime
    leafs["multicast-mode"] = defaultVrf.MulticastMode
    return leafs
}

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetBundleName() string { return "cisco_ios_xr" }

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetYangName() string { return "default-vrf" }

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) SetParent(parent types.Entity) { defaultVrf.parent = parent }

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetParent() types.Entity { return defaultVrf.parent }

func (defaultVrf *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf) GetParentYangName() string { return "igmp" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking
// IGMPv3 explicit host tracking
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable or disable, when value is TRUE or FALSE respectively. The type is
    // bool. This attribute is mandatory.
    Enable interface{}

    // Access list specifying tracking group range. The type is string with
    // length: 1..64.
    AccessListName interface{}
}

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetFilter() yfilter.YFilter { return explicitTracking.YFilter }

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) SetFilter(yf yfilter.YFilter) { explicitTracking.YFilter = yf }

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetSegmentPath() string {
    return "explicit-tracking"
}

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = explicitTracking.Enable
    leafs["access-list-name"] = explicitTracking.AccessListName
    return leafs
}

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetBundleName() string { return "cisco_ios_xr" }

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetYangName() string { return "explicit-tracking" }

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) SetParent(parent types.Entity) { explicitTracking.parent = parent }

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetParent() types.Entity { return explicitTracking.parent }

func (explicitTracking *DynamicTemplate_IpSubscribers_IpSubscriber_Igmp_DefaultVrf_ExplicitTracking) GetParentYangName() string { return "default-vrf" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network
// Interface IPv4 Network configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network struct {
    parent types.Entity
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

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetFilter() yfilter.YFilter { return ipv4Network.YFilter }

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) SetFilter(yf yfilter.YFilter) { ipv4Network.YFilter = yf }

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetGoName(yname string) string {
    if yname == "unnumbered" { return "Unnumbered" }
    if yname == "mtu" { return "Mtu" }
    if yname == "unreachables" { return "Unreachables" }
    if yname == "rpf" { return "Rpf" }
    return ""
}

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network"
}

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unnumbered"] = ipv4Network.Unnumbered
    leafs["mtu"] = ipv4Network.Mtu
    leafs["unreachables"] = ipv4Network.Unreachables
    leafs["rpf"] = ipv4Network.Rpf
    return leafs
}

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetYangName() string { return "ipv4-network" }

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) SetParent(parent types.Entity) { ipv4Network.parent = parent }

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetParent() types.Entity { return ipv4Network.parent }

func (ipv4Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv4Network) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network
// Interface IPv6 Network configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network struct {
    parent types.Entity
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

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetFilter() yfilter.YFilter { return ipv6Network.YFilter }

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) SetFilter(yf yfilter.YFilter) { ipv6Network.YFilter = yf }

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetGoName(yname string) string {
    if yname == "mtu" { return "Mtu" }
    if yname == "rpf" { return "Rpf" }
    if yname == "unreachables" { return "Unreachables" }
    if yname == "addresses" { return "Addresses" }
    return ""
}

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network"
}

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &ipv6Network.Addresses
    }
    return nil
}

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &ipv6Network.Addresses
    return children
}

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mtu"] = ipv6Network.Mtu
    leafs["rpf"] = ipv6Network.Rpf
    leafs["unreachables"] = ipv6Network.Unreachables
    return leafs
}

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetYangName() string { return "ipv6-network" }

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) SetParent(parent types.Entity) { ipv6Network.parent = parent }

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetParent() types.Entity { return ipv6Network.parent }

func (ipv6Network *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses
// Set the IPv6 address of an interface
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Auto IPv6 Interface Configuration.
    AutoConfiguration DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration
}

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetGoName(yname string) string {
    if yname == "auto-configuration" { return "AutoConfiguration" }
    return ""
}

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto-configuration" {
        return &addresses.AutoConfiguration
    }
    return nil
}

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto-configuration"] = &addresses.AutoConfiguration
    return children
}

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetYangName() string { return "addresses" }

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses) GetParentYangName() string { return "ipv6-network" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration
// Auto IPv6 Interface Configuration
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The flag to enable auto ipv6 interface configuration. The type is
    // interface{}.
    Enable interface{}
}

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetFilter() yfilter.YFilter { return autoConfiguration.YFilter }

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) SetFilter(yf yfilter.YFilter) { autoConfiguration.YFilter = yf }

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetSegmentPath() string {
    return "auto-configuration"
}

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = autoConfiguration.Enable
    return leafs
}

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetYangName() string { return "auto-configuration" }

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) SetParent(parent types.Entity) { autoConfiguration.parent = parent }

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetParent() types.Entity { return autoConfiguration.parent }

func (autoConfiguration *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Network_Addresses_AutoConfiguration) GetParentYangName() string { return "addresses" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor
// Interface IPv6 Network configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor struct {
    parent types.Entity
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

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetFilter() yfilter.YFilter { return ipv6Neighbor.YFilter }

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) SetFilter(yf yfilter.YFilter) { ipv6Neighbor.YFilter = yf }

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetGoName(yname string) string {
    if yname == "framed-prefix-pool" { return "FramedPrefixPool" }
    if yname == "managed-config" { return "ManagedConfig" }
    if yname == "other-config" { return "OtherConfig" }
    if yname == "start-ra-on-ipv6-enable" { return "StartRaOnIpv6Enable" }
    if yname == "nud-enable" { return "NudEnable" }
    if yname == "ra-lifetime" { return "RaLifetime" }
    if yname == "router-preference" { return "RouterPreference" }
    if yname == "ra-suppress" { return "RaSuppress" }
    if yname == "ra-unicast" { return "RaUnicast" }
    if yname == "ra-unspecify-hoplimit" { return "RaUnspecifyHoplimit" }
    if yname == "ra-suppress-mtu" { return "RaSuppressMtu" }
    if yname == "suppress-cache-learning" { return "SuppressCacheLearning" }
    if yname == "reachable-time" { return "ReachableTime" }
    if yname == "ns-interval" { return "NsInterval" }
    if yname == "ra-interval" { return "RaInterval" }
    if yname == "framed-prefix" { return "FramedPrefix" }
    if yname == "duplicate-address-detection" { return "DuplicateAddressDetection" }
    if yname == "ra-initial" { return "RaInitial" }
    return ""
}

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor"
}

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ra-interval" {
        return &ipv6Neighbor.RaInterval
    }
    if childYangName == "framed-prefix" {
        return &ipv6Neighbor.FramedPrefix
    }
    if childYangName == "duplicate-address-detection" {
        return &ipv6Neighbor.DuplicateAddressDetection
    }
    if childYangName == "ra-initial" {
        return &ipv6Neighbor.RaInitial
    }
    return nil
}

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ra-interval"] = &ipv6Neighbor.RaInterval
    children["framed-prefix"] = &ipv6Neighbor.FramedPrefix
    children["duplicate-address-detection"] = &ipv6Neighbor.DuplicateAddressDetection
    children["ra-initial"] = &ipv6Neighbor.RaInitial
    return children
}

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["framed-prefix-pool"] = ipv6Neighbor.FramedPrefixPool
    leafs["managed-config"] = ipv6Neighbor.ManagedConfig
    leafs["other-config"] = ipv6Neighbor.OtherConfig
    leafs["start-ra-on-ipv6-enable"] = ipv6Neighbor.StartRaOnIpv6Enable
    leafs["nud-enable"] = ipv6Neighbor.NudEnable
    leafs["ra-lifetime"] = ipv6Neighbor.RaLifetime
    leafs["router-preference"] = ipv6Neighbor.RouterPreference
    leafs["ra-suppress"] = ipv6Neighbor.RaSuppress
    leafs["ra-unicast"] = ipv6Neighbor.RaUnicast
    leafs["ra-unspecify-hoplimit"] = ipv6Neighbor.RaUnspecifyHoplimit
    leafs["ra-suppress-mtu"] = ipv6Neighbor.RaSuppressMtu
    leafs["suppress-cache-learning"] = ipv6Neighbor.SuppressCacheLearning
    leafs["reachable-time"] = ipv6Neighbor.ReachableTime
    leafs["ns-interval"] = ipv6Neighbor.NsInterval
    return leafs
}

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetYangName() string { return "ipv6-neighbor" }

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) SetParent(parent types.Entity) { ipv6Neighbor.parent = parent }

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetParent() types.Entity { return ipv6Neighbor.parent }

func (ipv6Neighbor *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval
// Set IPv6 Router Advertisement (RA) interval in
// seconds
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Maximum interface{}

    // Minimum RA interval in seconds. Must be less than 0.75 * maximum interval.
    // The type is interface{} with range: 3..1800. Units are second.
    Minimum interface{}
}

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetFilter() yfilter.YFilter { return raInterval.YFilter }

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) SetFilter(yf yfilter.YFilter) { raInterval.YFilter = yf }

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "minimum" { return "Minimum" }
    return ""
}

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetSegmentPath() string {
    return "ra-interval"
}

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = raInterval.Maximum
    leafs["minimum"] = raInterval.Minimum
    return leafs
}

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetBundleName() string { return "cisco_ios_xr" }

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetYangName() string { return "ra-interval" }

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) SetParent(parent types.Entity) { raInterval.parent = parent }

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetParent() types.Entity { return raInterval.parent }

func (raInterval *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInterval) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix
// Set the IPv6 framed ipv6 prefix for a
// subscriber interface 
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 framed prefix length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}

    // IPV6 framed prefix address. The type is string. This attribute is
    // mandatory.
    Prefix interface{}
}

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetFilter() yfilter.YFilter { return framedPrefix.YFilter }

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) SetFilter(yf yfilter.YFilter) { framedPrefix.YFilter = yf }

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetSegmentPath() string {
    return "framed-prefix"
}

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = framedPrefix.PrefixLength
    leafs["prefix"] = framedPrefix.Prefix
    return leafs
}

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetYangName() string { return "framed-prefix" }

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) SetParent(parent types.Entity) { framedPrefix.parent = parent }

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetParent() types.Entity { return framedPrefix.parent }

func (framedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_FramedPrefix) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection
// Duplicate Address Detection (DAD)
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set IPv6 duplicate address detection transmits. The type is interface{}
    // with range: 0..600.
    Attempts interface{}
}

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetFilter() yfilter.YFilter { return duplicateAddressDetection.YFilter }

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) SetFilter(yf yfilter.YFilter) { duplicateAddressDetection.YFilter = yf }

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetGoName(yname string) string {
    if yname == "attempts" { return "Attempts" }
    return ""
}

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetSegmentPath() string {
    return "duplicate-address-detection"
}

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attempts"] = duplicateAddressDetection.Attempts
    return leafs
}

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetBundleName() string { return "cisco_ios_xr" }

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetYangName() string { return "duplicate-address-detection" }

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) SetParent(parent types.Entity) { duplicateAddressDetection.parent = parent }

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetParent() types.Entity { return duplicateAddressDetection.parent }

func (duplicateAddressDetection *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_DuplicateAddressDetection) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial
// IPv6 ND RA Initial
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Initial RA count. The type is interface{} with range: 0..32. This attribute
    // is mandatory.
    Count interface{}

    // Initial RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Interval interface{}
}

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetFilter() yfilter.YFilter { return raInitial.YFilter }

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) SetFilter(yf yfilter.YFilter) { raInitial.YFilter = yf }

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetGoName(yname string) string {
    if yname == "count" { return "Count" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetSegmentPath() string {
    return "ra-initial"
}

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["count"] = raInitial.Count
    leafs["interval"] = raInitial.Interval
    return leafs
}

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetBundleName() string { return "cisco_ios_xr" }

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetYangName() string { return "ra-initial" }

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) SetParent(parent types.Entity) { raInitial.parent = parent }

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetParent() types.Entity { return raInitial.parent }

func (raInitial *DynamicTemplate_IpSubscribers_IpSubscriber_Ipv6Neighbor_RaInitial) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6
// Interface dhcpv6 configuration data
type DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dns IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DnsIpv6Address interface{}

    // Select proxy/server profile based on mode class name. The type is string.
    ModeClass interface{}

    // The pool to be used for Address assignment. The type is string.
    AddressPool interface{}

    // The pool to be used for Prefix Delegation. The type is string.
    DelegatedPrefixPool interface{}

    // The class to be used for proxy/server profile. The type is string.
    Class interface{}

    // Stateful IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StatefulAddress interface{}

    // The prefix to be used for Prefix Delegation.
    DelegatedPrefix DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix
}

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetFilter() yfilter.YFilter { return dhcpv6.YFilter }

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) SetFilter(yf yfilter.YFilter) { dhcpv6.YFilter = yf }

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetGoName(yname string) string {
    if yname == "dns-ipv6address" { return "DnsIpv6Address" }
    if yname == "mode-class" { return "ModeClass" }
    if yname == "address-pool" { return "AddressPool" }
    if yname == "delegated-prefix-pool" { return "DelegatedPrefixPool" }
    if yname == "class" { return "Class" }
    if yname == "stateful-address" { return "StatefulAddress" }
    if yname == "delegated-prefix" { return "DelegatedPrefix" }
    return ""
}

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-new-dhcpv6d-subscriber-cfg:dhcpv6"
}

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "delegated-prefix" {
        return &dhcpv6.DelegatedPrefix
    }
    return nil
}

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["delegated-prefix"] = &dhcpv6.DelegatedPrefix
    return children
}

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dns-ipv6address"] = dhcpv6.DnsIpv6Address
    leafs["mode-class"] = dhcpv6.ModeClass
    leafs["address-pool"] = dhcpv6.AddressPool
    leafs["delegated-prefix-pool"] = dhcpv6.DelegatedPrefixPool
    leafs["class"] = dhcpv6.Class
    leafs["stateful-address"] = dhcpv6.StatefulAddress
    return leafs
}

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetYangName() string { return "dhcpv6" }

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) SetParent(parent types.Entity) { dhcpv6.parent = parent }

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetParent() types.Entity { return dhcpv6.parent }

func (dhcpv6 *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix
// The prefix to be used for Prefix Delegation
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Prefix interface{}

    // PD Prefix Length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}
}

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetFilter() yfilter.YFilter { return delegatedPrefix.YFilter }

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) SetFilter(yf yfilter.YFilter) { delegatedPrefix.YFilter = yf }

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetSegmentPath() string {
    return "delegated-prefix"
}

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = delegatedPrefix.Prefix
    leafs["prefix-length"] = delegatedPrefix.PrefixLength
    return leafs
}

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetYangName() string { return "delegated-prefix" }

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) SetParent(parent types.Entity) { delegatedPrefix.parent = parent }

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetParent() types.Entity { return delegatedPrefix.parent }

func (delegatedPrefix *DynamicTemplate_IpSubscribers_IpSubscriber_Dhcpv6_DelegatedPrefix) GetParentYangName() string { return "dhcpv6" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Pbr
// Dynamic Template PBR configuration
type DynamicTemplate_IpSubscribers_IpSubscriber_Pbr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Class for subscriber ingress policy. The type is string.
    ServicePolicyIn interface{}

    // PBR service policy configuration.
    ServicePolicy DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy
}

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetFilter() yfilter.YFilter { return pbr.YFilter }

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) SetFilter(yf yfilter.YFilter) { pbr.YFilter = yf }

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetGoName(yname string) string {
    if yname == "service-policy-in" { return "ServicePolicyIn" }
    if yname == "service-policy" { return "ServicePolicy" }
    return ""
}

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetSegmentPath() string {
    return "Cisco-IOS-XR-pbr-subscriber-cfg:pbr"
}

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-policy" {
        return &pbr.ServicePolicy
    }
    return nil
}

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-policy"] = &pbr.ServicePolicy
    return children
}

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["service-policy-in"] = pbr.ServicePolicyIn
    return leafs
}

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetBundleName() string { return "cisco_ios_xr" }

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetYangName() string { return "pbr" }

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) SetParent(parent types.Entity) { pbr.parent = parent }

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetParent() types.Entity { return pbr.parent }

func (pbr *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy
// PBR service policy configuration
type DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ingress service policy. The type is string.
    Input interface{}
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetFilter() yfilter.YFilter { return servicePolicy.YFilter }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) SetFilter(yf yfilter.YFilter) { servicePolicy.YFilter = yf }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetSegmentPath() string {
    return "service-policy"
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input"] = servicePolicy.Input
    return leafs
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetYangName() string { return "service-policy" }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) SetParent(parent types.Entity) { servicePolicy.parent = parent }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetParent() types.Entity { return servicePolicy.parent }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Pbr_ServicePolicy) GetParentYangName() string { return "pbr" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos
// QoS dynamically applied configuration template
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service policy to be applied in ingress/egress direction.
    ServicePolicy DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy

    // QoS L2 overhead accounting.
    Account DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account

    // QoS to be applied in egress direction.
    Output DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output
}

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetFilter() yfilter.YFilter { return qos.YFilter }

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) SetFilter(yf yfilter.YFilter) { qos.YFilter = yf }

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetGoName(yname string) string {
    if yname == "service-policy" { return "ServicePolicy" }
    if yname == "account" { return "Account" }
    if yname == "output" { return "Output" }
    return ""
}

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetSegmentPath() string {
    return "Cisco-IOS-XR-qos-ma-bng-cfg:qos"
}

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-policy" {
        return &qos.ServicePolicy
    }
    if childYangName == "account" {
        return &qos.Account
    }
    if childYangName == "output" {
        return &qos.Output
    }
    return nil
}

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-policy"] = &qos.ServicePolicy
    children["account"] = &qos.Account
    children["output"] = &qos.Output
    return children
}

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetBundleName() string { return "cisco_ios_xr" }

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetYangName() string { return "qos" }

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) SetParent(parent types.Entity) { qos.parent = parent }

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetParent() types.Entity { return qos.parent }

func (qos *DynamicTemplate_IpSubscribers_IpSubscriber_Qos) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy
// Service policy to be applied in ingress/egress
// direction
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber ingress policy.
    Input DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input

    // Subscriber egress policy.
    Output DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetFilter() yfilter.YFilter { return servicePolicy.YFilter }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) SetFilter(yf yfilter.YFilter) { servicePolicy.YFilter = yf }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetSegmentPath() string {
    return "service-policy"
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &servicePolicy.Input
    }
    if childYangName == "output" {
        return &servicePolicy.Output
    }
    return nil
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &servicePolicy.Input
    children["output"] = &servicePolicy.Output
    return children
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetYangName() string { return "service-policy" }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) SetParent(parent types.Entity) { servicePolicy.parent = parent }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetParent() types.Entity { return servicePolicy.parent }

func (servicePolicy *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy) GetParentYangName() string { return "qos" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input
// Subscriber ingress policy
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "spi-name" { return "SpiName" }
    if yname == "merge" { return "Merge" }
    if yname == "merge-id" { return "MergeId" }
    if yname == "account-stats" { return "AccountStats" }
    return ""
}

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetSegmentPath() string {
    return "input"
}

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = input.PolicyName
    leafs["spi-name"] = input.SpiName
    leafs["merge"] = input.Merge
    leafs["merge-id"] = input.MergeId
    leafs["account-stats"] = input.AccountStats
    return leafs
}

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetYangName() string { return "input" }

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetParent() types.Entity { return input.parent }

func (input *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Input) GetParentYangName() string { return "service-policy" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output
// Subscriber egress policy
// This type is a presence type.
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "spi-name" { return "SpiName" }
    if yname == "merge" { return "Merge" }
    if yname == "merge-id" { return "MergeId" }
    if yname == "account-stats" { return "AccountStats" }
    return ""
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetSegmentPath() string {
    return "output"
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = output.PolicyName
    leafs["spi-name"] = output.SpiName
    leafs["merge"] = output.Merge
    leafs["merge-id"] = output.MergeId
    leafs["account-stats"] = output.AccountStats
    return leafs
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetYangName() string { return "output" }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetParent() types.Entity { return output.parent }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_ServicePolicy_Output) GetParentYangName() string { return "service-policy" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account
// QoS L2 overhead accounting
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account struct {
    parent types.Entity
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

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetFilter() yfilter.YFilter { return account.YFilter }

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) SetFilter(yf yfilter.YFilter) { account.YFilter = yf }

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetGoName(yname string) string {
    if yname == "aal" { return "Aal" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "atm-cell-tax" { return "AtmCellTax" }
    if yname == "user-defined" { return "UserDefined" }
    return ""
}

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetSegmentPath() string {
    return "account"
}

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aal"] = account.Aal
    leafs["encapsulation"] = account.Encapsulation
    leafs["atm-cell-tax"] = account.AtmCellTax
    leafs["user-defined"] = account.UserDefined
    return leafs
}

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetBundleName() string { return "cisco_ios_xr" }

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetYangName() string { return "account" }

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) SetParent(parent types.Entity) { account.parent = parent }

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetParent() types.Entity { return account.parent }

func (account *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Account) GetParentYangName() string { return "qos" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output
// QoS to be applied in egress direction
type DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum bandwidth value for the subscriber (in kbps). The type is
    // interface{} with range: 1..4294967295. Units are kbit/s.
    MinimumBandwidth interface{}
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetGoName(yname string) string {
    if yname == "minimum-bandwidth" { return "MinimumBandwidth" }
    return ""
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetSegmentPath() string {
    return "output"
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minimum-bandwidth"] = output.MinimumBandwidth
    return leafs
}

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetYangName() string { return "output" }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetParent() types.Entity { return output.parent }

func (output *DynamicTemplate_IpSubscribers_IpSubscriber_Qos_Output) GetParentYangName() string { return "qos" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Accounting
// Subscriber accounting dynamic-template commands
type DynamicTemplate_IpSubscribers_IpSubscriber_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber accounting prepaid feature. The type is string.
    PrepaidFeature interface{}

    // Subscriber accounting service accounting.
    ServiceAccounting DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting

    // Subscriber accounting session accounting.
    Session DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session

    // Subscriber accounting idle timeout.
    IdleTimeout DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout
}

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetGoName(yname string) string {
    if yname == "prepaid-feature" { return "PrepaidFeature" }
    if yname == "service-accounting" { return "ServiceAccounting" }
    if yname == "session" { return "Session" }
    if yname == "idle-timeout" { return "IdleTimeout" }
    return ""
}

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-accounting-cfg:accounting"
}

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-accounting" {
        return &accounting.ServiceAccounting
    }
    if childYangName == "session" {
        return &accounting.Session
    }
    if childYangName == "idle-timeout" {
        return &accounting.IdleTimeout
    }
    return nil
}

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-accounting"] = &accounting.ServiceAccounting
    children["session"] = &accounting.Session
    children["idle-timeout"] = &accounting.IdleTimeout
    return children
}

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prepaid-feature"] = accounting.PrepaidFeature
    return leafs
}

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetYangName() string { return "accounting" }

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting) GetParentYangName() string { return "ip-subscriber" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting
// Subscriber accounting service accounting
type DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service accounting method list name. The type is string.
    MethodListName interface{}

    // Accounting interim interval in minutes. The type is interface{} with range:
    // -2147483648..2147483647. Units are minute.
    AccountingInterimInterval interface{}
}

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetFilter() yfilter.YFilter { return serviceAccounting.YFilter }

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) SetFilter(yf yfilter.YFilter) { serviceAccounting.YFilter = yf }

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetGoName(yname string) string {
    if yname == "method-list-name" { return "MethodListName" }
    if yname == "accounting-interim-interval" { return "AccountingInterimInterval" }
    return ""
}

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetSegmentPath() string {
    return "service-accounting"
}

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method-list-name"] = serviceAccounting.MethodListName
    leafs["accounting-interim-interval"] = serviceAccounting.AccountingInterimInterval
    return leafs
}

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetBundleName() string { return "cisco_ios_xr" }

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetYangName() string { return "service-accounting" }

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) SetParent(parent types.Entity) { serviceAccounting.parent = parent }

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetParent() types.Entity { return serviceAccounting.parent }

func (serviceAccounting *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_ServiceAccounting) GetParentYangName() string { return "accounting" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session
// Subscriber accounting session accounting
type DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session accounting method list name. The type is string.
    MethodListName interface{}

    // Interim accounting interval in minutes. The type is interface{} with range:
    // -2147483648..2147483647. Units are minute.
    PeriodicInterval interface{}

    // Dual stack wait delay in seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    DualStackDelay interface{}

    // Hold Accounting start based on IA_PD. The type is interface{} with range:
    // -2147483648..2147483647.
    HoldAcctStart interface{}
}

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetGoName(yname string) string {
    if yname == "method-list-name" { return "MethodListName" }
    if yname == "periodic-interval" { return "PeriodicInterval" }
    if yname == "dual-stack-delay" { return "DualStackDelay" }
    if yname == "hold-acct-start" { return "HoldAcctStart" }
    return ""
}

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetSegmentPath() string {
    return "session"
}

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method-list-name"] = session.MethodListName
    leafs["periodic-interval"] = session.PeriodicInterval
    leafs["dual-stack-delay"] = session.DualStackDelay
    leafs["hold-acct-start"] = session.HoldAcctStart
    return leafs
}

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetYangName() string { return "session" }

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetParent() types.Entity { return session.parent }

func (session *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_Session) GetParentYangName() string { return "accounting" }

// DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout
// Subscriber accounting idle timeout
type DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout struct {
    parent types.Entity
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

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetFilter() yfilter.YFilter { return idleTimeout.YFilter }

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) SetFilter(yf yfilter.YFilter) { idleTimeout.YFilter = yf }

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetGoName(yname string) string {
    if yname == "timeout-value" { return "TimeoutValue" }
    if yname == "threshold" { return "Threshold" }
    if yname == "direction" { return "Direction" }
    return ""
}

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetSegmentPath() string {
    return "idle-timeout"
}

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout-value"] = idleTimeout.TimeoutValue
    leafs["threshold"] = idleTimeout.Threshold
    leafs["direction"] = idleTimeout.Direction
    return leafs
}

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetBundleName() string { return "cisco_ios_xr" }

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetYangName() string { return "idle-timeout" }

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) SetParent(parent types.Entity) { idleTimeout.parent = parent }

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetParent() types.Entity { return idleTimeout.parent }

func (idleTimeout *DynamicTemplate_IpSubscribers_IpSubscriber_Accounting_IdleTimeout) GetParentYangName() string { return "accounting" }

// DynamicTemplate_SubscriberServices
// The Service Type Template Table
type DynamicTemplate_SubscriberServices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Service Type Template . The type is slice of
    // DynamicTemplate_SubscriberServices_SubscriberService.
    SubscriberService []DynamicTemplate_SubscriberServices_SubscriberService
}

func (subscriberServices *DynamicTemplate_SubscriberServices) GetFilter() yfilter.YFilter { return subscriberServices.YFilter }

func (subscriberServices *DynamicTemplate_SubscriberServices) SetFilter(yf yfilter.YFilter) { subscriberServices.YFilter = yf }

func (subscriberServices *DynamicTemplate_SubscriberServices) GetGoName(yname string) string {
    if yname == "subscriber-service" { return "SubscriberService" }
    return ""
}

func (subscriberServices *DynamicTemplate_SubscriberServices) GetSegmentPath() string {
    return "subscriber-services"
}

func (subscriberServices *DynamicTemplate_SubscriberServices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subscriber-service" {
        for _, c := range subscriberServices.SubscriberService {
            if subscriberServices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DynamicTemplate_SubscriberServices_SubscriberService{}
        subscriberServices.SubscriberService = append(subscriberServices.SubscriberService, child)
        return &subscriberServices.SubscriberService[len(subscriberServices.SubscriberService)-1]
    }
    return nil
}

func (subscriberServices *DynamicTemplate_SubscriberServices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subscriberServices.SubscriberService {
        children[subscriberServices.SubscriberService[i].GetSegmentPath()] = &subscriberServices.SubscriberService[i]
    }
    return children
}

func (subscriberServices *DynamicTemplate_SubscriberServices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberServices *DynamicTemplate_SubscriberServices) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberServices *DynamicTemplate_SubscriberServices) GetYangName() string { return "subscriber-services" }

func (subscriberServices *DynamicTemplate_SubscriberServices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberServices *DynamicTemplate_SubscriberServices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberServices *DynamicTemplate_SubscriberServices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberServices *DynamicTemplate_SubscriberServices) SetParent(parent types.Entity) { subscriberServices.parent = parent }

func (subscriberServices *DynamicTemplate_SubscriberServices) GetParent() types.Entity { return subscriberServices.parent }

func (subscriberServices *DynamicTemplate_SubscriberServices) GetParentYangName() string { return "dynamic-template" }

// DynamicTemplate_SubscriberServices_SubscriberService
// A Service Type Template 
type DynamicTemplate_SubscriberServices_SubscriberService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the template. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

    // Dynamic Template PBR configuration.
    Pbr DynamicTemplate_SubscriberServices_SubscriberService_Pbr

    // QoS dynamically applied configuration template.
    Qos DynamicTemplate_SubscriberServices_SubscriberService_Qos

    // Subscriber accounting dynamic-template commands.
    Accounting DynamicTemplate_SubscriberServices_SubscriberService_Accounting
}

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetFilter() yfilter.YFilter { return subscriberService.YFilter }

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) SetFilter(yf yfilter.YFilter) { subscriberService.YFilter = yf }

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetGoName(yname string) string {
    if yname == "template-name" { return "TemplateName" }
    if yname == "vrf" { return "Vrf" }
    if yname == "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions" { return "SpanMonitorSessions" }
    if yname == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter" { return "Ipv4PacketFilter" }
    if yname == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter" { return "Ipv6PacketFilter" }
    if yname == "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network" { return "Ipv4Network" }
    if yname == "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network" { return "Ipv6Network" }
    if yname == "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor" { return "Ipv6Neighbor" }
    if yname == "Cisco-IOS-XR-pbr-subscriber-cfg:pbr" { return "Pbr" }
    if yname == "Cisco-IOS-XR-qos-ma-bng-cfg:qos" { return "Qos" }
    if yname == "Cisco-IOS-XR-subscriber-accounting-cfg:accounting" { return "Accounting" }
    return ""
}

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetSegmentPath() string {
    return "subscriber-service" + "[template-name='" + fmt.Sprintf("%v", subscriberService.TemplateName) + "']"
}

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions" {
        return &subscriberService.SpanMonitorSessions
    }
    if childYangName == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter" {
        return &subscriberService.Ipv4PacketFilter
    }
    if childYangName == "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter" {
        return &subscriberService.Ipv6PacketFilter
    }
    if childYangName == "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network" {
        return &subscriberService.Ipv4Network
    }
    if childYangName == "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network" {
        return &subscriberService.Ipv6Network
    }
    if childYangName == "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor" {
        return &subscriberService.Ipv6Neighbor
    }
    if childYangName == "Cisco-IOS-XR-pbr-subscriber-cfg:pbr" {
        return &subscriberService.Pbr
    }
    if childYangName == "Cisco-IOS-XR-qos-ma-bng-cfg:qos" {
        return &subscriberService.Qos
    }
    if childYangName == "Cisco-IOS-XR-subscriber-accounting-cfg:accounting" {
        return &subscriberService.Accounting
    }
    return nil
}

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions"] = &subscriberService.SpanMonitorSessions
    children["Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter"] = &subscriberService.Ipv4PacketFilter
    children["Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter"] = &subscriberService.Ipv6PacketFilter
    children["Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network"] = &subscriberService.Ipv4Network
    children["Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network"] = &subscriberService.Ipv6Network
    children["Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor"] = &subscriberService.Ipv6Neighbor
    children["Cisco-IOS-XR-pbr-subscriber-cfg:pbr"] = &subscriberService.Pbr
    children["Cisco-IOS-XR-qos-ma-bng-cfg:qos"] = &subscriberService.Qos
    children["Cisco-IOS-XR-subscriber-accounting-cfg:accounting"] = &subscriberService.Accounting
    return children
}

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["template-name"] = subscriberService.TemplateName
    leafs["vrf"] = subscriberService.Vrf
    return leafs
}

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetYangName() string { return "subscriber-service" }

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) SetParent(parent types.Entity) { subscriberService.parent = parent }

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetParent() types.Entity { return subscriberService.parent }

func (subscriberService *DynamicTemplate_SubscriberServices_SubscriberService) GetParentYangName() string { return "subscriber-services" }

// DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions
// Monitor Session container for this template
type DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a particular class of Monitor Session. The type is slice
    // of
    // DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession.
    SpanMonitorSession []DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession
}

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetFilter() yfilter.YFilter { return spanMonitorSessions.YFilter }

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) SetFilter(yf yfilter.YFilter) { spanMonitorSessions.YFilter = yf }

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetGoName(yname string) string {
    if yname == "span-monitor-session" { return "SpanMonitorSession" }
    return ""
}

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetSegmentPath() string {
    return "Cisco-IOS-XR-Ethernet-SPAN-subscriber-cfg:span-monitor-sessions"
}

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "span-monitor-session" {
        for _, c := range spanMonitorSessions.SpanMonitorSession {
            if spanMonitorSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession{}
        spanMonitorSessions.SpanMonitorSession = append(spanMonitorSessions.SpanMonitorSession, child)
        return &spanMonitorSessions.SpanMonitorSession[len(spanMonitorSessions.SpanMonitorSession)-1]
    }
    return nil
}

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range spanMonitorSessions.SpanMonitorSession {
        children[spanMonitorSessions.SpanMonitorSession[i].GetSegmentPath()] = &spanMonitorSessions.SpanMonitorSession[i]
    }
    return children
}

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetBundleName() string { return "cisco_ios_xr" }

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetYangName() string { return "span-monitor-sessions" }

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) SetParent(parent types.Entity) { spanMonitorSessions.parent = parent }

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetParent() types.Entity { return spanMonitorSessions.parent }

func (spanMonitorSessions *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions) GetParentYangName() string { return "subscriber-service" }

// DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession
// Configuration for a particular class of Monitor
// Session
type DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetFilter() yfilter.YFilter { return spanMonitorSession.YFilter }

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) SetFilter(yf yfilter.YFilter) { spanMonitorSession.YFilter = yf }

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetGoName(yname string) string {
    if yname == "session-class" { return "SessionClass" }
    if yname == "mirror-first" { return "MirrorFirst" }
    if yname == "mirror-interval" { return "MirrorInterval" }
    if yname == "attachment" { return "Attachment" }
    if yname == "acl" { return "Acl" }
    return ""
}

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetSegmentPath() string {
    return "span-monitor-session" + "[session-class='" + fmt.Sprintf("%v", spanMonitorSession.SessionClass) + "']"
}

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attachment" {
        return &spanMonitorSession.Attachment
    }
    if childYangName == "acl" {
        return &spanMonitorSession.Acl
    }
    return nil
}

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["attachment"] = &spanMonitorSession.Attachment
    children["acl"] = &spanMonitorSession.Acl
    return children
}

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-class"] = spanMonitorSession.SessionClass
    leafs["mirror-first"] = spanMonitorSession.MirrorFirst
    leafs["mirror-interval"] = spanMonitorSession.MirrorInterval
    return leafs
}

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetBundleName() string { return "cisco_ios_xr" }

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetYangName() string { return "span-monitor-session" }

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) SetParent(parent types.Entity) { spanMonitorSession.parent = parent }

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetParent() types.Entity { return spanMonitorSession.parent }

func (spanMonitorSession *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession) GetParentYangName() string { return "span-monitor-sessions" }

// DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment
// Attach the interface to a Monitor Session
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session Name. The type is string with length: 1..79. This attribute is
    // mandatory.
    SessionName interface{}

    // Specify the direction of traffic to replicate (optional). The type is
    // SpanTrafficDirection.
    Direction interface{}

    // Enable port level traffic mirroring. The type is interface{}.
    PortLevelEnable interface{}
}

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetFilter() yfilter.YFilter { return attachment.YFilter }

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) SetFilter(yf yfilter.YFilter) { attachment.YFilter = yf }

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetGoName(yname string) string {
    if yname == "session-name" { return "SessionName" }
    if yname == "direction" { return "Direction" }
    if yname == "port-level-enable" { return "PortLevelEnable" }
    return ""
}

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetSegmentPath() string {
    return "attachment"
}

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-name"] = attachment.SessionName
    leafs["direction"] = attachment.Direction
    leafs["port-level-enable"] = attachment.PortLevelEnable
    return leafs
}

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetBundleName() string { return "cisco_ios_xr" }

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetYangName() string { return "attachment" }

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) SetParent(parent types.Entity) { attachment.parent = parent }

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetParent() types.Entity { return attachment.parent }

func (attachment *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Attachment) GetParentYangName() string { return "span-monitor-session" }

// DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl
// Enable ACL matching for traffic mirroring
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ACL. The type is interface{}. This attribute is mandatory.
    AclEnable interface{}

    // ACL Name. The type is string with length: 1..80.
    AclName interface{}
}

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetFilter() yfilter.YFilter { return acl.YFilter }

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) SetFilter(yf yfilter.YFilter) { acl.YFilter = yf }

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetGoName(yname string) string {
    if yname == "acl-enable" { return "AclEnable" }
    if yname == "acl-name" { return "AclName" }
    return ""
}

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetSegmentPath() string {
    return "acl"
}

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["acl-enable"] = acl.AclEnable
    leafs["acl-name"] = acl.AclName
    return leafs
}

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetBundleName() string { return "cisco_ios_xr" }

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetYangName() string { return "acl" }

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) SetParent(parent types.Entity) { acl.parent = parent }

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetParent() types.Entity { return acl.parent }

func (acl *DynamicTemplate_SubscriberServices_SubscriberService_SpanMonitorSessions_SpanMonitorSession_Acl) GetParentYangName() string { return "span-monitor-session" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter
// IPv4 Packet Filtering configuration for the
// template
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound

    // IPv4 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound
}

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetFilter() yfilter.YFilter { return ipv4PacketFilter.YFilter }

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) SetFilter(yf yfilter.YFilter) { ipv4PacketFilter.YFilter = yf }

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetGoName(yname string) string {
    if yname == "outbound" { return "Outbound" }
    if yname == "inbound" { return "Inbound" }
    return ""
}

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv4-packet-filter"
}

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "outbound" {
        return &ipv4PacketFilter.Outbound
    }
    if childYangName == "inbound" {
        return &ipv4PacketFilter.Inbound
    }
    return nil
}

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["outbound"] = &ipv4PacketFilter.Outbound
    children["inbound"] = &ipv4PacketFilter.Inbound
    return children
}

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetYangName() string { return "ipv4-packet-filter" }

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) SetParent(parent types.Entity) { ipv4PacketFilter.parent = parent }

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetParent() types.Entity { return ipv4PacketFilter.parent }

func (ipv4PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter) GetParentYangName() string { return "subscriber-service" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound
// IPv4 Packet filter to be applied to outbound
// packets
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..65. This attribute is mandatory.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetFilter() yfilter.YFilter { return outbound.YFilter }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) SetFilter(yf yfilter.YFilter) { outbound.YFilter = yf }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "hardware-count" { return "HardwareCount" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetSegmentPath() string {
    return "outbound"
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = outbound.CommonAclName
    leafs["name"] = outbound.Name
    leafs["hardware-count"] = outbound.HardwareCount
    leafs["interface-statistics"] = outbound.InterfaceStatistics
    return leafs
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetBundleName() string { return "cisco_ios_xr" }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetYangName() string { return "outbound" }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) SetParent(parent types.Entity) { outbound.parent = parent }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetParent() types.Entity { return outbound.parent }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Outbound) GetParentYangName() string { return "ipv4-packet-filter" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound
// IPv4 Packet filter to be applied to inbound
// packets
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv4 Packet Filter Name to be applied to Inbound packets NOTE: This
    // parameter is mandatory if 'CommonACLName' is not specified. The type is
    // string with length: 1..65.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    HardwareCount interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetFilter() yfilter.YFilter { return inbound.YFilter }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) SetFilter(yf yfilter.YFilter) { inbound.YFilter = yf }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "hardware-count" { return "HardwareCount" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetSegmentPath() string {
    return "inbound"
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = inbound.CommonAclName
    leafs["name"] = inbound.Name
    leafs["hardware-count"] = inbound.HardwareCount
    leafs["interface-statistics"] = inbound.InterfaceStatistics
    return leafs
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetBundleName() string { return "cisco_ios_xr" }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetYangName() string { return "inbound" }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) SetParent(parent types.Entity) { inbound.parent = parent }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetParent() types.Entity { return inbound.parent }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4PacketFilter_Inbound) GetParentYangName() string { return "ipv4-packet-filter" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter
// IPv6 Packet Filtering configuration for the
// interface
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 Packet filter to be applied to inbound packets.
    Inbound DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound

    // IPv6 Packet filter to be applied to outbound packets.
    Outbound DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound
}

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetFilter() yfilter.YFilter { return ipv6PacketFilter.YFilter }

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) SetFilter(yf yfilter.YFilter) { ipv6PacketFilter.YFilter = yf }

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetGoName(yname string) string {
    if yname == "inbound" { return "Inbound" }
    if yname == "outbound" { return "Outbound" }
    return ""
}

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-pfilter-subscriber-cfg:ipv6-packet-filter"
}

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inbound" {
        return &ipv6PacketFilter.Inbound
    }
    if childYangName == "outbound" {
        return &ipv6PacketFilter.Outbound
    }
    return nil
}

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["inbound"] = &ipv6PacketFilter.Inbound
    children["outbound"] = &ipv6PacketFilter.Outbound
    return children
}

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetYangName() string { return "ipv6-packet-filter" }

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) SetParent(parent types.Entity) { ipv6PacketFilter.parent = parent }

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetParent() types.Entity { return ipv6PacketFilter.parent }

func (ipv6PacketFilter *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter) GetParentYangName() string { return "subscriber-service" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound
// IPv6 Packet filter to be applied to inbound
// packets
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Inbound  NOTE: This parameter is
    // mandatory if 'CommonACLName' is not specified. The type is string with
    // length: 1..65.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetFilter() yfilter.YFilter { return inbound.YFilter }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) SetFilter(yf yfilter.YFilter) { inbound.YFilter = yf }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetSegmentPath() string {
    return "inbound"
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = inbound.CommonAclName
    leafs["name"] = inbound.Name
    leafs["interface-statistics"] = inbound.InterfaceStatistics
    return leafs
}

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetBundleName() string { return "cisco_ios_xr" }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetYangName() string { return "inbound" }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) SetParent(parent types.Entity) { inbound.parent = parent }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetParent() types.Entity { return inbound.parent }

func (inbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Inbound) GetParentYangName() string { return "ipv6-packet-filter" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound
// IPv6 Packet filter to be applied to outbound
// packets
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Not supported (Leave unspecified). The type is string.
    CommonAclName interface{}

    // IPv6 Packet Filter Name to be applied to Outbound packets. The type is
    // string with length: 1..65. This attribute is mandatory.
    Name interface{}

    // Not supported (Leave unspecified). The type is interface{}.
    InterfaceStatistics interface{}
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetFilter() yfilter.YFilter { return outbound.YFilter }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) SetFilter(yf yfilter.YFilter) { outbound.YFilter = yf }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetGoName(yname string) string {
    if yname == "common-acl-name" { return "CommonAclName" }
    if yname == "name" { return "Name" }
    if yname == "interface-statistics" { return "InterfaceStatistics" }
    return ""
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetSegmentPath() string {
    return "outbound"
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["common-acl-name"] = outbound.CommonAclName
    leafs["name"] = outbound.Name
    leafs["interface-statistics"] = outbound.InterfaceStatistics
    return leafs
}

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetBundleName() string { return "cisco_ios_xr" }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetYangName() string { return "outbound" }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) SetParent(parent types.Entity) { outbound.parent = parent }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetParent() types.Entity { return outbound.parent }

func (outbound *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6PacketFilter_Outbound) GetParentYangName() string { return "ipv6-packet-filter" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network
// Interface IPv4 Network configuration data
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network struct {
    parent types.Entity
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

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetFilter() yfilter.YFilter { return ipv4Network.YFilter }

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) SetFilter(yf yfilter.YFilter) { ipv4Network.YFilter = yf }

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetGoName(yname string) string {
    if yname == "unnumbered" { return "Unnumbered" }
    if yname == "mtu" { return "Mtu" }
    if yname == "unreachables" { return "Unreachables" }
    if yname == "rpf" { return "Rpf" }
    return ""
}

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ma-subscriber-cfg:ipv4-network"
}

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unnumbered"] = ipv4Network.Unnumbered
    leafs["mtu"] = ipv4Network.Mtu
    leafs["unreachables"] = ipv4Network.Unreachables
    leafs["rpf"] = ipv4Network.Rpf
    return leafs
}

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetYangName() string { return "ipv4-network" }

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) SetParent(parent types.Entity) { ipv4Network.parent = parent }

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetParent() types.Entity { return ipv4Network.parent }

func (ipv4Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv4Network) GetParentYangName() string { return "subscriber-service" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network
// Interface IPv6 Network configuration data
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network struct {
    parent types.Entity
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

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetFilter() yfilter.YFilter { return ipv6Network.YFilter }

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) SetFilter(yf yfilter.YFilter) { ipv6Network.YFilter = yf }

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetGoName(yname string) string {
    if yname == "mtu" { return "Mtu" }
    if yname == "rpf" { return "Rpf" }
    if yname == "unreachables" { return "Unreachables" }
    if yname == "addresses" { return "Addresses" }
    return ""
}

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ma-subscriber-cfg:ipv6-network"
}

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "addresses" {
        return &ipv6Network.Addresses
    }
    return nil
}

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["addresses"] = &ipv6Network.Addresses
    return children
}

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mtu"] = ipv6Network.Mtu
    leafs["rpf"] = ipv6Network.Rpf
    leafs["unreachables"] = ipv6Network.Unreachables
    return leafs
}

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetYangName() string { return "ipv6-network" }

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) SetParent(parent types.Entity) { ipv6Network.parent = parent }

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetParent() types.Entity { return ipv6Network.parent }

func (ipv6Network *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network) GetParentYangName() string { return "subscriber-service" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses
// Set the IPv6 address of an interface
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Auto IPv6 Interface Configuration.
    AutoConfiguration DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration
}

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetFilter() yfilter.YFilter { return addresses.YFilter }

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) SetFilter(yf yfilter.YFilter) { addresses.YFilter = yf }

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetGoName(yname string) string {
    if yname == "auto-configuration" { return "AutoConfiguration" }
    return ""
}

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetSegmentPath() string {
    return "addresses"
}

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto-configuration" {
        return &addresses.AutoConfiguration
    }
    return nil
}

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["auto-configuration"] = &addresses.AutoConfiguration
    return children
}

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetYangName() string { return "addresses" }

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) SetParent(parent types.Entity) { addresses.parent = parent }

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetParent() types.Entity { return addresses.parent }

func (addresses *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses) GetParentYangName() string { return "ipv6-network" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration
// Auto IPv6 Interface Configuration
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The flag to enable auto ipv6 interface configuration. The type is
    // interface{}.
    Enable interface{}
}

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetFilter() yfilter.YFilter { return autoConfiguration.YFilter }

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) SetFilter(yf yfilter.YFilter) { autoConfiguration.YFilter = yf }

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetSegmentPath() string {
    return "auto-configuration"
}

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = autoConfiguration.Enable
    return leafs
}

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetYangName() string { return "auto-configuration" }

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) SetParent(parent types.Entity) { autoConfiguration.parent = parent }

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetParent() types.Entity { return autoConfiguration.parent }

func (autoConfiguration *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Network_Addresses_AutoConfiguration) GetParentYangName() string { return "addresses" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor
// Interface IPv6 Network configuration data
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor struct {
    parent types.Entity
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

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetFilter() yfilter.YFilter { return ipv6Neighbor.YFilter }

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) SetFilter(yf yfilter.YFilter) { ipv6Neighbor.YFilter = yf }

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetGoName(yname string) string {
    if yname == "framed-prefix-pool" { return "FramedPrefixPool" }
    if yname == "managed-config" { return "ManagedConfig" }
    if yname == "other-config" { return "OtherConfig" }
    if yname == "start-ra-on-ipv6-enable" { return "StartRaOnIpv6Enable" }
    if yname == "nud-enable" { return "NudEnable" }
    if yname == "ra-lifetime" { return "RaLifetime" }
    if yname == "router-preference" { return "RouterPreference" }
    if yname == "ra-suppress" { return "RaSuppress" }
    if yname == "ra-unicast" { return "RaUnicast" }
    if yname == "ra-unspecify-hoplimit" { return "RaUnspecifyHoplimit" }
    if yname == "ra-suppress-mtu" { return "RaSuppressMtu" }
    if yname == "suppress-cache-learning" { return "SuppressCacheLearning" }
    if yname == "reachable-time" { return "ReachableTime" }
    if yname == "ns-interval" { return "NsInterval" }
    if yname == "ra-interval" { return "RaInterval" }
    if yname == "framed-prefix" { return "FramedPrefix" }
    if yname == "duplicate-address-detection" { return "DuplicateAddressDetection" }
    if yname == "ra-initial" { return "RaInitial" }
    return ""
}

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-nd-subscriber-cfg:ipv6-neighbor"
}

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ra-interval" {
        return &ipv6Neighbor.RaInterval
    }
    if childYangName == "framed-prefix" {
        return &ipv6Neighbor.FramedPrefix
    }
    if childYangName == "duplicate-address-detection" {
        return &ipv6Neighbor.DuplicateAddressDetection
    }
    if childYangName == "ra-initial" {
        return &ipv6Neighbor.RaInitial
    }
    return nil
}

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ra-interval"] = &ipv6Neighbor.RaInterval
    children["framed-prefix"] = &ipv6Neighbor.FramedPrefix
    children["duplicate-address-detection"] = &ipv6Neighbor.DuplicateAddressDetection
    children["ra-initial"] = &ipv6Neighbor.RaInitial
    return children
}

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["framed-prefix-pool"] = ipv6Neighbor.FramedPrefixPool
    leafs["managed-config"] = ipv6Neighbor.ManagedConfig
    leafs["other-config"] = ipv6Neighbor.OtherConfig
    leafs["start-ra-on-ipv6-enable"] = ipv6Neighbor.StartRaOnIpv6Enable
    leafs["nud-enable"] = ipv6Neighbor.NudEnable
    leafs["ra-lifetime"] = ipv6Neighbor.RaLifetime
    leafs["router-preference"] = ipv6Neighbor.RouterPreference
    leafs["ra-suppress"] = ipv6Neighbor.RaSuppress
    leafs["ra-unicast"] = ipv6Neighbor.RaUnicast
    leafs["ra-unspecify-hoplimit"] = ipv6Neighbor.RaUnspecifyHoplimit
    leafs["ra-suppress-mtu"] = ipv6Neighbor.RaSuppressMtu
    leafs["suppress-cache-learning"] = ipv6Neighbor.SuppressCacheLearning
    leafs["reachable-time"] = ipv6Neighbor.ReachableTime
    leafs["ns-interval"] = ipv6Neighbor.NsInterval
    return leafs
}

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetYangName() string { return "ipv6-neighbor" }

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) SetParent(parent types.Entity) { ipv6Neighbor.parent = parent }

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetParent() types.Entity { return ipv6Neighbor.parent }

func (ipv6Neighbor *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor) GetParentYangName() string { return "subscriber-service" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval
// Set IPv6 Router Advertisement (RA) interval in
// seconds
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Maximum interface{}

    // Minimum RA interval in seconds. Must be less than 0.75 * maximum interval.
    // The type is interface{} with range: 3..1800. Units are second.
    Minimum interface{}
}

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetFilter() yfilter.YFilter { return raInterval.YFilter }

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) SetFilter(yf yfilter.YFilter) { raInterval.YFilter = yf }

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "minimum" { return "Minimum" }
    return ""
}

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetSegmentPath() string {
    return "ra-interval"
}

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = raInterval.Maximum
    leafs["minimum"] = raInterval.Minimum
    return leafs
}

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetBundleName() string { return "cisco_ios_xr" }

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetYangName() string { return "ra-interval" }

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) SetParent(parent types.Entity) { raInterval.parent = parent }

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetParent() types.Entity { return raInterval.parent }

func (raInterval *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInterval) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix
// Set the IPv6 framed ipv6 prefix for a
// subscriber interface 
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 framed prefix length. The type is interface{} with range: 0..128. This
    // attribute is mandatory.
    PrefixLength interface{}

    // IPV6 framed prefix address. The type is string. This attribute is
    // mandatory.
    Prefix interface{}
}

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetFilter() yfilter.YFilter { return framedPrefix.YFilter }

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) SetFilter(yf yfilter.YFilter) { framedPrefix.YFilter = yf }

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetSegmentPath() string {
    return "framed-prefix"
}

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = framedPrefix.PrefixLength
    leafs["prefix"] = framedPrefix.Prefix
    return leafs
}

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetYangName() string { return "framed-prefix" }

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) SetParent(parent types.Entity) { framedPrefix.parent = parent }

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetParent() types.Entity { return framedPrefix.parent }

func (framedPrefix *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_FramedPrefix) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection
// Duplicate Address Detection (DAD)
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set IPv6 duplicate address detection transmits. The type is interface{}
    // with range: 0..600.
    Attempts interface{}
}

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetFilter() yfilter.YFilter { return duplicateAddressDetection.YFilter }

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) SetFilter(yf yfilter.YFilter) { duplicateAddressDetection.YFilter = yf }

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetGoName(yname string) string {
    if yname == "attempts" { return "Attempts" }
    return ""
}

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetSegmentPath() string {
    return "duplicate-address-detection"
}

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["attempts"] = duplicateAddressDetection.Attempts
    return leafs
}

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetBundleName() string { return "cisco_ios_xr" }

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetYangName() string { return "duplicate-address-detection" }

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) SetParent(parent types.Entity) { duplicateAddressDetection.parent = parent }

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetParent() types.Entity { return duplicateAddressDetection.parent }

func (duplicateAddressDetection *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_DuplicateAddressDetection) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial
// IPv6 ND RA Initial
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Initial RA count. The type is interface{} with range: 0..32. This attribute
    // is mandatory.
    Count interface{}

    // Initial RA interval in seconds. The type is interface{} with range:
    // 4..1800. This attribute is mandatory. Units are second.
    Interval interface{}
}

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetFilter() yfilter.YFilter { return raInitial.YFilter }

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) SetFilter(yf yfilter.YFilter) { raInitial.YFilter = yf }

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetGoName(yname string) string {
    if yname == "count" { return "Count" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetSegmentPath() string {
    return "ra-initial"
}

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["count"] = raInitial.Count
    leafs["interval"] = raInitial.Interval
    return leafs
}

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetBundleName() string { return "cisco_ios_xr" }

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetYangName() string { return "ra-initial" }

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) SetParent(parent types.Entity) { raInitial.parent = parent }

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetParent() types.Entity { return raInitial.parent }

func (raInitial *DynamicTemplate_SubscriberServices_SubscriberService_Ipv6Neighbor_RaInitial) GetParentYangName() string { return "ipv6-neighbor" }

// DynamicTemplate_SubscriberServices_SubscriberService_Pbr
// Dynamic Template PBR configuration
type DynamicTemplate_SubscriberServices_SubscriberService_Pbr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Class for subscriber ingress policy. The type is string.
    ServicePolicyIn interface{}

    // PBR service policy configuration.
    ServicePolicy DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy
}

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetFilter() yfilter.YFilter { return pbr.YFilter }

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) SetFilter(yf yfilter.YFilter) { pbr.YFilter = yf }

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetGoName(yname string) string {
    if yname == "service-policy-in" { return "ServicePolicyIn" }
    if yname == "service-policy" { return "ServicePolicy" }
    return ""
}

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetSegmentPath() string {
    return "Cisco-IOS-XR-pbr-subscriber-cfg:pbr"
}

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-policy" {
        return &pbr.ServicePolicy
    }
    return nil
}

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-policy"] = &pbr.ServicePolicy
    return children
}

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["service-policy-in"] = pbr.ServicePolicyIn
    return leafs
}

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetBundleName() string { return "cisco_ios_xr" }

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetYangName() string { return "pbr" }

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) SetParent(parent types.Entity) { pbr.parent = parent }

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetParent() types.Entity { return pbr.parent }

func (pbr *DynamicTemplate_SubscriberServices_SubscriberService_Pbr) GetParentYangName() string { return "subscriber-service" }

// DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy
// PBR service policy configuration
type DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ingress service policy. The type is string.
    Input interface{}
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetFilter() yfilter.YFilter { return servicePolicy.YFilter }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) SetFilter(yf yfilter.YFilter) { servicePolicy.YFilter = yf }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetSegmentPath() string {
    return "service-policy"
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["input"] = servicePolicy.Input
    return leafs
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetYangName() string { return "service-policy" }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) SetParent(parent types.Entity) { servicePolicy.parent = parent }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetParent() types.Entity { return servicePolicy.parent }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Pbr_ServicePolicy) GetParentYangName() string { return "pbr" }

// DynamicTemplate_SubscriberServices_SubscriberService_Qos
// QoS dynamically applied configuration template
type DynamicTemplate_SubscriberServices_SubscriberService_Qos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service policy to be applied in ingress/egress direction.
    ServicePolicy DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy

    // QoS L2 overhead accounting.
    Account DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account

    // QoS to be applied in egress direction.
    Output DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output
}

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetFilter() yfilter.YFilter { return qos.YFilter }

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) SetFilter(yf yfilter.YFilter) { qos.YFilter = yf }

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetGoName(yname string) string {
    if yname == "service-policy" { return "ServicePolicy" }
    if yname == "account" { return "Account" }
    if yname == "output" { return "Output" }
    return ""
}

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetSegmentPath() string {
    return "Cisco-IOS-XR-qos-ma-bng-cfg:qos"
}

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-policy" {
        return &qos.ServicePolicy
    }
    if childYangName == "account" {
        return &qos.Account
    }
    if childYangName == "output" {
        return &qos.Output
    }
    return nil
}

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-policy"] = &qos.ServicePolicy
    children["account"] = &qos.Account
    children["output"] = &qos.Output
    return children
}

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetBundleName() string { return "cisco_ios_xr" }

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetYangName() string { return "qos" }

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) SetParent(parent types.Entity) { qos.parent = parent }

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetParent() types.Entity { return qos.parent }

func (qos *DynamicTemplate_SubscriberServices_SubscriberService_Qos) GetParentYangName() string { return "subscriber-service" }

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy
// Service policy to be applied in ingress/egress
// direction
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber ingress policy.
    Input DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input

    // Subscriber egress policy.
    Output DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetFilter() yfilter.YFilter { return servicePolicy.YFilter }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) SetFilter(yf yfilter.YFilter) { servicePolicy.YFilter = yf }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetSegmentPath() string {
    return "service-policy"
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &servicePolicy.Input
    }
    if childYangName == "output" {
        return &servicePolicy.Output
    }
    return nil
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &servicePolicy.Input
    children["output"] = &servicePolicy.Output
    return children
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetYangName() string { return "service-policy" }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) SetParent(parent types.Entity) { servicePolicy.parent = parent }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetParent() types.Entity { return servicePolicy.parent }

func (servicePolicy *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy) GetParentYangName() string { return "qos" }

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input
// Subscriber ingress policy
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "spi-name" { return "SpiName" }
    if yname == "merge" { return "Merge" }
    if yname == "merge-id" { return "MergeId" }
    if yname == "account-stats" { return "AccountStats" }
    return ""
}

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetSegmentPath() string {
    return "input"
}

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = input.PolicyName
    leafs["spi-name"] = input.SpiName
    leafs["merge"] = input.Merge
    leafs["merge-id"] = input.MergeId
    leafs["account-stats"] = input.AccountStats
    return leafs
}

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetYangName() string { return "input" }

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetParent() types.Entity { return input.parent }

func (input *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Input) GetParentYangName() string { return "service-policy" }

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output
// Subscriber egress policy
// This type is a presence type.
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "spi-name" { return "SpiName" }
    if yname == "merge" { return "Merge" }
    if yname == "merge-id" { return "MergeId" }
    if yname == "account-stats" { return "AccountStats" }
    return ""
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetSegmentPath() string {
    return "output"
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = output.PolicyName
    leafs["spi-name"] = output.SpiName
    leafs["merge"] = output.Merge
    leafs["merge-id"] = output.MergeId
    leafs["account-stats"] = output.AccountStats
    return leafs
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetYangName() string { return "output" }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetParent() types.Entity { return output.parent }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_ServicePolicy_Output) GetParentYangName() string { return "service-policy" }

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account
// QoS L2 overhead accounting
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account struct {
    parent types.Entity
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

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetFilter() yfilter.YFilter { return account.YFilter }

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) SetFilter(yf yfilter.YFilter) { account.YFilter = yf }

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetGoName(yname string) string {
    if yname == "aal" { return "Aal" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "atm-cell-tax" { return "AtmCellTax" }
    if yname == "user-defined" { return "UserDefined" }
    return ""
}

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetSegmentPath() string {
    return "account"
}

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aal"] = account.Aal
    leafs["encapsulation"] = account.Encapsulation
    leafs["atm-cell-tax"] = account.AtmCellTax
    leafs["user-defined"] = account.UserDefined
    return leafs
}

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetBundleName() string { return "cisco_ios_xr" }

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetYangName() string { return "account" }

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) SetParent(parent types.Entity) { account.parent = parent }

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetParent() types.Entity { return account.parent }

func (account *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Account) GetParentYangName() string { return "qos" }

// DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output
// QoS to be applied in egress direction
type DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum bandwidth value for the subscriber (in kbps). The type is
    // interface{} with range: 1..4294967295. Units are kbit/s.
    MinimumBandwidth interface{}
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetGoName(yname string) string {
    if yname == "minimum-bandwidth" { return "MinimumBandwidth" }
    return ""
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetSegmentPath() string {
    return "output"
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minimum-bandwidth"] = output.MinimumBandwidth
    return leafs
}

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetYangName() string { return "output" }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetParent() types.Entity { return output.parent }

func (output *DynamicTemplate_SubscriberServices_SubscriberService_Qos_Output) GetParentYangName() string { return "qos" }

// DynamicTemplate_SubscriberServices_SubscriberService_Accounting
// Subscriber accounting dynamic-template commands
type DynamicTemplate_SubscriberServices_SubscriberService_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber accounting prepaid feature. The type is string.
    PrepaidFeature interface{}

    // Subscriber accounting service accounting.
    ServiceAccounting DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting

    // Subscriber accounting session accounting.
    Session DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session

    // Subscriber accounting idle timeout.
    IdleTimeout DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout
}

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetGoName(yname string) string {
    if yname == "prepaid-feature" { return "PrepaidFeature" }
    if yname == "service-accounting" { return "ServiceAccounting" }
    if yname == "session" { return "Session" }
    if yname == "idle-timeout" { return "IdleTimeout" }
    return ""
}

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-accounting-cfg:accounting"
}

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service-accounting" {
        return &accounting.ServiceAccounting
    }
    if childYangName == "session" {
        return &accounting.Session
    }
    if childYangName == "idle-timeout" {
        return &accounting.IdleTimeout
    }
    return nil
}

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["service-accounting"] = &accounting.ServiceAccounting
    children["session"] = &accounting.Session
    children["idle-timeout"] = &accounting.IdleTimeout
    return children
}

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prepaid-feature"] = accounting.PrepaidFeature
    return leafs
}

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetYangName() string { return "accounting" }

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting) GetParentYangName() string { return "subscriber-service" }

// DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting
// Subscriber accounting service accounting
type DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service accounting method list name. The type is string.
    MethodListName interface{}

    // Accounting interim interval in minutes. The type is interface{} with range:
    // -2147483648..2147483647. Units are minute.
    AccountingInterimInterval interface{}
}

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetFilter() yfilter.YFilter { return serviceAccounting.YFilter }

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) SetFilter(yf yfilter.YFilter) { serviceAccounting.YFilter = yf }

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetGoName(yname string) string {
    if yname == "method-list-name" { return "MethodListName" }
    if yname == "accounting-interim-interval" { return "AccountingInterimInterval" }
    return ""
}

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetSegmentPath() string {
    return "service-accounting"
}

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method-list-name"] = serviceAccounting.MethodListName
    leafs["accounting-interim-interval"] = serviceAccounting.AccountingInterimInterval
    return leafs
}

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetBundleName() string { return "cisco_ios_xr" }

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetYangName() string { return "service-accounting" }

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) SetParent(parent types.Entity) { serviceAccounting.parent = parent }

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetParent() types.Entity { return serviceAccounting.parent }

func (serviceAccounting *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_ServiceAccounting) GetParentYangName() string { return "accounting" }

// DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session
// Subscriber accounting session accounting
type DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session accounting method list name. The type is string.
    MethodListName interface{}

    // Interim accounting interval in minutes. The type is interface{} with range:
    // -2147483648..2147483647. Units are minute.
    PeriodicInterval interface{}

    // Dual stack wait delay in seconds. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    DualStackDelay interface{}

    // Hold Accounting start based on IA_PD. The type is interface{} with range:
    // -2147483648..2147483647.
    HoldAcctStart interface{}
}

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetGoName(yname string) string {
    if yname == "method-list-name" { return "MethodListName" }
    if yname == "periodic-interval" { return "PeriodicInterval" }
    if yname == "dual-stack-delay" { return "DualStackDelay" }
    if yname == "hold-acct-start" { return "HoldAcctStart" }
    return ""
}

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetSegmentPath() string {
    return "session"
}

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["method-list-name"] = session.MethodListName
    leafs["periodic-interval"] = session.PeriodicInterval
    leafs["dual-stack-delay"] = session.DualStackDelay
    leafs["hold-acct-start"] = session.HoldAcctStart
    return leafs
}

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetYangName() string { return "session" }

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetParent() types.Entity { return session.parent }

func (session *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_Session) GetParentYangName() string { return "accounting" }

// DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout
// Subscriber accounting idle timeout
type DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout struct {
    parent types.Entity
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

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetFilter() yfilter.YFilter { return idleTimeout.YFilter }

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) SetFilter(yf yfilter.YFilter) { idleTimeout.YFilter = yf }

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetGoName(yname string) string {
    if yname == "timeout-value" { return "TimeoutValue" }
    if yname == "threshold" { return "Threshold" }
    if yname == "direction" { return "Direction" }
    return ""
}

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetSegmentPath() string {
    return "idle-timeout"
}

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout-value"] = idleTimeout.TimeoutValue
    leafs["threshold"] = idleTimeout.Threshold
    leafs["direction"] = idleTimeout.Direction
    return leafs
}

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetBundleName() string { return "cisco_ios_xr" }

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetYangName() string { return "idle-timeout" }

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) SetParent(parent types.Entity) { idleTimeout.parent = parent }

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetParent() types.Entity { return idleTimeout.parent }

func (idleTimeout *DynamicTemplate_SubscriberServices_SubscriberService_Accounting_IdleTimeout) GetParentYangName() string { return "accounting" }

