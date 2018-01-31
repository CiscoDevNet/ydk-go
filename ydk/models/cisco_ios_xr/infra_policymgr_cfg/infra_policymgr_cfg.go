// This module contains a collection of YANG definitions
// for Cisco IOS-XR ASR9k policy manager configuration.
//  
// Copyright (c) 2013, 2015-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_policymgr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_policymgr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-policymgr-cfg policy-manager}", reflect.TypeOf(PolicyManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-policymgr-cfg:policy-manager", reflect.TypeOf(PolicyManager{}))
}

// EventType represents Event type.
type EventType string

const (
    // Account logoff event.
    EventType_account_logoff EventType = "account-logoff"

    // Account logon event.
    EventType_account_logon EventType = "account-logon"

    // Authentication failure event.
    EventType_authentication_failure EventType = "authentication-failure"

    // Authentication no response event.
    EventType_authentication_no_response EventType = "authentication-no-response"

    // Authorization failure event.
    EventType_authorization_failure EventType = "authorization-failure"

    // Authorization no response event.
    EventType_authorization_no_response EventType = "authorization-no-response"

    // Credit exhaustion event.
    EventType_credit_exhausted EventType = "credit-exhausted"

    // Exception event.
    EventType_exception EventType = "exception"

    // Idle timeout event.
    EventType_idle_timeout EventType = "idle-timeout"

    // Quota depletion event.
    EventType_quota_depleted EventType = "quota-depleted"

    // Service start event.
    EventType_service_start EventType = "service-start"

    // Service stop event.
    EventType_service_stop EventType = "service-stop"

    // Session activate event.
    EventType_session_activate EventType = "session-activate"

    // Session start event.
    EventType_session_start EventType = "session-start"

    // Session stop event.
    EventType_session_stop EventType = "session-stop"

    // Timer expiry event.
    EventType_timer_expiry EventType = "timer-expiry"
)

// ClassMapType represents Policy manager class-map type.
type ClassMapType string

const (
    // QoS Classmap.
    ClassMapType_qos ClassMapType = "qos"

    // TRAFFIC Classmap.
    ClassMapType_traffic ClassMapType = "traffic"

    // Control Subscriber Classmap.
    ClassMapType_control ClassMapType = "control"
)

// ExecutionStrategy represents Executuion strategy.
type ExecutionStrategy string

const (
    // Do all actions.
    ExecutionStrategy_do_all ExecutionStrategy = "do-all"

    // Do all actions until failure.
    ExecutionStrategy_do_until_failure ExecutionStrategy = "do-until-failure"

    // Do all actions until success.
    ExecutionStrategy_do_until_success ExecutionStrategy = "do-until-success"
)

// AuthorizeIdentifier represents Authorize identifier.
type AuthorizeIdentifier string

const (
    // Authorize circuit ID.
    AuthorizeIdentifier_circuit_id AuthorizeIdentifier = "circuit-id"

    // Authorize dhcp client ID.
    AuthorizeIdentifier_dhcp_client_id AuthorizeIdentifier = "dhcp-client-id"

    // Authorize remote ID.
    AuthorizeIdentifier_remote_id AuthorizeIdentifier = "remote-id"

    // Authorize source IPv4 address.
    AuthorizeIdentifier_source_address_ipv4 AuthorizeIdentifier = "source-address-ipv4"

    // Authorize source IPv6 address.
    AuthorizeIdentifier_source_address_ipv6 AuthorizeIdentifier = "source-address-ipv6"

    // Authorize source MAC address.
    AuthorizeIdentifier_source_address_mac AuthorizeIdentifier = "source-address-mac"

    // Authorize username.
    AuthorizeIdentifier_username AuthorizeIdentifier = "username"
)

// PmapClassMapType represents Policy manager class-map type.
type PmapClassMapType string

const (
    // QoS Classmap.
    PmapClassMapType_qos PmapClassMapType = "qos"

    // TRAFFIC Classmap.
    PmapClassMapType_traffic PmapClassMapType = "traffic"

    // Subscriber Control Classmap.
    PmapClassMapType_subscriber_control PmapClassMapType = "subscriber-control"
)

// PolicyMapType represents Policy manager policy-map type.
type PolicyMapType string

const (
    // QoS Policymap
    PolicyMapType_qos PolicyMapType = "qos"

    // PBR Policymap
    PolicyMapType_pbr PolicyMapType = "pbr"

    // TRAFFIC Policymap
    PolicyMapType_traffic PolicyMapType = "traffic"

    // SUBSCRIBER-CONTROL Policymap
    PolicyMapType_subscriber_control PolicyMapType = "subscriber-control"

    // REDIRECT Policy map
    PolicyMapType_redirect PolicyMapType = "redirect"

    // FLOWMONITOR Policy map
    PolicyMapType_flow_monitor PolicyMapType = "flow-monitor"
)

// PolicyManager
// Global Policy Manager configuration.
type PolicyManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Class-maps configuration.
    ClassMaps PolicyManager_ClassMaps

    // Policy-maps configuration.
    PolicyMaps PolicyManager_PolicyMaps
}

func (policyManager *PolicyManager) GetFilter() yfilter.YFilter { return policyManager.YFilter }

func (policyManager *PolicyManager) SetFilter(yf yfilter.YFilter) { policyManager.YFilter = yf }

func (policyManager *PolicyManager) GetGoName(yname string) string {
    if yname == "class-maps" { return "ClassMaps" }
    if yname == "policy-maps" { return "PolicyMaps" }
    return ""
}

func (policyManager *PolicyManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager"
}

func (policyManager *PolicyManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class-maps" {
        return &policyManager.ClassMaps
    }
    if childYangName == "policy-maps" {
        return &policyManager.PolicyMaps
    }
    return nil
}

func (policyManager *PolicyManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["class-maps"] = &policyManager.ClassMaps
    children["policy-maps"] = &policyManager.PolicyMaps
    return children
}

func (policyManager *PolicyManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyManager *PolicyManager) GetBundleName() string { return "cisco_ios_xr" }

func (policyManager *PolicyManager) GetYangName() string { return "policy-manager" }

func (policyManager *PolicyManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyManager *PolicyManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyManager *PolicyManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyManager *PolicyManager) SetParent(parent types.Entity) { policyManager.parent = parent }

func (policyManager *PolicyManager) GetParent() types.Entity { return policyManager.parent }

func (policyManager *PolicyManager) GetParentYangName() string { return "Cisco-IOS-XR-infra-policymgr-cfg" }

// PolicyManager_ClassMaps
// Class-maps configuration.
type PolicyManager_ClassMaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Class-map configuration. The type is slice of
    // PolicyManager_ClassMaps_ClassMap.
    ClassMap []PolicyManager_ClassMaps_ClassMap
}

func (classMaps *PolicyManager_ClassMaps) GetFilter() yfilter.YFilter { return classMaps.YFilter }

func (classMaps *PolicyManager_ClassMaps) SetFilter(yf yfilter.YFilter) { classMaps.YFilter = yf }

func (classMaps *PolicyManager_ClassMaps) GetGoName(yname string) string {
    if yname == "class-map" { return "ClassMap" }
    return ""
}

func (classMaps *PolicyManager_ClassMaps) GetSegmentPath() string {
    return "class-maps"
}

func (classMaps *PolicyManager_ClassMaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class-map" {
        for _, c := range classMaps.ClassMap {
            if classMaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap{}
        classMaps.ClassMap = append(classMaps.ClassMap, child)
        return &classMaps.ClassMap[len(classMaps.ClassMap)-1]
    }
    return nil
}

func (classMaps *PolicyManager_ClassMaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classMaps.ClassMap {
        children[classMaps.ClassMap[i].GetSegmentPath()] = &classMaps.ClassMap[i]
    }
    return children
}

func (classMaps *PolicyManager_ClassMaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classMaps *PolicyManager_ClassMaps) GetBundleName() string { return "cisco_ios_xr" }

func (classMaps *PolicyManager_ClassMaps) GetYangName() string { return "class-maps" }

func (classMaps *PolicyManager_ClassMaps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classMaps *PolicyManager_ClassMaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classMaps *PolicyManager_ClassMaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classMaps *PolicyManager_ClassMaps) SetParent(parent types.Entity) { classMaps.parent = parent }

func (classMaps *PolicyManager_ClassMaps) GetParent() types.Entity { return classMaps.parent }

func (classMaps *PolicyManager_ClassMaps) GetParentYangName() string { return "policy-manager" }

// PolicyManager_ClassMaps_ClassMap
// Class-map configuration.
type PolicyManager_ClassMaps_ClassMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Type of class-map. The type is ClassMapType.
    Type interface{}

    // This attribute is a key. Name of class-map. The type is string with
    // pattern: [a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}.
    Name interface{}

    // Match all match criteria. The type is interface{}.
    ClassMapModeMatchAny interface{}

    // Match any match criteria. The type is interface{}.
    ClassMapModeMatchAll interface{}

    // Description for this policy-map. The type is string.
    Description interface{}

    // Match rules.
    Match PolicyManager_ClassMaps_ClassMap_Match

    // Match not rules.
    MatchNot PolicyManager_ClassMaps_ClassMap_MatchNot
}

func (classMap *PolicyManager_ClassMaps_ClassMap) GetFilter() yfilter.YFilter { return classMap.YFilter }

func (classMap *PolicyManager_ClassMaps_ClassMap) SetFilter(yf yfilter.YFilter) { classMap.YFilter = yf }

func (classMap *PolicyManager_ClassMaps_ClassMap) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "name" { return "Name" }
    if yname == "class-map-mode-match-any" { return "ClassMapModeMatchAny" }
    if yname == "class-map-mode-match-all" { return "ClassMapModeMatchAll" }
    if yname == "description" { return "Description" }
    if yname == "match" { return "Match" }
    if yname == "match-not" { return "MatchNot" }
    return ""
}

func (classMap *PolicyManager_ClassMaps_ClassMap) GetSegmentPath() string {
    return "class-map" + "[type='" + fmt.Sprintf("%v", classMap.Type) + "']" + "[name='" + fmt.Sprintf("%v", classMap.Name) + "']"
}

func (classMap *PolicyManager_ClassMaps_ClassMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "match" {
        return &classMap.Match
    }
    if childYangName == "match-not" {
        return &classMap.MatchNot
    }
    return nil
}

func (classMap *PolicyManager_ClassMaps_ClassMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["match"] = &classMap.Match
    children["match-not"] = &classMap.MatchNot
    return children
}

func (classMap *PolicyManager_ClassMaps_ClassMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = classMap.Type
    leafs["name"] = classMap.Name
    leafs["class-map-mode-match-any"] = classMap.ClassMapModeMatchAny
    leafs["class-map-mode-match-all"] = classMap.ClassMapModeMatchAll
    leafs["description"] = classMap.Description
    return leafs
}

func (classMap *PolicyManager_ClassMaps_ClassMap) GetBundleName() string { return "cisco_ios_xr" }

func (classMap *PolicyManager_ClassMaps_ClassMap) GetYangName() string { return "class-map" }

func (classMap *PolicyManager_ClassMaps_ClassMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classMap *PolicyManager_ClassMaps_ClassMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classMap *PolicyManager_ClassMaps_ClassMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classMap *PolicyManager_ClassMaps_ClassMap) SetParent(parent types.Entity) { classMap.parent = parent }

func (classMap *PolicyManager_ClassMaps_ClassMap) GetParent() types.Entity { return classMap.parent }

func (classMap *PolicyManager_ClassMaps_ClassMap) GetParentYangName() string { return "class-maps" }

// PolicyManager_ClassMaps_ClassMap_Match
// Match rules.
type PolicyManager_ClassMaps_ClassMap_Match struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Match IPv4 DSCP. The type is slice of string with pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(([0-9]|[1-5][0-9]|6[0-3])-([0-9]|[1-5][0-9]|6[0-3]))|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Ipv4Dscp []interface{}

    // Match IPv6 DSCP. The type is slice of string with pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(([0-9]|[1-5][0-9]|6[0-3])-([0-9]|[1-5][0-9]|6[0-3]))|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Ipv6Dscp []interface{}

    // Match DSCP. The type is slice of string with pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(([0-9]|[1-5][0-9]|6[0-3])-([0-9]|[1-5][0-9]|6[0-3]))|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Dscp []interface{}

    // Match IPv4 precedence. The type is one of the following types: slice of int
    // with range: 0..7, or slice of string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Ipv4Precedence []interface{}

    // Match IPv6 precedence. The type is one of the following types: slice of int
    // with range: 0..7, or slice of string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Ipv6Precedence []interface{}

    // Match precedence. The type is one of the following types: slice of int with
    // range: 0..7, or slice of string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Precedence []interface{}

    // Match QoS group. Should be value 0..512 or range. The type is slice of
    // string with pattern: (\d+)|(\d+\-\d+).
    QosGroup []interface{}

    // Match Traffic Class. Should be value 0..63 or range. The type is slice of
    // string with pattern: (\d+)|(\d+\-\d+).
    TrafficClass []interface{}

    // Match CoS. The type is slice of interface{} with range: 0..7.
    Cos []interface{}

    // Match inner CoS. The type is slice of interface{} with range: 0..7.
    InnerCos []interface{}

    // Match DEI bit. The type is interface{} with range: 0..1.
    Dei interface{}

    // Match DEI INNER  bit. The type is interface{} with range: 0..1.
    DeiInner interface{}

    // Match protocol. The type is slice of string with pattern:
    // ([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\-([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))|((ahp)|(dhcpv4)|(dhcpv6)|(eigrp)|(esp)|(gre)|(icmp)|(igmp)|(igrp)|(ipinip)|(ipv4)|(ipv6)|(ipv6icmp)|(mpls)|(nos)|(ospf)|(pcp)|(pim)|(ppp)|(sctp)|(tcp)|(udp)).
    Protocol []interface{}

    // Match IPv4 ACL. The type is string with length: 1..64.
    Ipv4Acl interface{}

    // Match IPv6 ACL. The type is string with length: 1..64.
    Ipv6Acl interface{}

    // Match Ethernet Services. The type is string with length: 1..64.
    EthernetServicesAcl interface{}

    // Match MPLS experimental topmost label. The type is slice of interface{}
    // with range: 0..7.
    MplsExperimentalTopmost []interface{}

    // Match MPLS experimental imposition label. The type is slice of interface{}
    // with range: 0..7.
    MplsExperimentalImposition []interface{}

    // Match discard class. The type is slice of interface{} with range: 0..7.
    DiscardClass []interface{}

    // Match IPv4 packet length. Should be value 0..65535 or range. The type is
    // slice of string with pattern: (\d+)|(\d+\-\d+).
    Ipv4PacketLength []interface{}

    // Match IPv6 packet length.  Should be value 0..65535 or range. The type is
    // slice of string with pattern: (\d+)|(\d+\-\d+).
    Ipv6PacketLength []interface{}

    // Match packet length.  Should be value 0..65535 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    PacketLength []interface{}

    // Match MPLS Label Disposition IPv4 access list. The type is string with
    // length: 1..32.
    MplsDispositionIpv4AccessList interface{}

    // Match MPLS Label Disposition IPv6 access list. The type is string with
    // length: 1..32.
    MplsDispositionIpv6AccessList interface{}

    // Match VLAN ID. The type is slice of string with pattern: (\d+)|(\d+\-\d+).
    Vlan []interface{}

    // Match inner VLAN ID. The type is slice of string with pattern:
    // (\d+)|(\d+\-\d+).
    InnerVlan []interface{}

    // Match flow-tag. Should be value 1..63 or range. The type is slice of string
    // with pattern: (\d+)|(\d+\-\d+).
    FlowTag []interface{}

    // Match Ethertype. The type is slice of string with pattern:
    // ((153[6-9]|15[4-9][0-9]|1[6-9][0-9][0-9]|[2-9][0-9][0-9][0-9])|([1-5][0-9][0-9][0-9][0-9]|6[0-4][0-9][0-9][0-9])|(65[0-4][0-9][0-9]|655[0-2][0-9]|6553[0-5]))|((arp)|(ipv4)|(ipv6)).
    Ethertype []interface{}

    // Match destination port.  Should be value 0..65535 or range. The type is
    // slice of string with pattern: (\d+)|(\d+\-\d+).
    DestinationPort []interface{}

    // Match fragment type for a packet. The type is slice of string with pattern:
    // (first-fragment)|(is-fragment)|(last-fragment).
    FragmentType []interface{}

    // Match frame-relay DLCI value.  Should be value 16..1007 or range. The type
    // is slice of string with pattern: (\d+)|(\d+\-\d+).
    FrameRelayDlci []interface{}

    // Set FrameRelay DE bit. The type is interface{} with range: 0..1.
    FrDe interface{}

    // Match IPv4 ICMP code.  Should be value 0..255 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    Icmpv4Code []interface{}

    // Match IPv4 ICMP type.  Should be value 0..255 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    Icmpv4Type []interface{}

    // Match IPv6 ICMP code.  Should be value 0..255 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    Icmpv6Code []interface{}

    // Match IPv6 ICMP type.  Should be value 0..255 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    Icmpv6Type []interface{}

    // Match source port.  Should be value 0..65535 or range. The type is slice of
    // string with pattern: (\d+)|(\d+\-\d+).
    SourcePort []interface{}

    // Match TCP flags. The type is interface{} with range: 0..4095.
    TcpFlag interface{}

    // Match authentication status. The type is string with pattern:
    // (authenticated)|(unauthenticated).
    AuthenStatus interface{}

    // Match Circuit ID. The type is slice of string with length: 1..32.
    CircuitId []interface{}

    // Match Circuit id regex. The type is slice of string with length: 1..32.
    CircuitIdRegex []interface{}

    // Match remote ID. The type is slice of string with length: 1..32.
    RemoteId []interface{}

    // Match remote id regex. The type is slice of string with length: 1..32.
    RemoteIdRegex []interface{}

    // Match servicve name. The type is slice of string with length: 1..32.
    ServiceName []interface{}

    // Match servicve name regular expression. The type is slice of string with
    // length: 1..32.
    ServiceNameRegex []interface{}

    // Match timer. The type is slice of string with length: 1..32.
    Timer []interface{}

    // Match timer regular expression. The type is slice of string with length:
    // 1..32.
    TimerRegex []interface{}

    // Match user name. The type is slice of string with length: 1..32.
    UserName []interface{}

    // Match user name regular expression. The type is slice of string with
    // length: 1..32.
    UserNameRegex []interface{}

    // Match source MAC address. The type is slice of string.
    SourceMac []interface{}

    // Match destination MAC address. The type is slice of string.
    DestinationMac []interface{}

    // Match VPLS control. The type is interface{}.
    VplsControl interface{}

    // Match VPLS Broadcast. The type is interface{}.
    VplsBroadcast interface{}

    // Match VPLS Multicast. The type is interface{}.
    VplsMulticast interface{}

    // Match VPLS Known. The type is interface{}.
    VplsKnown interface{}

    // Match VPLS Unknown. The type is interface{}.
    VplsUnknown interface{}

    // Match ATM CLP bit. The type is interface{} with range: 0..1.
    AtmClp interface{}

    // Match ATM OAM. The type is interface{}.
    AtmOam interface{}

    // Match CAC admitted. The type is interface{}.
    CacAdmit interface{}

    // Match CAC unadmitted. The type is interface{}.
    CacUnadmit interface{}

    // Match destination IPv4 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4.
    DestinationAddressIpv4 []PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4

    // Match destination IPv6 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6.
    DestinationAddressIpv6 []PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6

    // Match source IPv4 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4.
    SourceAddressIpv4 []PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4

    // Match source IPv6 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6.
    SourceAddressIpv6 []PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6

    // Match dhcp client ID. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId.
    DhcpClientId []PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId

    // Match dhcp client id regex. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex.
    DhcpClientIdRegex []PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex

    // Match domain name. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DomainName.
    DomainName []PolicyManager_ClassMaps_ClassMap_Match_DomainName

    // Match domain name. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex.
    DomainNameRegex []PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex

    // Match flow.
    Flow PolicyManager_ClassMaps_ClassMap_Match_Flow
}

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetFilter() yfilter.YFilter { return match.YFilter }

