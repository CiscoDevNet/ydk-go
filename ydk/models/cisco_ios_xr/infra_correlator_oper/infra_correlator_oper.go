// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-correlator package operational data.
// 
// This module contains definitions
// for the following management objects:
//   suppression: Suppression operational data
//   correlator: correlator
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_correlator_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_correlator_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-correlator-oper suppression}", reflect.TypeOf(Suppression{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-correlator-oper:suppression", reflect.TypeOf(Suppression{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-correlator-oper correlator}", reflect.TypeOf(Correlator{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-correlator-oper:correlator", reflect.TypeOf(Correlator{}))
}

// AlAlarmBistate represents Al alarm bistate
type AlAlarmBistate string

const (
    // not available
    AlAlarmBistate_not_available AlAlarmBistate = "not-available"

    // active
    AlAlarmBistate_active AlAlarmBistate = "active"

    // clear
    AlAlarmBistate_clear AlAlarmBistate = "clear"
)

// AlAlarmSeverity represents Al alarm severity
type AlAlarmSeverity string

const (
    // unknown
    AlAlarmSeverity_unknown AlAlarmSeverity = "unknown"

    // emergency
    AlAlarmSeverity_emergency AlAlarmSeverity = "emergency"

    // alert
    AlAlarmSeverity_alert AlAlarmSeverity = "alert"

    // critical
    AlAlarmSeverity_critical AlAlarmSeverity = "critical"

    // error
    AlAlarmSeverity_error_ AlAlarmSeverity = "error"

    // warning
    AlAlarmSeverity_warning AlAlarmSeverity = "warning"

    // notice
    AlAlarmSeverity_notice AlAlarmSeverity = "notice"

    // informational
    AlAlarmSeverity_informational AlAlarmSeverity = "informational"

    // debugging
    AlAlarmSeverity_debugging AlAlarmSeverity = "debugging"
)

// AcRuleState represents Ac rule state
type AcRuleState string

const (
    // Rule is in Unapplied state
    AcRuleState_rule_unapplied AcRuleState = "rule-unapplied"

    // Rule is Applied to specified RacksSlots,
    // Contexts and Sources
    AcRuleState_rule_applied AcRuleState = "rule-applied"

    // Rule is Applied to all of router
    AcRuleState_rule_applied_all AcRuleState = "rule-applied-all"
)

// Suppression
// Suppression operational data
type Suppression struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table that contains the database of suppression rule summary.
    RuleSummaries Suppression_RuleSummaries

    // Table that contains the database of suppression rule details.
    RuleDetails Suppression_RuleDetails
}

func (suppression *Suppression) GetEntityData() *types.CommonEntityData {
    suppression.EntityData.YFilter = suppression.YFilter
    suppression.EntityData.YangName = "suppression"
    suppression.EntityData.BundleName = "cisco_ios_xr"
    suppression.EntityData.ParentYangName = "Cisco-IOS-XR-infra-correlator-oper"
    suppression.EntityData.SegmentPath = "Cisco-IOS-XR-infra-correlator-oper:suppression"
    suppression.EntityData.AbsolutePath = suppression.EntityData.SegmentPath
    suppression.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppression.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppression.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppression.EntityData.Children = types.NewOrderedMap()
    suppression.EntityData.Children.Append("rule-summaries", types.YChild{"RuleSummaries", &suppression.RuleSummaries})
    suppression.EntityData.Children.Append("rule-details", types.YChild{"RuleDetails", &suppression.RuleDetails})
    suppression.EntityData.Leafs = types.NewOrderedMap()

    suppression.EntityData.YListKeys = []string {}

    return &(suppression.EntityData)
}

// Suppression_RuleSummaries
// Table that contains the database of suppression
// rule summary
type Suppression_RuleSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One of the suppression rules. The type is slice of
    // Suppression_RuleSummaries_RuleSummary.
    RuleSummary []*Suppression_RuleSummaries_RuleSummary
}

func (ruleSummaries *Suppression_RuleSummaries) GetEntityData() *types.CommonEntityData {
    ruleSummaries.EntityData.YFilter = ruleSummaries.YFilter
    ruleSummaries.EntityData.YangName = "rule-summaries"
    ruleSummaries.EntityData.BundleName = "cisco_ios_xr"
    ruleSummaries.EntityData.ParentYangName = "suppression"
    ruleSummaries.EntityData.SegmentPath = "rule-summaries"
    ruleSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:suppression/" + ruleSummaries.EntityData.SegmentPath
    ruleSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSummaries.EntityData.Children = types.NewOrderedMap()
    ruleSummaries.EntityData.Children.Append("rule-summary", types.YChild{"RuleSummary", nil})
    for i := range ruleSummaries.RuleSummary {
        ruleSummaries.EntityData.Children.Append(types.GetSegmentPath(ruleSummaries.RuleSummary[i]), types.YChild{"RuleSummary", ruleSummaries.RuleSummary[i]})
    }
    ruleSummaries.EntityData.Leafs = types.NewOrderedMap()

    ruleSummaries.EntityData.YListKeys = []string {}

    return &(ruleSummaries.EntityData)
}

// Suppression_RuleSummaries_RuleSummary
// One of the suppression rules
type Suppression_RuleSummaries_RuleSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Suppression Rule Name. The type is string with
    // length: 1..32.
    RuleName interface{}

    // Suppress Rule Name. The type is string.
    RuleNameXr interface{}

    // Applied state of the rule It could be not applied, applied or applied to
    // all. The type is AcRuleState.
    RuleState interface{}

    // Number of suppressed alarms associated with this rule. The type is
    // interface{} with range: 0..4294967295.
    SuppressedAlarmsCount interface{}
}

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetEntityData() *types.CommonEntityData {
    ruleSummary.EntityData.YFilter = ruleSummary.YFilter
    ruleSummary.EntityData.YangName = "rule-summary"
    ruleSummary.EntityData.BundleName = "cisco_ios_xr"
    ruleSummary.EntityData.ParentYangName = "rule-summaries"
    ruleSummary.EntityData.SegmentPath = "rule-summary" + types.AddKeyToken(ruleSummary.RuleName, "rule-name")
    ruleSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:suppression/rule-summaries/" + ruleSummary.EntityData.SegmentPath
    ruleSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSummary.EntityData.Children = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", ruleSummary.RuleName})
    ruleSummary.EntityData.Leafs.Append("rule-name-xr", types.YLeaf{"RuleNameXr", ruleSummary.RuleNameXr})
    ruleSummary.EntityData.Leafs.Append("rule-state", types.YLeaf{"RuleState", ruleSummary.RuleState})
    ruleSummary.EntityData.Leafs.Append("suppressed-alarms-count", types.YLeaf{"SuppressedAlarmsCount", ruleSummary.SuppressedAlarmsCount})

    ruleSummary.EntityData.YListKeys = []string {"RuleName"}

    return &(ruleSummary.EntityData)
}

