// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-correlator package operational data.
// 
// This module contains definitions
// for the following management objects:
//   suppression: Suppression operational data
//   correlator: correlator
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    AlAlarmSeverity_error AlAlarmSeverity = "error"

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Table that contains the database of suppression rule summary.
    RuleSummaries Suppression_RuleSummaries

    // Table that contains the database of suppression rule details.
    RuleDetails Suppression_RuleDetails
}

func (suppression *Suppression) GetFilter() yfilter.YFilter { return suppression.YFilter }

func (suppression *Suppression) SetFilter(yf yfilter.YFilter) { suppression.YFilter = yf }

func (suppression *Suppression) GetGoName(yname string) string {
    if yname == "rule-summaries" { return "RuleSummaries" }
    if yname == "rule-details" { return "RuleDetails" }
    return ""
}

func (suppression *Suppression) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-correlator-oper:suppression"
}

func (suppression *Suppression) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-summaries" {
        return &suppression.RuleSummaries
    }
    if childYangName == "rule-details" {
        return &suppression.RuleDetails
    }
    return nil
}

func (suppression *Suppression) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rule-summaries"] = &suppression.RuleSummaries
    children["rule-details"] = &suppression.RuleDetails
    return children
}

func (suppression *Suppression) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (suppression *Suppression) GetBundleName() string { return "cisco_ios_xr" }

func (suppression *Suppression) GetYangName() string { return "suppression" }

func (suppression *Suppression) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (suppression *Suppression) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (suppression *Suppression) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (suppression *Suppression) SetParent(parent types.Entity) { suppression.parent = parent }

func (suppression *Suppression) GetParent() types.Entity { return suppression.parent }

func (suppression *Suppression) GetParentYangName() string { return "Cisco-IOS-XR-infra-correlator-oper" }

// Suppression_RuleSummaries
// Table that contains the database of suppression
// rule summary
type Suppression_RuleSummaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One of the suppression rules. The type is slice of
    // Suppression_RuleSummaries_RuleSummary.
    RuleSummary []Suppression_RuleSummaries_RuleSummary
}

func (ruleSummaries *Suppression_RuleSummaries) GetFilter() yfilter.YFilter { return ruleSummaries.YFilter }

func (ruleSummaries *Suppression_RuleSummaries) SetFilter(yf yfilter.YFilter) { ruleSummaries.YFilter = yf }

func (ruleSummaries *Suppression_RuleSummaries) GetGoName(yname string) string {
    if yname == "rule-summary" { return "RuleSummary" }
    return ""
}

func (ruleSummaries *Suppression_RuleSummaries) GetSegmentPath() string {
    return "rule-summaries"
}

func (ruleSummaries *Suppression_RuleSummaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-summary" {
        for _, c := range ruleSummaries.RuleSummary {
            if ruleSummaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Suppression_RuleSummaries_RuleSummary{}
        ruleSummaries.RuleSummary = append(ruleSummaries.RuleSummary, child)
        return &ruleSummaries.RuleSummary[len(ruleSummaries.RuleSummary)-1]
    }
    return nil
}

func (ruleSummaries *Suppression_RuleSummaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleSummaries.RuleSummary {
        children[ruleSummaries.RuleSummary[i].GetSegmentPath()] = &ruleSummaries.RuleSummary[i]
    }
    return children
}

func (ruleSummaries *Suppression_RuleSummaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleSummaries *Suppression_RuleSummaries) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSummaries *Suppression_RuleSummaries) GetYangName() string { return "rule-summaries" }

func (ruleSummaries *Suppression_RuleSummaries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSummaries *Suppression_RuleSummaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSummaries *Suppression_RuleSummaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSummaries *Suppression_RuleSummaries) SetParent(parent types.Entity) { ruleSummaries.parent = parent }

func (ruleSummaries *Suppression_RuleSummaries) GetParent() types.Entity { return ruleSummaries.parent }

func (ruleSummaries *Suppression_RuleSummaries) GetParentYangName() string { return "suppression" }

// Suppression_RuleSummaries_RuleSummary
// One of the suppression rules
type Suppression_RuleSummaries_RuleSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetFilter() yfilter.YFilter { return ruleSummary.YFilter }

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) SetFilter(yf yfilter.YFilter) { ruleSummary.YFilter = yf }

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetGoName(yname string) string {
    if yname == "rule-name" { return "RuleName" }
    if yname == "rule-name-xr" { return "RuleNameXr" }
    if yname == "rule-state" { return "RuleState" }
    if yname == "suppressed-alarms-count" { return "SuppressedAlarmsCount" }
    return ""
}

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetSegmentPath() string {
    return "rule-summary" + "[rule-name='" + fmt.Sprintf("%v", ruleSummary.RuleName) + "']"
}

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name"] = ruleSummary.RuleName
    leafs["rule-name-xr"] = ruleSummary.RuleNameXr
    leafs["rule-state"] = ruleSummary.RuleState
    leafs["suppressed-alarms-count"] = ruleSummary.SuppressedAlarmsCount
    return leafs
}

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetYangName() string { return "rule-summary" }

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) SetParent(parent types.Entity) { ruleSummary.parent = parent }

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetParent() types.Entity { return ruleSummary.parent }

func (ruleSummary *Suppression_RuleSummaries_RuleSummary) GetParentYangName() string { return "rule-summaries" }

// Suppression_RuleDetails
// Table that contains the database of suppression
// rule details
type Suppression_RuleDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details of one of the suppression rules. The type is slice of
    // Suppression_RuleDetails_RuleDetail.
    RuleDetail []Suppression_RuleDetails_RuleDetail
}

func (ruleDetails *Suppression_RuleDetails) GetFilter() yfilter.YFilter { return ruleDetails.YFilter }

func (ruleDetails *Suppression_RuleDetails) SetFilter(yf yfilter.YFilter) { ruleDetails.YFilter = yf }

func (ruleDetails *Suppression_RuleDetails) GetGoName(yname string) string {
    if yname == "rule-detail" { return "RuleDetail" }
    return ""
}

