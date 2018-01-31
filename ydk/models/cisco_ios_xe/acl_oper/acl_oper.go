// This module contains a collection of YANG definitions
// for ACL statistical data.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package acl_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package acl_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-acl-oper access-lists}", reflect.TypeOf(AccessLists{}))
    ydk.RegisterEntity("Cisco-IOS-XE-acl-oper:access-lists", reflect.TypeOf(AccessLists{}))
}

// AccessLists
// This is top level container for Access Control Lists. It can have one
// or more Access Control List.
type AccessLists struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An access list (acl) is an ordered list of access list entries (ACE). Each
    // access control entries has a list of match criteria, and a list of actions.
    // Since there are several kinds of access control lists implemented with
    // different attributes for each and different for each vendor, this model
    // accommodates customizing access control lists for each kind and for each
    // vendor. The type is slice of AccessLists_AccessList.
    AccessList []AccessLists_AccessList
}

func (accessLists *AccessLists) GetFilter() yfilter.YFilter { return accessLists.YFilter }

func (accessLists *AccessLists) SetFilter(yf yfilter.YFilter) { accessLists.YFilter = yf }

func (accessLists *AccessLists) GetGoName(yname string) string {
    if yname == "access-list" { return "AccessList" }
    return ""
}

func (accessLists *AccessLists) GetSegmentPath() string {
    return "Cisco-IOS-XE-acl-oper:access-lists"
}

func (accessLists *AccessLists) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list" {
        for _, c := range accessLists.AccessList {
            if accessLists.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AccessLists_AccessList{}
        accessLists.AccessList = append(accessLists.AccessList, child)
        return &accessLists.AccessList[len(accessLists.AccessList)-1]
    }
    return nil
}

func (accessLists *AccessLists) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessLists.AccessList {
        children[accessLists.AccessList[i].GetSegmentPath()] = &accessLists.AccessList[i]
    }
    return children
}

func (accessLists *AccessLists) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessLists *AccessLists) GetBundleName() string { return "cisco_ios_xe" }

func (accessLists *AccessLists) GetYangName() string { return "access-lists" }

func (accessLists *AccessLists) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (accessLists *AccessLists) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (accessLists *AccessLists) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (accessLists *AccessLists) SetParent(parent types.Entity) { accessLists.parent = parent }

func (accessLists *AccessLists) GetParent() types.Entity { return accessLists.parent }

func (accessLists *AccessLists) GetParentYangName() string { return "Cisco-IOS-XE-acl-oper" }

// AccessLists_AccessList
// An access list (acl) is an ordered list of
// access list entries (ACE). Each access control entries has a
// list of match criteria, and a list of actions.
// Since there are several kinds of access control lists
// implemented with different attributes for
// each and different for each vendor, this
// model accommodates customizing access control lists for
// each kind and for each vendor.
type AccessLists_AccessList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of access-list. A device MAY restrict the
    // length and value of this name, possibly space and special characters are
    // not allowed. The type is string.
    AccessControlListName interface{}

    // access-list-entry(ACE) information.
    AccessListEntries AccessLists_AccessList_AccessListEntries
}

func (accessList *AccessLists_AccessList) GetFilter() yfilter.YFilter { return accessList.YFilter }

func (accessList *AccessLists_AccessList) SetFilter(yf yfilter.YFilter) { accessList.YFilter = yf }

func (accessList *AccessLists_AccessList) GetGoName(yname string) string {
    if yname == "access-control-list-name" { return "AccessControlListName" }
    if yname == "access-list-entries" { return "AccessListEntries" }
    return ""
}

func (accessList *AccessLists_AccessList) GetSegmentPath() string {
    return "access-list" + "[access-control-list-name='" + fmt.Sprintf("%v", accessList.AccessControlListName) + "']"
}

func (accessList *AccessLists_AccessList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-entries" {
        return &accessList.AccessListEntries
    }
    return nil
}

func (accessList *AccessLists_AccessList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list-entries"] = &accessList.AccessListEntries
    return children
}

func (accessList *AccessLists_AccessList) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-control-list-name"] = accessList.AccessControlListName
    return leafs
}

func (accessList *AccessLists_AccessList) GetBundleName() string { return "cisco_ios_xe" }

func (accessList *AccessLists_AccessList) GetYangName() string { return "access-list" }

