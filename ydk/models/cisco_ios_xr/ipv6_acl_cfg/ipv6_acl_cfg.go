// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-acl package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-acl-and-prefix-list: IPv6 ACL configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_acl_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_acl_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-acl-cfg ipv6-acl-and-prefix-list}", reflect.TypeOf(Ipv6AclAndPrefixList{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-acl-cfg:ipv6-acl-and-prefix-list", reflect.TypeOf(Ipv6AclAndPrefixList{}))
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

// Ipv6AclAndPrefixList
// IPv6 ACL configuration data
type Ipv6AclAndPrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of prefix lists.
    Prefixes Ipv6AclAndPrefixList_Prefixes

    // Control access lists log updates.
    LogUpdate Ipv6AclAndPrefixList_LogUpdate

    // Table of access lists.
    Accesses Ipv6AclAndPrefixList_Accesses
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetFilter() yfilter.YFilter { return ipv6AclAndPrefixList.YFilter }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) SetFilter(yf yfilter.YFilter) { ipv6AclAndPrefixList.YFilter = yf }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetGoName(yname string) string {
    if yname == "prefixes" { return "Prefixes" }
    if yname == "log-update" { return "LogUpdate" }
    if yname == "accesses" { return "Accesses" }
    return ""
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-acl-cfg:ipv6-acl-and-prefix-list"
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefixes" {
        return &ipv6AclAndPrefixList.Prefixes
    }
    if childYangName == "log-update" {
        return &ipv6AclAndPrefixList.LogUpdate
    }
    if childYangName == "accesses" {
        return &ipv6AclAndPrefixList.Accesses
    }
    return nil
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefixes"] = &ipv6AclAndPrefixList.Prefixes
    children["log-update"] = &ipv6AclAndPrefixList.LogUpdate
    children["accesses"] = &ipv6AclAndPrefixList.Accesses
    return children
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetYangName() string { return "ipv6-acl-and-prefix-list" }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) SetParent(parent types.Entity) { ipv6AclAndPrefixList.parent = parent }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetParent() types.Entity { return ipv6AclAndPrefixList.parent }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-acl-cfg" }

// Ipv6AclAndPrefixList_Prefixes
// Table of prefix lists
type Ipv6AclAndPrefixList_Prefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of a prefix list. The type is slice of
    // Ipv6AclAndPrefixList_Prefixes_Prefix.
    Prefix []Ipv6AclAndPrefixList_Prefixes_Prefix
}

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetFilter() yfilter.YFilter { return prefixes.YFilter }

func (prefixes *Ipv6AclAndPrefixList_Prefixes) SetFilter(yf yfilter.YFilter) { prefixes.YFilter = yf }

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetSegmentPath() string {
    return "prefixes"
}

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        for _, c := range prefixes.Prefix {
            if prefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_Prefixes_Prefix{}
        prefixes.Prefix = append(prefixes.Prefix, child)
        return &prefixes.Prefix[len(prefixes.Prefix)-1]
    }
    return nil
}

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixes.Prefix {
        children[prefixes.Prefix[i].GetSegmentPath()] = &prefixes.Prefix[i]
    }
    return children
}

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetBundleName() string { return "cisco_ios_xr" }

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetYangName() string { return "prefixes" }

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixes *Ipv6AclAndPrefixList_Prefixes) SetParent(parent types.Entity) { prefixes.parent = parent }

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetParent() types.Entity { return prefixes.parent }

func (prefixes *Ipv6AclAndPrefixList_Prefixes) GetParentYangName() string { return "ipv6-acl-and-prefix-list" }

// Ipv6AclAndPrefixList_Prefixes_Prefix
// Name of a prefix list
type Ipv6AclAndPrefixList_Prefixes_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of a prefix list. The type is string with
    // length: 1..65.
    Name interface{}

    // Sequence of entries forming a prefix list.
    PrefixListEntries Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries
}

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "prefix-list-entries" { return "PrefixListEntries" }
    return ""
}

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetSegmentPath() string {
    return "prefix" + "[name='" + fmt.Sprintf("%v", prefix.Name) + "']"
}

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list-entries" {
        return &prefix.PrefixListEntries
    }
    return nil
}

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-list-entries"] = &prefix.PrefixListEntries
    return children
}

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = prefix.Name
    return leafs
}

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetYangName() string { return "prefix" }

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *Ipv6AclAndPrefixList_Prefixes_Prefix) GetParentYangName() string { return "prefixes" }

// Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries
// Sequence of entries forming a prefix list
// This type is a presence type.
type Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A prefix list entry; either a description (remark) or a prefix to match
    // against. The type is slice of
    // Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry.
    PrefixListEntry []Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry
}

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetFilter() yfilter.YFilter { return prefixListEntries.YFilter }

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) SetFilter(yf yfilter.YFilter) { prefixListEntries.YFilter = yf }

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetGoName(yname string) string {
    if yname == "prefix-list-entry" { return "PrefixListEntry" }
    return ""
}

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetSegmentPath() string {
    return "prefix-list-entries"
}

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list-entry" {
        for _, c := range prefixListEntries.PrefixListEntry {
            if prefixListEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry{}
        prefixListEntries.PrefixListEntry = append(prefixListEntries.PrefixListEntry, child)
        return &prefixListEntries.PrefixListEntry[len(prefixListEntries.PrefixListEntry)-1]
    }
    return nil
}

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixListEntries.PrefixListEntry {
        children[prefixListEntries.PrefixListEntry[i].GetSegmentPath()] = &prefixListEntries.PrefixListEntry[i]
    }
    return children
}

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetYangName() string { return "prefix-list-entries" }

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) SetParent(parent types.Entity) { prefixListEntries.parent = parent }

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetParent() types.Entity { return prefixListEntries.parent }

func (prefixListEntries *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries) GetParentYangName() string { return "prefix" }

// Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry
// A prefix list entry; either a description
// (remark) or a prefix to match against
type Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence number of prefix list. The type is
    // interface{} with range: 1..2147483646.
    SequenceNumber interface{}

    // Whether to forward or drop packets matching the prefix list. The type is
    // Ipv6AclGrantEnum.
    Grant interface{}

    // The IPv6 address if entered  with the ZoneMutually exclusive with Prefix
    // and PrefixMask. The type is string.
    Ipv6AddressAsString interface{}

    // IPv6 Zone if entered  with the IPV6AddressMutually exclusive with Prefix
    // and PrefixMask. The type is string.
    Zone interface{}

    // IPv6 address prefix to match. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // MaskLength of IPv6 address prefix. The type is interface{} with range:
    // 0..128.
    PrefixMask interface{}

    // Set to perform an exact prefix length match. Item is mutually exclusive
    // with minimum and maximum length match items. The type is
    // Ipv6PrefixMatchExactLength.
    MatchExactLength interface{}

    // If exact prefix length matching specified, set the length of prefix to be
    // matched. The type is interface{} with range: 0..128.
    ExactPrefixLength interface{}

    // Set to perform a maximum length prefix match .  Item is mutually exclusive
    // with exact length match item. The type is Ipv6PrefixMatchMaxLength.
    MatchMaxLength interface{}

    // If maximum length prefix matching specified, set the maximum length of
    // prefix to be matched. The type is interface{} with range: 0..128.
    MaxPrefixLength interface{}

    // Set to perform a minimum length prefix match .  Item is mutually exclusive
    // with exact length match item. The type is Ipv6PrefixMatchMinLength.
    MatchMinLength interface{}

    // If minimum length prefix matching specified, set the minimum length of
    // prefix to be matched. The type is interface{} with range: 0..128.
    MinPrefixLength interface{}

    // Comments or a description for the prefix list.  Item is mutually exclusive
    // with all others in the object. The type is string.
    Remark interface{}
}

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetFilter() yfilter.YFilter { return prefixListEntry.YFilter }

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) SetFilter(yf yfilter.YFilter) { prefixListEntry.YFilter = yf }

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "grant" { return "Grant" }
    if yname == "ipv6-address-as-string" { return "Ipv6AddressAsString" }
    if yname == "zone" { return "Zone" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-mask" { return "PrefixMask" }
    if yname == "match-exact-length" { return "MatchExactLength" }
    if yname == "exact-prefix-length" { return "ExactPrefixLength" }
    if yname == "match-max-length" { return "MatchMaxLength" }
    if yname == "max-prefix-length" { return "MaxPrefixLength" }
    if yname == "match-min-length" { return "MatchMinLength" }
    if yname == "min-prefix-length" { return "MinPrefixLength" }
    if yname == "remark" { return "Remark" }
    return ""
}

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetSegmentPath() string {
    return "prefix-list-entry" + "[sequence-number='" + fmt.Sprintf("%v", prefixListEntry.SequenceNumber) + "']"
}

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = prefixListEntry.SequenceNumber
    leafs["grant"] = prefixListEntry.Grant
    leafs["ipv6-address-as-string"] = prefixListEntry.Ipv6AddressAsString
    leafs["zone"] = prefixListEntry.Zone
    leafs["prefix"] = prefixListEntry.Prefix
    leafs["prefix-mask"] = prefixListEntry.PrefixMask
    leafs["match-exact-length"] = prefixListEntry.MatchExactLength
    leafs["exact-prefix-length"] = prefixListEntry.ExactPrefixLength
    leafs["match-max-length"] = prefixListEntry.MatchMaxLength
    leafs["max-prefix-length"] = prefixListEntry.MaxPrefixLength
    leafs["match-min-length"] = prefixListEntry.MatchMinLength
    leafs["min-prefix-length"] = prefixListEntry.MinPrefixLength
    leafs["remark"] = prefixListEntry.Remark
    return leafs
}

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetYangName() string { return "prefix-list-entry" }

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) SetParent(parent types.Entity) { prefixListEntry.parent = parent }

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetParent() types.Entity { return prefixListEntry.parent }

func (prefixListEntry *Ipv6AclAndPrefixList_Prefixes_Prefix_PrefixListEntries_PrefixListEntry) GetParentYangName() string { return "prefix-list-entries" }

