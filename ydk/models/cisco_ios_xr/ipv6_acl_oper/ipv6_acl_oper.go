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

// Ipv6AclAndPrefixList
// Root class of IPv6 Oper schema tree
type Ipv6AclAndPrefixList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AccessListManager containing ACLs and prefix lists.
    AccessListManager Ipv6AclAndPrefixList_AccessListManager

    // Out Of Resources, Limits to the resources allocatable.
    Oor Ipv6AclAndPrefixList_Oor
}

func (ipv6AclAndPrefixList *Ipv6AclAndPrefixList) GetEntityData() *types.CommonEntityData {
    ipv6AclAndPrefixList.EntityData.YFilter = ipv6AclAndPrefixList.YFilter
    ipv6AclAndPrefixList.EntityData.YangName = "ipv6-acl-and-prefix-list"
    ipv6AclAndPrefixList.EntityData.BundleName = "cisco_ios_xr"
    ipv6AclAndPrefixList.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-acl-oper"
    ipv6AclAndPrefixList.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-acl-oper:ipv6-acl-and-prefix-list"
    ipv6AclAndPrefixList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6AclAndPrefixList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6AclAndPrefixList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6AclAndPrefixList.EntityData.Children = make(map[string]types.YChild)
    ipv6AclAndPrefixList.EntityData.Children["access-list-manager"] = types.YChild{"AccessListManager", &ipv6AclAndPrefixList.AccessListManager}
    ipv6AclAndPrefixList.EntityData.Children["oor"] = types.YChild{"Oor", &ipv6AclAndPrefixList.Oor}
    ipv6AclAndPrefixList.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6AclAndPrefixList.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager
// AccessListManager containing ACLs and prefix
// lists
type Ipv6AclAndPrefixList_AccessListManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of prefix lists.
    Prefixes Ipv6AclAndPrefixList_AccessListManager_Prefixes

    // Table of Usage statistics of ACLs at different nodes.
    Usages Ipv6AclAndPrefixList_AccessListManager_Usages

    // ACL class displaying Usage and Entries.
    Accesses Ipv6AclAndPrefixList_AccessListManager_Accesses
}

func (accessListManager *Ipv6AclAndPrefixList_AccessListManager) GetEntityData() *types.CommonEntityData {
    accessListManager.EntityData.YFilter = accessListManager.YFilter
    accessListManager.EntityData.YangName = "access-list-manager"
    accessListManager.EntityData.BundleName = "cisco_ios_xr"
    accessListManager.EntityData.ParentYangName = "ipv6-acl-and-prefix-list"
    accessListManager.EntityData.SegmentPath = "access-list-manager"
    accessListManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListManager.EntityData.Children = make(map[string]types.YChild)
    accessListManager.EntityData.Children["prefixes"] = types.YChild{"Prefixes", &accessListManager.Prefixes}
    accessListManager.EntityData.Children["usages"] = types.YChild{"Usages", &accessListManager.Usages}
    accessListManager.EntityData.Children["accesses"] = types.YChild{"Accesses", &accessListManager.Accesses}
    accessListManager.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessListManager.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Prefixes
// Table of prefix lists
type Ipv6AclAndPrefixList_AccessListManager_Prefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the prefix list. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix.
    Prefix []Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix
}

func (prefixes *Ipv6AclAndPrefixList_AccessListManager_Prefixes) GetEntityData() *types.CommonEntityData {
    prefixes.EntityData.YFilter = prefixes.YFilter
    prefixes.EntityData.YangName = "prefixes"
    prefixes.EntityData.BundleName = "cisco_ios_xr"
    prefixes.EntityData.ParentYangName = "access-list-manager"
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

// Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix
// Name of the prefix list
type Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the prefix list. The type is string with
    // length: 1..65.
    PrefixListName interface{}

    // Table of all the SequenceNumbers per prefix list.
    PrefixListSequences Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences
}

