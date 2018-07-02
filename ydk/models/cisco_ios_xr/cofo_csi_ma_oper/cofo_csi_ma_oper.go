// This module contains a collection of YANG definitions
// for Cisco IOS-XR cofo-csi-ma package operational data.
// 
// This module contains definitions
// for the following management objects:
//   cross-sdr-intf-ma: Csi-ma operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package cofo_csi_ma_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cofo_csi_ma_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cofo-csi-ma-oper cross-sdr-intf-ma}", reflect.TypeOf(CrossSdrIntfMa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cofo-csi-ma-oper:cross-sdr-intf-ma", reflect.TypeOf(CrossSdrIntfMa{}))
}

// CsiMaFoTimerState represents Fail Over Timer state
type CsiMaFoTimerState string

const (
    // Default state
    CsiMaFoTimerState_csi_ma_fo_default CsiMaFoTimerState = "csi-ma-fo-default"

    // Push self data
    CsiMaFoTimerState_csi_ma_fo_push_self_data CsiMaFoTimerState = "csi-ma-fo-push-self-data"

    // Sync remote data
    CsiMaFoTimerState_csi_ma_fo_sync_rem_data CsiMaFoTimerState = "csi-ma-fo-sync-rem-data"

    // Synced
    CsiMaFoTimerState_csi_ma_fo_synced CsiMaFoTimerState = "csi-ma-fo-synced"

    // Max state
    CsiMaFoTimerState_csi_ma_fo_max CsiMaFoTimerState = "csi-ma-fo-max"
)

// CsiMaAfi represents CSI MA Address family enum type
type CsiMaAfi string

const (
    // IPv4 address family
    CsiMaAfi_csi_afi_ipv4 CsiMaAfi = "csi-afi-ipv4"

    // IPv6 address family
    CsiMaAfi_csi_afi_ipv6 CsiMaAfi = "csi-afi-ipv6"

    // Invalid address family
    CsiMaAfi_csi_afi_invalid CsiMaAfi = "csi-afi-invalid"
)

// CsiMaItemState represents CSI MA item state enum type
type CsiMaItemState string

const (
    // Interface create requested, wating for ack
    CsiMaItemState_csi_ma_item_create_req CsiMaItemState = "csi-ma-item-create-req"

    // Requested attributes passed from pub-sub admin
    CsiMaItemState_csi_ma_item_attr_req CsiMaItemState = "csi-ma-item-attr-req"

    // Valid entry
    CsiMaItemState_csi_ma_item_valid CsiMaItemState = "csi-ma-item-valid"

    // Replicated and synced to all nodes
    CsiMaItemState_csi_ma_item_synced CsiMaItemState = "csi-ma-item-synced"

    // Marked for sweep at EOD
    CsiMaItemState_csi_ma_item_mark_ed CsiMaItemState = "csi-ma-item-mark-ed"

    // Marked for delete in IM and then purge entry
    CsiMaItemState_csi_ma_item_invalid CsiMaItemState = "csi-ma-item-invalid"

    // Interface delete requested, wating for ack
    CsiMaItemState_csi_ma_item_delete_req CsiMaItemState = "csi-ma-item-delete-req"

    // Invalid state
    CsiMaItemState_csi_ma_item_max_states CsiMaItemState = "csi-ma-item-max-states"
)

// CsiMaConnState represents CSI-MA connection state enum type
type CsiMaConnState string

const (
    // Default connection state
    CsiMaConnState_csi_ma_conn_default CsiMaConnState = "csi-ma-conn-default"

    // Connection closed
    CsiMaConnState_csi_ma_conn_closed CsiMaConnState = "csi-ma-conn-closed"

    // Connection opened
    CsiMaConnState_csi_ma_conn_opened CsiMaConnState = "csi-ma-conn-opened"

    // Connection synced
    CsiMaConnState_csi_ma_conn_synced CsiMaConnState = "csi-ma-conn-synced"

    // Unknown connection state
    CsiMaConnState_csi_ma_conn_max CsiMaConnState = "csi-ma-conn-max"
)

// CrossSdrIntfMa
// Csi-ma operational data
type CrossSdrIntfMa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific operational data.
    Nodes CrossSdrIntfMa_Nodes
}

