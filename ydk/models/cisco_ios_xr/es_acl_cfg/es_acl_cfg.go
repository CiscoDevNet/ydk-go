// This module contains a collection of YANG definitions
// for Cisco IOS-XR es-acl package configuration.
// 
// This module contains definitions
// for the following management objects:
//   es-acl: Layer 2 ACL configuration data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package es_acl_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package es_acl_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-es-acl-cfg es-acl}", reflect.TypeOf(EsAcl{}))
    ydk.RegisterEntity("Cisco-IOS-XR-es-acl-cfg:es-acl", reflect.TypeOf(EsAcl{}))
}

// EsAclGrantEnum represents ES ACL forwarding action.
type EsAclGrantEnum string

const (
    // Deny
    EsAclGrantEnum_deny EsAclGrantEnum = "deny"

    // Permit
    EsAclGrantEnum_permit EsAclGrantEnum = "permit"
)

// EsAcl
// Layer 2 ACL configuration data
type EsAcl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of access lists.
    Accesses EsAcl_Accesses
}

func (esAcl *EsAcl) GetEntityData() *types.CommonEntityData {
    esAcl.EntityData.YFilter = esAcl.YFilter
    esAcl.EntityData.YangName = "es-acl"
    esAcl.EntityData.BundleName = "cisco_ios_xr"
    esAcl.EntityData.ParentYangName = "Cisco-IOS-XR-es-acl-cfg"
    esAcl.EntityData.SegmentPath = "Cisco-IOS-XR-es-acl-cfg:es-acl"
    esAcl.EntityData.AbsolutePath = esAcl.EntityData.SegmentPath
    esAcl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esAcl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esAcl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esAcl.EntityData.Children = types.NewOrderedMap()
    esAcl.EntityData.Children.Append("accesses", types.YChild{"Accesses", &esAcl.Accesses})
    esAcl.EntityData.Leafs = types.NewOrderedMap()

    esAcl.EntityData.YListKeys = []string {}

    return &(esAcl.EntityData)
}

// EsAcl_Accesses
// Table of access lists
type EsAcl_Accesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An ACL. The type is slice of EsAcl_Accesses_Access.
    Access []*EsAcl_Accesses_Access
}

func (accesses *EsAcl_Accesses) GetEntityData() *types.CommonEntityData {
    accesses.EntityData.YFilter = accesses.YFilter
    accesses.EntityData.YangName = "accesses"
    accesses.EntityData.BundleName = "cisco_ios_xr"
    accesses.EntityData.ParentYangName = "es-acl"
    accesses.EntityData.SegmentPath = "accesses"
    accesses.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-cfg:es-acl/" + accesses.EntityData.SegmentPath
    accesses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accesses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accesses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accesses.EntityData.Children = types.NewOrderedMap()
    accesses.EntityData.Children.Append("access", types.YChild{"Access", nil})
    for i := range accesses.Access {
        accesses.EntityData.Children.Append(types.GetSegmentPath(accesses.Access[i]), types.YChild{"Access", accesses.Access[i]})
    }
    accesses.EntityData.Leafs = types.NewOrderedMap()

    accesses.EntityData.YListKeys = []string {}

    return &(accesses.EntityData)
}

// EsAcl_Accesses_Access
// An ACL
type EsAcl_Accesses_Access struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the access list. The type is string with
    // length: 1..64.
    Name interface{}

    // ACL entry table; contains list of access list entries.
    AccessListEntries EsAcl_Accesses_Access_AccessListEntries
}