// Suppression_RuleDetails
// Table that contains the database of suppression
// rule details
type Suppression_RuleDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Details of one of the suppression rules. The type is slice of
    // Suppression_RuleDetails_RuleDetail.
    RuleDetail []*Suppression_RuleDetails_RuleDetail
}

func (ruleDetails *Suppression_RuleDetails) GetEntityData() *types.CommonEntityData {
    ruleDetails.EntityData.YFilter = ruleDetails.YFilter
    ruleDetails.EntityData.YangName = "rule-details"
    ruleDetails.EntityData.BundleName = "cisco_ios_xr"
    ruleDetails.EntityData.ParentYangName = "suppression"
    ruleDetails.EntityData.SegmentPath = "rule-details"
    ruleDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:suppression/" + ruleDetails.EntityData.SegmentPath
    ruleDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleDetails.EntityData.Children = types.NewOrderedMap()
    ruleDetails.EntityData.Children.Append("rule-detail", types.YChild{"RuleDetail", nil})
    for i := range ruleDetails.RuleDetail {
        ruleDetails.EntityData.Children.Append(types.GetSegmentPath(ruleDetails.RuleDetail[i]), types.YChild{"RuleDetail", ruleDetails.RuleDetail[i]})
    }
    ruleDetails.EntityData.Leafs = types.NewOrderedMap()

    ruleDetails.EntityData.YListKeys = []string {}

    return &(ruleDetails.EntityData)
}

// Suppression_RuleDetails_RuleDetail
// Details of one of the suppression rules
type Suppression_RuleDetails_RuleDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Suppression Rule Name. The type is string with
    // length: 1..32.
    RuleName interface{}

    // Match any alarm. The type is bool.
    AllAlarms interface{}

    // Severity level to suppress. The type is AlAlarmSeverity.
    AlarmSeverity interface{}

    // Sources (R/S/M) to which the rule is applied. The type is slice of string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    ApplySource []interface{}

    // Rule summary, name, etc.
    RuleSummary Suppression_RuleDetails_RuleDetail_RuleSummary

    // Message codes defining the rule. The type is slice of
    // Suppression_RuleDetails_RuleDetail_Codes.
    Codes []*Suppression_RuleDetails_RuleDetail_Codes
}

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetEntityData() *types.CommonEntityData {
    ruleDetail.EntityData.YFilter = ruleDetail.YFilter
    ruleDetail.EntityData.YangName = "rule-detail"
    ruleDetail.EntityData.BundleName = "cisco_ios_xr"
    ruleDetail.EntityData.ParentYangName = "rule-details"
    ruleDetail.EntityData.SegmentPath = "rule-detail" + types.AddKeyToken(ruleDetail.RuleName, "rule-name")
    ruleDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:suppression/rule-details/" + ruleDetail.EntityData.SegmentPath
    ruleDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleDetail.EntityData.Children = types.NewOrderedMap()
    ruleDetail.EntityData.Children.Append("rule-summary", types.YChild{"RuleSummary", &ruleDetail.RuleSummary})
    ruleDetail.EntityData.Children.Append("codes", types.YChild{"Codes", nil})
    for i := range ruleDetail.Codes {
        types.SetYListKey(ruleDetail.Codes[i], i)
        ruleDetail.EntityData.Children.Append(types.GetSegmentPath(ruleDetail.Codes[i]), types.YChild{"Codes", ruleDetail.Codes[i]})
    }
    ruleDetail.EntityData.Leafs = types.NewOrderedMap()
    ruleDetail.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", ruleDetail.RuleName})
    ruleDetail.EntityData.Leafs.Append("all-alarms", types.YLeaf{"AllAlarms", ruleDetail.AllAlarms})
    ruleDetail.EntityData.Leafs.Append("alarm-severity", types.YLeaf{"AlarmSeverity", ruleDetail.AlarmSeverity})
    ruleDetail.EntityData.Leafs.Append("apply-source", types.YLeaf{"ApplySource", ruleDetail.ApplySource})

    ruleDetail.EntityData.YListKeys = []string {"RuleName"}

    return &(ruleDetail.EntityData)
}

// Suppression_RuleDetails_RuleDetail_RuleSummary
// Rule summary, name, etc
type Suppression_RuleDetails_RuleDetail_RuleSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Suppress Rule Name. The type is string.
    RuleNameXr interface{}

    // Applied state of the rule It could be not applied, applied or applied to
    // all. The type is AcRuleState.
    RuleState interface{}

    // Number of suppressed alarms associated with this rule. The type is
    // interface{} with range: 0..4294967295.
    SuppressedAlarmsCount interface{}
}

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetEntityData() *types.CommonEntityData {
    ruleSummary.EntityData.YFilter = ruleSummary.YFilter
    ruleSummary.EntityData.YangName = "rule-summary"
    ruleSummary.EntityData.BundleName = "cisco_ios_xr"
    ruleSummary.EntityData.ParentYangName = "rule-detail"
    ruleSummary.EntityData.SegmentPath = "rule-summary"
    ruleSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:suppression/rule-details/rule-detail/" + ruleSummary.EntityData.SegmentPath
    ruleSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSummary.EntityData.Children = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs.Append("rule-name-xr", types.YLeaf{"RuleNameXr", ruleSummary.RuleNameXr})
    ruleSummary.EntityData.Leafs.Append("rule-state", types.YLeaf{"RuleState", ruleSummary.RuleState})
    ruleSummary.EntityData.Leafs.Append("suppressed-alarms-count", types.YLeaf{"SuppressedAlarmsCount", ruleSummary.SuppressedAlarmsCount})

    ruleSummary.EntityData.YListKeys = []string {}

    return &(ruleSummary.EntityData)
}