// Ipv6AclAndPrefixList_LogUpdate
// Control access lists log updates
type Ipv6AclAndPrefixList_LogUpdate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log update threshold (number of hits). The type is interface{} with range:
    // 1..2147483647.
    Threshold interface{}

    // Log update rate (log messages per second). The type is interface{} with
    // range: 1..1000.
    Rate interface{}
}

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetFilter() yfilter.YFilter { return logUpdate.YFilter }

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) SetFilter(yf yfilter.YFilter) { logUpdate.YFilter = yf }

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    if yname == "rate" { return "Rate" }
    return ""
}

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetSegmentPath() string {
    return "log-update"
}

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold"] = logUpdate.Threshold
    leafs["rate"] = logUpdate.Rate
    return leafs
}

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetBundleName() string { return "cisco_ios_xr" }

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetYangName() string { return "log-update" }

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) SetParent(parent types.Entity) { logUpdate.parent = parent }

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetParent() types.Entity { return logUpdate.parent }

func (logUpdate *Ipv6AclAndPrefixList_LogUpdate) GetParentYangName() string { return "ipv6-acl-and-prefix-list" }

// Ipv6AclAndPrefixList_Accesses
// Table of access lists
type Ipv6AclAndPrefixList_Accesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An ACL. The type is slice of Ipv6AclAndPrefixList_Accesses_Access.
    Access []Ipv6AclAndPrefixList_Accesses_Access
}

func (accesses *Ipv6AclAndPrefixList_Accesses) GetFilter() yfilter.YFilter { return accesses.YFilter }

func (accesses *Ipv6AclAndPrefixList_Accesses) SetFilter(yf yfilter.YFilter) { accesses.YFilter = yf }

func (accesses *Ipv6AclAndPrefixList_Accesses) GetGoName(yname string) string {
    if yname == "access" { return "Access" }
    return ""
}

func (accesses *Ipv6AclAndPrefixList_Accesses) GetSegmentPath() string {
    return "accesses"
}

func (accesses *Ipv6AclAndPrefixList_Accesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access" {
        for _, c := range accesses.Access {
            if accesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_Accesses_Access{}
        accesses.Access = append(accesses.Access, child)
        return &accesses.Access[len(accesses.Access)-1]
    }
    return nil
}

func (accesses *Ipv6AclAndPrefixList_Accesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accesses.Access {
        children[accesses.Access[i].GetSegmentPath()] = &accesses.Access[i]
    }
    return children
}

func (accesses *Ipv6AclAndPrefixList_Accesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accesses *Ipv6AclAndPrefixList_Accesses) GetBundleName() string { return "cisco_ios_xr" }

func (accesses *Ipv6AclAndPrefixList_Accesses) GetYangName() string { return "accesses" }

func (accesses *Ipv6AclAndPrefixList_Accesses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accesses *Ipv6AclAndPrefixList_Accesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accesses *Ipv6AclAndPrefixList_Accesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accesses *Ipv6AclAndPrefixList_Accesses) SetParent(parent types.Entity) { accesses.parent = parent }

func (accesses *Ipv6AclAndPrefixList_Accesses) GetParent() types.Entity { return accesses.parent }

func (accesses *Ipv6AclAndPrefixList_Accesses) GetParentYangName() string { return "ipv6-acl-and-prefix-list" }

// Ipv6AclAndPrefixList_Accesses_Access
// An ACL
type Ipv6AclAndPrefixList_Accesses_Access struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the access list. The type is string with
    // length: 1..65.
    Name interface{}

    // ACL entry table; contains list of access list entries.
    AccessListEntries Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries
}

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetFilter() yfilter.YFilter { return access.YFilter }

func (access *Ipv6AclAndPrefixList_Accesses_Access) SetFilter(yf yfilter.YFilter) { access.YFilter = yf }

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "access-list-entries" { return "AccessListEntries" }
    return ""
}

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetSegmentPath() string {
    return "access" + "[name='" + fmt.Sprintf("%v", access.Name) + "']"
}

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-entries" {
        return &access.AccessListEntries
    }
    return nil
}

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list-entries"] = &access.AccessListEntries
    return children
}

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = access.Name
    return leafs
}

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetBundleName() string { return "cisco_ios_xr" }

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetYangName() string { return "access" }

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (access *Ipv6AclAndPrefixList_Accesses_Access) SetParent(parent types.Entity) { access.parent = parent }

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetParent() types.Entity { return access.parent }

