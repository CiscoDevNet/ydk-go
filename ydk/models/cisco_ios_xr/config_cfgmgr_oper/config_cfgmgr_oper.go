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
    parent types.Entity
    YFilter yfilter.YFilter

    // Global operational data.
    Global Config_Global
}

func (config *Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *Config) GetGoName(yname string) string {
    if yname == "global" { return "Global" }
    return ""
}

func (config *Config) GetSegmentPath() string {
    return "Cisco-IOS-XR-config-cfgmgr-oper:config"
}

func (config *Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global" {
        return &config.Global
    }
    return nil
}

func (config *Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global"] = &config.Global
    return children
}

func (config *Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (config *Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *Config) GetYangName() string { return "config" }

func (config *Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *Config) GetParent() types.Entity { return config.parent }

func (config *Config) GetParentYangName() string { return "Cisco-IOS-XR-config-cfgmgr-oper" }

// Config_Global
// Global operational data
type Config_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (global *Config_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Config_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Config_Global) GetGoName(yname string) string {
    return ""
}

func (global *Config_Global) GetSegmentPath() string {
    return "global"
}

func (global *Config_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (global *Config_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (global *Config_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (global *Config_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *Config_Global) GetYangName() string { return "global" }

func (global *Config_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *Config_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *Config_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *Config_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Config_Global) GetParent() types.Entity { return global.parent }

func (global *Config_Global) GetParentYangName() string { return "config" }