func (access *EsAcl_Accesses_Access) GetEntityData() *types.CommonEntityData {
    access.EntityData.YFilter = access.YFilter
    access.EntityData.YangName = "access"
    access.EntityData.BundleName = "cisco_ios_xr"
    access.EntityData.ParentYangName = "accesses"
    access.EntityData.SegmentPath = "access" + types.AddKeyToken(access.Name, "name")
    access.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-cfg:es-acl/accesses/" + access.EntityData.SegmentPath
    access.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    access.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    access.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    access.EntityData.Children = types.NewOrderedMap()
    access.EntityData.Children.Append("access-list-entries", types.YChild{"AccessListEntries", &access.AccessListEntries})
    access.EntityData.Leafs = types.NewOrderedMap()
    access.EntityData.Leafs.Append("name", types.YLeaf{"Name", access.Name})

    access.EntityData.YListKeys = []string {"Name"}

    return &(access.EntityData)
}

// EsAcl_Accesses_Access_AccessListEntries
// ACL entry table; contains list of access list
// entries
type EsAcl_Accesses_Access_AccessListEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An ACL entry; either a description (remark) or anAccess List Entry to match
    // against. The type is slice of
    // EsAcl_Accesses_Access_AccessListEntries_AccessListEntry.
    AccessListEntry []*EsAcl_Accesses_Access_AccessListEntries_AccessListEntry
}

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetEntityData() *types.CommonEntityData {
    accessListEntries.EntityData.YFilter = accessListEntries.YFilter
    accessListEntries.EntityData.YangName = "access-list-entries"
    accessListEntries.EntityData.BundleName = "cisco_ios_xr"
    accessListEntries.EntityData.ParentYangName = "access"
    accessListEntries.EntityData.SegmentPath = "access-list-entries"
    accessListEntries.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-cfg:es-acl/accesses/access/" + accessListEntries.EntityData.SegmentPath
    accessListEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListEntries.EntityData.Children = types.NewOrderedMap()
    accessListEntries.EntityData.Children.Append("access-list-entry", types.YChild{"AccessListEntry", nil})
    for i := range accessListEntries.AccessListEntry {
        accessListEntries.EntityData.Children.Append(types.GetSegmentPath(accessListEntries.AccessListEntry[i]), types.YChild{"AccessListEntry", accessListEntries.AccessListEntry[i]})
    }
    accessListEntries.EntityData.Leafs = types.NewOrderedMap()

    accessListEntries.EntityData.YListKeys = []string {}

    return &(accessListEntries.EntityData)
}

