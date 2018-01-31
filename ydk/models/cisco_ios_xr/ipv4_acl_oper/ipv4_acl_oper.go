// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-acl package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-acl-and-prefix-list: Root class of IPv4 Oper schema tree
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_acl_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_acl_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-acl-oper ipv4-acl-and-prefix-list}", reflect.TypeOf(Ipv4AclAndPrefixList{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list", reflect.TypeOf(Ipv4AclAndPrefixList{}))
}

// AclTcpflagsOperator represents Acl tcpflags operator
type AclTcpflagsOperator string

const (
    // Match None
    AclTcpflagsOperator_match_none AclTcpflagsOperator = "match-none"

    // Match All
    AclTcpflagsOperator_match_all AclTcpflagsOperator = "match-all"

    // Match any old
    AclTcpflagsOperator_match_any_old AclTcpflagsOperator = "match-any-old"

    // Match any
    AclTcpflagsOperator_match_any AclTcpflagsOperator = "match-any"
)

// AclPortOperator represents Acl port operator
type AclPortOperator string

const (
    // None
    AclPortOperator_none AclPortOperator = "none"

    // Equal
    AclPortOperator_eq AclPortOperator = "eq"

    // Greater than
    AclPortOperator_gt AclPortOperator = "gt"

    // Less than
    AclPortOperator_lt AclPortOperator = "lt"

    // Not Equal
    AclPortOperator_neq AclPortOperator = "neq"

    // Range
    AclPortOperator_range_ AclPortOperator = "range"

    // One Byte
    AclPortOperator_onebyte AclPortOperator = "onebyte"

    // Two Bytes
    AclPortOperator_twobytes AclPortOperator = "twobytes"
)

// AclAce1 represents ACE Types
type AclAce1 string

const (
    // This is Normal ACE
    AclAce1_normal AclAce1 = "normal"

    // This is Remark ACE
    AclAce1_remark AclAce1 = "remark"

    // This is ABF ACE
    AclAce1_abf AclAce1 = "abf"
)

// AclPortOperator represents Acl port operator
type AclPortOperator string

const (
    // None
    AclPortOperator_none AclPortOperator = "none"

    // Equal
    AclPortOperator_eq AclPortOperator = "eq"

    // Greater than
    AclPortOperator_gt AclPortOperator = "gt"

    // Less than
    AclPortOperator_lt AclPortOperator = "lt"

    // Not Equal
    AclPortOperator_neq AclPortOperator = "neq"

    // Range
    AclPortOperator_range_ AclPortOperator = "range"

    // One Byte
    AclPortOperator_onebyte AclPortOperator = "onebyte"

    // Two Bytes
    AclPortOperator_twobytes AclPortOperator = "twobytes"
)

// BagAclNhAtStatus represents Bag acl nh at status
type BagAclNhAtStatus string

const (
    // AT State Unknown
    BagAclNhAtStatus_unknown BagAclNhAtStatus = "unknown"

    // AT State UP
    BagAclNhAtStatus_up BagAclNhAtStatus = "up"

    // AT State DOWN
    BagAclNhAtStatus_down BagAclNhAtStatus = "down"

    // AT State Not Present
    BagAclNhAtStatus_not_present BagAclNhAtStatus = "not-present"

    // invalid status
    BagAclNhAtStatus_max BagAclNhAtStatus = "max"
)

// BagAclNh represents Bag acl nh
type BagAclNh string

const (
    // Next Hop None
    BagAclNh_nexthop_none BagAclNh = "nexthop-none"

    // Nexthop Default
    BagAclNh_nexthop_default BagAclNh = "nexthop-default"

    // Nexthop
    BagAclNh_nexthop BagAclNh = "nexthop"
)

// AclPortOperator represents Acl port operator
type AclPortOperator string

const (
    // None
    AclPortOperator_none AclPortOperator = "none"

    // Equal
    AclPortOperator_eq AclPortOperator = "eq"

    // Greater than
    AclPortOperator_gt AclPortOperator = "gt"

    // Less than
    AclPortOperator_lt AclPortOperator = "lt"

    // Not Equal
    AclPortOperator_neq AclPortOperator = "neq"

    // Range
    AclPortOperator_range_ AclPortOperator = "range"

    // One Byte
    AclPortOperator_onebyte AclPortOperator = "onebyte"

    // Two Bytes
    AclPortOperator_twobytes AclPortOperator = "twobytes"
)

// AclAction represents Acl action
type AclAction string

const (
    // Deny
    AclAction_deny AclAction = "deny"

    // Permit
    AclAction_permit AclAction = "permit"

    // Encrypt
    AclAction_encrypt AclAction = "encrypt"

    // Bypass
    AclAction_bypass AclAction = "bypass"

    // Fallthrough
    AclAction_fallthrough_ AclAction = "fallthrough"

    // Invalid
    AclAction_invalid AclAction = "invalid"
)

// BagAclNhStatus represents Bag acl nh status
type BagAclNhStatus string

const (
    // State Not Present
    BagAclNhStatus_not_present BagAclNhStatus = "not-present"

    // State Unknown
    BagAclNhStatus_unknown BagAclNhStatus = "unknown"

    // State DOWN
    BagAclNhStatus_down BagAclNhStatus = "down"

    // State UP
    BagAclNhStatus_up BagAclNhStatus = "up"

    // invalid status
    BagAclNhStatus_max BagAclNhStatus = "max"
)

// AclPortOperator represents Acl port operator
type AclPortOperator string

const (
    // None
    AclPortOperator_none AclPortOperator = "none"

    // Equal
    AclPortOperator_eq AclPortOperator = "eq"

    // Greater than
    AclPortOperator_gt AclPortOperator = "gt"

    // Less than
    AclPortOperator_lt AclPortOperator = "lt"

    // Not Equal
    AclPortOperator_neq AclPortOperator = "neq"

    // Range
    AclPortOperator_range_ AclPortOperator = "range"

    // One Byte
    AclPortOperator_onebyte AclPortOperator = "onebyte"

    // Two Bytes
    AclPortOperator_twobytes AclPortOperator = "twobytes"
)

// AclLog represents Acl log
type AclLog string

const (
    // Log None
    AclLog_log_none AclLog = "log-none"

    // Log Regular
    AclLog_log AclLog = "log"

    // Log Input
    AclLog_log_input AclLog = "log-input"
)

// AclAce1 represents ACE Types
type AclAce1 string

const (
    // This is Normal ACE
    AclAce1_normal AclAce1 = "normal"

    // This is Remark ACE
    AclAce1_remark AclAce1 = "remark"

    // This is ABF ACE
    AclAce1_abf AclAce1 = "abf"
)

// Ipv4AclAndPrefixList
// Root class of IPv4 Oper schema tree
type Ipv4AclAndPrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access list manager containing access lists and prefix lists.
    AccessListManager Ipv4AclAndPrefixList_AccessListManager

    // Out Of Resources, Limits to the resources allocatable.
    Oor Ipv4AclAndPrefixList_Oor
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetFilter() yfilter.YFilter { return ipv4AclAndPrefixList.YFilter }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) SetFilter(yf yfilter.YFilter) { ipv4AclAndPrefixList.YFilter = yf }

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetGoName(yname string) string {
    if yname == "access-list-manager" { return "AccessListManager" }
    if yname == "oor" { return "Oor" }
    return ""
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list"
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-manager" {
        return &ipv4AclAndPrefixList.AccessListManager
    }
    if childYangName == "oor" {
        return &ipv4AclAndPrefixList.Oor
    }
    return nil
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list-manager"] = &ipv4AclAndPrefixList.AccessListManager
    children["oor"] = &ipv4AclAndPrefixList.Oor
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

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-acl-oper" }

