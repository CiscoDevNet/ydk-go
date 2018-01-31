// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-acl package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-acl-and-prefix-list: IPv4 ACL configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_acl_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_acl_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-acl-cfg ipv4-acl-and-prefix-list}", reflect.TypeOf(Ipv4AclAndPrefixList{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-acl-cfg:ipv4-acl-and-prefix-list", reflect.TypeOf(Ipv4AclAndPrefixList{}))
}

// NextHopType represents Next-hop type.
type NextHopType string

const (
    // None next-hop.
    NextHopType_none_next_hop NextHopType = "none-next-hop"

    // Regular next-hop.
    NextHopType_regular_next_hop NextHopType = "regular-next-hop"

    // Default next-hop.
    NextHopType_default_next_hop NextHopType = "default-next-hop"
)

// Ipv4AclAndPrefixList
// IPv4 ACL configuration data
type Ipv4AclAndPrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of access lists.  Entries in this table and the
    // AccessListExistenceTable table must be kept consistent.
    Accesses Ipv4AclAndPrefixList_Accesses

    // Table of ACL prefix lists.  Entries in this table and the
    // PrefixListExistenceTable table must be kept consistent.
    Prefixes Ipv4AclAndPrefixList_Prefixes

    // Control access lists log updates.
    LogUpdate Ipv4AclAndPrefixList_LogUpdate
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetFilter() yfilter.YFilter { return ipv4AclAndPrefixList.YFilter }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) SetFilter(yf yfilter.YFilter) { ipv4AclAndPrefixList.YFilter = yf }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetGoName(yname string) string {
    if yname == "accesses" { return "Accesses" }
    if yname == "prefixes" { return "Prefixes" }
    if yname == "log-update" { return "LogUpdate" }
    return ""
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-acl-cfg:ipv4-acl-and-prefix-list"
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accesses" {
        return &ipv4AclAndPrefixList.Accesses
    }
    if childYangName == "prefixes" {
        return &ipv4AclAndPrefixList.Prefixes
    }
    if childYangName == "log-update" {
        return &ipv4AclAndPrefixList.LogUpdate
    }
    return nil
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accesses"] = &ipv4AclAndPrefixList.Accesses
    children["prefixes"] = &ipv4AclAndPrefixList.Prefixes
    children["log-update"] = &ipv4AclAndPrefixList.LogUpdate
    return children
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetYangName() string { return "ipv4-acl-and-prefix-list" }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) SetParent(parent types.Entity) { ipv4AclAndPrefixList.parent = parent }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetParent() types.Entity { return ipv4AclAndPrefixList.parent }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-acl-cfg" }

// Ipv4AclAndPrefixList_Accesses
// Table of access lists.  Entries in this table
// and the AccessListExistenceTable table must be
// kept consistent
type Ipv4AclAndPrefixList_Accesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An ACL. The type is slice of Ipv4AclAndPrefixList_Accesses_Access.
    Access []Ipv4AclAndPrefixList_Accesses_Access
}

func (accesses *Ipv4AclAndPrefixList_Accesses) GetFilter() yfilter.YFilter { return accesses.YFilter }

func (accesses *Ipv4AclAndPrefixList_Accesses) SetFilter(yf yfilter.YFilter) { accesses.YFilter = yf }

func (accesses *Ipv4AclAndPrefixList_Accesses) GetGoName(yname string) string {
    if yname == "access" { return "Access" }
    return ""
}

func (accesses *Ipv4AclAndPrefixList_Accesses) GetSegmentPath() string {
    return "accesses"
}

func (accesses *Ipv4AclAndPrefixList_Accesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access" {
        for _, c := range accesses.Access {
            if accesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_Accesses_Access{}
        accesses.Access = append(accesses.Access, child)
        return &accesses.Access[len(accesses.Access)-1]
    }
    return nil
}

func (accesses *Ipv4AclAndPrefixList_Accesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accesses.Access {
        children[accesses.Access[i].GetSegmentPath()] = &accesses.Access[i]
    }
    return children
}

func (accesses *Ipv4AclAndPrefixList_Accesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accesses *Ipv4AclAndPrefixList_Accesses) GetBundleName() string { return "cisco_ios_xr" }

func (accesses *Ipv4AclAndPrefixList_Accesses) GetYangName() string { return "accesses" }

func (accesses *Ipv4AclAndPrefixList_Accesses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accesses *Ipv4AclAndPrefixList_Accesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accesses *Ipv4AclAndPrefixList_Accesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accesses *Ipv4AclAndPrefixList_Accesses) SetParent(parent types.Entity) { accesses.parent = parent }

func (accesses *Ipv4AclAndPrefixList_Accesses) GetParent() types.Entity { return accesses.parent }

func (accesses *Ipv4AclAndPrefixList_Accesses) GetParentYangName() string { return "ipv4-acl-and-prefix-list" }

// Ipv4AclAndPrefixList_Accesses_Access
// An ACL
type Ipv4AclAndPrefixList_Accesses_Access struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Access list name - 64 characters max. The type is
    // string.
    AccessListName interface{}

    // ACL entry table; contains list of ACEs.
    AccessListEntries Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries
}

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetFilter() yfilter.YFilter { return access.YFilter }