func (ruleDetails *Suppression_RuleDetails) GetSegmentPath() string {
    return "rule-details"
}

func (ruleDetails *Suppression_RuleDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-detail" {
        for _, c := range ruleDetails.RuleDetail {
            if ruleDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Suppression_RuleDetails_RuleDetail{}
        ruleDetails.RuleDetail = append(ruleDetails.RuleDetail, child)
        return &ruleDetails.RuleDetail[len(ruleDetails.RuleDetail)-1]
    }
    return nil
}

func (ruleDetails *Suppression_RuleDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleDetails.RuleDetail {
        children[ruleDetails.RuleDetail[i].GetSegmentPath()] = &ruleDetails.RuleDetail[i]
    }
    return children
}

func (ruleDetails *Suppression_RuleDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleDetails *Suppression_RuleDetails) GetBundleName() string { return "cisco_ios_xr" }

func (ruleDetails *Suppression_RuleDetails) GetYangName() string { return "rule-details" }

func (ruleDetails *Suppression_RuleDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleDetails *Suppression_RuleDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleDetails *Suppression_RuleDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleDetails *Suppression_RuleDetails) SetParent(parent types.Entity) { ruleDetails.parent = parent }

func (ruleDetails *Suppression_RuleDetails) GetParent() types.Entity { return ruleDetails.parent }

func (ruleDetails *Suppression_RuleDetails) GetParentYangName() string { return "suppression" }

// Suppression_RuleDetails_RuleDetail
// Details of one of the suppression rules
type Suppression_RuleDetails_RuleDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    Codes []Suppression_RuleDetails_RuleDetail_Codes
}

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetFilter() yfilter.YFilter { return ruleDetail.YFilter }

func (ruleDetail *Suppression_RuleDetails_RuleDetail) SetFilter(yf yfilter.YFilter) { ruleDetail.YFilter = yf }

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetGoName(yname string) string {
    if yname == "rule-name" { return "RuleName" }
    if yname == "all-alarms" { return "AllAlarms" }
    if yname == "alarm-severity" { return "AlarmSeverity" }
    if yname == "apply-source" { return "ApplySource" }
    if yname == "rule-summary" { return "RuleSummary" }
    if yname == "codes" { return "Codes" }
    return ""
}

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetSegmentPath() string {
    return "rule-detail" + "[rule-name='" + fmt.Sprintf("%v", ruleDetail.RuleName) + "']"
}

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-summary" {
        return &ruleDetail.RuleSummary
    }
    if childYangName == "codes" {
        for _, c := range ruleDetail.Codes {
            if ruleDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Suppression_RuleDetails_RuleDetail_Codes{}
        ruleDetail.Codes = append(ruleDetail.Codes, child)
        return &ruleDetail.Codes[len(ruleDetail.Codes)-1]
    }
    return nil
}

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rule-summary"] = &ruleDetail.RuleSummary
    for i := range ruleDetail.Codes {
        children[ruleDetail.Codes[i].GetSegmentPath()] = &ruleDetail.Codes[i]
    }
    return children
}

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name"] = ruleDetail.RuleName
    leafs["all-alarms"] = ruleDetail.AllAlarms
    leafs["alarm-severity"] = ruleDetail.AlarmSeverity
    leafs["apply-source"] = ruleDetail.ApplySource
    return leafs
}

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetBundleName() string { return "cisco_ios_xr" }

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetYangName() string { return "rule-detail" }

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleDetail *Suppression_RuleDetails_RuleDetail) SetParent(parent types.Entity) { ruleDetail.parent = parent }

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetParent() types.Entity { return ruleDetail.parent }

func (ruleDetail *Suppression_RuleDetails_RuleDetail) GetParentYangName() string { return "rule-details" }

// Suppression_RuleDetails_RuleDetail_RuleSummary
// Rule summary, name, etc
type Suppression_RuleDetails_RuleDetail_RuleSummary struct {
    parent types.Entity
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

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetFilter() yfilter.YFilter { return ruleSummary.YFilter }

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) SetFilter(yf yfilter.YFilter) { ruleSummary.YFilter = yf }

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetGoName(yname string) string {
    if yname == "rule-name-xr" { return "RuleNameXr" }
    if yname == "rule-state" { return "RuleState" }
    if yname == "suppressed-alarms-count" { return "SuppressedAlarmsCount" }
    return ""
}

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetSegmentPath() string {
    return "rule-summary"
}

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name-xr"] = ruleSummary.RuleNameXr
    leafs["rule-state"] = ruleSummary.RuleState
    leafs["suppressed-alarms-count"] = ruleSummary.SuppressedAlarmsCount
    return leafs
}

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetYangName() string { return "rule-summary" }

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) SetParent(parent types.Entity) { ruleSummary.parent = parent }

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetParent() types.Entity { return ruleSummary.parent }

func (ruleSummary *Suppression_RuleDetails_RuleDetail_RuleSummary) GetParentYangName() string { return "rule-detail" }

// Suppression_RuleDetails_RuleDetail_Codes
// Message codes defining the rule.
type Suppression_RuleDetails_RuleDetail_Codes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of messages to which this alarm belongs. The type is string.
    Category interface{}

    // Group of messages to which this alarm belongs. The type is string.
    Group interface{}

    // Alarm code which further qualifies the alarm within a message group. The
    // type is string.
    Code interface{}
}

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetFilter() yfilter.YFilter { return codes.YFilter }

func (codes *Suppression_RuleDetails_RuleDetail_Codes) SetFilter(yf yfilter.YFilter) { codes.YFilter = yf }

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "code" { return "Code" }
    return ""
}

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetSegmentPath() string {
    return "codes"
}

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = codes.Category
    leafs["group"] = codes.Group
    leafs["code"] = codes.Code
    return leafs
}

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetBundleName() string { return "cisco_ios_xr" }

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetYangName() string { return "codes" }

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (codes *Suppression_RuleDetails_RuleDetail_Codes) SetParent(parent types.Entity) { codes.parent = parent }

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetParent() types.Entity { return codes.parent }

func (codes *Suppression_RuleDetails_RuleDetail_Codes) GetParentYangName() string { return "rule-detail" }

