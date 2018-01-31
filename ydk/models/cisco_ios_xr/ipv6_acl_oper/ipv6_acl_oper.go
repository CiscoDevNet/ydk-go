// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-acl package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-acl-and-prefix-list: Root class of IPv6 Oper schema tree
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_acl_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_acl_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-acl-oper ipv6-acl-and-prefix-list}", reflect.TypeOf(Ipv6AclAndPrefixList{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-acl-oper:ipv6-acl-and-prefix-list", reflect.TypeOf(Ipv6AclAndPrefixList{}))
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

// Ipv6AclAndPrefixList
// Root class of IPv6 Oper schema tree
type Ipv6AclAndPrefixList struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AccessListManager containing ACLs and prefix lists.
    AccessListManager Ipv6AclAndPrefixList_AccessListManager

    // Out Of Resources, Limits to the resources allocatable.
    Oor Ipv6AclAndPrefixList_Oor
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetFilter() yfilter.YFilter { return ipv6AclAndPrefixList.YFilter }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) SetFilter(yf yfilter.YFilter) { ipv6AclAndPrefixList.YFilter = yf }

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetGoName(yname string) string {
    if yname == "access-list-manager" { return "AccessListManager" }
    if yname == "oor" { return "Oor" }
    return ""
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-acl-oper:ipv6-acl-and-prefix-list"
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-manager" {
        return &ipv6AclAndPrefixList.AccessListManager
    }
    if childYangName == "oor" {
        return &ipv6AclAndPrefixList.Oor
    }
    return nil
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list-manager"] = &ipv6AclAndPrefixList.AccessListManager
    children["oor"] = &ipv6AclAndPrefixList.Oor
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

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-acl-oper" }

// Ipv6AclAndPrefixList_AccessListManager
// AccessListManager containing ACLs and prefix
// lists
type Ipv6AclAndPrefixList_AccessListManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of prefix lists.
    Prefixes Ipv6AclAndPrefixList_AccessListManager_Prefixes

    // Table of Usage statistics of ACLs at different nodes.
    Usages Ipv6AclAndPrefixList_AccessListManager_Usages

    // ACL class displaying Usage and Entries.
    Accesses Ipv6AclAndPrefixList_AccessListManager_Accesses
}

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetFilter() yfilter.YFilter { return accessListManager.YFilter }

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) SetFilter(yf yfilter.YFilter) { accessListManager.YFilter = yf }

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetGoName(yname string) string {
    if yname == "prefixes" { return "Prefixes" }
    if yname == "usages" { return "Usages" }
    if yname == "accesses" { return "Accesses" }
    return ""
}

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetSegmentPath() string {
    return "access-list-manager"
}

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefixes" {
        return &accessListManager.Prefixes
    }
    if childYangName == "usages" {
        return &accessListManager.Usages
    }
    if childYangName == "accesses" {
        return &accessListManager.Accesses
    }
    return nil
}

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefixes"] = &accessListManager.Prefixes
    children["usages"] = &accessListManager.Usages
    children["accesses"] = &accessListManager.Accesses
    return children
}

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetBundleName() string { return "cisco_ios_xr" }

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetYangName() string { return "access-list-manager" }

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) SetParent(parent types.Entity) { accessListManager.parent = parent }

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetParent() types.Entity { return accessListManager.parent }

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetParentYangName() string { return "ipv6-acl-and-prefix-list" }

// Ipv6AclAndPrefixList_AccessListManager_Prefixes
// Table of prefix lists
type Ipv6AclAndPrefixList_AccessListManager_Prefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the prefix list. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix.
    Prefix []Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix
}

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetFilter() yfilter.YFilter { return prefixes.YFilter }

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) SetFilter(yf yfilter.YFilter) { prefixes.YFilter = yf }

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetSegmentPath() string {
    return "prefixes"
}

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        for _, c := range prefixes.Prefix {
            if prefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix{}
        prefixes.Prefix = append(prefixes.Prefix, child)
        return &prefixes.Prefix[len(prefixes.Prefix)-1]
    }
    return nil
}

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixes.Prefix {
        children[prefixes.Prefix[i].GetSegmentPath()] = &prefixes.Prefix[i]
    }
    return children
}

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetBundleName() string { return "cisco_ios_xr" }

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetYangName() string { return "prefixes" }

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) SetParent(parent types.Entity) { prefixes.parent = parent }

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetParent() types.Entity { return prefixes.parent }

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetParentYangName() string { return "access-list-manager" }

// Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix
// Name of the prefix list
type Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the prefix list. The type is string with
    // length: 1..65.
    PrefixListName interface{}

    // Table of all the SequenceNumbers per prefix list.
    PrefixListSequences Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences
}

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetGoName(yname string) string {
    if yname == "prefix-list-name" { return "PrefixListName" }
    if yname == "prefix-list-sequences" { return "PrefixListSequences" }
    return ""
}

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetSegmentPath() string {
    return "prefix" + "[prefix-list-name='" + fmt.Sprintf("%v", prefix.PrefixListName) + "']"
}

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list-sequences" {
        return &prefix.PrefixListSequences
    }
    return nil
}

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix-list-sequences"] = &prefix.PrefixListSequences
    return children
}

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-list-name"] = prefix.PrefixListName
    return leafs
}

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetYangName() string { return "prefix" }

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetParentYangName() string { return "prefixes" }

// Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences
// Table of all the SequenceNumbers per prefix
// list
type Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sequence Number of a prefix list entry. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence.
    PrefixListSequence []Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence
}

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetFilter() yfilter.YFilter { return prefixListSequences.YFilter }

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) SetFilter(yf yfilter.YFilter) { prefixListSequences.YFilter = yf }

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetGoName(yname string) string {
    if yname == "prefix-list-sequence" { return "PrefixListSequence" }
    return ""
}

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetSegmentPath() string {
    return "prefix-list-sequences"
}

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-list-sequence" {
        for _, c := range prefixListSequences.PrefixListSequence {
            if prefixListSequences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence{}
        prefixListSequences.PrefixListSequence = append(prefixListSequences.PrefixListSequence, child)
        return &prefixListSequences.PrefixListSequence[len(prefixListSequences.PrefixListSequence)-1]
    }
    return nil
}

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixListSequences.PrefixListSequence {
        children[prefixListSequences.PrefixListSequence[i].GetSegmentPath()] = &prefixListSequences.PrefixListSequence[i]
    }
    return children
}

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetYangName() string { return "prefix-list-sequences" }

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) SetParent(parent types.Entity) { prefixListSequences.parent = parent }

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetParent() types.Entity { return prefixListSequences.parent }

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetParentYangName() string { return "prefix" }

// Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence
// Sequence Number of a prefix list entry
type Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence Number of the prefix list entry. The type
    // is interface{} with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (acl, remark). The type is AclAce1.
    IsAceType interface{}

    // ACLE sequence number. The type is interface{} with range: 0..4294967295.
    IsAceSequenceNumber interface{}

    // Grant value permit/deny . The type is AclAction.
    IsPacketAllowOrDeny interface{}

    // IPv6 prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IsAddressInNumbers interface{}

    // Prefix length . The type is interface{} with range: 0..4294967295.
    IsAddressMaskLength interface{}

    // Port Operator. The type is AclPortOperator.
    IsLengthOperator interface{}

    // Min length. The type is interface{} with range: 0..4294967295.
    IsPacketMinimumLength interface{}

    // Maximum length. The type is interface{} with range: 0..4294967295.
    IsPacketMaximumLength interface{}

    // Number of hits. The type is interface{} with range: 0..4294967295.
    Hits interface{}

    // Remark String. The type is string.
    IsCommentForEntry interface{}

    // ACL Name. The type is string.
    AclName interface{}
}

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetFilter() yfilter.YFilter { return prefixListSequence.YFilter }

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) SetFilter(yf yfilter.YFilter) { prefixListSequence.YFilter = yf }

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "is-ace-type" { return "IsAceType" }
    if yname == "is-ace-sequence-number" { return "IsAceSequenceNumber" }
    if yname == "is-packet-allow-or-deny" { return "IsPacketAllowOrDeny" }
    if yname == "is-address-in-numbers" { return "IsAddressInNumbers" }
    if yname == "is-address-mask-length" { return "IsAddressMaskLength" }
    if yname == "is-length-operator" { return "IsLengthOperator" }
    if yname == "is-packet-minimum-length" { return "IsPacketMinimumLength" }
    if yname == "is-packet-maximum-length" { return "IsPacketMaximumLength" }
    if yname == "hits" { return "Hits" }
    if yname == "is-comment-for-entry" { return "IsCommentForEntry" }
    if yname == "acl-name" { return "AclName" }
    return ""
}

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetSegmentPath() string {
    return "prefix-list-sequence" + "[sequence-number='" + fmt.Sprintf("%v", prefixListSequence.SequenceNumber) + "']"
}

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = prefixListSequence.SequenceNumber
    leafs["is-ace-type"] = prefixListSequence.IsAceType
    leafs["is-ace-sequence-number"] = prefixListSequence.IsAceSequenceNumber
    leafs["is-packet-allow-or-deny"] = prefixListSequence.IsPacketAllowOrDeny
    leafs["is-address-in-numbers"] = prefixListSequence.IsAddressInNumbers
    leafs["is-address-mask-length"] = prefixListSequence.IsAddressMaskLength
    leafs["is-length-operator"] = prefixListSequence.IsLengthOperator
    leafs["is-packet-minimum-length"] = prefixListSequence.IsPacketMinimumLength
    leafs["is-packet-maximum-length"] = prefixListSequence.IsPacketMaximumLength
    leafs["hits"] = prefixListSequence.Hits
    leafs["is-comment-for-entry"] = prefixListSequence.IsCommentForEntry
    leafs["acl-name"] = prefixListSequence.AclName
    return leafs
}

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetYangName() string { return "prefix-list-sequence" }

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) SetParent(parent types.Entity) { prefixListSequence.parent = parent }

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetParent() types.Entity { return prefixListSequence.parent }

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetParentYangName() string { return "prefix-list-sequences" }

// Ipv6AclAndPrefixList_AccessListManager_Usages
// Table of Usage statistics of ACLs at different
// nodes
type Ipv6AclAndPrefixList_AccessListManager_Usages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Usage statistics of an ACL at a node. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Usages_Usage.
    Usage []Ipv6AclAndPrefixList_AccessListManager_Usages_Usage
}

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetFilter() yfilter.YFilter { return usages.YFilter }

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) SetFilter(yf yfilter.YFilter) { usages.YFilter = yf }

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetGoName(yname string) string {
    if yname == "usage" { return "Usage" }
    return ""
}

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetSegmentPath() string {
    return "usages"
}

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usage" {
        for _, c := range usages.Usage {
            if usages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_AccessListManager_Usages_Usage{}
        usages.Usage = append(usages.Usage, child)
        return &usages.Usage[len(usages.Usage)-1]
    }
    return nil
}

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usages.Usage {
        children[usages.Usage[i].GetSegmentPath()] = &usages.Usage[i]
    }
    return children
}

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetBundleName() string { return "cisco_ios_xr" }

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetYangName() string { return "usages" }

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) SetParent(parent types.Entity) { usages.parent = parent }

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetParent() types.Entity { return usages.parent }

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetParentYangName() string { return "access-list-manager" }

// Ipv6AclAndPrefixList_AccessListManager_Usages_Usage
// Usage statistics of an ACL at a node
type Ipv6AclAndPrefixList_AccessListManager_Usages_Usage struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node where ACL is applied. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Application ID. The type is AclUsageAppIdEnum.
    ApplicationId interface{}

    // Name of the ACL. The type is string with length: 1..65.
    AccessListName interface{}

    // Usage Statistics Details. The type is string. This attribute is mandatory.
    UsageDetails interface{}
}

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetFilter() yfilter.YFilter { return usage.YFilter }

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) SetFilter(yf yfilter.YFilter) { usage.YFilter = yf }

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "application-id" { return "ApplicationId" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "usage-details" { return "UsageDetails" }
    return ""
}

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetSegmentPath() string {
    return "usage"
}

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = usage.NodeName
    leafs["application-id"] = usage.ApplicationId
    leafs["access-list-name"] = usage.AccessListName
    leafs["usage-details"] = usage.UsageDetails
    return leafs
}

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetBundleName() string { return "cisco_ios_xr" }

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetYangName() string { return "usage" }

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) SetParent(parent types.Entity) { usage.parent = parent }

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetParent() types.Entity { return usage.parent }

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetParentYangName() string { return "usages" }

