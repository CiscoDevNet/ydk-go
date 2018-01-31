// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-cfgmgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   cfgmgr: Cfgmgr configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package config_cfgmgr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package config_cfgmgr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-config-cfgmgr-cfg cfgmgr}", reflect.TypeOf(Cfgmgr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-config-cfgmgr-cfg:cfgmgr", reflect.TypeOf(Cfgmgr{}))
}

// Cfgmgr
// Cfgmgr configuration
type Cfgmgr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enabled or Disabled. The type is bool. The default value is true.
    ModeExclusive interface{}
}

func (cfgmgr *Cfgmgr) GetFilter() yfilter.YFilter { return cfgmgr.YFilter }

func (cfgmgr *Cfgmgr) SetFilter(yf yfilter.YFilter) { cfgmgr.YFilter = yf }

func (cfgmgr *Cfgmgr) GetGoName(yname string) string {
    if yname == "mode-exclusive" { return "ModeExclusive" }
    return ""
}

func (cfgmgr *Cfgmgr) GetSegmentPath() string {
    return "Cisco-IOS-XR-config-cfgmgr-cfg:cfgmgr"
}

func (cfgmgr *Cfgmgr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cfgmgr *Cfgmgr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cfgmgr *Cfgmgr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode-exclusive"] = cfgmgr.ModeExclusive
    return leafs
}

func (cfgmgr *Cfgmgr) GetBundleName() string { return "cisco_ios_xr" }

func (cfgmgr *Cfgmgr) GetYangName() string { return "cfgmgr" }

func (cfgmgr *Cfgmgr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cfgmgr *Cfgmgr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cfgmgr *Cfgmgr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cfgmgr *Cfgmgr) SetParent(parent types.Entity) { cfgmgr.parent = parent }

func (cfgmgr *Cfgmgr) GetParent() types.Entity { return cfgmgr.parent }

func (cfgmgr *Cfgmgr) GetParentYangName() string { return "Cisco-IOS-XR-config-cfgmgr-cfg" }