// Suppression_RuleDetails_RuleDetail_Codes
// Message codes defining the rule.
type Suppression_RuleDetails_RuleDetail_Codes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Category of messages to which this alarm belongs. The type is string.
    Category interface{}

    // Group of messages to which this alarm belongs. The type is string.
    Group interface{}

    // Alarm code which further qualifies the alarm within a message group. The
    // type is string.
    Code interface{}
}

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetEntityData() *types.CommonEntityData {
    codes.EntityData.YFilter = codes.YFilter
    codes.EntityData.YangName = "codes"
    codes.EntityData.BundleName = "cisco_ios_xr"
    codes.EntityData.ParentYangName = "rule-detail"
    codes.EntityData.SegmentPath = "codes" + types.AddNoKeyToken(codes)
    codes.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:suppression/rule-details/rule-detail/" + codes.EntityData.SegmentPath
    codes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    codes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    codes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    codes.EntityData.Children = types.NewOrderedMap()
    codes.EntityData.Leafs = types.NewOrderedMap()
    codes.EntityData.Leafs.Append("category", types.YLeaf{"Category", codes.Category})
    codes.EntityData.Leafs.Append("group", types.YLeaf{"Group", codes.Group})
    codes.EntityData.Leafs.Append("code", types.YLeaf{"Code", codes.Code})

    codes.EntityData.YListKeys = []string {}

    return &(codes.EntityData)
}

// Correlator
// correlator
type Correlator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table that contains the database of correlation rules.
    Rules Correlator_Rules

    // Describes buffer utilization and parameters configured.
    BufferStatus Correlator_BufferStatus

    // Correlated alarms Table.
    Alarms Correlator_Alarms

    // Table that contains the ruleset summary info.
    RuleSetSummaries Correlator_RuleSetSummaries

    // Table that contains the ruleset detail info.
    RuleSetDetails Correlator_RuleSetDetails

    // Table that contains the database of correlation rule details.
    RuleDetails Correlator_RuleDetails

    // Table that contains the database of correlation rule summary.
    RuleSummaries Correlator_RuleSummaries
}

func (correlator *Correlator) GetEntityData() *types.CommonEntityData {
    correlator.EntityData.YFilter = correlator.YFilter
    correlator.EntityData.YangName = "correlator"
    correlator.EntityData.BundleName = "cisco_ios_xr"
    correlator.EntityData.ParentYangName = "Cisco-IOS-XR-infra-correlator-oper"
    correlator.EntityData.SegmentPath = "Cisco-IOS-XR-infra-correlator-oper:correlator"
    correlator.EntityData.AbsolutePath = correlator.EntityData.SegmentPath
    correlator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    correlator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    correlator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    correlator.EntityData.Children = types.NewOrderedMap()
    correlator.EntityData.Children.Append("rules", types.YChild{"Rules", &correlator.Rules})
    correlator.EntityData.Children.Append("buffer-status", types.YChild{"BufferStatus", &correlator.BufferStatus})
    correlator.EntityData.Children.Append("alarms", types.YChild{"Alarms", &correlator.Alarms})
    correlator.EntityData.Children.Append("rule-set-summaries", types.YChild{"RuleSetSummaries", &correlator.RuleSetSummaries})
    correlator.EntityData.Children.Append("rule-set-details", types.YChild{"RuleSetDetails", &correlator.RuleSetDetails})
    correlator.EntityData.Children.Append("rule-details", types.YChild{"RuleDetails", &correlator.RuleDetails})
    correlator.EntityData.Children.Append("rule-summaries", types.YChild{"RuleSummaries", &correlator.RuleSummaries})
    correlator.EntityData.Leafs = types.NewOrderedMap()

    correlator.EntityData.YListKeys = []string {}

    return &(correlator.EntityData)
}

// Correlator_Rules
// Table that contains the database of correlation
// rules
type Correlator_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One of the correlation rules. The type is slice of Correlator_Rules_Rule.
    Rule []*Correlator_Rules_Rule
}

func (rules *Correlator_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "correlator"
    rules.EntityData.SegmentPath = "rules"
    rules.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/" + rules.EntityData.SegmentPath
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

// Correlator_Rules_Rule
// One of the correlation rules
type Correlator_Rules_Rule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Correlation Rule Name. The type is string with
    // length: 1..32.
    RuleName interface{}

    // Correlation Rule Name. The type is string.
    RuleNameXr interface{}

    // Time window (in ms) for which root/all messages are kept in correlater
    // before sending them to the logger. The type is interface{} with range:
    // 0..4294967295.
    Timeout interface{}

    // Applied state of the rule It could be not applied, applied or applied to
    // all. The type is AcRuleState.
    RuleState interface{}

    // Locations (R/S/M) to which the rule is  applied. The type is slice of
    // string with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    ApplyLocation []interface{}

    // Contexts (Interfaces) to which the rule is applied. The type is slice of
    // string with length: 0..33.
    ApplyContext []interface{}

    // Message codes defining the rule. The type is slice of
    // Correlator_Rules_Rule_Codes.
    Codes []*Correlator_Rules_Rule_Codes
}