// Ipv6AclAndPrefixList_AccessListManager_Accesses
// ACL class displaying Usage and Entries
type Ipv6AclAndPrefixList_AccessListManager_Accesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the Access List. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Accesses_Access.
    Access []Ipv6AclAndPrefixList_AccessListManager_Accesses_Access
}

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetFilter() yfilter.YFilter { return accesses.YFilter }

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) SetFilter(yf yfilter.YFilter) { accesses.YFilter = yf }

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetGoName(yname string) string {
    if yname == "access" { return "Access" }
    return ""
}

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetSegmentPath() string {
    return "accesses"
}

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access" {
        for _, c := range accesses.Access {
            if accesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_AccessListManager_Accesses_Access{}
        accesses.Access = append(accesses.Access, child)
        return &accesses.Access[len(accesses.Access)-1]
    }
    return nil
}

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accesses.Access {
        children[accesses.Access[i].GetSegmentPath()] = &accesses.Access[i]
    }
    return children
}

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetBundleName() string { return "cisco_ios_xr" }

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetYangName() string { return "accesses" }

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) SetParent(parent types.Entity) { accesses.parent = parent }

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetParent() types.Entity { return accesses.parent }

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetParentYangName() string { return "access-list-manager" }

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access
// Name of the Access List
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Access List. The type is string with
    // length: 1..65.
    AccessListName interface{}

    // Table of all the sequence numbers per ACL.
    AccessListSequences Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences
}

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetFilter() yfilter.YFilter { return access.YFilter }

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) SetFilter(yf yfilter.YFilter) { access.YFilter = yf }

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetGoName(yname string) string {
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "access-list-sequences" { return "AccessListSequences" }
    return ""
}

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetSegmentPath() string {
    return "access" + "[access-list-name='" + fmt.Sprintf("%v", access.AccessListName) + "']"
}

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-sequences" {
        return &access.AccessListSequences
    }
    return nil
}

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-list-sequences"] = &access.AccessListSequences
    return children
}

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-list-name"] = access.AccessListName
    return leafs
}

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetBundleName() string { return "cisco_ios_xr" }

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetYangName() string { return "access" }

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) SetParent(parent types.Entity) { access.parent = parent }

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetParent() types.Entity { return access.parent }

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetParentYangName() string { return "accesses" }

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences
// Table of all the sequence numbers per ACL
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sequence number of an ACL entry. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence.
    AccessListSequence []Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence
}

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetFilter() yfilter.YFilter { return accessListSequences.YFilter }

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) SetFilter(yf yfilter.YFilter) { accessListSequences.YFilter = yf }

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetGoName(yname string) string {
    if yname == "access-list-sequence" { return "AccessListSequence" }
    return ""
}

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetSegmentPath() string {
    return "access-list-sequences"
}

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-list-sequence" {
        for _, c := range accessListSequences.AccessListSequence {
            if accessListSequences.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence{}
        accessListSequences.AccessListSequence = append(accessListSequences.AccessListSequence, child)
        return &accessListSequences.AccessListSequence[len(accessListSequences.AccessListSequence)-1]
    }
    return nil
}

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessListSequences.AccessListSequence {
        children[accessListSequences.AccessListSequence[i].GetSegmentPath()] = &accessListSequences.AccessListSequence[i]
    }
    return children
}

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetBundleName() string { return "cisco_ios_xr" }

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetYangName() string { return "access-list-sequences" }

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) SetParent(parent types.Entity) { accessListSequences.parent = parent }

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetParent() types.Entity { return accessListSequences.parent }

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetParentYangName() string { return "access" }

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence
// Sequence number of an ACL entry
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ACL entry sequence number. The type is interface{}
    // with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (acl, remark). The type is AclAce1.
    IsAceType interface{}

    // ACLE sequence number. The type is interface{} with range: 0..4294967295.
    IsAceSequenceNumber interface{}

    // Grant value permit/deny . The type is AclAction.
    IsPacketAllowOrDeny interface{}

    // Protocol operator. The type is AclPortOperator.
    IsProtocolOperator interface{}

    // Protocol 1. The type is interface{} with range: -2147483648..2147483647.
    IsIpv6ProtocolType interface{}

    // Protocol 2. The type is interface{} with range: -2147483648..2147483647.
    IsIpv6Protocol2Type interface{}

    // IsSourceAddressInNumbers. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IsSourceAddressInNumbers interface{}

    // IsSourceAddressPrefixLength. The type is interface{} with range:
    // 0..4294967295.
    IsSourceAddressPrefixLength interface{}

    // Source Mask. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceMask interface{}

    // IsDestinationAddressInNumbers. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IsDestinationAddressInNumbers interface{}

    // IsDestinationAddressPrefixLength. The type is interface{} with range:
    // 0..4294967295.
    IsDestinationAddressPrefixLength interface{}

    // Destination Mask. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationMask interface{}

    // eq, ne, lt, etc... The type is AclPortOperator.
    IsSourceOperator interface{}

    // IsSourcePort1. The type is interface{} with range: 0..4294967295.
    IsSourcePort1 interface{}

    // IsSourcePort2. The type is interface{} with range: 0..4294967295.
    IsSourcePort2 interface{}

    // eq, ne, lt, etc... The type is AclPortOperator.
    IsDestinationOperator interface{}

    // IsDestinationPort1. The type is interface{} with range: 0..4294967295.
    IsDestinationPort1 interface{}

    // IsDestinationPort2. The type is interface{} with range: 0..4294967295.
    IsDestinationPort2 interface{}

    // IsLogOption. The type is AclLog.
    IsLogOption interface{}

    // Counter name. The type is string.
    CounterName interface{}

    // IsTCPBitsOperator. The type is AclTcpflagsOperator.
    IsTcpBitsOperator interface{}

    // IsTCPBits. The type is interface{} with range: 0..4294967295.
    IsTcpBits interface{}

    // IsTCPBitsMask. The type is interface{} with range: 0..4294967295.
    IsTcpBitsMask interface{}

    // IsDSCPPresent. The type is interface{} with range: -2147483648..2147483647.
    IsDscpPresent interface{}

    // IsDSCPValu. The type is interface{} with range: 0..4294967295.
    IsDscpValu interface{}

    // IsPrecedencePresent. The type is interface{} with range:
    // -2147483648..2147483647.
    IsPrecedencePresent interface{}

    // range from 0 to 7. The type is interface{} with range: 0..4294967295.
    IsPrecedenceValue interface{}

    // Match if routing header is presant. The type is interface{} with range:
    // 0..4294967295.
    IsHeaderMatches interface{}

    // Match if routing header is presant. The type is AclPortOperator.
    IsPacketLengthOperator interface{}

    // IsPacketLengthStart. The type is interface{} with range: 0..4294967295.
    IsPacketLengthStart interface{}

    // IsPacketLengthEnd. The type is interface{} with range: 0..4294967295.
    IsPacketLengthEnd interface{}

    // IsTimeToLiveOperator. The type is AclPortOperator.
    IsTimeToLiveOperator interface{}

    // IsTimeToLiveStart. The type is interface{} with range: 0..4294967295.
    IsTimeToLiveStart interface{}

    // IsTimeToLiveEnd. The type is interface{} with range: 0..4294967295.
    IsTimeToLiveEnd interface{}

    // no stats. The type is interface{} with range: -2147483648..2147483647.
    NoStats interface{}

    // hits. The type is interface{} with range: 0..18446744073709551615.
    Hits interface{}

    // Capture option, TRUE if enabled. The type is bool.
    Capture interface{}

    // Undetermined transport option, TRUE if enabled. The type is bool.
    UndeterminedTransport interface{}

    // Don't generate the icmp message. The type is interface{} with range:
    // -2147483648..2147483647.
    IsIcmpMessageOff interface{}

    // Set qos-group. The type is interface{} with range: 0..65535.
    QosGroup interface{}

    // IsCommentForEntry. The type is string.
    IsCommentForEntry interface{}

    // Next hop type. The type is BagAclNh.
    NextHopType interface{}

    // IsFlowId. The type is interface{} with range: 0..4294967295.
    IsFlowId interface{}

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

    // SetTTL. The type is interface{} with range: 0..65535.
    SetTtl interface{}

    // HW Next hop info.
    HwNextHopInfo Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo

    // Next hop info. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo.
    NextHopInfo []Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo

    // UDF. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf.
    Udf []Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf
}

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetFilter() yfilter.YFilter { return accessListSequence.YFilter }

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) SetFilter(yf yfilter.YFilter) { accessListSequence.YFilter = yf }

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "is-ace-type" { return "IsAceType" }
    if yname == "is-ace-sequence-number" { return "IsAceSequenceNumber" }
    if yname == "is-packet-allow-or-deny" { return "IsPacketAllowOrDeny" }
    if yname == "is-protocol-operator" { return "IsProtocolOperator" }
    if yname == "is-ipv6-protocol-type" { return "IsIpv6ProtocolType" }
    if yname == "is-ipv6-protocol2-type" { return "IsIpv6Protocol2Type" }
    if yname == "is-source-address-in-numbers" { return "IsSourceAddressInNumbers" }
    if yname == "is-source-address-prefix-length" { return "IsSourceAddressPrefixLength" }
    if yname == "source-mask" { return "SourceMask" }
    if yname == "is-destination-address-in-numbers" { return "IsDestinationAddressInNumbers" }
    if yname == "is-destination-address-prefix-length" { return "IsDestinationAddressPrefixLength" }
    if yname == "destination-mask" { return "DestinationMask" }
    if yname == "is-source-operator" { return "IsSourceOperator" }
    if yname == "is-source-port1" { return "IsSourcePort1" }
    if yname == "is-source-port2" { return "IsSourcePort2" }
    if yname == "is-destination-operator" { return "IsDestinationOperator" }
    if yname == "is-destination-port1" { return "IsDestinationPort1" }
    if yname == "is-destination-port2" { return "IsDestinationPort2" }
    if yname == "is-log-option" { return "IsLogOption" }
    if yname == "counter-name" { return "CounterName" }
    if yname == "is-tcp-bits-operator" { return "IsTcpBitsOperator" }
    if yname == "is-tcp-bits" { return "IsTcpBits" }
    if yname == "is-tcp-bits-mask" { return "IsTcpBitsMask" }
    if yname == "is-dscp-present" { return "IsDscpPresent" }
    if yname == "is-dscp-valu" { return "IsDscpValu" }
    if yname == "is-precedence-present" { return "IsPrecedencePresent" }
    if yname == "is-precedence-value" { return "IsPrecedenceValue" }
    if yname == "is-header-matches" { return "IsHeaderMatches" }
    if yname == "is-packet-length-operator" { return "IsPacketLengthOperator" }
    if yname == "is-packet-length-start" { return "IsPacketLengthStart" }
    if yname == "is-packet-length-end" { return "IsPacketLengthEnd" }
    if yname == "is-time-to-live-operator" { return "IsTimeToLiveOperator" }
    if yname == "is-time-to-live-start" { return "IsTimeToLiveStart" }
    if yname == "is-time-to-live-end" { return "IsTimeToLiveEnd" }
    if yname == "no-stats" { return "NoStats" }
    if yname == "hits" { return "Hits" }
    if yname == "capture" { return "Capture" }
    if yname == "undetermined-transport" { return "UndeterminedTransport" }
    if yname == "is-icmp-message-off" { return "IsIcmpMessageOff" }
    if yname == "qos-group" { return "QosGroup" }
    if yname == "is-comment-for-entry" { return "IsCommentForEntry" }
    if yname == "next-hop-type" { return "NextHopType" }
    if yname == "is-flow-id" { return "IsFlowId" }
    if yname == "source-prefix-group" { return "SourcePrefixGroup" }
    if yname == "destination-prefix-group" { return "DestinationPrefixGroup" }
    if yname == "source-port-group" { return "SourcePortGroup" }
    if yname == "destination-port-group" { return "DestinationPortGroup" }
    if yname == "acl-name" { return "AclName" }
    if yname == "sequence-str" { return "SequenceStr" }
    if yname == "set-ttl" { return "SetTtl" }
    if yname == "hw-next-hop-info" { return "HwNextHopInfo" }
    if yname == "next-hop-info" { return "NextHopInfo" }
    if yname == "udf" { return "Udf" }
    return ""
}

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetSegmentPath() string {
    return "access-list-sequence" + "[sequence-number='" + fmt.Sprintf("%v", accessListSequence.SequenceNumber) + "']"
}

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hw-next-hop-info" {
        return &accessListSequence.HwNextHopInfo
    }
    if childYangName == "next-hop-info" {
        for _, c := range accessListSequence.NextHopInfo {
            if accessListSequence.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo{}
        accessListSequence.NextHopInfo = append(accessListSequence.NextHopInfo, child)
        return &accessListSequence.NextHopInfo[len(accessListSequence.NextHopInfo)-1]
    }
    if childYangName == "udf" {
        for _, c := range accessListSequence.Udf {
            if accessListSequence.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf{}
        accessListSequence.Udf = append(accessListSequence.Udf, child)
        return &accessListSequence.Udf[len(accessListSequence.Udf)-1]
    }
    return nil
}

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetChildren() map[string]types.Entity {
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

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = accessListSequence.SequenceNumber
    leafs["is-ace-type"] = accessListSequence.IsAceType
    leafs["is-ace-sequence-number"] = accessListSequence.IsAceSequenceNumber
    leafs["is-packet-allow-or-deny"] = accessListSequence.IsPacketAllowOrDeny
    leafs["is-protocol-operator"] = accessListSequence.IsProtocolOperator
    leafs["is-ipv6-protocol-type"] = accessListSequence.IsIpv6ProtocolType
    leafs["is-ipv6-protocol2-type"] = accessListSequence.IsIpv6Protocol2Type
    leafs["is-source-address-in-numbers"] = accessListSequence.IsSourceAddressInNumbers
    leafs["is-source-address-prefix-length"] = accessListSequence.IsSourceAddressPrefixLength
    leafs["source-mask"] = accessListSequence.SourceMask
    leafs["is-destination-address-in-numbers"] = accessListSequence.IsDestinationAddressInNumbers
    leafs["is-destination-address-prefix-length"] = accessListSequence.IsDestinationAddressPrefixLength
    leafs["destination-mask"] = accessListSequence.DestinationMask
    leafs["is-source-operator"] = accessListSequence.IsSourceOperator
    leafs["is-source-port1"] = accessListSequence.IsSourcePort1
    leafs["is-source-port2"] = accessListSequence.IsSourcePort2
    leafs["is-destination-operator"] = accessListSequence.IsDestinationOperator
    leafs["is-destination-port1"] = accessListSequence.IsDestinationPort1
    leafs["is-destination-port2"] = accessListSequence.IsDestinationPort2
    leafs["is-log-option"] = accessListSequence.IsLogOption
    leafs["counter-name"] = accessListSequence.CounterName
    leafs["is-tcp-bits-operator"] = accessListSequence.IsTcpBitsOperator
    leafs["is-tcp-bits"] = accessListSequence.IsTcpBits
    leafs["is-tcp-bits-mask"] = accessListSequence.IsTcpBitsMask
    leafs["is-dscp-present"] = accessListSequence.IsDscpPresent
    leafs["is-dscp-valu"] = accessListSequence.IsDscpValu
    leafs["is-precedence-present"] = accessListSequence.IsPrecedencePresent
    leafs["is-precedence-value"] = accessListSequence.IsPrecedenceValue
    leafs["is-header-matches"] = accessListSequence.IsHeaderMatches
    leafs["is-packet-length-operator"] = accessListSequence.IsPacketLengthOperator
    leafs["is-packet-length-start"] = accessListSequence.IsPacketLengthStart
    leafs["is-packet-length-end"] = accessListSequence.IsPacketLengthEnd
    leafs["is-time-to-live-operator"] = accessListSequence.IsTimeToLiveOperator
    leafs["is-time-to-live-start"] = accessListSequence.IsTimeToLiveStart
    leafs["is-time-to-live-end"] = accessListSequence.IsTimeToLiveEnd
    leafs["no-stats"] = accessListSequence.NoStats
    leafs["hits"] = accessListSequence.Hits
    leafs["capture"] = accessListSequence.Capture
    leafs["undetermined-transport"] = accessListSequence.UndeterminedTransport
    leafs["is-icmp-message-off"] = accessListSequence.IsIcmpMessageOff
    leafs["qos-group"] = accessListSequence.QosGroup
    leafs["is-comment-for-entry"] = accessListSequence.IsCommentForEntry
    leafs["next-hop-type"] = accessListSequence.NextHopType
    leafs["is-flow-id"] = accessListSequence.IsFlowId
    leafs["source-prefix-group"] = accessListSequence.SourcePrefixGroup
    leafs["destination-prefix-group"] = accessListSequence.DestinationPrefixGroup
    leafs["source-port-group"] = accessListSequence.SourcePortGroup
    leafs["destination-port-group"] = accessListSequence.DestinationPortGroup
    leafs["acl-name"] = accessListSequence.AclName
    leafs["sequence-str"] = accessListSequence.SequenceStr
    leafs["set-ttl"] = accessListSequence.SetTtl
    return leafs
}

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetBundleName() string { return "cisco_ios_xr" }

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetYangName() string { return "access-list-sequence" }

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) SetParent(parent types.Entity) { accessListSequence.parent = parent }

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetParent() types.Entity { return accessListSequence.parent }

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetParentYangName() string { return "access-list-sequences" }

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo
// HW Next hop info
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The next-hop type. The type is BagAclNh.
    Type interface{}

    // The Next Hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Vrf Name. The type is string with length: 0..32.
    VrfName interface{}
}

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetFilter() yfilter.YFilter { return hwNextHopInfo.YFilter }

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) SetFilter(yf yfilter.YFilter) { hwNextHopInfo.YFilter = yf }

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "next-hop" { return "NextHop" }
    if yname == "table-id" { return "TableId" }
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetSegmentPath() string {
    return "hw-next-hop-info"
}

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = hwNextHopInfo.Type
    leafs["next-hop"] = hwNextHopInfo.NextHop
    leafs["table-id"] = hwNextHopInfo.TableId
    leafs["vrf-name"] = hwNextHopInfo.VrfName
    return leafs
}

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetBundleName() string { return "cisco_ios_xr" }

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetYangName() string { return "hw-next-hop-info" }

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) SetParent(parent types.Entity) { hwNextHopInfo.parent = parent }

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetParent() types.Entity { return hwNextHopInfo.parent }

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetParentYangName() string { return "access-list-sequence" }

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo
// Next hop info
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The next hop. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // Vrf Name. The type is string with length: 0..32.
    VrfName interface{}

    // Track name. The type is string with length: 0..33.
    TrackName interface{}

    // The next hop status. The type is BagAclNhStatus.
    Status interface{}

    // The next hop at status. The type is BagAclNhAtStatus.
    AtStatus interface{}

    // The nexthop exist. The type is interface{} with range:
    // -2147483648..2147483647.
    AclNhExist interface{}
}

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetFilter() yfilter.YFilter { return nextHopInfo.YFilter }

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) SetFilter(yf yfilter.YFilter) { nextHopInfo.YFilter = yf }

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "track-name" { return "TrackName" }
    if yname == "status" { return "Status" }
    if yname == "at-status" { return "AtStatus" }
    if yname == "acl-nh-exist" { return "AclNhExist" }
    return ""
}

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetSegmentPath() string {
    return "next-hop-info"
}

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = nextHopInfo.NextHop
    leafs["vrf-name"] = nextHopInfo.VrfName
    leafs["track-name"] = nextHopInfo.TrackName
    leafs["status"] = nextHopInfo.Status
    leafs["at-status"] = nextHopInfo.AtStatus
    leafs["acl-nh-exist"] = nextHopInfo.AclNhExist
    return leafs
}

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetBundleName() string { return "cisco_ios_xr" }

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetYangName() string { return "next-hop-info" }

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) SetParent(parent types.Entity) { nextHopInfo.parent = parent }

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetParent() types.Entity { return nextHopInfo.parent }

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetParentYangName() string { return "access-list-sequence" }

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf
// UDF
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // UDF Name. The type is string with length: 0..17.
    UdfName interface{}

    // UDF Value. The type is interface{} with range: 0..4294967295.
    UdfValue interface{}

    // UDF Mask. The type is interface{} with range: 0..4294967295.
    UdfMask interface{}
}

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetFilter() yfilter.YFilter { return udf.YFilter }

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) SetFilter(yf yfilter.YFilter) { udf.YFilter = yf }

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetGoName(yname string) string {
    if yname == "udf-name" { return "UdfName" }
    if yname == "udf-value" { return "UdfValue" }
    if yname == "udf-mask" { return "UdfMask" }
    return ""
}

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetSegmentPath() string {
    return "udf"
}

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udf-name"] = udf.UdfName
    leafs["udf-value"] = udf.UdfValue
    leafs["udf-mask"] = udf.UdfMask
    return leafs
}

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetBundleName() string { return "cisco_ios_xr" }

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetYangName() string { return "udf" }

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) SetParent(parent types.Entity) { udf.parent = parent }

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetParent() types.Entity { return udf.parent }

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetParentYangName() string { return "access-list-sequence" }

