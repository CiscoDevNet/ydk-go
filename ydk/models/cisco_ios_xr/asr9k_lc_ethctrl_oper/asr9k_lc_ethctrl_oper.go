// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-lc-ethctrl package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mlan: Management LAN Operational data space
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_lc_ethctrl_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_lc_ethctrl_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-lc-ethctrl-oper mlan}", reflect.TypeOf(Mlan{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-lc-ethctrl-oper:mlan", reflect.TypeOf(Mlan{}))
}

// Mlan
// Management LAN Operational data space
type Mlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes Mlan_Nodes
}

func (mlan *Mlan) GetEntityData() *types.CommonEntityData {
    mlan.EntityData.YFilter = mlan.YFilter
    mlan.EntityData.YangName = "mlan"
    mlan.EntityData.BundleName = "cisco_ios_xr"
    mlan.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-lc-ethctrl-oper"
    mlan.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-lc-ethctrl-oper:mlan"
    mlan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mlan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mlan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mlan.EntityData.Children = types.NewOrderedMap()
    mlan.EntityData.Children.Append("nodes", types.YChild{"Nodes", &mlan.Nodes})
    mlan.EntityData.Leafs = types.NewOrderedMap()

    mlan.EntityData.YListKeys = []string {}

    return &(mlan.EntityData)
}

// Mlan_Nodes
// Table of nodes
type Mlan_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of Mlan_Nodes_Node.
    Node []*Mlan_Nodes_Node
}

func (nodes *Mlan_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "mlan"
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

// Mlan_Nodes_Node
// Number
type Mlan_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. node number. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Table of port status.
    PortStatusNumbers Mlan_Nodes_Node_PortStatusNumbers

    // Table of switch status.
    SwitchStatusTable Mlan_Nodes_Node_SwitchStatusTable

    // Table of port counters.
    PortCountersNumbers Mlan_Nodes_Node_PortCountersNumbers

    // Table of switch ATU.
    AtuEntryNumbers Mlan_Nodes_Node_AtuEntryNumbers
}

func (node *Mlan_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Node, "node")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("port-status-numbers", types.YChild{"PortStatusNumbers", &node.PortStatusNumbers})
    node.EntityData.Children.Append("switch-status-table", types.YChild{"SwitchStatusTable", &node.SwitchStatusTable})
    node.EntityData.Children.Append("port-counters-numbers", types.YChild{"PortCountersNumbers", &node.PortCountersNumbers})
    node.EntityData.Children.Append("atu-entry-numbers", types.YChild{"AtuEntryNumbers", &node.AtuEntryNumbers})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node", types.YLeaf{"Node", node.Node})

    node.EntityData.YListKeys = []string {"Node"}

    return &(node.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers
// Table of port status
type Mlan_Nodes_Node_PortStatusNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber.
    PortStatusNumber []*Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber
}