func (access *Ipv6AclAndPrefixList_Accesses_Access) GetParentYangName() string { return "accesses" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries
// ACL entry table; contains list of access list
// entries
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An ACL entry; either a description (remark) or anAccess List Entry to match
    // against. The type is slice of
    // Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry.
    AccessListEntry []Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry
}

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetFilter() yfilter.YFilter { return accessListEntries.YFilter }

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) SetFilter(yf yfilter.YFilter) { accessListEntries.YFilter = yf }

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetGoName(yname string) string {
    if yname == "access-list-entry" { return "AccessListEntry" }
    return ""
}

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetSegmentPath() string {
    return "access-list-entries"
}

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-entry" {
        for _, c := range accessListEntries.AccessListEntry {
            if accessListEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry{}
        accessListEntries.AccessListEntry = append(accessListEntries.AccessListEntry, child)
        return &accessListEntries.AccessListEntry[len(accessListEntries.AccessListEntry)-1]
    }
    return nil
}

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessListEntries.AccessListEntry {
        children[accessListEntries.AccessListEntry[i].GetSegmentPath()] = &accessListEntries.AccessListEntry[i]
    }
    return children
}

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetBundleName() string { return "cisco_ios_xr" }

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetYangName() string { return "access-list-entries" }

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) SetParent(parent types.Entity) { accessListEntries.parent = parent }

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetParent() types.Entity { return accessListEntries.parent }

func (accessListEntries *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries) GetParentYangName() string { return "access" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry
// An ACL entry; either a description (remark)
// or anAccess List Entry to match against
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence number of access list entry. The type is
    // interface{} with range: 1..2147483643.
    SequenceNumber interface{}

    // Whether to forward or drop packets matching the  ACE. The type is
    // Ipv6AclGrantEnum.
    Grant interface{}

    // Protocol operator. Leave unspecified if no protocol comparison is to be
    // done. The type is Ipv6AclOperatorEnum.
    ProtocolOperator interface{}

    // Protocol to match. The type is one of the following types: enumeration
    // Ipv6AclProtocolNumber, or int with range: 0..255.
    Protocol interface{}

    // Protocol2 to match. The type is one of the following types: enumeration
    // Ipv6AclProtocolNumber, or int with range: 0..255.
    Protocol2 interface{}

    // DSCP value to match (if a protocol was specified), leave unspecified if
    // DSCP comparion is not to be performed. The type is one of the following
    // types: enumeration Ipv6AclDscpNumber, or int with range: 0..64.
    Dscp interface{}

    // Precedence value to match (if a protocol was  specified), leave unspecified
    // if precedence  comparion is not to be performed. The type is one of the
    // following types: enumeration Ipv6AclPrecedenceNumber, or int with range:
    // 0..7.
    Precedence interface{}

    // Counter name. The type is string.
    CounterName interface{}

    // Whether and how to log matches against this  entry. The type is
    // Ipv6AclLoggingEnum.
    LogOption interface{}

    // Enable capture. The type is bool.
    Capture interface{}

    // Enable undetermined-transport. The type is bool.
    UndeterminedTransport interface{}

    // To turn off ICMP generation for deny ACEs. The type is interface{}.
    IcmpOff interface{}

    // Set qos-group number. The type is interface{} with range: 0..512.
    QosGroup interface{}

    // Set TTL Value. Ranges from 0-255. The type is interface{} with range:
    // 0..255.
    SetTtl interface{}

    // Comments or a description for the access list. The type is string.
    Remark interface{}

    // IPv6 source network object group name. The type is string with length:
    // 1..64.
    SourcePrefixGroup interface{}

    // IPv6 destination network object group name. The type is string with length:
    // 1..64.
    DestinationPrefixGroup interface{}

    // Source port object group name. The type is string with length: 1..64.
    SourcePortGroup interface{}

    // Destination port object group name. The type is string with length: 1..64.
    DestinationPortGroup interface{}

    // Sequence String for the ace. The type is string with length: 1..64.
    SequenceStr interface{}

    // Source network settings.
    SourceNetwork Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork

    // Destination network settings.
    DestinationNetwork Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork

    // Source port settings.
    SourcePort Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort

    // Destination port settings.
    DestinationPort Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort

    // ICMP settings.
    Icmp Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp

    // TCP settings.
    Tcp Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp

    // Packet length settings.
    PacketLength Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength

    // TTL settings.
    TimeToLive Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive

    // Next-hop settings.
    NextHop Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop

    // Match if header-flag is present.
    HeaderFlags Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags
}

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetFilter() yfilter.YFilter { return accessListEntry.YFilter }

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) SetFilter(yf yfilter.YFilter) { accessListEntry.YFilter = yf }

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "grant" { return "Grant" }
    if yname == "protocol-operator" { return "ProtocolOperator" }
    if yname == "protocol" { return "Protocol" }
    if yname == "protocol2" { return "Protocol2" }
    if yname == "dscp" { return "Dscp" }
    if yname == "precedence" { return "Precedence" }
    if yname == "counter-name" { return "CounterName" }
    if yname == "log-option" { return "LogOption" }
    if yname == "capture" { return "Capture" }
    if yname == "undetermined-transport" { return "UndeterminedTransport" }
    if yname == "icmp-off" { return "IcmpOff" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "set-ttl" { return "SetTtl" }
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
    if yname == "next-hop" { return "NextHop" }
    if yname == "header-flags" { return "HeaderFlags" }
    return ""
}

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetSegmentPath() string {
    return "access-list-entry" + "[sequence-number='" + fmt.Sprintf("%v", accessListEntry.SequenceNumber) + "']"
}

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
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
    if childYangName == "next-hop" {
        return &accessListEntry.NextHop
    }
    if childYangName == "header-flags" {
        return &accessListEntry.HeaderFlags
    }
    return nil
}

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source-network"] = &accessListEntry.SourceNetwork
    children["destination-network"] = &accessListEntry.DestinationNetwork
    children["source-port"] = &accessListEntry.SourcePort
    children["destination-port"] = &accessListEntry.DestinationPort
    children["icmp"] = &accessListEntry.Icmp
    children["tcp"] = &accessListEntry.Tcp
    children["packet-length"] = &accessListEntry.PacketLength
    children["time-to-live"] = &accessListEntry.TimeToLive
    children["next-hop"] = &accessListEntry.NextHop
    children["header-flags"] = &accessListEntry.HeaderFlags
    return children
}

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = accessListEntry.SequenceNumber
    leafs["grant"] = accessListEntry.Grant
    leafs["protocol-operator"] = accessListEntry.ProtocolOperator
    leafs["protocol"] = accessListEntry.Protocol
    leafs["protocol2"] = accessListEntry.Protocol2
    leafs["dscp"] = accessListEntry.Dscp
    leafs["precedence"] = accessListEntry.Precedence
    leafs["counter-name"] = accessListEntry.CounterName
    leafs["log-option"] = accessListEntry.LogOption
    leafs["capture"] = accessListEntry.Capture
    leafs["undetermined-transport"] = accessListEntry.UndeterminedTransport
    leafs["icmp-off"] = accessListEntry.IcmpOff
    leafs["qos-group"] = accessListEntry.QosGroup
    leafs["set-ttl"] = accessListEntry.SetTtl
    leafs["remark"] = accessListEntry.Remark
    leafs["source-prefix-group"] = accessListEntry.SourcePrefixGroup
    leafs["destination-prefix-group"] = accessListEntry.DestinationPrefixGroup
    leafs["source-port-group"] = accessListEntry.SourcePortGroup
    leafs["destination-port-group"] = accessListEntry.DestinationPortGroup
    leafs["sequence-str"] = accessListEntry.SequenceStr
    return leafs
}

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetBundleName() string { return "cisco_ios_xr" }

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetYangName() string { return "access-list-entry" }

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) SetParent(parent types.Entity) { accessListEntry.parent = parent }

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetParent() types.Entity { return accessListEntry.parent }

