// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-nacm package operational data.
// 
// This module contains definitions
// for the following management objects:
//   aaa-nacm: AAA Nacm Information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package aaa_nacm_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_nacm_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-aaa-nacm-oper aaa-nacm}", reflect.TypeOf(AaaNacm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-aaa-nacm-oper:aaa-nacm", reflect.TypeOf(AaaNacm{}))
}

// AaaNacm
// AAA Nacm Information
type AaaNacm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA NACM summary.
    Counters AaaNacm_Counters

    // AAA NACM User summary.
    Users AaaNacm_Users

    // AAA NACM summary.
    Summary AaaNacm_Summary

    // AAA NACM Rulelist summary.
    Rules AaaNacm_Rules

    // AAA NACM Group summary.
    Groups AaaNacm_Groups
}

func (aaaNacm *AaaNacm) GetEntityData() *types.CommonEntityData {
    aaaNacm.EntityData.YFilter = aaaNacm.YFilter
    aaaNacm.EntityData.YangName = "aaa-nacm"
    aaaNacm.EntityData.BundleName = "cisco_ios_xr"
    aaaNacm.EntityData.ParentYangName = "Cisco-IOS-XR-aaa-nacm-oper"
    aaaNacm.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-nacm-oper:aaa-nacm"
    aaaNacm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaaNacm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaaNacm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaaNacm.EntityData.Children = types.NewOrderedMap()
    aaaNacm.EntityData.Children.Append("counters", types.YChild{"Counters", &aaaNacm.Counters})
    aaaNacm.EntityData.Children.Append("users", types.YChild{"Users", &aaaNacm.Users})
    aaaNacm.EntityData.Children.Append("summary", types.YChild{"Summary", &aaaNacm.Summary})
    aaaNacm.EntityData.Children.Append("rules", types.YChild{"Rules", &aaaNacm.Rules})
    aaaNacm.EntityData.Children.Append("groups", types.YChild{"Groups", &aaaNacm.Groups})
    aaaNacm.EntityData.Leafs = types.NewOrderedMap()

    aaaNacm.EntityData.YListKeys = []string {}

    return &(aaaNacm.EntityData)
}

// AaaNacm_Counters
// AAA NACM summary
type AaaNacm_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Denied Operations. The type is interface{} with range:
    // -2147483648..2147483647.
    DeniedOperations interface{}

    // Denied Data Writes. The type is interface{} with range:
    // -2147483648..2147483647.
    DeniedDataWrites interface{}

    // Denied Notifications. The type is interface{} with range:
    // -2147483648..2147483647.
    DeniedNotifications interface{}
}

func (counters *AaaNacm_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "aaa-nacm"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Leafs = types.NewOrderedMap()
    counters.EntityData.Leafs.Append("denied-operations", types.YLeaf{"DeniedOperations", counters.DeniedOperations})
    counters.EntityData.Leafs.Append("denied-data-writes", types.YLeaf{"DeniedDataWrites", counters.DeniedDataWrites})
    counters.EntityData.Leafs.Append("denied-notifications", types.YLeaf{"DeniedNotifications", counters.DeniedNotifications})

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// AaaNacm_Users
// AAA NACM User summary
type AaaNacm_Users struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA NACM User detail. The type is slice of AaaNacm_Users_User.
    User []*AaaNacm_Users_User
}

func (users *AaaNacm_Users) GetEntityData() *types.CommonEntityData {
    users.EntityData.YFilter = users.YFilter
    users.EntityData.YangName = "users"
    users.EntityData.BundleName = "cisco_ios_xr"
    users.EntityData.ParentYangName = "aaa-nacm"
    users.EntityData.SegmentPath = "users"
    users.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    users.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    users.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    users.EntityData.Children = types.NewOrderedMap()
    users.EntityData.Children.Append("user", types.YChild{"User", nil})
    for i := range users.User {
        users.EntityData.Children.Append(types.GetSegmentPath(users.User[i]), types.YChild{"User", users.User[i]})
    }
    users.EntityData.Leafs = types.NewOrderedMap()

    users.EntityData.YListKeys = []string {}

    return &(users.EntityData)
}

// AaaNacm_Users_User
// AAA NACM User detail
type AaaNacm_Users_User struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. User. The type is string with length: 1..256.
    User interface{}

    // User Name. The type is string with length: 0..256.
    UserName interface{}

    // Group Name List. The type is slice of AaaNacm_Users_User_GroupName.
    GroupName []*AaaNacm_Users_User_GroupName
}