// Ipv4AclAndPrefixList_AccessListManager
// Access list manager containing access lists and
// prefix lists
type Ipv4AclAndPrefixList_AccessListManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of prefix lists.
    Prefixes Ipv4AclAndPrefixList_AccessListManager_Prefixes

    // Access listL class displaying Usage and Entries.
    Accesses Ipv4AclAndPrefixList_AccessListManager_Accesses

    // Table of Usage statistics of access lists at different nodes.
    Usages Ipv4AclAndPrefixList_AccessListManager_Usages
}

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetFilter() yfilter.YFilter { return accessListManager.YFilter }

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) SetFilter(yf yfilter.YFilter) { accessListManager.YFilter = yf }

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetGoName(yname string) string {
    if yname == "prefixes" { return "Prefixes" }
    if yname == "accesses" { return "Accesses" }
    if yname == "usages" { return "Usages" }
    return ""
}

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetSegmentPath() string {
    return "access-list-manager"
}

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefixes" {
        return &accessListManager.Prefixes
    }
    if childYangName == "accesses" {
        return &accessListManager.Accesses
    }
    if childYangName == "usages" {
        return &accessListManager.Usages
    }
    return nil
}

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefixes"] = &accessListManager.Prefixes
    children["accesses"] = &accessListManager.Accesses
    children["usages"] = &accessListManager.Usages
    return children
}

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetBundleName() string { return "cisco_ios_xr" }

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetYangName() string { return "access-list-manager" }

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) SetParent(parent types.Entity) { accessListManager.parent = parent }

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetParent() types.Entity { return accessListManager.parent }

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetParentYangName() string { return "ipv4-acl-and-prefix-list" }

// Ipv4AclAndPrefixList_AccessListManager_Prefixes
// Table of prefix lists
type Ipv4AclAndPrefixList_AccessListManager_Prefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the prefix list. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix.
    Prefix []Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix
}

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetFilter() yfilter.YFilter { return prefixes.YFilter }

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) SetFilter(yf yfilter.YFilter) { prefixes.YFilter = yf }

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetSegmentPath() string {
    return "prefixes"
}

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        for _, c := range prefixes.Prefix {
            if prefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix{}
        prefixes.Prefix = append(prefixes.Prefix, child)
        return &prefixes.Prefix[len(prefixes.Prefix)-1]
    }
    return nil
}

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixes.Prefix {
        children[prefixes.Prefix[i].GetSegmentPath()] = &prefixes.Prefix[i]
    }
    return children
}

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetBundleName() string { return "cisco_ios_xr" }

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetYangName() string { return "prefixes" }

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) SetParent(parent types.Entity) { prefixes.parent = parent }

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetParent() types.Entity { return prefixes.parent }

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetParentYangName() string { return "access-list-manager" }

// Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix
// Name of the prefix list
type Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the prefix list. The type is string.
    PrefixListName interface{}

    // Table of all the SequenceNumbers per prefix list.
    PrefixListSequences Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences
}

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetGoName(yname string) string {
    if yname == "prefix-list-name" { return "PrefixListName" }
    if yname == "prefix-list-sequences" { return "PrefixListSequences" }
    return ""
}

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetSegmentPath() string {
    return "prefix" + "[prefix-list-name='" + fmt.Sprintf("%v", prefix.PrefixListName) + "']"
}

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list-sequences" {
        return &prefix.PrefixListSequences
    }
    return nil
}

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-list-sequences"] = &prefix.PrefixListSequences
    return children
}

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-list-name"] = prefix.PrefixListName
    return leafs
}

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetYangName() string { return "prefix" }

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetParentYangName() string { return "prefixes" }

// Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences
// Table of all the SequenceNumbers per prefix
// list
type Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sequence Number of a prefix list entry. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence.
    PrefixListSequence []Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence
}

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetFilter() yfilter.YFilter { return prefixListSequences.YFilter }

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) SetFilter(yf yfilter.YFilter) { prefixListSequences.YFilter = yf }

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetGoName(yname string) string {
    if yname == "prefix-list-sequence" { return "PrefixListSequence" }
    return ""
}

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetSegmentPath() string {
    return "prefix-list-sequences"
}

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list-sequence" {
        for _, c := range prefixListSequences.PrefixListSequence {
            if prefixListSequences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence{}
        prefixListSequences.PrefixListSequence = append(prefixListSequences.PrefixListSequence, child)
        return &prefixListSequences.PrefixListSequence[len(prefixListSequences.PrefixListSequence)-1]
    }
    return nil
}

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixListSequences.PrefixListSequence {
        children[prefixListSequences.PrefixListSequence[i].GetSegmentPath()] = &prefixListSequences.PrefixListSequence[i]
    }
    return children
}

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetYangName() string { return "prefix-list-sequences" }

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) SetParent(parent types.Entity) { prefixListSequences.parent = parent }

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetParent() types.Entity { return prefixListSequences.parent }

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetParentYangName() string { return "prefix" }

// Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence
// Sequence Number of a prefix list entry
type Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence Number of the prefix list entry. The type
    // is interface{} with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (prefix, remark). The type is AclAce1.
    ItemType interface{}

    // ACLE sequence number. The type is interface{} with range: 0..4294967295.
    Sequence interface{}

    // Grant value permit/deny . The type is AclAction.
    Grant interface{}

    // Prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length . The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Port Operator. The type is AclPortOperator.
    Operator interface{}

    // Min length. The type is interface{} with range: 0..4294967295.
    MinimumLength interface{}

    // Maximum length. The type is interface{} with range: 0..4294967295.
    MaximumLength interface{}

    // Number of hits. The type is interface{} with range: 0..4294967295.
    Hits interface{}

    // Remark String. The type is string.
    Remark interface{}

    // ACL Name. The type is string.
    AclName interface{}
}

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetFilter() yfilter.YFilter { return prefixListSequence.YFilter }

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) SetFilter(yf yfilter.YFilter) { prefixListSequence.YFilter = yf }

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "item-type" { return "ItemType" }
    if yname == "sequence" { return "Sequence" }
    if yname == "grant" { return "Grant" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "operator" { return "Operator" }
    if yname == "minimum-length" { return "MinimumLength" }
    if yname == "maximum-length" { return "MaximumLength" }
    if yname == "hits" { return "Hits" }
    if yname == "remark" { return "Remark" }
    if yname == "acl-name" { return "AclName" }
    return ""
}

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetSegmentPath() string {
    return "prefix-list-sequence" + "[sequence-number='" + fmt.Sprintf("%v", prefixListSequence.SequenceNumber) + "']"
}

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = prefixListSequence.SequenceNumber
    leafs["item-type"] = prefixListSequence.ItemType
    leafs["sequence"] = prefixListSequence.Sequence
    leafs["grant"] = prefixListSequence.Grant
    leafs["prefix"] = prefixListSequence.Prefix
    leafs["prefix-length"] = prefixListSequence.PrefixLength
    leafs["operator"] = prefixListSequence.Operator
    leafs["minimum-length"] = prefixListSequence.MinimumLength
    leafs["maximum-length"] = prefixListSequence.MaximumLength
    leafs["hits"] = prefixListSequence.Hits
    leafs["remark"] = prefixListSequence.Remark
    leafs["acl-name"] = prefixListSequence.AclName
    return leafs
}

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetYangName() string { return "prefix-list-sequence" }

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) SetParent(parent types.Entity) { prefixListSequence.parent = parent }

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetParent() types.Entity { return prefixListSequence.parent }

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetParentYangName() string { return "prefix-list-sequences" }

