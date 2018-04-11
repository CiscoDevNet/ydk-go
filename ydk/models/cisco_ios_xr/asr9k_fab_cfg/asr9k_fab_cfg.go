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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mode Type.
    Mode FabVqiConfig_Mode
}

func (fabVqiConfig *FabVqiConfig) GetEntityData() *types.CommonEntityData {
    fabVqiConfig.EntityData.YFilter = fabVqiConfig.YFilter
    fabVqiConfig.EntityData.YangName = "fab-vqi-config"
    fabVqiConfig.EntityData.BundleName = "cisco_ios_xr"
    fabVqiConfig.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-fab-cfg"
    fabVqiConfig.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-fab-cfg:fab-vqi-config"
    fabVqiConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabVqiConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabVqiConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabVqiConfig.EntityData.Children = make(map[string]types.YChild)
    fabVqiConfig.EntityData.Children["mode"] = types.YChild{"Mode", &fabVqiConfig.Mode}
    fabVqiConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fabVqiConfig.EntityData)
}

// FabVqiConfig_Mode
// Mode Type
type FabVqiConfig_Mode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mode Type. The type is Asr9kFabMode.
    FabModeTypeXr interface{}

    // Mode Type. The type is Asr9kFabMode.
    FabModeType interface{}
}

func (mode *FabVqiConfig_Mode) GetEntityData() *types.CommonEntityData {
    mode.EntityData.YFilter = mode.YFilter
    mode.EntityData.YangName = "mode"
    mode.EntityData.BundleName = "cisco_ios_xr"
    mode.EntityData.ParentYangName = "fab-vqi-config"
    mode.EntityData.SegmentPath = "mode"
    mode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mode.EntityData.Children = make(map[string]types.YChild)
    mode.EntityData.Leafs = make(map[string]types.YLeaf)
    mode.EntityData.Leafs["fab-mode-type-xr"] = types.YLeaf{"FabModeTypeXr", mode.FabModeTypeXr}
    mode.EntityData.Leafs["fab-mode-type"] = types.YLeaf{"FabModeType", mode.FabModeType}
    return &(mode.EntityData)
}

