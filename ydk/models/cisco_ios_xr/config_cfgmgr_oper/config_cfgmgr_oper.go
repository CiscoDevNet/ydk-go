// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-cfgmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   config: Configuration-related operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package config_cfgmgr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package config_cfgmgr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-config-cfgmgr-oper config}", reflect.TypeOf(Config{}))
    ydk.RegisterEntity("Cisco-IOS-XR-config-cfgmgr-oper:config", reflect.TypeOf(Config{}))
}

// Config
// Configuration-related operational data
type Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global operational data.
    Global Config_Global
}

func (config *Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "Cisco-IOS-XR-config-cfgmgr-oper"
    config.EntityData.SegmentPath = "Cisco-IOS-XR-config-cfgmgr-oper:config"
    config.EntityData.AbsolutePath = config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("global", types.YChild{"Global", &config.Global})
    config.EntityData.Leafs = types.NewOrderedMap()

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Config_Global
// Global operational data
type Config_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration validation operational data.
    Validation Config_Global_Validation
}

func (global *Config_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "config"
    global.EntityData.SegmentPath = "global"
    global.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-oper:config/" + global.EntityData.SegmentPath
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("Cisco-IOS-XR-config-valid-ccv-oper:validation", types.YChild{"Validation", &global.Validation})
    global.EntityData.Leafs = types.NewOrderedMap()

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// Config_Global_Validation
// Configuration validation operational data
type Config_Global_Validation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unsupported config warnings present in running configuration.
    UnsupportedConfigs Config_Global_Validation_UnsupportedConfigs

    // Validation failures present in running configuration.
    PersistentFailures Config_Global_Validation_PersistentFailures
}

func (validation *Config_Global_Validation) GetEntityData() *types.CommonEntityData {
    validation.EntityData.YFilter = validation.YFilter
    validation.EntityData.YangName = "validation"
    validation.EntityData.BundleName = "cisco_ios_xr"
    validation.EntityData.ParentYangName = "global"
    validation.EntityData.SegmentPath = "Cisco-IOS-XR-config-valid-ccv-oper:validation"
    validation.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-oper:config/global/" + validation.EntityData.SegmentPath
    validation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    validation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    validation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    validation.EntityData.Children = types.NewOrderedMap()
    validation.EntityData.Children.Append("unsupported-configs", types.YChild{"UnsupportedConfigs", &validation.UnsupportedConfigs})
    validation.EntityData.Children.Append("persistent-failures", types.YChild{"PersistentFailures", &validation.PersistentFailures})
    validation.EntityData.Leafs = types.NewOrderedMap()

    validation.EntityData.YListKeys = []string {}

    return &(validation.EntityData)
}

// Config_Global_Validation_UnsupportedConfigs
// Unsupported config warnings present in running
// configuration
type Config_Global_Validation_UnsupportedConfigs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about an unsupported warning. The type is slice of
    // Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig.
    UnsupportedConfig []*Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig
}

func (unsupportedConfigs *Config_Global_Validation_UnsupportedConfigs) GetEntityData() *types.CommonEntityData {
    unsupportedConfigs.EntityData.YFilter = unsupportedConfigs.YFilter
    unsupportedConfigs.EntityData.YangName = "unsupported-configs"
    unsupportedConfigs.EntityData.BundleName = "cisco_ios_xr"
    unsupportedConfigs.EntityData.ParentYangName = "validation"
    unsupportedConfigs.EntityData.SegmentPath = "unsupported-configs"
    unsupportedConfigs.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-oper:config/global/Cisco-IOS-XR-config-valid-ccv-oper:validation/" + unsupportedConfigs.EntityData.SegmentPath
    unsupportedConfigs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unsupportedConfigs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unsupportedConfigs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unsupportedConfigs.EntityData.Children = types.NewOrderedMap()
    unsupportedConfigs.EntityData.Children.Append("unsupported-config", types.YChild{"UnsupportedConfig", nil})
    for i := range unsupportedConfigs.UnsupportedConfig {
        unsupportedConfigs.EntityData.Children.Append(types.GetSegmentPath(unsupportedConfigs.UnsupportedConfig[i]), types.YChild{"UnsupportedConfig", unsupportedConfigs.UnsupportedConfig[i]})
    }
    unsupportedConfigs.EntityData.Leafs = types.NewOrderedMap()

    unsupportedConfigs.EntityData.YListKeys = []string {}

    return &(unsupportedConfigs.EntityData)
}

// Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig
// Information about an unsupported warning
type Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. XPath of the unsupported configuration. The type
    // is string.
    Xpath interface{}

    // The configuration group that this item is inherited from, if any. The type
    // is string.
    GroupName interface{}

    // Validation failures for this configuration item. The type is slice of
    // Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig_Failure.
    Failure []*Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig_Failure
}

