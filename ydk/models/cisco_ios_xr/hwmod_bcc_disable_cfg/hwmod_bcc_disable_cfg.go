// This module contains a collection of YANG definitions
// for Cisco IOS-XR hwmod-bcc-disable package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hardware-module-bcc-disable: HW module BCC Disable config
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package hwmod_bcc_disable_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package hwmod_bcc_disable_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-hwmod-bcc-disable-cfg hardware-module-bcc-disable}", reflect.TypeOf(HardwareModuleBccDisable{}))
    ydk.RegisterEntity("Cisco-IOS-XR-hwmod-bcc-disable-cfg:hardware-module-bcc-disable", reflect.TypeOf(HardwareModuleBccDisable{}))
}

// HardwareModuleBccDisable
// HW module BCC Disable config
type HardwareModuleBccDisable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bundle configuration.
    Bcc HardwareModuleBccDisable_Bcc
}

func (hardwareModuleBccDisable *HardwareModuleBccDisable) GetEntityData() *types.CommonEntityData {
    hardwareModuleBccDisable.EntityData.YFilter = hardwareModuleBccDisable.YFilter
    hardwareModuleBccDisable.EntityData.YangName = "hardware-module-bcc-disable"
    hardwareModuleBccDisable.EntityData.BundleName = "cisco_ios_xr"
    hardwareModuleBccDisable.EntityData.ParentYangName = "Cisco-IOS-XR-hwmod-bcc-disable-cfg"
    hardwareModuleBccDisable.EntityData.SegmentPath = "Cisco-IOS-XR-hwmod-bcc-disable-cfg:hardware-module-bcc-disable"
    hardwareModuleBccDisable.EntityData.AbsolutePath = hardwareModuleBccDisable.EntityData.SegmentPath
    hardwareModuleBccDisable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareModuleBccDisable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareModuleBccDisable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareModuleBccDisable.EntityData.Children = types.NewOrderedMap()
    hardwareModuleBccDisable.EntityData.Children.Append("bcc", types.YChild{"Bcc", &hardwareModuleBccDisable.Bcc})
    hardwareModuleBccDisable.EntityData.Leafs = types.NewOrderedMap()

    hardwareModuleBccDisable.EntityData.YListKeys = []string {}

    return &(hardwareModuleBccDisable.EntityData)
}

// HardwareModuleBccDisable_Bcc
// bundle configuration
type HardwareModuleBccDisable_Bcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node.
    Node HardwareModuleBccDisable_Bcc_Node
}

func (bcc *HardwareModuleBccDisable_Bcc) GetEntityData() *types.CommonEntityData {
    bcc.EntityData.YFilter = bcc.YFilter
    bcc.EntityData.YangName = "bcc"
    bcc.EntityData.BundleName = "cisco_ios_xr"
    bcc.EntityData.ParentYangName = "hardware-module-bcc-disable"
    bcc.EntityData.SegmentPath = "bcc"
    bcc.EntityData.AbsolutePath = "Cisco-IOS-XR-hwmod-bcc-disable-cfg:hardware-module-bcc-disable/" + bcc.EntityData.SegmentPath
    bcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bcc.EntityData.Children = types.NewOrderedMap()
    bcc.EntityData.Children.Append("node", types.YChild{"Node", &bcc.Node})
    bcc.EntityData.Leafs = types.NewOrderedMap()

    bcc.EntityData.YListKeys = []string {}

    return &(bcc.EntityData)
}

// HardwareModuleBccDisable_Bcc_Node
// Node
type HardwareModuleBccDisable_Bcc_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // all node configuration.
    All HardwareModuleBccDisable_Bcc_Node_All
}

func (node *HardwareModuleBccDisable_Bcc_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "bcc"
    node.EntityData.SegmentPath = "node"
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-hwmod-bcc-disable-cfg:hardware-module-bcc-disable/bcc/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("all", types.YChild{"All", &node.All})
    node.EntityData.Leafs = types.NewOrderedMap()

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// HardwareModuleBccDisable_Bcc_Node_All
// all node configuration
type HardwareModuleBccDisable_Bcc_Node_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bcc disable config. The type is interface{}.
    Disable interface{}
}

func (all *HardwareModuleBccDisable_Bcc_Node_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "node"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-hwmod-bcc-disable-cfg:hardware-module-bcc-disable/bcc/node/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", all.Disable})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