// Ipv6AclAndPrefixList_Oor
// Out Of Resources, Limits to the resources
// allocatable
type Ipv6AclAndPrefixList_Oor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details of the overall out of resource limit.
    Details Ipv6AclAndPrefixList_Oor_Details

    // Summary of the prefix Lists resource utilization.
    PrefixListSummary Ipv6AclAndPrefixList_Oor_PrefixListSummary

    // Resource occupation details for ACLs.
    OorAccesses Ipv6AclAndPrefixList_Oor_OorAccesses

    // Resource occupation details for prefix lists.
    OorPrefixes Ipv6AclAndPrefixList_Oor_OorPrefixes

    // Resource Limits pertaining to ACLs only.
    AccessListSummary Ipv6AclAndPrefixList_Oor_AccessListSummary
}

func (oor *Ipv6AclAndPrefixList_Oor) GetFilter() yfilter.YFilter { return oor.YFilter }

func (oor *Ipv6AclAndPrefixList_Oor) SetFilter(yf yfilter.YFilter) { oor.YFilter = yf }

func (oor *Ipv6AclAndPrefixList_Oor) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    if yname == "prefix-list-summary" { return "PrefixListSummary" }
    if yname == "oor-accesses" { return "OorAccesses" }
    if yname == "oor-prefixes" { return "OorPrefixes" }
    if yname == "access-list-summary" { return "AccessListSummary" }
    return ""
}

