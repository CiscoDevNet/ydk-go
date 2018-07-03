// This module contains a collection of YANG definitions
// for Cisco IOS-XR ethernet-lldp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lldp: Link Layer Discovery Protocol operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_lldp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_lldp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ethernet-lldp-oper lldp}", reflect.TypeOf(Lldp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ethernet-lldp-oper:lldp", reflect.TypeOf(Lldp{}))
}

// LldpL3AddrProtocol represents Lldp l3 addr protocol
type LldpL3AddrProtocol string

const (
    // IPv4
    LldpL3AddrProtocol_ipv4 LldpL3AddrProtocol = "ipv4"

    // IPv6
    LldpL3AddrProtocol_ipv6 LldpL3AddrProtocol = "ipv6"
)

// Lldp
// Link Layer Discovery Protocol operational data
type Lldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global LLDP data.
    GlobalLldp Lldp_GlobalLldp

    // Per node LLDP operational data.
    Nodes Lldp_Nodes
}

func (lldp *Lldp) GetEntityData() *types.CommonEntityData {
    lldp.EntityData.YFilter = lldp.YFilter
    lldp.EntityData.YangName = "lldp"
    lldp.EntityData.BundleName = "cisco_ios_xr"
    lldp.EntityData.ParentYangName = "Cisco-IOS-XR-ethernet-lldp-oper"
    lldp.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-lldp-oper:lldp"
    lldp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldp.EntityData.Children = types.NewOrderedMap()
    lldp.EntityData.Children.Append("global-lldp", types.YChild{"GlobalLldp", &lldp.GlobalLldp})
    lldp.EntityData.Children.Append("nodes", types.YChild{"Nodes", &lldp.Nodes})
    lldp.EntityData.Leafs = types.NewOrderedMap()

    lldp.EntityData.YListKeys = []string {}

    return &(lldp.EntityData)
}

// Lldp_GlobalLldp
// Global LLDP data
type Lldp_GlobalLldp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The LLDP Global Information of this box.
    LldpInfo Lldp_GlobalLldp_LldpInfo
}

func (globalLldp *Lldp_GlobalLldp) GetEntityData() *types.CommonEntityData {
    globalLldp.EntityData.YFilter = globalLldp.YFilter
    globalLldp.EntityData.YangName = "global-lldp"
    globalLldp.EntityData.BundleName = "cisco_ios_xr"
    globalLldp.EntityData.ParentYangName = "lldp"
    globalLldp.EntityData.SegmentPath = "global-lldp"
    globalLldp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalLldp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalLldp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalLldp.EntityData.Children = types.NewOrderedMap()
    globalLldp.EntityData.Children.Append("lldp-info", types.YChild{"LldpInfo", &globalLldp.LldpInfo})
    globalLldp.EntityData.Leafs = types.NewOrderedMap()

    globalLldp.EntityData.YListKeys = []string {}

    return &(globalLldp.EntityData)
}

// Lldp_GlobalLldp_LldpInfo
// The LLDP Global Information of this box
type Lldp_GlobalLldp_LldpInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Chassis identifier. The type is string.
    ChassisId interface{}

    // Chassis ID sub type. The type is interface{} with range: 0..255.
    ChassisIdSubType interface{}

    // System Name. The type is string.
    SystemName interface{}

    // Rate at which LLDP packets re sent (in sec). The type is interface{} with
    // range: 0..4294967295.
    Timer interface{}

    // Length  of time  (in sec)that receiver must keep thispacket. The type is
    // interface{} with range: 0..4294967295.
    HoldTime interface{}

    // Delay (in sec) for LLDPinitialization on anyinterface. The type is
    // interface{} with range: 0..4294967295.
    ReInit interface{}
}

func (lldpInfo *Lldp_GlobalLldp_LldpInfo) GetEntityData() *types.CommonEntityData {
    lldpInfo.EntityData.YFilter = lldpInfo.YFilter
    lldpInfo.EntityData.YangName = "lldp-info"
    lldpInfo.EntityData.BundleName = "cisco_ios_xr"
    lldpInfo.EntityData.ParentYangName = "global-lldp"
    lldpInfo.EntityData.SegmentPath = "lldp-info"
    lldpInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpInfo.EntityData.Children = types.NewOrderedMap()
    lldpInfo.EntityData.Leafs = types.NewOrderedMap()
    lldpInfo.EntityData.Leafs.Append("chassis-id", types.YLeaf{"ChassisId", lldpInfo.ChassisId})
    lldpInfo.EntityData.Leafs.Append("chassis-id-sub-type", types.YLeaf{"ChassisIdSubType", lldpInfo.ChassisIdSubType})
    lldpInfo.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", lldpInfo.SystemName})
    lldpInfo.EntityData.Leafs.Append("timer", types.YLeaf{"Timer", lldpInfo.Timer})
    lldpInfo.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", lldpInfo.HoldTime})
    lldpInfo.EntityData.Leafs.Append("re-init", types.YLeaf{"ReInit", lldpInfo.ReInit})

    lldpInfo.EntityData.YListKeys = []string {}

    return &(lldpInfo.EntityData)
}

// Lldp_Nodes
// Per node LLDP operational data
type Lldp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The LLDP operational data for a particular node. The type is slice of
    // Lldp_Nodes_Node.
    Node []*Lldp_Nodes_Node
}

func (nodes *Lldp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "lldp"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Lldp_Nodes_Node
// The LLDP operational data for a particular node
type Lldp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // The LLDP neighbor tables on this node.
    Neighbors Lldp_Nodes_Node_Neighbors

    // The table of interfaces on which LLDP is running on this node.
    Interfaces Lldp_Nodes_Node_Interfaces

    // The LLDP traffic statistics for this node.
    Statistics Lldp_Nodes_Node_Statistics
}

func (node *Lldp_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &node.Neighbors})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Children.Append("statistics", types.YChild{"Statistics", &node.Statistics})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Lldp_Nodes_Node_Neighbors
// The LLDP neighbor tables on this node
type Lldp_Nodes_Node_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The detailed LLDP neighbor table on this device.
    Devices Lldp_Nodes_Node_Neighbors_Devices

    // The detailed LLDP neighbor table.
    Details Lldp_Nodes_Node_Neighbors_Details

    // The LLDP neighbor summary table.
    Summaries Lldp_Nodes_Node_Neighbors_Summaries
}

func (neighbors *Lldp_Nodes_Node_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "node"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("devices", types.YChild{"Devices", &neighbors.Devices})
    neighbors.EntityData.Children.Append("details", types.YChild{"Details", &neighbors.Details})
    neighbors.EntityData.Children.Append("summaries", types.YChild{"Summaries", &neighbors.Summaries})
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices
// The detailed LLDP neighbor table on this
// device
type Lldp_Nodes_Node_Neighbors_Devices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information about a LLDP neighbor entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device.
    Device []*Lldp_Nodes_Node_Neighbors_Devices_Device
}

