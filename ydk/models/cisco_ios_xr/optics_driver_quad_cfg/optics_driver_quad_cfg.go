// This module contains a collection of YANG definitions
// for Cisco IOS-XR optics-driver-quad package configuration.
// 
// This module contains definitions
// for the following management objects:
//   node: HW module Quad Config
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package optics_driver_quad_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package optics_driver_quad_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-optics-driver-quad-cfg node}", reflect.TypeOf(Node{}))
    ydk.RegisterEntity("Cisco-IOS-XR-optics-driver-quad-cfg:node", reflect.TypeOf(Node{}))
}

// Node
// HW module Quad Config
type Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none.
    Acts Node_Acts
}

func (node *Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "Cisco-IOS-XR-optics-driver-quad-cfg"
    node.EntityData.SegmentPath = "Cisco-IOS-XR-optics-driver-quad-cfg:node"
    node.EntityData.AbsolutePath = node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("acts", types.YChild{"Acts", &node.Acts})
    node.EntityData.Leafs = types.NewOrderedMap()

    node.EntityData.YListKeys = []string {}

    return &(node.EntityData)
}

// Node_Acts
// none
type Node_Acts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nodename. The type is slice of Node_Acts_Act.
    Act []*Node_Acts_Act
}

func (acts *Node_Acts) GetEntityData() *types.CommonEntityData {
    acts.EntityData.YFilter = acts.YFilter
    acts.EntityData.YangName = "acts"
    acts.EntityData.BundleName = "cisco_ios_xr"
    acts.EntityData.ParentYangName = "node"
    acts.EntityData.SegmentPath = "acts"
    acts.EntityData.AbsolutePath = "Cisco-IOS-XR-optics-driver-quad-cfg:node/" + acts.EntityData.SegmentPath
    acts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acts.EntityData.Children = types.NewOrderedMap()
    acts.EntityData.Children.Append("act", types.YChild{"Act", nil})
    for i := range acts.Act {
        acts.EntityData.Children.Append(types.GetSegmentPath(acts.Act[i]), types.YChild{"Act", acts.Act[i]})
    }
    acts.EntityData.Leafs = types.NewOrderedMap()

    acts.EntityData.YListKeys = []string {}

    return &(acts.EntityData)
}

// Node_Acts_Act
// Nodename
type Node_Acts_Act struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NodeName. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    NodeName interface{}

    // quad configuration.
    QuadConfigs Node_Acts_Act_QuadConfigs
}

func (act *Node_Acts_Act) GetEntityData() *types.CommonEntityData {
    act.EntityData.YFilter = act.YFilter
    act.EntityData.YangName = "act"
    act.EntityData.BundleName = "cisco_ios_xr"
    act.EntityData.ParentYangName = "acts"
    act.EntityData.SegmentPath = "act" + types.AddKeyToken(act.NodeName, "node-name")
    act.EntityData.AbsolutePath = "Cisco-IOS-XR-optics-driver-quad-cfg:node/acts/" + act.EntityData.SegmentPath
    act.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    act.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    act.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    act.EntityData.Children = types.NewOrderedMap()
    act.EntityData.Children.Append("quad-configs", types.YChild{"QuadConfigs", &act.QuadConfigs})
    act.EntityData.Leafs = types.NewOrderedMap()
    act.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", act.NodeName})

    act.EntityData.YListKeys = []string {"NodeName"}

    return &(act.EntityData)
}

// Node_Acts_Act_QuadConfigs
// quad configuration
type Node_Acts_Act_QuadConfigs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of Node_Acts_Act_QuadConfigs_QuadConfig.
    QuadConfig []*Node_Acts_Act_QuadConfigs_QuadConfig
}