func (accessListEntry *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry) GetParentYangName() string { return "access-list-entries" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork
// Source network settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source IPv6 address, leave unspecified if inputting as IPv6 address with
    // wildcarding. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Wildcard bits to apply to source-address (if specified), leave unspecified
    // for no wildcarding. The type is interface{} with range: 0..128.
    SourceWildCardBits interface{}

    // Source address mask. Either source-wild-card-bits or source-mask is.
    // supported, not both. Leave unspecified. for any. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceMask interface{}
}

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetFilter() yfilter.YFilter { return sourceNetwork.YFilter }

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) SetFilter(yf yfilter.YFilter) { sourceNetwork.YFilter = yf }

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-wild-card-bits" { return "SourceWildCardBits" }
    if yname == "source-mask" { return "SourceMask" }
    return ""
}

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetSegmentPath() string {
    return "source-network"
}

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = sourceNetwork.SourceAddress
    leafs["source-wild-card-bits"] = sourceNetwork.SourceWildCardBits
    leafs["source-mask"] = sourceNetwork.SourceMask
    return leafs
}

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetBundleName() string { return "cisco_ios_xr" }

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetYangName() string { return "source-network" }

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) SetParent(parent types.Entity) { sourceNetwork.parent = parent }

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetParent() types.Entity { return sourceNetwork.parent }

func (sourceNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourceNetwork) GetParentYangName() string { return "access-list-entry" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork
// Destination network settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination IPv6 address, leave unspecified if inputting as IPv6 address
    // with  wildcarding. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Wildcard bits to apply to destination  destination-address (if specified), 
    // leave unspecified for no wildcarding. The type is interface{} with range:
    // 0..128.
    DestinationWildCardBits interface{}

    // Destination address mask. Either destination-wild-card-bits or
    // destination-mask. is supported, not both. Leave unspecified for any. The
    // type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationMask interface{}
}

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetFilter() yfilter.YFilter { return destinationNetwork.YFilter }

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) SetFilter(yf yfilter.YFilter) { destinationNetwork.YFilter = yf }

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetGoName(yname string) string {
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-wild-card-bits" { return "DestinationWildCardBits" }
    if yname == "destination-mask" { return "DestinationMask" }
    return ""
}

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetSegmentPath() string {
    return "destination-network"
}

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-address"] = destinationNetwork.DestinationAddress
    leafs["destination-wild-card-bits"] = destinationNetwork.DestinationWildCardBits
    leafs["destination-mask"] = destinationNetwork.DestinationMask
    return leafs
}

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetBundleName() string { return "cisco_ios_xr" }

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetYangName() string { return "destination-network" }

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) SetParent(parent types.Entity) { destinationNetwork.parent = parent }

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetParent() types.Entity { return destinationNetwork.parent }