func (user *AaaNacm_Users_User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xr"
    user.EntityData.ParentYangName = "users"
    user.EntityData.SegmentPath = "user" + types.AddKeyToken(user.User, "user")
    user.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    user.EntityData.Children = types.NewOrderedMap()
    user.EntityData.Children.Append("group-name", types.YChild{"GroupName", nil})
    for i := range user.GroupName {
        user.EntityData.Children.Append(types.GetSegmentPath(user.GroupName[i]), types.YChild{"GroupName", user.GroupName[i]})
    }
    user.EntityData.Leafs = types.NewOrderedMap()
    user.EntityData.Leafs.Append("user", types.YLeaf{"User", user.User})
    user.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", user.UserName})

    user.EntityData.YListKeys = []string {"User"}

    return &(user.EntityData)
}

// AaaNacm_Users_User_GroupName
// Group Name List
type AaaNacm_Users_User_GroupName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is string with length: 0..256.
    Name interface{}
}

func (groupName *AaaNacm_Users_User_GroupName) GetEntityData() *types.CommonEntityData {
    groupName.EntityData.YFilter = groupName.YFilter
    groupName.EntityData.YangName = "group-name"
    groupName.EntityData.BundleName = "cisco_ios_xr"
    groupName.EntityData.ParentYangName = "user"
    groupName.EntityData.SegmentPath = "group-name"
    groupName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groupName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groupName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groupName.EntityData.Children = types.NewOrderedMap()
    groupName.EntityData.Leafs = types.NewOrderedMap()
    groupName.EntityData.Leafs.Append("name", types.YLeaf{"Name", groupName.Name})

    groupName.EntityData.YListKeys = []string {}

    return &(groupName.EntityData)
}

// AaaNacm_Summary
// AAA NACM summary
type AaaNacm_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Groups. The type is interface{} with range: -2147483648..2147483647.
    Groups interface{}

    // Users. The type is interface{} with range: -2147483648..2147483647.
    Users interface{}

    // Rulelist. The type is interface{} with range: -2147483648..2147483647.
    Rulelist interface{}

    // Rules. The type is interface{} with range: -2147483648..2147483647.
    Rules interface{}

    // Read Default. The type is string with length: 0..16.
    ReadDefault interface{}

    // Write Default. The type is string with length: 0..16.
    WriteDefault interface{}

    // Exec Default. The type is string with length: 0..16.
    ExecDefault interface{}

    // Enable Nacm. The type is string with length: 0..16.
    EnableNacm interface{}

    // Enable External Groups. The type is string with length: 0..16.
    EnableExternalGroups interface{}
}

func (summary *AaaNacm_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "aaa-nacm"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("groups", types.YLeaf{"Groups", summary.Groups})
    summary.EntityData.Leafs.Append("users", types.YLeaf{"Users", summary.Users})
    summary.EntityData.Leafs.Append("rulelist", types.YLeaf{"Rulelist", summary.Rulelist})
    summary.EntityData.Leafs.Append("rules", types.YLeaf{"Rules", summary.Rules})
    summary.EntityData.Leafs.Append("read-default", types.YLeaf{"ReadDefault", summary.ReadDefault})
    summary.EntityData.Leafs.Append("write-default", types.YLeaf{"WriteDefault", summary.WriteDefault})
    summary.EntityData.Leafs.Append("exec-default", types.YLeaf{"ExecDefault", summary.ExecDefault})
    summary.EntityData.Leafs.Append("enable-nacm", types.YLeaf{"EnableNacm", summary.EnableNacm})
    summary.EntityData.Leafs.Append("enable-external-groups", types.YLeaf{"EnableExternalGroups", summary.EnableExternalGroups})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// AaaNacm_Rules
// AAA NACM Rulelist summary
type AaaNacm_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA NACM Rulelist detail. The type is slice of AaaNacm_Rules_Rule.
    Rule []*AaaNacm_Rules_Rule
}

func (rules *AaaNacm_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "aaa-nacm"
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

// AaaNacm_Rules_Rule
// AAA NACM Rulelist detail
type AaaNacm_Rules_Rule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rulelist Ordering Index. The type is string with length: 1..256.
    OrderingIndex interface{}

    // Rulelist Name. The type is string with length: 1..256.
    RulelistName interface{}

    // AAA NACM Rulelist detail.
    RulelistRules AaaNacm_Rules_Rule_RulelistRules
}

