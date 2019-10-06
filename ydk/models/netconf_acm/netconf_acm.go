// NETCONF Access Control Model.
// 
// Copyright (c) 2012 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD
// License set forth in Section 4.c of the IETF Trust's
// Legal Provisions Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC 6536; see
// the RFC itself for full legal notices.
package netconf_acm

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package netconf_acm"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-netconf-acm nacm}", reflect.TypeOf(Nacm{}))
    ydk.RegisterEntity("ietf-netconf-acm:nacm", reflect.TypeOf(Nacm{}))
}

// ActionType represents rule matches.
type ActionType string

const (
    // Requested action is permitted.
    ActionType_permit ActionType = "permit"

    // Requested action is denied.
    ActionType_deny ActionType = "deny"
)

// Nacm
// Parameters for NETCONF Access Control Model.
type Nacm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables or disables all NETCONF access control enforcement.  If 'true',
    // then enforcement is enabled.  If 'false', then enforcement is disabled. The
    // type is bool. The default value is true.
    EnableNacm interface{}

    // Controls whether read access is granted if no appropriate rule is found for
    // a particular read request. The type is ActionType. The default value is
    // permit.
    ReadDefault interface{}

    // Controls whether create, update, or delete access is granted if no
    // appropriate rule is found for a particular write request. The type is
    // ActionType. The default value is deny.
    WriteDefault interface{}

    // Controls whether exec access is granted if no appropriate rule is found for
    // a particular protocol operation request. The type is ActionType. The
    // default value is permit.
    ExecDefault interface{}

    // Controls whether the server uses the groups reported by the NETCONF
    // transport layer when it assigns the user to a set of NACM groups.  If this
    // leaf has the value 'false', any group names reported by the transport layer
    // are ignored by the server. The type is bool. The default value is true.
    EnableExternalGroups interface{}

    // Number of times since the server last restarted that a protocol operation
    // request was denied. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    DeniedOperations interface{}

    // Number of times since the server last restarted that a protocol operation
    // request to alter a configuration datastore was denied. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    DeniedDataWrites interface{}

    // Number of times since the server last restarted that a notification was
    // dropped for a subscription because access to the event type was denied. The
    // type is interface{} with range: 0..4294967295. This attribute is mandatory.
    DeniedNotifications interface{}

    // NETCONF Access Control Groups.
    Groups Nacm_Groups

    // An ordered collection of access control rules. The type is slice of
    // Nacm_RuleList.
    RuleList []*Nacm_RuleList
}

func (nacm *Nacm) GetEntityData() *types.CommonEntityData {
    nacm.EntityData.YFilter = nacm.YFilter
    nacm.EntityData.YangName = "nacm"
    nacm.EntityData.BundleName = "ietf"
    nacm.EntityData.ParentYangName = "ietf-netconf-acm"
    nacm.EntityData.SegmentPath = "ietf-netconf-acm:nacm"
    nacm.EntityData.AbsolutePath = nacm.EntityData.SegmentPath
    nacm.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    nacm.EntityData.NamespaceTable = ietf.GetNamespaces()
    nacm.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    nacm.EntityData.Children = types.NewOrderedMap()
    nacm.EntityData.Children.Append("groups", types.YChild{"Groups", &nacm.Groups})
    nacm.EntityData.Children.Append("rule-list", types.YChild{"RuleList", nil})
    for i := range nacm.RuleList {
        nacm.EntityData.Children.Append(types.GetSegmentPath(nacm.RuleList[i]), types.YChild{"RuleList", nacm.RuleList[i]})
    }
    nacm.EntityData.Leafs = types.NewOrderedMap()
    nacm.EntityData.Leafs.Append("enable-nacm", types.YLeaf{"EnableNacm", nacm.EnableNacm})
    nacm.EntityData.Leafs.Append("read-default", types.YLeaf{"ReadDefault", nacm.ReadDefault})
    nacm.EntityData.Leafs.Append("write-default", types.YLeaf{"WriteDefault", nacm.WriteDefault})
    nacm.EntityData.Leafs.Append("exec-default", types.YLeaf{"ExecDefault", nacm.ExecDefault})
    nacm.EntityData.Leafs.Append("enable-external-groups", types.YLeaf{"EnableExternalGroups", nacm.EnableExternalGroups})
    nacm.EntityData.Leafs.Append("denied-operations", types.YLeaf{"DeniedOperations", nacm.DeniedOperations})
    nacm.EntityData.Leafs.Append("denied-data-writes", types.YLeaf{"DeniedDataWrites", nacm.DeniedDataWrites})
    nacm.EntityData.Leafs.Append("denied-notifications", types.YLeaf{"DeniedNotifications", nacm.DeniedNotifications})

    nacm.EntityData.YListKeys = []string {}

    return &(nacm.EntityData)
}