func (destinationNetwork *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationNetwork) GetParentYangName() string { return "access-list-entry" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort
// Source port settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source comparison operator. Leave unspecified if no source port comparison
    // is to be done. The type is Ipv6AclOperatorEnum.
    SourceOperator interface{}

    // First source port for comparison,  leave unspecified if source port
    // comparison is not to be performed. The type is one of the following types:
    // enumeration Ipv6AclPortNumber, or int with range: 0..65535.
    FirstSourcePort interface{}

    // Second source port for comparion,  leave unspecified if source port
    // comparison is not to be performed. The type is one of the following types:
    // enumeration Ipv6AclPortNumber, or int with range: 0..65535.
    SecondSourcePort interface{}
}

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetFilter() yfilter.YFilter { return sourcePort.YFilter }

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) SetFilter(yf yfilter.YFilter) { sourcePort.YFilter = yf }

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetGoName(yname string) string {
    if yname == "source-operator" { return "SourceOperator" }
    if yname == "first-source-port" { return "FirstSourcePort" }
    if yname == "second-source-port" { return "SecondSourcePort" }
    return ""
}

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetSegmentPath() string {
    return "source-port"
}

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-operator"] = sourcePort.SourceOperator
    leafs["first-source-port"] = sourcePort.FirstSourcePort
    leafs["second-source-port"] = sourcePort.SecondSourcePort
    return leafs
}

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetBundleName() string { return "cisco_ios_xr" }

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetYangName() string { return "source-port" }

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) SetParent(parent types.Entity) { sourcePort.parent = parent }

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetParent() types.Entity { return sourcePort.parent }

func (sourcePort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_SourcePort) GetParentYangName() string { return "access-list-entry" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort
// Destination port settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Destination comparison operator. Leave  unspecified if no destination port
    // comparison  is to be done. The type is Ipv6AclOperatorEnum.
    DestinationOperator interface{}

    // First destination port for comparison, leave  unspecified if destination
    // port comparison is not to be performed. The type is one of the following
    // types: enumeration Ipv6AclPortNumber, or int with range: 0..65535.
    FirstDestinationPort interface{}

    // Second destination port for comparion, leave  unspecified if destination
    // port comparison is not to be performed. The type is one of the following
    // types: enumeration Ipv6AclPortNumber, or int with range: 0..65535.
    SecondDestinationPort interface{}
}

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetFilter() yfilter.YFilter { return destinationPort.YFilter }

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) SetFilter(yf yfilter.YFilter) { destinationPort.YFilter = yf }

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetGoName(yname string) string {
    if yname == "destination-operator" { return "DestinationOperator" }
    if yname == "first-destination-port" { return "FirstDestinationPort" }
    if yname == "second-destination-port" { return "SecondDestinationPort" }
    return ""
}

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetSegmentPath() string {
    return "destination-port"
}

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["destination-operator"] = destinationPort.DestinationOperator
    leafs["first-destination-port"] = destinationPort.FirstDestinationPort
    leafs["second-destination-port"] = destinationPort.SecondDestinationPort
    return leafs
}

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetBundleName() string { return "cisco_ios_xr" }

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetYangName() string { return "destination-port" }

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) SetParent(parent types.Entity) { destinationPort.parent = parent }

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetParent() types.Entity { return destinationPort.parent }

func (destinationPort *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_DestinationPort) GetParentYangName() string { return "access-list-entry" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp
// ICMP settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Well known ICMP message code types to match,  leave unspecified if ICMP
    // message code type  comparion is not to be performed. The type is
    // Ipv6AclIcmpTypeCodeEnum.
    IcmpTypeCode interface{}
}

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetFilter() yfilter.YFilter { return icmp.YFilter }

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) SetFilter(yf yfilter.YFilter) { icmp.YFilter = yf }

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetGoName(yname string) string {
    if yname == "icmp-type-code" { return "IcmpTypeCode" }
    return ""
}

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetSegmentPath() string {
    return "icmp"
}

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["icmp-type-code"] = icmp.IcmpTypeCode
    return leafs
}

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetBundleName() string { return "cisco_ios_xr" }

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetYangName() string { return "icmp" }

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) SetParent(parent types.Entity) { icmp.parent = parent }

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetParent() types.Entity { return icmp.parent }