func (access *Ipv4AclAndPrefixList_Accesses_Access) SetFilter(yf yfilter.YFilter) { access.YFilter = yf }

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetGoName(yname string) string {
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "access-list-entries" { return "AccessListEntries" }
    return ""
}

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetSegmentPath() string {
    return "access" + "[access-list-name='" + fmt.Sprintf("%v", access.AccessListName) + "']"
}

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-entries" {
        return &access.AccessListEntries
    }
    return nil
}

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list-entries"] = &access.AccessListEntries
    return children
}

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-list-name"] = access.AccessListName
    return leafs
}

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetBundleName() string { return "cisco_ios_xr" }

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetYangName() string { return "access" }

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (access *Ipv4AclAndPrefixList_Accesses_Access) SetParent(parent types.Entity) { access.parent = parent }

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetParent() types.Entity { return access.parent }

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetParentYangName() string { return "accesses" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries
// ACL entry table; contains list of ACEs
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An ACL entry; either a description (remark) or an ACE to match against. The
    // type is slice of
    // Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry.
    AccessListEntry []Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry
}

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetFilter() yfilter.YFilter { return accessListEntries.YFilter }

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) SetFilter(yf yfilter.YFilter) { accessListEntries.YFilter = yf }

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetGoName(yname string) string {
    if yname == "access-list-entry" { return "AccessListEntry" }
    return ""
}

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetSegmentPath() string {
    return "access-list-entries"
}

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-entry" {
        for _, c := range accessListEntries.AccessListEntry {
            if accessListEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry{}
        accessListEntries.AccessListEntry = append(accessListEntries.AccessListEntry, child)
        return &accessListEntries.AccessListEntry[len(accessListEntries.AccessListEntry)-1]
    }
    return nil
}

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessListEntries.AccessListEntry {
        children[accessListEntries.AccessListEntry[i].GetSegmentPath()] = &accessListEntries.AccessListEntry[i]
    }
    return children
}

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetBundleName() string { return "cisco_ios_xr" }

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetYangName() string { return "access-list-entries" }

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) SetParent(parent types.Entity) { accessListEntries.parent = parent }

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetParent() types.Entity { return accessListEntries.parent }

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetParentYangName() string { return "access" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry
// An ACL entry; either a description (remark)
// or an ACE to match against
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence number for this entry. The type is
    // interface{} with range: 1..2147483643.
    SequenceNumber interface{}

    // Whether to forward or drop packets matching the  ACE. The type is
    // Ipv4AclGrantEnum.
    Grant interface{}

    // Protocol operator. Leave unspecified if no protocol comparison is to be
    // done. The type is Ipv4AclOperatorEnum.
    ProtocolOperator interface{}

    // Protocol to match. The type is one of the following types: enumeration
    // Ipv4AclProtocolNumber, or int with range: 0..255.
    Protocol interface{}

    // Protocol2 to match. The type is one of the following types: enumeration
    // Ipv4AclProtocolNumber, or int with range: 0..255.
    Protocol2 interface{}

    // Fragment flags, such as dont-fragment, is-fragment, first-fragment and
    // last-fragment. The type is Ipv4AclFragFlags.
    FragmentType interface{}

    // Counter name. The type is string.
    CounterName interface{}

    // IGMP message type to match. Leave unspecified if  no message type
    // comparison is to be done. The type is one of the following types:
    // enumeration Ipv4AclIgmpNumber, or int with range: 0..255.
    IgmpMessageType interface{}

    // Precedence value to match (if a protocol was  specified), leave unspecified
    // if precedence  comparion is not to be performed. The type is one of the
    // following types: enumeration Ipv4AclPrecedenceNumber, or int with range:
    // 0..7.
    Precedence interface{}

    // Whether and how to log matches against this  entry. The type is
    // Ipv4AclLoggingEnum.
    LogOption interface{}

    // Enable capture. The type is bool.
    Capture interface{}

    // To turn off ICMP generation for deny ACEs. The type is interface{}.
    IcmpOff interface{}

    // Set qos-group number. The type is interface{} with range: 0..512.
    QosGroup interface{}

    // Set TTL Value. Ranges from 0-255. The type is interface{} with range:
    // 0..255.
    SetTtl interface{}

    // Check non-initial fragments. Item is mutually  exclusive with TCP, SCTP,
    // UDP, IGMP and ICMP  comparions and with logging. The type is interface{}.
    Fragments interface{}

    // Comments or a description for the access list. The type is string.
    Remark interface{}

    // IPv4 source network object group name. The type is string with length:
    // 1..64.
    SourcePrefixGroup interface{}

    // IPv4 destination network object group name. The type is string with length:
    // 1..64.
    DestinationPrefixGroup interface{}

    // Source port object group name. The type is string with length: 1..64.
    SourcePortGroup interface{}

    // Destination port object group name. The type is string with length: 1..64.
    DestinationPortGroup interface{}

    // Sequence String for the ace. The type is string with length: 1..64.
    SequenceStr interface{}

    // Source network settings.
    SourceNetwork Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork

    // Destination network settings.
    DestinationNetwork Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork

    // Source port settings.
    SourcePort Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort

    // Destination port settings.
    DestinationPort Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort

    // ICMP settings.
    Icmp Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp

    // TCP settings.
    Tcp Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp

    // Packet length settings.
    PacketLength Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength

    // TTL settings.
    TimeToLive Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive

    // Fragment-offset settings.
    FragmentOffset Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset

    // Next-hop settings.
    NextHop Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop

    // DSCP settings.
    Dscp Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp
}

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetFilter() yfilter.YFilter { return accessListEntry.YFilter }

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) SetFilter(yf yfilter.YFilter) { accessListEntry.YFilter = yf }

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "grant" { return "Grant" }
    if yname == "protocol-operator" { return "ProtocolOperator" }
    if yname == "protocol" { return "Protocol" }
    if yname == "protocol2" { return "Protocol2" }
    if yname == "fragment-type" { return "FragmentType" }
    if yname == "counter-name" { return "CounterName" }
    if yname == "igmp-message-type" { return "IgmpMessageType" }
    if yname == "precedence" { return "Precedence" }
    if yname == "log-option" { return "LogOption" }
    if yname == "capture" { return "Capture" }
    if yname == "icmp-off" { return "IcmpOff" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "set-ttl" { return "SetTtl" }
    if yname == "fragments" { return "Fragments" }
    if yname == "remark" { return "Remark" }
    if yname == "source-prefix-group" { return "SourcePrefixGroup" }
    if yname == "destination-prefix-group" { return "DestinationPrefixGroup" }
    if yname == "source-port-group" { return "SourcePortGroup" }
    if yname == "destination-port-group" { return "DestinationPortGroup" }
    if yname == "sequence-str" { return "SequenceStr" }
    if yname == "source-network" { return "SourceNetwork" }
    if yname == "destination-network" { return "DestinationNetwork" }
    if yname == "source-port" { return "SourcePort" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "icmp" { return "Icmp" }
    if yname == "tcp" { return "Tcp" }
    if yname == "packet-length" { return "PacketLength" }
    if yname == "time-to-live" { return "TimeToLive" }
    if yname == "fragment-offset" { return "FragmentOffset" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetSegmentPath() string {
    return "access-list-entry" + "[sequence-number='" + fmt.Sprintf("%v", accessListEntry.SequenceNumber) + "']"
}

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source-network" {
        return &accessListEntry.SourceNetwork
    }
    if childYangName == "destination-network" {
        return &accessListEntry.DestinationNetwork
    }
    if childYangName == "source-port" {
        return &accessListEntry.SourcePort
    }
    if childYangName == "destination-port" {
        return &accessListEntry.DestinationPort
    }
    if childYangName == "icmp" {
        return &accessListEntry.Icmp
    }
    if childYangName == "tcp" {
        return &accessListEntry.Tcp
    }
    if childYangName == "packet-length" {
        return &accessListEntry.PacketLength
    }
    if childYangName == "time-to-live" {
        return &accessListEntry.TimeToLive
    }
    if childYangName == "fragment-offset" {
        return &accessListEntry.FragmentOffset
    }
    if childYangName == "next-hop" {
        return &accessListEntry.NextHop
    }
    if childYangName == "dscp" {
        return &accessListEntry.Dscp
    }
    return nil
}

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source-network"] = &accessListEntry.SourceNetwork
    children["destination-network"] = &accessListEntry.DestinationNetwork
    children["source-port"] = &accessListEntry.SourcePort
    children["destination-port"] = &accessListEntry.DestinationPort
    children["icmp"] = &accessListEntry.Icmp
    children["tcp"] = &accessListEntry.Tcp
    children["packet-length"] = &accessListEntry.PacketLength
    children["time-to-live"] = &accessListEntry.TimeToLive
    children["fragment-offset"] = &accessListEntry.FragmentOffset
    children["next-hop"] = &accessListEntry.NextHop
    children["dscp"] = &accessListEntry.Dscp
    return children
}

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = accessListEntry.SequenceNumber
    leafs["grant"] = accessListEntry.Grant
    leafs["protocol-operator"] = accessListEntry.ProtocolOperator
    leafs["protocol"] = accessListEntry.Protocol
    leafs["protocol2"] = accessListEntry.Protocol2
    leafs["fragment-type"] = accessListEntry.FragmentType
    leafs["counter-name"] = accessListEntry.CounterName
    leafs["igmp-message-type"] = accessListEntry.IgmpMessageType
    leafs["precedence"] = accessListEntry.Precedence
    leafs["log-option"] = accessListEntry.LogOption
    leafs["capture"] = accessListEntry.Capture
    leafs["icmp-off"] = accessListEntry.IcmpOff
    leafs["qos-group"] = accessListEntry.QosGroup
    leafs["set-ttl"] = accessListEntry.SetTtl
    leafs["fragments"] = accessListEntry.Fragments
    leafs["remark"] = accessListEntry.Remark
    leafs["source-prefix-group"] = accessListEntry.SourcePrefixGroup
    leafs["destination-prefix-group"] = accessListEntry.DestinationPrefixGroup
    leafs["source-port-group"] = accessListEntry.SourcePortGroup
    leafs["destination-port-group"] = accessListEntry.DestinationPortGroup
    leafs["sequence-str"] = accessListEntry.SequenceStr
    return leafs
}

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetBundleName() string { return "cisco_ios_xr" }

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetYangName() string { return "access-list-entry" }

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) SetParent(parent types.Entity) { accessListEntry.parent = parent }

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetParent() types.Entity { return accessListEntry.parent }

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetParentYangName() string { return "access-list-entries" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork
// Source network settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source IPv4 address to match, leave unspecified for any. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Wildcard bits to apply to source address  (if specified), leave unspecified
    // for no  wildcarding. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceWildCardBits interface{}

    // Prefix length to apply to source address  (if specified), leave unspecified
    // for no  wildcarding. The type is interface{} with range: 0..32.
    SourcePrefixLength interface{}
}

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetFilter() yfilter.YFilter { return sourceNetwork.YFilter }

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) SetFilter(yf yfilter.YFilter) { sourceNetwork.YFilter = yf }

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-wild-card-bits" { return "SourceWildCardBits" }
    if yname == "source-prefix-length" { return "SourcePrefixLength" }
    return ""
}

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetSegmentPath() string {
    return "source-network"
}

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = sourceNetwork.SourceAddress
    leafs["source-wild-card-bits"] = sourceNetwork.SourceWildCardBits
    leafs["source-prefix-length"] = sourceNetwork.SourcePrefixLength
    return leafs
}

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetBundleName() string { return "cisco_ios_xr" }

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetYangName() string { return "source-network" }

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) SetParent(parent types.Entity) { sourceNetwork.parent = parent }

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetParent() types.Entity { return sourceNetwork.parent }

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork
// Destination network settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination IPv4 address to match (if a protocol was specified), leave
    // unspecified for any. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Wildcard bits to apply to destination address (if specified), leave
    // unspecified for no  wildcarding. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationWildCardBits interface{}

    // Prefix length to apply to destination address  (if specified), leave
    // unspecified for no  wildcarding. The type is interface{} with range: 0..32.
    DestinationPrefixLength interface{}
}

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetFilter() yfilter.YFilter { return destinationNetwork.YFilter }

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) SetFilter(yf yfilter.YFilter) { destinationNetwork.YFilter = yf }

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetGoName(yname string) string {
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-wild-card-bits" { return "DestinationWildCardBits" }
    if yname == "destination-prefix-length" { return "DestinationPrefixLength" }
    return ""
}

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetSegmentPath() string {
    return "destination-network"
}

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-address"] = destinationNetwork.DestinationAddress
    leafs["destination-wild-card-bits"] = destinationNetwork.DestinationWildCardBits
    leafs["destination-prefix-length"] = destinationNetwork.DestinationPrefixLength
    return leafs
}

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetBundleName() string { return "cisco_ios_xr" }

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetYangName() string { return "destination-network" }

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) SetParent(parent types.Entity) { destinationNetwork.parent = parent }

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetParent() types.Entity { return destinationNetwork.parent }

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort
// Source port settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source comparison operator . Leave unspecified  if no source port
    // comparison is to be done. The type is Ipv4AclOperatorEnum.
    SourceOperator interface{}

    // First source port for comparison, leave  unspecified if source port
    // comparison is not to be performed. The type is one of the following types:
    // enumeration Ipv4AclPortNumber, or int with range: 0..65535.
    FirstSourcePort interface{}

    // Second source port for comparion, leave  unspecified if source port
    // comparison is not to be performed. The type is one of the following types:
    // enumeration Ipv4AclPortNumber, or int with range: 0..65535.
    SecondSourcePort interface{}
}

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetFilter() yfilter.YFilter { return sourcePort.YFilter }

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) SetFilter(yf yfilter.YFilter) { sourcePort.YFilter = yf }

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetGoName(yname string) string {
    if yname == "source-operator" { return "SourceOperator" }
    if yname == "first-source-port" { return "FirstSourcePort" }
    if yname == "second-source-port" { return "SecondSourcePort" }
    return ""
}

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetSegmentPath() string {
    return "source-port"
}

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-operator"] = sourcePort.SourceOperator
    leafs["first-source-port"] = sourcePort.FirstSourcePort
    leafs["second-source-port"] = sourcePort.SecondSourcePort
    return leafs
}

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetBundleName() string { return "cisco_ios_xr" }

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetYangName() string { return "source-port" }

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) SetParent(parent types.Entity) { sourcePort.parent = parent }

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetParent() types.Entity { return sourcePort.parent }

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort
// Destination port settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination comparison operator. Leave  unspecified if no destination port
    // comparison is to be done. The type is Ipv4AclOperatorEnum.
    DestinationOperator interface{}

    // First destination port for comparison, leave unspecified if destination
    // port comparison is not to be performed. The type is one of the following
    // types: enumeration Ipv4AclPortNumber, or int with range: 0..65535.
    FirstDestinationPort interface{}

    // Second destination port for comparion, leave unspecified if destination
    // port comparison is not to be performed. The type is one of the following
    // types: enumeration Ipv4AclPortNumber, or int with range: 0..65535.
    SecondDestinationPort interface{}
}

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetFilter() yfilter.YFilter { return destinationPort.YFilter }

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) SetFilter(yf yfilter.YFilter) { destinationPort.YFilter = yf }

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetGoName(yname string) string {
    if yname == "destination-operator" { return "DestinationOperator" }
    if yname == "first-destination-port" { return "FirstDestinationPort" }
    if yname == "second-destination-port" { return "SecondDestinationPort" }
    return ""
}

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetSegmentPath() string {
    return "destination-port"
}

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-operator"] = destinationPort.DestinationOperator
    leafs["first-destination-port"] = destinationPort.FirstDestinationPort
    leafs["second-destination-port"] = destinationPort.SecondDestinationPort
    return leafs
}

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetBundleName() string { return "cisco_ios_xr" }

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetYangName() string { return "destination-port" }

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) SetParent(parent types.Entity) { destinationPort.parent = parent }

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetParent() types.Entity { return destinationPort.parent }

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp
// ICMP settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Well known ICMP message code types to match,  leave unspecified if ICMP
    // message code type  comparion is not to be performed. The type is
    // Ipv4AclIcmpTypeCodeEnum.
    IcmpTypeCode interface{}
}

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetFilter() yfilter.YFilter { return icmp.YFilter }

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) SetFilter(yf yfilter.YFilter) { icmp.YFilter = yf }

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetGoName(yname string) string {
    if yname == "icmp-type-code" { return "IcmpTypeCode" }
    return ""
}

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetSegmentPath() string {
    return "icmp"
}

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["icmp-type-code"] = icmp.IcmpTypeCode
    return leafs
}

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetBundleName() string { return "cisco_ios_xr" }

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetYangName() string { return "icmp" }

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) SetParent(parent types.Entity) { icmp.parent = parent }

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetParent() types.Entity { return icmp.parent }

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp
// TCP settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP Bits match operator. Leave unspecified if  flexible comparison of TCP
    // bits is not  required. The type is Ipv4AclTcpMatchOperatorEnum.
    TcpBitsMatchOperator interface{}

    // TCP bits to match. Leave unspecified if comparison of TCP bits is not
    // required. The type is one of the following types: enumeration
    // Ipv4AclTcpBitsNumber, or int with range: 0..63.
    TcpBits interface{}

    // TCP bits mask to use for flexible TCP matching. Leave unspecified if
    // tcp-bits-match-operator is  unspecified. The type is one of the following
    // types: enumeration Ipv4AclTcpBitsNumber, or int with range: 0..63.
    TcpBitsMask interface{}
}

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetGoName(yname string) string {
    if yname == "tcp-bits-match-operator" { return "TcpBitsMatchOperator" }
    if yname == "tcp-bits" { return "TcpBits" }
    if yname == "tcp-bits-mask" { return "TcpBitsMask" }
    return ""
}

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetSegmentPath() string {
    return "tcp"
}

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-bits-match-operator"] = tcp.TcpBitsMatchOperator
    leafs["tcp-bits"] = tcp.TcpBits
    leafs["tcp-bits-mask"] = tcp.TcpBitsMask
    return leafs
}

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetBundleName() string { return "cisco_ios_xr" }

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetYangName() string { return "tcp" }

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength
// Packet length settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet length operator applicable if Packet  length is to be compared.
    // Leave unspecified if  no packet length comparison is to be done. The type
    // is Ipv4AclOperatorEnum.
    PacketLengthOperator interface{}

    // Minimum packet length for comparison, leave  unspecified if packet length
    // comparison is not  to be performed or if only the maximum packet  length
    // should be considered. The type is interface{} with range: 0..65535.
    PacketLengthMin interface{}

    // Maximum packet length for comparion, leave  unspecified if packet length
    // comparison is not  to be performed or if only the minimum packet  length
    // should be considered. The type is interface{} with range: 0..65535.
    PacketLengthMax interface{}
}

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetFilter() yfilter.YFilter { return packetLength.YFilter }

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) SetFilter(yf yfilter.YFilter) { packetLength.YFilter = yf }

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetGoName(yname string) string {
    if yname == "packet-length-operator" { return "PacketLengthOperator" }
    if yname == "packet-length-min" { return "PacketLengthMin" }
    if yname == "packet-length-max" { return "PacketLengthMax" }
    return ""
}

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetSegmentPath() string {
    return "packet-length"
}

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packet-length-operator"] = packetLength.PacketLengthOperator
    leafs["packet-length-min"] = packetLength.PacketLengthMin
    leafs["packet-length-max"] = packetLength.PacketLengthMax
    return leafs
}

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetBundleName() string { return "cisco_ios_xr" }

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetYangName() string { return "packet-length" }

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) SetParent(parent types.Entity) { packetLength.parent = parent }

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetParent() types.Entity { return packetLength.parent }

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive
// TTL settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TTL operator is applicable if TTL is to be  compared. Leave unspecified if
    // TTL  classification is not required. The type is Ipv4AclOperatorEnum.
    TimeToLiveOperator interface{}

    // TTL value for comparison OR Minimum TTL value  for TTL range comparision,
    // leave unspecified if TTL classification is not required. The type is
    // interface{} with range: 0..255.
    TimeToLiveMin interface{}

    // Maximum TTL for comparion, leave unspecified if  TTL comparison is not to
    // be performed or if only the minimum TTL should be considered. The type is
    // interface{} with range: 0..255.
    TimeToLiveMax interface{}
}

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetFilter() yfilter.YFilter { return timeToLive.YFilter }

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) SetFilter(yf yfilter.YFilter) { timeToLive.YFilter = yf }

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetGoName(yname string) string {
    if yname == "time-to-live-operator" { return "TimeToLiveOperator" }
    if yname == "time-to-live-min" { return "TimeToLiveMin" }
    if yname == "time-to-live-max" { return "TimeToLiveMax" }
    return ""
}

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetSegmentPath() string {
    return "time-to-live"
}

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-to-live-operator"] = timeToLive.TimeToLiveOperator
    leafs["time-to-live-min"] = timeToLive.TimeToLiveMin
    leafs["time-to-live-max"] = timeToLive.TimeToLiveMax
    return leafs
}

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetBundleName() string { return "cisco_ios_xr" }

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetYangName() string { return "time-to-live" }

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) SetParent(parent types.Entity) { timeToLive.parent = parent }

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetParent() types.Entity { return timeToLive.parent }

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset
// Fragment-offset settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fragment-offset operator if fragment-offset is to be compared. Leave
    // unspecified if fragment-offset classification is not required. The type is
    // Ipv4AclOperatorEnum.
    FragmentOffsetOperator interface{}

    // Fragment-offset value for comparison or first  fragment-offset value for
    // fragment-offset range  comparision, leave unspecified if fragment-offset
    // classification is not required. The type is interface{} with range:
    // 0..8191.
    FragmentOffset1 interface{}

    // Second fragment-offset value for comparion,  leave unspecified if
    // fragment-offset comparison is not to be performed or if only the first
    // fragment-offset should be considered. The type is interface{} with range:
    // 0..8191.
    FragmentOffset2 interface{}
}

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetFilter() yfilter.YFilter { return fragmentOffset.YFilter }

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) SetFilter(yf yfilter.YFilter) { fragmentOffset.YFilter = yf }

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetGoName(yname string) string {
    if yname == "fragment-offset-operator" { return "FragmentOffsetOperator" }
    if yname == "fragment-offset-1" { return "FragmentOffset1" }
    if yname == "fragment-offset-2" { return "FragmentOffset2" }
    return ""
}

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetSegmentPath() string {
    return "fragment-offset"
}

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fragment-offset-operator"] = fragmentOffset.FragmentOffsetOperator
    leafs["fragment-offset-1"] = fragmentOffset.FragmentOffset1
    leafs["fragment-offset-2"] = fragmentOffset.FragmentOffset2
    return leafs
}

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetBundleName() string { return "cisco_ios_xr" }

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetYangName() string { return "fragment-offset" }

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) SetParent(parent types.Entity) { fragmentOffset.parent = parent }

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetParent() types.Entity { return fragmentOffset.parent }

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop
// Next-hop settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The nexthop type. The type is NextHopType.
    NextHopType interface{}

    // The first next-hop settings.
    NextHop1 Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1

    // The second next-hop settings.
    NextHop2 Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2

    // The third next-hop settings.
    NextHop3 Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3
}

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetGoName(yname string) string {
    if yname == "next-hop-type" { return "NextHopType" }
    if yname == "next-hop-1" { return "NextHop1" }
    if yname == "next-hop-2" { return "NextHop2" }
    if yname == "next-hop-3" { return "NextHop3" }
    return ""
}

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "next-hop-1" {
        return &nextHop.NextHop1
    }
    if childYangName == "next-hop-2" {
        return &nextHop.NextHop2
    }
    if childYangName == "next-hop-3" {
        return &nextHop.NextHop3
    }
    return nil
}

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop-1"] = &nextHop.NextHop1
    children["next-hop-2"] = &nextHop.NextHop2
    children["next-hop-3"] = &nextHop.NextHop3
    return children
}

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop-type"] = nextHop.NextHopType
    return leafs
}

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1
// The first next-hop settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IPv4 address of the next-hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // The VRF name of the next-hop. The type is string.
    VrfName interface{}

    // The object tracking name for the next-hop. The type is string.
    TrackName interface{}
}

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetFilter() yfilter.YFilter { return nextHop1.YFilter }

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) SetFilter(yf yfilter.YFilter) { nextHop1.YFilter = yf }

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "track-name" { return "TrackName" }
    return ""
}

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetSegmentPath() string {
    return "next-hop-1"
}

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHop1.NextHop
    leafs["vrf-name"] = nextHop1.VrfName
    leafs["track-name"] = nextHop1.TrackName
    return leafs
}

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetYangName() string { return "next-hop-1" }

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) SetParent(parent types.Entity) { nextHop1.parent = parent }

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetParent() types.Entity { return nextHop1.parent }

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetParentYangName() string { return "next-hop" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2
// The second next-hop settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IPv4 address of the next-hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // The VRF name of the next-hop. The type is string.
    VrfName interface{}

    // The object tracking name for the next-hop. The type is string.
    TrackName interface{}
}

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetFilter() yfilter.YFilter { return nextHop2.YFilter }

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) SetFilter(yf yfilter.YFilter) { nextHop2.YFilter = yf }

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "track-name" { return "TrackName" }
    return ""
}

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetSegmentPath() string {
    return "next-hop-2"
}

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHop2.NextHop
    leafs["vrf-name"] = nextHop2.VrfName
    leafs["track-name"] = nextHop2.TrackName
    return leafs
}

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetYangName() string { return "next-hop-2" }

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) SetParent(parent types.Entity) { nextHop2.parent = parent }

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetParent() types.Entity { return nextHop2.parent }

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetParentYangName() string { return "next-hop" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3
// The third next-hop settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IPv4 address of the next-hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // The VRF name of the next-hop. The type is string.
    VrfName interface{}

    // The object tracking name for the next-hop. The type is string.
    TrackName interface{}
}

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetFilter() yfilter.YFilter { return nextHop3.YFilter }

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) SetFilter(yf yfilter.YFilter) { nextHop3.YFilter = yf }

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "track-name" { return "TrackName" }
    return ""
}

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetSegmentPath() string {
    return "next-hop-3"
}

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHop3.NextHop
    leafs["vrf-name"] = nextHop3.VrfName
    leafs["track-name"] = nextHop3.TrackName
    return leafs
}

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetYangName() string { return "next-hop-3" }

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) SetParent(parent types.Entity) { nextHop3.parent = parent }

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetParent() types.Entity { return nextHop3.parent }

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetParentYangName() string { return "next-hop" }

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp
// DSCP settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DSCP operator is applicable only when DSCP  range is configured. Leave
    // unspecified if  DSCP range is not required. The type is
    // Ipv4AclOperatorEnum.
    DscpOperator interface{}

    // DSCP value to match or minimum DSCP value  for DSCP range comparison, leave
    // unspecified  if DSCP comparion is not to be performed. The type is one of
    // the following types: enumeration Ipv4AclDscpNumber, or int with range:
    // 0..63.
    DscpMin interface{}

    // Maximum DSCP value for comparion, leave  unspecified if DSCP comparison is
    // not to  be performed or if only the minimum DSCP should be considered. The
    // type is one of the following types: enumeration Ipv4AclDscpNumber, or int
    // with range: 0..63.
    DscpMax interface{}
}

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetFilter() yfilter.YFilter { return dscp.YFilter }

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) SetFilter(yf yfilter.YFilter) { dscp.YFilter = yf }

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetGoName(yname string) string {
    if yname == "dscp-operator" { return "DscpOperator" }
    if yname == "dscp-min" { return "DscpMin" }
    if yname == "dscp-max" { return "DscpMax" }
    return ""
}

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetSegmentPath() string {
    return "dscp"
}

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp-operator"] = dscp.DscpOperator
    leafs["dscp-min"] = dscp.DscpMin
    leafs["dscp-max"] = dscp.DscpMax
    return leafs
}

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetBundleName() string { return "cisco_ios_xr" }

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetYangName() string { return "dscp" }

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) SetParent(parent types.Entity) { dscp.parent = parent }

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetParent() types.Entity { return dscp.parent }

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetParentYangName() string { return "access-list-entry" }

