// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-cfgmgr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   config: Configuration-related operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
}

func (global *Config_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "config"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Leafs = types.NewOrderedMap()

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