func (icmp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Icmp) GetParentYangName() string { return "access-list-entry" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp
// TCP settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP Bits match operator. Leave unspecified if  flexible comparison of TCP
    // bits is not  required. The type is Ipv6AclTcpMatchOperatorEnum.
    TcpBitsMatchOperator interface{}

    // TCP bits to match. Leave unspecified if  comparison of TCP bits is not
    // required. The type is one of the following types: enumeration
    // Ipv6AclTcpBitsNumber, or int with range: 0..63.
    TcpBits interface{}

    // TCP bits mask to use for flexible TCP matching. Leave unspecified if it is
    // not required. The type is one of the following types: enumeration
    // Ipv6AclTcpBitsNumber, or int with range: 0..63.
    TcpBitsMask interface{}
}

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetGoName(yname string) string {
    if yname == "tcp-bits-match-operator" { return "TcpBitsMatchOperator" }
    if yname == "tcp-bits" { return "TcpBits" }
    if yname == "tcp-bits-mask" { return "TcpBitsMask" }
    return ""
}

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetSegmentPath() string {
    return "tcp"
}

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-bits-match-operator"] = tcp.TcpBitsMatchOperator
    leafs["tcp-bits"] = tcp.TcpBits
    leafs["tcp-bits-mask"] = tcp.TcpBitsMask
    return leafs
}

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetBundleName() string { return "cisco_ios_xr" }

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetYangName() string { return "tcp" }

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_Tcp) GetParentYangName() string { return "access-list-entry" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength
// Packet length settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packet length operator applicable if packet  length is to be compared.
    // Leave unspecified if no Packet length comparison is to be done. The type is
    // Ipv6AclOperatorEnum.
    PacketLengthOperator interface{}

    // Minimum packet length for comparison, leave  unspecified if packet length
    // comparison is not to be performed or if only the maximum packet length
    // should be considered. The type is interface{} with range: 0..65535.
    PacketLengthMin interface{}

    // Maximum packet length for comparion, leave  unspecified if packet length
    // comparison is not to be performed or if only the minimum packet  length
    // should be considered. The type is interface{} with range: 0..65535.
    PacketLengthMax interface{}
}

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetFilter() yfilter.YFilter { return packetLength.YFilter }

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) SetFilter(yf yfilter.YFilter) { packetLength.YFilter = yf }

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetGoName(yname string) string {
    if yname == "packet-length-operator" { return "PacketLengthOperator" }
    if yname == "packet-length-min" { return "PacketLengthMin" }
    if yname == "packet-length-max" { return "PacketLengthMax" }
    return ""
}

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetSegmentPath() string {
    return "packet-length"
}

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packet-length-operator"] = packetLength.PacketLengthOperator
    leafs["packet-length-min"] = packetLength.PacketLengthMin
    leafs["packet-length-max"] = packetLength.PacketLengthMax
    return leafs
}

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetBundleName() string { return "cisco_ios_xr" }

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetYangName() string { return "packet-length" }

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) SetParent(parent types.Entity) { packetLength.parent = parent }

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetParent() types.Entity { return packetLength.parent }

func (packetLength *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_PacketLength) GetParentYangName() string { return "access-list-entry" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive
// TTL settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TTL operator is applicable if TTL is to be  compared. Leave unspecified if
    // TTL  classification is not required. The type is Ipv6AclOperatorEnum.
    TimeToLiveOperator interface{}

    // TTL value for comparison OR Minimum TTL value  for TTL range comparision,
    // leave unspecified if TTL classification is not required. The type is
    // interface{} with range: 0..255.
    TimeToLiveMin interface{}

    // Maximum TTL for comparion, leave unspecified if TTL comparison is not to be
    // performed or if only the minimum TTL should be considered. The type is
    // interface{} with range: 0..255.
    TimeToLiveMax interface{}
}

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetFilter() yfilter.YFilter { return timeToLive.YFilter }

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) SetFilter(yf yfilter.YFilter) { timeToLive.YFilter = yf }

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetGoName(yname string) string {
    if yname == "time-to-live-operator" { return "TimeToLiveOperator" }
    if yname == "time-to-live-min" { return "TimeToLiveMin" }
    if yname == "time-to-live-max" { return "TimeToLiveMax" }
    return ""
}

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetSegmentPath() string {
    return "time-to-live"
}

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-to-live-operator"] = timeToLive.TimeToLiveOperator
    leafs["time-to-live-min"] = timeToLive.TimeToLiveMin
    leafs["time-to-live-max"] = timeToLive.TimeToLiveMax
    return leafs
}

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetBundleName() string { return "cisco_ios_xr" }

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetYangName() string { return "time-to-live" }

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) SetParent(parent types.Entity) { timeToLive.parent = parent }

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetParent() types.Entity { return timeToLive.parent }

func (timeToLive *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_TimeToLive) GetParentYangName() string { return "access-list-entry" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop
// Next-hop settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The nexthop type. The type is NextHopType.
    NextHopType interface{}

    // The first next-hop settings.
    NextHop1 Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1

    // The second next-hop settings.
    NextHop2 Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2

    // The third next-hop settings.
    NextHop3 Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3
}

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetFilter() yfilter.YFilter { return nextHop.YFilter }

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) SetFilter(yf yfilter.YFilter) { nextHop.YFilter = yf }

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetGoName(yname string) string {
    if yname == "next-hop-type" { return "NextHopType" }
    if yname == "next-hop-1" { return "NextHop1" }
    if yname == "next-hop-2" { return "NextHop2" }
    if yname == "next-hop-3" { return "NextHop3" }
    return ""
}

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetSegmentPath() string {
    return "next-hop"
}

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetChildByName(childYangName string, segmentPath string) types.Entity {
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

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["next-hop-1"] = &nextHop.NextHop1
    children["next-hop-2"] = &nextHop.NextHop2
    children["next-hop-3"] = &nextHop.NextHop3
    return children
}

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop-type"] = nextHop.NextHopType
    return leafs
}

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetYangName() string { return "next-hop" }

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) SetParent(parent types.Entity) { nextHop.parent = parent }

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetParent() types.Entity { return nextHop.parent }

