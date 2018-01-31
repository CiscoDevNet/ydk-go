// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2016 by Cisco Systems, Inc.
// All rights reserved.
package cfgmgr_rollback_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cfgmgr_rollback_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cfgmgr-rollback-act roll-back-configuration-last}", reflect.TypeOf(RollBackConfigurationLast{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-last", reflect.TypeOf(RollBackConfigurationLast{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cfgmgr-rollback-act roll-back-configuration-to}", reflect.TypeOf(RollBackConfigurationTo{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-to", reflect.TypeOf(RollBackConfigurationTo{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cfgmgr-rollback-act roll-back-configuration-to-exclude}", reflect.TypeOf(RollBackConfigurationToExclude{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-to-exclude", reflect.TypeOf(RollBackConfigurationToExclude{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cfgmgr-rollback-act roll-back-configuration}", reflect.TypeOf(RollBackConfiguration{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration", reflect.TypeOf(RollBackConfiguration{}))
}

// RollBackConfigurationLast
// Rollback last <n> commits made
type RollBackConfigurationLast struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RollBackConfigurationLast_Input
}

func (rollBackConfigurationLast *RollBackConfigurationLast) GetFilter() yfilter.YFilter { return rollBackConfigurationLast.YFilter }

func (rollBackConfigurationLast *RollBackConfigurationLast) SetFilter(yf yfilter.YFilter) { rollBackConfigurationLast.YFilter = yf }

func (rollBackConfigurationLast *RollBackConfigurationLast) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (rollBackConfigurationLast *RollBackConfigurationLast) GetSegmentPath() string {
    return "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-last"
}

func (rollBackConfigurationLast *RollBackConfigurationLast) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &rollBackConfigurationLast.Input
    }
    return nil
}

func (rollBackConfigurationLast *RollBackConfigurationLast) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &rollBackConfigurationLast.Input
    return children
}

func (rollBackConfigurationLast *RollBackConfigurationLast) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rollBackConfigurationLast *RollBackConfigurationLast) GetBundleName() string { return "cisco_ios_xr" }

func (rollBackConfigurationLast *RollBackConfigurationLast) GetYangName() string { return "roll-back-configuration-last" }

func (rollBackConfigurationLast *RollBackConfigurationLast) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rollBackConfigurationLast *RollBackConfigurationLast) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rollBackConfigurationLast *RollBackConfigurationLast) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rollBackConfigurationLast *RollBackConfigurationLast) SetParent(parent types.Entity) { rollBackConfigurationLast.parent = parent }

func (rollBackConfigurationLast *RollBackConfigurationLast) GetParent() types.Entity { return rollBackConfigurationLast.parent }

func (rollBackConfigurationLast *RollBackConfigurationLast) GetParentYangName() string { return "Cisco-IOS-XR-cfgmgr-rollback-act" }

// RollBackConfigurationLast_Input
type RollBackConfigurationLast_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of commits to rollback. The type is interface{} with range: 1..100.
    // This attribute is mandatory.
    Count interface{}

    // Override commit blocks. The type is bool. The default value is false.
    Force interface{}

    // Rollback via best-effort operation. The type is bool. The default value is
    // false.
    BestEffort interface{}

    // Assign a label to this rollback. The type is string.
    Label interface{}

    // Assign a comment to this rollback. The type is string.
    Comment interface{}
}

func (input *RollBackConfigurationLast_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RollBackConfigurationLast_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RollBackConfigurationLast_Input) GetGoName(yname string) string {
    if yname == "count" { return "Count" }
    if yname == "force" { return "Force" }
    if yname == "best-effort" { return "BestEffort" }
    if yname == "label" { return "Label" }
    if yname == "comment" { return "Comment" }
    return ""
}

func (input *RollBackConfigurationLast_Input) GetSegmentPath() string {
    return "input"
}

func (input *RollBackConfigurationLast_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RollBackConfigurationLast_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RollBackConfigurationLast_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["count"] = input.Count
    leafs["force"] = input.Force
    leafs["best-effort"] = input.BestEffort
    leafs["label"] = input.Label
    leafs["comment"] = input.Comment
    return leafs
}

func (input *RollBackConfigurationLast_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RollBackConfigurationLast_Input) GetYangName() string { return "input" }

func (input *RollBackConfigurationLast_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RollBackConfigurationLast_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RollBackConfigurationLast_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RollBackConfigurationLast_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RollBackConfigurationLast_Input) GetParent() types.Entity { return input.parent }

func (input *RollBackConfigurationLast_Input) GetParentYangName() string { return "roll-back-configuration-last" }

// RollBackConfigurationTo
// Rollback up to (and including) a specific commit
type RollBackConfigurationTo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RollBackConfigurationTo_Input
}

