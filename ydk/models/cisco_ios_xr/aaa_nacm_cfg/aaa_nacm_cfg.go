// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-nacm package configuration.
// 
// This module contains definitions
// for the following management objects:
//   nacm: Parameters for NETCONF Access Control Model
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package aaa_nacm_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_nacm_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-aaa-nacm-cfg nacm}", reflect.TypeOf(Nacm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-aaa-nacm-cfg:nacm", reflect.TypeOf(Nacm{}))
}

// NacmAction represents Nacm action
type NacmAction string

const (
    // Permit
    NacmAction_permit NacmAction = "permit"

    // Deny
    NacmAction_deny NacmAction = "deny"
)

// NacmRule represents Nacm rule
type NacmRule string

const (
    // Protocoloperation
    NacmRule_protocol_operation NacmRule = "protocol-operation"

    // Datanode
    NacmRule_data_node NacmRule = "data-node"

    // Notification
    NacmRule_notification NacmRule = "notification"
)

// Nacm
// Parameters for NETCONF Access Control Model
type Nacm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enables or Disables all NETCONF access control enforcement. The type is
    // bool.
    EnableNacm interface{}

    // Controls write access if no appropriate rule is found. The type is
    // NacmAction.
    WriteDefault interface{}

    // Controls exec access if no appropriate rule is found. The type is
    // NacmAction.
    ExecDefault interface{}

    // Controls whether the server uses the groups reported by NETCONF transport
    // layer. The type is bool.
    EnableExternalGroups interface{}

    // Controls read access if no appropriate rule is found. The type is
    // NacmAction.
    ReadDefault interface{}

    // NETCONF Access Control Groups.
    Groups Nacm_Groups

    // Contains all rule lists of NACM.
    RulelistClasses Nacm_RulelistClasses
}

func (nacm *Nacm) GetEntityData() *types.CommonEntityData {
    nacm.EntityData.YFilter = nacm.YFilter
    nacm.EntityData.YangName = "nacm"
    nacm.EntityData.BundleName = "cisco_ios_xr"
    nacm.EntityData.ParentYangName = "Cisco-IOS-XR-aaa-nacm-cfg"
    nacm.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-nacm-cfg:nacm"
    nacm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nacm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nacm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nacm.EntityData.Children = types.NewOrderedMap()
    nacm.EntityData.Children.Append("groups", types.YChild{"Groups", &nacm.Groups})
    nacm.EntityData.Children.Append("rulelist-classes", types.YChild{"RulelistClasses", &nacm.RulelistClasses})
    nacm.EntityData.Leafs = types.NewOrderedMap()
    nacm.EntityData.Leafs.Append("enable-nacm", types.YLeaf{"EnableNacm", nacm.EnableNacm})
    nacm.EntityData.Leafs.Append("write-default", types.YLeaf{"WriteDefault", nacm.WriteDefault})
    nacm.EntityData.Leafs.Append("exec-default", types.YLeaf{"ExecDefault", nacm.ExecDefault})
    nacm.EntityData.Leafs.Append("enable-external-groups", types.YLeaf{"EnableExternalGroups", nacm.EnableExternalGroups})
    nacm.EntityData.Leafs.Append("read-default", types.YLeaf{"ReadDefault", nacm.ReadDefault})

    nacm.EntityData.YListKeys = []string {}

    return &(nacm.EntityData)
}

// Nacm_Groups
// NETCONF Access Control Groups
type Nacm_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One NACM Group Entry. The type is slice of Nacm_Groups_Group.
    Group []*Nacm_Groups_Group
}

func (groups *Nacm_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "nacm"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

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
// One NACM Group Entry
type Nacm_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. User group name. The type is string with length:
    // 1..63.
    GroupName interface{}

    // User name. The type is slice of string with length: 1..63.
    UserName []interface{}
}

func (group *Nacm_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.GroupName, "group-name")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", group.GroupName})
    group.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", group.UserName})

    group.EntityData.YListKeys = []string {"GroupName"}

    return &(group.EntityData)
}

// Nacm_RulelistClasses
// Contains all rule lists of NACM
type Nacm_RulelistClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each rule list of NACM. The type is slice of
    // Nacm_RulelistClasses_RulelistClass.
    RulelistClass []*Nacm_RulelistClasses_RulelistClass
}

func (rulelistClasses *Nacm_RulelistClasses) GetEntityData() *types.CommonEntityData {
    rulelistClasses.EntityData.YFilter = rulelistClasses.YFilter
    rulelistClasses.EntityData.YangName = "rulelist-classes"
    rulelistClasses.EntityData.BundleName = "cisco_ios_xr"
    rulelistClasses.EntityData.ParentYangName = "nacm"
    rulelistClasses.EntityData.SegmentPath = "rulelist-classes"
    rulelistClasses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rulelistClasses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rulelistClasses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rulelistClasses.EntityData.Children = types.NewOrderedMap()
    rulelistClasses.EntityData.Children.Append("rulelist-class", types.YChild{"RulelistClass", nil})
    for i := range rulelistClasses.RulelistClass {
        rulelistClasses.EntityData.Children.Append(types.GetSegmentPath(rulelistClasses.RulelistClass[i]), types.YChild{"RulelistClass", rulelistClasses.RulelistClass[i]})
    }
    rulelistClasses.EntityData.Leafs = types.NewOrderedMap()

    rulelistClasses.EntityData.YListKeys = []string {}

    return &(rulelistClasses.EntityData)
}