func (oor *Ipv6AclAndPrefixList_Oor) GetSegmentPath() string {
    return "oor"
}

func (oor *Ipv6AclAndPrefixList_Oor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &oor.Details
    }
    if childYangName == "prefix-list-summary" {
        return &oor.PrefixListSummary
    }
    if childYangName == "oor-accesses" {
        return &oor.OorAccesses
    }
    if childYangName == "oor-prefixes" {
        return &oor.OorPrefixes
    }
    if childYangName == "access-list-summary" {
        return &oor.AccessListSummary
    }
    return nil
}

func (oor *Ipv6AclAndPrefixList_Oor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &oor.Details
    children["prefix-list-summary"] = &oor.PrefixListSummary
    children["oor-accesses"] = &oor.OorAccesses
    children["oor-prefixes"] = &oor.OorPrefixes
    children["access-list-summary"] = &oor.AccessListSummary
    return children
}

func (oor *Ipv6AclAndPrefixList_Oor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oor *Ipv6AclAndPrefixList_Oor) GetBundleName() string { return "cisco_ios_xr" }

func (oor *Ipv6AclAndPrefixList_Oor) GetYangName() string { return "oor" }

func (oor *Ipv6AclAndPrefixList_Oor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oor *Ipv6AclAndPrefixList_Oor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oor *Ipv6AclAndPrefixList_Oor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oor *Ipv6AclAndPrefixList_Oor) SetParent(parent types.Entity) { oor.parent = parent }

