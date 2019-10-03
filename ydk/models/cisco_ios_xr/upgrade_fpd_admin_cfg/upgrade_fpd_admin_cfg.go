// This module contains a collection of YANG definitions
// for Cisco IOS-XR upgrade-fpd package admin-plane configuration.
// 
// This module contains definitions
// for the following management objects:
//   fpd: Configuration for fpd auto-upgrade
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package upgrade_fpd_admin_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package upgrade_fpd_admin_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-upgrade-fpd-admin-cfg fpd}", reflect.TypeOf(Fpd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-upgrade-fpd-admin-cfg:fpd", reflect.TypeOf(Fpd{}))
}

// Fpd
// Configuration for fpd auto-upgrade
type Fpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Variable for fpd auto-upgrade enable/disable. The type is bool.
    AutoUpgrade interface{}

    // Variable for fpd auto-reload enable/disable. The type is bool.
    AutoReload interface{}
}

func (fpd *Fpd) GetEntityData() *types.CommonEntityData {
    fpd.EntityData.YFilter = fpd.YFilter
    fpd.EntityData.YangName = "fpd"
    fpd.EntityData.BundleName = "cisco_ios_xr"
    fpd.EntityData.ParentYangName = "Cisco-IOS-XR-upgrade-fpd-admin-cfg"
    fpd.EntityData.SegmentPath = "Cisco-IOS-XR-upgrade-fpd-admin-cfg:fpd"
    fpd.EntityData.AbsolutePath = fpd.EntityData.SegmentPath
    fpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpd.EntityData.Children = types.NewOrderedMap()
    fpd.EntityData.Leafs = types.NewOrderedMap()
    fpd.EntityData.Leafs.Append("auto-upgrade", types.YLeaf{"AutoUpgrade", fpd.AutoUpgrade})
    fpd.EntityData.Leafs.Append("auto-reload", types.YLeaf{"AutoReload", fpd.AutoReload})

    fpd.EntityData.YListKeys = []string {}

    return &(fpd.EntityData)
}