func (crossSdrIntfMa *CrossSdrIntfMa) GetEntityData() *types.CommonEntityData {
    crossSdrIntfMa.EntityData.YFilter = crossSdrIntfMa.YFilter
    crossSdrIntfMa.EntityData.YangName = "cross-sdr-intf-ma"
    crossSdrIntfMa.EntityData.BundleName = "cisco_ios_xr"
    crossSdrIntfMa.EntityData.ParentYangName = "Cisco-IOS-XR-cofo-csi-ma-oper"
    crossSdrIntfMa.EntityData.SegmentPath = "Cisco-IOS-XR-cofo-csi-ma-oper:cross-sdr-intf-ma"
    crossSdrIntfMa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crossSdrIntfMa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crossSdrIntfMa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crossSdrIntfMa.EntityData.Children = types.NewOrderedMap()
    crossSdrIntfMa.EntityData.Children.Append("nodes", types.YChild{"Nodes", &crossSdrIntfMa.Nodes})
    crossSdrIntfMa.EntityData.Leafs = types.NewOrderedMap()

    crossSdrIntfMa.EntityData.YListKeys = []string {}

    return &(crossSdrIntfMa.EntityData)
}

// CrossSdrIntfMa_Nodes
// Node-specific operational data
type CrossSdrIntfMa_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Operational data for a particular node. The type is slice of
    // CrossSdrIntfMa_Nodes_Node.
    Node []*CrossSdrIntfMa_Nodes_Node
}

func (nodes *CrossSdrIntfMa_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "cross-sdr-intf-ma"
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

// CrossSdrIntfMa_Nodes_Node
// Operational data for a particular node
type CrossSdrIntfMa_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Connected node slices information.
    RackIds CrossSdrIntfMa_Nodes_Node_RackIds

    // Admin data for csi-index.
    CsiIndexes CrossSdrIntfMa_Nodes_Node_CsiIndexes

    // Interface data for csi interfaces.
    InterfaceNames CrossSdrIntfMa_Nodes_Node_InterfaceNames

    // Global data associated with csi-ma.
    Global CrossSdrIntfMa_Nodes_Node_Global

    // Remote data for SDR ID.
    SdrIds CrossSdrIntfMa_Nodes_Node_SdrIds
}

func (node *CrossSdrIntfMa_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("rack-ids", types.YChild{"RackIds", &node.RackIds})
    node.EntityData.Children.Append("csi-indexes", types.YChild{"CsiIndexes", &node.CsiIndexes})
    node.EntityData.Children.Append("interface-names", types.YChild{"InterfaceNames", &node.InterfaceNames})
    node.EntityData.Children.Append("global", types.YChild{"Global", &node.Global})
    node.EntityData.Children.Append("sdr-ids", types.YChild{"SdrIds", &node.SdrIds})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_RackIds
// Connected node slices information
type CrossSdrIntfMa_Nodes_Node_RackIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Connected node rack-id. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_RackIds_RackId.
    RackId []*CrossSdrIntfMa_Nodes_Node_RackIds_RackId
}

