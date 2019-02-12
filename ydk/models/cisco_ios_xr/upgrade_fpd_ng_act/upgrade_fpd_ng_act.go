// This module contains a collection of YANG definitions
// for Cisco IOS-XR upgrade-fpd package action data.
// 
// This module contains definitions
// for the following management objects:
//   fpd: Field programmable device (FPD) operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package upgrade_fpd_ng_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package upgrade_fpd_ng_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-upgrade-fpd-ng-act upgrade-fpd}", reflect.TypeOf(UpgradeFpd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-upgrade-fpd-ng-act:upgrade-fpd", reflect.TypeOf(UpgradeFpd{}))
}

// UpgradeFpd
// Execute FPD upgrade
type UpgradeFpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input UpgradeFpd_Input
}

func (upgradeFpd *UpgradeFpd) GetEntityData() *types.CommonEntityData {
    upgradeFpd.EntityData.YFilter = upgradeFpd.YFilter
    upgradeFpd.EntityData.YangName = "upgrade-fpd"
    upgradeFpd.EntityData.BundleName = "cisco_ios_xr"
    upgradeFpd.EntityData.ParentYangName = "Cisco-IOS-XR-upgrade-fpd-ng-act"
    upgradeFpd.EntityData.SegmentPath = "Cisco-IOS-XR-upgrade-fpd-ng-act:upgrade-fpd"
    upgradeFpd.EntityData.AbsolutePath = upgradeFpd.EntityData.SegmentPath
    upgradeFpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    upgradeFpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    upgradeFpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    upgradeFpd.EntityData.Children = types.NewOrderedMap()
    upgradeFpd.EntityData.Children.Append("input", types.YChild{"Input", &upgradeFpd.Input})
    upgradeFpd.EntityData.Leafs = types.NewOrderedMap()

    upgradeFpd.EntityData.YListKeys = []string {}

    return &(upgradeFpd.EntityData)
}

// UpgradeFpd_Input
type UpgradeFpd_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location of the FPD to be upgraded. The type is string. This attribute is
    // mandatory.
    Location interface{}

    // name of the fpd to be upgraded. The type is string. This attribute is
    // mandatory.
    Fpd interface{}

    // Force the upgrade process. The type is interface{}.
    Force interface{}
}

func (input *UpgradeFpd_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "upgrade-fpd"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-upgrade-fpd-ng-act:upgrade-fpd/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("location", types.YLeaf{"Location", input.Location})
    input.EntityData.Leafs.Append("fpd", types.YLeaf{"Fpd", input.Fpd})
    input.EntityData.Leafs.Append("force", types.YLeaf{"Force", input.Force})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