func (accessList *AccessLists_AccessList) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (accessList *AccessLists_AccessList) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (accessList *AccessLists_AccessList) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (accessList *AccessLists_AccessList) SetParent(parent types.Entity) { accessList.parent = parent }

func (accessList *AccessLists_AccessList) GetParent() types.Entity { return accessList.parent }

func (accessList *AccessLists_AccessList) GetParentYangName() string { return "access-lists" }

// AccessLists_AccessList_AccessListEntries
// access-list-entry(ACE) information
type AccessLists_AccessList_AccessListEntries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A list of ACEs. The type is slice of
    // AccessLists_AccessList_AccessListEntries_AccessListEntry.
    AccessListEntry []AccessLists_AccessList_AccessListEntries_AccessListEntry
}

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetFilter() yfilter.YFilter { return accessListEntries.YFilter }

func (accessListEntries *AccessLists_AccessList_AccessListEntries) SetFilter(yf yfilter.YFilter) { accessListEntries.YFilter = yf }

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetGoName(yname string) string {
    if yname == "access-list-entry" { return "AccessListEntry" }
    return ""
}

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetSegmentPath() string {
    return "access-list-entries"
}

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-entry" {
        for _, c := range accessListEntries.AccessListEntry {
            if accessListEntries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AccessLists_AccessList_AccessListEntries_AccessListEntry{}
        accessListEntries.AccessListEntry = append(accessListEntries.AccessListEntry, child)
        return &accessListEntries.AccessListEntry[len(accessListEntries.AccessListEntry)-1]
    }
    return nil
}

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessListEntries.AccessListEntry {
        children[accessListEntries.AccessListEntry[i].GetSegmentPath()] = &accessListEntries.AccessListEntry[i]
    }
    return children
}

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetBundleName() string { return "cisco_ios_xe" }

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetYangName() string { return "access-list-entries" }

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (accessListEntries *AccessLists_AccessList_AccessListEntries) SetParent(parent types.Entity) { accessListEntries.parent = parent }

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetParent() types.Entity { return accessListEntries.parent }

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetParentYangName() string { return "access-list" }

// AccessLists_AccessList_AccessListEntries_AccessListEntry
// A list of ACEs
type AccessLists_AccessList_AccessListEntries_AccessListEntry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Entry number. The type is interface{} with range:
    // 0..4294967295.
    RuleName interface{}

    // Per access list entries operational data.
    AccessListEntriesOperData AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData
}

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetFilter() yfilter.YFilter { return accessListEntry.YFilter }

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) SetFilter(yf yfilter.YFilter) { accessListEntry.YFilter = yf }

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetGoName(yname string) string {
    if yname == "rule-name" { return "RuleName" }
    if yname == "access-list-entries-oper-data" { return "AccessListEntriesOperData" }
    return ""
}

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetSegmentPath() string {
    return "access-list-entry" + "[rule-name='" + fmt.Sprintf("%v", accessListEntry.RuleName) + "']"
}

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-entries-oper-data" {
        return &accessListEntry.AccessListEntriesOperData
    }
    return nil
}

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list-entries-oper-data"] = &accessListEntry.AccessListEntriesOperData
    return children
}

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name"] = accessListEntry.RuleName
    return leafs
}

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetBundleName() string { return "cisco_ios_xe" }

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetYangName() string { return "access-list-entry" }

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) SetParent(parent types.Entity) { accessListEntry.parent = parent }

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetParent() types.Entity { return accessListEntry.parent }

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetParentYangName() string { return "access-list-entries" }

// AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData
// Per access list entries operational data
type AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of matches for an access list entry. The type is interface{} with
    // range: 0..18446744073709551615.
    MatchCounter interface{}
}

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetFilter() yfilter.YFilter { return accessListEntriesOperData.YFilter }

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) SetFilter(yf yfilter.YFilter) { accessListEntriesOperData.YFilter = yf }

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetGoName(yname string) string {
    if yname == "match-counter" { return "MatchCounter" }
    return ""
}

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetSegmentPath() string {
    return "access-list-entries-oper-data"
}

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["match-counter"] = accessListEntriesOperData.MatchCounter
    return leafs
}

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetBundleName() string { return "cisco_ios_xe" }

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetYangName() string { return "access-list-entries-oper-data" }

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) SetParent(parent types.Entity) { accessListEntriesOperData.parent = parent }

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetParent() types.Entity { return accessListEntriesOperData.parent }

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetParentYangName() string { return "access-list-entry" }