func (rollBackConfigurationTo *RollBackConfigurationTo) GetFilter() yfilter.YFilter { return rollBackConfigurationTo.YFilter }

func (rollBackConfigurationTo *RollBackConfigurationTo) SetFilter(yf yfilter.YFilter) { rollBackConfigurationTo.YFilter = yf }

func (rollBackConfigurationTo *RollBackConfigurationTo) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (rollBackConfigurationTo *RollBackConfigurationTo) GetSegmentPath() string {
    return "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-to"
}

func (rollBackConfigurationTo *RollBackConfigurationTo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &rollBackConfigurationTo.Input
    }
    return nil
}

func (rollBackConfigurationTo *RollBackConfigurationTo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &rollBackConfigurationTo.Input
    return children
}

func (rollBackConfigurationTo *RollBackConfigurationTo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rollBackConfigurationTo *RollBackConfigurationTo) GetBundleName() string { return "cisco_ios_xr" }

func (rollBackConfigurationTo *RollBackConfigurationTo) GetYangName() string { return "roll-back-configuration-to" }

func (rollBackConfigurationTo *RollBackConfigurationTo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rollBackConfigurationTo *RollBackConfigurationTo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rollBackConfigurationTo *RollBackConfigurationTo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rollBackConfigurationTo *RollBackConfigurationTo) SetParent(parent types.Entity) { rollBackConfigurationTo.parent = parent }

func (rollBackConfigurationTo *RollBackConfigurationTo) GetParent() types.Entity { return rollBackConfigurationTo.parent }

func (rollBackConfigurationTo *RollBackConfigurationTo) GetParentYangName() string { return "Cisco-IOS-XR-cfgmgr-rollback-act" }

// RollBackConfigurationTo_Input
type RollBackConfigurationTo_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Commit ID. The type is string. This attribute is mandatory.
    CommitId interface{}

    // Override commit blocks. The type is bool. The default value is false.
    Force interface{}

    // Rollback via best-effort operation. The type is bool. The default value is
    // false.
    BestEffort interface{}

    // Assign a label to this rollback. The type is string.
    Label interface{}

    // Assign a comment to this rollback. The type is string.
    Comment interface{}
}

func (input *RollBackConfigurationTo_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RollBackConfigurationTo_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RollBackConfigurationTo_Input) GetGoName(yname string) string {
    if yname == "commit-id" { return "CommitId" }
    if yname == "force" { return "Force" }
    if yname == "best-effort" { return "BestEffort" }
    if yname == "label" { return "Label" }
    if yname == "comment" { return "Comment" }
    return ""
}

func (input *RollBackConfigurationTo_Input) GetSegmentPath() string {
    return "input"
}

func (input *RollBackConfigurationTo_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RollBackConfigurationTo_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RollBackConfigurationTo_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["commit-id"] = input.CommitId
    leafs["force"] = input.Force
    leafs["best-effort"] = input.BestEffort
    leafs["label"] = input.Label
    leafs["comment"] = input.Comment
    return leafs
}

func (input *RollBackConfigurationTo_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RollBackConfigurationTo_Input) GetYangName() string { return "input" }

func (input *RollBackConfigurationTo_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RollBackConfigurationTo_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RollBackConfigurationTo_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RollBackConfigurationTo_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RollBackConfigurationTo_Input) GetParent() types.Entity { return input.parent }

func (input *RollBackConfigurationTo_Input) GetParentYangName() string { return "roll-back-configuration-to" }

// RollBackConfigurationToExclude
// Rollback up to (and excluding) a specific commit
type RollBackConfigurationToExclude struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RollBackConfigurationToExclude_Input
}

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetFilter() yfilter.YFilter { return rollBackConfigurationToExclude.YFilter }

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) SetFilter(yf yfilter.YFilter) { rollBackConfigurationToExclude.YFilter = yf }

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetSegmentPath() string {
    return "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-to-exclude"
}

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &rollBackConfigurationToExclude.Input
    }
    return nil
}

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &rollBackConfigurationToExclude.Input
    return children
}

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetBundleName() string { return "cisco_ios_xr" }

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetYangName() string { return "roll-back-configuration-to-exclude" }

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) SetParent(parent types.Entity) { rollBackConfigurationToExclude.parent = parent }

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetParent() types.Entity { return rollBackConfigurationToExclude.parent }

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetParentYangName() string { return "Cisco-IOS-XR-cfgmgr-rollback-act" }

// RollBackConfigurationToExclude_Input
type RollBackConfigurationToExclude_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Commit ID. The type is string. This attribute is mandatory.
    CommitId interface{}

    // Override commit blocks. The type is bool. The default value is false.
    Force interface{}

    // Rollback via best-effort operation. The type is bool. The default value is
    // false.
    BestEffort interface{}

    // Assign a label to this rollback. The type is string.
    Label interface{}

    // Assign a comment to this rollback. The type is string.
    Comment interface{}
}

