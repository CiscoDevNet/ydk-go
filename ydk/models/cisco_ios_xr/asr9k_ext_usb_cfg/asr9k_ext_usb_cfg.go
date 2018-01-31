// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-ext-usb package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-ext-usb: External USB configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_ext_usb_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_ext_usb_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-ext-usb-cfg hardware-module-ext-usb}", reflect.TypeOf(HardwareModuleExtUsb{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-ext-usb-cfg:hardware-module-ext-usb", reflect.TypeOf(HardwareModuleExtUsb{}))
}

// HardwareModuleExtUsb
// External USB configuration
type HardwareModuleExtUsb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // External USB disable. The type is interface{}.
    Disable interface{}
}

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetFilter() yfilter.YFilter { return hardwareModuleExtUsb.YFilter }

func (hardwareModuleExtUsb *HardwareModuleExtUsb) SetFilter(yf yfilter.YFilter) { hardwareModuleExtUsb.YFilter = yf }

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-ext-usb-cfg:hardware-module-ext-usb"
}

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = hardwareModuleExtUsb.Disable
    return leafs
}

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetBundleName() string { return "cisco_ios_xr" }

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetYangName() string { return "hardware-module-ext-usb" }

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hardwareModuleExtUsb *HardwareModuleExtUsb) SetParent(parent types.Entity) { hardwareModuleExtUsb.parent = parent }

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetParent() types.Entity { return hardwareModuleExtUsb.parent }

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-ext-usb-cfg" }

