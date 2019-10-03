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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An access list (acl) is an ordered list of access list entries (ACE). Each
    // access control entries has a list of match criteria, and a list of actions.
    // Since there are several kinds of access control lists implemented with
    // different attributes for each and different for each vendor, this model
    // accommodates customizing access control lists for each kind and for each
    // vendor. The type is slice of AccessLists_AccessList.
    AccessList []*AccessLists_AccessList
}

func (accessLists *AccessLists) GetEntityData() *types.CommonEntityData {
    accessLists.EntityData.YFilter = accessLists.YFilter
    accessLists.EntityData.YangName = "access-lists"
    accessLists.EntityData.BundleName = "cisco_ios_xe"
    accessLists.EntityData.ParentYangName = "Cisco-IOS-XE-acl-oper"
    accessLists.EntityData.SegmentPath = "Cisco-IOS-XE-acl-oper:access-lists"
    accessLists.EntityData.AbsolutePath = accessLists.EntityData.SegmentPath
    accessLists.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessLists.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessLists.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessLists.EntityData.Children = types.NewOrderedMap()
    accessLists.EntityData.Children.Append("access-list", types.YChild{"AccessList", nil})
    for i := range accessLists.AccessList {
        accessLists.EntityData.Children.Append(types.GetSegmentPath(accessLists.AccessList[i]), types.YChild{"AccessList", accessLists.AccessList[i]})
    }
    accessLists.EntityData.Leafs = types.NewOrderedMap()

    accessLists.EntityData.YListKeys = []string {}

    return &(accessLists.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The name of access-list. A device MAY restrict the
    // length and value of this name, possibly space and special characters are
    // not allowed. The type is string.
    AccessControlListName interface{}

    // access-list-entry(ACE) information.
    AccessListEntries AccessLists_AccessList_AccessListEntries
}

func (accessList *AccessLists_AccessList) GetEntityData() *types.CommonEntityData {
    accessList.EntityData.YFilter = accessList.YFilter
    accessList.EntityData.YangName = "access-list"
    accessList.EntityData.BundleName = "cisco_ios_xe"
    accessList.EntityData.ParentYangName = "access-lists"
    accessList.EntityData.SegmentPath = "access-list" + types.AddKeyToken(accessList.AccessControlListName, "access-control-list-name")
    accessList.EntityData.AbsolutePath = "Cisco-IOS-XE-acl-oper:access-lists/" + accessList.EntityData.SegmentPath
    accessList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessList.EntityData.Children = types.NewOrderedMap()
    accessList.EntityData.Children.Append("access-list-entries", types.YChild{"AccessListEntries", &accessList.AccessListEntries})
    accessList.EntityData.Leafs = types.NewOrderedMap()
    accessList.EntityData.Leafs.Append("access-control-list-name", types.YLeaf{"AccessControlListName", accessList.AccessControlListName})

    accessList.EntityData.YListKeys = []string {"AccessControlListName"}

    return &(accessList.EntityData)
}

// AccessLists_AccessList_AccessListEntries
// access-list-entry(ACE) information
type AccessLists_AccessList_AccessListEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of ACEs. The type is slice of
    // AccessLists_AccessList_AccessListEntries_AccessListEntry.
    AccessListEntry []*AccessLists_AccessList_AccessListEntries_AccessListEntry
}

