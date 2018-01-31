// This module contains a collection of YANG definitions
// for Cisco IOS-XR es-acl package configuration.
// 
// This module contains definitions
// for the following management objects:
//   es-acl: Layer 2 ACL configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

// EsAclGrantEnum represents ES acl grant enum
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of access lists.
    Accesses EsAcl_Accesses
}

func (esAcl *EsAcl) GetFilter() yfilter.YFilter { return esAcl.YFilter }

func (esAcl *EsAcl) SetFilter(yf yfilter.YFilter) { esAcl.YFilter = yf }

func (esAcl *EsAcl) GetGoName(yname string) string {
    if yname == "accesses" { return "Accesses" }
    return ""
}

func (esAcl *EsAcl) GetSegmentPath() string {
    return "Cisco-IOS-XR-es-acl-cfg:es-acl"
}

func (esAcl *EsAcl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accesses" {
        return &esAcl.Accesses
    }
    return nil
}

func (esAcl *EsAcl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accesses"] = &esAcl.Accesses
    return children
}

func (esAcl *EsAcl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (esAcl *EsAcl) GetBundleName() string { return "cisco_ios_xr" }

func (esAcl *EsAcl) GetYangName() string { return "es-acl" }

func (esAcl *EsAcl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (esAcl *EsAcl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (esAcl *EsAcl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (esAcl *EsAcl) SetParent(parent types.Entity) { esAcl.parent = parent }

func (esAcl *EsAcl) GetParent() types.Entity { return esAcl.parent }

func (esAcl *EsAcl) GetParentYangName() string { return "Cisco-IOS-XR-es-acl-cfg" }

// EsAcl_Accesses
// Table of access lists
type EsAcl_Accesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An ACL. The type is slice of EsAcl_Accesses_Access.
    Access []EsAcl_Accesses_Access
}

func (accesses *EsAcl_Accesses) GetFilter() yfilter.YFilter { return accesses.YFilter }

func (accesses *EsAcl_Accesses) SetFilter(yf yfilter.YFilter) { accesses.YFilter = yf }

func (accesses *EsAcl_Accesses) GetGoName(yname string) string {
    if yname == "access" { return "Access" }
    return ""
}

func (accesses *EsAcl_Accesses) GetSegmentPath() string {
    return "accesses"
}

func (accesses *EsAcl_Accesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access" {
        for _, c := range accesses.Access {
            if accesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EsAcl_Accesses_Access{}
        accesses.Access = append(accesses.Access, child)
        return &accesses.Access[len(accesses.Access)-1]
    }
    return nil
}

func (accesses *EsAcl_Accesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accesses.Access {
        children[accesses.Access[i].GetSegmentPath()] = &accesses.Access[i]
    }
    return children
}

func (accesses *EsAcl_Accesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accesses *EsAcl_Accesses) GetBundleName() string { return "cisco_ios_xr" }

func (accesses *EsAcl_Accesses) GetYangName() string { return "accesses" }

func (accesses *EsAcl_Accesses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accesses *EsAcl_Accesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accesses *EsAcl_Accesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accesses *EsAcl_Accesses) SetParent(parent types.Entity) { accesses.parent = parent }

func (accesses *EsAcl_Accesses) GetParent() types.Entity { return accesses.parent }

func (accesses *EsAcl_Accesses) GetParentYangName() string { return "es-acl" }

// EsAcl_Accesses_Access
// An ACL
type EsAcl_Accesses_Access struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the access list. The type is string with
    // length: 1..65.
    Name interface{}

    // ACL entry table; contains list of access list entries.
    AccessListEntries EsAcl_Accesses_Access_AccessListEntries
}

func (access *EsAcl_Accesses_Access) GetFilter() yfilter.YFilter { return access.YFilter }

func (access *EsAcl_Accesses_Access) SetFilter(yf yfilter.YFilter) { access.YFilter = yf }

func (access *EsAcl_Accesses_Access) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "access-list-entries" { return "AccessListEntries" }
    return ""
}

func (access *EsAcl_Accesses_Access) GetSegmentPath() string {
    return "access" + "[name='" + fmt.Sprintf("%v", access.Name) + "']"
}

func (access *EsAcl_Accesses_Access) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-entries" {
        return &access.AccessListEntries
    }
    return nil
}

func (access *EsAcl_Accesses_Access) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list-entries"] = &access.AccessListEntries
    return children
}

func (access *EsAcl_Accesses_Access) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = access.Name
    return leafs
}

func (access *EsAcl_Accesses_Access) GetBundleName() string { return "cisco_ios_xr" }

func (access *EsAcl_Accesses_Access) GetYangName() string { return "access" }

