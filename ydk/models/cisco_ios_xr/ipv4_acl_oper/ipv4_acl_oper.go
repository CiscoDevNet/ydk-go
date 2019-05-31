// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-acl package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-acl-and-prefix-list: Root class of IPv4 Oper schema tree
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// AclPortOperator_ represents Acl port operator
type AclPortOperator_ string

const (
    // None
    AclPortOperator__none AclPortOperator_ = "none"

    // Equal
    AclPortOperator__eq AclPortOperator_ = "eq"

    // Greater than
    AclPortOperator__gt AclPortOperator_ = "gt"

    // Less than
    AclPortOperator__lt AclPortOperator_ = "lt"

    // Not Equal
    AclPortOperator__neq AclPortOperator_ = "neq"

    // Range
    AclPortOperator__range_ AclPortOperator_ = "range"

    // One Byte
    AclPortOperator__onebyte AclPortOperator_ = "onebyte"

    // Two Bytes
    AclPortOperator__twobytes AclPortOperator_ = "twobytes"
)

// AclPortOperator_ represents Acl port operator
type AclPortOperator_ string

const (
    // None
    AclPortOperator__none AclPortOperator_ = "none"

    // Equal
    AclPortOperator__eq AclPortOperator_ = "eq"

    // Greater than
    AclPortOperator__gt AclPortOperator_ = "gt"

    // Less than
    AclPortOperator__lt AclPortOperator_ = "lt"

    // Not Equal
    AclPortOperator__neq AclPortOperator_ = "neq"

    // Range
    AclPortOperator__range_ AclPortOperator_ = "range"

    // One Byte
    AclPortOperator__onebyte AclPortOperator_ = "onebyte"

    // Two Bytes
    AclPortOperator__twobytes AclPortOperator_ = "twobytes"
)

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

// AclPortOperator_ represents Acl port operator
type AclPortOperator_ string

const (
    // None
    AclPortOperator__none AclPortOperator_ = "none"

    // Equal
    AclPortOperator__eq AclPortOperator_ = "eq"

    // Greater than
    AclPortOperator__gt AclPortOperator_ = "gt"

    // Less than
    AclPortOperator__lt AclPortOperator_ = "lt"

    // Not Equal
    AclPortOperator__neq AclPortOperator_ = "neq"

    // Range
    AclPortOperator__range_ AclPortOperator_ = "range"

    // One Byte
    AclPortOperator__onebyte AclPortOperator_ = "onebyte"

    // Two Bytes
    AclPortOperator__twobytes AclPortOperator_ = "twobytes"
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

// AclAce1_ represents ACE Types
type AclAce1_ string

const (
    // This is Normal ACE
    AclAce1__normal AclAce1_ = "normal"

    // This is Remark ACE
    AclAce1__remark AclAce1_ = "remark"

    // This is ABF ACE
    AclAce1__abf AclAce1_ = "abf"
)

// Ipv4AclAndPrefixList
// Root class of IPv4 Oper schema tree
type Ipv4AclAndPrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access list manager containing access lists and prefix lists.
    AccessListManager Ipv4AclAndPrefixList_AccessListManager

    // Out Of Resources, Limits to the resources allocatable.
    Oor Ipv4AclAndPrefixList_Oor
}