func (portStatusNumbers *Mlan_Nodes_Node_PortStatusNumbers) GetEntityData() *types.CommonEntityData {
    portStatusNumbers.EntityData.YFilter = portStatusNumbers.YFilter
    portStatusNumbers.EntityData.YangName = "port-status-numbers"
    portStatusNumbers.EntityData.BundleName = "cisco_ios_xr"
    portStatusNumbers.EntityData.ParentYangName = "node"
    portStatusNumbers.EntityData.SegmentPath = "port-status-numbers"
    portStatusNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portStatusNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portStatusNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portStatusNumbers.EntityData.Children = types.NewOrderedMap()
    portStatusNumbers.EntityData.Children.Append("port-status-number", types.YChild{"PortStatusNumber", nil})
    for i := range portStatusNumbers.PortStatusNumber {
        portStatusNumbers.EntityData.Children.Append(types.GetSegmentPath(portStatusNumbers.PortStatusNumber[i]), types.YChild{"PortStatusNumber", portStatusNumbers.PortStatusNumber[i]})
    }
    portStatusNumbers.EntityData.Leafs = types.NewOrderedMap()

    portStatusNumbers.EntityData.YListKeys = []string {}

    return &(portStatusNumbers.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber
// Number
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. port number. The type is interface{} with range:
    // 0..4294967295.
    Number interface{}

    // mlan port status info.
    PortStatus Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus
}

func (portStatusNumber *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber) GetEntityData() *types.CommonEntityData {
    portStatusNumber.EntityData.YFilter = portStatusNumber.YFilter
    portStatusNumber.EntityData.YangName = "port-status-number"
    portStatusNumber.EntityData.BundleName = "cisco_ios_xr"
    portStatusNumber.EntityData.ParentYangName = "port-status-numbers"
    portStatusNumber.EntityData.SegmentPath = "port-status-number" + types.AddKeyToken(portStatusNumber.Number, "number")
    portStatusNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portStatusNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portStatusNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portStatusNumber.EntityData.Children = types.NewOrderedMap()
    portStatusNumber.EntityData.Children.Append("port-status", types.YChild{"PortStatus", &portStatusNumber.PortStatus})
    portStatusNumber.EntityData.Leafs = types.NewOrderedMap()
    portStatusNumber.EntityData.Leafs.Append("number", types.YLeaf{"Number", portStatusNumber.Number})

    portStatusNumber.EntityData.YListKeys = []string {"Number"}

    return &(portStatusNumber.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus
// mlan port status info
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port Number. The type is interface{} with range: 0..4294967295.
    PortNum interface{}

    // PHY data valid. The type is interface{} with range: 0..4294967295.
    PhyValid interface{}

    // SERDES data valid. The type is interface{} with range: 0..4294967295.
    SerdesValid interface{}

    // MAC data valid. The type is interface{} with range: 0..4294967295.
    MacValid interface{}

    // Configuration Data.
    Config Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config

    // PHY Registers.
    Phy Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy

    // SERDES Registers.
    Serdes Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes

    // MAC Registers.
    Mac Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac
}

func (portStatus *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus) GetEntityData() *types.CommonEntityData {
    portStatus.EntityData.YFilter = portStatus.YFilter
    portStatus.EntityData.YangName = "port-status"
    portStatus.EntityData.BundleName = "cisco_ios_xr"
    portStatus.EntityData.ParentYangName = "port-status-number"
    portStatus.EntityData.SegmentPath = "port-status"
    portStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portStatus.EntityData.Children = types.NewOrderedMap()
    portStatus.EntityData.Children.Append("config", types.YChild{"Config", &portStatus.Config})
    portStatus.EntityData.Children.Append("phy", types.YChild{"Phy", &portStatus.Phy})
    portStatus.EntityData.Children.Append("serdes", types.YChild{"Serdes", &portStatus.Serdes})
    portStatus.EntityData.Children.Append("mac", types.YChild{"Mac", &portStatus.Mac})
    portStatus.EntityData.Leafs = types.NewOrderedMap()
    portStatus.EntityData.Leafs.Append("port-num", types.YLeaf{"PortNum", portStatus.PortNum})
    portStatus.EntityData.Leafs.Append("phy-valid", types.YLeaf{"PhyValid", portStatus.PhyValid})
    portStatus.EntityData.Leafs.Append("serdes-valid", types.YLeaf{"SerdesValid", portStatus.SerdesValid})
    portStatus.EntityData.Leafs.Append("mac-valid", types.YLeaf{"MacValid", portStatus.MacValid})

    portStatus.EntityData.YListKeys = []string {}

    return &(portStatus.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config
// Configuration Data
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // speed. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // duplex. The type is interface{} with range: 0..4294967295.
    Duplex interface{}

    // pauseEn. The type is interface{} with range: 0..65535.
    Pause interface{}

    // myPause. The type is interface{} with range: 0..65535.
    MyPause interface{}

    // loopback. The type is interface{} with range: 0..4294967295.
    Loopback interface{}
}

func (config *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xr"
    config.EntityData.ParentYangName = "port-status"
    config.EntityData.SegmentPath = "config"
    config.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", config.Speed})
    config.EntityData.Leafs.Append("duplex", types.YLeaf{"Duplex", config.Duplex})
    config.EntityData.Leafs.Append("pause", types.YLeaf{"Pause", config.Pause})
    config.EntityData.Leafs.Append("my-pause", types.YLeaf{"MyPause", config.MyPause})
    config.EntityData.Leafs.Append("loopback", types.YLeaf{"Loopback", config.Loopback})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy
// PHY Registers
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // reg. The type is slice of
    // Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy_Reg.
    Reg []*Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy_Reg
}

func (phy *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy) GetEntityData() *types.CommonEntityData {
    phy.EntityData.YFilter = phy.YFilter
    phy.EntityData.YangName = "phy"
    phy.EntityData.BundleName = "cisco_ios_xr"
    phy.EntityData.ParentYangName = "port-status"
    phy.EntityData.SegmentPath = "phy"
    phy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    phy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    phy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    phy.EntityData.Children = types.NewOrderedMap()
    phy.EntityData.Children.Append("reg", types.YChild{"Reg", nil})
    for i := range phy.Reg {
        phy.EntityData.Children.Append(types.GetSegmentPath(phy.Reg[i]), types.YChild{"Reg", phy.Reg[i]})
    }
    phy.EntityData.Leafs = types.NewOrderedMap()

    phy.EntityData.YListKeys = []string {}

    return &(phy.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy_Reg
// reg
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy_Reg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535.
    Entry interface{}
}

func (reg *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Phy_Reg) GetEntityData() *types.CommonEntityData {
    reg.EntityData.YFilter = reg.YFilter
    reg.EntityData.YangName = "reg"
    reg.EntityData.BundleName = "cisco_ios_xr"
    reg.EntityData.ParentYangName = "phy"
    reg.EntityData.SegmentPath = "reg"
    reg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reg.EntityData.Children = types.NewOrderedMap()
    reg.EntityData.Leafs = types.NewOrderedMap()
    reg.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", reg.Entry})

    reg.EntityData.YListKeys = []string {}

    return &(reg.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes
// SERDES Registers
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // reg. The type is slice of
    // Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes_Reg.
    Reg []*Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes_Reg
}

func (serdes *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes) GetEntityData() *types.CommonEntityData {
    serdes.EntityData.YFilter = serdes.YFilter
    serdes.EntityData.YangName = "serdes"
    serdes.EntityData.BundleName = "cisco_ios_xr"
    serdes.EntityData.ParentYangName = "port-status"
    serdes.EntityData.SegmentPath = "serdes"
    serdes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serdes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serdes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serdes.EntityData.Children = types.NewOrderedMap()
    serdes.EntityData.Children.Append("reg", types.YChild{"Reg", nil})
    for i := range serdes.Reg {
        serdes.EntityData.Children.Append(types.GetSegmentPath(serdes.Reg[i]), types.YChild{"Reg", serdes.Reg[i]})
    }
    serdes.EntityData.Leafs = types.NewOrderedMap()

    serdes.EntityData.YListKeys = []string {}

    return &(serdes.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes_Reg
// reg
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes_Reg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535.
    Entry interface{}
}

func (reg *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Serdes_Reg) GetEntityData() *types.CommonEntityData {
    reg.EntityData.YFilter = reg.YFilter
    reg.EntityData.YangName = "reg"
    reg.EntityData.BundleName = "cisco_ios_xr"
    reg.EntityData.ParentYangName = "serdes"
    reg.EntityData.SegmentPath = "reg"
    reg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reg.EntityData.Children = types.NewOrderedMap()
    reg.EntityData.Leafs = types.NewOrderedMap()
    reg.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", reg.Entry})

    reg.EntityData.YListKeys = []string {}

    return &(reg.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac
// MAC Registers
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // reg. The type is slice of
    // Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac_Reg.
    Reg []*Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac_Reg
}

func (mac *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "port-status"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("reg", types.YChild{"Reg", nil})
    for i := range mac.Reg {
        mac.EntityData.Children.Append(types.GetSegmentPath(mac.Reg[i]), types.YChild{"Reg", mac.Reg[i]})
    }
    mac.EntityData.Leafs = types.NewOrderedMap()

    mac.EntityData.YListKeys = []string {}

    return &(mac.EntityData)
}

// Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac_Reg
// reg
type Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac_Reg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535.
    Entry interface{}
}

func (reg *Mlan_Nodes_Node_PortStatusNumbers_PortStatusNumber_PortStatus_Mac_Reg) GetEntityData() *types.CommonEntityData {
    reg.EntityData.YFilter = reg.YFilter
    reg.EntityData.YangName = "reg"
    reg.EntityData.BundleName = "cisco_ios_xr"
    reg.EntityData.ParentYangName = "mac"
    reg.EntityData.SegmentPath = "reg"
    reg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reg.EntityData.Children = types.NewOrderedMap()
    reg.EntityData.Leafs = types.NewOrderedMap()
    reg.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", reg.Entry})

    reg.EntityData.YListKeys = []string {}

    return &(reg.EntityData)
}

// Mlan_Nodes_Node_SwitchStatusTable
// Table of switch status
type Mlan_Nodes_Node_SwitchStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // mlan switch status info.
    SwitchStatus Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus
}

func (switchStatusTable *Mlan_Nodes_Node_SwitchStatusTable) GetEntityData() *types.CommonEntityData {
    switchStatusTable.EntityData.YFilter = switchStatusTable.YFilter
    switchStatusTable.EntityData.YangName = "switch-status-table"
    switchStatusTable.EntityData.BundleName = "cisco_ios_xr"
    switchStatusTable.EntityData.ParentYangName = "node"
    switchStatusTable.EntityData.SegmentPath = "switch-status-table"
    switchStatusTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchStatusTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchStatusTable.EntityData.Children = types.NewOrderedMap()
    switchStatusTable.EntityData.Children.Append("switch-status", types.YChild{"SwitchStatus", &switchStatusTable.SwitchStatus})
    switchStatusTable.EntityData.Leafs = types.NewOrderedMap()

    switchStatusTable.EntityData.YListKeys = []string {}

    return &(switchStatusTable.EntityData)
}

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus
// mlan switch status info
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CPU Interface Rate Limit. The type is interface{} with range:
    // -2147483648..2147483647.
    RateLimit interface{}

    // Switch Global Registers.
    SwReg1 Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1

    // Switch Global Registers.
    SwReg2 Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2

    // Switch Status Data.
    SwStatus Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus
}

func (switchStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus) GetEntityData() *types.CommonEntityData {
    switchStatus.EntityData.YFilter = switchStatus.YFilter
    switchStatus.EntityData.YangName = "switch-status"
    switchStatus.EntityData.BundleName = "cisco_ios_xr"
    switchStatus.EntityData.ParentYangName = "switch-status-table"
    switchStatus.EntityData.SegmentPath = "switch-status"
    switchStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchStatus.EntityData.Children = types.NewOrderedMap()
    switchStatus.EntityData.Children.Append("sw-reg-1", types.YChild{"SwReg1", &switchStatus.SwReg1})
    switchStatus.EntityData.Children.Append("sw-reg-2", types.YChild{"SwReg2", &switchStatus.SwReg2})
    switchStatus.EntityData.Children.Append("sw-status", types.YChild{"SwStatus", &switchStatus.SwStatus})
    switchStatus.EntityData.Leafs = types.NewOrderedMap()
    switchStatus.EntityData.Leafs.Append("rate-limit", types.YLeaf{"RateLimit", switchStatus.RateLimit})

    switchStatus.EntityData.YListKeys = []string {}

    return &(switchStatus.EntityData)
}

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1
// Switch Global Registers
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // reg. The type is slice of
    // Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1_Reg.
    Reg []*Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1_Reg
}

func (swReg1 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1) GetEntityData() *types.CommonEntityData {
    swReg1.EntityData.YFilter = swReg1.YFilter
    swReg1.EntityData.YangName = "sw-reg-1"
    swReg1.EntityData.BundleName = "cisco_ios_xr"
    swReg1.EntityData.ParentYangName = "switch-status"
    swReg1.EntityData.SegmentPath = "sw-reg-1"
    swReg1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    swReg1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    swReg1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    swReg1.EntityData.Children = types.NewOrderedMap()
    swReg1.EntityData.Children.Append("reg", types.YChild{"Reg", nil})
    for i := range swReg1.Reg {
        swReg1.EntityData.Children.Append(types.GetSegmentPath(swReg1.Reg[i]), types.YChild{"Reg", swReg1.Reg[i]})
    }
    swReg1.EntityData.Leafs = types.NewOrderedMap()

    swReg1.EntityData.YListKeys = []string {}

    return &(swReg1.EntityData)
}

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1_Reg
// reg
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1_Reg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535.
    Entry interface{}
}

func (reg *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg1_Reg) GetEntityData() *types.CommonEntityData {
    reg.EntityData.YFilter = reg.YFilter
    reg.EntityData.YangName = "reg"
    reg.EntityData.BundleName = "cisco_ios_xr"
    reg.EntityData.ParentYangName = "sw-reg-1"
    reg.EntityData.SegmentPath = "reg"
    reg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reg.EntityData.Children = types.NewOrderedMap()
    reg.EntityData.Leafs = types.NewOrderedMap()
    reg.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", reg.Entry})

    reg.EntityData.YListKeys = []string {}

    return &(reg.EntityData)
}

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2
// Switch Global Registers
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // reg. The type is slice of
    // Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2_Reg.
    Reg []*Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2_Reg
}

func (swReg2 *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2) GetEntityData() *types.CommonEntityData {
    swReg2.EntityData.YFilter = swReg2.YFilter
    swReg2.EntityData.YangName = "sw-reg-2"
    swReg2.EntityData.BundleName = "cisco_ios_xr"
    swReg2.EntityData.ParentYangName = "switch-status"
    swReg2.EntityData.SegmentPath = "sw-reg-2"
    swReg2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    swReg2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    swReg2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    swReg2.EntityData.Children = types.NewOrderedMap()
    swReg2.EntityData.Children.Append("reg", types.YChild{"Reg", nil})
    for i := range swReg2.Reg {
        swReg2.EntityData.Children.Append(types.GetSegmentPath(swReg2.Reg[i]), types.YChild{"Reg", swReg2.Reg[i]})
    }
    swReg2.EntityData.Leafs = types.NewOrderedMap()

    swReg2.EntityData.YListKeys = []string {}

    return &(swReg2.EntityData)
}

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2_Reg
// reg
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2_Reg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535.
    Entry interface{}
}

