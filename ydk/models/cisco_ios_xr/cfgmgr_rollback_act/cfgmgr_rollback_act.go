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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RollBackConfigurationLast_Input
}

func (rollBackConfigurationLast *RollBackConfigurationLast) GetEntityData() *types.CommonEntityData {
    rollBackConfigurationLast.EntityData.YFilter = rollBackConfigurationLast.YFilter
    rollBackConfigurationLast.EntityData.YangName = "roll-back-configuration-last"
    rollBackConfigurationLast.EntityData.BundleName = "cisco_ios_xr"
    rollBackConfigurationLast.EntityData.ParentYangName = "Cisco-IOS-XR-cfgmgr-rollback-act"
    rollBackConfigurationLast.EntityData.SegmentPath = "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-last"
    rollBackConfigurationLast.EntityData.AbsolutePath = rollBackConfigurationLast.EntityData.SegmentPath
    rollBackConfigurationLast.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rollBackConfigurationLast.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rollBackConfigurationLast.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rollBackConfigurationLast.EntityData.Children = types.NewOrderedMap()
    rollBackConfigurationLast.EntityData.Children.Append("input", types.YChild{"Input", &rollBackConfigurationLast.Input})
    rollBackConfigurationLast.EntityData.Leafs = types.NewOrderedMap()

    rollBackConfigurationLast.EntityData.YListKeys = []string {}

    return &(rollBackConfigurationLast.EntityData)
}

