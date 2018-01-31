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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input UpgradeFpd_Input
}

func (upgradeFpd *UpgradeFpd) GetFilter() yfilter.YFilter { return upgradeFpd.YFilter }

func (upgradeFpd *UpgradeFpd) SetFilter(yf yfilter.YFilter) { upgradeFpd.YFilter = yf }

func (upgradeFpd *UpgradeFpd) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (upgradeFpd *UpgradeFpd) GetSegmentPath() string {
    return "Cisco-IOS-XR-upgrade-fpd-ng-act:upgrade-fpd"
}

func (upgradeFpd *UpgradeFpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &upgradeFpd.Input
    }
    return nil
}

func (upgradeFpd *UpgradeFpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &upgradeFpd.Input
    return children
}

func (upgradeFpd *UpgradeFpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (upgradeFpd *UpgradeFpd) GetBundleName() string { return "cisco_ios_xr" }

func (upgradeFpd *UpgradeFpd) GetYangName() string { return "upgrade-fpd" }

func (upgradeFpd *UpgradeFpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (upgradeFpd *UpgradeFpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (upgradeFpd *UpgradeFpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (upgradeFpd *UpgradeFpd) SetParent(parent types.Entity) { upgradeFpd.parent = parent }

func (upgradeFpd *UpgradeFpd) GetParent() types.Entity { return upgradeFpd.parent }

func (upgradeFpd *UpgradeFpd) GetParentYangName() string { return "Cisco-IOS-XR-upgrade-fpd-ng-act" }

// UpgradeFpd_Input
type UpgradeFpd_Input struct {
    parent types.Entity
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

func (input *UpgradeFpd_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *UpgradeFpd_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *UpgradeFpd_Input) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    if yname == "fpd" { return "Fpd" }
    if yname == "force" { return "Force" }
    return ""
}

func (input *UpgradeFpd_Input) GetSegmentPath() string {
    return "input"
}

func (input *UpgradeFpd_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *UpgradeFpd_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *UpgradeFpd_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = input.Location
    leafs["fpd"] = input.Fpd
    leafs["force"] = input.Force
    return leafs
}

func (input *UpgradeFpd_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *UpgradeFpd_Input) GetYangName() string { return "input" }

func (input *UpgradeFpd_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *UpgradeFpd_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *UpgradeFpd_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *UpgradeFpd_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *UpgradeFpd_Input) GetParent() types.Entity { return input.parent }

func (input *UpgradeFpd_Input) GetParentYangName() string { return "upgrade-fpd" }