func (input *RollBackConfigurationToExclude_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RollBackConfigurationToExclude_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RollBackConfigurationToExclude_Input) GetGoName(yname string) string {
    if yname == "commit-id" { return "CommitId" }
    if yname == "force" { return "Force" }
    if yname == "best-effort" { return "BestEffort" }
    if yname == "label" { return "Label" }
    if yname == "comment" { return "Comment" }
    return ""
}

func (input *RollBackConfigurationToExclude_Input) GetSegmentPath() string {
    return "input"
}

func (input *RollBackConfigurationToExclude_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RollBackConfigurationToExclude_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RollBackConfigurationToExclude_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["commit-id"] = input.CommitId
    leafs["force"] = input.Force
    leafs["best-effort"] = input.BestEffort
    leafs["label"] = input.Label
    leafs["comment"] = input.Comment
    return leafs
}

func (input *RollBackConfigurationToExclude_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RollBackConfigurationToExclude_Input) GetYangName() string { return "input" }

func (input *RollBackConfigurationToExclude_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RollBackConfigurationToExclude_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RollBackConfigurationToExclude_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RollBackConfigurationToExclude_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RollBackConfigurationToExclude_Input) GetParent() types.Entity { return input.parent }

func (input *RollBackConfigurationToExclude_Input) GetParentYangName() string { return "roll-back-configuration-to-exclude" }

// RollBackConfiguration
// Rollback a specific commit
type RollBackConfiguration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RollBackConfiguration_Input
}

func (rollBackConfiguration *RollBackConfiguration) GetFilter() yfilter.YFilter { return rollBackConfiguration.YFilter }

func (rollBackConfiguration *RollBackConfiguration) SetFilter(yf yfilter.YFilter) { rollBackConfiguration.YFilter = yf }

func (rollBackConfiguration *RollBackConfiguration) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (rollBackConfiguration *RollBackConfiguration) GetSegmentPath() string {
    return "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration"
}

func (rollBackConfiguration *RollBackConfiguration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &rollBackConfiguration.Input
    }
    return nil
}

func (rollBackConfiguration *RollBackConfiguration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &rollBackConfiguration.Input
    return children
}

func (rollBackConfiguration *RollBackConfiguration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rollBackConfiguration *RollBackConfiguration) GetBundleName() string { return "cisco_ios_xr" }

func (rollBackConfiguration *RollBackConfiguration) GetYangName() string { return "roll-back-configuration" }

func (rollBackConfiguration *RollBackConfiguration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rollBackConfiguration *RollBackConfiguration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rollBackConfiguration *RollBackConfiguration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rollBackConfiguration *RollBackConfiguration) SetParent(parent types.Entity) { rollBackConfiguration.parent = parent }

func (rollBackConfiguration *RollBackConfiguration) GetParent() types.Entity { return rollBackConfiguration.parent }

func (rollBackConfiguration *RollBackConfiguration) GetParentYangName() string { return "Cisco-IOS-XR-cfgmgr-rollback-act" }

// RollBackConfiguration_Input
type RollBackConfiguration_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Commit ID. The type is string. This attribute is mandatory.
    CommitId interface{}

    // Override commit blocks. The type is bool. The default value is false.
    Force interface{}

    // Rollback via best-effort operation. The type is bool. The default value is
    // false.
    BestEffort interface{}

    // Assign a label to this rollback. The type is string.
    Label interface{}

    // Assign a comment to this rollback. The type is string.
    Comment interface{}
}

func (input *RollBackConfiguration_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RollBackConfiguration_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RollBackConfiguration_Input) GetGoName(yname string) string {
    if yname == "commit-id" { return "CommitId" }
    if yname == "force" { return "Force" }
    if yname == "best-effort" { return "BestEffort" }
    if yname == "label" { return "Label" }
    if yname == "comment" { return "Comment" }
    return ""
}

func (input *RollBackConfiguration_Input) GetSegmentPath() string {
    return "input"
}

func (input *RollBackConfiguration_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RollBackConfiguration_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RollBackConfiguration_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["commit-id"] = input.CommitId
    leafs["force"] = input.Force
    leafs["best-effort"] = input.BestEffort
    leafs["label"] = input.Label
    leafs["comment"] = input.Comment
    return leafs
}

func (input *RollBackConfiguration_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *RollBackConfiguration_Input) GetYangName() string { return "input" }

func (input *RollBackConfiguration_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *RollBackConfiguration_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *RollBackConfiguration_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *RollBackConfiguration_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RollBackConfiguration_Input) GetParent() types.Entity { return input.parent }

func (input *RollBackConfiguration_Input) GetParentYangName() string { return "roll-back-configuration" }