func (rule *Correlator_Rules_Rule) GetEntityData() *types.CommonEntityData {
    rule.EntityData.YFilter = rule.YFilter
    rule.EntityData.YangName = "rule"
    rule.EntityData.BundleName = "cisco_ios_xr"
    rule.EntityData.ParentYangName = "rules"
    rule.EntityData.SegmentPath = "rule" + types.AddKeyToken(rule.RuleName, "rule-name")
    rule.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/rules/" + rule.EntityData.SegmentPath
    rule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rule.EntityData.Children = types.NewOrderedMap()
    rule.EntityData.Children.Append("codes", types.YChild{"Codes", nil})
    for i := range rule.Codes {
        types.SetYListKey(rule.Codes[i], i)
        rule.EntityData.Children.Append(types.GetSegmentPath(rule.Codes[i]), types.YChild{"Codes", rule.Codes[i]})
    }
    rule.EntityData.Leafs = types.NewOrderedMap()
    rule.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", rule.RuleName})
    rule.EntityData.Leafs.Append("rule-name-xr", types.YLeaf{"RuleNameXr", rule.RuleNameXr})
    rule.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", rule.Timeout})
    rule.EntityData.Leafs.Append("rule-state", types.YLeaf{"RuleState", rule.RuleState})
    rule.EntityData.Leafs.Append("apply-location", types.YLeaf{"ApplyLocation", rule.ApplyLocation})
    rule.EntityData.Leafs.Append("apply-context", types.YLeaf{"ApplyContext", rule.ApplyContext})

    rule.EntityData.YListKeys = []string {"RuleName"}

    return &(rule.EntityData)
}

// Correlator_Rules_Rule_Codes
// Message codes defining the rule.
type Correlator_Rules_Rule_Codes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Category of messages to which this alarm belongs. The type is string.
    Category interface{}

    // Group of messages to which this alarm belongs. The type is string.
    Group interface{}

    // Alarm code which further qualifies the alarm within a message group. The
    // type is string.
    Code interface{}
}

func (codes *Correlator_Rules_Rule_Codes) GetEntityData() *types.CommonEntityData {
    codes.EntityData.YFilter = codes.YFilter
    codes.EntityData.YangName = "codes"
    codes.EntityData.BundleName = "cisco_ios_xr"
    codes.EntityData.ParentYangName = "rule"
    codes.EntityData.SegmentPath = "codes" + types.AddNoKeyToken(codes)
    codes.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/rules/rule/" + codes.EntityData.SegmentPath
    codes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    codes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    codes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    codes.EntityData.Children = types.NewOrderedMap()
    codes.EntityData.Leafs = types.NewOrderedMap()
    codes.EntityData.Leafs.Append("category", types.YLeaf{"Category", codes.Category})
    codes.EntityData.Leafs.Append("group", types.YLeaf{"Group", codes.Group})
    codes.EntityData.Leafs.Append("code", types.YLeaf{"Code", codes.Code})

    codes.EntityData.YListKeys = []string {}

    return &(codes.EntityData)
}

// Correlator_BufferStatus
// Describes buffer utilization and parameters
// configured
type Correlator_BufferStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current buffer usage. The type is interface{} with range: 0..4294967295.
    CurrentSize interface{}

    // Configured buffer size. The type is interface{} with range: 0..4294967295.
    ConfiguredSize interface{}
}

func (bufferStatus *Correlator_BufferStatus) GetEntityData() *types.CommonEntityData {
    bufferStatus.EntityData.YFilter = bufferStatus.YFilter
    bufferStatus.EntityData.YangName = "buffer-status"
    bufferStatus.EntityData.BundleName = "cisco_ios_xr"
    bufferStatus.EntityData.ParentYangName = "correlator"
    bufferStatus.EntityData.SegmentPath = "buffer-status"
    bufferStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/" + bufferStatus.EntityData.SegmentPath
    bufferStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bufferStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bufferStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bufferStatus.EntityData.Children = types.NewOrderedMap()
    bufferStatus.EntityData.Leafs = types.NewOrderedMap()
    bufferStatus.EntityData.Leafs.Append("current-size", types.YLeaf{"CurrentSize", bufferStatus.CurrentSize})
    bufferStatus.EntityData.Leafs.Append("configured-size", types.YLeaf{"ConfiguredSize", bufferStatus.ConfiguredSize})

    bufferStatus.EntityData.YListKeys = []string {}

    return &(bufferStatus.EntityData)
}

// Correlator_Alarms
// Correlated alarms Table
type Correlator_Alarms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One of the correlated alarms. The type is slice of Correlator_Alarms_Alarm.
    Alarm []*Correlator_Alarms_Alarm
}

func (alarms *Correlator_Alarms) GetEntityData() *types.CommonEntityData {
    alarms.EntityData.YFilter = alarms.YFilter
    alarms.EntityData.YangName = "alarms"
    alarms.EntityData.BundleName = "cisco_ios_xr"
    alarms.EntityData.ParentYangName = "correlator"
    alarms.EntityData.SegmentPath = "alarms"
    alarms.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/" + alarms.EntityData.SegmentPath
    alarms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarms.EntityData.Children = types.NewOrderedMap()
    alarms.EntityData.Children.Append("alarm", types.YChild{"Alarm", nil})
    for i := range alarms.Alarm {
        alarms.EntityData.Children.Append(types.GetSegmentPath(alarms.Alarm[i]), types.YChild{"Alarm", alarms.Alarm[i]})
    }
    alarms.EntityData.Leafs = types.NewOrderedMap()

    alarms.EntityData.YListKeys = []string {}

    return &(alarms.EntityData)
}

// Correlator_Alarms_Alarm
// One of the correlated alarms
type Correlator_Alarms_Alarm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Alarm ID. The type is interface{} with range:
    // 0..4294967295.
    AlarmId interface{}

    // Correlation rule name. The type is string.
    RuleName interface{}

    // Context string  for the alarm. The type is string.
    Context interface{}

    // Correlated alarm information.
    AlarmInfo Correlator_Alarms_Alarm_AlarmInfo
}

