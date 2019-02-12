// This module contains a collection of YANG definitions
// for Cisco IOS-XR config-cfgmgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   cfgmgr: Cfgmgr configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enabled or Disabled. The type is bool. The default value is true.
    ModeExclusive interface{}
}

func (cfgmgr *Cfgmgr) GetEntityData() *types.CommonEntityData {
    cfgmgr.EntityData.YFilter = cfgmgr.YFilter
    cfgmgr.EntityData.YangName = "cfgmgr"
    cfgmgr.EntityData.BundleName = "cisco_ios_xr"
    cfgmgr.EntityData.ParentYangName = "Cisco-IOS-XR-config-cfgmgr-cfg"
    cfgmgr.EntityData.SegmentPath = "Cisco-IOS-XR-config-cfgmgr-cfg:cfgmgr"
    cfgmgr.EntityData.AbsolutePath = cfgmgr.EntityData.SegmentPath
    cfgmgr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cfgmgr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cfgmgr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cfgmgr.EntityData.Children = types.NewOrderedMap()
    cfgmgr.EntityData.Leafs = types.NewOrderedMap()
    cfgmgr.EntityData.Leafs.Append("mode-exclusive", types.YLeaf{"ModeExclusive", cfgmgr.ModeExclusive})

    cfgmgr.EntityData.YListKeys = []string {}

    return &(cfgmgr.EntityData)
}