// EsAcl_Accesses_Access_AccessListEntries_AccessListEntry
// An ACL entry; either a description (remark)
// or anAccess List Entry to match against
type EsAcl_Accesses_Access_AccessListEntries_AccessListEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sequence number of access list entry. The type is
    // interface{} with range: 1..2147483646.
    SequenceNumber interface{}

    // Forwarding action for the packet. This is required  for any non-remark ACE.
    // Leave unspecified otherwise. The type is EsAclGrantEnum.
    Grant interface{}

    // This 12-bit VLAN-ID in the VLAN Tag header uniquely identifies the VLAN. It
    // can be used for the lower bound (in range) or single value. Any value not
    // in the permissible range will be rejected. The type is interface{} with
    // range: 0..4095.
    Vlan1 interface{}

    // This 12 bit VLAN-ID in the VLAN Tag header uniquely identifies the VLAN. It
    // is used in the upper bound (in range). Any value not in the permissible
    // range will be rejected. The type is interface{} with range: 0..4095.
    Vlan2 interface{}

    // Class of Service value. Any value not in the  permissible range will be
    // rejected. The type is interface{} with range: 0..7.
    Cos interface{}

    // Discard Eligibility Indication bit. User can specify  1 to indicate the bit
    // is set. Leave unspecified otherwise. The type is interface{} with range:
    // 0..1.
    Dei interface{}

    // This represents the QinQ vlan identifier. It can be used for the lower
    // bound (in range) or single value. Any value not  in the permissible range
    // will be rejected. The type is interface{} with range: 0..4095.
    InnerVlan1 interface{}

    // This represents the QinQ vlan identifier. It is used in  the upper bound
    // (in range). Any value not in the permissible range will be rejected. The
    // type is interface{} with range: 0..4095.
    InnerVlan2 interface{}

    // Class of Service of Inner Header. Any value not in the permissible range
    // will be rejected. The type is interface{} with range: 0..7.
    InnerCos interface{}

    // Discard Eligibility Indication for Inner Header. User  can specify 1 to
    // indicate the bit is set. Leave  unspecified otherwise. The type is
    // interface{} with range: 0..1.
    InnerDei interface{}

    // Description for the access-list-entry/rule. The type is string with length:
    // 0..255.
    Remark interface{}

    // Ethernet type Number in Hex. Any value not in the permissible range will be
    // rejected. The type is string with pattern: b'[0-9a-fA-F]{1,4}'.
    EtherTypeNumber interface{}

    // Enable capture if set to TRUE. The type is bool.
    Capture interface{}

    // Log the packet on this access-list-entry/rule. User can specify 1 to enable
    // logging the match, leave unspecified otherwise. The type is interface{}
    // with range: 0..1.
    LogOption interface{}

    // Sequence String for the ace. The type is string with length: 1..64.
    SequenceStr interface{}

    // Source network settings.
    SourceNetwork EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork

    // Destination network settings.
    DestinationNetwork EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork
}

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetEntityData() *types.CommonEntityData {
    accessListEntry.EntityData.YFilter = accessListEntry.YFilter
    accessListEntry.EntityData.YangName = "access-list-entry"
    accessListEntry.EntityData.BundleName = "cisco_ios_xr"
    accessListEntry.EntityData.ParentYangName = "access-list-entries"
    accessListEntry.EntityData.SegmentPath = "access-list-entry" + types.AddKeyToken(accessListEntry.SequenceNumber, "sequence-number")
    accessListEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-cfg:es-acl/accesses/access/access-list-entries/" + accessListEntry.EntityData.SegmentPath
    accessListEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListEntry.EntityData.Children = types.NewOrderedMap()
    accessListEntry.EntityData.Children.Append("source-network", types.YChild{"SourceNetwork", &accessListEntry.SourceNetwork})
    accessListEntry.EntityData.Children.Append("destination-network", types.YChild{"DestinationNetwork", &accessListEntry.DestinationNetwork})
    accessListEntry.EntityData.Leafs = types.NewOrderedMap()
    accessListEntry.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", accessListEntry.SequenceNumber})
    accessListEntry.EntityData.Leafs.Append("grant", types.YLeaf{"Grant", accessListEntry.Grant})
    accessListEntry.EntityData.Leafs.Append("vlan1", types.YLeaf{"Vlan1", accessListEntry.Vlan1})
    accessListEntry.EntityData.Leafs.Append("vlan2", types.YLeaf{"Vlan2", accessListEntry.Vlan2})
    accessListEntry.EntityData.Leafs.Append("cos", types.YLeaf{"Cos", accessListEntry.Cos})
    accessListEntry.EntityData.Leafs.Append("dei", types.YLeaf{"Dei", accessListEntry.Dei})
    accessListEntry.EntityData.Leafs.Append("inner-vlan1", types.YLeaf{"InnerVlan1", accessListEntry.InnerVlan1})
    accessListEntry.EntityData.Leafs.Append("inner-vlan2", types.YLeaf{"InnerVlan2", accessListEntry.InnerVlan2})
    accessListEntry.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", accessListEntry.InnerCos})
    accessListEntry.EntityData.Leafs.Append("inner-dei", types.YLeaf{"InnerDei", accessListEntry.InnerDei})
    accessListEntry.EntityData.Leafs.Append("remark", types.YLeaf{"Remark", accessListEntry.Remark})
    accessListEntry.EntityData.Leafs.Append("ether-type-number", types.YLeaf{"EtherTypeNumber", accessListEntry.EtherTypeNumber})
    accessListEntry.EntityData.Leafs.Append("capture", types.YLeaf{"Capture", accessListEntry.Capture})
    accessListEntry.EntityData.Leafs.Append("log-option", types.YLeaf{"LogOption", accessListEntry.LogOption})
    accessListEntry.EntityData.Leafs.Append("sequence-str", types.YLeaf{"SequenceStr", accessListEntry.SequenceStr})

    accessListEntry.EntityData.YListKeys = []string {"SequenceNumber"}

    return &(accessListEntry.EntityData)
}

// EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork
// Source network settings.
type EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source address to match, leave unspecified for any. The type is string with
    // pattern: b'([0-9a-fA-F]{1,4}(\\.[0-9a-fA-F]{1,4}){2})'.
    SourceAddress interface{}

    // Wildcard bits to apply to source address (if specified), leave unspecified
    // for no wildcarding. The type is string with pattern:
    // b'([0-9a-fA-F]{1,4}(\\.[0-9a-fA-F]{1,4}){2})'.
    SourceWildCardBits interface{}
}

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetEntityData() *types.CommonEntityData {
    sourceNetwork.EntityData.YFilter = sourceNetwork.YFilter
    sourceNetwork.EntityData.YangName = "source-network"
    sourceNetwork.EntityData.BundleName = "cisco_ios_xr"
    sourceNetwork.EntityData.ParentYangName = "access-list-entry"
    sourceNetwork.EntityData.SegmentPath = "source-network"
    sourceNetwork.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-cfg:es-acl/accesses/access/access-list-entries/access-list-entry/" + sourceNetwork.EntityData.SegmentPath
    sourceNetwork.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceNetwork.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceNetwork.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceNetwork.EntityData.Children = types.NewOrderedMap()
    sourceNetwork.EntityData.Leafs = types.NewOrderedMap()
    sourceNetwork.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", sourceNetwork.SourceAddress})
    sourceNetwork.EntityData.Leafs.Append("source-wild-card-bits", types.YLeaf{"SourceWildCardBits", sourceNetwork.SourceWildCardBits})

    sourceNetwork.EntityData.YListKeys = []string {}

    return &(sourceNetwork.EntityData)
}

// EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork
// Destination network settings.
type EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination address to match (if a protocol was specified), leave
    // unspecified for any. The type is string with pattern:
    // b'([0-9a-fA-F]{1,4}(\\.[0-9a-fA-F]{1,4}){2})'.
    DestinationAddress interface{}

    // Wildcard bits to apply to destination address (if specified), leave
    // unspecified for no wildcarding. The type is string with pattern:
    // b'([0-9a-fA-F]{1,4}(\\.[0-9a-fA-F]{1,4}){2})'.
    DestinationWildCardBits interface{}
}

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetEntityData() *types.CommonEntityData {
    destinationNetwork.EntityData.YFilter = destinationNetwork.YFilter
    destinationNetwork.EntityData.YangName = "destination-network"
    destinationNetwork.EntityData.BundleName = "cisco_ios_xr"
    destinationNetwork.EntityData.ParentYangName = "access-list-entry"
    destinationNetwork.EntityData.SegmentPath = "destination-network"
    destinationNetwork.EntityData.AbsolutePath = "Cisco-IOS-XR-es-acl-cfg:es-acl/accesses/access/access-list-entries/access-list-entry/" + destinationNetwork.EntityData.SegmentPath
    destinationNetwork.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationNetwork.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationNetwork.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationNetwork.EntityData.Children = types.NewOrderedMap()
    destinationNetwork.EntityData.Leafs = types.NewOrderedMap()
    destinationNetwork.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", destinationNetwork.DestinationAddress})
    destinationNetwork.EntityData.Leafs.Append("destination-wild-card-bits", types.YLeaf{"DestinationWildCardBits", destinationNetwork.DestinationWildCardBits})

    destinationNetwork.EntityData.YListKeys = []string {}

    return &(destinationNetwork.EntityData)
}