func (access *EsAcl_Accesses_Access) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (access *EsAcl_Accesses_Access) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (access *EsAcl_Accesses_Access) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (access *EsAcl_Accesses_Access) SetParent(parent types.Entity) { access.parent = parent }

func (access *EsAcl_Accesses_Access) GetParent() types.Entity { return access.parent }

func (access *EsAcl_Accesses_Access) GetParentYangName() string { return "accesses" }

// EsAcl_Accesses_Access_AccessListEntries
// ACL entry table; contains list of access list
// entries
type EsAcl_Accesses_Access_AccessListEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An ACL entry; either a description (remark) or anAccess List Entry to match
    // against. The type is slice of
    // EsAcl_Accesses_Access_AccessListEntries_AccessListEntry.
    AccessListEntry []EsAcl_Accesses_Access_AccessListEntries_AccessListEntry
}

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetFilter() yfilter.YFilter { return accessListEntries.YFilter }

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) SetFilter(yf yfilter.YFilter) { accessListEntries.YFilter = yf }

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetGoName(yname string) string {
    if yname == "access-list-entry" { return "AccessListEntry" }
    return ""
}

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetSegmentPath() string {
    return "access-list-entries"
}

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-entry" {
        for _, c := range accessListEntries.AccessListEntry {
            if accessListEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EsAcl_Accesses_Access_AccessListEntries_AccessListEntry{}
        accessListEntries.AccessListEntry = append(accessListEntries.AccessListEntry, child)
        return &accessListEntries.AccessListEntry[len(accessListEntries.AccessListEntry)-1]
    }
    return nil
}

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessListEntries.AccessListEntry {
        children[accessListEntries.AccessListEntry[i].GetSegmentPath()] = &accessListEntries.AccessListEntry[i]
    }
    return children
}

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetBundleName() string { return "cisco_ios_xr" }

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetYangName() string { return "access-list-entries" }

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) SetParent(parent types.Entity) { accessListEntries.parent = parent }

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetParent() types.Entity { return accessListEntries.parent }

func (accessListEntries *EsAcl_Accesses_Access_AccessListEntries) GetParentYangName() string { return "access" }

// EsAcl_Accesses_Access_AccessListEntries_AccessListEntry
// An ACL entry; either a description (remark)
// or anAccess List Entry to match against
type EsAcl_Accesses_Access_AccessListEntries_AccessListEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence number of access list entry. The type is
    // interface{} with range: 1..2147483646.
    SequenceNumber interface{}

    // Whether to forward or drop packets matching the ACE. The type is
    // EsAclGrantEnum.
    Grant interface{}

    // VLAN ID/range lower limit. The type is interface{} with range: 0..65535.
    Vlan1 interface{}

    // VLAN ID range higher limit. The type is interface{} with range: 0..65535.
    Vlan2 interface{}

    // COS value. The type is interface{} with range: 0..255.
    Cos interface{}

    // DEI bit. The type is interface{} with range: 0..255.
    Dei interface{}

    // Inner VLAN ID/range lower limit. The type is interface{} with range:
    // 0..65535.
    InnerVlan1 interface{}

    // Inner VLAN ID range higher limit. The type is interface{} with range:
    // 0..65535.
    InnerVlan2 interface{}

    // Inner COS value. The type is interface{} with range: 0..255.
    InnerCos interface{}

    // Inner DEI bit. The type is interface{} with range: 0..255.
    InnerDei interface{}

    // Comments or a description for the access list. The type is string.
    Remark interface{}

    // Ethernet type Number. The type is interface{} with range: 0..65535.
    EtherTypeNumber interface{}

    // Enable capture. The type is bool.
    Capture interface{}

    // Whether and how to log matches against this entry. The type is interface{}
    // with range: 0..255.
    LogOption interface{}

    // Sequence String for the ace. The type is string with length: 1..64.
    SequenceStr interface{}

    // Source network settings.
    SourceNetwork EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork

    // Destination network settings.
    DestinationNetwork EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork
}

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetFilter() yfilter.YFilter { return accessListEntry.YFilter }

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) SetFilter(yf yfilter.YFilter) { accessListEntry.YFilter = yf }

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "grant" { return "Grant" }
    if yname == "vlan1" { return "Vlan1" }
    if yname == "vlan2" { return "Vlan2" }
    if yname == "cos" { return "Cos" }
    if yname == "dei" { return "Dei" }
    if yname == "inner-vlan1" { return "InnerVlan1" }
    if yname == "inner-vlan2" { return "InnerVlan2" }
    if yname == "inner-cos" { return "InnerCos" }
    if yname == "inner-dei" { return "InnerDei" }
    if yname == "remark" { return "Remark" }
    if yname == "ether-type-number" { return "EtherTypeNumber" }
    if yname == "capture" { return "Capture" }
    if yname == "log-option" { return "LogOption" }
    if yname == "sequence-str" { return "SequenceStr" }
    if yname == "source-network" { return "SourceNetwork" }
    if yname == "destination-network" { return "DestinationNetwork" }
    return ""
}

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetSegmentPath() string {
    return "access-list-entry" + "[sequence-number='" + fmt.Sprintf("%v", accessListEntry.SequenceNumber) + "']"
}

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source-network" {
        return &accessListEntry.SourceNetwork
    }
    if childYangName == "destination-network" {
        return &accessListEntry.DestinationNetwork
    }
    return nil
}

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source-network"] = &accessListEntry.SourceNetwork
    children["destination-network"] = &accessListEntry.DestinationNetwork
    return children
}

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = accessListEntry.SequenceNumber
    leafs["grant"] = accessListEntry.Grant
    leafs["vlan1"] = accessListEntry.Vlan1
    leafs["vlan2"] = accessListEntry.Vlan2
    leafs["cos"] = accessListEntry.Cos
    leafs["dei"] = accessListEntry.Dei
    leafs["inner-vlan1"] = accessListEntry.InnerVlan1
    leafs["inner-vlan2"] = accessListEntry.InnerVlan2
    leafs["inner-cos"] = accessListEntry.InnerCos
    leafs["inner-dei"] = accessListEntry.InnerDei
    leafs["remark"] = accessListEntry.Remark
    leafs["ether-type-number"] = accessListEntry.EtherTypeNumber
    leafs["capture"] = accessListEntry.Capture
    leafs["log-option"] = accessListEntry.LogOption
    leafs["sequence-str"] = accessListEntry.SequenceStr
    return leafs
}

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetBundleName() string { return "cisco_ios_xr" }

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetYangName() string { return "access-list-entry" }

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) SetParent(parent types.Entity) { accessListEntry.parent = parent }

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetParent() types.Entity { return accessListEntry.parent }