func (rackIds *CrossSdrIntfMa_Nodes_Node_RackIds) GetEntityData() *types.CommonEntityData {
    rackIds.EntityData.YFilter = rackIds.YFilter
    rackIds.EntityData.YangName = "rack-ids"
    rackIds.EntityData.BundleName = "cisco_ios_xr"
    rackIds.EntityData.ParentYangName = "node"
    rackIds.EntityData.SegmentPath = "rack-ids"
    rackIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackIds.EntityData.Children = types.NewOrderedMap()
    rackIds.EntityData.Children.Append("rack-id", types.YChild{"RackId", nil})
    for i := range rackIds.RackId {
        rackIds.EntityData.Children.Append(types.GetSegmentPath(rackIds.RackId[i]), types.YChild{"RackId", rackIds.RackId[i]})
    }
    rackIds.EntityData.Leafs = types.NewOrderedMap()

    rackIds.EntityData.YListKeys = []string {}

    return &(rackIds.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_RackIds_RackId
// Connected node rack-id
type CrossSdrIntfMa_Nodes_Node_RackIds_RackId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack Id. The type is interface{} with range:
    // 0..65535.
    RackId interface{}

    // NodeInfo for a NodeId. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId.
    SlotId []*CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId
}

func (rackId *CrossSdrIntfMa_Nodes_Node_RackIds_RackId) GetEntityData() *types.CommonEntityData {
    rackId.EntityData.YFilter = rackId.YFilter
    rackId.EntityData.YangName = "rack-id"
    rackId.EntityData.BundleName = "cisco_ios_xr"
    rackId.EntityData.ParentYangName = "rack-ids"
    rackId.EntityData.SegmentPath = "rack-id" + types.AddKeyToken(rackId.RackId, "rack-id")
    rackId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rackId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rackId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rackId.EntityData.Children = types.NewOrderedMap()
    rackId.EntityData.Children.Append("slot-id", types.YChild{"SlotId", nil})
    for i := range rackId.SlotId {
        rackId.EntityData.Children.Append(types.GetSegmentPath(rackId.SlotId[i]), types.YChild{"SlotId", rackId.SlotId[i]})
    }
    rackId.EntityData.Leafs = types.NewOrderedMap()
    rackId.EntityData.Leafs.Append("rack-id", types.YLeaf{"RackId", rackId.RackId})

    rackId.EntityData.YListKeys = []string {"RackId"}

    return &(rackId.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId
// NodeInfo for a NodeId
type CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Slot Id. The type is interface{} with range:
    // 0..65535.
    SlotId interface{}

    // Rack slot ID. The type is interface{} with range: 0..18446744073709551615.
    RackSlotId interface{}

    // Node id. The type is interface{} with range: 0..4294967295.
    NodeId interface{}

    // Node up flag. The type is bool.
    NodeUp interface{}

    // Slice array associated with node. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId_SliceArr.
    SliceArr []*CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId_SliceArr
}

func (slotId *CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId) GetEntityData() *types.CommonEntityData {
    slotId.EntityData.YFilter = slotId.YFilter
    slotId.EntityData.YangName = "slot-id"
    slotId.EntityData.BundleName = "cisco_ios_xr"
    slotId.EntityData.ParentYangName = "rack-id"
    slotId.EntityData.SegmentPath = "slot-id" + types.AddKeyToken(slotId.SlotId, "slot-id")
    slotId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slotId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slotId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slotId.EntityData.Children = types.NewOrderedMap()
    slotId.EntityData.Children.Append("slice-arr", types.YChild{"SliceArr", nil})
    for i := range slotId.SliceArr {
        slotId.EntityData.Children.Append(types.GetSegmentPath(slotId.SliceArr[i]), types.YChild{"SliceArr", slotId.SliceArr[i]})
    }
    slotId.EntityData.Leafs = types.NewOrderedMap()
    slotId.EntityData.Leafs.Append("slot-id", types.YLeaf{"SlotId", slotId.SlotId})
    slotId.EntityData.Leafs.Append("rack-slot-id", types.YLeaf{"RackSlotId", slotId.RackSlotId})
    slotId.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", slotId.NodeId})
    slotId.EntityData.Leafs.Append("node-up", types.YLeaf{"NodeUp", slotId.NodeUp})

    slotId.EntityData.YListKeys = []string {"SlotId"}

    return &(slotId.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId_SliceArr
// Slice array associated with node
type CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId_SliceArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slice Node ID. The type is interface{} with range: 0..4294967295.
    SliceNodeId interface{}

    // Admin state UP flag. The type is bool.
    AdminUp interface{}

    // Oper state UP flag. The type is bool.
    OperUp interface{}

    // PIC value. The type is interface{} with range: 0..18446744073709551615.
    Pic interface{}

    // CSI PIC array. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    CsiPicArr []interface{}
}

func (sliceArr *CrossSdrIntfMa_Nodes_Node_RackIds_RackId_SlotId_SliceArr) GetEntityData() *types.CommonEntityData {
    sliceArr.EntityData.YFilter = sliceArr.YFilter
    sliceArr.EntityData.YangName = "slice-arr"
    sliceArr.EntityData.BundleName = "cisco_ios_xr"
    sliceArr.EntityData.ParentYangName = "slot-id"
    sliceArr.EntityData.SegmentPath = "slice-arr"
    sliceArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sliceArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sliceArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sliceArr.EntityData.Children = types.NewOrderedMap()
    sliceArr.EntityData.Leafs = types.NewOrderedMap()
    sliceArr.EntityData.Leafs.Append("slice-node-id", types.YLeaf{"SliceNodeId", sliceArr.SliceNodeId})
    sliceArr.EntityData.Leafs.Append("admin-up", types.YLeaf{"AdminUp", sliceArr.AdminUp})
    sliceArr.EntityData.Leafs.Append("oper-up", types.YLeaf{"OperUp", sliceArr.OperUp})
    sliceArr.EntityData.Leafs.Append("pic", types.YLeaf{"Pic", sliceArr.Pic})
    sliceArr.EntityData.Leafs.Append("csi-pic-arr", types.YLeaf{"CsiPicArr", sliceArr.CsiPicArr})

    sliceArr.EntityData.YListKeys = []string {}

    return &(sliceArr.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_CsiIndexes
// Admin data for csi-index
type CrossSdrIntfMa_Nodes_Node_CsiIndexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Admin data for a csi-index. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_CsiIndexes_CsiIndex.
    CsiIndex []*CrossSdrIntfMa_Nodes_Node_CsiIndexes_CsiIndex
}

func (csiIndexes *CrossSdrIntfMa_Nodes_Node_CsiIndexes) GetEntityData() *types.CommonEntityData {
    csiIndexes.EntityData.YFilter = csiIndexes.YFilter
    csiIndexes.EntityData.YangName = "csi-indexes"
    csiIndexes.EntityData.BundleName = "cisco_ios_xr"
    csiIndexes.EntityData.ParentYangName = "node"
    csiIndexes.EntityData.SegmentPath = "csi-indexes"
    csiIndexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csiIndexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csiIndexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csiIndexes.EntityData.Children = types.NewOrderedMap()
    csiIndexes.EntityData.Children.Append("csi-index", types.YChild{"CsiIndex", nil})
    for i := range csiIndexes.CsiIndex {
        csiIndexes.EntityData.Children.Append(types.GetSegmentPath(csiIndexes.CsiIndex[i]), types.YChild{"CsiIndex", csiIndexes.CsiIndex[i]})
    }
    csiIndexes.EntityData.Leafs = types.NewOrderedMap()

    csiIndexes.EntityData.YListKeys = []string {}

    return &(csiIndexes.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_CsiIndexes_CsiIndex
// Admin data for a csi-index
type CrossSdrIntfMa_Nodes_Node_CsiIndexes_CsiIndex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. CSI Index. The type is interface{} with range:
    // 1..15.
    CsiIndex interface{}

    // CSI Index. The type is interface{} with range: 0..4294967295.
    CsiIndexXr interface{}

    // Local SDR ID. The type is interface{} with range: 0..4294967295.
    SdrId interface{}

    // Peer SDR ID. The type is interface{} with range: 0..4294967295.
    PeerSdrId interface{}

    // Peer SDR name. The type is string.
    PeerSdrName interface{}

    // Item State. The type is CsiMaItemState.
    ItemState interface{}

    // CSI Helper reg flag. The type is bool.
    CsiHelperReg interface{}
}

func (csiIndex *CrossSdrIntfMa_Nodes_Node_CsiIndexes_CsiIndex) GetEntityData() *types.CommonEntityData {
    csiIndex.EntityData.YFilter = csiIndex.YFilter
    csiIndex.EntityData.YangName = "csi-index"
    csiIndex.EntityData.BundleName = "cisco_ios_xr"
    csiIndex.EntityData.ParentYangName = "csi-indexes"
    csiIndex.EntityData.SegmentPath = "csi-index" + types.AddKeyToken(csiIndex.CsiIndex, "csi-index")
    csiIndex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csiIndex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csiIndex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csiIndex.EntityData.Children = types.NewOrderedMap()
    csiIndex.EntityData.Leafs = types.NewOrderedMap()
    csiIndex.EntityData.Leafs.Append("csi-index", types.YLeaf{"CsiIndex", csiIndex.CsiIndex})
    csiIndex.EntityData.Leafs.Append("csi-index-xr", types.YLeaf{"CsiIndexXr", csiIndex.CsiIndexXr})
    csiIndex.EntityData.Leafs.Append("sdr-id", types.YLeaf{"SdrId", csiIndex.SdrId})
    csiIndex.EntityData.Leafs.Append("peer-sdr-id", types.YLeaf{"PeerSdrId", csiIndex.PeerSdrId})
    csiIndex.EntityData.Leafs.Append("peer-sdr-name", types.YLeaf{"PeerSdrName", csiIndex.PeerSdrName})
    csiIndex.EntityData.Leafs.Append("item-state", types.YLeaf{"ItemState", csiIndex.ItemState})
    csiIndex.EntityData.Leafs.Append("csi-helper-reg", types.YLeaf{"CsiHelperReg", csiIndex.CsiHelperReg})

    csiIndex.EntityData.YListKeys = []string {"CsiIndex"}

    return &(csiIndex.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_InterfaceNames
// Interface data for csi interfaces
type CrossSdrIntfMa_Nodes_Node_InterfaceNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface data for an Interface. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName.
    InterfaceName []*CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName
}

func (interfaceNames *CrossSdrIntfMa_Nodes_Node_InterfaceNames) GetEntityData() *types.CommonEntityData {
    interfaceNames.EntityData.YFilter = interfaceNames.YFilter
    interfaceNames.EntityData.YangName = "interface-names"
    interfaceNames.EntityData.BundleName = "cisco_ios_xr"
    interfaceNames.EntityData.ParentYangName = "node"
    interfaceNames.EntityData.SegmentPath = "interface-names"
    interfaceNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceNames.EntityData.Children = types.NewOrderedMap()
    interfaceNames.EntityData.Children.Append("interface-name", types.YChild{"InterfaceName", nil})
    for i := range interfaceNames.InterfaceName {
        interfaceNames.EntityData.Children.Append(types.GetSegmentPath(interfaceNames.InterfaceName[i]), types.YChild{"InterfaceName", interfaceNames.InterfaceName[i]})
    }
    interfaceNames.EntityData.Leafs = types.NewOrderedMap()

    interfaceNames.EntityData.YListKeys = []string {}

    return &(interfaceNames.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName
// Interface data for an Interface
type CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. CSI Index. The type is interface{} with range:
    // 1..15.
    CsiIndex interface{}

    // Interface name. The type is string.
    Name interface{}

    // Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    Handle interface{}

    // CSI Index. The type is interface{} with range: 0..4294967295.
    CsiIndexXr interface{}

    // SDR ID. The type is interface{} with range: 0..4294967295.
    SdrId interface{}

    // Peer SDR ID. The type is interface{} with range: 0..4294967295.
    PeerSdrId interface{}

    // Replication handle. The type is interface{} with range: 0..4294967295.
    ReplHandle interface{}

    // Replication fail count. The type is interface{} with range: 0..4294967295.
    ReplFailCount interface{}

    // Item state. The type is CsiMaItemState.
    ItemState interface{}

    // Interface state. The type is interface{} with range: 0..4294967295.
    IfState interface{}

    // Rem slice attribute. The type is bool.
    RemSliceAttr interface{}

    // Rem IP attribute. The type is bool.
    RemIpAttr interface{}

    // Remote published data.
    PubData CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData

    // Local IP addresses. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_LocalIpArr.
    LocalIpArr []*CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_LocalIpArr

    // Peer IP addresses. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PeerIpArr.
    PeerIpArr []*CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PeerIpArr
}

func (interfaceName *CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName) GetEntityData() *types.CommonEntityData {
    interfaceName.EntityData.YFilter = interfaceName.YFilter
    interfaceName.EntityData.YangName = "interface-name"
    interfaceName.EntityData.BundleName = "cisco_ios_xr"
    interfaceName.EntityData.ParentYangName = "interface-names"
    interfaceName.EntityData.SegmentPath = "interface-name" + types.AddKeyToken(interfaceName.CsiIndex, "csi-index")
    interfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceName.EntityData.Children = types.NewOrderedMap()
    interfaceName.EntityData.Children.Append("pub-data", types.YChild{"PubData", &interfaceName.PubData})
    interfaceName.EntityData.Children.Append("local-ip-arr", types.YChild{"LocalIpArr", nil})
    for i := range interfaceName.LocalIpArr {
        interfaceName.EntityData.Children.Append(types.GetSegmentPath(interfaceName.LocalIpArr[i]), types.YChild{"LocalIpArr", interfaceName.LocalIpArr[i]})
    }
    interfaceName.EntityData.Children.Append("peer-ip-arr", types.YChild{"PeerIpArr", nil})
    for i := range interfaceName.PeerIpArr {
        interfaceName.EntityData.Children.Append(types.GetSegmentPath(interfaceName.PeerIpArr[i]), types.YChild{"PeerIpArr", interfaceName.PeerIpArr[i]})
    }
    interfaceName.EntityData.Leafs = types.NewOrderedMap()
    interfaceName.EntityData.Leafs.Append("csi-index", types.YLeaf{"CsiIndex", interfaceName.CsiIndex})
    interfaceName.EntityData.Leafs.Append("name", types.YLeaf{"Name", interfaceName.Name})
    interfaceName.EntityData.Leafs.Append("handle", types.YLeaf{"Handle", interfaceName.Handle})
    interfaceName.EntityData.Leafs.Append("csi-index-xr", types.YLeaf{"CsiIndexXr", interfaceName.CsiIndexXr})
    interfaceName.EntityData.Leafs.Append("sdr-id", types.YLeaf{"SdrId", interfaceName.SdrId})
    interfaceName.EntityData.Leafs.Append("peer-sdr-id", types.YLeaf{"PeerSdrId", interfaceName.PeerSdrId})
    interfaceName.EntityData.Leafs.Append("repl-handle", types.YLeaf{"ReplHandle", interfaceName.ReplHandle})
    interfaceName.EntityData.Leafs.Append("repl-fail-count", types.YLeaf{"ReplFailCount", interfaceName.ReplFailCount})
    interfaceName.EntityData.Leafs.Append("item-state", types.YLeaf{"ItemState", interfaceName.ItemState})
    interfaceName.EntityData.Leafs.Append("if-state", types.YLeaf{"IfState", interfaceName.IfState})
    interfaceName.EntityData.Leafs.Append("rem-slice-attr", types.YLeaf{"RemSliceAttr", interfaceName.RemSliceAttr})
    interfaceName.EntityData.Leafs.Append("rem-ip-attr", types.YLeaf{"RemIpAttr", interfaceName.RemIpAttr})

    interfaceName.EntityData.YListKeys = []string {"CsiIndex"}

    return &(interfaceName.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData
// Remote published data
type CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CSI PIC Array. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    CsiPicArr []interface{}

    // CSI Index - IP Array. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr.
    CsiInfoArr []*CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr
}

func (pubData *CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData) GetEntityData() *types.CommonEntityData {
    pubData.EntityData.YFilter = pubData.YFilter
    pubData.EntityData.YangName = "pub-data"
    pubData.EntityData.BundleName = "cisco_ios_xr"
    pubData.EntityData.ParentYangName = "interface-name"
    pubData.EntityData.SegmentPath = "pub-data"
    pubData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pubData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pubData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pubData.EntityData.Children = types.NewOrderedMap()
    pubData.EntityData.Children.Append("csi-info-arr", types.YChild{"CsiInfoArr", nil})
    for i := range pubData.CsiInfoArr {
        pubData.EntityData.Children.Append(types.GetSegmentPath(pubData.CsiInfoArr[i]), types.YChild{"CsiInfoArr", pubData.CsiInfoArr[i]})
    }
    pubData.EntityData.Leafs = types.NewOrderedMap()
    pubData.EntityData.Leafs.Append("csi-pic-arr", types.YLeaf{"CsiPicArr", pubData.CsiPicArr})

    pubData.EntityData.YListKeys = []string {}

    return &(pubData.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr
// CSI Index - IP Array
type CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CSI Index. The type is interface{} with range: 0..4294967295.
    CsiIndex interface{}

    // CSI IP Array. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr_IpArr.
    IpArr []*CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr_IpArr
}

func (csiInfoArr *CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr) GetEntityData() *types.CommonEntityData {
    csiInfoArr.EntityData.YFilter = csiInfoArr.YFilter
    csiInfoArr.EntityData.YangName = "csi-info-arr"
    csiInfoArr.EntityData.BundleName = "cisco_ios_xr"
    csiInfoArr.EntityData.ParentYangName = "pub-data"
    csiInfoArr.EntityData.SegmentPath = "csi-info-arr"
    csiInfoArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csiInfoArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csiInfoArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csiInfoArr.EntityData.Children = types.NewOrderedMap()
    csiInfoArr.EntityData.Children.Append("ip-arr", types.YChild{"IpArr", nil})
    for i := range csiInfoArr.IpArr {
        csiInfoArr.EntityData.Children.Append(types.GetSegmentPath(csiInfoArr.IpArr[i]), types.YChild{"IpArr", csiInfoArr.IpArr[i]})
    }
    csiInfoArr.EntityData.Leafs = types.NewOrderedMap()
    csiInfoArr.EntityData.Leafs.Append("csi-index", types.YLeaf{"CsiIndex", csiInfoArr.CsiIndex})

    csiInfoArr.EntityData.YListKeys = []string {}

    return &(csiInfoArr.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr_IpArr
// CSI IP Array
type CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr_IpArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is CsiMaAfi.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipArr *CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PubData_CsiInfoArr_IpArr) GetEntityData() *types.CommonEntityData {
    ipArr.EntityData.YFilter = ipArr.YFilter
    ipArr.EntityData.YangName = "ip-arr"
    ipArr.EntityData.BundleName = "cisco_ios_xr"
    ipArr.EntityData.ParentYangName = "csi-info-arr"
    ipArr.EntityData.SegmentPath = "ip-arr"
    ipArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipArr.EntityData.Children = types.NewOrderedMap()
    ipArr.EntityData.Leafs = types.NewOrderedMap()
    ipArr.EntityData.Leafs.Append("af", types.YLeaf{"Af", ipArr.Af})
    ipArr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipArr.Ipv4})
    ipArr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipArr.Ipv6})

    ipArr.EntityData.YListKeys = []string {}

    return &(ipArr.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_LocalIpArr
// Local IP addresses
type CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_LocalIpArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is CsiMaAfi.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (localIpArr *CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_LocalIpArr) GetEntityData() *types.CommonEntityData {
    localIpArr.EntityData.YFilter = localIpArr.YFilter
    localIpArr.EntityData.YangName = "local-ip-arr"
    localIpArr.EntityData.BundleName = "cisco_ios_xr"
    localIpArr.EntityData.ParentYangName = "interface-name"
    localIpArr.EntityData.SegmentPath = "local-ip-arr"
    localIpArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localIpArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localIpArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localIpArr.EntityData.Children = types.NewOrderedMap()
    localIpArr.EntityData.Leafs = types.NewOrderedMap()
    localIpArr.EntityData.Leafs.Append("af", types.YLeaf{"Af", localIpArr.Af})
    localIpArr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", localIpArr.Ipv4})
    localIpArr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", localIpArr.Ipv6})

    localIpArr.EntityData.YListKeys = []string {}

    return &(localIpArr.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PeerIpArr
// Peer IP addresses
type CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PeerIpArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is CsiMaAfi.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (peerIpArr *CrossSdrIntfMa_Nodes_Node_InterfaceNames_InterfaceName_PeerIpArr) GetEntityData() *types.CommonEntityData {
    peerIpArr.EntityData.YFilter = peerIpArr.YFilter
    peerIpArr.EntityData.YangName = "peer-ip-arr"
    peerIpArr.EntityData.BundleName = "cisco_ios_xr"
    peerIpArr.EntityData.ParentYangName = "interface-name"
    peerIpArr.EntityData.SegmentPath = "peer-ip-arr"
    peerIpArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerIpArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerIpArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerIpArr.EntityData.Children = types.NewOrderedMap()
    peerIpArr.EntityData.Leafs = types.NewOrderedMap()
    peerIpArr.EntityData.Leafs.Append("af", types.YLeaf{"Af", peerIpArr.Af})
    peerIpArr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", peerIpArr.Ipv4})
    peerIpArr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", peerIpArr.Ipv6})

    peerIpArr.EntityData.YListKeys = []string {}

    return &(peerIpArr.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_Global
// Global data associated with csi-ma
type CrossSdrIntfMa_Nodes_Node_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory manager connection state. The type is CsiMaConnState.
    InvmgrConnState interface{}

    // Fail Over Timer state for remote data. The type is CsiMaFoTimerState.
    FailOverTimerState interface{}

    // Owner channel IM connection state. The type is CsiMaConnState.
    OwnImConnState interface{}

    // GDP IM connection state. The type is CsiMaConnState.
    GdpImConnState interface{}

    // L3P IM connection state. The type is CsiMaConnState.
    L3pImConnState interface{}
}

func (global *CrossSdrIntfMa_Nodes_Node_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "node"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("invmgr-conn-state", types.YLeaf{"InvmgrConnState", global.InvmgrConnState})
    global.EntityData.Leafs.Append("fail-over-timer-state", types.YLeaf{"FailOverTimerState", global.FailOverTimerState})
    global.EntityData.Leafs.Append("own-im-conn-state", types.YLeaf{"OwnImConnState", global.OwnImConnState})
    global.EntityData.Leafs.Append("gdp-im-conn-state", types.YLeaf{"GdpImConnState", global.GdpImConnState})
    global.EntityData.Leafs.Append("l3p-im-conn-state", types.YLeaf{"L3pImConnState", global.L3pImConnState})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_SdrIds
// Remote data for SDR ID
type CrossSdrIntfMa_Nodes_Node_SdrIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote data for sdr-id. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId.
    SdrId []*CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId
}

func (sdrIds *CrossSdrIntfMa_Nodes_Node_SdrIds) GetEntityData() *types.CommonEntityData {
    sdrIds.EntityData.YFilter = sdrIds.YFilter
    sdrIds.EntityData.YangName = "sdr-ids"
    sdrIds.EntityData.BundleName = "cisco_ios_xr"
    sdrIds.EntityData.ParentYangName = "node"
    sdrIds.EntityData.SegmentPath = "sdr-ids"
    sdrIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrIds.EntityData.Children = types.NewOrderedMap()
    sdrIds.EntityData.Children.Append("sdr-id", types.YChild{"SdrId", nil})
    for i := range sdrIds.SdrId {
        sdrIds.EntityData.Children.Append(types.GetSegmentPath(sdrIds.SdrId[i]), types.YChild{"SdrId", sdrIds.SdrId[i]})
    }
    sdrIds.EntityData.Leafs = types.NewOrderedMap()

    sdrIds.EntityData.YListKeys = []string {}

    return &(sdrIds.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId
// Remote data for sdr-id
type CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Sdr Id. The type is interface{} with range:
    // 1..65535.
    SdrId interface{}

    // SDR ID. The type is interface{} with range: 0..4294967295.
    SdrIdXr interface{}

    // Remote data entry state. The type is CsiMaItemState.
    ItemState interface{}

    // Published data.
    PubData CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData
}

func (sdrId *CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId) GetEntityData() *types.CommonEntityData {
    sdrId.EntityData.YFilter = sdrId.YFilter
    sdrId.EntityData.YangName = "sdr-id"
    sdrId.EntityData.BundleName = "cisco_ios_xr"
    sdrId.EntityData.ParentYangName = "sdr-ids"
    sdrId.EntityData.SegmentPath = "sdr-id" + types.AddKeyToken(sdrId.SdrId, "sdr-id")
    sdrId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrId.EntityData.Children = types.NewOrderedMap()
    sdrId.EntityData.Children.Append("pub-data", types.YChild{"PubData", &sdrId.PubData})
    sdrId.EntityData.Leafs = types.NewOrderedMap()
    sdrId.EntityData.Leafs.Append("sdr-id", types.YLeaf{"SdrId", sdrId.SdrId})
    sdrId.EntityData.Leafs.Append("sdr-id-xr", types.YLeaf{"SdrIdXr", sdrId.SdrIdXr})
    sdrId.EntityData.Leafs.Append("item-state", types.YLeaf{"ItemState", sdrId.ItemState})

    sdrId.EntityData.YListKeys = []string {"SdrId"}

    return &(sdrId.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData
// Published data
type CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CSI PIC Array. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    CsiPicArr []interface{}

    // CSI Index - IP Array. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr.
    CsiInfoArr []*CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr
}

func (pubData *CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData) GetEntityData() *types.CommonEntityData {
    pubData.EntityData.YFilter = pubData.YFilter
    pubData.EntityData.YangName = "pub-data"
    pubData.EntityData.BundleName = "cisco_ios_xr"
    pubData.EntityData.ParentYangName = "sdr-id"
    pubData.EntityData.SegmentPath = "pub-data"
    pubData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pubData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pubData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pubData.EntityData.Children = types.NewOrderedMap()
    pubData.EntityData.Children.Append("csi-info-arr", types.YChild{"CsiInfoArr", nil})
    for i := range pubData.CsiInfoArr {
        pubData.EntityData.Children.Append(types.GetSegmentPath(pubData.CsiInfoArr[i]), types.YChild{"CsiInfoArr", pubData.CsiInfoArr[i]})
    }
    pubData.EntityData.Leafs = types.NewOrderedMap()
    pubData.EntityData.Leafs.Append("csi-pic-arr", types.YLeaf{"CsiPicArr", pubData.CsiPicArr})

    pubData.EntityData.YListKeys = []string {}

    return &(pubData.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr
// CSI Index - IP Array
type CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CSI Index. The type is interface{} with range: 0..4294967295.
    CsiIndex interface{}

    // CSI IP Array. The type is slice of
    // CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr_IpArr.
    IpArr []*CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr_IpArr
}

func (csiInfoArr *CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr) GetEntityData() *types.CommonEntityData {
    csiInfoArr.EntityData.YFilter = csiInfoArr.YFilter
    csiInfoArr.EntityData.YangName = "csi-info-arr"
    csiInfoArr.EntityData.BundleName = "cisco_ios_xr"
    csiInfoArr.EntityData.ParentYangName = "pub-data"
    csiInfoArr.EntityData.SegmentPath = "csi-info-arr"
    csiInfoArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    csiInfoArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    csiInfoArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    csiInfoArr.EntityData.Children = types.NewOrderedMap()
    csiInfoArr.EntityData.Children.Append("ip-arr", types.YChild{"IpArr", nil})
    for i := range csiInfoArr.IpArr {
        csiInfoArr.EntityData.Children.Append(types.GetSegmentPath(csiInfoArr.IpArr[i]), types.YChild{"IpArr", csiInfoArr.IpArr[i]})
    }
    csiInfoArr.EntityData.Leafs = types.NewOrderedMap()
    csiInfoArr.EntityData.Leafs.Append("csi-index", types.YLeaf{"CsiIndex", csiInfoArr.CsiIndex})

    csiInfoArr.EntityData.YListKeys = []string {}

    return &(csiInfoArr.EntityData)
}

// CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr_IpArr
// CSI IP Array
type CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr_IpArr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is CsiMaAfi.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipArr *CrossSdrIntfMa_Nodes_Node_SdrIds_SdrId_PubData_CsiInfoArr_IpArr) GetEntityData() *types.CommonEntityData {
    ipArr.EntityData.YFilter = ipArr.YFilter
    ipArr.EntityData.YangName = "ip-arr"
    ipArr.EntityData.BundleName = "cisco_ios_xr"
    ipArr.EntityData.ParentYangName = "csi-info-arr"
    ipArr.EntityData.SegmentPath = "ip-arr"
    ipArr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipArr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipArr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipArr.EntityData.Children = types.NewOrderedMap()
    ipArr.EntityData.Leafs = types.NewOrderedMap()
    ipArr.EntityData.Leafs.Append("af", types.YLeaf{"Af", ipArr.Af})
    ipArr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipArr.Ipv4})
    ipArr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipArr.Ipv6})

    ipArr.EntityData.YListKeys = []string {}

    return &(ipArr.EntityData)
}