// Ipv4AclAndPrefixList_AccessListManager_Accesses
// Access listL class displaying Usage and Entries
type Ipv4AclAndPrefixList_AccessListManager_Accesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the Access List. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Accesses_Access.
    Access []Ipv4AclAndPrefixList_AccessListManager_Accesses_Access
}

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetFilter() yfilter.YFilter { return accesses.YFilter }

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) SetFilter(yf yfilter.YFilter) { accesses.YFilter = yf }

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetGoName(yname string) string {
    if yname == "access" { return "Access" }
    return ""
}

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetSegmentPath() string {
    return "accesses"
}

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access" {
        for _, c := range accesses.Access {
            if accesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_AccessListManager_Accesses_Access{}
        accesses.Access = append(accesses.Access, child)
        return &accesses.Access[len(accesses.Access)-1]
    }
    return nil
}

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accesses.Access {
        children[accesses.Access[i].GetSegmentPath()] = &accesses.Access[i]
    }
    return children
}

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetBundleName() string { return "cisco_ios_xr" }

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetYangName() string { return "accesses" }

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) SetParent(parent types.Entity) { accesses.parent = parent }

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetParent() types.Entity { return accesses.parent }

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetParentYangName() string { return "access-list-manager" }

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access
// Name of the Access List
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Access List. The type is string.
    AccessListName interface{}

    // Table of all the SequenceNumbers per access list.
    AccessListSequences Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences

    // Object Group in an Access list.
    ObjectGroup Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup
}

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetFilter() yfilter.YFilter { return access.YFilter }

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) SetFilter(yf yfilter.YFilter) { access.YFilter = yf }

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetGoName(yname string) string {
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "access-list-sequences" { return "AccessListSequences" }
    if yname == "object-group" { return "ObjectGroup" }
    return ""
}

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetSegmentPath() string {
    return "access" + "[access-list-name='" + fmt.Sprintf("%v", access.AccessListName) + "']"
}

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-sequences" {
        return &access.AccessListSequences
    }
    if childYangName == "object-group" {
        return &access.ObjectGroup
    }
    return nil
}

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list-sequences"] = &access.AccessListSequences
    children["object-group"] = &access.ObjectGroup
    return children
}

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-list-name"] = access.AccessListName
    return leafs
}

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetBundleName() string { return "cisco_ios_xr" }

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetYangName() string { return "access" }

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) SetParent(parent types.Entity) { access.parent = parent }

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetParent() types.Entity { return access.parent }

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetParentYangName() string { return "accesses" }

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences
// Table of all the SequenceNumbers per access
// list
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sequence Number of an access list entry. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence.
    AccessListSequence []Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence
}

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetFilter() yfilter.YFilter { return accessListSequences.YFilter }

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) SetFilter(yf yfilter.YFilter) { accessListSequences.YFilter = yf }

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetGoName(yname string) string {
    if yname == "access-list-sequence" { return "AccessListSequence" }
    return ""
}

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetSegmentPath() string {
    return "access-list-sequences"
}

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-sequence" {
        for _, c := range accessListSequences.AccessListSequence {
            if accessListSequences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence{}
        accessListSequences.AccessListSequence = append(accessListSequences.AccessListSequence, child)
        return &accessListSequences.AccessListSequence[len(accessListSequences.AccessListSequence)-1]
    }
    return nil
}

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessListSequences.AccessListSequence {
        children[accessListSequences.AccessListSequence[i].GetSegmentPath()] = &accessListSequences.AccessListSequence[i]
    }
    return children
}

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetBundleName() string { return "cisco_ios_xr" }

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetYangName() string { return "access-list-sequences" }

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) SetParent(parent types.Entity) { accessListSequences.parent = parent }

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetParent() types.Entity { return accessListSequences.parent }

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetParentYangName() string { return "access" }

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence
// Sequence Number of an access list entry
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ACLEntry Sequence Number. The type is interface{}
    // with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (acl, remark). The type is AclAce1.
    ItemType interface{}

    // ACLE sequence number. The type is interface{} with range: 0..4294967295.
    Sequence interface{}

    // Permit/deny. The type is AclAction.
    Grant interface{}

    // IPv4 protocol operator. The type is interface{} with range: 0..65535.
    ProtocolOperator interface{}

    // IPv4 protocol type. The type is interface{} with range: 0..65535.
    Protocol interface{}

    // IPv4 protocol 2. The type is interface{} with range: 0..65535.
    Protocol2 interface{}

    // Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Source mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddressMask interface{}

    // Destination address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Destination mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddressMask interface{}

    // Source operator. The type is AclPortOperator.
    SourceOperator interface{}

    // Source port 1. The type is interface{} with range: 0..65535.
    SourcePort1 interface{}

    // Source port 2. The type is interface{} with range: 0..65535.
    SourcePort2 interface{}

    // Deprecated by Source operator. The type is AclPortOperator.
    SorceOperator interface{}

    // Deprecated by SourcePort1. The type is interface{} with range: 0..65535.
    SorcePort1 interface{}

    // Deprecated by SourcePort2. The type is interface{} with range: 0..65535.
    SorcePort2 interface{}

    // Destination operator. The type is AclPortOperator.
    DestinationOperator interface{}

    // Destination port 1. The type is interface{} with range: 0..65535.
    DestinationPort1 interface{}

    // Destination port 2. The type is interface{} with range: 0..65535.
    DestinationPort2 interface{}

    // Log option. The type is AclLog.
    LogOption interface{}

    // Counter name. The type is string.
    CounterName interface{}

    // Capture option, TRUE if enabled. The type is bool.
    Capture interface{}

    // DSCP present. The type is bool.
    DscpPresent interface{}

    // DSCP or DSCP range start. The type is interface{} with range: 0..255.
    Dscp interface{}

    // DSCP Range End. The type is interface{} with range: 0..255.
    Dscp2 interface{}

    // DSCP Operator. The type is interface{} with range: 0..255.
    DscpOperator interface{}

    // Precedence present. The type is bool.
    PrecedencePresent interface{}

    // Precedence. The type is interface{} with range: 0..255.
    Precedence interface{}

    // TCP flags operator. The type is AclTcpflagsOperator.
    TcpFlagsOperator interface{}

    // TCP flags. The type is interface{} with range: 0..255.
    TcpFlags interface{}

    // TCP flags mask. The type is interface{} with range: 0..255.
    TcpFlagsMask interface{}

    // Fragments. The type is interface{} with range: 0..255.
    Fragments interface{}

    // Packet length operator. The type is AclPortOperator.
    PacketLengthOperator interface{}

    // Packet length 1. The type is interface{} with range: 0..65535.
    PacketLength1 interface{}

    // Packet length 2. The type is interface{} with range: 0..65535.
    PacketLength2 interface{}

    // TTL operator. The type is AclPortOperator.
    TtlOperator interface{}

    // TTL 1. The type is interface{} with range: 0..65535.
    Ttl1 interface{}

    // TTL 2. The type is interface{} with range: 0..65535.
    Ttl2 interface{}

    // No stats. The type is bool.
    NoStats interface{}

    // Number of hits. The type is interface{} with range:
    // 0..18446744073709551615.
    Hits interface{}

    // True if ICMP off. The type is bool.
    IsIcmpOff interface{}

    // Qos group number. The type is interface{} with range: 0..65535.
    QosGroup interface{}

    // Next hop type. The type is BagAclNh.
    NextHopType interface{}

    // Remark String. The type is string.
    Remark interface{}

    // Is dynamic ACE. The type is bool.
    Dynamic interface{}

    // Source prefix object-group. The type is string.
    SourcePrefixGroup interface{}

    // Destination prefix object-group. The type is string.
    DestinationPrefixGroup interface{}

    // Source port object-group. The type is string.
    SourcePortGroup interface{}

    // Destination port object-group. The type is string.
    DestinationPortGroup interface{}

    // ACL Name. The type is string.
    AclName interface{}

    // Sequence String. The type is string.
    SequenceStr interface{}

    // Fragment offset operator. The type is AclPortOperator.
    FragmentOffsetOperator interface{}

    // Fragment offset 1. The type is interface{} with range: 0..65535.
    FragmentOffset1 interface{}

    // Fragment offset 2. The type is interface{} with range: 0..65535.
    FragmentOffset2 interface{}

    // SET TTL. The type is interface{} with range: 0..65535.
    SetTtl interface{}

    // Fragment flags. The type is interface{} with range: 0..255.
    FragmentFlags interface{}

    // HW Next hop info.
    HwNextHopInfo Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo

    // Next hop info. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo.
    NextHopInfo []Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo

    // UDF BAG. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf.
    Udf []Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf
}

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetFilter() yfilter.YFilter { return accessListSequence.YFilter }

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) SetFilter(yf yfilter.YFilter) { accessListSequence.YFilter = yf }

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "item-type" { return "ItemType" }
    if yname == "sequence" { return "Sequence" }
    if yname == "grant" { return "Grant" }
    if yname == "protocol-operator" { return "ProtocolOperator" }
    if yname == "protocol" { return "Protocol" }
    if yname == "protocol2" { return "Protocol2" }
    if yname == "source-address" { return "SourceAddress" }
    if yname == "source-address-mask" { return "SourceAddressMask" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "destination-address-mask" { return "DestinationAddressMask" }
    if yname == "source-operator" { return "SourceOperator" }
    if yname == "source-port1" { return "SourcePort1" }
    if yname == "source-port2" { return "SourcePort2" }
    if yname == "sorce-operator" { return "SorceOperator" }
    if yname == "sorce-port1" { return "SorcePort1" }
    if yname == "sorce-port2" { return "SorcePort2" }
    if yname == "destination-operator" { return "DestinationOperator" }
    if yname == "destination-port1" { return "DestinationPort1" }
    if yname == "destination-port2" { return "DestinationPort2" }
    if yname == "log-option" { return "LogOption" }
    if yname == "counter-name" { return "CounterName" }
    if yname == "capture" { return "Capture" }
    if yname == "dscp-present" { return "DscpPresent" }
    if yname == "dscp" { return "Dscp" }
    if yname == "dscp2" { return "Dscp2" }
    if yname == "dscp-operator" { return "DscpOperator" }
    if yname == "precedence-present" { return "PrecedencePresent" }
    if yname == "precedence" { return "Precedence" }
    if yname == "tcp-flags-operator" { return "TcpFlagsOperator" }
    if yname == "tcp-flags" { return "TcpFlags" }
    if yname == "tcp-flags-mask" { return "TcpFlagsMask" }
    if yname == "fragments" { return "Fragments" }
    if yname == "packet-length-operator" { return "PacketLengthOperator" }
    if yname == "packet-length1" { return "PacketLength1" }
    if yname == "packet-length2" { return "PacketLength2" }
    if yname == "ttl-operator" { return "TtlOperator" }
    if yname == "ttl1" { return "Ttl1" }
    if yname == "ttl2" { return "Ttl2" }
    if yname == "no-stats" { return "NoStats" }
    if yname == "hits" { return "Hits" }
    if yname == "is-icmp-off" { return "IsIcmpOff" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "next-hop-type" { return "NextHopType" }
    if yname == "remark" { return "Remark" }
    if yname == "dynamic" { return "Dynamic" }
    if yname == "source-prefix-group" { return "SourcePrefixGroup" }
    if yname == "destination-prefix-group" { return "DestinationPrefixGroup" }
    if yname == "source-port-group" { return "SourcePortGroup" }
    if yname == "destination-port-group" { return "DestinationPortGroup" }
    if yname == "acl-name" { return "AclName" }
    if yname == "sequence-str" { return "SequenceStr" }
    if yname == "fragment-offset-operator" { return "FragmentOffsetOperator" }
    if yname == "fragment-offset1" { return "FragmentOffset1" }
    if yname == "fragment-offset2" { return "FragmentOffset2" }
    if yname == "set-ttl" { return "SetTtl" }
    if yname == "fragment-flags" { return "FragmentFlags" }
    if yname == "hw-next-hop-info" { return "HwNextHopInfo" }
    if yname == "next-hop-info" { return "NextHopInfo" }
    if yname == "udf" { return "Udf" }
    return ""
}

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetSegmentPath() string {
    return "access-list-sequence" + "[sequence-number='" + fmt.Sprintf("%v", accessListSequence.SequenceNumber) + "']"
}

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-next-hop-info" {
        return &accessListSequence.HwNextHopInfo
    }
    if childYangName == "next-hop-info" {
        for _, c := range accessListSequence.NextHopInfo {
            if accessListSequence.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo{}
        accessListSequence.NextHopInfo = append(accessListSequence.NextHopInfo, child)
        return &accessListSequence.NextHopInfo[len(accessListSequence.NextHopInfo)-1]
    }
    if childYangName == "udf" {
        for _, c := range accessListSequence.Udf {
            if accessListSequence.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf{}
        accessListSequence.Udf = append(accessListSequence.Udf, child)
        return &accessListSequence.Udf[len(accessListSequence.Udf)-1]
    }
    return nil
}

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hw-next-hop-info"] = &accessListSequence.HwNextHopInfo
    for i := range accessListSequence.NextHopInfo {
        children[accessListSequence.NextHopInfo[i].GetSegmentPath()] = &accessListSequence.NextHopInfo[i]
    }
    for i := range accessListSequence.Udf {
        children[accessListSequence.Udf[i].GetSegmentPath()] = &accessListSequence.Udf[i]
    }
    return children
}

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = accessListSequence.SequenceNumber
    leafs["item-type"] = accessListSequence.ItemType
    leafs["sequence"] = accessListSequence.Sequence
    leafs["grant"] = accessListSequence.Grant
    leafs["protocol-operator"] = accessListSequence.ProtocolOperator
    leafs["protocol"] = accessListSequence.Protocol
    leafs["protocol2"] = accessListSequence.Protocol2
    leafs["source-address"] = accessListSequence.SourceAddress
    leafs["source-address-mask"] = accessListSequence.SourceAddressMask
    leafs["destination-address"] = accessListSequence.DestinationAddress
    leafs["destination-address-mask"] = accessListSequence.DestinationAddressMask
    leafs["source-operator"] = accessListSequence.SourceOperator
    leafs["source-port1"] = accessListSequence.SourcePort1
    leafs["source-port2"] = accessListSequence.SourcePort2
    leafs["sorce-operator"] = accessListSequence.SorceOperator
    leafs["sorce-port1"] = accessListSequence.SorcePort1
    leafs["sorce-port2"] = accessListSequence.SorcePort2
    leafs["destination-operator"] = accessListSequence.DestinationOperator
    leafs["destination-port1"] = accessListSequence.DestinationPort1
    leafs["destination-port2"] = accessListSequence.DestinationPort2
    leafs["log-option"] = accessListSequence.LogOption
    leafs["counter-name"] = accessListSequence.CounterName
    leafs["capture"] = accessListSequence.Capture
    leafs["dscp-present"] = accessListSequence.DscpPresent
    leafs["dscp"] = accessListSequence.Dscp
    leafs["dscp2"] = accessListSequence.Dscp2
    leafs["dscp-operator"] = accessListSequence.DscpOperator
    leafs["precedence-present"] = accessListSequence.PrecedencePresent
    leafs["precedence"] = accessListSequence.Precedence
    leafs["tcp-flags-operator"] = accessListSequence.TcpFlagsOperator
    leafs["tcp-flags"] = accessListSequence.TcpFlags
    leafs["tcp-flags-mask"] = accessListSequence.TcpFlagsMask
    leafs["fragments"] = accessListSequence.Fragments
    leafs["packet-length-operator"] = accessListSequence.PacketLengthOperator
    leafs["packet-length1"] = accessListSequence.PacketLength1
    leafs["packet-length2"] = accessListSequence.PacketLength2
    leafs["ttl-operator"] = accessListSequence.TtlOperator
    leafs["ttl1"] = accessListSequence.Ttl1
    leafs["ttl2"] = accessListSequence.Ttl2
    leafs["no-stats"] = accessListSequence.NoStats
    leafs["hits"] = accessListSequence.Hits
    leafs["is-icmp-off"] = accessListSequence.IsIcmpOff
    leafs["qos-group"] = accessListSequence.QosGroup
    leafs["next-hop-type"] = accessListSequence.NextHopType
    leafs["remark"] = accessListSequence.Remark
    leafs["dynamic"] = accessListSequence.Dynamic
    leafs["source-prefix-group"] = accessListSequence.SourcePrefixGroup
    leafs["destination-prefix-group"] = accessListSequence.DestinationPrefixGroup
    leafs["source-port-group"] = accessListSequence.SourcePortGroup
    leafs["destination-port-group"] = accessListSequence.DestinationPortGroup
    leafs["acl-name"] = accessListSequence.AclName
    leafs["sequence-str"] = accessListSequence.SequenceStr
    leafs["fragment-offset-operator"] = accessListSequence.FragmentOffsetOperator
    leafs["fragment-offset1"] = accessListSequence.FragmentOffset1
    leafs["fragment-offset2"] = accessListSequence.FragmentOffset2
    leafs["set-ttl"] = accessListSequence.SetTtl
    leafs["fragment-flags"] = accessListSequence.FragmentFlags
    return leafs
}

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetBundleName() string { return "cisco_ios_xr" }

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetYangName() string { return "access-list-sequence" }

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) SetParent(parent types.Entity) { accessListSequence.parent = parent }

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetParent() types.Entity { return accessListSequence.parent }

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetParentYangName() string { return "access-list-sequences" }

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo
// HW Next hop info
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The Next Hop. The type is interface{} with range: 0..4294967295.
    NextHop interface{}

    // the next-hop type. The type is BagAclNh.
    Type interface{}

    // VRF name. The type is string with length: 0..32.
    VrfName interface{}
}

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetFilter() yfilter.YFilter { return hwNextHopInfo.YFilter }

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) SetFilter(yf yfilter.YFilter) { hwNextHopInfo.YFilter = yf }

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "type" { return "Type" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetSegmentPath() string {
    return "hw-next-hop-info"
}

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = hwNextHopInfo.NextHop
    leafs["type"] = hwNextHopInfo.Type
    leafs["vrf-name"] = hwNextHopInfo.VrfName
    return leafs
}

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetBundleName() string { return "cisco_ios_xr" }

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetYangName() string { return "hw-next-hop-info" }

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) SetParent(parent types.Entity) { hwNextHopInfo.parent = parent }

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetParent() types.Entity { return hwNextHopInfo.parent }

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetParentYangName() string { return "access-list-sequence" }

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo
// Next hop info
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The next hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Track name. The type is string.
    TrackName interface{}

    // The next hop status. The type is BagAclNhStatus.
    Status interface{}

    // The next hop at status. The type is BagAclNhAtStatus.
    AtStatus interface{}

    // The nexthop exist. The type is bool.
    IsAclNextHopExist interface{}
}

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetFilter() yfilter.YFilter { return nextHopInfo.YFilter }

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) SetFilter(yf yfilter.YFilter) { nextHopInfo.YFilter = yf }

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "track-name" { return "TrackName" }
    if yname == "status" { return "Status" }
    if yname == "at-status" { return "AtStatus" }
    if yname == "is-acl-next-hop-exist" { return "IsAclNextHopExist" }
    return ""
}

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetSegmentPath() string {
    return "next-hop-info"
}

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHopInfo.NextHop
    leafs["track-name"] = nextHopInfo.TrackName
    leafs["status"] = nextHopInfo.Status
    leafs["at-status"] = nextHopInfo.AtStatus
    leafs["is-acl-next-hop-exist"] = nextHopInfo.IsAclNextHopExist
    return leafs
}

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetBundleName() string { return "cisco_ios_xr" }

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetYangName() string { return "next-hop-info" }

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) SetParent(parent types.Entity) { nextHopInfo.parent = parent }

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetParent() types.Entity { return nextHopInfo.parent }

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetParentYangName() string { return "access-list-sequence" }

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf
// UDF BAG
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // UDF Name. The type is string with length: 0..17.
    UdfName interface{}

    // UDF Value. The type is interface{} with range: 0..4294967295.
    UdfValue interface{}

    // UDF Mask. The type is interface{} with range: 0..4294967295.
    UdfMask interface{}
}

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetFilter() yfilter.YFilter { return udf.YFilter }

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) SetFilter(yf yfilter.YFilter) { udf.YFilter = yf }

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetGoName(yname string) string {
    if yname == "udf-name" { return "UdfName" }
    if yname == "udf-value" { return "UdfValue" }
    if yname == "udf-mask" { return "UdfMask" }
    return ""
}

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetSegmentPath() string {
    return "udf"
}

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udf-name"] = udf.UdfName
    leafs["udf-value"] = udf.UdfValue
    leafs["udf-mask"] = udf.UdfMask
    return leafs
}

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetBundleName() string { return "cisco_ios_xr" }

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetYangName() string { return "udf" }

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) SetParent(parent types.Entity) { udf.parent = parent }

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetParent() types.Entity { return udf.parent }

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetParentYangName() string { return "access-list-sequence" }

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup
// Object Group in an Access list
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object-group info. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo.
    ObjGrpInfo []Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo
}

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetFilter() yfilter.YFilter { return objectGroup.YFilter }

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) SetFilter(yf yfilter.YFilter) { objectGroup.YFilter = yf }

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetGoName(yname string) string {
    if yname == "obj-grp-info" { return "ObjGrpInfo" }
    return ""
}

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetSegmentPath() string {
    return "object-group"
}

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "obj-grp-info" {
        for _, c := range objectGroup.ObjGrpInfo {
            if objectGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo{}
        objectGroup.ObjGrpInfo = append(objectGroup.ObjGrpInfo, child)
        return &objectGroup.ObjGrpInfo[len(objectGroup.ObjGrpInfo)-1]
    }
    return nil
}

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range objectGroup.ObjGrpInfo {
        children[objectGroup.ObjGrpInfo[i].GetSegmentPath()] = &objectGroup.ObjGrpInfo[i]
    }
    return children
}

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetBundleName() string { return "cisco_ios_xr" }

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetYangName() string { return "object-group" }

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) SetParent(parent types.Entity) { objectGroup.parent = parent }

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetParent() types.Entity { return objectGroup.parent }

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetParentYangName() string { return "access" }

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo
// Object-group info
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object-group name. The type is string with length: 0..64.
    ObjGrpName interface{}

    // Object-group Type. The type is interface{} with range: 0..4294967295.
    ObjGrpType interface{}
}

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetFilter() yfilter.YFilter { return objGrpInfo.YFilter }

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) SetFilter(yf yfilter.YFilter) { objGrpInfo.YFilter = yf }

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetGoName(yname string) string {
    if yname == "obj-grp-name" { return "ObjGrpName" }
    if yname == "obj-grp-type" { return "ObjGrpType" }
    return ""
}

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetSegmentPath() string {
    return "obj-grp-info"
}

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["obj-grp-name"] = objGrpInfo.ObjGrpName
    leafs["obj-grp-type"] = objGrpInfo.ObjGrpType
    return leafs
}

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetBundleName() string { return "cisco_ios_xr" }

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetYangName() string { return "obj-grp-info" }

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) SetParent(parent types.Entity) { objGrpInfo.parent = parent }

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetParent() types.Entity { return objGrpInfo.parent }

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetParentYangName() string { return "object-group" }