func (alarm *Correlator_Alarms_Alarm) GetEntityData() *types.CommonEntityData {
    alarm.EntityData.YFilter = alarm.YFilter
    alarm.EntityData.YangName = "alarm"
    alarm.EntityData.BundleName = "cisco_ios_xr"
    alarm.EntityData.ParentYangName = "alarms"
    alarm.EntityData.SegmentPath = "alarm" + types.AddKeyToken(alarm.AlarmId, "alarm-id")
    alarm.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/alarms/" + alarm.EntityData.SegmentPath
    alarm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarm.EntityData.Children = types.NewOrderedMap()
    alarm.EntityData.Children.Append("alarm-info", types.YChild{"AlarmInfo", &alarm.AlarmInfo})
    alarm.EntityData.Leafs = types.NewOrderedMap()
    alarm.EntityData.Leafs.Append("alarm-id", types.YLeaf{"AlarmId", alarm.AlarmId})
    alarm.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", alarm.RuleName})
    alarm.EntityData.Leafs.Append("context", types.YLeaf{"Context", alarm.Context})

    alarm.EntityData.YListKeys = []string {"AlarmId"}

    return &(alarm.EntityData)
}

// Correlator_Alarms_Alarm_AlarmInfo
// Correlated alarm information
type Correlator_Alarms_Alarm_AlarmInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source Identifier(Location).Indicates the node in which the alarm was
    // generated. The type is string.
    SourceId interface{}

    // Time when the alarm was generated. It is expressed in number of
    // milliseconds since 00:00 :00 UTC, January 1, 1970. The type is interface{}
    // with range: 0..18446744073709551615. Units are millisecond.
    Timestamp interface{}

    // Category of the alarm. The type is string.
    Category interface{}

    // Group of messages to which this alarm belongs to. The type is string.
    Group interface{}

    // Alarm code which further qualifies the alarm within a message group. The
    // type is string.
    Code interface{}

    // Severity of the alarm. The type is AlAlarmSeverity.
    Severity interface{}

    // State of the alarm (bistate alarms only). The type is AlAlarmBistate.
    State interface{}

    // Correlation Identifier. The type is interface{} with range: 0..4294967295.
    CorrelationId interface{}

    // Indicates the event id admin-level. The type is bool.
    IsAdmin interface{}

    // Full text of the Alarm. The type is string.
    AdditionalText interface{}
}

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetEntityData() *types.CommonEntityData {
    alarmInfo.EntityData.YFilter = alarmInfo.YFilter
    alarmInfo.EntityData.YangName = "alarm-info"
    alarmInfo.EntityData.BundleName = "cisco_ios_xr"
    alarmInfo.EntityData.ParentYangName = "alarm"
    alarmInfo.EntityData.SegmentPath = "alarm-info"
    alarmInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/alarms/alarm/" + alarmInfo.EntityData.SegmentPath
    alarmInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarmInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarmInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarmInfo.EntityData.Children = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs = types.NewOrderedMap()
    alarmInfo.EntityData.Leafs.Append("source-id", types.YLeaf{"SourceId", alarmInfo.SourceId})
    alarmInfo.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", alarmInfo.Timestamp})
    alarmInfo.EntityData.Leafs.Append("category", types.YLeaf{"Category", alarmInfo.Category})
    alarmInfo.EntityData.Leafs.Append("group", types.YLeaf{"Group", alarmInfo.Group})
    alarmInfo.EntityData.Leafs.Append("code", types.YLeaf{"Code", alarmInfo.Code})
    alarmInfo.EntityData.Leafs.Append("severity", types.YLeaf{"Severity", alarmInfo.Severity})
    alarmInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", alarmInfo.State})
    alarmInfo.EntityData.Leafs.Append("correlation-id", types.YLeaf{"CorrelationId", alarmInfo.CorrelationId})
    alarmInfo.EntityData.Leafs.Append("is-admin", types.YLeaf{"IsAdmin", alarmInfo.IsAdmin})
    alarmInfo.EntityData.Leafs.Append("additional-text", types.YLeaf{"AdditionalText", alarmInfo.AdditionalText})

    alarmInfo.EntityData.YListKeys = []string {}

    return &(alarmInfo.EntityData)
}

// Correlator_RuleSetSummaries
// Table that contains the ruleset summary info
type Correlator_RuleSetSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of one of the correlation rulesets. The type is slice of
    // Correlator_RuleSetSummaries_RuleSetSummary.
    RuleSetSummary []*Correlator_RuleSetSummaries_RuleSetSummary
}

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetEntityData() *types.CommonEntityData {
    ruleSetSummaries.EntityData.YFilter = ruleSetSummaries.YFilter
    ruleSetSummaries.EntityData.YangName = "rule-set-summaries"
    ruleSetSummaries.EntityData.BundleName = "cisco_ios_xr"
    ruleSetSummaries.EntityData.ParentYangName = "correlator"
    ruleSetSummaries.EntityData.SegmentPath = "rule-set-summaries"
    ruleSetSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/" + ruleSetSummaries.EntityData.SegmentPath
    ruleSetSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSetSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSetSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSetSummaries.EntityData.Children = types.NewOrderedMap()
    ruleSetSummaries.EntityData.Children.Append("rule-set-summary", types.YChild{"RuleSetSummary", nil})
    for i := range ruleSetSummaries.RuleSetSummary {
        ruleSetSummaries.EntityData.Children.Append(types.GetSegmentPath(ruleSetSummaries.RuleSetSummary[i]), types.YChild{"RuleSetSummary", ruleSetSummaries.RuleSetSummary[i]})
    }
    ruleSetSummaries.EntityData.Leafs = types.NewOrderedMap()

    ruleSetSummaries.EntityData.YListKeys = []string {}

    return &(ruleSetSummaries.EntityData)
}

// Correlator_RuleSetSummaries_RuleSetSummary
// Summary of one of the correlation rulesets
type Correlator_RuleSetSummaries_RuleSetSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Ruleset Name. The type is string with length:
    // 1..32.
    RuleSetName interface{}

    // Ruleset Name. The type is string.
    RuleSetNameXr interface{}
}

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetEntityData() *types.CommonEntityData {
    ruleSetSummary.EntityData.YFilter = ruleSetSummary.YFilter
    ruleSetSummary.EntityData.YangName = "rule-set-summary"
    ruleSetSummary.EntityData.BundleName = "cisco_ios_xr"
    ruleSetSummary.EntityData.ParentYangName = "rule-set-summaries"
    ruleSetSummary.EntityData.SegmentPath = "rule-set-summary" + types.AddKeyToken(ruleSetSummary.RuleSetName, "rule-set-name")
    ruleSetSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/rule-set-summaries/" + ruleSetSummary.EntityData.SegmentPath
    ruleSetSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSetSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSetSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSetSummary.EntityData.Children = types.NewOrderedMap()
    ruleSetSummary.EntityData.Leafs = types.NewOrderedMap()
    ruleSetSummary.EntityData.Leafs.Append("rule-set-name", types.YLeaf{"RuleSetName", ruleSetSummary.RuleSetName})
    ruleSetSummary.EntityData.Leafs.Append("rule-set-name-xr", types.YLeaf{"RuleSetNameXr", ruleSetSummary.RuleSetNameXr})

    ruleSetSummary.EntityData.YListKeys = []string {"RuleSetName"}

    return &(ruleSetSummary.EntityData)
}

