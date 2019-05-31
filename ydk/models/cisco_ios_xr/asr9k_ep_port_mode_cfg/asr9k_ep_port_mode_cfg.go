// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-ep-port-mode package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hw-module-ep-port-mode: HW Module EP port-mode configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_ep_port_mode_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_ep_port_mode_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-ep-port-mode-cfg hw-module-ep-port-mode}", reflect.TypeOf(HwModuleEpPortMode{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-ep-port-mode-cfg:hw-module-ep-port-mode", reflect.TypeOf(HwModuleEpPortMode{}))
}

// HwModuleEpIfPortMode represents Hw module ep if port mode
type HwModuleEpIfPortMode string

const (
    // 2xHundredGigE 16QAM mode configuration
    HwModuleEpIfPortMode_Y_2xhundredgige_16qam HwModuleEpIfPortMode = "2xhundredgige-16qam"

    // 2xHundredGigE 8QAM mode configuration
    HwModuleEpIfPortMode_Y_2xhundredgige_8qam HwModuleEpIfPortMode = "2xhundredgige-8qam"
)

// HwModuleEpPortMode
// HW Module EP port-mode configuration
type HwModuleEpPortMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // active or pre configuration. The type is slice of
    // HwModuleEpPortMode_EpPortModeConfiguration.
    EpPortModeConfiguration []*HwModuleEpPortMode_EpPortModeConfiguration
}

func (hwModuleEpPortMode *HwModuleEpPortMode) GetEntityData() *types.CommonEntityData {
    hwModuleEpPortMode.EntityData.YFilter = hwModuleEpPortMode.YFilter
    hwModuleEpPortMode.EntityData.YangName = "hw-module-ep-port-mode"
    hwModuleEpPortMode.EntityData.BundleName = "cisco_ios_xr"
    hwModuleEpPortMode.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-ep-port-mode-cfg"
    hwModuleEpPortMode.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-ep-port-mode-cfg:hw-module-ep-port-mode"
    hwModuleEpPortMode.EntityData.AbsolutePath = hwModuleEpPortMode.EntityData.SegmentPath
    hwModuleEpPortMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwModuleEpPortMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwModuleEpPortMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwModuleEpPortMode.EntityData.Children = types.NewOrderedMap()
    hwModuleEpPortMode.EntityData.Children.Append("ep-port-mode-configuration", types.YChild{"EpPortModeConfiguration", nil})
    for i := range hwModuleEpPortMode.EpPortModeConfiguration {
        hwModuleEpPortMode.EntityData.Children.Append(types.GetSegmentPath(hwModuleEpPortMode.EpPortModeConfiguration[i]), types.YChild{"EpPortModeConfiguration", hwModuleEpPortMode.EpPortModeConfiguration[i]})
    }
    hwModuleEpPortMode.EntityData.Leafs = types.NewOrderedMap()

    hwModuleEpPortMode.EntityData.YListKeys = []string {}

    return &(hwModuleEpPortMode.EntityData)
}

// HwModuleEpPortMode_EpPortModeConfiguration
// active or pre configuration
type HwModuleEpPortMode_EpPortModeConfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. act or pre configuration. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Active interface{}

    // line-card node location. The type is slice of
    // HwModuleEpPortMode_EpPortModeConfiguration_Node.
    Node []*HwModuleEpPortMode_EpPortModeConfiguration_Node
}

func (epPortModeConfiguration *HwModuleEpPortMode_EpPortModeConfiguration) GetEntityData() *types.CommonEntityData {
    epPortModeConfiguration.EntityData.YFilter = epPortModeConfiguration.YFilter
    epPortModeConfiguration.EntityData.YangName = "ep-port-mode-configuration"
    epPortModeConfiguration.EntityData.BundleName = "cisco_ios_xr"
    epPortModeConfiguration.EntityData.ParentYangName = "hw-module-ep-port-mode"
    epPortModeConfiguration.EntityData.SegmentPath = "ep-port-mode-configuration" + types.AddKeyToken(epPortModeConfiguration.Active, "active")
    epPortModeConfiguration.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-ep-port-mode-cfg:hw-module-ep-port-mode/" + epPortModeConfiguration.EntityData.SegmentPath
    epPortModeConfiguration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    epPortModeConfiguration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    epPortModeConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    epPortModeConfiguration.EntityData.Children = types.NewOrderedMap()
    epPortModeConfiguration.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range epPortModeConfiguration.Node {
        epPortModeConfiguration.EntityData.Children.Append(types.GetSegmentPath(epPortModeConfiguration.Node[i]), types.YChild{"Node", epPortModeConfiguration.Node[i]})
    }
    epPortModeConfiguration.EntityData.Leafs = types.NewOrderedMap()
    epPortModeConfiguration.EntityData.Leafs.Append("active", types.YLeaf{"Active", epPortModeConfiguration.Active})

    epPortModeConfiguration.EntityData.YListKeys = []string {"Active"}

    return &(epPortModeConfiguration.EntityData)
}