func (unsupportedConfig *Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig) GetEntityData() *types.CommonEntityData {
    unsupportedConfig.EntityData.YFilter = unsupportedConfig.YFilter
    unsupportedConfig.EntityData.YangName = "unsupported-config"
    unsupportedConfig.EntityData.BundleName = "cisco_ios_xr"
    unsupportedConfig.EntityData.ParentYangName = "unsupported-configs"
    unsupportedConfig.EntityData.SegmentPath = "unsupported-config" + types.AddKeyToken(unsupportedConfig.Xpath, "xpath")
    unsupportedConfig.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-oper:config/global/Cisco-IOS-XR-config-valid-ccv-oper:validation/unsupported-configs/" + unsupportedConfig.EntityData.SegmentPath
    unsupportedConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unsupportedConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unsupportedConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unsupportedConfig.EntityData.Children = types.NewOrderedMap()
    unsupportedConfig.EntityData.Children.Append("failure", types.YChild{"Failure", nil})
    for i := range unsupportedConfig.Failure {
        types.SetYListKey(unsupportedConfig.Failure[i], i)
        unsupportedConfig.EntityData.Children.Append(types.GetSegmentPath(unsupportedConfig.Failure[i]), types.YChild{"Failure", unsupportedConfig.Failure[i]})
    }
    unsupportedConfig.EntityData.Leafs = types.NewOrderedMap()
    unsupportedConfig.EntityData.Leafs.Append("xpath", types.YLeaf{"Xpath", unsupportedConfig.Xpath})
    unsupportedConfig.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", unsupportedConfig.GroupName})

    unsupportedConfig.EntityData.YListKeys = []string {"Xpath"}

    return &(unsupportedConfig.EntityData)
}

// Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig_Failure
// Validation failures for this configuration item
type Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig_Failure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // A unique string that identifies the error; equivalent to error-app-tag in
    // RFC 6241. The type is string.
    ErrorAppTag interface{}

    // The error message; equivalent to error-message in RFC 6241. The type is
    // string.
    ErrorMessage interface{}

    // The severity of the error; equivalent to error-severity in RFC 6241. The
    // type is string.
    ErrorSeverity interface{}
}

func (failure *Config_Global_Validation_UnsupportedConfigs_UnsupportedConfig_Failure) GetEntityData() *types.CommonEntityData {
    failure.EntityData.YFilter = failure.YFilter
    failure.EntityData.YangName = "failure"
    failure.EntityData.BundleName = "cisco_ios_xr"
    failure.EntityData.ParentYangName = "unsupported-config"
    failure.EntityData.SegmentPath = "failure" + types.AddNoKeyToken(failure)
    failure.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-oper:config/global/Cisco-IOS-XR-config-valid-ccv-oper:validation/unsupported-configs/unsupported-config/" + failure.EntityData.SegmentPath
    failure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    failure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    failure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    failure.EntityData.Children = types.NewOrderedMap()
    failure.EntityData.Leafs = types.NewOrderedMap()
    failure.EntityData.Leafs.Append("error-app-tag", types.YLeaf{"ErrorAppTag", failure.ErrorAppTag})
    failure.EntityData.Leafs.Append("error-message", types.YLeaf{"ErrorMessage", failure.ErrorMessage})
    failure.EntityData.Leafs.Append("error-severity", types.YLeaf{"ErrorSeverity", failure.ErrorSeverity})

    failure.EntityData.YListKeys = []string {}

    return &(failure.EntityData)
}

// Config_Global_Validation_PersistentFailures
// Validation failures present in running
// configuration
type Config_Global_Validation_PersistentFailures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a validation failure. The type is slice of
    // Config_Global_Validation_PersistentFailures_PersistentFailure.
    PersistentFailure []*Config_Global_Validation_PersistentFailures_PersistentFailure
}

func (persistentFailures *Config_Global_Validation_PersistentFailures) GetEntityData() *types.CommonEntityData {
    persistentFailures.EntityData.YFilter = persistentFailures.YFilter
    persistentFailures.EntityData.YangName = "persistent-failures"
    persistentFailures.EntityData.BundleName = "cisco_ios_xr"
    persistentFailures.EntityData.ParentYangName = "validation"
    persistentFailures.EntityData.SegmentPath = "persistent-failures"
    persistentFailures.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-oper:config/global/Cisco-IOS-XR-config-valid-ccv-oper:validation/" + persistentFailures.EntityData.SegmentPath
    persistentFailures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    persistentFailures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    persistentFailures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    persistentFailures.EntityData.Children = types.NewOrderedMap()
    persistentFailures.EntityData.Children.Append("persistent-failure", types.YChild{"PersistentFailure", nil})
    for i := range persistentFailures.PersistentFailure {
        persistentFailures.EntityData.Children.Append(types.GetSegmentPath(persistentFailures.PersistentFailure[i]), types.YChild{"PersistentFailure", persistentFailures.PersistentFailure[i]})
    }
    persistentFailures.EntityData.Leafs = types.NewOrderedMap()

    persistentFailures.EntityData.YListKeys = []string {}

    return &(persistentFailures.EntityData)
}