// Correlator_RuleSetDetails
// Table that contains the ruleset detail info
type Correlator_RuleSetDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail of one of the correlation rulesets. The type is slice of
    // Correlator_RuleSetDetails_RuleSetDetail.
    RuleSetDetail []*Correlator_RuleSetDetails_RuleSetDetail
}

func (ruleSetDetails *Correlator_RuleSetDetails) GetEntityData() *types.CommonEntityData {
    ruleSetDetails.EntityData.YFilter = ruleSetDetails.YFilter
    ruleSetDetails.EntityData.YangName = "rule-set-details"
    ruleSetDetails.EntityData.BundleName = "cisco_ios_xr"
    ruleSetDetails.EntityData.ParentYangName = "correlator"
    ruleSetDetails.EntityData.SegmentPath = "rule-set-details"
    ruleSetDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/" + ruleSetDetails.EntityData.SegmentPath
    ruleSetDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSetDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSetDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSetDetails.EntityData.Children = types.NewOrderedMap()
    ruleSetDetails.EntityData.Children.Append("rule-set-detail", types.YChild{"RuleSetDetail", nil})
    for i := range ruleSetDetails.RuleSetDetail {
        ruleSetDetails.EntityData.Children.Append(types.GetSegmentPath(ruleSetDetails.RuleSetDetail[i]), types.YChild{"RuleSetDetail", ruleSetDetails.RuleSetDetail[i]})
    }
    ruleSetDetails.EntityData.Leafs = types.NewOrderedMap()

    ruleSetDetails.EntityData.YListKeys = []string {}

    return &(ruleSetDetails.EntityData)
}

// Correlator_RuleSetDetails_RuleSetDetail
// Detail of one of the correlation rulesets
type Correlator_RuleSetDetails_RuleSetDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Ruleset Name. The type is string with length:
    // 1..32.
    RuleSetName interface{}

    // Ruleset Name. The type is string.
    RuleSetNameXr interface{}

    // Rules contained in a ruleset. The type is slice of
    // Correlator_RuleSetDetails_RuleSetDetail_Rules.
    Rules []*Correlator_RuleSetDetails_RuleSetDetail_Rules
}

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetEntityData() *types.CommonEntityData {
    ruleSetDetail.EntityData.YFilter = ruleSetDetail.YFilter
    ruleSetDetail.EntityData.YangName = "rule-set-detail"
    ruleSetDetail.EntityData.BundleName = "cisco_ios_xr"
    ruleSetDetail.EntityData.ParentYangName = "rule-set-details"
    ruleSetDetail.EntityData.SegmentPath = "rule-set-detail" + types.AddKeyToken(ruleSetDetail.RuleSetName, "rule-set-name")
    ruleSetDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/rule-set-details/" + ruleSetDetail.EntityData.SegmentPath
    ruleSetDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSetDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSetDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSetDetail.EntityData.Children = types.NewOrderedMap()
    ruleSetDetail.EntityData.Children.Append("rules", types.YChild{"Rules", nil})
    for i := range ruleSetDetail.Rules {
        types.SetYListKey(ruleSetDetail.Rules[i], i)
        ruleSetDetail.EntityData.Children.Append(types.GetSegmentPath(ruleSetDetail.Rules[i]), types.YChild{"Rules", ruleSetDetail.Rules[i]})
    }
    ruleSetDetail.EntityData.Leafs = types.NewOrderedMap()
    ruleSetDetail.EntityData.Leafs.Append("rule-set-name", types.YLeaf{"RuleSetName", ruleSetDetail.RuleSetName})
    ruleSetDetail.EntityData.Leafs.Append("rule-set-name-xr", types.YLeaf{"RuleSetNameXr", ruleSetDetail.RuleSetNameXr})

    ruleSetDetail.EntityData.YListKeys = []string {"RuleSetName"}

    return &(ruleSetDetail.EntityData)
}

// Correlator_RuleSetDetails_RuleSetDetail_Rules
// Rules contained in a ruleset
type Correlator_RuleSetDetails_RuleSetDetail_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Correlation Rule Name. The type is string.
    RuleNameXr interface{}

    // Whether the rule is stateful. The type is bool.
    Stateful interface{}

    // Applied state of the rule It could be not applied, applied or applied to
    // all. The type is AcRuleState.
    RuleState interface{}

    // Number of buffered alarms correlated to this rule. The type is interface{}
    // with range: 0..4294967295.
    BufferedAlarmsCount interface{}
}

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "rule-set-detail"
    rules.EntityData.SegmentPath = "rules" + types.AddNoKeyToken(rules)
    rules.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/rule-set-details/rule-set-detail/" + rules.EntityData.SegmentPath
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = types.NewOrderedMap()
    rules.EntityData.Leafs = types.NewOrderedMap()
    rules.EntityData.Leafs.Append("rule-name-xr", types.YLeaf{"RuleNameXr", rules.RuleNameXr})
    rules.EntityData.Leafs.Append("stateful", types.YLeaf{"Stateful", rules.Stateful})
    rules.EntityData.Leafs.Append("rule-state", types.YLeaf{"RuleState", rules.RuleState})
    rules.EntityData.Leafs.Append("buffered-alarms-count", types.YLeaf{"BufferedAlarmsCount", rules.BufferedAlarmsCount})

    rules.EntityData.YListKeys = []string {}

    return &(rules.EntityData)
}