// Correlator
// correlator
type Correlator struct {
    parent types.Entity
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

func (correlator *Correlator) GetFilter() yfilter.YFilter { return correlator.YFilter }

func (correlator *Correlator) SetFilter(yf yfilter.YFilter) { correlator.YFilter = yf }

func (correlator *Correlator) GetGoName(yname string) string {
    if yname == "rules" { return "Rules" }
    if yname == "buffer-status" { return "BufferStatus" }
    if yname == "alarms" { return "Alarms" }
    if yname == "rule-set-summaries" { return "RuleSetSummaries" }
    if yname == "rule-set-details" { return "RuleSetDetails" }
    if yname == "rule-details" { return "RuleDetails" }
    if yname == "rule-summaries" { return "RuleSummaries" }
    return ""
}

func (correlator *Correlator) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-correlator-oper:correlator"
}

func (correlator *Correlator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rules" {
        return &correlator.Rules
    }
    if childYangName == "buffer-status" {
        return &correlator.BufferStatus
    }
    if childYangName == "alarms" {
        return &correlator.Alarms
    }
    if childYangName == "rule-set-summaries" {
        return &correlator.RuleSetSummaries
    }
    if childYangName == "rule-set-details" {
        return &correlator.RuleSetDetails
    }
    if childYangName == "rule-details" {
        return &correlator.RuleDetails
    }
    if childYangName == "rule-summaries" {
        return &correlator.RuleSummaries
    }
    return nil
}

func (correlator *Correlator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rules"] = &correlator.Rules
    children["buffer-status"] = &correlator.BufferStatus
    children["alarms"] = &correlator.Alarms
    children["rule-set-summaries"] = &correlator.RuleSetSummaries
    children["rule-set-details"] = &correlator.RuleSetDetails
    children["rule-details"] = &correlator.RuleDetails
    children["rule-summaries"] = &correlator.RuleSummaries
    return children
}