// Ipv4AclAndPrefixList_AccessListManager_Usages
// Table of Usage statistics of access lists at
// different nodes
type Ipv4AclAndPrefixList_AccessListManager_Usages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Usage statistics of an access list at a node. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Usages_Usage.
    Usage []Ipv4AclAndPrefixList_AccessListManager_Usages_Usage
}

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetFilter() yfilter.YFilter { return usages.YFilter }

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) SetFilter(yf yfilter.YFilter) { usages.YFilter = yf }

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetGoName(yname string) string {
    if yname == "usage" { return "Usage" }
    return ""
}

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetSegmentPath() string {
    return "usages"
}

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usage" {
        for _, c := range usages.Usage {
            if usages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_AccessListManager_Usages_Usage{}
        usages.Usage = append(usages.Usage, child)
        return &usages.Usage[len(usages.Usage)-1]
    }
    return nil
}

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usages.Usage {
        children[usages.Usage[i].GetSegmentPath()] = &usages.Usage[i]
    }
    return children
}

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetBundleName() string { return "cisco_ios_xr" }

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetYangName() string { return "usages" }

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) SetParent(parent types.Entity) { usages.parent = parent }

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetParent() types.Entity { return usages.parent }

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetParentYangName() string { return "access-list-manager" }

// Ipv4AclAndPrefixList_AccessListManager_Usages_Usage
// Usage statistics of an access list at a node
type Ipv4AclAndPrefixList_AccessListManager_Usages_Usage struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node where access list is applied. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Application ID. The type is AclUsageAppIdEnum.
    ApplicationId interface{}

    // Name of the access list. The type is string.
    AccessListName interface{}

    // Usage Statistics Details. The type is string. This attribute is mandatory.
    UsageDetails interface{}
}

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetFilter() yfilter.YFilter { return usage.YFilter }

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) SetFilter(yf yfilter.YFilter) { usage.YFilter = yf }

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "application-id" { return "ApplicationId" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "usage-details" { return "UsageDetails" }
    return ""
}

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetSegmentPath() string {
    return "usage"
}

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = usage.NodeName
    leafs["application-id"] = usage.ApplicationId
    leafs["access-list-name"] = usage.AccessListName
    leafs["usage-details"] = usage.UsageDetails
    return leafs
}

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetBundleName() string { return "cisco_ios_xr" }

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetYangName() string { return "usage" }

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) SetParent(parent types.Entity) { usage.parent = parent }

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetParent() types.Entity { return usage.parent }

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetParentYangName() string { return "usages" }

