// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs5500-coherent-portmode package operational data.
// 
// This module contains definitions
// for the following management objects:
//   controller-port-mode: Coherent PortMode  operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs5500_coherent_portmode_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs5500_coherent_portmode_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs5500-coherent-portmode-oper controller-port-mode}", reflect.TypeOf(ControllerPortMode{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs5500-coherent-portmode-oper:controller-port-mode", reflect.TypeOf(ControllerPortMode{}))
}

// ControllerPortMode
// Coherent PortMode  operational data
type ControllerPortMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of optics controller. The type is slice of
    // ControllerPortMode_OpticsName.
    OpticsName []ControllerPortMode_OpticsName
}

func (controllerPortMode *ControllerPortMode) GetEntityData() *types.CommonEntityData {
    controllerPortMode.EntityData.YFilter = controllerPortMode.YFilter
    controllerPortMode.EntityData.YangName = "controller-port-mode"
    controllerPortMode.EntityData.BundleName = "cisco_ios_xr"
    controllerPortMode.EntityData.ParentYangName = "Cisco-IOS-XR-ncs5500-coherent-portmode-oper"
    controllerPortMode.EntityData.SegmentPath = "Cisco-IOS-XR-ncs5500-coherent-portmode-oper:controller-port-mode"
    controllerPortMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllerPortMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllerPortMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllerPortMode.EntityData.Children = make(map[string]types.YChild)
    controllerPortMode.EntityData.Children["optics-name"] = types.YChild{"OpticsName", nil}
    for i := range controllerPortMode.OpticsName {
        controllerPortMode.EntityData.Children[types.GetSegmentPath(&controllerPortMode.OpticsName[i])] = types.YChild{"OpticsName", &controllerPortMode.OpticsName[i]}
    }
    controllerPortMode.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(controllerPortMode.EntityData)
}

// ControllerPortMode_OpticsName
// Name of optics controller
type ControllerPortMode_OpticsName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // PortMode  operational data.
    PortModeInfo ControllerPortMode_OpticsName_PortModeInfo
}

func (opticsName *ControllerPortMode_OpticsName) GetEntityData() *types.CommonEntityData {
    opticsName.EntityData.YFilter = opticsName.YFilter
    opticsName.EntityData.YangName = "optics-name"
    opticsName.EntityData.BundleName = "cisco_ios_xr"
    opticsName.EntityData.ParentYangName = "controller-port-mode"
    opticsName.EntityData.SegmentPath = "optics-name" + "[interface-name='" + fmt.Sprintf("%v", opticsName.InterfaceName) + "']"
    opticsName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticsName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticsName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticsName.EntityData.Children = make(map[string]types.YChild)
    opticsName.EntityData.Children["port-mode-info"] = types.YChild{"PortModeInfo", &opticsName.PortModeInfo}
    opticsName.EntityData.Leafs = make(map[string]types.YLeaf)
    opticsName.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", opticsName.InterfaceName}
    return &(opticsName.EntityData)
}

// ControllerPortMode_OpticsName_PortModeInfo
// PortMode  operational data
type ControllerPortMode_OpticsName_PortModeInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // intf name. The type is string with length: 0..128.
    IntfName interface{}

    // speed. The type is string with length: 0..128.
    Speed interface{}

    // fec. The type is string with length: 0..128.
    Fec interface{}

    // diff. The type is string with length: 0..128.
    Diff interface{}

    // modulation. The type is string with length: 0..128.
    Modulation interface{}
}

func (portModeInfo *ControllerPortMode_OpticsName_PortModeInfo) GetEntityData() *types.CommonEntityData {
    portModeInfo.EntityData.YFilter = portModeInfo.YFilter
    portModeInfo.EntityData.YangName = "port-mode-info"
    portModeInfo.EntityData.BundleName = "cisco_ios_xr"
    portModeInfo.EntityData.ParentYangName = "optics-name"
    portModeInfo.EntityData.SegmentPath = "port-mode-info"
    portModeInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portModeInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portModeInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portModeInfo.EntityData.Children = make(map[string]types.YChild)
    portModeInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    portModeInfo.EntityData.Leafs["intf-name"] = types.YLeaf{"IntfName", portModeInfo.IntfName}
    portModeInfo.EntityData.Leafs["speed"] = types.YLeaf{"Speed", portModeInfo.Speed}
    portModeInfo.EntityData.Leafs["fec"] = types.YLeaf{"Fec", portModeInfo.Fec}
    portModeInfo.EntityData.Leafs["diff"] = types.YLeaf{"Diff", portModeInfo.Diff}
    portModeInfo.EntityData.Leafs["modulation"] = types.YLeaf{"Modulation", portModeInfo.Modulation}
    return &(portModeInfo.EntityData)
}