func (oor *Ipv6AclAndPrefixList_Oor) GetParent() types.Entity { return oor.parent }

func (oor *Ipv6AclAndPrefixList_Oor) GetParentYangName() string { return "ipv6-acl-and-prefix-list" }

// Ipv6AclAndPrefixList_Oor_Details
// Details of the overall out of resource limit
type Ipv6AclAndPrefixList_Oor_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAces interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAcls interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAces interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcEs interface{}
}

func (details *Ipv6AclAndPrefixList_Oor_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Ipv6AclAndPrefixList_Oor_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Ipv6AclAndPrefixList_Oor_Details) GetGoName(yname string) string {
    if yname == "is-default-maximum-configurable-ac-ls" { return "IsDefaultMaximumConfigurableAcLs" }
    if yname == "is-default-maximum-configurable-ac-es" { return "IsDefaultMaximumConfigurableAcEs" }
    if yname == "is-current-configured-ac-ls" { return "IsCurrentConfiguredAcLs" }
    if yname == "is-current-configured-aces" { return "IsCurrentConfiguredAces" }
    if yname == "is-current-maximum-configurable-acls" { return "IsCurrentMaximumConfigurableAcls" }
    if yname == "is-current-maximum-configurable-aces" { return "IsCurrentMaximumConfigurableAces" }
    if yname == "is-maximum-configurable-ac-ls" { return "IsMaximumConfigurableAcLs" }
    if yname == "is-maximum-configurable-ac-es" { return "IsMaximumConfigurableAcEs" }
    return ""
}

func (details *Ipv6AclAndPrefixList_Oor_Details) GetSegmentPath() string {
    return "details"
}

func (details *Ipv6AclAndPrefixList_Oor_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (details *Ipv6AclAndPrefixList_Oor_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (details *Ipv6AclAndPrefixList_Oor_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-default-maximum-configurable-ac-ls"] = details.IsDefaultMaximumConfigurableAcLs
    leafs["is-default-maximum-configurable-ac-es"] = details.IsDefaultMaximumConfigurableAcEs
    leafs["is-current-configured-ac-ls"] = details.IsCurrentConfiguredAcLs
    leafs["is-current-configured-aces"] = details.IsCurrentConfiguredAces
    leafs["is-current-maximum-configurable-acls"] = details.IsCurrentMaximumConfigurableAcls
    leafs["is-current-maximum-configurable-aces"] = details.IsCurrentMaximumConfigurableAces
    leafs["is-maximum-configurable-ac-ls"] = details.IsMaximumConfigurableAcLs
    leafs["is-maximum-configurable-ac-es"] = details.IsMaximumConfigurableAcEs
    return leafs
}

func (details *Ipv6AclAndPrefixList_Oor_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Ipv6AclAndPrefixList_Oor_Details) GetYangName() string { return "details" }

func (details *Ipv6AclAndPrefixList_Oor_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Ipv6AclAndPrefixList_Oor_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Ipv6AclAndPrefixList_Oor_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Ipv6AclAndPrefixList_Oor_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Ipv6AclAndPrefixList_Oor_Details) GetParent() types.Entity { return details.parent }

func (details *Ipv6AclAndPrefixList_Oor_Details) GetParentYangName() string { return "oor" }

// Ipv6AclAndPrefixList_Oor_PrefixListSummary
// Summary of the prefix Lists resource
// utilization
type Ipv6AclAndPrefixList_Oor_PrefixListSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary Detail of the prefix list Resource Utilisation.
    Details Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details
}

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetFilter() yfilter.YFilter { return prefixListSummary.YFilter }

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) SetFilter(yf yfilter.YFilter) { prefixListSummary.YFilter = yf }

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    return ""
}

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetSegmentPath() string {
    return "prefix-list-summary"
}

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &prefixListSummary.Details
    }
    return nil
}

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &prefixListSummary.Details
    return children
}

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetBundleName() string { return "cisco_ios_xr" }

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetYangName() string { return "prefix-list-summary" }

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) SetParent(parent types.Entity) { prefixListSummary.parent = parent }

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetParent() types.Entity { return prefixListSummary.parent }

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetParentYangName() string { return "oor" }

// Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details
// Summary Detail of the prefix list Resource
// Utilisation
type Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAces interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAcls interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAces interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcEs interface{}
}

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetGoName(yname string) string {
    if yname == "is-default-maximum-configurable-ac-ls" { return "IsDefaultMaximumConfigurableAcLs" }
    if yname == "is-default-maximum-configurable-ac-es" { return "IsDefaultMaximumConfigurableAcEs" }
    if yname == "is-current-configured-ac-ls" { return "IsCurrentConfiguredAcLs" }
    if yname == "is-current-configured-aces" { return "IsCurrentConfiguredAces" }
    if yname == "is-current-maximum-configurable-acls" { return "IsCurrentMaximumConfigurableAcls" }
    if yname == "is-current-maximum-configurable-aces" { return "IsCurrentMaximumConfigurableAces" }
    if yname == "is-maximum-configurable-ac-ls" { return "IsMaximumConfigurableAcLs" }
    if yname == "is-maximum-configurable-ac-es" { return "IsMaximumConfigurableAcEs" }
    return ""
}

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetSegmentPath() string {
    return "details"
}

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-default-maximum-configurable-ac-ls"] = details.IsDefaultMaximumConfigurableAcLs
    leafs["is-default-maximum-configurable-ac-es"] = details.IsDefaultMaximumConfigurableAcEs
    leafs["is-current-configured-ac-ls"] = details.IsCurrentConfiguredAcLs
    leafs["is-current-configured-aces"] = details.IsCurrentConfiguredAces
    leafs["is-current-maximum-configurable-acls"] = details.IsCurrentMaximumConfigurableAcls
    leafs["is-current-maximum-configurable-aces"] = details.IsCurrentMaximumConfigurableAces
    leafs["is-maximum-configurable-ac-ls"] = details.IsMaximumConfigurableAcLs
    leafs["is-maximum-configurable-ac-es"] = details.IsMaximumConfigurableAcEs
    return leafs
}

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetYangName() string { return "details" }

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetParent() types.Entity { return details.parent }

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetParentYangName() string { return "prefix-list-summary" }

// Ipv6AclAndPrefixList_Oor_OorAccesses
// Resource occupation details for ACLs
type Ipv6AclAndPrefixList_Oor_OorAccesses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource occupation details for a particular ACL. The type is slice of
    // Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess.
    OorAccess []Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess
}

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetFilter() yfilter.YFilter { return oorAccesses.YFilter }

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) SetFilter(yf yfilter.YFilter) { oorAccesses.YFilter = yf }

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetGoName(yname string) string {
    if yname == "oor-access" { return "OorAccess" }
    return ""
}

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetSegmentPath() string {
    return "oor-accesses"
}

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oor-access" {
        for _, c := range oorAccesses.OorAccess {
            if oorAccesses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess{}
        oorAccesses.OorAccess = append(oorAccesses.OorAccess, child)
        return &oorAccesses.OorAccess[len(oorAccesses.OorAccess)-1]
    }
    return nil
}

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range oorAccesses.OorAccess {
        children[oorAccesses.OorAccess[i].GetSegmentPath()] = &oorAccesses.OorAccess[i]
    }
    return children
}

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetBundleName() string { return "cisco_ios_xr" }

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetYangName() string { return "oor-accesses" }

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) SetParent(parent types.Entity) { oorAccesses.parent = parent }

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetParent() types.Entity { return oorAccesses.parent }

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetParentYangName() string { return "oor" }

// Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess
// Resource occupation details for a particular
// ACL
type Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Access List. The type is string with
    // length: 1..65.
    AccessListName interface{}

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAces interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAcls interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAces interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcEs interface{}
}

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetFilter() yfilter.YFilter { return oorAccess.YFilter }

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) SetFilter(yf yfilter.YFilter) { oorAccess.YFilter = yf }

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetGoName(yname string) string {
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "is-default-maximum-configurable-ac-ls" { return "IsDefaultMaximumConfigurableAcLs" }
    if yname == "is-default-maximum-configurable-ac-es" { return "IsDefaultMaximumConfigurableAcEs" }
    if yname == "is-current-configured-ac-ls" { return "IsCurrentConfiguredAcLs" }
    if yname == "is-current-configured-aces" { return "IsCurrentConfiguredAces" }
    if yname == "is-current-maximum-configurable-acls" { return "IsCurrentMaximumConfigurableAcls" }
    if yname == "is-current-maximum-configurable-aces" { return "IsCurrentMaximumConfigurableAces" }
    if yname == "is-maximum-configurable-ac-ls" { return "IsMaximumConfigurableAcLs" }
    if yname == "is-maximum-configurable-ac-es" { return "IsMaximumConfigurableAcEs" }
    return ""
}

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetSegmentPath() string {
    return "oor-access" + "[access-list-name='" + fmt.Sprintf("%v", oorAccess.AccessListName) + "']"
}

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["access-list-name"] = oorAccess.AccessListName
    leafs["is-default-maximum-configurable-ac-ls"] = oorAccess.IsDefaultMaximumConfigurableAcLs
    leafs["is-default-maximum-configurable-ac-es"] = oorAccess.IsDefaultMaximumConfigurableAcEs
    leafs["is-current-configured-ac-ls"] = oorAccess.IsCurrentConfiguredAcLs
    leafs["is-current-configured-aces"] = oorAccess.IsCurrentConfiguredAces
    leafs["is-current-maximum-configurable-acls"] = oorAccess.IsCurrentMaximumConfigurableAcls
    leafs["is-current-maximum-configurable-aces"] = oorAccess.IsCurrentMaximumConfigurableAces
    leafs["is-maximum-configurable-ac-ls"] = oorAccess.IsMaximumConfigurableAcLs
    leafs["is-maximum-configurable-ac-es"] = oorAccess.IsMaximumConfigurableAcEs
    return leafs
}

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetBundleName() string { return "cisco_ios_xr" }

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetYangName() string { return "oor-access" }

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) SetParent(parent types.Entity) { oorAccess.parent = parent }

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetParent() types.Entity { return oorAccess.parent }

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetParentYangName() string { return "oor-accesses" }

// Ipv6AclAndPrefixList_Oor_OorPrefixes
// Resource occupation details for prefix lists
type Ipv6AclAndPrefixList_Oor_OorPrefixes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Resource occupation details for a particular prefix list. The type is slice
    // of Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix.
    OorPrefix []Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix
}

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetFilter() yfilter.YFilter { return oorPrefixes.YFilter }

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) SetFilter(yf yfilter.YFilter) { oorPrefixes.YFilter = yf }

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetGoName(yname string) string {
    if yname == "oor-prefix" { return "OorPrefix" }
    return ""
}

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetSegmentPath() string {
    return "oor-prefixes"
}

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oor-prefix" {
        for _, c := range oorPrefixes.OorPrefix {
            if oorPrefixes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix{}
        oorPrefixes.OorPrefix = append(oorPrefixes.OorPrefix, child)
        return &oorPrefixes.OorPrefix[len(oorPrefixes.OorPrefix)-1]
    }
    return nil
}

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range oorPrefixes.OorPrefix {
        children[oorPrefixes.OorPrefix[i].GetSegmentPath()] = &oorPrefixes.OorPrefix[i]
    }
    return children
}

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetBundleName() string { return "cisco_ios_xr" }

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetYangName() string { return "oor-prefixes" }

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) SetParent(parent types.Entity) { oorPrefixes.parent = parent }

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetParent() types.Entity { return oorPrefixes.parent }

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetParentYangName() string { return "oor" }

// Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix
// Resource occupation details for a particular
// prefix list
type Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of a prefix list. The type is string with
    // length: 1..65.
    PrefixListName interface{}

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAces interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAcls interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAces interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcEs interface{}
}

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetFilter() yfilter.YFilter { return oorPrefix.YFilter }

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) SetFilter(yf yfilter.YFilter) { oorPrefix.YFilter = yf }

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetGoName(yname string) string {
    if yname == "prefix-list-name" { return "PrefixListName" }
    if yname == "is-default-maximum-configurable-ac-ls" { return "IsDefaultMaximumConfigurableAcLs" }
    if yname == "is-default-maximum-configurable-ac-es" { return "IsDefaultMaximumConfigurableAcEs" }
    if yname == "is-current-configured-ac-ls" { return "IsCurrentConfiguredAcLs" }
    if yname == "is-current-configured-aces" { return "IsCurrentConfiguredAces" }
    if yname == "is-current-maximum-configurable-acls" { return "IsCurrentMaximumConfigurableAcls" }
    if yname == "is-current-maximum-configurable-aces" { return "IsCurrentMaximumConfigurableAces" }
    if yname == "is-maximum-configurable-ac-ls" { return "IsMaximumConfigurableAcLs" }
    if yname == "is-maximum-configurable-ac-es" { return "IsMaximumConfigurableAcEs" }
    return ""
}

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetSegmentPath() string {
    return "oor-prefix" + "[prefix-list-name='" + fmt.Sprintf("%v", oorPrefix.PrefixListName) + "']"
}

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-list-name"] = oorPrefix.PrefixListName
    leafs["is-default-maximum-configurable-ac-ls"] = oorPrefix.IsDefaultMaximumConfigurableAcLs
    leafs["is-default-maximum-configurable-ac-es"] = oorPrefix.IsDefaultMaximumConfigurableAcEs
    leafs["is-current-configured-ac-ls"] = oorPrefix.IsCurrentConfiguredAcLs
    leafs["is-current-configured-aces"] = oorPrefix.IsCurrentConfiguredAces
    leafs["is-current-maximum-configurable-acls"] = oorPrefix.IsCurrentMaximumConfigurableAcls
    leafs["is-current-maximum-configurable-aces"] = oorPrefix.IsCurrentMaximumConfigurableAces
    leafs["is-maximum-configurable-ac-ls"] = oorPrefix.IsMaximumConfigurableAcLs
    leafs["is-maximum-configurable-ac-es"] = oorPrefix.IsMaximumConfigurableAcEs
    return leafs
}

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetYangName() string { return "oor-prefix" }

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) SetParent(parent types.Entity) { oorPrefix.parent = parent }

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetParent() types.Entity { return oorPrefix.parent }

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetParentYangName() string { return "oor-prefixes" }

// Ipv6AclAndPrefixList_Oor_AccessListSummary
// Resource Limits pertaining to ACLs only
type Ipv6AclAndPrefixList_Oor_AccessListSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details containing the resource limits of the ACLs.
    Details Ipv6AclAndPrefixList_Oor_AccessListSummary_Details
}

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetFilter() yfilter.YFilter { return accessListSummary.YFilter }

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) SetFilter(yf yfilter.YFilter) { accessListSummary.YFilter = yf }

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    return ""
}

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetSegmentPath() string {
    return "access-list-summary"
}

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &accessListSummary.Details
    }
    return nil
}

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &accessListSummary.Details
    return children
}

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetBundleName() string { return "cisco_ios_xr" }

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetYangName() string { return "access-list-summary" }

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) SetParent(parent types.Entity) { accessListSummary.parent = parent }

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetParent() types.Entity { return accessListSummary.parent }

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetParentYangName() string { return "oor" }

// Ipv6AclAndPrefixList_Oor_AccessListSummary_Details
// Details containing the resource limits of the
// ACLs
type Ipv6AclAndPrefixList_Oor_AccessListSummary_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // default max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcLs interface{}

    // default max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsDefaultMaximumConfigurableAcEs interface{}

    // Current configured acls. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAcLs interface{}

    // Current configured aces. The type is interface{} with range: 0..4294967295.
    IsCurrentConfiguredAces interface{}

    // Current max configurable acls. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAcls interface{}

    // Current max configurable aces. The type is interface{} with range:
    // 0..4294967295.
    IsCurrentMaximumConfigurableAces interface{}

    // max configurable acls. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcLs interface{}

    // max configurable aces. The type is interface{} with range: 0..4294967295.
    IsMaximumConfigurableAcEs interface{}
}

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetGoName(yname string) string {
    if yname == "is-default-maximum-configurable-ac-ls" { return "IsDefaultMaximumConfigurableAcLs" }
    if yname == "is-default-maximum-configurable-ac-es" { return "IsDefaultMaximumConfigurableAcEs" }
    if yname == "is-current-configured-ac-ls" { return "IsCurrentConfiguredAcLs" }
    if yname == "is-current-configured-aces" { return "IsCurrentConfiguredAces" }
    if yname == "is-current-maximum-configurable-acls" { return "IsCurrentMaximumConfigurableAcls" }
    if yname == "is-current-maximum-configurable-aces" { return "IsCurrentMaximumConfigurableAces" }
    if yname == "is-maximum-configurable-ac-ls" { return "IsMaximumConfigurableAcLs" }
    if yname == "is-maximum-configurable-ac-es" { return "IsMaximumConfigurableAcEs" }
    return ""
}

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetSegmentPath() string {
    return "details"
}

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-default-maximum-configurable-ac-ls"] = details.IsDefaultMaximumConfigurableAcLs
    leafs["is-default-maximum-configurable-ac-es"] = details.IsDefaultMaximumConfigurableAcEs
    leafs["is-current-configured-ac-ls"] = details.IsCurrentConfiguredAcLs
    leafs["is-current-configured-aces"] = details.IsCurrentConfiguredAces
    leafs["is-current-maximum-configurable-acls"] = details.IsCurrentMaximumConfigurableAcls
    leafs["is-current-maximum-configurable-aces"] = details.IsCurrentMaximumConfigurableAces
    leafs["is-maximum-configurable-ac-ls"] = details.IsMaximumConfigurableAcLs
    leafs["is-maximum-configurable-ac-es"] = details.IsMaximumConfigurableAcEs
    return leafs
}

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetYangName() string { return "details" }

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetParent() types.Entity { return details.parent }

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetParentYangName() string { return "access-list-summary" }