// Correlator_RuleDetails
// Table that contains the database of correlation
// rule details
type Correlator_RuleDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Details of one of the correlation rules. The type is slice of
    // Correlator_RuleDetails_RuleDetail.
    RuleDetail []*Correlator_RuleDetails_RuleDetail
}

func (ruleDetails *Correlator_RuleDetails) GetEntityData() *types.CommonEntityData {
    ruleDetails.EntityData.YFilter = ruleDetails.YFilter
    ruleDetails.EntityData.YangName = "rule-details"
    ruleDetails.EntityData.BundleName = "cisco_ios_xr"
    ruleDetails.EntityData.ParentYangName = "correlator"
    ruleDetails.EntityData.SegmentPath = "rule-details"
    ruleDetails.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/" + ruleDetails.EntityData.SegmentPath
    ruleDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleDetails.EntityData.Children = types.NewOrderedMap()
    ruleDetails.EntityData.Children.Append("rule-detail", types.YChild{"RuleDetail", nil})
    for i := range ruleDetails.RuleDetail {
        ruleDetails.EntityData.Children.Append(types.GetSegmentPath(ruleDetails.RuleDetail[i]), types.YChild{"RuleDetail", ruleDetails.RuleDetail[i]})
    }
    ruleDetails.EntityData.Leafs = types.NewOrderedMap()

    ruleDetails.EntityData.YListKeys = []string {}

    return &(ruleDetails.EntityData)
}

// Correlator_RuleDetails_RuleDetail
// Details of one of the correlation rules
type Correlator_RuleDetails_RuleDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Correlation Rule Name. The type is string with
    // length: 1..32.
    RuleName interface{}

    // Time window (in ms) for which root/all messages are kept in correlater
    // before sending them to the logger. The type is interface{} with range:
    // 0..4294967295.
    Timeout interface{}

    // Timeout before root cause alarm. The type is interface{} with range:
    // 0..4294967295.
    RootCauseTimeout interface{}

    // True if the rule is internal. The type is bool.
    Internal interface{}

    // Whether to reissue non-bistate alarms. The type is bool.
    ReissueNonBistate interface{}

    // Reparent. The type is bool.
    Reparent interface{}

    // Whether context correlation is enabled. The type is bool.
    ContextCorrelation interface{}

    // Locations (R/S/M) to which the rule is applied. The type is slice of string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    ApplyLocation []interface{}

    // Contexts (Interfaces) to which the rule is applied. The type is slice of
    // string with length: 0..33.
    ApplyContext []interface{}

    // Rule summary, name, etc.
    RuleSummary Correlator_RuleDetails_RuleDetail_RuleSummary

    // Message codes defining the rule. The type is slice of
    // Correlator_RuleDetails_RuleDetail_Codes.
    Codes []*Correlator_RuleDetails_RuleDetail_Codes
}

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetEntityData() *types.CommonEntityData {
    ruleDetail.EntityData.YFilter = ruleDetail.YFilter
    ruleDetail.EntityData.YangName = "rule-detail"
    ruleDetail.EntityData.BundleName = "cisco_ios_xr"
    ruleDetail.EntityData.ParentYangName = "rule-details"
    ruleDetail.EntityData.SegmentPath = "rule-detail" + types.AddKeyToken(ruleDetail.RuleName, "rule-name")
    ruleDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/rule-details/" + ruleDetail.EntityData.SegmentPath
    ruleDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleDetail.EntityData.Children = types.NewOrderedMap()
    ruleDetail.EntityData.Children.Append("rule-summary", types.YChild{"RuleSummary", &ruleDetail.RuleSummary})
    ruleDetail.EntityData.Children.Append("codes", types.YChild{"Codes", nil})
    for i := range ruleDetail.Codes {
        types.SetYListKey(ruleDetail.Codes[i], i)
        ruleDetail.EntityData.Children.Append(types.GetSegmentPath(ruleDetail.Codes[i]), types.YChild{"Codes", ruleDetail.Codes[i]})
    }
    ruleDetail.EntityData.Leafs = types.NewOrderedMap()
    ruleDetail.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", ruleDetail.RuleName})
    ruleDetail.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", ruleDetail.Timeout})
    ruleDetail.EntityData.Leafs.Append("root-cause-timeout", types.YLeaf{"RootCauseTimeout", ruleDetail.RootCauseTimeout})
    ruleDetail.EntityData.Leafs.Append("internal", types.YLeaf{"Internal", ruleDetail.Internal})
    ruleDetail.EntityData.Leafs.Append("reissue-non-bistate", types.YLeaf{"ReissueNonBistate", ruleDetail.ReissueNonBistate})
    ruleDetail.EntityData.Leafs.Append("reparent", types.YLeaf{"Reparent", ruleDetail.Reparent})
    ruleDetail.EntityData.Leafs.Append("context-correlation", types.YLeaf{"ContextCorrelation", ruleDetail.ContextCorrelation})
    ruleDetail.EntityData.Leafs.Append("apply-location", types.YLeaf{"ApplyLocation", ruleDetail.ApplyLocation})
    ruleDetail.EntityData.Leafs.Append("apply-context", types.YLeaf{"ApplyContext", ruleDetail.ApplyContext})

    ruleDetail.EntityData.YListKeys = []string {"RuleName"}

    return &(ruleDetail.EntityData)
}