// Nacm_RulelistClasses_RulelistClass
// Each rule list of NACM
type Nacm_RulelistClasses_RulelistClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the rulelists in the order of
    // precedence. The type is string with length: 1..15.
    OrderingIndex interface{}

    // This attribute is a key. Rulelist key name. The type is string with length:
    // 1..63.
    RulelistName interface{}

    // List of groups that will be assigned with the rule.
    GroupNames Nacm_RulelistClasses_RulelistClass_GroupNames

    // Set of rules in a rulelist.
    Rules Nacm_RulelistClasses_RulelistClass_Rules
}

func (rulelistClass *Nacm_RulelistClasses_RulelistClass) GetEntityData() *types.CommonEntityData {
    rulelistClass.EntityData.YFilter = rulelistClass.YFilter
    rulelistClass.EntityData.YangName = "rulelist-class"
    rulelistClass.EntityData.BundleName = "cisco_ios_xr"
    rulelistClass.EntityData.ParentYangName = "rulelist-classes"
    rulelistClass.EntityData.SegmentPath = "rulelist-class" + types.AddKeyToken(rulelistClass.OrderingIndex, "ordering-index") + types.AddKeyToken(rulelistClass.RulelistName, "rulelist-name")
    rulelistClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rulelistClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rulelistClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rulelistClass.EntityData.Children = types.NewOrderedMap()
    rulelistClass.EntityData.Children.Append("group-names", types.YChild{"GroupNames", &rulelistClass.GroupNames})
    rulelistClass.EntityData.Children.Append("rules", types.YChild{"Rules", &rulelistClass.Rules})
    rulelistClass.EntityData.Leafs = types.NewOrderedMap()
    rulelistClass.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", rulelistClass.OrderingIndex})
    rulelistClass.EntityData.Leafs.Append("rulelist-name", types.YLeaf{"RulelistName", rulelistClass.RulelistName})

    rulelistClass.EntityData.YListKeys = []string {"OrderingIndex", "RulelistName"}

    return &(rulelistClass.EntityData)
}

// Nacm_RulelistClasses_RulelistClass_GroupNames
// List of groups that will be assigned with the
// rule
type Nacm_RulelistClasses_RulelistClass_GroupNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group name. The type is slice of string with length: 1..63.
    GroupName []interface{}
}

func (groupNames *Nacm_RulelistClasses_RulelistClass_GroupNames) GetEntityData() *types.CommonEntityData {
    groupNames.EntityData.YFilter = groupNames.YFilter
    groupNames.EntityData.YangName = "group-names"
    groupNames.EntityData.BundleName = "cisco_ios_xr"
    groupNames.EntityData.ParentYangName = "rulelist-class"
    groupNames.EntityData.SegmentPath = "group-names"
    groupNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupNames.EntityData.Children = types.NewOrderedMap()
    groupNames.EntityData.Leafs = types.NewOrderedMap()
    groupNames.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", groupNames.GroupName})

    groupNames.EntityData.YListKeys = []string {}

    return &(groupNames.EntityData)
}

// Nacm_RulelistClasses_RulelistClass_Rules
// Set of rules in a rulelist
type Nacm_RulelistClasses_RulelistClass_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each rule in a rulelist. The type is slice of
    // Nacm_RulelistClasses_RulelistClass_Rules_Rule.
    Rule []*Nacm_RulelistClasses_RulelistClass_Rules_Rule
}

func (rules *Nacm_RulelistClasses_RulelistClass_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "rulelist-class"
    rules.EntityData.SegmentPath = "rules"
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = types.NewOrderedMap()
    rules.EntityData.Children.Append("rule", types.YChild{"Rule", nil})
    for i := range rules.Rule {
        rules.EntityData.Children.Append(types.GetSegmentPath(rules.Rule[i]), types.YChild{"Rule", rules.Rule[i]})
    }
    rules.EntityData.Leafs = types.NewOrderedMap()

    rules.EntityData.YListKeys = []string {}

    return &(rules.EntityData)
}

// Nacm_RulelistClasses_RulelistClass_Rules_Rule
// Each rule in a rulelist
type Nacm_RulelistClasses_RulelistClass_Rules_Rule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This is used to sort the rules in the order of
    // precedence. The type is string with length: 1..15.
    OrderingIndex interface{}

    // This attribute is a key. Rule name. The type is string with length: 1..63.
    RuleName interface{}

    // Name of the module associated with this rule. The type is string with
    // length: 1..63.
    ModuleName interface{}

    // The access control action associated with the rule. The type is NacmAction.
    // This attribute is mandatory.
    Action interface{}

    // Textual description of the access rule. The type is string with length:
    // 1..255.
    Comment interface{}

    // Rule Type associated with this rule.
    RuleType Nacm_RulelistClasses_RulelistClass_Rules_Rule_RuleType

    // Access operations associated with this rule.
    AccessOperations Nacm_RulelistClasses_RulelistClass_Rules_Rule_AccessOperations
}