func (correlator *Correlator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (correlator *Correlator) GetBundleName() string { return "cisco_ios_xr" }

func (correlator *Correlator) GetYangName() string { return "correlator" }

func (correlator *Correlator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (correlator *Correlator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (correlator *Correlator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (correlator *Correlator) SetParent(parent types.Entity) { correlator.parent = parent }

func (correlator *Correlator) GetParent() types.Entity { return correlator.parent }

func (correlator *Correlator) GetParentYangName() string { return "Cisco-IOS-XR-infra-correlator-oper" }

// Correlator_Rules
// Table that contains the database of correlation
// rules
type Correlator_Rules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One of the correlation rules. The type is slice of Correlator_Rules_Rule.
    Rule []Correlator_Rules_Rule
}

func (rules *Correlator_Rules) GetFilter() yfilter.YFilter { return rules.YFilter }

func (rules *Correlator_Rules) SetFilter(yf yfilter.YFilter) { rules.YFilter = yf }

func (rules *Correlator_Rules) GetGoName(yname string) string {
    if yname == "rule" { return "Rule" }
    return ""
}

func (rules *Correlator_Rules) GetSegmentPath() string {
    return "rules"
}

func (rules *Correlator_Rules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule" {
        for _, c := range rules.Rule {
            if rules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Correlator_Rules_Rule{}
        rules.Rule = append(rules.Rule, child)
        return &rules.Rule[len(rules.Rule)-1]
    }
    return nil
}

func (rules *Correlator_Rules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rules.Rule {
        children[rules.Rule[i].GetSegmentPath()] = &rules.Rule[i]
    }
    return children
}

func (rules *Correlator_Rules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rules *Correlator_Rules) GetBundleName() string { return "cisco_ios_xr" }

func (rules *Correlator_Rules) GetYangName() string { return "rules" }

func (rules *Correlator_Rules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rules *Correlator_Rules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rules *Correlator_Rules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rules *Correlator_Rules) SetParent(parent types.Entity) { rules.parent = parent }

func (rules *Correlator_Rules) GetParent() types.Entity { return rules.parent }

func (rules *Correlator_Rules) GetParentYangName() string { return "correlator" }

// Correlator_Rules_Rule
// One of the correlation rules
type Correlator_Rules_Rule struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    Codes []Correlator_Rules_Rule_Codes
}

func (rule *Correlator_Rules_Rule) GetFilter() yfilter.YFilter { return rule.YFilter }

func (rule *Correlator_Rules_Rule) SetFilter(yf yfilter.YFilter) { rule.YFilter = yf }

func (rule *Correlator_Rules_Rule) GetGoName(yname string) string {
    if yname == "rule-name" { return "RuleName" }
    if yname == "rule-name-xr" { return "RuleNameXr" }
    if yname == "timeout" { return "Timeout" }
    if yname == "rule-state" { return "RuleState" }
    if yname == "apply-location" { return "ApplyLocation" }
    if yname == "apply-context" { return "ApplyContext" }
    if yname == "codes" { return "Codes" }
    return ""
}

func (rule *Correlator_Rules_Rule) GetSegmentPath() string {
    return "rule" + "[rule-name='" + fmt.Sprintf("%v", rule.RuleName) + "']"
}

func (rule *Correlator_Rules_Rule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "codes" {
        for _, c := range rule.Codes {
            if rule.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Correlator_Rules_Rule_Codes{}
        rule.Codes = append(rule.Codes, child)
        return &rule.Codes[len(rule.Codes)-1]
    }
    return nil
}

func (rule *Correlator_Rules_Rule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rule.Codes {
        children[rule.Codes[i].GetSegmentPath()] = &rule.Codes[i]
    }
    return children
}

func (rule *Correlator_Rules_Rule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name"] = rule.RuleName
    leafs["rule-name-xr"] = rule.RuleNameXr
    leafs["timeout"] = rule.Timeout
    leafs["rule-state"] = rule.RuleState
    leafs["apply-location"] = rule.ApplyLocation
    leafs["apply-context"] = rule.ApplyContext
    return leafs
}

func (rule *Correlator_Rules_Rule) GetBundleName() string { return "cisco_ios_xr" }

func (rule *Correlator_Rules_Rule) GetYangName() string { return "rule" }

func (rule *Correlator_Rules_Rule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rule *Correlator_Rules_Rule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rule *Correlator_Rules_Rule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rule *Correlator_Rules_Rule) SetParent(parent types.Entity) { rule.parent = parent }

func (rule *Correlator_Rules_Rule) GetParent() types.Entity { return rule.parent }

func (rule *Correlator_Rules_Rule) GetParentYangName() string { return "rules" }

// Correlator_Rules_Rule_Codes
// Message codes defining the rule.
type Correlator_Rules_Rule_Codes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of messages to which this alarm belongs. The type is string.
    Category interface{}

    // Group of messages to which this alarm belongs. The type is string.
    Group interface{}

    // Alarm code which further qualifies the alarm within a message group. The
    // type is string.
    Code interface{}
}

func (codes *Correlator_Rules_Rule_Codes) GetFilter() yfilter.YFilter { return codes.YFilter }

func (codes *Correlator_Rules_Rule_Codes) SetFilter(yf yfilter.YFilter) { codes.YFilter = yf }

func (codes *Correlator_Rules_Rule_Codes) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "code" { return "Code" }
    return ""
}

func (codes *Correlator_Rules_Rule_Codes) GetSegmentPath() string {
    return "codes"
}

func (codes *Correlator_Rules_Rule_Codes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (codes *Correlator_Rules_Rule_Codes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (codes *Correlator_Rules_Rule_Codes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = codes.Category
    leafs["group"] = codes.Group
    leafs["code"] = codes.Code
    return leafs
}

func (codes *Correlator_Rules_Rule_Codes) GetBundleName() string { return "cisco_ios_xr" }

func (codes *Correlator_Rules_Rule_Codes) GetYangName() string { return "codes" }

func (codes *Correlator_Rules_Rule_Codes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (codes *Correlator_Rules_Rule_Codes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (codes *Correlator_Rules_Rule_Codes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (codes *Correlator_Rules_Rule_Codes) SetParent(parent types.Entity) { codes.parent = parent }

func (codes *Correlator_Rules_Rule_Codes) GetParent() types.Entity { return codes.parent }

func (codes *Correlator_Rules_Rule_Codes) GetParentYangName() string { return "rule" }

// Correlator_BufferStatus
// Describes buffer utilization and parameters
// configured
type Correlator_BufferStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current buffer usage. The type is interface{} with range: 0..4294967295.
    CurrentSize interface{}

    // Configured buffer size. The type is interface{} with range: 0..4294967295.
    ConfiguredSize interface{}
}

func (bufferStatus *Correlator_BufferStatus) GetFilter() yfilter.YFilter { return bufferStatus.YFilter }

func (bufferStatus *Correlator_BufferStatus) SetFilter(yf yfilter.YFilter) { bufferStatus.YFilter = yf }

func (bufferStatus *Correlator_BufferStatus) GetGoName(yname string) string {
    if yname == "current-size" { return "CurrentSize" }
    if yname == "configured-size" { return "ConfiguredSize" }
    return ""
}

func (bufferStatus *Correlator_BufferStatus) GetSegmentPath() string {
    return "buffer-status"
}

func (bufferStatus *Correlator_BufferStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bufferStatus *Correlator_BufferStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bufferStatus *Correlator_BufferStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["current-size"] = bufferStatus.CurrentSize
    leafs["configured-size"] = bufferStatus.ConfiguredSize
    return leafs
}

func (bufferStatus *Correlator_BufferStatus) GetBundleName() string { return "cisco_ios_xr" }

func (bufferStatus *Correlator_BufferStatus) GetYangName() string { return "buffer-status" }

func (bufferStatus *Correlator_BufferStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bufferStatus *Correlator_BufferStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bufferStatus *Correlator_BufferStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bufferStatus *Correlator_BufferStatus) SetParent(parent types.Entity) { bufferStatus.parent = parent }

func (bufferStatus *Correlator_BufferStatus) GetParent() types.Entity { return bufferStatus.parent }

func (bufferStatus *Correlator_BufferStatus) GetParentYangName() string { return "correlator" }

// Correlator_Alarms
// Correlated alarms Table
type Correlator_Alarms struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One of the correlated alarms. The type is slice of Correlator_Alarms_Alarm.
    Alarm []Correlator_Alarms_Alarm
}

func (alarms *Correlator_Alarms) GetFilter() yfilter.YFilter { return alarms.YFilter }

func (alarms *Correlator_Alarms) SetFilter(yf yfilter.YFilter) { alarms.YFilter = yf }

func (alarms *Correlator_Alarms) GetGoName(yname string) string {
    if yname == "alarm" { return "Alarm" }
    return ""
}

func (alarms *Correlator_Alarms) GetSegmentPath() string {
    return "alarms"
}

func (alarms *Correlator_Alarms) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm" {
        for _, c := range alarms.Alarm {
            if alarms.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Correlator_Alarms_Alarm{}
        alarms.Alarm = append(alarms.Alarm, child)
        return &alarms.Alarm[len(alarms.Alarm)-1]
    }
    return nil
}

func (alarms *Correlator_Alarms) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range alarms.Alarm {
        children[alarms.Alarm[i].GetSegmentPath()] = &alarms.Alarm[i]
    }
    return children
}

func (alarms *Correlator_Alarms) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (alarms *Correlator_Alarms) GetBundleName() string { return "cisco_ios_xr" }

func (alarms *Correlator_Alarms) GetYangName() string { return "alarms" }

func (alarms *Correlator_Alarms) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarms *Correlator_Alarms) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarms *Correlator_Alarms) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarms *Correlator_Alarms) SetParent(parent types.Entity) { alarms.parent = parent }

func (alarms *Correlator_Alarms) GetParent() types.Entity { return alarms.parent }

func (alarms *Correlator_Alarms) GetParentYangName() string { return "correlator" }

