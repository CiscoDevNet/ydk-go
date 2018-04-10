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
    EntityData types.CommonEntityData
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

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetEntityData() *types.CommonEntityData {
    ipv4AclAndPrefixList.EntityData.YFilter = ipv4AclAndPrefixList.YFilter
    ipv4AclAndPrefixList.EntityData.YangName = "ipv4-acl-and-prefix-list"
    ipv4AclAndPrefixList.EntityData.BundleName = "cisco_ios_xr"
    ipv4AclAndPrefixList.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-acl-cfg"
    ipv4AclAndPrefixList.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-acl-cfg:ipv4-acl-and-prefix-list"
    ipv4AclAndPrefixList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AclAndPrefixList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AclAndPrefixList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AclAndPrefixList.EntityData.Children = make(map[string]types.YChild)
    ipv4AclAndPrefixList.EntityData.Children["accesses"] = types.YChild{"Accesses", &ipv4AclAndPrefixList.Accesses}
    ipv4AclAndPrefixList.EntityData.Children["prefixes"] = types.YChild{"Prefixes", &ipv4AclAndPrefixList.Prefixes}
    ipv4AclAndPrefixList.EntityData.Children["log-update"] = types.YChild{"LogUpdate", &ipv4AclAndPrefixList.LogUpdate}
    ipv4AclAndPrefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4AclAndPrefixList.EntityData)
}

// Ipv4AclAndPrefixList_Accesses
// Table of access lists.  Entries in this table
// and the AccessListExistenceTable table must be
// kept consistent
type Ipv4AclAndPrefixList_Accesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An ACL. The type is slice of Ipv4AclAndPrefixList_Accesses_Access.
    Access []Ipv4AclAndPrefixList_Accesses_Access
}

func (accesses *Ipv4AclAndPrefixList_Accesses) GetEntityData() *types.CommonEntityData {
    accesses.EntityData.YFilter = accesses.YFilter
    accesses.EntityData.YangName = "accesses"
    accesses.EntityData.BundleName = "cisco_ios_xr"
    accesses.EntityData.ParentYangName = "ipv4-acl-and-prefix-list"
    accesses.EntityData.SegmentPath = "accesses"
    accesses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accesses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accesses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accesses.EntityData.Children = make(map[string]types.YChild)
    accesses.EntityData.Children["access"] = types.YChild{"Access", nil}
    for i := range accesses.Access {
        accesses.EntityData.Children[types.GetSegmentPath(&accesses.Access[i])] = types.YChild{"Access", &accesses.Access[i]}
    }
    accesses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accesses.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access
// An ACL
type Ipv4AclAndPrefixList_Accesses_Access struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Access list name - 64 characters max. The type is
    // string.
    AccessListName interface{}

    // ACL entry table; contains list of ACEs.
    AccessListEntries Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries
}