// Correlator_RuleDetails_RuleDetail_RuleSummary
// Rule summary, name, etc
type Correlator_RuleDetails_RuleDetail_RuleSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Correlation Rule Name. The type is string.
    RuleNameXr interface{}

    // Whether the rule is stateful. The type is bool.
    Stateful interface{}

    // Applied state of the rule It could be not applied, applied or applied to
    // all. The type is AcRuleState.
    RuleState interface{}

    // Number of buffered alarms correlated to this rule. The type is interface{}
    // with range: 0..4294967295.
    BufferedAlarmsCount interface{}
}

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetEntityData() *types.CommonEntityData {
    ruleSummary.EntityData.YFilter = ruleSummary.YFilter
    ruleSummary.EntityData.YangName = "rule-summary"
    ruleSummary.EntityData.BundleName = "cisco_ios_xr"
    ruleSummary.EntityData.ParentYangName = "rule-detail"
    ruleSummary.EntityData.SegmentPath = "rule-summary"
    ruleSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/rule-details/rule-detail/" + ruleSummary.EntityData.SegmentPath
    ruleSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSummary.EntityData.Children = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs.Append("rule-name-xr", types.YLeaf{"RuleNameXr", ruleSummary.RuleNameXr})
    ruleSummary.EntityData.Leafs.Append("stateful", types.YLeaf{"Stateful", ruleSummary.Stateful})
    ruleSummary.EntityData.Leafs.Append("rule-state", types.YLeaf{"RuleState", ruleSummary.RuleState})
    ruleSummary.EntityData.Leafs.Append("buffered-alarms-count", types.YLeaf{"BufferedAlarmsCount", ruleSummary.BufferedAlarmsCount})

    ruleSummary.EntityData.YListKeys = []string {}

    return &(ruleSummary.EntityData)
}

// Correlator_RuleDetails_RuleDetail_Codes
// Message codes defining the rule.
type Correlator_RuleDetails_RuleDetail_Codes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Category of messages to which this alarm belongs. The type is string.
    Category interface{}

    // Group of messages to which this alarm belongs. The type is string.
    Group interface{}

    // Alarm code which further qualifies the alarm within a message group. The
    // type is string.
    Code interface{}
}

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetEntityData() *types.CommonEntityData {
    codes.EntityData.YFilter = codes.YFilter
    codes.EntityData.YangName = "codes"
    codes.EntityData.BundleName = "cisco_ios_xr"
    codes.EntityData.ParentYangName = "rule-detail"
    codes.EntityData.SegmentPath = "codes" + types.AddNoKeyToken(codes)
    codes.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/rule-details/rule-detail/" + codes.EntityData.SegmentPath
    codes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    codes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    codes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    codes.EntityData.Children = types.NewOrderedMap()
    codes.EntityData.Leafs = types.NewOrderedMap()
    codes.EntityData.Leafs.Append("category", types.YLeaf{"Category", codes.Category})
    codes.EntityData.Leafs.Append("group", types.YLeaf{"Group", codes.Group})
    codes.EntityData.Leafs.Append("code", types.YLeaf{"Code", codes.Code})

    codes.EntityData.YListKeys = []string {}

    return &(codes.EntityData)
}

// Correlator_RuleSummaries
// Table that contains the database of correlation
// rule summary
type Correlator_RuleSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One of the correlation rules. The type is slice of
    // Correlator_RuleSummaries_RuleSummary.
    RuleSummary []*Correlator_RuleSummaries_RuleSummary
}

func (ruleSummaries *Correlator_RuleSummaries) GetEntityData() *types.CommonEntityData {
    ruleSummaries.EntityData.YFilter = ruleSummaries.YFilter
    ruleSummaries.EntityData.YangName = "rule-summaries"
    ruleSummaries.EntityData.BundleName = "cisco_ios_xr"
    ruleSummaries.EntityData.ParentYangName = "correlator"
    ruleSummaries.EntityData.SegmentPath = "rule-summaries"
    ruleSummaries.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/" + ruleSummaries.EntityData.SegmentPath
    ruleSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSummaries.EntityData.Children = types.NewOrderedMap()
    ruleSummaries.EntityData.Children.Append("rule-summary", types.YChild{"RuleSummary", nil})
    for i := range ruleSummaries.RuleSummary {
        ruleSummaries.EntityData.Children.Append(types.GetSegmentPath(ruleSummaries.RuleSummary[i]), types.YChild{"RuleSummary", ruleSummaries.RuleSummary[i]})
    }
    ruleSummaries.EntityData.Leafs = types.NewOrderedMap()

    ruleSummaries.EntityData.YListKeys = []string {}

    return &(ruleSummaries.EntityData)
}

// Correlator_RuleSummaries_RuleSummary
// One of the correlation rules
type Correlator_RuleSummaries_RuleSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Correlation Rule Name. The type is string with
    // length: 1..32.
    RuleName interface{}

    // Correlation Rule Name. The type is string.
    RuleNameXr interface{}

    // Whether the rule is stateful. The type is bool.
    Stateful interface{}

    // Applied state of the rule It could be not applied, applied or applied to
    // all. The type is AcRuleState.
    RuleState interface{}

    // Number of buffered alarms correlated to this rule. The type is interface{}
    // with range: 0..4294967295.
    BufferedAlarmsCount interface{}
}

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetEntityData() *types.CommonEntityData {
    ruleSummary.EntityData.YFilter = ruleSummary.YFilter
    ruleSummary.EntityData.YangName = "rule-summary"
    ruleSummary.EntityData.BundleName = "cisco_ios_xr"
    ruleSummary.EntityData.ParentYangName = "rule-summaries"
    ruleSummary.EntityData.SegmentPath = "rule-summary" + types.AddKeyToken(ruleSummary.RuleName, "rule-name")
    ruleSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-correlator-oper:correlator/rule-summaries/" + ruleSummary.EntityData.SegmentPath
    ruleSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSummary.EntityData.Children = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs = types.NewOrderedMap()
    ruleSummary.EntityData.Leafs.Append("rule-name", types.YLeaf{"RuleName", ruleSummary.RuleName})
    ruleSummary.EntityData.Leafs.Append("rule-name-xr", types.YLeaf{"RuleNameXr", ruleSummary.RuleNameXr})
    ruleSummary.EntityData.Leafs.Append("stateful", types.YLeaf{"Stateful", ruleSummary.Stateful})
    ruleSummary.EntityData.Leafs.Append("rule-state", types.YLeaf{"RuleState", ruleSummary.RuleState})
    ruleSummary.EntityData.Leafs.Append("buffered-alarms-count", types.YLeaf{"BufferedAlarmsCount", ruleSummary.BufferedAlarmsCount})

    ruleSummary.EntityData.YListKeys = []string {"RuleName"}

    return &(ruleSummary.EntityData)
}