// Correlator_Alarms_Alarm
// One of the correlated alarms
type Correlator_Alarms_Alarm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Alarm ID. The type is interface{} with range:
    // -2147483648..2147483647.
    AlarmId interface{}

    // Correlation rule name. The type is string.
    RuleName interface{}

    // Context string  for the alarm. The type is string.
    Context interface{}

    // Correlated alarm information.
    AlarmInfo Correlator_Alarms_Alarm_AlarmInfo
}

func (alarm *Correlator_Alarms_Alarm) GetFilter() yfilter.YFilter { return alarm.YFilter }

func (alarm *Correlator_Alarms_Alarm) SetFilter(yf yfilter.YFilter) { alarm.YFilter = yf }

func (alarm *Correlator_Alarms_Alarm) GetGoName(yname string) string {
    if yname == "alarm-id" { return "AlarmId" }
    if yname == "rule-name" { return "RuleName" }
    if yname == "context" { return "Context" }
    if yname == "alarm-info" { return "AlarmInfo" }
    return ""
}

func (alarm *Correlator_Alarms_Alarm) GetSegmentPath() string {
    return "alarm" + "[alarm-id='" + fmt.Sprintf("%v", alarm.AlarmId) + "']"
}

func (alarm *Correlator_Alarms_Alarm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "alarm-info" {
        return &alarm.AlarmInfo
    }
    return nil
}

func (alarm *Correlator_Alarms_Alarm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["alarm-info"] = &alarm.AlarmInfo
    return children
}

func (alarm *Correlator_Alarms_Alarm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["alarm-id"] = alarm.AlarmId
    leafs["rule-name"] = alarm.RuleName
    leafs["context"] = alarm.Context
    return leafs
}

func (alarm *Correlator_Alarms_Alarm) GetBundleName() string { return "cisco_ios_xr" }

func (alarm *Correlator_Alarms_Alarm) GetYangName() string { return "alarm" }

func (alarm *Correlator_Alarms_Alarm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarm *Correlator_Alarms_Alarm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarm *Correlator_Alarms_Alarm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarm *Correlator_Alarms_Alarm) SetParent(parent types.Entity) { alarm.parent = parent }

func (alarm *Correlator_Alarms_Alarm) GetParent() types.Entity { return alarm.parent }

func (alarm *Correlator_Alarms_Alarm) GetParentYangName() string { return "alarms" }

// Correlator_Alarms_Alarm_AlarmInfo
// Correlated alarm information
type Correlator_Alarms_Alarm_AlarmInfo struct {
    parent types.Entity
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

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetFilter() yfilter.YFilter { return alarmInfo.YFilter }

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) SetFilter(yf yfilter.YFilter) { alarmInfo.YFilter = yf }

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetGoName(yname string) string {
    if yname == "source-id" { return "SourceId" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "code" { return "Code" }
    if yname == "severity" { return "Severity" }
    if yname == "state" { return "State" }
    if yname == "correlation-id" { return "CorrelationId" }
    if yname == "is-admin" { return "IsAdmin" }
    if yname == "additional-text" { return "AdditionalText" }
    return ""
}

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetSegmentPath() string {
    return "alarm-info"
}

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-id"] = alarmInfo.SourceId
    leafs["timestamp"] = alarmInfo.Timestamp
    leafs["category"] = alarmInfo.Category
    leafs["group"] = alarmInfo.Group
    leafs["code"] = alarmInfo.Code
    leafs["severity"] = alarmInfo.Severity
    leafs["state"] = alarmInfo.State
    leafs["correlation-id"] = alarmInfo.CorrelationId
    leafs["is-admin"] = alarmInfo.IsAdmin
    leafs["additional-text"] = alarmInfo.AdditionalText
    return leafs
}

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetBundleName() string { return "cisco_ios_xr" }

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetYangName() string { return "alarm-info" }

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) SetParent(parent types.Entity) { alarmInfo.parent = parent }

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetParent() types.Entity { return alarmInfo.parent }

func (alarmInfo *Correlator_Alarms_Alarm_AlarmInfo) GetParentYangName() string { return "alarm" }

// Correlator_RuleSetSummaries
// Table that contains the ruleset summary info
type Correlator_RuleSetSummaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary of one of the correlation rulesets. The type is slice of
    // Correlator_RuleSetSummaries_RuleSetSummary.
    RuleSetSummary []Correlator_RuleSetSummaries_RuleSetSummary
}

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetFilter() yfilter.YFilter { return ruleSetSummaries.YFilter }

func (ruleSetSummaries *Correlator_RuleSetSummaries) SetFilter(yf yfilter.YFilter) { ruleSetSummaries.YFilter = yf }

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetGoName(yname string) string {
    if yname == "rule-set-summary" { return "RuleSetSummary" }
    return ""
}

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetSegmentPath() string {
    return "rule-set-summaries"
}

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-set-summary" {
        for _, c := range ruleSetSummaries.RuleSetSummary {
            if ruleSetSummaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Correlator_RuleSetSummaries_RuleSetSummary{}
        ruleSetSummaries.RuleSetSummary = append(ruleSetSummaries.RuleSetSummary, child)
        return &ruleSetSummaries.RuleSetSummary[len(ruleSetSummaries.RuleSetSummary)-1]
    }
    return nil
}

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleSetSummaries.RuleSetSummary {
        children[ruleSetSummaries.RuleSetSummary[i].GetSegmentPath()] = &ruleSetSummaries.RuleSetSummary[i]
    }
    return children
}

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetYangName() string { return "rule-set-summaries" }

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSetSummaries *Correlator_RuleSetSummaries) SetParent(parent types.Entity) { ruleSetSummaries.parent = parent }

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetParent() types.Entity { return ruleSetSummaries.parent }

func (ruleSetSummaries *Correlator_RuleSetSummaries) GetParentYangName() string { return "correlator" }