func (nextHop *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop) GetParentYangName() string { return "access-list-entry" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1
// The first next-hop settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IPv6 address of the next-hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // The VRF name of the next-hop. The type is string.
    VrfName interface{}

    // The object tracking name for the next-hop. The type is string.
    TrackName interface{}
}

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetFilter() yfilter.YFilter { return nextHop1.YFilter }

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) SetFilter(yf yfilter.YFilter) { nextHop1.YFilter = yf }

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "track-name" { return "TrackName" }
    return ""
}

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetSegmentPath() string {
    return "next-hop-1"
}

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHop1.NextHop
    leafs["vrf-name"] = nextHop1.VrfName
    leafs["track-name"] = nextHop1.TrackName
    return leafs
}

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetYangName() string { return "next-hop-1" }

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) SetParent(parent types.Entity) { nextHop1.parent = parent }

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetParent() types.Entity { return nextHop1.parent }

func (nextHop1 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop1) GetParentYangName() string { return "next-hop" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2
// The second next-hop settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IPv6 address of the next-hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // The VRF name of the next-hop. The type is string.
    VrfName interface{}

    // The object tracking name for the next-hop. The type is string.
    TrackName interface{}
}

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetFilter() yfilter.YFilter { return nextHop2.YFilter }

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) SetFilter(yf yfilter.YFilter) { nextHop2.YFilter = yf }

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "track-name" { return "TrackName" }
    return ""
}

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetSegmentPath() string {
    return "next-hop-2"
}

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHop2.NextHop
    leafs["vrf-name"] = nextHop2.VrfName
    leafs["track-name"] = nextHop2.TrackName
    return leafs
}

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetYangName() string { return "next-hop-2" }

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) SetParent(parent types.Entity) { nextHop2.parent = parent }

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetParent() types.Entity { return nextHop2.parent }

func (nextHop2 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop2) GetParentYangName() string { return "next-hop" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3
// The third next-hop settings.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The IPv6 address of the next-hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // The VRF name of the next-hop. The type is string.
    VrfName interface{}

    // The object tracking name for the next-hop. The type is string.
    TrackName interface{}
}

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetFilter() yfilter.YFilter { return nextHop3.YFilter }

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) SetFilter(yf yfilter.YFilter) { nextHop3.YFilter = yf }

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "track-name" { return "TrackName" }
    return ""
}

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetSegmentPath() string {
    return "next-hop-3"
}

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHop3.NextHop
    leafs["vrf-name"] = nextHop3.VrfName
    leafs["track-name"] = nextHop3.TrackName
    return leafs
}

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetBundleName() string { return "cisco_ios_xr" }

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetYangName() string { return "next-hop-3" }

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) SetParent(parent types.Entity) { nextHop3.parent = parent }

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetParent() types.Entity { return nextHop3.parent }

func (nextHop3 *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_NextHop_NextHop3) GetParentYangName() string { return "next-hop" }

// Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags
// Match if header-flag is present.
type Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Match if routing header is present. The type is interface{}.
    Routing interface{}

    // Match if destops header is present. The type is interface{}.
    Destopts interface{}

    // Match if hop-by-hop header is present. The type is interface{}.
    HopByHop interface{}

    // Match if fragments header is present. The type is interface{}.
    Fragments interface{}

    // Match if authen header is present. The type is interface{}.
    Authen interface{}
}

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetFilter() yfilter.YFilter { return headerFlags.YFilter }

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) SetFilter(yf yfilter.YFilter) { headerFlags.YFilter = yf }

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetGoName(yname string) string {
    if yname == "routing" { return "Routing" }
    if yname == "destopts" { return "Destopts" }
    if yname == "hop-by-hop" { return "HopByHop" }
    if yname == "fragments" { return "Fragments" }
    if yname == "authen" { return "Authen" }
    return ""
}

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetSegmentPath() string {
    return "header-flags"
}

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["routing"] = headerFlags.Routing
    leafs["destopts"] = headerFlags.Destopts
    leafs["hop-by-hop"] = headerFlags.HopByHop
    leafs["fragments"] = headerFlags.Fragments
    leafs["authen"] = headerFlags.Authen
    return leafs
}

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetBundleName() string { return "cisco_ios_xr" }

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetYangName() string { return "header-flags" }

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) SetParent(parent types.Entity) { headerFlags.parent = parent }

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetParent() types.Entity { return headerFlags.parent }

func (headerFlags *Ipv6AclAndPrefixList_Accesses_Access_AccessListEntries_AccessListEntry_HeaderFlags) GetParentYangName() string { return "access-list-entry" }