func (prefix *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefixes"
    prefix.EntityData.SegmentPath = "prefix" + "[prefix-list-name='" + fmt.Sprintf("%v", prefix.PrefixListName) + "']"
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = make(map[string]types.YChild)
    prefix.EntityData.Children["prefix-list-sequences"] = types.YChild{"PrefixListSequences", &prefix.PrefixListSequences}
    prefix.EntityData.Leafs = make(map[string]types.YLeaf)
    prefix.EntityData.Leafs["prefix-list-name"] = types.YLeaf{"PrefixListName", prefix.PrefixListName}
    return &(prefix.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences
// Table of all the SequenceNumbers per prefix
// list
type Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sequence Number of a prefix list entry. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence.
    PrefixListSequence []Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence
}

func (prefixListSequences *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences) GetEntityData() *types.CommonEntityData {
    prefixListSequences.EntityData.YFilter = prefixListSequences.YFilter
    prefixListSequences.EntityData.YangName = "prefix-list-sequences"
    prefixListSequences.EntityData.BundleName = "cisco_ios_xr"
    prefixListSequences.EntityData.ParentYangName = "prefix"
    prefixListSequences.EntityData.SegmentPath = "prefix-list-sequences"
    prefixListSequences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixListSequences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixListSequences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixListSequences.EntityData.Children = make(map[string]types.YChild)
    prefixListSequences.EntityData.Children["prefix-list-sequence"] = types.YChild{"PrefixListSequence", nil}
    for i := range prefixListSequences.PrefixListSequence {
        prefixListSequences.EntityData.Children[types.GetSegmentPath(&prefixListSequences.PrefixListSequence[i])] = types.YChild{"PrefixListSequence", &prefixListSequences.PrefixListSequence[i]}
    }
    prefixListSequences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixListSequences.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence
// Sequence Number of a prefix list entry
type Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Sequence Number of the prefix list entry. The type
    // is interface{} with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (acl, remark). The type is AclAce1_.
    IsAceType interface{}

    // ACLE sequence number. The type is interface{} with range: 0..4294967295.
    IsAceSequenceNumber interface{}

    // Grant value permit/deny . The type is AclAction.
    IsPacketAllowOrDeny interface{}

    // IPv6 prefix. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IsAddressInNumbers interface{}

    // Prefix length . The type is interface{} with range: 0..4294967295.
    IsAddressMaskLength interface{}

    // Port Operator. The type is AclPortOperator_.
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

func (prefixListSequence *Ipv6AclAndPrefixList_AccessListManager_Prefixes_Prefix_PrefixListSequences_PrefixListSequence) GetEntityData() *types.CommonEntityData {
    prefixListSequence.EntityData.YFilter = prefixListSequence.YFilter
    prefixListSequence.EntityData.YangName = "prefix-list-sequence"
    prefixListSequence.EntityData.BundleName = "cisco_ios_xr"
    prefixListSequence.EntityData.ParentYangName = "prefix-list-sequences"
    prefixListSequence.EntityData.SegmentPath = "prefix-list-sequence" + "[sequence-number='" + fmt.Sprintf("%v", prefixListSequence.SequenceNumber) + "']"
    prefixListSequence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixListSequence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixListSequence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixListSequence.EntityData.Children = make(map[string]types.YChild)
    prefixListSequence.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixListSequence.EntityData.Leafs["sequence-number"] = types.YLeaf{"SequenceNumber", prefixListSequence.SequenceNumber}
    prefixListSequence.EntityData.Leafs["is-ace-type"] = types.YLeaf{"IsAceType", prefixListSequence.IsAceType}
    prefixListSequence.EntityData.Leafs["is-ace-sequence-number"] = types.YLeaf{"IsAceSequenceNumber", prefixListSequence.IsAceSequenceNumber}
    prefixListSequence.EntityData.Leafs["is-packet-allow-or-deny"] = types.YLeaf{"IsPacketAllowOrDeny", prefixListSequence.IsPacketAllowOrDeny}
    prefixListSequence.EntityData.Leafs["is-address-in-numbers"] = types.YLeaf{"IsAddressInNumbers", prefixListSequence.IsAddressInNumbers}
    prefixListSequence.EntityData.Leafs["is-address-mask-length"] = types.YLeaf{"IsAddressMaskLength", prefixListSequence.IsAddressMaskLength}
    prefixListSequence.EntityData.Leafs["is-length-operator"] = types.YLeaf{"IsLengthOperator", prefixListSequence.IsLengthOperator}
    prefixListSequence.EntityData.Leafs["is-packet-minimum-length"] = types.YLeaf{"IsPacketMinimumLength", prefixListSequence.IsPacketMinimumLength}
    prefixListSequence.EntityData.Leafs["is-packet-maximum-length"] = types.YLeaf{"IsPacketMaximumLength", prefixListSequence.IsPacketMaximumLength}
    prefixListSequence.EntityData.Leafs["hits"] = types.YLeaf{"Hits", prefixListSequence.Hits}
    prefixListSequence.EntityData.Leafs["is-comment-for-entry"] = types.YLeaf{"IsCommentForEntry", prefixListSequence.IsCommentForEntry}
    prefixListSequence.EntityData.Leafs["acl-name"] = types.YLeaf{"AclName", prefixListSequence.AclName}
    return &(prefixListSequence.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Usages
// Table of Usage statistics of ACLs at different
// nodes
type Ipv6AclAndPrefixList_AccessListManager_Usages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Usage statistics of an ACL at a node. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Usages_Usage.
    Usage []Ipv6AclAndPrefixList_AccessListManager_Usages_Usage
}

func (usages *Ipv6AclAndPrefixList_AccessListManager_Usages) GetEntityData() *types.CommonEntityData {
    usages.EntityData.YFilter = usages.YFilter
    usages.EntityData.YangName = "usages"
    usages.EntityData.BundleName = "cisco_ios_xr"
    usages.EntityData.ParentYangName = "access-list-manager"
    usages.EntityData.SegmentPath = "usages"
    usages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usages.EntityData.Children = make(map[string]types.YChild)
    usages.EntityData.Children["usage"] = types.YChild{"Usage", nil}
    for i := range usages.Usage {
        usages.EntityData.Children[types.GetSegmentPath(&usages.Usage[i])] = types.YChild{"Usage", &usages.Usage[i]}
    }
    usages.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usages.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Usages_Usage
// Usage statistics of an ACL at a node
type Ipv6AclAndPrefixList_AccessListManager_Usages_Usage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node where ACL is applied. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Application ID. The type is AclUsageAppIdEnum.
    ApplicationId interface{}

    // Name of the ACL. The type is string with length: 1..65.
    AccessListName interface{}

    // Usage Statistics Details. The type is string. This attribute is mandatory.
    UsageDetails interface{}
}

func (usage *Ipv6AclAndPrefixList_AccessListManager_Usages_Usage) GetEntityData() *types.CommonEntityData {
    usage.EntityData.YFilter = usage.YFilter
    usage.EntityData.YangName = "usage"
    usage.EntityData.BundleName = "cisco_ios_xr"
    usage.EntityData.ParentYangName = "usages"
    usage.EntityData.SegmentPath = "usage"
    usage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usage.EntityData.Children = make(map[string]types.YChild)
    usage.EntityData.Leafs = make(map[string]types.YLeaf)
    usage.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", usage.NodeName}
    usage.EntityData.Leafs["application-id"] = types.YLeaf{"ApplicationId", usage.ApplicationId}
    usage.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", usage.AccessListName}
    usage.EntityData.Leafs["usage-details"] = types.YLeaf{"UsageDetails", usage.UsageDetails}
    return &(usage.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Accesses
// ACL class displaying Usage and Entries
type Ipv6AclAndPrefixList_AccessListManager_Accesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the Access List. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Accesses_Access.
    Access []Ipv6AclAndPrefixList_AccessListManager_Accesses_Access
}

func (accesses *Ipv6AclAndPrefixList_AccessListManager_Accesses) GetEntityData() *types.CommonEntityData {
    accesses.EntityData.YFilter = accesses.YFilter
    accesses.EntityData.YangName = "accesses"
    accesses.EntityData.BundleName = "cisco_ios_xr"
    accesses.EntityData.ParentYangName = "access-list-manager"
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

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access
// Name of the Access List
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the Access List. The type is string with
    // length: 1..65.
    AccessListName interface{}

    // Table of all the sequence numbers per ACL.
    AccessListSequences Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences
}

func (access *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access) GetEntityData() *types.CommonEntityData {
    access.EntityData.YFilter = access.YFilter
    access.EntityData.YangName = "access"
    access.EntityData.BundleName = "cisco_ios_xr"
    access.EntityData.ParentYangName = "accesses"
    access.EntityData.SegmentPath = "access" + "[access-list-name='" + fmt.Sprintf("%v", access.AccessListName) + "']"
    access.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    access.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    access.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    access.EntityData.Children = make(map[string]types.YChild)
    access.EntityData.Children["access-list-sequences"] = types.YChild{"AccessListSequences", &access.AccessListSequences}
    access.EntityData.Leafs = make(map[string]types.YLeaf)
    access.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", access.AccessListName}
    return &(access.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences
// Table of all the sequence numbers per ACL
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sequence number of an ACL entry. The type is slice of
    // Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence.
    AccessListSequence []Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence
}

func (accessListSequences *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences) GetEntityData() *types.CommonEntityData {
    accessListSequences.EntityData.YFilter = accessListSequences.YFilter
    accessListSequences.EntityData.YangName = "access-list-sequences"
    accessListSequences.EntityData.BundleName = "cisco_ios_xr"
    accessListSequences.EntityData.ParentYangName = "access"
    accessListSequences.EntityData.SegmentPath = "access-list-sequences"
    accessListSequences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListSequences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListSequences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListSequences.EntityData.Children = make(map[string]types.YChild)
    accessListSequences.EntityData.Children["access-list-sequence"] = types.YChild{"AccessListSequence", nil}
    for i := range accessListSequences.AccessListSequence {
        accessListSequences.EntityData.Children[types.GetSegmentPath(&accessListSequences.AccessListSequence[i])] = types.YChild{"AccessListSequence", &accessListSequences.AccessListSequence[i]}
    }
    accessListSequences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessListSequences.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence
// Sequence number of an ACL entry
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ACL entry sequence number. The type is interface{}
    // with range: 1..2147483646.
    SequenceNumber interface{}

    // ACE type (acl, remark). The type is AclAce1_.
    IsAceType interface{}

    // ACLE sequence number. The type is interface{} with range: 0..4294967295.
    IsAceSequenceNumber interface{}

    // Grant value permit/deny . The type is AclAction.
    IsPacketAllowOrDeny interface{}

    // Protocol operator. The type is AclPortOperator_.
    IsProtocolOperator interface{}

    // Protocol 1. The type is interface{} with range: -2147483648..2147483647.
    IsIpv6ProtocolType interface{}

    // Protocol 2. The type is interface{} with range: -2147483648..2147483647.
    IsIpv6Protocol2Type interface{}

    // IsSourceAddressInNumbers. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IsSourceAddressInNumbers interface{}

    // IsSourceAddressPrefixLength. The type is interface{} with range:
    // 0..4294967295.
    IsSourceAddressPrefixLength interface{}

    // Source Mask. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceMask interface{}

    // IsDestinationAddressInNumbers. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IsDestinationAddressInNumbers interface{}

    // IsDestinationAddressPrefixLength. The type is interface{} with range:
    // 0..4294967295.
    IsDestinationAddressPrefixLength interface{}

    // Destination Mask. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DestinationMask interface{}

    // eq, ne, lt, etc... The type is AclPortOperator_.
    IsSourceOperator interface{}

    // IsSourcePort1. The type is interface{} with range: 0..4294967295.
    IsSourcePort1 interface{}

    // IsSourcePort2. The type is interface{} with range: 0..4294967295.
    IsSourcePort2 interface{}

    // eq, ne, lt, etc... The type is AclPortOperator_.
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

    // Match if routing header is presant. The type is AclPortOperator_.
    IsPacketLengthOperator interface{}

    // IsPacketLengthStart. The type is interface{} with range: 0..4294967295.
    IsPacketLengthStart interface{}

    // IsPacketLengthEnd. The type is interface{} with range: 0..4294967295.
    IsPacketLengthEnd interface{}

    // IsTimeToLiveOperator. The type is AclPortOperator_.
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

func (accessListSequence *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence) GetEntityData() *types.CommonEntityData {
    accessListSequence.EntityData.YFilter = accessListSequence.YFilter
    accessListSequence.EntityData.YangName = "access-list-sequence"
    accessListSequence.EntityData.BundleName = "cisco_ios_xr"
    accessListSequence.EntityData.ParentYangName = "access-list-sequences"
    accessListSequence.EntityData.SegmentPath = "access-list-sequence" + "[sequence-number='" + fmt.Sprintf("%v", accessListSequence.SequenceNumber) + "']"
    accessListSequence.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListSequence.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListSequence.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListSequence.EntityData.Children = make(map[string]types.YChild)
    accessListSequence.EntityData.Children["hw-next-hop-info"] = types.YChild{"HwNextHopInfo", &accessListSequence.HwNextHopInfo}
    accessListSequence.EntityData.Children["next-hop-info"] = types.YChild{"NextHopInfo", nil}
    for i := range accessListSequence.NextHopInfo {
        accessListSequence.EntityData.Children[types.GetSegmentPath(&accessListSequence.NextHopInfo[i])] = types.YChild{"NextHopInfo", &accessListSequence.NextHopInfo[i]}
    }
    accessListSequence.EntityData.Children["udf"] = types.YChild{"Udf", nil}
    for i := range accessListSequence.Udf {
        accessListSequence.EntityData.Children[types.GetSegmentPath(&accessListSequence.Udf[i])] = types.YChild{"Udf", &accessListSequence.Udf[i]}
    }
    accessListSequence.EntityData.Leafs = make(map[string]types.YLeaf)
    accessListSequence.EntityData.Leafs["sequence-number"] = types.YLeaf{"SequenceNumber", accessListSequence.SequenceNumber}
    accessListSequence.EntityData.Leafs["is-ace-type"] = types.YLeaf{"IsAceType", accessListSequence.IsAceType}
    accessListSequence.EntityData.Leafs["is-ace-sequence-number"] = types.YLeaf{"IsAceSequenceNumber", accessListSequence.IsAceSequenceNumber}
    accessListSequence.EntityData.Leafs["is-packet-allow-or-deny"] = types.YLeaf{"IsPacketAllowOrDeny", accessListSequence.IsPacketAllowOrDeny}
    accessListSequence.EntityData.Leafs["is-protocol-operator"] = types.YLeaf{"IsProtocolOperator", accessListSequence.IsProtocolOperator}
    accessListSequence.EntityData.Leafs["is-ipv6-protocol-type"] = types.YLeaf{"IsIpv6ProtocolType", accessListSequence.IsIpv6ProtocolType}
    accessListSequence.EntityData.Leafs["is-ipv6-protocol2-type"] = types.YLeaf{"IsIpv6Protocol2Type", accessListSequence.IsIpv6Protocol2Type}
    accessListSequence.EntityData.Leafs["is-source-address-in-numbers"] = types.YLeaf{"IsSourceAddressInNumbers", accessListSequence.IsSourceAddressInNumbers}
    accessListSequence.EntityData.Leafs["is-source-address-prefix-length"] = types.YLeaf{"IsSourceAddressPrefixLength", accessListSequence.IsSourceAddressPrefixLength}
    accessListSequence.EntityData.Leafs["source-mask"] = types.YLeaf{"SourceMask", accessListSequence.SourceMask}
    accessListSequence.EntityData.Leafs["is-destination-address-in-numbers"] = types.YLeaf{"IsDestinationAddressInNumbers", accessListSequence.IsDestinationAddressInNumbers}
    accessListSequence.EntityData.Leafs["is-destination-address-prefix-length"] = types.YLeaf{"IsDestinationAddressPrefixLength", accessListSequence.IsDestinationAddressPrefixLength}
    accessListSequence.EntityData.Leafs["destination-mask"] = types.YLeaf{"DestinationMask", accessListSequence.DestinationMask}
    accessListSequence.EntityData.Leafs["is-source-operator"] = types.YLeaf{"IsSourceOperator", accessListSequence.IsSourceOperator}
    accessListSequence.EntityData.Leafs["is-source-port1"] = types.YLeaf{"IsSourcePort1", accessListSequence.IsSourcePort1}
    accessListSequence.EntityData.Leafs["is-source-port2"] = types.YLeaf{"IsSourcePort2", accessListSequence.IsSourcePort2}
    accessListSequence.EntityData.Leafs["is-destination-operator"] = types.YLeaf{"IsDestinationOperator", accessListSequence.IsDestinationOperator}
    accessListSequence.EntityData.Leafs["is-destination-port1"] = types.YLeaf{"IsDestinationPort1", accessListSequence.IsDestinationPort1}
    accessListSequence.EntityData.Leafs["is-destination-port2"] = types.YLeaf{"IsDestinationPort2", accessListSequence.IsDestinationPort2}
    accessListSequence.EntityData.Leafs["is-log-option"] = types.YLeaf{"IsLogOption", accessListSequence.IsLogOption}
    accessListSequence.EntityData.Leafs["counter-name"] = types.YLeaf{"CounterName", accessListSequence.CounterName}
    accessListSequence.EntityData.Leafs["is-tcp-bits-operator"] = types.YLeaf{"IsTcpBitsOperator", accessListSequence.IsTcpBitsOperator}
    accessListSequence.EntityData.Leafs["is-tcp-bits"] = types.YLeaf{"IsTcpBits", accessListSequence.IsTcpBits}
    accessListSequence.EntityData.Leafs["is-tcp-bits-mask"] = types.YLeaf{"IsTcpBitsMask", accessListSequence.IsTcpBitsMask}
    accessListSequence.EntityData.Leafs["is-dscp-present"] = types.YLeaf{"IsDscpPresent", accessListSequence.IsDscpPresent}
    accessListSequence.EntityData.Leafs["is-dscp-valu"] = types.YLeaf{"IsDscpValu", accessListSequence.IsDscpValu}
    accessListSequence.EntityData.Leafs["is-precedence-present"] = types.YLeaf{"IsPrecedencePresent", accessListSequence.IsPrecedencePresent}
    accessListSequence.EntityData.Leafs["is-precedence-value"] = types.YLeaf{"IsPrecedenceValue", accessListSequence.IsPrecedenceValue}
    accessListSequence.EntityData.Leafs["is-header-matches"] = types.YLeaf{"IsHeaderMatches", accessListSequence.IsHeaderMatches}
    accessListSequence.EntityData.Leafs["is-packet-length-operator"] = types.YLeaf{"IsPacketLengthOperator", accessListSequence.IsPacketLengthOperator}
    accessListSequence.EntityData.Leafs["is-packet-length-start"] = types.YLeaf{"IsPacketLengthStart", accessListSequence.IsPacketLengthStart}
    accessListSequence.EntityData.Leafs["is-packet-length-end"] = types.YLeaf{"IsPacketLengthEnd", accessListSequence.IsPacketLengthEnd}
    accessListSequence.EntityData.Leafs["is-time-to-live-operator"] = types.YLeaf{"IsTimeToLiveOperator", accessListSequence.IsTimeToLiveOperator}
    accessListSequence.EntityData.Leafs["is-time-to-live-start"] = types.YLeaf{"IsTimeToLiveStart", accessListSequence.IsTimeToLiveStart}
    accessListSequence.EntityData.Leafs["is-time-to-live-end"] = types.YLeaf{"IsTimeToLiveEnd", accessListSequence.IsTimeToLiveEnd}
    accessListSequence.EntityData.Leafs["no-stats"] = types.YLeaf{"NoStats", accessListSequence.NoStats}
    accessListSequence.EntityData.Leafs["hits"] = types.YLeaf{"Hits", accessListSequence.Hits}
    accessListSequence.EntityData.Leafs["capture"] = types.YLeaf{"Capture", accessListSequence.Capture}
    accessListSequence.EntityData.Leafs["undetermined-transport"] = types.YLeaf{"UndeterminedTransport", accessListSequence.UndeterminedTransport}
    accessListSequence.EntityData.Leafs["is-icmp-message-off"] = types.YLeaf{"IsIcmpMessageOff", accessListSequence.IsIcmpMessageOff}
    accessListSequence.EntityData.Leafs["qos-group"] = types.YLeaf{"QosGroup", accessListSequence.QosGroup}
    accessListSequence.EntityData.Leafs["is-comment-for-entry"] = types.YLeaf{"IsCommentForEntry", accessListSequence.IsCommentForEntry}
    accessListSequence.EntityData.Leafs["next-hop-type"] = types.YLeaf{"NextHopType", accessListSequence.NextHopType}
    accessListSequence.EntityData.Leafs["is-flow-id"] = types.YLeaf{"IsFlowId", accessListSequence.IsFlowId}
    accessListSequence.EntityData.Leafs["source-prefix-group"] = types.YLeaf{"SourcePrefixGroup", accessListSequence.SourcePrefixGroup}
    accessListSequence.EntityData.Leafs["destination-prefix-group"] = types.YLeaf{"DestinationPrefixGroup", accessListSequence.DestinationPrefixGroup}
    accessListSequence.EntityData.Leafs["source-port-group"] = types.YLeaf{"SourcePortGroup", accessListSequence.SourcePortGroup}
    accessListSequence.EntityData.Leafs["destination-port-group"] = types.YLeaf{"DestinationPortGroup", accessListSequence.DestinationPortGroup}
    accessListSequence.EntityData.Leafs["acl-name"] = types.YLeaf{"AclName", accessListSequence.AclName}
    accessListSequence.EntityData.Leafs["sequence-str"] = types.YLeaf{"SequenceStr", accessListSequence.SequenceStr}
    accessListSequence.EntityData.Leafs["set-ttl"] = types.YLeaf{"SetTtl", accessListSequence.SetTtl}
    return &(accessListSequence.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo
// HW Next hop info
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The next-hop type. The type is BagAclNh.
    Type_ interface{}

    // The Next Hop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHop interface{}

    // Table ID. The type is interface{} with range: 0..4294967295.
    TableId interface{}

    // Vrf Name. The type is string with length: 0..32.
    VrfName interface{}
}

func (hwNextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_HwNextHopInfo) GetEntityData() *types.CommonEntityData {
    hwNextHopInfo.EntityData.YFilter = hwNextHopInfo.YFilter
    hwNextHopInfo.EntityData.YangName = "hw-next-hop-info"
    hwNextHopInfo.EntityData.BundleName = "cisco_ios_xr"
    hwNextHopInfo.EntityData.ParentYangName = "access-list-sequence"
    hwNextHopInfo.EntityData.SegmentPath = "hw-next-hop-info"
    hwNextHopInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwNextHopInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwNextHopInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwNextHopInfo.EntityData.Children = make(map[string]types.YChild)
    hwNextHopInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    hwNextHopInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", hwNextHopInfo.Type_}
    hwNextHopInfo.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", hwNextHopInfo.NextHop}
    hwNextHopInfo.EntityData.Leafs["table-id"] = types.YLeaf{"TableId", hwNextHopInfo.TableId}
    hwNextHopInfo.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", hwNextHopInfo.VrfName}
    return &(hwNextHopInfo.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo
// Next hop info
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The next hop. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (nextHopInfo *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_NextHopInfo) GetEntityData() *types.CommonEntityData {
    nextHopInfo.EntityData.YFilter = nextHopInfo.YFilter
    nextHopInfo.EntityData.YangName = "next-hop-info"
    nextHopInfo.EntityData.BundleName = "cisco_ios_xr"
    nextHopInfo.EntityData.ParentYangName = "access-list-sequence"
    nextHopInfo.EntityData.SegmentPath = "next-hop-info"
    nextHopInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHopInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHopInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHopInfo.EntityData.Children = make(map[string]types.YChild)
    nextHopInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHopInfo.EntityData.Leafs["next-hop"] = types.YLeaf{"NextHop", nextHopInfo.NextHop}
    nextHopInfo.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", nextHopInfo.VrfName}
    nextHopInfo.EntityData.Leafs["track-name"] = types.YLeaf{"TrackName", nextHopInfo.TrackName}
    nextHopInfo.EntityData.Leafs["status"] = types.YLeaf{"Status", nextHopInfo.Status}
    nextHopInfo.EntityData.Leafs["at-status"] = types.YLeaf{"AtStatus", nextHopInfo.AtStatus}
    nextHopInfo.EntityData.Leafs["acl-nh-exist"] = types.YLeaf{"AclNhExist", nextHopInfo.AclNhExist}
    return &(nextHopInfo.EntityData)
}

// Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf
// UDF
type Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UDF Name. The type is string with length: 0..17.
    UdfName interface{}

    // UDF Value. The type is interface{} with range: 0..4294967295.
    UdfValue interface{}

    // UDF Mask. The type is interface{} with range: 0..4294967295.
    UdfMask interface{}
}

func (udf *Ipv6AclAndPrefixList_AccessListManager_Accesses_Access_AccessListSequences_AccessListSequence_Udf) GetEntityData() *types.CommonEntityData {
    udf.EntityData.YFilter = udf.YFilter
    udf.EntityData.YangName = "udf"
    udf.EntityData.BundleName = "cisco_ios_xr"
    udf.EntityData.ParentYangName = "access-list-sequence"
    udf.EntityData.SegmentPath = "udf"
    udf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    udf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    udf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    udf.EntityData.Children = make(map[string]types.YChild)
    udf.EntityData.Leafs = make(map[string]types.YLeaf)
    udf.EntityData.Leafs["udf-name"] = types.YLeaf{"UdfName", udf.UdfName}
    udf.EntityData.Leafs["udf-value"] = types.YLeaf{"UdfValue", udf.UdfValue}
    udf.EntityData.Leafs["udf-mask"] = types.YLeaf{"UdfMask", udf.UdfMask}
    return &(udf.EntityData)
}

// Ipv6AclAndPrefixList_Oor
// Out Of Resources, Limits to the resources
// allocatable
type Ipv6AclAndPrefixList_Oor struct {
    EntityData types.CommonEntityData
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

func (oor *Ipv6AclAndPrefixList_Oor) GetEntityData() *types.CommonEntityData {
    oor.EntityData.YFilter = oor.YFilter
    oor.EntityData.YangName = "oor"
    oor.EntityData.BundleName = "cisco_ios_xr"
    oor.EntityData.ParentYangName = "ipv6-acl-and-prefix-list"
    oor.EntityData.SegmentPath = "oor"
    oor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oor.EntityData.Children = make(map[string]types.YChild)
    oor.EntityData.Children["details"] = types.YChild{"Details", &oor.Details}
    oor.EntityData.Children["prefix-list-summary"] = types.YChild{"PrefixListSummary", &oor.PrefixListSummary}
    oor.EntityData.Children["oor-accesses"] = types.YChild{"OorAccesses", &oor.OorAccesses}
    oor.EntityData.Children["oor-prefixes"] = types.YChild{"OorPrefixes", &oor.OorPrefixes}
    oor.EntityData.Children["access-list-summary"] = types.YChild{"AccessListSummary", &oor.AccessListSummary}
    oor.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oor.EntityData)
}

// Ipv6AclAndPrefixList_Oor_Details
// Details of the overall out of resource limit
type Ipv6AclAndPrefixList_Oor_Details struct {
    EntityData types.CommonEntityData
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

func (details *Ipv6AclAndPrefixList_Oor_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "oor"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    details.EntityData.Leafs["is-default-maximum-configurable-ac-ls"] = types.YLeaf{"IsDefaultMaximumConfigurableAcLs", details.IsDefaultMaximumConfigurableAcLs}
    details.EntityData.Leafs["is-default-maximum-configurable-ac-es"] = types.YLeaf{"IsDefaultMaximumConfigurableAcEs", details.IsDefaultMaximumConfigurableAcEs}
    details.EntityData.Leafs["is-current-configured-ac-ls"] = types.YLeaf{"IsCurrentConfiguredAcLs", details.IsCurrentConfiguredAcLs}
    details.EntityData.Leafs["is-current-configured-aces"] = types.YLeaf{"IsCurrentConfiguredAces", details.IsCurrentConfiguredAces}
    details.EntityData.Leafs["is-current-maximum-configurable-acls"] = types.YLeaf{"IsCurrentMaximumConfigurableAcls", details.IsCurrentMaximumConfigurableAcls}
    details.EntityData.Leafs["is-current-maximum-configurable-aces"] = types.YLeaf{"IsCurrentMaximumConfigurableAces", details.IsCurrentMaximumConfigurableAces}
    details.EntityData.Leafs["is-maximum-configurable-ac-ls"] = types.YLeaf{"IsMaximumConfigurableAcLs", details.IsMaximumConfigurableAcLs}
    details.EntityData.Leafs["is-maximum-configurable-ac-es"] = types.YLeaf{"IsMaximumConfigurableAcEs", details.IsMaximumConfigurableAcEs}
    return &(details.EntityData)
}

// Ipv6AclAndPrefixList_Oor_PrefixListSummary
// Summary of the prefix Lists resource
// utilization
type Ipv6AclAndPrefixList_Oor_PrefixListSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary Detail of the prefix list Resource Utilisation.
    Details Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details
}

func (prefixListSummary *Ipv6AclAndPrefixList_Oor_PrefixListSummary) GetEntityData() *types.CommonEntityData {
    prefixListSummary.EntityData.YFilter = prefixListSummary.YFilter
    prefixListSummary.EntityData.YangName = "prefix-list-summary"
    prefixListSummary.EntityData.BundleName = "cisco_ios_xr"
    prefixListSummary.EntityData.ParentYangName = "oor"
    prefixListSummary.EntityData.SegmentPath = "prefix-list-summary"
    prefixListSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixListSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixListSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixListSummary.EntityData.Children = make(map[string]types.YChild)
    prefixListSummary.EntityData.Children["details"] = types.YChild{"Details", &prefixListSummary.Details}
    prefixListSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixListSummary.EntityData)
}

// Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details
// Summary Detail of the prefix list Resource
// Utilisation
type Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details struct {
    EntityData types.CommonEntityData
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

func (details *Ipv6AclAndPrefixList_Oor_PrefixListSummary_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "prefix-list-summary"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    details.EntityData.Leafs["is-default-maximum-configurable-ac-ls"] = types.YLeaf{"IsDefaultMaximumConfigurableAcLs", details.IsDefaultMaximumConfigurableAcLs}
    details.EntityData.Leafs["is-default-maximum-configurable-ac-es"] = types.YLeaf{"IsDefaultMaximumConfigurableAcEs", details.IsDefaultMaximumConfigurableAcEs}
    details.EntityData.Leafs["is-current-configured-ac-ls"] = types.YLeaf{"IsCurrentConfiguredAcLs", details.IsCurrentConfiguredAcLs}
    details.EntityData.Leafs["is-current-configured-aces"] = types.YLeaf{"IsCurrentConfiguredAces", details.IsCurrentConfiguredAces}
    details.EntityData.Leafs["is-current-maximum-configurable-acls"] = types.YLeaf{"IsCurrentMaximumConfigurableAcls", details.IsCurrentMaximumConfigurableAcls}
    details.EntityData.Leafs["is-current-maximum-configurable-aces"] = types.YLeaf{"IsCurrentMaximumConfigurableAces", details.IsCurrentMaximumConfigurableAces}
    details.EntityData.Leafs["is-maximum-configurable-ac-ls"] = types.YLeaf{"IsMaximumConfigurableAcLs", details.IsMaximumConfigurableAcLs}
    details.EntityData.Leafs["is-maximum-configurable-ac-es"] = types.YLeaf{"IsMaximumConfigurableAcEs", details.IsMaximumConfigurableAcEs}
    return &(details.EntityData)
}

// Ipv6AclAndPrefixList_Oor_OorAccesses
// Resource occupation details for ACLs
type Ipv6AclAndPrefixList_Oor_OorAccesses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource occupation details for a particular ACL. The type is slice of
    // Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess.
    OorAccess []Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess
}

func (oorAccesses *Ipv6AclAndPrefixList_Oor_OorAccesses) GetEntityData() *types.CommonEntityData {
    oorAccesses.EntityData.YFilter = oorAccesses.YFilter
    oorAccesses.EntityData.YangName = "oor-accesses"
    oorAccesses.EntityData.BundleName = "cisco_ios_xr"
    oorAccesses.EntityData.ParentYangName = "oor"
    oorAccesses.EntityData.SegmentPath = "oor-accesses"
    oorAccesses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorAccesses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorAccesses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorAccesses.EntityData.Children = make(map[string]types.YChild)
    oorAccesses.EntityData.Children["oor-access"] = types.YChild{"OorAccess", nil}
    for i := range oorAccesses.OorAccess {
        oorAccesses.EntityData.Children[types.GetSegmentPath(&oorAccesses.OorAccess[i])] = types.YChild{"OorAccess", &oorAccesses.OorAccess[i]}
    }
    oorAccesses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oorAccesses.EntityData)
}

// Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess
// Resource occupation details for a particular
// ACL
type Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess struct {
    EntityData types.CommonEntityData
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

func (oorAccess *Ipv6AclAndPrefixList_Oor_OorAccesses_OorAccess) GetEntityData() *types.CommonEntityData {
    oorAccess.EntityData.YFilter = oorAccess.YFilter
    oorAccess.EntityData.YangName = "oor-access"
    oorAccess.EntityData.BundleName = "cisco_ios_xr"
    oorAccess.EntityData.ParentYangName = "oor-accesses"
    oorAccess.EntityData.SegmentPath = "oor-access" + "[access-list-name='" + fmt.Sprintf("%v", oorAccess.AccessListName) + "']"
    oorAccess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorAccess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorAccess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorAccess.EntityData.Children = make(map[string]types.YChild)
    oorAccess.EntityData.Leafs = make(map[string]types.YLeaf)
    oorAccess.EntityData.Leafs["access-list-name"] = types.YLeaf{"AccessListName", oorAccess.AccessListName}
    oorAccess.EntityData.Leafs["is-default-maximum-configurable-ac-ls"] = types.YLeaf{"IsDefaultMaximumConfigurableAcLs", oorAccess.IsDefaultMaximumConfigurableAcLs}
    oorAccess.EntityData.Leafs["is-default-maximum-configurable-ac-es"] = types.YLeaf{"IsDefaultMaximumConfigurableAcEs", oorAccess.IsDefaultMaximumConfigurableAcEs}
    oorAccess.EntityData.Leafs["is-current-configured-ac-ls"] = types.YLeaf{"IsCurrentConfiguredAcLs", oorAccess.IsCurrentConfiguredAcLs}
    oorAccess.EntityData.Leafs["is-current-configured-aces"] = types.YLeaf{"IsCurrentConfiguredAces", oorAccess.IsCurrentConfiguredAces}
    oorAccess.EntityData.Leafs["is-current-maximum-configurable-acls"] = types.YLeaf{"IsCurrentMaximumConfigurableAcls", oorAccess.IsCurrentMaximumConfigurableAcls}
    oorAccess.EntityData.Leafs["is-current-maximum-configurable-aces"] = types.YLeaf{"IsCurrentMaximumConfigurableAces", oorAccess.IsCurrentMaximumConfigurableAces}
    oorAccess.EntityData.Leafs["is-maximum-configurable-ac-ls"] = types.YLeaf{"IsMaximumConfigurableAcLs", oorAccess.IsMaximumConfigurableAcLs}
    oorAccess.EntityData.Leafs["is-maximum-configurable-ac-es"] = types.YLeaf{"IsMaximumConfigurableAcEs", oorAccess.IsMaximumConfigurableAcEs}
    return &(oorAccess.EntityData)
}

// Ipv6AclAndPrefixList_Oor_OorPrefixes
// Resource occupation details for prefix lists
type Ipv6AclAndPrefixList_Oor_OorPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Resource occupation details for a particular prefix list. The type is slice
    // of Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix.
    OorPrefix []Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix
}

func (oorPrefixes *Ipv6AclAndPrefixList_Oor_OorPrefixes) GetEntityData() *types.CommonEntityData {
    oorPrefixes.EntityData.YFilter = oorPrefixes.YFilter
    oorPrefixes.EntityData.YangName = "oor-prefixes"
    oorPrefixes.EntityData.BundleName = "cisco_ios_xr"
    oorPrefixes.EntityData.ParentYangName = "oor"
    oorPrefixes.EntityData.SegmentPath = "oor-prefixes"
    oorPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorPrefixes.EntityData.Children = make(map[string]types.YChild)
    oorPrefixes.EntityData.Children["oor-prefix"] = types.YChild{"OorPrefix", nil}
    for i := range oorPrefixes.OorPrefix {
        oorPrefixes.EntityData.Children[types.GetSegmentPath(&oorPrefixes.OorPrefix[i])] = types.YChild{"OorPrefix", &oorPrefixes.OorPrefix[i]}
    }
    oorPrefixes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oorPrefixes.EntityData)
}

// Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix
// Resource occupation details for a particular
// prefix list
type Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix struct {
    EntityData types.CommonEntityData
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

func (oorPrefix *Ipv6AclAndPrefixList_Oor_OorPrefixes_OorPrefix) GetEntityData() *types.CommonEntityData {
    oorPrefix.EntityData.YFilter = oorPrefix.YFilter
    oorPrefix.EntityData.YangName = "oor-prefix"
    oorPrefix.EntityData.BundleName = "cisco_ios_xr"
    oorPrefix.EntityData.ParentYangName = "oor-prefixes"
    oorPrefix.EntityData.SegmentPath = "oor-prefix" + "[prefix-list-name='" + fmt.Sprintf("%v", oorPrefix.PrefixListName) + "']"
    oorPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorPrefix.EntityData.Children = make(map[string]types.YChild)
    oorPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    oorPrefix.EntityData.Leafs["prefix-list-name"] = types.YLeaf{"PrefixListName", oorPrefix.PrefixListName}
    oorPrefix.EntityData.Leafs["is-default-maximum-configurable-ac-ls"] = types.YLeaf{"IsDefaultMaximumConfigurableAcLs", oorPrefix.IsDefaultMaximumConfigurableAcLs}
    oorPrefix.EntityData.Leafs["is-default-maximum-configurable-ac-es"] = types.YLeaf{"IsDefaultMaximumConfigurableAcEs", oorPrefix.IsDefaultMaximumConfigurableAcEs}
    oorPrefix.EntityData.Leafs["is-current-configured-ac-ls"] = types.YLeaf{"IsCurrentConfiguredAcLs", oorPrefix.IsCurrentConfiguredAcLs}
    oorPrefix.EntityData.Leafs["is-current-configured-aces"] = types.YLeaf{"IsCurrentConfiguredAces", oorPrefix.IsCurrentConfiguredAces}
    oorPrefix.EntityData.Leafs["is-current-maximum-configurable-acls"] = types.YLeaf{"IsCurrentMaximumConfigurableAcls", oorPrefix.IsCurrentMaximumConfigurableAcls}
    oorPrefix.EntityData.Leafs["is-current-maximum-configurable-aces"] = types.YLeaf{"IsCurrentMaximumConfigurableAces", oorPrefix.IsCurrentMaximumConfigurableAces}
    oorPrefix.EntityData.Leafs["is-maximum-configurable-ac-ls"] = types.YLeaf{"IsMaximumConfigurableAcLs", oorPrefix.IsMaximumConfigurableAcLs}
    oorPrefix.EntityData.Leafs["is-maximum-configurable-ac-es"] = types.YLeaf{"IsMaximumConfigurableAcEs", oorPrefix.IsMaximumConfigurableAcEs}
    return &(oorPrefix.EntityData)
}

// Ipv6AclAndPrefixList_Oor_AccessListSummary
// Resource Limits pertaining to ACLs only
type Ipv6AclAndPrefixList_Oor_AccessListSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Details containing the resource limits of the ACLs.
    Details Ipv6AclAndPrefixList_Oor_AccessListSummary_Details
}

func (accessListSummary *Ipv6AclAndPrefixList_Oor_AccessListSummary) GetEntityData() *types.CommonEntityData {
    accessListSummary.EntityData.YFilter = accessListSummary.YFilter
    accessListSummary.EntityData.YangName = "access-list-summary"
    accessListSummary.EntityData.BundleName = "cisco_ios_xr"
    accessListSummary.EntityData.ParentYangName = "oor"
    accessListSummary.EntityData.SegmentPath = "access-list-summary"
    accessListSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessListSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessListSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessListSummary.EntityData.Children = make(map[string]types.YChild)
    accessListSummary.EntityData.Children["details"] = types.YChild{"Details", &accessListSummary.Details}
    accessListSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessListSummary.EntityData)
}

// Ipv6AclAndPrefixList_Oor_AccessListSummary_Details
// Details containing the resource limits of the
// ACLs
type Ipv6AclAndPrefixList_Oor_AccessListSummary_Details struct {
    EntityData types.CommonEntityData
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

func (details *Ipv6AclAndPrefixList_Oor_AccessListSummary_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "access-list-summary"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    details.EntityData.Leafs["is-default-maximum-configurable-ac-ls"] = types.YLeaf{"IsDefaultMaximumConfigurableAcLs", details.IsDefaultMaximumConfigurableAcLs}
    details.EntityData.Leafs["is-default-maximum-configurable-ac-es"] = types.YLeaf{"IsDefaultMaximumConfigurableAcEs", details.IsDefaultMaximumConfigurableAcEs}
    details.EntityData.Leafs["is-current-configured-ac-ls"] = types.YLeaf{"IsCurrentConfiguredAcLs", details.IsCurrentConfiguredAcLs}
    details.EntityData.Leafs["is-current-configured-aces"] = types.YLeaf{"IsCurrentConfiguredAces", details.IsCurrentConfiguredAces}
    details.EntityData.Leafs["is-current-maximum-configurable-acls"] = types.YLeaf{"IsCurrentMaximumConfigurableAcls", details.IsCurrentMaximumConfigurableAcls}
    details.EntityData.Leafs["is-current-maximum-configurable-aces"] = types.YLeaf{"IsCurrentMaximumConfigurableAces", details.IsCurrentMaximumConfigurableAces}
    details.EntityData.Leafs["is-maximum-configurable-ac-ls"] = types.YLeaf{"IsMaximumConfigurableAcLs", details.IsMaximumConfigurableAcLs}
    details.EntityData.Leafs["is-maximum-configurable-ac-es"] = types.YLeaf{"IsMaximumConfigurableAcEs", details.IsMaximumConfigurableAcEs}
    return &(details.EntityData)
}