// Correlator_RuleSetSummaries_RuleSetSummary
// Summary of one of the correlation rulesets
type Correlator_RuleSetSummaries_RuleSetSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Ruleset Name. The type is string with length:
    // 1..32.
    RuleSetName interface{}

    // Ruleset Name. The type is string.
    RuleSetNameXr interface{}
}

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetFilter() yfilter.YFilter { return ruleSetSummary.YFilter }

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) SetFilter(yf yfilter.YFilter) { ruleSetSummary.YFilter = yf }

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetGoName(yname string) string {
    if yname == "rule-set-name" { return "RuleSetName" }
    if yname == "rule-set-name-xr" { return "RuleSetNameXr" }
    return ""
}

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetSegmentPath() string {
    return "rule-set-summary" + "[rule-set-name='" + fmt.Sprintf("%v", ruleSetSummary.RuleSetName) + "']"
}

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-set-name"] = ruleSetSummary.RuleSetName
    leafs["rule-set-name-xr"] = ruleSetSummary.RuleSetNameXr
    return leafs
}

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetYangName() string { return "rule-set-summary" }

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) SetParent(parent types.Entity) { ruleSetSummary.parent = parent }

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetParent() types.Entity { return ruleSetSummary.parent }

func (ruleSetSummary *Correlator_RuleSetSummaries_RuleSetSummary) GetParentYangName() string { return "rule-set-summaries" }

// Correlator_RuleSetDetails
// Table that contains the ruleset detail info
type Correlator_RuleSetDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detail of one of the correlation rulesets. The type is slice of
    // Correlator_RuleSetDetails_RuleSetDetail.
    RuleSetDetail []Correlator_RuleSetDetails_RuleSetDetail
}

func (ruleSetDetails *Correlator_RuleSetDetails) GetFilter() yfilter.YFilter { return ruleSetDetails.YFilter }

func (ruleSetDetails *Correlator_RuleSetDetails) SetFilter(yf yfilter.YFilter) { ruleSetDetails.YFilter = yf }

func (ruleSetDetails *Correlator_RuleSetDetails) GetGoName(yname string) string {
    if yname == "rule-set-detail" { return "RuleSetDetail" }
    return ""
}

func (ruleSetDetails *Correlator_RuleSetDetails) GetSegmentPath() string {
    return "rule-set-details"
}

func (ruleSetDetails *Correlator_RuleSetDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-set-detail" {
        for _, c := range ruleSetDetails.RuleSetDetail {
            if ruleSetDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Correlator_RuleSetDetails_RuleSetDetail{}
        ruleSetDetails.RuleSetDetail = append(ruleSetDetails.RuleSetDetail, child)
        return &ruleSetDetails.RuleSetDetail[len(ruleSetDetails.RuleSetDetail)-1]
    }
    return nil
}

func (ruleSetDetails *Correlator_RuleSetDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleSetDetails.RuleSetDetail {
        children[ruleSetDetails.RuleSetDetail[i].GetSegmentPath()] = &ruleSetDetails.RuleSetDetail[i]
    }
    return children
}

func (ruleSetDetails *Correlator_RuleSetDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleSetDetails *Correlator_RuleSetDetails) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSetDetails *Correlator_RuleSetDetails) GetYangName() string { return "rule-set-details" }

func (ruleSetDetails *Correlator_RuleSetDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSetDetails *Correlator_RuleSetDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSetDetails *Correlator_RuleSetDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSetDetails *Correlator_RuleSetDetails) SetParent(parent types.Entity) { ruleSetDetails.parent = parent }

func (ruleSetDetails *Correlator_RuleSetDetails) GetParent() types.Entity { return ruleSetDetails.parent }

func (ruleSetDetails *Correlator_RuleSetDetails) GetParentYangName() string { return "correlator" }

// Correlator_RuleSetDetails_RuleSetDetail
// Detail of one of the correlation rulesets
type Correlator_RuleSetDetails_RuleSetDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Ruleset Name. The type is string with length:
    // 1..32.
    RuleSetName interface{}

    // Ruleset Name. The type is string.
    RuleSetNameXr interface{}

    // Rules contained in a ruleset. The type is slice of
    // Correlator_RuleSetDetails_RuleSetDetail_Rules.
    Rules []Correlator_RuleSetDetails_RuleSetDetail_Rules
}

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetFilter() yfilter.YFilter { return ruleSetDetail.YFilter }

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) SetFilter(yf yfilter.YFilter) { ruleSetDetail.YFilter = yf }

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetGoName(yname string) string {
    if yname == "rule-set-name" { return "RuleSetName" }
    if yname == "rule-set-name-xr" { return "RuleSetNameXr" }
    if yname == "rules" { return "Rules" }
    return ""
}

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetSegmentPath() string {
    return "rule-set-detail" + "[rule-set-name='" + fmt.Sprintf("%v", ruleSetDetail.RuleSetName) + "']"
}

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rules" {
        for _, c := range ruleSetDetail.Rules {
            if ruleSetDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Correlator_RuleSetDetails_RuleSetDetail_Rules{}
        ruleSetDetail.Rules = append(ruleSetDetail.Rules, child)
        return &ruleSetDetail.Rules[len(ruleSetDetail.Rules)-1]
    }
    return nil
}

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleSetDetail.Rules {
        children[ruleSetDetail.Rules[i].GetSegmentPath()] = &ruleSetDetail.Rules[i]
    }
    return children
}

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-set-name"] = ruleSetDetail.RuleSetName
    leafs["rule-set-name-xr"] = ruleSetDetail.RuleSetNameXr
    return leafs
}

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetYangName() string { return "rule-set-detail" }

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) SetParent(parent types.Entity) { ruleSetDetail.parent = parent }

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetParent() types.Entity { return ruleSetDetail.parent }

func (ruleSetDetail *Correlator_RuleSetDetails_RuleSetDetail) GetParentYangName() string { return "rule-set-details" }