func (reg *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwReg2_Reg) GetEntityData() *types.CommonEntityData {
    reg.EntityData.YFilter = reg.YFilter
    reg.EntityData.YangName = "reg"
    reg.EntityData.BundleName = "cisco_ios_xr"
    reg.EntityData.ParentYangName = "sw-reg-2"
    reg.EntityData.SegmentPath = "reg"
    reg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reg.EntityData.Children = types.NewOrderedMap()
    reg.EntityData.Leafs = types.NewOrderedMap()
    reg.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", reg.Entry})

    reg.EntityData.YListKeys = []string {}

    return &(reg.EntityData)
}

// Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus
// Switch Status Data
type Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ppu. The type is interface{} with range: 0..4294967295.
    Ppu interface{}

    // mtu. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // mac. The type is string with length: 0..6.
    Mac interface{}

    // cpu port. The type is interface{} with range: 0..65535.
    CpuPort interface{}

    // cpu mac. The type is interface{} with range: 0..65535.
    CpuMac interface{}

    // initialized. The type is interface{} with range: 0..65535.
    Initialized interface{}

    // restarted. The type is interface{} with range: 0..65535.
    Restarted interface{}
}

func (swStatus *Mlan_Nodes_Node_SwitchStatusTable_SwitchStatus_SwStatus) GetEntityData() *types.CommonEntityData {
    swStatus.EntityData.YFilter = swStatus.YFilter
    swStatus.EntityData.YangName = "sw-status"
    swStatus.EntityData.BundleName = "cisco_ios_xr"
    swStatus.EntityData.ParentYangName = "switch-status"
    swStatus.EntityData.SegmentPath = "sw-status"
    swStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    swStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    swStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    swStatus.EntityData.Children = types.NewOrderedMap()
    swStatus.EntityData.Leafs = types.NewOrderedMap()
    swStatus.EntityData.Leafs.Append("ppu", types.YLeaf{"Ppu", swStatus.Ppu})
    swStatus.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", swStatus.Mtu})
    swStatus.EntityData.Leafs.Append("mac", types.YLeaf{"Mac", swStatus.Mac})
    swStatus.EntityData.Leafs.Append("cpu-port", types.YLeaf{"CpuPort", swStatus.CpuPort})
    swStatus.EntityData.Leafs.Append("cpu-mac", types.YLeaf{"CpuMac", swStatus.CpuMac})
    swStatus.EntityData.Leafs.Append("initialized", types.YLeaf{"Initialized", swStatus.Initialized})
    swStatus.EntityData.Leafs.Append("restarted", types.YLeaf{"Restarted", swStatus.Restarted})

    swStatus.EntityData.YListKeys = []string {}

    return &(swStatus.EntityData)
}