// RollBackConfigurationLast_Input
type RollBackConfigurationLast_Input struct {
    EntityData types.CommonEntityData
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

func (input *RollBackConfigurationLast_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "roll-back-configuration-last"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-last/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("count", types.YLeaf{"Count", input.Count})
    input.EntityData.Leafs.Append("force", types.YLeaf{"Force", input.Force})
    input.EntityData.Leafs.Append("best-effort", types.YLeaf{"BestEffort", input.BestEffort})
    input.EntityData.Leafs.Append("label", types.YLeaf{"Label", input.Label})
    input.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", input.Comment})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RollBackConfigurationTo
// Rollback up to (and including) a specific commit
type RollBackConfigurationTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RollBackConfigurationTo_Input
}

func (rollBackConfigurationTo *RollBackConfigurationTo) GetEntityData() *types.CommonEntityData {
    rollBackConfigurationTo.EntityData.YFilter = rollBackConfigurationTo.YFilter
    rollBackConfigurationTo.EntityData.YangName = "roll-back-configuration-to"
    rollBackConfigurationTo.EntityData.BundleName = "cisco_ios_xr"
    rollBackConfigurationTo.EntityData.ParentYangName = "Cisco-IOS-XR-cfgmgr-rollback-act"
    rollBackConfigurationTo.EntityData.SegmentPath = "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-to"
    rollBackConfigurationTo.EntityData.AbsolutePath = rollBackConfigurationTo.EntityData.SegmentPath
    rollBackConfigurationTo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rollBackConfigurationTo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rollBackConfigurationTo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rollBackConfigurationTo.EntityData.Children = types.NewOrderedMap()
    rollBackConfigurationTo.EntityData.Children.Append("input", types.YChild{"Input", &rollBackConfigurationTo.Input})
    rollBackConfigurationTo.EntityData.Leafs = types.NewOrderedMap()

    rollBackConfigurationTo.EntityData.YListKeys = []string {}

    return &(rollBackConfigurationTo.EntityData)
}

// RollBackConfigurationTo_Input
type RollBackConfigurationTo_Input struct {
    EntityData types.CommonEntityData
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

func (input *RollBackConfigurationTo_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "roll-back-configuration-to"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-to/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("commit-id", types.YLeaf{"CommitId", input.CommitId})
    input.EntityData.Leafs.Append("force", types.YLeaf{"Force", input.Force})
    input.EntityData.Leafs.Append("best-effort", types.YLeaf{"BestEffort", input.BestEffort})
    input.EntityData.Leafs.Append("label", types.YLeaf{"Label", input.Label})
    input.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", input.Comment})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RollBackConfigurationToExclude
// Rollback up to (and excluding) a specific commit
type RollBackConfigurationToExclude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RollBackConfigurationToExclude_Input
}

func (rollBackConfigurationToExclude *RollBackConfigurationToExclude) GetEntityData() *types.CommonEntityData {
    rollBackConfigurationToExclude.EntityData.YFilter = rollBackConfigurationToExclude.YFilter
    rollBackConfigurationToExclude.EntityData.YangName = "roll-back-configuration-to-exclude"
    rollBackConfigurationToExclude.EntityData.BundleName = "cisco_ios_xr"
    rollBackConfigurationToExclude.EntityData.ParentYangName = "Cisco-IOS-XR-cfgmgr-rollback-act"
    rollBackConfigurationToExclude.EntityData.SegmentPath = "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-to-exclude"
    rollBackConfigurationToExclude.EntityData.AbsolutePath = rollBackConfigurationToExclude.EntityData.SegmentPath
    rollBackConfigurationToExclude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rollBackConfigurationToExclude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rollBackConfigurationToExclude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rollBackConfigurationToExclude.EntityData.Children = types.NewOrderedMap()
    rollBackConfigurationToExclude.EntityData.Children.Append("input", types.YChild{"Input", &rollBackConfigurationToExclude.Input})
    rollBackConfigurationToExclude.EntityData.Leafs = types.NewOrderedMap()

    rollBackConfigurationToExclude.EntityData.YListKeys = []string {}

    return &(rollBackConfigurationToExclude.EntityData)
}

// RollBackConfigurationToExclude_Input
type RollBackConfigurationToExclude_Input struct {
    EntityData types.CommonEntityData
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

func (input *RollBackConfigurationToExclude_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "roll-back-configuration-to-exclude"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration-to-exclude/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("commit-id", types.YLeaf{"CommitId", input.CommitId})
    input.EntityData.Leafs.Append("force", types.YLeaf{"Force", input.Force})
    input.EntityData.Leafs.Append("best-effort", types.YLeaf{"BestEffort", input.BestEffort})
    input.EntityData.Leafs.Append("label", types.YLeaf{"Label", input.Label})
    input.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", input.Comment})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RollBackConfiguration
// Rollback a specific commit
type RollBackConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RollBackConfiguration_Input
}

func (rollBackConfiguration *RollBackConfiguration) GetEntityData() *types.CommonEntityData {
    rollBackConfiguration.EntityData.YFilter = rollBackConfiguration.YFilter
    rollBackConfiguration.EntityData.YangName = "roll-back-configuration"
    rollBackConfiguration.EntityData.BundleName = "cisco_ios_xr"
    rollBackConfiguration.EntityData.ParentYangName = "Cisco-IOS-XR-cfgmgr-rollback-act"
    rollBackConfiguration.EntityData.SegmentPath = "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration"
    rollBackConfiguration.EntityData.AbsolutePath = rollBackConfiguration.EntityData.SegmentPath
    rollBackConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rollBackConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rollBackConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rollBackConfiguration.EntityData.Children = types.NewOrderedMap()
    rollBackConfiguration.EntityData.Children.Append("input", types.YChild{"Input", &rollBackConfiguration.Input})
    rollBackConfiguration.EntityData.Leafs = types.NewOrderedMap()

    rollBackConfiguration.EntityData.YListKeys = []string {}

    return &(rollBackConfiguration.EntityData)
}

// RollBackConfiguration_Input
type RollBackConfiguration_Input struct {
    EntityData types.CommonEntityData
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

func (input *RollBackConfiguration_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "roll-back-configuration"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-cfgmgr-rollback-act:roll-back-configuration/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("commit-id", types.YLeaf{"CommitId", input.CommitId})
    input.EntityData.Leafs.Append("force", types.YLeaf{"Force", input.Force})
    input.EntityData.Leafs.Append("best-effort", types.YLeaf{"BestEffort", input.BestEffort})
    input.EntityData.Leafs.Append("label", types.YLeaf{"Label", input.Label})
    input.EntityData.Leafs.Append("comment", types.YLeaf{"Comment", input.Comment})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