// Correlator_RuleSetDetails_RuleSetDetail_Rules
// Rules contained in a ruleset
type Correlator_RuleSetDetails_RuleSetDetail_Rules struct {
    parent types.Entity
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

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetFilter() yfilter.YFilter { return rules.YFilter }

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) SetFilter(yf yfilter.YFilter) { rules.YFilter = yf }

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetGoName(yname string) string {
    if yname == "rule-name-xr" { return "RuleNameXr" }
    if yname == "stateful" { return "Stateful" }
    if yname == "rule-state" { return "RuleState" }
    if yname == "buffered-alarms-count" { return "BufferedAlarmsCount" }
    return ""
}

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetSegmentPath() string {
    return "rules"
}

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name-xr"] = rules.RuleNameXr
    leafs["stateful"] = rules.Stateful
    leafs["rule-state"] = rules.RuleState
    leafs["buffered-alarms-count"] = rules.BufferedAlarmsCount
    return leafs
}

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetBundleName() string { return "cisco_ios_xr" }

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetYangName() string { return "rules" }

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) SetParent(parent types.Entity) { rules.parent = parent }

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetParent() types.Entity { return rules.parent }

func (rules *Correlator_RuleSetDetails_RuleSetDetail_Rules) GetParentYangName() string { return "rule-set-detail" }

// Correlator_RuleDetails
// Table that contains the database of correlation
// rule details
type Correlator_RuleDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details of one of the correlation rules. The type is slice of
    // Correlator_RuleDetails_RuleDetail.
    RuleDetail []Correlator_RuleDetails_RuleDetail
}

func (ruleDetails *Correlator_RuleDetails) GetFilter() yfilter.YFilter { return ruleDetails.YFilter }

func (ruleDetails *Correlator_RuleDetails) SetFilter(yf yfilter.YFilter) { ruleDetails.YFilter = yf }

func (ruleDetails *Correlator_RuleDetails) GetGoName(yname string) string {
    if yname == "rule-detail" { return "RuleDetail" }
    return ""
}

func (ruleDetails *Correlator_RuleDetails) GetSegmentPath() string {
    return "rule-details"
}

func (ruleDetails *Correlator_RuleDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-detail" {
        for _, c := range ruleDetails.RuleDetail {
            if ruleDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Correlator_RuleDetails_RuleDetail{}
        ruleDetails.RuleDetail = append(ruleDetails.RuleDetail, child)
        return &ruleDetails.RuleDetail[len(ruleDetails.RuleDetail)-1]
    }
    return nil
}

func (ruleDetails *Correlator_RuleDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleDetails.RuleDetail {
        children[ruleDetails.RuleDetail[i].GetSegmentPath()] = &ruleDetails.RuleDetail[i]
    }
    return children
}

func (ruleDetails *Correlator_RuleDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleDetails *Correlator_RuleDetails) GetBundleName() string { return "cisco_ios_xr" }

func (ruleDetails *Correlator_RuleDetails) GetYangName() string { return "rule-details" }

func (ruleDetails *Correlator_RuleDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleDetails *Correlator_RuleDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleDetails *Correlator_RuleDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleDetails *Correlator_RuleDetails) SetParent(parent types.Entity) { ruleDetails.parent = parent }

func (ruleDetails *Correlator_RuleDetails) GetParent() types.Entity { return ruleDetails.parent }

func (ruleDetails *Correlator_RuleDetails) GetParentYangName() string { return "correlator" }

// Correlator_RuleDetails_RuleDetail
// Details of one of the correlation rules
type Correlator_RuleDetails_RuleDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    Codes []Correlator_RuleDetails_RuleDetail_Codes
}

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetFilter() yfilter.YFilter { return ruleDetail.YFilter }

func (ruleDetail *Correlator_RuleDetails_RuleDetail) SetFilter(yf yfilter.YFilter) { ruleDetail.YFilter = yf }

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetGoName(yname string) string {
    if yname == "rule-name" { return "RuleName" }
    if yname == "timeout" { return "Timeout" }
    if yname == "root-cause-timeout" { return "RootCauseTimeout" }
    if yname == "internal" { return "Internal" }
    if yname == "reissue-non-bistate" { return "ReissueNonBistate" }
    if yname == "reparent" { return "Reparent" }
    if yname == "context-correlation" { return "ContextCorrelation" }
    if yname == "apply-location" { return "ApplyLocation" }
    if yname == "apply-context" { return "ApplyContext" }
    if yname == "rule-summary" { return "RuleSummary" }
    if yname == "codes" { return "Codes" }
    return ""
}

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetSegmentPath() string {
    return "rule-detail" + "[rule-name='" + fmt.Sprintf("%v", ruleDetail.RuleName) + "']"
}

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-summary" {
        return &ruleDetail.RuleSummary
    }
    if childYangName == "codes" {
        for _, c := range ruleDetail.Codes {
            if ruleDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Correlator_RuleDetails_RuleDetail_Codes{}
        ruleDetail.Codes = append(ruleDetail.Codes, child)
        return &ruleDetail.Codes[len(ruleDetail.Codes)-1]
    }
    return nil
}

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rule-summary"] = &ruleDetail.RuleSummary
    for i := range ruleDetail.Codes {
        children[ruleDetail.Codes[i].GetSegmentPath()] = &ruleDetail.Codes[i]
    }
    return children
}

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name"] = ruleDetail.RuleName
    leafs["timeout"] = ruleDetail.Timeout
    leafs["root-cause-timeout"] = ruleDetail.RootCauseTimeout
    leafs["internal"] = ruleDetail.Internal
    leafs["reissue-non-bistate"] = ruleDetail.ReissueNonBistate
    leafs["reparent"] = ruleDetail.Reparent
    leafs["context-correlation"] = ruleDetail.ContextCorrelation
    leafs["apply-location"] = ruleDetail.ApplyLocation
    leafs["apply-context"] = ruleDetail.ApplyContext
    return leafs
}

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetBundleName() string { return "cisco_ios_xr" }

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetYangName() string { return "rule-detail" }

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleDetail *Correlator_RuleDetails_RuleDetail) SetParent(parent types.Entity) { ruleDetail.parent = parent }

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetParent() types.Entity { return ruleDetail.parent }

func (ruleDetail *Correlator_RuleDetails_RuleDetail) GetParentYangName() string { return "rule-details" }

