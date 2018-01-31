// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-fab package configuration.
// 
// This module contains definitions
// for the following management objects:
//   fab-vqi-config: Configure Fabric Operation Mode
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_fab_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_fab_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-fab-cfg fab-vqi-config}", reflect.TypeOf(FabVqiConfig{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-fab-cfg:fab-vqi-config", reflect.TypeOf(FabVqiConfig{}))
}

// Asr9kFabMode represents Asr9k fab mode
type Asr9kFabMode string

const (
    // A99 High bandwidth mode
    Asr9kFabMode_a99_highbandwidth Asr9kFabMode = "a99-highbandwidth"
)

// FabVqiConfig
// Configure Fabric Operation Mode
type FabVqiConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mode Type.
    Mode FabVqiConfig_Mode
}

func (fabVqiConfig *FabVqiConfig) GetFilter() yfilter.YFilter { return fabVqiConfig.YFilter }

func (fabVqiConfig *FabVqiConfig) SetFilter(yf yfilter.YFilter) { fabVqiConfig.YFilter = yf }

func (fabVqiConfig *FabVqiConfig) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    return ""
}

func (fabVqiConfig *FabVqiConfig) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-fab-cfg:fab-vqi-config"
}

func (fabVqiConfig *FabVqiConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mode" {
        return &fabVqiConfig.Mode
    }
    return nil
}

func (fabVqiConfig *FabVqiConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mode"] = &fabVqiConfig.Mode
    return children
}

func (fabVqiConfig *FabVqiConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fabVqiConfig *FabVqiConfig) GetBundleName() string { return "cisco_ios_xr" }

func (fabVqiConfig *FabVqiConfig) GetYangName() string { return "fab-vqi-config" }

func (fabVqiConfig *FabVqiConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fabVqiConfig *FabVqiConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fabVqiConfig *FabVqiConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fabVqiConfig *FabVqiConfig) SetParent(parent types.Entity) { fabVqiConfig.parent = parent }

func (fabVqiConfig *FabVqiConfig) GetParent() types.Entity { return fabVqiConfig.parent }

func (fabVqiConfig *FabVqiConfig) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-fab-cfg" }

// FabVqiConfig_Mode
// Mode Type
type FabVqiConfig_Mode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mode Type. The type is Asr9kFabMode.
    FabModeTypeXr interface{}

    // Mode Type. The type is Asr9kFabMode.
    FabModeType interface{}
}

func (mode *FabVqiConfig_Mode) GetFilter() yfilter.YFilter { return mode.YFilter }

func (mode *FabVqiConfig_Mode) SetFilter(yf yfilter.YFilter) { mode.YFilter = yf }

func (mode *FabVqiConfig_Mode) GetGoName(yname string) string {
    if yname == "fab-mode-type-xr" { return "FabModeTypeXr" }
    if yname == "fab-mode-type" { return "FabModeType" }
    return ""
}

func (mode *FabVqiConfig_Mode) GetSegmentPath() string {
    return "mode"
}

func (mode *FabVqiConfig_Mode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mode *FabVqiConfig_Mode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mode *FabVqiConfig_Mode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fab-mode-type-xr"] = mode.FabModeTypeXr
    leafs["fab-mode-type"] = mode.FabModeType
    return leafs
}

func (mode *FabVqiConfig_Mode) GetBundleName() string { return "cisco_ios_xr" }

func (mode *FabVqiConfig_Mode) GetYangName() string { return "mode" }

func (mode *FabVqiConfig_Mode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mode *FabVqiConfig_Mode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mode *FabVqiConfig_Mode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mode *FabVqiConfig_Mode) SetParent(parent types.Entity) { mode.parent = parent }

func (mode *FabVqiConfig_Mode) GetParent() types.Entity { return mode.parent }

func (mode *FabVqiConfig_Mode) GetParentYangName() string { return "fab-vqi-config" }