func (accessListEntries *AccessLists_AccessList_AccessListEntries) GetEntityData() *types.CommonEntityData {
    accessListEntries.EntityData.YFilter = accessListEntries.YFilter
    accessListEntries.EntityData.YangName = "access-list-entries"
    accessListEntries.EntityData.BundleName = "cisco_ios_xe"
    accessListEntries.EntityData.ParentYangName = "access-list"
    accessListEntries.EntityData.SegmentPath = "access-list-entries"
    accessListEntries.EntityData.AbsolutePath = "Cisco-IOS-XE-acl-oper:access-lists/access-list/" + accessListEntries.EntityData.SegmentPath
    accessListEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessListEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessListEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessListEntries.EntityData.Children = types.NewOrderedMap()
    accessListEntries.EntityData.Children.Append("access-list-entry", types.YChild{"AccessListEntry", nil})
    for i := range accessListEntries.AccessListEntry {
        accessListEntries.EntityData.Children.Append(types.GetSegmentPath(accessListEntries.AccessListEntry[i]), types.YChild{"AccessListEntry", accessListEntries.AccessListEntry[i]})
    }
    accessListEntries.EntityData.Leafs = types.NewOrderedMap()

    accessListEntries.EntityData.YListKeys = []string {}

    return &(accessListEntries.EntityData)
}

// AccessLists_AccessList_AccessListEntries_AccessListEntry
// A list of ACEs
type AccessLists_AccessList_AccessListEntries_AccessListEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Entry number. The type is interface{} with range:
    // 0..4294967295.
    RuleName interface{}

    // Per access list entries operational data.
    AccessListEntriesOperData AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData
}

func (accessListEntry *AccessLists_AccessList_AccessListEntries_AccessListEntry) GetEntityData() *types.CommonEntityData {
    accessListEntry.EntityData.YFilter = accessListEntry.YFilter
    accessListEntry.EntityData.YangName = "access-list-entry"
    accessListEntry.EntityData.BundleName = "cisco_ios_xe"
    accessListEntry.EntityData.ParentYangName = "access-list-entries"
    accessListEntry.EntityData.SegmentPath = "access-list-entry" + types.AddKeyToken(accessListEntry.RuleName, "rule-name")
    accessListEntry.EntityData.AbsolutePath = "Cisco-IOS-XE-acl-oper:access-lists/access-list/access-list-entries/" + accessListEntry.EntityData.SegmentPath
    accessListEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessListEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessListEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessListEntry.EntityData.Children = types.NewOrderedMap()
    accessListEntry.EntityData.Children.Append("access-list-entries-oper-data", types.YChild{"AccessListEntriesOperData", &accessListEntry.AccessListEntriesOperData})
    accessListEntry.EntityData.Leafs = types.NewOrderedMap()
    accessListEntry.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", accessListEntry.RuleName})

    accessListEntry.EntityData.YListKeys = []string {"RuleName"}

    return &(accessListEntry.EntityData)
}

// AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData
// Per access list entries operational data
type AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of matches for an access list entry. The type is interface{} with
    // range: 0..18446744073709551615.
    MatchCounter interface{}
}

func (accessListEntriesOperData *AccessLists_AccessList_AccessListEntries_AccessListEntry_AccessListEntriesOperData) GetEntityData() *types.CommonEntityData {
    accessListEntriesOperData.EntityData.YFilter = accessListEntriesOperData.YFilter
    accessListEntriesOperData.EntityData.YangName = "access-list-entries-oper-data"
    accessListEntriesOperData.EntityData.BundleName = "cisco_ios_xe"
    accessListEntriesOperData.EntityData.ParentYangName = "access-list-entry"
    accessListEntriesOperData.EntityData.SegmentPath = "access-list-entries-oper-data"
    accessListEntriesOperData.EntityData.AbsolutePath = "Cisco-IOS-XE-acl-oper:access-lists/access-list/access-list-entries/access-list-entry/" + accessListEntriesOperData.EntityData.SegmentPath
    accessListEntriesOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessListEntriesOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessListEntriesOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessListEntriesOperData.EntityData.Children = types.NewOrderedMap()
    accessListEntriesOperData.EntityData.Leafs = types.NewOrderedMap()
    accessListEntriesOperData.EntityData.Leafs.Append("match-counter", types.YLeaf{"MatchCounter", accessListEntriesOperData.MatchCounter})

    accessListEntriesOperData.EntityData.YListKeys = []string {}

    return &(accessListEntriesOperData.EntityData)
}