// Ipv4AclAndPrefixList_Prefixes
// Table of ACL prefix lists.  Entries in this
// table and the PrefixListExistenceTable table
// must be kept consistent
type Ipv4AclAndPrefixList_Prefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of a prefix list. The type is slice of
    // Ipv4AclAndPrefixList_Prefixes_Prefix.
    Prefix []Ipv4AclAndPrefixList_Prefixes_Prefix
}

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetFilter() yfilter.YFilter { return prefixes.YFilter }

func (prefixes *Ipv4AclAndPrefixList_Prefixes) SetFilter(yf yfilter.YFilter) { prefixes.YFilter = yf }

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetSegmentPath() string {
    return "prefixes"
}

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        for _, c := range prefixes.Prefix {
            if prefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_Prefixes_Prefix{}
        prefixes.Prefix = append(prefixes.Prefix, child)
        return &prefixes.Prefix[len(prefixes.Prefix)-1]
    }
    return nil
}

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixes.Prefix {
        children[prefixes.Prefix[i].GetSegmentPath()] = &prefixes.Prefix[i]
    }
    return children
}

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetBundleName() string { return "cisco_ios_xr" }

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetYangName() string { return "prefixes" }

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixes *Ipv4AclAndPrefixList_Prefixes) SetParent(parent types.Entity) { prefixes.parent = parent }

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetParent() types.Entity { return prefixes.parent }

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetParentYangName() string { return "ipv4-acl-and-prefix-list" }

