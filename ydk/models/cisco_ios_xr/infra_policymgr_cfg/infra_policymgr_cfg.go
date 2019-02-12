// This module contains a collection of YANG definitions
// for Cisco IOS-XR ASR9k policy manager configuration.
//  
// Copyright (c) 2013, 2015-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class-maps configuration.
    ClassMaps PolicyManager_ClassMaps

    // Policy-maps configuration.
    PolicyMaps PolicyManager_PolicyMaps
}

func (policyManager *PolicyManager) GetEntityData() *types.CommonEntityData {
    policyManager.EntityData.YFilter = policyManager.YFilter
    policyManager.EntityData.YangName = "policy-manager"
    policyManager.EntityData.BundleName = "cisco_ios_xr"
    policyManager.EntityData.ParentYangName = "Cisco-IOS-XR-infra-policymgr-cfg"
    policyManager.EntityData.SegmentPath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager"
    policyManager.EntityData.AbsolutePath = policyManager.EntityData.SegmentPath
    policyManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyManager.EntityData.Children = types.NewOrderedMap()
    policyManager.EntityData.Children.Append("class-maps", types.YChild{"ClassMaps", &policyManager.ClassMaps})
    policyManager.EntityData.Children.Append("policy-maps", types.YChild{"PolicyMaps", &policyManager.PolicyMaps})
    policyManager.EntityData.Leafs = types.NewOrderedMap()

    policyManager.EntityData.YListKeys = []string {}

    return &(policyManager.EntityData)
}

// PolicyManager_ClassMaps
// Class-maps configuration.
type PolicyManager_ClassMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Class-map configuration. The type is slice of
    // PolicyManager_ClassMaps_ClassMap.
    ClassMap []*PolicyManager_ClassMaps_ClassMap
}

func (classMaps *PolicyManager_ClassMaps) GetEntityData() *types.CommonEntityData {
    classMaps.EntityData.YFilter = classMaps.YFilter
    classMaps.EntityData.YangName = "class-maps"
    classMaps.EntityData.BundleName = "cisco_ios_xr"
    classMaps.EntityData.ParentYangName = "policy-manager"
    classMaps.EntityData.SegmentPath = "class-maps"
    classMaps.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/" + classMaps.EntityData.SegmentPath
    classMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMaps.EntityData.Children = types.NewOrderedMap()
    classMaps.EntityData.Children.Append("class-map", types.YChild{"ClassMap", nil})
    for i := range classMaps.ClassMap {
        classMaps.EntityData.Children.Append(types.GetSegmentPath(classMaps.ClassMap[i]), types.YChild{"ClassMap", classMaps.ClassMap[i]})
    }
    classMaps.EntityData.Leafs = types.NewOrderedMap()

    classMaps.EntityData.YListKeys = []string {}

    return &(classMaps.EntityData)
}