// HwModuleEpPortMode_EpPortModeConfiguration_Node
// line-card node location
type HwModuleEpPortMode_EpPortModeConfiguration_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Fully qualified line-card location. The type is
    // string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Location interface{}

    // port-mode configuration for EP bay number.
    Bays HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays
}

func (node *HwModuleEpPortMode_EpPortModeConfiguration_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "ep-port-mode-configuration"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Location, "location")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-ep-port-mode-cfg:hw-module-ep-port-mode/ep-port-mode-configuration/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("bays", types.YChild{"Bays", &node.Bays})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("location", types.YLeaf{"Location", node.Location})

    node.EntityData.YListKeys = []string {"Location"}

    return &(node.EntityData)
}

// HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays
// port-mode configuration for EP bay number
type HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EP Bay number. The type is slice of
    // HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay.
    Bay []*HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay
}

func (bays *HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays) GetEntityData() *types.CommonEntityData {
    bays.EntityData.YFilter = bays.YFilter
    bays.EntityData.YangName = "bays"
    bays.EntityData.BundleName = "cisco_ios_xr"
    bays.EntityData.ParentYangName = "node"
    bays.EntityData.SegmentPath = "bays"
    bays.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-ep-port-mode-cfg:hw-module-ep-port-mode/ep-port-mode-configuration/node/" + bays.EntityData.SegmentPath
    bays.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bays.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bays.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bays.EntityData.Children = types.NewOrderedMap()
    bays.EntityData.Children.Append("bay", types.YChild{"Bay", nil})
    for i := range bays.Bay {
        bays.EntityData.Children.Append(types.GetSegmentPath(bays.Bay[i]), types.YChild{"Bay", bays.Bay[i]})
    }
    bays.EntityData.Leafs = types.NewOrderedMap()

    bays.EntityData.YListKeys = []string {}

    return &(bays.EntityData)
}

// HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay
// EP Bay number
type HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. bay number. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    BayNumber interface{}

    // port-mode configuration for port number.
    Ports HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay_Ports
}

func (bay *HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay) GetEntityData() *types.CommonEntityData {
    bay.EntityData.YFilter = bay.YFilter
    bay.EntityData.YangName = "bay"
    bay.EntityData.BundleName = "cisco_ios_xr"
    bay.EntityData.ParentYangName = "bays"
    bay.EntityData.SegmentPath = "bay" + types.AddKeyToken(bay.BayNumber, "bay-number")
    bay.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-ep-port-mode-cfg:hw-module-ep-port-mode/ep-port-mode-configuration/node/bays/" + bay.EntityData.SegmentPath
    bay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bay.EntityData.Children = types.NewOrderedMap()
    bay.EntityData.Children.Append("ports", types.YChild{"Ports", &bay.Ports})
    bay.EntityData.Leafs = types.NewOrderedMap()
    bay.EntityData.Leafs.Append("bay-number", types.YLeaf{"BayNumber", bay.BayNumber})

    bay.EntityData.YListKeys = []string {"BayNumber"}

    return &(bay.EntityData)
}

// HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay_Ports
// port-mode configuration for port number
type HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay_Ports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Optics port number. The type is slice of
    // HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay_Ports_Port.
    Port []*HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay_Ports_Port
}

func (ports *HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay_Ports) GetEntityData() *types.CommonEntityData {
    ports.EntityData.YFilter = ports.YFilter
    ports.EntityData.YangName = "ports"
    ports.EntityData.BundleName = "cisco_ios_xr"
    ports.EntityData.ParentYangName = "bay"
    ports.EntityData.SegmentPath = "ports"
    ports.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-ep-port-mode-cfg:hw-module-ep-port-mode/ep-port-mode-configuration/node/bays/bay/" + ports.EntityData.SegmentPath
    ports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ports.EntityData.Children = types.NewOrderedMap()
    ports.EntityData.Children.Append("port", types.YChild{"Port", nil})
    for i := range ports.Port {
        ports.EntityData.Children.Append(types.GetSegmentPath(ports.Port[i]), types.YChild{"Port", ports.Port[i]})
    }
    ports.EntityData.Leafs = types.NewOrderedMap()

    ports.EntityData.YListKeys = []string {}

    return &(ports.EntityData)
}

// HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay_Ports_Port
// Optics port number
type HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay_Ports_Port struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Optics port number. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    PortNumber interface{}

    // port-mode type. The type is HwModuleEpIfPortMode.
    IfPortMode interface{}
}

func (port *HwModuleEpPortMode_EpPortModeConfiguration_Node_Bays_Bay_Ports_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "ports"
    port.EntityData.SegmentPath = "port" + types.AddKeyToken(port.PortNumber, "port-number")
    port.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-ep-port-mode-cfg:hw-module-ep-port-mode/ep-port-mode-configuration/node/bays/bay/ports/" + port.EntityData.SegmentPath
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = types.NewOrderedMap()
    port.EntityData.Leafs = types.NewOrderedMap()
    port.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", port.PortNumber})
    port.EntityData.Leafs.Append("if-port-mode", types.YLeaf{"IfPortMode", port.IfPortMode})

    port.EntityData.YListKeys = []string {"PortNumber"}

    return &(port.EntityData)
}