// Ipv4AclAndPrefixList_Oor
// Out Of Resources, Limits to the resources
// allocatable
type Ipv4AclAndPrefixList_Oor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details of the Overall Out Of Resources Limits.
    Details Ipv4AclAndPrefixList_Oor_Details

    // Resource occupation details for prefix lists.
    OorPrefixes Ipv4AclAndPrefixList_Oor_OorPrefixes

    // Resource occupation details for access lists.
    OorAccesses Ipv4AclAndPrefixList_Oor_OorAccesses

    // Resource limits pertaining to access lists only.
    AccessListSummary Ipv4AclAndPrefixList_Oor_AccessListSummary

    // Summary of the prefix Lists resource utilization.
    PrefixListSummary Ipv4AclAndPrefixList_Oor_PrefixListSummary
}

func (oor *Ipv4AclAndPrefixList_Oor) GetFilter() yfilter.YFilter { return oor.YFilter }

func (oor *Ipv4AclAndPrefixList_Oor) SetFilter(yf yfilter.YFilter) { oor.YFilter = yf }

func (oor *Ipv4AclAndPrefixList_Oor) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    if yname == "oor-prefixes" { return "OorPrefixes" }
    if yname == "oor-accesses" { return "OorAccesses" }
    if yname == "access-list-summary" { return "AccessListSummary" }
    if yname == "prefix-list-summary" { return "PrefixListSummary" }
    return ""
}

