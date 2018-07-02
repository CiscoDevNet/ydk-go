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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // External USB disable. The type is interface{}.
    Disable interface{}
}

func (hardwareModuleExtUsb *HardwareModuleExtUsb) GetEntityData() *types.CommonEntityData {
    hardwareModuleExtUsb.EntityData.YFilter = hardwareModuleExtUsb.YFilter
    hardwareModuleExtUsb.EntityData.YangName = "hardware-module-ext-usb"
    hardwareModuleExtUsb.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleExtUsb.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-ext-usb-cfg"
    hardwareModuleExtUsb.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-ext-usb-cfg:hardware-module-ext-usb"
    hardwareModuleExtUsb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleExtUsb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleExtUsb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleExtUsb.EntityData.Children = types.NewOrderedMap()
    hardwareModuleExtUsb.EntityData.Leafs = types.NewOrderedMap()
    hardwareModuleExtUsb.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", hardwareModuleExtUsb.Disable})

    hardwareModuleExtUsb.EntityData.YListKeys = []string {}

    return &(hardwareModuleExtUsb.EntityData)
}