func (accessListEntry *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry) GetParentYangName() string { return "access-list-entries" }

// EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork
// Source network settings.
type EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source address to match, leave unspecified for any. The type is string with
    // pattern: ([0-9a-fA-F]{1,4}(\.[0-9a-fA-F]{1,4}){2}).
    SourceAddress interface{}

    // Wildcard bits to apply to source address (if specified), leave unspecified
    // for no wildcarding. The type is string with pattern:
    // ([0-9a-fA-F]{1,4}(\.[0-9a-fA-F]{1,4}){2}).
    SourceWildCardBits interface{}
}

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetFilter() yfilter.YFilter { return sourceNetwork.YFilter }

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) SetFilter(yf yfilter.YFilter) { sourceNetwork.YFilter = yf }

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-wild-card-bits" { return "SourceWildCardBits" }
    return ""
}

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetSegmentPath() string {
    return "source-network"
}

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = sourceNetwork.SourceAddress
    leafs["source-wild-card-bits"] = sourceNetwork.SourceWildCardBits
    return leafs
}

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetBundleName() string { return "cisco_ios_xr" }

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetYangName() string { return "source-network" }

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) SetParent(parent types.Entity) { sourceNetwork.parent = parent }

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetParent() types.Entity { return sourceNetwork.parent }

func (sourceNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetParentYangName() string { return "access-list-entry" }

// EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork
// Destination network settings.
type EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination address to match (if a protocol was specified), leave
    // unspecified for any. The type is string with pattern:
    // ([0-9a-fA-F]{1,4}(\.[0-9a-fA-F]{1,4}){2}).
    DestinationAddress interface{}

    // Wildcard bits to apply to destination address (if specified), leave
    // unspecified for no wildcarding. The type is string with pattern:
    // ([0-9a-fA-F]{1,4}(\.[0-9a-fA-F]{1,4}){2}).
    DestinationWildCardBits interface{}
}

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetFilter() yfilter.YFilter { return destinationNetwork.YFilter }

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) SetFilter(yf yfilter.YFilter) { destinationNetwork.YFilter = yf }

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetGoName(yname string) string {
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-wild-card-bits" { return "DestinationWildCardBits" }
    return ""
}

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetSegmentPath() string {
    return "destination-network"
}

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-address"] = destinationNetwork.DestinationAddress
    leafs["destination-wild-card-bits"] = destinationNetwork.DestinationWildCardBits
    return leafs
}

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetBundleName() string { return "cisco_ios_xr" }

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetYangName() string { return "destination-network" }

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) SetParent(parent types.Entity) { destinationNetwork.parent = parent }

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetParent() types.Entity { return destinationNetwork.parent }

func (destinationNetwork *EsAcl_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetParentYangName() string { return "access-list-entry" }