func (rule *AaaNacm_Rules_Rule) GetEntityData() *types.CommonEntityData {
    rule.EntityData.YFilter = rule.YFilter
    rule.EntityData.YangName = "rule"
    rule.EntityData.BundleName = "cisco_ios_xr"
    rule.EntityData.ParentYangName = "rules"
    rule.EntityData.SegmentPath = "rule"
    rule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rule.EntityData.Children = types.NewOrderedMap()
    rule.EntityData.Children.Append("rulelist-rules", types.YChild{"RulelistRules", &rule.RulelistRules})
    rule.EntityData.Leafs = types.NewOrderedMap()
    rule.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", rule.OrderingIndex})
    rule.EntityData.Leafs.Append("rulelist-name", types.YLeaf{"RulelistName", rule.RulelistName})

    rule.EntityData.YListKeys = []string {}

    return &(rule.EntityData)
}

// AaaNacm_Rules_Rule_RulelistRules
// AAA NACM Rulelist detail
type AaaNacm_Rules_Rule_RulelistRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA NACM Rulelist detail. The type is slice of
    // AaaNacm_Rules_Rule_RulelistRules_RulelistRule.
    RulelistRule []*AaaNacm_Rules_Rule_RulelistRules_RulelistRule
}

func (rulelistRules *AaaNacm_Rules_Rule_RulelistRules) GetEntityData() *types.CommonEntityData {
    rulelistRules.EntityData.YFilter = rulelistRules.YFilter
    rulelistRules.EntityData.YangName = "rulelist-rules"
    rulelistRules.EntityData.BundleName = "cisco_ios_xr"
    rulelistRules.EntityData.ParentYangName = "rule"
    rulelistRules.EntityData.SegmentPath = "rulelist-rules"
    rulelistRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rulelistRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rulelistRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rulelistRules.EntityData.Children = types.NewOrderedMap()
    rulelistRules.EntityData.Children.Append("rulelist-rule", types.YChild{"RulelistRule", nil})
    for i := range rulelistRules.RulelistRule {
        rulelistRules.EntityData.Children.Append(types.GetSegmentPath(rulelistRules.RulelistRule[i]), types.YChild{"RulelistRule", rulelistRules.RulelistRule[i]})
    }
    rulelistRules.EntityData.Leafs = types.NewOrderedMap()

    rulelistRules.EntityData.YListKeys = []string {}

    return &(rulelistRules.EntityData)
}

// AaaNacm_Rules_Rule_RulelistRules_RulelistRule
// AAA NACM Rulelist detail
type AaaNacm_Rules_Rule_RulelistRules_RulelistRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rule. The type is string with length: 1..256.
    Rule interface{}

    // Rule Name. The type is string with length: 0..512.
    RuleName interface{}

    // Rule Index. The type is string with length: 0..16.
    RuleIndex interface{}

    // Rulelist Index. The type is string with length: 0..16.
    RulelistIndex interface{}

    // Module Name. The type is string with length: 0..512.
    ModuleName interface{}

    // Action. The type is string with length: 0..256.
    Action interface{}

    // Rule Type. The type is string with length: 0..256.
    RuleType interface{}

    // Comment. The type is string with length: 0..256.
    Comment interface{}

    // Access Operations. The type is string with length: 0..256.
    AccessOperations interface{}

    // Rule Value. The type is string with length: 0..512.
    RuleValue interface{}

    // Hit Count. The type is interface{} with range: -2147483648..2147483647.
    HitCount interface{}
}

