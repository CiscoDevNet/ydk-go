// This module contains a collection of YANG definitions
// for Cisco IOS-XR spirit-fpd-infra package configuration.
// 
// This module contains definitions
// for the following management objects:
//   fpd: Configuration for fpd auto-upgrade enable/disable 
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package spirit_fpd_infra_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package spirit_fpd_infra_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-fpd-infra-cfg fpd}", reflect.TypeOf(Fpd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-fpd-infra-cfg:fpd", reflect.TypeOf(Fpd{}))
}

// AutoUpgrade represents Auto upgrade
type AutoUpgrade string

const (
    // fpd auto-upgrade disable
    AutoUpgrade_disable AutoUpgrade = "disable"

    // fpd auto-upgrade enable
    AutoUpgrade_enable AutoUpgrade = "enable"
)

// Fpd
// Configuration for fpd auto-upgrade enable/disable
type Fpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Variable for fpd auto-upgrade enable/disable. The type is AutoUpgrade.
    AutoUpgrade interface{}
}

func (fpd *Fpd) GetEntityData() *types.CommonEntityData {
    fpd.EntityData.YFilter = fpd.YFilter
    fpd.EntityData.YangName = "fpd"
    fpd.EntityData.BundleName = "cisco_ios_xr"
    fpd.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-fpd-infra-cfg"
    fpd.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-fpd-infra-cfg:fpd"
    fpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpd.EntityData.Children = make(map[string]types.YChild)
    fpd.EntityData.Leafs = make(map[string]types.YLeaf)
    fpd.EntityData.Leafs["auto-upgrade"] = types.YLeaf{"AutoUpgrade", fpd.AutoUpgrade}
    return &(fpd.EntityData)
}