func (devices *Lldp_Nodes_Node_Neighbors_Devices) GetEntityData() *types.CommonEntityData {
    devices.EntityData.YFilter = devices.YFilter
    devices.EntityData.YangName = "devices"
    devices.EntityData.BundleName = "cisco_ios_xr"
    devices.EntityData.ParentYangName = "neighbors"
    devices.EntityData.SegmentPath = "devices"
    devices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    devices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    devices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    devices.EntityData.Children = types.NewOrderedMap()
    devices.EntityData.Children.Append("device", types.YChild{"Device", nil})
    for i := range devices.Device {
        devices.EntityData.Children.Append(types.GetSegmentPath(devices.Device[i]), types.YChild{"Device", devices.Device[i]})
    }
    devices.EntityData.Leafs = types.NewOrderedMap()

    devices.EntityData.YListKeys = []string {}

    return &(devices.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device
// Detailed information about a LLDP neighbor
// entry
type Lldp_Nodes_Node_Neighbors_Devices_Device struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // The interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // lldp neighbor. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor.
    LldpNeighbor []*Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor
}

func (device *Lldp_Nodes_Node_Neighbors_Devices_Device) GetEntityData() *types.CommonEntityData {
    device.EntityData.YFilter = device.YFilter
    device.EntityData.YangName = "device"
    device.EntityData.BundleName = "cisco_ios_xr"
    device.EntityData.ParentYangName = "devices"
    device.EntityData.SegmentPath = "device"
    device.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    device.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    device.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    device.EntityData.Children = types.NewOrderedMap()
    device.EntityData.Children.Append("lldp-neighbor", types.YChild{"LldpNeighbor", nil})
    for i := range device.LldpNeighbor {
        device.EntityData.Children.Append(types.GetSegmentPath(device.LldpNeighbor[i]), types.YChild{"LldpNeighbor", device.LldpNeighbor[i]})
    }
    device.EntityData.Leafs = types.NewOrderedMap()
    device.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", device.DeviceId})
    device.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", device.InterfaceName})

    device.EntityData.YListKeys = []string {}

    return &(device.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor
// lldp neighbor
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    ReceivingInterfaceName interface{}

    // Parent Interface the neighbor entry was received on . The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    ReceivingParentInterfaceName interface{}

    // Device identifier. The type is string.
    DeviceId interface{}

    // Chassis id. The type is string.
    ChassisId interface{}

    // Outgoing port identifier. The type is string.
    PortIdDetail interface{}

    // Version number. The type is interface{} with range: 0..255.
    HeaderVersion interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Platform type. The type is string.
    Platform interface{}

    // Detailed neighbor info.
    Detail Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail

    // MIB nieghbor info.
    Mib Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor) GetEntityData() *types.CommonEntityData {
    lldpNeighbor.EntityData.YFilter = lldpNeighbor.YFilter
    lldpNeighbor.EntityData.YangName = "lldp-neighbor"
    lldpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    lldpNeighbor.EntityData.ParentYangName = "device"
    lldpNeighbor.EntityData.SegmentPath = "lldp-neighbor"
    lldpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpNeighbor.EntityData.Children = types.NewOrderedMap()
    lldpNeighbor.EntityData.Children.Append("detail", types.YChild{"Detail", &lldpNeighbor.Detail})
    lldpNeighbor.EntityData.Children.Append("mib", types.YChild{"Mib", &lldpNeighbor.Mib})
    lldpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    lldpNeighbor.EntityData.Leafs.Append("receiving-interface-name", types.YLeaf{"ReceivingInterfaceName", lldpNeighbor.ReceivingInterfaceName})
    lldpNeighbor.EntityData.Leafs.Append("receiving-parent-interface-name", types.YLeaf{"ReceivingParentInterfaceName", lldpNeighbor.ReceivingParentInterfaceName})
    lldpNeighbor.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", lldpNeighbor.DeviceId})
    lldpNeighbor.EntityData.Leafs.Append("chassis-id", types.YLeaf{"ChassisId", lldpNeighbor.ChassisId})
    lldpNeighbor.EntityData.Leafs.Append("port-id-detail", types.YLeaf{"PortIdDetail", lldpNeighbor.PortIdDetail})
    lldpNeighbor.EntityData.Leafs.Append("header-version", types.YLeaf{"HeaderVersion", lldpNeighbor.HeaderVersion})
    lldpNeighbor.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", lldpNeighbor.HoldTime})
    lldpNeighbor.EntityData.Leafs.Append("enabled-capabilities", types.YLeaf{"EnabledCapabilities", lldpNeighbor.EnabledCapabilities})
    lldpNeighbor.EntityData.Leafs.Append("platform", types.YLeaf{"Platform", lldpNeighbor.Platform})

    lldpNeighbor.EntityData.YListKeys = []string {}

    return &(lldpNeighbor.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail
// Detailed neighbor info
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer Mac Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PeerMacAddress interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // System Name. The type is string.
    SystemName interface{}

    // System Description. The type is string.
    SystemDescription interface{}

    // Time remaining. The type is interface{} with range: 0..4294967295.
    TimeRemaining interface{}

    // System Capabilities. The type is string.
    SystemCapabilities interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Auto Negotiation. The type is string.
    AutoNegotiation interface{}

    // Physical media capabilities. The type is string.
    PhysicalMediaCapabilities interface{}

    // Media Attachment Unit type. The type is interface{} with range:
    // 0..4294967295.
    MediaAttachmentUnitType interface{}

    // Vlan ID. The type is interface{} with range: 0..4294967295.
    PortVlanId interface{}

    // Management Addresses.
    NetworkAddresses Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses
}

func (detail *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "lldp-neighbor"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("network-addresses", types.YChild{"NetworkAddresses", &detail.NetworkAddresses})
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("peer-mac-address", types.YLeaf{"PeerMacAddress", detail.PeerMacAddress})
    detail.EntityData.Leafs.Append("port-description", types.YLeaf{"PortDescription", detail.PortDescription})
    detail.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", detail.SystemName})
    detail.EntityData.Leafs.Append("system-description", types.YLeaf{"SystemDescription", detail.SystemDescription})
    detail.EntityData.Leafs.Append("time-remaining", types.YLeaf{"TimeRemaining", detail.TimeRemaining})
    detail.EntityData.Leafs.Append("system-capabilities", types.YLeaf{"SystemCapabilities", detail.SystemCapabilities})
    detail.EntityData.Leafs.Append("enabled-capabilities", types.YLeaf{"EnabledCapabilities", detail.EnabledCapabilities})
    detail.EntityData.Leafs.Append("auto-negotiation", types.YLeaf{"AutoNegotiation", detail.AutoNegotiation})
    detail.EntityData.Leafs.Append("physical-media-capabilities", types.YLeaf{"PhysicalMediaCapabilities", detail.PhysicalMediaCapabilities})
    detail.EntityData.Leafs.Append("media-attachment-unit-type", types.YLeaf{"MediaAttachmentUnitType", detail.MediaAttachmentUnitType})
    detail.EntityData.Leafs.Append("port-vlan-id", types.YLeaf{"PortVlanId", detail.PortVlanId})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses
// Management Addresses
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []*Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "detail"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = types.NewOrderedMap()
    networkAddresses.EntityData.Children.Append("lldp-addr-entry", types.YChild{"LldpAddrEntry", nil})
    for i := range networkAddresses.LldpAddrEntry {
        networkAddresses.EntityData.Children.Append(types.GetSegmentPath(networkAddresses.LldpAddrEntry[i]), types.YChild{"LldpAddrEntry", networkAddresses.LldpAddrEntry[i]})
    }
    networkAddresses.EntityData.Leafs = types.NewOrderedMap()

    networkAddresses.EntityData.YListKeys = []string {}

    return &(networkAddresses.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
// lldp addr entry
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetEntityData() *types.CommonEntityData {
    lldpAddrEntry.EntityData.YFilter = lldpAddrEntry.YFilter
    lldpAddrEntry.EntityData.YangName = "lldp-addr-entry"
    lldpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpAddrEntry.EntityData.ParentYangName = "network-addresses"
    lldpAddrEntry.EntityData.SegmentPath = "lldp-addr-entry"
    lldpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpAddrEntry.EntityData.Children = types.NewOrderedMap()
    lldpAddrEntry.EntityData.Children.Append("address", types.YChild{"Address", &lldpAddrEntry.Address})
    lldpAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpAddrEntry.EntityData.Leafs.Append("ma-subtype", types.YLeaf{"MaSubtype", lldpAddrEntry.MaSubtype})
    lldpAddrEntry.EntityData.Leafs.Append("if-num", types.YLeaf{"IfNum", lldpAddrEntry.IfNum})

    lldpAddrEntry.EntityData.YListKeys = []string {}

    return &(lldpAddrEntry.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
// Network layer address
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "lldp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib
// MIB nieghbor info
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeFilter. The type is interface{} with range: 0..4294967295.
    RemTimeMark interface{}

    // LldpPortNumber. The type is interface{} with range: 0..4294967295.
    RemLocalPortNum interface{}

    // lldpRemIndex. The type is interface{} with range: 0..4294967295.
    RemIndex interface{}

    // Chassis ID sub type. The type is interface{} with range: 0..255.
    ChassisIdSubType interface{}

    // Chassis ID length. The type is interface{} with range: 0..65535.
    ChassisIdLen interface{}

    // Port ID sub type. The type is interface{} with range: 0..255.
    PortIdSubType interface{}

    // Port ID length. The type is interface{} with range: 0..65535.
    PortIdLen interface{}

    // Supported and combined cpabilities. The type is interface{} with range:
    // 0..4294967295.
    CombinedCapabilities interface{}

    // Unknown TLV list.
    UnknownTlvList Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList

    // Org Def TLV list.
    OrgDefTlvList Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList
}

func (mib *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib) GetEntityData() *types.CommonEntityData {
    mib.EntityData.YFilter = mib.YFilter
    mib.EntityData.YangName = "mib"
    mib.EntityData.BundleName = "cisco_ios_xr"
    mib.EntityData.ParentYangName = "lldp-neighbor"
    mib.EntityData.SegmentPath = "mib"
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = types.NewOrderedMap()
    mib.EntityData.Children.Append("unknown-tlv-list", types.YChild{"UnknownTlvList", &mib.UnknownTlvList})
    mib.EntityData.Children.Append("org-def-tlv-list", types.YChild{"OrgDefTlvList", &mib.OrgDefTlvList})
    mib.EntityData.Leafs = types.NewOrderedMap()
    mib.EntityData.Leafs.Append("rem-time-mark", types.YLeaf{"RemTimeMark", mib.RemTimeMark})
    mib.EntityData.Leafs.Append("rem-local-port-num", types.YLeaf{"RemLocalPortNum", mib.RemLocalPortNum})
    mib.EntityData.Leafs.Append("rem-index", types.YLeaf{"RemIndex", mib.RemIndex})
    mib.EntityData.Leafs.Append("chassis-id-sub-type", types.YLeaf{"ChassisIdSubType", mib.ChassisIdSubType})
    mib.EntityData.Leafs.Append("chassis-id-len", types.YLeaf{"ChassisIdLen", mib.ChassisIdLen})
    mib.EntityData.Leafs.Append("port-id-sub-type", types.YLeaf{"PortIdSubType", mib.PortIdSubType})
    mib.EntityData.Leafs.Append("port-id-len", types.YLeaf{"PortIdLen", mib.PortIdLen})
    mib.EntityData.Leafs.Append("combined-capabilities", types.YLeaf{"CombinedCapabilities", mib.CombinedCapabilities})

    mib.EntityData.YListKeys = []string {}

    return &(mib.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList
// Unknown TLV list
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp unknown tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry.
    LldpUnknownTlvEntry []*Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList) GetEntityData() *types.CommonEntityData {
    unknownTlvList.EntityData.YFilter = unknownTlvList.YFilter
    unknownTlvList.EntityData.YangName = "unknown-tlv-list"
    unknownTlvList.EntityData.BundleName = "cisco_ios_xr"
    unknownTlvList.EntityData.ParentYangName = "mib"
    unknownTlvList.EntityData.SegmentPath = "unknown-tlv-list"
    unknownTlvList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownTlvList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownTlvList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownTlvList.EntityData.Children = types.NewOrderedMap()
    unknownTlvList.EntityData.Children.Append("lldp-unknown-tlv-entry", types.YChild{"LldpUnknownTlvEntry", nil})
    for i := range unknownTlvList.LldpUnknownTlvEntry {
        unknownTlvList.EntityData.Children.Append(types.GetSegmentPath(unknownTlvList.LldpUnknownTlvEntry[i]), types.YChild{"LldpUnknownTlvEntry", unknownTlvList.LldpUnknownTlvEntry[i]})
    }
    unknownTlvList.EntityData.Leafs = types.NewOrderedMap()

    unknownTlvList.EntityData.YListKeys = []string {}

    return &(unknownTlvList.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
// lldp unknown tlv entry
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown TLV type. The type is interface{} with range: 0..255.
    TlvType interface{}

    // Unknown TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetEntityData() *types.CommonEntityData {
    lldpUnknownTlvEntry.EntityData.YFilter = lldpUnknownTlvEntry.YFilter
    lldpUnknownTlvEntry.EntityData.YangName = "lldp-unknown-tlv-entry"
    lldpUnknownTlvEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpUnknownTlvEntry.EntityData.ParentYangName = "unknown-tlv-list"
    lldpUnknownTlvEntry.EntityData.SegmentPath = "lldp-unknown-tlv-entry"
    lldpUnknownTlvEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpUnknownTlvEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpUnknownTlvEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpUnknownTlvEntry.EntityData.Children = types.NewOrderedMap()
    lldpUnknownTlvEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpUnknownTlvEntry.EntityData.Leafs.Append("tlv-type", types.YLeaf{"TlvType", lldpUnknownTlvEntry.TlvType})
    lldpUnknownTlvEntry.EntityData.Leafs.Append("tlv-value", types.YLeaf{"TlvValue", lldpUnknownTlvEntry.TlvValue})

    lldpUnknownTlvEntry.EntityData.YListKeys = []string {}

    return &(lldpUnknownTlvEntry.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList
// Org Def TLV list
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp org def tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry.
    LldpOrgDefTlvEntry []*Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList) GetEntityData() *types.CommonEntityData {
    orgDefTlvList.EntityData.YFilter = orgDefTlvList.YFilter
    orgDefTlvList.EntityData.YangName = "org-def-tlv-list"
    orgDefTlvList.EntityData.BundleName = "cisco_ios_xr"
    orgDefTlvList.EntityData.ParentYangName = "mib"
    orgDefTlvList.EntityData.SegmentPath = "org-def-tlv-list"
    orgDefTlvList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    orgDefTlvList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    orgDefTlvList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    orgDefTlvList.EntityData.Children = types.NewOrderedMap()
    orgDefTlvList.EntityData.Children.Append("lldp-org-def-tlv-entry", types.YChild{"LldpOrgDefTlvEntry", nil})
    for i := range orgDefTlvList.LldpOrgDefTlvEntry {
        orgDefTlvList.EntityData.Children.Append(types.GetSegmentPath(orgDefTlvList.LldpOrgDefTlvEntry[i]), types.YChild{"LldpOrgDefTlvEntry", orgDefTlvList.LldpOrgDefTlvEntry[i]})
    }
    orgDefTlvList.EntityData.Leafs = types.NewOrderedMap()

    orgDefTlvList.EntityData.YListKeys = []string {}

    return &(orgDefTlvList.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
// lldp org def tlv entry
type Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Organizationally Unique Identifier. The type is interface{} with range:
    // 0..4294967295.
    Oui interface{}

    // Org Def TLV subtype. The type is interface{} with range: 0..255.
    TlvSubtype interface{}

    // lldpRemOrgDefInfoIndex. The type is interface{} with range: 0..4294967295.
    TlvInfoIndes interface{}

    // Org Def TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Devices_Device_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetEntityData() *types.CommonEntityData {
    lldpOrgDefTlvEntry.EntityData.YFilter = lldpOrgDefTlvEntry.YFilter
    lldpOrgDefTlvEntry.EntityData.YangName = "lldp-org-def-tlv-entry"
    lldpOrgDefTlvEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpOrgDefTlvEntry.EntityData.ParentYangName = "org-def-tlv-list"
    lldpOrgDefTlvEntry.EntityData.SegmentPath = "lldp-org-def-tlv-entry"
    lldpOrgDefTlvEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpOrgDefTlvEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpOrgDefTlvEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpOrgDefTlvEntry.EntityData.Children = types.NewOrderedMap()
    lldpOrgDefTlvEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("oui", types.YLeaf{"Oui", lldpOrgDefTlvEntry.Oui})
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("tlv-subtype", types.YLeaf{"TlvSubtype", lldpOrgDefTlvEntry.TlvSubtype})
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("tlv-info-indes", types.YLeaf{"TlvInfoIndes", lldpOrgDefTlvEntry.TlvInfoIndes})
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("tlv-value", types.YLeaf{"TlvValue", lldpOrgDefTlvEntry.TlvValue})

    lldpOrgDefTlvEntry.EntityData.YListKeys = []string {}

    return &(lldpOrgDefTlvEntry.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details
// The detailed LLDP neighbor table
type Lldp_Nodes_Node_Neighbors_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information about a LLDP neighbor entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail.
    Detail []*Lldp_Nodes_Node_Neighbors_Details_Detail
}

func (details *Lldp_Nodes_Node_Neighbors_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "neighbors"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = types.NewOrderedMap()
    details.EntityData.Children.Append("detail", types.YChild{"Detail", nil})
    for i := range details.Detail {
        details.EntityData.Children.Append(types.GetSegmentPath(details.Detail[i]), types.YChild{"Detail", details.Detail[i]})
    }
    details.EntityData.Leafs = types.NewOrderedMap()

    details.EntityData.YListKeys = []string {}

    return &(details.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail
// Detailed information about a LLDP neighbor
// entry
type Lldp_Nodes_Node_Neighbors_Details_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // lldp neighbor. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor.
    LldpNeighbor []*Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "details"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("lldp-neighbor", types.YChild{"LldpNeighbor", nil})
    for i := range detail.LldpNeighbor {
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.LldpNeighbor[i]), types.YChild{"LldpNeighbor", detail.LldpNeighbor[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", detail.InterfaceName})
    detail.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", detail.DeviceId})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor
// lldp neighbor
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    ReceivingInterfaceName interface{}

    // Parent Interface the neighbor entry was received on . The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    ReceivingParentInterfaceName interface{}

    // Device identifier. The type is string.
    DeviceId interface{}

    // Chassis id. The type is string.
    ChassisId interface{}

    // Outgoing port identifier. The type is string.
    PortIdDetail interface{}

    // Version number. The type is interface{} with range: 0..255.
    HeaderVersion interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Platform type. The type is string.
    Platform interface{}

    // Detailed neighbor info.
    Detail Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail

    // MIB nieghbor info.
    Mib Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor) GetEntityData() *types.CommonEntityData {
    lldpNeighbor.EntityData.YFilter = lldpNeighbor.YFilter
    lldpNeighbor.EntityData.YangName = "lldp-neighbor"
    lldpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    lldpNeighbor.EntityData.ParentYangName = "detail"
    lldpNeighbor.EntityData.SegmentPath = "lldp-neighbor"
    lldpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpNeighbor.EntityData.Children = types.NewOrderedMap()
    lldpNeighbor.EntityData.Children.Append("detail", types.YChild{"Detail", &lldpNeighbor.Detail})
    lldpNeighbor.EntityData.Children.Append("mib", types.YChild{"Mib", &lldpNeighbor.Mib})
    lldpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    lldpNeighbor.EntityData.Leafs.Append("receiving-interface-name", types.YLeaf{"ReceivingInterfaceName", lldpNeighbor.ReceivingInterfaceName})
    lldpNeighbor.EntityData.Leafs.Append("receiving-parent-interface-name", types.YLeaf{"ReceivingParentInterfaceName", lldpNeighbor.ReceivingParentInterfaceName})
    lldpNeighbor.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", lldpNeighbor.DeviceId})
    lldpNeighbor.EntityData.Leafs.Append("chassis-id", types.YLeaf{"ChassisId", lldpNeighbor.ChassisId})
    lldpNeighbor.EntityData.Leafs.Append("port-id-detail", types.YLeaf{"PortIdDetail", lldpNeighbor.PortIdDetail})
    lldpNeighbor.EntityData.Leafs.Append("header-version", types.YLeaf{"HeaderVersion", lldpNeighbor.HeaderVersion})
    lldpNeighbor.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", lldpNeighbor.HoldTime})
    lldpNeighbor.EntityData.Leafs.Append("enabled-capabilities", types.YLeaf{"EnabledCapabilities", lldpNeighbor.EnabledCapabilities})
    lldpNeighbor.EntityData.Leafs.Append("platform", types.YLeaf{"Platform", lldpNeighbor.Platform})

    lldpNeighbor.EntityData.YListKeys = []string {}

    return &(lldpNeighbor.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail
// Detailed neighbor info
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer Mac Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PeerMacAddress interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // System Name. The type is string.
    SystemName interface{}

    // System Description. The type is string.
    SystemDescription interface{}

    // Time remaining. The type is interface{} with range: 0..4294967295.
    TimeRemaining interface{}

    // System Capabilities. The type is string.
    SystemCapabilities interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Auto Negotiation. The type is string.
    AutoNegotiation interface{}

    // Physical media capabilities. The type is string.
    PhysicalMediaCapabilities interface{}

    // Media Attachment Unit type. The type is interface{} with range:
    // 0..4294967295.
    MediaAttachmentUnitType interface{}

    // Vlan ID. The type is interface{} with range: 0..4294967295.
    PortVlanId interface{}

    // Management Addresses.
    NetworkAddresses Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses
}

func (detail *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "lldp-neighbor"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("network-addresses", types.YChild{"NetworkAddresses", &detail.NetworkAddresses})
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("peer-mac-address", types.YLeaf{"PeerMacAddress", detail.PeerMacAddress})
    detail.EntityData.Leafs.Append("port-description", types.YLeaf{"PortDescription", detail.PortDescription})
    detail.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", detail.SystemName})
    detail.EntityData.Leafs.Append("system-description", types.YLeaf{"SystemDescription", detail.SystemDescription})
    detail.EntityData.Leafs.Append("time-remaining", types.YLeaf{"TimeRemaining", detail.TimeRemaining})
    detail.EntityData.Leafs.Append("system-capabilities", types.YLeaf{"SystemCapabilities", detail.SystemCapabilities})
    detail.EntityData.Leafs.Append("enabled-capabilities", types.YLeaf{"EnabledCapabilities", detail.EnabledCapabilities})
    detail.EntityData.Leafs.Append("auto-negotiation", types.YLeaf{"AutoNegotiation", detail.AutoNegotiation})
    detail.EntityData.Leafs.Append("physical-media-capabilities", types.YLeaf{"PhysicalMediaCapabilities", detail.PhysicalMediaCapabilities})
    detail.EntityData.Leafs.Append("media-attachment-unit-type", types.YLeaf{"MediaAttachmentUnitType", detail.MediaAttachmentUnitType})
    detail.EntityData.Leafs.Append("port-vlan-id", types.YLeaf{"PortVlanId", detail.PortVlanId})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses
// Management Addresses
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []*Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "detail"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = types.NewOrderedMap()
    networkAddresses.EntityData.Children.Append("lldp-addr-entry", types.YChild{"LldpAddrEntry", nil})
    for i := range networkAddresses.LldpAddrEntry {
        networkAddresses.EntityData.Children.Append(types.GetSegmentPath(networkAddresses.LldpAddrEntry[i]), types.YChild{"LldpAddrEntry", networkAddresses.LldpAddrEntry[i]})
    }
    networkAddresses.EntityData.Leafs = types.NewOrderedMap()

    networkAddresses.EntityData.YListKeys = []string {}

    return &(networkAddresses.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
// lldp addr entry
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetEntityData() *types.CommonEntityData {
    lldpAddrEntry.EntityData.YFilter = lldpAddrEntry.YFilter
    lldpAddrEntry.EntityData.YangName = "lldp-addr-entry"
    lldpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpAddrEntry.EntityData.ParentYangName = "network-addresses"
    lldpAddrEntry.EntityData.SegmentPath = "lldp-addr-entry"
    lldpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpAddrEntry.EntityData.Children = types.NewOrderedMap()
    lldpAddrEntry.EntityData.Children.Append("address", types.YChild{"Address", &lldpAddrEntry.Address})
    lldpAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpAddrEntry.EntityData.Leafs.Append("ma-subtype", types.YLeaf{"MaSubtype", lldpAddrEntry.MaSubtype})
    lldpAddrEntry.EntityData.Leafs.Append("if-num", types.YLeaf{"IfNum", lldpAddrEntry.IfNum})

    lldpAddrEntry.EntityData.YListKeys = []string {}

    return &(lldpAddrEntry.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
// Network layer address
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "lldp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib
// MIB nieghbor info
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeFilter. The type is interface{} with range: 0..4294967295.
    RemTimeMark interface{}

    // LldpPortNumber. The type is interface{} with range: 0..4294967295.
    RemLocalPortNum interface{}

    // lldpRemIndex. The type is interface{} with range: 0..4294967295.
    RemIndex interface{}

    // Chassis ID sub type. The type is interface{} with range: 0..255.
    ChassisIdSubType interface{}

    // Chassis ID length. The type is interface{} with range: 0..65535.
    ChassisIdLen interface{}

    // Port ID sub type. The type is interface{} with range: 0..255.
    PortIdSubType interface{}

    // Port ID length. The type is interface{} with range: 0..65535.
    PortIdLen interface{}

    // Supported and combined cpabilities. The type is interface{} with range:
    // 0..4294967295.
    CombinedCapabilities interface{}

    // Unknown TLV list.
    UnknownTlvList Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList

    // Org Def TLV list.
    OrgDefTlvList Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList
}

func (mib *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib) GetEntityData() *types.CommonEntityData {
    mib.EntityData.YFilter = mib.YFilter
    mib.EntityData.YangName = "mib"
    mib.EntityData.BundleName = "cisco_ios_xr"
    mib.EntityData.ParentYangName = "lldp-neighbor"
    mib.EntityData.SegmentPath = "mib"
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = types.NewOrderedMap()
    mib.EntityData.Children.Append("unknown-tlv-list", types.YChild{"UnknownTlvList", &mib.UnknownTlvList})
    mib.EntityData.Children.Append("org-def-tlv-list", types.YChild{"OrgDefTlvList", &mib.OrgDefTlvList})
    mib.EntityData.Leafs = types.NewOrderedMap()
    mib.EntityData.Leafs.Append("rem-time-mark", types.YLeaf{"RemTimeMark", mib.RemTimeMark})
    mib.EntityData.Leafs.Append("rem-local-port-num", types.YLeaf{"RemLocalPortNum", mib.RemLocalPortNum})
    mib.EntityData.Leafs.Append("rem-index", types.YLeaf{"RemIndex", mib.RemIndex})
    mib.EntityData.Leafs.Append("chassis-id-sub-type", types.YLeaf{"ChassisIdSubType", mib.ChassisIdSubType})
    mib.EntityData.Leafs.Append("chassis-id-len", types.YLeaf{"ChassisIdLen", mib.ChassisIdLen})
    mib.EntityData.Leafs.Append("port-id-sub-type", types.YLeaf{"PortIdSubType", mib.PortIdSubType})
    mib.EntityData.Leafs.Append("port-id-len", types.YLeaf{"PortIdLen", mib.PortIdLen})
    mib.EntityData.Leafs.Append("combined-capabilities", types.YLeaf{"CombinedCapabilities", mib.CombinedCapabilities})

    mib.EntityData.YListKeys = []string {}

    return &(mib.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList
// Unknown TLV list
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp unknown tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry.
    LldpUnknownTlvEntry []*Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList) GetEntityData() *types.CommonEntityData {
    unknownTlvList.EntityData.YFilter = unknownTlvList.YFilter
    unknownTlvList.EntityData.YangName = "unknown-tlv-list"
    unknownTlvList.EntityData.BundleName = "cisco_ios_xr"
    unknownTlvList.EntityData.ParentYangName = "mib"
    unknownTlvList.EntityData.SegmentPath = "unknown-tlv-list"
    unknownTlvList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownTlvList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownTlvList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownTlvList.EntityData.Children = types.NewOrderedMap()
    unknownTlvList.EntityData.Children.Append("lldp-unknown-tlv-entry", types.YChild{"LldpUnknownTlvEntry", nil})
    for i := range unknownTlvList.LldpUnknownTlvEntry {
        unknownTlvList.EntityData.Children.Append(types.GetSegmentPath(unknownTlvList.LldpUnknownTlvEntry[i]), types.YChild{"LldpUnknownTlvEntry", unknownTlvList.LldpUnknownTlvEntry[i]})
    }
    unknownTlvList.EntityData.Leafs = types.NewOrderedMap()

    unknownTlvList.EntityData.YListKeys = []string {}

    return &(unknownTlvList.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
// lldp unknown tlv entry
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown TLV type. The type is interface{} with range: 0..255.
    TlvType interface{}

    // Unknown TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetEntityData() *types.CommonEntityData {
    lldpUnknownTlvEntry.EntityData.YFilter = lldpUnknownTlvEntry.YFilter
    lldpUnknownTlvEntry.EntityData.YangName = "lldp-unknown-tlv-entry"
    lldpUnknownTlvEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpUnknownTlvEntry.EntityData.ParentYangName = "unknown-tlv-list"
    lldpUnknownTlvEntry.EntityData.SegmentPath = "lldp-unknown-tlv-entry"
    lldpUnknownTlvEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpUnknownTlvEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpUnknownTlvEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpUnknownTlvEntry.EntityData.Children = types.NewOrderedMap()
    lldpUnknownTlvEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpUnknownTlvEntry.EntityData.Leafs.Append("tlv-type", types.YLeaf{"TlvType", lldpUnknownTlvEntry.TlvType})
    lldpUnknownTlvEntry.EntityData.Leafs.Append("tlv-value", types.YLeaf{"TlvValue", lldpUnknownTlvEntry.TlvValue})

    lldpUnknownTlvEntry.EntityData.YListKeys = []string {}

    return &(lldpUnknownTlvEntry.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList
// Org Def TLV list
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp org def tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry.
    LldpOrgDefTlvEntry []*Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList) GetEntityData() *types.CommonEntityData {
    orgDefTlvList.EntityData.YFilter = orgDefTlvList.YFilter
    orgDefTlvList.EntityData.YangName = "org-def-tlv-list"
    orgDefTlvList.EntityData.BundleName = "cisco_ios_xr"
    orgDefTlvList.EntityData.ParentYangName = "mib"
    orgDefTlvList.EntityData.SegmentPath = "org-def-tlv-list"
    orgDefTlvList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    orgDefTlvList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    orgDefTlvList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    orgDefTlvList.EntityData.Children = types.NewOrderedMap()
    orgDefTlvList.EntityData.Children.Append("lldp-org-def-tlv-entry", types.YChild{"LldpOrgDefTlvEntry", nil})
    for i := range orgDefTlvList.LldpOrgDefTlvEntry {
        orgDefTlvList.EntityData.Children.Append(types.GetSegmentPath(orgDefTlvList.LldpOrgDefTlvEntry[i]), types.YChild{"LldpOrgDefTlvEntry", orgDefTlvList.LldpOrgDefTlvEntry[i]})
    }
    orgDefTlvList.EntityData.Leafs = types.NewOrderedMap()

    orgDefTlvList.EntityData.YListKeys = []string {}

    return &(orgDefTlvList.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
// lldp org def tlv entry
type Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Organizationally Unique Identifier. The type is interface{} with range:
    // 0..4294967295.
    Oui interface{}

    // Org Def TLV subtype. The type is interface{} with range: 0..255.
    TlvSubtype interface{}

    // lldpRemOrgDefInfoIndex. The type is interface{} with range: 0..4294967295.
    TlvInfoIndes interface{}

    // Org Def TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Details_Detail_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetEntityData() *types.CommonEntityData {
    lldpOrgDefTlvEntry.EntityData.YFilter = lldpOrgDefTlvEntry.YFilter
    lldpOrgDefTlvEntry.EntityData.YangName = "lldp-org-def-tlv-entry"
    lldpOrgDefTlvEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpOrgDefTlvEntry.EntityData.ParentYangName = "org-def-tlv-list"
    lldpOrgDefTlvEntry.EntityData.SegmentPath = "lldp-org-def-tlv-entry"
    lldpOrgDefTlvEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpOrgDefTlvEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpOrgDefTlvEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpOrgDefTlvEntry.EntityData.Children = types.NewOrderedMap()
    lldpOrgDefTlvEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("oui", types.YLeaf{"Oui", lldpOrgDefTlvEntry.Oui})
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("tlv-subtype", types.YLeaf{"TlvSubtype", lldpOrgDefTlvEntry.TlvSubtype})
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("tlv-info-indes", types.YLeaf{"TlvInfoIndes", lldpOrgDefTlvEntry.TlvInfoIndes})
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("tlv-value", types.YLeaf{"TlvValue", lldpOrgDefTlvEntry.TlvValue})

    lldpOrgDefTlvEntry.EntityData.YListKeys = []string {}

    return &(lldpOrgDefTlvEntry.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries
// The LLDP neighbor summary table
type Lldp_Nodes_Node_Neighbors_Summaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information about a LLDP neighbor entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary.
    Summary []*Lldp_Nodes_Node_Neighbors_Summaries_Summary
}

func (summaries *Lldp_Nodes_Node_Neighbors_Summaries) GetEntityData() *types.CommonEntityData {
    summaries.EntityData.YFilter = summaries.YFilter
    summaries.EntityData.YangName = "summaries"
    summaries.EntityData.BundleName = "cisco_ios_xr"
    summaries.EntityData.ParentYangName = "neighbors"
    summaries.EntityData.SegmentPath = "summaries"
    summaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaries.EntityData.Children = types.NewOrderedMap()
    summaries.EntityData.Children.Append("summary", types.YChild{"Summary", nil})
    for i := range summaries.Summary {
        summaries.EntityData.Children.Append(types.GetSegmentPath(summaries.Summary[i]), types.YChild{"Summary", summaries.Summary[i]})
    }
    summaries.EntityData.Leafs = types.NewOrderedMap()

    summaries.EntityData.YListKeys = []string {}

    return &(summaries.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary
// Brief information about a LLDP neighbor
// entry
type Lldp_Nodes_Node_Neighbors_Summaries_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // The neighboring device identifier. The type is string.
    DeviceId interface{}

    // lldp neighbor. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor.
    LldpNeighbor []*Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor
}

func (summary *Lldp_Nodes_Node_Neighbors_Summaries_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "summaries"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("lldp-neighbor", types.YChild{"LldpNeighbor", nil})
    for i := range summary.LldpNeighbor {
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.LldpNeighbor[i]), types.YChild{"LldpNeighbor", summary.LldpNeighbor[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", summary.InterfaceName})
    summary.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", summary.DeviceId})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor
// lldp neighbor
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface the neighbor entry was received on . The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    ReceivingInterfaceName interface{}

    // Parent Interface the neighbor entry was received on . The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    ReceivingParentInterfaceName interface{}

    // Device identifier. The type is string.
    DeviceId interface{}

    // Chassis id. The type is string.
    ChassisId interface{}

    // Outgoing port identifier. The type is string.
    PortIdDetail interface{}

    // Version number. The type is interface{} with range: 0..255.
    HeaderVersion interface{}

    // Remaining hold time. The type is interface{} with range: 0..65535.
    HoldTime interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Platform type. The type is string.
    Platform interface{}

    // Detailed neighbor info.
    Detail Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail

    // MIB nieghbor info.
    Mib Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib
}

func (lldpNeighbor *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor) GetEntityData() *types.CommonEntityData {
    lldpNeighbor.EntityData.YFilter = lldpNeighbor.YFilter
    lldpNeighbor.EntityData.YangName = "lldp-neighbor"
    lldpNeighbor.EntityData.BundleName = "cisco_ios_xr"
    lldpNeighbor.EntityData.ParentYangName = "summary"
    lldpNeighbor.EntityData.SegmentPath = "lldp-neighbor"
    lldpNeighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpNeighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpNeighbor.EntityData.Children = types.NewOrderedMap()
    lldpNeighbor.EntityData.Children.Append("detail", types.YChild{"Detail", &lldpNeighbor.Detail})
    lldpNeighbor.EntityData.Children.Append("mib", types.YChild{"Mib", &lldpNeighbor.Mib})
    lldpNeighbor.EntityData.Leafs = types.NewOrderedMap()
    lldpNeighbor.EntityData.Leafs.Append("receiving-interface-name", types.YLeaf{"ReceivingInterfaceName", lldpNeighbor.ReceivingInterfaceName})
    lldpNeighbor.EntityData.Leafs.Append("receiving-parent-interface-name", types.YLeaf{"ReceivingParentInterfaceName", lldpNeighbor.ReceivingParentInterfaceName})
    lldpNeighbor.EntityData.Leafs.Append("device-id", types.YLeaf{"DeviceId", lldpNeighbor.DeviceId})
    lldpNeighbor.EntityData.Leafs.Append("chassis-id", types.YLeaf{"ChassisId", lldpNeighbor.ChassisId})
    lldpNeighbor.EntityData.Leafs.Append("port-id-detail", types.YLeaf{"PortIdDetail", lldpNeighbor.PortIdDetail})
    lldpNeighbor.EntityData.Leafs.Append("header-version", types.YLeaf{"HeaderVersion", lldpNeighbor.HeaderVersion})
    lldpNeighbor.EntityData.Leafs.Append("hold-time", types.YLeaf{"HoldTime", lldpNeighbor.HoldTime})
    lldpNeighbor.EntityData.Leafs.Append("enabled-capabilities", types.YLeaf{"EnabledCapabilities", lldpNeighbor.EnabledCapabilities})
    lldpNeighbor.EntityData.Leafs.Append("platform", types.YLeaf{"Platform", lldpNeighbor.Platform})

    lldpNeighbor.EntityData.YListKeys = []string {}

    return &(lldpNeighbor.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail
// Detailed neighbor info
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer Mac Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    PeerMacAddress interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // System Name. The type is string.
    SystemName interface{}

    // System Description. The type is string.
    SystemDescription interface{}

    // Time remaining. The type is interface{} with range: 0..4294967295.
    TimeRemaining interface{}

    // System Capabilities. The type is string.
    SystemCapabilities interface{}

    // Enabled Capabilities. The type is string.
    EnabledCapabilities interface{}

    // Auto Negotiation. The type is string.
    AutoNegotiation interface{}

    // Physical media capabilities. The type is string.
    PhysicalMediaCapabilities interface{}

    // Media Attachment Unit type. The type is interface{} with range:
    // 0..4294967295.
    MediaAttachmentUnitType interface{}

    // Vlan ID. The type is interface{} with range: 0..4294967295.
    PortVlanId interface{}

    // Management Addresses.
    NetworkAddresses Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses
}

func (detail *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "lldp-neighbor"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("network-addresses", types.YChild{"NetworkAddresses", &detail.NetworkAddresses})
    detail.EntityData.Leafs = types.NewOrderedMap()
    detail.EntityData.Leafs.Append("peer-mac-address", types.YLeaf{"PeerMacAddress", detail.PeerMacAddress})
    detail.EntityData.Leafs.Append("port-description", types.YLeaf{"PortDescription", detail.PortDescription})
    detail.EntityData.Leafs.Append("system-name", types.YLeaf{"SystemName", detail.SystemName})
    detail.EntityData.Leafs.Append("system-description", types.YLeaf{"SystemDescription", detail.SystemDescription})
    detail.EntityData.Leafs.Append("time-remaining", types.YLeaf{"TimeRemaining", detail.TimeRemaining})
    detail.EntityData.Leafs.Append("system-capabilities", types.YLeaf{"SystemCapabilities", detail.SystemCapabilities})
    detail.EntityData.Leafs.Append("enabled-capabilities", types.YLeaf{"EnabledCapabilities", detail.EnabledCapabilities})
    detail.EntityData.Leafs.Append("auto-negotiation", types.YLeaf{"AutoNegotiation", detail.AutoNegotiation})
    detail.EntityData.Leafs.Append("physical-media-capabilities", types.YLeaf{"PhysicalMediaCapabilities", detail.PhysicalMediaCapabilities})
    detail.EntityData.Leafs.Append("media-attachment-unit-type", types.YLeaf{"MediaAttachmentUnitType", detail.MediaAttachmentUnitType})
    detail.EntityData.Leafs.Append("port-vlan-id", types.YLeaf{"PortVlanId", detail.PortVlanId})

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses
// Management Addresses
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []*Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
}

func (networkAddresses *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses) GetEntityData() *types.CommonEntityData {
    networkAddresses.EntityData.YFilter = networkAddresses.YFilter
    networkAddresses.EntityData.YangName = "network-addresses"
    networkAddresses.EntityData.BundleName = "cisco_ios_xr"
    networkAddresses.EntityData.ParentYangName = "detail"
    networkAddresses.EntityData.SegmentPath = "network-addresses"
    networkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAddresses.EntityData.Children = types.NewOrderedMap()
    networkAddresses.EntityData.Children.Append("lldp-addr-entry", types.YChild{"LldpAddrEntry", nil})
    for i := range networkAddresses.LldpAddrEntry {
        networkAddresses.EntityData.Children.Append(types.GetSegmentPath(networkAddresses.LldpAddrEntry[i]), types.YChild{"LldpAddrEntry", networkAddresses.LldpAddrEntry[i]})
    }
    networkAddresses.EntityData.Leafs = types.NewOrderedMap()

    networkAddresses.EntityData.YListKeys = []string {}

    return &(networkAddresses.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry
// lldp addr entry
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry) GetEntityData() *types.CommonEntityData {
    lldpAddrEntry.EntityData.YFilter = lldpAddrEntry.YFilter
    lldpAddrEntry.EntityData.YangName = "lldp-addr-entry"
    lldpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpAddrEntry.EntityData.ParentYangName = "network-addresses"
    lldpAddrEntry.EntityData.SegmentPath = "lldp-addr-entry"
    lldpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpAddrEntry.EntityData.Children = types.NewOrderedMap()
    lldpAddrEntry.EntityData.Children.Append("address", types.YChild{"Address", &lldpAddrEntry.Address})
    lldpAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpAddrEntry.EntityData.Leafs.Append("ma-subtype", types.YLeaf{"MaSubtype", lldpAddrEntry.MaSubtype})
    lldpAddrEntry.EntityData.Leafs.Append("if-num", types.YLeaf{"IfNum", lldpAddrEntry.IfNum})

    lldpAddrEntry.EntityData.YListKeys = []string {}

    return &(lldpAddrEntry.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address
// Network layer address
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Detail_NetworkAddresses_LldpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "lldp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib
// MIB nieghbor info
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TimeFilter. The type is interface{} with range: 0..4294967295.
    RemTimeMark interface{}

    // LldpPortNumber. The type is interface{} with range: 0..4294967295.
    RemLocalPortNum interface{}

    // lldpRemIndex. The type is interface{} with range: 0..4294967295.
    RemIndex interface{}

    // Chassis ID sub type. The type is interface{} with range: 0..255.
    ChassisIdSubType interface{}

    // Chassis ID length. The type is interface{} with range: 0..65535.
    ChassisIdLen interface{}

    // Port ID sub type. The type is interface{} with range: 0..255.
    PortIdSubType interface{}

    // Port ID length. The type is interface{} with range: 0..65535.
    PortIdLen interface{}

    // Supported and combined cpabilities. The type is interface{} with range:
    // 0..4294967295.
    CombinedCapabilities interface{}

    // Unknown TLV list.
    UnknownTlvList Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList

    // Org Def TLV list.
    OrgDefTlvList Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList
}

func (mib *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib) GetEntityData() *types.CommonEntityData {
    mib.EntityData.YFilter = mib.YFilter
    mib.EntityData.YangName = "mib"
    mib.EntityData.BundleName = "cisco_ios_xr"
    mib.EntityData.ParentYangName = "lldp-neighbor"
    mib.EntityData.SegmentPath = "mib"
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = types.NewOrderedMap()
    mib.EntityData.Children.Append("unknown-tlv-list", types.YChild{"UnknownTlvList", &mib.UnknownTlvList})
    mib.EntityData.Children.Append("org-def-tlv-list", types.YChild{"OrgDefTlvList", &mib.OrgDefTlvList})
    mib.EntityData.Leafs = types.NewOrderedMap()
    mib.EntityData.Leafs.Append("rem-time-mark", types.YLeaf{"RemTimeMark", mib.RemTimeMark})
    mib.EntityData.Leafs.Append("rem-local-port-num", types.YLeaf{"RemLocalPortNum", mib.RemLocalPortNum})
    mib.EntityData.Leafs.Append("rem-index", types.YLeaf{"RemIndex", mib.RemIndex})
    mib.EntityData.Leafs.Append("chassis-id-sub-type", types.YLeaf{"ChassisIdSubType", mib.ChassisIdSubType})
    mib.EntityData.Leafs.Append("chassis-id-len", types.YLeaf{"ChassisIdLen", mib.ChassisIdLen})
    mib.EntityData.Leafs.Append("port-id-sub-type", types.YLeaf{"PortIdSubType", mib.PortIdSubType})
    mib.EntityData.Leafs.Append("port-id-len", types.YLeaf{"PortIdLen", mib.PortIdLen})
    mib.EntityData.Leafs.Append("combined-capabilities", types.YLeaf{"CombinedCapabilities", mib.CombinedCapabilities})

    mib.EntityData.YListKeys = []string {}

    return &(mib.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList
// Unknown TLV list
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp unknown tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry.
    LldpUnknownTlvEntry []*Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
}

func (unknownTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList) GetEntityData() *types.CommonEntityData {
    unknownTlvList.EntityData.YFilter = unknownTlvList.YFilter
    unknownTlvList.EntityData.YangName = "unknown-tlv-list"
    unknownTlvList.EntityData.BundleName = "cisco_ios_xr"
    unknownTlvList.EntityData.ParentYangName = "mib"
    unknownTlvList.EntityData.SegmentPath = "unknown-tlv-list"
    unknownTlvList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unknownTlvList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unknownTlvList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unknownTlvList.EntityData.Children = types.NewOrderedMap()
    unknownTlvList.EntityData.Children.Append("lldp-unknown-tlv-entry", types.YChild{"LldpUnknownTlvEntry", nil})
    for i := range unknownTlvList.LldpUnknownTlvEntry {
        unknownTlvList.EntityData.Children.Append(types.GetSegmentPath(unknownTlvList.LldpUnknownTlvEntry[i]), types.YChild{"LldpUnknownTlvEntry", unknownTlvList.LldpUnknownTlvEntry[i]})
    }
    unknownTlvList.EntityData.Leafs = types.NewOrderedMap()

    unknownTlvList.EntityData.YListKeys = []string {}

    return &(unknownTlvList.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry
// lldp unknown tlv entry
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unknown TLV type. The type is interface{} with range: 0..255.
    TlvType interface{}

    // Unknown TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpUnknownTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_UnknownTlvList_LldpUnknownTlvEntry) GetEntityData() *types.CommonEntityData {
    lldpUnknownTlvEntry.EntityData.YFilter = lldpUnknownTlvEntry.YFilter
    lldpUnknownTlvEntry.EntityData.YangName = "lldp-unknown-tlv-entry"
    lldpUnknownTlvEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpUnknownTlvEntry.EntityData.ParentYangName = "unknown-tlv-list"
    lldpUnknownTlvEntry.EntityData.SegmentPath = "lldp-unknown-tlv-entry"
    lldpUnknownTlvEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpUnknownTlvEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpUnknownTlvEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpUnknownTlvEntry.EntityData.Children = types.NewOrderedMap()
    lldpUnknownTlvEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpUnknownTlvEntry.EntityData.Leafs.Append("tlv-type", types.YLeaf{"TlvType", lldpUnknownTlvEntry.TlvType})
    lldpUnknownTlvEntry.EntityData.Leafs.Append("tlv-value", types.YLeaf{"TlvValue", lldpUnknownTlvEntry.TlvValue})

    lldpUnknownTlvEntry.EntityData.YListKeys = []string {}

    return &(lldpUnknownTlvEntry.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList
// Org Def TLV list
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp org def tlv entry. The type is slice of
    // Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry.
    LldpOrgDefTlvEntry []*Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
}

func (orgDefTlvList *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList) GetEntityData() *types.CommonEntityData {
    orgDefTlvList.EntityData.YFilter = orgDefTlvList.YFilter
    orgDefTlvList.EntityData.YangName = "org-def-tlv-list"
    orgDefTlvList.EntityData.BundleName = "cisco_ios_xr"
    orgDefTlvList.EntityData.ParentYangName = "mib"
    orgDefTlvList.EntityData.SegmentPath = "org-def-tlv-list"
    orgDefTlvList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    orgDefTlvList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    orgDefTlvList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    orgDefTlvList.EntityData.Children = types.NewOrderedMap()
    orgDefTlvList.EntityData.Children.Append("lldp-org-def-tlv-entry", types.YChild{"LldpOrgDefTlvEntry", nil})
    for i := range orgDefTlvList.LldpOrgDefTlvEntry {
        orgDefTlvList.EntityData.Children.Append(types.GetSegmentPath(orgDefTlvList.LldpOrgDefTlvEntry[i]), types.YChild{"LldpOrgDefTlvEntry", orgDefTlvList.LldpOrgDefTlvEntry[i]})
    }
    orgDefTlvList.EntityData.Leafs = types.NewOrderedMap()

    orgDefTlvList.EntityData.YListKeys = []string {}

    return &(orgDefTlvList.EntityData)
}

// Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry
// lldp org def tlv entry
type Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Organizationally Unique Identifier. The type is interface{} with range:
    // 0..4294967295.
    Oui interface{}

    // Org Def TLV subtype. The type is interface{} with range: 0..255.
    TlvSubtype interface{}

    // lldpRemOrgDefInfoIndex. The type is interface{} with range: 0..4294967295.
    TlvInfoIndes interface{}

    // Org Def TLV payload. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    TlvValue interface{}
}

func (lldpOrgDefTlvEntry *Lldp_Nodes_Node_Neighbors_Summaries_Summary_LldpNeighbor_Mib_OrgDefTlvList_LldpOrgDefTlvEntry) GetEntityData() *types.CommonEntityData {
    lldpOrgDefTlvEntry.EntityData.YFilter = lldpOrgDefTlvEntry.YFilter
    lldpOrgDefTlvEntry.EntityData.YangName = "lldp-org-def-tlv-entry"
    lldpOrgDefTlvEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpOrgDefTlvEntry.EntityData.ParentYangName = "org-def-tlv-list"
    lldpOrgDefTlvEntry.EntityData.SegmentPath = "lldp-org-def-tlv-entry"
    lldpOrgDefTlvEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpOrgDefTlvEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpOrgDefTlvEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpOrgDefTlvEntry.EntityData.Children = types.NewOrderedMap()
    lldpOrgDefTlvEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("oui", types.YLeaf{"Oui", lldpOrgDefTlvEntry.Oui})
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("tlv-subtype", types.YLeaf{"TlvSubtype", lldpOrgDefTlvEntry.TlvSubtype})
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("tlv-info-indes", types.YLeaf{"TlvInfoIndes", lldpOrgDefTlvEntry.TlvInfoIndes})
    lldpOrgDefTlvEntry.EntityData.Leafs.Append("tlv-value", types.YLeaf{"TlvValue", lldpOrgDefTlvEntry.TlvValue})

    lldpOrgDefTlvEntry.EntityData.YListKeys = []string {}

    return &(lldpOrgDefTlvEntry.EntityData)
}

// Lldp_Nodes_Node_Interfaces
// The table of interfaces on which LLDP is
// running on this node
type Lldp_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for an interface on which LLDP is running. The type is
    // slice of Lldp_Nodes_Node_Interfaces_Interface.
    Interface []*Lldp_Nodes_Node_Interfaces_Interface
}

func (interfaces *Lldp_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Lldp_Nodes_Node_Interfaces_Interface
// Operational data for an interface on which
// LLDP is running
type Lldp_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceNameXr interface{}

    // TX Enabled. The type is interface{} with range: 0..255.
    TxEnabled interface{}

    // RX Enabled. The type is interface{} with range: 0..255.
    RxEnabled interface{}

    // TX State. The type is string.
    TxState interface{}

    // RX State. The type is string.
    RxState interface{}

    // ifIndex. The type is interface{} with range: 0..4294967295.
    IfIndex interface{}

    // Outgoing port identifier. The type is string.
    PortId interface{}

    // Port ID sub type. The type is interface{} with range: 0..255.
    PortIdSubType interface{}

    // Port Description. The type is string.
    PortDescription interface{}

    // Local Management Addresses.
    LocalNetworkAddresses Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses
}

func (self *Lldp_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("local-network-addresses", types.YChild{"LocalNetworkAddresses", &self.LocalNetworkAddresses})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("interface-name-xr", types.YLeaf{"InterfaceNameXr", self.InterfaceNameXr})
    self.EntityData.Leafs.Append("tx-enabled", types.YLeaf{"TxEnabled", self.TxEnabled})
    self.EntityData.Leafs.Append("rx-enabled", types.YLeaf{"RxEnabled", self.RxEnabled})
    self.EntityData.Leafs.Append("tx-state", types.YLeaf{"TxState", self.TxState})
    self.EntityData.Leafs.Append("rx-state", types.YLeaf{"RxState", self.RxState})
    self.EntityData.Leafs.Append("if-index", types.YLeaf{"IfIndex", self.IfIndex})
    self.EntityData.Leafs.Append("port-id", types.YLeaf{"PortId", self.PortId})
    self.EntityData.Leafs.Append("port-id-sub-type", types.YLeaf{"PortIdSubType", self.PortIdSubType})
    self.EntityData.Leafs.Append("port-description", types.YLeaf{"PortDescription", self.PortDescription})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses
// Local Management Addresses
type Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // lldp addr entry. The type is slice of
    // Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry.
    LldpAddrEntry []*Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry
}

func (localNetworkAddresses *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses) GetEntityData() *types.CommonEntityData {
    localNetworkAddresses.EntityData.YFilter = localNetworkAddresses.YFilter
    localNetworkAddresses.EntityData.YangName = "local-network-addresses"
    localNetworkAddresses.EntityData.BundleName = "cisco_ios_xr"
    localNetworkAddresses.EntityData.ParentYangName = "interface"
    localNetworkAddresses.EntityData.SegmentPath = "local-network-addresses"
    localNetworkAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localNetworkAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localNetworkAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localNetworkAddresses.EntityData.Children = types.NewOrderedMap()
    localNetworkAddresses.EntityData.Children.Append("lldp-addr-entry", types.YChild{"LldpAddrEntry", nil})
    for i := range localNetworkAddresses.LldpAddrEntry {
        localNetworkAddresses.EntityData.Children.Append(types.GetSegmentPath(localNetworkAddresses.LldpAddrEntry[i]), types.YChild{"LldpAddrEntry", localNetworkAddresses.LldpAddrEntry[i]})
    }
    localNetworkAddresses.EntityData.Leafs = types.NewOrderedMap()

    localNetworkAddresses.EntityData.YListKeys = []string {}

    return &(localNetworkAddresses.EntityData)
}

// Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry
// lldp addr entry
type Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MA sub type. The type is interface{} with range: 0..255.
    MaSubtype interface{}

    // Interface num. The type is interface{} with range: 0..4294967295.
    IfNum interface{}

    // Network layer address.
    Address Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address
}

func (lldpAddrEntry *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry) GetEntityData() *types.CommonEntityData {
    lldpAddrEntry.EntityData.YFilter = lldpAddrEntry.YFilter
    lldpAddrEntry.EntityData.YangName = "lldp-addr-entry"
    lldpAddrEntry.EntityData.BundleName = "cisco_ios_xr"
    lldpAddrEntry.EntityData.ParentYangName = "local-network-addresses"
    lldpAddrEntry.EntityData.SegmentPath = "lldp-addr-entry"
    lldpAddrEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldpAddrEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldpAddrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldpAddrEntry.EntityData.Children = types.NewOrderedMap()
    lldpAddrEntry.EntityData.Children.Append("address", types.YChild{"Address", &lldpAddrEntry.Address})
    lldpAddrEntry.EntityData.Leafs = types.NewOrderedMap()
    lldpAddrEntry.EntityData.Leafs.Append("ma-subtype", types.YLeaf{"MaSubtype", lldpAddrEntry.MaSubtype})
    lldpAddrEntry.EntityData.Leafs.Append("if-num", types.YLeaf{"IfNum", lldpAddrEntry.IfNum})

    lldpAddrEntry.EntityData.YListKeys = []string {}

    return &(lldpAddrEntry.EntityData)
}

// Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address
// Network layer address
type Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is LldpL3AddrProtocol.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (address *Lldp_Nodes_Node_Interfaces_Interface_LocalNetworkAddresses_LldpAddrEntry_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "lldp-addr-entry"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lldp_Nodes_Node_Statistics
// The LLDP traffic statistics for this node
type Lldp_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transmitted packets. The type is interface{} with range: 0..4294967295.
    TransmittedPackets interface{}

    // Aged out entries. The type is interface{} with range: 0..4294967295.
    AgedOutEntries interface{}

    // Discarded packets. The type is interface{} with range: 0..4294967295.
    DiscardedPackets interface{}

    // Bad packet received and dropped. The type is interface{} with range:
    // 0..4294967295.
    BadPackets interface{}

    // Received packets. The type is interface{} with range: 0..4294967295.
    ReceivedPackets interface{}

    // Discarded TLVs. The type is interface{} with range: 0..4294967295.
    DiscardedTlVs interface{}

    // Unrecognized TLVs. The type is interface{} with range: 0..4294967295.
    UnrecognizedTlVs interface{}

    // Out-of-memory conditions. The type is interface{} with range:
    // 0..4294967295.
    OutOfMemoryErrors interface{}

    // Transmission errors. The type is interface{} with range: 0..4294967295.
    EncapsulationErrors interface{}

    // Queue overflows. The type is interface{} with range: 0..4294967295.
    QueueOverflowErrors interface{}

    // Table overflows. The type is interface{} with range: 0..4294967295.
    TableOverflowErrors interface{}
}

func (statistics *Lldp_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("transmitted-packets", types.YLeaf{"TransmittedPackets", statistics.TransmittedPackets})
    statistics.EntityData.Leafs.Append("aged-out-entries", types.YLeaf{"AgedOutEntries", statistics.AgedOutEntries})
    statistics.EntityData.Leafs.Append("discarded-packets", types.YLeaf{"DiscardedPackets", statistics.DiscardedPackets})
    statistics.EntityData.Leafs.Append("bad-packets", types.YLeaf{"BadPackets", statistics.BadPackets})
    statistics.EntityData.Leafs.Append("received-packets", types.YLeaf{"ReceivedPackets", statistics.ReceivedPackets})
    statistics.EntityData.Leafs.Append("discarded-tl-vs", types.YLeaf{"DiscardedTlVs", statistics.DiscardedTlVs})
    statistics.EntityData.Leafs.Append("unrecognized-tl-vs", types.YLeaf{"UnrecognizedTlVs", statistics.UnrecognizedTlVs})
    statistics.EntityData.Leafs.Append("out-of-memory-errors", types.YLeaf{"OutOfMemoryErrors", statistics.OutOfMemoryErrors})
    statistics.EntityData.Leafs.Append("encapsulation-errors", types.YLeaf{"EncapsulationErrors", statistics.EncapsulationErrors})
    statistics.EntityData.Leafs.Append("queue-overflow-errors", types.YLeaf{"QueueOverflowErrors", statistics.QueueOverflowErrors})
    statistics.EntityData.Leafs.Append("table-overflow-errors", types.YLeaf{"TableOverflowErrors", statistics.TableOverflowErrors})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