func (access *Ipv4AclAndPrefixList_Accesses_Access) GetEntityData() *types.CommonEntityData {
    access.EntityData.YFilter = access.YFilter
    access.EntityData.YangName = "access"
    access.EntityData.BundleName = "cisco_ios_xr"
    access.EntityData.ParentYangName = "accesses"
    access.EntityData.SegmentPath = "access" + "[access-list-name='" + fmt.Sprintf("%v", access.AccessListName) + "']"
    access.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    access.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    access.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    access.EntityData.Children = make(map[string]types.YChild)
    access.EntityData.Children["access-list-entries"] = types.YChild{"AccessListEntries", &access.AccessListEntries}
    access.EntityData.Leafs = make(map[string]types.YLeaf)
    access.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", access.AccessListName}
    return &(access.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries
// ACL entry table; contains list of ACEs
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An ACL entry; either a description (remark) or an ACE to match against. The
    // type is slice of
    // Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry.
    AccessListEntry []Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry
}

func (accessListEntries *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries) GetEntityData() *types.CommonEntityData {
    accessListEntries.EntityData.YFilter = accessListEntries.YFilter
    accessListEntries.EntityData.YangName = "access-list-entries"
    accessListEntries.EntityData.BundleName = "cisco_ios_xr"
    accessListEntries.EntityData.ParentYangName = "access"
    accessListEntries.EntityData.SegmentPath = "access-list-entries"
    accessListEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListEntries.EntityData.Children = make(map[string]types.YChild)
    accessListEntries.EntityData.Children["access-list-entry"] = types.YChild{"AccessListEntry", nil}
    for i := range accessListEntries.AccessListEntry {
        accessListEntries.EntityData.Children[types.GetSegmentPath(&accessListEntries.AccessListEntry[i])] = types.YChild{"AccessListEntry", &accessListEntries.AccessListEntry[i]}
    }
    accessListEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessListEntries.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry
// An ACL entry; either a description (remark)
// or an ACE to match against
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry struct {
    EntityData types.CommonEntityData
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

    // Fragment flags, such as dont-fragment, is-fragment, first-fragment, and
    // last-fragment. The type is one of the following types: enumeration
    // Ipv4AclFragFlags, or int with range: 1..9.
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

func (accessListEntry *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetEntityData() *types.CommonEntityData {
    accessListEntry.EntityData.YFilter = accessListEntry.YFilter
    accessListEntry.EntityData.YangName = "access-list-entry"
    accessListEntry.EntityData.BundleName = "cisco_ios_xr"
    accessListEntry.EntityData.ParentYangName = "access-list-entries"
    accessListEntry.EntityData.SegmentPath = "access-list-entry" + "[sequence-number='" + fmt.Sprintf("%v", accessListEntry.SequenceNumber) + "']"
    accessListEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListEntry.EntityData.Children = make(map[string]types.YChild)
    accessListEntry.EntityData.Children["source-network"] = types.YChild{"SourceNetwork", &accessListEntry.SourceNetwork}
    accessListEntry.EntityData.Children["destination-network"] = types.YChild{"DestinationNetwork", &accessListEntry.DestinationNetwork}
    accessListEntry.EntityData.Children["source-port"] = types.YChild{"SourcePort", &accessListEntry.SourcePort}
    accessListEntry.EntityData.Children["destination-port"] = types.YChild{"DestinationPort", &accessListEntry.DestinationPort}
    accessListEntry.EntityData.Children["icmp"] = types.YChild{"Icmp", &accessListEntry.Icmp}
    accessListEntry.EntityData.Children["tcp"] = types.YChild{"Tcp", &accessListEntry.Tcp}
    accessListEntry.EntityData.Children["packet-length"] = types.YChild{"PacketLength", &accessListEntry.PacketLength}
    accessListEntry.EntityData.Children["time-to-live"] = types.YChild{"TimeToLive", &accessListEntry.TimeToLive}
    accessListEntry.EntityData.Children["fragment-offset"] = types.YChild{"FragmentOffset", &accessListEntry.FragmentOffset}
    accessListEntry.EntityData.Children["next-hop"] = types.YChild{"NextHop", &accessListEntry.NextHop}
    accessListEntry.EntityData.Children["dscp"] = types.YChild{"Dscp", &accessListEntry.Dscp}
    accessListEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    accessListEntry.EntityData.Leafs["sequence-number"] = types.YLeaf{"SequenceNumber", accessListEntry.SequenceNumber}
    accessListEntry.EntityData.Leafs["grant"] = types.YLeaf{"Grant", accessListEntry.Grant}
    accessListEntry.EntityData.Leafs["protocol-operator"] = types.YLeaf{"ProtocolOperator", accessListEntry.ProtocolOperator}
    accessListEntry.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", accessListEntry.Protocol}
    accessListEntry.EntityData.Leafs["protocol2"] = types.YLeaf{"Protocol2", accessListEntry.Protocol2}
    accessListEntry.EntityData.Leafs["fragment-type"] = types.YLeaf{"FragmentType", accessListEntry.FragmentType}
    accessListEntry.EntityData.Leafs["counter-name"] = types.YLeaf{"CounterName", accessListEntry.CounterName}
    accessListEntry.EntityData.Leafs["igmp-message-type"] = types.YLeaf{"IgmpMessageType", accessListEntry.IgmpMessageType}
    accessListEntry.EntityData.Leafs["precedence"] = types.YLeaf{"Precedence", accessListEntry.Precedence}
    accessListEntry.EntityData.Leafs["log-option"] = types.YLeaf{"LogOption", accessListEntry.LogOption}
    accessListEntry.EntityData.Leafs["capture"] = types.YLeaf{"Capture", accessListEntry.Capture}
    accessListEntry.EntityData.Leafs["icmp-off"] = types.YLeaf{"IcmpOff", accessListEntry.IcmpOff}
    accessListEntry.EntityData.Leafs["qos-group"] = types.YLeaf{"QosGroup", accessListEntry.QosGroup}
    accessListEntry.EntityData.Leafs["set-ttl"] = types.YLeaf{"SetTtl", accessListEntry.SetTtl}
    accessListEntry.EntityData.Leafs["fragments"] = types.YLeaf{"Fragments", accessListEntry.Fragments}
    accessListEntry.EntityData.Leafs["remark"] = types.YLeaf{"Remark", accessListEntry.Remark}
    accessListEntry.EntityData.Leafs["source-prefix-group"] = types.YLeaf{"SourcePrefixGroup", accessListEntry.SourcePrefixGroup}
    accessListEntry.EntityData.Leafs["destination-prefix-group"] = types.YLeaf{"DestinationPrefixGroup", accessListEntry.DestinationPrefixGroup}
    accessListEntry.EntityData.Leafs["source-port-group"] = types.YLeaf{"SourcePortGroup", accessListEntry.SourcePortGroup}
    accessListEntry.EntityData.Leafs["destination-port-group"] = types.YLeaf{"DestinationPortGroup", accessListEntry.DestinationPortGroup}
    accessListEntry.EntityData.Leafs["sequence-str"] = types.YLeaf{"SequenceStr", accessListEntry.SequenceStr}
    return &(accessListEntry.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork
// Source network settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source IPv4 address to match, leave unspecified for any. The type is string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Wildcard bits to apply to source address  (if specified), leave unspecified
    // for no  wildcarding. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceWildCardBits interface{}

    // Prefix length to apply to source address  (if specified), leave unspecified
    // for no  wildcarding. The type is interface{} with range: 0..32.
    SourcePrefixLength interface{}
}

func (sourceNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetEntityData() *types.CommonEntityData {
    sourceNetwork.EntityData.YFilter = sourceNetwork.YFilter
    sourceNetwork.EntityData.YangName = "source-network"
    sourceNetwork.EntityData.BundleName = "cisco_ios_xr"
    sourceNetwork.EntityData.ParentYangName = "access-list-entry"
    sourceNetwork.EntityData.SegmentPath = "source-network"
    sourceNetwork.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceNetwork.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceNetwork.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceNetwork.EntityData.Children = make(map[string]types.YChild)
    sourceNetwork.EntityData.Leafs = make(map[string]types.YLeaf)
    sourceNetwork.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", sourceNetwork.SourceAddress}
    sourceNetwork.EntityData.Leafs["source-wild-card-bits"] = types.YLeaf{"SourceWildCardBits", sourceNetwork.SourceWildCardBits}
    sourceNetwork.EntityData.Leafs["source-prefix-length"] = types.YLeaf{"SourcePrefixLength", sourceNetwork.SourcePrefixLength}
    return &(sourceNetwork.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork
// Destination network settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination IPv4 address to match (if a protocol was specified), leave
    // unspecified for any. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Wildcard bits to apply to destination address (if specified), leave
    // unspecified for no  wildcarding. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationWildCardBits interface{}

    // Prefix length to apply to destination address  (if specified), leave
    // unspecified for no  wildcarding. The type is interface{} with range: 0..32.
    DestinationPrefixLength interface{}
}

func (destinationNetwork *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetEntityData() *types.CommonEntityData {
    destinationNetwork.EntityData.YFilter = destinationNetwork.YFilter
    destinationNetwork.EntityData.YangName = "destination-network"
    destinationNetwork.EntityData.BundleName = "cisco_ios_xr"
    destinationNetwork.EntityData.ParentYangName = "access-list-entry"
    destinationNetwork.EntityData.SegmentPath = "destination-network"
    destinationNetwork.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationNetwork.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationNetwork.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationNetwork.EntityData.Children = make(map[string]types.YChild)
    destinationNetwork.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationNetwork.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", destinationNetwork.DestinationAddress}
    destinationNetwork.EntityData.Leafs["destination-wild-card-bits"] = types.YLeaf{"DestinationWildCardBits", destinationNetwork.DestinationWildCardBits}
    destinationNetwork.EntityData.Leafs["destination-prefix-length"] = types.YLeaf{"DestinationPrefixLength", destinationNetwork.DestinationPrefixLength}
    return &(destinationNetwork.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort
// Source port settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort struct {
    EntityData types.CommonEntityData
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

func (sourcePort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetEntityData() *types.CommonEntityData {
    sourcePort.EntityData.YFilter = sourcePort.YFilter
    sourcePort.EntityData.YangName = "source-port"
    sourcePort.EntityData.BundleName = "cisco_ios_xr"
    sourcePort.EntityData.ParentYangName = "access-list-entry"
    sourcePort.EntityData.SegmentPath = "source-port"
    sourcePort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourcePort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourcePort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourcePort.EntityData.Children = make(map[string]types.YChild)
    sourcePort.EntityData.Leafs = make(map[string]types.YLeaf)
    sourcePort.EntityData.Leafs["source-operator"] = types.YLeaf{"SourceOperator", sourcePort.SourceOperator}
    sourcePort.EntityData.Leafs["first-source-port"] = types.YLeaf{"FirstSourcePort", sourcePort.FirstSourcePort}
    sourcePort.EntityData.Leafs["second-source-port"] = types.YLeaf{"SecondSourcePort", sourcePort.SecondSourcePort}
    return &(sourcePort.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort
// Destination port settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort struct {
    EntityData types.CommonEntityData
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

func (destinationPort *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetEntityData() *types.CommonEntityData {
    destinationPort.EntityData.YFilter = destinationPort.YFilter
    destinationPort.EntityData.YangName = "destination-port"
    destinationPort.EntityData.BundleName = "cisco_ios_xr"
    destinationPort.EntityData.ParentYangName = "access-list-entry"
    destinationPort.EntityData.SegmentPath = "destination-port"
    destinationPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationPort.EntityData.Children = make(map[string]types.YChild)
    destinationPort.EntityData.Leafs = make(map[string]types.YLeaf)
    destinationPort.EntityData.Leafs["destination-operator"] = types.YLeaf{"DestinationOperator", destinationPort.DestinationOperator}
    destinationPort.EntityData.Leafs["first-destination-port"] = types.YLeaf{"FirstDestinationPort", destinationPort.FirstDestinationPort}
    destinationPort.EntityData.Leafs["second-destination-port"] = types.YLeaf{"SecondDestinationPort", destinationPort.SecondDestinationPort}
    return &(destinationPort.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp
// ICMP settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Well known ICMP message code types to match,  leave unspecified if ICMP
    // message code type  comparion is not to be performed. The type is
    // Ipv4AclIcmpTypeCodeEnum.
    IcmpTypeCode interface{}
}

func (icmp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetEntityData() *types.CommonEntityData {
    icmp.EntityData.YFilter = icmp.YFilter
    icmp.EntityData.YangName = "icmp"
    icmp.EntityData.BundleName = "cisco_ios_xr"
    icmp.EntityData.ParentYangName = "access-list-entry"
    icmp.EntityData.SegmentPath = "icmp"
    icmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    icmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    icmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    icmp.EntityData.Children = make(map[string]types.YChild)
    icmp.EntityData.Leafs = make(map[string]types.YLeaf)
    icmp.EntityData.Leafs["icmp-type-code"] = types.YLeaf{"IcmpTypeCode", icmp.IcmpTypeCode}
    return &(icmp.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp
// TCP settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp struct {
    EntityData types.CommonEntityData
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

func (tcp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xr"
    tcp.EntityData.ParentYangName = "access-list-entry"
    tcp.EntityData.SegmentPath = "tcp"
    tcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcp.EntityData.Children = make(map[string]types.YChild)
    tcp.EntityData.Leafs = make(map[string]types.YLeaf)
    tcp.EntityData.Leafs["tcp-bits-match-operator"] = types.YLeaf{"TcpBitsMatchOperator", tcp.TcpBitsMatchOperator}
    tcp.EntityData.Leafs["tcp-bits"] = types.YLeaf{"TcpBits", tcp.TcpBits}
    tcp.EntityData.Leafs["tcp-bits-mask"] = types.YLeaf{"TcpBitsMask", tcp.TcpBitsMask}
    return &(tcp.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength
// Packet length settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength struct {
    EntityData types.CommonEntityData
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

func (packetLength *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetEntityData() *types.CommonEntityData {
    packetLength.EntityData.YFilter = packetLength.YFilter
    packetLength.EntityData.YangName = "packet-length"
    packetLength.EntityData.BundleName = "cisco_ios_xr"
    packetLength.EntityData.ParentYangName = "access-list-entry"
    packetLength.EntityData.SegmentPath = "packet-length"
    packetLength.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetLength.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetLength.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetLength.EntityData.Children = make(map[string]types.YChild)
    packetLength.EntityData.Leafs = make(map[string]types.YLeaf)
    packetLength.EntityData.Leafs["packet-length-operator"] = types.YLeaf{"PacketLengthOperator", packetLength.PacketLengthOperator}
    packetLength.EntityData.Leafs["packet-length-min"] = types.YLeaf{"PacketLengthMin", packetLength.PacketLengthMin}
    packetLength.EntityData.Leafs["packet-length-max"] = types.YLeaf{"PacketLengthMax", packetLength.PacketLengthMax}
    return &(packetLength.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive
// TTL settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TTL operator is applicable if TTL is to be  compared. Leave unspecified if
    // TTL  classification is not required. The type is Ipv4AclOperatorEnum.
    TimeToLiveOperator interface{}

    // TTL value for comparison OR Minimum TTL value  for TTL range comparision,
    // leave unspecified if  TTL classification is not required. The type is
    // interface{} with range: 0..255.
    TimeToLiveMin interface{}

    // Maximum TTL for comparion, leave unspecified if  TTL comparison is not to
    // be performed or if only the minimum TTL should be considered. The type is
    // interface{} with range: 0..255.
    TimeToLiveMax interface{}
}

func (timeToLive *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetEntityData() *types.CommonEntityData {
    timeToLive.EntityData.YFilter = timeToLive.YFilter
    timeToLive.EntityData.YangName = "time-to-live"
    timeToLive.EntityData.BundleName = "cisco_ios_xr"
    timeToLive.EntityData.ParentYangName = "access-list-entry"
    timeToLive.EntityData.SegmentPath = "time-to-live"
    timeToLive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeToLive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeToLive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeToLive.EntityData.Children = make(map[string]types.YChild)
    timeToLive.EntityData.Leafs = make(map[string]types.YLeaf)
    timeToLive.EntityData.Leafs["time-to-live-operator"] = types.YLeaf{"TimeToLiveOperator", timeToLive.TimeToLiveOperator}
    timeToLive.EntityData.Leafs["time-to-live-min"] = types.YLeaf{"TimeToLiveMin", timeToLive.TimeToLiveMin}
    timeToLive.EntityData.Leafs["time-to-live-max"] = types.YLeaf{"TimeToLiveMax", timeToLive.TimeToLiveMax}
    return &(timeToLive.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset
// Fragment-offset settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fragment-offset operator is applicable if fragment-offset is to be
    // compared. Leave unspecified if fragment-offset classification is not
    // required. The type is Ipv4AclOperatorEnum.
    FragmentOffsetOperator interface{}

    // Fragment-offset value for comparison OR first fragment-offset value for
    // fragment-offset range comparision, leave unspecified if fragment-offset
    // classification is not required. The type is interface{} with range:
    // 0..8191.
    FragmentOffset1 interface{}

    // Second fragment-offset value for comparion, leave unspecified if
    // fragment-offset comparison is not to be performed or if only the first
    // fragment-offset value should be considered. The type is interface{} with
    // range: 0..8191.
    FragmentOffset2 interface{}
}

func (fragmentOffset *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_FragmentOffset) GetEntityData() *types.CommonEntityData {
    fragmentOffset.EntityData.YFilter = fragmentOffset.YFilter
    fragmentOffset.EntityData.YangName = "fragment-offset"
    fragmentOffset.EntityData.BundleName = "cisco_ios_xr"
    fragmentOffset.EntityData.ParentYangName = "access-list-entry"
    fragmentOffset.EntityData.SegmentPath = "fragment-offset"
    fragmentOffset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fragmentOffset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fragmentOffset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fragmentOffset.EntityData.Children = make(map[string]types.YChild)
    fragmentOffset.EntityData.Leafs = make(map[string]types.YLeaf)
    fragmentOffset.EntityData.Leafs["fragment-offset-operator"] = types.YLeaf{"FragmentOffsetOperator", fragmentOffset.FragmentOffsetOperator}
    fragmentOffset.EntityData.Leafs["fragment-offset-1"] = types.YLeaf{"FragmentOffset1", fragmentOffset.FragmentOffset1}
    fragmentOffset.EntityData.Leafs["fragment-offset-2"] = types.YLeaf{"FragmentOffset2", fragmentOffset.FragmentOffset2}
    return &(fragmentOffset.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop
// Next-hop settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop struct {
    EntityData types.CommonEntityData
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

func (nextHop *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "access-list-entry"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Children["next-hop-1"] = types.YChild{"NextHop1", &nextHop.NextHop1}
    nextHop.EntityData.Children["next-hop-2"] = types.YChild{"NextHop2", &nextHop.NextHop2}
    nextHop.EntityData.Children["next-hop-3"] = types.YChild{"NextHop3", &nextHop.NextHop3}
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["next-hop-type"] = types.YLeaf{"NextHopType", nextHop.NextHopType}
    return &(nextHop.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1
// The first next-hop settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IPv4 address of the next-hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // The VRF name of the next-hop. The type is string.
    VrfName interface{}

    // The object tracking name for the next-hop. The type is string.
    TrackName interface{}
}

func (nextHop1 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetEntityData() *types.CommonEntityData {
    nextHop1.EntityData.YFilter = nextHop1.YFilter
    nextHop1.EntityData.YangName = "next-hop-1"
    nextHop1.EntityData.BundleName = "cisco_ios_xr"
    nextHop1.EntityData.ParentYangName = "next-hop"
    nextHop1.EntityData.SegmentPath = "next-hop-1"
    nextHop1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop1.EntityData.Children = make(map[string]types.YChild)
    nextHop1.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop1.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", nextHop1.NextHop}
    nextHop1.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", nextHop1.VrfName}
    nextHop1.EntityData.Leafs["track-name"] = types.YLeaf{"TrackName", nextHop1.TrackName}
    return &(nextHop1.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2
// The second next-hop settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IPv4 address of the next-hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // The VRF name of the next-hop. The type is string.
    VrfName interface{}

    // The object tracking name for the next-hop. The type is string.
    TrackName interface{}
}

func (nextHop2 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetEntityData() *types.CommonEntityData {
    nextHop2.EntityData.YFilter = nextHop2.YFilter
    nextHop2.EntityData.YangName = "next-hop-2"
    nextHop2.EntityData.BundleName = "cisco_ios_xr"
    nextHop2.EntityData.ParentYangName = "next-hop"
    nextHop2.EntityData.SegmentPath = "next-hop-2"
    nextHop2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop2.EntityData.Children = make(map[string]types.YChild)
    nextHop2.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop2.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", nextHop2.NextHop}
    nextHop2.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", nextHop2.VrfName}
    nextHop2.EntityData.Leafs["track-name"] = types.YLeaf{"TrackName", nextHop2.TrackName}
    return &(nextHop2.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3
// The third next-hop settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IPv4 address of the next-hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // The VRF name of the next-hop. The type is string.
    VrfName interface{}

    // The object tracking name for the next-hop. The type is string.
    TrackName interface{}
}

func (nextHop3 *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetEntityData() *types.CommonEntityData {
    nextHop3.EntityData.YFilter = nextHop3.YFilter
    nextHop3.EntityData.YangName = "next-hop-3"
    nextHop3.EntityData.BundleName = "cisco_ios_xr"
    nextHop3.EntityData.ParentYangName = "next-hop"
    nextHop3.EntityData.SegmentPath = "next-hop-3"
    nextHop3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop3.EntityData.Children = make(map[string]types.YChild)
    nextHop3.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop3.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", nextHop3.NextHop}
    nextHop3.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", nextHop3.VrfName}
    nextHop3.EntityData.Leafs["track-name"] = types.YLeaf{"TrackName", nextHop3.TrackName}
    return &(nextHop3.EntityData)
}

// Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp
// DSCP settings.
type Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp struct {
    EntityData types.CommonEntityData
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

func (dscp *Ipv4AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Dscp) GetEntityData() *types.CommonEntityData {
    dscp.EntityData.YFilter = dscp.YFilter
    dscp.EntityData.YangName = "dscp"
    dscp.EntityData.BundleName = "cisco_ios_xr"
    dscp.EntityData.ParentYangName = "access-list-entry"
    dscp.EntityData.SegmentPath = "dscp"
    dscp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dscp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dscp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dscp.EntityData.Children = make(map[string]types.YChild)
    dscp.EntityData.Leafs = make(map[string]types.YLeaf)
    dscp.EntityData.Leafs["dscp-operator"] = types.YLeaf{"DscpOperator", dscp.DscpOperator}
    dscp.EntityData.Leafs["dscp-min"] = types.YLeaf{"DscpMin", dscp.DscpMin}
    dscp.EntityData.Leafs["dscp-max"] = types.YLeaf{"DscpMax", dscp.DscpMax}
    return &(dscp.EntityData)
}

// Ipv4AclAndPrefixList_Prefixes
// Table of ACL prefix lists.  Entries in this
// table and the PrefixListExistenceTable table
// must be kept consistent
type Ipv4AclAndPrefixList_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of a prefix list. The type is slice of
    // Ipv4AclAndPrefixList_Prefixes_Prefix.
    Prefix []Ipv4AclAndPrefixList_Prefixes_Prefix
}

func (prefixes *Ipv4AclAndPrefixList_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "cisco_ios_xr"
    prefixes.EntityData.ParentYangName = "ipv4-acl-and-prefix-list"
    prefixes.EntityData.SegmentPath = "prefixes"
    prefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixes.EntityData.Children = make(map[string]types.YChild)
    prefixes.EntityData.Children["prefix"] = types.YChild{"Prefix", nil}
    for i := range prefixes.Prefix {
        prefixes.EntityData.Children[types.GetSegmentPath(&prefixes.Prefix[i])] = types.YChild{"Prefix", &prefixes.Prefix[i]}
    }
    prefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixes.EntityData)
}

// Ipv4AclAndPrefixList_Prefixes_Prefix
// Name of a prefix list
type Ipv4AclAndPrefixList_Prefixes_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Prefix list name - max 32 characters. The type is
    // string.
    PrefixListName interface{}

    // Sequence of entries forming a prefix list.
    PrefixListEntries Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries
}

func (prefix *Ipv4AclAndPrefixList_Prefixes_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefixes"
    prefix.EntityData.SegmentPath = "prefix" + "[prefix-list-name='" + fmt.Sprintf("%v", prefix.PrefixListName) + "']"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Children["prefix-list-entries"] = types.YChild{"PrefixListEntries", &prefix.PrefixListEntries}
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["prefix-list-name"] = types.YLeaf{"PrefixListName", prefix.PrefixListName}
    return &(prefix.EntityData)
}

// Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries
// Sequence of entries forming a prefix list
// This type is a presence type.
type Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A prefix list entry; either a description (remark) or a prefix to match
    // against. The type is slice of
    // Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry.
    PrefixListEntry []Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry
}

func (prefixListEntries *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetEntityData() *types.CommonEntityData {
    prefixListEntries.EntityData.YFilter = prefixListEntries.YFilter
    prefixListEntries.EntityData.YangName = "prefix-list-entries"
    prefixListEntries.EntityData.BundleName = "cisco_ios_xr"
    prefixListEntries.EntityData.ParentYangName = "prefix"
    prefixListEntries.EntityData.SegmentPath = "prefix-list-entries"
    prefixListEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixListEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixListEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixListEntries.EntityData.Children = make(map[string]types.YChild)
    prefixListEntries.EntityData.Children["prefix-list-entry"] = types.YChild{"PrefixListEntry", nil}
    for i := range prefixListEntries.PrefixListEntry {
        prefixListEntries.EntityData.Children[types.GetSegmentPath(&prefixListEntries.PrefixListEntry[i])] = types.YChild{"PrefixListEntry", &prefixListEntries.PrefixListEntry[i]}
    }
    prefixListEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixListEntries.EntityData)
}

// Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry
// A prefix list entry; either a description
// (remark) or a prefix to match against
type Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence number of prefix list. The type is
    // interface{} with range: 1..2147483646.
    SequenceNumber interface{}

    // Whether to forward or drop packets matching the prefix list. The type is
    // Ipv4AclGrantEnum.
    Grant interface{}

    // IPv4 address prefix to match. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Mask of IPv4 address prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (prefixListEntry *Ipv4AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetEntityData() *types.CommonEntityData {
    prefixListEntry.EntityData.YFilter = prefixListEntry.YFilter
    prefixListEntry.EntityData.YangName = "prefix-list-entry"
    prefixListEntry.EntityData.BundleName = "cisco_ios_xr"
    prefixListEntry.EntityData.ParentYangName = "prefix-list-entries"
    prefixListEntry.EntityData.SegmentPath = "prefix-list-entry" + "[sequence-number='" + fmt.Sprintf("%v", prefixListEntry.SequenceNumber) + "']"
    prefixListEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixListEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixListEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixListEntry.EntityData.Children = make(map[string]types.YChild)
    prefixListEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixListEntry.EntityData.Leafs["sequence-number"] = types.YLeaf{"SequenceNumber", prefixListEntry.SequenceNumber}
    prefixListEntry.EntityData.Leafs["grant"] = types.YLeaf{"Grant", prefixListEntry.Grant}
    prefixListEntry.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", prefixListEntry.Prefix}
    prefixListEntry.EntityData.Leafs["netmask"] = types.YLeaf{"Netmask", prefixListEntry.Netmask}
    prefixListEntry.EntityData.Leafs["match-exact-length"] = types.YLeaf{"MatchExactLength", prefixListEntry.MatchExactLength}
    prefixListEntry.EntityData.Leafs["exact-prefix-length"] = types.YLeaf{"ExactPrefixLength", prefixListEntry.ExactPrefixLength}
    prefixListEntry.EntityData.Leafs["match-max-length"] = types.YLeaf{"MatchMaxLength", prefixListEntry.MatchMaxLength}
    prefixListEntry.EntityData.Leafs["max-prefix-length"] = types.YLeaf{"MaxPrefixLength", prefixListEntry.MaxPrefixLength}
    prefixListEntry.EntityData.Leafs["match-min-length"] = types.YLeaf{"MatchMinLength", prefixListEntry.MatchMinLength}
    prefixListEntry.EntityData.Leafs["min-prefix-length"] = types.YLeaf{"MinPrefixLength", prefixListEntry.MinPrefixLength}
    prefixListEntry.EntityData.Leafs["remark"] = types.YLeaf{"Remark", prefixListEntry.Remark}
    return &(prefixListEntry.EntityData)
}

// Ipv4AclAndPrefixList_LogUpdate
// Control access lists log updates
type Ipv4AclAndPrefixList_LogUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log update threshold (number of hits). The type is interface{} with range:
    // 1..2147483647.
    Threshold interface{}

    // Log update rate (log msgs per second). The type is interface{} with range:
    // 1..1000.
    Rate interface{}
}

func (logUpdate *Ipv4AclAndPrefixList_LogUpdate) GetEntityData() *types.CommonEntityData {
    logUpdate.EntityData.YFilter = logUpdate.YFilter
    logUpdate.EntityData.YangName = "log-update"
    logUpdate.EntityData.BundleName = "cisco_ios_xr"
    logUpdate.EntityData.ParentYangName = "ipv4-acl-and-prefix-list"
    logUpdate.EntityData.SegmentPath = "log-update"
    logUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logUpdate.EntityData.Children = make(map[string]types.YChild)
    logUpdate.EntityData.Leafs = make(map[string]types.YLeaf)
    logUpdate.EntityData.Leafs["threshold"] = types.YLeaf{"Threshold", logUpdate.Threshold}
    logUpdate.EntityData.Leafs["rate"] = types.YLeaf{"Rate", logUpdate.Rate}
    return &(logUpdate.EntityData)
}

