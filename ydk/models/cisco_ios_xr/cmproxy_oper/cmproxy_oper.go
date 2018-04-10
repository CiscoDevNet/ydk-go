// This module contains a collection of YANG definitions
// for Cisco IOS-XR cmproxy package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sdr-inventory-vm: Platform VM information
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    sdrInventoryVm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrInventoryVm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrInventoryVm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrInventoryVm.EntityData.Children = make(map[string]types.YChild)
    sdrInventoryVm.EntityData.Children["nodes"] = types.YChild{"Nodes", &sdrInventoryVm.Nodes}
    sdrInventoryVm.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sdrInventoryVm.EntityData)
}

// SdrInventoryVm_Nodes
// Node directory
type SdrInventoryVm_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node name. The type is slice of SdrInventoryVm_Nodes_Node.
    Node []SdrInventoryVm_Nodes_Node
}

func (nodes *SdrInventoryVm_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "sdr-inventory-vm"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// SdrInventoryVm_Nodes_Node
// Node name
type SdrInventoryVm_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}

    // VM Information.
    NodeEntries SdrInventoryVm_Nodes_Node_NodeEntries
}

func (node *SdrInventoryVm_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[name='" + fmt.Sprintf("%v", node.Name) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["node-entries"] = types.YChild{"NodeEntries", &node.NodeEntries}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["name"] = types.YLeaf{"Name", node.Name}
    return &(node.EntityData)
}

// SdrInventoryVm_Nodes_Node_NodeEntries
// VM Information
type SdrInventoryVm_Nodes_Node_NodeEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VM information for a node. The type is slice of
    // SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry.
    NodeEntry []SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry
}

func (nodeEntries *SdrInventoryVm_Nodes_Node_NodeEntries) GetEntityData() *types.CommonEntityData {
    nodeEntries.EntityData.YFilter = nodeEntries.YFilter
    nodeEntries.EntityData.YangName = "node-entries"
    nodeEntries.EntityData.BundleName = "cisco_ios_xr"
    nodeEntries.EntityData.ParentYangName = "node"
    nodeEntries.EntityData.SegmentPath = "node-entries"
    nodeEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeEntries.EntityData.Children = make(map[string]types.YChild)
    nodeEntries.EntityData.Children["node-entry"] = types.YChild{"NodeEntry", nil}
    for i := range nodeEntries.NodeEntry {
        nodeEntries.EntityData.Children[types.GetSegmentPath(&nodeEntries.NodeEntry[i])] = types.YChild{"NodeEntry", &nodeEntries.NodeEntry[i]}
    }
    nodeEntries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodeEntries.EntityData)
}

// SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry
// VM information for a node
type SdrInventoryVm_Nodes_Node_NodeEntries_NodeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    nodeEntry.EntityData.SegmentPath = "node-entry" + "[name='" + fmt.Sprintf("%v", nodeEntry.Name) + "']"
    nodeEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeEntry.EntityData.Children = make(map[string]types.YChild)
    nodeEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeEntry.EntityData.Leafs["name"] = types.YLeaf{"Name", nodeEntry.Name}
    nodeEntry.EntityData.Leafs["valid"] = types.YLeaf{"Valid", nodeEntry.Valid}
    nodeEntry.EntityData.Leafs["card-type"] = types.YLeaf{"CardType", nodeEntry.CardType}
    nodeEntry.EntityData.Leafs["card-type-string"] = types.YLeaf{"CardTypeString", nodeEntry.CardTypeString}
    nodeEntry.EntityData.Leafs["nodeid"] = types.YLeaf{"Nodeid", nodeEntry.Nodeid}
    nodeEntry.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", nodeEntry.NodeName}
    nodeEntry.EntityData.Leafs["partner-id"] = types.YLeaf{"PartnerId", nodeEntry.PartnerId}
    nodeEntry.EntityData.Leafs["partner-name"] = types.YLeaf{"PartnerName", nodeEntry.PartnerName}
    nodeEntry.EntityData.Leafs["red-state"] = types.YLeaf{"RedState", nodeEntry.RedState}
    nodeEntry.EntityData.Leafs["red-state-string"] = types.YLeaf{"RedStateString", nodeEntry.RedStateString}
    nodeEntry.EntityData.Leafs["node-sw-state"] = types.YLeaf{"NodeSwState", nodeEntry.NodeSwState}
    nodeEntry.EntityData.Leafs["node-sw-state-string"] = types.YLeaf{"NodeSwStateString", nodeEntry.NodeSwStateString}
    nodeEntry.EntityData.Leafs["prev-sw-state"] = types.YLeaf{"PrevSwState", nodeEntry.PrevSwState}
    nodeEntry.EntityData.Leafs["prev-sw-state-string"] = types.YLeaf{"PrevSwStateString", nodeEntry.PrevSwStateString}
    nodeEntry.EntityData.Leafs["node-ip"] = types.YLeaf{"NodeIp", nodeEntry.NodeIp}
    nodeEntry.EntityData.Leafs["node-ipv4-string"] = types.YLeaf{"NodeIpv4String", nodeEntry.NodeIpv4String}
    return &(nodeEntry.EntityData)
}