func (oor *Ipv4AclAndPrefixList_Oor) GetSegmentPath() string {
    return "oor"
}

func (oor *Ipv4AclAndPrefixList_Oor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &oor.Details
    }
    if childYangName == "oor-prefixes" {
        return &oor.OorPrefixes
    }
    if childYangName == "oor-accesses" {
        return &oor.OorAccesses
    }
    if childYangName == "access-list-summary" {
        return &oor.AccessListSummary
    }
    if childYangName == "prefix-list-summary" {
        return &oor.PrefixListSummary
    }
    return nil
}

func (oor *Ipv4AclAndPrefixList_Oor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &oor.Details
    children["oor-prefixes"] = &oor.OorPrefixes
    children["oor-accesses"] = &oor.OorAccesses
    children["access-list-summary"] = &oor.AccessListSummary
    children["prefix-list-summary"] = &oor.PrefixListSummary
    return children
}

func (oor *Ipv4AclAndPrefixList_Oor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oor *Ipv4AclAndPrefixList_Oor) GetBundleName() string { return "cisco_ios_xr" }

func (oor *Ipv4AclAndPrefixList_Oor) GetYangName() string { return "oor" }

func (oor *Ipv4AclAndPrefixList_Oor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oor *Ipv4AclAndPrefixList_Oor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oor *Ipv4AclAndPrefixList_Oor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oor *Ipv4AclAndPrefixList_Oor) SetParent(parent types.Entity) { oor.parent = parent }

func (oor *Ipv4AclAndPrefixList_Oor) GetParent() types.Entity { return oor.parent }

func (oor *Ipv4AclAndPrefixList_Oor) GetParentYangName() string { return "ipv4-acl-and-prefix-list" }

// Ipv4AclAndPrefixList_Oor_Details
// Details of the Overall Out Of Resources Limits
type Ipv4AclAndPrefixList_Oor_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcEs interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcLs interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcEs interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcEs interface{}
}

func (details *Ipv4AclAndPrefixList_Oor_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Ipv4AclAndPrefixList_Oor_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Ipv4AclAndPrefixList_Oor_Details) GetGoName(yname string) string {
    if yname == "default-max-ac-ls" { return "DefaultMaxAcLs" }
    if yname == "default-max-ac-es" { return "DefaultMaxAcEs" }
    if yname == "current-configured-ac-ls" { return "CurrentConfiguredAcLs" }
    if yname == "current-configured-ac-es" { return "CurrentConfiguredAcEs" }
    if yname == "current-max-configurable-ac-ls" { return "CurrentMaxConfigurableAcLs" }
    if yname == "current-max-configurable-ac-es" { return "CurrentMaxConfigurableAcEs" }
    if yname == "max-configurable-ac-ls" { return "MaxConfigurableAcLs" }
    if yname == "max-configurable-ac-es" { return "MaxConfigurableAcEs" }
    return ""
}

