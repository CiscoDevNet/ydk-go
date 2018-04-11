// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the top level container for
// all 'external-usb' commands for Sysadmin.
// 
// Copyright(c) 2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_external_usb

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_external_usb"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-external-usb external-usb}", reflect.TypeOf(ExternalUsb{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-external-usb:external-usb", reflect.TypeOf(ExternalUsb{}))
}

// ExternalUsb
type ExternalUsb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Config ExternalUsb_Config
}

func (externalUsb *ExternalUsb) GetEntityData() *types.CommonEntityData {
    externalUsb.EntityData.YFilter = externalUsb.YFilter
    externalUsb.EntityData.YangName = "external-usb"
    externalUsb.EntityData.BundleName = "cisco_ios_xr"
    externalUsb.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-external-usb"
    externalUsb.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-external-usb:external-usb"
    externalUsb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    externalUsb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    externalUsb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    externalUsb.EntityData.Children = make(map[string]types.YChild)
    externalUsb.EntityData.Children["config"] = types.YChild{"Config", &externalUsb.Config}
    externalUsb.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(externalUsb.EntityData)
}

// ExternalUsb_Config
type ExternalUsb_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Disable interface{}
}

func (config *ExternalUsb_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "external-usb"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = make(map[string]types.YChild)
    config.EntityData.Leafs = make(map[string]types.YLeaf)
    config.EntityData.Leafs["disable"] = types.YLeaf{"Disable", config.Disable}
    return &(config.EntityData)
}

