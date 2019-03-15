// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-fia package configuration.
// 
// This module contains definitions
// for the following management objects:
//   fabric-fia-config: Configure Global Fabric Fia Settings
// 
// This YANG module augments the
//   Cisco-IOS-XR-config-mda-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_fia_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_fia_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-fia-cfg fabric-fia-config}", reflect.TypeOf(FabricFiaConfig{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-fia-cfg:fabric-fia-config", reflect.TypeOf(FabricFiaConfig{}))
}

// FabricFiaConfig
// Configure Global Fabric Fia Settings
type FabricFiaConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FIA interface rate-limiter on 7-Fabric LC.
    FiaIntfPolicer FabricFiaConfig_FiaIntfPolicer
}

func (fabricFiaConfig *FabricFiaConfig) GetEntityData() *types.CommonEntityData {
    fabricFiaConfig.EntityData.YFilter = fabricFiaConfig.YFilter
    fabricFiaConfig.EntityData.YangName = "fabric-fia-config"
    fabricFiaConfig.EntityData.BundleName = "cisco_ios_xr"
    fabricFiaConfig.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-fia-cfg"
    fabricFiaConfig.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-fia-cfg:fabric-fia-config"
    fabricFiaConfig.EntityData.AbsolutePath = fabricFiaConfig.EntityData.SegmentPath
    fabricFiaConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabricFiaConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabricFiaConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabricFiaConfig.EntityData.Children = types.NewOrderedMap()
    fabricFiaConfig.EntityData.Children.Append("fia-intf-policer", types.YChild{"FiaIntfPolicer", &fabricFiaConfig.FiaIntfPolicer})
    fabricFiaConfig.EntityData.Leafs = types.NewOrderedMap()

    fabricFiaConfig.EntityData.YListKeys = []string {}

    return &(fabricFiaConfig.EntityData)
}

// FabricFiaConfig_FiaIntfPolicer
// FIA interface rate-limiter on 7-Fabric LC
type FabricFiaConfig_FiaIntfPolicer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // disable FIA interface policer . The type is bool.
    Disable interface{}
}

func (fiaIntfPolicer *FabricFiaConfig_FiaIntfPolicer) GetEntityData() *types.CommonEntityData {
    fiaIntfPolicer.EntityData.YFilter = fiaIntfPolicer.YFilter
    fiaIntfPolicer.EntityData.YangName = "fia-intf-policer"
    fiaIntfPolicer.EntityData.BundleName = "cisco_ios_xr"
    fiaIntfPolicer.EntityData.ParentYangName = "fabric-fia-config"
    fiaIntfPolicer.EntityData.SegmentPath = "fia-intf-policer"
    fiaIntfPolicer.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-fia-cfg:fabric-fia-config/" + fiaIntfPolicer.EntityData.SegmentPath
    fiaIntfPolicer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fiaIntfPolicer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fiaIntfPolicer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fiaIntfPolicer.EntityData.Children = types.NewOrderedMap()
    fiaIntfPolicer.EntityData.Leafs = types.NewOrderedMap()
    fiaIntfPolicer.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", fiaIntfPolicer.Disable})

    fiaIntfPolicer.EntityData.YListKeys = []string {}

    return &(fiaIntfPolicer.EntityData)
}