func (quadConfigs *Node_Acts_Act_QuadConfigs) GetEntityData() *types.CommonEntityData {
    quadConfigs.EntityData.YFilter = quadConfigs.YFilter
    quadConfigs.EntityData.YangName = "quad-configs"
    quadConfigs.EntityData.BundleName = "cisco_ios_xr"
    quadConfigs.EntityData.ParentYangName = "act"
    quadConfigs.EntityData.SegmentPath = "quad-configs"
    quadConfigs.EntityData.AbsolutePath = "Cisco-IOS-XR-optics-driver-quad-cfg:node/acts/act/" + quadConfigs.EntityData.SegmentPath
    quadConfigs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    quadConfigs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    quadConfigs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    quadConfigs.EntityData.Children = types.NewOrderedMap()
    quadConfigs.EntityData.Children.Append("quad-config", types.YChild{"QuadConfig", nil})
    for i := range quadConfigs.QuadConfig {
        quadConfigs.EntityData.Children.Append(types.GetSegmentPath(quadConfigs.QuadConfig[i]), types.YChild{"QuadConfig", quadConfigs.QuadConfig[i]})
    }
    quadConfigs.EntityData.Leafs = types.NewOrderedMap()

    quadConfigs.EntityData.YListKeys = []string {}

    return &(quadConfigs.EntityData)
}

// Node_Acts_Act_QuadConfigs_QuadConfig
// none
type Node_Acts_Act_QuadConfigs_QuadConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. none. The type is interface{} with range:
    // 0..4294967295.
    Id1 interface{}

    // select mode 10g or 25g for a quad(group of 4 ports).
    Mode Node_Acts_Act_QuadConfigs_QuadConfig_Mode
}

func (quadConfig *Node_Acts_Act_QuadConfigs_QuadConfig) GetEntityData() *types.CommonEntityData {
    quadConfig.EntityData.YFilter = quadConfig.YFilter
    quadConfig.EntityData.YangName = "quad-config"
    quadConfig.EntityData.BundleName = "cisco_ios_xr"
    quadConfig.EntityData.ParentYangName = "quad-configs"
    quadConfig.EntityData.SegmentPath = "quad-config" + types.AddKeyToken(quadConfig.Id1, "id1")
    quadConfig.EntityData.AbsolutePath = "Cisco-IOS-XR-optics-driver-quad-cfg:node/acts/act/quad-configs/" + quadConfig.EntityData.SegmentPath
    quadConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    quadConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    quadConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    quadConfig.EntityData.Children = types.NewOrderedMap()
    quadConfig.EntityData.Children.Append("mode", types.YChild{"Mode", &quadConfig.Mode})
    quadConfig.EntityData.Leafs = types.NewOrderedMap()
    quadConfig.EntityData.Leafs.Append("id1", types.YLeaf{"Id1", quadConfig.Id1})

    quadConfig.EntityData.YListKeys = []string {"Id1"}

    return &(quadConfig.EntityData)
}

// Node_Acts_Act_QuadConfigs_QuadConfig_Mode
// select mode 10g or 25g for a quad(group of 4
// ports).
type Node_Acts_Act_QuadConfigs_QuadConfig_Mode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // speed 10g or 25g. The type is string.
    Speed interface{}
}

func (mode *Node_Acts_Act_QuadConfigs_QuadConfig_Mode) GetEntityData() *types.CommonEntityData {
    mode.EntityData.YFilter = mode.YFilter
    mode.EntityData.YangName = "mode"
    mode.EntityData.BundleName = "cisco_ios_xr"
    mode.EntityData.ParentYangName = "quad-config"
    mode.EntityData.SegmentPath = "mode"
    mode.EntityData.AbsolutePath = "Cisco-IOS-XR-optics-driver-quad-cfg:node/acts/act/quad-configs/quad-config/" + mode.EntityData.SegmentPath
    mode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mode.EntityData.Children = types.NewOrderedMap()
    mode.EntityData.Leafs = types.NewOrderedMap()
    mode.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", mode.Speed})

    mode.EntityData.YListKeys = []string {}

    return &(mode.EntityData)
}