func (rulelistRule *AaaNacm_Rules_Rule_RulelistRules_RulelistRule) GetEntityData() *types.CommonEntityData {
    rulelistRule.EntityData.YFilter = rulelistRule.YFilter
    rulelistRule.EntityData.YangName = "rulelist-rule"
    rulelistRule.EntityData.BundleName = "cisco_ios_xr"
    rulelistRule.EntityData.ParentYangName = "rulelist-rules"
    rulelistRule.EntityData.SegmentPath = "rulelist-rule" + types.AddKeyToken(rulelistRule.Rule, "rule")
    rulelistRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rulelistRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rulelistRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rulelistRule.EntityData.Children = types.NewOrderedMap()
    rulelistRule.EntityData.Leafs = types.NewOrderedMap()
    rulelistRule.EntityData.Leafs.Append("rule", types.YLeaf{"Rule", rulelistRule.Rule})
    rulelistRule.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", rulelistRule.RuleName})
    rulelistRule.EntityData.Leafs.Append("rule-index", types.YLeaf{"RuleIndex", rulelistRule.RuleIndex})
    rulelistRule.EntityData.Leafs.Append("rulelist-index", types.YLeaf{"RulelistIndex", rulelistRule.RulelistIndex})
    rulelistRule.EntityData.Leafs.Append("module-name", types.YLeaf{"ModuleName", rulelistRule.ModuleName})
    rulelistRule.EntityData.Leafs.Append("action", types.YLeaf{"Action", rulelistRule.Action})
    rulelistRule.EntityData.Leafs.Append("rule-type", types.YLeaf{"RuleType", rulelistRule.RuleType})
    rulelistRule.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", rulelistRule.Comment})
    rulelistRule.EntityData.Leafs.Append("access-operations", types.YLeaf{"AccessOperations", rulelistRule.AccessOperations})
    rulelistRule.EntityData.Leafs.Append("rule-value", types.YLeaf{"RuleValue", rulelistRule.RuleValue})
    rulelistRule.EntityData.Leafs.Append("hit-count", types.YLeaf{"HitCount", rulelistRule.HitCount})

    rulelistRule.EntityData.YListKeys = []string {"Rule"}

    return &(rulelistRule.EntityData)
}

// AaaNacm_Groups
// AAA NACM Group summary
type AaaNacm_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA NACM Group detail. The type is slice of AaaNacm_Groups_Group.
    Group []*AaaNacm_Groups_Group
}

func (groups *AaaNacm_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "aaa-nacm"
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

// AaaNacm_Groups_Group
// AAA NACM Group detail
type AaaNacm_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group. The type is string with length: 1..256.
    Group interface{}

    // Group Name. The type is string with length: 0..256.
    GroupName interface{}

    // Users Name List. The type is slice of AaaNacm_Groups_Group_UserName.
    UserName []*AaaNacm_Groups_Group_UserName

    // Rules Name List. The type is slice of AaaNacm_Groups_Group_RuleName.
    RuleName []*AaaNacm_Groups_Group_RuleName
}

func (group *AaaNacm_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.Group, "group")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Children.Append("user-name", types.YChild{"UserName", nil})
    for i := range group.UserName {
        group.EntityData.Children.Append(types.GetSegmentPath(group.UserName[i]), types.YChild{"UserName", group.UserName[i]})
    }
    group.EntityData.Children.Append("rule-name", types.YChild{"RuleName", nil})
    for i := range group.RuleName {
        group.EntityData.Children.Append(types.GetSegmentPath(group.RuleName[i]), types.YChild{"RuleName", group.RuleName[i]})
    }
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("group", types.YLeaf{"Group", group.Group})
    group.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", group.GroupName})

    group.EntityData.YListKeys = []string {"Group"}

    return &(group.EntityData)
}

// AaaNacm_Groups_Group_UserName
// Users Name List
type AaaNacm_Groups_Group_UserName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is string with length: 0..256.
    Name interface{}
}

func (userName *AaaNacm_Groups_Group_UserName) GetEntityData() *types.CommonEntityData {
    userName.EntityData.YFilter = userName.YFilter
    userName.EntityData.YangName = "user-name"
    userName.EntityData.BundleName = "cisco_ios_xr"
    userName.EntityData.ParentYangName = "group"
    userName.EntityData.SegmentPath = "user-name"
    userName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userName.EntityData.Children = types.NewOrderedMap()
    userName.EntityData.Leafs = types.NewOrderedMap()
    userName.EntityData.Leafs.Append("name", types.YLeaf{"Name", userName.Name})

    userName.EntityData.YListKeys = []string {}

    return &(userName.EntityData)
}

// AaaNacm_Groups_Group_RuleName
// Rules Name List
type AaaNacm_Groups_Group_RuleName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name. The type is string with length: 0..256.
    Name interface{}
}

func (ruleName *AaaNacm_Groups_Group_RuleName) GetEntityData() *types.CommonEntityData {
    ruleName.EntityData.YFilter = ruleName.YFilter
    ruleName.EntityData.YangName = "rule-name"
    ruleName.EntityData.BundleName = "cisco_ios_xr"
    ruleName.EntityData.ParentYangName = "group"
    ruleName.EntityData.SegmentPath = "rule-name"
    ruleName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleName.EntityData.Children = types.NewOrderedMap()
    ruleName.EntityData.Leafs = types.NewOrderedMap()
    ruleName.EntityData.Leafs.Append("name", types.YLeaf{"Name", ruleName.Name})

    ruleName.EntityData.YListKeys = []string {}

    return &(ruleName.EntityData)
}