// Mlan_Nodes_Node_PortCountersNumbers
// Table of port counters
type Mlan_Nodes_Node_PortCountersNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of
    // Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber.
    PortCountersNumber []*Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber
}

func (portCountersNumbers *Mlan_Nodes_Node_PortCountersNumbers) GetEntityData() *types.CommonEntityData {
    portCountersNumbers.EntityData.YFilter = portCountersNumbers.YFilter
    portCountersNumbers.EntityData.YangName = "port-counters-numbers"
    portCountersNumbers.EntityData.BundleName = "cisco_ios_xr"
    portCountersNumbers.EntityData.ParentYangName = "node"
    portCountersNumbers.EntityData.SegmentPath = "port-counters-numbers"
    portCountersNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portCountersNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portCountersNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portCountersNumbers.EntityData.Children = types.NewOrderedMap()
    portCountersNumbers.EntityData.Children.Append("port-counters-number", types.YChild{"PortCountersNumber", nil})
    for i := range portCountersNumbers.PortCountersNumber {
        portCountersNumbers.EntityData.Children.Append(types.GetSegmentPath(portCountersNumbers.PortCountersNumber[i]), types.YChild{"PortCountersNumber", portCountersNumbers.PortCountersNumber[i]})
    }
    portCountersNumbers.EntityData.Leafs = types.NewOrderedMap()

    portCountersNumbers.EntityData.YListKeys = []string {}

    return &(portCountersNumbers.EntityData)
}

// Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber
// Number
type Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. port number. The type is interface{} with range:
    // 0..4294967295.
    Number interface{}

    // mlan port counters info.
    PortCounters Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters
}

func (portCountersNumber *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber) GetEntityData() *types.CommonEntityData {
    portCountersNumber.EntityData.YFilter = portCountersNumber.YFilter
    portCountersNumber.EntityData.YangName = "port-counters-number"
    portCountersNumber.EntityData.BundleName = "cisco_ios_xr"
    portCountersNumber.EntityData.ParentYangName = "port-counters-numbers"
    portCountersNumber.EntityData.SegmentPath = "port-counters-number" + types.AddKeyToken(portCountersNumber.Number, "number")
    portCountersNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portCountersNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portCountersNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portCountersNumber.EntityData.Children = types.NewOrderedMap()
    portCountersNumber.EntityData.Children.Append("port-counters", types.YChild{"PortCounters", &portCountersNumber.PortCounters})
    portCountersNumber.EntityData.Leafs = types.NewOrderedMap()
    portCountersNumber.EntityData.Leafs.Append("number", types.YLeaf{"Number", portCountersNumber.Number})

    portCountersNumber.EntityData.YListKeys = []string {"Number"}

    return &(portCountersNumber.EntityData)
}

// Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters
// mlan port counters info
type Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port Number. The type is interface{} with range: 0..4294967295.
    PortNum interface{}

    // Switch Port Statistics.
    MlanStats Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats
}

func (portCounters *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters) GetEntityData() *types.CommonEntityData {
    portCounters.EntityData.YFilter = portCounters.YFilter
    portCounters.EntityData.YangName = "port-counters"
    portCounters.EntityData.BundleName = "cisco_ios_xr"
    portCounters.EntityData.ParentYangName = "port-counters-number"
    portCounters.EntityData.SegmentPath = "port-counters"
    portCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portCounters.EntityData.Children = types.NewOrderedMap()
    portCounters.EntityData.Children.Append("mlan-stats", types.YChild{"MlanStats", &portCounters.MlanStats})
    portCounters.EntityData.Leafs = types.NewOrderedMap()
    portCounters.EntityData.Leafs.Append("port-num", types.YLeaf{"PortNum", portCounters.PortNum})

    portCounters.EntityData.YListKeys = []string {}

    return &(portCounters.EntityData)
}

// Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats
// Switch Port Statistics
type Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // inGoodOctets hi. The type is interface{} with range: 0..4294967295.
    InGoodOctetsHi interface{}

    // inGoodOctets. The type is interface{} with range: 0..4294967295.
    InGoodOctets interface{}

    // inBadOctets. The type is interface{} with range: 0..4294967295.
    InBadOctets interface{}

    // inUnicastPkt. The type is interface{} with range: 0..4294967295.
    InUnicastPkt interface{}

    // inBcastPkt. The type is interface{} with range: 0..4294967295.
    InBcastPkt interface{}

    // inMcastPkt. The type is interface{} with range: 0..4294967295.
    InMcastPkt interface{}

    // inPausePkt. The type is interface{} with range: 0..4294967295.
    InPausePkt interface{}

    // inUndersizePkt. The type is interface{} with range: 0..4294967295.
    InUndersizePkt interface{}

    // inFragments. The type is interface{} with range: 0..4294967295.
    InFragments interface{}

    // inOversize. The type is interface{} with range: 0..4294967295.
    InOversize interface{}

    // inJabber. The type is interface{} with range: 0..4294967295.
    InJabber interface{}

    // inRxErr. The type is interface{} with range: 0..4294967295.
    InRxErr interface{}

    // inFcsErr. The type is interface{} with range: 0..4294967295.
    InFcsErr interface{}

    // outOctets hi. The type is interface{} with range: 0..4294967295.
    OutOctetsHi interface{}

    // outOctets. The type is interface{} with range: 0..4294967295.
    OutOctets interface{}

    // outUnicastPkt. The type is interface{} with range: 0..4294967295.
    OutUnicastPkt interface{}

    // outBcastPkt. The type is interface{} with range: 0..4294967295.
    OutBcastPkt interface{}

    // outMcastPkt. The type is interface{} with range: 0..4294967295.
    OutMcastPkt interface{}

    // outPausePkt. The type is interface{} with range: 0..4294967295.
    OutPausePkt interface{}

    // excessive. The type is interface{} with range: 0..4294967295.
    Excessive interface{}

    // collisions. The type is interface{} with range: 0..4294967295.
    Collisions interface{}

    // deferred. The type is interface{} with range: 0..4294967295.
    Deferred interface{}

    // single. The type is interface{} with range: 0..4294967295.
    Single interface{}

    // multiple. The type is interface{} with range: 0..4294967295.
    Multiple interface{}

    // outFcsErr. The type is interface{} with range: 0..4294967295.
    OutFcsErr interface{}

    // late. The type is interface{} with range: 0..4294967295.
    Late interface{}

    // rx tx 64 Octets. The type is interface{} with range: 0..4294967295.
    RxTx64Octets interface{}

    // rx tx 65 127 Octets. The type is interface{} with range: 0..4294967295.
    RxTx65127Octets interface{}

    // rx tx 128 255 Octets. The type is interface{} with range: 0..4294967295.
    RxTx128255Octets interface{}

    // rx tx 256 511 Octets. The type is interface{} with range: 0..4294967295.
    RxTx256511Octets interface{}

    // rx tx 512 1023 Octets. The type is interface{} with range: 0..4294967295.
    RxTx5121023Octets interface{}

    // rx tx 1024 Max Octets. The type is interface{} with range: 0..4294967295.
    RxTx1024MaxOctets interface{}

    // inDiscards. The type is interface{} with range: 0..4294967295.
    InDiscards interface{}

    // inFiltered. The type is interface{} with range: 0..4294967295.
    InFiltered interface{}

    // outFiltered. The type is interface{} with range: 0..4294967295.
    OutFiltered interface{}
}

