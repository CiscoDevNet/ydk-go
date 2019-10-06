// This module contains a collection of YANG definitions
// for Cisco IOS-XR cmproxy package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sdr-inventory-vm: Platform VM information
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package cmproxy_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cmproxy_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cmproxy-oper sdr-inventory-vm}", reflect.TypeOf(SdrInventoryVm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cmproxy-oper:sdr-inventory-vm", reflect.TypeOf(SdrInventoryVm{}))
}

// SdrInventoryVm
// Platform VM information
type SdrInventoryVm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node directory.
    Nodes SdrInventoryVm_Nodes
}

func (sdrInventoryVm *SdrInventoryVm) GetEntityData() *types.CommonEntityData {
    sdrInventoryVm.EntityData.YFilter = sdrInventoryVm.YFilter
    sdrInventoryVm.EntityData.YangName = "sdr-inventory-vm"
    sdrInventoryVm.EntityData.BundleName = "cisco_ios_xr"
    sdrInventoryVm.EntityData.ParentYangName = "Cisco-IOS-XR-cmproxy-oper"
    sdrInventoryVm.EntityData.SegmentPath = "Cisco-IOS-XR-cmproxy-oper:sdr-inventory-vm"
    sdrInventoryVm.EntityData.AbsolutePath = sdrInventoryVm.EntityData.SegmentPath
    sdrInventoryVm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrInventoryVm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrInventoryVm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrInventoryVm.EntityData.Children = types.NewOrderedMap()
    sdrInventoryVm.EntityData.Children.Append("nodes", types.YChild{"Nodes", &sdrInventoryVm.Nodes})
    sdrInventoryVm.EntityData.Leafs = types.NewOrderedMap()

    sdrInventoryVm.EntityData.YListKeys = []string {}

    return &(sdrInventoryVm.EntityData)
}

// SdrInventoryVm_Nodes
// Node directory
type SdrInventoryVm_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node name. The type is slice of SdrInventoryVm_Nodes_Node.
    Node []*SdrInventoryVm_Nodes_Node
}

func (nodes *SdrInventoryVm_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "sdr-inventory-vm"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-cmproxy-oper:sdr-inventory-vm/" + nodes.EntityData.SegmentPath
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

// SdrInventoryVm_Nodes_Node
// Node name
type SdrInventoryVm_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // VM Information.
    NodeEntries SdrInventoryVm_Nodes_Node_NodeEntries
}

func (node *SdrInventoryVm_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Name, "name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-cmproxy-oper:sdr-inventory-vm/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("node-entries", types.YChild{"NodeEntries", &node.NodeEntries})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("name", types.YLeaf{"Name", node.Name})

    node.EntityData.YListKeys = []string {"Name"}

    return &(node.EntityData)
}

// SdrInventoryVm_Nodes_Node_NodeEntries
// VM Information
type SdrInventoryVm_Nodes_Node_NodeEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VM information for a node. The type is slice of
    // SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry.
    NodeEntry []*SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry
}

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetEntityData() *types.CommonEntityData {
    nodeEntries.EntityData.YFilter = nodeEntries.YFilter
    nodeEntries.EntityData.YangName = "node-entries"
    nodeEntries.EntityData.BundleName = "cisco_ios_xr"
    nodeEntries.EntityData.ParentYangName = "node"
    nodeEntries.EntityData.SegmentPath = "node-entries"
    nodeEntries.EntityData.AbsolutePath = "Cisco-IOS-XR-cmproxy-oper:sdr-inventory-vm/nodes/node/" + nodeEntries.EntityData.SegmentPath
    nodeEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeEntries.EntityData.Children = types.NewOrderedMap()
    nodeEntries.EntityData.Children.Append("node-entry", types.YChild{"NodeEntry", nil})
    for i := range nodeEntries.NodeEntry {
        nodeEntries.EntityData.Children.Append(types.GetSegmentPath(nodeEntries.NodeEntry[i]), types.YChild{"NodeEntry", nodeEntries.NodeEntry[i]})
    }
    nodeEntries.EntityData.Leafs = types.NewOrderedMap()

    nodeEntries.EntityData.YListKeys = []string {}

    return &(nodeEntries.EntityData)
}

// SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry
// VM information for a node
type SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // valid flag. The type is interface{} with range: 0..4294967295.
    Valid interface{}

    // card type. The type is interface{} with range: 0..4294967295.
    CardType interface{}

    // card type string. The type is string with length: 0..32.
    CardTypeString interface{}

    // node ID. The type is interface{} with range: -2147483648..2147483647.
    Nodeid interface{}

    // node name string. The type is string with length: 0..32.
    NodeName interface{}

    // partner node id. The type is interface{} with range:
    // -2147483648..2147483647.
    PartnerId interface{}

    // partner name string. The type is string with length: 0..32.
    PartnerName interface{}

    // redundancy state. The type is interface{} with range: 0..4294967295.
    RedState interface{}

    // redundancy state string. The type is string with length: 0..32.
    RedStateString interface{}

    // current software state. The type is interface{} with range: 0..4294967295.
    NodeSwState interface{}

    // current software state string. The type is string with length: 0..32.
    NodeSwStateString interface{}

    // previous software state. The type is interface{} with range: 0..4294967295.
    PrevSwState interface{}

    // previous software state string. The type is string with length: 0..32.
    PrevSwStateString interface{}

    // node IP address. The type is interface{} with range: 0..4294967295.
    NodeIp interface{}

    // node IPv4 address string. The type is string with length: 0..16.
    NodeIpv4String interface{}
}

func (nodeEntry *SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry) GetEntityData() *types.CommonEntityData {
    nodeEntry.EntityData.YFilter = nodeEntry.YFilter
    nodeEntry.EntityData.YangName = "node-entry"
    nodeEntry.EntityData.BundleName = "cisco_ios_xr"
    nodeEntry.EntityData.ParentYangName = "node-entries"
    nodeEntry.EntityData.SegmentPath = "node-entry" + types.AddKeyToken(nodeEntry.Name, "name")
    nodeEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-cmproxy-oper:sdr-inventory-vm/nodes/node/node-entries/" + nodeEntry.EntityData.SegmentPath
    nodeEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeEntry.EntityData.Children = types.NewOrderedMap()
    nodeEntry.EntityData.Leafs = types.NewOrderedMap()
    nodeEntry.EntityData.Leafs.Append("name", types.YLeaf{"Name", nodeEntry.Name})
    nodeEntry.EntityData.Leafs.Append("valid", types.YLeaf{"Valid", nodeEntry.Valid})
    nodeEntry.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", nodeEntry.CardType})
    nodeEntry.EntityData.Leafs.Append("card-type-string", types.YLeaf{"CardTypeString", nodeEntry.CardTypeString})
    nodeEntry.EntityData.Leafs.Append("nodeid", types.YLeaf{"Nodeid", nodeEntry.Nodeid})
    nodeEntry.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", nodeEntry.NodeName})
    nodeEntry.EntityData.Leafs.Append("partner-id", types.YLeaf{"PartnerId", nodeEntry.PartnerId})
    nodeEntry.EntityData.Leafs.Append("partner-name", types.YLeaf{"PartnerName", nodeEntry.PartnerName})
    nodeEntry.EntityData.Leafs.Append("red-state", types.YLeaf{"RedState", nodeEntry.RedState})
    nodeEntry.EntityData.Leafs.Append("red-state-string", types.YLeaf{"RedStateString", nodeEntry.RedStateString})
    nodeEntry.EntityData.Leafs.Append("node-sw-state", types.YLeaf{"NodeSwState", nodeEntry.NodeSwState})
    nodeEntry.EntityData.Leafs.Append("node-sw-state-string", types.YLeaf{"NodeSwStateString", nodeEntry.NodeSwStateString})
    nodeEntry.EntityData.Leafs.Append("prev-sw-state", types.YLeaf{"PrevSwState", nodeEntry.PrevSwState})
    nodeEntry.EntityData.Leafs.Append("prev-sw-state-string", types.YLeaf{"PrevSwStateString", nodeEntry.PrevSwStateString})
    nodeEntry.EntityData.Leafs.Append("node-ip", types.YLeaf{"NodeIp", nodeEntry.NodeIp})
    nodeEntry.EntityData.Leafs.Append("node-ipv4-string", types.YLeaf{"NodeIpv4String", nodeEntry.NodeIpv4String})

    nodeEntry.EntityData.YListKeys = []string {"Name"}

    return &(nodeEntry.EntityData)
}