// Nacm_Groups
// NETCONF Access Control Groups.
type Nacm_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One NACM Group Entry.  This list will only contain configured entries, not
    // any entries learned from any transport protocols. The type is slice of
    // Nacm_Groups_Group.
    Group []*Nacm_Groups_Group
}

func (groups *Nacm_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "ietf"
    groups.EntityData.ParentYangName = "nacm"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.AbsolutePath = "ietf-netconf-acm:nacm/" + groups.EntityData.SegmentPath
    groups.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    groups.EntityData.NamespaceTable = ietf.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Nacm_Groups_Group
// One NACM Group Entry.  This list will only contain
// configured entries, not any entries learned from
// any transport protocols.
type Nacm_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Group name associated with this entry. The type is
    // string with length: 1..18446744073709551615.
    Name interface{}

    // Each entry identifies the username of a member of the group associated with
    // this entry. The type is slice of string with length:
    // 1..18446744073709551615.
    UserName []interface{}
}

func (group *Nacm_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "ietf"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.Name, "name")
    group.EntityData.AbsolutePath = "ietf-netconf-acm:nacm/groups/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    group.EntityData.NamespaceTable = ietf.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("name", types.YLeaf{"Name", group.Name})
    group.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", group.UserName})

    group.EntityData.YListKeys = []string {"Name"}

    return &(group.EntityData)
}

// Nacm_RuleList
// An ordered collection of access control rules.
type Nacm_RuleList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Arbitrary name assigned to the rule-list. The type
    // is string with length: 1..18446744073709551615.
    Name interface{}

    // List of administrative groups that will be assigned the associated access
    // rights defined by the 'rule' list.  The string '*' indicates that all
    // groups apply to the entry. The type is one of the following types: slice of
    // string with pattern: \*, or slice of string with length:
    // 1..18446744073709551615.
    Group []interface{}

    // One access control rule.  Rules are processed in user-defined order until a
    // match is found.  A rule matches if 'module-name', 'rule-type', and
    // 'access-operations' match the request.  If a rule matches, the 'action'
    // leaf determines if access is granted or not. The type is slice of
    // Nacm_RuleList_Rule.
    Rule []*Nacm_RuleList_Rule
}

func (ruleList *Nacm_RuleList) GetEntityData() *types.CommonEntityData {
    ruleList.EntityData.YFilter = ruleList.YFilter
    ruleList.EntityData.YangName = "rule-list"
    ruleList.EntityData.BundleName = "ietf"
    ruleList.EntityData.ParentYangName = "nacm"
    ruleList.EntityData.SegmentPath = "rule-list" + types.AddKeyToken(ruleList.Name, "name")
    ruleList.EntityData.AbsolutePath = "ietf-netconf-acm:nacm/" + ruleList.EntityData.SegmentPath
    ruleList.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    ruleList.EntityData.NamespaceTable = ietf.GetNamespaces()
    ruleList.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    ruleList.EntityData.Children = types.NewOrderedMap()
    ruleList.EntityData.Children.Append("rule", types.YChild{"Rule", nil})
    for i := range ruleList.Rule {
        ruleList.EntityData.Children.Append(types.GetSegmentPath(ruleList.Rule[i]), types.YChild{"Rule", ruleList.Rule[i]})
    }
    ruleList.EntityData.Leafs = types.NewOrderedMap()
    ruleList.EntityData.Leafs.Append("name", types.YLeaf{"Name", ruleList.Name})
    ruleList.EntityData.Leafs.Append("group", types.YLeaf{"Group", ruleList.Group})

    ruleList.EntityData.YListKeys = []string {"Name"}

    return &(ruleList.EntityData)
}