func (match *PolicyManager_ClassMaps_ClassMap_Match) SetFilter(yf yfilter.YFilter) { match.YFilter = yf }

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetGoName(yname string) string {
    if yname == "ipv4-dscp" { return "Ipv4Dscp" }
    if yname == "ipv6-dscp" { return "Ipv6Dscp" }
    if yname == "dscp" { return "Dscp" }
    if yname == "ipv4-precedence" { return "Ipv4Precedence" }
    if yname == "ipv6-precedence" { return "Ipv6Precedence" }
    if yname == "precedence" { return "Precedence" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "cos" { return "Cos" }
    if yname == "inner-cos" { return "InnerCos" }
    if yname == "dei" { return "Dei" }
    if yname == "dei-inner" { return "DeiInner" }
    if yname == "protocol" { return "Protocol" }
    if yname == "ipv4-acl" { return "Ipv4Acl" }
    if yname == "ipv6-acl" { return "Ipv6Acl" }
    if yname == "ethernet-services-acl" { return "EthernetServicesAcl" }
    if yname == "mpls-experimental-topmost" { return "MplsExperimentalTopmost" }
    if yname == "mpls-experimental-imposition" { return "MplsExperimentalImposition" }
    if yname == "discard-class" { return "DiscardClass" }
    if yname == "ipv4-packet-length" { return "Ipv4PacketLength" }
    if yname == "ipv6-packet-length" { return "Ipv6PacketLength" }
    if yname == "packet-length" { return "PacketLength" }
    if yname == "mpls-disposition-ipv4-access-list" { return "MplsDispositionIpv4AccessList" }
    if yname == "mpls-disposition-ipv6-access-list" { return "MplsDispositionIpv6AccessList" }
    if yname == "vlan" { return "Vlan" }
    if yname == "inner-vlan" { return "InnerVlan" }
    if yname == "flow-tag" { return "FlowTag" }
    if yname == "ethertype" { return "Ethertype" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "fragment-type" { return "FragmentType" }
    if yname == "frame-relay-dlci" { return "FrameRelayDlci" }
    if yname == "fr-de" { return "FrDe" }
    if yname == "icmpv4-code" { return "Icmpv4Code" }
    if yname == "icmpv4-type" { return "Icmpv4Type" }
    if yname == "icmpv6-code" { return "Icmpv6Code" }
    if yname == "icmpv6-type" { return "Icmpv6Type" }
    if yname == "source-port" { return "SourcePort" }
    if yname == "tcp-flag" { return "TcpFlag" }
    if yname == "authen-status" { return "AuthenStatus" }
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "circuit-id-regex" { return "CircuitIdRegex" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "remote-id-regex" { return "RemoteIdRegex" }
    if yname == "service-name" { return "ServiceName" }
    if yname == "service-name-regex" { return "ServiceNameRegex" }
    if yname == "timer" { return "Timer" }
    if yname == "timer-regex" { return "TimerRegex" }
    if yname == "user-name" { return "UserName" }
    if yname == "user-name-regex" { return "UserNameRegex" }
    if yname == "source-mac" { return "SourceMac" }
    if yname == "destination-mac" { return "DestinationMac" }
    if yname == "vpls-control" { return "VplsControl" }
    if yname == "vpls-broadcast" { return "VplsBroadcast" }
    if yname == "vpls-multicast" { return "VplsMulticast" }
    if yname == "vpls-known" { return "VplsKnown" }
    if yname == "vpls-unknown" { return "VplsUnknown" }
    if yname == "atm-clp" { return "AtmClp" }
    if yname == "atm-oam" { return "AtmOam" }
    if yname == "cac-admit" { return "CacAdmit" }
    if yname == "cac-unadmit" { return "CacUnadmit" }
    if yname == "destination-address-ipv4" { return "DestinationAddressIpv4" }
    if yname == "destination-address-ipv6" { return "DestinationAddressIpv6" }
    if yname == "source-address-ipv4" { return "SourceAddressIpv4" }
    if yname == "source-address-ipv6" { return "SourceAddressIpv6" }
    if yname == "dhcp-client-id" { return "DhcpClientId" }
    if yname == "dhcp-client-id-regex" { return "DhcpClientIdRegex" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "domain-name-regex" { return "DomainNameRegex" }
    if yname == "flow" { return "Flow" }
    return ""
}

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetSegmentPath() string {
    return "match"
}

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-address-ipv4" {
        for _, c := range match.DestinationAddressIpv4 {
            if match.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4{}
        match.DestinationAddressIpv4 = append(match.DestinationAddressIpv4, child)
        return &match.DestinationAddressIpv4[len(match.DestinationAddressIpv4)-1]
    }
    if childYangName == "destination-address-ipv6" {
        for _, c := range match.DestinationAddressIpv6 {
            if match.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6{}
        match.DestinationAddressIpv6 = append(match.DestinationAddressIpv6, child)
        return &match.DestinationAddressIpv6[len(match.DestinationAddressIpv6)-1]
    }
    if childYangName == "source-address-ipv4" {
        for _, c := range match.SourceAddressIpv4 {
            if match.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4{}
        match.SourceAddressIpv4 = append(match.SourceAddressIpv4, child)
        return &match.SourceAddressIpv4[len(match.SourceAddressIpv4)-1]
    }
    if childYangName == "source-address-ipv6" {
        for _, c := range match.SourceAddressIpv6 {
            if match.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6{}
        match.SourceAddressIpv6 = append(match.SourceAddressIpv6, child)
        return &match.SourceAddressIpv6[len(match.SourceAddressIpv6)-1]
    }
    if childYangName == "dhcp-client-id" {
        for _, c := range match.DhcpClientId {
            if match.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId{}
        match.DhcpClientId = append(match.DhcpClientId, child)
        return &match.DhcpClientId[len(match.DhcpClientId)-1]
    }
    if childYangName == "dhcp-client-id-regex" {
        for _, c := range match.DhcpClientIdRegex {
            if match.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex{}
        match.DhcpClientIdRegex = append(match.DhcpClientIdRegex, child)
        return &match.DhcpClientIdRegex[len(match.DhcpClientIdRegex)-1]
    }
    if childYangName == "domain-name" {
        for _, c := range match.DomainName {
            if match.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_Match_DomainName{}
        match.DomainName = append(match.DomainName, child)
        return &match.DomainName[len(match.DomainName)-1]
    }
    if childYangName == "domain-name-regex" {
        for _, c := range match.DomainNameRegex {
            if match.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex{}
        match.DomainNameRegex = append(match.DomainNameRegex, child)
        return &match.DomainNameRegex[len(match.DomainNameRegex)-1]
    }
    if childYangName == "flow" {
        return &match.Flow
    }
    return nil
}

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range match.DestinationAddressIpv4 {
        children[match.DestinationAddressIpv4[i].GetSegmentPath()] = &match.DestinationAddressIpv4[i]
    }
    for i := range match.DestinationAddressIpv6 {
        children[match.DestinationAddressIpv6[i].GetSegmentPath()] = &match.DestinationAddressIpv6[i]
    }
    for i := range match.SourceAddressIpv4 {
        children[match.SourceAddressIpv4[i].GetSegmentPath()] = &match.SourceAddressIpv4[i]
    }
    for i := range match.SourceAddressIpv6 {
        children[match.SourceAddressIpv6[i].GetSegmentPath()] = &match.SourceAddressIpv6[i]
    }
    for i := range match.DhcpClientId {
        children[match.DhcpClientId[i].GetSegmentPath()] = &match.DhcpClientId[i]
    }
    for i := range match.DhcpClientIdRegex {
        children[match.DhcpClientIdRegex[i].GetSegmentPath()] = &match.DhcpClientIdRegex[i]
    }
    for i := range match.DomainName {
        children[match.DomainName[i].GetSegmentPath()] = &match.DomainName[i]
    }
    for i := range match.DomainNameRegex {
        children[match.DomainNameRegex[i].GetSegmentPath()] = &match.DomainNameRegex[i]
    }
    children["flow"] = &match.Flow
    return children
}

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-dscp"] = match.Ipv4Dscp
    leafs["ipv6-dscp"] = match.Ipv6Dscp
    leafs["dscp"] = match.Dscp
    leafs["ipv4-precedence"] = match.Ipv4Precedence
    leafs["ipv6-precedence"] = match.Ipv6Precedence
    leafs["precedence"] = match.Precedence
    leafs["qos-group"] = match.QosGroup
    leafs["traffic-class"] = match.TrafficClass
    leafs["cos"] = match.Cos
    leafs["inner-cos"] = match.InnerCos
    leafs["dei"] = match.Dei
    leafs["dei-inner"] = match.DeiInner
    leafs["protocol"] = match.Protocol
    leafs["ipv4-acl"] = match.Ipv4Acl
    leafs["ipv6-acl"] = match.Ipv6Acl
    leafs["ethernet-services-acl"] = match.EthernetServicesAcl
    leafs["mpls-experimental-topmost"] = match.MplsExperimentalTopmost
    leafs["mpls-experimental-imposition"] = match.MplsExperimentalImposition
    leafs["discard-class"] = match.DiscardClass
    leafs["ipv4-packet-length"] = match.Ipv4PacketLength
    leafs["ipv6-packet-length"] = match.Ipv6PacketLength
    leafs["packet-length"] = match.PacketLength
    leafs["mpls-disposition-ipv4-access-list"] = match.MplsDispositionIpv4AccessList
    leafs["mpls-disposition-ipv6-access-list"] = match.MplsDispositionIpv6AccessList
    leafs["vlan"] = match.Vlan
    leafs["inner-vlan"] = match.InnerVlan
    leafs["flow-tag"] = match.FlowTag
    leafs["ethertype"] = match.Ethertype
    leafs["destination-port"] = match.DestinationPort
    leafs["fragment-type"] = match.FragmentType
    leafs["frame-relay-dlci"] = match.FrameRelayDlci
    leafs["fr-de"] = match.FrDe
    leafs["icmpv4-code"] = match.Icmpv4Code
    leafs["icmpv4-type"] = match.Icmpv4Type
    leafs["icmpv6-code"] = match.Icmpv6Code
    leafs["icmpv6-type"] = match.Icmpv6Type
    leafs["source-port"] = match.SourcePort
    leafs["tcp-flag"] = match.TcpFlag
    leafs["authen-status"] = match.AuthenStatus
    leafs["circuit-id"] = match.CircuitId
    leafs["circuit-id-regex"] = match.CircuitIdRegex
    leafs["remote-id"] = match.RemoteId
    leafs["remote-id-regex"] = match.RemoteIdRegex
    leafs["service-name"] = match.ServiceName
    leafs["service-name-regex"] = match.ServiceNameRegex
    leafs["timer"] = match.Timer
    leafs["timer-regex"] = match.TimerRegex
    leafs["user-name"] = match.UserName
    leafs["user-name-regex"] = match.UserNameRegex
    leafs["source-mac"] = match.SourceMac
    leafs["destination-mac"] = match.DestinationMac
    leafs["vpls-control"] = match.VplsControl
    leafs["vpls-broadcast"] = match.VplsBroadcast
    leafs["vpls-multicast"] = match.VplsMulticast
    leafs["vpls-known"] = match.VplsKnown
    leafs["vpls-unknown"] = match.VplsUnknown
    leafs["atm-clp"] = match.AtmClp
    leafs["atm-oam"] = match.AtmOam
    leafs["cac-admit"] = match.CacAdmit
    leafs["cac-unadmit"] = match.CacUnadmit
    return leafs
}

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetBundleName() string { return "cisco_ios_xr" }

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetYangName() string { return "match" }

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (match *PolicyManager_ClassMaps_ClassMap_Match) SetParent(parent types.Entity) { match.parent = parent }

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetParent() types.Entity { return match.parent }

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetParentYangName() string { return "class-map" }

// PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4
// Match destination IPv4 address.
type PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv4 netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetFilter() yfilter.YFilter { return destinationAddressIpv4.YFilter }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) SetFilter(yf yfilter.YFilter) { destinationAddressIpv4.YFilter = yf }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "netmask" { return "Netmask" }
    return ""
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetSegmentPath() string {
    return "destination-address-ipv4" + "[address='" + fmt.Sprintf("%v", destinationAddressIpv4.Address) + "']" + "[netmask='" + fmt.Sprintf("%v", destinationAddressIpv4.Netmask) + "']"
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = destinationAddressIpv4.Address
    leafs["netmask"] = destinationAddressIpv4.Netmask
    return leafs
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetYangName() string { return "destination-address-ipv4" }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) SetParent(parent types.Entity) { destinationAddressIpv4.parent = parent }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetParent() types.Entity { return destinationAddressIpv4.parent }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetParentYangName() string { return "match" }

// PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6
// Match destination IPv6 address.
type PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetFilter() yfilter.YFilter { return destinationAddressIpv6.YFilter }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) SetFilter(yf yfilter.YFilter) { destinationAddressIpv6.YFilter = yf }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetSegmentPath() string {
    return "destination-address-ipv6" + "[address='" + fmt.Sprintf("%v", destinationAddressIpv6.Address) + "']" + "[prefix-length='" + fmt.Sprintf("%v", destinationAddressIpv6.PrefixLength) + "']"
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = destinationAddressIpv6.Address
    leafs["prefix-length"] = destinationAddressIpv6.PrefixLength
    return leafs
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetYangName() string { return "destination-address-ipv6" }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) SetParent(parent types.Entity) { destinationAddressIpv6.parent = parent }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetParent() types.Entity { return destinationAddressIpv6.parent }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetParentYangName() string { return "match" }

// PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4
// Match source IPv4 address.
type PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv4 netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetFilter() yfilter.YFilter { return sourceAddressIpv4.YFilter }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) SetFilter(yf yfilter.YFilter) { sourceAddressIpv4.YFilter = yf }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "netmask" { return "Netmask" }
    return ""
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetSegmentPath() string {
    return "source-address-ipv4" + "[address='" + fmt.Sprintf("%v", sourceAddressIpv4.Address) + "']" + "[netmask='" + fmt.Sprintf("%v", sourceAddressIpv4.Netmask) + "']"
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = sourceAddressIpv4.Address
    leafs["netmask"] = sourceAddressIpv4.Netmask
    return leafs
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetYangName() string { return "source-address-ipv4" }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) SetParent(parent types.Entity) { sourceAddressIpv4.parent = parent }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetParent() types.Entity { return sourceAddressIpv4.parent }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetParentYangName() string { return "match" }

// PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6
// Match source IPv6 address.
type PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetFilter() yfilter.YFilter { return sourceAddressIpv6.YFilter }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) SetFilter(yf yfilter.YFilter) { sourceAddressIpv6.YFilter = yf }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetSegmentPath() string {
    return "source-address-ipv6" + "[address='" + fmt.Sprintf("%v", sourceAddressIpv6.Address) + "']" + "[prefix-length='" + fmt.Sprintf("%v", sourceAddressIpv6.PrefixLength) + "']"
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = sourceAddressIpv6.Address
    leafs["prefix-length"] = sourceAddressIpv6.PrefixLength
    return leafs
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetYangName() string { return "source-address-ipv6" }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) SetParent(parent types.Entity) { sourceAddressIpv6.parent = parent }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetParent() types.Entity { return sourceAddressIpv6.parent }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetParentYangName() string { return "match" }

// PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId
// Match dhcp client ID.
type PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Dhcp client Id. The type is string with length:
    // 1..32.
    Value interface{}

    // This attribute is a key. Dhcp client id Ascii/Hex. The type is string with
    // pattern: (none)|(ascii)|(hex).
    Flag interface{}
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetFilter() yfilter.YFilter { return dhcpClientId.YFilter }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) SetFilter(yf yfilter.YFilter) { dhcpClientId.YFilter = yf }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "flag" { return "Flag" }
    return ""
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetSegmentPath() string {
    return "dhcp-client-id" + "[value='" + fmt.Sprintf("%v", dhcpClientId.Value) + "']" + "[flag='" + fmt.Sprintf("%v", dhcpClientId.Flag) + "']"
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = dhcpClientId.Value
    leafs["flag"] = dhcpClientId.Flag
    return leafs
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetYangName() string { return "dhcp-client-id" }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) SetParent(parent types.Entity) { dhcpClientId.parent = parent }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetParent() types.Entity { return dhcpClientId.parent }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetParentYangName() string { return "match" }

// PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex
// Match dhcp client id regex.
type PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Dhcp client id regular expression. The type is
    // string with length: 1..32.
    Value interface{}

    // This attribute is a key. Dhcp client Id regex Ascii/Hex. The type is string
    // with pattern: (none)|(ascii)|(hex).
    Flag interface{}
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetFilter() yfilter.YFilter { return dhcpClientIdRegex.YFilter }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) SetFilter(yf yfilter.YFilter) { dhcpClientIdRegex.YFilter = yf }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "flag" { return "Flag" }
    return ""
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetSegmentPath() string {
    return "dhcp-client-id-regex" + "[value='" + fmt.Sprintf("%v", dhcpClientIdRegex.Value) + "']" + "[flag='" + fmt.Sprintf("%v", dhcpClientIdRegex.Flag) + "']"
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = dhcpClientIdRegex.Value
    leafs["flag"] = dhcpClientIdRegex.Flag
    return leafs
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetYangName() string { return "dhcp-client-id-regex" }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) SetParent(parent types.Entity) { dhcpClientIdRegex.parent = parent }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetParent() types.Entity { return dhcpClientIdRegex.parent }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetParentYangName() string { return "match" }

// PolicyManager_ClassMaps_ClassMap_Match_DomainName
// Match domain name.
type PolicyManager_ClassMaps_ClassMap_Match_DomainName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Domain name or regular expression. The type is
    // string with length: 1..32.
    Name interface{}

    // This attribute is a key. Domain-format name. The type is string with
    // length: 1..32.
    Format interface{}
}

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetFilter() yfilter.YFilter { return domainName.YFilter }

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) SetFilter(yf yfilter.YFilter) { domainName.YFilter = yf }

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "format" { return "Format" }
    return ""
}

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetSegmentPath() string {
    return "domain-name" + "[name='" + fmt.Sprintf("%v", domainName.Name) + "']" + "[format='" + fmt.Sprintf("%v", domainName.Format) + "']"
}

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = domainName.Name
    leafs["format"] = domainName.Format
    return leafs
}

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetBundleName() string { return "cisco_ios_xr" }

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetYangName() string { return "domain-name" }

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) SetParent(parent types.Entity) { domainName.parent = parent }

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetParent() types.Entity { return domainName.parent }

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetParentYangName() string { return "match" }

// PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex
// Match domain name.
type PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Domain name or regular expression. The type is
    // string with length: 1..32.
    Regex interface{}

    // This attribute is a key. Domain-format name. The type is string with
    // length: 1..32.
    Format interface{}
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetFilter() yfilter.YFilter { return domainNameRegex.YFilter }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) SetFilter(yf yfilter.YFilter) { domainNameRegex.YFilter = yf }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetGoName(yname string) string {
    if yname == "regex" { return "Regex" }
    if yname == "format" { return "Format" }
    return ""
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetSegmentPath() string {
    return "domain-name-regex" + "[regex='" + fmt.Sprintf("%v", domainNameRegex.Regex) + "']" + "[format='" + fmt.Sprintf("%v", domainNameRegex.Format) + "']"
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["regex"] = domainNameRegex.Regex
    leafs["format"] = domainNameRegex.Format
    return leafs
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetBundleName() string { return "cisco_ios_xr" }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetYangName() string { return "domain-name-regex" }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) SetParent(parent types.Entity) { domainNameRegex.parent = parent }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetParent() types.Entity { return domainNameRegex.parent }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetParentYangName() string { return "match" }

// PolicyManager_ClassMaps_ClassMap_Match_Flow
// Match flow.
type PolicyManager_ClassMaps_ClassMap_Match_Flow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the flow-key parameters. The type is slice of string with
    // pattern: (SourceIP)|(DestinationIP)|(5Tuple).
    FlowKey []interface{}

    // Configure the flow-cache parameters.
    FlowCache PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache
}

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetFilter() yfilter.YFilter { return flow.YFilter }

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) SetFilter(yf yfilter.YFilter) { flow.YFilter = yf }

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetGoName(yname string) string {
    if yname == "flow-key" { return "FlowKey" }
    if yname == "flow-cache" { return "FlowCache" }
    return ""
}

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetSegmentPath() string {
    return "flow"
}

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-cache" {
        return &flow.FlowCache
    }
    return nil
}

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flow-cache"] = &flow.FlowCache
    return children
}

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-key"] = flow.FlowKey
    return leafs
}

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetBundleName() string { return "cisco_ios_xr" }

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetYangName() string { return "flow" }

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) SetParent(parent types.Entity) { flow.parent = parent }

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetParent() types.Entity { return flow.parent }

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetParentYangName() string { return "match" }

// PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache
// Configure the flow-cache parameters
type PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum time of inactivity for a flow. The type is one of the following
    // types: int with range: 10..2550, or string with pattern: (None)|(none).
    IdleTimeout interface{}
}

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetFilter() yfilter.YFilter { return flowCache.YFilter }

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) SetFilter(yf yfilter.YFilter) { flowCache.YFilter = yf }

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetGoName(yname string) string {
    if yname == "idle-timeout" { return "IdleTimeout" }
    return ""
}

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetSegmentPath() string {
    return "flow-cache"
}

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["idle-timeout"] = flowCache.IdleTimeout
    return leafs
}

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetBundleName() string { return "cisco_ios_xr" }

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetYangName() string { return "flow-cache" }

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) SetParent(parent types.Entity) { flowCache.parent = parent }

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetParent() types.Entity { return flowCache.parent }

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetParentYangName() string { return "flow" }

// PolicyManager_ClassMaps_ClassMap_MatchNot
// Match not rules.
type PolicyManager_ClassMaps_ClassMap_MatchNot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Match IPv4 DSCP. The type is slice of string with pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(([0-9]|[1-5][0-9]|6[0-3])-([0-9]|[1-5][0-9]|6[0-3]))|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Ipv4Dscp []interface{}

    // Match IPv6 DSCP. The type is slice of string with pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(([0-9]|[1-5][0-9]|6[0-3])-([0-9]|[1-5][0-9]|6[0-3]))|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Ipv6Dscp []interface{}

    // Match DSCP. The type is slice of string with pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(([0-9]|[1-5][0-9]|6[0-3])-([0-9]|[1-5][0-9]|6[0-3]))|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Dscp []interface{}

    // Match IPv4 precedence. The type is one of the following types: slice of int
    // with range: 0..7, or slice of string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Ipv4Precedence []interface{}

    // Match IPv6 precedence. The type is one of the following types: slice of int
    // with range: 0..7, or slice of string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Ipv6Precedence []interface{}

    // Match precedence. The type is one of the following types: slice of int with
    // range: 0..7, or slice of string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Precedence []interface{}

    // Match QoS group. Should be value 0..512 or range. The type is slice of
    // string with pattern: (\d+)|(\d+\-\d+).
    QosGroup []interface{}

    // Match Traffic Class. Should be value 0..63 or range. The type is slice of
    // string with pattern: (\d+)|(\d+\-\d+).
    TrafficClass []interface{}

    // Match CoS. The type is slice of interface{} with range: 0..7.
    Cos []interface{}

    // Match inner CoS. The type is slice of interface{} with range: 0..7.
    InnerCos []interface{}

    // Match DEI bit. The type is interface{} with range: 0..1.
    Dei interface{}

    // Match DEI INNER  bit. The type is interface{} with range: 0..1.
    DeiInner interface{}

    // Match protocol. The type is slice of string with pattern:
    // ([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\-([1-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5]))|((ahp)|(dhcpv4)|(dhcpv6)|(eigrp)|(esp)|(gre)|(icmp)|(igmp)|(igrp)|(ipinip)|(ipv4)|(ipv6)|(ipv6icmp)|(mpls)|(nos)|(ospf)|(pcp)|(pim)|(ppp)|(sctp)|(tcp)|(udp)).
    Protocol []interface{}

    // Match IPv4 ACL. The type is string with length: 1..64.
    Ipv4Acl interface{}

    // Match IPv6 ACL. The type is string with length: 1..64.
    Ipv6Acl interface{}

    // Match Ethernet Services. The type is string with length: 1..64.
    EthernetServicesAcl interface{}

    // Match MPLS experimental topmost label. The type is slice of interface{}
    // with range: 0..7.
    MplsExperimentalTopmost []interface{}

    // Match MPLS experimental imposition label. The type is slice of interface{}
    // with range: 0..7.
    MplsExperimentalImposition []interface{}

    // Match discard class. The type is slice of interface{} with range: 0..7.
    DiscardClass []interface{}

    // Match IPv4 packet length. Should be value 0..65535 or range. The type is
    // slice of string with pattern: (\d+)|(\d+\-\d+).
    Ipv4PacketLength []interface{}

    // Match IPv6 packet length.  Should be value 0..65535 or range. The type is
    // slice of string with pattern: (\d+)|(\d+\-\d+).
    Ipv6PacketLength []interface{}

    // Match packet length.  Should be value 0..65535 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    PacketLength []interface{}

    // Match MPLS Label Disposition IPv4 access list. The type is string with
    // length: 1..32.
    MplsDispositionIpv4AccessList interface{}

    // Match MPLS Label Disposition IPv6 access list. The type is string with
    // length: 1..32.
    MplsDispositionIpv6AccessList interface{}

    // Match VLAN ID. The type is slice of string with pattern: (\d+)|(\d+\-\d+).
    Vlan []interface{}

    // Match inner VLAN ID. The type is slice of string with pattern:
    // (\d+)|(\d+\-\d+).
    InnerVlan []interface{}

    // Match flow-tag. Should be value 1..63 or range. The type is slice of string
    // with pattern: (\d+)|(\d+\-\d+).
    FlowTag []interface{}

    // Match Ethertype. The type is slice of string with pattern:
    // ((153[6-9]|15[4-9][0-9]|1[6-9][0-9][0-9]|[2-9][0-9][0-9][0-9])|([1-5][0-9][0-9][0-9][0-9]|6[0-4][0-9][0-9][0-9])|(65[0-4][0-9][0-9]|655[0-2][0-9]|6553[0-5]))|((arp)|(ipv4)|(ipv6)).
    Ethertype []interface{}

    // Match destination port.  Should be value 0..65535 or range. The type is
    // slice of string with pattern: (\d+)|(\d+\-\d+).
    DestinationPort []interface{}

    // Match fragment type for a packet. The type is slice of string with pattern:
    // (first-fragment)|(is-fragment)|(last-fragment).
    FragmentType []interface{}

    // Match frame-relay DLCI value.  Should be value 16..1007 or range. The type
    // is slice of string with pattern: (\d+)|(\d+\-\d+).
    FrameRelayDlci []interface{}

    // Set FrameRelay DE bit. The type is interface{} with range: 0..1.
    FrDe interface{}

    // Match IPv4 ICMP code.  Should be value 0..255 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    Icmpv4Code []interface{}

    // Match IPv4 ICMP type.  Should be value 0..255 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    Icmpv4Type []interface{}

    // Match IPv6 ICMP code.  Should be value 0..255 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    Icmpv6Code []interface{}

    // Match IPv6 ICMP type.  Should be value 0..255 or range. The type is slice
    // of string with pattern: (\d+)|(\d+\-\d+).
    Icmpv6Type []interface{}

    // Match source port.  Should be value 0..65535 or range. The type is slice of
    // string with pattern: (\d+)|(\d+\-\d+).
    SourcePort []interface{}

    // Match TCP flags. The type is interface{} with range: 0..4095.
    TcpFlag interface{}

    // Match authentication status. The type is string with pattern:
    // (authenticated)|(unauthenticated).
    AuthenStatus interface{}

    // Match Circuit ID. The type is slice of string with length: 1..32.
    CircuitId []interface{}

    // Match Circuit id regex. The type is slice of string with length: 1..32.
    CircuitIdRegex []interface{}

    // Match remote ID. The type is slice of string with length: 1..32.
    RemoteId []interface{}

    // Match remote id regex. The type is slice of string with length: 1..32.
    RemoteIdRegex []interface{}

    // Match servicve name. The type is slice of string with length: 1..32.
    ServiceName []interface{}

    // Match servicve name regular expression. The type is slice of string with
    // length: 1..32.
    ServiceNameRegex []interface{}

    // Match timer. The type is slice of string with length: 1..32.
    Timer []interface{}

    // Match timer regular expression. The type is slice of string with length:
    // 1..32.
    TimerRegex []interface{}

    // Match user name. The type is slice of string with length: 1..32.
    UserName []interface{}

    // Match user name regular expression. The type is slice of string with
    // length: 1..32.
    UserNameRegex []interface{}

    // Match source MAC address. The type is slice of string.
    SourceMac []interface{}

    // Match destination MAC address. The type is slice of string.
    DestinationMac []interface{}

    // Match VPLS control. The type is interface{}.
    VplsControl interface{}

    // Match VPLS Broadcast. The type is interface{}.
    VplsBroadcast interface{}

    // Match VPLS Multicast. The type is interface{}.
    VplsMulticast interface{}

    // Match VPLS Known. The type is interface{}.
    VplsKnown interface{}

    // Match VPLS Unknown. The type is interface{}.
    VplsUnknown interface{}

    // Match destination IPv4 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4.
    DestinationAddressIpv4 []PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4

    // Match destination IPv6 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6.
    DestinationAddressIpv6 []PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6

    // Match source IPv4 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4.
    SourceAddressIpv4 []PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4

    // Match source IPv6 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6.
    SourceAddressIpv6 []PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6

    // Match dhcp client ID. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId.
    DhcpClientId []PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId

    // Match dhcp client id regex. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex.
    DhcpClientIdRegex []PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex

    // Match domain name. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName.
    DomainName []PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName

    // Match domain name. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex.
    DomainNameRegex []PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex

    // Match flow.
    Flow PolicyManager_ClassMaps_ClassMap_MatchNot_Flow
}

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetFilter() yfilter.YFilter { return matchNot.YFilter }

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) SetFilter(yf yfilter.YFilter) { matchNot.YFilter = yf }

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetGoName(yname string) string {
    if yname == "ipv4-dscp" { return "Ipv4Dscp" }
    if yname == "ipv6-dscp" { return "Ipv6Dscp" }
    if yname == "dscp" { return "Dscp" }
    if yname == "ipv4-precedence" { return "Ipv4Precedence" }
    if yname == "ipv6-precedence" { return "Ipv6Precedence" }
    if yname == "precedence" { return "Precedence" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "cos" { return "Cos" }
    if yname == "inner-cos" { return "InnerCos" }
    if yname == "dei" { return "Dei" }
    if yname == "dei-inner" { return "DeiInner" }
    if yname == "protocol" { return "Protocol" }
    if yname == "ipv4-acl" { return "Ipv4Acl" }
    if yname == "ipv6-acl" { return "Ipv6Acl" }
    if yname == "ethernet-services-acl" { return "EthernetServicesAcl" }
    if yname == "mpls-experimental-topmost" { return "MplsExperimentalTopmost" }
    if yname == "mpls-experimental-imposition" { return "MplsExperimentalImposition" }
    if yname == "discard-class" { return "DiscardClass" }
    if yname == "ipv4-packet-length" { return "Ipv4PacketLength" }
    if yname == "ipv6-packet-length" { return "Ipv6PacketLength" }
    if yname == "packet-length" { return "PacketLength" }
    if yname == "mpls-disposition-ipv4-access-list" { return "MplsDispositionIpv4AccessList" }
    if yname == "mpls-disposition-ipv6-access-list" { return "MplsDispositionIpv6AccessList" }
    if yname == "vlan" { return "Vlan" }
    if yname == "inner-vlan" { return "InnerVlan" }
    if yname == "flow-tag" { return "FlowTag" }
    if yname == "ethertype" { return "Ethertype" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "fragment-type" { return "FragmentType" }
    if yname == "frame-relay-dlci" { return "FrameRelayDlci" }
    if yname == "fr-de" { return "FrDe" }
    if yname == "icmpv4-code" { return "Icmpv4Code" }
    if yname == "icmpv4-type" { return "Icmpv4Type" }
    if yname == "icmpv6-code" { return "Icmpv6Code" }
    if yname == "icmpv6-type" { return "Icmpv6Type" }
    if yname == "source-port" { return "SourcePort" }
    if yname == "tcp-flag" { return "TcpFlag" }
    if yname == "authen-status" { return "AuthenStatus" }
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "circuit-id-regex" { return "CircuitIdRegex" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "remote-id-regex" { return "RemoteIdRegex" }
    if yname == "service-name" { return "ServiceName" }
    if yname == "service-name-regex" { return "ServiceNameRegex" }
    if yname == "timer" { return "Timer" }
    if yname == "timer-regex" { return "TimerRegex" }
    if yname == "user-name" { return "UserName" }
    if yname == "user-name-regex" { return "UserNameRegex" }
    if yname == "source-mac" { return "SourceMac" }
    if yname == "destination-mac" { return "DestinationMac" }
    if yname == "vpls-control" { return "VplsControl" }
    if yname == "vpls-broadcast" { return "VplsBroadcast" }
    if yname == "vpls-multicast" { return "VplsMulticast" }
    if yname == "vpls-known" { return "VplsKnown" }
    if yname == "vpls-unknown" { return "VplsUnknown" }
    if yname == "destination-address-ipv4" { return "DestinationAddressIpv4" }
    if yname == "destination-address-ipv6" { return "DestinationAddressIpv6" }
    if yname == "source-address-ipv4" { return "SourceAddressIpv4" }
    if yname == "source-address-ipv6" { return "SourceAddressIpv6" }
    if yname == "dhcp-client-id" { return "DhcpClientId" }
    if yname == "dhcp-client-id-regex" { return "DhcpClientIdRegex" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "domain-name-regex" { return "DomainNameRegex" }
    if yname == "flow" { return "Flow" }
    return ""
}

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetSegmentPath() string {
    return "match-not"
}

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "destination-address-ipv4" {
        for _, c := range matchNot.DestinationAddressIpv4 {
            if matchNot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4{}
        matchNot.DestinationAddressIpv4 = append(matchNot.DestinationAddressIpv4, child)
        return &matchNot.DestinationAddressIpv4[len(matchNot.DestinationAddressIpv4)-1]
    }
    if childYangName == "destination-address-ipv6" {
        for _, c := range matchNot.DestinationAddressIpv6 {
            if matchNot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6{}
        matchNot.DestinationAddressIpv6 = append(matchNot.DestinationAddressIpv6, child)
        return &matchNot.DestinationAddressIpv6[len(matchNot.DestinationAddressIpv6)-1]
    }
    if childYangName == "source-address-ipv4" {
        for _, c := range matchNot.SourceAddressIpv4 {
            if matchNot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4{}
        matchNot.SourceAddressIpv4 = append(matchNot.SourceAddressIpv4, child)
        return &matchNot.SourceAddressIpv4[len(matchNot.SourceAddressIpv4)-1]
    }
    if childYangName == "source-address-ipv6" {
        for _, c := range matchNot.SourceAddressIpv6 {
            if matchNot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6{}
        matchNot.SourceAddressIpv6 = append(matchNot.SourceAddressIpv6, child)
        return &matchNot.SourceAddressIpv6[len(matchNot.SourceAddressIpv6)-1]
    }
    if childYangName == "dhcp-client-id" {
        for _, c := range matchNot.DhcpClientId {
            if matchNot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId{}
        matchNot.DhcpClientId = append(matchNot.DhcpClientId, child)
        return &matchNot.DhcpClientId[len(matchNot.DhcpClientId)-1]
    }
    if childYangName == "dhcp-client-id-regex" {
        for _, c := range matchNot.DhcpClientIdRegex {
            if matchNot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex{}
        matchNot.DhcpClientIdRegex = append(matchNot.DhcpClientIdRegex, child)
        return &matchNot.DhcpClientIdRegex[len(matchNot.DhcpClientIdRegex)-1]
    }
    if childYangName == "domain-name" {
        for _, c := range matchNot.DomainName {
            if matchNot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName{}
        matchNot.DomainName = append(matchNot.DomainName, child)
        return &matchNot.DomainName[len(matchNot.DomainName)-1]
    }
    if childYangName == "domain-name-regex" {
        for _, c := range matchNot.DomainNameRegex {
            if matchNot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex{}
        matchNot.DomainNameRegex = append(matchNot.DomainNameRegex, child)
        return &matchNot.DomainNameRegex[len(matchNot.DomainNameRegex)-1]
    }
    if childYangName == "flow" {
        return &matchNot.Flow
    }
    return nil
}

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range matchNot.DestinationAddressIpv4 {
        children[matchNot.DestinationAddressIpv4[i].GetSegmentPath()] = &matchNot.DestinationAddressIpv4[i]
    }
    for i := range matchNot.DestinationAddressIpv6 {
        children[matchNot.DestinationAddressIpv6[i].GetSegmentPath()] = &matchNot.DestinationAddressIpv6[i]
    }
    for i := range matchNot.SourceAddressIpv4 {
        children[matchNot.SourceAddressIpv4[i].GetSegmentPath()] = &matchNot.SourceAddressIpv4[i]
    }
    for i := range matchNot.SourceAddressIpv6 {
        children[matchNot.SourceAddressIpv6[i].GetSegmentPath()] = &matchNot.SourceAddressIpv6[i]
    }
    for i := range matchNot.DhcpClientId {
        children[matchNot.DhcpClientId[i].GetSegmentPath()] = &matchNot.DhcpClientId[i]
    }
    for i := range matchNot.DhcpClientIdRegex {
        children[matchNot.DhcpClientIdRegex[i].GetSegmentPath()] = &matchNot.DhcpClientIdRegex[i]
    }
    for i := range matchNot.DomainName {
        children[matchNot.DomainName[i].GetSegmentPath()] = &matchNot.DomainName[i]
    }
    for i := range matchNot.DomainNameRegex {
        children[matchNot.DomainNameRegex[i].GetSegmentPath()] = &matchNot.DomainNameRegex[i]
    }
    children["flow"] = &matchNot.Flow
    return children
}

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-dscp"] = matchNot.Ipv4Dscp
    leafs["ipv6-dscp"] = matchNot.Ipv6Dscp
    leafs["dscp"] = matchNot.Dscp
    leafs["ipv4-precedence"] = matchNot.Ipv4Precedence
    leafs["ipv6-precedence"] = matchNot.Ipv6Precedence
    leafs["precedence"] = matchNot.Precedence
    leafs["qos-group"] = matchNot.QosGroup
    leafs["traffic-class"] = matchNot.TrafficClass
    leafs["cos"] = matchNot.Cos
    leafs["inner-cos"] = matchNot.InnerCos
    leafs["dei"] = matchNot.Dei
    leafs["dei-inner"] = matchNot.DeiInner
    leafs["protocol"] = matchNot.Protocol
    leafs["ipv4-acl"] = matchNot.Ipv4Acl
    leafs["ipv6-acl"] = matchNot.Ipv6Acl
    leafs["ethernet-services-acl"] = matchNot.EthernetServicesAcl
    leafs["mpls-experimental-topmost"] = matchNot.MplsExperimentalTopmost
    leafs["mpls-experimental-imposition"] = matchNot.MplsExperimentalImposition
    leafs["discard-class"] = matchNot.DiscardClass
    leafs["ipv4-packet-length"] = matchNot.Ipv4PacketLength
    leafs["ipv6-packet-length"] = matchNot.Ipv6PacketLength
    leafs["packet-length"] = matchNot.PacketLength
    leafs["mpls-disposition-ipv4-access-list"] = matchNot.MplsDispositionIpv4AccessList
    leafs["mpls-disposition-ipv6-access-list"] = matchNot.MplsDispositionIpv6AccessList
    leafs["vlan"] = matchNot.Vlan
    leafs["inner-vlan"] = matchNot.InnerVlan
    leafs["flow-tag"] = matchNot.FlowTag
    leafs["ethertype"] = matchNot.Ethertype
    leafs["destination-port"] = matchNot.DestinationPort
    leafs["fragment-type"] = matchNot.FragmentType
    leafs["frame-relay-dlci"] = matchNot.FrameRelayDlci
    leafs["fr-de"] = matchNot.FrDe
    leafs["icmpv4-code"] = matchNot.Icmpv4Code
    leafs["icmpv4-type"] = matchNot.Icmpv4Type
    leafs["icmpv6-code"] = matchNot.Icmpv6Code
    leafs["icmpv6-type"] = matchNot.Icmpv6Type
    leafs["source-port"] = matchNot.SourcePort
    leafs["tcp-flag"] = matchNot.TcpFlag
    leafs["authen-status"] = matchNot.AuthenStatus
    leafs["circuit-id"] = matchNot.CircuitId
    leafs["circuit-id-regex"] = matchNot.CircuitIdRegex
    leafs["remote-id"] = matchNot.RemoteId
    leafs["remote-id-regex"] = matchNot.RemoteIdRegex
    leafs["service-name"] = matchNot.ServiceName
    leafs["service-name-regex"] = matchNot.ServiceNameRegex
    leafs["timer"] = matchNot.Timer
    leafs["timer-regex"] = matchNot.TimerRegex
    leafs["user-name"] = matchNot.UserName
    leafs["user-name-regex"] = matchNot.UserNameRegex
    leafs["source-mac"] = matchNot.SourceMac
    leafs["destination-mac"] = matchNot.DestinationMac
    leafs["vpls-control"] = matchNot.VplsControl
    leafs["vpls-broadcast"] = matchNot.VplsBroadcast
    leafs["vpls-multicast"] = matchNot.VplsMulticast
    leafs["vpls-known"] = matchNot.VplsKnown
    leafs["vpls-unknown"] = matchNot.VplsUnknown
    return leafs
}

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetBundleName() string { return "cisco_ios_xr" }

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetYangName() string { return "match-not" }

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) SetParent(parent types.Entity) { matchNot.parent = parent }

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetParent() types.Entity { return matchNot.parent }

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetParentYangName() string { return "class-map" }

// PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4
// Match destination IPv4 address.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv4 netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetFilter() yfilter.YFilter { return destinationAddressIpv4.YFilter }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) SetFilter(yf yfilter.YFilter) { destinationAddressIpv4.YFilter = yf }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "netmask" { return "Netmask" }
    return ""
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetSegmentPath() string {
    return "destination-address-ipv4" + "[address='" + fmt.Sprintf("%v", destinationAddressIpv4.Address) + "']" + "[netmask='" + fmt.Sprintf("%v", destinationAddressIpv4.Netmask) + "']"
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = destinationAddressIpv4.Address
    leafs["netmask"] = destinationAddressIpv4.Netmask
    return leafs
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetYangName() string { return "destination-address-ipv4" }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) SetParent(parent types.Entity) { destinationAddressIpv4.parent = parent }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetParent() types.Entity { return destinationAddressIpv4.parent }

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetParentYangName() string { return "match-not" }

// PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6
// Match destination IPv6 address.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetFilter() yfilter.YFilter { return destinationAddressIpv6.YFilter }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) SetFilter(yf yfilter.YFilter) { destinationAddressIpv6.YFilter = yf }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetSegmentPath() string {
    return "destination-address-ipv6" + "[address='" + fmt.Sprintf("%v", destinationAddressIpv6.Address) + "']" + "[prefix-length='" + fmt.Sprintf("%v", destinationAddressIpv6.PrefixLength) + "']"
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = destinationAddressIpv6.Address
    leafs["prefix-length"] = destinationAddressIpv6.PrefixLength
    return leafs
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetYangName() string { return "destination-address-ipv6" }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) SetParent(parent types.Entity) { destinationAddressIpv6.parent = parent }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetParent() types.Entity { return destinationAddressIpv6.parent }

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetParentYangName() string { return "match-not" }

// PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4
// Match source IPv4 address.
type PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv4 netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetFilter() yfilter.YFilter { return sourceAddressIpv4.YFilter }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) SetFilter(yf yfilter.YFilter) { sourceAddressIpv4.YFilter = yf }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "netmask" { return "Netmask" }
    return ""
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetSegmentPath() string {
    return "source-address-ipv4" + "[address='" + fmt.Sprintf("%v", sourceAddressIpv4.Address) + "']" + "[netmask='" + fmt.Sprintf("%v", sourceAddressIpv4.Netmask) + "']"
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = sourceAddressIpv4.Address
    leafs["netmask"] = sourceAddressIpv4.Netmask
    return leafs
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetBundleName() string { return "cisco_ios_xr" }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetYangName() string { return "source-address-ipv4" }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) SetParent(parent types.Entity) { sourceAddressIpv4.parent = parent }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetParent() types.Entity { return sourceAddressIpv4.parent }

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetParentYangName() string { return "match-not" }

// PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6
// Match source IPv6 address.
type PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetFilter() yfilter.YFilter { return sourceAddressIpv6.YFilter }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) SetFilter(yf yfilter.YFilter) { sourceAddressIpv6.YFilter = yf }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetSegmentPath() string {
    return "source-address-ipv6" + "[address='" + fmt.Sprintf("%v", sourceAddressIpv6.Address) + "']" + "[prefix-length='" + fmt.Sprintf("%v", sourceAddressIpv6.PrefixLength) + "']"
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = sourceAddressIpv6.Address
    leafs["prefix-length"] = sourceAddressIpv6.PrefixLength
    return leafs
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetBundleName() string { return "cisco_ios_xr" }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetYangName() string { return "source-address-ipv6" }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) SetParent(parent types.Entity) { sourceAddressIpv6.parent = parent }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetParent() types.Entity { return sourceAddressIpv6.parent }

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetParentYangName() string { return "match-not" }

// PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId
// Match dhcp client ID.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Dhcp client Id. The type is string with length:
    // 1..32.
    Value interface{}

    // This attribute is a key. Dhcp client id Ascii/Hex. The type is string with
    // pattern: (none)|(ascii)|(hex).
    Flag interface{}
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetFilter() yfilter.YFilter { return dhcpClientId.YFilter }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) SetFilter(yf yfilter.YFilter) { dhcpClientId.YFilter = yf }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "flag" { return "Flag" }
    return ""
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetSegmentPath() string {
    return "dhcp-client-id" + "[value='" + fmt.Sprintf("%v", dhcpClientId.Value) + "']" + "[flag='" + fmt.Sprintf("%v", dhcpClientId.Flag) + "']"
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = dhcpClientId.Value
    leafs["flag"] = dhcpClientId.Flag
    return leafs
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetYangName() string { return "dhcp-client-id" }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) SetParent(parent types.Entity) { dhcpClientId.parent = parent }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetParent() types.Entity { return dhcpClientId.parent }

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetParentYangName() string { return "match-not" }

// PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex
// Match dhcp client id regex.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Dhcp client id regular expression. The type is
    // string with length: 1..32.
    Value interface{}

    // This attribute is a key. Dhcp client Id regex Ascii/Hex. The type is string
    // with pattern: (none)|(ascii)|(hex).
    Flag interface{}
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetFilter() yfilter.YFilter { return dhcpClientIdRegex.YFilter }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) SetFilter(yf yfilter.YFilter) { dhcpClientIdRegex.YFilter = yf }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "flag" { return "Flag" }
    return ""
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetSegmentPath() string {
    return "dhcp-client-id-regex" + "[value='" + fmt.Sprintf("%v", dhcpClientIdRegex.Value) + "']" + "[flag='" + fmt.Sprintf("%v", dhcpClientIdRegex.Flag) + "']"
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = dhcpClientIdRegex.Value
    leafs["flag"] = dhcpClientIdRegex.Flag
    return leafs
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetYangName() string { return "dhcp-client-id-regex" }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) SetParent(parent types.Entity) { dhcpClientIdRegex.parent = parent }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetParent() types.Entity { return dhcpClientIdRegex.parent }

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetParentYangName() string { return "match-not" }

// PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName
// Match domain name.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Domain name or regular expression. The type is
    // string with length: 1..32.
    Name interface{}

    // This attribute is a key. Domain-format name. The type is string with
    // length: 1..32.
    Format interface{}
}

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetFilter() yfilter.YFilter { return domainName.YFilter }

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) SetFilter(yf yfilter.YFilter) { domainName.YFilter = yf }

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "format" { return "Format" }
    return ""
}

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetSegmentPath() string {
    return "domain-name" + "[name='" + fmt.Sprintf("%v", domainName.Name) + "']" + "[format='" + fmt.Sprintf("%v", domainName.Format) + "']"
}

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = domainName.Name
    leafs["format"] = domainName.Format
    return leafs
}

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetBundleName() string { return "cisco_ios_xr" }

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetYangName() string { return "domain-name" }

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) SetParent(parent types.Entity) { domainName.parent = parent }

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetParent() types.Entity { return domainName.parent }

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetParentYangName() string { return "match-not" }

// PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex
// Match domain name.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Domain name or regular expression. The type is
    // string with length: 1..32.
    Regex interface{}

    // This attribute is a key. Domain-format name. The type is string with
    // length: 1..32.
    Format interface{}
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetFilter() yfilter.YFilter { return domainNameRegex.YFilter }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) SetFilter(yf yfilter.YFilter) { domainNameRegex.YFilter = yf }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetGoName(yname string) string {
    if yname == "regex" { return "Regex" }
    if yname == "format" { return "Format" }
    return ""
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetSegmentPath() string {
    return "domain-name-regex" + "[regex='" + fmt.Sprintf("%v", domainNameRegex.Regex) + "']" + "[format='" + fmt.Sprintf("%v", domainNameRegex.Format) + "']"
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["regex"] = domainNameRegex.Regex
    leafs["format"] = domainNameRegex.Format
    return leafs
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetBundleName() string { return "cisco_ios_xr" }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetYangName() string { return "domain-name-regex" }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) SetParent(parent types.Entity) { domainNameRegex.parent = parent }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetParent() types.Entity { return domainNameRegex.parent }

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetParentYangName() string { return "match-not" }

// PolicyManager_ClassMaps_ClassMap_MatchNot_Flow
// Match flow.
type PolicyManager_ClassMaps_ClassMap_MatchNot_Flow struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the flow-tag parameters. The type is slice of interface{} with
    // range: 1..63.
    FlowTag []interface{}
}

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetFilter() yfilter.YFilter { return flow.YFilter }

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) SetFilter(yf yfilter.YFilter) { flow.YFilter = yf }

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetGoName(yname string) string {
    if yname == "flow-tag" { return "FlowTag" }
    return ""
}

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetSegmentPath() string {
    return "flow"
}

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-tag"] = flow.FlowTag
    return leafs
}

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetBundleName() string { return "cisco_ios_xr" }

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetYangName() string { return "flow" }

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) SetParent(parent types.Entity) { flow.parent = parent }

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetParent() types.Entity { return flow.parent }

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetParentYangName() string { return "match-not" }

// PolicyManager_PolicyMaps
// Policy-maps configuration.
type PolicyManager_PolicyMaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy-map configuration. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap.
    PolicyMap []PolicyManager_PolicyMaps_PolicyMap
}

func (policyMaps *PolicyManager_PolicyMaps) GetFilter() yfilter.YFilter { return policyMaps.YFilter }

func (policyMaps *PolicyManager_PolicyMaps) SetFilter(yf yfilter.YFilter) { policyMaps.YFilter = yf }

func (policyMaps *PolicyManager_PolicyMaps) GetGoName(yname string) string {
    if yname == "policy-map" { return "PolicyMap" }
    return ""
}

func (policyMaps *PolicyManager_PolicyMaps) GetSegmentPath() string {
    return "policy-maps"
}

func (policyMaps *PolicyManager_PolicyMaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-map" {
        for _, c := range policyMaps.PolicyMap {
            if policyMaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_PolicyMaps_PolicyMap{}
        policyMaps.PolicyMap = append(policyMaps.PolicyMap, child)
        return &policyMaps.PolicyMap[len(policyMaps.PolicyMap)-1]
    }
    return nil
}

func (policyMaps *PolicyManager_PolicyMaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policyMaps.PolicyMap {
        children[policyMaps.PolicyMap[i].GetSegmentPath()] = &policyMaps.PolicyMap[i]
    }
    return children
}

func (policyMaps *PolicyManager_PolicyMaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policyMaps *PolicyManager_PolicyMaps) GetBundleName() string { return "cisco_ios_xr" }

func (policyMaps *PolicyManager_PolicyMaps) GetYangName() string { return "policy-maps" }

func (policyMaps *PolicyManager_PolicyMaps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyMaps *PolicyManager_PolicyMaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyMaps *PolicyManager_PolicyMaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyMaps *PolicyManager_PolicyMaps) SetParent(parent types.Entity) { policyMaps.parent = parent }

func (policyMaps *PolicyManager_PolicyMaps) GetParent() types.Entity { return policyMaps.parent }

func (policyMaps *PolicyManager_PolicyMaps) GetParentYangName() string { return "policy-manager" }

// PolicyManager_PolicyMaps_PolicyMap
// Policy-map configuration.
type PolicyManager_PolicyMaps_PolicyMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Type of policy-map. The type is PolicyMapType.
    Type interface{}

    // This attribute is a key. Name of policy-map. The type is string with
    // pattern: [a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}.
    Name interface{}

    // Description for this policy-map. The type is string.
    Description interface{}

    // Policy event. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap_Event.
    Event []PolicyManager_PolicyMaps_PolicyMap_Event

    // Class-map rule. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule.
    PolicyMapRule []PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule
}

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetFilter() yfilter.YFilter { return policyMap.YFilter }

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) SetFilter(yf yfilter.YFilter) { policyMap.YFilter = yf }

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "name" { return "Name" }
    if yname == "description" { return "Description" }
    if yname == "event" { return "Event" }
    if yname == "policy-map-rule" { return "PolicyMapRule" }
    return ""
}

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetSegmentPath() string {
    return "policy-map" + "[type='" + fmt.Sprintf("%v", policyMap.Type) + "']" + "[name='" + fmt.Sprintf("%v", policyMap.Name) + "']"
}

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "event" {
        for _, c := range policyMap.Event {
            if policyMap.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_PolicyMaps_PolicyMap_Event{}
        policyMap.Event = append(policyMap.Event, child)
        return &policyMap.Event[len(policyMap.Event)-1]
    }
    if childYangName == "policy-map-rule" {
        for _, c := range policyMap.PolicyMapRule {
            if policyMap.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule{}
        policyMap.PolicyMapRule = append(policyMap.PolicyMapRule, child)
        return &policyMap.PolicyMapRule[len(policyMap.PolicyMapRule)-1]
    }
    return nil
}

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policyMap.Event {
        children[policyMap.Event[i].GetSegmentPath()] = &policyMap.Event[i]
    }
    for i := range policyMap.PolicyMapRule {
        children[policyMap.PolicyMapRule[i].GetSegmentPath()] = &policyMap.PolicyMapRule[i]
    }
    return children
}

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = policyMap.Type
    leafs["name"] = policyMap.Name
    leafs["description"] = policyMap.Description
    return leafs
}

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetBundleName() string { return "cisco_ios_xr" }

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetYangName() string { return "policy-map" }

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) SetParent(parent types.Entity) { policyMap.parent = parent }

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetParent() types.Entity { return policyMap.parent }

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetParentYangName() string { return "policy-maps" }

// PolicyManager_PolicyMaps_PolicyMap_Event
// Policy event.
type PolicyManager_PolicyMaps_PolicyMap_Event struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Event type. The type is EventType.
    EventType interface{}

    // Execute all the matched classes. The type is interface{}.
    EventModeMatchAll interface{}

    // Execute only the first matched class. The type is interface{}.
    EventModeMatchFirst interface{}

    // Class-map rule. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap_Event_Class.
    Class []PolicyManager_PolicyMaps_PolicyMap_Event_Class
}

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetFilter() yfilter.YFilter { return event.YFilter }

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) SetFilter(yf yfilter.YFilter) { event.YFilter = yf }

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetGoName(yname string) string {
    if yname == "event-type" { return "EventType" }
    if yname == "event-mode-match-all" { return "EventModeMatchAll" }
    if yname == "event-mode-match-first" { return "EventModeMatchFirst" }
    if yname == "class" { return "Class" }
    return ""
}

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetSegmentPath() string {
    return "event" + "[event-type='" + fmt.Sprintf("%v", event.EventType) + "']"
}

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range event.Class {
            if event.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_PolicyMaps_PolicyMap_Event_Class{}
        event.Class = append(event.Class, child)
        return &event.Class[len(event.Class)-1]
    }
    return nil
}

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range event.Class {
        children[event.Class[i].GetSegmentPath()] = &event.Class[i]
    }
    return children
}

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-type"] = event.EventType
    leafs["event-mode-match-all"] = event.EventModeMatchAll
    leafs["event-mode-match-first"] = event.EventModeMatchFirst
    return leafs
}

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetBundleName() string { return "cisco_ios_xr" }

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetYangName() string { return "event" }

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) SetParent(parent types.Entity) { event.parent = parent }

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetParent() types.Entity { return event.parent }

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetParentYangName() string { return "policy-map" }

