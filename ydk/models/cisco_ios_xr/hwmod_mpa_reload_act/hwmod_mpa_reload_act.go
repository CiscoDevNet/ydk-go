// This module contains a collection of YANG definitions
// for Cisco IOS-XR MPA reload action package configuration.
// 
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package hwmod_mpa_reload_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package hwmod_mpa_reload_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-hwmod-mpa-reload-act hw-module-subslot}", reflect.TypeOf(HwModuleSubslot{}))
    ydk.RegisterEntity("Cisco-IOS-XR-hwmod-mpa-reload-act:hw-module-subslot", reflect.TypeOf(HwModuleSubslot{}))
}

// HwModuleSubslot
// Execute subslot h/w module operations
type HwModuleSubslot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input HwModuleSubslot_Input
}

func (hwModuleSubslot *HwModuleSubslot) GetFilter() yfilter.YFilter { return hwModuleSubslot.YFilter }

func (hwModuleSubslot *HwModuleSubslot) SetFilter(yf yfilter.YFilter) { hwModuleSubslot.YFilter = yf }

func (hwModuleSubslot *HwModuleSubslot) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (hwModuleSubslot *HwModuleSubslot) GetSegmentPath() string {
    return "Cisco-IOS-XR-hwmod-mpa-reload-act:hw-module-subslot"
}

func (hwModuleSubslot *HwModuleSubslot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &hwModuleSubslot.Input
    }
    return nil
}

func (hwModuleSubslot *HwModuleSubslot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &hwModuleSubslot.Input
    return children
}

func (hwModuleSubslot *HwModuleSubslot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hwModuleSubslot *HwModuleSubslot) GetBundleName() string { return "cisco_ios_xr" }

func (hwModuleSubslot *HwModuleSubslot) GetYangName() string { return "hw-module-subslot" }

func (hwModuleSubslot *HwModuleSubslot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwModuleSubslot *HwModuleSubslot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwModuleSubslot *HwModuleSubslot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwModuleSubslot *HwModuleSubslot) SetParent(parent types.Entity) { hwModuleSubslot.parent = parent }

func (hwModuleSubslot *HwModuleSubslot) GetParent() types.Entity { return hwModuleSubslot.parent }

func (hwModuleSubslot *HwModuleSubslot) GetParentYangName() string { return "Cisco-IOS-XR-hwmod-mpa-reload-act" }

// HwModuleSubslot_Input
type HwModuleSubslot_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fully qualified location specification. The type is string.
    Subslot interface{}

    // Cycle subslot h/w reset. The type is interface{}.
    Reload interface{}
}

func (input *HwModuleSubslot_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *HwModuleSubslot_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *HwModuleSubslot_Input) GetGoName(yname string) string {
    if yname == "subslot" { return "Subslot" }
    if yname == "reload" { return "Reload" }
    return ""
}

func (input *HwModuleSubslot_Input) GetSegmentPath() string {
    return "input"
}

func (input *HwModuleSubslot_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *HwModuleSubslot_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *HwModuleSubslot_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subslot"] = input.Subslot
    leafs["reload"] = input.Reload
    return leafs
}

func (input *HwModuleSubslot_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *HwModuleSubslot_Input) GetYangName() string { return "input" }

func (input *HwModuleSubslot_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *HwModuleSubslot_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *HwModuleSubslot_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *HwModuleSubslot_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *HwModuleSubslot_Input) GetParent() types.Entity { return input.parent }

func (input *HwModuleSubslot_Input) GetParentYangName() string { return "hw-module-subslot" }