// Config_Global_Validation_PersistentFailures_PersistentFailure
// Information about a validation failure
type Config_Global_Validation_PersistentFailures_PersistentFailure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. XPath of the failed configuration. The type is
    // string.
    Xpath interface{}

    // The configuration group that this item is inherited from, if any. The type
    // is string.
    GroupName interface{}

    // Validation failures for this configuration item. The type is slice of
    // Config_Global_Validation_PersistentFailures_PersistentFailure_Failure.
    Failure []*Config_Global_Validation_PersistentFailures_PersistentFailure_Failure
}

func (persistentFailure *Config_Global_Validation_PersistentFailures_PersistentFailure) GetEntityData() *types.CommonEntityData {
    persistentFailure.EntityData.YFilter = persistentFailure.YFilter
    persistentFailure.EntityData.YangName = "persistent-failure"
    persistentFailure.EntityData.BundleName = "cisco_ios_xr"
    persistentFailure.EntityData.ParentYangName = "persistent-failures"
    persistentFailure.EntityData.SegmentPath = "persistent-failure" + types.AddKeyToken(persistentFailure.Xpath, "xpath")
    persistentFailure.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-oper:config/global/Cisco-IOS-XR-config-valid-ccv-oper:validation/persistent-failures/" + persistentFailure.EntityData.SegmentPath
    persistentFailure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    persistentFailure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    persistentFailure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    persistentFailure.EntityData.Children = types.NewOrderedMap()
    persistentFailure.EntityData.Children.Append("failure", types.YChild{"Failure", nil})
    for i := range persistentFailure.Failure {
        types.SetYListKey(persistentFailure.Failure[i], i)
        persistentFailure.EntityData.Children.Append(types.GetSegmentPath(persistentFailure.Failure[i]), types.YChild{"Failure", persistentFailure.Failure[i]})
    }
    persistentFailure.EntityData.Leafs = types.NewOrderedMap()
    persistentFailure.EntityData.Leafs.Append("xpath", types.YLeaf{"Xpath", persistentFailure.Xpath})
    persistentFailure.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", persistentFailure.GroupName})

    persistentFailure.EntityData.YListKeys = []string {"Xpath"}

    return &(persistentFailure.EntityData)
}

// Config_Global_Validation_PersistentFailures_PersistentFailure_Failure
// Validation failures for this configuration item
type Config_Global_Validation_PersistentFailures_PersistentFailure_Failure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // A unique string that identifies the error; equivalent to error-app-tag in
    // RFC 6241. The type is string.
    ErrorAppTag interface{}

    // The error message; equivalent to error-message in RFC 6241. The type is
    // string.
    ErrorMessage interface{}

    // The severity of the error; equivalent to error-severity in RFC 6241. The
    // type is string.
    ErrorSeverity interface{}
}

func (failure *Config_Global_Validation_PersistentFailures_PersistentFailure_Failure) GetEntityData() *types.CommonEntityData {
    failure.EntityData.YFilter = failure.YFilter
    failure.EntityData.YangName = "failure"
    failure.EntityData.BundleName = "cisco_ios_xr"
    failure.EntityData.ParentYangName = "persistent-failure"
    failure.EntityData.SegmentPath = "failure" + types.AddNoKeyToken(failure)
    failure.EntityData.AbsolutePath = "Cisco-IOS-XR-config-cfgmgr-oper:config/global/Cisco-IOS-XR-config-valid-ccv-oper:validation/persistent-failures/persistent-failure/" + failure.EntityData.SegmentPath
    failure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    failure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    failure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    failure.EntityData.Children = types.NewOrderedMap()
    failure.EntityData.Leafs = types.NewOrderedMap()
    failure.EntityData.Leafs.Append("error-app-tag", types.YLeaf{"ErrorAppTag", failure.ErrorAppTag})
    failure.EntityData.Leafs.Append("error-message", types.YLeaf{"ErrorMessage", failure.ErrorMessage})
    failure.EntityData.Leafs.Append("error-severity", types.YLeaf{"ErrorSeverity", failure.ErrorSeverity})

    failure.EntityData.YListKeys = []string {}

    return &(failure.EntityData)
}