func (mlanStats *Mlan_Nodes_Node_PortCountersNumbers_PortCountersNumber_PortCounters_MlanStats) GetEntityData() *types.CommonEntityData {
    mlanStats.EntityData.YFilter = mlanStats.YFilter
    mlanStats.EntityData.YangName = "mlan-stats"
    mlanStats.EntityData.BundleName = "cisco_ios_xr"
    mlanStats.EntityData.ParentYangName = "port-counters"
    mlanStats.EntityData.SegmentPath = "mlan-stats"
    mlanStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mlanStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mlanStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mlanStats.EntityData.Children = types.NewOrderedMap()
    mlanStats.EntityData.Leafs = types.NewOrderedMap()
    mlanStats.EntityData.Leafs.Append("in-good-octets-hi", types.YLeaf{"InGoodOctetsHi", mlanStats.InGoodOctetsHi})
    mlanStats.EntityData.Leafs.Append("in-good-octets", types.YLeaf{"InGoodOctets", mlanStats.InGoodOctets})
    mlanStats.EntityData.Leafs.Append("in-bad-octets", types.YLeaf{"InBadOctets", mlanStats.InBadOctets})
    mlanStats.EntityData.Leafs.Append("in-unicast-pkt", types.YLeaf{"InUnicastPkt", mlanStats.InUnicastPkt})
    mlanStats.EntityData.Leafs.Append("in-bcast-pkt", types.YLeaf{"InBcastPkt", mlanStats.InBcastPkt})
    mlanStats.EntityData.Leafs.Append("in-mcast-pkt", types.YLeaf{"InMcastPkt", mlanStats.InMcastPkt})
    mlanStats.EntityData.Leafs.Append("in-pause-pkt", types.YLeaf{"InPausePkt", mlanStats.InPausePkt})
    mlanStats.EntityData.Leafs.Append("in-undersize-pkt", types.YLeaf{"InUndersizePkt", mlanStats.InUndersizePkt})
    mlanStats.EntityData.Leafs.Append("in-fragments", types.YLeaf{"InFragments", mlanStats.InFragments})
    mlanStats.EntityData.Leafs.Append("in-oversize", types.YLeaf{"InOversize", mlanStats.InOversize})
    mlanStats.EntityData.Leafs.Append("in-jabber", types.YLeaf{"InJabber", mlanStats.InJabber})
    mlanStats.EntityData.Leafs.Append("in-rx-err", types.YLeaf{"InRxErr", mlanStats.InRxErr})
    mlanStats.EntityData.Leafs.Append("in-fcs-err", types.YLeaf{"InFcsErr", mlanStats.InFcsErr})
    mlanStats.EntityData.Leafs.Append("out-octets-hi", types.YLeaf{"OutOctetsHi", mlanStats.OutOctetsHi})
    mlanStats.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", mlanStats.OutOctets})
    mlanStats.EntityData.Leafs.Append("out-unicast-pkt", types.YLeaf{"OutUnicastPkt", mlanStats.OutUnicastPkt})
    mlanStats.EntityData.Leafs.Append("out-bcast-pkt", types.YLeaf{"OutBcastPkt", mlanStats.OutBcastPkt})
    mlanStats.EntityData.Leafs.Append("out-mcast-pkt", types.YLeaf{"OutMcastPkt", mlanStats.OutMcastPkt})
    mlanStats.EntityData.Leafs.Append("out-pause-pkt", types.YLeaf{"OutPausePkt", mlanStats.OutPausePkt})
    mlanStats.EntityData.Leafs.Append("excessive", types.YLeaf{"Excessive", mlanStats.Excessive})
    mlanStats.EntityData.Leafs.Append("collisions", types.YLeaf{"Collisions", mlanStats.Collisions})
    mlanStats.EntityData.Leafs.Append("deferred", types.YLeaf{"Deferred", mlanStats.Deferred})
    mlanStats.EntityData.Leafs.Append("single", types.YLeaf{"Single", mlanStats.Single})
    mlanStats.EntityData.Leafs.Append("multiple", types.YLeaf{"Multiple", mlanStats.Multiple})
    mlanStats.EntityData.Leafs.Append("out-fcs-err", types.YLeaf{"OutFcsErr", mlanStats.OutFcsErr})
    mlanStats.EntityData.Leafs.Append("late", types.YLeaf{"Late", mlanStats.Late})
    mlanStats.EntityData.Leafs.Append("rx-tx-64-octets", types.YLeaf{"RxTx64Octets", mlanStats.RxTx64Octets})
    mlanStats.EntityData.Leafs.Append("rx-tx-65-127-octets", types.YLeaf{"RxTx65127Octets", mlanStats.RxTx65127Octets})
    mlanStats.EntityData.Leafs.Append("rx-tx-128-255-octets", types.YLeaf{"RxTx128255Octets", mlanStats.RxTx128255Octets})
    mlanStats.EntityData.Leafs.Append("rx-tx-256-511-octets", types.YLeaf{"RxTx256511Octets", mlanStats.RxTx256511Octets})
    mlanStats.EntityData.Leafs.Append("rx-tx-512-1023-octets", types.YLeaf{"RxTx5121023Octets", mlanStats.RxTx5121023Octets})
    mlanStats.EntityData.Leafs.Append("rx-tx-1024-max-octets", types.YLeaf{"RxTx1024MaxOctets", mlanStats.RxTx1024MaxOctets})
    mlanStats.EntityData.Leafs.Append("in-discards", types.YLeaf{"InDiscards", mlanStats.InDiscards})
    mlanStats.EntityData.Leafs.Append("in-filtered", types.YLeaf{"InFiltered", mlanStats.InFiltered})
    mlanStats.EntityData.Leafs.Append("out-filtered", types.YLeaf{"OutFiltered", mlanStats.OutFiltered})

    mlanStats.EntityData.YListKeys = []string {}

    return &(mlanStats.EntityData)
}

// Mlan_Nodes_Node_AtuEntryNumbers
// Table of switch ATU
type Mlan_Nodes_Node_AtuEntryNumbers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Entry number. The type is slice of
    // Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber.
    AtuEntryNumber []*Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber
}

func (atuEntryNumbers *Mlan_Nodes_Node_AtuEntryNumbers) GetEntityData() *types.CommonEntityData {
    atuEntryNumbers.EntityData.YFilter = atuEntryNumbers.YFilter
    atuEntryNumbers.EntityData.YangName = "atu-entry-numbers"
    atuEntryNumbers.EntityData.BundleName = "cisco_ios_xr"
    atuEntryNumbers.EntityData.ParentYangName = "node"
    atuEntryNumbers.EntityData.SegmentPath = "atu-entry-numbers"
    atuEntryNumbers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    atuEntryNumbers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    atuEntryNumbers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    atuEntryNumbers.EntityData.Children = types.NewOrderedMap()
    atuEntryNumbers.EntityData.Children.Append("atu-entry-number", types.YChild{"AtuEntryNumber", nil})
    for i := range atuEntryNumbers.AtuEntryNumber {
        atuEntryNumbers.EntityData.Children.Append(types.GetSegmentPath(atuEntryNumbers.AtuEntryNumber[i]), types.YChild{"AtuEntryNumber", atuEntryNumbers.AtuEntryNumber[i]})
    }
    atuEntryNumbers.EntityData.Leafs = types.NewOrderedMap()

    atuEntryNumbers.EntityData.YListKeys = []string {}

    return &(atuEntryNumbers.EntityData)
}

// Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber
// Entry number
type Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. entry number. The type is interface{} with range:
    // 0..4294967295.
    Entry interface{}

    // mlan switch counters info.
    SwitchCounters Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters
}

func (atuEntryNumber *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber) GetEntityData() *types.CommonEntityData {
    atuEntryNumber.EntityData.YFilter = atuEntryNumber.YFilter
    atuEntryNumber.EntityData.YangName = "atu-entry-number"
    atuEntryNumber.EntityData.BundleName = "cisco_ios_xr"
    atuEntryNumber.EntityData.ParentYangName = "atu-entry-numbers"
    atuEntryNumber.EntityData.SegmentPath = "atu-entry-number" + types.AddKeyToken(atuEntryNumber.Entry, "entry")
    atuEntryNumber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    atuEntryNumber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    atuEntryNumber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    atuEntryNumber.EntityData.Children = types.NewOrderedMap()
    atuEntryNumber.EntityData.Children.Append("switch-counters", types.YChild{"SwitchCounters", &atuEntryNumber.SwitchCounters})
    atuEntryNumber.EntityData.Leafs = types.NewOrderedMap()
    atuEntryNumber.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", atuEntryNumber.Entry})

    atuEntryNumber.EntityData.YListKeys = []string {"Entry"}

    return &(atuEntryNumber.EntityData)
}

// Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters
// mlan switch counters info
type Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index of ATU Entry. The type is interface{} with range: 0..4294967295.
    EntryNum interface{}

    // Switch ATU Data.
    Atu Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu
}

func (switchCounters *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters) GetEntityData() *types.CommonEntityData {
    switchCounters.EntityData.YFilter = switchCounters.YFilter
    switchCounters.EntityData.YangName = "switch-counters"
    switchCounters.EntityData.BundleName = "cisco_ios_xr"
    switchCounters.EntityData.ParentYangName = "atu-entry-number"
    switchCounters.EntityData.SegmentPath = "switch-counters"
    switchCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchCounters.EntityData.Children = types.NewOrderedMap()
    switchCounters.EntityData.Children.Append("atu", types.YChild{"Atu", &switchCounters.Atu})
    switchCounters.EntityData.Leafs = types.NewOrderedMap()
    switchCounters.EntityData.Leafs.Append("entry-num", types.YLeaf{"EntryNum", switchCounters.EntryNum})

    switchCounters.EntityData.YListKeys = []string {}

    return &(switchCounters.EntityData)
}

// Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu
// Switch ATU Data
type Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dbNum. The type is interface{} with range: 0..65535.
    DbNum interface{}

    // priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // trunk. The type is bool.
    Trunk interface{}

    // dpv. The type is interface{} with range: 0..255.
    Dpv interface{}

    // es. The type is interface{} with range: 0..255.
    Es interface{}

    // macaddr. The type is slice of
    // Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu_Macaddr.
    Macaddr []*Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu_Macaddr
}

func (atu *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu) GetEntityData() *types.CommonEntityData {
    atu.EntityData.YFilter = atu.YFilter
    atu.EntityData.YangName = "atu"
    atu.EntityData.BundleName = "cisco_ios_xr"
    atu.EntityData.ParentYangName = "switch-counters"
    atu.EntityData.SegmentPath = "atu"
    atu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    atu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    atu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    atu.EntityData.Children = types.NewOrderedMap()
    atu.EntityData.Children.Append("macaddr", types.YChild{"Macaddr", nil})
    for i := range atu.Macaddr {
        atu.EntityData.Children.Append(types.GetSegmentPath(atu.Macaddr[i]), types.YChild{"Macaddr", atu.Macaddr[i]})
    }
    atu.EntityData.Leafs = types.NewOrderedMap()
    atu.EntityData.Leafs.Append("db-num", types.YLeaf{"DbNum", atu.DbNum})
    atu.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", atu.Priority})
    atu.EntityData.Leafs.Append("trunk", types.YLeaf{"Trunk", atu.Trunk})
    atu.EntityData.Leafs.Append("dpv", types.YLeaf{"Dpv", atu.Dpv})
    atu.EntityData.Leafs.Append("es", types.YLeaf{"Es", atu.Es})

    atu.EntityData.YListKeys = []string {}

    return &(atu.EntityData)
}

// Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu_Macaddr
// macaddr
type Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu_Macaddr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..65535.
    Entry interface{}
}

func (macaddr *Mlan_Nodes_Node_AtuEntryNumbers_AtuEntryNumber_SwitchCounters_Atu_Macaddr) GetEntityData() *types.CommonEntityData {
    macaddr.EntityData.YFilter = macaddr.YFilter
    macaddr.EntityData.YangName = "macaddr"
    macaddr.EntityData.BundleName = "cisco_ios_xr"
    macaddr.EntityData.ParentYangName = "atu"
    macaddr.EntityData.SegmentPath = "macaddr"
    macaddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macaddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macaddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macaddr.EntityData.Children = types.NewOrderedMap()
    macaddr.EntityData.Leafs = types.NewOrderedMap()
    macaddr.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", macaddr.Entry})

    macaddr.EntityData.YListKeys = []string {}

    return &(macaddr.EntityData)
}