func (rule *Nacm_RulelistClasses_RulelistClass_Rules_Rule) GetEntityData() *types.CommonEntityData {
    rule.EntityData.YFilter = rule.YFilter
    rule.EntityData.YangName = "rule"
    rule.EntityData.BundleName = "cisco_ios_xr"
    rule.EntityData.ParentYangName = "rules"
    rule.EntityData.SegmentPath = "rule" + types.AddKeyToken(rule.OrderingIndex, "ordering-index") + types.AddKeyToken(rule.RuleName, "rule-name")
    rule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rule.EntityData.Children = types.NewOrderedMap()
    rule.EntityData.Children.Append("rule-type", types.YChild{"RuleType", &rule.RuleType})
    rule.EntityData.Children.Append("access-operations", types.YChild{"AccessOperations", &rule.AccessOperations})
    rule.EntityData.Leafs = types.NewOrderedMap()
    rule.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", rule.OrderingIndex})
    rule.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", rule.RuleName})
    rule.EntityData.Leafs.Append("module-name", types.YLeaf{"ModuleName", rule.ModuleName})
    rule.EntityData.Leafs.Append("action", types.YLeaf{"Action", rule.Action})
    rule.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", rule.Comment})

    rule.EntityData.YListKeys = []string {"OrderingIndex", "RuleName"}

    return &(rule.EntityData)
}

// Nacm_RulelistClasses_RulelistClass_Rules_Rule_RuleType
// Rule Type associated with this rule
type Nacm_RulelistClasses_RulelistClass_Rules_Rule_RuleType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rule Type. The type is NacmRule.
    Type interface{}

    // Rule Value. The type is string with length: 1..511.
    Value interface{}
}

func (ruleType *Nacm_RulelistClasses_RulelistClass_Rules_Rule_RuleType) GetEntityData() *types.CommonEntityData {
    ruleType.EntityData.YFilter = ruleType.YFilter
    ruleType.EntityData.YangName = "rule-type"
    ruleType.EntityData.BundleName = "cisco_ios_xr"
    ruleType.EntityData.ParentYangName = "rule"
    ruleType.EntityData.SegmentPath = "rule-type"
    ruleType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleType.EntityData.Children = types.NewOrderedMap()
    ruleType.EntityData.Leafs = types.NewOrderedMap()
    ruleType.EntityData.Leafs.Append("type", types.YLeaf{"Type", ruleType.Type})
    ruleType.EntityData.Leafs.Append("value", types.YLeaf{"Value", ruleType.Value})

    ruleType.EntityData.YListKeys = []string {}

    return &(ruleType.EntityData)
}

// Nacm_RulelistClasses_RulelistClass_Rules_Rule_AccessOperations
// Access operations associated with this rule
type Nacm_RulelistClasses_RulelistClass_Rules_Rule_AccessOperations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Create. The type is interface{} with range: 0..4294967295.
    Create interface{}

    // Enable Read. The type is interface{} with range: 0..4294967295.
    Read interface{}

    // Enable Update. The type is interface{} with range: 0..4294967295.
    Update interface{}

    // Enable Delete. The type is interface{} with range: 0..4294967295.
    Delete interface{}

    // Enable Exec. The type is interface{} with range: 0..4294967295.
    Exec interface{}

    // Enable All permissions. The type is interface{} with range: 0..4294967295.
    All interface{}
}

func (accessOperations *Nacm_RulelistClasses_RulelistClass_Rules_Rule_AccessOperations) GetEntityData() *types.CommonEntityData {
    accessOperations.EntityData.YFilter = accessOperations.YFilter
    accessOperations.EntityData.YangName = "access-operations"
    accessOperations.EntityData.BundleName = "cisco_ios_xr"
    accessOperations.EntityData.ParentYangName = "rule"
    accessOperations.EntityData.SegmentPath = "access-operations"
    accessOperations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessOperations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessOperations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessOperations.EntityData.Children = types.NewOrderedMap()
    accessOperations.EntityData.Leafs = types.NewOrderedMap()
    accessOperations.EntityData.Leafs.Append("create", types.YLeaf{"Create", accessOperations.Create})
    accessOperations.EntityData.Leafs.Append("read", types.YLeaf{"Read", accessOperations.Read})
    accessOperations.EntityData.Leafs.Append("update", types.YLeaf{"Update", accessOperations.Update})
    accessOperations.EntityData.Leafs.Append("delete", types.YLeaf{"Delete", accessOperations.Delete})
    accessOperations.EntityData.Leafs.Append("exec", types.YLeaf{"Exec", accessOperations.Exec})
    accessOperations.EntityData.Leafs.Append("all", types.YLeaf{"All", accessOperations.All})

    accessOperations.EntityData.YListKeys = []string {}

    return &(accessOperations.EntityData)
}