// PolicyManager_PolicyMaps_PolicyMap_Event_Class
// Class-map rule.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of class. The type is string with pattern:
    // [a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}.
    ClassName interface{}

    // This attribute is a key. Type of class. The type is PmapClassMapType.
    ClassType interface{}

    // Class execution strategy. The type is ExecutionStrategy.
    ClassExecutionStrategy interface{}

    // Action rule. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule.
    ActionRule []PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule
}

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "class-type" { return "ClassType" }
    if yname == "class-execution-strategy" { return "ClassExecutionStrategy" }
    if yname == "action-rule" { return "ActionRule" }
    return ""
}

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetSegmentPath() string {
    return "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']" + "[class-type='" + fmt.Sprintf("%v", class.ClassType) + "']"
}

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "action-rule" {
        for _, c := range class.ActionRule {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule{}
        class.ActionRule = append(class.ActionRule, child)
        return &class.ActionRule[len(class.ActionRule)-1]
    }
    return nil
}

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range class.ActionRule {
        children[class.ActionRule[i].GetSegmentPath()] = &class.ActionRule[i]
    }
    return children
}

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = class.ClassName
    leafs["class-type"] = class.ClassType
    leafs["class-execution-strategy"] = class.ClassExecutionStrategy
    return leafs
}

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetYangName() string { return "class" }

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetParent() types.Entity { return class.parent }

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetParentYangName() string { return "event" }

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule
// Action rule.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence number for this action. The type is
    // interface{} with range: 1..65535.
    ActionSequenceNumber interface{}

    // Disconnect session. The type is interface{}.
    Disconnect interface{}

    // Monitor session. The type is interface{}.
    Monitor interface{}

    // Activate dynamic templates.
    ActivateDynamicTemplate PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate

    // Authentication related configuration.
    Authenticate PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate

    // Authorize.
    Authorize PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize

    // Deactivate dynamic templates.
    DeactivateDynamicTemplate PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate

    // Set a timer to execute a rule on its  expiry.
    SetTimer PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer

    // Disable timer before it expires.
    StopTimer PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer
}

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetFilter() yfilter.YFilter { return actionRule.YFilter }

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) SetFilter(yf yfilter.YFilter) { actionRule.YFilter = yf }

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetGoName(yname string) string {
    if yname == "action-sequence-number" { return "ActionSequenceNumber" }
    if yname == "disconnect" { return "Disconnect" }
    if yname == "monitor" { return "Monitor" }
    if yname == "activate-dynamic-template" { return "ActivateDynamicTemplate" }
    if yname == "authenticate" { return "Authenticate" }
    if yname == "authorize" { return "Authorize" }
    if yname == "deactivate-dynamic-template" { return "DeactivateDynamicTemplate" }
    if yname == "set-timer" { return "SetTimer" }
    if yname == "stop-timer" { return "StopTimer" }
    return ""
}

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetSegmentPath() string {
    return "action-rule" + "[action-sequence-number='" + fmt.Sprintf("%v", actionRule.ActionSequenceNumber) + "']"
}

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "activate-dynamic-template" {
        return &actionRule.ActivateDynamicTemplate
    }
    if childYangName == "authenticate" {
        return &actionRule.Authenticate
    }
    if childYangName == "authorize" {
        return &actionRule.Authorize
    }
    if childYangName == "deactivate-dynamic-template" {
        return &actionRule.DeactivateDynamicTemplate
    }
    if childYangName == "set-timer" {
        return &actionRule.SetTimer
    }
    if childYangName == "stop-timer" {
        return &actionRule.StopTimer
    }
    return nil
}

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["activate-dynamic-template"] = &actionRule.ActivateDynamicTemplate
    children["authenticate"] = &actionRule.Authenticate
    children["authorize"] = &actionRule.Authorize
    children["deactivate-dynamic-template"] = &actionRule.DeactivateDynamicTemplate
    children["set-timer"] = &actionRule.SetTimer
    children["stop-timer"] = &actionRule.StopTimer
    return children
}

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-sequence-number"] = actionRule.ActionSequenceNumber
    leafs["disconnect"] = actionRule.Disconnect
    leafs["monitor"] = actionRule.Monitor
    return leafs
}

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetBundleName() string { return "cisco_ios_xr" }

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetYangName() string { return "action-rule" }

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) SetParent(parent types.Entity) { actionRule.parent = parent }

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetParent() types.Entity { return actionRule.parent }

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetParentYangName() string { return "class" }

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate
// Activate dynamic templates.
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dynamic template name. The type is string. This attribute is mandatory.
    Name interface{}

    // Name of the AAA method list. The type is string.
    AaaList interface{}
}

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetFilter() yfilter.YFilter { return activateDynamicTemplate.YFilter }

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) SetFilter(yf yfilter.YFilter) { activateDynamicTemplate.YFilter = yf }

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "aaa-list" { return "AaaList" }
    return ""
}

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetSegmentPath() string {
    return "activate-dynamic-template"
}

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = activateDynamicTemplate.Name
    leafs["aaa-list"] = activateDynamicTemplate.AaaList
    return leafs
}

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetYangName() string { return "activate-dynamic-template" }

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) SetParent(parent types.Entity) { activateDynamicTemplate.parent = parent }

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetParent() types.Entity { return activateDynamicTemplate.parent }

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetParentYangName() string { return "action-rule" }

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate
// Authentication related configuration.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the AAA method list. The type is string.
    AaaList interface{}
}

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetFilter() yfilter.YFilter { return authenticate.YFilter }

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) SetFilter(yf yfilter.YFilter) { authenticate.YFilter = yf }

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetGoName(yname string) string {
    if yname == "aaa-list" { return "AaaList" }
    return ""
}

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetSegmentPath() string {
    return "authenticate"
}

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aaa-list"] = authenticate.AaaList
    return leafs
}

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetBundleName() string { return "cisco_ios_xr" }

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetYangName() string { return "authenticate" }

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) SetParent(parent types.Entity) { authenticate.parent = parent }

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetParent() types.Entity { return authenticate.parent }

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetParentYangName() string { return "action-rule" }

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize
// Authorize.
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the AAA method list. The type is string. This attribute is
    // mandatory.
    AaaList interface{}

    // Specify an Authorize format name. The type is string.
    Format interface{}

    // Specify an Authorize format name. The type is AuthorizeIdentifier.
    Identifier interface{}

    // Specify a password to be used for AAA request. The type is string. This
    // attribute is mandatory.
    Password interface{}
}

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetFilter() yfilter.YFilter { return authorize.YFilter }

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) SetFilter(yf yfilter.YFilter) { authorize.YFilter = yf }

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetGoName(yname string) string {
    if yname == "aaa-list" { return "AaaList" }
    if yname == "format" { return "Format" }
    if yname == "identifier" { return "Identifier" }
    if yname == "password" { return "Password" }
    return ""
}

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetSegmentPath() string {
    return "authorize"
}

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aaa-list"] = authorize.AaaList
    leafs["format"] = authorize.Format
    leafs["identifier"] = authorize.Identifier
    leafs["password"] = authorize.Password
    return leafs
}

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetBundleName() string { return "cisco_ios_xr" }

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetYangName() string { return "authorize" }

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) SetParent(parent types.Entity) { authorize.parent = parent }

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetParent() types.Entity { return authorize.parent }

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetParentYangName() string { return "action-rule" }

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate
// Deactivate dynamic templates.
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dynamic template name. The type is string. This attribute is mandatory.
    Name interface{}

    // Name of the AAA method list. The type is string.
    AaaList interface{}
}

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetFilter() yfilter.YFilter { return deactivateDynamicTemplate.YFilter }

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) SetFilter(yf yfilter.YFilter) { deactivateDynamicTemplate.YFilter = yf }

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "aaa-list" { return "AaaList" }
    return ""
}

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetSegmentPath() string {
    return "deactivate-dynamic-template"
}

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = deactivateDynamicTemplate.Name
    leafs["aaa-list"] = deactivateDynamicTemplate.AaaList
    return leafs
}

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetBundleName() string { return "cisco_ios_xr" }

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetYangName() string { return "deactivate-dynamic-template" }

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) SetParent(parent types.Entity) { deactivateDynamicTemplate.parent = parent }

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetParent() types.Entity { return deactivateDynamicTemplate.parent }

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetParentYangName() string { return "action-rule" }

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer
// Set a timer to execute a rule on its 
// expiry
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the timer. The type is string. This attribute is mandatory.
    TimerName interface{}

    // Timer value in minutes. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory. Units are minutes.
    TimerValue interface{}
}

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetFilter() yfilter.YFilter { return setTimer.YFilter }

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) SetFilter(yf yfilter.YFilter) { setTimer.YFilter = yf }

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetGoName(yname string) string {
    if yname == "timer-name" { return "TimerName" }
    if yname == "timer-value" { return "TimerValue" }
    return ""
}

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetSegmentPath() string {
    return "set-timer"
}

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timer-name"] = setTimer.TimerName
    leafs["timer-value"] = setTimer.TimerValue
    return leafs
}

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetBundleName() string { return "cisco_ios_xr" }

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetYangName() string { return "set-timer" }

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) SetParent(parent types.Entity) { setTimer.parent = parent }

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetParent() types.Entity { return setTimer.parent }

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetParentYangName() string { return "action-rule" }

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer
// Disable timer before it expires.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the timer. The type is string.
    TimerName interface{}
}

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetFilter() yfilter.YFilter { return stopTimer.YFilter }

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) SetFilter(yf yfilter.YFilter) { stopTimer.YFilter = yf }

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetGoName(yname string) string {
    if yname == "timer-name" { return "TimerName" }
    return ""
}

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetSegmentPath() string {
    return "stop-timer"
}

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timer-name"] = stopTimer.TimerName
    return leafs
}

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetBundleName() string { return "cisco_ios_xr" }

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetYangName() string { return "stop-timer" }

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) SetParent(parent types.Entity) { stopTimer.parent = parent }

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetParent() types.Entity { return stopTimer.parent }

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetParentYangName() string { return "action-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule
// Class-map rule.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of class-map. The type is string with
    // pattern: [a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}.
    ClassName interface{}

    // This attribute is a key. Type of class-map. The type is PmapClassMapType.
    ClassType interface{}

    // Priority level. The type is interface{} with range: 1..7.
    PriorityLevel interface{}

    // Default random early detection. The type is interface{}.
    DefaultRed interface{}

    // ECN based random early detection. The type is interface{}.
    EcnRed interface{}

    // Policy action http redirect. Redirect to this url. The type is string.
    HttpRedirect interface{}

    // Policy action PBR transmit. The type is interface{}.
    PbrTransmit interface{}

    // Policy action PBR drop. The type is interface{}.
    PbrDrop interface{}

    // Policy action DECAP GRE. The type is interface{}.
    DecapGre interface{}

    // Policy action service fragment.  Service fragment name. The type is string.
    ServiceFragment interface{}

    // Policy action fragment. Fragment name. The type is string.
    Fragment interface{}

    // Policy action shape.
    Shape PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape

    // Policy action minimum bandwidth queue.
    MinBandwidth PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth

    // Policy action bandwidth remaining queue.
    BandwidthRemaining PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining

    // Policy action queue limit.
    QueueLimit PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit

    // Policy action pfc.
    Pfc PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc

    // Random early detection. All RED profiles in a class must be based on the
    // same field. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect.
    RandomDetect []PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect

    // Policy action packet marking.
    Set PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set

    // Configures traffic policing action.
    Police PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police

    // Configure a child service policy.
    ServicePolicy PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy

    // Policy action CAC.
    CacLocal PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal

    // Policy flow monitoring action.
    FlowParams PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams

    // Policy IP-CBR metric action.
    MetricsIpcbr PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr

    // Policy action react.
    React PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React

    // Policy action redirect.
    PbrRedirect PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect

    // Policy action PBR forward.
    PbrForward PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward

    // Policy action service function path.
    ServiceFunctionPath PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath
}

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetFilter() yfilter.YFilter { return policyMapRule.YFilter }

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) SetFilter(yf yfilter.YFilter) { policyMapRule.YFilter = yf }

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "class-type" { return "ClassType" }
    if yname == "priority-level" { return "PriorityLevel" }
    if yname == "default-red" { return "DefaultRed" }
    if yname == "ecn-red" { return "EcnRed" }
    if yname == "http-redirect" { return "HttpRedirect" }
    if yname == "pbr-transmit" { return "PbrTransmit" }
    if yname == "pbr-drop" { return "PbrDrop" }
    if yname == "decap-gre" { return "DecapGre" }
    if yname == "service-fragment" { return "ServiceFragment" }
    if yname == "fragment" { return "Fragment" }
    if yname == "shape" { return "Shape" }
    if yname == "min-bandwidth" { return "MinBandwidth" }
    if yname == "bandwidth-remaining" { return "BandwidthRemaining" }
    if yname == "queue-limit" { return "QueueLimit" }
    if yname == "pfc" { return "Pfc" }
    if yname == "random-detect" { return "RandomDetect" }
    if yname == "set" { return "Set" }
    if yname == "police" { return "Police" }
    if yname == "service-policy" { return "ServicePolicy" }
    if yname == "cac-local" { return "CacLocal" }
    if yname == "flow-params" { return "FlowParams" }
    if yname == "metrics-ipcbr" { return "MetricsIpcbr" }
    if yname == "react" { return "React" }
    if yname == "pbr-redirect" { return "PbrRedirect" }
    if yname == "pbr-forward" { return "PbrForward" }
    if yname == "service-function-path" { return "ServiceFunctionPath" }
    return ""
}

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetSegmentPath() string {
    return "policy-map-rule" + "[class-name='" + fmt.Sprintf("%v", policyMapRule.ClassName) + "']" + "[class-type='" + fmt.Sprintf("%v", policyMapRule.ClassType) + "']"
}

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "shape" {
        return &policyMapRule.Shape
    }
    if childYangName == "min-bandwidth" {
        return &policyMapRule.MinBandwidth
    }
    if childYangName == "bandwidth-remaining" {
        return &policyMapRule.BandwidthRemaining
    }
    if childYangName == "queue-limit" {
        return &policyMapRule.QueueLimit
    }
    if childYangName == "pfc" {
        return &policyMapRule.Pfc
    }
    if childYangName == "random-detect" {
        for _, c := range policyMapRule.RandomDetect {
            if policyMapRule.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect{}
        policyMapRule.RandomDetect = append(policyMapRule.RandomDetect, child)
        return &policyMapRule.RandomDetect[len(policyMapRule.RandomDetect)-1]
    }
    if childYangName == "set" {
        return &policyMapRule.Set
    }
    if childYangName == "police" {
        return &policyMapRule.Police
    }
    if childYangName == "service-policy" {
        return &policyMapRule.ServicePolicy
    }
    if childYangName == "cac-local" {
        return &policyMapRule.CacLocal
    }
    if childYangName == "flow-params" {
        return &policyMapRule.FlowParams
    }
    if childYangName == "metrics-ipcbr" {
        return &policyMapRule.MetricsIpcbr
    }
    if childYangName == "react" {
        return &policyMapRule.React
    }
    if childYangName == "pbr-redirect" {
        return &policyMapRule.PbrRedirect
    }
    if childYangName == "pbr-forward" {
        return &policyMapRule.PbrForward
    }
    if childYangName == "service-function-path" {
        return &policyMapRule.ServiceFunctionPath
    }
    return nil
}

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["shape"] = &policyMapRule.Shape
    children["min-bandwidth"] = &policyMapRule.MinBandwidth
    children["bandwidth-remaining"] = &policyMapRule.BandwidthRemaining
    children["queue-limit"] = &policyMapRule.QueueLimit
    children["pfc"] = &policyMapRule.Pfc
    for i := range policyMapRule.RandomDetect {
        children[policyMapRule.RandomDetect[i].GetSegmentPath()] = &policyMapRule.RandomDetect[i]
    }
    children["set"] = &policyMapRule.Set
    children["police"] = &policyMapRule.Police
    children["service-policy"] = &policyMapRule.ServicePolicy
    children["cac-local"] = &policyMapRule.CacLocal
    children["flow-params"] = &policyMapRule.FlowParams
    children["metrics-ipcbr"] = &policyMapRule.MetricsIpcbr
    children["react"] = &policyMapRule.React
    children["pbr-redirect"] = &policyMapRule.PbrRedirect
    children["pbr-forward"] = &policyMapRule.PbrForward
    children["service-function-path"] = &policyMapRule.ServiceFunctionPath
    return children
}

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = policyMapRule.ClassName
    leafs["class-type"] = policyMapRule.ClassType
    leafs["priority-level"] = policyMapRule.PriorityLevel
    leafs["default-red"] = policyMapRule.DefaultRed
    leafs["ecn-red"] = policyMapRule.EcnRed
    leafs["http-redirect"] = policyMapRule.HttpRedirect
    leafs["pbr-transmit"] = policyMapRule.PbrTransmit
    leafs["pbr-drop"] = policyMapRule.PbrDrop
    leafs["decap-gre"] = policyMapRule.DecapGre
    leafs["service-fragment"] = policyMapRule.ServiceFragment
    leafs["fragment"] = policyMapRule.Fragment
    return leafs
}

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetBundleName() string { return "cisco_ios_xr" }

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetYangName() string { return "policy-map-rule" }

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) SetParent(parent types.Entity) { policyMapRule.parent = parent }

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetParent() types.Entity { return policyMapRule.parent }

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetParentYangName() string { return "policy-map" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape
// Policy action shape.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rate configuration.
    Rate PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate

    // Burst size configuration.
    Burst PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst
}

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetFilter() yfilter.YFilter { return shape.YFilter }

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) SetFilter(yf yfilter.YFilter) { shape.YFilter = yf }

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetGoName(yname string) string {
    if yname == "rate" { return "Rate" }
    if yname == "burst" { return "Burst" }
    return ""
}

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetSegmentPath() string {
    return "shape"
}

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rate" {
        return &shape.Rate
    }
    if childYangName == "burst" {
        return &shape.Burst
    }
    return nil
}

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rate"] = &shape.Rate
    children["burst"] = &shape.Burst
    return children
}

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetBundleName() string { return "cisco_ios_xr" }

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetYangName() string { return "shape" }

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) SetParent(parent types.Entity) { shape.parent = parent }

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetParent() types.Entity { return shape.parent }

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate
// Rate configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Shape bandwidth value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Shape bandwidth units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(percent)|(per-million)|(per-thousand).
    Unit interface{}
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetFilter() yfilter.YFilter { return rate.YFilter }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) SetFilter(yf yfilter.YFilter) { rate.YFilter = yf }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetSegmentPath() string {
    return "rate"
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = rate.Value
    leafs["unit"] = rate.Unit
    return leafs
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetBundleName() string { return "cisco_ios_xr" }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetYangName() string { return "rate" }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) SetParent(parent types.Entity) { rate.parent = parent }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetParent() types.Entity { return rate.parent }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetParentYangName() string { return "shape" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst
// Burst size configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Burst size value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Burst size units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Units interface{}
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetFilter() yfilter.YFilter { return burst.YFilter }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) SetFilter(yf yfilter.YFilter) { burst.YFilter = yf }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "units" { return "Units" }
    return ""
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetSegmentPath() string {
    return "burst"
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = burst.Value
    leafs["units"] = burst.Units
    return leafs
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetBundleName() string { return "cisco_ios_xr" }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetYangName() string { return "burst" }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) SetParent(parent types.Entity) { burst.parent = parent }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetParent() types.Entity { return burst.parent }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetParentYangName() string { return "shape" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth
// Policy action minimum bandwidth queue.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum bandwidth value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Minimum bandwidth units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(percent)|(per-million)|(per-thousand).
    Unit interface{}
}

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetFilter() yfilter.YFilter { return minBandwidth.YFilter }

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) SetFilter(yf yfilter.YFilter) { minBandwidth.YFilter = yf }

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetSegmentPath() string {
    return "min-bandwidth"
}

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = minBandwidth.Value
    leafs["unit"] = minBandwidth.Unit
    return leafs
}

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetYangName() string { return "min-bandwidth" }

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) SetParent(parent types.Entity) { minBandwidth.parent = parent }

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetParent() types.Entity { return minBandwidth.parent }

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining
// Policy action bandwidth remaining queue.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Remaining bandwidth value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Remaining bandwidth units. The type is string with pattern:
    // (percent)|(ratio).
    Unit interface{}
}

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetFilter() yfilter.YFilter { return bandwidthRemaining.YFilter }

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) SetFilter(yf yfilter.YFilter) { bandwidthRemaining.YFilter = yf }

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetSegmentPath() string {
    return "bandwidth-remaining"
}

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bandwidthRemaining.Value
    leafs["unit"] = bandwidthRemaining.Unit
    return leafs
}

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetYangName() string { return "bandwidth-remaining" }

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) SetParent(parent types.Entity) { bandwidthRemaining.parent = parent }

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetParent() types.Entity { return bandwidthRemaining.parent }

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit
// Policy action queue limit.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Remaining bandwidth value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Remaining bandwidth units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells)|(percent).
    Unit interface{}
}

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetFilter() yfilter.YFilter { return queueLimit.YFilter }

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) SetFilter(yf yfilter.YFilter) { queueLimit.YFilter = yf }

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetSegmentPath() string {
    return "queue-limit"
}

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = queueLimit.Value
    leafs["unit"] = queueLimit.Unit
    return leafs
}

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetBundleName() string { return "cisco_ios_xr" }

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetYangName() string { return "queue-limit" }

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) SetParent(parent types.Entity) { queueLimit.parent = parent }

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetParent() types.Entity { return queueLimit.parent }

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc
// Policy action pfc.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pfc Pause set value. The type is interface{}.
    PfcPauseSet interface{}

    
    PfcBufferSize PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize

    
    PfcPauseThreshold PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold

    
    PfcResumeThreshold PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold
}

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetFilter() yfilter.YFilter { return pfc.YFilter }

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) SetFilter(yf yfilter.YFilter) { pfc.YFilter = yf }

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetGoName(yname string) string {
    if yname == "pfc-pause-set" { return "PfcPauseSet" }
    if yname == "pfc-buffer-size" { return "PfcBufferSize" }
    if yname == "pfc-pause-threshold" { return "PfcPauseThreshold" }
    if yname == "pfc-resume-threshold" { return "PfcResumeThreshold" }
    return ""
}

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetSegmentPath() string {
    return "pfc"
}

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pfc-buffer-size" {
        return &pfc.PfcBufferSize
    }
    if childYangName == "pfc-pause-threshold" {
        return &pfc.PfcPauseThreshold
    }
    if childYangName == "pfc-resume-threshold" {
        return &pfc.PfcResumeThreshold
    }
    return nil
}

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pfc-buffer-size"] = &pfc.PfcBufferSize
    children["pfc-pause-threshold"] = &pfc.PfcPauseThreshold
    children["pfc-resume-threshold"] = &pfc.PfcResumeThreshold
    return children
}

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pfc-pause-set"] = pfc.PfcPauseSet
    return leafs
}

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetBundleName() string { return "cisco_ios_xr" }

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetYangName() string { return "pfc" }

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) SetParent(parent types.Entity) { pfc.parent = parent }

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetParent() types.Entity { return pfc.parent }

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pfc buffer size value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Pfc buffer size units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Unit interface{}
}

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetFilter() yfilter.YFilter { return pfcBufferSize.YFilter }

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) SetFilter(yf yfilter.YFilter) { pfcBufferSize.YFilter = yf }

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetSegmentPath() string {
    return "pfc-buffer-size"
}

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pfcBufferSize.Value
    leafs["unit"] = pfcBufferSize.Unit
    return leafs
}

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetBundleName() string { return "cisco_ios_xr" }

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetYangName() string { return "pfc-buffer-size" }

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) SetParent(parent types.Entity) { pfcBufferSize.parent = parent }

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetParent() types.Entity { return pfcBufferSize.parent }

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetParentYangName() string { return "pfc" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pfc pause threshold value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Pfc pause threshold units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Unit interface{}
}

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetFilter() yfilter.YFilter { return pfcPauseThreshold.YFilter }

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) SetFilter(yf yfilter.YFilter) { pfcPauseThreshold.YFilter = yf }

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetSegmentPath() string {
    return "pfc-pause-threshold"
}

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pfcPauseThreshold.Value
    leafs["unit"] = pfcPauseThreshold.Unit
    return leafs
}

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetYangName() string { return "pfc-pause-threshold" }

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) SetParent(parent types.Entity) { pfcPauseThreshold.parent = parent }

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetParent() types.Entity { return pfcPauseThreshold.parent }

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetParentYangName() string { return "pfc" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pfc resume threshold value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Pfc resume threshold units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Unit interface{}
}

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetFilter() yfilter.YFilter { return pfcResumeThreshold.YFilter }

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) SetFilter(yf yfilter.YFilter) { pfcResumeThreshold.YFilter = yf }

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetSegmentPath() string {
    return "pfc-resume-threshold"
}

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pfcResumeThreshold.Value
    leafs["unit"] = pfcResumeThreshold.Unit
    return leafs
}

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetYangName() string { return "pfc-resume-threshold" }

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) SetParent(parent types.Entity) { pfcResumeThreshold.parent = parent }

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetParent() types.Entity { return pfcResumeThreshold.parent }

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetParentYangName() string { return "pfc" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect
// Random early detection.
// All RED profiles in a class must be based
// on the same field.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Minimum RED threshold value. The type is
    // interface{} with range: 0..4294967295.
    ThresholdMinValue interface{}

    // This attribute is a key. Minimum RED threshold units. The type is string
    // with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    ThresholdMinUnits interface{}

    // This attribute is a key. Maximum RED threshold value. The type is
    // interface{} with range: 0..4294967295.
    ThresholdMaxValue interface{}

    // This attribute is a key. Maximum RED threshold units. The type is string
    // with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    ThresholdMaxUnits interface{}

    // WRED based on CoS. The type is slice of string with pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(([0-9]|[1-5][0-9]|6[0-3])-([0-9]|[1-5][0-9]|6[0-3]))|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Cos []interface{}

    // WRED based on discard class. The type is slice of interface{} with range:
    // 0..7.
    DiscardClass []interface{}

    // WRED based on DSCP. The type is slice of string with pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(([0-9]|[1-5][0-9]|6[0-3])-([0-9]|[1-5][0-9]|6[0-3]))|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Dscp []interface{}

    // MPLS Experimental value based WRED. The type is slice of interface{} with
    // range: 0..7.
    MplsExp []interface{}

    // WRED based on precedence. The type is one of the following types: slice of
    // int with range: 0..7, or slice of string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Precedence []interface{}

    // DEI based WRED. The type is interface{} with range: 0..1.
    Dei interface{}

    // ECN based WRED. The type is interface{}.
    Ecn interface{}
}

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetFilter() yfilter.YFilter { return randomDetect.YFilter }

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) SetFilter(yf yfilter.YFilter) { randomDetect.YFilter = yf }

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetGoName(yname string) string {
    if yname == "threshold-min-value" { return "ThresholdMinValue" }
    if yname == "threshold-min-units" { return "ThresholdMinUnits" }
    if yname == "threshold-max-value" { return "ThresholdMaxValue" }
    if yname == "threshold-max-units" { return "ThresholdMaxUnits" }
    if yname == "cos" { return "Cos" }
    if yname == "discard-class" { return "DiscardClass" }
    if yname == "dscp" { return "Dscp" }
    if yname == "mpls-exp" { return "MplsExp" }
    if yname == "precedence" { return "Precedence" }
    if yname == "dei" { return "Dei" }
    if yname == "ecn" { return "Ecn" }
    return ""
}

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetSegmentPath() string {
    return "random-detect" + "[threshold-min-value='" + fmt.Sprintf("%v", randomDetect.ThresholdMinValue) + "']" + "[threshold-min-units='" + fmt.Sprintf("%v", randomDetect.ThresholdMinUnits) + "']" + "[threshold-max-value='" + fmt.Sprintf("%v", randomDetect.ThresholdMaxValue) + "']" + "[threshold-max-units='" + fmt.Sprintf("%v", randomDetect.ThresholdMaxUnits) + "']"
}

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-min-value"] = randomDetect.ThresholdMinValue
    leafs["threshold-min-units"] = randomDetect.ThresholdMinUnits
    leafs["threshold-max-value"] = randomDetect.ThresholdMaxValue
    leafs["threshold-max-units"] = randomDetect.ThresholdMaxUnits
    leafs["cos"] = randomDetect.Cos
    leafs["discard-class"] = randomDetect.DiscardClass
    leafs["dscp"] = randomDetect.Dscp
    leafs["mpls-exp"] = randomDetect.MplsExp
    leafs["precedence"] = randomDetect.Precedence
    leafs["dei"] = randomDetect.Dei
    leafs["ecn"] = randomDetect.Ecn
    return leafs
}

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetBundleName() string { return "cisco_ios_xr" }

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetYangName() string { return "random-detect" }

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) SetParent(parent types.Entity) { randomDetect.parent = parent }

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetParent() types.Entity { return randomDetect.parent }

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set
// Policy action packet marking.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Marks a packet by setting the DSCP in the ToS byte. The type is string with
    // pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Dscp interface{}

    // Sets the QoS group identifiers on IPv4 or MPLS packets. The set qos-group
    // is supported only on an ingress policy. The type is interface{} with range:
    // 0..512.
    QosGroup interface{}

    // Sets the Traffic class identifiers on IPv4 or MPLS packets. The type is
    // interface{} with range: 0..63.
    TrafficClass interface{}

    // Sets the discard class on IPv4 or MPLS packets. The discard-class can be
    // used only in service policies  that are attached in the ingress policy. The
    // type is interface{} with range: 0..7.
    DiscardClass interface{}

    // Sets the forward class. The type is interface{} with range: 0..7.
    ForwardClass interface{}

    // Set DF bit. The type is interface{} with range: 0..1.
    Df interface{}

    // Sets the specific IEEE 802.1Q Layer 2 CoS value of an outgoing packet. This
    // command should be used by a router if a user wants to mark a packet that is
    // being sent to a switch.  Switches can leverage Layer 2 header information, 
    // including a CoS value marking. Packets entering an  interface cannot be set
    // with a CoS value. The type is interface{} with range: 0..7.
    Cos interface{}

    // Set inner cos. The type is interface{} with range: 0..7.
    InnerCos interface{}

    // Sets the precedence value in the IP header. The type is one of the
    // following types: int with range: 0..7, or string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Precedence interface{}

    // Sets the precedence tunnel value for ipsec. The type is one of the
    // following types: int with range: 0..7, or string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    PrecedenceTunnel interface{}

    // Sets the experimental value of the MPLS packet top-most labels. The type is
    // interface{} with range: 0..7.
    MplsExperimentalTopMost interface{}

    // Sets the experimental value of the MPLS packet  imposition labels.
    // Imposition can be used only in service policies that  are attached in the
    // ingress policy. The type is interface{} with range: 0..7.
    MplsExperimentalImposition interface{}

    // Sets the spatial reuse protocol priority value of an  outgoing packet. The
    // type is interface{} with range: 0..7.
    SrpPriority interface{}

    // Set FrameRelay DE bit. The type is interface{} with range: 0..1.
    FrDe interface{}

    // Set DEI bit. The type is interface{} with range: 0..1.
    Dei interface{}

    // Set DEI imposition bit. The type is interface{} with range: 0..1.
    DeiImposition interface{}

    // Source IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "discard-class" { return "DiscardClass" }
    if yname == "forward-class" { return "ForwardClass" }
    if yname == "df" { return "Df" }
    if yname == "cos" { return "Cos" }
    if yname == "inner-cos" { return "InnerCos" }
    if yname == "precedence" { return "Precedence" }
    if yname == "precedence-tunnel" { return "PrecedenceTunnel" }
    if yname == "mpls-experimental-top-most" { return "MplsExperimentalTopMost" }
    if yname == "mpls-experimental-imposition" { return "MplsExperimentalImposition" }
    if yname == "srp-priority" { return "SrpPriority" }
    if yname == "fr-de" { return "FrDe" }
    if yname == "dei" { return "Dei" }
    if yname == "dei-imposition" { return "DeiImposition" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    return ""
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetSegmentPath() string {
    return "set"
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = set.Dscp
    leafs["qos-group"] = set.QosGroup
    leafs["traffic-class"] = set.TrafficClass
    leafs["discard-class"] = set.DiscardClass
    leafs["forward-class"] = set.ForwardClass
    leafs["df"] = set.Df
    leafs["cos"] = set.Cos
    leafs["inner-cos"] = set.InnerCos
    leafs["precedence"] = set.Precedence
    leafs["precedence-tunnel"] = set.PrecedenceTunnel
    leafs["mpls-experimental-top-most"] = set.MplsExperimentalTopMost
    leafs["mpls-experimental-imposition"] = set.MplsExperimentalImposition
    leafs["srp-priority"] = set.SrpPriority
    leafs["fr-de"] = set.FrDe
    leafs["dei"] = set.Dei
    leafs["dei-imposition"] = set.DeiImposition
    leafs["source-address"] = set.SourceAddress
    leafs["destination-address"] = set.DestinationAddress
    return leafs
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetYangName() string { return "set" }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetParent() types.Entity { return set.parent }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police
// Configures traffic policing action.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rate configuration.
    Rate PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate

    // Peak rate configuration.
    PeakRate PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate

    // Burst configuration.
    Burst PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst

    // Peak burst configuration.
    PeakBurst PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst

    // Configures the action to take on packets that conform  to the rate limit.
    ConformAction PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction

    // Configures the action to take on packets that exceed  the rate limit.
    ExceedAction PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction

    // Configures the action to take on packets that violate the rate limit.
    ViolateAction PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction
}

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetGoName(yname string) string {
    if yname == "rate" { return "Rate" }
    if yname == "peak-rate" { return "PeakRate" }
    if yname == "burst" { return "Burst" }
    if yname == "peak-burst" { return "PeakBurst" }
    if yname == "conform-action" { return "ConformAction" }
    if yname == "exceed-action" { return "ExceedAction" }
    if yname == "violate-action" { return "ViolateAction" }
    return ""
}

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetSegmentPath() string {
    return "police"
}

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rate" {
        return &police.Rate
    }
    if childYangName == "peak-rate" {
        return &police.PeakRate
    }
    if childYangName == "burst" {
        return &police.Burst
    }
    if childYangName == "peak-burst" {
        return &police.PeakBurst
    }
    if childYangName == "conform-action" {
        return &police.ConformAction
    }
    if childYangName == "exceed-action" {
        return &police.ExceedAction
    }
    if childYangName == "violate-action" {
        return &police.ViolateAction
    }
    return nil
}

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rate"] = &police.Rate
    children["peak-rate"] = &police.PeakRate
    children["burst"] = &police.Burst
    children["peak-burst"] = &police.PeakBurst
    children["conform-action"] = &police.ConformAction
    children["exceed-action"] = &police.ExceedAction
    children["violate-action"] = &police.ViolateAction
    return children
}

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetYangName() string { return "police" }

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetParent() types.Entity { return police.parent }

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate
// Rate configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rate value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Rate units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(pps)|(percent)|(cellsps).
    Units interface{}
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetFilter() yfilter.YFilter { return rate.YFilter }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) SetFilter(yf yfilter.YFilter) { rate.YFilter = yf }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "units" { return "Units" }
    return ""
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetSegmentPath() string {
    return "rate"
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = rate.Value
    leafs["units"] = rate.Units
    return leafs
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetBundleName() string { return "cisco_ios_xr" }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetYangName() string { return "rate" }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) SetParent(parent types.Entity) { rate.parent = parent }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetParent() types.Entity { return rate.parent }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetParentYangName() string { return "police" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate
// Peak rate configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peak rate value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Peak rate units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(pps)|(percent)|(cellsps).
    Units interface{}
}

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetFilter() yfilter.YFilter { return peakRate.YFilter }

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) SetFilter(yf yfilter.YFilter) { peakRate.YFilter = yf }

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "units" { return "Units" }
    return ""
}

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetSegmentPath() string {
    return "peak-rate"
}

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = peakRate.Value
    leafs["units"] = peakRate.Units
    return leafs
}

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetBundleName() string { return "cisco_ios_xr" }

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetYangName() string { return "peak-rate" }

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) SetParent(parent types.Entity) { peakRate.parent = parent }

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetParent() types.Entity { return peakRate.parent }

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetParentYangName() string { return "police" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst
// Burst configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Burst value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Burst units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Units interface{}
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetFilter() yfilter.YFilter { return burst.YFilter }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) SetFilter(yf yfilter.YFilter) { burst.YFilter = yf }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "units" { return "Units" }
    return ""
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetSegmentPath() string {
    return "burst"
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = burst.Value
    leafs["units"] = burst.Units
    return leafs
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetBundleName() string { return "cisco_ios_xr" }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetYangName() string { return "burst" }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) SetParent(parent types.Entity) { burst.parent = parent }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetParent() types.Entity { return burst.parent }

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetParentYangName() string { return "police" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst
// Peak burst configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peak burst value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Peak burst units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Units interface{}
}

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetFilter() yfilter.YFilter { return peakBurst.YFilter }

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) SetFilter(yf yfilter.YFilter) { peakBurst.YFilter = yf }

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "units" { return "Units" }
    return ""
}

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetSegmentPath() string {
    return "peak-burst"
}

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = peakBurst.Value
    leafs["units"] = peakBurst.Units
    return leafs
}

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetBundleName() string { return "cisco_ios_xr" }

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetYangName() string { return "peak-burst" }

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) SetParent(parent types.Entity) { peakBurst.parent = parent }

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetParent() types.Entity { return peakBurst.parent }

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetParentYangName() string { return "police" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction
// Configures the action to take on packets that conform 
// to the rate limit.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Police action transmit. The type is interface{}.
    Transmit interface{}

    // Police action drop. The type is interface{}.
    Drop interface{}

    // Police action packet marking.
    Set PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set
}

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetFilter() yfilter.YFilter { return conformAction.YFilter }

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) SetFilter(yf yfilter.YFilter) { conformAction.YFilter = yf }

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetGoName(yname string) string {
    if yname == "Transmit" { return "Transmit" }
    if yname == "drop" { return "Drop" }
    if yname == "set" { return "Set" }
    return ""
}

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetSegmentPath() string {
    return "conform-action"
}

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        return &conformAction.Set
    }
    return nil
}

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["set"] = &conformAction.Set
    return children
}

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["Transmit"] = conformAction.Transmit
    leafs["drop"] = conformAction.Drop
    return leafs
}

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetBundleName() string { return "cisco_ios_xr" }

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetYangName() string { return "conform-action" }

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) SetParent(parent types.Entity) { conformAction.parent = parent }

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetParent() types.Entity { return conformAction.parent }

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetParentYangName() string { return "police" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set
// Police action packet marking.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Marks a packet by setting the DSCP in the ToS byte. The type is string with
    // pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Dscp interface{}

    // Sets the QoS group identifiers on IPv4 or MPLS packets. The set qos-group
    // is supported only on an ingress policy. The type is interface{} with range:
    // 0..512.
    QosGroup interface{}

    // Sets the Traffic class identifiers on IPv4 or MPLS packets. The type is
    // interface{} with range: 0..63.
    TrafficClass interface{}

    // Sets the discard class on IPv4 or MPLS packets. The discard-class can be
    // used only in service policies  that are attached in the ingress policy. The
    // type is interface{} with range: 0..7.
    DiscardClass interface{}

    // Sets the forward class. The type is interface{} with range: 0..7.
    ForwardClass interface{}

    // Set DF bit. The type is interface{} with range: 0..1.
    Df interface{}

    // Sets the specific IEEE 802.1Q Layer 2 CoS value of an outgoing packet. This
    // command should be used by a router if a user wants to mark a packet that is
    // being sent to a switch.  Switches can leverage Layer 2 header information, 
    // including a CoS value marking. Packets entering an  interface cannot be set
    // with a CoS value. The type is interface{} with range: 0..7.
    Cos interface{}

    // Set inner cos. The type is interface{} with range: 0..7.
    InnerCos interface{}

    // Sets the precedence value in the IP header. The type is one of the
    // following types: int with range: 0..7, or string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Precedence interface{}

    // Sets the precedence tunnel value for ipsec. The type is one of the
    // following types: int with range: 0..7, or string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    PrecedenceTunnel interface{}

    // Sets the experimental value of the MPLS packet top-most labels. The type is
    // interface{} with range: 0..7.
    MplsExperimentalTopMost interface{}

    // Sets the experimental value of the MPLS packet  imposition labels.
    // Imposition can be used only in service policies that  are attached in the
    // ingress policy. The type is interface{} with range: 0..7.
    MplsExperimentalImposition interface{}

    // Sets the spatial reuse protocol priority value of an  outgoing packet. The
    // type is interface{} with range: 0..7.
    SrpPriority interface{}

    // Set FrameRelay DE bit. The type is interface{} with range: 0..1.
    FrDe interface{}

    // Set DEI bit. The type is interface{} with range: 0..1.
    Dei interface{}

    // Set DEI imposition bit. The type is interface{} with range: 0..1.
    DeiImposition interface{}

    // Source IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "discard-class" { return "DiscardClass" }
    if yname == "forward-class" { return "ForwardClass" }
    if yname == "df" { return "Df" }
    if yname == "cos" { return "Cos" }
    if yname == "inner-cos" { return "InnerCos" }
    if yname == "precedence" { return "Precedence" }
    if yname == "precedence-tunnel" { return "PrecedenceTunnel" }
    if yname == "mpls-experimental-top-most" { return "MplsExperimentalTopMost" }
    if yname == "mpls-experimental-imposition" { return "MplsExperimentalImposition" }
    if yname == "srp-priority" { return "SrpPriority" }
    if yname == "fr-de" { return "FrDe" }
    if yname == "dei" { return "Dei" }
    if yname == "dei-imposition" { return "DeiImposition" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    return ""
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetSegmentPath() string {
    return "set"
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = set.Dscp
    leafs["qos-group"] = set.QosGroup
    leafs["traffic-class"] = set.TrafficClass
    leafs["discard-class"] = set.DiscardClass
    leafs["forward-class"] = set.ForwardClass
    leafs["df"] = set.Df
    leafs["cos"] = set.Cos
    leafs["inner-cos"] = set.InnerCos
    leafs["precedence"] = set.Precedence
    leafs["precedence-tunnel"] = set.PrecedenceTunnel
    leafs["mpls-experimental-top-most"] = set.MplsExperimentalTopMost
    leafs["mpls-experimental-imposition"] = set.MplsExperimentalImposition
    leafs["srp-priority"] = set.SrpPriority
    leafs["fr-de"] = set.FrDe
    leafs["dei"] = set.Dei
    leafs["dei-imposition"] = set.DeiImposition
    leafs["source-address"] = set.SourceAddress
    leafs["destination-address"] = set.DestinationAddress
    return leafs
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetYangName() string { return "set" }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetParent() types.Entity { return set.parent }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetParentYangName() string { return "conform-action" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction
// Configures the action to take on packets that exceed 
// the rate limit.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Police action transmit. The type is interface{}.
    Transmit interface{}

    // Police action drop. The type is interface{}.
    Drop interface{}

    // Police action packet marking.
    Set PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set
}

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetFilter() yfilter.YFilter { return exceedAction.YFilter }

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) SetFilter(yf yfilter.YFilter) { exceedAction.YFilter = yf }

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetGoName(yname string) string {
    if yname == "Transmit" { return "Transmit" }
    if yname == "drop" { return "Drop" }
    if yname == "set" { return "Set" }
    return ""
}

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetSegmentPath() string {
    return "exceed-action"
}

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        return &exceedAction.Set
    }
    return nil
}

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["set"] = &exceedAction.Set
    return children
}

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["Transmit"] = exceedAction.Transmit
    leafs["drop"] = exceedAction.Drop
    return leafs
}

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetBundleName() string { return "cisco_ios_xr" }

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetYangName() string { return "exceed-action" }

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) SetParent(parent types.Entity) { exceedAction.parent = parent }

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetParent() types.Entity { return exceedAction.parent }

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetParentYangName() string { return "police" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set
// Police action packet marking.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Marks a packet by setting the DSCP in the ToS byte. The type is string with
    // pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Dscp interface{}

    // Sets the QoS group identifiers on IPv4 or MPLS packets. The set qos-group
    // is supported only on an ingress policy. The type is interface{} with range:
    // 0..512.
    QosGroup interface{}

    // Sets the Traffic class identifiers on IPv4 or MPLS packets. The type is
    // interface{} with range: 0..63.
    TrafficClass interface{}

    // Sets the discard class on IPv4 or MPLS packets. The discard-class can be
    // used only in service policies  that are attached in the ingress policy. The
    // type is interface{} with range: 0..7.
    DiscardClass interface{}

    // Sets the forward class. The type is interface{} with range: 0..7.
    ForwardClass interface{}

    // Set DF bit. The type is interface{} with range: 0..1.
    Df interface{}

    // Sets the specific IEEE 802.1Q Layer 2 CoS value of an outgoing packet. This
    // command should be used by a router if a user wants to mark a packet that is
    // being sent to a switch.  Switches can leverage Layer 2 header information, 
    // including a CoS value marking. Packets entering an  interface cannot be set
    // with a CoS value. The type is interface{} with range: 0..7.
    Cos interface{}

    // Set inner cos. The type is interface{} with range: 0..7.
    InnerCos interface{}

    // Sets the precedence value in the IP header. The type is one of the
    // following types: int with range: 0..7, or string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Precedence interface{}

    // Sets the precedence tunnel value for ipsec. The type is one of the
    // following types: int with range: 0..7, or string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    PrecedenceTunnel interface{}

    // Sets the experimental value of the MPLS packet top-most labels. The type is
    // interface{} with range: 0..7.
    MplsExperimentalTopMost interface{}

    // Sets the experimental value of the MPLS packet  imposition labels.
    // Imposition can be used only in service policies that  are attached in the
    // ingress policy. The type is interface{} with range: 0..7.
    MplsExperimentalImposition interface{}

    // Sets the spatial reuse protocol priority value of an  outgoing packet. The
    // type is interface{} with range: 0..7.
    SrpPriority interface{}

    // Set FrameRelay DE bit. The type is interface{} with range: 0..1.
    FrDe interface{}

    // Set DEI bit. The type is interface{} with range: 0..1.
    Dei interface{}

    // Set DEI imposition bit. The type is interface{} with range: 0..1.
    DeiImposition interface{}

    // Source IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "discard-class" { return "DiscardClass" }
    if yname == "forward-class" { return "ForwardClass" }
    if yname == "df" { return "Df" }
    if yname == "cos" { return "Cos" }
    if yname == "inner-cos" { return "InnerCos" }
    if yname == "precedence" { return "Precedence" }
    if yname == "precedence-tunnel" { return "PrecedenceTunnel" }
    if yname == "mpls-experimental-top-most" { return "MplsExperimentalTopMost" }
    if yname == "mpls-experimental-imposition" { return "MplsExperimentalImposition" }
    if yname == "srp-priority" { return "SrpPriority" }
    if yname == "fr-de" { return "FrDe" }
    if yname == "dei" { return "Dei" }
    if yname == "dei-imposition" { return "DeiImposition" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    return ""
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetSegmentPath() string {
    return "set"
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = set.Dscp
    leafs["qos-group"] = set.QosGroup
    leafs["traffic-class"] = set.TrafficClass
    leafs["discard-class"] = set.DiscardClass
    leafs["forward-class"] = set.ForwardClass
    leafs["df"] = set.Df
    leafs["cos"] = set.Cos
    leafs["inner-cos"] = set.InnerCos
    leafs["precedence"] = set.Precedence
    leafs["precedence-tunnel"] = set.PrecedenceTunnel
    leafs["mpls-experimental-top-most"] = set.MplsExperimentalTopMost
    leafs["mpls-experimental-imposition"] = set.MplsExperimentalImposition
    leafs["srp-priority"] = set.SrpPriority
    leafs["fr-de"] = set.FrDe
    leafs["dei"] = set.Dei
    leafs["dei-imposition"] = set.DeiImposition
    leafs["source-address"] = set.SourceAddress
    leafs["destination-address"] = set.DestinationAddress
    return leafs
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetYangName() string { return "set" }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetParent() types.Entity { return set.parent }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetParentYangName() string { return "exceed-action" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction
// Configures the action to take on packets that violate
// the rate limit.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Police action transmit. The type is interface{}.
    Transmit interface{}

    // Police action drop. The type is interface{}.
    Drop interface{}

    // Police action packet marking.
    Set PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set
}

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetFilter() yfilter.YFilter { return violateAction.YFilter }

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) SetFilter(yf yfilter.YFilter) { violateAction.YFilter = yf }

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetGoName(yname string) string {
    if yname == "Transmit" { return "Transmit" }
    if yname == "drop" { return "Drop" }
    if yname == "set" { return "Set" }
    return ""
}

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetSegmentPath() string {
    return "violate-action"
}

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set" {
        return &violateAction.Set
    }
    return nil
}

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["set"] = &violateAction.Set
    return children
}

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["Transmit"] = violateAction.Transmit
    leafs["drop"] = violateAction.Drop
    return leafs
}

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetBundleName() string { return "cisco_ios_xr" }

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetYangName() string { return "violate-action" }

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) SetParent(parent types.Entity) { violateAction.parent = parent }

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetParent() types.Entity { return violateAction.parent }

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetParentYangName() string { return "police" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set
// Police action packet marking.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Marks a packet by setting the DSCP in the ToS byte. The type is string with
    // pattern:
    // ([0-9]|[1-5][0-9]|6[0-3])|(af11)|(af12)|(af13)|(af21)|(af22)|(af23)|(af31)|(af32)|(af33)|(af41)|(af42)|(af43)|(ef)|(default)|(cs1)|(cs2)|(cs3)|(cs4)|(cs5)|(cs6)|(cs7).
    Dscp interface{}

    // Sets the QoS group identifiers on IPv4 or MPLS packets. The set qos-group
    // is supported only on an ingress policy. The type is interface{} with range:
    // 0..512.
    QosGroup interface{}

    // Sets the Traffic class identifiers on IPv4 or MPLS packets. The type is
    // interface{} with range: 0..63.
    TrafficClass interface{}

    // Sets the discard class on IPv4 or MPLS packets. The discard-class can be
    // used only in service policies  that are attached in the ingress policy. The
    // type is interface{} with range: 0..7.
    DiscardClass interface{}

    // Sets the forward class. The type is interface{} with range: 0..7.
    ForwardClass interface{}

    // Set DF bit. The type is interface{} with range: 0..1.
    Df interface{}

    // Sets the specific IEEE 802.1Q Layer 2 CoS value of an outgoing packet. This
    // command should be used by a router if a user wants to mark a packet that is
    // being sent to a switch.  Switches can leverage Layer 2 header information, 
    // including a CoS value marking. Packets entering an  interface cannot be set
    // with a CoS value. The type is interface{} with range: 0..7.
    Cos interface{}

    // Set inner cos. The type is interface{} with range: 0..7.
    InnerCos interface{}

    // Sets the precedence value in the IP header. The type is one of the
    // following types: int with range: 0..7, or string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    Precedence interface{}

    // Sets the precedence tunnel value for ipsec. The type is one of the
    // following types: int with range: 0..7, or string with pattern:
    // (critical)|(flash)|(flash-override)|(immediate)|(internet)|(network)|(priority)|(routine).
    PrecedenceTunnel interface{}

    // Sets the experimental value of the MPLS packet top-most labels. The type is
    // interface{} with range: 0..7.
    MplsExperimentalTopMost interface{}

    // Sets the experimental value of the MPLS packet  imposition labels.
    // Imposition can be used only in service policies that  are attached in the
    // ingress policy. The type is interface{} with range: 0..7.
    MplsExperimentalImposition interface{}

    // Sets the spatial reuse protocol priority value of an  outgoing packet. The
    // type is interface{} with range: 0..7.
    SrpPriority interface{}

    // Set FrameRelay DE bit. The type is interface{} with range: 0..1.
    FrDe interface{}

    // Set DEI bit. The type is interface{} with range: 0..1.
    Dei interface{}

    // Set DEI imposition bit. The type is interface{} with range: 0..1.
    DeiImposition interface{}

    // Source IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetFilter() yfilter.YFilter { return set.YFilter }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) SetFilter(yf yfilter.YFilter) { set.YFilter = yf }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "discard-class" { return "DiscardClass" }
    if yname == "forward-class" { return "ForwardClass" }
    if yname == "df" { return "Df" }
    if yname == "cos" { return "Cos" }
    if yname == "inner-cos" { return "InnerCos" }
    if yname == "precedence" { return "Precedence" }
    if yname == "precedence-tunnel" { return "PrecedenceTunnel" }
    if yname == "mpls-experimental-top-most" { return "MplsExperimentalTopMost" }
    if yname == "mpls-experimental-imposition" { return "MplsExperimentalImposition" }
    if yname == "srp-priority" { return "SrpPriority" }
    if yname == "fr-de" { return "FrDe" }
    if yname == "dei" { return "Dei" }
    if yname == "dei-imposition" { return "DeiImposition" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    return ""
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetSegmentPath() string {
    return "set"
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = set.Dscp
    leafs["qos-group"] = set.QosGroup
    leafs["traffic-class"] = set.TrafficClass
    leafs["discard-class"] = set.DiscardClass
    leafs["forward-class"] = set.ForwardClass
    leafs["df"] = set.Df
    leafs["cos"] = set.Cos
    leafs["inner-cos"] = set.InnerCos
    leafs["precedence"] = set.Precedence
    leafs["precedence-tunnel"] = set.PrecedenceTunnel
    leafs["mpls-experimental-top-most"] = set.MplsExperimentalTopMost
    leafs["mpls-experimental-imposition"] = set.MplsExperimentalImposition
    leafs["srp-priority"] = set.SrpPriority
    leafs["fr-de"] = set.FrDe
    leafs["dei"] = set.Dei
    leafs["dei-imposition"] = set.DeiImposition
    leafs["source-address"] = set.SourceAddress
    leafs["destination-address"] = set.DestinationAddress
    return leafs
}

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetBundleName() string { return "cisco_ios_xr" }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetYangName() string { return "set" }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) SetParent(parent types.Entity) { set.parent = parent }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetParent() types.Entity { return set.parent }

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetParentYangName() string { return "violate-action" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy
// Configure a child service policy.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of service-policy. The type is string with pattern:
    // [a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}.
    PolicyName interface{}

    // Type of service-policy. The type is string with pattern:
    // (PBR)|(QOS)|(REDIRECT)|(TRAFFIC)|(pbr)|(qos)|(redirect)|(traffic).
    Type interface{}
}

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetFilter() yfilter.YFilter { return servicePolicy.YFilter }

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) SetFilter(yf yfilter.YFilter) { servicePolicy.YFilter = yf }

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "type" { return "Type" }
    return ""
}

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetSegmentPath() string {
    return "service-policy"
}

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = servicePolicy.PolicyName
    leafs["type"] = servicePolicy.Type
    return leafs
}

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetBundleName() string { return "cisco_ios_xr" }

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetYangName() string { return "service-policy" }

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) SetParent(parent types.Entity) { servicePolicy.parent = parent }

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetParent() types.Entity { return servicePolicy.parent }

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal
// Policy action CAC.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interval after which a flow is removed,  if there is no activity. If
    // timeout is 0 this flow does not expire. The type is one of the following
    // types: int with range: 10..2550, or string with pattern: (None)|(none).
    FlowIdleTimeout interface{}

    // The rate allocated for all flows.
    Rate PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate

    // The rate allocated per flow.
    FlowRate PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate
}

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetFilter() yfilter.YFilter { return cacLocal.YFilter }

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) SetFilter(yf yfilter.YFilter) { cacLocal.YFilter = yf }

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetGoName(yname string) string {
    if yname == "flow-idle-timeout" { return "FlowIdleTimeout" }
    if yname == "rate" { return "Rate" }
    if yname == "flow-rate" { return "FlowRate" }
    return ""
}

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetSegmentPath() string {
    return "cac-local"
}

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rate" {
        return &cacLocal.Rate
    }
    if childYangName == "flow-rate" {
        return &cacLocal.FlowRate
    }
    return nil
}

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rate"] = &cacLocal.Rate
    children["flow-rate"] = &cacLocal.FlowRate
    return children
}

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flow-idle-timeout"] = cacLocal.FlowIdleTimeout
    return leafs
}

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetBundleName() string { return "cisco_ios_xr" }

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetYangName() string { return "cac-local" }

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) SetParent(parent types.Entity) { cacLocal.parent = parent }

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetParent() types.Entity { return cacLocal.parent }

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate
// The rate allocated for all flows.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rate value. The type is interface{} with range: 1..4294967295.
    Value interface{}

    // Rate units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(cellsps).
    Units interface{}
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetFilter() yfilter.YFilter { return rate.YFilter }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) SetFilter(yf yfilter.YFilter) { rate.YFilter = yf }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "units" { return "Units" }
    return ""
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetSegmentPath() string {
    return "rate"
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = rate.Value
    leafs["units"] = rate.Units
    return leafs
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetBundleName() string { return "cisco_ios_xr" }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetYangName() string { return "rate" }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) SetParent(parent types.Entity) { rate.parent = parent }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetParent() types.Entity { return rate.parent }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetParentYangName() string { return "cac-local" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate
// The rate allocated per flow.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rate value. The type is interface{} with range: 1..4294967295.
    Value interface{}

    // Rate units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(cellsps).
    Units interface{}
}

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetFilter() yfilter.YFilter { return flowRate.YFilter }

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) SetFilter(yf yfilter.YFilter) { flowRate.YFilter = yf }

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "units" { return "Units" }
    return ""
}

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetSegmentPath() string {
    return "flow-rate"
}

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = flowRate.Value
    leafs["units"] = flowRate.Units
    return leafs
}

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetBundleName() string { return "cisco_ios_xr" }

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetYangName() string { return "flow-rate" }

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) SetParent(parent types.Entity) { flowRate.parent = parent }

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetParent() types.Entity { return flowRate.parent }

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetParentYangName() string { return "cac-local" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams
// Policy flow monitoring action.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Max simultaneous flows monitored per policy class. The type is interface{}
    // with range: 0..4096.
    MaxFlow interface{}

    // Monitored interval duration. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    IntervalDuration interface{}

    // Keep stats/metrics on box for so many intervals. The type is interface{}
    // with range: 0..4294967295.
    History interface{}

    // Declare a flow dead if no packets received in so much time. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Timeout interface{}
}

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetFilter() yfilter.YFilter { return flowParams.YFilter }

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) SetFilter(yf yfilter.YFilter) { flowParams.YFilter = yf }

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetGoName(yname string) string {
    if yname == "max-flow" { return "MaxFlow" }
    if yname == "interval-duration" { return "IntervalDuration" }
    if yname == "history" { return "History" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetSegmentPath() string {
    return "flow-params"
}

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-flow"] = flowParams.MaxFlow
    leafs["interval-duration"] = flowParams.IntervalDuration
    leafs["history"] = flowParams.History
    leafs["timeout"] = flowParams.Timeout
    return leafs
}

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetBundleName() string { return "cisco_ios_xr" }

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetYangName() string { return "flow-params" }

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) SetParent(parent types.Entity) { flowParams.parent = parent }

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetParent() types.Entity { return flowParams.parent }

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr
// Policy IP-CBR metric action.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Nominal per-flow data rate.
    Rate PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate

    // Media-packet structure.
    MediaPacket PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket
}

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetFilter() yfilter.YFilter { return metricsIpcbr.YFilter }

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) SetFilter(yf yfilter.YFilter) { metricsIpcbr.YFilter = yf }

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetGoName(yname string) string {
    if yname == "rate" { return "Rate" }
    if yname == "media-packet" { return "MediaPacket" }
    return ""
}

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetSegmentPath() string {
    return "metrics-ipcbr"
}

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rate" {
        return &metricsIpcbr.Rate
    }
    if childYangName == "media-packet" {
        return &metricsIpcbr.MediaPacket
    }
    return nil
}

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rate"] = &metricsIpcbr.Rate
    children["media-packet"] = &metricsIpcbr.MediaPacket
    return children
}

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetBundleName() string { return "cisco_ios_xr" }

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetYangName() string { return "metrics-ipcbr" }

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) SetParent(parent types.Entity) { metricsIpcbr.parent = parent }

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetParent() types.Entity { return metricsIpcbr.parent }

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate
// Nominal per-flow data rate.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Nominal rate specified at the L3 (IP). The type is interface{} with range:
    // 0..4294967295. Units are bps.
    Layer3 interface{}

    // Nominal IP layer packet rate (in pps). The type is interface{} with range:
    // 0..4294967295. Units are pps.
    Packet interface{}

    // Nominal data rate of the media flow (ip payload). The type is interface{}
    // with range: 1..3000000000. Units are bps.
    Media interface{}
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetFilter() yfilter.YFilter { return rate.YFilter }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) SetFilter(yf yfilter.YFilter) { rate.YFilter = yf }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetGoName(yname string) string {
    if yname == "layer3" { return "Layer3" }
    if yname == "packet" { return "Packet" }
    if yname == "media" { return "Media" }
    return ""
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetSegmentPath() string {
    return "rate"
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["layer3"] = rate.Layer3
    leafs["packet"] = rate.Packet
    leafs["media"] = rate.Media
    return leafs
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetBundleName() string { return "cisco_ios_xr" }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetYangName() string { return "rate" }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) SetParent(parent types.Entity) { rate.parent = parent }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetParent() types.Entity { return rate.parent }

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetParentYangName() string { return "metrics-ipcbr" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket
// Media-packet structure.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Nominal size of the media-packet. The type is interface{} with range:
    // 0..65535. Units are bytes.
    Size interface{}

    // Nominal number of media packets in an IP payload. The type is interface{}
    // with range: 1..64. Units are packets.
    CountInLayer3 interface{}
}

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetFilter() yfilter.YFilter { return mediaPacket.YFilter }

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) SetFilter(yf yfilter.YFilter) { mediaPacket.YFilter = yf }

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "count-in-layer3" { return "CountInLayer3" }
    return ""
}

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetSegmentPath() string {
    return "media-packet"
}

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = mediaPacket.Size
    leafs["count-in-layer3"] = mediaPacket.CountInLayer3
    return leafs
}

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetBundleName() string { return "cisco_ios_xr" }

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetYangName() string { return "media-packet" }

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) SetParent(parent types.Entity) { mediaPacket.parent = parent }

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetParent() types.Entity { return mediaPacket.parent }

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetParentYangName() string { return "metrics-ipcbr" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React
// Policy action react.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // String describing the react statement. The type is string.
    Descrition interface{}

    // React criterion delay factor. The type is interface{}.
    CriterionDelayFactor interface{}

    // React criterion media stop. The type is interface{}.
    CriterionMediaStop interface{}

    // React criterion mrv. The type is interface{}.
    CriterionMrv interface{}

    // React criterion flow count. The type is interface{}.
    CriterionFlowCount interface{}

    // React criterion packet rate. The type is interface{}.
    CriterionPacketRate interface{}

    // Action on alert.
    Action PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action

    // Alaram settings.
    Alarm PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm

    // Alarm threshold settings.
    Threshold PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold
}

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetFilter() yfilter.YFilter { return react.YFilter }

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) SetFilter(yf yfilter.YFilter) { react.YFilter = yf }

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetGoName(yname string) string {
    if yname == "descrition" { return "Descrition" }
    if yname == "criterion-delay-factor" { return "CriterionDelayFactor" }
    if yname == "criterion-media-stop" { return "CriterionMediaStop" }
    if yname == "criterion-mrv" { return "CriterionMrv" }
    if yname == "criterion-flow-count" { return "CriterionFlowCount" }
    if yname == "criterion-packet-rate" { return "CriterionPacketRate" }
    if yname == "action" { return "Action" }
    if yname == "alarm" { return "Alarm" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetSegmentPath() string {
    return "react"
}

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "action" {
        return &react.Action
    }
    if childYangName == "alarm" {
        return &react.Alarm
    }
    if childYangName == "threshold" {
        return &react.Threshold
    }
    return nil
}

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["action"] = &react.Action
    children["alarm"] = &react.Alarm
    children["threshold"] = &react.Threshold
    return children
}

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["descrition"] = react.Descrition
    leafs["criterion-delay-factor"] = react.CriterionDelayFactor
    leafs["criterion-media-stop"] = react.CriterionMediaStop
    leafs["criterion-mrv"] = react.CriterionMrv
    leafs["criterion-flow-count"] = react.CriterionFlowCount
    leafs["criterion-packet-rate"] = react.CriterionPacketRate
    return leafs
}

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetBundleName() string { return "cisco_ios_xr" }

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetYangName() string { return "react" }

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) SetParent(parent types.Entity) { react.parent = parent }

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetParent() types.Entity { return react.parent }

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action
// Action on alert.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Syslog. The type is interface{}.
    Syslog interface{}

    // SNMP. The type is interface{}.
    Snmp interface{}
}

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetFilter() yfilter.YFilter { return action.YFilter }

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) SetFilter(yf yfilter.YFilter) { action.YFilter = yf }

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetGoName(yname string) string {
    if yname == "syslog" { return "Syslog" }
    if yname == "snmp" { return "Snmp" }
    return ""
}

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetSegmentPath() string {
    return "action"
}

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["syslog"] = action.Syslog
    leafs["snmp"] = action.Snmp
    return leafs
}

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetBundleName() string { return "cisco_ios_xr" }

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetYangName() string { return "action" }

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) SetParent(parent types.Entity) { action.parent = parent }

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetParent() types.Entity { return action.parent }

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetParentYangName() string { return "react" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm
// Alaram settings.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Severity of the alarm. The type is string with pattern:
    // (informational)|(notification)|(warning)|(error)|(critical)|(alert)|(emergency).
    Severity interface{}

    // Alarm type.
    Type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type
}

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetFilter() yfilter.YFilter { return alarm.YFilter }

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) SetFilter(yf yfilter.YFilter) { alarm.YFilter = yf }

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetGoName(yname string) string {
    if yname == "severity" { return "Severity" }
    if yname == "type" { return "Type" }
    return ""
}

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetSegmentPath() string {
    return "alarm"
}

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "type" {
        return &alarm.Type
    }
    return nil
}

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["type"] = &alarm.Type
    return children
}

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["severity"] = alarm.Severity
    return leafs
}

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetBundleName() string { return "cisco_ios_xr" }

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetYangName() string { return "alarm" }

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) SetParent(parent types.Entity) { alarm.parent = parent }

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetParent() types.Entity { return alarm.parent }

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetParentYangName() string { return "react" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type
// Alarm type.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Discrete alarm type. The type is interface{}.
    Discrete interface{}

    // Number of flows to reach before  triggering alarm. The type is interface{}
    // with range: 0..65535. Units are number of flows.
    GroupCount interface{}

    // Percent to reach before triggering alarm. The type is interface{} with
    // range: 0..65535. Units are percentage.
    GroupPercent interface{}
}

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetGoName(yname string) string {
    if yname == "discrete" { return "Discrete" }
    if yname == "group-count" { return "GroupCount" }
    if yname == "group-percent" { return "GroupPercent" }
    return ""
}

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetSegmentPath() string {
    return "type"
}

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discrete"] = self.Discrete
    leafs["group-count"] = self.GroupCount
    leafs["group-percent"] = self.GroupPercent
    return leafs
}

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetBundleName() string { return "cisco_ios_xr" }

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetYangName() string { return "type" }

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) SetParent(parent types.Entity) { self.parent = parent }

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetParent() types.Entity { return self.parent }

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetParentYangName() string { return "alarm" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold
// Alarm threshold settings.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Alarm trigger value settings.
    TriggerValue PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue

    // Alarm trigger type settings.
    TriggerType PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType
}

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetGoName(yname string) string {
    if yname == "trigger-value" { return "TriggerValue" }
    if yname == "trigger-type" { return "TriggerType" }
    return ""
}

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trigger-value" {
        return &threshold.TriggerValue
    }
    if childYangName == "trigger-type" {
        return &threshold.TriggerType
    }
    return nil
}

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["trigger-value"] = &threshold.TriggerValue
    children["trigger-type"] = &threshold.TriggerType
    return children
}

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetYangName() string { return "threshold" }

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetParentYangName() string { return "react" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue
// Alarm trigger value settings.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Greater than. The type is string.
    GreaterThan interface{}

    // Greater than equal. The type is string.
    GreaterThanEqual interface{}

    // Less than. The type is string.
    LessThan interface{}

    // Less than equal. The type is string.
    LessThanEqual interface{}

    // Range. The type is string.
    Range interface{}
}

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetFilter() yfilter.YFilter { return triggerValue.YFilter }

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) SetFilter(yf yfilter.YFilter) { triggerValue.YFilter = yf }

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetGoName(yname string) string {
    if yname == "greater-than" { return "GreaterThan" }
    if yname == "greater-than-equal" { return "GreaterThanEqual" }
    if yname == "less-than" { return "LessThan" }
    if yname == "less-than-equal" { return "LessThanEqual" }
    if yname == "range" { return "Range" }
    return ""
}

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetSegmentPath() string {
    return "trigger-value"
}

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["greater-than"] = triggerValue.GreaterThan
    leafs["greater-than-equal"] = triggerValue.GreaterThanEqual
    leafs["less-than"] = triggerValue.LessThan
    leafs["less-than-equal"] = triggerValue.LessThanEqual
    leafs["range"] = triggerValue.Range
    return leafs
}

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetBundleName() string { return "cisco_ios_xr" }

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetYangName() string { return "trigger-value" }

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) SetParent(parent types.Entity) { triggerValue.parent = parent }

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetParent() types.Entity { return triggerValue.parent }

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetParentYangName() string { return "threshold" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType
// Alarm trigger type settings.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Immediate trigger. The type is interface{}.
    Immediate interface{}

    // Trigger averaged over N intervals. The type is interface{} with range:
    // 0..4294967295.
    Average interface{}
}

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetFilter() yfilter.YFilter { return triggerType.YFilter }

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) SetFilter(yf yfilter.YFilter) { triggerType.YFilter = yf }

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetGoName(yname string) string {
    if yname == "immediate" { return "Immediate" }
    if yname == "average" { return "Average" }
    return ""
}

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetSegmentPath() string {
    return "trigger-type"
}

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["immediate"] = triggerType.Immediate
    leafs["average"] = triggerType.Average
    return leafs
}

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetBundleName() string { return "cisco_ios_xr" }

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetYangName() string { return "trigger-type" }

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) SetParent(parent types.Entity) { triggerType.parent = parent }

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetParent() types.Entity { return triggerType.parent }

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetParentYangName() string { return "threshold" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect
// Policy action redirect
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy action redirect IPv4.
    Ipv4 PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4

    // Policy action redirect IPv6.
    Ipv6 PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6

    // Next hop address.
    NextHop PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop
}

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetFilter() yfilter.YFilter { return pbrRedirect.YFilter }

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) SetFilter(yf yfilter.YFilter) { pbrRedirect.YFilter = yf }

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetGoName(yname string) string {
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetSegmentPath() string {
    return "pbr-redirect"
}

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &pbrRedirect.Ipv4
    }
    if childYangName == "ipv6" {
        return &pbrRedirect.Ipv6
    }
    if childYangName == "next-hop" {
        return &pbrRedirect.NextHop
    }
    return nil
}

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &pbrRedirect.Ipv4
    children["ipv6"] = &pbrRedirect.Ipv6
    children["next-hop"] = &pbrRedirect.NextHop
    return children
}

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetBundleName() string { return "cisco_ios_xr" }

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetYangName() string { return "pbr-redirect" }

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) SetParent(parent types.Entity) { pbrRedirect.parent = parent }

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetParent() types.Entity { return pbrRedirect.parent }

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4
// Policy action redirect IPv4
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4NextHop interface{}

    // IPv4 VRF. The type is string.
    Vrf interface{}
}

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetGoName(yname string) string {
    if yname == "ipv4-next-hop" { return "Ipv4NextHop" }
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv4-next-hop"] = ipv4.Ipv4NextHop
    leafs["vrf"] = ipv4.Vrf
    return leafs
}

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetParentYangName() string { return "pbr-redirect" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6
// Policy action redirect IPv6
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6NextHop interface{}

    // IPv6 VRF. The type is string.
    Vrf interface{}
}

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetGoName(yname string) string {
    if yname == "ipv6-next-hop" { return "Ipv6NextHop" }
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-next-hop"] = ipv6.Ipv6NextHop
    leafs["vrf"] = ipv6.Vrf
    return leafs
}

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetParentYangName() string { return "pbr-redirect" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop
// Next hop address.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route Target.
    RouteTarget PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetGoName(yname string) string {
    if yname == "route-target" { return "RouteTarget" }
    return ""
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route-target" {
        return &nextHop.RouteTarget
    }
    return nil
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route-target"] = &nextHop.RouteTarget
    return children
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetParentYangName() string { return "pbr-redirect" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget
// Route Target
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 2-byte/4-byte AS number. The type is interface{} with range: 1..4294967295.
    AsNumber interface{}

    // ASN2:index 2/4 byte (hex or decimal format). The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // IPv4 address.
    Ipv4Address PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address
}

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetFilter() yfilter.YFilter { return routeTarget.YFilter }

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) SetFilter(yf yfilter.YFilter) { routeTarget.YFilter = yf }

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetGoName(yname string) string {
    if yname == "as-number" { return "AsNumber" }
    if yname == "index" { return "Index" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    return ""
}

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetSegmentPath() string {
    return "route-target"
}

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-address" {
        return &routeTarget.Ipv4Address
    }
    return nil
}

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-address"] = &routeTarget.Ipv4Address
    return children
}

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-number"] = routeTarget.AsNumber
    leafs["index"] = routeTarget.Index
    return leafs
}

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetYangName() string { return "route-target" }

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) SetParent(parent types.Entity) { routeTarget.parent = parent }

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetParent() types.Entity { return routeTarget.parent }

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetParentYangName() string { return "next-hop" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address
// IPv4 address.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // IPv4 netmask. The type is string.
    Netmask interface{}
}

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetFilter() yfilter.YFilter { return ipv4Address.YFilter }

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) SetFilter(yf yfilter.YFilter) { ipv4Address.YFilter = yf }

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "netmask" { return "Netmask" }
    return ""
}

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetSegmentPath() string {
    return "ipv4-address"
}

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipv4Address.Address
    leafs["netmask"] = ipv4Address.Netmask
    return leafs
}

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetYangName() string { return "ipv4-address" }

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) SetParent(parent types.Entity) { ipv4Address.parent = parent }

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetParent() types.Entity { return ipv4Address.parent }

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetParentYangName() string { return "route-target" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward
// Policy action PBR forward.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Use system default routing table. The type is interface{}.
    Default interface{}

    // Use specific next-hop. Here we present 5 different combination  for the pbf
    // next-hop.  1. vrf with v6 address  2. vrf with v4 address  3. vrf   4. v4
    // address  5. v6 address.
    NextHop PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop
}

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetFilter() yfilter.YFilter { return pbrForward.YFilter }

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) SetFilter(yf yfilter.YFilter) { pbrForward.YFilter = yf }

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetGoName(yname string) string {
    if yname == "default" { return "Default" }
    if yname == "next-hop" { return "NextHop" }
    return ""
}

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetSegmentPath() string {
    return "pbr-forward"
}

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop" {
        return &pbrForward.NextHop
    }
    return nil
}

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop"] = &pbrForward.NextHop
    return children
}

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default"] = pbrForward.Default
    return leafs
}

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetBundleName() string { return "cisco_ios_xr" }

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetYangName() string { return "pbr-forward" }

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) SetParent(parent types.Entity) { pbrForward.parent = parent }

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetParent() types.Entity { return pbrForward.parent }

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetParentYangName() string { return "policy-map-rule" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop
// Use specific next-hop.
// Here we present 5 different combination 
// for the pbf next-hop.
//  1. vrf with v6 address
//  2. vrf with v4 address
//  3. vrf 
//  4. v4 address
//  5. v6 address
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is string.
    Vrf interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf"] = nextHop.Vrf
    leafs["ipv4-address"] = nextHop.Ipv4Address
    leafs["ipv6-address"] = nextHop.Ipv6Address
    return leafs
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetParentYangName() string { return "pbr-forward" }

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath
// Policy action service function path.
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service function path id. The type is interface{} with range: 1..16777215.
    // This attribute is mandatory.
    PathId interface{}

    // Service function path index. The type is interface{} with range: 1..255.
    // This attribute is mandatory.
    Index interface{}

    // Service function path metadata name. The type is string.
    Metadata interface{}
}

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetFilter() yfilter.YFilter { return serviceFunctionPath.YFilter }

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) SetFilter(yf yfilter.YFilter) { serviceFunctionPath.YFilter = yf }

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "index" { return "Index" }
    if yname == "metadata" { return "Metadata" }
    return ""
}

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetSegmentPath() string {
    return "service-function-path"
}

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = serviceFunctionPath.PathId
    leafs["index"] = serviceFunctionPath.Index
    leafs["metadata"] = serviceFunctionPath.Metadata
    return leafs
}

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetBundleName() string { return "cisco_ios_xr" }

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetYangName() string { return "service-function-path" }

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) SetParent(parent types.Entity) { serviceFunctionPath.parent = parent }

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetParent() types.Entity { return serviceFunctionPath.parent }

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetParentYangName() string { return "policy-map-rule" }