func (ipv4AclAndPrefixList *Ipv4AclAndPrefixList) GetEntityData() *types.CommonEntityData {
    ipv4AclAndPrefixList.EntityData.YFilter = ipv4AclAndPrefixList.YFilter
    ipv4AclAndPrefixList.EntityData.YangName = "ipv4-acl-and-prefix-list"
    ipv4AclAndPrefixList.EntityData.BundleName = "cisco_ios_xr"
    ipv4AclAndPrefixList.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-acl-oper"
    ipv4AclAndPrefixList.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list"
    ipv4AclAndPrefixList.EntityData.AbsolutePath = ipv4AclAndPrefixList.EntityData.SegmentPath
    ipv4AclAndPrefixList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AclAndPrefixList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AclAndPrefixList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AclAndPrefixList.EntityData.Children = types.NewOrderedMap()
    ipv4AclAndPrefixList.EntityData.Children.Append("access-list-manager", types.YChild{"AccessListManager", &ipv4AclAndPrefixList.AccessListManager})
    ipv4AclAndPrefixList.EntityData.Children.Append("oor", types.YChild{"Oor", &ipv4AclAndPrefixList.Oor})
    ipv4AclAndPrefixList.EntityData.Leafs = types.NewOrderedMap()

    ipv4AclAndPrefixList.EntityData.YListKeys = []string {}

    return &(ipv4AclAndPrefixList.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager
// Access list manager containing access lists and
// prefix lists
type Ipv4AclAndPrefixList_AccessListManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of prefix lists.
    Prefixes Ipv4AclAndPrefixList_AccessListManager_Prefixes

    // Access listL class displaying Usage and Entries.
    Accesses Ipv4AclAndPrefixList_AccessListManager_Accesses

    // Table of Usage statistics of access lists at different nodes.
    Usages Ipv4AclAndPrefixList_AccessListManager_Usages
}

func (accessListManager *Ipv4AclAndPrefixList_AccessListManager) GetEntityData() *types.CommonEntityData {
    accessListManager.EntityData.YFilter = accessListManager.YFilter
    accessListManager.EntityData.YangName = "access-list-manager"
    accessListManager.EntityData.BundleName = "cisco_ios_xr"
    accessListManager.EntityData.ParentYangName = "ipv4-acl-and-prefix-list"
    accessListManager.EntityData.SegmentPath = "access-list-manager"
    accessListManager.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/" + accessListManager.EntityData.SegmentPath
    accessListManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListManager.EntityData.Children = types.NewOrderedMap()
    accessListManager.EntityData.Children.Append("prefixes", types.YChild{"Prefixes", &accessListManager.Prefixes})
    accessListManager.EntityData.Children.Append("accesses", types.YChild{"Accesses", &accessListManager.Accesses})
    accessListManager.EntityData.Children.Append("usages", types.YChild{"Usages", &accessListManager.Usages})
    accessListManager.EntityData.Leafs = types.NewOrderedMap()

    accessListManager.EntityData.YListKeys = []string {}

    return &(accessListManager.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Prefixes
// Table of prefix lists
type Ipv4AclAndPrefixList_AccessListManager_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the prefix list. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix.
    Prefix []*Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix
}

func (prefixes *Ipv4AclAndPrefixList_AccessListManager_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "cisco_ios_xr"
    prefixes.EntityData.ParentYangName = "access-list-manager"
    prefixes.EntityData.SegmentPath = "prefixes"
    prefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/" + prefixes.EntityData.SegmentPath
    prefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixes.EntityData.Children = types.NewOrderedMap()
    prefixes.EntityData.Children.Append("prefix", types.YChild{"Prefix", nil})
    for i := range prefixes.Prefix {
        prefixes.EntityData.Children.Append(types.GetSegmentPath(prefixes.Prefix[i]), types.YChild{"Prefix", prefixes.Prefix[i]})
    }
    prefixes.EntityData.Leafs = types.NewOrderedMap()

    prefixes.EntityData.YListKeys = []string {}

    return &(prefixes.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix
// Name of the prefix list
type Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the prefix list. The type is string.
    PrefixListName interface{}

    // Table of all the SequenceNumbers per prefix list.
    PrefixListSequences Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences
}

func (prefix *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefixes"
    prefix.EntityData.SegmentPath = "prefix" + types.AddKeyToken(prefix.PrefixListName, "prefix-list-name")
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/prefixes/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Children.Append("prefix-list-sequences", types.YChild{"PrefixListSequences", &prefix.PrefixListSequences})
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("prefix-list-name", types.YLeaf{"PrefixListName", prefix.PrefixListName})

    prefix.EntityData.YListKeys = []string {"PrefixListName"}

    return &(prefix.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences
// Table of all the SequenceNumbers per prefix
// list
type Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sequence Number of a prefix list entry. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence.
    PrefixListSequence []*Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence
}

func (prefixListSequences *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetEntityData() *types.CommonEntityData {
    prefixListSequences.EntityData.YFilter = prefixListSequences.YFilter
    prefixListSequences.EntityData.YangName = "prefix-list-sequences"
    prefixListSequences.EntityData.BundleName = "cisco_ios_xr"
    prefixListSequences.EntityData.ParentYangName = "prefix"
    prefixListSequences.EntityData.SegmentPath = "prefix-list-sequences"
    prefixListSequences.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/prefixes/prefix/" + prefixListSequences.EntityData.SegmentPath
    prefixListSequences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixListSequences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixListSequences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixListSequences.EntityData.Children = types.NewOrderedMap()
    prefixListSequences.EntityData.Children.Append("prefix-list-sequence", types.YChild{"PrefixListSequence", nil})
    for i := range prefixListSequences.PrefixListSequence {
        prefixListSequences.EntityData.Children.Append(types.GetSegmentPath(prefixListSequences.PrefixListSequence[i]), types.YChild{"PrefixListSequence", prefixListSequences.PrefixListSequence[i]})
    }
    prefixListSequences.EntityData.Leafs = types.NewOrderedMap()

    prefixListSequences.EntityData.YListKeys = []string {}

    return &(prefixListSequences.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence
// Sequence Number of a prefix list entry
type Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Sequence Number of the prefix list entry. The type
    // is interface{} with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (prefix, remark). The type is AclAce1_.
    ItemType interface{}

    // ACLE sequence number. The type is interface{} with range: 0..4294967295.
    Sequence interface{}

    // Grant value permit/deny . The type is AclAction.
    Grant interface{}

    // Prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length . The type is interface{} with range: 0..4294967295.
    PrefixLength interface{}

    // Port Operator. The type is AclPortOperator_.
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

func (prefixListSequence *Ipv4AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetEntityData() *types.CommonEntityData {
    prefixListSequence.EntityData.YFilter = prefixListSequence.YFilter
    prefixListSequence.EntityData.YangName = "prefix-list-sequence"
    prefixListSequence.EntityData.BundleName = "cisco_ios_xr"
    prefixListSequence.EntityData.ParentYangName = "prefix-list-sequences"
    prefixListSequence.EntityData.SegmentPath = "prefix-list-sequence" + types.AddKeyToken(prefixListSequence.SequenceNumber, "sequence-number")
    prefixListSequence.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/prefixes/prefix/prefix-list-sequences/" + prefixListSequence.EntityData.SegmentPath
    prefixListSequence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixListSequence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixListSequence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixListSequence.EntityData.Children = types.NewOrderedMap()
    prefixListSequence.EntityData.Leafs = types.NewOrderedMap()
    prefixListSequence.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", prefixListSequence.SequenceNumber})
    prefixListSequence.EntityData.Leafs.Append("item-type", types.YLeaf{"ItemType", prefixListSequence.ItemType})
    prefixListSequence.EntityData.Leafs.Append("sequence", types.YLeaf{"Sequence", prefixListSequence.Sequence})
    prefixListSequence.EntityData.Leafs.Append("grant", types.YLeaf{"Grant", prefixListSequence.Grant})
    prefixListSequence.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefixListSequence.Prefix})
    prefixListSequence.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefixListSequence.PrefixLength})
    prefixListSequence.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", prefixListSequence.Operator})
    prefixListSequence.EntityData.Leafs.Append("minimum-length", types.YLeaf{"MinimumLength", prefixListSequence.MinimumLength})
    prefixListSequence.EntityData.Leafs.Append("maximum-length", types.YLeaf{"MaximumLength", prefixListSequence.MaximumLength})
    prefixListSequence.EntityData.Leafs.Append("hits", types.YLeaf{"Hits", prefixListSequence.Hits})
    prefixListSequence.EntityData.Leafs.Append("remark", types.YLeaf{"Remark", prefixListSequence.Remark})
    prefixListSequence.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", prefixListSequence.AclName})

    prefixListSequence.EntityData.YListKeys = []string {"SequenceNumber"}

    return &(prefixListSequence.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Accesses
// Access listL class displaying Usage and Entries
type Ipv4AclAndPrefixList_AccessListManager_Accesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Access List. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Accesses_Access.
    Access []*Ipv4AclAndPrefixList_AccessListManager_Accesses_Access
}

func (accesses *Ipv4AclAndPrefixList_AccessListManager_Accesses) GetEntityData() *types.CommonEntityData {
    accesses.EntityData.YFilter = accesses.YFilter
    accesses.EntityData.YangName = "accesses"
    accesses.EntityData.BundleName = "cisco_ios_xr"
    accesses.EntityData.ParentYangName = "access-list-manager"
    accesses.EntityData.SegmentPath = "accesses"
    accesses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/" + accesses.EntityData.SegmentPath
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

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access
// Name of the Access List
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the Access List. The type is string.
    AccessListName interface{}

    // Table of all the SequenceNumbers per access list.
    AccessListSequences Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences

    // Object Group in an Access list.
    ObjectGroup Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup
}

func (access *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access) GetEntityData() *types.CommonEntityData {
    access.EntityData.YFilter = access.YFilter
    access.EntityData.YangName = "access"
    access.EntityData.BundleName = "cisco_ios_xr"
    access.EntityData.ParentYangName = "accesses"
    access.EntityData.SegmentPath = "access" + types.AddKeyToken(access.AccessListName, "access-list-name")
    access.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/accesses/" + access.EntityData.SegmentPath
    access.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    access.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    access.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    access.EntityData.Children = types.NewOrderedMap()
    access.EntityData.Children.Append("access-list-sequences", types.YChild{"AccessListSequences", &access.AccessListSequences})
    access.EntityData.Children.Append("object-group", types.YChild{"ObjectGroup", &access.ObjectGroup})
    access.EntityData.Leafs = types.NewOrderedMap()
    access.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", access.AccessListName})

    access.EntityData.YListKeys = []string {"AccessListName"}

    return &(access.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences
// Table of all the SequenceNumbers per access
// list
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sequence Number of an access list entry. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence.
    AccessListSequence []*Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence
}

func (accessListSequences *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetEntityData() *types.CommonEntityData {
    accessListSequences.EntityData.YFilter = accessListSequences.YFilter
    accessListSequences.EntityData.YangName = "access-list-sequences"
    accessListSequences.EntityData.BundleName = "cisco_ios_xr"
    accessListSequences.EntityData.ParentYangName = "access"
    accessListSequences.EntityData.SegmentPath = "access-list-sequences"
    accessListSequences.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/accesses/access/" + accessListSequences.EntityData.SegmentPath
    accessListSequences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListSequences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListSequences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListSequences.EntityData.Children = types.NewOrderedMap()
    accessListSequences.EntityData.Children.Append("access-list-sequence", types.YChild{"AccessListSequence", nil})
    for i := range accessListSequences.AccessListSequence {
        accessListSequences.EntityData.Children.Append(types.GetSegmentPath(accessListSequences.AccessListSequence[i]), types.YChild{"AccessListSequence", accessListSequences.AccessListSequence[i]})
    }
    accessListSequences.EntityData.Leafs = types.NewOrderedMap()

    accessListSequences.EntityData.YListKeys = []string {}

    return &(accessListSequences.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence
// Sequence Number of an access list entry
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. ACLEntry Sequence Number. The type is interface{}
    // with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (acl, remark). The type is AclAce1_.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Source mask. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddressMask interface{}

    // Destination address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Destination mask. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddressMask interface{}

    // Source operator. The type is AclPortOperator_.
    SourceOperator interface{}

    // Source port 1. The type is interface{} with range: 0..65535.
    SourcePort1 interface{}

    // Source port 2. The type is interface{} with range: 0..65535.
    SourcePort2 interface{}

    // Deprecated by Source operator. The type is AclPortOperator_.
    SorceOperator interface{}

    // Deprecated by SourcePort1. The type is interface{} with range: 0..65535.
    SorcePort1 interface{}

    // Deprecated by SourcePort2. The type is interface{} with range: 0..65535.
    SorcePort2 interface{}

    // Destination operator. The type is AclPortOperator_.
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

    // Packet length operator. The type is AclPortOperator_.
    PacketLengthOperator interface{}

    // Packet length 1. The type is interface{} with range: 0..65535.
    PacketLength1 interface{}

    // Packet length 2. The type is interface{} with range: 0..65535.
    PacketLength2 interface{}

    // TTL operator. The type is AclPortOperator_.
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

    // Fragment offset operator. The type is AclPortOperator_.
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
    NextHopInfo []*Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo

    // UDF BAG. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf.
    Udf []*Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf
}

func (accessListSequence *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetEntityData() *types.CommonEntityData {
    accessListSequence.EntityData.YFilter = accessListSequence.YFilter
    accessListSequence.EntityData.YangName = "access-list-sequence"
    accessListSequence.EntityData.BundleName = "cisco_ios_xr"
    accessListSequence.EntityData.ParentYangName = "access-list-sequences"
    accessListSequence.EntityData.SegmentPath = "access-list-sequence" + types.AddKeyToken(accessListSequence.SequenceNumber, "sequence-number")
    accessListSequence.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/accesses/access/access-list-sequences/" + accessListSequence.EntityData.SegmentPath
    accessListSequence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListSequence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListSequence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListSequence.EntityData.Children = types.NewOrderedMap()
    accessListSequence.EntityData.Children.Append("hw-next-hop-info", types.YChild{"HwNextHopInfo", &accessListSequence.HwNextHopInfo})
    accessListSequence.EntityData.Children.Append("next-hop-info", types.YChild{"NextHopInfo", nil})
    for i := range accessListSequence.NextHopInfo {
        types.SetYListKey(accessListSequence.NextHopInfo[i], i)
        accessListSequence.EntityData.Children.Append(types.GetSegmentPath(accessListSequence.NextHopInfo[i]), types.YChild{"NextHopInfo", accessListSequence.NextHopInfo[i]})
    }
    accessListSequence.EntityData.Children.Append("udf", types.YChild{"Udf", nil})
    for i := range accessListSequence.Udf {
        types.SetYListKey(accessListSequence.Udf[i], i)
        accessListSequence.EntityData.Children.Append(types.GetSegmentPath(accessListSequence.Udf[i]), types.YChild{"Udf", accessListSequence.Udf[i]})
    }
    accessListSequence.EntityData.Leafs = types.NewOrderedMap()
    accessListSequence.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", accessListSequence.SequenceNumber})
    accessListSequence.EntityData.Leafs.Append("item-type", types.YLeaf{"ItemType", accessListSequence.ItemType})
    accessListSequence.EntityData.Leafs.Append("sequence", types.YLeaf{"Sequence", accessListSequence.Sequence})
    accessListSequence.EntityData.Leafs.Append("grant", types.YLeaf{"Grant", accessListSequence.Grant})
    accessListSequence.EntityData.Leafs.Append("protocol-operator", types.YLeaf{"ProtocolOperator", accessListSequence.ProtocolOperator})
    accessListSequence.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", accessListSequence.Protocol})
    accessListSequence.EntityData.Leafs.Append("protocol2", types.YLeaf{"Protocol2", accessListSequence.Protocol2})
    accessListSequence.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", accessListSequence.SourceAddress})
    accessListSequence.EntityData.Leafs.Append("source-address-mask", types.YLeaf{"SourceAddressMask", accessListSequence.SourceAddressMask})
    accessListSequence.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", accessListSequence.DestinationAddress})
    accessListSequence.EntityData.Leafs.Append("destination-address-mask", types.YLeaf{"DestinationAddressMask", accessListSequence.DestinationAddressMask})
    accessListSequence.EntityData.Leafs.Append("source-operator", types.YLeaf{"SourceOperator", accessListSequence.SourceOperator})
    accessListSequence.EntityData.Leafs.Append("source-port1", types.YLeaf{"SourcePort1", accessListSequence.SourcePort1})
    accessListSequence.EntityData.Leafs.Append("source-port2", types.YLeaf{"SourcePort2", accessListSequence.SourcePort2})
    accessListSequence.EntityData.Leafs.Append("sorce-operator", types.YLeaf{"SorceOperator", accessListSequence.SorceOperator})
    accessListSequence.EntityData.Leafs.Append("sorce-port1", types.YLeaf{"SorcePort1", accessListSequence.SorcePort1})
    accessListSequence.EntityData.Leafs.Append("sorce-port2", types.YLeaf{"SorcePort2", accessListSequence.SorcePort2})
    accessListSequence.EntityData.Leafs.Append("destination-operator", types.YLeaf{"DestinationOperator", accessListSequence.DestinationOperator})
    accessListSequence.EntityData.Leafs.Append("destination-port1", types.YLeaf{"DestinationPort1", accessListSequence.DestinationPort1})
    accessListSequence.EntityData.Leafs.Append("destination-port2", types.YLeaf{"DestinationPort2", accessListSequence.DestinationPort2})
    accessListSequence.EntityData.Leafs.Append("log-option", types.YLeaf{"LogOption", accessListSequence.LogOption})
    accessListSequence.EntityData.Leafs.Append("counter-name", types.YLeaf{"CounterName", accessListSequence.CounterName})
    accessListSequence.EntityData.Leafs.Append("capture", types.YLeaf{"Capture", accessListSequence.Capture})
    accessListSequence.EntityData.Leafs.Append("dscp-present", types.YLeaf{"DscpPresent", accessListSequence.DscpPresent})
    accessListSequence.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", accessListSequence.Dscp})
    accessListSequence.EntityData.Leafs.Append("dscp2", types.YLeaf{"Dscp2", accessListSequence.Dscp2})
    accessListSequence.EntityData.Leafs.Append("dscp-operator", types.YLeaf{"DscpOperator", accessListSequence.DscpOperator})
    accessListSequence.EntityData.Leafs.Append("precedence-present", types.YLeaf{"PrecedencePresent", accessListSequence.PrecedencePresent})
    accessListSequence.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", accessListSequence.Precedence})
    accessListSequence.EntityData.Leafs.Append("tcp-flags-operator", types.YLeaf{"TcpFlagsOperator", accessListSequence.TcpFlagsOperator})
    accessListSequence.EntityData.Leafs.Append("tcp-flags", types.YLeaf{"TcpFlags", accessListSequence.TcpFlags})
    accessListSequence.EntityData.Leafs.Append("tcp-flags-mask", types.YLeaf{"TcpFlagsMask", accessListSequence.TcpFlagsMask})
    accessListSequence.EntityData.Leafs.Append("fragments", types.YLeaf{"Fragments", accessListSequence.Fragments})
    accessListSequence.EntityData.Leafs.Append("packet-length-operator", types.YLeaf{"PacketLengthOperator", accessListSequence.PacketLengthOperator})
    accessListSequence.EntityData.Leafs.Append("packet-length1", types.YLeaf{"PacketLength1", accessListSequence.PacketLength1})
    accessListSequence.EntityData.Leafs.Append("packet-length2", types.YLeaf{"PacketLength2", accessListSequence.PacketLength2})
    accessListSequence.EntityData.Leafs.Append("ttl-operator", types.YLeaf{"TtlOperator", accessListSequence.TtlOperator})
    accessListSequence.EntityData.Leafs.Append("ttl1", types.YLeaf{"Ttl1", accessListSequence.Ttl1})
    accessListSequence.EntityData.Leafs.Append("ttl2", types.YLeaf{"Ttl2", accessListSequence.Ttl2})
    accessListSequence.EntityData.Leafs.Append("no-stats", types.YLeaf{"NoStats", accessListSequence.NoStats})
    accessListSequence.EntityData.Leafs.Append("hits", types.YLeaf{"Hits", accessListSequence.Hits})
    accessListSequence.EntityData.Leafs.Append("is-icmp-off", types.YLeaf{"IsIcmpOff", accessListSequence.IsIcmpOff})
    accessListSequence.EntityData.Leafs.Append("qos-group", types.YLeaf{"QosGroup", accessListSequence.QosGroup})
    accessListSequence.EntityData.Leafs.Append("next-hop-type", types.YLeaf{"NextHopType", accessListSequence.NextHopType})
    accessListSequence.EntityData.Leafs.Append("remark", types.YLeaf{"Remark", accessListSequence.Remark})
    accessListSequence.EntityData.Leafs.Append("dynamic", types.YLeaf{"Dynamic", accessListSequence.Dynamic})
    accessListSequence.EntityData.Leafs.Append("source-prefix-group", types.YLeaf{"SourcePrefixGroup", accessListSequence.SourcePrefixGroup})
    accessListSequence.EntityData.Leafs.Append("destination-prefix-group", types.YLeaf{"DestinationPrefixGroup", accessListSequence.DestinationPrefixGroup})
    accessListSequence.EntityData.Leafs.Append("source-port-group", types.YLeaf{"SourcePortGroup", accessListSequence.SourcePortGroup})
    accessListSequence.EntityData.Leafs.Append("destination-port-group", types.YLeaf{"DestinationPortGroup", accessListSequence.DestinationPortGroup})
    accessListSequence.EntityData.Leafs.Append("acl-name", types.YLeaf{"AclName", accessListSequence.AclName})
    accessListSequence.EntityData.Leafs.Append("sequence-str", types.YLeaf{"SequenceStr", accessListSequence.SequenceStr})
    accessListSequence.EntityData.Leafs.Append("fragment-offset-operator", types.YLeaf{"FragmentOffsetOperator", accessListSequence.FragmentOffsetOperator})
    accessListSequence.EntityData.Leafs.Append("fragment-offset1", types.YLeaf{"FragmentOffset1", accessListSequence.FragmentOffset1})
    accessListSequence.EntityData.Leafs.Append("fragment-offset2", types.YLeaf{"FragmentOffset2", accessListSequence.FragmentOffset2})
    accessListSequence.EntityData.Leafs.Append("set-ttl", types.YLeaf{"SetTtl", accessListSequence.SetTtl})
    accessListSequence.EntityData.Leafs.Append("fragment-flags", types.YLeaf{"FragmentFlags", accessListSequence.FragmentFlags})

    accessListSequence.EntityData.YListKeys = []string {"SequenceNumber"}

    return &(accessListSequence.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo
// HW Next hop info
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Next Hop. The type is interface{} with range: 0..4294967295.
    NextHop interface{}

    // the next-hop type. The type is BagAclNh.
    Type interface{}

    // VRF name. The type is string with length: 0..32.
    VrfName interface{}
}

func (hwNextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetEntityData() *types.CommonEntityData {
    hwNextHopInfo.EntityData.YFilter = hwNextHopInfo.YFilter
    hwNextHopInfo.EntityData.YangName = "hw-next-hop-info"
    hwNextHopInfo.EntityData.BundleName = "cisco_ios_xr"
    hwNextHopInfo.EntityData.ParentYangName = "access-list-sequence"
    hwNextHopInfo.EntityData.SegmentPath = "hw-next-hop-info"
    hwNextHopInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/accesses/access/access-list-sequences/access-list-sequence/" + hwNextHopInfo.EntityData.SegmentPath
    hwNextHopInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwNextHopInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwNextHopInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwNextHopInfo.EntityData.Children = types.NewOrderedMap()
    hwNextHopInfo.EntityData.Leafs = types.NewOrderedMap()
    hwNextHopInfo.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", hwNextHopInfo.NextHop})
    hwNextHopInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", hwNextHopInfo.Type})
    hwNextHopInfo.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", hwNextHopInfo.VrfName})

    hwNextHopInfo.EntityData.YListKeys = []string {}

    return &(hwNextHopInfo.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo
// Next hop info
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The next hop. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (nextHopInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetEntityData() *types.CommonEntityData {
    nextHopInfo.EntityData.YFilter = nextHopInfo.YFilter
    nextHopInfo.EntityData.YangName = "next-hop-info"
    nextHopInfo.EntityData.BundleName = "cisco_ios_xr"
    nextHopInfo.EntityData.ParentYangName = "access-list-sequence"
    nextHopInfo.EntityData.SegmentPath = "next-hop-info" + types.AddNoKeyToken(nextHopInfo)
    nextHopInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/accesses/access/access-list-sequences/access-list-sequence/" + nextHopInfo.EntityData.SegmentPath
    nextHopInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHopInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHopInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHopInfo.EntityData.Children = types.NewOrderedMap()
    nextHopInfo.EntityData.Leafs = types.NewOrderedMap()
    nextHopInfo.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", nextHopInfo.NextHop})
    nextHopInfo.EntityData.Leafs.Append("track-name", types.YLeaf{"TrackName", nextHopInfo.TrackName})
    nextHopInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", nextHopInfo.Status})
    nextHopInfo.EntityData.Leafs.Append("at-status", types.YLeaf{"AtStatus", nextHopInfo.AtStatus})
    nextHopInfo.EntityData.Leafs.Append("is-acl-next-hop-exist", types.YLeaf{"IsAclNextHopExist", nextHopInfo.IsAclNextHopExist})

    nextHopInfo.EntityData.YListKeys = []string {}

    return &(nextHopInfo.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf
// UDF BAG
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // UDF Name. The type is string with length: 0..17.
    UdfName interface{}

    // UDF Value. The type is interface{} with range: 0..4294967295.
    UdfValue interface{}

    // UDF Mask. The type is interface{} with range: 0..4294967295.
    UdfMask interface{}
}

func (udf *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetEntityData() *types.CommonEntityData {
    udf.EntityData.YFilter = udf.YFilter
    udf.EntityData.YangName = "udf"
    udf.EntityData.BundleName = "cisco_ios_xr"
    udf.EntityData.ParentYangName = "access-list-sequence"
    udf.EntityData.SegmentPath = "udf" + types.AddNoKeyToken(udf)
    udf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/accesses/access/access-list-sequences/access-list-sequence/" + udf.EntityData.SegmentPath
    udf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udf.EntityData.Children = types.NewOrderedMap()
    udf.EntityData.Leafs = types.NewOrderedMap()
    udf.EntityData.Leafs.Append("udf-name", types.YLeaf{"UdfName", udf.UdfName})
    udf.EntityData.Leafs.Append("udf-value", types.YLeaf{"UdfValue", udf.UdfValue})
    udf.EntityData.Leafs.Append("udf-mask", types.YLeaf{"UdfMask", udf.UdfMask})

    udf.EntityData.YListKeys = []string {}

    return &(udf.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup
// Object Group in an Access list
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object-group info. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo.
    ObjGrpInfo []*Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo
}

func (objectGroup *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup) GetEntityData() *types.CommonEntityData {
    objectGroup.EntityData.YFilter = objectGroup.YFilter
    objectGroup.EntityData.YangName = "object-group"
    objectGroup.EntityData.BundleName = "cisco_ios_xr"
    objectGroup.EntityData.ParentYangName = "access"
    objectGroup.EntityData.SegmentPath = "object-group"
    objectGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/accesses/access/" + objectGroup.EntityData.SegmentPath
    objectGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectGroup.EntityData.Children = types.NewOrderedMap()
    objectGroup.EntityData.Children.Append("obj-grp-info", types.YChild{"ObjGrpInfo", nil})
    for i := range objectGroup.ObjGrpInfo {
        types.SetYListKey(objectGroup.ObjGrpInfo[i], i)
        objectGroup.EntityData.Children.Append(types.GetSegmentPath(objectGroup.ObjGrpInfo[i]), types.YChild{"ObjGrpInfo", objectGroup.ObjGrpInfo[i]})
    }
    objectGroup.EntityData.Leafs = types.NewOrderedMap()

    objectGroup.EntityData.YListKeys = []string {}

    return &(objectGroup.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo
// Object-group info
type Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Object-group name. The type is string with length: 0..64.
    ObjGrpName interface{}

    // Object-group Type. The type is interface{} with range: 0..4294967295.
    ObjGrpType interface{}
}

func (objGrpInfo *Ipv4AclAndPrefixList_AccessListManager_Accesses_Access_ObjectGroup_ObjGrpInfo) GetEntityData() *types.CommonEntityData {
    objGrpInfo.EntityData.YFilter = objGrpInfo.YFilter
    objGrpInfo.EntityData.YangName = "obj-grp-info"
    objGrpInfo.EntityData.BundleName = "cisco_ios_xr"
    objGrpInfo.EntityData.ParentYangName = "object-group"
    objGrpInfo.EntityData.SegmentPath = "obj-grp-info" + types.AddNoKeyToken(objGrpInfo)
    objGrpInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/accesses/access/object-group/" + objGrpInfo.EntityData.SegmentPath
    objGrpInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objGrpInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objGrpInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objGrpInfo.EntityData.Children = types.NewOrderedMap()
    objGrpInfo.EntityData.Leafs = types.NewOrderedMap()
    objGrpInfo.EntityData.Leafs.Append("obj-grp-name", types.YLeaf{"ObjGrpName", objGrpInfo.ObjGrpName})
    objGrpInfo.EntityData.Leafs.Append("obj-grp-type", types.YLeaf{"ObjGrpType", objGrpInfo.ObjGrpType})

    objGrpInfo.EntityData.YListKeys = []string {}

    return &(objGrpInfo.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Usages
// Table of Usage statistics of access lists at
// different nodes
type Ipv4AclAndPrefixList_AccessListManager_Usages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Usage statistics of an access list at a node. The type is slice of
    // Ipv4AclAndPrefixList_AccessListManager_Usages_Usage.
    Usage []*Ipv4AclAndPrefixList_AccessListManager_Usages_Usage
}

func (usages *Ipv4AclAndPrefixList_AccessListManager_Usages) GetEntityData() *types.CommonEntityData {
    usages.EntityData.YFilter = usages.YFilter
    usages.EntityData.YangName = "usages"
    usages.EntityData.BundleName = "cisco_ios_xr"
    usages.EntityData.ParentYangName = "access-list-manager"
    usages.EntityData.SegmentPath = "usages"
    usages.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/" + usages.EntityData.SegmentPath
    usages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usages.EntityData.Children = types.NewOrderedMap()
    usages.EntityData.Children.Append("usage", types.YChild{"Usage", nil})
    for i := range usages.Usage {
        types.SetYListKey(usages.Usage[i], i)
        usages.EntityData.Children.Append(types.GetSegmentPath(usages.Usage[i]), types.YChild{"Usage", usages.Usage[i]})
    }
    usages.EntityData.Leafs = types.NewOrderedMap()

    usages.EntityData.YListKeys = []string {}

    return &(usages.EntityData)
}

// Ipv4AclAndPrefixList_AccessListManager_Usages_Usage
// Usage statistics of an access list at a node
type Ipv4AclAndPrefixList_AccessListManager_Usages_Usage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Node where access list is applied. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Application ID. The type is AclUsageAppIdEnum.
    ApplicationId interface{}

    // Name of the access list. The type is string.
    AccessListName interface{}

    // Usage Statistics Details. The type is string. This attribute is mandatory.
    UsageDetails interface{}
}

func (usage *Ipv4AclAndPrefixList_AccessListManager_Usages_Usage) GetEntityData() *types.CommonEntityData {
    usage.EntityData.YFilter = usage.YFilter
    usage.EntityData.YangName = "usage"
    usage.EntityData.BundleName = "cisco_ios_xr"
    usage.EntityData.ParentYangName = "usages"
    usage.EntityData.SegmentPath = "usage" + types.AddNoKeyToken(usage)
    usage.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/access-list-manager/usages/" + usage.EntityData.SegmentPath
    usage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usage.EntityData.Children = types.NewOrderedMap()
    usage.EntityData.Leafs = types.NewOrderedMap()
    usage.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", usage.NodeName})
    usage.EntityData.Leafs.Append("application-id", types.YLeaf{"ApplicationId", usage.ApplicationId})
    usage.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", usage.AccessListName})
    usage.EntityData.Leafs.Append("usage-details", types.YLeaf{"UsageDetails", usage.UsageDetails})

    usage.EntityData.YListKeys = []string {}

    return &(usage.EntityData)
}

// Ipv4AclAndPrefixList_Oor
// Out Of Resources, Limits to the resources
// allocatable
type Ipv4AclAndPrefixList_Oor struct {
    EntityData types.CommonEntityData
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

func (oor *Ipv4AclAndPrefixList_Oor) GetEntityData() *types.CommonEntityData {
    oor.EntityData.YFilter = oor.YFilter
    oor.EntityData.YangName = "oor"
    oor.EntityData.BundleName = "cisco_ios_xr"
    oor.EntityData.ParentYangName = "ipv4-acl-and-prefix-list"
    oor.EntityData.SegmentPath = "oor"
    oor.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/" + oor.EntityData.SegmentPath
    oor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oor.EntityData.Children = types.NewOrderedMap()
    oor.EntityData.Children.Append("details", types.YChild{"Details", &oor.Details})
    oor.EntityData.Children.Append("oor-prefixes", types.YChild{"OorPrefixes", &oor.OorPrefixes})
    oor.EntityData.Children.Append("oor-accesses", types.YChild{"OorAccesses", &oor.OorAccesses})
    oor.EntityData.Children.Append("access-list-summary", types.YChild{"AccessListSummary", &oor.AccessListSummary})
    oor.EntityData.Children.Append("prefix-list-summary", types.YChild{"PrefixListSummary", &oor.PrefixListSummary})
    oor.EntityData.Leafs = types.NewOrderedMap()

    oor.EntityData.YListKeys = []string {}

    return &(oor.EntityData)
}

// Ipv4AclAndPrefixList_Oor_Details
// Details of the Overall Out Of Resources Limits
type Ipv4AclAndPrefixList_Oor_Details struct {
    EntityData types.CommonEntityData
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

func (details *Ipv4AclAndPrefixList_Oor_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "oor"
    details.EntityData.SegmentPath = "details"
    details.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/oor/" + details.EntityData.SegmentPath
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = types.NewOrderedMap()
    details.EntityData.Leafs = types.NewOrderedMap()
    details.EntityData.Leafs.Append("default-max-ac-ls", types.YLeaf{"DefaultMaxAcLs", details.DefaultMaxAcLs})
    details.EntityData.Leafs.Append("default-max-ac-es", types.YLeaf{"DefaultMaxAcEs", details.DefaultMaxAcEs})
    details.EntityData.Leafs.Append("current-configured-ac-ls", types.YLeaf{"CurrentConfiguredAcLs", details.CurrentConfiguredAcLs})
    details.EntityData.Leafs.Append("current-configured-ac-es", types.YLeaf{"CurrentConfiguredAcEs", details.CurrentConfiguredAcEs})
    details.EntityData.Leafs.Append("current-max-configurable-ac-ls", types.YLeaf{"CurrentMaxConfigurableAcLs", details.CurrentMaxConfigurableAcLs})
    details.EntityData.Leafs.Append("current-max-configurable-ac-es", types.YLeaf{"CurrentMaxConfigurableAcEs", details.CurrentMaxConfigurableAcEs})
    details.EntityData.Leafs.Append("max-configurable-ac-ls", types.YLeaf{"MaxConfigurableAcLs", details.MaxConfigurableAcLs})
    details.EntityData.Leafs.Append("max-configurable-ac-es", types.YLeaf{"MaxConfigurableAcEs", details.MaxConfigurableAcEs})

    details.EntityData.YListKeys = []string {}

    return &(details.EntityData)
}

// Ipv4AclAndPrefixList_Oor_OorPrefixes
// Resource occupation details for prefix lists
type Ipv4AclAndPrefixList_Oor_OorPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource occupation details for a particular prefix list. The type is slice
    // of Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix.
    OorPrefix []*Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix
}

func (oorPrefixes *Ipv4AclAndPrefixList_Oor_OorPrefixes) GetEntityData() *types.CommonEntityData {
    oorPrefixes.EntityData.YFilter = oorPrefixes.YFilter
    oorPrefixes.EntityData.YangName = "oor-prefixes"
    oorPrefixes.EntityData.BundleName = "cisco_ios_xr"
    oorPrefixes.EntityData.ParentYangName = "oor"
    oorPrefixes.EntityData.SegmentPath = "oor-prefixes"
    oorPrefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/oor/" + oorPrefixes.EntityData.SegmentPath
    oorPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorPrefixes.EntityData.Children = types.NewOrderedMap()
    oorPrefixes.EntityData.Children.Append("oor-prefix", types.YChild{"OorPrefix", nil})
    for i := range oorPrefixes.OorPrefix {
        oorPrefixes.EntityData.Children.Append(types.GetSegmentPath(oorPrefixes.OorPrefix[i]), types.YChild{"OorPrefix", oorPrefixes.OorPrefix[i]})
    }
    oorPrefixes.EntityData.Leafs = types.NewOrderedMap()

    oorPrefixes.EntityData.YListKeys = []string {}

    return &(oorPrefixes.EntityData)
}

// Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix
// Resource occupation details for a particular
// prefix list
type Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (oorPrefix *Ipv4AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetEntityData() *types.CommonEntityData {
    oorPrefix.EntityData.YFilter = oorPrefix.YFilter
    oorPrefix.EntityData.YangName = "oor-prefix"
    oorPrefix.EntityData.BundleName = "cisco_ios_xr"
    oorPrefix.EntityData.ParentYangName = "oor-prefixes"
    oorPrefix.EntityData.SegmentPath = "oor-prefix" + types.AddKeyToken(oorPrefix.PrefixListName, "prefix-list-name")
    oorPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/oor/oor-prefixes/" + oorPrefix.EntityData.SegmentPath
    oorPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorPrefix.EntityData.Children = types.NewOrderedMap()
    oorPrefix.EntityData.Leafs = types.NewOrderedMap()
    oorPrefix.EntityData.Leafs.Append("prefix-list-name", types.YLeaf{"PrefixListName", oorPrefix.PrefixListName})
    oorPrefix.EntityData.Leafs.Append("default-max-ac-ls", types.YLeaf{"DefaultMaxAcLs", oorPrefix.DefaultMaxAcLs})
    oorPrefix.EntityData.Leafs.Append("default-max-ac-es", types.YLeaf{"DefaultMaxAcEs", oorPrefix.DefaultMaxAcEs})
    oorPrefix.EntityData.Leafs.Append("current-configured-ac-ls", types.YLeaf{"CurrentConfiguredAcLs", oorPrefix.CurrentConfiguredAcLs})
    oorPrefix.EntityData.Leafs.Append("current-configured-ac-es", types.YLeaf{"CurrentConfiguredAcEs", oorPrefix.CurrentConfiguredAcEs})
    oorPrefix.EntityData.Leafs.Append("current-max-configurable-ac-ls", types.YLeaf{"CurrentMaxConfigurableAcLs", oorPrefix.CurrentMaxConfigurableAcLs})
    oorPrefix.EntityData.Leafs.Append("current-max-configurable-ac-es", types.YLeaf{"CurrentMaxConfigurableAcEs", oorPrefix.CurrentMaxConfigurableAcEs})
    oorPrefix.EntityData.Leafs.Append("max-configurable-ac-ls", types.YLeaf{"MaxConfigurableAcLs", oorPrefix.MaxConfigurableAcLs})
    oorPrefix.EntityData.Leafs.Append("max-configurable-ac-es", types.YLeaf{"MaxConfigurableAcEs", oorPrefix.MaxConfigurableAcEs})

    oorPrefix.EntityData.YListKeys = []string {"PrefixListName"}

    return &(oorPrefix.EntityData)
}

// Ipv4AclAndPrefixList_Oor_OorAccesses
// Resource occupation details for access lists
type Ipv4AclAndPrefixList_Oor_OorAccesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource occupation details for a particular access list. The type is slice
    // of Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess.
    OorAccess []*Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess
}

func (oorAccesses *Ipv4AclAndPrefixList_Oor_OorAccesses) GetEntityData() *types.CommonEntityData {
    oorAccesses.EntityData.YFilter = oorAccesses.YFilter
    oorAccesses.EntityData.YangName = "oor-accesses"
    oorAccesses.EntityData.BundleName = "cisco_ios_xr"
    oorAccesses.EntityData.ParentYangName = "oor"
    oorAccesses.EntityData.SegmentPath = "oor-accesses"
    oorAccesses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/oor/" + oorAccesses.EntityData.SegmentPath
    oorAccesses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorAccesses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorAccesses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorAccesses.EntityData.Children = types.NewOrderedMap()
    oorAccesses.EntityData.Children.Append("oor-access", types.YChild{"OorAccess", nil})
    for i := range oorAccesses.OorAccess {
        oorAccesses.EntityData.Children.Append(types.GetSegmentPath(oorAccesses.OorAccess[i]), types.YChild{"OorAccess", oorAccesses.OorAccess[i]})
    }
    oorAccesses.EntityData.Leafs = types.NewOrderedMap()

    oorAccesses.EntityData.YListKeys = []string {}

    return &(oorAccesses.EntityData)
}

// Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess
// Resource occupation details for a particular
// access list
type Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (oorAccess *Ipv4AclAndPrefixList_Oor_OorAccesses_OorAccess) GetEntityData() *types.CommonEntityData {
    oorAccess.EntityData.YFilter = oorAccess.YFilter
    oorAccess.EntityData.YangName = "oor-access"
    oorAccess.EntityData.BundleName = "cisco_ios_xr"
    oorAccess.EntityData.ParentYangName = "oor-accesses"
    oorAccess.EntityData.SegmentPath = "oor-access" + types.AddKeyToken(oorAccess.AccessListName, "access-list-name")
    oorAccess.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/oor/oor-accesses/" + oorAccess.EntityData.SegmentPath
    oorAccess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorAccess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorAccess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorAccess.EntityData.Children = types.NewOrderedMap()
    oorAccess.EntityData.Leafs = types.NewOrderedMap()
    oorAccess.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", oorAccess.AccessListName})
    oorAccess.EntityData.Leafs.Append("default-max-ac-ls", types.YLeaf{"DefaultMaxAcLs", oorAccess.DefaultMaxAcLs})
    oorAccess.EntityData.Leafs.Append("default-max-ac-es", types.YLeaf{"DefaultMaxAcEs", oorAccess.DefaultMaxAcEs})
    oorAccess.EntityData.Leafs.Append("current-configured-ac-ls", types.YLeaf{"CurrentConfiguredAcLs", oorAccess.CurrentConfiguredAcLs})
    oorAccess.EntityData.Leafs.Append("current-configured-ac-es", types.YLeaf{"CurrentConfiguredAcEs", oorAccess.CurrentConfiguredAcEs})
    oorAccess.EntityData.Leafs.Append("current-max-configurable-ac-ls", types.YLeaf{"CurrentMaxConfigurableAcLs", oorAccess.CurrentMaxConfigurableAcLs})
    oorAccess.EntityData.Leafs.Append("current-max-configurable-ac-es", types.YLeaf{"CurrentMaxConfigurableAcEs", oorAccess.CurrentMaxConfigurableAcEs})
    oorAccess.EntityData.Leafs.Append("max-configurable-ac-ls", types.YLeaf{"MaxConfigurableAcLs", oorAccess.MaxConfigurableAcLs})
    oorAccess.EntityData.Leafs.Append("max-configurable-ac-es", types.YLeaf{"MaxConfigurableAcEs", oorAccess.MaxConfigurableAcEs})

    oorAccess.EntityData.YListKeys = []string {"AccessListName"}

    return &(oorAccess.EntityData)
}

// Ipv4AclAndPrefixList_Oor_AccessListSummary
// Resource limits pertaining to access lists only
type Ipv4AclAndPrefixList_Oor_AccessListSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Details containing the resource limits of the access lists.
    Details Ipv4AclAndPrefixList_Oor_AccessListSummary_Details
}

func (accessListSummary *Ipv4AclAndPrefixList_Oor_AccessListSummary) GetEntityData() *types.CommonEntityData {
    accessListSummary.EntityData.YFilter = accessListSummary.YFilter
    accessListSummary.EntityData.YangName = "access-list-summary"
    accessListSummary.EntityData.BundleName = "cisco_ios_xr"
    accessListSummary.EntityData.ParentYangName = "oor"
    accessListSummary.EntityData.SegmentPath = "access-list-summary"
    accessListSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/oor/" + accessListSummary.EntityData.SegmentPath
    accessListSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListSummary.EntityData.Children = types.NewOrderedMap()
    accessListSummary.EntityData.Children.Append("details", types.YChild{"Details", &accessListSummary.Details})
    accessListSummary.EntityData.Leafs = types.NewOrderedMap()

    accessListSummary.EntityData.YListKeys = []string {}

    return &(accessListSummary.EntityData)
}

// Ipv4AclAndPrefixList_Oor_AccessListSummary_Details
// Details containing the resource limits of the
// access lists
type Ipv4AclAndPrefixList_Oor_AccessListSummary_Details struct {
    EntityData types.CommonEntityData
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

func (details *Ipv4AclAndPrefixList_Oor_AccessListSummary_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "access-list-summary"
    details.EntityData.SegmentPath = "details"
    details.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/oor/access-list-summary/" + details.EntityData.SegmentPath
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = types.NewOrderedMap()
    details.EntityData.Leafs = types.NewOrderedMap()
    details.EntityData.Leafs.Append("default-max-ac-ls", types.YLeaf{"DefaultMaxAcLs", details.DefaultMaxAcLs})
    details.EntityData.Leafs.Append("default-max-ac-es", types.YLeaf{"DefaultMaxAcEs", details.DefaultMaxAcEs})
    details.EntityData.Leafs.Append("current-configured-ac-ls", types.YLeaf{"CurrentConfiguredAcLs", details.CurrentConfiguredAcLs})
    details.EntityData.Leafs.Append("current-configured-ac-es", types.YLeaf{"CurrentConfiguredAcEs", details.CurrentConfiguredAcEs})
    details.EntityData.Leafs.Append("current-max-configurable-ac-ls", types.YLeaf{"CurrentMaxConfigurableAcLs", details.CurrentMaxConfigurableAcLs})
    details.EntityData.Leafs.Append("current-max-configurable-ac-es", types.YLeaf{"CurrentMaxConfigurableAcEs", details.CurrentMaxConfigurableAcEs})
    details.EntityData.Leafs.Append("max-configurable-ac-ls", types.YLeaf{"MaxConfigurableAcLs", details.MaxConfigurableAcLs})
    details.EntityData.Leafs.Append("max-configurable-ac-es", types.YLeaf{"MaxConfigurableAcEs", details.MaxConfigurableAcEs})

    details.EntityData.YListKeys = []string {}

    return &(details.EntityData)
}

// Ipv4AclAndPrefixList_Oor_PrefixListSummary
// Summary of the prefix Lists resource
// utilization
type Ipv4AclAndPrefixList_Oor_PrefixListSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary Detail of the prefix list Resource Utilisation.
    Details Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details
}

func (prefixListSummary *Ipv4AclAndPrefixList_Oor_PrefixListSummary) GetEntityData() *types.CommonEntityData {
    prefixListSummary.EntityData.YFilter = prefixListSummary.YFilter
    prefixListSummary.EntityData.YangName = "prefix-list-summary"
    prefixListSummary.EntityData.BundleName = "cisco_ios_xr"
    prefixListSummary.EntityData.ParentYangName = "oor"
    prefixListSummary.EntityData.SegmentPath = "prefix-list-summary"
    prefixListSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/oor/" + prefixListSummary.EntityData.SegmentPath
    prefixListSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixListSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixListSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixListSummary.EntityData.Children = types.NewOrderedMap()
    prefixListSummary.EntityData.Children.Append("details", types.YChild{"Details", &prefixListSummary.Details})
    prefixListSummary.EntityData.Leafs = types.NewOrderedMap()

    prefixListSummary.EntityData.YListKeys = []string {}

    return &(prefixListSummary.EntityData)
}

// Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details
// Summary Detail of the prefix list Resource
// Utilisation
type Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details struct {
    EntityData types.CommonEntityData
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

func (details *Ipv4AclAndPrefixList_Oor_PrefixListSummary_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "prefix-list-summary"
    details.EntityData.SegmentPath = "details"
    details.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-acl-oper:ipv4-acl-and-prefix-list/oor/prefix-list-summary/" + details.EntityData.SegmentPath
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = types.NewOrderedMap()
    details.EntityData.Leafs = types.NewOrderedMap()
    details.EntityData.Leafs.Append("default-max-ac-ls", types.YLeaf{"DefaultMaxAcLs", details.DefaultMaxAcLs})
    details.EntityData.Leafs.Append("default-max-ac-es", types.YLeaf{"DefaultMaxAcEs", details.DefaultMaxAcEs})
    details.EntityData.Leafs.Append("current-configured-ac-ls", types.YLeaf{"CurrentConfiguredAcLs", details.CurrentConfiguredAcLs})
    details.EntityData.Leafs.Append("current-configured-ac-es", types.YLeaf{"CurrentConfiguredAcEs", details.CurrentConfiguredAcEs})
    details.EntityData.Leafs.Append("current-max-configurable-ac-ls", types.YLeaf{"CurrentMaxConfigurableAcLs", details.CurrentMaxConfigurableAcLs})
    details.EntityData.Leafs.Append("current-max-configurable-ac-es", types.YLeaf{"CurrentMaxConfigurableAcEs", details.CurrentMaxConfigurableAcEs})
    details.EntityData.Leafs.Append("max-configurable-ac-ls", types.YLeaf{"MaxConfigurableAcLs", details.MaxConfigurableAcLs})
    details.EntityData.Leafs.Append("max-configurable-ac-es", types.YLeaf{"MaxConfigurableAcEs", details.MaxConfigurableAcEs})

    details.EntityData.YListKeys = []string {}

    return &(details.EntityData)
}

