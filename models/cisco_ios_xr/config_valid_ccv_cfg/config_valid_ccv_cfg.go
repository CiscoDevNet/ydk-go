// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-valid-ccv package configuration.
// 
// This module contains definitions
// for the following management objects:
//   configurationvalidation: Configuration for the Configuration
//     Validation feature
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package config_valid_ccv_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package config_valid_ccv_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-config-valid-ccv-cfg configurationvalidation}", reflect.TypeOf(Configurationvalidation{}))
    ydk.RegisterEntity("Cisco-IOS-XR-config-valid-ccv-cfg:configurationvalidation", reflect.TypeOf(Configurationvalidation{}))
}

// Failure represents Failure
type Failure string

const (
    // Unsupported failure type
    Failure_unsupported Failure = "unsupported"
)

// FailureAction represents Failure action
type FailureAction string

const (
    // Report this failure type
    FailureAction_report FailureAction = "report"
)

// Configurationvalidation
// Configuration for the Configuration Validation
// feature
type Configurationvalidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable configuration validation. The type is interface{}.
    Enable interface{}

    // Table for failure type actions.
    FailureTypeActions Configurationvalidation_FailureTypeActions
}

func (configurationvalidation *Configurationvalidation) GetEntityData() *types.CommonEntityData {
    configurationvalidation.EntityData.YFilter = configurationvalidation.YFilter
    configurationvalidation.EntityData.YangName = "configurationvalidation"
    configurationvalidation.EntityData.BundleName = "cisco_ios_xr"
    configurationvalidation.EntityData.ParentYangName = "Cisco-IOS-XR-config-valid-ccv-cfg"
    configurationvalidation.EntityData.SegmentPath = "Cisco-IOS-XR-config-valid-ccv-cfg:configurationvalidation"
    configurationvalidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configurationvalidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configurationvalidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configurationvalidation.EntityData.Children = types.NewOrderedMap()
    configurationvalidation.EntityData.Children.Append("failure-type-actions", types.YChild{"FailureTypeActions", &configurationvalidation.FailureTypeActions})
    configurationvalidation.EntityData.Leafs = types.NewOrderedMap()
    configurationvalidation.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", configurationvalidation.Enable})

    configurationvalidation.EntityData.YListKeys = []string {}

    return &(configurationvalidation.EntityData)
}

// Configurationvalidation_FailureTypeActions
// Table for failure type actions
type Configurationvalidation_FailureTypeActions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Failure type action configuration. The type is slice of
    // Configurationvalidation_FailureTypeActions_FailureTypeAction.
    FailureTypeAction []*Configurationvalidation_FailureTypeActions_FailureTypeAction
}

func (failureTypeActions *Configurationvalidation_FailureTypeActions) GetEntityData() *types.CommonEntityData {
    failureTypeActions.EntityData.YFilter = failureTypeActions.YFilter
    failureTypeActions.EntityData.YangName = "failure-type-actions"
    failureTypeActions.EntityData.BundleName = "cisco_ios_xr"
    failureTypeActions.EntityData.ParentYangName = "configurationvalidation"
    failureTypeActions.EntityData.SegmentPath = "failure-type-actions"
    failureTypeActions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    failureTypeActions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    failureTypeActions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    failureTypeActions.EntityData.Children = types.NewOrderedMap()
    failureTypeActions.EntityData.Children.Append("failure-type-action", types.YChild{"FailureTypeAction", nil})
    for i := range failureTypeActions.FailureTypeAction {
        failureTypeActions.EntityData.Children.Append(types.GetSegmentPath(failureTypeActions.FailureTypeAction[i]), types.YChild{"FailureTypeAction", failureTypeActions.FailureTypeAction[i]})
    }
    failureTypeActions.EntityData.Leafs = types.NewOrderedMap()

    failureTypeActions.EntityData.YListKeys = []string {}

    return &(failureTypeActions.EntityData)
}

// Configurationvalidation_FailureTypeActions_FailureTypeAction
// Failure type action configuration
type Configurationvalidation_FailureTypeActions_FailureTypeAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Failure type. The type is Failure.
    Failure interface{}

    // Action configuration for this failure type. The type is FailureAction.
    Action interface{}
}

func (failureTypeAction *Configurationvalidation_FailureTypeActions_FailureTypeAction) GetEntityData() *types.CommonEntityData {
    failureTypeAction.EntityData.YFilter = failureTypeAction.YFilter
    failureTypeAction.EntityData.YangName = "failure-type-action"
    failureTypeAction.EntityData.BundleName = "cisco_ios_xr"
    failureTypeAction.EntityData.ParentYangName = "failure-type-actions"
    failureTypeAction.EntityData.SegmentPath = "failure-type-action" + types.AddKeyToken(failureTypeAction.Failure, "failure")
    failureTypeAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    failureTypeAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    failureTypeAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    failureTypeAction.EntityData.Children = types.NewOrderedMap()
    failureTypeAction.EntityData.Leafs = types.NewOrderedMap()
    failureTypeAction.EntityData.Leafs.Append("failure", types.YLeaf{"Failure", failureTypeAction.Failure})
    failureTypeAction.EntityData.Leafs.Append("action", types.YLeaf{"Action", failureTypeAction.Action})

    failureTypeAction.EntityData.YListKeys = []string {"Failure"}

    return &(failureTypeAction.EntityData)
}