// Correlator_RuleDetails_RuleDetail_RuleSummary
// Rule summary, name, etc
type Correlator_RuleDetails_RuleDetail_RuleSummary struct {
    parent types.Entity
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

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetFilter() yfilter.YFilter { return ruleSummary.YFilter }

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) SetFilter(yf yfilter.YFilter) { ruleSummary.YFilter = yf }

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetGoName(yname string) string {
    if yname == "rule-name-xr" { return "RuleNameXr" }
    if yname == "stateful" { return "Stateful" }
    if yname == "rule-state" { return "RuleState" }
    if yname == "buffered-alarms-count" { return "BufferedAlarmsCount" }
    return ""
}

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetSegmentPath() string {
    return "rule-summary"
}

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name-xr"] = ruleSummary.RuleNameXr
    leafs["stateful"] = ruleSummary.Stateful
    leafs["rule-state"] = ruleSummary.RuleState
    leafs["buffered-alarms-count"] = ruleSummary.BufferedAlarmsCount
    return leafs
}

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetYangName() string { return "rule-summary" }

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) SetParent(parent types.Entity) { ruleSummary.parent = parent }

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetParent() types.Entity { return ruleSummary.parent }

func (ruleSummary *Correlator_RuleDetails_RuleDetail_RuleSummary) GetParentYangName() string { return "rule-detail" }

// Correlator_RuleDetails_RuleDetail_Codes
// Message codes defining the rule.
type Correlator_RuleDetails_RuleDetail_Codes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Category of messages to which this alarm belongs. The type is string.
    Category interface{}

    // Group of messages to which this alarm belongs. The type is string.
    Group interface{}

    // Alarm code which further qualifies the alarm within a message group. The
    // type is string.
    Code interface{}
}

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetFilter() yfilter.YFilter { return codes.YFilter }

func (codes *Correlator_RuleDetails_RuleDetail_Codes) SetFilter(yf yfilter.YFilter) { codes.YFilter = yf }

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetGoName(yname string) string {
    if yname == "category" { return "Category" }
    if yname == "group" { return "Group" }
    if yname == "code" { return "Code" }
    return ""
}

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetSegmentPath() string {
    return "codes"
}

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["category"] = codes.Category
    leafs["group"] = codes.Group
    leafs["code"] = codes.Code
    return leafs
}

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetBundleName() string { return "cisco_ios_xr" }

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetYangName() string { return "codes" }

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (codes *Correlator_RuleDetails_RuleDetail_Codes) SetParent(parent types.Entity) { codes.parent = parent }

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetParent() types.Entity { return codes.parent }

func (codes *Correlator_RuleDetails_RuleDetail_Codes) GetParentYangName() string { return "rule-detail" }

// Correlator_RuleSummaries
// Table that contains the database of correlation
// rule summary
type Correlator_RuleSummaries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // One of the correlation rules. The type is slice of
    // Correlator_RuleSummaries_RuleSummary.
    RuleSummary []Correlator_RuleSummaries_RuleSummary
}

func (ruleSummaries *Correlator_RuleSummaries) GetFilter() yfilter.YFilter { return ruleSummaries.YFilter }

func (ruleSummaries *Correlator_RuleSummaries) SetFilter(yf yfilter.YFilter) { ruleSummaries.YFilter = yf }

func (ruleSummaries *Correlator_RuleSummaries) GetGoName(yname string) string {
    if yname == "rule-summary" { return "RuleSummary" }
    return ""
}

func (ruleSummaries *Correlator_RuleSummaries) GetSegmentPath() string {
    return "rule-summaries"
}

func (ruleSummaries *Correlator_RuleSummaries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-summary" {
        for _, c := range ruleSummaries.RuleSummary {
            if ruleSummaries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Correlator_RuleSummaries_RuleSummary{}
        ruleSummaries.RuleSummary = append(ruleSummaries.RuleSummary, child)
        return &ruleSummaries.RuleSummary[len(ruleSummaries.RuleSummary)-1]
    }
    return nil
}

func (ruleSummaries *Correlator_RuleSummaries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleSummaries.RuleSummary {
        children[ruleSummaries.RuleSummary[i].GetSegmentPath()] = &ruleSummaries.RuleSummary[i]
    }
    return children
}

func (ruleSummaries *Correlator_RuleSummaries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleSummaries *Correlator_RuleSummaries) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSummaries *Correlator_RuleSummaries) GetYangName() string { return "rule-summaries" }

func (ruleSummaries *Correlator_RuleSummaries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSummaries *Correlator_RuleSummaries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSummaries *Correlator_RuleSummaries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSummaries *Correlator_RuleSummaries) SetParent(parent types.Entity) { ruleSummaries.parent = parent }

func (ruleSummaries *Correlator_RuleSummaries) GetParent() types.Entity { return ruleSummaries.parent }

func (ruleSummaries *Correlator_RuleSummaries) GetParentYangName() string { return "correlator" }

// Correlator_RuleSummaries_RuleSummary
// One of the correlation rules
type Correlator_RuleSummaries_RuleSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetFilter() yfilter.YFilter { return ruleSummary.YFilter }

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) SetFilter(yf yfilter.YFilter) { ruleSummary.YFilter = yf }

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetGoName(yname string) string {
    if yname == "rule-name" { return "RuleName" }
    if yname == "rule-name-xr" { return "RuleNameXr" }
    if yname == "stateful" { return "Stateful" }
    if yname == "rule-state" { return "RuleState" }
    if yname == "buffered-alarms-count" { return "BufferedAlarmsCount" }
    return ""
}

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetSegmentPath() string {
    return "rule-summary" + "[rule-name='" + fmt.Sprintf("%v", ruleSummary.RuleName) + "']"
}

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rule-name"] = ruleSummary.RuleName
    leafs["rule-name-xr"] = ruleSummary.RuleNameXr
    leafs["stateful"] = ruleSummary.Stateful
    leafs["rule-state"] = ruleSummary.RuleState
    leafs["buffered-alarms-count"] = ruleSummary.BufferedAlarmsCount
    return leafs
}

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetYangName() string { return "rule-summary" }

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) SetParent(parent types.Entity) { ruleSummary.parent = parent }

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetParent() types.Entity { return ruleSummary.parent }

func (ruleSummary *Correlator_RuleSummaries_RuleSummary) GetParentYangName() string { return "rule-summaries" }