// Ipv4AclAndPrefixList_Prefixes_Prefix
// Name of a prefix list
type Ipv4AclAndPrefixList_Prefixes_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix list name - max 32 characters. The type is
    // string.
    PrefixListName interface{}

    // Sequence of entries forming a prefix list.
    PrefixListEntries Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries
}

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetGoName(yname string) string {
    if yname == "prefix-list-name" { return "PrefixListName" }
    if yname == "prefix-list-entries" { return "PrefixListEntries" }
    return ""
}

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetSegmentPath() string {
    return "prefix" + "[prefix-list-name='" + fmt.Sprintf("%v", prefix.PrefixListName) + "']"
}

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list-entries" {
        return &prefix.PrefixListEntries
    }
    return nil
}

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-list-entries"] = &prefix.PrefixListEntries
    return children
}

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-list-name"] = prefix.PrefixListName
    return leafs
}

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetYangName() string { return "prefix" }

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetParentYangName() string { return "prefixes" }

// Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries
// Sequence of entries forming a prefix list
// This type is a presence type.
type Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A prefix list entry; either a description (remark) or a prefix to match
    // against. The type is slice of
    // Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry.
    PrefixListEntry []Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry
}

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetFilter() yfilter.YFilter { return prefixListEntries.YFilter }

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) SetFilter(yf yfilter.YFilter) { prefixListEntries.YFilter = yf }

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetGoName(yname string) string {
    if yname == "prefix-list-entry" { return "PrefixListEntry" }
    return ""
}

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetSegmentPath() string {
    return "prefix-list-entries"
}

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list-entry" {
        for _, c := range prefixListEntries.PrefixListEntry {
            if prefixListEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry{}
        prefixListEntries.PrefixListEntry = append(prefixListEntries.PrefixListEntry, child)
        return &prefixListEntries.PrefixListEntry[len(prefixListEntries.PrefixListEntry)-1]
    }
    return nil
}

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixListEntries.PrefixListEntry {
        children[prefixListEntries.PrefixListEntry[i].GetSegmentPath()] = &prefixListEntries.PrefixListEntry[i]
    }
    return children
}

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetYangName() string { return "prefix-list-entries" }

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) SetParent(parent types.Entity) { prefixListEntries.parent = parent }

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetParent() types.Entity { return prefixListEntries.parent }

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetParentYangName() string { return "prefix" }

// Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry
// A prefix list entry; either a description
// (remark) or a prefix to match against
type Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence number of prefix list. The type is
    // interface{} with range: 1..2147483646.
    SequenceNumber interface{}

    // Whether to forward or drop packets matching the prefix list. The type is
    // Ipv4AclGrantEnum.
    Grant interface{}

    // IPv4 address prefix to match. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Mask of IPv4 address prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Netmask interface{}

    // Set to perform an exact prefix length match. Item is mutually exclusive
    // with minimum and maximum length match items. The type is interface{}.
    MatchExactLength interface{}

    // If exact prefix length matching specified, set the length of prefix to be
    // matched. The type is interface{} with range: 0..32.
    ExactPrefixLength interface{}

    // Set to perform a maximum length prefix match .  Item is mutually exclusive
    // with exact length match item. The type is interface{}.
    MatchMaxLength interface{}

    // If maximum length prefix matching specified, set the maximum length of
    // prefix to be matched. The type is interface{} with range: 0..32.
    MaxPrefixLength interface{}

    // Set to perform a minimum length prefix match .  Item is mutually exclusive
    // with exact length match item. The type is interface{}.
    MatchMinLength interface{}

    // If minimum length prefix matching specified, set the minimum length of
    // prefix to be matched. The type is interface{} with range: 0..32.
    MinPrefixLength interface{}

    // Comments or a description for the prefix list.  Item is mutually exclusive
    // with all others in the object. The type is string.
    Remark interface{}
}

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetFilter() yfilter.YFilter { return prefixListEntry.YFilter }

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) SetFilter(yf yfilter.YFilter) { prefixListEntry.YFilter = yf }

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "grant" { return "Grant" }
    if yname == "prefix" { return "Prefix" }
    if yname == "netmask" { return "Netmask" }
    if yname == "match-exact-length" { return "MatchExactLength" }
    if yname == "exact-prefix-length" { return "ExactPrefixLength" }
    if yname == "match-max-length" { return "MatchMaxLength" }
    if yname == "max-prefix-length" { return "MaxPrefixLength" }
    if yname == "match-min-length" { return "MatchMinLength" }
    if yname == "min-prefix-length" { return "MinPrefixLength" }
    if yname == "remark" { return "Remark" }
    return ""
}

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetSegmentPath() string {
    return "prefix-list-entry" + "[sequence-number='" + fmt.Sprintf("%v", prefixListEntry.SequenceNumber) + "']"
}

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = prefixListEntry.SequenceNumber
    leafs["grant"] = prefixListEntry.Grant
    leafs["prefix"] = prefixListEntry.Prefix
    leafs["netmask"] = prefixListEntry.Netmask
    leafs["match-exact-length"] = prefixListEntry.MatchExactLength
    leafs["exact-prefix-length"] = prefixListEntry.ExactPrefixLength
    leafs["match-max-length"] = prefixListEntry.MatchMaxLength
    leafs["max-prefix-length"] = prefixListEntry.MaxPrefixLength
    leafs["match-min-length"] = prefixListEntry.MatchMinLength
    leafs["min-prefix-length"] = prefixListEntry.MinPrefixLength
    leafs["remark"] = prefixListEntry.Remark
    return leafs
}

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetYangName() string { return "prefix-list-entry" }

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) SetParent(parent types.Entity) { prefixListEntry.parent = parent }

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetParent() types.Entity { return prefixListEntry.parent }

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetParentYangName() string { return "prefix-list-entries" }

// Ipv4AclAndPrefixList_LogUpdate
// Control access lists log updates
type Ipv4AclAndPrefixList_LogUpdate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log update threshold (number of hits). The type is interface{} with range:
    // 1..2147483647.
    Threshold interface{}

    // Log update rate (log msgs per second). The type is interface{} with range:
    // 1..1000.
    Rate interface{}
}

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetFilter() yfilter.YFilter { return logUpdate.YFilter }

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) SetFilter(yf yfilter.YFilter) { logUpdate.YFilter = yf }

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    if yname == "rate" { return "Rate" }
    return ""
}

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetSegmentPath() string {
    return "log-update"
}

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold"] = logUpdate.Threshold
    leafs["rate"] = logUpdate.Rate
    return leafs
}

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetBundleName() string { return "cisco_ios_xr" }

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetYangName() string { return "log-update" }

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) SetParent(parent types.Entity) { logUpdate.parent = parent }

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetParent() types.Entity { return logUpdate.parent }

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetParentYangName() string { return "ipv4-acl-and-prefix-list" }