func (details *Ipv4AclAndPrefixList_Oor_Details) GetSegmentPath() string {
    return "details"
}

func (details *Ipv4AclAndPrefixList_Oor_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (details *Ipv4AclAndPrefixList_Oor_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (details *Ipv4AclAndPrefixList_Oor_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default-max-ac-ls"] = details.DefaultMaxAcLs
    leafs["default-max-ac-es"] = details.DefaultMaxAcEs
    leafs["current-configured-ac-ls"] = details.CurrentConfiguredAcLs
    leafs["current-configured-ac-es"] = details.CurrentConfiguredAcEs
    leafs["current-max-configurable-ac-ls"] = details.CurrentMaxConfigurableAcLs
    leafs["current-max-configurable-ac-es"] = details.CurrentMaxConfigurableAcEs
    leafs["max-configurable-ac-ls"] = details.MaxConfigurableAcLs
    leafs["max-configurable-ac-es"] = details.MaxConfigurableAcEs
    return leafs
}

func (details *Ipv4AclAndPrefixList_Oor_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Ipv4AclAndPrefixList_Oor_Details) GetYangName() string { return "details" }

func (details *Ipv4AclAndPrefixList_Oor_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Ipv4AclAndPrefixList_Oor_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Ipv4AclAndPrefixList_Oor_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Ipv4AclAndPrefixList_Oor_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Ipv4AclAndPrefixList_Oor_Details) GetParent() types.Entity { return details.parent }

func (details *Ipv4AclAndPrefixList_Oor_Details) GetParentYangName() string { return "oor" }

// Ipv4AclAndPrefixList_Oor_OorPrefixes
// Resource occupation details for prefix lists
type Ipv4AclAndPrefixList_Oor_OorPrefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource occupation details for a particular prefix list. The type is slice
    // of Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix.
    OorPrefix []Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix
}

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetFilter() yfilter.YFilter { return oorPrefixes.YFilter }

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) SetFilter(yf yfilter.YFilter) { oorPrefixes.YFilter = yf }

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetGoName(yname string) string {
    if yname == "oor-prefix" { return "OorPrefix" }
    return ""
}

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetSegmentPath() string {
    return "oor-prefixes"
}

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oor-prefix" {
        for _, c := range oorPrefixes.OorPrefix {
            if oorPrefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix{}
        oorPrefixes.OorPrefix = append(oorPrefixes.OorPrefix, child)
        return &oorPrefixes.OorPrefix[len(oorPrefixes.OorPrefix)-1]
    }
    return nil
}

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range oorPrefixes.OorPrefix {
        children[oorPrefixes.OorPrefix[i].GetSegmentPath()] = &oorPrefixes.OorPrefix[i]
    }
    return children
}

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetBundleName() string { return "cisco_ios_xr" }

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetYangName() string { return "oor-prefixes" }

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) SetParent(parent types.Entity) { oorPrefixes.parent = parent }

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetParent() types.Entity { return oorPrefixes.parent }

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetParentYangName() string { return "oor" }

// Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix
// Resource occupation details for a particular
// prefix list
type Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of a prefix list. The type is string.
    PrefixListName interface{}

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcEs interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcLs interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcEs interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcEs interface{}
}

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetFilter() yfilter.YFilter { return oorPrefix.YFilter }

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) SetFilter(yf yfilter.YFilter) { oorPrefix.YFilter = yf }

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetGoName(yname string) string {
    if yname == "prefix-list-name" { return "PrefixListName" }
    if yname == "default-max-ac-ls" { return "DefaultMaxAcLs" }
    if yname == "default-max-ac-es" { return "DefaultMaxAcEs" }
    if yname == "current-configured-ac-ls" { return "CurrentConfiguredAcLs" }
    if yname == "current-configured-ac-es" { return "CurrentConfiguredAcEs" }
    if yname == "current-max-configurable-ac-ls" { return "CurrentMaxConfigurableAcLs" }
    if yname == "current-max-configurable-ac-es" { return "CurrentMaxConfigurableAcEs" }
    if yname == "max-configurable-ac-ls" { return "MaxConfigurableAcLs" }
    if yname == "max-configurable-ac-es" { return "MaxConfigurableAcEs" }
    return ""
}

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetSegmentPath() string {
    return "oor-prefix" + "[prefix-list-name='" + fmt.Sprintf("%v", oorPrefix.PrefixListName) + "']"
}

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-list-name"] = oorPrefix.PrefixListName
    leafs["default-max-ac-ls"] = oorPrefix.DefaultMaxAcLs
    leafs["default-max-ac-es"] = oorPrefix.DefaultMaxAcEs
    leafs["current-configured-ac-ls"] = oorPrefix.CurrentConfiguredAcLs
    leafs["current-configured-ac-es"] = oorPrefix.CurrentConfiguredAcEs
    leafs["current-max-configurable-ac-ls"] = oorPrefix.CurrentMaxConfigurableAcLs
    leafs["current-max-configurable-ac-es"] = oorPrefix.CurrentMaxConfigurableAcEs
    leafs["max-configurable-ac-ls"] = oorPrefix.MaxConfigurableAcLs
    leafs["max-configurable-ac-es"] = oorPrefix.MaxConfigurableAcEs
    return leafs
}

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetYangName() string { return "oor-prefix" }

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) SetParent(parent types.Entity) { oorPrefix.parent = parent }

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetParent() types.Entity { return oorPrefix.parent }

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetParentYangName() string { return "oor-prefixes" }

// Ipv4AclAndPrefixList_Oor_OorAccesses
// Resource occupation details for access lists
type Ipv4AclAndPrefixList_Oor_OorAccesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource occupation details for a particular access list. The type is slice
    // of Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess.
    OorAccess []Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess
}

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetFilter() yfilter.YFilter { return oorAccesses.YFilter }

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) SetFilter(yf yfilter.YFilter) { oorAccesses.YFilter = yf }

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetGoName(yname string) string {
    if yname == "oor-access" { return "OorAccess" }
    return ""
}

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetSegmentPath() string {
    return "oor-accesses"
}

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oor-access" {
        for _, c := range oorAccesses.OorAccess {
            if oorAccesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess{}
        oorAccesses.OorAccess = append(oorAccesses.OorAccess, child)
        return &oorAccesses.OorAccess[len(oorAccesses.OorAccess)-1]
    }
    return nil
}

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range oorAccesses.OorAccess {
        children[oorAccesses.OorAccess[i].GetSegmentPath()] = &oorAccesses.OorAccess[i]
    }
    return children
}

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetBundleName() string { return "cisco_ios_xr" }

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetYangName() string { return "oor-accesses" }

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) SetParent(parent types.Entity) { oorAccesses.parent = parent }

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetParent() types.Entity { return oorAccesses.parent }

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetParentYangName() string { return "oor" }

// Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess
// Resource occupation details for a particular
// access list
type Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Access List. The type is string.
    AccessListName interface{}

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcEs interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcLs interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcEs interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcEs interface{}
}

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetFilter() yfilter.YFilter { return oorAccess.YFilter }

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) SetFilter(yf yfilter.YFilter) { oorAccess.YFilter = yf }

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetGoName(yname string) string {
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "default-max-ac-ls" { return "DefaultMaxAcLs" }
    if yname == "default-max-ac-es" { return "DefaultMaxAcEs" }
    if yname == "current-configured-ac-ls" { return "CurrentConfiguredAcLs" }
    if yname == "current-configured-ac-es" { return "CurrentConfiguredAcEs" }
    if yname == "current-max-configurable-ac-ls" { return "CurrentMaxConfigurableAcLs" }
    if yname == "current-max-configurable-ac-es" { return "CurrentMaxConfigurableAcEs" }
    if yname == "max-configurable-ac-ls" { return "MaxConfigurableAcLs" }
    if yname == "max-configurable-ac-es" { return "MaxConfigurableAcEs" }
    return ""
}

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetSegmentPath() string {
    return "oor-access" + "[access-list-name='" + fmt.Sprintf("%v", oorAccess.AccessListName) + "']"
}

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-list-name"] = oorAccess.AccessListName
    leafs["default-max-ac-ls"] = oorAccess.DefaultMaxAcLs
    leafs["default-max-ac-es"] = oorAccess.DefaultMaxAcEs
    leafs["current-configured-ac-ls"] = oorAccess.CurrentConfiguredAcLs
    leafs["current-configured-ac-es"] = oorAccess.CurrentConfiguredAcEs
    leafs["current-max-configurable-ac-ls"] = oorAccess.CurrentMaxConfigurableAcLs
    leafs["current-max-configurable-ac-es"] = oorAccess.CurrentMaxConfigurableAcEs
    leafs["max-configurable-ac-ls"] = oorAccess.MaxConfigurableAcLs
    leafs["max-configurable-ac-es"] = oorAccess.MaxConfigurableAcEs
    return leafs
}

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetBundleName() string { return "cisco_ios_xr" }

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetYangName() string { return "oor-access" }

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) SetParent(parent types.Entity) { oorAccess.parent = parent }

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetParent() types.Entity { return oorAccess.parent }

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetParentYangName() string { return "oor-accesses" }

// Ipv4AclAndPrefixList_Oor_AccessListSummary
// Resource limits pertaining to access lists only
type Ipv4AclAndPrefixList_Oor_AccessListSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details containing the resource limits of the access lists.
    Details Ipv4AclAndPrefixList_Oor_AccessListSummary_Details
}

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetFilter() yfilter.YFilter { return accessListSummary.YFilter }

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) SetFilter(yf yfilter.YFilter) { accessListSummary.YFilter = yf }

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    return ""
}

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetSegmentPath() string {
    return "access-list-summary"
}

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &accessListSummary.Details
    }
    return nil
}

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &accessListSummary.Details
    return children
}

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetBundleName() string { return "cisco_ios_xr" }

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetYangName() string { return "access-list-summary" }

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) SetParent(parent types.Entity) { accessListSummary.parent = parent }

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetParent() types.Entity { return accessListSummary.parent }

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetParentYangName() string { return "oor" }

// Ipv4AclAndPrefixList_Oor_AccessListSummary_Details
// Details containing the resource limits of the
// access lists
type Ipv4AclAndPrefixList_Oor_AccessListSummary_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcEs interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcLs interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcEs interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcEs interface{}
}

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetGoName(yname string) string {
    if yname == "default-max-ac-ls" { return "DefaultMaxAcLs" }
    if yname == "default-max-ac-es" { return "DefaultMaxAcEs" }
    if yname == "current-configured-ac-ls" { return "CurrentConfiguredAcLs" }
    if yname == "current-configured-ac-es" { return "CurrentConfiguredAcEs" }
    if yname == "current-max-configurable-ac-ls" { return "CurrentMaxConfigurableAcLs" }
    if yname == "current-max-configurable-ac-es" { return "CurrentMaxConfigurableAcEs" }
    if yname == "max-configurable-ac-ls" { return "MaxConfigurableAcLs" }
    if yname == "max-configurable-ac-es" { return "MaxConfigurableAcEs" }
    return ""
}

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetSegmentPath() string {
    return "details"
}

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default-max-ac-ls"] = details.DefaultMaxAcLs
    leafs["default-max-ac-es"] = details.DefaultMaxAcEs
    leafs["current-configured-ac-ls"] = details.CurrentConfiguredAcLs
    leafs["current-configured-ac-es"] = details.CurrentConfiguredAcEs
    leafs["current-max-configurable-ac-ls"] = details.CurrentMaxConfigurableAcLs
    leafs["current-max-configurable-ac-es"] = details.CurrentMaxConfigurableAcEs
    leafs["max-configurable-ac-ls"] = details.MaxConfigurableAcLs
    leafs["max-configurable-ac-es"] = details.MaxConfigurableAcEs
    return leafs
}

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetYangName() string { return "details" }

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetParent() types.Entity { return details.parent }

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetParentYangName() string { return "access-list-summary" }

// Ipv4AclAndPrefixList_Oor_PrefixListSummary
// Summary of the prefix Lists resource
// utilization
type Ipv4AclAndPrefixList_Oor_PrefixListSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary Detail of the prefix list Resource Utilisation.
    Details Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details
}

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetFilter() yfilter.YFilter { return prefixListSummary.YFilter }

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) SetFilter(yf yfilter.YFilter) { prefixListSummary.YFilter = yf }

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    return ""
}

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetSegmentPath() string {
    return "prefix-list-summary"
}

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &prefixListSummary.Details
    }
    return nil
}

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &prefixListSummary.Details
    return children
}

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetYangName() string { return "prefix-list-summary" }

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) SetParent(parent types.Entity) { prefixListSummary.parent = parent }

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetParent() types.Entity { return prefixListSummary.parent }

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetParentYangName() string { return "oor" }

// Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details
// Summary Detail of the prefix list Resource
// Utilisation
type Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    DefaultMaxAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    CurrentConfiguredAcEs interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcLs interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    CurrentMaxConfigurableAcEs interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    MaxConfigurableAcEs interface{}
}

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetGoName(yname string) string {
    if yname == "default-max-ac-ls" { return "DefaultMaxAcLs" }
    if yname == "default-max-ac-es" { return "DefaultMaxAcEs" }
    if yname == "current-configured-ac-ls" { return "CurrentConfiguredAcLs" }
    if yname == "current-configured-ac-es" { return "CurrentConfiguredAcEs" }
    if yname == "current-max-configurable-ac-ls" { return "CurrentMaxConfigurableAcLs" }
    if yname == "current-max-configurable-ac-es" { return "CurrentMaxConfigurableAcEs" }
    if yname == "max-configurable-ac-ls" { return "MaxConfigurableAcLs" }
    if yname == "max-configurable-ac-es" { return "MaxConfigurableAcEs" }
    return ""
}

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetSegmentPath() string {
    return "details"
}

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default-max-ac-ls"] = details.DefaultMaxAcLs
    leafs["default-max-ac-es"] = details.DefaultMaxAcEs
    leafs["current-configured-ac-ls"] = details.CurrentConfiguredAcLs
    leafs["current-configured-ac-es"] = details.CurrentConfiguredAcEs
    leafs["current-max-configurable-ac-ls"] = details.CurrentMaxConfigurableAcLs
    leafs["current-max-configurable-ac-es"] = details.CurrentMaxConfigurableAcEs
    leafs["max-configurable-ac-ls"] = details.MaxConfigurableAcLs
    leafs["max-configurable-ac-es"] = details.MaxConfigurableAcEs
    return leafs
}

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetYangName() string { return "details" }

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetParent() types.Entity { return details.parent }

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetParentYangName() string { return "prefix-list-summary" }