// Nacm_RuleList_Rule
// One access control rule.
// 
// Rules are processed in user-defined order until a match is
// found.  A rule matches if 'module-name', 'rule-type', and
// 'access-operations' match the request.  If a rule
// matches, the 'action' leaf determines if access is granted
// or not.
type Nacm_RuleList_Rule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Arbitrary name assigned to the rule. The type is
    // string with length: 1..18446744073709551615.
    Name interface{}

    // Name of the module associated with this rule.  This leaf matches if it has
    // the value '*' or if the object being accessed is defined in the module with
    // the specified module name. The type is one of the following types: string
    // with pattern: \* The default value is *., or string The default value is
    // *..
    ModuleName interface{}

    // This leaf matches if it has the value '*' or if its value equals the
    // requested protocol operation name. The type is one of the following types:
    // string with pattern: \*, or string.
    RpcName interface{}

    // This leaf matches if it has the value '*' or if its value equals the
    // requested notification name. The type is one of the following types: string
    // with pattern: \*, or string.
    NotificationName interface{}

    // Data Node Instance Identifier associated with the data node controlled by
    // this rule.  Configuration data or state data instance identifiers start
    // with a top-level data node.  A complete instance identifier is required for
    // this type of path value.  The special value '/' refers to all possible
    // datastore contents. The type is string. This attribute is mandatory.
    Path interface{}

    // Access operations associated with this rule.  This leaf matches if it has
    // the value '*' or if the bit corresponding to the requested operation is
    // set. The type is one of the following types: string with pattern: \* The
    // default value is *., or :go:struct:`Bits <ydk/models/netconf_acm/Bits>` The
    // default value is *..
    AccessOperations interface{}

    // The access control action associated with the rule.  If a rule is
    // determined to match a particular request, then this object is used to
    // determine whether to permit or deny the request. The type is ActionType.
    // This attribute is mandatory.
    Action interface{}

    // A textual description of the access rule. The type is string.
    Comment interface{}
}

func (rule *Nacm_RuleList_Rule) GetEntityData() *types.CommonEntityData {
    rule.EntityData.YFilter = rule.YFilter
    rule.EntityData.YangName = "rule"
    rule.EntityData.BundleName = "ietf"
    rule.EntityData.ParentYangName = "rule-list"
    rule.EntityData.SegmentPath = "rule" + types.AddKeyToken(rule.Name, "name")
    rule.EntityData.AbsolutePath = "ietf-netconf-acm:nacm/rule-list/" + rule.EntityData.SegmentPath
    rule.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    rule.EntityData.NamespaceTable = ietf.GetNamespaces()
    rule.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    rule.EntityData.Children = types.NewOrderedMap()
    rule.EntityData.Leafs = types.NewOrderedMap()
    rule.EntityData.Leafs.Append("name", types.YLeaf{"Name", rule.Name})
    rule.EntityData.Leafs.Append("module-name", types.YLeaf{"ModuleName", rule.ModuleName})
    rule.EntityData.Leafs.Append("rpc-name", types.YLeaf{"RpcName", rule.RpcName})
    rule.EntityData.Leafs.Append("notification-name", types.YLeaf{"NotificationName", rule.NotificationName})
    rule.EntityData.Leafs.Append("path", types.YLeaf{"Path", rule.Path})
    rule.EntityData.Leafs.Append("access-operations", types.YLeaf{"AccessOperations", rule.AccessOperations})
    rule.EntityData.Leafs.Append("action", types.YLeaf{"Action", rule.Action})
    rule.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", rule.Comment})

    rule.EntityData.YListKeys = []string {"Name"}

    return &(rule.EntityData)
}