// PolicyManager_ClassMaps_ClassMap
// Class-map configuration.
type PolicyManager_ClassMaps_ClassMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (classMap *PolicyManager_ClassMaps_ClassMap) GetEntityData() *types.CommonEntityData {
    classMap.EntityData.YFilter = classMap.YFilter
    classMap.EntityData.YangName = "class-map"
    classMap.EntityData.BundleName = "cisco_ios_xr"
    classMap.EntityData.ParentYangName = "class-maps"
    classMap.EntityData.SegmentPath = "class-map" + types.AddKeyToken(classMap.Type, "type") + types.AddKeyToken(classMap.Name, "name")
    classMap.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/" + classMap.EntityData.SegmentPath
    classMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMap.EntityData.Children = types.NewOrderedMap()
    classMap.EntityData.Children.Append("match", types.YChild{"Match", &classMap.Match})
    classMap.EntityData.Children.Append("match-not", types.YChild{"MatchNot", &classMap.MatchNot})
    classMap.EntityData.Leafs = types.NewOrderedMap()
    classMap.EntityData.Leafs.Append("type", types.YLeaf{"Type", classMap.Type})
    classMap.EntityData.Leafs.Append("name", types.YLeaf{"Name", classMap.Name})
    classMap.EntityData.Leafs.Append("class-map-mode-match-any", types.YLeaf{"ClassMapModeMatchAny", classMap.ClassMapModeMatchAny})
    classMap.EntityData.Leafs.Append("class-map-mode-match-all", types.YLeaf{"ClassMapModeMatchAll", classMap.ClassMapModeMatchAll})
    classMap.EntityData.Leafs.Append("description", types.YLeaf{"Description", classMap.Description})

    classMap.EntityData.YListKeys = []string {"Type", "Name"}

    return &(classMap.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match
// Match rules.
type PolicyManager_ClassMaps_ClassMap_Match struct {
    EntityData types.CommonEntityData
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
    DestinationAddressIpv4 []*PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4

    // Match destination IPv6 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6.
    DestinationAddressIpv6 []*PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6

    // Match source IPv4 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4.
    SourceAddressIpv4 []*PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4

    // Match source IPv6 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6.
    SourceAddressIpv6 []*PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6

    // Match dhcp client ID. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId.
    DhcpClientId []*PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId

    // Match dhcp client id regex. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex.
    DhcpClientIdRegex []*PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex

    // Match domain name. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DomainName.
    DomainName []*PolicyManager_ClassMaps_ClassMap_Match_DomainName

    // Match domain name. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex.
    DomainNameRegex []*PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex

    // Match flow.
    Flow PolicyManager_ClassMaps_ClassMap_Match_Flow
}

func (match *PolicyManager_ClassMaps_ClassMap_Match) GetEntityData() *types.CommonEntityData {
    match.EntityData.YFilter = match.YFilter
    match.EntityData.YangName = "match"
    match.EntityData.BundleName = "cisco_ios_xr"
    match.EntityData.ParentYangName = "class-map"
    match.EntityData.SegmentPath = "match"
    match.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/" + match.EntityData.SegmentPath
    match.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    match.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    match.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    match.EntityData.Children = types.NewOrderedMap()
    match.EntityData.Children.Append("destination-address-ipv4", types.YChild{"DestinationAddressIpv4", nil})
    for i := range match.DestinationAddressIpv4 {
        match.EntityData.Children.Append(types.GetSegmentPath(match.DestinationAddressIpv4[i]), types.YChild{"DestinationAddressIpv4", match.DestinationAddressIpv4[i]})
    }
    match.EntityData.Children.Append("destination-address-ipv6", types.YChild{"DestinationAddressIpv6", nil})
    for i := range match.DestinationAddressIpv6 {
        match.EntityData.Children.Append(types.GetSegmentPath(match.DestinationAddressIpv6[i]), types.YChild{"DestinationAddressIpv6", match.DestinationAddressIpv6[i]})
    }
    match.EntityData.Children.Append("source-address-ipv4", types.YChild{"SourceAddressIpv4", nil})
    for i := range match.SourceAddressIpv4 {
        match.EntityData.Children.Append(types.GetSegmentPath(match.SourceAddressIpv4[i]), types.YChild{"SourceAddressIpv4", match.SourceAddressIpv4[i]})
    }
    match.EntityData.Children.Append("source-address-ipv6", types.YChild{"SourceAddressIpv6", nil})
    for i := range match.SourceAddressIpv6 {
        match.EntityData.Children.Append(types.GetSegmentPath(match.SourceAddressIpv6[i]), types.YChild{"SourceAddressIpv6", match.SourceAddressIpv6[i]})
    }
    match.EntityData.Children.Append("dhcp-client-id", types.YChild{"DhcpClientId", nil})
    for i := range match.DhcpClientId {
        match.EntityData.Children.Append(types.GetSegmentPath(match.DhcpClientId[i]), types.YChild{"DhcpClientId", match.DhcpClientId[i]})
    }
    match.EntityData.Children.Append("dhcp-client-id-regex", types.YChild{"DhcpClientIdRegex", nil})
    for i := range match.DhcpClientIdRegex {
        match.EntityData.Children.Append(types.GetSegmentPath(match.DhcpClientIdRegex[i]), types.YChild{"DhcpClientIdRegex", match.DhcpClientIdRegex[i]})
    }
    match.EntityData.Children.Append("domain-name", types.YChild{"DomainName", nil})
    for i := range match.DomainName {
        match.EntityData.Children.Append(types.GetSegmentPath(match.DomainName[i]), types.YChild{"DomainName", match.DomainName[i]})
    }
    match.EntityData.Children.Append("domain-name-regex", types.YChild{"DomainNameRegex", nil})
    for i := range match.DomainNameRegex {
        match.EntityData.Children.Append(types.GetSegmentPath(match.DomainNameRegex[i]), types.YChild{"DomainNameRegex", match.DomainNameRegex[i]})
    }
    match.EntityData.Children.Append("flow", types.YChild{"Flow", &match.Flow})
    match.EntityData.Leafs = types.NewOrderedMap()
    match.EntityData.Leafs.Append("ipv4-dscp", types.YLeaf{"Ipv4Dscp", match.Ipv4Dscp})
    match.EntityData.Leafs.Append("ipv6-dscp", types.YLeaf{"Ipv6Dscp", match.Ipv6Dscp})
    match.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", match.Dscp})
    match.EntityData.Leafs.Append("ipv4-precedence", types.YLeaf{"Ipv4Precedence", match.Ipv4Precedence})
    match.EntityData.Leafs.Append("ipv6-precedence", types.YLeaf{"Ipv6Precedence", match.Ipv6Precedence})
    match.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", match.Precedence})
    match.EntityData.Leafs.Append("qos-group", types.YLeaf{"QosGroup", match.QosGroup})
    match.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", match.TrafficClass})
    match.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", match.Cos})
    match.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", match.InnerCos})
    match.EntityData.Leafs.Append("dei", types.YLeaf{"Dei", match.Dei})
    match.EntityData.Leafs.Append("dei-inner", types.YLeaf{"DeiInner", match.DeiInner})
    match.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", match.Protocol})
    match.EntityData.Leafs.Append("ipv4-acl", types.YLeaf{"Ipv4Acl", match.Ipv4Acl})
    match.EntityData.Leafs.Append("ipv6-acl", types.YLeaf{"Ipv6Acl", match.Ipv6Acl})
    match.EntityData.Leafs.Append("ethernet-services-acl", types.YLeaf{"EthernetServicesAcl", match.EthernetServicesAcl})
    match.EntityData.Leafs.Append("mpls-experimental-topmost", types.YLeaf{"MplsExperimentalTopmost", match.MplsExperimentalTopmost})
    match.EntityData.Leafs.Append("mpls-experimental-imposition", types.YLeaf{"MplsExperimentalImposition", match.MplsExperimentalImposition})
    match.EntityData.Leafs.Append("discard-class", types.YLeaf{"DiscardClass", match.DiscardClass})
    match.EntityData.Leafs.Append("ipv4-packet-length", types.YLeaf{"Ipv4PacketLength", match.Ipv4PacketLength})
    match.EntityData.Leafs.Append("ipv6-packet-length", types.YLeaf{"Ipv6PacketLength", match.Ipv6PacketLength})
    match.EntityData.Leafs.Append("packet-length", types.YLeaf{"PacketLength", match.PacketLength})
    match.EntityData.Leafs.Append("mpls-disposition-ipv4-access-list", types.YLeaf{"MplsDispositionIpv4AccessList", match.MplsDispositionIpv4AccessList})
    match.EntityData.Leafs.Append("mpls-disposition-ipv6-access-list", types.YLeaf{"MplsDispositionIpv6AccessList", match.MplsDispositionIpv6AccessList})
    match.EntityData.Leafs.Append("vlan", types.YLeaf{"Vlan", match.Vlan})
    match.EntityData.Leafs.Append("inner-vlan", types.YLeaf{"InnerVlan", match.InnerVlan})
    match.EntityData.Leafs.Append("flow-tag", types.YLeaf{"FlowTag", match.FlowTag})
    match.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", match.Ethertype})
    match.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", match.DestinationPort})
    match.EntityData.Leafs.Append("fragment-type", types.YLeaf{"FragmentType", match.FragmentType})
    match.EntityData.Leafs.Append("frame-relay-dlci", types.YLeaf{"FrameRelayDlci", match.FrameRelayDlci})
    match.EntityData.Leafs.Append("fr-de", types.YLeaf{"FrDe", match.FrDe})
    match.EntityData.Leafs.Append("icmpv4-code", types.YLeaf{"Icmpv4Code", match.Icmpv4Code})
    match.EntityData.Leafs.Append("icmpv4-type", types.YLeaf{"Icmpv4Type", match.Icmpv4Type})
    match.EntityData.Leafs.Append("icmpv6-code", types.YLeaf{"Icmpv6Code", match.Icmpv6Code})
    match.EntityData.Leafs.Append("icmpv6-type", types.YLeaf{"Icmpv6Type", match.Icmpv6Type})
    match.EntityData.Leafs.Append("source-port", types.YLeaf{"SourcePort", match.SourcePort})
    match.EntityData.Leafs.Append("tcp-flag", types.YLeaf{"TcpFlag", match.TcpFlag})
    match.EntityData.Leafs.Append("authen-status", types.YLeaf{"AuthenStatus", match.AuthenStatus})
    match.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", match.CircuitId})
    match.EntityData.Leafs.Append("circuit-id-regex", types.YLeaf{"CircuitIdRegex", match.CircuitIdRegex})
    match.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", match.RemoteId})
    match.EntityData.Leafs.Append("remote-id-regex", types.YLeaf{"RemoteIdRegex", match.RemoteIdRegex})
    match.EntityData.Leafs.Append("service-name", types.YLeaf{"ServiceName", match.ServiceName})
    match.EntityData.Leafs.Append("service-name-regex", types.YLeaf{"ServiceNameRegex", match.ServiceNameRegex})
    match.EntityData.Leafs.Append("timer", types.YLeaf{"Timer", match.Timer})
    match.EntityData.Leafs.Append("timer-regex", types.YLeaf{"TimerRegex", match.TimerRegex})
    match.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", match.UserName})
    match.EntityData.Leafs.Append("user-name-regex", types.YLeaf{"UserNameRegex", match.UserNameRegex})
    match.EntityData.Leafs.Append("source-mac", types.YLeaf{"SourceMac", match.SourceMac})
    match.EntityData.Leafs.Append("destination-mac", types.YLeaf{"DestinationMac", match.DestinationMac})
    match.EntityData.Leafs.Append("vpls-control", types.YLeaf{"VplsControl", match.VplsControl})
    match.EntityData.Leafs.Append("vpls-broadcast", types.YLeaf{"VplsBroadcast", match.VplsBroadcast})
    match.EntityData.Leafs.Append("vpls-multicast", types.YLeaf{"VplsMulticast", match.VplsMulticast})
    match.EntityData.Leafs.Append("vpls-known", types.YLeaf{"VplsKnown", match.VplsKnown})
    match.EntityData.Leafs.Append("vpls-unknown", types.YLeaf{"VplsUnknown", match.VplsUnknown})
    match.EntityData.Leafs.Append("atm-clp", types.YLeaf{"AtmClp", match.AtmClp})
    match.EntityData.Leafs.Append("atm-oam", types.YLeaf{"AtmOam", match.AtmOam})
    match.EntityData.Leafs.Append("cac-admit", types.YLeaf{"CacAdmit", match.CacAdmit})
    match.EntityData.Leafs.Append("cac-unadmit", types.YLeaf{"CacUnadmit", match.CacUnadmit})

    match.EntityData.YListKeys = []string {}

    return &(match.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4
// Match destination IPv4 address.
type PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv4 netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv4) GetEntityData() *types.CommonEntityData {
    destinationAddressIpv4.EntityData.YFilter = destinationAddressIpv4.YFilter
    destinationAddressIpv4.EntityData.YangName = "destination-address-ipv4"
    destinationAddressIpv4.EntityData.BundleName = "cisco_ios_xr"
    destinationAddressIpv4.EntityData.ParentYangName = "match"
    destinationAddressIpv4.EntityData.SegmentPath = "destination-address-ipv4" + types.AddKeyToken(destinationAddressIpv4.Address, "address") + types.AddKeyToken(destinationAddressIpv4.Netmask, "netmask")
    destinationAddressIpv4.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/" + destinationAddressIpv4.EntityData.SegmentPath
    destinationAddressIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationAddressIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationAddressIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationAddressIpv4.EntityData.Children = types.NewOrderedMap()
    destinationAddressIpv4.EntityData.Leafs = types.NewOrderedMap()
    destinationAddressIpv4.EntityData.Leafs.Append("address", types.YLeaf{"Address", destinationAddressIpv4.Address})
    destinationAddressIpv4.EntityData.Leafs.Append("netmask", types.YLeaf{"Netmask", destinationAddressIpv4.Netmask})

    destinationAddressIpv4.EntityData.YListKeys = []string {"Address", "Netmask"}

    return &(destinationAddressIpv4.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6
// Match destination IPv6 address.
type PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_DestinationAddressIpv6) GetEntityData() *types.CommonEntityData {
    destinationAddressIpv6.EntityData.YFilter = destinationAddressIpv6.YFilter
    destinationAddressIpv6.EntityData.YangName = "destination-address-ipv6"
    destinationAddressIpv6.EntityData.BundleName = "cisco_ios_xr"
    destinationAddressIpv6.EntityData.ParentYangName = "match"
    destinationAddressIpv6.EntityData.SegmentPath = "destination-address-ipv6" + types.AddKeyToken(destinationAddressIpv6.Address, "address") + types.AddKeyToken(destinationAddressIpv6.PrefixLength, "prefix-length")
    destinationAddressIpv6.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/" + destinationAddressIpv6.EntityData.SegmentPath
    destinationAddressIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationAddressIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationAddressIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationAddressIpv6.EntityData.Children = types.NewOrderedMap()
    destinationAddressIpv6.EntityData.Leafs = types.NewOrderedMap()
    destinationAddressIpv6.EntityData.Leafs.Append("address", types.YLeaf{"Address", destinationAddressIpv6.Address})
    destinationAddressIpv6.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", destinationAddressIpv6.PrefixLength})

    destinationAddressIpv6.EntityData.YListKeys = []string {"Address", "PrefixLength"}

    return &(destinationAddressIpv6.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4
// Match source IPv4 address.
type PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv4 netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv4) GetEntityData() *types.CommonEntityData {
    sourceAddressIpv4.EntityData.YFilter = sourceAddressIpv4.YFilter
    sourceAddressIpv4.EntityData.YangName = "source-address-ipv4"
    sourceAddressIpv4.EntityData.BundleName = "cisco_ios_xr"
    sourceAddressIpv4.EntityData.ParentYangName = "match"
    sourceAddressIpv4.EntityData.SegmentPath = "source-address-ipv4" + types.AddKeyToken(sourceAddressIpv4.Address, "address") + types.AddKeyToken(sourceAddressIpv4.Netmask, "netmask")
    sourceAddressIpv4.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/" + sourceAddressIpv4.EntityData.SegmentPath
    sourceAddressIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddressIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddressIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddressIpv4.EntityData.Children = types.NewOrderedMap()
    sourceAddressIpv4.EntityData.Leafs = types.NewOrderedMap()
    sourceAddressIpv4.EntityData.Leafs.Append("address", types.YLeaf{"Address", sourceAddressIpv4.Address})
    sourceAddressIpv4.EntityData.Leafs.Append("netmask", types.YLeaf{"Netmask", sourceAddressIpv4.Netmask})

    sourceAddressIpv4.EntityData.YListKeys = []string {"Address", "Netmask"}

    return &(sourceAddressIpv4.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6
// Match source IPv6 address.
type PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_Match_SourceAddressIpv6) GetEntityData() *types.CommonEntityData {
    sourceAddressIpv6.EntityData.YFilter = sourceAddressIpv6.YFilter
    sourceAddressIpv6.EntityData.YangName = "source-address-ipv6"
    sourceAddressIpv6.EntityData.BundleName = "cisco_ios_xr"
    sourceAddressIpv6.EntityData.ParentYangName = "match"
    sourceAddressIpv6.EntityData.SegmentPath = "source-address-ipv6" + types.AddKeyToken(sourceAddressIpv6.Address, "address") + types.AddKeyToken(sourceAddressIpv6.PrefixLength, "prefix-length")
    sourceAddressIpv6.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/" + sourceAddressIpv6.EntityData.SegmentPath
    sourceAddressIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddressIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddressIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddressIpv6.EntityData.Children = types.NewOrderedMap()
    sourceAddressIpv6.EntityData.Leafs = types.NewOrderedMap()
    sourceAddressIpv6.EntityData.Leafs.Append("address", types.YLeaf{"Address", sourceAddressIpv6.Address})
    sourceAddressIpv6.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", sourceAddressIpv6.PrefixLength})

    sourceAddressIpv6.EntityData.YListKeys = []string {"Address", "PrefixLength"}

    return &(sourceAddressIpv6.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId
// Match dhcp client ID.
type PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Dhcp client Id. The type is string with length:
    // 1..32.
    Value interface{}

    // This attribute is a key. Dhcp client id Ascii/Hex. The type is string with
    // pattern: (none)|(ascii)|(hex).
    Flag interface{}
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientId) GetEntityData() *types.CommonEntityData {
    dhcpClientId.EntityData.YFilter = dhcpClientId.YFilter
    dhcpClientId.EntityData.YangName = "dhcp-client-id"
    dhcpClientId.EntityData.BundleName = "cisco_ios_xr"
    dhcpClientId.EntityData.ParentYangName = "match"
    dhcpClientId.EntityData.SegmentPath = "dhcp-client-id" + types.AddKeyToken(dhcpClientId.Value, "value") + types.AddKeyToken(dhcpClientId.Flag, "flag")
    dhcpClientId.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/" + dhcpClientId.EntityData.SegmentPath
    dhcpClientId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpClientId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpClientId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpClientId.EntityData.Children = types.NewOrderedMap()
    dhcpClientId.EntityData.Leafs = types.NewOrderedMap()
    dhcpClientId.EntityData.Leafs.Append("value", types.YLeaf{"Value", dhcpClientId.Value})
    dhcpClientId.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", dhcpClientId.Flag})

    dhcpClientId.EntityData.YListKeys = []string {"Value", "Flag"}

    return &(dhcpClientId.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex
// Match dhcp client id regex.
type PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Dhcp client id regular expression. The type is
    // string with length: 1..32.
    Value interface{}

    // This attribute is a key. Dhcp client Id regex Ascii/Hex. The type is string
    // with pattern: (none)|(ascii)|(hex).
    Flag interface{}
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_Match_DhcpClientIdRegex) GetEntityData() *types.CommonEntityData {
    dhcpClientIdRegex.EntityData.YFilter = dhcpClientIdRegex.YFilter
    dhcpClientIdRegex.EntityData.YangName = "dhcp-client-id-regex"
    dhcpClientIdRegex.EntityData.BundleName = "cisco_ios_xr"
    dhcpClientIdRegex.EntityData.ParentYangName = "match"
    dhcpClientIdRegex.EntityData.SegmentPath = "dhcp-client-id-regex" + types.AddKeyToken(dhcpClientIdRegex.Value, "value") + types.AddKeyToken(dhcpClientIdRegex.Flag, "flag")
    dhcpClientIdRegex.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/" + dhcpClientIdRegex.EntityData.SegmentPath
    dhcpClientIdRegex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpClientIdRegex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpClientIdRegex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpClientIdRegex.EntityData.Children = types.NewOrderedMap()
    dhcpClientIdRegex.EntityData.Leafs = types.NewOrderedMap()
    dhcpClientIdRegex.EntityData.Leafs.Append("value", types.YLeaf{"Value", dhcpClientIdRegex.Value})
    dhcpClientIdRegex.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", dhcpClientIdRegex.Flag})

    dhcpClientIdRegex.EntityData.YListKeys = []string {"Value", "Flag"}

    return &(dhcpClientIdRegex.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_DomainName
// Match domain name.
type PolicyManager_ClassMaps_ClassMap_Match_DomainName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Domain name or regular expression. The type is
    // string with length: 1..32.
    Name interface{}

    // This attribute is a key. Domain-format name. The type is string with
    // length: 1..32.
    Format interface{}
}

func (domainName *PolicyManager_ClassMaps_ClassMap_Match_DomainName) GetEntityData() *types.CommonEntityData {
    domainName.EntityData.YFilter = domainName.YFilter
    domainName.EntityData.YangName = "domain-name"
    domainName.EntityData.BundleName = "cisco_ios_xr"
    domainName.EntityData.ParentYangName = "match"
    domainName.EntityData.SegmentPath = "domain-name" + types.AddKeyToken(domainName.Name, "name") + types.AddKeyToken(domainName.Format, "format")
    domainName.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/" + domainName.EntityData.SegmentPath
    domainName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domainName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domainName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domainName.EntityData.Children = types.NewOrderedMap()
    domainName.EntityData.Leafs = types.NewOrderedMap()
    domainName.EntityData.Leafs.Append("name", types.YLeaf{"Name", domainName.Name})
    domainName.EntityData.Leafs.Append("format", types.YLeaf{"Format", domainName.Format})

    domainName.EntityData.YListKeys = []string {"Name", "Format"}

    return &(domainName.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex
// Match domain name.
type PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Domain name or regular expression. The type is
    // string with length: 1..32.
    Regex interface{}

    // This attribute is a key. Domain-format name. The type is string with
    // length: 1..32.
    Format interface{}
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_Match_DomainNameRegex) GetEntityData() *types.CommonEntityData {
    domainNameRegex.EntityData.YFilter = domainNameRegex.YFilter
    domainNameRegex.EntityData.YangName = "domain-name-regex"
    domainNameRegex.EntityData.BundleName = "cisco_ios_xr"
    domainNameRegex.EntityData.ParentYangName = "match"
    domainNameRegex.EntityData.SegmentPath = "domain-name-regex" + types.AddKeyToken(domainNameRegex.Regex, "regex") + types.AddKeyToken(domainNameRegex.Format, "format")
    domainNameRegex.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/" + domainNameRegex.EntityData.SegmentPath
    domainNameRegex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domainNameRegex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domainNameRegex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domainNameRegex.EntityData.Children = types.NewOrderedMap()
    domainNameRegex.EntityData.Leafs = types.NewOrderedMap()
    domainNameRegex.EntityData.Leafs.Append("regex", types.YLeaf{"Regex", domainNameRegex.Regex})
    domainNameRegex.EntityData.Leafs.Append("format", types.YLeaf{"Format", domainNameRegex.Format})

    domainNameRegex.EntityData.YListKeys = []string {"Regex", "Format"}

    return &(domainNameRegex.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_Flow
// Match flow.
type PolicyManager_ClassMaps_ClassMap_Match_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the flow-key parameters. The type is slice of string with
    // pattern: (SourceIP)|(DestinationIP)|(5Tuple).
    FlowKey []interface{}

    // Configure the flow-cache parameters.
    FlowCache PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache
}

func (flow *PolicyManager_ClassMaps_ClassMap_Match_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xr"
    flow.EntityData.ParentYangName = "match"
    flow.EntityData.SegmentPath = "flow"
    flow.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/" + flow.EntityData.SegmentPath
    flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flow.EntityData.Children = types.NewOrderedMap()
    flow.EntityData.Children.Append("flow-cache", types.YChild{"FlowCache", &flow.FlowCache})
    flow.EntityData.Leafs = types.NewOrderedMap()
    flow.EntityData.Leafs.Append("flow-key", types.YLeaf{"FlowKey", flow.FlowKey})

    flow.EntityData.YListKeys = []string {}

    return &(flow.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache
// Configure the flow-cache parameters
type PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum time of inactivity for a flow. The type is one of the following
    // types: int with range: 10..2550, or string with pattern: (None)|(none).
    IdleTimeout interface{}
}

func (flowCache *PolicyManager_ClassMaps_ClassMap_Match_Flow_FlowCache) GetEntityData() *types.CommonEntityData {
    flowCache.EntityData.YFilter = flowCache.YFilter
    flowCache.EntityData.YangName = "flow-cache"
    flowCache.EntityData.BundleName = "cisco_ios_xr"
    flowCache.EntityData.ParentYangName = "flow"
    flowCache.EntityData.SegmentPath = "flow-cache"
    flowCache.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match/flow/" + flowCache.EntityData.SegmentPath
    flowCache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowCache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowCache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowCache.EntityData.Children = types.NewOrderedMap()
    flowCache.EntityData.Leafs = types.NewOrderedMap()
    flowCache.EntityData.Leafs.Append("idle-timeout", types.YLeaf{"IdleTimeout", flowCache.IdleTimeout})

    flowCache.EntityData.YListKeys = []string {}

    return &(flowCache.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot
// Match not rules.
type PolicyManager_ClassMaps_ClassMap_MatchNot struct {
    EntityData types.CommonEntityData
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
    DestinationAddressIpv4 []*PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4

    // Match destination IPv6 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6.
    DestinationAddressIpv6 []*PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6

    // Match source IPv4 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4.
    SourceAddressIpv4 []*PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4

    // Match source IPv6 address. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6.
    SourceAddressIpv6 []*PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6

    // Match dhcp client ID. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId.
    DhcpClientId []*PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId

    // Match dhcp client id regex. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex.
    DhcpClientIdRegex []*PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex

    // Match domain name. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName.
    DomainName []*PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName

    // Match domain name. The type is slice of
    // PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex.
    DomainNameRegex []*PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex

    // Match flow.
    Flow PolicyManager_ClassMaps_ClassMap_MatchNot_Flow
}

func (matchNot *PolicyManager_ClassMaps_ClassMap_MatchNot) GetEntityData() *types.CommonEntityData {
    matchNot.EntityData.YFilter = matchNot.YFilter
    matchNot.EntityData.YangName = "match-not"
    matchNot.EntityData.BundleName = "cisco_ios_xr"
    matchNot.EntityData.ParentYangName = "class-map"
    matchNot.EntityData.SegmentPath = "match-not"
    matchNot.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/" + matchNot.EntityData.SegmentPath
    matchNot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    matchNot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    matchNot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    matchNot.EntityData.Children = types.NewOrderedMap()
    matchNot.EntityData.Children.Append("destination-address-ipv4", types.YChild{"DestinationAddressIpv4", nil})
    for i := range matchNot.DestinationAddressIpv4 {
        matchNot.EntityData.Children.Append(types.GetSegmentPath(matchNot.DestinationAddressIpv4[i]), types.YChild{"DestinationAddressIpv4", matchNot.DestinationAddressIpv4[i]})
    }
    matchNot.EntityData.Children.Append("destination-address-ipv6", types.YChild{"DestinationAddressIpv6", nil})
    for i := range matchNot.DestinationAddressIpv6 {
        matchNot.EntityData.Children.Append(types.GetSegmentPath(matchNot.DestinationAddressIpv6[i]), types.YChild{"DestinationAddressIpv6", matchNot.DestinationAddressIpv6[i]})
    }
    matchNot.EntityData.Children.Append("source-address-ipv4", types.YChild{"SourceAddressIpv4", nil})
    for i := range matchNot.SourceAddressIpv4 {
        matchNot.EntityData.Children.Append(types.GetSegmentPath(matchNot.SourceAddressIpv4[i]), types.YChild{"SourceAddressIpv4", matchNot.SourceAddressIpv4[i]})
    }
    matchNot.EntityData.Children.Append("source-address-ipv6", types.YChild{"SourceAddressIpv6", nil})
    for i := range matchNot.SourceAddressIpv6 {
        matchNot.EntityData.Children.Append(types.GetSegmentPath(matchNot.SourceAddressIpv6[i]), types.YChild{"SourceAddressIpv6", matchNot.SourceAddressIpv6[i]})
    }
    matchNot.EntityData.Children.Append("dhcp-client-id", types.YChild{"DhcpClientId", nil})
    for i := range matchNot.DhcpClientId {
        matchNot.EntityData.Children.Append(types.GetSegmentPath(matchNot.DhcpClientId[i]), types.YChild{"DhcpClientId", matchNot.DhcpClientId[i]})
    }
    matchNot.EntityData.Children.Append("dhcp-client-id-regex", types.YChild{"DhcpClientIdRegex", nil})
    for i := range matchNot.DhcpClientIdRegex {
        matchNot.EntityData.Children.Append(types.GetSegmentPath(matchNot.DhcpClientIdRegex[i]), types.YChild{"DhcpClientIdRegex", matchNot.DhcpClientIdRegex[i]})
    }
    matchNot.EntityData.Children.Append("domain-name", types.YChild{"DomainName", nil})
    for i := range matchNot.DomainName {
        matchNot.EntityData.Children.Append(types.GetSegmentPath(matchNot.DomainName[i]), types.YChild{"DomainName", matchNot.DomainName[i]})
    }
    matchNot.EntityData.Children.Append("domain-name-regex", types.YChild{"DomainNameRegex", nil})
    for i := range matchNot.DomainNameRegex {
        matchNot.EntityData.Children.Append(types.GetSegmentPath(matchNot.DomainNameRegex[i]), types.YChild{"DomainNameRegex", matchNot.DomainNameRegex[i]})
    }
    matchNot.EntityData.Children.Append("flow", types.YChild{"Flow", &matchNot.Flow})
    matchNot.EntityData.Leafs = types.NewOrderedMap()
    matchNot.EntityData.Leafs.Append("ipv4-dscp", types.YLeaf{"Ipv4Dscp", matchNot.Ipv4Dscp})
    matchNot.EntityData.Leafs.Append("ipv6-dscp", types.YLeaf{"Ipv6Dscp", matchNot.Ipv6Dscp})
    matchNot.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", matchNot.Dscp})
    matchNot.EntityData.Leafs.Append("ipv4-precedence", types.YLeaf{"Ipv4Precedence", matchNot.Ipv4Precedence})
    matchNot.EntityData.Leafs.Append("ipv6-precedence", types.YLeaf{"Ipv6Precedence", matchNot.Ipv6Precedence})
    matchNot.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", matchNot.Precedence})
    matchNot.EntityData.Leafs.Append("qos-group", types.YLeaf{"QosGroup", matchNot.QosGroup})
    matchNot.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", matchNot.TrafficClass})
    matchNot.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", matchNot.Cos})
    matchNot.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", matchNot.InnerCos})
    matchNot.EntityData.Leafs.Append("dei", types.YLeaf{"Dei", matchNot.Dei})
    matchNot.EntityData.Leafs.Append("dei-inner", types.YLeaf{"DeiInner", matchNot.DeiInner})
    matchNot.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", matchNot.Protocol})
    matchNot.EntityData.Leafs.Append("ipv4-acl", types.YLeaf{"Ipv4Acl", matchNot.Ipv4Acl})
    matchNot.EntityData.Leafs.Append("ipv6-acl", types.YLeaf{"Ipv6Acl", matchNot.Ipv6Acl})
    matchNot.EntityData.Leafs.Append("ethernet-services-acl", types.YLeaf{"EthernetServicesAcl", matchNot.EthernetServicesAcl})
    matchNot.EntityData.Leafs.Append("mpls-experimental-topmost", types.YLeaf{"MplsExperimentalTopmost", matchNot.MplsExperimentalTopmost})
    matchNot.EntityData.Leafs.Append("mpls-experimental-imposition", types.YLeaf{"MplsExperimentalImposition", matchNot.MplsExperimentalImposition})
    matchNot.EntityData.Leafs.Append("discard-class", types.YLeaf{"DiscardClass", matchNot.DiscardClass})
    matchNot.EntityData.Leafs.Append("ipv4-packet-length", types.YLeaf{"Ipv4PacketLength", matchNot.Ipv4PacketLength})
    matchNot.EntityData.Leafs.Append("ipv6-packet-length", types.YLeaf{"Ipv6PacketLength", matchNot.Ipv6PacketLength})
    matchNot.EntityData.Leafs.Append("packet-length", types.YLeaf{"PacketLength", matchNot.PacketLength})
    matchNot.EntityData.Leafs.Append("mpls-disposition-ipv4-access-list", types.YLeaf{"MplsDispositionIpv4AccessList", matchNot.MplsDispositionIpv4AccessList})
    matchNot.EntityData.Leafs.Append("mpls-disposition-ipv6-access-list", types.YLeaf{"MplsDispositionIpv6AccessList", matchNot.MplsDispositionIpv6AccessList})
    matchNot.EntityData.Leafs.Append("vlan", types.YLeaf{"Vlan", matchNot.Vlan})
    matchNot.EntityData.Leafs.Append("inner-vlan", types.YLeaf{"InnerVlan", matchNot.InnerVlan})
    matchNot.EntityData.Leafs.Append("flow-tag", types.YLeaf{"FlowTag", matchNot.FlowTag})
    matchNot.EntityData.Leafs.Append("ethertype", types.YLeaf{"Ethertype", matchNot.Ethertype})
    matchNot.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", matchNot.DestinationPort})
    matchNot.EntityData.Leafs.Append("fragment-type", types.YLeaf{"FragmentType", matchNot.FragmentType})
    matchNot.EntityData.Leafs.Append("frame-relay-dlci", types.YLeaf{"FrameRelayDlci", matchNot.FrameRelayDlci})
    matchNot.EntityData.Leafs.Append("fr-de", types.YLeaf{"FrDe", matchNot.FrDe})
    matchNot.EntityData.Leafs.Append("icmpv4-code", types.YLeaf{"Icmpv4Code", matchNot.Icmpv4Code})
    matchNot.EntityData.Leafs.Append("icmpv4-type", types.YLeaf{"Icmpv4Type", matchNot.Icmpv4Type})
    matchNot.EntityData.Leafs.Append("icmpv6-code", types.YLeaf{"Icmpv6Code", matchNot.Icmpv6Code})
    matchNot.EntityData.Leafs.Append("icmpv6-type", types.YLeaf{"Icmpv6Type", matchNot.Icmpv6Type})
    matchNot.EntityData.Leafs.Append("source-port", types.YLeaf{"SourcePort", matchNot.SourcePort})
    matchNot.EntityData.Leafs.Append("tcp-flag", types.YLeaf{"TcpFlag", matchNot.TcpFlag})
    matchNot.EntityData.Leafs.Append("authen-status", types.YLeaf{"AuthenStatus", matchNot.AuthenStatus})
    matchNot.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", matchNot.CircuitId})
    matchNot.EntityData.Leafs.Append("circuit-id-regex", types.YLeaf{"CircuitIdRegex", matchNot.CircuitIdRegex})
    matchNot.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", matchNot.RemoteId})
    matchNot.EntityData.Leafs.Append("remote-id-regex", types.YLeaf{"RemoteIdRegex", matchNot.RemoteIdRegex})
    matchNot.EntityData.Leafs.Append("service-name", types.YLeaf{"ServiceName", matchNot.ServiceName})
    matchNot.EntityData.Leafs.Append("service-name-regex", types.YLeaf{"ServiceNameRegex", matchNot.ServiceNameRegex})
    matchNot.EntityData.Leafs.Append("timer", types.YLeaf{"Timer", matchNot.Timer})
    matchNot.EntityData.Leafs.Append("timer-regex", types.YLeaf{"TimerRegex", matchNot.TimerRegex})
    matchNot.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", matchNot.UserName})
    matchNot.EntityData.Leafs.Append("user-name-regex", types.YLeaf{"UserNameRegex", matchNot.UserNameRegex})
    matchNot.EntityData.Leafs.Append("source-mac", types.YLeaf{"SourceMac", matchNot.SourceMac})
    matchNot.EntityData.Leafs.Append("destination-mac", types.YLeaf{"DestinationMac", matchNot.DestinationMac})
    matchNot.EntityData.Leafs.Append("vpls-control", types.YLeaf{"VplsControl", matchNot.VplsControl})
    matchNot.EntityData.Leafs.Append("vpls-broadcast", types.YLeaf{"VplsBroadcast", matchNot.VplsBroadcast})
    matchNot.EntityData.Leafs.Append("vpls-multicast", types.YLeaf{"VplsMulticast", matchNot.VplsMulticast})
    matchNot.EntityData.Leafs.Append("vpls-known", types.YLeaf{"VplsKnown", matchNot.VplsKnown})
    matchNot.EntityData.Leafs.Append("vpls-unknown", types.YLeaf{"VplsUnknown", matchNot.VplsUnknown})

    matchNot.EntityData.YListKeys = []string {}

    return &(matchNot.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4
// Match destination IPv4 address.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv4 netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}
}

func (destinationAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv4) GetEntityData() *types.CommonEntityData {
    destinationAddressIpv4.EntityData.YFilter = destinationAddressIpv4.YFilter
    destinationAddressIpv4.EntityData.YangName = "destination-address-ipv4"
    destinationAddressIpv4.EntityData.BundleName = "cisco_ios_xr"
    destinationAddressIpv4.EntityData.ParentYangName = "match-not"
    destinationAddressIpv4.EntityData.SegmentPath = "destination-address-ipv4" + types.AddKeyToken(destinationAddressIpv4.Address, "address") + types.AddKeyToken(destinationAddressIpv4.Netmask, "netmask")
    destinationAddressIpv4.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match-not/" + destinationAddressIpv4.EntityData.SegmentPath
    destinationAddressIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationAddressIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationAddressIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationAddressIpv4.EntityData.Children = types.NewOrderedMap()
    destinationAddressIpv4.EntityData.Leafs = types.NewOrderedMap()
    destinationAddressIpv4.EntityData.Leafs.Append("address", types.YLeaf{"Address", destinationAddressIpv4.Address})
    destinationAddressIpv4.EntityData.Leafs.Append("netmask", types.YLeaf{"Netmask", destinationAddressIpv4.Netmask})

    destinationAddressIpv4.EntityData.YListKeys = []string {"Address", "Netmask"}

    return &(destinationAddressIpv4.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6
// Match destination IPv6 address.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}
}

func (destinationAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_DestinationAddressIpv6) GetEntityData() *types.CommonEntityData {
    destinationAddressIpv6.EntityData.YFilter = destinationAddressIpv6.YFilter
    destinationAddressIpv6.EntityData.YangName = "destination-address-ipv6"
    destinationAddressIpv6.EntityData.BundleName = "cisco_ios_xr"
    destinationAddressIpv6.EntityData.ParentYangName = "match-not"
    destinationAddressIpv6.EntityData.SegmentPath = "destination-address-ipv6" + types.AddKeyToken(destinationAddressIpv6.Address, "address") + types.AddKeyToken(destinationAddressIpv6.PrefixLength, "prefix-length")
    destinationAddressIpv6.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match-not/" + destinationAddressIpv6.EntityData.SegmentPath
    destinationAddressIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationAddressIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationAddressIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationAddressIpv6.EntityData.Children = types.NewOrderedMap()
    destinationAddressIpv6.EntityData.Leafs = types.NewOrderedMap()
    destinationAddressIpv6.EntityData.Leafs.Append("address", types.YLeaf{"Address", destinationAddressIpv6.Address})
    destinationAddressIpv6.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", destinationAddressIpv6.PrefixLength})

    destinationAddressIpv6.EntityData.YListKeys = []string {"Address", "PrefixLength"}

    return &(destinationAddressIpv6.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4
// Match source IPv4 address.
type PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv4 netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}
}

func (sourceAddressIpv4 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv4) GetEntityData() *types.CommonEntityData {
    sourceAddressIpv4.EntityData.YFilter = sourceAddressIpv4.YFilter
    sourceAddressIpv4.EntityData.YangName = "source-address-ipv4"
    sourceAddressIpv4.EntityData.BundleName = "cisco_ios_xr"
    sourceAddressIpv4.EntityData.ParentYangName = "match-not"
    sourceAddressIpv4.EntityData.SegmentPath = "source-address-ipv4" + types.AddKeyToken(sourceAddressIpv4.Address, "address") + types.AddKeyToken(sourceAddressIpv4.Netmask, "netmask")
    sourceAddressIpv4.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match-not/" + sourceAddressIpv4.EntityData.SegmentPath
    sourceAddressIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddressIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddressIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddressIpv4.EntityData.Children = types.NewOrderedMap()
    sourceAddressIpv4.EntityData.Leafs = types.NewOrderedMap()
    sourceAddressIpv4.EntityData.Leafs.Append("address", types.YLeaf{"Address", sourceAddressIpv4.Address})
    sourceAddressIpv4.EntityData.Leafs.Append("netmask", types.YLeaf{"Netmask", sourceAddressIpv4.Netmask})

    sourceAddressIpv4.EntityData.YListKeys = []string {"Address", "Netmask"}

    return &(sourceAddressIpv4.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6
// Match source IPv6 address.
type PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. IPv6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}
}

func (sourceAddressIpv6 *PolicyManager_ClassMaps_ClassMap_MatchNot_SourceAddressIpv6) GetEntityData() *types.CommonEntityData {
    sourceAddressIpv6.EntityData.YFilter = sourceAddressIpv6.YFilter
    sourceAddressIpv6.EntityData.YangName = "source-address-ipv6"
    sourceAddressIpv6.EntityData.BundleName = "cisco_ios_xr"
    sourceAddressIpv6.EntityData.ParentYangName = "match-not"
    sourceAddressIpv6.EntityData.SegmentPath = "source-address-ipv6" + types.AddKeyToken(sourceAddressIpv6.Address, "address") + types.AddKeyToken(sourceAddressIpv6.PrefixLength, "prefix-length")
    sourceAddressIpv6.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match-not/" + sourceAddressIpv6.EntityData.SegmentPath
    sourceAddressIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddressIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddressIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddressIpv6.EntityData.Children = types.NewOrderedMap()
    sourceAddressIpv6.EntityData.Leafs = types.NewOrderedMap()
    sourceAddressIpv6.EntityData.Leafs.Append("address", types.YLeaf{"Address", sourceAddressIpv6.Address})
    sourceAddressIpv6.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", sourceAddressIpv6.PrefixLength})

    sourceAddressIpv6.EntityData.YListKeys = []string {"Address", "PrefixLength"}

    return &(sourceAddressIpv6.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId
// Match dhcp client ID.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Dhcp client Id. The type is string with length:
    // 1..32.
    Value interface{}

    // This attribute is a key. Dhcp client id Ascii/Hex. The type is string with
    // pattern: (none)|(ascii)|(hex).
    Flag interface{}
}

func (dhcpClientId *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientId) GetEntityData() *types.CommonEntityData {
    dhcpClientId.EntityData.YFilter = dhcpClientId.YFilter
    dhcpClientId.EntityData.YangName = "dhcp-client-id"
    dhcpClientId.EntityData.BundleName = "cisco_ios_xr"
    dhcpClientId.EntityData.ParentYangName = "match-not"
    dhcpClientId.EntityData.SegmentPath = "dhcp-client-id" + types.AddKeyToken(dhcpClientId.Value, "value") + types.AddKeyToken(dhcpClientId.Flag, "flag")
    dhcpClientId.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match-not/" + dhcpClientId.EntityData.SegmentPath
    dhcpClientId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpClientId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpClientId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpClientId.EntityData.Children = types.NewOrderedMap()
    dhcpClientId.EntityData.Leafs = types.NewOrderedMap()
    dhcpClientId.EntityData.Leafs.Append("value", types.YLeaf{"Value", dhcpClientId.Value})
    dhcpClientId.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", dhcpClientId.Flag})

    dhcpClientId.EntityData.YListKeys = []string {"Value", "Flag"}

    return &(dhcpClientId.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex
// Match dhcp client id regex.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Dhcp client id regular expression. The type is
    // string with length: 1..32.
    Value interface{}

    // This attribute is a key. Dhcp client Id regex Ascii/Hex. The type is string
    // with pattern: (none)|(ascii)|(hex).
    Flag interface{}
}

func (dhcpClientIdRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DhcpClientIdRegex) GetEntityData() *types.CommonEntityData {
    dhcpClientIdRegex.EntityData.YFilter = dhcpClientIdRegex.YFilter
    dhcpClientIdRegex.EntityData.YangName = "dhcp-client-id-regex"
    dhcpClientIdRegex.EntityData.BundleName = "cisco_ios_xr"
    dhcpClientIdRegex.EntityData.ParentYangName = "match-not"
    dhcpClientIdRegex.EntityData.SegmentPath = "dhcp-client-id-regex" + types.AddKeyToken(dhcpClientIdRegex.Value, "value") + types.AddKeyToken(dhcpClientIdRegex.Flag, "flag")
    dhcpClientIdRegex.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match-not/" + dhcpClientIdRegex.EntityData.SegmentPath
    dhcpClientIdRegex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpClientIdRegex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpClientIdRegex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpClientIdRegex.EntityData.Children = types.NewOrderedMap()
    dhcpClientIdRegex.EntityData.Leafs = types.NewOrderedMap()
    dhcpClientIdRegex.EntityData.Leafs.Append("value", types.YLeaf{"Value", dhcpClientIdRegex.Value})
    dhcpClientIdRegex.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", dhcpClientIdRegex.Flag})

    dhcpClientIdRegex.EntityData.YListKeys = []string {"Value", "Flag"}

    return &(dhcpClientIdRegex.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName
// Match domain name.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Domain name or regular expression. The type is
    // string with length: 1..32.
    Name interface{}

    // This attribute is a key. Domain-format name. The type is string with
    // length: 1..32.
    Format interface{}
}

func (domainName *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainName) GetEntityData() *types.CommonEntityData {
    domainName.EntityData.YFilter = domainName.YFilter
    domainName.EntityData.YangName = "domain-name"
    domainName.EntityData.BundleName = "cisco_ios_xr"
    domainName.EntityData.ParentYangName = "match-not"
    domainName.EntityData.SegmentPath = "domain-name" + types.AddKeyToken(domainName.Name, "name") + types.AddKeyToken(domainName.Format, "format")
    domainName.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match-not/" + domainName.EntityData.SegmentPath
    domainName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domainName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domainName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domainName.EntityData.Children = types.NewOrderedMap()
    domainName.EntityData.Leafs = types.NewOrderedMap()
    domainName.EntityData.Leafs.Append("name", types.YLeaf{"Name", domainName.Name})
    domainName.EntityData.Leafs.Append("format", types.YLeaf{"Format", domainName.Format})

    domainName.EntityData.YListKeys = []string {"Name", "Format"}

    return &(domainName.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex
// Match domain name.
type PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Domain name or regular expression. The type is
    // string with length: 1..32.
    Regex interface{}

    // This attribute is a key. Domain-format name. The type is string with
    // length: 1..32.
    Format interface{}
}

func (domainNameRegex *PolicyManager_ClassMaps_ClassMap_MatchNot_DomainNameRegex) GetEntityData() *types.CommonEntityData {
    domainNameRegex.EntityData.YFilter = domainNameRegex.YFilter
    domainNameRegex.EntityData.YangName = "domain-name-regex"
    domainNameRegex.EntityData.BundleName = "cisco_ios_xr"
    domainNameRegex.EntityData.ParentYangName = "match-not"
    domainNameRegex.EntityData.SegmentPath = "domain-name-regex" + types.AddKeyToken(domainNameRegex.Regex, "regex") + types.AddKeyToken(domainNameRegex.Format, "format")
    domainNameRegex.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match-not/" + domainNameRegex.EntityData.SegmentPath
    domainNameRegex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domainNameRegex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domainNameRegex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domainNameRegex.EntityData.Children = types.NewOrderedMap()
    domainNameRegex.EntityData.Leafs = types.NewOrderedMap()
    domainNameRegex.EntityData.Leafs.Append("regex", types.YLeaf{"Regex", domainNameRegex.Regex})
    domainNameRegex.EntityData.Leafs.Append("format", types.YLeaf{"Format", domainNameRegex.Format})

    domainNameRegex.EntityData.YListKeys = []string {"Regex", "Format"}

    return &(domainNameRegex.EntityData)
}

// PolicyManager_ClassMaps_ClassMap_MatchNot_Flow
// Match flow.
type PolicyManager_ClassMaps_ClassMap_MatchNot_Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the flow-tag parameters. The type is slice of interface{} with
    // range: 1..63.
    FlowTag []interface{}
}

func (flow *PolicyManager_ClassMaps_ClassMap_MatchNot_Flow) GetEntityData() *types.CommonEntityData {
    flow.EntityData.YFilter = flow.YFilter
    flow.EntityData.YangName = "flow"
    flow.EntityData.BundleName = "cisco_ios_xr"
    flow.EntityData.ParentYangName = "match-not"
    flow.EntityData.SegmentPath = "flow"
    flow.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/class-maps/class-map/match-not/" + flow.EntityData.SegmentPath
    flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flow.EntityData.Children = types.NewOrderedMap()
    flow.EntityData.Leafs = types.NewOrderedMap()
    flow.EntityData.Leafs.Append("flow-tag", types.YLeaf{"FlowTag", flow.FlowTag})

    flow.EntityData.YListKeys = []string {}

    return &(flow.EntityData)
}

// PolicyManager_PolicyMaps
// Policy-maps configuration.
type PolicyManager_PolicyMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy-map configuration. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap.
    PolicyMap []*PolicyManager_PolicyMaps_PolicyMap
}

func (policyMaps *PolicyManager_PolicyMaps) GetEntityData() *types.CommonEntityData {
    policyMaps.EntityData.YFilter = policyMaps.YFilter
    policyMaps.EntityData.YangName = "policy-maps"
    policyMaps.EntityData.BundleName = "cisco_ios_xr"
    policyMaps.EntityData.ParentYangName = "policy-manager"
    policyMaps.EntityData.SegmentPath = "policy-maps"
    policyMaps.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/" + policyMaps.EntityData.SegmentPath
    policyMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyMaps.EntityData.Children = types.NewOrderedMap()
    policyMaps.EntityData.Children.Append("policy-map", types.YChild{"PolicyMap", nil})
    for i := range policyMaps.PolicyMap {
        policyMaps.EntityData.Children.Append(types.GetSegmentPath(policyMaps.PolicyMap[i]), types.YChild{"PolicyMap", policyMaps.PolicyMap[i]})
    }
    policyMaps.EntityData.Leafs = types.NewOrderedMap()

    policyMaps.EntityData.YListKeys = []string {}

    return &(policyMaps.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap
// Policy-map configuration.
type PolicyManager_PolicyMaps_PolicyMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Type of policy-map. The type is PolicyMapType.
    Type interface{}

    // This attribute is a key. Name of policy-map. The type is string with
    // pattern: [a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}.
    Name interface{}

    // Description for this policy-map. The type is string.
    Description interface{}

    // Policy event. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap_Event.
    Event []*PolicyManager_PolicyMaps_PolicyMap_Event

    // Class-map rule. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule.
    PolicyMapRule []*PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule
}

func (policyMap *PolicyManager_PolicyMaps_PolicyMap) GetEntityData() *types.CommonEntityData {
    policyMap.EntityData.YFilter = policyMap.YFilter
    policyMap.EntityData.YangName = "policy-map"
    policyMap.EntityData.BundleName = "cisco_ios_xr"
    policyMap.EntityData.ParentYangName = "policy-maps"
    policyMap.EntityData.SegmentPath = "policy-map" + types.AddKeyToken(policyMap.Type, "type") + types.AddKeyToken(policyMap.Name, "name")
    policyMap.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/" + policyMap.EntityData.SegmentPath
    policyMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyMap.EntityData.Children = types.NewOrderedMap()
    policyMap.EntityData.Children.Append("event", types.YChild{"Event", nil})
    for i := range policyMap.Event {
        policyMap.EntityData.Children.Append(types.GetSegmentPath(policyMap.Event[i]), types.YChild{"Event", policyMap.Event[i]})
    }
    policyMap.EntityData.Children.Append("policy-map-rule", types.YChild{"PolicyMapRule", nil})
    for i := range policyMap.PolicyMapRule {
        policyMap.EntityData.Children.Append(types.GetSegmentPath(policyMap.PolicyMapRule[i]), types.YChild{"PolicyMapRule", policyMap.PolicyMapRule[i]})
    }
    policyMap.EntityData.Leafs = types.NewOrderedMap()
    policyMap.EntityData.Leafs.Append("type", types.YLeaf{"Type", policyMap.Type})
    policyMap.EntityData.Leafs.Append("name", types.YLeaf{"Name", policyMap.Name})
    policyMap.EntityData.Leafs.Append("description", types.YLeaf{"Description", policyMap.Description})

    policyMap.EntityData.YListKeys = []string {"Type", "Name"}

    return &(policyMap.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_Event
// Policy event.
type PolicyManager_PolicyMaps_PolicyMap_Event struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Event type. The type is EventType.
    EventType interface{}

    // Execute all the matched classes. The type is interface{}.
    EventModeMatchAll interface{}

    // Execute only the first matched class. The type is interface{}.
    EventModeMatchFirst interface{}

    // Class-map rule. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap_Event_Class.
    Class []*PolicyManager_PolicyMaps_PolicyMap_Event_Class
}

func (event *PolicyManager_PolicyMaps_PolicyMap_Event) GetEntityData() *types.CommonEntityData {
    event.EntityData.YFilter = event.YFilter
    event.EntityData.YangName = "event"
    event.EntityData.BundleName = "cisco_ios_xr"
    event.EntityData.ParentYangName = "policy-map"
    event.EntityData.SegmentPath = "event" + types.AddKeyToken(event.EventType, "event-type")
    event.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/" + event.EntityData.SegmentPath
    event.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    event.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    event.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    event.EntityData.Children = types.NewOrderedMap()
    event.EntityData.Children.Append("class", types.YChild{"Class", nil})
    for i := range event.Class {
        event.EntityData.Children.Append(types.GetSegmentPath(event.Class[i]), types.YChild{"Class", event.Class[i]})
    }
    event.EntityData.Leafs = types.NewOrderedMap()
    event.EntityData.Leafs.Append("event-type", types.YLeaf{"EventType", event.EventType})
    event.EntityData.Leafs.Append("event-mode-match-all", types.YLeaf{"EventModeMatchAll", event.EventModeMatchAll})
    event.EntityData.Leafs.Append("event-mode-match-first", types.YLeaf{"EventModeMatchFirst", event.EventModeMatchFirst})

    event.EntityData.YListKeys = []string {"EventType"}

    return &(event.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_Event_Class
// Class-map rule.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of class. The type is string with pattern:
    // [a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}.
    ClassName interface{}

    // This attribute is a key. Type of class. The type is PmapClassMapType.
    ClassType interface{}

    // Class execution strategy. The type is ExecutionStrategy.
    ClassExecutionStrategy interface{}

    // Action rule. The type is slice of
    // PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule.
    ActionRule []*PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule
}

func (class *PolicyManager_PolicyMaps_PolicyMap_Event_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "event"
    class.EntityData.SegmentPath = "class" + types.AddKeyToken(class.ClassName, "class-name") + types.AddKeyToken(class.ClassType, "class-type")
    class.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/event/" + class.EntityData.SegmentPath
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = types.NewOrderedMap()
    class.EntityData.Children.Append("action-rule", types.YChild{"ActionRule", nil})
    for i := range class.ActionRule {
        class.EntityData.Children.Append(types.GetSegmentPath(class.ActionRule[i]), types.YChild{"ActionRule", class.ActionRule[i]})
    }
    class.EntityData.Leafs = types.NewOrderedMap()
    class.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", class.ClassName})
    class.EntityData.Leafs.Append("class-type", types.YLeaf{"ClassType", class.ClassType})
    class.EntityData.Leafs.Append("class-execution-strategy", types.YLeaf{"ClassExecutionStrategy", class.ClassExecutionStrategy})

    class.EntityData.YListKeys = []string {"ClassName", "ClassType"}

    return &(class.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule
// Action rule.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (actionRule *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule) GetEntityData() *types.CommonEntityData {
    actionRule.EntityData.YFilter = actionRule.YFilter
    actionRule.EntityData.YangName = "action-rule"
    actionRule.EntityData.BundleName = "cisco_ios_xr"
    actionRule.EntityData.ParentYangName = "class"
    actionRule.EntityData.SegmentPath = "action-rule" + types.AddKeyToken(actionRule.ActionSequenceNumber, "action-sequence-number")
    actionRule.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/event/class/" + actionRule.EntityData.SegmentPath
    actionRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actionRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actionRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actionRule.EntityData.Children = types.NewOrderedMap()
    actionRule.EntityData.Children.Append("activate-dynamic-template", types.YChild{"ActivateDynamicTemplate", &actionRule.ActivateDynamicTemplate})
    actionRule.EntityData.Children.Append("authenticate", types.YChild{"Authenticate", &actionRule.Authenticate})
    actionRule.EntityData.Children.Append("authorize", types.YChild{"Authorize", &actionRule.Authorize})
    actionRule.EntityData.Children.Append("deactivate-dynamic-template", types.YChild{"DeactivateDynamicTemplate", &actionRule.DeactivateDynamicTemplate})
    actionRule.EntityData.Children.Append("set-timer", types.YChild{"SetTimer", &actionRule.SetTimer})
    actionRule.EntityData.Children.Append("stop-timer", types.YChild{"StopTimer", &actionRule.StopTimer})
    actionRule.EntityData.Leafs = types.NewOrderedMap()
    actionRule.EntityData.Leafs.Append("action-sequence-number", types.YLeaf{"ActionSequenceNumber", actionRule.ActionSequenceNumber})
    actionRule.EntityData.Leafs.Append("disconnect", types.YLeaf{"Disconnect", actionRule.Disconnect})
    actionRule.EntityData.Leafs.Append("monitor", types.YLeaf{"Monitor", actionRule.Monitor})

    actionRule.EntityData.YListKeys = []string {"ActionSequenceNumber"}

    return &(actionRule.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate
// Activate dynamic templates.
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Dynamic template name. The type is string. This attribute is mandatory.
    Name interface{}

    // Name of the AAA method list. The type is string.
    AaaList interface{}
}

func (activateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_ActivateDynamicTemplate) GetEntityData() *types.CommonEntityData {
    activateDynamicTemplate.EntityData.YFilter = activateDynamicTemplate.YFilter
    activateDynamicTemplate.EntityData.YangName = "activate-dynamic-template"
    activateDynamicTemplate.EntityData.BundleName = "cisco_ios_xr"
    activateDynamicTemplate.EntityData.ParentYangName = "action-rule"
    activateDynamicTemplate.EntityData.SegmentPath = "activate-dynamic-template"
    activateDynamicTemplate.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/event/class/action-rule/" + activateDynamicTemplate.EntityData.SegmentPath
    activateDynamicTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activateDynamicTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activateDynamicTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activateDynamicTemplate.EntityData.Children = types.NewOrderedMap()
    activateDynamicTemplate.EntityData.Leafs = types.NewOrderedMap()
    activateDynamicTemplate.EntityData.Leafs.Append("name", types.YLeaf{"Name", activateDynamicTemplate.Name})
    activateDynamicTemplate.EntityData.Leafs.Append("aaa-list", types.YLeaf{"AaaList", activateDynamicTemplate.AaaList})

    activateDynamicTemplate.EntityData.YListKeys = []string {}

    return &(activateDynamicTemplate.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate
// Authentication related configuration.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the AAA method list. The type is string.
    AaaList interface{}
}

func (authenticate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authenticate) GetEntityData() *types.CommonEntityData {
    authenticate.EntityData.YFilter = authenticate.YFilter
    authenticate.EntityData.YangName = "authenticate"
    authenticate.EntityData.BundleName = "cisco_ios_xr"
    authenticate.EntityData.ParentYangName = "action-rule"
    authenticate.EntityData.SegmentPath = "authenticate"
    authenticate.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/event/class/action-rule/" + authenticate.EntityData.SegmentPath
    authenticate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticate.EntityData.Children = types.NewOrderedMap()
    authenticate.EntityData.Leafs = types.NewOrderedMap()
    authenticate.EntityData.Leafs.Append("aaa-list", types.YLeaf{"AaaList", authenticate.AaaList})

    authenticate.EntityData.YListKeys = []string {}

    return &(authenticate.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize
// Authorize.
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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

func (authorize *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_Authorize) GetEntityData() *types.CommonEntityData {
    authorize.EntityData.YFilter = authorize.YFilter
    authorize.EntityData.YangName = "authorize"
    authorize.EntityData.BundleName = "cisco_ios_xr"
    authorize.EntityData.ParentYangName = "action-rule"
    authorize.EntityData.SegmentPath = "authorize"
    authorize.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/event/class/action-rule/" + authorize.EntityData.SegmentPath
    authorize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorize.EntityData.Children = types.NewOrderedMap()
    authorize.EntityData.Leafs = types.NewOrderedMap()
    authorize.EntityData.Leafs.Append("aaa-list", types.YLeaf{"AaaList", authorize.AaaList})
    authorize.EntityData.Leafs.Append("format", types.YLeaf{"Format", authorize.Format})
    authorize.EntityData.Leafs.Append("identifier", types.YLeaf{"Identifier", authorize.Identifier})
    authorize.EntityData.Leafs.Append("password", types.YLeaf{"Password", authorize.Password})

    authorize.EntityData.YListKeys = []string {}

    return &(authorize.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate
// Deactivate dynamic templates.
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Dynamic template name. The type is string. This attribute is mandatory.
    Name interface{}

    // Name of the AAA method list. The type is string.
    AaaList interface{}
}

func (deactivateDynamicTemplate *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_DeactivateDynamicTemplate) GetEntityData() *types.CommonEntityData {
    deactivateDynamicTemplate.EntityData.YFilter = deactivateDynamicTemplate.YFilter
    deactivateDynamicTemplate.EntityData.YangName = "deactivate-dynamic-template"
    deactivateDynamicTemplate.EntityData.BundleName = "cisco_ios_xr"
    deactivateDynamicTemplate.EntityData.ParentYangName = "action-rule"
    deactivateDynamicTemplate.EntityData.SegmentPath = "deactivate-dynamic-template"
    deactivateDynamicTemplate.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/event/class/action-rule/" + deactivateDynamicTemplate.EntityData.SegmentPath
    deactivateDynamicTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    deactivateDynamicTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    deactivateDynamicTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    deactivateDynamicTemplate.EntityData.Children = types.NewOrderedMap()
    deactivateDynamicTemplate.EntityData.Leafs = types.NewOrderedMap()
    deactivateDynamicTemplate.EntityData.Leafs.Append("name", types.YLeaf{"Name", deactivateDynamicTemplate.Name})
    deactivateDynamicTemplate.EntityData.Leafs.Append("aaa-list", types.YLeaf{"AaaList", deactivateDynamicTemplate.AaaList})

    deactivateDynamicTemplate.EntityData.YListKeys = []string {}

    return &(deactivateDynamicTemplate.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer
// Set a timer to execute a rule on its 
// expiry
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Name of the timer. The type is string. This attribute is mandatory.
    TimerName interface{}

    // Timer value in minutes. The type is interface{} with range: 0..4294967295.
    // This attribute is mandatory. Units are minutes.
    TimerValue interface{}
}

func (setTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_SetTimer) GetEntityData() *types.CommonEntityData {
    setTimer.EntityData.YFilter = setTimer.YFilter
    setTimer.EntityData.YangName = "set-timer"
    setTimer.EntityData.BundleName = "cisco_ios_xr"
    setTimer.EntityData.ParentYangName = "action-rule"
    setTimer.EntityData.SegmentPath = "set-timer"
    setTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/event/class/action-rule/" + setTimer.EntityData.SegmentPath
    setTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    setTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    setTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    setTimer.EntityData.Children = types.NewOrderedMap()
    setTimer.EntityData.Leafs = types.NewOrderedMap()
    setTimer.EntityData.Leafs.Append("timer-name", types.YLeaf{"TimerName", setTimer.TimerName})
    setTimer.EntityData.Leafs.Append("timer-value", types.YLeaf{"TimerValue", setTimer.TimerValue})

    setTimer.EntityData.YListKeys = []string {}

    return &(setTimer.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer
// Disable timer before it expires.
type PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the timer. The type is string.
    TimerName interface{}
}

func (stopTimer *PolicyManager_PolicyMaps_PolicyMap_Event_Class_ActionRule_StopTimer) GetEntityData() *types.CommonEntityData {
    stopTimer.EntityData.YFilter = stopTimer.YFilter
    stopTimer.EntityData.YangName = "stop-timer"
    stopTimer.EntityData.BundleName = "cisco_ios_xr"
    stopTimer.EntityData.ParentYangName = "action-rule"
    stopTimer.EntityData.SegmentPath = "stop-timer"
    stopTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/event/class/action-rule/" + stopTimer.EntityData.SegmentPath
    stopTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stopTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stopTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stopTimer.EntityData.Children = types.NewOrderedMap()
    stopTimer.EntityData.Leafs = types.NewOrderedMap()
    stopTimer.EntityData.Leafs.Append("timer-name", types.YLeaf{"TimerName", stopTimer.TimerName})

    stopTimer.EntityData.YListKeys = []string {}

    return &(stopTimer.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule
// Class-map rule.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    RandomDetect []*PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect

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

    // HTTP Enrichment action.
    HttpEnrichment PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_HttpEnrichment
}

func (policyMapRule *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule) GetEntityData() *types.CommonEntityData {
    policyMapRule.EntityData.YFilter = policyMapRule.YFilter
    policyMapRule.EntityData.YangName = "policy-map-rule"
    policyMapRule.EntityData.BundleName = "cisco_ios_xr"
    policyMapRule.EntityData.ParentYangName = "policy-map"
    policyMapRule.EntityData.SegmentPath = "policy-map-rule" + types.AddKeyToken(policyMapRule.ClassName, "class-name") + types.AddKeyToken(policyMapRule.ClassType, "class-type")
    policyMapRule.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/" + policyMapRule.EntityData.SegmentPath
    policyMapRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyMapRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyMapRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyMapRule.EntityData.Children = types.NewOrderedMap()
    policyMapRule.EntityData.Children.Append("shape", types.YChild{"Shape", &policyMapRule.Shape})
    policyMapRule.EntityData.Children.Append("min-bandwidth", types.YChild{"MinBandwidth", &policyMapRule.MinBandwidth})
    policyMapRule.EntityData.Children.Append("bandwidth-remaining", types.YChild{"BandwidthRemaining", &policyMapRule.BandwidthRemaining})
    policyMapRule.EntityData.Children.Append("queue-limit", types.YChild{"QueueLimit", &policyMapRule.QueueLimit})
    policyMapRule.EntityData.Children.Append("pfc", types.YChild{"Pfc", &policyMapRule.Pfc})
    policyMapRule.EntityData.Children.Append("random-detect", types.YChild{"RandomDetect", nil})
    for i := range policyMapRule.RandomDetect {
        policyMapRule.EntityData.Children.Append(types.GetSegmentPath(policyMapRule.RandomDetect[i]), types.YChild{"RandomDetect", policyMapRule.RandomDetect[i]})
    }
    policyMapRule.EntityData.Children.Append("set", types.YChild{"Set", &policyMapRule.Set})
    policyMapRule.EntityData.Children.Append("police", types.YChild{"Police", &policyMapRule.Police})
    policyMapRule.EntityData.Children.Append("service-policy", types.YChild{"ServicePolicy", &policyMapRule.ServicePolicy})
    policyMapRule.EntityData.Children.Append("cac-local", types.YChild{"CacLocal", &policyMapRule.CacLocal})
    policyMapRule.EntityData.Children.Append("flow-params", types.YChild{"FlowParams", &policyMapRule.FlowParams})
    policyMapRule.EntityData.Children.Append("metrics-ipcbr", types.YChild{"MetricsIpcbr", &policyMapRule.MetricsIpcbr})
    policyMapRule.EntityData.Children.Append("react", types.YChild{"React", &policyMapRule.React})
    policyMapRule.EntityData.Children.Append("pbr-redirect", types.YChild{"PbrRedirect", &policyMapRule.PbrRedirect})
    policyMapRule.EntityData.Children.Append("pbr-forward", types.YChild{"PbrForward", &policyMapRule.PbrForward})
    policyMapRule.EntityData.Children.Append("service-function-path", types.YChild{"ServiceFunctionPath", &policyMapRule.ServiceFunctionPath})
    policyMapRule.EntityData.Children.Append("http-enrichment", types.YChild{"HttpEnrichment", &policyMapRule.HttpEnrichment})
    policyMapRule.EntityData.Leafs = types.NewOrderedMap()
    policyMapRule.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", policyMapRule.ClassName})
    policyMapRule.EntityData.Leafs.Append("class-type", types.YLeaf{"ClassType", policyMapRule.ClassType})
    policyMapRule.EntityData.Leafs.Append("priority-level", types.YLeaf{"PriorityLevel", policyMapRule.PriorityLevel})
    policyMapRule.EntityData.Leafs.Append("default-red", types.YLeaf{"DefaultRed", policyMapRule.DefaultRed})
    policyMapRule.EntityData.Leafs.Append("ecn-red", types.YLeaf{"EcnRed", policyMapRule.EcnRed})
    policyMapRule.EntityData.Leafs.Append("http-redirect", types.YLeaf{"HttpRedirect", policyMapRule.HttpRedirect})
    policyMapRule.EntityData.Leafs.Append("pbr-transmit", types.YLeaf{"PbrTransmit", policyMapRule.PbrTransmit})
    policyMapRule.EntityData.Leafs.Append("pbr-drop", types.YLeaf{"PbrDrop", policyMapRule.PbrDrop})
    policyMapRule.EntityData.Leafs.Append("decap-gre", types.YLeaf{"DecapGre", policyMapRule.DecapGre})
    policyMapRule.EntityData.Leafs.Append("service-fragment", types.YLeaf{"ServiceFragment", policyMapRule.ServiceFragment})
    policyMapRule.EntityData.Leafs.Append("fragment", types.YLeaf{"Fragment", policyMapRule.Fragment})

    policyMapRule.EntityData.YListKeys = []string {"ClassName", "ClassType"}

    return &(policyMapRule.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape
// Policy action shape.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rate configuration.
    Rate PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate

    // Burst size configuration.
    Burst PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst
}

func (shape *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape) GetEntityData() *types.CommonEntityData {
    shape.EntityData.YFilter = shape.YFilter
    shape.EntityData.YangName = "shape"
    shape.EntityData.BundleName = "cisco_ios_xr"
    shape.EntityData.ParentYangName = "policy-map-rule"
    shape.EntityData.SegmentPath = "shape"
    shape.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + shape.EntityData.SegmentPath
    shape.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shape.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shape.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shape.EntityData.Children = types.NewOrderedMap()
    shape.EntityData.Children.Append("rate", types.YChild{"Rate", &shape.Rate})
    shape.EntityData.Children.Append("burst", types.YChild{"Burst", &shape.Burst})
    shape.EntityData.Leafs = types.NewOrderedMap()

    shape.EntityData.YListKeys = []string {}

    return &(shape.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate
// Rate configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Shape bandwidth value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Shape bandwidth units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(percent)|(per-million)|(per-thousand).
    Unit interface{}
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "shape"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/shape/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Leafs = types.NewOrderedMap()
    rate.EntityData.Leafs.Append("value", types.YLeaf{"Value", rate.Value})
    rate.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", rate.Unit})

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst
// Burst size configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Burst size value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Burst size units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Units interface{}
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Shape_Burst) GetEntityData() *types.CommonEntityData {
    burst.EntityData.YFilter = burst.YFilter
    burst.EntityData.YangName = "burst"
    burst.EntityData.BundleName = "cisco_ios_xr"
    burst.EntityData.ParentYangName = "shape"
    burst.EntityData.SegmentPath = "burst"
    burst.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/shape/" + burst.EntityData.SegmentPath
    burst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    burst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    burst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    burst.EntityData.Children = types.NewOrderedMap()
    burst.EntityData.Leafs = types.NewOrderedMap()
    burst.EntityData.Leafs.Append("value", types.YLeaf{"Value", burst.Value})
    burst.EntityData.Leafs.Append("units", types.YLeaf{"Units", burst.Units})

    burst.EntityData.YListKeys = []string {}

    return &(burst.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth
// Policy action minimum bandwidth queue.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum bandwidth value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Minimum bandwidth units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(percent)|(per-million)|(per-thousand).
    Unit interface{}
}

func (minBandwidth *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MinBandwidth) GetEntityData() *types.CommonEntityData {
    minBandwidth.EntityData.YFilter = minBandwidth.YFilter
    minBandwidth.EntityData.YangName = "min-bandwidth"
    minBandwidth.EntityData.BundleName = "cisco_ios_xr"
    minBandwidth.EntityData.ParentYangName = "policy-map-rule"
    minBandwidth.EntityData.SegmentPath = "min-bandwidth"
    minBandwidth.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + minBandwidth.EntityData.SegmentPath
    minBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minBandwidth.EntityData.Children = types.NewOrderedMap()
    minBandwidth.EntityData.Leafs = types.NewOrderedMap()
    minBandwidth.EntityData.Leafs.Append("value", types.YLeaf{"Value", minBandwidth.Value})
    minBandwidth.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", minBandwidth.Unit})

    minBandwidth.EntityData.YListKeys = []string {}

    return &(minBandwidth.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining
// Policy action bandwidth remaining queue.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remaining bandwidth value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Remaining bandwidth units. The type is string with pattern:
    // (percent)|(ratio).
    Unit interface{}
}

func (bandwidthRemaining *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_BandwidthRemaining) GetEntityData() *types.CommonEntityData {
    bandwidthRemaining.EntityData.YFilter = bandwidthRemaining.YFilter
    bandwidthRemaining.EntityData.YangName = "bandwidth-remaining"
    bandwidthRemaining.EntityData.BundleName = "cisco_ios_xr"
    bandwidthRemaining.EntityData.ParentYangName = "policy-map-rule"
    bandwidthRemaining.EntityData.SegmentPath = "bandwidth-remaining"
    bandwidthRemaining.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + bandwidthRemaining.EntityData.SegmentPath
    bandwidthRemaining.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidthRemaining.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidthRemaining.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidthRemaining.EntityData.Children = types.NewOrderedMap()
    bandwidthRemaining.EntityData.Leafs = types.NewOrderedMap()
    bandwidthRemaining.EntityData.Leafs.Append("value", types.YLeaf{"Value", bandwidthRemaining.Value})
    bandwidthRemaining.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", bandwidthRemaining.Unit})

    bandwidthRemaining.EntityData.YListKeys = []string {}

    return &(bandwidthRemaining.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit
// Policy action queue limit.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remaining bandwidth value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Remaining bandwidth units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells)|(percent).
    Unit interface{}
}

func (queueLimit *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_QueueLimit) GetEntityData() *types.CommonEntityData {
    queueLimit.EntityData.YFilter = queueLimit.YFilter
    queueLimit.EntityData.YangName = "queue-limit"
    queueLimit.EntityData.BundleName = "cisco_ios_xr"
    queueLimit.EntityData.ParentYangName = "policy-map-rule"
    queueLimit.EntityData.SegmentPath = "queue-limit"
    queueLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + queueLimit.EntityData.SegmentPath
    queueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queueLimit.EntityData.Children = types.NewOrderedMap()
    queueLimit.EntityData.Leafs = types.NewOrderedMap()
    queueLimit.EntityData.Leafs.Append("value", types.YLeaf{"Value", queueLimit.Value})
    queueLimit.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", queueLimit.Unit})

    queueLimit.EntityData.YListKeys = []string {}

    return &(queueLimit.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc
// Policy action pfc.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pfc Pause set value. The type is interface{}.
    PfcPauseSet interface{}

    
    PfcBufferSize PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize

    
    PfcPauseThreshold PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold

    
    PfcResumeThreshold PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold
}

func (pfc *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc) GetEntityData() *types.CommonEntityData {
    pfc.EntityData.YFilter = pfc.YFilter
    pfc.EntityData.YangName = "pfc"
    pfc.EntityData.BundleName = "cisco_ios_xr"
    pfc.EntityData.ParentYangName = "policy-map-rule"
    pfc.EntityData.SegmentPath = "pfc"
    pfc.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + pfc.EntityData.SegmentPath
    pfc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pfc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pfc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pfc.EntityData.Children = types.NewOrderedMap()
    pfc.EntityData.Children.Append("pfc-buffer-size", types.YChild{"PfcBufferSize", &pfc.PfcBufferSize})
    pfc.EntityData.Children.Append("pfc-pause-threshold", types.YChild{"PfcPauseThreshold", &pfc.PfcPauseThreshold})
    pfc.EntityData.Children.Append("pfc-resume-threshold", types.YChild{"PfcResumeThreshold", &pfc.PfcResumeThreshold})
    pfc.EntityData.Leafs = types.NewOrderedMap()
    pfc.EntityData.Leafs.Append("pfc-pause-set", types.YLeaf{"PfcPauseSet", pfc.PfcPauseSet})

    pfc.EntityData.YListKeys = []string {}

    return &(pfc.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pfc buffer size value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Pfc buffer size units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Unit interface{}
}

func (pfcBufferSize *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcBufferSize) GetEntityData() *types.CommonEntityData {
    pfcBufferSize.EntityData.YFilter = pfcBufferSize.YFilter
    pfcBufferSize.EntityData.YangName = "pfc-buffer-size"
    pfcBufferSize.EntityData.BundleName = "cisco_ios_xr"
    pfcBufferSize.EntityData.ParentYangName = "pfc"
    pfcBufferSize.EntityData.SegmentPath = "pfc-buffer-size"
    pfcBufferSize.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/pfc/" + pfcBufferSize.EntityData.SegmentPath
    pfcBufferSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pfcBufferSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pfcBufferSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pfcBufferSize.EntityData.Children = types.NewOrderedMap()
    pfcBufferSize.EntityData.Leafs = types.NewOrderedMap()
    pfcBufferSize.EntityData.Leafs.Append("value", types.YLeaf{"Value", pfcBufferSize.Value})
    pfcBufferSize.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", pfcBufferSize.Unit})

    pfcBufferSize.EntityData.YListKeys = []string {}

    return &(pfcBufferSize.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pfc pause threshold value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Pfc pause threshold units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Unit interface{}
}

func (pfcPauseThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcPauseThreshold) GetEntityData() *types.CommonEntityData {
    pfcPauseThreshold.EntityData.YFilter = pfcPauseThreshold.YFilter
    pfcPauseThreshold.EntityData.YangName = "pfc-pause-threshold"
    pfcPauseThreshold.EntityData.BundleName = "cisco_ios_xr"
    pfcPauseThreshold.EntityData.ParentYangName = "pfc"
    pfcPauseThreshold.EntityData.SegmentPath = "pfc-pause-threshold"
    pfcPauseThreshold.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/pfc/" + pfcPauseThreshold.EntityData.SegmentPath
    pfcPauseThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pfcPauseThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pfcPauseThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pfcPauseThreshold.EntityData.Children = types.NewOrderedMap()
    pfcPauseThreshold.EntityData.Leafs = types.NewOrderedMap()
    pfcPauseThreshold.EntityData.Leafs.Append("value", types.YLeaf{"Value", pfcPauseThreshold.Value})
    pfcPauseThreshold.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", pfcPauseThreshold.Unit})

    pfcPauseThreshold.EntityData.YListKeys = []string {}

    return &(pfcPauseThreshold.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pfc resume threshold value. The type is interface{} with range:
    // 0..4294967295.
    Value interface{}

    // Pfc resume threshold units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Unit interface{}
}

func (pfcResumeThreshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Pfc_PfcResumeThreshold) GetEntityData() *types.CommonEntityData {
    pfcResumeThreshold.EntityData.YFilter = pfcResumeThreshold.YFilter
    pfcResumeThreshold.EntityData.YangName = "pfc-resume-threshold"
    pfcResumeThreshold.EntityData.BundleName = "cisco_ios_xr"
    pfcResumeThreshold.EntityData.ParentYangName = "pfc"
    pfcResumeThreshold.EntityData.SegmentPath = "pfc-resume-threshold"
    pfcResumeThreshold.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/pfc/" + pfcResumeThreshold.EntityData.SegmentPath
    pfcResumeThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pfcResumeThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pfcResumeThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pfcResumeThreshold.EntityData.Children = types.NewOrderedMap()
    pfcResumeThreshold.EntityData.Leafs = types.NewOrderedMap()
    pfcResumeThreshold.EntityData.Leafs.Append("value", types.YLeaf{"Value", pfcResumeThreshold.Value})
    pfcResumeThreshold.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", pfcResumeThreshold.Unit})

    pfcResumeThreshold.EntityData.YListKeys = []string {}

    return &(pfcResumeThreshold.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect
// Random early detection.
// All RED profiles in a class must be based
// on the same field.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (randomDetect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_RandomDetect) GetEntityData() *types.CommonEntityData {
    randomDetect.EntityData.YFilter = randomDetect.YFilter
    randomDetect.EntityData.YangName = "random-detect"
    randomDetect.EntityData.BundleName = "cisco_ios_xr"
    randomDetect.EntityData.ParentYangName = "policy-map-rule"
    randomDetect.EntityData.SegmentPath = "random-detect" + types.AddKeyToken(randomDetect.ThresholdMinValue, "threshold-min-value") + types.AddKeyToken(randomDetect.ThresholdMinUnits, "threshold-min-units") + types.AddKeyToken(randomDetect.ThresholdMaxValue, "threshold-max-value") + types.AddKeyToken(randomDetect.ThresholdMaxUnits, "threshold-max-units")
    randomDetect.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + randomDetect.EntityData.SegmentPath
    randomDetect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    randomDetect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    randomDetect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    randomDetect.EntityData.Children = types.NewOrderedMap()
    randomDetect.EntityData.Leafs = types.NewOrderedMap()
    randomDetect.EntityData.Leafs.Append("threshold-min-value", types.YLeaf{"ThresholdMinValue", randomDetect.ThresholdMinValue})
    randomDetect.EntityData.Leafs.Append("threshold-min-units", types.YLeaf{"ThresholdMinUnits", randomDetect.ThresholdMinUnits})
    randomDetect.EntityData.Leafs.Append("threshold-max-value", types.YLeaf{"ThresholdMaxValue", randomDetect.ThresholdMaxValue})
    randomDetect.EntityData.Leafs.Append("threshold-max-units", types.YLeaf{"ThresholdMaxUnits", randomDetect.ThresholdMaxUnits})
    randomDetect.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", randomDetect.Cos})
    randomDetect.EntityData.Leafs.Append("discard-class", types.YLeaf{"DiscardClass", randomDetect.DiscardClass})
    randomDetect.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", randomDetect.Dscp})
    randomDetect.EntityData.Leafs.Append("mpls-exp", types.YLeaf{"MplsExp", randomDetect.MplsExp})
    randomDetect.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", randomDetect.Precedence})
    randomDetect.EntityData.Leafs.Append("dei", types.YLeaf{"Dei", randomDetect.Dei})
    randomDetect.EntityData.Leafs.Append("ecn", types.YLeaf{"Ecn", randomDetect.Ecn})

    randomDetect.EntityData.YListKeys = []string {"ThresholdMinValue", "ThresholdMinUnits", "ThresholdMaxValue", "ThresholdMaxUnits"}

    return &(randomDetect.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set
// Policy action packet marking.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set struct {
    EntityData types.CommonEntityData
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

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "policy-map-rule"
    set.EntityData.SegmentPath = "set"
    set.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + set.EntityData.SegmentPath
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", set.Dscp})
    set.EntityData.Leafs.Append("qos-group", types.YLeaf{"QosGroup", set.QosGroup})
    set.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", set.TrafficClass})
    set.EntityData.Leafs.Append("discard-class", types.YLeaf{"DiscardClass", set.DiscardClass})
    set.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", set.ForwardClass})
    set.EntityData.Leafs.Append("df", types.YLeaf{"Df", set.Df})
    set.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", set.Cos})
    set.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", set.InnerCos})
    set.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", set.Precedence})
    set.EntityData.Leafs.Append("precedence-tunnel", types.YLeaf{"PrecedenceTunnel", set.PrecedenceTunnel})
    set.EntityData.Leafs.Append("mpls-experimental-top-most", types.YLeaf{"MplsExperimentalTopMost", set.MplsExperimentalTopMost})
    set.EntityData.Leafs.Append("mpls-experimental-imposition", types.YLeaf{"MplsExperimentalImposition", set.MplsExperimentalImposition})
    set.EntityData.Leafs.Append("srp-priority", types.YLeaf{"SrpPriority", set.SrpPriority})
    set.EntityData.Leafs.Append("fr-de", types.YLeaf{"FrDe", set.FrDe})
    set.EntityData.Leafs.Append("dei", types.YLeaf{"Dei", set.Dei})
    set.EntityData.Leafs.Append("dei-imposition", types.YLeaf{"DeiImposition", set.DeiImposition})
    set.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", set.SourceAddress})
    set.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", set.DestinationAddress})

    set.EntityData.YListKeys = []string {}

    return &(set.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police
// Configures traffic policing action.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police struct {
    EntityData types.CommonEntityData
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

func (police *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police) GetEntityData() *types.CommonEntityData {
    police.EntityData.YFilter = police.YFilter
    police.EntityData.YangName = "police"
    police.EntityData.BundleName = "cisco_ios_xr"
    police.EntityData.ParentYangName = "policy-map-rule"
    police.EntityData.SegmentPath = "police"
    police.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + police.EntityData.SegmentPath
    police.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    police.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    police.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    police.EntityData.Children = types.NewOrderedMap()
    police.EntityData.Children.Append("rate", types.YChild{"Rate", &police.Rate})
    police.EntityData.Children.Append("peak-rate", types.YChild{"PeakRate", &police.PeakRate})
    police.EntityData.Children.Append("burst", types.YChild{"Burst", &police.Burst})
    police.EntityData.Children.Append("peak-burst", types.YChild{"PeakBurst", &police.PeakBurst})
    police.EntityData.Children.Append("conform-action", types.YChild{"ConformAction", &police.ConformAction})
    police.EntityData.Children.Append("exceed-action", types.YChild{"ExceedAction", &police.ExceedAction})
    police.EntityData.Children.Append("violate-action", types.YChild{"ViolateAction", &police.ViolateAction})
    police.EntityData.Leafs = types.NewOrderedMap()

    police.EntityData.YListKeys = []string {}

    return &(police.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate
// Rate configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rate value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Rate units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(pps)|(percent)|(cellsps).
    Units interface{}
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "police"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Leafs = types.NewOrderedMap()
    rate.EntityData.Leafs.Append("value", types.YLeaf{"Value", rate.Value})
    rate.EntityData.Leafs.Append("units", types.YLeaf{"Units", rate.Units})

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate
// Peak rate configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peak rate value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Peak rate units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(pps)|(percent)|(cellsps).
    Units interface{}
}

func (peakRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakRate) GetEntityData() *types.CommonEntityData {
    peakRate.EntityData.YFilter = peakRate.YFilter
    peakRate.EntityData.YangName = "peak-rate"
    peakRate.EntityData.BundleName = "cisco_ios_xr"
    peakRate.EntityData.ParentYangName = "police"
    peakRate.EntityData.SegmentPath = "peak-rate"
    peakRate.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/" + peakRate.EntityData.SegmentPath
    peakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakRate.EntityData.Children = types.NewOrderedMap()
    peakRate.EntityData.Leafs = types.NewOrderedMap()
    peakRate.EntityData.Leafs.Append("value", types.YLeaf{"Value", peakRate.Value})
    peakRate.EntityData.Leafs.Append("units", types.YLeaf{"Units", peakRate.Units})

    peakRate.EntityData.YListKeys = []string {}

    return &(peakRate.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst
// Burst configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Burst value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Burst units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Units interface{}
}

func (burst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_Burst) GetEntityData() *types.CommonEntityData {
    burst.EntityData.YFilter = burst.YFilter
    burst.EntityData.YangName = "burst"
    burst.EntityData.BundleName = "cisco_ios_xr"
    burst.EntityData.ParentYangName = "police"
    burst.EntityData.SegmentPath = "burst"
    burst.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/" + burst.EntityData.SegmentPath
    burst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    burst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    burst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    burst.EntityData.Children = types.NewOrderedMap()
    burst.EntityData.Leafs = types.NewOrderedMap()
    burst.EntityData.Leafs.Append("value", types.YLeaf{"Value", burst.Value})
    burst.EntityData.Leafs.Append("units", types.YLeaf{"Units", burst.Units})

    burst.EntityData.YListKeys = []string {}

    return &(burst.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst
// Peak burst configuration.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peak burst value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Peak burst units. The type is string with pattern:
    // (bytes)|(kbytes)|(mbytes)|(gbytes)|(us)|(ms)|(packets)|(cells).
    Units interface{}
}

func (peakBurst *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_PeakBurst) GetEntityData() *types.CommonEntityData {
    peakBurst.EntityData.YFilter = peakBurst.YFilter
    peakBurst.EntityData.YangName = "peak-burst"
    peakBurst.EntityData.BundleName = "cisco_ios_xr"
    peakBurst.EntityData.ParentYangName = "police"
    peakBurst.EntityData.SegmentPath = "peak-burst"
    peakBurst.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/" + peakBurst.EntityData.SegmentPath
    peakBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peakBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peakBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peakBurst.EntityData.Children = types.NewOrderedMap()
    peakBurst.EntityData.Leafs = types.NewOrderedMap()
    peakBurst.EntityData.Leafs.Append("value", types.YLeaf{"Value", peakBurst.Value})
    peakBurst.EntityData.Leafs.Append("units", types.YLeaf{"Units", peakBurst.Units})

    peakBurst.EntityData.YListKeys = []string {}

    return &(peakBurst.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction
// Configures the action to take on packets that conform 
// to the rate limit.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Police action transmit. The type is interface{}.
    Transmit interface{}

    // Police action drop. The type is interface{}.
    Drop interface{}

    // Police action packet marking.
    Set PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set
}

func (conformAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction) GetEntityData() *types.CommonEntityData {
    conformAction.EntityData.YFilter = conformAction.YFilter
    conformAction.EntityData.YangName = "conform-action"
    conformAction.EntityData.BundleName = "cisco_ios_xr"
    conformAction.EntityData.ParentYangName = "police"
    conformAction.EntityData.SegmentPath = "conform-action"
    conformAction.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/" + conformAction.EntityData.SegmentPath
    conformAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformAction.EntityData.Children = types.NewOrderedMap()
    conformAction.EntityData.Children.Append("set", types.YChild{"Set", &conformAction.Set})
    conformAction.EntityData.Leafs = types.NewOrderedMap()
    conformAction.EntityData.Leafs.Append("Transmit", types.YLeaf{"Transmit", conformAction.Transmit})
    conformAction.EntityData.Leafs.Append("drop", types.YLeaf{"Drop", conformAction.Drop})

    conformAction.EntityData.YListKeys = []string {}

    return &(conformAction.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set
// Police action packet marking.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set struct {
    EntityData types.CommonEntityData
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

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ConformAction_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "conform-action"
    set.EntityData.SegmentPath = "set"
    set.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/conform-action/" + set.EntityData.SegmentPath
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", set.Dscp})
    set.EntityData.Leafs.Append("qos-group", types.YLeaf{"QosGroup", set.QosGroup})
    set.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", set.TrafficClass})
    set.EntityData.Leafs.Append("discard-class", types.YLeaf{"DiscardClass", set.DiscardClass})
    set.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", set.ForwardClass})
    set.EntityData.Leafs.Append("df", types.YLeaf{"Df", set.Df})
    set.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", set.Cos})
    set.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", set.InnerCos})
    set.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", set.Precedence})
    set.EntityData.Leafs.Append("precedence-tunnel", types.YLeaf{"PrecedenceTunnel", set.PrecedenceTunnel})
    set.EntityData.Leafs.Append("mpls-experimental-top-most", types.YLeaf{"MplsExperimentalTopMost", set.MplsExperimentalTopMost})
    set.EntityData.Leafs.Append("mpls-experimental-imposition", types.YLeaf{"MplsExperimentalImposition", set.MplsExperimentalImposition})
    set.EntityData.Leafs.Append("srp-priority", types.YLeaf{"SrpPriority", set.SrpPriority})
    set.EntityData.Leafs.Append("fr-de", types.YLeaf{"FrDe", set.FrDe})
    set.EntityData.Leafs.Append("dei", types.YLeaf{"Dei", set.Dei})
    set.EntityData.Leafs.Append("dei-imposition", types.YLeaf{"DeiImposition", set.DeiImposition})
    set.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", set.SourceAddress})
    set.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", set.DestinationAddress})

    set.EntityData.YListKeys = []string {}

    return &(set.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction
// Configures the action to take on packets that exceed 
// the rate limit.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Police action transmit. The type is interface{}.
    Transmit interface{}

    // Police action drop. The type is interface{}.
    Drop interface{}

    // Police action packet marking.
    Set PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set
}

func (exceedAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction) GetEntityData() *types.CommonEntityData {
    exceedAction.EntityData.YFilter = exceedAction.YFilter
    exceedAction.EntityData.YangName = "exceed-action"
    exceedAction.EntityData.BundleName = "cisco_ios_xr"
    exceedAction.EntityData.ParentYangName = "police"
    exceedAction.EntityData.SegmentPath = "exceed-action"
    exceedAction.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/" + exceedAction.EntityData.SegmentPath
    exceedAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedAction.EntityData.Children = types.NewOrderedMap()
    exceedAction.EntityData.Children.Append("set", types.YChild{"Set", &exceedAction.Set})
    exceedAction.EntityData.Leafs = types.NewOrderedMap()
    exceedAction.EntityData.Leafs.Append("Transmit", types.YLeaf{"Transmit", exceedAction.Transmit})
    exceedAction.EntityData.Leafs.Append("drop", types.YLeaf{"Drop", exceedAction.Drop})

    exceedAction.EntityData.YListKeys = []string {}

    return &(exceedAction.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set
// Police action packet marking.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set struct {
    EntityData types.CommonEntityData
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

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ExceedAction_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "exceed-action"
    set.EntityData.SegmentPath = "set"
    set.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/exceed-action/" + set.EntityData.SegmentPath
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", set.Dscp})
    set.EntityData.Leafs.Append("qos-group", types.YLeaf{"QosGroup", set.QosGroup})
    set.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", set.TrafficClass})
    set.EntityData.Leafs.Append("discard-class", types.YLeaf{"DiscardClass", set.DiscardClass})
    set.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", set.ForwardClass})
    set.EntityData.Leafs.Append("df", types.YLeaf{"Df", set.Df})
    set.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", set.Cos})
    set.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", set.InnerCos})
    set.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", set.Precedence})
    set.EntityData.Leafs.Append("precedence-tunnel", types.YLeaf{"PrecedenceTunnel", set.PrecedenceTunnel})
    set.EntityData.Leafs.Append("mpls-experimental-top-most", types.YLeaf{"MplsExperimentalTopMost", set.MplsExperimentalTopMost})
    set.EntityData.Leafs.Append("mpls-experimental-imposition", types.YLeaf{"MplsExperimentalImposition", set.MplsExperimentalImposition})
    set.EntityData.Leafs.Append("srp-priority", types.YLeaf{"SrpPriority", set.SrpPriority})
    set.EntityData.Leafs.Append("fr-de", types.YLeaf{"FrDe", set.FrDe})
    set.EntityData.Leafs.Append("dei", types.YLeaf{"Dei", set.Dei})
    set.EntityData.Leafs.Append("dei-imposition", types.YLeaf{"DeiImposition", set.DeiImposition})
    set.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", set.SourceAddress})
    set.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", set.DestinationAddress})

    set.EntityData.YListKeys = []string {}

    return &(set.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction
// Configures the action to take on packets that violate
// the rate limit.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Police action transmit. The type is interface{}.
    Transmit interface{}

    // Police action drop. The type is interface{}.
    Drop interface{}

    // Police action packet marking.
    Set PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set
}

func (violateAction *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction) GetEntityData() *types.CommonEntityData {
    violateAction.EntityData.YFilter = violateAction.YFilter
    violateAction.EntityData.YangName = "violate-action"
    violateAction.EntityData.BundleName = "cisco_ios_xr"
    violateAction.EntityData.ParentYangName = "police"
    violateAction.EntityData.SegmentPath = "violate-action"
    violateAction.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/" + violateAction.EntityData.SegmentPath
    violateAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    violateAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    violateAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    violateAction.EntityData.Children = types.NewOrderedMap()
    violateAction.EntityData.Children.Append("set", types.YChild{"Set", &violateAction.Set})
    violateAction.EntityData.Leafs = types.NewOrderedMap()
    violateAction.EntityData.Leafs.Append("Transmit", types.YLeaf{"Transmit", violateAction.Transmit})
    violateAction.EntityData.Leafs.Append("drop", types.YLeaf{"Drop", violateAction.Drop})

    violateAction.EntityData.YListKeys = []string {}

    return &(violateAction.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set
// Police action packet marking.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set struct {
    EntityData types.CommonEntityData
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

func (set *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_Police_ViolateAction_Set) GetEntityData() *types.CommonEntityData {
    set.EntityData.YFilter = set.YFilter
    set.EntityData.YangName = "set"
    set.EntityData.BundleName = "cisco_ios_xr"
    set.EntityData.ParentYangName = "violate-action"
    set.EntityData.SegmentPath = "set"
    set.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/police/violate-action/" + set.EntityData.SegmentPath
    set.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    set.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    set.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    set.EntityData.Children = types.NewOrderedMap()
    set.EntityData.Leafs = types.NewOrderedMap()
    set.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", set.Dscp})
    set.EntityData.Leafs.Append("qos-group", types.YLeaf{"QosGroup", set.QosGroup})
    set.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", set.TrafficClass})
    set.EntityData.Leafs.Append("discard-class", types.YLeaf{"DiscardClass", set.DiscardClass})
    set.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", set.ForwardClass})
    set.EntityData.Leafs.Append("df", types.YLeaf{"Df", set.Df})
    set.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", set.Cos})
    set.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", set.InnerCos})
    set.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", set.Precedence})
    set.EntityData.Leafs.Append("precedence-tunnel", types.YLeaf{"PrecedenceTunnel", set.PrecedenceTunnel})
    set.EntityData.Leafs.Append("mpls-experimental-top-most", types.YLeaf{"MplsExperimentalTopMost", set.MplsExperimentalTopMost})
    set.EntityData.Leafs.Append("mpls-experimental-imposition", types.YLeaf{"MplsExperimentalImposition", set.MplsExperimentalImposition})
    set.EntityData.Leafs.Append("srp-priority", types.YLeaf{"SrpPriority", set.SrpPriority})
    set.EntityData.Leafs.Append("fr-de", types.YLeaf{"FrDe", set.FrDe})
    set.EntityData.Leafs.Append("dei", types.YLeaf{"Dei", set.Dei})
    set.EntityData.Leafs.Append("dei-imposition", types.YLeaf{"DeiImposition", set.DeiImposition})
    set.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", set.SourceAddress})
    set.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", set.DestinationAddress})

    set.EntityData.YListKeys = []string {}

    return &(set.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy
// Configure a child service policy.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of service-policy. The type is string with pattern:
    // [a-zA-Z0-9][a-zA-Z0-9\._@$%+#:=<>\-]{0,62}.
    PolicyName interface{}

    // Type of service-policy. The type is string with pattern:
    // (PBR)|(QOS)|(REDIRECT)|(TRAFFIC)|(pbr)|(qos)|(redirect)|(traffic).
    Type interface{}
}

func (servicePolicy *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServicePolicy) GetEntityData() *types.CommonEntityData {
    servicePolicy.EntityData.YFilter = servicePolicy.YFilter
    servicePolicy.EntityData.YangName = "service-policy"
    servicePolicy.EntityData.BundleName = "cisco_ios_xr"
    servicePolicy.EntityData.ParentYangName = "policy-map-rule"
    servicePolicy.EntityData.SegmentPath = "service-policy"
    servicePolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + servicePolicy.EntityData.SegmentPath
    servicePolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servicePolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servicePolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servicePolicy.EntityData.Children = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs = types.NewOrderedMap()
    servicePolicy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", servicePolicy.PolicyName})
    servicePolicy.EntityData.Leafs.Append("type", types.YLeaf{"Type", servicePolicy.Type})

    servicePolicy.EntityData.YListKeys = []string {}

    return &(servicePolicy.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal
// Policy action CAC.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal struct {
    EntityData types.CommonEntityData
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

func (cacLocal *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal) GetEntityData() *types.CommonEntityData {
    cacLocal.EntityData.YFilter = cacLocal.YFilter
    cacLocal.EntityData.YangName = "cac-local"
    cacLocal.EntityData.BundleName = "cisco_ios_xr"
    cacLocal.EntityData.ParentYangName = "policy-map-rule"
    cacLocal.EntityData.SegmentPath = "cac-local"
    cacLocal.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + cacLocal.EntityData.SegmentPath
    cacLocal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cacLocal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cacLocal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cacLocal.EntityData.Children = types.NewOrderedMap()
    cacLocal.EntityData.Children.Append("rate", types.YChild{"Rate", &cacLocal.Rate})
    cacLocal.EntityData.Children.Append("flow-rate", types.YChild{"FlowRate", &cacLocal.FlowRate})
    cacLocal.EntityData.Leafs = types.NewOrderedMap()
    cacLocal.EntityData.Leafs.Append("flow-idle-timeout", types.YLeaf{"FlowIdleTimeout", cacLocal.FlowIdleTimeout})

    cacLocal.EntityData.YListKeys = []string {}

    return &(cacLocal.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate
// The rate allocated for all flows.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rate value. The type is interface{} with range: 1..4294967295.
    Value interface{}

    // Rate units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(cellsps).
    Units interface{}
}

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "cac-local"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/cac-local/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Leafs = types.NewOrderedMap()
    rate.EntityData.Leafs.Append("value", types.YLeaf{"Value", rate.Value})
    rate.EntityData.Leafs.Append("units", types.YLeaf{"Units", rate.Units})

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate
// The rate allocated per flow.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rate value. The type is interface{} with range: 1..4294967295.
    Value interface{}

    // Rate units. The type is string with pattern:
    // (bps)|(kbps)|(mbps)|(gbps)|(cellsps).
    Units interface{}
}

func (flowRate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_CacLocal_FlowRate) GetEntityData() *types.CommonEntityData {
    flowRate.EntityData.YFilter = flowRate.YFilter
    flowRate.EntityData.YangName = "flow-rate"
    flowRate.EntityData.BundleName = "cisco_ios_xr"
    flowRate.EntityData.ParentYangName = "cac-local"
    flowRate.EntityData.SegmentPath = "flow-rate"
    flowRate.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/cac-local/" + flowRate.EntityData.SegmentPath
    flowRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowRate.EntityData.Children = types.NewOrderedMap()
    flowRate.EntityData.Leafs = types.NewOrderedMap()
    flowRate.EntityData.Leafs.Append("value", types.YLeaf{"Value", flowRate.Value})
    flowRate.EntityData.Leafs.Append("units", types.YLeaf{"Units", flowRate.Units})

    flowRate.EntityData.YListKeys = []string {}

    return &(flowRate.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams
// Policy flow monitoring action.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams struct {
    EntityData types.CommonEntityData
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

func (flowParams *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_FlowParams) GetEntityData() *types.CommonEntityData {
    flowParams.EntityData.YFilter = flowParams.YFilter
    flowParams.EntityData.YangName = "flow-params"
    flowParams.EntityData.BundleName = "cisco_ios_xr"
    flowParams.EntityData.ParentYangName = "policy-map-rule"
    flowParams.EntityData.SegmentPath = "flow-params"
    flowParams.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + flowParams.EntityData.SegmentPath
    flowParams.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flowParams.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flowParams.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flowParams.EntityData.Children = types.NewOrderedMap()
    flowParams.EntityData.Leafs = types.NewOrderedMap()
    flowParams.EntityData.Leafs.Append("max-flow", types.YLeaf{"MaxFlow", flowParams.MaxFlow})
    flowParams.EntityData.Leafs.Append("interval-duration", types.YLeaf{"IntervalDuration", flowParams.IntervalDuration})
    flowParams.EntityData.Leafs.Append("history", types.YLeaf{"History", flowParams.History})
    flowParams.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", flowParams.Timeout})

    flowParams.EntityData.YListKeys = []string {}

    return &(flowParams.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr
// Policy IP-CBR metric action.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nominal per-flow data rate.
    Rate PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate

    // Media-packet structure.
    MediaPacket PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket
}

func (metricsIpcbr *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr) GetEntityData() *types.CommonEntityData {
    metricsIpcbr.EntityData.YFilter = metricsIpcbr.YFilter
    metricsIpcbr.EntityData.YangName = "metrics-ipcbr"
    metricsIpcbr.EntityData.BundleName = "cisco_ios_xr"
    metricsIpcbr.EntityData.ParentYangName = "policy-map-rule"
    metricsIpcbr.EntityData.SegmentPath = "metrics-ipcbr"
    metricsIpcbr.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + metricsIpcbr.EntityData.SegmentPath
    metricsIpcbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metricsIpcbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metricsIpcbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metricsIpcbr.EntityData.Children = types.NewOrderedMap()
    metricsIpcbr.EntityData.Children.Append("rate", types.YChild{"Rate", &metricsIpcbr.Rate})
    metricsIpcbr.EntityData.Children.Append("media-packet", types.YChild{"MediaPacket", &metricsIpcbr.MediaPacket})
    metricsIpcbr.EntityData.Leafs = types.NewOrderedMap()

    metricsIpcbr.EntityData.YListKeys = []string {}

    return &(metricsIpcbr.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate
// Nominal per-flow data rate.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate struct {
    EntityData types.CommonEntityData
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

func (rate *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_Rate) GetEntityData() *types.CommonEntityData {
    rate.EntityData.YFilter = rate.YFilter
    rate.EntityData.YangName = "rate"
    rate.EntityData.BundleName = "cisco_ios_xr"
    rate.EntityData.ParentYangName = "metrics-ipcbr"
    rate.EntityData.SegmentPath = "rate"
    rate.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/metrics-ipcbr/" + rate.EntityData.SegmentPath
    rate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rate.EntityData.Children = types.NewOrderedMap()
    rate.EntityData.Leafs = types.NewOrderedMap()
    rate.EntityData.Leafs.Append("layer3", types.YLeaf{"Layer3", rate.Layer3})
    rate.EntityData.Leafs.Append("packet", types.YLeaf{"Packet", rate.Packet})
    rate.EntityData.Leafs.Append("media", types.YLeaf{"Media", rate.Media})

    rate.EntityData.YListKeys = []string {}

    return &(rate.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket
// Media-packet structure.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nominal size of the media-packet. The type is interface{} with range:
    // 0..65535. Units are bytes.
    Size interface{}

    // Nominal number of media packets in an IP payload. The type is interface{}
    // with range: 1..64. Units are packets.
    CountInLayer3 interface{}
}

func (mediaPacket *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_MetricsIpcbr_MediaPacket) GetEntityData() *types.CommonEntityData {
    mediaPacket.EntityData.YFilter = mediaPacket.YFilter
    mediaPacket.EntityData.YangName = "media-packet"
    mediaPacket.EntityData.BundleName = "cisco_ios_xr"
    mediaPacket.EntityData.ParentYangName = "metrics-ipcbr"
    mediaPacket.EntityData.SegmentPath = "media-packet"
    mediaPacket.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/metrics-ipcbr/" + mediaPacket.EntityData.SegmentPath
    mediaPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mediaPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mediaPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mediaPacket.EntityData.Children = types.NewOrderedMap()
    mediaPacket.EntityData.Leafs = types.NewOrderedMap()
    mediaPacket.EntityData.Leafs.Append("size", types.YLeaf{"Size", mediaPacket.Size})
    mediaPacket.EntityData.Leafs.Append("count-in-layer3", types.YLeaf{"CountInLayer3", mediaPacket.CountInLayer3})

    mediaPacket.EntityData.YListKeys = []string {}

    return &(mediaPacket.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React
// Policy action react.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React struct {
    EntityData types.CommonEntityData
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

func (react *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React) GetEntityData() *types.CommonEntityData {
    react.EntityData.YFilter = react.YFilter
    react.EntityData.YangName = "react"
    react.EntityData.BundleName = "cisco_ios_xr"
    react.EntityData.ParentYangName = "policy-map-rule"
    react.EntityData.SegmentPath = "react"
    react.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + react.EntityData.SegmentPath
    react.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    react.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    react.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    react.EntityData.Children = types.NewOrderedMap()
    react.EntityData.Children.Append("action", types.YChild{"Action", &react.Action})
    react.EntityData.Children.Append("alarm", types.YChild{"Alarm", &react.Alarm})
    react.EntityData.Children.Append("threshold", types.YChild{"Threshold", &react.Threshold})
    react.EntityData.Leafs = types.NewOrderedMap()
    react.EntityData.Leafs.Append("descrition", types.YLeaf{"Descrition", react.Descrition})
    react.EntityData.Leafs.Append("criterion-delay-factor", types.YLeaf{"CriterionDelayFactor", react.CriterionDelayFactor})
    react.EntityData.Leafs.Append("criterion-media-stop", types.YLeaf{"CriterionMediaStop", react.CriterionMediaStop})
    react.EntityData.Leafs.Append("criterion-mrv", types.YLeaf{"CriterionMrv", react.CriterionMrv})
    react.EntityData.Leafs.Append("criterion-flow-count", types.YLeaf{"CriterionFlowCount", react.CriterionFlowCount})
    react.EntityData.Leafs.Append("criterion-packet-rate", types.YLeaf{"CriterionPacketRate", react.CriterionPacketRate})

    react.EntityData.YListKeys = []string {}

    return &(react.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action
// Action on alert.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Syslog. The type is interface{}.
    Syslog interface{}

    // SNMP. The type is interface{}.
    Snmp interface{}
}

func (action *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Action) GetEntityData() *types.CommonEntityData {
    action.EntityData.YFilter = action.YFilter
    action.EntityData.YangName = "action"
    action.EntityData.BundleName = "cisco_ios_xr"
    action.EntityData.ParentYangName = "react"
    action.EntityData.SegmentPath = "action"
    action.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/react/" + action.EntityData.SegmentPath
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = types.NewOrderedMap()
    action.EntityData.Leafs = types.NewOrderedMap()
    action.EntityData.Leafs.Append("syslog", types.YLeaf{"Syslog", action.Syslog})
    action.EntityData.Leafs.Append("snmp", types.YLeaf{"Snmp", action.Snmp})

    action.EntityData.YListKeys = []string {}

    return &(action.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm
// Alaram settings.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Severity of the alarm. The type is string with pattern:
    // (informational)|(notification)|(warning)|(error)|(critical)|(alert)|(emergency).
    Severity interface{}

    // Alarm type.
    Type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type
}

func (alarm *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm) GetEntityData() *types.CommonEntityData {
    alarm.EntityData.YFilter = alarm.YFilter
    alarm.EntityData.YangName = "alarm"
    alarm.EntityData.BundleName = "cisco_ios_xr"
    alarm.EntityData.ParentYangName = "react"
    alarm.EntityData.SegmentPath = "alarm"
    alarm.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/react/" + alarm.EntityData.SegmentPath
    alarm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarm.EntityData.Children = types.NewOrderedMap()
    alarm.EntityData.Children.Append("type", types.YChild{"Type", &alarm.Type})
    alarm.EntityData.Leafs = types.NewOrderedMap()
    alarm.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarm.Severity})

    alarm.EntityData.YListKeys = []string {}

    return &(alarm.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type
// Alarm type.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type struct {
    EntityData types.CommonEntityData
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

func (self *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Alarm_Type) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "type"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "alarm"
    self.EntityData.SegmentPath = "type"
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/react/alarm/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("discrete", types.YLeaf{"Discrete", self.Discrete})
    self.EntityData.Leafs.Append("group-count", types.YLeaf{"GroupCount", self.GroupCount})
    self.EntityData.Leafs.Append("group-percent", types.YLeaf{"GroupPercent", self.GroupPercent})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold
// Alarm threshold settings.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Alarm trigger value settings.
    TriggerValue PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue

    // Alarm trigger type settings.
    TriggerType PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType
}

func (threshold *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "react"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/react/" + threshold.EntityData.SegmentPath
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Children.Append("trigger-value", types.YChild{"TriggerValue", &threshold.TriggerValue})
    threshold.EntityData.Children.Append("trigger-type", types.YChild{"TriggerType", &threshold.TriggerType})
    threshold.EntityData.Leafs = types.NewOrderedMap()

    threshold.EntityData.YListKeys = []string {}

    return &(threshold.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue
// Alarm trigger value settings.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue struct {
    EntityData types.CommonEntityData
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

func (triggerValue *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerValue) GetEntityData() *types.CommonEntityData {
    triggerValue.EntityData.YFilter = triggerValue.YFilter
    triggerValue.EntityData.YangName = "trigger-value"
    triggerValue.EntityData.BundleName = "cisco_ios_xr"
    triggerValue.EntityData.ParentYangName = "threshold"
    triggerValue.EntityData.SegmentPath = "trigger-value"
    triggerValue.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/react/threshold/" + triggerValue.EntityData.SegmentPath
    triggerValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerValue.EntityData.Children = types.NewOrderedMap()
    triggerValue.EntityData.Leafs = types.NewOrderedMap()
    triggerValue.EntityData.Leafs.Append("greater-than", types.YLeaf{"GreaterThan", triggerValue.GreaterThan})
    triggerValue.EntityData.Leafs.Append("greater-than-equal", types.YLeaf{"GreaterThanEqual", triggerValue.GreaterThanEqual})
    triggerValue.EntityData.Leafs.Append("less-than", types.YLeaf{"LessThan", triggerValue.LessThan})
    triggerValue.EntityData.Leafs.Append("less-than-equal", types.YLeaf{"LessThanEqual", triggerValue.LessThanEqual})
    triggerValue.EntityData.Leafs.Append("range", types.YLeaf{"Range", triggerValue.Range})

    triggerValue.EntityData.YListKeys = []string {}

    return &(triggerValue.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType
// Alarm trigger type settings.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Immediate trigger. The type is interface{}.
    Immediate interface{}

    // Trigger averaged over N intervals. The type is interface{} with range:
    // 0..4294967295.
    Average interface{}
}

func (triggerType *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_React_Threshold_TriggerType) GetEntityData() *types.CommonEntityData {
    triggerType.EntityData.YFilter = triggerType.YFilter
    triggerType.EntityData.YangName = "trigger-type"
    triggerType.EntityData.BundleName = "cisco_ios_xr"
    triggerType.EntityData.ParentYangName = "threshold"
    triggerType.EntityData.SegmentPath = "trigger-type"
    triggerType.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/react/threshold/" + triggerType.EntityData.SegmentPath
    triggerType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    triggerType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    triggerType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    triggerType.EntityData.Children = types.NewOrderedMap()
    triggerType.EntityData.Leafs = types.NewOrderedMap()
    triggerType.EntityData.Leafs.Append("immediate", types.YLeaf{"Immediate", triggerType.Immediate})
    triggerType.EntityData.Leafs.Append("average", types.YLeaf{"Average", triggerType.Average})

    triggerType.EntityData.YListKeys = []string {}

    return &(triggerType.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect
// Policy action redirect
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy action redirect IPv4.
    Ipv4 PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4

    // Policy action redirect IPv6.
    Ipv6 PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6

    // Next hop address.
    NextHop PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop
}

func (pbrRedirect *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect) GetEntityData() *types.CommonEntityData {
    pbrRedirect.EntityData.YFilter = pbrRedirect.YFilter
    pbrRedirect.EntityData.YangName = "pbr-redirect"
    pbrRedirect.EntityData.BundleName = "cisco_ios_xr"
    pbrRedirect.EntityData.ParentYangName = "policy-map-rule"
    pbrRedirect.EntityData.SegmentPath = "pbr-redirect"
    pbrRedirect.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + pbrRedirect.EntityData.SegmentPath
    pbrRedirect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbrRedirect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbrRedirect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbrRedirect.EntityData.Children = types.NewOrderedMap()
    pbrRedirect.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &pbrRedirect.Ipv4})
    pbrRedirect.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &pbrRedirect.Ipv6})
    pbrRedirect.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &pbrRedirect.NextHop})
    pbrRedirect.EntityData.Leafs = types.NewOrderedMap()

    pbrRedirect.EntityData.YListKeys = []string {}

    return &(pbrRedirect.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4
// Policy action redirect IPv4
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4NextHop interface{}

    // IPv4 VRF. The type is string.
    Vrf interface{}
}

func (ipv4 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "pbr-redirect"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/pbr-redirect/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("ipv4-next-hop", types.YLeaf{"Ipv4NextHop", ipv4.Ipv4NextHop})
    ipv4.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", ipv4.Vrf})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6
// Policy action redirect IPv6
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6NextHop interface{}

    // IPv6 VRF. The type is string.
    Vrf interface{}
}

func (ipv6 *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "pbr-redirect"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/pbr-redirect/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("ipv6-next-hop", types.YLeaf{"Ipv6NextHop", ipv6.Ipv6NextHop})
    ipv6.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", ipv6.Vrf})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop
// Next hop address.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route Target.
    RouteTarget PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget
}

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "pbr-redirect"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/pbr-redirect/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("route-target", types.YChild{"RouteTarget", &nextHop.RouteTarget})
    nextHop.EntityData.Leafs = types.NewOrderedMap()

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget
// Route Target
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 2-byte/4-byte AS number. The type is interface{} with range: 1..4294967295.
    AsNumber interface{}

    // ASN2:index 2/4 byte (hex or decimal format). The type is interface{} with
    // range: 0..4294967295.
    Index interface{}

    // IPv4 address.
    Ipv4Address PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address
}

func (routeTarget *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget) GetEntityData() *types.CommonEntityData {
    routeTarget.EntityData.YFilter = routeTarget.YFilter
    routeTarget.EntityData.YangName = "route-target"
    routeTarget.EntityData.BundleName = "cisco_ios_xr"
    routeTarget.EntityData.ParentYangName = "next-hop"
    routeTarget.EntityData.SegmentPath = "route-target"
    routeTarget.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/pbr-redirect/next-hop/" + routeTarget.EntityData.SegmentPath
    routeTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    routeTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    routeTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    routeTarget.EntityData.Children = types.NewOrderedMap()
    routeTarget.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", &routeTarget.Ipv4Address})
    routeTarget.EntityData.Leafs = types.NewOrderedMap()
    routeTarget.EntityData.Leafs.Append("as-number", types.YLeaf{"AsNumber", routeTarget.AsNumber})
    routeTarget.EntityData.Leafs.Append("index", types.YLeaf{"Index", routeTarget.Index})

    routeTarget.EntityData.YListKeys = []string {}

    return &(routeTarget.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address
// IPv4 address.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // IPv4 netmask. The type is string.
    Netmask interface{}
}

func (ipv4Address *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrRedirect_NextHop_RouteTarget_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "route-target"
    ipv4Address.EntityData.SegmentPath = "ipv4-address"
    ipv4Address.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/pbr-redirect/next-hop/route-target/" + ipv4Address.EntityData.SegmentPath
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Address.Address})
    ipv4Address.EntityData.Leafs.Append("netmask", types.YLeaf{"Netmask", ipv4Address.Netmask})

    ipv4Address.EntityData.YListKeys = []string {}

    return &(ipv4Address.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward
// Policy action PBR forward.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use system default routing table. The type is interface{}.
    Default interface{}

    // Use specific next-hop. Here we present 5 different combination  for the pbf
    // next-hop.  1. vrf with v6 address  2. vrf with v4 address  3. vrf   4. v4
    // address  5. v6 address.
    NextHop PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop
}

func (pbrForward *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward) GetEntityData() *types.CommonEntityData {
    pbrForward.EntityData.YFilter = pbrForward.YFilter
    pbrForward.EntityData.YangName = "pbr-forward"
    pbrForward.EntityData.BundleName = "cisco_ios_xr"
    pbrForward.EntityData.ParentYangName = "policy-map-rule"
    pbrForward.EntityData.SegmentPath = "pbr-forward"
    pbrForward.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + pbrForward.EntityData.SegmentPath
    pbrForward.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbrForward.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbrForward.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbrForward.EntityData.Children = types.NewOrderedMap()
    pbrForward.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &pbrForward.NextHop})
    pbrForward.EntityData.Leafs = types.NewOrderedMap()
    pbrForward.EntityData.Leafs.Append("default", types.YLeaf{"Default", pbrForward.Default})

    pbrForward.EntityData.YListKeys = []string {}

    return &(pbrForward.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (nextHop *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_PbrForward_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "pbr-forward"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/pbr-forward/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", nextHop.Vrf})
    nextHop.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", nextHop.Ipv4Address})
    nextHop.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", nextHop.Ipv6Address})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath
// Policy action service function path.
// This type is a presence type.
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Service function path id. The type is interface{} with range: 1..16777215.
    // This attribute is mandatory.
    PathId interface{}

    // Service function path index. The type is interface{} with range: 1..255.
    // This attribute is mandatory.
    Index interface{}

    // Service function path metadata name. The type is string.
    Metadata interface{}
}

func (serviceFunctionPath *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_ServiceFunctionPath) GetEntityData() *types.CommonEntityData {
    serviceFunctionPath.EntityData.YFilter = serviceFunctionPath.YFilter
    serviceFunctionPath.EntityData.YangName = "service-function-path"
    serviceFunctionPath.EntityData.BundleName = "cisco_ios_xr"
    serviceFunctionPath.EntityData.ParentYangName = "policy-map-rule"
    serviceFunctionPath.EntityData.SegmentPath = "service-function-path"
    serviceFunctionPath.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + serviceFunctionPath.EntityData.SegmentPath
    serviceFunctionPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceFunctionPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceFunctionPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceFunctionPath.EntityData.Children = types.NewOrderedMap()
    serviceFunctionPath.EntityData.Leafs = types.NewOrderedMap()
    serviceFunctionPath.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", serviceFunctionPath.PathId})
    serviceFunctionPath.EntityData.Leafs.Append("index", types.YLeaf{"Index", serviceFunctionPath.Index})
    serviceFunctionPath.EntityData.Leafs.Append("metadata", types.YLeaf{"Metadata", serviceFunctionPath.Metadata})

    serviceFunctionPath.EntityData.YListKeys = []string {}

    return &(serviceFunctionPath.EntityData)
}

// PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_HttpEnrichment
// HTTP Enrichment action
type PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_HttpEnrichment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber Mac. The type is interface{}.
    Subscribermac interface{}

    // Subscriber IP. The type is interface{}.
    Subscriberip interface{}

    // Hostname. The type is interface{}.
    Hostname interface{}

    // Bng Identifier Interface. The type is interface{}.
    Bngidentifierinterface interface{}
}

func (httpEnrichment *PolicyManager_PolicyMaps_PolicyMap_PolicyMapRule_HttpEnrichment) GetEntityData() *types.CommonEntityData {
    httpEnrichment.EntityData.YFilter = httpEnrichment.YFilter
    httpEnrichment.EntityData.YangName = "http-enrichment"
    httpEnrichment.EntityData.BundleName = "cisco_ios_xr"
    httpEnrichment.EntityData.ParentYangName = "policy-map-rule"
    httpEnrichment.EntityData.SegmentPath = "http-enrichment"
    httpEnrichment.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-policymgr-cfg:policy-manager/policy-maps/policy-map/policy-map-rule/" + httpEnrichment.EntityData.SegmentPath
    httpEnrichment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    httpEnrichment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    httpEnrichment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    httpEnrichment.EntityData.Children = types.NewOrderedMap()
    httpEnrichment.EntityData.Leafs = types.NewOrderedMap()
    httpEnrichment.EntityData.Leafs.Append("subscribermac", types.YLeaf{"Subscribermac", httpEnrichment.Subscribermac})
    httpEnrichment.EntityData.Leafs.Append("subscriberip", types.YLeaf{"Subscriberip", httpEnrichment.Subscriberip})
    httpEnrichment.EntityData.Leafs.Append("hostname", types.YLeaf{"Hostname", httpEnrichment.Hostname})
    httpEnrichment.EntityData.Leafs.Append("bngidentifierinterface", types.YLeaf{"Bngidentifierinterface", httpEnrichment.Bngidentifierinterface})

    httpEnrichment.EntityData.YListKeys = []string {}

    return &(httpEnrichment.EntityData)
}

